// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// MasterChefV2PoolInfo is an auto generated low-level Go binding around an user-defined struct.
type MasterChefV2PoolInfo struct {
	AccCakePerShare   *big.Int
	LastRewardBlock   *big.Int
	AllocPoint        *big.Int
	TotalBoostedShare *big.Int
	IsRegular         bool
}

// MasterChefV2MetaData contains all meta data concerning the MasterChefV2 contract.
var MasterChefV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIMasterChef\",\"name\":\"_MASTER_CHEF\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_cake\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_MASTER_PID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_burnAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"contractIBEP20\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRegular\",\"type\":\"bool\"}],\"name\":\"AddPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Init\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"}],\"name\":\"SetPool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"boostContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxBoostVal\",\"type\":\"uint256\"}],\"name\":\"UpdateBoostContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"UpdateBurnAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"regularFarmRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"specialFarmRate\",\"type\":\"uint256\"}],\"name\":\"UpdateCakeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"}],\"name\":\"UpdatePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBoostVal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBoostVal\",\"type\":\"uint256\"}],\"name\":\"UpdateUserBoost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"name\":\"UpdateWhiteList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CAKE\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_CHEF\",\"outputs\":[{\"internalType\":\"contractIMasterChef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_PID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"contractIBEP20\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isRegular\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"boostContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnCake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isRegular\",\"type\":\"bool\"}],\"name\":\"cakePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakePerBlockToBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakeRateToBurn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakeRateToRegularFarm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cakeRateToSpecialFarm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"getBoostValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestFromMasterChef\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"dummyToken\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastBurnedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"massUpdatePools\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBoostValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"pendingCake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBoostedShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRegular\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pools\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_withUpdate\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRegularAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSpecialAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBoostContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxBoostVal\",\"type\":\"uint256\"}],\"name\":\"updateBoostContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"updateBurnAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_burnRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_regularFarmRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_specialFarmRate\",\"type\":\"uint256\"}],\"name\":\"updateCakeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updatePool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"accCakePerShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBoostedShare\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isRegular\",\"type\":\"bool\"}],\"internalType\":\"structMasterChefV2.PoolInfo\",\"name\":\"pool\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"updateUserBoost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isValid\",\"type\":\"bool\"}],\"name\":\"updateWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDebt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteListMapping\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MasterChefV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterChefV2MetaData.ABI instead.
var MasterChefV2ABI = MasterChefV2MetaData.ABI

// MasterChefV2 is an auto generated Go binding around an Ethereum contract.
type MasterChefV2 struct {
	MasterChefV2Caller     // Read-only binding to the contract
	MasterChefV2Transactor // Write-only binding to the contract
	MasterChefV2Filterer   // Log filterer for contract events
}

// MasterChefV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MasterChefV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterChefV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterChefV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterChefV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterChefV2Session struct {
	Contract     *MasterChefV2     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterChefV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterChefV2CallerSession struct {
	Contract *MasterChefV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MasterChefV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterChefV2TransactorSession struct {
	Contract     *MasterChefV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MasterChefV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MasterChefV2Raw struct {
	Contract *MasterChefV2 // Generic contract binding to access the raw methods on
}

// MasterChefV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterChefV2CallerRaw struct {
	Contract *MasterChefV2Caller // Generic read-only contract binding to access the raw methods on
}

// MasterChefV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterChefV2TransactorRaw struct {
	Contract *MasterChefV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterChefV2 creates a new instance of MasterChefV2, bound to a specific deployed contract.
func NewMasterChefV2(address common.Address, backend bind.ContractBackend) (*MasterChefV2, error) {
	contract, err := bindMasterChefV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2{MasterChefV2Caller: MasterChefV2Caller{contract: contract}, MasterChefV2Transactor: MasterChefV2Transactor{contract: contract}, MasterChefV2Filterer: MasterChefV2Filterer{contract: contract}}, nil
}

// NewMasterChefV2Caller creates a new read-only instance of MasterChefV2, bound to a specific deployed contract.
func NewMasterChefV2Caller(address common.Address, caller bind.ContractCaller) (*MasterChefV2Caller, error) {
	contract, err := bindMasterChefV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2Caller{contract: contract}, nil
}

// NewMasterChefV2Transactor creates a new write-only instance of MasterChefV2, bound to a specific deployed contract.
func NewMasterChefV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MasterChefV2Transactor, error) {
	contract, err := bindMasterChefV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2Transactor{contract: contract}, nil
}

// NewMasterChefV2Filterer creates a new log filterer instance of MasterChefV2, bound to a specific deployed contract.
func NewMasterChefV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MasterChefV2Filterer, error) {
	contract, err := bindMasterChefV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2Filterer{contract: contract}, nil
}

// bindMasterChefV2 binds a generic wrapper to an already deployed contract.
func bindMasterChefV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MasterChefV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterChefV2 *MasterChefV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterChefV2.Contract.MasterChefV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterChefV2 *MasterChefV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefV2.Contract.MasterChefV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterChefV2 *MasterChefV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterChefV2.Contract.MasterChefV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterChefV2 *MasterChefV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterChefV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterChefV2 *MasterChefV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterChefV2 *MasterChefV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterChefV2.Contract.contract.Transact(opts, method, params...)
}

// CAKE is a free data retrieval call binding the contract method 0x4ca6ef28.
//
// Solidity: function CAKE() view returns(address)
func (_MasterChefV2 *MasterChefV2Caller) CAKE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "CAKE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CAKE is a free data retrieval call binding the contract method 0x4ca6ef28.
//
// Solidity: function CAKE() view returns(address)
func (_MasterChefV2 *MasterChefV2Session) CAKE() (common.Address, error) {
	return _MasterChefV2.Contract.CAKE(&_MasterChefV2.CallOpts)
}

// CAKE is a free data retrieval call binding the contract method 0x4ca6ef28.
//
// Solidity: function CAKE() view returns(address)
func (_MasterChefV2 *MasterChefV2CallerSession) CAKE() (common.Address, error) {
	return _MasterChefV2.Contract.CAKE(&_MasterChefV2.CallOpts)
}

// MASTERCHEF is a free data retrieval call binding the contract method 0xedd8b170.
//
// Solidity: function MASTER_CHEF() view returns(address)
func (_MasterChefV2 *MasterChefV2Caller) MASTERCHEF(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "MASTER_CHEF")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MASTERCHEF is a free data retrieval call binding the contract method 0xedd8b170.
//
// Solidity: function MASTER_CHEF() view returns(address)
func (_MasterChefV2 *MasterChefV2Session) MASTERCHEF() (common.Address, error) {
	return _MasterChefV2.Contract.MASTERCHEF(&_MasterChefV2.CallOpts)
}

// MASTERCHEF is a free data retrieval call binding the contract method 0xedd8b170.
//
// Solidity: function MASTER_CHEF() view returns(address)
func (_MasterChefV2 *MasterChefV2CallerSession) MASTERCHEF() (common.Address, error) {
	return _MasterChefV2.Contract.MASTERCHEF(&_MasterChefV2.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) MASTERPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "MASTER_PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) MASTERPID() (*big.Int, error) {
	return _MasterChefV2.Contract.MASTERPID(&_MasterChefV2.CallOpts)
}

// MASTERPID is a free data retrieval call binding the contract method 0x61621aaa.
//
// Solidity: function MASTER_PID() view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) MASTERPID() (*big.Int, error) {
	return _MasterChefV2.Contract.MASTERPID(&_MasterChefV2.CallOpts)
}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_MasterChefV2 *MasterChefV2Caller) BoostContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "boostContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_MasterChefV2 *MasterChefV2Session) BoostContract() (common.Address, error) {
	return _MasterChefV2.Contract.BoostContract(&_MasterChefV2.CallOpts)
}

