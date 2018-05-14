// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package mytoken

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// MyTokenABI is the input ABI used to generate the binding from.
const MyTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint64\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"spentAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":false,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"s\",\"type\":\"string\"}],\"name\":\"StringArg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"i8\",\"type\":\"int8\"},{\"indexed\":true,\"name\":\"i64\",\"type\":\"int64\"}],\"name\":\"IntArg\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"ReceiveApproval\",\"type\":\"event\"}]"

// MyTokenBin is the compiled bytecode used for deploying new contracts.
const MyTokenBin = `0x6060604052341561000f57600080fd5b604051610a07380380610a0783398101604052808051919060200180518201919060200180519190602001805190910190505b600160a060020a033316600090815260036020526040812085905583805161006e92916020019061009c565b50600181805161008292916020019061009c565b506002805460ff191660ff84161790555b5050505061013c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100dd57805160ff191683800117855561010a565b8280016001018555821561010a579182015b8281111561010a5782518255916020019190600101906100ef565b5b5061011792915061011b565b5090565b61013991905b808211156101175760008155600101610121565b5090565b90565b6108bc8061014b6000396000f300606060405236156100965763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100ae57806323b872dd14610139578063313ce567146101755780635d359fbd1461019e57806370a08231146101cc57806395d89b41146101fd578063cae9ca5114610288578063dc3080f214610301578063dd62ed3e14610338575b34156100a157600080fd5b6100ac5b600080fd5b565b005b34156100b957600080fd5b6100c161036f565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100fe5780820151818401525b6020016100e5565b50505050905090810190601f16801561012b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561014457600080fd5b610161600160a060020a036004358116906024351660443561040d565b604051901515815260200160405180910390f35b341561018057600080fd5b610188610535565b60405160ff909116815260200160405180910390f35b34156101a957600080fd5b6100ac600160a060020a036004351667ffffffffffffffff6024351661053e565b005b34156101d757600080fd5b6101eb600160a060020a03600435166106b4565b60405190815260200160405180910390f35b341561020857600080fd5b6100c16106c6565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100fe5780820151818401525b6020016100e5565b50505050905090810190601f16801561012b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561029357600080fd5b61016160048035600160a060020a03169060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061076495505050505050565b604051901515815260200160405180910390f35b341561030c57600080fd5b6101eb600160a060020a0360043581169060243516610856565b60405190815260200160405180910390f35b341561034357600080fd5b6101eb600160a060020a0360043581169060243516610873565b60405190815260200160405180910390f35b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104055780601f106103da57610100808354040283529160200191610405565b820191906000526020600020905b8154815290600101906020018083116103e857829003601f168201915b505050505081565b600160a060020a0383166000908152600360205260408120548290101561043357600080fd5b600160a060020a038316600090815260036020526040902054828101101561045a57600080fd5b600160a060020a0380851660008181526004602090815260408083203390951680845294825280832054938352600582528083209483529390529190912054830111156104a657600080fd5b600160a060020a03808516600081815260036020908152604080832080548890039055878516808452818420805489019055938352600582528083203390951680845294909152908190208054860190559091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35b9392505050565b60025460ff1681565b600160a060020a03331660009081526003602052604090205467ffffffffffffffff821690101561056e57600080fd5b600160a060020a03821660009081526003602052604090205467ffffffffffffffff82168101101561059f57600080fd5b600160a060020a03338116600081815260036020526040808220805467ffffffffffffffff871690819003909155938616808352918190208054909401909355917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9084905167ffffffffffffffff909116815260200160405180910390a36040517f7465737400000000000000000000000000000000000000000000000000000000815260040160405180910390207f32aa877c63867cefeb8e6762a30e9d3624d28d9b2876bce43cba3b69fd5a308060405160405180910390a28060070b8160000b7f61e876266a7198592d33f12d6c52e88c9e5002c1ad9317a38343df20deef7a5460405160405180910390a35b5050565b60036020526000908152604090205481565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104055780601f106103da57610100808354040283529160200191610405565b820191906000526020600020905b8154815290600101906020018083116103e857829003601f168201915b505050505081565b33600160a060020a0381811660009081526004602090815260408083209388168352929052818120859055917fc580975836a6474723779cc3df05b753a9d649f7125258b85e3a94ad063a98e79185903090869051600160a060020a038086168252602082018590528316604082015260806060820181815290820183818151815260200191508051906020019080838360005b838110156108115780820151818401525b6020016107f8565b50505050905090810190601f16801561083e5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a15b9392505050565b600560209081526000928352604080842090915290825290205481565b6004602090815260009283526040808420909152908252902054815600a165627a7a72305820048a926722b87eaf795a29da0520cdb5aa57068d60cc363d51a5097e082195470029`

