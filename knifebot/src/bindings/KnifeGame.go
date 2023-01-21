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

// KnifeGameMetaData contains all meta data concerning the KnifeGame contract.
var KnifeGameMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gameStart\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_spyNft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_knifeNFT\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DumbMove\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoWhales\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooPoor\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"KnifePurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Shouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hitman\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"victim\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"knifeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spyId\",\"type\":\"uint256\"}],\"name\":\"SpyKilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"SpyPurchasedETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SpyPurchasedMoo\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_PURCHASE_SPY_ETH_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MULTISIG\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHOUTS_FUNDS_RECIPIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFreeMoo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasUserClaimedFreeMooTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasUserPrepurchased\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReimbursed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_knifeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_spyId\",\"type\":\"uint256\"}],\"name\":\"killSpy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"knifeLVRGDA\",\"outputs\":[{\"internalType\":\"contractLVRGDA\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"knifeNFT\",\"outputs\":[{\"internalType\":\"contractKnifeNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"knifePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"knivesMintedFromMoo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPrice\",\"type\":\"uint256\"}],\"name\":\"mintKnifeFromMoolah\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"knifeId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPrice\",\"type\":\"uint256\"}],\"name\":\"mintSpyFromMoolah\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spyId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"purchaseSpy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spyId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_reimburse\",\"type\":\"address[]\"}],\"name\":\"reimburseSpy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"shout\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spiesMintedFromMoo\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spyLVRGDA\",\"outputs\":[{\"internalType\":\"contractLVRGDA\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spyNFT\",\"outputs\":[{\"internalType\":\"contractSpyNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spyPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"spyPriceETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_targetPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_priceDecayPercent\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_timeScale\",\"type\":\"int256\"}],\"name\":\"updateKnifeVRGDA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_targetPrice\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_priceDecayPercent\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_timeScale\",\"type\":\"int256\"}],\"name\":\"updateSpyVRGDA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userPurchasesOnDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KnifeGameABI is the input ABI used to generate the binding from.
// Deprecated: Use KnifeGameMetaData.ABI instead.
var KnifeGameABI = KnifeGameMetaData.ABI

// KnifeGame is an auto generated Go binding around an Ethereum contract.
type KnifeGame struct {
	KnifeGameCaller     // Read-only binding to the contract
	KnifeGameTransactor // Write-only binding to the contract
	KnifeGameFilterer   // Log filterer for contract events
}