// BoostContract is a free data retrieval call binding the contract method 0xdfcedeee.
//
// Solidity: function boostContract() view returns(address)
func (_MasterChefV2 *MasterChefV2CallerSession) BoostContract() (common.Address, error) {
	return _MasterChefV2.Contract.BoostContract(&_MasterChefV2.CallOpts)
}

// BurnAdmin is a free data retrieval call binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() view returns(address)
func (_MasterChefV2 *MasterChefV2Caller) BurnAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "burnAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BurnAdmin is a free data retrieval call binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() view returns(address)
func (_MasterChefV2 *MasterChefV2Session) BurnAdmin() (common.Address, error) {
	return _MasterChefV2.Contract.BurnAdmin(&_MasterChefV2.CallOpts)
}

// BurnAdmin is a free data retrieval call binding the contract method 0x81bdf98c.
//
// Solidity: function burnAdmin() view returns(address)
func (_MasterChefV2 *MasterChefV2CallerSession) BurnAdmin() (common.Address, error) {
	return _MasterChefV2.Contract.BurnAdmin(&_MasterChefV2.CallOpts)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x1e9b828b.
//
// Solidity: function cakePerBlock(bool _isRegular) view returns(uint256 amount)
func (_MasterChefV2 *MasterChefV2Caller) CakePerBlock(opts *bind.CallOpts, _isRegular bool) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "cakePerBlock", _isRegular)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakePerBlock is a free data retrieval call binding the contract method 0x1e9b828b.
//
// Solidity: function cakePerBlock(bool _isRegular) view returns(uint256 amount)
func (_MasterChefV2 *MasterChefV2Session) CakePerBlock(_isRegular bool) (*big.Int, error) {
	return _MasterChefV2.Contract.CakePerBlock(&_MasterChefV2.CallOpts, _isRegular)
}

// CakePerBlock is a free data retrieval call binding the contract method 0x1e9b828b.
//
// Solidity: function cakePerBlock(bool _isRegular) view returns(uint256 amount)
func (_MasterChefV2 *MasterChefV2CallerSession) CakePerBlock(_isRegular bool) (*big.Int, error) {
	return _MasterChefV2.Contract.CakePerBlock(&_MasterChefV2.CallOpts, _isRegular)
}

// CakePerBlockToBurn is a free data retrieval call binding the contract method 0x9dcc1b5f.
//
// Solidity: function cakePerBlockToBurn() view returns(uint256 amount)
func (_MasterChefV2 *MasterChefV2Caller) CakePerBlockToBurn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "cakePerBlockToBurn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakePerBlockToBurn is a free data retrieval call binding the contract method 0x9dcc1b5f.
//
// Solidity: function cakePerBlockToBurn() view returns(uint256 amount)
func (_MasterChefV2 *MasterChefV2Session) CakePerBlockToBurn() (*big.Int, error) {
	return _MasterChefV2.Contract.CakePerBlockToBurn(&_MasterChefV2.CallOpts)
}

// CakePerBlockToBurn is a free data retrieval call binding the contract method 0x9dcc1b5f.
//
// Solidity: function cakePerBlockToBurn() view returns(uint256 amount)
func (_MasterChefV2 *MasterChefV2CallerSession) CakePerBlockToBurn() (*big.Int, error) {
	return _MasterChefV2.Contract.CakePerBlockToBurn(&_MasterChefV2.CallOpts)
}

// CakeRateToBurn is a free data retrieval call binding the contract method 0xe0f91f6c.
//
// Solidity: function cakeRateToBurn() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) CakeRateToBurn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "cakeRateToBurn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakeRateToBurn is a free data retrieval call binding the contract method 0xe0f91f6c.
//
// Solidity: function cakeRateToBurn() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) CakeRateToBurn() (*big.Int, error) {
	return _MasterChefV2.Contract.CakeRateToBurn(&_MasterChefV2.CallOpts)
}

// CakeRateToBurn is a free data retrieval call binding the contract method 0xe0f91f6c.
//
// Solidity: function cakeRateToBurn() view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) CakeRateToBurn() (*big.Int, error) {
	return _MasterChefV2.Contract.CakeRateToBurn(&_MasterChefV2.CallOpts)
}

// CakeRateToRegularFarm is a free data retrieval call binding the contract method 0xaa47bc8e.
//
// Solidity: function cakeRateToRegularFarm() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) CakeRateToRegularFarm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "cakeRateToRegularFarm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakeRateToRegularFarm is a free data retrieval call binding the contract method 0xaa47bc8e.
//
// Solidity: function cakeRateToRegularFarm() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) CakeRateToRegularFarm() (*big.Int, error) {
	return _MasterChefV2.Contract.CakeRateToRegularFarm(&_MasterChefV2.CallOpts)
}

// CakeRateToRegularFarm is a free data retrieval call binding the contract method 0xaa47bc8e.
//
// Solidity: function cakeRateToRegularFarm() view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) CakeRateToRegularFarm() (*big.Int, error) {
	return _MasterChefV2.Contract.CakeRateToRegularFarm(&_MasterChefV2.CallOpts)
}

// CakeRateToSpecialFarm is a free data retrieval call binding the contract method 0x1ce06d57.
//
// Solidity: function cakeRateToSpecialFarm() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) CakeRateToSpecialFarm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "cakeRateToSpecialFarm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CakeRateToSpecialFarm is a free data retrieval call binding the contract method 0x1ce06d57.
//
// Solidity: function cakeRateToSpecialFarm() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) CakeRateToSpecialFarm() (*big.Int, error) {
	return _MasterChefV2.Contract.CakeRateToSpecialFarm(&_MasterChefV2.CallOpts)
}

// CakeRateToSpecialFarm is a free data retrieval call binding the contract method 0x1ce06d57.
//
// Solidity: function cakeRateToSpecialFarm() view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) CakeRateToSpecialFarm() (*big.Int, error) {
	return _MasterChefV2.Contract.CakeRateToSpecialFarm(&_MasterChefV2.CallOpts)
}

// GetBoostValue is a free data retrieval call binding the contract method 0x2c2c6983.
//
// Solidity: function getBoostValue(address _user, uint256 _pid) view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) GetBoostValue(opts *bind.CallOpts, _user common.Address, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "getBoostValue", _user, _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBoostValue is a free data retrieval call binding the contract method 0x2c2c6983.
//
// Solidity: function getBoostValue(address _user, uint256 _pid) view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) GetBoostValue(_user common.Address, _pid *big.Int) (*big.Int, error) {
	return _MasterChefV2.Contract.GetBoostValue(&_MasterChefV2.CallOpts, _user, _pid)
}

// GetBoostValue is a free data retrieval call binding the contract method 0x2c2c6983.
//
// Solidity: function getBoostValue(address _user, uint256 _pid) view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) GetBoostValue(_user common.Address, _pid *big.Int) (*big.Int, error) {
	return _MasterChefV2.Contract.GetBoostValue(&_MasterChefV2.CallOpts, _user, _pid)
}

// LastBurnedBlock is a free data retrieval call binding the contract method 0x78db4c34.
//
// Solidity: function lastBurnedBlock() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) LastBurnedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "lastBurnedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBurnedBlock is a free data retrieval call binding the contract method 0x78db4c34.
//
// Solidity: function lastBurnedBlock() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) LastBurnedBlock() (*big.Int, error) {
	return _MasterChefV2.Contract.LastBurnedBlock(&_MasterChefV2.CallOpts)
}

