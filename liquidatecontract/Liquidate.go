// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquidatecontract

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

// LiquidatecontractMetaData contains all meta data concerning the Liquidatecontract contract.
var LiquidatecontractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"_addressProvider\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"_uniswapV2Router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"stringFailure\",\"type\":\"string\"}],\"name\":\"ErrorHandled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetToLiquidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_flashAmt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_userToLiquidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_swapPath\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_gasCost\",\"type\":\"uint256\"}],\"name\":\"executeFlashLoans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"premiums\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidate_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_userToLiquidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_receiveaToken\",\"type\":\"bool\"}],\"name\":\"liquidateLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"swapPath\",\"type\":\"address[]\"}],\"name\":\"swapToBarrowedAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// LiquidatecontractABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidatecontractMetaData.ABI instead.
var LiquidatecontractABI = LiquidatecontractMetaData.ABI

// Liquidatecontract is an auto generated Go binding around an Ethereum contract.
type Liquidatecontract struct {
	LiquidatecontractCaller     // Read-only binding to the contract
	LiquidatecontractTransactor // Write-only binding to the contract
	LiquidatecontractFilterer   // Log filterer for contract events
}

// LiquidatecontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidatecontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendpoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidatecontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendpoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidatecontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendpoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidatecontractSession struct {
	Contract     *Liquidatecontract        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendpoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidatecontractCallerSession struct {
	Contract *LiquidatecontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LendpoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidatecontractTransactorSession struct {
	Contract     *LiquidatecontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LendpoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidatecontractRaw struct {
	Contract *Liquidatecontract // Generic contract binding to access the raw methods on
}

// LendpoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidatecontractCallerRaw struct {
	Contract *LiquidatecontractCaller // Generic read-only contract binding to access the raw methods on
}

// LendpoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidatecontractTransactorRaw struct {
	Contract *LiquidatecontractTransactor // Generic write-only contract binding to access the raw methods on
}
// NewLiquidatecontract creates a new instance of Liquidatecontract, bound to a specific deployed contract.
func NewLiquidatecontract(address common.Address, backend bind.ContractBackend) (*Liquidatecontract, error) {
	contract, err := bindLiquidatecontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Liquidatecontract{LiquidatecontractCaller: LiquidatecontractCaller{contract: contract}, LiquidatecontractTransactor: LiquidatecontractTransactor{contract: contract}, LiquidatecontractFilterer: LiquidatecontractFilterer{contract: contract}}, nil
}

// NewLiquidatecontractCaller creates a new read-only instance of Liquidatecontract, bound to a specific deployed contract.
func NewLiquidatecontractCaller(address common.Address, caller bind.ContractCaller) (*LiquidatecontractCaller, error) {
	contract, err := bindLiquidatecontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidatecontractCaller{contract: contract}, nil
}

// NewLiquidatecontractTransactor creates a new write-only instance of Liquidatecontract, bound to a specific deployed contract.
func NewLiquidatecontractTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidatecontractTransactor, error) {
	contract, err := bindLiquidatecontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidatecontractTransactor{contract: contract}, nil
}

// NewLiquidatecontractFilterer creates a new log filterer instance of Liquidatecontract, bound to a specific deployed contract.
func NewLiquidatecontractFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidatecontractFilterer, error) {
	contract, err := bindLiquidatecontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidatecontractFilterer{contract: contract}, nil
}