// KnifeGameCaller is an auto generated read-only Go binding around an Ethereum contract.
type KnifeGameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KnifeGameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KnifeGameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KnifeGameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KnifeGameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KnifeGameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KnifeGameSession struct {
	Contract     *KnifeGame        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KnifeGameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KnifeGameCallerSession struct {
	Contract *KnifeGameCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KnifeGameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KnifeGameTransactorSession struct {
	Contract     *KnifeGameTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KnifeGameRaw is an auto generated low-level Go binding around an Ethereum contract.
type KnifeGameRaw struct {
	Contract *KnifeGame // Generic contract binding to access the raw methods on
}

// KnifeGameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KnifeGameCallerRaw struct {
	Contract *KnifeGameCaller // Generic read-only contract binding to access the raw methods on
}

// KnifeGameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KnifeGameTransactorRaw struct {
	Contract *KnifeGameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKnifeGame creates a new instance of KnifeGame, bound to a specific deployed contract.
func NewKnifeGame(address common.Address, backend bind.ContractBackend) (*KnifeGame, error) {
	contract, err := bindKnifeGame(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KnifeGame{KnifeGameCaller: KnifeGameCaller{contract: contract}, KnifeGameTransactor: KnifeGameTransactor{contract: contract}, KnifeGameFilterer: KnifeGameFilterer{contract: contract}}, nil
}

// NewKnifeGameCaller creates a new read-only instance of KnifeGame, bound to a specific deployed contract.
func NewKnifeGameCaller(address common.Address, caller bind.ContractCaller) (*KnifeGameCaller, error) {
	contract, err := bindKnifeGame(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KnifeGameCaller{contract: contract}, nil
}

// NewKnifeGameTransactor creates a new write-only instance of KnifeGame, bound to a specific deployed contract.
func NewKnifeGameTransactor(address common.Address, transactor bind.ContractTransactor) (*KnifeGameTransactor, error) {
	contract, err := bindKnifeGame(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KnifeGameTransactor{contract: contract}, nil
}

// NewKnifeGameFilterer creates a new log filterer instance of KnifeGame, bound to a specific deployed contract.
func NewKnifeGameFilterer(address common.Address, filterer bind.ContractFilterer) (*KnifeGameFilterer, error) {
	contract, err := bindKnifeGame(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KnifeGameFilterer{contract: contract}, nil
}

// bindKnifeGame binds a generic wrapper to an already deployed contract.
func bindKnifeGame(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KnifeGameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KnifeGame *KnifeGameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KnifeGame.Contract.KnifeGameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KnifeGame *KnifeGameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KnifeGame.Contract.KnifeGameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KnifeGame *KnifeGameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KnifeGame.Contract.KnifeGameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KnifeGame *KnifeGameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KnifeGame.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KnifeGame *KnifeGameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KnifeGame.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KnifeGame *KnifeGameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KnifeGame.Contract.contract.Transact(opts, method, params...)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_KnifeGame *KnifeGameCaller) BURNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "BURN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_KnifeGame *KnifeGameSession) BURNADDRESS() (common.Address, error) {
	return _KnifeGame.Contract.BURNADDRESS(&_KnifeGame.CallOpts)
}

// BURNADDRESS is a free data retrieval call binding the contract method 0xfccc2813.
//
// Solidity: function BURN_ADDRESS() view returns(address)
func (_KnifeGame *KnifeGameCallerSession) BURNADDRESS() (common.Address, error) {
	return _KnifeGame.Contract.BURNADDRESS(&_KnifeGame.CallOpts)
}

// INITIALPURCHASESPYETHPRICE is a free data retrieval call binding the contract method 0xdd416943.
//
// Solidity: function INITIAL_PURCHASE_SPY_ETH_PRICE() view returns(uint256)
func (_KnifeGame *KnifeGameCaller) INITIALPURCHASESPYETHPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "INITIAL_PURCHASE_SPY_ETH_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALPURCHASESPYETHPRICE is a free data retrieval call binding the contract method 0xdd416943.
//
// Solidity: function INITIAL_PURCHASE_SPY_ETH_PRICE() view returns(uint256)
func (_KnifeGame *KnifeGameSession) INITIALPURCHASESPYETHPRICE() (*big.Int, error) {
	return _KnifeGame.Contract.INITIALPURCHASESPYETHPRICE(&_KnifeGame.CallOpts)
}

// INITIALPURCHASESPYETHPRICE is a free data retrieval call binding the contract method 0xdd416943.
//
// Solidity: function INITIAL_PURCHASE_SPY_ETH_PRICE() view returns(uint256)
func (_KnifeGame *KnifeGameCallerSession) INITIALPURCHASESPYETHPRICE() (*big.Int, error) {
	return _KnifeGame.Contract.INITIALPURCHASESPYETHPRICE(&_KnifeGame.CallOpts)
}

// MULTISIG is a free data retrieval call binding the contract method 0x2530b145.
//
// Solidity: function MULTISIG() view returns(address)
func (_KnifeGame *KnifeGameCaller) MULTISIG(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "MULTISIG")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MULTISIG is a free data retrieval call binding the contract method 0x2530b145.
//
// Solidity: function MULTISIG() view returns(address)
func (_KnifeGame *KnifeGameSession) MULTISIG() (common.Address, error) {
	return _KnifeGame.Contract.MULTISIG(&_KnifeGame.CallOpts)
}

// MULTISIG is a free data retrieval call binding the contract method 0x2530b145.
//
// Solidity: function MULTISIG() view returns(address)
func (_KnifeGame *KnifeGameCallerSession) MULTISIG() (common.Address, error) {
	return _KnifeGame.Contract.MULTISIG(&_KnifeGame.CallOpts)
}

// SHOUTSFUNDSRECIPIENT is a free data retrieval call binding the contract method 0xd5ed95c1.
//
// Solidity: function SHOUTS_FUNDS_RECIPIENT() view returns(address)
func (_KnifeGame *KnifeGameCaller) SHOUTSFUNDSRECIPIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "SHOUTS_FUNDS_RECIPIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SHOUTSFUNDSRECIPIENT is a free data retrieval call binding the contract method 0xd5ed95c1.
//
// Solidity: function SHOUTS_FUNDS_RECIPIENT() view returns(address)
func (_KnifeGame *KnifeGameSession) SHOUTSFUNDSRECIPIENT() (common.Address, error) {
	return _KnifeGame.Contract.SHOUTSFUNDSRECIPIENT(&_KnifeGame.CallOpts)
}

// SHOUTSFUNDSRECIPIENT is a free data retrieval call binding the contract method 0xd5ed95c1.
//
// Solidity: function SHOUTS_FUNDS_RECIPIENT() view returns(address)
func (_KnifeGame *KnifeGameCallerSession) SHOUTSFUNDSRECIPIENT() (common.Address, error) {
	return _KnifeGame.Contract.SHOUTSFUNDSRECIPIENT(&_KnifeGame.CallOpts)
}

// GameStart is a free data retrieval call binding the contract method 0x3218b99d.
//
// Solidity: function gameStart() view returns(uint256)
func (_KnifeGame *KnifeGameCaller) GameStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "gameStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GameStart is a free data retrieval call binding the contract method 0x3218b99d.
//
// Solidity: function gameStart() view returns(uint256)
func (_KnifeGame *KnifeGameSession) GameStart() (*big.Int, error) {
	return _KnifeGame.Contract.GameStart(&_KnifeGame.CallOpts)
}

// GameStart is a free data retrieval call binding the contract method 0x3218b99d.
//
// Solidity: function gameStart() view returns(uint256)
func (_KnifeGame *KnifeGameCallerSession) GameStart() (*big.Int, error) {
	return _KnifeGame.Contract.GameStart(&_KnifeGame.CallOpts)
}

// HasUserClaimedFreeMooTokens is a free data retrieval call binding the contract method 0x4c27b6ca.
//
// Solidity: function hasUserClaimedFreeMooTokens(address ) view returns(bool)
func (_KnifeGame *KnifeGameCaller) HasUserClaimedFreeMooTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "hasUserClaimedFreeMooTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserClaimedFreeMooTokens is a free data retrieval call binding the contract method 0x4c27b6ca.
//
// Solidity: function hasUserClaimedFreeMooTokens(address ) view returns(bool)
func (_KnifeGame *KnifeGameSession) HasUserClaimedFreeMooTokens(arg0 common.Address) (bool, error) {
	return _KnifeGame.Contract.HasUserClaimedFreeMooTokens(&_KnifeGame.CallOpts, arg0)
}

// HasUserClaimedFreeMooTokens is a free data retrieval call binding the contract method 0x4c27b6ca.
//
// Solidity: function hasUserClaimedFreeMooTokens(address ) view returns(bool)
func (_KnifeGame *KnifeGameCallerSession) HasUserClaimedFreeMooTokens(arg0 common.Address) (bool, error) {
	return _KnifeGame.Contract.HasUserClaimedFreeMooTokens(&_KnifeGame.CallOpts, arg0)
}

// HasUserPrepurchased is a free data retrieval call binding the contract method 0x1bb47402.
//
// Solidity: function hasUserPrepurchased(address ) view returns(bool)
func (_KnifeGame *KnifeGameCaller) HasUserPrepurchased(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "hasUserPrepurchased", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUserPrepurchased is a free data retrieval call binding the contract method 0x1bb47402.
//
// Solidity: function hasUserPrepurchased(address ) view returns(bool)
func (_KnifeGame *KnifeGameSession) HasUserPrepurchased(arg0 common.Address) (bool, error) {
	return _KnifeGame.Contract.HasUserPrepurchased(&_KnifeGame.CallOpts, arg0)
}

// HasUserPrepurchased is a free data retrieval call binding the contract method 0x1bb47402.
//
// Solidity: function hasUserPrepurchased(address ) view returns(bool)
func (_KnifeGame *KnifeGameCallerSession) HasUserPrepurchased(arg0 common.Address) (bool, error) {
	return _KnifeGame.Contract.HasUserPrepurchased(&_KnifeGame.CallOpts, arg0)
}

// IsReimbursed is a free data retrieval call binding the contract method 0xb6bef898.
//
// Solidity: function isReimbursed() view returns(bool)
func (_KnifeGame *KnifeGameCaller) IsReimbursed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "isReimbursed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReimbursed is a free data retrieval call binding the contract method 0xb6bef898.
//
// Solidity: function isReimbursed() view returns(bool)
func (_KnifeGame *KnifeGameSession) IsReimbursed() (bool, error) {
	return _KnifeGame.Contract.IsReimbursed(&_KnifeGame.CallOpts)
}

// IsReimbursed is a free data retrieval call binding the contract method 0xb6bef898.
//
// Solidity: function isReimbursed() view returns(bool)
func (_KnifeGame *KnifeGameCallerSession) IsReimbursed() (bool, error) {
	return _KnifeGame.Contract.IsReimbursed(&_KnifeGame.CallOpts)
}

// KnifeLVRGDA is a free data retrieval call binding the contract method 0xde28179e.
//
// Solidity: function knifeLVRGDA() view returns(address)
func (_KnifeGame *KnifeGameCaller) KnifeLVRGDA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "knifeLVRGDA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KnifeLVRGDA is a free data retrieval call binding the contract method 0xde28179e.
//
// Solidity: function knifeLVRGDA() view returns(address)
func (_KnifeGame *KnifeGameSession) KnifeLVRGDA() (common.Address, error) {
	return _KnifeGame.Contract.KnifeLVRGDA(&_KnifeGame.CallOpts)
}

// KnifeLVRGDA is a free data retrieval call binding the contract method 0xde28179e.
//
// Solidity: function knifeLVRGDA() view returns(address)
func (_KnifeGame *KnifeGameCallerSession) KnifeLVRGDA() (common.Address, error) {
	return _KnifeGame.Contract.KnifeLVRGDA(&_KnifeGame.CallOpts)
}

// KnifeNFT is a free data retrieval call binding the contract method 0x07e05f97.
//
// Solidity: function knifeNFT() view returns(address)
func (_KnifeGame *KnifeGameCaller) KnifeNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "knifeNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KnifeNFT is a free data retrieval call binding the contract method 0x07e05f97.
//
// Solidity: function knifeNFT() view returns(address)
func (_KnifeGame *KnifeGameSession) KnifeNFT() (common.Address, error) {
	return _KnifeGame.Contract.KnifeNFT(&_KnifeGame.CallOpts)
}

// KnifeNFT is a free data retrieval call binding the contract method 0x07e05f97.
//
// Solidity: function knifeNFT() view returns(address)
func (_KnifeGame *KnifeGameCallerSession) KnifeNFT() (common.Address, error) {
	return _KnifeGame.Contract.KnifeNFT(&_KnifeGame.CallOpts)
}

// KnifePrice is a free data retrieval call binding the contract method 0x1a40084d.
//
// Solidity: function knifePrice() view returns(uint256)
func (_KnifeGame *KnifeGameCaller) KnifePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "knifePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KnifePrice is a free data retrieval call binding the contract method 0x1a40084d.
//
// Solidity: function knifePrice() view returns(uint256)
func (_KnifeGame *KnifeGameSession) KnifePrice() (*big.Int, error) {
	return _KnifeGame.Contract.KnifePrice(&_KnifeGame.CallOpts)
}

// KnifePrice is a free data retrieval call binding the contract method 0x1a40084d.
//
// Solidity: function knifePrice() view returns(uint256)
func (_KnifeGame *KnifeGameCallerSession) KnifePrice() (*big.Int, error) {
	return _KnifeGame.Contract.KnifePrice(&_KnifeGame.CallOpts)
}

// KnivesMintedFromMoo is a free data retrieval call binding the contract method 0x8b56f487.
//
// Solidity: function knivesMintedFromMoo() view returns(uint128)
func (_KnifeGame *KnifeGameCaller) KnivesMintedFromMoo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "knivesMintedFromMoo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KnivesMintedFromMoo is a free data retrieval call binding the contract method 0x8b56f487.
//
// Solidity: function knivesMintedFromMoo() view returns(uint128)
func (_KnifeGame *KnifeGameSession) KnivesMintedFromMoo() (*big.Int, error) {
	return _KnifeGame.Contract.KnivesMintedFromMoo(&_KnifeGame.CallOpts)
}

// KnivesMintedFromMoo is a free data retrieval call binding the contract method 0x8b56f487.
//
// Solidity: function knivesMintedFromMoo() view returns(uint128)
func (_KnifeGame *KnifeGameCallerSession) KnivesMintedFromMoo() (*big.Int, error) {
	return _KnifeGame.Contract.KnivesMintedFromMoo(&_KnifeGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KnifeGame *KnifeGameCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KnifeGame *KnifeGameSession) Owner() (common.Address, error) {
	return _KnifeGame.Contract.Owner(&_KnifeGame.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KnifeGame *KnifeGameCallerSession) Owner() (common.Address, error) {
	return _KnifeGame.Contract.Owner(&_KnifeGame.CallOpts)
}

// SpiesMintedFromMoo is a free data retrieval call binding the contract method 0xc3a48baa.
//
// Solidity: function spiesMintedFromMoo() view returns(uint128)
func (_KnifeGame *KnifeGameCaller) SpiesMintedFromMoo(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "spiesMintedFromMoo")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpiesMintedFromMoo is a free data retrieval call binding the contract method 0xc3a48baa.
//
// Solidity: function spiesMintedFromMoo() view returns(uint128)
func (_KnifeGame *KnifeGameSession) SpiesMintedFromMoo() (*big.Int, error) {
	return _KnifeGame.Contract.SpiesMintedFromMoo(&_KnifeGame.CallOpts)
}

// SpiesMintedFromMoo is a free data retrieval call binding the contract method 0xc3a48baa.
//
// Solidity: function spiesMintedFromMoo() view returns(uint128)
func (_KnifeGame *KnifeGameCallerSession) SpiesMintedFromMoo() (*big.Int, error) {
	return _KnifeGame.Contract.SpiesMintedFromMoo(&_KnifeGame.CallOpts)
}

// SpyLVRGDA is a free data retrieval call binding the contract method 0xb3579862.
//
// Solidity: function spyLVRGDA() view returns(address)
func (_KnifeGame *KnifeGameCaller) SpyLVRGDA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "spyLVRGDA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpyLVRGDA is a free data retrieval call binding the contract method 0xb3579862.
//
// Solidity: function spyLVRGDA() view returns(address)
func (_KnifeGame *KnifeGameSession) SpyLVRGDA() (common.Address, error) {
	return _KnifeGame.Contract.SpyLVRGDA(&_KnifeGame.CallOpts)
}

// SpyLVRGDA is a free data retrieval call binding the contract method 0xb3579862.
//
// Solidity: function spyLVRGDA() view returns(address)
func (_KnifeGame *KnifeGameCallerSession) SpyLVRGDA() (common.Address, error) {
	return _KnifeGame.Contract.SpyLVRGDA(&_KnifeGame.CallOpts)
}

// SpyNFT is a free data retrieval call binding the contract method 0x95e9143f.
//
// Solidity: function spyNFT() view returns(address)
func (_KnifeGame *KnifeGameCaller) SpyNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "spyNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SpyNFT is a free data retrieval call binding the contract method 0x95e9143f.
//
// Solidity: function spyNFT() view returns(address)
func (_KnifeGame *KnifeGameSession) SpyNFT() (common.Address, error) {
	return _KnifeGame.Contract.SpyNFT(&_KnifeGame.CallOpts)
}

// SpyNFT is a free data retrieval call binding the contract method 0x95e9143f.
//
// Solidity: function spyNFT() view returns(address)
func (_KnifeGame *KnifeGameCallerSession) SpyNFT() (common.Address, error) {
	return _KnifeGame.Contract.SpyNFT(&_KnifeGame.CallOpts)
}

// SpyPrice is a free data retrieval call binding the contract method 0xc11339a4.
//
// Solidity: function spyPrice() view returns(uint256)
func (_KnifeGame *KnifeGameCaller) SpyPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "spyPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpyPrice is a free data retrieval call binding the contract method 0xc11339a4.
//
// Solidity: function spyPrice() view returns(uint256)
func (_KnifeGame *KnifeGameSession) SpyPrice() (*big.Int, error) {
	return _KnifeGame.Contract.SpyPrice(&_KnifeGame.CallOpts)
}

// SpyPrice is a free data retrieval call binding the contract method 0xc11339a4.
//
// Solidity: function spyPrice() view returns(uint256)
func (_KnifeGame *KnifeGameCallerSession) SpyPrice() (*big.Int, error) {
	return _KnifeGame.Contract.SpyPrice(&_KnifeGame.CallOpts)
}

// SpyPriceETH is a free data retrieval call binding the contract method 0x7e04e55f.
//
// Solidity: function spyPriceETH(address _user) view returns(uint256)
func (_KnifeGame *KnifeGameCaller) SpyPriceETH(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "spyPriceETH", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpyPriceETH is a free data retrieval call binding the contract method 0x7e04e55f.
//
// Solidity: function spyPriceETH(address _user) view returns(uint256)
func (_KnifeGame *KnifeGameSession) SpyPriceETH(_user common.Address) (*big.Int, error) {
	return _KnifeGame.Contract.SpyPriceETH(&_KnifeGame.CallOpts, _user)
}

// SpyPriceETH is a free data retrieval call binding the contract method 0x7e04e55f.
//
// Solidity: function spyPriceETH(address _user) view returns(uint256)
func (_KnifeGame *KnifeGameCallerSession) SpyPriceETH(_user common.Address) (*big.Int, error) {
	return _KnifeGame.Contract.SpyPriceETH(&_KnifeGame.CallOpts, _user)
}

// UserPurchasesOnDay is a free data retrieval call binding the contract method 0x277ddbaa.
//
// Solidity: function userPurchasesOnDay(address , uint256 ) view returns(uint256)
func (_KnifeGame *KnifeGameCaller) UserPurchasesOnDay(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _KnifeGame.contract.Call(opts, &out, "userPurchasesOnDay", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPurchasesOnDay is a free data retrieval call binding the contract method 0x277ddbaa.
//
// Solidity: function userPurchasesOnDay(address , uint256 ) view returns(uint256)
func (_KnifeGame *KnifeGameSession) UserPurchasesOnDay(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _KnifeGame.Contract.UserPurchasesOnDay(&_KnifeGame.CallOpts, arg0, arg1)
}

// UserPurchasesOnDay is a free data retrieval call binding the contract method 0x277ddbaa.
//
// Solidity: function userPurchasesOnDay(address , uint256 ) view returns(uint256)
func (_KnifeGame *KnifeGameCallerSession) UserPurchasesOnDay(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _KnifeGame.Contract.UserPurchasesOnDay(&_KnifeGame.CallOpts, arg0, arg1)
}

// ClaimFreeMoo is a paid mutator transaction binding the contract method 0x23f02f87.
//
// Solidity: function claimFreeMoo() returns()
func (_KnifeGame *KnifeGameTransactor) ClaimFreeMoo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "claimFreeMoo")
}

// ClaimFreeMoo is a paid mutator transaction binding the contract method 0x23f02f87.
//
// Solidity: function claimFreeMoo() returns()
func (_KnifeGame *KnifeGameSession) ClaimFreeMoo() (*types.Transaction, error) {
	return _KnifeGame.Contract.ClaimFreeMoo(&_KnifeGame.TransactOpts)
}

// ClaimFreeMoo is a paid mutator transaction binding the contract method 0x23f02f87.
//
// Solidity: function claimFreeMoo() returns()
func (_KnifeGame *KnifeGameTransactorSession) ClaimFreeMoo() (*types.Transaction, error) {
	return _KnifeGame.Contract.ClaimFreeMoo(&_KnifeGame.TransactOpts)
}

// KillSpy is a paid mutator transaction binding the contract method 0xa9ad5e6f.
//
// Solidity: function killSpy(uint256 _knifeId, uint256 _spyId) returns()
func (_KnifeGame *KnifeGameTransactor) KillSpy(opts *bind.TransactOpts, _knifeId *big.Int, _spyId *big.Int) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "killSpy", _knifeId, _spyId)
}

// KillSpy is a paid mutator transaction binding the contract method 0xa9ad5e6f.
//
// Solidity: function killSpy(uint256 _knifeId, uint256 _spyId) returns()
func (_KnifeGame *KnifeGameSession) KillSpy(_knifeId *big.Int, _spyId *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.KillSpy(&_KnifeGame.TransactOpts, _knifeId, _spyId)
}

// KillSpy is a paid mutator transaction binding the contract method 0xa9ad5e6f.
//
// Solidity: function killSpy(uint256 _knifeId, uint256 _spyId) returns()
func (_KnifeGame *KnifeGameTransactorSession) KillSpy(_knifeId *big.Int, _spyId *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.KillSpy(&_KnifeGame.TransactOpts, _knifeId, _spyId)
}

// MintKnifeFromMoolah is a paid mutator transaction binding the contract method 0xd5493aa7.
//
// Solidity: function mintKnifeFromMoolah(uint256 _maxPrice) returns(uint256 knifeId)
func (_KnifeGame *KnifeGameTransactor) MintKnifeFromMoolah(opts *bind.TransactOpts, _maxPrice *big.Int) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "mintKnifeFromMoolah", _maxPrice)
}

// MintKnifeFromMoolah is a paid mutator transaction binding the contract method 0xd5493aa7.
//
// Solidity: function mintKnifeFromMoolah(uint256 _maxPrice) returns(uint256 knifeId)
func (_KnifeGame *KnifeGameSession) MintKnifeFromMoolah(_maxPrice *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.MintKnifeFromMoolah(&_KnifeGame.TransactOpts, _maxPrice)
}

// MintKnifeFromMoolah is a paid mutator transaction binding the contract method 0xd5493aa7.
//
// Solidity: function mintKnifeFromMoolah(uint256 _maxPrice) returns(uint256 knifeId)
func (_KnifeGame *KnifeGameTransactorSession) MintKnifeFromMoolah(_maxPrice *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.MintKnifeFromMoolah(&_KnifeGame.TransactOpts, _maxPrice)
}

// MintSpyFromMoolah is a paid mutator transaction binding the contract method 0x55ab5fbd.
//
// Solidity: function mintSpyFromMoolah(uint256 _maxPrice) returns(uint256 spyId)
func (_KnifeGame *KnifeGameTransactor) MintSpyFromMoolah(opts *bind.TransactOpts, _maxPrice *big.Int) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "mintSpyFromMoolah", _maxPrice)
}

// MintSpyFromMoolah is a paid mutator transaction binding the contract method 0x55ab5fbd.
//
// Solidity: function mintSpyFromMoolah(uint256 _maxPrice) returns(uint256 spyId)
func (_KnifeGame *KnifeGameSession) MintSpyFromMoolah(_maxPrice *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.MintSpyFromMoolah(&_KnifeGame.TransactOpts, _maxPrice)
}

// MintSpyFromMoolah is a paid mutator transaction binding the contract method 0x55ab5fbd.
//
// Solidity: function mintSpyFromMoolah(uint256 _maxPrice) returns(uint256 spyId)
func (_KnifeGame *KnifeGameTransactorSession) MintSpyFromMoolah(_maxPrice *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.MintSpyFromMoolah(&_KnifeGame.TransactOpts, _maxPrice)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_KnifeGame *KnifeGameTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_KnifeGame *KnifeGameSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _KnifeGame.Contract.OnERC721Received(&_KnifeGame.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_KnifeGame *KnifeGameTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _KnifeGame.Contract.OnERC721Received(&_KnifeGame.TransactOpts, arg0, arg1, arg2, arg3)
}

// PurchaseSpy is a paid mutator transaction binding the contract method 0xbe501955.
//
// Solidity: function purchaseSpy(address user) payable returns(uint256 spyId)
func (_KnifeGame *KnifeGameTransactor) PurchaseSpy(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "purchaseSpy", user)
}

// PurchaseSpy is a paid mutator transaction binding the contract method 0xbe501955.
//
// Solidity: function purchaseSpy(address user) payable returns(uint256 spyId)
func (_KnifeGame *KnifeGameSession) PurchaseSpy(user common.Address) (*types.Transaction, error) {
	return _KnifeGame.Contract.PurchaseSpy(&_KnifeGame.TransactOpts, user)
}

// PurchaseSpy is a paid mutator transaction binding the contract method 0xbe501955.
//
// Solidity: function purchaseSpy(address user) payable returns(uint256 spyId)
func (_KnifeGame *KnifeGameTransactorSession) PurchaseSpy(user common.Address) (*types.Transaction, error) {
	return _KnifeGame.Contract.PurchaseSpy(&_KnifeGame.TransactOpts, user)
}

// ReimburseSpy is a paid mutator transaction binding the contract method 0xa7caef3f.
//
// Solidity: function reimburseSpy(address[] _reimburse) returns()
func (_KnifeGame *KnifeGameTransactor) ReimburseSpy(opts *bind.TransactOpts, _reimburse []common.Address) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "reimburseSpy", _reimburse)
}

// ReimburseSpy is a paid mutator transaction binding the contract method 0xa7caef3f.
//
// Solidity: function reimburseSpy(address[] _reimburse) returns()
func (_KnifeGame *KnifeGameSession) ReimburseSpy(_reimburse []common.Address) (*types.Transaction, error) {
	return _KnifeGame.Contract.ReimburseSpy(&_KnifeGame.TransactOpts, _reimburse)
}

// ReimburseSpy is a paid mutator transaction binding the contract method 0xa7caef3f.
//
// Solidity: function reimburseSpy(address[] _reimburse) returns()
func (_KnifeGame *KnifeGameTransactorSession) ReimburseSpy(_reimburse []common.Address) (*types.Transaction, error) {
	return _KnifeGame.Contract.ReimburseSpy(&_KnifeGame.TransactOpts, _reimburse)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KnifeGame *KnifeGameTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KnifeGame *KnifeGameSession) RenounceOwnership() (*types.Transaction, error) {
	return _KnifeGame.Contract.RenounceOwnership(&_KnifeGame.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KnifeGame *KnifeGameTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KnifeGame.Contract.RenounceOwnership(&_KnifeGame.TransactOpts)
}

// Shout is a paid mutator transaction binding the contract method 0x625d2055.
//
// Solidity: function shout(string message) payable returns()
func (_KnifeGame *KnifeGameTransactor) Shout(opts *bind.TransactOpts, message string) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "shout", message)
}

