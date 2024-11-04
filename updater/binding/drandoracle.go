// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package binding

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IDrandOracleRandom is an auto generated low-level Go binding around an user-defined struct.
type IDrandOracleRandom struct {
	Round      uint64
	Timestamp  uint64
	Randomness [32]byte
	Signature  []byte
}

// BindingMetaData contains all meta data concerning the Binding contract.
var BindingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CHAIN_HASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"earliestRound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRandomnessFromRound\",\"inputs\":[{\"name\":\"_round\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDrandOracle.Random\",\"components\":[{\"name\":\"round\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRandomnessFromTimestamp\",\"inputs\":[{\"name\":\"_timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDrandOracle.Random\",\"components\":[{\"name\":\"round\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestRound\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rounds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"round\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRandomness\",\"inputs\":[{\"name\":\"_random\",\"type\":\"tuple\",\"internalType\":\"structIDrandOracle.Random\",\"components\":[{\"name\":\"round\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSigner\",\"inputs\":[{\"name\":\"_newSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"timestamps\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"round\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RandomnessUpdated\",\"inputs\":[{\"name\":\"round\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerUpdated\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInput\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoundTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// BindingABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingMetaData.ABI instead.
var BindingABI = BindingMetaData.ABI

// Binding is an auto generated Go binding around an Ethereum contract.
type Binding struct {
	BindingCaller     // Read-only binding to the contract
	BindingTransactor // Write-only binding to the contract
	BindingFilterer   // Log filterer for contract events
}

// BindingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingSession struct {
	Contract     *Binding          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingCallerSession struct {
	Contract *BindingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BindingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingTransactorSession struct {
	Contract     *BindingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BindingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingRaw struct {
	Contract *Binding // Generic contract binding to access the raw methods on
}

// BindingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingCallerRaw struct {
	Contract *BindingCaller // Generic read-only contract binding to access the raw methods on
}

// BindingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingTransactorRaw struct {
	Contract *BindingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBinding creates a new instance of Binding, bound to a specific deployed contract.
func NewBinding(address common.Address, backend bind.ContractBackend) (*Binding, error) {
	contract, err := bindBinding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Binding{BindingCaller: BindingCaller{contract: contract}, BindingTransactor: BindingTransactor{contract: contract}, BindingFilterer: BindingFilterer{contract: contract}}, nil
}

// NewBindingCaller creates a new read-only instance of Binding, bound to a specific deployed contract.
func NewBindingCaller(address common.Address, caller bind.ContractCaller) (*BindingCaller, error) {
	contract, err := bindBinding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingCaller{contract: contract}, nil
}

// NewBindingTransactor creates a new write-only instance of Binding, bound to a specific deployed contract.
func NewBindingTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingTransactor, error) {
	contract, err := bindBinding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingTransactor{contract: contract}, nil
}

// NewBindingFilterer creates a new log filterer instance of Binding, bound to a specific deployed contract.
func NewBindingFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingFilterer, error) {
	contract, err := bindBinding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingFilterer{contract: contract}, nil
}

// bindBinding binds a generic wrapper to an already deployed contract.
func bindBinding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Binding *BindingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Binding.Contract.BindingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Binding *BindingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.Contract.BindingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Binding *BindingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Binding.Contract.BindingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Binding *BindingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Binding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Binding *BindingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Binding *BindingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Binding.Contract.contract.Transact(opts, method, params...)
}

// CHAINHASH is a free data retrieval call binding the contract method 0x9acbc4ca.
//
// Solidity: function CHAIN_HASH() view returns(bytes32)
func (_Binding *BindingCaller) CHAINHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "CHAIN_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CHAINHASH is a free data retrieval call binding the contract method 0x9acbc4ca.
//
// Solidity: function CHAIN_HASH() view returns(bytes32)
func (_Binding *BindingSession) CHAINHASH() ([32]byte, error) {
	return _Binding.Contract.CHAINHASH(&_Binding.CallOpts)
}

// CHAINHASH is a free data retrieval call binding the contract method 0x9acbc4ca.
//
// Solidity: function CHAIN_HASH() view returns(bytes32)
func (_Binding *BindingCallerSession) CHAINHASH() ([32]byte, error) {
	return _Binding.Contract.CHAINHASH(&_Binding.CallOpts)
}

