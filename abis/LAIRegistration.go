// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lairegistration

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

// LairegistrationMetaData contains all meta data concerning the Lairegistration contract.
var LairegistrationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tgId\",\"type\":\"uint256\"}],\"name\":\"RegistrationSuccessfulTG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"RegistrationSuccessfulUI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"apiKeys\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"apiKeysTG\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"apiKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userWallet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"changeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tgId\",\"type\":\"uint256\"}],\"name\":\"getTGApiKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getUIApiKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tgId\",\"type\":\"uint256\"}],\"name\":\"registration\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LairegistrationABI is the input ABI used to generate the binding from.
// Deprecated: Use LairegistrationMetaData.ABI instead.
var LairegistrationABI = LairegistrationMetaData.ABI

// Lairegistration is an auto generated Go binding around an Ethereum contract.
type Lairegistration struct {
	LairegistrationCaller     // Read-only binding to the contract
	LairegistrationTransactor // Write-only binding to the contract
	LairegistrationFilterer   // Log filterer for contract events
}

// LairegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type LairegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LairegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LairegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LairegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LairegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LairegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LairegistrationSession struct {
	Contract     *Lairegistration  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LairegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LairegistrationCallerSession struct {
	Contract *LairegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// LairegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LairegistrationTransactorSession struct {
	Contract     *LairegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// LairegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type LairegistrationRaw struct {
	Contract *Lairegistration // Generic contract binding to access the raw methods on
}

// LairegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LairegistrationCallerRaw struct {
	Contract *LairegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// LairegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LairegistrationTransactorRaw struct {
	Contract *LairegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLairegistration creates a new instance of Lairegistration, bound to a specific deployed contract.
func NewLairegistration(address common.Address, backend bind.ContractBackend) (*Lairegistration, error) {
	contract, err := bindLairegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lairegistration{LairegistrationCaller: LairegistrationCaller{contract: contract}, LairegistrationTransactor: LairegistrationTransactor{contract: contract}, LairegistrationFilterer: LairegistrationFilterer{contract: contract}}, nil
}

// NewLairegistrationCaller creates a new read-only instance of Lairegistration, bound to a specific deployed contract.
func NewLairegistrationCaller(address common.Address, caller bind.ContractCaller) (*LairegistrationCaller, error) {
	contract, err := bindLairegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LairegistrationCaller{contract: contract}, nil
}

// NewLairegistrationTransactor creates a new write-only instance of Lairegistration, bound to a specific deployed contract.
func NewLairegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*LairegistrationTransactor, error) {
	contract, err := bindLairegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LairegistrationTransactor{contract: contract}, nil
}

// NewLairegistrationFilterer creates a new log filterer instance of Lairegistration, bound to a specific deployed contract.
func NewLairegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*LairegistrationFilterer, error) {
	contract, err := bindLairegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LairegistrationFilterer{contract: contract}, nil
}

// bindLairegistration binds a generic wrapper to an already deployed contract.
func bindLairegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LairegistrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lairegistration *LairegistrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lairegistration.Contract.LairegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lairegistration *LairegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lairegistration.Contract.LairegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lairegistration *LairegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lairegistration.Contract.LairegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lairegistration *LairegistrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lairegistration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lairegistration *LairegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lairegistration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lairegistration *LairegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lairegistration.Contract.contract.Transact(opts, method, params...)
}