// Shout is a paid mutator transaction binding the contract method 0x625d2055.
//
// Solidity: function shout(string message) payable returns()
func (_KnifeGame *KnifeGameSession) Shout(message string) (*types.Transaction, error) {
	return _KnifeGame.Contract.Shout(&_KnifeGame.TransactOpts, message)
}

// Shout is a paid mutator transaction binding the contract method 0x625d2055.
//
// Solidity: function shout(string message) payable returns()
func (_KnifeGame *KnifeGameTransactorSession) Shout(message string) (*types.Transaction, error) {
	return _KnifeGame.Contract.Shout(&_KnifeGame.TransactOpts, message)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KnifeGame *KnifeGameTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KnifeGame *KnifeGameSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KnifeGame.Contract.TransferOwnership(&_KnifeGame.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KnifeGame *KnifeGameTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KnifeGame.Contract.TransferOwnership(&_KnifeGame.TransactOpts, newOwner)
}

// UpdateKnifeVRGDA is a paid mutator transaction binding the contract method 0x317be1b7.
//
// Solidity: function updateKnifeVRGDA(int256 _targetPrice, int256 _priceDecayPercent, uint256 _maxAmount, int256 _timeScale) returns()
func (_KnifeGame *KnifeGameTransactor) UpdateKnifeVRGDA(opts *bind.TransactOpts, _targetPrice *big.Int, _priceDecayPercent *big.Int, _maxAmount *big.Int, _timeScale *big.Int) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "updateKnifeVRGDA", _targetPrice, _priceDecayPercent, _maxAmount, _timeScale)
}