// LastBurnedBlock is a free data retrieval call binding the contract method 0x78db4c34.
//
// Solidity: function lastBurnedBlock() view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) LastBurnedBlock() (*big.Int, error) {
	return _MasterChefV2.Contract.LastBurnedBlock(&_MasterChefV2.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_MasterChefV2 *MasterChefV2Caller) LpToken(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "lpToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_MasterChefV2 *MasterChefV2Session) LpToken(arg0 *big.Int) (common.Address, error) {
	return _MasterChefV2.Contract.LpToken(&_MasterChefV2.CallOpts, arg0)
}

// LpToken is a free data retrieval call binding the contract method 0x78ed5d1f.
//
// Solidity: function lpToken(uint256 ) view returns(address)
func (_MasterChefV2 *MasterChefV2CallerSession) LpToken(arg0 *big.Int) (common.Address, error) {
	return _MasterChefV2.Contract.LpToken(&_MasterChefV2.CallOpts, arg0)
}

// MaxBoostValue is a free data retrieval call binding the contract method 0x1fa2cfba.
//
// Solidity: function maxBoostValue() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) MaxBoostValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "maxBoostValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBoostValue is a free data retrieval call binding the contract method 0x1fa2cfba.
//
// Solidity: function maxBoostValue() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) MaxBoostValue() (*big.Int, error) {
	return _MasterChefV2.Contract.MaxBoostValue(&_MasterChefV2.CallOpts)
}

// MaxBoostValue is a free data retrieval call binding the contract method 0x1fa2cfba.
//
// Solidity: function maxBoostValue() view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) MaxBoostValue() (*big.Int, error) {
	return _MasterChefV2.Contract.MaxBoostValue(&_MasterChefV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefV2 *MasterChefV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefV2 *MasterChefV2Session) Owner() (common.Address, error) {
	return _MasterChefV2.Contract.Owner(&_MasterChefV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MasterChefV2 *MasterChefV2CallerSession) Owner() (common.Address, error) {
	return _MasterChefV2.Contract.Owner(&_MasterChefV2.CallOpts)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) PendingCake(opts *bind.CallOpts, _pid *big.Int, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "pendingCake", _pid, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _MasterChefV2.Contract.PendingCake(&_MasterChefV2.CallOpts, _pid, _user)
}

// PendingCake is a free data retrieval call binding the contract method 0x1175a1dd.
//
// Solidity: function pendingCake(uint256 _pid, address _user) view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) PendingCake(_pid *big.Int, _user common.Address) (*big.Int, error) {
	return _MasterChefV2.Contract.PendingCake(&_MasterChefV2.CallOpts, _pid, _user)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint256 accCakePerShare, uint256 lastRewardBlock, uint256 allocPoint, uint256 totalBoostedShare, bool isRegular)
func (_MasterChefV2 *MasterChefV2Caller) PoolInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	AccCakePerShare   *big.Int
	LastRewardBlock   *big.Int
	AllocPoint        *big.Int
	TotalBoostedShare *big.Int
	IsRegular         bool
}, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "poolInfo", arg0)

	outstruct := new(struct {
		AccCakePerShare   *big.Int
		LastRewardBlock   *big.Int
		AllocPoint        *big.Int
		TotalBoostedShare *big.Int
		IsRegular         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AccCakePerShare = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AllocPoint = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalBoostedShare = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsRegular = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint256 accCakePerShare, uint256 lastRewardBlock, uint256 allocPoint, uint256 totalBoostedShare, bool isRegular)
func (_MasterChefV2 *MasterChefV2Session) PoolInfo(arg0 *big.Int) (struct {
	AccCakePerShare   *big.Int
	LastRewardBlock   *big.Int
	AllocPoint        *big.Int
	TotalBoostedShare *big.Int
	IsRegular         bool
}, error) {
	return _MasterChefV2.Contract.PoolInfo(&_MasterChefV2.CallOpts, arg0)
}

// PoolInfo is a free data retrieval call binding the contract method 0x1526fe27.
//
// Solidity: function poolInfo(uint256 ) view returns(uint256 accCakePerShare, uint256 lastRewardBlock, uint256 allocPoint, uint256 totalBoostedShare, bool isRegular)
func (_MasterChefV2 *MasterChefV2CallerSession) PoolInfo(arg0 *big.Int) (struct {
	AccCakePerShare   *big.Int
	LastRewardBlock   *big.Int
	AllocPoint        *big.Int
	TotalBoostedShare *big.Int
	IsRegular         bool
}, error) {
	return _MasterChefV2.Contract.PoolInfo(&_MasterChefV2.CallOpts, arg0)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_MasterChefV2 *MasterChefV2Caller) PoolLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "poolLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_MasterChefV2 *MasterChefV2Session) PoolLength() (*big.Int, error) {
	return _MasterChefV2.Contract.PoolLength(&_MasterChefV2.CallOpts)
}

// PoolLength is a free data retrieval call binding the contract method 0x081e3eda.
//
// Solidity: function poolLength() view returns(uint256 pools)
func (_MasterChefV2 *MasterChefV2CallerSession) PoolLength() (*big.Int, error) {
	return _MasterChefV2.Contract.PoolLength(&_MasterChefV2.CallOpts)
}

// TotalRegularAllocPoint is a free data retrieval call binding the contract method 0xc40d337b.
//
// Solidity: function totalRegularAllocPoint() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) TotalRegularAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "totalRegularAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRegularAllocPoint is a free data retrieval call binding the contract method 0xc40d337b.
//
// Solidity: function totalRegularAllocPoint() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) TotalRegularAllocPoint() (*big.Int, error) {
	return _MasterChefV2.Contract.TotalRegularAllocPoint(&_MasterChefV2.CallOpts)
}

// TotalRegularAllocPoint is a free data retrieval call binding the contract method 0xc40d337b.
//
// Solidity: function totalRegularAllocPoint() view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) TotalRegularAllocPoint() (*big.Int, error) {
	return _MasterChefV2.Contract.TotalRegularAllocPoint(&_MasterChefV2.CallOpts)
}

// TotalSpecialAllocPoint is a free data retrieval call binding the contract method 0x99d7e84a.
//
// Solidity: function totalSpecialAllocPoint() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Caller) TotalSpecialAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "totalSpecialAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSpecialAllocPoint is a free data retrieval call binding the contract method 0x99d7e84a.
//
// Solidity: function totalSpecialAllocPoint() view returns(uint256)
func (_MasterChefV2 *MasterChefV2Session) TotalSpecialAllocPoint() (*big.Int, error) {
	return _MasterChefV2.Contract.TotalSpecialAllocPoint(&_MasterChefV2.CallOpts)
}