// bindLiquidatecontract binds a generic wrapper to an already deployed contract.
func bindLiquidatecontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LiquidatecontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidatecontract *LiquidatecontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidatecontract.Contract.LiquidatecontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidatecontract *LiquidatecontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.LiquidatecontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidatecontract *LiquidatecontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.LiquidatecontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidatecontract *LiquidatecontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidatecontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidatecontract *LiquidatecontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidatecontract *LiquidatecontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Liquidatecontract *LiquidatecontractCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Liquidatecontract.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Liquidatecontract *LiquidatecontractSession) IsOwner() (bool, error) {
	return _Liquidatecontract.Contract.IsOwner(&_Liquidatecontract.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Liquidatecontract *LiquidatecontractCallerSession) IsOwner() (bool, error) {
	return _Liquidatecontract.Contract.IsOwner(&_Liquidatecontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquidatecontract *LiquidatecontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidatecontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquidatecontract *LiquidatecontractSession) Owner() (common.Address, error) {
	return _Liquidatecontract.Contract.Owner(&_Liquidatecontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquidatecontract *LiquidatecontractCallerSession) Owner() (common.Address, error) {
	return _Liquidatecontract.Contract.Owner(&_Liquidatecontract.CallOpts)
}

// ExecuteFlashLoans is a paid mutator transaction binding the contract method 0x72e59479.
//
// Solidity: function executeFlashLoans(address _assetToLiquidate, uint256 _flashAmt, address _collateral, address _userToLiquidate, uint256 _amountOutMin, address[] _swapPath, uint256 _gasCost) returns()
func (_Liquidatecontract *LiquidatecontractTransactor) ExecuteFlashLoans(opts *bind.TransactOpts, _assetToLiquidate common.Address, _flashAmt *big.Int, _collateral common.Address, _userToLiquidate common.Address, _amountOutMin *big.Int, _swapPath []common.Address, _gasCost *big.Int) (*types.Transaction, error) {
	return _Liquidatecontract.contract.Transact(opts, "executeFlashLoans", _assetToLiquidate, _flashAmt, _collateral, _userToLiquidate, _amountOutMin, _swapPath, _gasCost)
}

// ExecuteFlashLoans is a paid mutator transaction binding the contract method 0x72e59479.
//
// Solidity: function executeFlashLoans(address _assetToLiquidate, uint256 _flashAmt, address _collateral, address _userToLiquidate, uint256 _amountOutMin, address[] _swapPath, uint256 _gasCost) returns()
func (_Liquidatecontract *LiquidatecontractSession) ExecuteFlashLoans(_assetToLiquidate common.Address, _flashAmt *big.Int, _collateral common.Address, _userToLiquidate common.Address, _amountOutMin *big.Int, _swapPath []common.Address, _gasCost *big.Int) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.ExecuteFlashLoans(&_Liquidatecontract.TransactOpts, _assetToLiquidate, _flashAmt, _collateral, _userToLiquidate, _amountOutMin, _swapPath, _gasCost)
}

// ExecuteFlashLoans is a paid mutator transaction binding the contract method 0x72e59479.
//
// Solidity: function executeFlashLoans(address _assetToLiquidate, uint256 _flashAmt, address _collateral, address _userToLiquidate, uint256 _amountOutMin, address[] _swapPath, uint256 _gasCost) returns()
func (_Liquidatecontract *LiquidatecontractTransactorSession) ExecuteFlashLoans(_assetToLiquidate common.Address, _flashAmt *big.Int, _collateral common.Address, _userToLiquidate common.Address, _amountOutMin *big.Int, _swapPath []common.Address, _gasCost *big.Int) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.ExecuteFlashLoans(&_Liquidatecontract.TransactOpts, _assetToLiquidate, _flashAmt, _collateral, _userToLiquidate, _amountOutMin, _swapPath, _gasCost)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address , bytes params) returns(bool)
func (_Liquidatecontract *LiquidatecontractTransactor) ExecuteOperation(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, premiums []*big.Int, arg3 common.Address, params []byte) (*types.Transaction, error) {
	return _Liquidatecontract.contract.Transact(opts, "executeOperation", assets, amounts, premiums, arg3, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address , bytes params) returns(bool)
func (_Liquidatecontract *LiquidatecontractSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, arg3 common.Address, params []byte) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.ExecuteOperation(&_Liquidatecontract.TransactOpts, assets, amounts, premiums, arg3, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address , bytes params) returns(bool)
func (_Liquidatecontract *LiquidatecontractTransactorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, arg3 common.Address, params []byte) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.ExecuteOperation(&_Liquidatecontract.TransactOpts, assets, amounts, premiums, arg3, params)
}

// LiquidateLoan is a paid mutator transaction binding the contract method 0x39908e92.
//
// Solidity: function liquidateLoan(address _collateral, address _liquidate_asset, address _userToLiquidate, uint256 _amount, bool _receiveaToken) returns()
func (_Liquidatecontract *LiquidatecontractTransactor) LiquidateLoan(opts *bind.TransactOpts, _collateral common.Address, _liquidate_asset common.Address, _userToLiquidate common.Address, _amount *big.Int, _receiveaToken bool) (*types.Transaction, error) {
	return _Liquidatecontract.contract.Transact(opts, "liquidateLoan", _collateral, _liquidate_asset, _userToLiquidate, _amount, _receiveaToken)
}

// LiquidateLoan is a paid mutator transaction binding the contract method 0x39908e92.
//
// Solidity: function liquidateLoan(address _collateral, address _liquidate_asset, address _userToLiquidate, uint256 _amount, bool _receiveaToken) returns()
func (_Liquidatecontract *LiquidatecontractSession) LiquidateLoan(_collateral common.Address, _liquidate_asset common.Address, _userToLiquidate common.Address, _amount *big.Int, _receiveaToken bool) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.LiquidateLoan(&_Liquidatecontract.TransactOpts, _collateral, _liquidate_asset, _userToLiquidate, _amount, _receiveaToken)
}

