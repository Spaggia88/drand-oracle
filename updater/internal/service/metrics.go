package service

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/drand/drand/chain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	// Label names
	labelChainHash      = "chain_hash"
	labelChainID        = "chain_id"
	labelOracleAddress  = "oracle_address"
	labelUpdaterAddress = "updater_address"

	// Drand info metric labels
	labelPublicKey   = "public_key"
	labelID          = "id"
	labelPeriod      = "period"
	labelScheme      = "scheme"
	labelGenesisTime = "genesis_time"
	labelGenesisSeed = "genesis_seed"
)

// Metrics holds all Prometheus metrics for the updater
type Metrics struct {
	drandRoundTotal           *prometheus.GaugeVec
	oracleRoundTotal          *prometheus.GaugeVec
	setRandomnessSuccessTotal *prometheus.CounterVec
	setRandomnessFailureTotal *prometheus.CounterVec
	updaterBalance            *prometheus.GaugeVec

	// New info metric
	drandInfo *prometheus.GaugeVec

	// Store label values
	chainHash      string
	chainID        int64
	oracleAddress  common.Address
	updaterAddress common.Address
}

// NewMetrics creates and registers all Prometheus metrics
func NewMetrics(chainID int64, oracleAddress common.Address, updaterAddress common.Address, drandInfo *chain.Info) *Metrics {
	m := &Metrics{
		chainHash:      drandInfo.HashString(),
		chainID:        chainID,
		oracleAddress:  oracleAddress,
		updaterAddress: updaterAddress,
	}

	m.drandRoundTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "drand_round_number_network",
		Help: "Current round number from the Drand network",
	}, []string{labelChainHash})

	m.oracleRoundTotal = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "drand_round_number_oracle",
		Help: "Current round number processed by the Oracle",
	}, []string{labelChainID, labelOracleAddress})

	m.setRandomnessSuccessTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "drand_set_randomness_success_total",
		Help: "Total number of successful SetRandomness transactions",
	}, []string{labelChainID, labelOracleAddress})

	m.setRandomnessFailureTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "drand_set_randomness_failure_total",
		Help: "Total number of failed SetRandomness transactions",
	}, []string{labelChainID, labelOracleAddress})

	// Add info metric
	m.drandInfo = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "drand_network_info",
		Help: "Static information about the Drand network configuration",
	}, []string{
		labelChainHash,
		labelPublicKey,
		labelPeriod,
		labelScheme,
		labelGenesisTime,
		labelGenesisSeed,
	})

	// Set the info metric with a constant value of 1
	m.drandInfo.WithLabelValues(
		m.chainHash,
		drandInfo.PublicKey.String(),
		drandInfo.Period.String(),
		drandInfo.Scheme,
		fmt.Sprintf("%d", drandInfo.GenesisTime),
		hex.EncodeToString(drandInfo.GenesisSeed),
	).Set(1)

	// Add the balance metric
	m.updaterBalance = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "drand_updater_balance_wei",
		Help: "Current balance of the updater address in wei",
	}, []string{labelChainID, labelOracleAddress, labelUpdaterAddress})

	return m
}

// Helper methods for setting metrics with labels
func (m *Metrics) SetDrandRound(round float64) {
	m.drandRoundTotal.WithLabelValues(
		m.chainHash,
	).Set(round)
}

func (m *Metrics) SetOracleRound(round float64) {
	m.oracleRoundTotal.WithLabelValues(
		fmt.Sprintf("%d", m.chainID),
		m.oracleAddress.Hex(),
	).Set(round)
}

func (m *Metrics) IncSetRandomnessSuccess() {
	m.setRandomnessSuccessTotal.WithLabelValues(
		fmt.Sprintf("%d", m.chainID),
		m.oracleAddress.Hex(),
	).Inc()
}

func (m *Metrics) IncSetRandomnessFailure() {
	m.setRandomnessFailureTotal.WithLabelValues(
		fmt.Sprintf("%d", m.chainID),
		m.oracleAddress.Hex(),
	).Inc()
}

func (m *Metrics) SetUpdaterBalance(wei string) {
	balance, _ := new(big.Float).SetString(wei)
	b, _ := balance.Float64()

	m.updaterBalance.WithLabelValues(
		fmt.Sprintf("%d", m.chainID),
		m.oracleAddress.Hex(),
		m.updaterAddress.Hex(),
	).Set(b)
}