// DeployMyToken deploys a new Ethereum contract, binding an instance of MyToken to it.
func DeployMyToken(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string) (common.Address, *types.Transaction, *MyToken, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyTokenBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}}, nil
}

// MyToken is an auto generated Go binding around an Ethereum contract.
type MyToken struct {
	MyTokenCaller     // Read-only binding to the contract
	MyTokenTransactor // Write-only binding to the contract
}

// MyTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyTokenSession struct {
	Contract     *MyToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyTokenCallerSession struct {
	Contract *MyTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyTokenTransactorSession struct {
	Contract     *MyTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyTokenRaw struct {
	Contract *MyToken // Generic contract binding to access the raw methods on
}

// MyTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyTokenCallerRaw struct {
	Contract *MyTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MyTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyTokenTransactorRaw struct {
	Contract *MyTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyToken creates a new instance of MyToken, bound to a specific deployed contract.
func NewMyToken(address common.Address, backend bind.ContractBackend) (*MyToken, error) {
	contract, err := bindMyToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyToken{MyTokenCaller: MyTokenCaller{contract: contract}, MyTokenTransactor: MyTokenTransactor{contract: contract}}, nil
}

// NewMyTokenCaller creates a new read-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenCaller(address common.Address, caller bind.ContractCaller) (*MyTokenCaller, error) {
	contract, err := bindMyToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MyTokenCaller{contract: contract}, nil
}

// NewMyTokenTransactor creates a new write-only instance of MyToken, bound to a specific deployed contract.
func NewMyTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MyTokenTransactor, error) {
	contract, err := bindMyToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MyTokenTransactor{contract: contract}, nil
}

// bindMyToken binds a generic wrapper to an already deployed contract.
func bindMyToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.MyTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyToken.Contract.MyTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyToken *MyTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyToken *MyTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyToken *MyTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyToken.Contract.Allowance(&_MyToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyToken.Contract.Allowance(&_MyToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyToken *MyTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyToken *MyTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_MyToken *MyTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _MyToken.Contract.BalanceOf(&_MyToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyToken *MyTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyToken *MyTokenSession) Decimals() (uint8, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyToken *MyTokenCallerSession) Decimals() (uint8, error) {
	return _MyToken.Contract.Decimals(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyToken *MyTokenCallerSession) Name() (string, error) {
	return _MyToken.Contract.Name(&_MyToken.CallOpts)
}

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenCaller) SpentAllowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "spentAllowance", arg0, arg1)
	return *ret0, err
}

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenSession) SpentAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyToken.Contract.SpentAllowance(&_MyToken.CallOpts, arg0, arg1)
}

