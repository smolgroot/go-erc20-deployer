// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

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

// GrootCoinMetaData contains all meta data concerning the GrootCoin contract.
var GrootCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526a396c41bd9e54ada380000060025534801561001f57600080fd5b506002546000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d2b806100746000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063313ce56711610066578063313ce5671461013457806370a082311461015257806395d89b4114610182578063a9059cbb146101a0578063dd62ed3e146101d057610093565b806306fdde0314610098578063095ea7b3146100b657806318160ddd146100e657806323b872dd14610104575b600080fd5b6100a0610200565b6040516100ad9190610a0c565b60405180910390f35b6100d060048036038101906100cb9190610ac7565b610239565b6040516100dd9190610b22565b60405180910390f35b6100ee61032b565b6040516100fb9190610b4c565b60405180910390f35b61011e60048036038101906101199190610b67565b610335565b60405161012b9190610b22565b60405180910390f35b61013c61069b565b6040516101499190610bd6565b60405180910390f35b61016c60048036038101906101679190610bf1565b6106a0565b6040516101799190610b4c565b60405180910390f35b61018a6106e8565b6040516101979190610a0c565b60405180910390f35b6101ba60048036038101906101b59190610ac7565b610721565b6040516101c79190610b22565b60405180910390f35b6101ea60048036038101906101e59190610c1e565b6108f5565b6040516101f79190610b4c565b60405180910390f35b6040518060400160405280600981526020017f47726f6f74436f696e000000000000000000000000000000000000000000000081525081565b600081600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516103199190610b4c565b60405180910390a36001905092915050565b6000600254905090565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111561038257600080fd5b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111561040b57600080fd5b816000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546104559190610c8d565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461051f9190610c8d565b600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105e99190610cc1565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516106889190610b4c565b60405180910390a3600190509392505050565b601281565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6040518060400160405280600581526020017f47524f4f5400000000000000000000000000000000000000000000000000000081525081565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111561076e57600080fd5b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546107b89190610c8d565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546108449190610cc1565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108e39190610b4c565b60405180910390a36001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109b657808201518184015260208101905061099b565b60008484015250505050565b6000601f19601f8301169050919050565b60006109de8261097c565b6109e88185610987565b93506109f8818560208601610998565b610a01816109c2565b840191505092915050565b60006020820190508181036000830152610a2681846109d3565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a5e82610a33565b9050919050565b610a6e81610a53565b8114610a7957600080fd5b50565b600081359050610a8b81610a65565b92915050565b6000819050919050565b610aa481610a91565b8114610aaf57600080fd5b50565b600081359050610ac181610a9b565b92915050565b60008060408385031215610ade57610add610a2e565b5b6000610aec85828601610a7c565b9250506020610afd85828601610ab2565b9150509250929050565b60008115159050919050565b610b1c81610b07565b82525050565b6000602082019050610b376000830184610b13565b92915050565b610b4681610a91565b82525050565b6000602082019050610b616000830184610b3d565b92915050565b600080600060608486031215610b8057610b7f610a2e565b5b6000610b8e86828701610a7c565b9350506020610b9f86828701610a7c565b9250506040610bb086828701610ab2565b9150509250925092565b600060ff82169050919050565b610bd081610bba565b82525050565b6000602082019050610beb6000830184610bc7565b92915050565b600060208284031215610c0757610c06610a2e565b5b6000610c1584828501610a7c565b91505092915050565b60008060408385031215610c3557610c34610a2e565b5b6000610c4385828601610a7c565b9250506020610c5485828601610a7c565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c9882610a91565b9150610ca383610a91565b9250828203905081811115610cbb57610cba610c5e565b5b92915050565b6000610ccc82610a91565b9150610cd783610a91565b9250828201905080821115610cef57610cee610c5e565b5b9291505056fea2646970667358221220db4440c69e820c04b6e30e30f9288c90bd98478753f11976ef151e8440ccccd664736f6c63430008130033",
}

// GrootCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use GrootCoinMetaData.ABI instead.
var GrootCoinABI = GrootCoinMetaData.ABI

// GrootCoinBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GrootCoinMetaData.Bin instead.
var GrootCoinBin = GrootCoinMetaData.Bin

// DeployGrootCoin deploys a new Ethereum contract, binding an instance of GrootCoin to it.
func DeployGrootCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GrootCoin, error) {
	parsed, err := GrootCoinMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GrootCoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GrootCoin{GrootCoinCaller: GrootCoinCaller{contract: contract}, GrootCoinTransactor: GrootCoinTransactor{contract: contract}, GrootCoinFilterer: GrootCoinFilterer{contract: contract}}, nil
}

// GrootCoin is an auto generated Go binding around an Ethereum contract.
type GrootCoin struct {
	GrootCoinCaller     // Read-only binding to the contract
	GrootCoinTransactor // Write-only binding to the contract
	GrootCoinFilterer   // Log filterer for contract events
}

// GrootCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type GrootCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrootCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GrootCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrootCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GrootCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GrootCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GrootCoinSession struct {
	Contract     *GrootCoin        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GrootCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GrootCoinCallerSession struct {
	Contract *GrootCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GrootCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GrootCoinTransactorSession struct {
	Contract     *GrootCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GrootCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type GrootCoinRaw struct {
	Contract *GrootCoin // Generic contract binding to access the raw methods on
}

// GrootCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GrootCoinCallerRaw struct {
	Contract *GrootCoinCaller // Generic read-only contract binding to access the raw methods on
}

// GrootCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GrootCoinTransactorRaw struct {
	Contract *GrootCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGrootCoin creates a new instance of GrootCoin, bound to a specific deployed contract.
func NewGrootCoin(address common.Address, backend bind.ContractBackend) (*GrootCoin, error) {
	contract, err := bindGrootCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GrootCoin{GrootCoinCaller: GrootCoinCaller{contract: contract}, GrootCoinTransactor: GrootCoinTransactor{contract: contract}, GrootCoinFilterer: GrootCoinFilterer{contract: contract}}, nil
}

// NewGrootCoinCaller creates a new read-only instance of GrootCoin, bound to a specific deployed contract.
func NewGrootCoinCaller(address common.Address, caller bind.ContractCaller) (*GrootCoinCaller, error) {
	contract, err := bindGrootCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GrootCoinCaller{contract: contract}, nil
}

// NewGrootCoinTransactor creates a new write-only instance of GrootCoin, bound to a specific deployed contract.
func NewGrootCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*GrootCoinTransactor, error) {
	contract, err := bindGrootCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GrootCoinTransactor{contract: contract}, nil
}