// UpdateKnifeVRGDA is a paid mutator transaction binding the contract method 0x317be1b7.
//
// Solidity: function updateKnifeVRGDA(int256 _targetPrice, int256 _priceDecayPercent, uint256 _maxAmount, int256 _timeScale) returns()
func (_KnifeGame *KnifeGameSession) UpdateKnifeVRGDA(_targetPrice *big.Int, _priceDecayPercent *big.Int, _maxAmount *big.Int, _timeScale *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.UpdateKnifeVRGDA(&_KnifeGame.TransactOpts, _targetPrice, _priceDecayPercent, _maxAmount, _timeScale)
}

// UpdateKnifeVRGDA is a paid mutator transaction binding the contract method 0x317be1b7.
//
// Solidity: function updateKnifeVRGDA(int256 _targetPrice, int256 _priceDecayPercent, uint256 _maxAmount, int256 _timeScale) returns()
func (_KnifeGame *KnifeGameTransactorSession) UpdateKnifeVRGDA(_targetPrice *big.Int, _priceDecayPercent *big.Int, _maxAmount *big.Int, _timeScale *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.UpdateKnifeVRGDA(&_KnifeGame.TransactOpts, _targetPrice, _priceDecayPercent, _maxAmount, _timeScale)
}

// UpdateSpyVRGDA is a paid mutator transaction binding the contract method 0xce9b0ec0.
//
// Solidity: function updateSpyVRGDA(int256 _targetPrice, int256 _priceDecayPercent, uint256 _maxAmount, int256 _timeScale) returns()
func (_KnifeGame *KnifeGameTransactor) UpdateSpyVRGDA(opts *bind.TransactOpts, _targetPrice *big.Int, _priceDecayPercent *big.Int, _maxAmount *big.Int, _timeScale *big.Int) (*types.Transaction, error) {
	return _KnifeGame.contract.Transact(opts, "updateSpyVRGDA", _targetPrice, _priceDecayPercent, _maxAmount, _timeScale)
}

