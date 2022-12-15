// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// SpyNFTMetaData contains all meta data concerning the SpyNFT contract.
var SpyNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TooEarly\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGooBalance\",\"type\":\"uint256\"}],\"name\":\"GooBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EMISSION_MULTIPLE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getUserData\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"lastBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"lastTimestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"gooBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"sudoTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gooAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumGooBalanceUpdateType\",\"name\":\"updateType\",\"type\":\"uint8\"}],\"name\":\"updateUserGooBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SpyNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use SpyNFTMetaData.ABI instead.
var SpyNFTABI = SpyNFTMetaData.ABI

// SpyNFT is an auto generated Go binding around an Ethereum contract.
type SpyNFT struct {
	SpyNFTCaller     // Read-only binding to the contract
	SpyNFTTransactor // Write-only binding to the contract
	SpyNFTFilterer   // Log filterer for contract events
}

// SpyNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpyNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpyNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpyNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpyNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpyNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpyNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpyNFTSession struct {
	Contract     *SpyNFT           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpyNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpyNFTCallerSession struct {
	Contract *SpyNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SpyNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpyNFTTransactorSession struct {
	Contract     *SpyNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpyNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpyNFTRaw struct {
	Contract *SpyNFT // Generic contract binding to access the raw methods on
}

// SpyNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpyNFTCallerRaw struct {
	Contract *SpyNFTCaller // Generic read-only contract binding to access the raw methods on
}

// SpyNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpyNFTTransactorRaw struct {
	Contract *SpyNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpyNFT creates a new instance of SpyNFT, bound to a specific deployed contract.
func NewSpyNFT(address common.Address, backend bind.ContractBackend) (*SpyNFT, error) {
	contract, err := bindSpyNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SpyNFT{SpyNFTCaller: SpyNFTCaller{contract: contract}, SpyNFTTransactor: SpyNFTTransactor{contract: contract}, SpyNFTFilterer: SpyNFTFilterer{contract: contract}}, nil
}

// NewSpyNFTCaller creates a new read-only instance of SpyNFT, bound to a specific deployed contract.
func NewSpyNFTCaller(address common.Address, caller bind.ContractCaller) (*SpyNFTCaller, error) {
	contract, err := bindSpyNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpyNFTCaller{contract: contract}, nil
}

// NewSpyNFTTransactor creates a new write-only instance of SpyNFT, bound to a specific deployed contract.
func NewSpyNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*SpyNFTTransactor, error) {
	contract, err := bindSpyNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpyNFTTransactor{contract: contract}, nil
}

// NewSpyNFTFilterer creates a new log filterer instance of SpyNFT, bound to a specific deployed contract.
func NewSpyNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*SpyNFTFilterer, error) {
	contract, err := bindSpyNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpyNFTFilterer{contract: contract}, nil
}

// bindSpyNFT binds a generic wrapper to an already deployed contract.
func bindSpyNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SpyNFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpyNFT *SpyNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpyNFT.Contract.SpyNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpyNFT *SpyNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpyNFT.Contract.SpyNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpyNFT *SpyNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpyNFT.Contract.SpyNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpyNFT *SpyNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpyNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpyNFT *SpyNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpyNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpyNFT *SpyNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpyNFT.Contract.contract.Transact(opts, method, params...)
}

// EMISSIONMULTIPLE is a free data retrieval call binding the contract method 0xcb1a29b5.
//
// Solidity: function EMISSION_MULTIPLE() view returns(uint256)
func (_SpyNFT *SpyNFTCaller) EMISSIONMULTIPLE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "EMISSION_MULTIPLE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EMISSIONMULTIPLE is a free data retrieval call binding the contract method 0xcb1a29b5.
//
// Solidity: function EMISSION_MULTIPLE() view returns(uint256)
func (_SpyNFT *SpyNFTSession) EMISSIONMULTIPLE() (*big.Int, error) {
	return _SpyNFT.Contract.EMISSIONMULTIPLE(&_SpyNFT.CallOpts)
}