// SpentAllowance is a free data retrieval call binding the contract method 0xdc3080f2.
//
// Solidity: function spentAllowance( address,  address) constant returns(uint256)
func (_MyToken *MyTokenCallerSession) SpentAllowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MyToken.Contract.SpentAllowance(&_MyToken.CallOpts, arg0, arg1)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyToken *MyTokenCallerSession) Symbol() (string, error) {
	return _MyToken.Contract.Symbol(&_MyToken.CallOpts)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyToken *MyTokenTransactor) ApproveAndCall(opts *bind.TransactOpts, _spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "approveAndCall", _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyToken *MyTokenSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyToken.Contract.ApproveAndCall(&_MyToken.TransactOpts, _spender, _value, _extraData)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(_spender address, _value uint256, _extraData bytes) returns(success bool)
func (_MyToken *MyTokenTransactorSession) ApproveAndCall(_spender common.Address, _value *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _MyToken.Contract.ApproveAndCall(&_MyToken.TransactOpts, _spender, _value, _extraData)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(_to address, _value uint64) returns()
func (_MyToken *MyTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value uint64) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(_to address, _value uint64) returns()
func (_MyToken *MyTokenSession) Transfer(_to common.Address, _value uint64) (*types.Transaction, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0x5d359fbd.
//
// Solidity: function transfer(_to address, _value uint64) returns()
func (_MyToken *MyTokenTransactorSession) Transfer(_to common.Address, _value uint64) (*types.Transaction, error) {
	return _MyToken.Contract.Transfer(&_MyToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyToken *MyTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyToken *MyTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.TransferFrom(&_MyToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_MyToken *MyTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MyToken.Contract.TransferFrom(&_MyToken.TransactOpts, _from, _to, _value)
}

func (_MyToken *MyTokenCaller) EventSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber, eventName string) (<-chan types.Log, ethereum.Subscription, error) {
	parsed, err := abi.JSON(strings.NewReader(MyTokenABI))
	if err != nil {
		return nil, nil, err
	}
	var q ethereum.FilterQuery
	q.FromBlock = big.NewInt(int64(fromBlock))
	if toBlock == rpc.LatestBlockNumber {
		q.ToBlock = nil
	} else {
		q.ToBlock = big.NewInt(int64(toBlock))
	}
	q.Topics = [][]common.Hash{
		{parsed.Events[eventName].Id()}, //event signature
	}
	return _MyToken.contract.SubscribeFilterLogs(opts, q)
}

/*
   get all event IntArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCaller) EventIntArgSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.EventSubscribe(opts, fromBlock, toBlock, "IntArg")
}

/*
   get all event IntArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenSession) EventIntArgSubscribe(fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventIntArgSubscribe(&t.CallOpts, fromBlock, toBlock)
}

/*
   get all event IntArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCallerSession) EventIntArgSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventIntArgSubscribe(&t.CallOpts, fromBlock, toBlock)
}

/*
   get all event ReceiveApproval happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCaller) EventReceiveApprovalSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.EventSubscribe(opts, fromBlock, toBlock, "ReceiveApproval")
}

/*
   get all event ReceiveApproval happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenSession) EventReceiveApprovalSubscribe(fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventReceiveApprovalSubscribe(&t.CallOpts, fromBlock, toBlock)
}

/*
   get all event ReceiveApproval happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCallerSession) EventReceiveApprovalSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventReceiveApprovalSubscribe(&t.CallOpts, fromBlock, toBlock)
}

/*
   get all event StringArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCaller) EventStringArgSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.EventSubscribe(opts, fromBlock, toBlock, "StringArg")
}

/*
   get all event StringArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenSession) EventStringArgSubscribe(fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventStringArgSubscribe(&t.CallOpts, fromBlock, toBlock)
}

/*
   get all event StringArg happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCallerSession) EventStringArgSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventStringArgSubscribe(&t.CallOpts, fromBlock, toBlock)
}

/*
   get all event Transfer happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCaller) EventTransferSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.EventSubscribe(opts, fromBlock, toBlock, "Transfer")
}

/*
   get all event Transfer happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenSession) EventTransferSubscribe(fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventTransferSubscribe(&t.CallOpts, fromBlock, toBlock)
}

/*
   get all event Transfer happened from [fromBlock] to [toBlock]
   if [toBlock] is -1, you can get all the events that will happen later.
   you can cancel event listenging through Subscription's Unsubscribe
*/
func (t *MyTokenCallerSession) EventTransferSubscribe(opts *bind.CallOpts, fromBlock rpc.BlockNumber,
	toBlock rpc.BlockNumber) (<-chan types.Log, ethereum.Subscription, error) {
	return t.Contract.EventTransferSubscribe(&t.CallOpts, fromBlock, toBlock)
}