// UpdateSpyVRGDA is a paid mutator transaction binding the contract method 0xce9b0ec0.
//
// Solidity: function updateSpyVRGDA(int256 _targetPrice, int256 _priceDecayPercent, uint256 _maxAmount, int256 _timeScale) returns()
func (_KnifeGame *KnifeGameSession) UpdateSpyVRGDA(_targetPrice *big.Int, _priceDecayPercent *big.Int, _maxAmount *big.Int, _timeScale *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.UpdateSpyVRGDA(&_KnifeGame.TransactOpts, _targetPrice, _priceDecayPercent, _maxAmount, _timeScale)
}

// UpdateSpyVRGDA is a paid mutator transaction binding the contract method 0xce9b0ec0.
//
// Solidity: function updateSpyVRGDA(int256 _targetPrice, int256 _priceDecayPercent, uint256 _maxAmount, int256 _timeScale) returns()
func (_KnifeGame *KnifeGameTransactorSession) UpdateSpyVRGDA(_targetPrice *big.Int, _priceDecayPercent *big.Int, _maxAmount *big.Int, _timeScale *big.Int) (*types.Transaction, error) {
	return _KnifeGame.Contract.UpdateSpyVRGDA(&_KnifeGame.TransactOpts, _targetPrice, _priceDecayPercent, _maxAmount, _timeScale)
}