// TotalSpecialAllocPoint is a free data retrieval call binding the contract method 0x99d7e84a.
//
// Solidity: function totalSpecialAllocPoint() view returns(uint256)
func (_MasterChefV2 *MasterChefV2CallerSession) TotalSpecialAllocPoint() (*big.Int, error) {
	return _MasterChefV2.Contract.TotalSpecialAllocPoint(&_MasterChefV2.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_MasterChefV2 *MasterChefV2Caller) UserInfo(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "userInfo", arg0, arg1)

	outstruct := new(struct {
		Amount     *big.Int
		RewardDebt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RewardDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_MasterChefV2 *MasterChefV2Session) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _MasterChefV2.Contract.UserInfo(&_MasterChefV2.CallOpts, arg0, arg1)
}

// UserInfo is a free data retrieval call binding the contract method 0x93f1a40b.
//
// Solidity: function userInfo(uint256 , address ) view returns(uint256 amount, uint256 rewardDebt)
func (_MasterChefV2 *MasterChefV2CallerSession) UserInfo(arg0 *big.Int, arg1 common.Address) (struct {
	Amount     *big.Int
	RewardDebt *big.Int
}, error) {
	return _MasterChefV2.Contract.UserInfo(&_MasterChefV2.CallOpts, arg0, arg1)
}

// WhiteListMapping is a free data retrieval call binding the contract method 0x3a7f4e63.
//
// Solidity: function whiteListMapping(address ) view returns(bool)
func (_MasterChefV2 *MasterChefV2Caller) WhiteListMapping(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MasterChefV2.contract.Call(opts, &out, "whiteListMapping", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteListMapping is a free data retrieval call binding the contract method 0x3a7f4e63.
//
// Solidity: function whiteListMapping(address ) view returns(bool)
func (_MasterChefV2 *MasterChefV2Session) WhiteListMapping(arg0 common.Address) (bool, error) {
	return _MasterChefV2.Contract.WhiteListMapping(&_MasterChefV2.CallOpts, arg0)
}

// WhiteListMapping is a free data retrieval call binding the contract method 0x3a7f4e63.
//
// Solidity: function whiteListMapping(address ) view returns(bool)
func (_MasterChefV2 *MasterChefV2CallerSession) WhiteListMapping(arg0 common.Address) (bool, error) {
	return _MasterChefV2.Contract.WhiteListMapping(&_MasterChefV2.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _isRegular, bool _withUpdate) returns()
func (_MasterChefV2 *MasterChefV2Transactor) Add(opts *bind.TransactOpts, _allocPoint *big.Int, _lpToken common.Address, _isRegular bool, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "add", _allocPoint, _lpToken, _isRegular, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _isRegular, bool _withUpdate) returns()
func (_MasterChefV2 *MasterChefV2Session) Add(_allocPoint *big.Int, _lpToken common.Address, _isRegular bool, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Add(&_MasterChefV2.TransactOpts, _allocPoint, _lpToken, _isRegular, _withUpdate)
}

// Add is a paid mutator transaction binding the contract method 0xc507aeaa.
//
// Solidity: function add(uint256 _allocPoint, address _lpToken, bool _isRegular, bool _withUpdate) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) Add(_allocPoint *big.Int, _lpToken common.Address, _isRegular bool, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Add(&_MasterChefV2.TransactOpts, _allocPoint, _lpToken, _isRegular, _withUpdate)
}

// BurnCake is a paid mutator transaction binding the contract method 0x5f4112f9.
//
// Solidity: function burnCake() returns()
func (_MasterChefV2 *MasterChefV2Transactor) BurnCake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "burnCake")
}

// BurnCake is a paid mutator transaction binding the contract method 0x5f4112f9.
//
// Solidity: function burnCake() returns()
func (_MasterChefV2 *MasterChefV2Session) BurnCake() (*types.Transaction, error) {
	return _MasterChefV2.Contract.BurnCake(&_MasterChefV2.TransactOpts)
}

// BurnCake is a paid mutator transaction binding the contract method 0x5f4112f9.
//
// Solidity: function burnCake() returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) BurnCake() (*types.Transaction, error) {
	return _MasterChefV2.Contract.BurnCake(&_MasterChefV2.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChefV2 *MasterChefV2Transactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChefV2 *MasterChefV2Session) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Deposit(&_MasterChefV2.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Deposit(&_MasterChefV2.TransactOpts, _pid, _amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefV2 *MasterChefV2Transactor) EmergencyWithdraw(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "emergencyWithdraw", _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefV2 *MasterChefV2Session) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.EmergencyWithdraw(&_MasterChefV2.TransactOpts, _pid)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0x5312ea8e.
//
// Solidity: function emergencyWithdraw(uint256 _pid) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) EmergencyWithdraw(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.EmergencyWithdraw(&_MasterChefV2.TransactOpts, _pid)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_MasterChefV2 *MasterChefV2Transactor) HarvestFromMasterChef(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "harvestFromMasterChef")
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_MasterChefV2 *MasterChefV2Session) HarvestFromMasterChef() (*types.Transaction, error) {
	return _MasterChefV2.Contract.HarvestFromMasterChef(&_MasterChefV2.TransactOpts)
}

// HarvestFromMasterChef is a paid mutator transaction binding the contract method 0x4f70b15a.
//
// Solidity: function harvestFromMasterChef() returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) HarvestFromMasterChef() (*types.Transaction, error) {
	return _MasterChefV2.Contract.HarvestFromMasterChef(&_MasterChefV2.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_MasterChefV2 *MasterChefV2Transactor) Init(opts *bind.TransactOpts, dummyToken common.Address) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "init", dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_MasterChefV2 *MasterChefV2Session) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Init(&_MasterChefV2.TransactOpts, dummyToken)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address dummyToken) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) Init(dummyToken common.Address) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Init(&_MasterChefV2.TransactOpts, dummyToken)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefV2 *MasterChefV2Transactor) MassUpdatePools(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "massUpdatePools")
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefV2 *MasterChefV2Session) MassUpdatePools() (*types.Transaction, error) {
	return _MasterChefV2.Contract.MassUpdatePools(&_MasterChefV2.TransactOpts)
}

// MassUpdatePools is a paid mutator transaction binding the contract method 0x630b5ba1.
//
// Solidity: function massUpdatePools() returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) MassUpdatePools() (*types.Transaction, error) {
	return _MasterChefV2.Contract.MassUpdatePools(&_MasterChefV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefV2 *MasterChefV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefV2 *MasterChefV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _MasterChefV2.Contract.RenounceOwnership(&_MasterChefV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MasterChefV2.Contract.RenounceOwnership(&_MasterChefV2.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_MasterChefV2 *MasterChefV2Transactor) Set(opts *bind.TransactOpts, _pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "set", _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_MasterChefV2 *MasterChefV2Session) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Set(&_MasterChefV2.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// Set is a paid mutator transaction binding the contract method 0x64482f79.
//
// Solidity: function set(uint256 _pid, uint256 _allocPoint, bool _withUpdate) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) Set(_pid *big.Int, _allocPoint *big.Int, _withUpdate bool) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Set(&_MasterChefV2.TransactOpts, _pid, _allocPoint, _withUpdate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefV2 *MasterChefV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefV2 *MasterChefV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefV2.Contract.TransferOwnership(&_MasterChefV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MasterChefV2.Contract.TransferOwnership(&_MasterChefV2.TransactOpts, newOwner)
}

// UpdateBoostContract is a paid mutator transaction binding the contract method 0xcaa7a10d.
//
// Solidity: function updateBoostContract(address _newBoostContract, uint256 _maxBoostVal) returns()
func (_MasterChefV2 *MasterChefV2Transactor) UpdateBoostContract(opts *bind.TransactOpts, _newBoostContract common.Address, _maxBoostVal *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "updateBoostContract", _newBoostContract, _maxBoostVal)
}

// UpdateBoostContract is a paid mutator transaction binding the contract method 0xcaa7a10d.
//
// Solidity: function updateBoostContract(address _newBoostContract, uint256 _maxBoostVal) returns()
func (_MasterChefV2 *MasterChefV2Session) UpdateBoostContract(_newBoostContract common.Address, _maxBoostVal *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateBoostContract(&_MasterChefV2.TransactOpts, _newBoostContract, _maxBoostVal)
}

// UpdateBoostContract is a paid mutator transaction binding the contract method 0xcaa7a10d.
//
// Solidity: function updateBoostContract(address _newBoostContract, uint256 _maxBoostVal) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) UpdateBoostContract(_newBoostContract common.Address, _maxBoostVal *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateBoostContract(&_MasterChefV2.TransactOpts, _newBoostContract, _maxBoostVal)
}

// UpdateBurnAdmin is a paid mutator transaction binding the contract method 0x0bb844bc.
//
// Solidity: function updateBurnAdmin(address _newAdmin) returns()
func (_MasterChefV2 *MasterChefV2Transactor) UpdateBurnAdmin(opts *bind.TransactOpts, _newAdmin common.Address) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "updateBurnAdmin", _newAdmin)
}

