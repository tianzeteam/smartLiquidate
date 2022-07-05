// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bakecontract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BakecontractABI is the input ABI used to generate the binding from.
const BakecontractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WBNB\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WBNB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBNBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidityBNB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBNB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveOut\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBNBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityBNB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBNB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBNBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityBNBSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBNB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBNBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityBNBWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBNB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBNBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityBNBWithPermitSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountBNB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"approveMax\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"removeLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapBNBForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactBNBForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactBNBForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForBNB\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForBNBSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactBNB\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Bakecontract is an auto generated Go binding around an Ethereum contract.
type Bakecontract struct {
	BakecontractCaller     // Read-only binding to the contract
	BakecontractTransactor // Write-only binding to the contract
	BakecontractFilterer   // Log filterer for contract events
}

// BakecontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BakecontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BakecontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BakecontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BakecontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BakecontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BakecontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BakecontractSession struct {
	Contract     *Bakecontract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BakecontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BakecontractCallerSession struct {
	Contract *BakecontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BakecontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BakecontractTransactorSession struct {
	Contract     *BakecontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BakecontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BakecontractRaw struct {
	Contract *Bakecontract // Generic contract binding to access the raw methods on
}

// BakecontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BakecontractCallerRaw struct {
	Contract *BakecontractCaller // Generic read-only contract binding to access the raw methods on
}

// BakecontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BakecontractTransactorRaw struct {
	Contract *BakecontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBakecontract creates a new instance of Bakecontract, bound to a specific deployed contract.
func NewBakecontract(address common.Address, backend bind.ContractBackend) (*Bakecontract, error) {
	contract, err := bindBakecontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bakecontract{BakecontractCaller: BakecontractCaller{contract: contract}, BakecontractTransactor: BakecontractTransactor{contract: contract}, BakecontractFilterer: BakecontractFilterer{contract: contract}}, nil
}

// NewBakecontractCaller creates a new read-only instance of Bakecontract, bound to a specific deployed contract.
func NewBakecontractCaller(address common.Address, caller bind.ContractCaller) (*BakecontractCaller, error) {
	contract, err := bindBakecontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BakecontractCaller{contract: contract}, nil
}

// NewBakecontractTransactor creates a new write-only instance of Bakecontract, bound to a specific deployed contract.
func NewBakecontractTransactor(address common.Address, transactor bind.ContractTransactor) (*BakecontractTransactor, error) {
	contract, err := bindBakecontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BakecontractTransactor{contract: contract}, nil
}

// NewBakecontractFilterer creates a new log filterer instance of Bakecontract, bound to a specific deployed contract.
func NewBakecontractFilterer(address common.Address, filterer bind.ContractFilterer) (*BakecontractFilterer, error) {
	contract, err := bindBakecontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BakecontractFilterer{contract: contract}, nil
}

// bindBakecontract binds a generic wrapper to an already deployed contract.
func bindBakecontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BakecontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bakecontract *BakecontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bakecontract.Contract.BakecontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bakecontract *BakecontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bakecontract.Contract.BakecontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bakecontract *BakecontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bakecontract.Contract.BakecontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bakecontract *BakecontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bakecontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bakecontract *BakecontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bakecontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bakecontract *BakecontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bakecontract.Contract.contract.Transact(opts, method, params...)
}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_Bakecontract *BakecontractCaller) WBNB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bakecontract.contract.Call(opts, &out, "WBNB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_Bakecontract *BakecontractSession) WBNB() (common.Address, error) {
	return _Bakecontract.Contract.WBNB(&_Bakecontract.CallOpts)
}