// KnifeGameKnifePurchasedIterator is returned from FilterKnifePurchased and is used to iterate over the raw logs and unpacked data for KnifePurchased events raised by the KnifeGame contract.
type KnifeGameKnifePurchasedIterator struct {
	Event *KnifeGameKnifePurchased // Event containing the contract specifics and raw log

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
func (it *KnifeGameKnifePurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KnifeGameKnifePurchased)
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
		it.Event = new(KnifeGameKnifePurchased)
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
func (it *KnifeGameKnifePurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KnifeGameKnifePurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KnifeGameKnifePurchased represents a KnifePurchased event raised by the KnifeGame contract.
type KnifeGameKnifePurchased struct {
	Recipient common.Address
	Id        *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterKnifePurchased is a free log retrieval operation binding the contract event 0x63ea169cbcbcf8d736b710d6ef4ff9da8ececc2cc5f5403f2b12b47d5f965e00.
//
// Solidity: event KnifePurchased(address indexed recipient, uint256 indexed id, uint256 price)
func (_KnifeGame *KnifeGameFilterer) FilterKnifePurchased(opts *bind.FilterOpts, recipient []common.Address, id []*big.Int) (*KnifeGameKnifePurchasedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _KnifeGame.contract.FilterLogs(opts, "KnifePurchased", recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &KnifeGameKnifePurchasedIterator{contract: _KnifeGame.contract, event: "KnifePurchased", logs: logs, sub: sub}, nil
}

// WatchKnifePurchased is a free log subscription operation binding the contract event 0x63ea169cbcbcf8d736b710d6ef4ff9da8ececc2cc5f5403f2b12b47d5f965e00.
//
// Solidity: event KnifePurchased(address indexed recipient, uint256 indexed id, uint256 price)
func (_KnifeGame *KnifeGameFilterer) WatchKnifePurchased(opts *bind.WatchOpts, sink chan<- *KnifeGameKnifePurchased, recipient []common.Address, id []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _KnifeGame.contract.WatchLogs(opts, "KnifePurchased", recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KnifeGameKnifePurchased)
				if err := _KnifeGame.contract.UnpackLog(event, "KnifePurchased", log); err != nil {
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

// ParseKnifePurchased is a log parse operation binding the contract event 0x63ea169cbcbcf8d736b710d6ef4ff9da8ececc2cc5f5403f2b12b47d5f965e00.
//
// Solidity: event KnifePurchased(address indexed recipient, uint256 indexed id, uint256 price)
func (_KnifeGame *KnifeGameFilterer) ParseKnifePurchased(log types.Log) (*KnifeGameKnifePurchased, error) {
	event := new(KnifeGameKnifePurchased)
	if err := _KnifeGame.contract.UnpackLog(event, "KnifePurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KnifeGameOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KnifeGame contract.
type KnifeGameOwnershipTransferredIterator struct {
	Event *KnifeGameOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KnifeGameOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KnifeGameOwnershipTransferred)
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
		it.Event = new(KnifeGameOwnershipTransferred)
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
func (it *KnifeGameOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KnifeGameOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KnifeGameOwnershipTransferred represents a OwnershipTransferred event raised by the KnifeGame contract.
type KnifeGameOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KnifeGame *KnifeGameFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KnifeGameOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KnifeGame.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KnifeGameOwnershipTransferredIterator{contract: _KnifeGame.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KnifeGame *KnifeGameFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KnifeGameOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KnifeGame.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KnifeGameOwnershipTransferred)
				if err := _KnifeGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KnifeGame *KnifeGameFilterer) ParseOwnershipTransferred(log types.Log) (*KnifeGameOwnershipTransferred, error) {
	event := new(KnifeGameOwnershipTransferred)
	if err := _KnifeGame.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KnifeGameShoutedIterator is returned from FilterShouted and is used to iterate over the raw logs and unpacked data for Shouted events raised by the KnifeGame contract.
type KnifeGameShoutedIterator struct {
	Event *KnifeGameShouted // Event containing the contract specifics and raw log

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
func (it *KnifeGameShoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KnifeGameShouted)
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
		it.Event = new(KnifeGameShouted)
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
func (it *KnifeGameShoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KnifeGameShoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KnifeGameShouted represents a Shouted event raised by the KnifeGame contract.
type KnifeGameShouted struct {
	Sender  common.Address
	Message string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterShouted is a free log retrieval operation binding the contract event 0x2a61c0d7ddc4a70821f0286174956bf3d0d434e921873e1c34a3896e96bb6f45.
//
// Solidity: event Shouted(address indexed sender, string message)
func (_KnifeGame *KnifeGameFilterer) FilterShouted(opts *bind.FilterOpts, sender []common.Address) (*KnifeGameShoutedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KnifeGame.contract.FilterLogs(opts, "Shouted", senderRule)
	if err != nil {
		return nil, err
	}
	return &KnifeGameShoutedIterator{contract: _KnifeGame.contract, event: "Shouted", logs: logs, sub: sub}, nil
}

// WatchShouted is a free log subscription operation binding the contract event 0x2a61c0d7ddc4a70821f0286174956bf3d0d434e921873e1c34a3896e96bb6f45.
//
// Solidity: event Shouted(address indexed sender, string message)
func (_KnifeGame *KnifeGameFilterer) WatchShouted(opts *bind.WatchOpts, sink chan<- *KnifeGameShouted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KnifeGame.contract.WatchLogs(opts, "Shouted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KnifeGameShouted)
				if err := _KnifeGame.contract.UnpackLog(event, "Shouted", log); err != nil {
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

// ParseShouted is a log parse operation binding the contract event 0x2a61c0d7ddc4a70821f0286174956bf3d0d434e921873e1c34a3896e96bb6f45.
//
// Solidity: event Shouted(address indexed sender, string message)
func (_KnifeGame *KnifeGameFilterer) ParseShouted(log types.Log) (*KnifeGameShouted, error) {
	event := new(KnifeGameShouted)
	if err := _KnifeGame.contract.UnpackLog(event, "Shouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KnifeGameSpyKilledIterator is returned from FilterSpyKilled and is used to iterate over the raw logs and unpacked data for SpyKilled events raised by the KnifeGame contract.
type KnifeGameSpyKilledIterator struct {
	Event *KnifeGameSpyKilled // Event containing the contract specifics and raw log

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
func (it *KnifeGameSpyKilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KnifeGameSpyKilled)
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
		it.Event = new(KnifeGameSpyKilled)
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
func (it *KnifeGameSpyKilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KnifeGameSpyKilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KnifeGameSpyKilled represents a SpyKilled event raised by the KnifeGame contract.
type KnifeGameSpyKilled struct {
	Hitman  common.Address
	Victim  common.Address
	KnifeId *big.Int
	SpyId   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpyKilled is a free log retrieval operation binding the contract event 0x45e465566601be3e6eb4963d1b7cefe8cba5f8fda38b9107329955fea6ce9467.
//
// Solidity: event SpyKilled(address indexed hitman, address indexed victim, uint256 knifeId, uint256 spyId)
func (_KnifeGame *KnifeGameFilterer) FilterSpyKilled(opts *bind.FilterOpts, hitman []common.Address, victim []common.Address) (*KnifeGameSpyKilledIterator, error) {

	var hitmanRule []interface{}
	for _, hitmanItem := range hitman {
		hitmanRule = append(hitmanRule, hitmanItem)
	}
	var victimRule []interface{}
	for _, victimItem := range victim {
		victimRule = append(victimRule, victimItem)
	}

	logs, sub, err := _KnifeGame.contract.FilterLogs(opts, "SpyKilled", hitmanRule, victimRule)
	if err != nil {
		return nil, err
	}
	return &KnifeGameSpyKilledIterator{contract: _KnifeGame.contract, event: "SpyKilled", logs: logs, sub: sub}, nil
}

// WatchSpyKilled is a free log subscription operation binding the contract event 0x45e465566601be3e6eb4963d1b7cefe8cba5f8fda38b9107329955fea6ce9467.
//
// Solidity: event SpyKilled(address indexed hitman, address indexed victim, uint256 knifeId, uint256 spyId)
func (_KnifeGame *KnifeGameFilterer) WatchSpyKilled(opts *bind.WatchOpts, sink chan<- *KnifeGameSpyKilled, hitman []common.Address, victim []common.Address) (event.Subscription, error) {

	var hitmanRule []interface{}
	for _, hitmanItem := range hitman {
		hitmanRule = append(hitmanRule, hitmanItem)
	}
	var victimRule []interface{}
	for _, victimItem := range victim {
		victimRule = append(victimRule, victimItem)
	}

	logs, sub, err := _KnifeGame.contract.WatchLogs(opts, "SpyKilled", hitmanRule, victimRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KnifeGameSpyKilled)
				if err := _KnifeGame.contract.UnpackLog(event, "SpyKilled", log); err != nil {
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

// ParseSpyKilled is a log parse operation binding the contract event 0x45e465566601be3e6eb4963d1b7cefe8cba5f8fda38b9107329955fea6ce9467.
//
// Solidity: event SpyKilled(address indexed hitman, address indexed victim, uint256 knifeId, uint256 spyId)
func (_KnifeGame *KnifeGameFilterer) ParseSpyKilled(log types.Log) (*KnifeGameSpyKilled, error) {
	event := new(KnifeGameSpyKilled)
	if err := _KnifeGame.contract.UnpackLog(event, "SpyKilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KnifeGameSpyPurchasedETHIterator is returned from FilterSpyPurchasedETH and is used to iterate over the raw logs and unpacked data for SpyPurchasedETH events raised by the KnifeGame contract.
type KnifeGameSpyPurchasedETHIterator struct {
	Event *KnifeGameSpyPurchasedETH // Event containing the contract specifics and raw log

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
func (it *KnifeGameSpyPurchasedETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KnifeGameSpyPurchasedETH)
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
		it.Event = new(KnifeGameSpyPurchasedETH)
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
func (it *KnifeGameSpyPurchasedETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KnifeGameSpyPurchasedETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KnifeGameSpyPurchasedETH represents a SpyPurchasedETH event raised by the KnifeGame contract.
type KnifeGameSpyPurchasedETH struct {
	Recipient common.Address
	Id        *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSpyPurchasedETH is a free log retrieval operation binding the contract event 0x81901b6ecf8f0aabcd532317ec6918cbd7c92fe3639de618fc5f6a55c11a5671.
//
// Solidity: event SpyPurchasedETH(address indexed recipient, uint256 indexed id)
func (_KnifeGame *KnifeGameFilterer) FilterSpyPurchasedETH(opts *bind.FilterOpts, recipient []common.Address, id []*big.Int) (*KnifeGameSpyPurchasedETHIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _KnifeGame.contract.FilterLogs(opts, "SpyPurchasedETH", recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &KnifeGameSpyPurchasedETHIterator{contract: _KnifeGame.contract, event: "SpyPurchasedETH", logs: logs, sub: sub}, nil
}

// WatchSpyPurchasedETH is a free log subscription operation binding the contract event 0x81901b6ecf8f0aabcd532317ec6918cbd7c92fe3639de618fc5f6a55c11a5671.
//
// Solidity: event SpyPurchasedETH(address indexed recipient, uint256 indexed id)
func (_KnifeGame *KnifeGameFilterer) WatchSpyPurchasedETH(opts *bind.WatchOpts, sink chan<- *KnifeGameSpyPurchasedETH, recipient []common.Address, id []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _KnifeGame.contract.WatchLogs(opts, "SpyPurchasedETH", recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KnifeGameSpyPurchasedETH)
				if err := _KnifeGame.contract.UnpackLog(event, "SpyPurchasedETH", log); err != nil {
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

// ParseSpyPurchasedETH is a log parse operation binding the contract event 0x81901b6ecf8f0aabcd532317ec6918cbd7c92fe3639de618fc5f6a55c11a5671.
//
// Solidity: event SpyPurchasedETH(address indexed recipient, uint256 indexed id)
func (_KnifeGame *KnifeGameFilterer) ParseSpyPurchasedETH(log types.Log) (*KnifeGameSpyPurchasedETH, error) {
	event := new(KnifeGameSpyPurchasedETH)
	if err := _KnifeGame.contract.UnpackLog(event, "SpyPurchasedETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KnifeGameSpyPurchasedMooIterator is returned from FilterSpyPurchasedMoo and is used to iterate over the raw logs and unpacked data for SpyPurchasedMoo events raised by the KnifeGame contract.
type KnifeGameSpyPurchasedMooIterator struct {
	Event *KnifeGameSpyPurchasedMoo // Event containing the contract specifics and raw log

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
func (it *KnifeGameSpyPurchasedMooIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KnifeGameSpyPurchasedMoo)
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
		it.Event = new(KnifeGameSpyPurchasedMoo)
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
func (it *KnifeGameSpyPurchasedMooIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KnifeGameSpyPurchasedMooIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KnifeGameSpyPurchasedMoo represents a SpyPurchasedMoo event raised by the KnifeGame contract.
type KnifeGameSpyPurchasedMoo struct {
	Recipient common.Address
	Id        *big.Int
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSpyPurchasedMoo is a free log retrieval operation binding the contract event 0xf387c02ea0b7fe0642d2110ca2b211586dc019957d013bc7af259a30a20b5413.
//
// Solidity: event SpyPurchasedMoo(address indexed recipient, uint256 indexed id, uint256 price)
func (_KnifeGame *KnifeGameFilterer) FilterSpyPurchasedMoo(opts *bind.FilterOpts, recipient []common.Address, id []*big.Int) (*KnifeGameSpyPurchasedMooIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _KnifeGame.contract.FilterLogs(opts, "SpyPurchasedMoo", recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &KnifeGameSpyPurchasedMooIterator{contract: _KnifeGame.contract, event: "SpyPurchasedMoo", logs: logs, sub: sub}, nil
}

// WatchSpyPurchasedMoo is a free log subscription operation binding the contract event 0xf387c02ea0b7fe0642d2110ca2b211586dc019957d013bc7af259a30a20b5413.
//
// Solidity: event SpyPurchasedMoo(address indexed recipient, uint256 indexed id, uint256 price)
func (_KnifeGame *KnifeGameFilterer) WatchSpyPurchasedMoo(opts *bind.WatchOpts, sink chan<- *KnifeGameSpyPurchasedMoo, recipient []common.Address, id []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _KnifeGame.contract.WatchLogs(opts, "SpyPurchasedMoo", recipientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KnifeGameSpyPurchasedMoo)
				if err := _KnifeGame.contract.UnpackLog(event, "SpyPurchasedMoo", log); err != nil {
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

// ParseSpyPurchasedMoo is a log parse operation binding the contract event 0xf387c02ea0b7fe0642d2110ca2b211586dc019957d013bc7af259a30a20b5413.
//
// Solidity: event SpyPurchasedMoo(address indexed recipient, uint256 indexed id, uint256 price)
func (_KnifeGame *KnifeGameFilterer) ParseSpyPurchasedMoo(log types.Log) (*KnifeGameSpyPurchasedMoo, error) {
	event := new(KnifeGameSpyPurchasedMoo)
	if err := _KnifeGame.contract.UnpackLog(event, "SpyPurchasedMoo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
