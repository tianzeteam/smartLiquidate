// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oraclecontract

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

// OraclecontractMetaData contains all meta data concerning the Oraclecontract contract.
var OraclecontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getAssetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OraclecontractABI is the input ABI used to generate the binding from.
// Deprecated: Use OraclecontractMetaData.ABI instead.
var OraclecontractABI = OraclecontractMetaData.ABI

// Oraclecontract is an auto generated Go binding around an Ethereum contract.
type Oraclecontract struct {
	OraclecontractCaller     // Read-only binding to the contract
	OraclecontractTransactor // Write-only binding to the contract
	OraclecontractFilterer   // Log filterer for contract events
}

// OraclecontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type OraclecontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclecontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OraclecontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclecontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OraclecontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OraclecontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OraclecontractSession struct {
	Contract     *Oraclecontract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OraclecontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OraclecontractCallerSession struct {
	Contract *OraclecontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OraclecontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OraclecontractTransactorSession struct {
	Contract     *OraclecontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OraclecontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type OraclecontractRaw struct {
	Contract *Oraclecontract // Generic contract binding to access the raw methods on
}

// OraclecontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OraclecontractCallerRaw struct {
	Contract *OraclecontractCaller // Generic read-only contract binding to access the raw methods on
}

// OraclecontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OraclecontractTransactorRaw struct {
	Contract *OraclecontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOraclecontract creates a new instance of Oraclecontract, bound to a specific deployed contract.
func NewOraclecontract(address common.Address, backend bind.ContractBackend) (*Oraclecontract, error) {
	contract, err := bindOraclecontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Oraclecontract{OraclecontractCaller: OraclecontractCaller{contract: contract}, OraclecontractTransactor: OraclecontractTransactor{contract: contract}, OraclecontractFilterer: OraclecontractFilterer{contract: contract}}, nil
}

// NewOraclecontractCaller creates a new read-only instance of Oraclecontract, bound to a specific deployed contract.
func NewOraclecontractCaller(address common.Address, caller bind.ContractCaller) (*OraclecontractCaller, error) {
	contract, err := bindOraclecontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OraclecontractCaller{contract: contract}, nil
}

// NewOraclecontractTransactor creates a new write-only instance of Oraclecontract, bound to a specific deployed contract.
func NewOraclecontractTransactor(address common.Address, transactor bind.ContractTransactor) (*OraclecontractTransactor, error) {
	contract, err := bindOraclecontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OraclecontractTransactor{contract: contract}, nil
}

// NewOraclecontractFilterer creates a new log filterer instance of Oraclecontract, bound to a specific deployed contract.
func NewOraclecontractFilterer(address common.Address, filterer bind.ContractFilterer) (*OraclecontractFilterer, error) {
	contract, err := bindOraclecontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OraclecontractFilterer{contract: contract}, nil
}

// bindOraclecontract binds a generic wrapper to an already deployed contract.
func bindOraclecontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OraclecontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oraclecontract *OraclecontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oraclecontract.Contract.OraclecontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oraclecontract *OraclecontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oraclecontract.Contract.OraclecontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oraclecontract *OraclecontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oraclecontract.Contract.OraclecontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Oraclecontract *OraclecontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Oraclecontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Oraclecontract *OraclecontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Oraclecontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Oraclecontract *OraclecontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Oraclecontract.Contract.contract.Transact(opts, method, params...)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Oraclecontract *OraclecontractCaller) GetAssetPrice(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Oraclecontract.contract.Call(opts, &out, "getAssetPrice", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Oraclecontract *OraclecontractSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _Oraclecontract.Contract.GetAssetPrice(&_Oraclecontract.CallOpts, asset)
}

// GetAssetPrice is a free data retrieval call binding the contract method 0xb3596f07.
//
// Solidity: function getAssetPrice(address asset) view returns(uint256)
func (_Oraclecontract *OraclecontractCallerSession) GetAssetPrice(asset common.Address) (*big.Int, error) {
	return _Oraclecontract.Contract.GetAssetPrice(&_Oraclecontract.CallOpts, asset)
}