// ApiKeys is a free data retrieval call binding the contract method 0x498ced46.
//
// Solidity: function apiKeys(address ) view returns(string)
func (_Lairegistration *LairegistrationCaller) ApiKeys(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Lairegistration.contract.Call(opts, &out, "apiKeys", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiKeys is a free data retrieval call binding the contract method 0x498ced46.
//
// Solidity: function apiKeys(address ) view returns(string)
func (_Lairegistration *LairegistrationSession) ApiKeys(arg0 common.Address) (string, error) {
	return _Lairegistration.Contract.ApiKeys(&_Lairegistration.CallOpts, arg0)
}

// ApiKeys is a free data retrieval call binding the contract method 0x498ced46.
//
// Solidity: function apiKeys(address ) view returns(string)
func (_Lairegistration *LairegistrationCallerSession) ApiKeys(arg0 common.Address) (string, error) {
	return _Lairegistration.Contract.ApiKeys(&_Lairegistration.CallOpts, arg0)
}

// ApiKeysTG is a free data retrieval call binding the contract method 0x5fb1f907.
//
// Solidity: function apiKeysTG(uint256 ) view returns(string apiKey, address userWallet)
func (_Lairegistration *LairegistrationCaller) ApiKeysTG(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ApiKey     string
	UserWallet common.Address
}, error) {
	var out []interface{}
	err := _Lairegistration.contract.Call(opts, &out, "apiKeysTG", arg0)

	outstruct := new(struct {
		ApiKey     string
		UserWallet common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ApiKey = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.UserWallet = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ApiKeysTG is a free data retrieval call binding the contract method 0x5fb1f907.
//
// Solidity: function apiKeysTG(uint256 ) view returns(string apiKey, address userWallet)
func (_Lairegistration *LairegistrationSession) ApiKeysTG(arg0 *big.Int) (struct {
	ApiKey     string
	UserWallet common.Address
}, error) {
	return _Lairegistration.Contract.ApiKeysTG(&_Lairegistration.CallOpts, arg0)
}

// ApiKeysTG is a free data retrieval call binding the contract method 0x5fb1f907.
//
// Solidity: function apiKeysTG(uint256 ) view returns(string apiKey, address userWallet)
func (_Lairegistration *LairegistrationCallerSession) ApiKeysTG(arg0 *big.Int) (struct {
	ApiKey     string
	UserWallet common.Address
}, error) {
	return _Lairegistration.Contract.ApiKeysTG(&_Lairegistration.CallOpts, arg0)
}

// GetTGApiKey is a free data retrieval call binding the contract method 0x010b51ca.
//
// Solidity: function getTGApiKey(uint256 tgId) view returns(string, address)
func (_Lairegistration *LairegistrationCaller) GetTGApiKey(opts *bind.CallOpts, tgId *big.Int) (string, common.Address, error) {
	var out []interface{}
	err := _Lairegistration.contract.Call(opts, &out, "getTGApiKey", tgId)

	if err != nil {
		return *new(string), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetTGApiKey is a free data retrieval call binding the contract method 0x010b51ca.
//
// Solidity: function getTGApiKey(uint256 tgId) view returns(string, address)
func (_Lairegistration *LairegistrationSession) GetTGApiKey(tgId *big.Int) (string, common.Address, error) {
	return _Lairegistration.Contract.GetTGApiKey(&_Lairegistration.CallOpts, tgId)
}

// GetTGApiKey is a free data retrieval call binding the contract method 0x010b51ca.
//
// Solidity: function getTGApiKey(uint256 tgId) view returns(string, address)
func (_Lairegistration *LairegistrationCallerSession) GetTGApiKey(tgId *big.Int) (string, common.Address, error) {
	return _Lairegistration.Contract.GetTGApiKey(&_Lairegistration.CallOpts, tgId)
}

// GetUIApiKey is a free data retrieval call binding the contract method 0x61e9e810.
//
// Solidity: function getUIApiKey(address wallet) view returns(string)
func (_Lairegistration *LairegistrationCaller) GetUIApiKey(opts *bind.CallOpts, wallet common.Address) (string, error) {
	var out []interface{}
	err := _Lairegistration.contract.Call(opts, &out, "getUIApiKey", wallet)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUIApiKey is a free data retrieval call binding the contract method 0x61e9e810.
//
// Solidity: function getUIApiKey(address wallet) view returns(string)
func (_Lairegistration *LairegistrationSession) GetUIApiKey(wallet common.Address) (string, error) {
	return _Lairegistration.Contract.GetUIApiKey(&_Lairegistration.CallOpts, wallet)
}

// GetUIApiKey is a free data retrieval call binding the contract method 0x61e9e810.
//
// Solidity: function getUIApiKey(address wallet) view returns(string)
func (_Lairegistration *LairegistrationCallerSession) GetUIApiKey(wallet common.Address) (string, error) {
	return _Lairegistration.Contract.GetUIApiKey(&_Lairegistration.CallOpts, wallet)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lairegistration *LairegistrationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lairegistration.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lairegistration *LairegistrationSession) Owner() (common.Address, error) {
	return _Lairegistration.Contract.Owner(&_Lairegistration.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lairegistration *LairegistrationCallerSession) Owner() (common.Address, error) {
	return _Lairegistration.Contract.Owner(&_Lairegistration.CallOpts)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _newFee) returns()
func (_Lairegistration *LairegistrationTransactor) ChangeFee(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Lairegistration.contract.Transact(opts, "changeFee", _newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _newFee) returns()
func (_Lairegistration *LairegistrationSession) ChangeFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Lairegistration.Contract.ChangeFee(&_Lairegistration.TransactOpts, _newFee)
}

// ChangeFee is a paid mutator transaction binding the contract method 0x6a1db1bf.
//
// Solidity: function changeFee(uint256 _newFee) returns()
func (_Lairegistration *LairegistrationTransactorSession) ChangeFee(_newFee *big.Int) (*types.Transaction, error) {
	return _Lairegistration.Contract.ChangeFee(&_Lairegistration.TransactOpts, _newFee)
}

// Registration is a paid mutator transaction binding the contract method 0x8ee1d304.
//
// Solidity: function registration(string key, uint256 tgId) payable returns()
func (_Lairegistration *LairegistrationTransactor) Registration(opts *bind.TransactOpts, key string, tgId *big.Int) (*types.Transaction, error) {
	return _Lairegistration.contract.Transact(opts, "registration", key, tgId)
}

// Registration is a paid mutator transaction binding the contract method 0x8ee1d304.
//
// Solidity: function registration(string key, uint256 tgId) payable returns()
func (_Lairegistration *LairegistrationSession) Registration(key string, tgId *big.Int) (*types.Transaction, error) {
	return _Lairegistration.Contract.Registration(&_Lairegistration.TransactOpts, key, tgId)
}

// Registration is a paid mutator transaction binding the contract method 0x8ee1d304.
//
// Solidity: function registration(string key, uint256 tgId) payable returns()
func (_Lairegistration *LairegistrationTransactorSession) Registration(key string, tgId *big.Int) (*types.Transaction, error) {
	return _Lairegistration.Contract.Registration(&_Lairegistration.TransactOpts, key, tgId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lairegistration *LairegistrationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lairegistration.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lairegistration *LairegistrationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lairegistration.Contract.RenounceOwnership(&_Lairegistration.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lairegistration *LairegistrationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lairegistration.Contract.RenounceOwnership(&_Lairegistration.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lairegistration *LairegistrationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lairegistration.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lairegistration *LairegistrationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lairegistration.Contract.TransferOwnership(&_Lairegistration.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lairegistration *LairegistrationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lairegistration.Contract.TransferOwnership(&_Lairegistration.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lairegistration *LairegistrationTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lairegistration.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lairegistration *LairegistrationSession) Withdraw() (*types.Transaction, error) {
	return _Lairegistration.Contract.Withdraw(&_Lairegistration.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Lairegistration *LairegistrationTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Lairegistration.Contract.Withdraw(&_Lairegistration.TransactOpts)
}

// LairegistrationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lairegistration contract.
type LairegistrationOwnershipTransferredIterator struct {
	Event *LairegistrationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LairegistrationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LairegistrationOwnershipTransferred)
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
		it.Event = new(LairegistrationOwnershipTransferred)
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
func (it *LairegistrationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LairegistrationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LairegistrationOwnershipTransferred represents a OwnershipTransferred event raised by the Lairegistration contract.
type LairegistrationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lairegistration *LairegistrationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LairegistrationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lairegistration.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LairegistrationOwnershipTransferredIterator{contract: _Lairegistration.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lairegistration *LairegistrationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LairegistrationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lairegistration.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LairegistrationOwnershipTransferred)
				if err := _Lairegistration.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Lairegistration *LairegistrationFilterer) ParseOwnershipTransferred(log types.Log) (*LairegistrationOwnershipTransferred, error) {
	event := new(LairegistrationOwnershipTransferred)
	if err := _Lairegistration.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LairegistrationRegistrationSuccessfulTGIterator is returned from FilterRegistrationSuccessfulTG and is used to iterate over the raw logs and unpacked data for RegistrationSuccessfulTG events raised by the Lairegistration contract.
type LairegistrationRegistrationSuccessfulTGIterator struct {
	Event *LairegistrationRegistrationSuccessfulTG // Event containing the contract specifics and raw log

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
func (it *LairegistrationRegistrationSuccessfulTGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LairegistrationRegistrationSuccessfulTG)
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
		it.Event = new(LairegistrationRegistrationSuccessfulTG)
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
func (it *LairegistrationRegistrationSuccessfulTGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LairegistrationRegistrationSuccessfulTGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LairegistrationRegistrationSuccessfulTG represents a RegistrationSuccessfulTG event raised by the Lairegistration contract.
type LairegistrationRegistrationSuccessfulTG struct {
	Wallet common.Address
	Key    string
	TgId   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegistrationSuccessfulTG is a free log retrieval operation binding the contract event 0x92105da43b232c654d7cb657b1a6abc36f473f5e82eb1214a6b59f7aa73f213a.
//
// Solidity: event RegistrationSuccessfulTG(address indexed wallet, string key, uint256 tgId)
func (_Lairegistration *LairegistrationFilterer) FilterRegistrationSuccessfulTG(opts *bind.FilterOpts, wallet []common.Address) (*LairegistrationRegistrationSuccessfulTGIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Lairegistration.contract.FilterLogs(opts, "RegistrationSuccessfulTG", walletRule)
	if err != nil {
		return nil, err
	}
	return &LairegistrationRegistrationSuccessfulTGIterator{contract: _Lairegistration.contract, event: "RegistrationSuccessfulTG", logs: logs, sub: sub}, nil
}

// WatchRegistrationSuccessfulTG is a free log subscription operation binding the contract event 0x92105da43b232c654d7cb657b1a6abc36f473f5e82eb1214a6b59f7aa73f213a.
//
// Solidity: event RegistrationSuccessfulTG(address indexed wallet, string key, uint256 tgId)
func (_Lairegistration *LairegistrationFilterer) WatchRegistrationSuccessfulTG(opts *bind.WatchOpts, sink chan<- *LairegistrationRegistrationSuccessfulTG, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Lairegistration.contract.WatchLogs(opts, "RegistrationSuccessfulTG", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LairegistrationRegistrationSuccessfulTG)
				if err := _Lairegistration.contract.UnpackLog(event, "RegistrationSuccessfulTG", log); err != nil {
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

// ParseRegistrationSuccessfulTG is a log parse operation binding the contract event 0x92105da43b232c654d7cb657b1a6abc36f473f5e82eb1214a6b59f7aa73f213a.
//
// Solidity: event RegistrationSuccessfulTG(address indexed wallet, string key, uint256 tgId)
func (_Lairegistration *LairegistrationFilterer) ParseRegistrationSuccessfulTG(log types.Log) (*LairegistrationRegistrationSuccessfulTG, error) {
	event := new(LairegistrationRegistrationSuccessfulTG)
	if err := _Lairegistration.contract.UnpackLog(event, "RegistrationSuccessfulTG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LairegistrationRegistrationSuccessfulUIIterator is returned from FilterRegistrationSuccessfulUI and is used to iterate over the raw logs and unpacked data for RegistrationSuccessfulUI events raised by the Lairegistration contract.
type LairegistrationRegistrationSuccessfulUIIterator struct {
	Event *LairegistrationRegistrationSuccessfulUI // Event containing the contract specifics and raw log

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
func (it *LairegistrationRegistrationSuccessfulUIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LairegistrationRegistrationSuccessfulUI)
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
		it.Event = new(LairegistrationRegistrationSuccessfulUI)
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
func (it *LairegistrationRegistrationSuccessfulUIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LairegistrationRegistrationSuccessfulUIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LairegistrationRegistrationSuccessfulUI represents a RegistrationSuccessfulUI event raised by the Lairegistration contract.
type LairegistrationRegistrationSuccessfulUI struct {
	Wallet common.Address
	Key    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRegistrationSuccessfulUI is a free log retrieval operation binding the contract event 0xcc1ad402c343c58809d16c5a46d95a1fe7a4dfae5dad2e3d15dbbf4afb6c8abb.
//
// Solidity: event RegistrationSuccessfulUI(address indexed wallet, string key)
func (_Lairegistration *LairegistrationFilterer) FilterRegistrationSuccessfulUI(opts *bind.FilterOpts, wallet []common.Address) (*LairegistrationRegistrationSuccessfulUIIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Lairegistration.contract.FilterLogs(opts, "RegistrationSuccessfulUI", walletRule)
	if err != nil {
		return nil, err
	}
	return &LairegistrationRegistrationSuccessfulUIIterator{contract: _Lairegistration.contract, event: "RegistrationSuccessfulUI", logs: logs, sub: sub}, nil
}

// WatchRegistrationSuccessfulUI is a free log subscription operation binding the contract event 0xcc1ad402c343c58809d16c5a46d95a1fe7a4dfae5dad2e3d15dbbf4afb6c8abb.
//
// Solidity: event RegistrationSuccessfulUI(address indexed wallet, string key)
func (_Lairegistration *LairegistrationFilterer) WatchRegistrationSuccessfulUI(opts *bind.WatchOpts, sink chan<- *LairegistrationRegistrationSuccessfulUI, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _Lairegistration.contract.WatchLogs(opts, "RegistrationSuccessfulUI", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LairegistrationRegistrationSuccessfulUI)
				if err := _Lairegistration.contract.UnpackLog(event, "RegistrationSuccessfulUI", log); err != nil {
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

// ParseRegistrationSuccessfulUI is a log parse operation binding the contract event 0xcc1ad402c343c58809d16c5a46d95a1fe7a4dfae5dad2e3d15dbbf4afb6c8abb.
//
// Solidity: event RegistrationSuccessfulUI(address indexed wallet, string key)
func (_Lairegistration *LairegistrationFilterer) ParseRegistrationSuccessfulUI(log types.Log) (*LairegistrationRegistrationSuccessfulUI, error) {
	event := new(LairegistrationRegistrationSuccessfulUI)
	if err := _Lairegistration.contract.UnpackLog(event, "RegistrationSuccessfulUI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
