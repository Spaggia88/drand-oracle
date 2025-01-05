package service

import (
	"bytes"
	"context"
	"drand-oracle-updater/binding"
	"drand-oracle-updater/sender"
	"drand-oracle-updater/signer"
	"encoding/hex"
	"errors"
	"math"
	"sync"
	"time"

	"github.com/drand/drand/chain"
	"github.com/drand/drand/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type Updater struct {
	// drandClient is the Drand HTTP client
	drandClient client.Client

	// drandInfo is the Drand info
	drandInfo *chain.Info

	// rpcClient is the Ethereum RPC client
	rpcClient *ethclient.Client

	// setRandomnessGasLimit is the gas limit for the setRandomness transaction
	setRandomnessGasLimit uint64

	// binding is the Drand Oracle contract binding
	binding *binding.Binding

	// chainID is the chain ID
	chainID int64

	// oracleAddress is the address of the Drand Oracle contract
	oracleAddress common.Address

	// genesisRound is the round at which oracle starts tracking
	genesisRound uint64

	// roundChan is the channel for processing rounds
	roundChan chan *roundData

	// latestOracleRound keeps track of the latest round processed by the Oracle
	latestOracleRound      uint64
	latestOracleRoundMutex sync.RWMutex

	// latestDrandRound keeps track of the latest round from the Drand network
	latestDrandRound      uint64
	latestDrandRoundMutex sync.RWMutex

	// signer is the signer for the Drand Oracle contract
	signer *signer.Signer

	// sender is the sender for the Drand Oracle contract
	sender *sender.Sender

	// Metrics instance
	metrics *Metrics
}

const balanceUpdateInterval = 1 * time.Minute

type roundData struct {
	round      uint64
	randomness []byte
	signature  []byte
}

func NewUpdater(
	drandClient client.Client,
	rpcClient *ethclient.Client,
	setRandomnessGasLimit uint64,
	chainID int64,
	oracleAddress common.Address,
	binding *binding.Binding,
	genesisRound uint64,
	signer *signer.Signer,
	sender *sender.Sender,
) (*Updater, error) {
	// Set a timeout for the Drand info request
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get chain hash from drand client
	drandInfo, err := drandClient.Info(ctx)
	if err != nil {
		return nil, err
	}

	updater := &Updater{
		drandClient:           drandClient,
		rpcClient:             rpcClient,
		setRandomnessGasLimit: setRandomnessGasLimit,
		chainID:               chainID,
		oracleAddress:         oracleAddress,
		binding:               binding,
		genesisRound:          genesisRound,
		roundChan:             make(chan *roundData, 1),
		latestOracleRound:     0,
		latestDrandRound:      0,
		signer:                signer,
		sender:                sender,
		metrics: NewMetrics(
			chainID,
			oracleAddress,
			sender.Address(),
			drandInfo,
		),
	}
	return updater, nil
}

func (u *Updater) Start(ctx context.Context) error {
	// Get the earliest and latest round from the Drand Oracle contract
	earliestRound, err := u.binding.EarliestRound(nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get earliest round from Drand Oracle contract")
		return err
	}
	latestRound, err := u.binding.LatestRound(nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get latest round from Drand Oracle contract")
		return err
	}
	u.latestOracleRoundMutex.Lock()
	u.latestOracleRound = latestRound
	u.latestOracleRoundMutex.Unlock()
	log.Info().Msgf("Oracle: Earliest round: %d, Latest round: %d", earliestRound, latestRound)

	// Get the latest round from the Drand network
	latestDrandRound, err := u.drandClient.Get(ctx, 0)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get latest round from Drand network")
		return err
	}
	u.latestDrandRoundMutex.Lock()
	u.latestDrandRound = latestDrandRound.Round()
	u.latestDrandRoundMutex.Unlock()
	log.Info().Msgf("Drand: Latest round: %d", u.latestDrandRound)

	// Get and validate the Drand info against the Oracle contract
	u.drandInfo, err = u.drandClient.Info(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get Drand info")
		return err
	}
	chainHash, err := u.binding.CHAINHASH(nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get chain hash from Drand Oracle contract")
		return err
	}
	if !bytes.Equal(chainHash[:], u.drandInfo.Hash()) {
		err = errors.New("chain hash mismatch")
		return err
	}

	// Start the updater goroutines
	errg, ctx := errgroup.WithContext(ctx)
	errg.Go(func() error {
		return u.processRounds(ctx)
	})
	errg.Go(func() error {
		return u.catchUp(ctx)
	})
	errg.Go(func() error {
		return u.watchNewRounds(ctx)
	})
	errg.Go(func() error {
		return u.monitorBalance(ctx)
	})
	return errg.Wait()
}

