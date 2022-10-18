// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package helpcontract

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
)

// IProtocolDataProviderTokenData is an auto generated low-level Go binding around an user-defined struct.
type IProtocolDataProviderTokenData struct {
	Symbol       string
	TokenAddress common.Address
}

// HelpcontractMetaData contains all meta data concerning the Helpcontract contract.
var HelpcontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getAllReservesTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structIProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"stableRateLastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HelpcontractABI is the input ABI used to generate the binding from.
// Deprecated: Use HelpcontractMetaData.ABI instead.
var HelpcontractABI = HelpcontractMetaData.ABI

// Helpcontract is an auto generated Go binding around an Ethereum contract.
type Helpcontract struct {
	HelpcontractCaller     // Read-only binding to the contract
	HelpcontractTransactor // Write-only binding to the contract
	HelpcontractFilterer   // Log filterer for contract events
}

// HelpcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelpcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelpcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelpcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelpcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelpcontractSession struct {
	Contract     *Helpcontract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelpcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelpcontractCallerSession struct {
	Contract *HelpcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HelpcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelpcontractTransactorSession struct {
	Contract     *HelpcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HelpcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelpcontractRaw struct {
	Contract *Helpcontract // Generic contract binding to access the raw methods on
}

// HelpcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelpcontractCallerRaw struct {
	Contract *HelpcontractCaller // Generic read-only contract binding to access the raw methods on
}

// HelpcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelpcontractTransactorRaw struct {
	Contract *HelpcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelpcontract creates a new instance of Helpcontract, bound to a specific deployed contract.
func NewHelpcontract(address common.Address, backend bind.ContractBackend) (*Helpcontract, error) {
	contract, err := bindHelpcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helpcontract{HelpcontractCaller: HelpcontractCaller{contract: contract}, HelpcontractTransactor: HelpcontractTransactor{contract: contract}, HelpcontractFilterer: HelpcontractFilterer{contract: contract}}, nil
}

// NewHelpcontractCaller creates a new read-only instance of Helpcontract, bound to a specific deployed contract.
func NewHelpcontractCaller(address common.Address, caller bind.ContractCaller) (*HelpcontractCaller, error) {
	contract, err := bindHelpcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelpcontractCaller{contract: contract}, nil
}

// NewHelpcontractTransactor creates a new write-only instance of Helpcontract, bound to a specific deployed contract.
func NewHelpcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*HelpcontractTransactor, error) {
	contract, err := bindHelpcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelpcontractTransactor{contract: contract}, nil
}

// NewHelpcontractFilterer creates a new log filterer instance of Helpcontract, bound to a specific deployed contract.
func NewHelpcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*HelpcontractFilterer, error) {
	contract, err := bindHelpcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelpcontractFilterer{contract: contract}, nil
}

// bindHelpcontract binds a generic wrapper to an already deployed contract.
func bindHelpcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelpcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helpcontract *HelpcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helpcontract.Contract.HelpcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helpcontract *HelpcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helpcontract.Contract.HelpcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helpcontract *HelpcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helpcontract.Contract.HelpcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helpcontract *HelpcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helpcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helpcontract *HelpcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helpcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helpcontract *HelpcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helpcontract.Contract.contract.Transact(opts, method, params...)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_Helpcontract *HelpcontractCaller) GetAllReservesTokens(opts *bind.CallOpts) ([]IProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _Helpcontract.contract.Call(opts, &out, "getAllReservesTokens")

	if err != nil {
		return *new([]IProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]IProtocolDataProviderTokenData)).(*[]IProtocolDataProviderTokenData)

	return out0, err

}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_Helpcontract *HelpcontractSession) GetAllReservesTokens() ([]IProtocolDataProviderTokenData, error) {
	return _Helpcontract.Contract.GetAllReservesTokens(&_Helpcontract.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_Helpcontract *HelpcontractCallerSession) GetAllReservesTokens() ([]IProtocolDataProviderTokenData, error) {
	return _Helpcontract.Contract.GetAllReservesTokens(&_Helpcontract.CallOpts)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_Helpcontract *HelpcontractCaller) GetUserReserveData(opts *bind.CallOpts, asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _Helpcontract.contract.Call(opts, &out, "getUserReserveData", asset, user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentStableDebt        *big.Int
		CurrentVariableDebt      *big.Int
		PrincipalStableDebt      *big.Int
		ScaledVariableDebt       *big.Int
		StableBorrowRate         *big.Int
		LiquidityRate            *big.Int
		StableRateLastUpdated    *big.Int
		UsageAsCollateralEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentATokenBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentStableDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentVariableDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PrincipalStableDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ScaledVariableDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StableRateLastUpdated = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_Helpcontract *HelpcontractSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _Helpcontract.Contract.GetUserReserveData(&_Helpcontract.CallOpts, asset, user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_Helpcontract *HelpcontractCallerSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _Helpcontract.Contract.GetUserReserveData(&_Helpcontract.CallOpts, asset, user)
}