// EMISSIONMULTIPLE is a free data retrieval call binding the contract method 0xcb1a29b5.
//
// Solidity: function EMISSION_MULTIPLE() view returns(uint256)
func (_SpyNFT *SpyNFTCallerSession) EMISSIONMULTIPLE() (*big.Int, error) {
	return _SpyNFT.Contract.EMISSIONMULTIPLE(&_SpyNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SpyNFT *SpyNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SpyNFT *SpyNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SpyNFT.Contract.BalanceOf(&_SpyNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SpyNFT *SpyNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SpyNFT.Contract.BalanceOf(&_SpyNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SpyNFT *SpyNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SpyNFT *SpyNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SpyNFT.Contract.GetApproved(&_SpyNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SpyNFT *SpyNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SpyNFT.Contract.GetApproved(&_SpyNFT.CallOpts, tokenId)
}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address ) view returns(uint128 lastBalance, uint64 lastTimestamp)
func (_SpyNFT *SpyNFTCaller) GetUserData(opts *bind.CallOpts, arg0 common.Address) (struct {
	LastBalance   *big.Int
	LastTimestamp uint64
}, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "getUserData", arg0)

	outstruct := new(struct {
		LastBalance   *big.Int
		LastTimestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastTimestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address ) view returns(uint128 lastBalance, uint64 lastTimestamp)
func (_SpyNFT *SpyNFTSession) GetUserData(arg0 common.Address) (struct {
	LastBalance   *big.Int
	LastTimestamp uint64
}, error) {
	return _SpyNFT.Contract.GetUserData(&_SpyNFT.CallOpts, arg0)
}

// GetUserData is a free data retrieval call binding the contract method 0xffc9896b.
//
// Solidity: function getUserData(address ) view returns(uint128 lastBalance, uint64 lastTimestamp)
func (_SpyNFT *SpyNFTCallerSession) GetUserData(arg0 common.Address) (struct {
	LastBalance   *big.Int
	LastTimestamp uint64
}, error) {
	return _SpyNFT.Contract.GetUserData(&_SpyNFT.CallOpts, arg0)
}

// GooBalance is a free data retrieval call binding the contract method 0xd075fbba.
//
// Solidity: function gooBalance(address user) view returns(uint256)
func (_SpyNFT *SpyNFTCaller) GooBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "gooBalance", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GooBalance is a free data retrieval call binding the contract method 0xd075fbba.
//
// Solidity: function gooBalance(address user) view returns(uint256)
func (_SpyNFT *SpyNFTSession) GooBalance(user common.Address) (*big.Int, error) {
	return _SpyNFT.Contract.GooBalance(&_SpyNFT.CallOpts, user)
}

// GooBalance is a free data retrieval call binding the contract method 0xd075fbba.
//
// Solidity: function gooBalance(address user) view returns(uint256)
func (_SpyNFT *SpyNFTCallerSession) GooBalance(user common.Address) (*big.Int, error) {
	return _SpyNFT.Contract.GooBalance(&_SpyNFT.CallOpts, user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SpyNFT *SpyNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SpyNFT *SpyNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SpyNFT.Contract.IsApprovedForAll(&_SpyNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SpyNFT *SpyNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SpyNFT.Contract.IsApprovedForAll(&_SpyNFT.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SpyNFT *SpyNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SpyNFT *SpyNFTSession) Name() (string, error) {
	return _SpyNFT.Contract.Name(&_SpyNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SpyNFT *SpyNFTCallerSession) Name() (string, error) {
	return _SpyNFT.Contract.Name(&_SpyNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpyNFT *SpyNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpyNFT *SpyNFTSession) Owner() (common.Address, error) {
	return _SpyNFT.Contract.Owner(&_SpyNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpyNFT *SpyNFTCallerSession) Owner() (common.Address, error) {
	return _SpyNFT.Contract.Owner(&_SpyNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SpyNFT *SpyNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SpyNFT *SpyNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SpyNFT.Contract.OwnerOf(&_SpyNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SpyNFT *SpyNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SpyNFT.Contract.OwnerOf(&_SpyNFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SpyNFT *SpyNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SpyNFT *SpyNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SpyNFT.Contract.SupportsInterface(&_SpyNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SpyNFT *SpyNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SpyNFT.Contract.SupportsInterface(&_SpyNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SpyNFT *SpyNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SpyNFT *SpyNFTSession) Symbol() (string, error) {
	return _SpyNFT.Contract.Symbol(&_SpyNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SpyNFT *SpyNFTCallerSession) Symbol() (string, error) {
	return _SpyNFT.Contract.Symbol(&_SpyNFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SpyNFT *SpyNFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SpyNFT *SpyNFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SpyNFT.Contract.TokenByIndex(&_SpyNFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_SpyNFT *SpyNFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _SpyNFT.Contract.TokenByIndex(&_SpyNFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SpyNFT *SpyNFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SpyNFT *SpyNFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SpyNFT.Contract.TokenOfOwnerByIndex(&_SpyNFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_SpyNFT *SpyNFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _SpyNFT.Contract.TokenOfOwnerByIndex(&_SpyNFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) pure returns(string)
func (_SpyNFT *SpyNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) pure returns(string)
func (_SpyNFT *SpyNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SpyNFT.Contract.TokenURI(&_SpyNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) pure returns(string)
func (_SpyNFT *SpyNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SpyNFT.Contract.TokenURI(&_SpyNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SpyNFT *SpyNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SpyNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SpyNFT *SpyNFTSession) TotalSupply() (*big.Int, error) {
	return _SpyNFT.Contract.TotalSupply(&_SpyNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SpyNFT *SpyNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _SpyNFT.Contract.TotalSupply(&_SpyNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.Approve(&_SpyNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.Approve(&_SpyNFT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_SpyNFT *SpyNFTTransactor) Burn(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "burn", id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_SpyNFT *SpyNFTSession) Burn(id *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.Burn(&_SpyNFT.TransactOpts, id)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 id) returns()
func (_SpyNFT *SpyNFTTransactorSession) Burn(id *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.Burn(&_SpyNFT.TransactOpts, id)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 timestamp) returns(uint256 id)
func (_SpyNFT *SpyNFTTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "mint", recipient, timestamp)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 timestamp) returns(uint256 id)
func (_SpyNFT *SpyNFTSession) Mint(recipient common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.Mint(&_SpyNFT.TransactOpts, recipient, timestamp)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 timestamp) returns(uint256 id)
func (_SpyNFT *SpyNFTTransactorSession) Mint(recipient common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.Mint(&_SpyNFT.TransactOpts, recipient, timestamp)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address recipient) returns(uint256 id)
func (_SpyNFT *SpyNFTTransactor) Mint0(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "mint0", recipient)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address recipient) returns(uint256 id)
func (_SpyNFT *SpyNFTSession) Mint0(recipient common.Address) (*types.Transaction, error) {
	return _SpyNFT.Contract.Mint0(&_SpyNFT.TransactOpts, recipient)
}

// Mint0 is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address recipient) returns(uint256 id)
func (_SpyNFT *SpyNFTTransactorSession) Mint0(recipient common.Address) (*types.Transaction, error) {
	return _SpyNFT.Contract.Mint0(&_SpyNFT.TransactOpts, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SpyNFT *SpyNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SpyNFT *SpyNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _SpyNFT.Contract.RenounceOwnership(&_SpyNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SpyNFT *SpyNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SpyNFT.Contract.RenounceOwnership(&_SpyNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.SafeTransferFrom(&_SpyNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.SafeTransferFrom(&_SpyNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SpyNFT *SpyNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SpyNFT *SpyNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SpyNFT.Contract.SafeTransferFrom0(&_SpyNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SpyNFT *SpyNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SpyNFT.Contract.SafeTransferFrom0(&_SpyNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SpyNFT *SpyNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SpyNFT *SpyNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SpyNFT.Contract.SetApprovalForAll(&_SpyNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SpyNFT *SpyNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SpyNFT.Contract.SetApprovalForAll(&_SpyNFT.TransactOpts, operator, approved)
}

// SudoTransferFrom is a paid mutator transaction binding the contract method 0x20dd4c02.
//
// Solidity: function sudoTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTTransactor) SudoTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "sudoTransferFrom", from, to, tokenId)
}

// SudoTransferFrom is a paid mutator transaction binding the contract method 0x20dd4c02.
//
// Solidity: function sudoTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTSession) SudoTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.SudoTransferFrom(&_SpyNFT.TransactOpts, from, to, tokenId)
}

// SudoTransferFrom is a paid mutator transaction binding the contract method 0x20dd4c02.
//
// Solidity: function sudoTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTTransactorSession) SudoTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.SudoTransferFrom(&_SpyNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.TransferFrom(&_SpyNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SpyNFT *SpyNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpyNFT.Contract.TransferFrom(&_SpyNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SpyNFT *SpyNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SpyNFT *SpyNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SpyNFT.Contract.TransferOwnership(&_SpyNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SpyNFT *SpyNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SpyNFT.Contract.TransferOwnership(&_SpyNFT.TransactOpts, newOwner)
}

// UpdateUserGooBalance is a paid mutator transaction binding the contract method 0xf66e9ac3.
//
// Solidity: function updateUserGooBalance(address user, uint256 gooAmount, uint8 updateType) returns()
func (_SpyNFT *SpyNFTTransactor) UpdateUserGooBalance(opts *bind.TransactOpts, user common.Address, gooAmount *big.Int, updateType uint8) (*types.Transaction, error) {
	return _SpyNFT.contract.Transact(opts, "updateUserGooBalance", user, gooAmount, updateType)
}

// UpdateUserGooBalance is a paid mutator transaction binding the contract method 0xf66e9ac3.
//
// Solidity: function updateUserGooBalance(address user, uint256 gooAmount, uint8 updateType) returns()
func (_SpyNFT *SpyNFTSession) UpdateUserGooBalance(user common.Address, gooAmount *big.Int, updateType uint8) (*types.Transaction, error) {
	return _SpyNFT.Contract.UpdateUserGooBalance(&_SpyNFT.TransactOpts, user, gooAmount, updateType)
}

// UpdateUserGooBalance is a paid mutator transaction binding the contract method 0xf66e9ac3.
//
// Solidity: function updateUserGooBalance(address user, uint256 gooAmount, uint8 updateType) returns()
func (_SpyNFT *SpyNFTTransactorSession) UpdateUserGooBalance(user common.Address, gooAmount *big.Int, updateType uint8) (*types.Transaction, error) {
	return _SpyNFT.Contract.UpdateUserGooBalance(&_SpyNFT.TransactOpts, user, gooAmount, updateType)
}

// SpyNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SpyNFT contract.
type SpyNFTApprovalIterator struct {
	Event *SpyNFTApproval // Event containing the contract specifics and raw log

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
func (it *SpyNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpyNFTApproval)
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
		it.Event = new(SpyNFTApproval)
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
func (it *SpyNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpyNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpyNFTApproval represents a Approval event raised by the SpyNFT contract.
type SpyNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SpyNFT *SpyNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SpyNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SpyNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SpyNFTApprovalIterator{contract: _SpyNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SpyNFT *SpyNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SpyNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SpyNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpyNFTApproval)
				if err := _SpyNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SpyNFT *SpyNFTFilterer) ParseApproval(log types.Log) (*SpyNFTApproval, error) {
	event := new(SpyNFTApproval)
	if err := _SpyNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpyNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SpyNFT contract.
type SpyNFTApprovalForAllIterator struct {
	Event *SpyNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SpyNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpyNFTApprovalForAll)
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
		it.Event = new(SpyNFTApprovalForAll)
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
func (it *SpyNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpyNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpyNFTApprovalForAll represents a ApprovalForAll event raised by the SpyNFT contract.
type SpyNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SpyNFT *SpyNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SpyNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SpyNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SpyNFTApprovalForAllIterator{contract: _SpyNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SpyNFT *SpyNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SpyNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SpyNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpyNFTApprovalForAll)
				if err := _SpyNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SpyNFT *SpyNFTFilterer) ParseApprovalForAll(log types.Log) (*SpyNFTApprovalForAll, error) {
	event := new(SpyNFTApprovalForAll)
	if err := _SpyNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpyNFTGooBalanceUpdatedIterator is returned from FilterGooBalanceUpdated and is used to iterate over the raw logs and unpacked data for GooBalanceUpdated events raised by the SpyNFT contract.
type SpyNFTGooBalanceUpdatedIterator struct {
	Event *SpyNFTGooBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *SpyNFTGooBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpyNFTGooBalanceUpdated)
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
		it.Event = new(SpyNFTGooBalanceUpdated)
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
func (it *SpyNFTGooBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpyNFTGooBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpyNFTGooBalanceUpdated represents a GooBalanceUpdated event raised by the SpyNFT contract.
type SpyNFTGooBalanceUpdated struct {
	User          common.Address
	NewGooBalance *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGooBalanceUpdated is a free log retrieval operation binding the contract event 0x2171f1d8b6b7927c42287fd11040aa8c3569c5c2040b8453104ced365ed84cd4.
//
// Solidity: event GooBalanceUpdated(address indexed user, uint256 newGooBalance)
func (_SpyNFT *SpyNFTFilterer) FilterGooBalanceUpdated(opts *bind.FilterOpts, user []common.Address) (*SpyNFTGooBalanceUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SpyNFT.contract.FilterLogs(opts, "GooBalanceUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &SpyNFTGooBalanceUpdatedIterator{contract: _SpyNFT.contract, event: "GooBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchGooBalanceUpdated is a free log subscription operation binding the contract event 0x2171f1d8b6b7927c42287fd11040aa8c3569c5c2040b8453104ced365ed84cd4.
//
// Solidity: event GooBalanceUpdated(address indexed user, uint256 newGooBalance)
func (_SpyNFT *SpyNFTFilterer) WatchGooBalanceUpdated(opts *bind.WatchOpts, sink chan<- *SpyNFTGooBalanceUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SpyNFT.contract.WatchLogs(opts, "GooBalanceUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpyNFTGooBalanceUpdated)
				if err := _SpyNFT.contract.UnpackLog(event, "GooBalanceUpdated", log); err != nil {
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

// ParseGooBalanceUpdated is a log parse operation binding the contract event 0x2171f1d8b6b7927c42287fd11040aa8c3569c5c2040b8453104ced365ed84cd4.
//
// Solidity: event GooBalanceUpdated(address indexed user, uint256 newGooBalance)
func (_SpyNFT *SpyNFTFilterer) ParseGooBalanceUpdated(log types.Log) (*SpyNFTGooBalanceUpdated, error) {
	event := new(SpyNFTGooBalanceUpdated)
	if err := _SpyNFT.contract.UnpackLog(event, "GooBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpyNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SpyNFT contract.
type SpyNFTOwnershipTransferredIterator struct {
	Event *SpyNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SpyNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpyNFTOwnershipTransferred)
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
		it.Event = new(SpyNFTOwnershipTransferred)
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
func (it *SpyNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpyNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpyNFTOwnershipTransferred represents a OwnershipTransferred event raised by the SpyNFT contract.
type SpyNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SpyNFT *SpyNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SpyNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SpyNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SpyNFTOwnershipTransferredIterator{contract: _SpyNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SpyNFT *SpyNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SpyNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SpyNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpyNFTOwnershipTransferred)
				if err := _SpyNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SpyNFT *SpyNFTFilterer) ParseOwnershipTransferred(log types.Log) (*SpyNFTOwnershipTransferred, error) {
	event := new(SpyNFTOwnershipTransferred)
	if err := _SpyNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpyNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SpyNFT contract.
type SpyNFTTransferIterator struct {
	Event *SpyNFTTransfer // Event containing the contract specifics and raw log

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
func (it *SpyNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpyNFTTransfer)
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
		it.Event = new(SpyNFTTransfer)
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
func (it *SpyNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpyNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpyNFTTransfer represents a Transfer event raised by the SpyNFT contract.
type SpyNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SpyNFT *SpyNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SpyNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SpyNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SpyNFTTransferIterator{contract: _SpyNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SpyNFT *SpyNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SpyNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SpyNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpyNFTTransfer)
				if err := _SpyNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SpyNFT *SpyNFTFilterer) ParseTransfer(log types.Log) (*SpyNFTTransfer, error) {
	event := new(SpyNFTTransfer)
	if err := _SpyNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