// UpdateBurnAdmin is a paid mutator transaction binding the contract method 0x0bb844bc.
//
// Solidity: function updateBurnAdmin(address _newAdmin) returns()
func (_MasterChefV2 *MasterChefV2Session) UpdateBurnAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateBurnAdmin(&_MasterChefV2.TransactOpts, _newAdmin)
}

// UpdateBurnAdmin is a paid mutator transaction binding the contract method 0x0bb844bc.
//
// Solidity: function updateBurnAdmin(address _newAdmin) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) UpdateBurnAdmin(_newAdmin common.Address) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateBurnAdmin(&_MasterChefV2.TransactOpts, _newAdmin)
}

// UpdateCakeRate is a paid mutator transaction binding the contract method 0xdcd23463.
//
// Solidity: function updateCakeRate(uint256 _burnRate, uint256 _regularFarmRate, uint256 _specialFarmRate) returns()
func (_MasterChefV2 *MasterChefV2Transactor) UpdateCakeRate(opts *bind.TransactOpts, _burnRate *big.Int, _regularFarmRate *big.Int, _specialFarmRate *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "updateCakeRate", _burnRate, _regularFarmRate, _specialFarmRate)
}

// UpdateCakeRate is a paid mutator transaction binding the contract method 0xdcd23463.
//
// Solidity: function updateCakeRate(uint256 _burnRate, uint256 _regularFarmRate, uint256 _specialFarmRate) returns()
func (_MasterChefV2 *MasterChefV2Session) UpdateCakeRate(_burnRate *big.Int, _regularFarmRate *big.Int, _specialFarmRate *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateCakeRate(&_MasterChefV2.TransactOpts, _burnRate, _regularFarmRate, _specialFarmRate)
}

// UpdateCakeRate is a paid mutator transaction binding the contract method 0xdcd23463.
//
// Solidity: function updateCakeRate(uint256 _burnRate, uint256 _regularFarmRate, uint256 _specialFarmRate) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) UpdateCakeRate(_burnRate *big.Int, _regularFarmRate *big.Int, _specialFarmRate *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateCakeRate(&_MasterChefV2.TransactOpts, _burnRate, _regularFarmRate, _specialFarmRate)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns((uint256,uint256,uint256,uint256,bool) pool)
func (_MasterChefV2 *MasterChefV2Transactor) UpdatePool(opts *bind.TransactOpts, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "updatePool", _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns((uint256,uint256,uint256,uint256,bool) pool)
func (_MasterChefV2 *MasterChefV2Session) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdatePool(&_MasterChefV2.TransactOpts, _pid)
}

// UpdatePool is a paid mutator transaction binding the contract method 0x51eb05a6.
//
// Solidity: function updatePool(uint256 _pid) returns((uint256,uint256,uint256,uint256,bool) pool)
func (_MasterChefV2 *MasterChefV2TransactorSession) UpdatePool(_pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdatePool(&_MasterChefV2.TransactOpts, _pid)
}

// UpdateUserBoost is a paid mutator transaction binding the contract method 0x8046cc94.
//
// Solidity: function updateUserBoost(address _user, uint256 _pid) returns()
func (_MasterChefV2 *MasterChefV2Transactor) UpdateUserBoost(opts *bind.TransactOpts, _user common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "updateUserBoost", _user, _pid)
}

// UpdateUserBoost is a paid mutator transaction binding the contract method 0x8046cc94.
//
// Solidity: function updateUserBoost(address _user, uint256 _pid) returns()
func (_MasterChefV2 *MasterChefV2Session) UpdateUserBoost(_user common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateUserBoost(&_MasterChefV2.TransactOpts, _user, _pid)
}

// UpdateUserBoost is a paid mutator transaction binding the contract method 0x8046cc94.
//
// Solidity: function updateUserBoost(address _user, uint256 _pid) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) UpdateUserBoost(_user common.Address, _pid *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateUserBoost(&_MasterChefV2.TransactOpts, _user, _pid)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(address _user, bool _isValid) returns()
func (_MasterChefV2 *MasterChefV2Transactor) UpdateWhiteList(opts *bind.TransactOpts, _user common.Address, _isValid bool) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "updateWhiteList", _user, _isValid)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(address _user, bool _isValid) returns()
func (_MasterChefV2 *MasterChefV2Session) UpdateWhiteList(_user common.Address, _isValid bool) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateWhiteList(&_MasterChefV2.TransactOpts, _user, _isValid)
}

// UpdateWhiteList is a paid mutator transaction binding the contract method 0xac1d0609.
//
// Solidity: function updateWhiteList(address _user, bool _isValid) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) UpdateWhiteList(_user common.Address, _isValid bool) (*types.Transaction, error) {
	return _MasterChefV2.Contract.UpdateWhiteList(&_MasterChefV2.TransactOpts, _user, _isValid)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefV2 *MasterChefV2Transactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefV2 *MasterChefV2Session) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Withdraw(&_MasterChefV2.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_MasterChefV2 *MasterChefV2TransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _MasterChefV2.Contract.Withdraw(&_MasterChefV2.TransactOpts, _pid, _amount)
}