// LiquidateLoan is a paid mutator transaction binding the contract method 0x39908e92.
//
// Solidity: function liquidateLoan(address _collateral, address _liquidate_asset, address _userToLiquidate, uint256 _amount, bool _receiveaToken) returns()
func (_Liquidatecontract *LiquidatecontractTransactorSession) LiquidateLoan(_collateral common.Address, _liquidate_asset common.Address, _userToLiquidate common.Address, _amount *big.Int, _receiveaToken bool) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.LiquidateLoan(&_Liquidatecontract.TransactOpts, _collateral, _liquidate_asset, _userToLiquidate, _amount, _receiveaToken)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquidatecontract *LiquidatecontractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidatecontract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquidatecontract *LiquidatecontractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Liquidatecontract.Contract.RenounceOwnership(&_Liquidatecontract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquidatecontract *LiquidatecontractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Liquidatecontract.Contract.RenounceOwnership(&_Liquidatecontract.TransactOpts)
}

// SwapToBarrowedAsset is a paid mutator transaction binding the contract method 0xa1f424e6.
//
// Solidity: function swapToBarrowedAsset(address asset_from, uint256 amountOutMin, address[] swapPath) returns()
func (_Liquidatecontract *LiquidatecontractTransactor) SwapToBarrowedAsset(opts *bind.TransactOpts, asset_from common.Address, amountOutMin *big.Int, swapPath []common.Address) (*types.Transaction, error) {
	return _Liquidatecontract.contract.Transact(opts, "swapToBarrowedAsset", asset_from, amountOutMin, swapPath)
}

// SwapToBarrowedAsset is a paid mutator transaction binding the contract method 0xa1f424e6.
//
// Solidity: function swapToBarrowedAsset(address asset_from, uint256 amountOutMin, address[] swapPath) returns()
func (_Liquidatecontract *LiquidatecontractSession) SwapToBarrowedAsset(asset_from common.Address, amountOutMin *big.Int, swapPath []common.Address) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.SwapToBarrowedAsset(&_Liquidatecontract.TransactOpts, asset_from, amountOutMin, swapPath)
}