// WBNB is a free data retrieval call binding the contract method 0x8dd95002.
//
// Solidity: function WBNB() view returns(address)
func (_Bakecontract *BakecontractCallerSession) WBNB() (common.Address, error) {
	return _Bakecontract.Contract.WBNB(&_Bakecontract.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Bakecontract *BakecontractCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bakecontract.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Bakecontract *BakecontractSession) Factory() (common.Address, error) {
	return _Bakecontract.Contract.Factory(&_Bakecontract.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Bakecontract *BakecontractCallerSession) Factory() (common.Address, error) {
	return _Bakecontract.Contract.Factory(&_Bakecontract.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Bakecontract *BakecontractCaller) GetAmountIn(opts *bind.CallOpts, amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bakecontract.contract.Call(opts, &out, "getAmountIn", amountOut, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Bakecontract *BakecontractSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Bakecontract.Contract.GetAmountIn(&_Bakecontract.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountIn is a free data retrieval call binding the contract method 0x85f8c259.
//
// Solidity: function getAmountIn(uint256 amountOut, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountIn)
func (_Bakecontract *BakecontractCallerSession) GetAmountIn(amountOut *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Bakecontract.Contract.GetAmountIn(&_Bakecontract.CallOpts, amountOut, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Bakecontract *BakecontractCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bakecontract.contract.Call(opts, &out, "getAmountOut", amountIn, reserveIn, reserveOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Bakecontract *BakecontractSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Bakecontract.Contract.GetAmountOut(&_Bakecontract.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x054d50d4.
//
// Solidity: function getAmountOut(uint256 amountIn, uint256 reserveIn, uint256 reserveOut) pure returns(uint256 amountOut)
func (_Bakecontract *BakecontractCallerSession) GetAmountOut(amountIn *big.Int, reserveIn *big.Int, reserveOut *big.Int) (*big.Int, error) {
	return _Bakecontract.Contract.GetAmountOut(&_Bakecontract.CallOpts, amountIn, reserveIn, reserveOut)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Bakecontract *BakecontractCaller) GetAmountsIn(opts *bind.CallOpts, amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Bakecontract.contract.Call(opts, &out, "getAmountsIn", amountOut, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Bakecontract *BakecontractSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Bakecontract.Contract.GetAmountsIn(&_Bakecontract.CallOpts, amountOut, path)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0x1f00ca74.
//
// Solidity: function getAmountsIn(uint256 amountOut, address[] path) view returns(uint256[] amounts)
func (_Bakecontract *BakecontractCallerSession) GetAmountsIn(amountOut *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Bakecontract.Contract.GetAmountsIn(&_Bakecontract.CallOpts, amountOut, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Bakecontract *BakecontractCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Bakecontract.contract.Call(opts, &out, "getAmountsOut", amountIn, path)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Bakecontract *BakecontractSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Bakecontract.Contract.GetAmountsOut(&_Bakecontract.CallOpts, amountIn, path)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0xd06ca61f.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path) view returns(uint256[] amounts)
func (_Bakecontract *BakecontractCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address) ([]*big.Int, error) {
	return _Bakecontract.Contract.GetAmountsOut(&_Bakecontract.CallOpts, amountIn, path)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Bakecontract *BakecontractCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bakecontract.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Bakecontract *BakecontractSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Bakecontract.Contract.Quote(&_Bakecontract.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_Bakecontract *BakecontractCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _Bakecontract.Contract.Quote(&_Bakecontract.CallOpts, amountA, reserveA, reserveB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Bakecontract *BakecontractTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Bakecontract *BakecontractSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.AddLiquidity(&_Bakecontract.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8e33700.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Bakecontract *BakecontractTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.AddLiquidity(&_Bakecontract.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline)
}

// AddLiquidityBNB is a paid mutator transaction binding the contract method 0xeaaed442.
//
// Solidity: function addLiquidityBNB(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountBNB, uint256 liquidity)
func (_Bakecontract *BakecontractTransactor) AddLiquidityBNB(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "addLiquidityBNB", token, amountTokenDesired, amountTokenMin, amountBNBMin, to, deadline)
}

// AddLiquidityBNB is a paid mutator transaction binding the contract method 0xeaaed442.
//
// Solidity: function addLiquidityBNB(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountBNB, uint256 liquidity)
func (_Bakecontract *BakecontractSession) AddLiquidityBNB(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.AddLiquidityBNB(&_Bakecontract.TransactOpts, token, amountTokenDesired, amountTokenMin, amountBNBMin, to, deadline)
}

// AddLiquidityBNB is a paid mutator transaction binding the contract method 0xeaaed442.
//
// Solidity: function addLiquidityBNB(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) payable returns(uint256 amountToken, uint256 amountBNB, uint256 liquidity)
func (_Bakecontract *BakecontractTransactorSession) AddLiquidityBNB(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.AddLiquidityBNB(&_Bakecontract.TransactOpts, token, amountTokenDesired, amountTokenMin, amountBNBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Bakecontract *BakecontractTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Bakecontract *BakecontractSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidity(&_Bakecontract.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xbaa2abde.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline) returns(uint256 amountA, uint256 amountB)
func (_Bakecontract *BakecontractTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidity(&_Bakecontract.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline)
}

// RemoveLiquidityBNB is a paid mutator transaction binding the contract method 0xe0588488.
//
// Solidity: function removeLiquidityBNB(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountBNB)
func (_Bakecontract *BakecontractTransactor) RemoveLiquidityBNB(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "removeLiquidityBNB", token, liquidity, amountTokenMin, amountBNBMin, to, deadline)
}

// RemoveLiquidityBNB is a paid mutator transaction binding the contract method 0xe0588488.
//
// Solidity: function removeLiquidityBNB(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountBNB)
func (_Bakecontract *BakecontractSession) RemoveLiquidityBNB(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityBNB(&_Bakecontract.TransactOpts, token, liquidity, amountTokenMin, amountBNBMin, to, deadline)
}

// RemoveLiquidityBNB is a paid mutator transaction binding the contract method 0xe0588488.
//
// Solidity: function removeLiquidityBNB(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) returns(uint256 amountToken, uint256 amountBNB)
func (_Bakecontract *BakecontractTransactorSession) RemoveLiquidityBNB(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityBNB(&_Bakecontract.TransactOpts, token, liquidity, amountTokenMin, amountBNBMin, to, deadline)
}

// RemoveLiquidityBNBSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb7e65949.
//
// Solidity: function removeLiquidityBNBSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) returns(uint256 amountBNB)
func (_Bakecontract *BakecontractTransactor) RemoveLiquidityBNBSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "removeLiquidityBNBSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountBNBMin, to, deadline)
}

// RemoveLiquidityBNBSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb7e65949.
//
// Solidity: function removeLiquidityBNBSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) returns(uint256 amountBNB)
func (_Bakecontract *BakecontractSession) RemoveLiquidityBNBSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityBNBSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, token, liquidity, amountTokenMin, amountBNBMin, to, deadline)
}

// RemoveLiquidityBNBSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xb7e65949.
//
// Solidity: function removeLiquidityBNBSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline) returns(uint256 amountBNB)
func (_Bakecontract *BakecontractTransactorSession) RemoveLiquidityBNBSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityBNBSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, token, liquidity, amountTokenMin, amountBNBMin, to, deadline)
}

// RemoveLiquidityBNBWithPermit is a paid mutator transaction binding the contract method 0x34a0772d.
//
// Solidity: function removeLiquidityBNBWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountBNB)
func (_Bakecontract *BakecontractTransactor) RemoveLiquidityBNBWithPermit(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "removeLiquidityBNBWithPermit", token, liquidity, amountTokenMin, amountBNBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityBNBWithPermit is a paid mutator transaction binding the contract method 0x34a0772d.
//
// Solidity: function removeLiquidityBNBWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountBNB)
func (_Bakecontract *BakecontractSession) RemoveLiquidityBNBWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityBNBWithPermit(&_Bakecontract.TransactOpts, token, liquidity, amountTokenMin, amountBNBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityBNBWithPermit is a paid mutator transaction binding the contract method 0x34a0772d.
//
// Solidity: function removeLiquidityBNBWithPermit(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountToken, uint256 amountBNB)
func (_Bakecontract *BakecontractTransactorSession) RemoveLiquidityBNBWithPermit(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityBNBWithPermit(&_Bakecontract.TransactOpts, token, liquidity, amountTokenMin, amountBNBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityBNBWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x685a0a34.
//
// Solidity: function removeLiquidityBNBWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountBNB)
func (_Bakecontract *BakecontractTransactor) RemoveLiquidityBNBWithPermitSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "removeLiquidityBNBWithPermitSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountBNBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityBNBWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x685a0a34.
//
// Solidity: function removeLiquidityBNBWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountBNB)
func (_Bakecontract *BakecontractSession) RemoveLiquidityBNBWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityBNBWithPermitSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, token, liquidity, amountTokenMin, amountBNBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityBNBWithPermitSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x685a0a34.
//
// Solidity: function removeLiquidityBNBWithPermitSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountBNBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountBNB)
func (_Bakecontract *BakecontractTransactorSession) RemoveLiquidityBNBWithPermitSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountBNBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityBNBWithPermitSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, token, liquidity, amountTokenMin, amountBNBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Bakecontract *BakecontractTransactor) RemoveLiquidityWithPermit(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "removeLiquidityWithPermit", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Bakecontract *BakecontractSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityWithPermit(&_Bakecontract.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// RemoveLiquidityWithPermit is a paid mutator transaction binding the contract method 0x2195995c.
//
// Solidity: function removeLiquidityWithPermit(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool approveMax, uint8 v, bytes32 r, bytes32 s) returns(uint256 amountA, uint256 amountB)
func (_Bakecontract *BakecontractTransactorSession) RemoveLiquidityWithPermit(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, approveMax bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Bakecontract.Contract.RemoveLiquidityWithPermit(&_Bakecontract.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, approveMax, v, r, s)
}

// SwapBNBForExactTokens is a paid mutator transaction binding the contract method 0x8332a963.
//
// Solidity: function swapBNBForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactor) SwapBNBForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapBNBForExactTokens", amountOut, path, to, deadline)
}

// SwapBNBForExactTokens is a paid mutator transaction binding the contract method 0x8332a963.
//
// Solidity: function swapBNBForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Bakecontract *BakecontractSession) SwapBNBForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapBNBForExactTokens(&_Bakecontract.TransactOpts, amountOut, path, to, deadline)
}

// SwapBNBForExactTokens is a paid mutator transaction binding the contract method 0x8332a963.
//
// Solidity: function swapBNBForExactTokens(uint256 amountOut, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactorSession) SwapBNBForExactTokens(amountOut *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapBNBForExactTokens(&_Bakecontract.TransactOpts, amountOut, path, to, deadline)
}

// SwapExactBNBForTokens is a paid mutator transaction binding the contract method 0x9cf68911.
//
// Solidity: function swapExactBNBForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactor) SwapExactBNBForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapExactBNBForTokens", amountOutMin, path, to, deadline)
}

// SwapExactBNBForTokens is a paid mutator transaction binding the contract method 0x9cf68911.
//
// Solidity: function swapExactBNBForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Bakecontract *BakecontractSession) SwapExactBNBForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactBNBForTokens(&_Bakecontract.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactBNBForTokens is a paid mutator transaction binding the contract method 0x9cf68911.
//
// Solidity: function swapExactBNBForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactorSession) SwapExactBNBForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactBNBForTokens(&_Bakecontract.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactBNBForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x50e27df3.
//
// Solidity: function swapExactBNBForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Bakecontract *BakecontractTransactor) SwapExactBNBForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapExactBNBForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline)
}

// SwapExactBNBForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x50e27df3.
//
// Solidity: function swapExactBNBForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Bakecontract *BakecontractSession) SwapExactBNBForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactBNBForTokensSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactBNBForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x50e27df3.
//
// Solidity: function swapExactBNBForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline) payable returns()
func (_Bakecontract *BakecontractTransactorSession) SwapExactBNBForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactBNBForTokensSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, amountOutMin, path, to, deadline)
}

// SwapExactTokensForBNB is a paid mutator transaction binding the contract method 0x5d616c5b.
//
// Solidity: function swapExactTokensForBNB(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactor) SwapExactTokensForBNB(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapExactTokensForBNB", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForBNB is a paid mutator transaction binding the contract method 0x5d616c5b.
//
// Solidity: function swapExactTokensForBNB(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractSession) SwapExactTokensForBNB(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactTokensForBNB(&_Bakecontract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForBNB is a paid mutator transaction binding the contract method 0x5d616c5b.
//
// Solidity: function swapExactTokensForBNB(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactorSession) SwapExactTokensForBNB(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactTokensForBNB(&_Bakecontract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForBNBSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xd46d2f83.
//
// Solidity: function swapExactTokensForBNBSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Bakecontract *BakecontractTransactor) SwapExactTokensForBNBSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapExactTokensForBNBSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForBNBSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xd46d2f83.
//
// Solidity: function swapExactTokensForBNBSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Bakecontract *BakecontractSession) SwapExactTokensForBNBSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactTokensForBNBSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForBNBSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xd46d2f83.
//
// Solidity: function swapExactTokensForBNBSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Bakecontract *BakecontractTransactorSession) SwapExactTokensForBNBSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactTokensForBNBSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactTokensForTokens(&_Bakecontract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactTokensForTokens(&_Bakecontract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Bakecontract *BakecontractTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Bakecontract *BakecontractSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x5c11d795.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline) returns()
func (_Bakecontract *BakecontractTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_Bakecontract.TransactOpts, amountIn, amountOutMin, path, to, deadline)
}

// SwapTokensForExactBNB is a paid mutator transaction binding the contract method 0xd67b571e.
//
// Solidity: function swapTokensForExactBNB(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactor) SwapTokensForExactBNB(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapTokensForExactBNB", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactBNB is a paid mutator transaction binding the contract method 0xd67b571e.
//
// Solidity: function swapTokensForExactBNB(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractSession) SwapTokensForExactBNB(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapTokensForExactBNB(&_Bakecontract.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactBNB is a paid mutator transaction binding the contract method 0xd67b571e.
//
// Solidity: function swapTokensForExactBNB(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactorSession) SwapTokensForExactBNB(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapTokensForExactBNB(&_Bakecontract.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapTokensForExactTokens(&_Bakecontract.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path, address to, uint256 deadline) returns(uint256[] amounts)
func (_Bakecontract *BakecontractTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address, to common.Address, deadline *big.Int) (*types.Transaction, error) {
	return _Bakecontract.Contract.SwapTokensForExactTokens(&_Bakecontract.TransactOpts, amountOut, amountInMax, path, to, deadline)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bakecontract *BakecontractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bakecontract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bakecontract *BakecontractSession) Receive() (*types.Transaction, error) {
	return _Bakecontract.Contract.Receive(&_Bakecontract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bakecontract *BakecontractTransactorSession) Receive() (*types.Transaction, error) {
	return _Bakecontract.Contract.Receive(&_Bakecontract.TransactOpts)
}