// MasterChefV2AddPoolIterator is returned from FilterAddPool and is used to iterate over the raw logs and unpacked data for AddPool events raised by the MasterChefV2 contract.
type MasterChefV2AddPoolIterator struct {
	Event *MasterChefV2AddPool // Event containing the contract specifics and raw log

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
func (it *MasterChefV2AddPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2AddPool)
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
		it.Event = new(MasterChefV2AddPool)
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
func (it *MasterChefV2AddPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2AddPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2AddPool represents a AddPool event raised by the MasterChefV2 contract.
type MasterChefV2AddPool struct {
	Pid        *big.Int
	AllocPoint *big.Int
	LpToken    common.Address
	IsRegular  bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddPool is a free log retrieval operation binding the contract event 0x18caa0724a26384928efe604ae6ddc99c242548876259770fc88fcb7e719d8fa.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, bool isRegular)
func (_MasterChefV2 *MasterChefV2Filterer) FilterAddPool(opts *bind.FilterOpts, pid []*big.Int, lpToken []common.Address) (*MasterChefV2AddPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "AddPool", pidRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2AddPoolIterator{contract: _MasterChefV2.contract, event: "AddPool", logs: logs, sub: sub}, nil
}

// WatchAddPool is a free log subscription operation binding the contract event 0x18caa0724a26384928efe604ae6ddc99c242548876259770fc88fcb7e719d8fa.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, bool isRegular)
func (_MasterChefV2 *MasterChefV2Filterer) WatchAddPool(opts *bind.WatchOpts, sink chan<- *MasterChefV2AddPool, pid []*big.Int, lpToken []common.Address) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "AddPool", pidRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2AddPool)
				if err := _MasterChefV2.contract.UnpackLog(event, "AddPool", log); err != nil {
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

// ParseAddPool is a log parse operation binding the contract event 0x18caa0724a26384928efe604ae6ddc99c242548876259770fc88fcb7e719d8fa.
//
// Solidity: event AddPool(uint256 indexed pid, uint256 allocPoint, address indexed lpToken, bool isRegular)
func (_MasterChefV2 *MasterChefV2Filterer) ParseAddPool(log types.Log) (*MasterChefV2AddPool, error) {
	event := new(MasterChefV2AddPool)
	if err := _MasterChefV2.contract.UnpackLog(event, "AddPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MasterChefV2 contract.
type MasterChefV2DepositIterator struct {
	Event *MasterChefV2Deposit // Event containing the contract specifics and raw log

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
func (it *MasterChefV2DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2Deposit)
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
		it.Event = new(MasterChefV2Deposit)
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
func (it *MasterChefV2DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2Deposit represents a Deposit event raised by the MasterChefV2 contract.
type MasterChefV2Deposit struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefV2DepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2DepositIterator{contract: _MasterChefV2.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MasterChefV2Deposit, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "Deposit", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2Deposit)
				if err := _MasterChefV2.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) ParseDeposit(log types.Log) (*MasterChefV2Deposit, error) {
	event := new(MasterChefV2Deposit)
	if err := _MasterChefV2.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2EmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the MasterChefV2 contract.
type MasterChefV2EmergencyWithdrawIterator struct {
	Event *MasterChefV2EmergencyWithdraw // Event containing the contract specifics and raw log

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
func (it *MasterChefV2EmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2EmergencyWithdraw)
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
		it.Event = new(MasterChefV2EmergencyWithdraw)
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
func (it *MasterChefV2EmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2EmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2EmergencyWithdraw represents a EmergencyWithdraw event raised by the MasterChefV2 contract.
type MasterChefV2EmergencyWithdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefV2EmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2EmergencyWithdrawIterator{contract: _MasterChefV2.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *MasterChefV2EmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2EmergencyWithdraw)
				if err := _MasterChefV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
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

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) ParseEmergencyWithdraw(log types.Log) (*MasterChefV2EmergencyWithdraw, error) {
	event := new(MasterChefV2EmergencyWithdraw)
	if err := _MasterChefV2.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2InitIterator is returned from FilterInit and is used to iterate over the raw logs and unpacked data for Init events raised by the MasterChefV2 contract.
type MasterChefV2InitIterator struct {
	Event *MasterChefV2Init // Event containing the contract specifics and raw log

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
func (it *MasterChefV2InitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2Init)
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
		it.Event = new(MasterChefV2Init)
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
func (it *MasterChefV2InitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2InitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2Init represents a Init event raised by the MasterChefV2 contract.
type MasterChefV2Init struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInit is a free log retrieval operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_MasterChefV2 *MasterChefV2Filterer) FilterInit(opts *bind.FilterOpts) (*MasterChefV2InitIterator, error) {

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return &MasterChefV2InitIterator{contract: _MasterChefV2.contract, event: "Init", logs: logs, sub: sub}, nil
}

// WatchInit is a free log subscription operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_MasterChefV2 *MasterChefV2Filterer) WatchInit(opts *bind.WatchOpts, sink chan<- *MasterChefV2Init) (event.Subscription, error) {

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "Init")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2Init)
				if err := _MasterChefV2.contract.UnpackLog(event, "Init", log); err != nil {
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

// ParseInit is a log parse operation binding the contract event 0x57a86f7d14ccde89e22870afe839e3011216827daa9b24e18629f0a1e9d6cc14.
//
// Solidity: event Init()
func (_MasterChefV2 *MasterChefV2Filterer) ParseInit(log types.Log) (*MasterChefV2Init, error) {
	event := new(MasterChefV2Init)
	if err := _MasterChefV2.contract.UnpackLog(event, "Init", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MasterChefV2 contract.
type MasterChefV2OwnershipTransferredIterator struct {
	Event *MasterChefV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MasterChefV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2OwnershipTransferred)
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
		it.Event = new(MasterChefV2OwnershipTransferred)
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
func (it *MasterChefV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2OwnershipTransferred represents a OwnershipTransferred event raised by the MasterChefV2 contract.
type MasterChefV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterChefV2 *MasterChefV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MasterChefV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2OwnershipTransferredIterator{contract: _MasterChefV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MasterChefV2 *MasterChefV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MasterChefV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2OwnershipTransferred)
				if err := _MasterChefV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MasterChefV2 *MasterChefV2Filterer) ParseOwnershipTransferred(log types.Log) (*MasterChefV2OwnershipTransferred, error) {
	event := new(MasterChefV2OwnershipTransferred)
	if err := _MasterChefV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2SetPoolIterator is returned from FilterSetPool and is used to iterate over the raw logs and unpacked data for SetPool events raised by the MasterChefV2 contract.
type MasterChefV2SetPoolIterator struct {
	Event *MasterChefV2SetPool // Event containing the contract specifics and raw log

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
func (it *MasterChefV2SetPoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2SetPool)
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
		it.Event = new(MasterChefV2SetPool)
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
func (it *MasterChefV2SetPoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2SetPoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2SetPool represents a SetPool event raised by the MasterChefV2 contract.
type MasterChefV2SetPool struct {
	Pid        *big.Int
	AllocPoint *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetPool is a free log retrieval operation binding the contract event 0xc0cfd54d2de2b55f1e6e108d3ec53ff0a1abe6055401d32c61e9433b747ef9f8.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint)
func (_MasterChefV2 *MasterChefV2Filterer) FilterSetPool(opts *bind.FilterOpts, pid []*big.Int) (*MasterChefV2SetPoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "SetPool", pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2SetPoolIterator{contract: _MasterChefV2.contract, event: "SetPool", logs: logs, sub: sub}, nil
}

// WatchSetPool is a free log subscription operation binding the contract event 0xc0cfd54d2de2b55f1e6e108d3ec53ff0a1abe6055401d32c61e9433b747ef9f8.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint)
func (_MasterChefV2 *MasterChefV2Filterer) WatchSetPool(opts *bind.WatchOpts, sink chan<- *MasterChefV2SetPool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "SetPool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2SetPool)
				if err := _MasterChefV2.contract.UnpackLog(event, "SetPool", log); err != nil {
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

// ParseSetPool is a log parse operation binding the contract event 0xc0cfd54d2de2b55f1e6e108d3ec53ff0a1abe6055401d32c61e9433b747ef9f8.
//
// Solidity: event SetPool(uint256 indexed pid, uint256 allocPoint)
func (_MasterChefV2 *MasterChefV2Filterer) ParseSetPool(log types.Log) (*MasterChefV2SetPool, error) {
	event := new(MasterChefV2SetPool)
	if err := _MasterChefV2.contract.UnpackLog(event, "SetPool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2UpdateBoostContractIterator is returned from FilterUpdateBoostContract and is used to iterate over the raw logs and unpacked data for UpdateBoostContract events raised by the MasterChefV2 contract.
type MasterChefV2UpdateBoostContractIterator struct {
	Event *MasterChefV2UpdateBoostContract // Event containing the contract specifics and raw log

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
func (it *MasterChefV2UpdateBoostContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2UpdateBoostContract)
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
		it.Event = new(MasterChefV2UpdateBoostContract)
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
func (it *MasterChefV2UpdateBoostContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2UpdateBoostContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2UpdateBoostContract represents a UpdateBoostContract event raised by the MasterChefV2 contract.
type MasterChefV2UpdateBoostContract struct {
	BoostContract common.Address
	MaxBoostVal   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateBoostContract is a free log retrieval operation binding the contract event 0x7bb3a7b27bc06fbc4c272d8978c8520bdf5274b5f7683b54e0ee6dba9ad49416.
//
// Solidity: event UpdateBoostContract(address indexed boostContract, uint256 maxBoostVal)
func (_MasterChefV2 *MasterChefV2Filterer) FilterUpdateBoostContract(opts *bind.FilterOpts, boostContract []common.Address) (*MasterChefV2UpdateBoostContractIterator, error) {

	var boostContractRule []interface{}
	for _, boostContractItem := range boostContract {
		boostContractRule = append(boostContractRule, boostContractItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "UpdateBoostContract", boostContractRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2UpdateBoostContractIterator{contract: _MasterChefV2.contract, event: "UpdateBoostContract", logs: logs, sub: sub}, nil
}

// WatchUpdateBoostContract is a free log subscription operation binding the contract event 0x7bb3a7b27bc06fbc4c272d8978c8520bdf5274b5f7683b54e0ee6dba9ad49416.
//
// Solidity: event UpdateBoostContract(address indexed boostContract, uint256 maxBoostVal)
func (_MasterChefV2 *MasterChefV2Filterer) WatchUpdateBoostContract(opts *bind.WatchOpts, sink chan<- *MasterChefV2UpdateBoostContract, boostContract []common.Address) (event.Subscription, error) {

	var boostContractRule []interface{}
	for _, boostContractItem := range boostContract {
		boostContractRule = append(boostContractRule, boostContractItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "UpdateBoostContract", boostContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2UpdateBoostContract)
				if err := _MasterChefV2.contract.UnpackLog(event, "UpdateBoostContract", log); err != nil {
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

// ParseUpdateBoostContract is a log parse operation binding the contract event 0x7bb3a7b27bc06fbc4c272d8978c8520bdf5274b5f7683b54e0ee6dba9ad49416.
//
// Solidity: event UpdateBoostContract(address indexed boostContract, uint256 maxBoostVal)
func (_MasterChefV2 *MasterChefV2Filterer) ParseUpdateBoostContract(log types.Log) (*MasterChefV2UpdateBoostContract, error) {
	event := new(MasterChefV2UpdateBoostContract)
	if err := _MasterChefV2.contract.UnpackLog(event, "UpdateBoostContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2UpdateBurnAdminIterator is returned from FilterUpdateBurnAdmin and is used to iterate over the raw logs and unpacked data for UpdateBurnAdmin events raised by the MasterChefV2 contract.
type MasterChefV2UpdateBurnAdminIterator struct {
	Event *MasterChefV2UpdateBurnAdmin // Event containing the contract specifics and raw log

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
func (it *MasterChefV2UpdateBurnAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2UpdateBurnAdmin)
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
		it.Event = new(MasterChefV2UpdateBurnAdmin)
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
func (it *MasterChefV2UpdateBurnAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2UpdateBurnAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2UpdateBurnAdmin represents a UpdateBurnAdmin event raised by the MasterChefV2 contract.
type MasterChefV2UpdateBurnAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdateBurnAdmin is a free log retrieval operation binding the contract event 0xd146fe330fdddf682413850a35b28edfccd4c4b53cfee802fd24950de5be1dbe.
//
// Solidity: event UpdateBurnAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_MasterChefV2 *MasterChefV2Filterer) FilterUpdateBurnAdmin(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*MasterChefV2UpdateBurnAdminIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "UpdateBurnAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2UpdateBurnAdminIterator{contract: _MasterChefV2.contract, event: "UpdateBurnAdmin", logs: logs, sub: sub}, nil
}

// WatchUpdateBurnAdmin is a free log subscription operation binding the contract event 0xd146fe330fdddf682413850a35b28edfccd4c4b53cfee802fd24950de5be1dbe.
//
// Solidity: event UpdateBurnAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_MasterChefV2 *MasterChefV2Filterer) WatchUpdateBurnAdmin(opts *bind.WatchOpts, sink chan<- *MasterChefV2UpdateBurnAdmin, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "UpdateBurnAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2UpdateBurnAdmin)
				if err := _MasterChefV2.contract.UnpackLog(event, "UpdateBurnAdmin", log); err != nil {
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

// ParseUpdateBurnAdmin is a log parse operation binding the contract event 0xd146fe330fdddf682413850a35b28edfccd4c4b53cfee802fd24950de5be1dbe.
//
// Solidity: event UpdateBurnAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_MasterChefV2 *MasterChefV2Filterer) ParseUpdateBurnAdmin(log types.Log) (*MasterChefV2UpdateBurnAdmin, error) {
	event := new(MasterChefV2UpdateBurnAdmin)
	if err := _MasterChefV2.contract.UnpackLog(event, "UpdateBurnAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2UpdateCakeRateIterator is returned from FilterUpdateCakeRate and is used to iterate over the raw logs and unpacked data for UpdateCakeRate events raised by the MasterChefV2 contract.
type MasterChefV2UpdateCakeRateIterator struct {
	Event *MasterChefV2UpdateCakeRate // Event containing the contract specifics and raw log

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
func (it *MasterChefV2UpdateCakeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2UpdateCakeRate)
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
		it.Event = new(MasterChefV2UpdateCakeRate)
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
func (it *MasterChefV2UpdateCakeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2UpdateCakeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2UpdateCakeRate represents a UpdateCakeRate event raised by the MasterChefV2 contract.
type MasterChefV2UpdateCakeRate struct {
	BurnRate        *big.Int
	RegularFarmRate *big.Int
	SpecialFarmRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateCakeRate is a free log retrieval operation binding the contract event 0xae2d2e7d1ae84564fc72227253ce0f36a007209f7fd5ec414dea80e5af2fb5b0.
//
// Solidity: event UpdateCakeRate(uint256 burnRate, uint256 regularFarmRate, uint256 specialFarmRate)
func (_MasterChefV2 *MasterChefV2Filterer) FilterUpdateCakeRate(opts *bind.FilterOpts) (*MasterChefV2UpdateCakeRateIterator, error) {

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "UpdateCakeRate")
	if err != nil {
		return nil, err
	}
	return &MasterChefV2UpdateCakeRateIterator{contract: _MasterChefV2.contract, event: "UpdateCakeRate", logs: logs, sub: sub}, nil
}

// WatchUpdateCakeRate is a free log subscription operation binding the contract event 0xae2d2e7d1ae84564fc72227253ce0f36a007209f7fd5ec414dea80e5af2fb5b0.
//
// Solidity: event UpdateCakeRate(uint256 burnRate, uint256 regularFarmRate, uint256 specialFarmRate)
func (_MasterChefV2 *MasterChefV2Filterer) WatchUpdateCakeRate(opts *bind.WatchOpts, sink chan<- *MasterChefV2UpdateCakeRate) (event.Subscription, error) {

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "UpdateCakeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2UpdateCakeRate)
				if err := _MasterChefV2.contract.UnpackLog(event, "UpdateCakeRate", log); err != nil {
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

// ParseUpdateCakeRate is a log parse operation binding the contract event 0xae2d2e7d1ae84564fc72227253ce0f36a007209f7fd5ec414dea80e5af2fb5b0.
//
// Solidity: event UpdateCakeRate(uint256 burnRate, uint256 regularFarmRate, uint256 specialFarmRate)
func (_MasterChefV2 *MasterChefV2Filterer) ParseUpdateCakeRate(log types.Log) (*MasterChefV2UpdateCakeRate, error) {
	event := new(MasterChefV2UpdateCakeRate)
	if err := _MasterChefV2.contract.UnpackLog(event, "UpdateCakeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2UpdatePoolIterator is returned from FilterUpdatePool and is used to iterate over the raw logs and unpacked data for UpdatePool events raised by the MasterChefV2 contract.
type MasterChefV2UpdatePoolIterator struct {
	Event *MasterChefV2UpdatePool // Event containing the contract specifics and raw log

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
func (it *MasterChefV2UpdatePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2UpdatePool)
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
		it.Event = new(MasterChefV2UpdatePool)
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
func (it *MasterChefV2UpdatePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2UpdatePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2UpdatePool represents a UpdatePool event raised by the MasterChefV2 contract.
type MasterChefV2UpdatePool struct {
	Pid             *big.Int
	LastRewardBlock *big.Int
	LpSupply        *big.Int
	AccCakePerShare *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatePool is a free log retrieval operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardBlock, uint256 lpSupply, uint256 accCakePerShare)
func (_MasterChefV2 *MasterChefV2Filterer) FilterUpdatePool(opts *bind.FilterOpts, pid []*big.Int) (*MasterChefV2UpdatePoolIterator, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2UpdatePoolIterator{contract: _MasterChefV2.contract, event: "UpdatePool", logs: logs, sub: sub}, nil
}

// WatchUpdatePool is a free log subscription operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardBlock, uint256 lpSupply, uint256 accCakePerShare)
func (_MasterChefV2 *MasterChefV2Filterer) WatchUpdatePool(opts *bind.WatchOpts, sink chan<- *MasterChefV2UpdatePool, pid []*big.Int) (event.Subscription, error) {

	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "UpdatePool", pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2UpdatePool)
				if err := _MasterChefV2.contract.UnpackLog(event, "UpdatePool", log); err != nil {
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

// ParseUpdatePool is a log parse operation binding the contract event 0x3be3541fc42237d611b30329040bfa4569541d156560acdbbae57640d20b8f46.
//
// Solidity: event UpdatePool(uint256 indexed pid, uint256 lastRewardBlock, uint256 lpSupply, uint256 accCakePerShare)
func (_MasterChefV2 *MasterChefV2Filterer) ParseUpdatePool(log types.Log) (*MasterChefV2UpdatePool, error) {
	event := new(MasterChefV2UpdatePool)
	if err := _MasterChefV2.contract.UnpackLog(event, "UpdatePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2UpdateUserBoostIterator is returned from FilterUpdateUserBoost and is used to iterate over the raw logs and unpacked data for UpdateUserBoost events raised by the MasterChefV2 contract.
type MasterChefV2UpdateUserBoostIterator struct {
	Event *MasterChefV2UpdateUserBoost // Event containing the contract specifics and raw log

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
func (it *MasterChefV2UpdateUserBoostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2UpdateUserBoost)
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
		it.Event = new(MasterChefV2UpdateUserBoost)
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
func (it *MasterChefV2UpdateUserBoostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2UpdateUserBoostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2UpdateUserBoost represents a UpdateUserBoost event raised by the MasterChefV2 contract.
type MasterChefV2UpdateUserBoost struct {
	User        common.Address
	OldBoostVal *big.Int
	NewBoostVal *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateUserBoost is a free log retrieval operation binding the contract event 0x7fdfbbd6d5d1be9ee22c079ae0b35c46044fbc9adede5056e28b3dd2f08b6c75.
//
// Solidity: event UpdateUserBoost(address indexed user, uint256 oldBoostVal, uint256 newBoostVal)
func (_MasterChefV2 *MasterChefV2Filterer) FilterUpdateUserBoost(opts *bind.FilterOpts, user []common.Address) (*MasterChefV2UpdateUserBoostIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "UpdateUserBoost", userRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2UpdateUserBoostIterator{contract: _MasterChefV2.contract, event: "UpdateUserBoost", logs: logs, sub: sub}, nil
}

// WatchUpdateUserBoost is a free log subscription operation binding the contract event 0x7fdfbbd6d5d1be9ee22c079ae0b35c46044fbc9adede5056e28b3dd2f08b6c75.
//
// Solidity: event UpdateUserBoost(address indexed user, uint256 oldBoostVal, uint256 newBoostVal)
func (_MasterChefV2 *MasterChefV2Filterer) WatchUpdateUserBoost(opts *bind.WatchOpts, sink chan<- *MasterChefV2UpdateUserBoost, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "UpdateUserBoost", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2UpdateUserBoost)
				if err := _MasterChefV2.contract.UnpackLog(event, "UpdateUserBoost", log); err != nil {
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

// ParseUpdateUserBoost is a log parse operation binding the contract event 0x7fdfbbd6d5d1be9ee22c079ae0b35c46044fbc9adede5056e28b3dd2f08b6c75.
//
// Solidity: event UpdateUserBoost(address indexed user, uint256 oldBoostVal, uint256 newBoostVal)
func (_MasterChefV2 *MasterChefV2Filterer) ParseUpdateUserBoost(log types.Log) (*MasterChefV2UpdateUserBoost, error) {
	event := new(MasterChefV2UpdateUserBoost)
	if err := _MasterChefV2.contract.UnpackLog(event, "UpdateUserBoost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2UpdateWhiteListIterator is returned from FilterUpdateWhiteList and is used to iterate over the raw logs and unpacked data for UpdateWhiteList events raised by the MasterChefV2 contract.
type MasterChefV2UpdateWhiteListIterator struct {
	Event *MasterChefV2UpdateWhiteList // Event containing the contract specifics and raw log

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
func (it *MasterChefV2UpdateWhiteListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2UpdateWhiteList)
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
		it.Event = new(MasterChefV2UpdateWhiteList)
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
func (it *MasterChefV2UpdateWhiteListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2UpdateWhiteListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2UpdateWhiteList represents a UpdateWhiteList event raised by the MasterChefV2 contract.
type MasterChefV2UpdateWhiteList struct {
	User    common.Address
	IsValid bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateWhiteList is a free log retrieval operation binding the contract event 0xc551bbb22d0406dbfb8b6b7740cc521bcf44e1106029cf899c19b6a8e4c99d51.
//
// Solidity: event UpdateWhiteList(address indexed user, bool isValid)
func (_MasterChefV2 *MasterChefV2Filterer) FilterUpdateWhiteList(opts *bind.FilterOpts, user []common.Address) (*MasterChefV2UpdateWhiteListIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "UpdateWhiteList", userRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2UpdateWhiteListIterator{contract: _MasterChefV2.contract, event: "UpdateWhiteList", logs: logs, sub: sub}, nil
}

// WatchUpdateWhiteList is a free log subscription operation binding the contract event 0xc551bbb22d0406dbfb8b6b7740cc521bcf44e1106029cf899c19b6a8e4c99d51.
//
// Solidity: event UpdateWhiteList(address indexed user, bool isValid)
func (_MasterChefV2 *MasterChefV2Filterer) WatchUpdateWhiteList(opts *bind.WatchOpts, sink chan<- *MasterChefV2UpdateWhiteList, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "UpdateWhiteList", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2UpdateWhiteList)
				if err := _MasterChefV2.contract.UnpackLog(event, "UpdateWhiteList", log); err != nil {
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

// ParseUpdateWhiteList is a log parse operation binding the contract event 0xc551bbb22d0406dbfb8b6b7740cc521bcf44e1106029cf899c19b6a8e4c99d51.
//
// Solidity: event UpdateWhiteList(address indexed user, bool isValid)
func (_MasterChefV2 *MasterChefV2Filterer) ParseUpdateWhiteList(log types.Log) (*MasterChefV2UpdateWhiteList, error) {
	event := new(MasterChefV2UpdateWhiteList)
	if err := _MasterChefV2.contract.UnpackLog(event, "UpdateWhiteList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MasterChefV2WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the MasterChefV2 contract.
type MasterChefV2WithdrawIterator struct {
	Event *MasterChefV2Withdraw // Event containing the contract specifics and raw log

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
func (it *MasterChefV2WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MasterChefV2Withdraw)
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
		it.Event = new(MasterChefV2Withdraw)
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
func (it *MasterChefV2WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MasterChefV2WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MasterChefV2Withdraw represents a Withdraw event raised by the MasterChefV2 contract.
type MasterChefV2Withdraw struct {
	User   common.Address
	Pid    *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*MasterChefV2WithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.FilterLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &MasterChefV2WithdrawIterator{contract: _MasterChefV2.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *MasterChefV2Withdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _MasterChefV2.contract.WatchLogs(opts, "Withdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MasterChefV2Withdraw)
				if err := _MasterChefV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed pid, uint256 amount)
func (_MasterChefV2 *MasterChefV2Filterer) ParseWithdraw(log types.Log) (*MasterChefV2Withdraw, error) {
	event := new(MasterChefV2Withdraw)
	if err := _MasterChefV2.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