// SwapToBarrowedAsset is a paid mutator transaction binding the contract method 0xa1f424e6.
//
// Solidity: function swapToBarrowedAsset(address asset_from, uint256 amountOutMin, address[] swapPath) returns()
func (_Liquidatecontract *LiquidatecontractTransactorSession) SwapToBarrowedAsset(asset_from common.Address, amountOutMin *big.Int, swapPath []common.Address) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.SwapToBarrowedAsset(&_Liquidatecontract.TransactOpts, asset_from, amountOutMin, swapPath)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquidatecontract *LiquidatecontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Liquidatecontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquidatecontract *LiquidatecontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.TransferOwnership(&_Liquidatecontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquidatecontract *LiquidatecontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Liquidatecontract.Contract.TransferOwnership(&_Liquidatecontract.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquidatecontract *LiquidatecontractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidatecontract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquidatecontract *LiquidatecontractSession) Receive() (*types.Transaction, error) {
	return _Liquidatecontract.Contract.Receive(&_Liquidatecontract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Liquidatecontract *LiquidatecontractTransactorSession) Receive() (*types.Transaction, error) {
	return _Liquidatecontract.Contract.Receive(&_Liquidatecontract.TransactOpts)
}

// LiquidatecontractErrorHandledIterator is returned from FilterErrorHandled and is used to iterate over the raw logs and unpacked data for ErrorHandled events raised by the Liquidatecontract contract.
type LiquidatecontractErrorHandledIterator struct {
	Event *LiquidatecontractErrorHandled // Event containing the contract specifics and raw log

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
func (it *LiquidatecontractErrorHandledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatecontractErrorHandled)
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
		it.Event = new(LiquidatecontractErrorHandled)
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
func (it *LiquidatecontractErrorHandledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatecontractErrorHandledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatecontractErrorHandled represents a ErrorHandled event raised by the Liquidatecontract contract.
type LiquidatecontractErrorHandled struct {
	StringFailure string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterErrorHandled is a free log retrieval operation binding the contract event 0x4f17d5cf0427b4fe5e5649c407d3021315cfaa7216c9e58c369ef3937c40801a.
//
// Solidity: event ErrorHandled(string stringFailure)
func (_Liquidatecontract *LiquidatecontractFilterer) FilterErrorHandled(opts *bind.FilterOpts) (*LiquidatecontractErrorHandledIterator, error) {

	logs, sub, err := _Liquidatecontract.contract.FilterLogs(opts, "ErrorHandled")
	if err != nil {
		return nil, err
	}
	return &LiquidatecontractErrorHandledIterator{contract: _Liquidatecontract.contract, event: "ErrorHandled", logs: logs, sub: sub}, nil
}

// WatchErrorHandled is a free log subscription operation binding the contract event 0x4f17d5cf0427b4fe5e5649c407d3021315cfaa7216c9e58c369ef3937c40801a.
//
// Solidity: event ErrorHandled(string stringFailure)
func (_Liquidatecontract *LiquidatecontractFilterer) WatchErrorHandled(opts *bind.WatchOpts, sink chan<- *LiquidatecontractErrorHandled) (event.Subscription, error) {

	logs, sub, err := _Liquidatecontract.contract.WatchLogs(opts, "ErrorHandled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatecontractErrorHandled)
				if err := _Liquidatecontract.contract.UnpackLog(event, "ErrorHandled", log); err != nil {
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

// ParseErrorHandled is a log parse operation binding the contract event 0x4f17d5cf0427b4fe5e5649c407d3021315cfaa7216c9e58c369ef3937c40801a.
//
// Solidity: event ErrorHandled(string stringFailure)
func (_Liquidatecontract *LiquidatecontractFilterer) ParseErrorHandled(log types.Log) (*LiquidatecontractErrorHandled, error) {
	event := new(LiquidatecontractErrorHandled)
	if err := _Liquidatecontract.contract.UnpackLog(event, "ErrorHandled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatecontractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Liquidatecontract contract.
type LiquidatecontractOwnershipTransferredIterator struct {
	Event *LiquidatecontractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LiquidatecontractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatecontractOwnershipTransferred)
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
		it.Event = new(LiquidatecontractOwnershipTransferred)
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
func (it *LiquidatecontractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatecontractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatecontractOwnershipTransferred represents a OwnershipTransferred event raised by the Liquidatecontract contract.
type LiquidatecontractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liquidatecontract *LiquidatecontractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidatecontractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Liquidatecontract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidatecontractOwnershipTransferredIterator{contract: _Liquidatecontract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liquidatecontract *LiquidatecontractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidatecontractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Liquidatecontract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatecontractOwnershipTransferred)
				if err := _Liquidatecontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Liquidatecontract *LiquidatecontractFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidatecontractOwnershipTransferred, error) {
	event := new(LiquidatecontractOwnershipTransferred)
	if err := _Liquidatecontract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