// EarliestRound is a free data retrieval call binding the contract method 0x67eb66cb.
//
// Solidity: function earliestRound() view returns(uint64)
func (_Binding *BindingCaller) EarliestRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "earliestRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// EarliestRound is a free data retrieval call binding the contract method 0x67eb66cb.
//
// Solidity: function earliestRound() view returns(uint64)
func (_Binding *BindingSession) EarliestRound() (uint64, error) {
	return _Binding.Contract.EarliestRound(&_Binding.CallOpts)
}

// EarliestRound is a free data retrieval call binding the contract method 0x67eb66cb.
//
// Solidity: function earliestRound() view returns(uint64)
func (_Binding *BindingCallerSession) EarliestRound() (uint64, error) {
	return _Binding.Contract.EarliestRound(&_Binding.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Binding *BindingCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Binding *BindingSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Binding.Contract.Eip712Domain(&_Binding.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Binding *BindingCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Binding.Contract.Eip712Domain(&_Binding.CallOpts)
}

// GetRandomnessFromRound is a free data retrieval call binding the contract method 0xfc4f57c5.
//
// Solidity: function getRandomnessFromRound(uint64 _round) view returns((uint64,uint64,bytes32,bytes))
func (_Binding *BindingCaller) GetRandomnessFromRound(opts *bind.CallOpts, _round uint64) (IDrandOracleRandom, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "getRandomnessFromRound", _round)

	if err != nil {
		return *new(IDrandOracleRandom), err
	}

	out0 := *abi.ConvertType(out[0], new(IDrandOracleRandom)).(*IDrandOracleRandom)

	return out0, err

}

// GetRandomnessFromRound is a free data retrieval call binding the contract method 0xfc4f57c5.
//
// Solidity: function getRandomnessFromRound(uint64 _round) view returns((uint64,uint64,bytes32,bytes))
func (_Binding *BindingSession) GetRandomnessFromRound(_round uint64) (IDrandOracleRandom, error) {
	return _Binding.Contract.GetRandomnessFromRound(&_Binding.CallOpts, _round)
}

// GetRandomnessFromRound is a free data retrieval call binding the contract method 0xfc4f57c5.
//
// Solidity: function getRandomnessFromRound(uint64 _round) view returns((uint64,uint64,bytes32,bytes))
func (_Binding *BindingCallerSession) GetRandomnessFromRound(_round uint64) (IDrandOracleRandom, error) {
	return _Binding.Contract.GetRandomnessFromRound(&_Binding.CallOpts, _round)
}

// GetRandomnessFromTimestamp is a free data retrieval call binding the contract method 0x167ef5fb.
//
// Solidity: function getRandomnessFromTimestamp(uint64 _timestamp) view returns((uint64,uint64,bytes32,bytes))
func (_Binding *BindingCaller) GetRandomnessFromTimestamp(opts *bind.CallOpts, _timestamp uint64) (IDrandOracleRandom, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "getRandomnessFromTimestamp", _timestamp)

	if err != nil {
		return *new(IDrandOracleRandom), err
	}

	out0 := *abi.ConvertType(out[0], new(IDrandOracleRandom)).(*IDrandOracleRandom)

	return out0, err

}

// GetRandomnessFromTimestamp is a free data retrieval call binding the contract method 0x167ef5fb.
//
// Solidity: function getRandomnessFromTimestamp(uint64 _timestamp) view returns((uint64,uint64,bytes32,bytes))
func (_Binding *BindingSession) GetRandomnessFromTimestamp(_timestamp uint64) (IDrandOracleRandom, error) {
	return _Binding.Contract.GetRandomnessFromTimestamp(&_Binding.CallOpts, _timestamp)
}

// GetRandomnessFromTimestamp is a free data retrieval call binding the contract method 0x167ef5fb.
//
// Solidity: function getRandomnessFromTimestamp(uint64 _timestamp) view returns((uint64,uint64,bytes32,bytes))
func (_Binding *BindingCallerSession) GetRandomnessFromTimestamp(_timestamp uint64) (IDrandOracleRandom, error) {
	return _Binding.Contract.GetRandomnessFromTimestamp(&_Binding.CallOpts, _timestamp)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint64)