func (u *Updater) catchUp(ctx context.Context) error {
	for {
		u.latestDrandRoundMutex.Lock()
		latestDrandRound := u.latestDrandRound
		u.latestDrandRoundMutex.Unlock()

		u.latestOracleRoundMutex.Lock()
		latestOracleRound := u.latestOracleRound
		u.latestOracleRoundMutex.Unlock()

		if latestDrandRound == latestOracleRound {
			log.Info().Msg("Caught up, exiting catch up goroutine")
			break
		}

		var currentRound uint64
		if latestOracleRound == 0 {
			currentRound = u.genesisRound
		} else {
			currentRound = latestOracleRound + 1
		}

		for currentRound <= latestDrandRound {
			result, err := u.drandClient.Get(ctx, currentRound)
			if err != nil {
				log.Error().Err(err).Uint64("round", currentRound).Msg("Failed to get round from Drand network")
				return err
			}

			u.roundChan <- &roundData{
				round:      result.Round(),
				randomness: result.Randomness(),
				signature:  result.Signature(),
			}
			currentRound++
		}
	}
	return nil
}

func (u *Updater) watchNewRounds(ctx context.Context) error {
	for result := range u.drandClient.Watch(ctx) {
		u.latestDrandRoundMutex.Lock()
		u.latestDrandRound = result.Round()
		u.metrics.SetDrandRound(float64(result.Round()))
		u.latestDrandRoundMutex.Unlock()
		u.roundChan <- &roundData{
			round:      result.Round(),
			randomness: result.Randomness(),
			signature:  result.Signature(),
		}
	}
	return nil
}

func (u *Updater) processRounds(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case rd := <-u.roundChan:
			// Add retry logic with exponential backoff
			maxRetries := 5
			var err error
			for attempt := 0; attempt < maxRetries; attempt++ {
				err = u.processRound(ctx, rd.round, rd.randomness, rd.signature)
				if err == nil {
					break
				}

				if attempt < maxRetries-1 {
					backoffDuration := time.Duration(math.Pow(2, float64(attempt))) * time.Second
					log.Warn().
						Err(err).
						Uint64("round", rd.round).
						Int("attempt", attempt+1).
						Dur("backoff", backoffDuration).
						Msg("Retrying round processing after backoff")

					select {
					case <-ctx.Done():
						return ctx.Err()
					case <-time.After(backoffDuration):
						continue
					}
				}
			}

			if err != nil {
				log.Error().
					Err(err).
					Uint64("round", rd.round).
					Msg("Failed to process round after all retries")
				return err
			}
		}
	}
}

func (u *Updater) processRound(
	ctx context.Context,
	round uint64,
	randomness []byte,
	signature []byte,
) error {
	u.latestOracleRoundMutex.Lock()
	defer u.latestOracleRoundMutex.Unlock()
	if round != u.genesisRound && u.latestOracleRound+1 != round {
		log.Info().
			Uint64("latestOracleRound", u.latestOracleRound).
			Uint64("round", round).
			Msg("Skipping irrelevant round")
		return nil
	}

	roundTimestamp := uint64(u.drandInfo.GenesisTime) + uint64(round-1)*uint64(u.drandInfo.Period.Seconds())

	log.Info().
		Uint64("round", round).
		Time("timestamp", time.Unix(int64(roundTimestamp), 0)).
		Str("randomness", hex.EncodeToString(randomness)).
		Str("signature", hex.EncodeToString(signature)).
		Msg("Processing round")

	eip712Signature, err := u.signer.SignSetRandomness(round, roundTimestamp, [32]byte(randomness), signature)
	if err != nil {
		log.Error().Err(err).Msg("Failed to sign set randomness")
		return err
	}

	// Get current gas price suggestion from the network
	gasPrice, err := u.rpcClient.SuggestGasPrice(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get suggested gas price")
		return err
	}

	tx, err := u.binding.SetRandomness(
		&bind.TransactOpts{
			From:     u.sender.Address(),
			Signer:   u.sender.SignerFn(),
			GasLimit: u.setRandomnessGasLimit,
			GasPrice: gasPrice,
		},
		binding.IDrandOracleRandom{
			Round:      round,
			Timestamp:  roundTimestamp,
			Randomness: [32]byte(randomness),
			Signature:  signature,
		},
		eip712Signature,
	)
	if err != nil {
		return err
	}
	receipt, err := bind.WaitMined(ctx, u.rpcClient, tx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to wait for transaction to be mined")
		return err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		u.metrics.IncSetRandomnessFailure()
		err = errors.New("set randomness transaction failed")
		return err
	} else {
		log.Info().Uint64("round", round).Str("hash", tx.Hash().Hex()).Msg("Set randomness transaction successful")
		u.latestOracleRound = round
		u.metrics.SetOracleRound(float64(round))
		u.metrics.IncSetRandomnessSuccess()
	}
	return nil
}

// Add a getter method for safe access
func (u *Updater) GetLatestOracleRound() uint64 {
	u.latestOracleRoundMutex.Lock()
	defer u.latestOracleRoundMutex.Unlock()
	return u.latestOracleRound
}

func (u *Updater) monitorBalance(ctx context.Context) error {
	ticker := time.NewTicker(balanceUpdateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			balance, err := u.rpcClient.BalanceAt(ctx, u.sender.Address(), nil)
			if err != nil {
				log.Error().Err(err).Msg("Failed to get updater balance")
				continue
			}

			u.metrics.SetUpdaterBalance(balance.Int64())

			log.Debug().
				Str("address", u.sender.Address().Hex()).
				Str("balance", balance.String()).
				Msg("Updated balance metric")
		}
	}
}