// NewGrootCoinFilterer creates a new log filterer instance of GrootCoin, bound to a specific deployed contract.
func NewGrootCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*GrootCoinFilterer, error) {
	contract, err := bindGrootCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GrootCoinFilterer{contract: contract}, nil
}

// bindGrootCoin binds a generic wrapper to an already deployed contract.
func bindGrootCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GrootCoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GrootCoin *GrootCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GrootCoin.Contract.GrootCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GrootCoin *GrootCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GrootCoin.Contract.GrootCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GrootCoin *GrootCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GrootCoin.Contract.GrootCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GrootCoin *GrootCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GrootCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GrootCoin *GrootCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GrootCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GrootCoin *GrootCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GrootCoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_GrootCoin *GrootCoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, delegate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GrootCoin.contract.Call(opts, &out, "allowance", owner, delegate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_GrootCoin *GrootCoinSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _GrootCoin.Contract.Allowance(&_GrootCoin.CallOpts, owner, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_GrootCoin *GrootCoinCallerSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _GrootCoin.Contract.Allowance(&_GrootCoin.CallOpts, owner, delegate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_GrootCoin *GrootCoinCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GrootCoin.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_GrootCoin *GrootCoinSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _GrootCoin.Contract.BalanceOf(&_GrootCoin.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_GrootCoin *GrootCoinCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _GrootCoin.Contract.BalanceOf(&_GrootCoin.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GrootCoin *GrootCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GrootCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GrootCoin *GrootCoinSession) Decimals() (uint8, error) {
	return _GrootCoin.Contract.Decimals(&_GrootCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GrootCoin *GrootCoinCallerSession) Decimals() (uint8, error) {
	return _GrootCoin.Contract.Decimals(&_GrootCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GrootCoin *GrootCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GrootCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GrootCoin *GrootCoinSession) Name() (string, error) {
	return _GrootCoin.Contract.Name(&_GrootCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GrootCoin *GrootCoinCallerSession) Name() (string, error) {
	return _GrootCoin.Contract.Name(&_GrootCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GrootCoin *GrootCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GrootCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GrootCoin *GrootCoinSession) Symbol() (string, error) {
	return _GrootCoin.Contract.Symbol(&_GrootCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GrootCoin *GrootCoinCallerSession) Symbol() (string, error) {
	return _GrootCoin.Contract.Symbol(&_GrootCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GrootCoin *GrootCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GrootCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GrootCoin *GrootCoinSession) TotalSupply() (*big.Int, error) {
	return _GrootCoin.Contract.TotalSupply(&_GrootCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GrootCoin *GrootCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _GrootCoin.Contract.TotalSupply(&_GrootCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.contract.Transact(opts, "approve", delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.Contract.Approve(&_GrootCoin.TransactOpts, delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinTransactorSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.Contract.Approve(&_GrootCoin.TransactOpts, delegate, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.contract.Transact(opts, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.Contract.Transfer(&_GrootCoin.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinTransactorSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.Contract.Transfer(&_GrootCoin.TransactOpts, receiver, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.contract.Transact(opts, "transferFrom", owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.Contract.TransferFrom(&_GrootCoin.TransactOpts, owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_GrootCoin *GrootCoinTransactorSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _GrootCoin.Contract.TransferFrom(&_GrootCoin.TransactOpts, owner, buyer, numTokens)
}

// GrootCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GrootCoin contract.
type GrootCoinApprovalIterator struct {
	Event *GrootCoinApproval // Event containing the contract specifics and raw log

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
func (it *GrootCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GrootCoinApproval)
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
		it.Event = new(GrootCoinApproval)
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
func (it *GrootCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GrootCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GrootCoinApproval represents a Approval event raised by the GrootCoin contract.
type GrootCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GrootCoin *GrootCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GrootCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GrootCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GrootCoinApprovalIterator{contract: _GrootCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GrootCoin *GrootCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GrootCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GrootCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GrootCoinApproval)
				if err := _GrootCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GrootCoin *GrootCoinFilterer) ParseApproval(log types.Log) (*GrootCoinApproval, error) {
	event := new(GrootCoinApproval)
	if err := _GrootCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GrootCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GrootCoin contract.
type GrootCoinTransferIterator struct {
	Event *GrootCoinTransfer // Event containing the contract specifics and raw log

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
func (it *GrootCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GrootCoinTransfer)
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
		it.Event = new(GrootCoinTransfer)
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
func (it *GrootCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GrootCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GrootCoinTransfer represents a Transfer event raised by the GrootCoin contract.
type GrootCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GrootCoin *GrootCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GrootCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GrootCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GrootCoinTransferIterator{contract: _GrootCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GrootCoin *GrootCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GrootCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GrootCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GrootCoinTransfer)
				if err := _GrootCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GrootCoin *GrootCoinFilterer) ParseTransfer(log types.Log) (*GrootCoinTransfer, error) {
	event := new(GrootCoinTransfer)
	if err := _GrootCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