func (_Binding *BindingCaller) LatestRound(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint64)
func (_Binding *BindingSession) LatestRound() (uint64, error) {
	return _Binding.Contract.LatestRound(&_Binding.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint64)
func (_Binding *BindingCallerSession) LatestRound() (uint64, error) {
	return _Binding.Contract.LatestRound(&_Binding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Binding *BindingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Binding *BindingSession) Owner() (common.Address, error) {
	return _Binding.Contract.Owner(&_Binding.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Binding *BindingCallerSession) Owner() (common.Address, error) {
	return _Binding.Contract.Owner(&_Binding.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Binding *BindingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Binding *BindingSession) Paused() (bool, error) {
	return _Binding.Contract.Paused(&_Binding.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Binding *BindingCallerSession) Paused() (bool, error) {
	return _Binding.Contract.Paused(&_Binding.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Binding *BindingCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Binding *BindingSession) PendingOwner() (common.Address, error) {
	return _Binding.Contract.PendingOwner(&_Binding.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Binding *BindingCallerSession) PendingOwner() (common.Address, error) {
	return _Binding.Contract.PendingOwner(&_Binding.CallOpts)
}

// Rounds is a free data retrieval call binding the contract method 0xbc3ac875.
//
// Solidity: function rounds(uint64 ) view returns(uint64 round, uint64 timestamp, bytes32 randomness, bytes signature)
func (_Binding *BindingCaller) Rounds(opts *bind.CallOpts, arg0 uint64) (struct {
	Round      uint64
	Timestamp  uint64
	Randomness [32]byte
	Signature  []byte
}, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "rounds", arg0)

	outstruct := new(struct {
		Round      uint64
		Timestamp  uint64
		Randomness [32]byte
		Signature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Round = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Randomness = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Signature = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Rounds is a free data retrieval call binding the contract method 0xbc3ac875.
//
// Solidity: function rounds(uint64 ) view returns(uint64 round, uint64 timestamp, bytes32 randomness, bytes signature)
func (_Binding *BindingSession) Rounds(arg0 uint64) (struct {
	Round      uint64
	Timestamp  uint64
	Randomness [32]byte
	Signature  []byte
}, error) {
	return _Binding.Contract.Rounds(&_Binding.CallOpts, arg0)
}

// Rounds is a free data retrieval call binding the contract method 0xbc3ac875.
//
// Solidity: function rounds(uint64 ) view returns(uint64 round, uint64 timestamp, bytes32 randomness, bytes signature)
func (_Binding *BindingCallerSession) Rounds(arg0 uint64) (struct {
	Round      uint64
	Timestamp  uint64
	Randomness [32]byte
	Signature  []byte
}, error) {
	return _Binding.Contract.Rounds(&_Binding.CallOpts, arg0)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Binding *BindingCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Binding *BindingSession) Signer() (common.Address, error) {
	return _Binding.Contract.Signer(&_Binding.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Binding *BindingCallerSession) Signer() (common.Address, error) {
	return _Binding.Contract.Signer(&_Binding.CallOpts)
}

// Timestamps is a free data retrieval call binding the contract method 0xdeaf3ba6.
//
// Solidity: function timestamps(uint64 ) view returns(uint64 round, uint64 timestamp, bytes32 randomness, bytes signature)
func (_Binding *BindingCaller) Timestamps(opts *bind.CallOpts, arg0 uint64) (struct {
	Round      uint64
	Timestamp  uint64
	Randomness [32]byte
	Signature  []byte
}, error) {
	var out []interface{}
	err := _Binding.contract.Call(opts, &out, "timestamps", arg0)

	outstruct := new(struct {
		Round      uint64
		Timestamp  uint64
		Randomness [32]byte
		Signature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Round = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Randomness = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.Signature = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// Timestamps is a free data retrieval call binding the contract method 0xdeaf3ba6.
//
// Solidity: function timestamps(uint64 ) view returns(uint64 round, uint64 timestamp, bytes32 randomness, bytes signature)
func (_Binding *BindingSession) Timestamps(arg0 uint64) (struct {
	Round      uint64
	Timestamp  uint64
	Randomness [32]byte
	Signature  []byte
}, error) {
	return _Binding.Contract.Timestamps(&_Binding.CallOpts, arg0)
}

// Timestamps is a free data retrieval call binding the contract method 0xdeaf3ba6.
//
// Solidity: function timestamps(uint64 ) view returns(uint64 round, uint64 timestamp, bytes32 randomness, bytes signature)
func (_Binding *BindingCallerSession) Timestamps(arg0 uint64) (struct {
	Round      uint64
	Timestamp  uint64
	Randomness [32]byte
	Signature  []byte
}, error) {
	return _Binding.Contract.Timestamps(&_Binding.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Binding *BindingTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Binding *BindingSession) AcceptOwnership() (*types.Transaction, error) {
	return _Binding.Contract.AcceptOwnership(&_Binding.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Binding *BindingTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Binding.Contract.AcceptOwnership(&_Binding.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Binding *BindingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Binding *BindingSession) Pause() (*types.Transaction, error) {
	return _Binding.Contract.Pause(&_Binding.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Binding *BindingTransactorSession) Pause() (*types.Transaction, error) {
	return _Binding.Contract.Pause(&_Binding.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Binding *BindingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Binding *BindingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Binding.Contract.RenounceOwnership(&_Binding.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Binding *BindingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Binding.Contract.RenounceOwnership(&_Binding.TransactOpts)
}

// SetRandomness is a paid mutator transaction binding the contract method 0x2e1f8309.
//
// Solidity: function setRandomness((uint64,uint64,bytes32,bytes) _random, bytes _signature) returns()
func (_Binding *BindingTransactor) SetRandomness(opts *bind.TransactOpts, _random IDrandOracleRandom, _signature []byte) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "setRandomness", _random, _signature)
}

// SetRandomness is a paid mutator transaction binding the contract method 0x2e1f8309.
//
// Solidity: function setRandomness((uint64,uint64,bytes32,bytes) _random, bytes _signature) returns()
func (_Binding *BindingSession) SetRandomness(_random IDrandOracleRandom, _signature []byte) (*types.Transaction, error) {
	return _Binding.Contract.SetRandomness(&_Binding.TransactOpts, _random, _signature)
}

// SetRandomness is a paid mutator transaction binding the contract method 0x2e1f8309.
//
// Solidity: function setRandomness((uint64,uint64,bytes32,bytes) _random, bytes _signature) returns()
func (_Binding *BindingTransactorSession) SetRandomness(_random IDrandOracleRandom, _signature []byte) (*types.Transaction, error) {
	return _Binding.Contract.SetRandomness(&_Binding.TransactOpts, _random, _signature)
}

// SetSigner is a paid mutator transaction binding the contract method 0x03c0737f.
//
// Solidity: function setSigner(address _newSigner, bytes _signature) returns()
func (_Binding *BindingTransactor) SetSigner(opts *bind.TransactOpts, _newSigner common.Address, _signature []byte) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "setSigner", _newSigner, _signature)
}

// SetSigner is a paid mutator transaction binding the contract method 0x03c0737f.
//
// Solidity: function setSigner(address _newSigner, bytes _signature) returns()
func (_Binding *BindingSession) SetSigner(_newSigner common.Address, _signature []byte) (*types.Transaction, error) {
	return _Binding.Contract.SetSigner(&_Binding.TransactOpts, _newSigner, _signature)
}

// SetSigner is a paid mutator transaction binding the contract method 0x03c0737f.
//
// Solidity: function setSigner(address _newSigner, bytes _signature) returns()
func (_Binding *BindingTransactorSession) SetSigner(_newSigner common.Address, _signature []byte) (*types.Transaction, error) {
	return _Binding.Contract.SetSigner(&_Binding.TransactOpts, _newSigner, _signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Binding *BindingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Binding *BindingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Binding.Contract.TransferOwnership(&_Binding.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Binding *BindingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Binding.Contract.TransferOwnership(&_Binding.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Binding *BindingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Binding.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Binding *BindingSession) Unpause() (*types.Transaction, error) {
	return _Binding.Contract.Unpause(&_Binding.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Binding *BindingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Binding.Contract.Unpause(&_Binding.TransactOpts)
}

// BindingEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Binding contract.
type BindingEIP712DomainChangedIterator struct {
	Event *BindingEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingEIP712DomainChanged represents a EIP712DomainChanged event raised by the Binding contract.
type BindingEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Binding *BindingFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*BindingEIP712DomainChangedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &BindingEIP712DomainChangedIterator{contract: _Binding.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Binding *BindingFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *BindingEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingEIP712DomainChanged)
				if err := _Binding.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Binding *BindingFilterer) ParseEIP712DomainChanged(log types.Log) (*BindingEIP712DomainChanged, error) {
	event := new(BindingEIP712DomainChanged)
	if err := _Binding.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Binding contract.
type BindingOwnershipTransferStartedIterator struct {
	Event *BindingOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Binding contract.
type BindingOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Binding.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingOwnershipTransferStartedIterator{contract: _Binding.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *BindingOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Binding.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingOwnershipTransferStarted)
				if err := _Binding.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) ParseOwnershipTransferStarted(log types.Log) (*BindingOwnershipTransferStarted, error) {
	event := new(BindingOwnershipTransferStarted)
	if err := _Binding.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Binding contract.
type BindingOwnershipTransferredIterator struct {
	Event *BindingOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingOwnershipTransferred represents a OwnershipTransferred event raised by the Binding contract.
type BindingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BindingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Binding.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BindingOwnershipTransferredIterator{contract: _Binding.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BindingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Binding.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingOwnershipTransferred)
				if err := _Binding.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Binding *BindingFilterer) ParseOwnershipTransferred(log types.Log) (*BindingOwnershipTransferred, error) {
	event := new(BindingOwnershipTransferred)
	if err := _Binding.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Binding contract.
type BindingPausedIterator struct {
	Event *BindingPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingPaused represents a Paused event raised by the Binding contract.
type BindingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Binding *BindingFilterer) FilterPaused(opts *bind.FilterOpts) (*BindingPausedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BindingPausedIterator{contract: _Binding.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Binding *BindingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BindingPaused) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingPaused)
				if err := _Binding.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Binding *BindingFilterer) ParsePaused(log types.Log) (*BindingPaused, error) {
	event := new(BindingPaused)
	if err := _Binding.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingRandomnessUpdatedIterator is returned from FilterRandomnessUpdated and is used to iterate over the raw logs and unpacked data for RandomnessUpdated events raised by the Binding contract.
type BindingRandomnessUpdatedIterator struct {
	Event *BindingRandomnessUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingRandomnessUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingRandomnessUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingRandomnessUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingRandomnessUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingRandomnessUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingRandomnessUpdated represents a RandomnessUpdated event raised by the Binding contract.
type BindingRandomnessUpdated struct {
	Round      uint64
	Randomness [32]byte
	Signature  []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRandomnessUpdated is a free log retrieval operation binding the contract event 0xb2f2d8fc12d36205a87db37f52aa7754f67d078758e4c76ef43c0a420b399724.
//
// Solidity: event RandomnessUpdated(uint64 round, bytes32 randomness, bytes signature)
func (_Binding *BindingFilterer) FilterRandomnessUpdated(opts *bind.FilterOpts) (*BindingRandomnessUpdatedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "RandomnessUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingRandomnessUpdatedIterator{contract: _Binding.contract, event: "RandomnessUpdated", logs: logs, sub: sub}, nil
}

// WatchRandomnessUpdated is a free log subscription operation binding the contract event 0xb2f2d8fc12d36205a87db37f52aa7754f67d078758e4c76ef43c0a420b399724.
//
// Solidity: event RandomnessUpdated(uint64 round, bytes32 randomness, bytes signature)
func (_Binding *BindingFilterer) WatchRandomnessUpdated(opts *bind.WatchOpts, sink chan<- *BindingRandomnessUpdated) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "RandomnessUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingRandomnessUpdated)
				if err := _Binding.contract.UnpackLog(event, "RandomnessUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomnessUpdated is a log parse operation binding the contract event 0xb2f2d8fc12d36205a87db37f52aa7754f67d078758e4c76ef43c0a420b399724.
//
// Solidity: event RandomnessUpdated(uint64 round, bytes32 randomness, bytes signature)
func (_Binding *BindingFilterer) ParseRandomnessUpdated(log types.Log) (*BindingRandomnessUpdated, error) {
	event := new(BindingRandomnessUpdated)
	if err := _Binding.contract.UnpackLog(event, "RandomnessUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the Binding contract.
type BindingSignerUpdatedIterator struct {
	Event *BindingSignerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingSignerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingSignerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingSignerUpdated represents a SignerUpdated event raised by the Binding contract.
type BindingSignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_Binding *BindingFilterer) FilterSignerUpdated(opts *bind.FilterOpts) (*BindingSignerUpdatedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return &BindingSignerUpdatedIterator{contract: _Binding.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_Binding *BindingFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *BindingSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingSignerUpdated)
				if err := _Binding.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerUpdated is a log parse operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_Binding *BindingFilterer) ParseSignerUpdated(log types.Log) (*BindingSignerUpdated, error) {
	event := new(BindingSignerUpdated)
	if err := _Binding.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BindingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Binding contract.
type BindingUnpausedIterator struct {
	Event *BindingUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingUnpaused represents a Unpaused event raised by the Binding contract.
type BindingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Binding *BindingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BindingUnpausedIterator, error) {

	logs, sub, err := _Binding.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BindingUnpausedIterator{contract: _Binding.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Binding *BindingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BindingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Binding.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingUnpaused)
				if err := _Binding.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Binding *BindingFilterer) ParseUnpaused(log types.Log) (*BindingUnpaused, error) {
	event := new(BindingUnpaused)
	if err := _Binding.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
