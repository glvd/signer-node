// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package video

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VideoABI is the input ABI used to generate the binding from.
const VideoABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"idx\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"internalType\":\"structvideoStorage.Video\",\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"replaceHot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"internalType\":\"structvideoStorage.Video\",\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"addHot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"}],\"name\":\"retrieval\",\"outputs\":[{\"internalType\":\"string[4]\",\"name\":\"\",\"type\":\"string[4]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"idx\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"internalType\":\"structvideoStorage.Video\",\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"replaceNew\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"internalType\":\"structvideoStorage.Video\",\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"replace\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"}],\"name\":\"getNewVideos\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"internalType\":\"structvideoStorage.Video[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVideoCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotVideos\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"internalType\":\"structvideoStorage.Video\",\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"add\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"internalType\":\"structvideoStorage.Video\",\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"addNew\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"videosCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHotVideos\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"bangumi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cover\",\"type\":\"string\"}],\"internalType\":\"structvideoStorage.Video[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Video is an auto generated Go binding around an Ethereum contract.
type Video struct {
	VideoCaller     // Read-only binding to the contract
	VideoTransactor // Write-only binding to the contract
	VideoFilterer   // Log filterer for contract events
}

// VideoCaller is an auto generated read-only Go binding around an Ethereum contract.
type VideoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VideoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VideoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VideoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VideoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VideoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VideoSession struct {
	Contract     *Video            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VideoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VideoCallerSession struct {
	Contract *VideoCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VideoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VideoTransactorSession struct {
	Contract     *VideoTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VideoRaw is an auto generated low-level Go binding around an Ethereum contract.
type VideoRaw struct {
	Contract *Video // Generic contract binding to access the raw methods on
}

// VideoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VideoCallerRaw struct {
	Contract *VideoCaller // Generic read-only contract binding to access the raw methods on
}

// VideoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VideoTransactorRaw struct {
	Contract *VideoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVideo creates a new instance of Video, bound to a specific deployed contract.
func NewVideo(address common.Address, backend bind.ContractBackend) (*Video, error) {
	contract, err := bindVideo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Video{VideoCaller: VideoCaller{contract: contract}, VideoTransactor: VideoTransactor{contract: contract}, VideoFilterer: VideoFilterer{contract: contract}}, nil
}

// NewVideoCaller creates a new read-only instance of Video, bound to a specific deployed contract.
func NewVideoCaller(address common.Address, caller bind.ContractCaller) (*VideoCaller, error) {
	contract, err := bindVideo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VideoCaller{contract: contract}, nil
}

// NewVideoTransactor creates a new write-only instance of Video, bound to a specific deployed contract.
func NewVideoTransactor(address common.Address, transactor bind.ContractTransactor) (*VideoTransactor, error) {
	contract, err := bindVideo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VideoTransactor{contract: contract}, nil
}

// NewVideoFilterer creates a new log filterer instance of Video, bound to a specific deployed contract.
func NewVideoFilterer(address common.Address, filterer bind.ContractFilterer) (*VideoFilterer, error) {
	contract, err := bindVideo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VideoFilterer{contract: contract}, nil
}

// bindVideo binds a generic wrapper to an already deployed contract.
func bindVideo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VideoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Video *VideoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Video.Contract.VideoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Video *VideoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Video.Contract.VideoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Video *VideoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Video.Contract.VideoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Video *VideoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Video.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Video *VideoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Video.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Video *VideoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Video.Contract.contract.Transact(opts, method, params...)
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Bangumi string
	Name    string
	Hash    string
	Cover   string
}

// GetHotVideos is a free data retrieval call binding the contract method 0xf07b6b79.
//
// Solidity: function getHotVideos() constant returns([]Struct0)
func (_Video *VideoCaller) GetHotVideos(opts *bind.CallOpts) ([]Struct0, error) {
	var (
		ret0 = new([]Struct0)
	)
	out := ret0
	err := _Video.contract.Call(opts, out, "getHotVideos")
	return *ret0, err
}

// GetHotVideos is a free data retrieval call binding the contract method 0xf07b6b79.
//
// Solidity: function getHotVideos() constant returns([]Struct0)
func (_Video *VideoSession) GetHotVideos() ([]Struct0, error) {
	return _Video.Contract.GetHotVideos(&_Video.CallOpts)
}

// GetHotVideos is a free data retrieval call binding the contract method 0xf07b6b79.
//
// Solidity: function getHotVideos() constant returns([]Struct0)
func (_Video *VideoCallerSession) GetHotVideos() ([]Struct0, error) {
	return _Video.Contract.GetHotVideos(&_Video.CallOpts)
}

// GetNewVideos is a free data retrieval call binding the contract method 0xa20ddc12.
//
// Solidity: function getNewVideos(string date) constant returns([]Struct0)
func (_Video *VideoCaller) GetNewVideos(opts *bind.CallOpts, date string) ([]Struct0, error) {
	var (
		ret0 = new([]Struct0)
	)
	out := ret0
	err := _Video.contract.Call(opts, out, "getNewVideos", date)
	return *ret0, err
}

// GetNewVideos is a free data retrieval call binding the contract method 0xa20ddc12.
//
// Solidity: function getNewVideos(string date) constant returns([]Struct0)
func (_Video *VideoSession) GetNewVideos(date string) ([]Struct0, error) {
	return _Video.Contract.GetNewVideos(&_Video.CallOpts, date)
}

// GetNewVideos is a free data retrieval call binding the contract method 0xa20ddc12.
//
// Solidity: function getNewVideos(string date) constant returns([]Struct0)
func (_Video *VideoCallerSession) GetNewVideos(date string) ([]Struct0, error) {
	return _Video.Contract.GetNewVideos(&_Video.CallOpts, date)
}

// GetVideoCount is a free data retrieval call binding the contract method 0xa286c0b9.
//
// Solidity: function getVideoCount() constant returns(uint32)
func (_Video *VideoCaller) GetVideoCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Video.contract.Call(opts, out, "getVideoCount")
	return *ret0, err
}

// GetVideoCount is a free data retrieval call binding the contract method 0xa286c0b9.
//
// Solidity: function getVideoCount() constant returns(uint32)
func (_Video *VideoSession) GetVideoCount() (uint32, error) {
	return _Video.Contract.GetVideoCount(&_Video.CallOpts)
}

// GetVideoCount is a free data retrieval call binding the contract method 0xa286c0b9.
//
// Solidity: function getVideoCount() constant returns(uint32)
func (_Video *VideoCallerSession) GetVideoCount() (uint32, error) {
	return _Video.Contract.GetVideoCount(&_Video.CallOpts)
}

// HotVideos is a free data retrieval call binding the contract method 0xb332a7bc.
//
// Solidity: function hotVideos(uint256 ) constant returns(string bangumi, string name, string hash, string cover)
func (_Video *VideoCaller) HotVideos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Bangumi string
	Name    string
	Hash    string
	Cover   string
}, error) {
	ret := new(struct {
		Bangumi string
		Name    string
		Hash    string
		Cover   string
	})
	out := ret
	err := _Video.contract.Call(opts, out, "hotVideos", arg0)
	return *ret, err
}

// HotVideos is a free data retrieval call binding the contract method 0xb332a7bc.
//
// Solidity: function hotVideos(uint256 ) constant returns(string bangumi, string name, string hash, string cover)
func (_Video *VideoSession) HotVideos(arg0 *big.Int) (struct {
	Bangumi string
	Name    string
	Hash    string
	Cover   string
}, error) {
	return _Video.Contract.HotVideos(&_Video.CallOpts, arg0)
}

// HotVideos is a free data retrieval call binding the contract method 0xb332a7bc.
//
// Solidity: function hotVideos(uint256 ) constant returns(string bangumi, string name, string hash, string cover)
func (_Video *VideoCallerSession) HotVideos(arg0 *big.Int) (struct {
	Bangumi string
	Name    string
	Hash    string
	Cover   string
}, error) {
	return _Video.Contract.HotVideos(&_Video.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Video *VideoCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Video.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Video *VideoSession) IsOwner() (bool, error) {
	return _Video.Contract.IsOwner(&_Video.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Video *VideoCallerSession) IsOwner() (bool, error) {
	return _Video.Contract.IsOwner(&_Video.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Video *VideoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Video.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Video *VideoSession) Owner() (common.Address, error) {
	return _Video.Contract.Owner(&_Video.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Video *VideoCallerSession) Owner() (common.Address, error) {
	return _Video.Contract.Owner(&_Video.CallOpts)
}

// Retrieval is a free data retrieval call binding the contract method 0x4478bacf.
//
// Solidity: function retrieval(string bangumi) constant returns(string[4])
func (_Video *VideoCaller) Retrieval(opts *bind.CallOpts, bangumi string) ([4]string, error) {
	var (
		ret0 = new([4]string)
	)
	out := ret0
	err := _Video.contract.Call(opts, out, "retrieval", bangumi)
	return *ret0, err
}

// Retrieval is a free data retrieval call binding the contract method 0x4478bacf.
//
// Solidity: function retrieval(string bangumi) constant returns(string[4])
func (_Video *VideoSession) Retrieval(bangumi string) ([4]string, error) {
	return _Video.Contract.Retrieval(&_Video.CallOpts, bangumi)
}

// Retrieval is a free data retrieval call binding the contract method 0x4478bacf.
//
// Solidity: function retrieval(string bangumi) constant returns(string[4])
func (_Video *VideoCallerSession) Retrieval(bangumi string) ([4]string, error) {
	return _Video.Contract.Retrieval(&_Video.CallOpts, bangumi)
}

// VideosCount is a free data retrieval call binding the contract method 0xef08523b.
//
// Solidity: function videosCount() constant returns(uint32)
func (_Video *VideoCaller) VideosCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Video.contract.Call(opts, out, "videosCount")
	return *ret0, err
}

// VideosCount is a free data retrieval call binding the contract method 0xef08523b.
//
// Solidity: function videosCount() constant returns(uint32)
func (_Video *VideoSession) VideosCount() (uint32, error) {
	return _Video.Contract.VideosCount(&_Video.CallOpts)
}

// VideosCount is a free data retrieval call binding the contract method 0xef08523b.
//
// Solidity: function videosCount() constant returns(uint32)
func (_Video *VideoCallerSession) VideosCount() (uint32, error) {
	return _Video.Contract.VideosCount(&_Video.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xbd18cabf.
//
// Solidity: function add(Struct0 v) returns(bool)
func (_Video *VideoTransactor) Add(opts *bind.TransactOpts, v Struct0) (*types.Transaction, error) {
	return _Video.contract.Transact(opts, "add", v)
}

// Add is a paid mutator transaction binding the contract method 0xbd18cabf.
//
// Solidity: function add(Struct0 v) returns(bool)
func (_Video *VideoSession) Add(v Struct0) (*types.Transaction, error) {
	return _Video.Contract.Add(&_Video.TransactOpts, v)
}

// Add is a paid mutator transaction binding the contract method 0xbd18cabf.
//
// Solidity: function add(Struct0 v) returns(bool)
func (_Video *VideoTransactorSession) Add(v Struct0) (*types.Transaction, error) {
	return _Video.Contract.Add(&_Video.TransactOpts, v)
}

// AddHot is a paid mutator transaction binding the contract method 0x41556a30.
//
// Solidity: function addHot(Struct0 v) returns(bool)
func (_Video *VideoTransactor) AddHot(opts *bind.TransactOpts, v Struct0) (*types.Transaction, error) {
	return _Video.contract.Transact(opts, "addHot", v)
}

// AddHot is a paid mutator transaction binding the contract method 0x41556a30.
//
// Solidity: function addHot(Struct0 v) returns(bool)
func (_Video *VideoSession) AddHot(v Struct0) (*types.Transaction, error) {
	return _Video.Contract.AddHot(&_Video.TransactOpts, v)
}

// AddHot is a paid mutator transaction binding the contract method 0x41556a30.
//
// Solidity: function addHot(Struct0 v) returns(bool)
func (_Video *VideoTransactorSession) AddHot(v Struct0) (*types.Transaction, error) {
	return _Video.Contract.AddHot(&_Video.TransactOpts, v)
}

// AddNew is a paid mutator transaction binding the contract method 0xdcf38949.
//
// Solidity: function addNew(string date, Struct0 v) returns(bool)
func (_Video *VideoTransactor) AddNew(opts *bind.TransactOpts, date string, v Struct0) (*types.Transaction, error) {
	return _Video.contract.Transact(opts, "addNew", date, v)
}

// AddNew is a paid mutator transaction binding the contract method 0xdcf38949.
//
// Solidity: function addNew(string date, Struct0 v) returns(bool)
func (_Video *VideoSession) AddNew(date string, v Struct0) (*types.Transaction, error) {
	return _Video.Contract.AddNew(&_Video.TransactOpts, date, v)
}

// AddNew is a paid mutator transaction binding the contract method 0xdcf38949.
//
// Solidity: function addNew(string date, Struct0 v) returns(bool)
func (_Video *VideoTransactorSession) AddNew(date string, v Struct0) (*types.Transaction, error) {
	return _Video.Contract.AddNew(&_Video.TransactOpts, date, v)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Video *VideoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Video.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Video *VideoSession) RenounceOwnership() (*types.Transaction, error) {
	return _Video.Contract.RenounceOwnership(&_Video.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Video *VideoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Video.Contract.RenounceOwnership(&_Video.TransactOpts)
}

// Replace is a paid mutator transaction binding the contract method 0x5df6dcb0.
//
// Solidity: function replace(Struct0 v) returns(bool)
func (_Video *VideoTransactor) Replace(opts *bind.TransactOpts, v Struct0) (*types.Transaction, error) {
	return _Video.contract.Transact(opts, "replace", v)
}

// Replace is a paid mutator transaction binding the contract method 0x5df6dcb0.
//
// Solidity: function replace(Struct0 v) returns(bool)
func (_Video *VideoSession) Replace(v Struct0) (*types.Transaction, error) {
	return _Video.Contract.Replace(&_Video.TransactOpts, v)
}

// Replace is a paid mutator transaction binding the contract method 0x5df6dcb0.
//
// Solidity: function replace(Struct0 v) returns(bool)
func (_Video *VideoTransactorSession) Replace(v Struct0) (*types.Transaction, error) {
	return _Video.Contract.Replace(&_Video.TransactOpts, v)
}

// ReplaceHot is a paid mutator transaction binding the contract method 0x3ca01096.
//
// Solidity: function replaceHot(uint32 idx, Struct0 v) returns()
func (_Video *VideoTransactor) ReplaceHot(opts *bind.TransactOpts, idx uint32, v Struct0) (*types.Transaction, error) {
	return _Video.contract.Transact(opts, "replaceHot", idx, v)
}

// ReplaceHot is a paid mutator transaction binding the contract method 0x3ca01096.
//
// Solidity: function replaceHot(uint32 idx, Struct0 v) returns()
func (_Video *VideoSession) ReplaceHot(idx uint32, v Struct0) (*types.Transaction, error) {
	return _Video.Contract.ReplaceHot(&_Video.TransactOpts, idx, v)
}

// ReplaceHot is a paid mutator transaction binding the contract method 0x3ca01096.
//
// Solidity: function replaceHot(uint32 idx, Struct0 v) returns()
func (_Video *VideoTransactorSession) ReplaceHot(idx uint32, v Struct0) (*types.Transaction, error) {
	return _Video.Contract.ReplaceHot(&_Video.TransactOpts, idx, v)
}

// ReplaceNew is a paid mutator transaction binding the contract method 0x4d1d90fe.
//
// Solidity: function replaceNew(string date, uint32 idx, Struct0 v) returns()
func (_Video *VideoTransactor) ReplaceNew(opts *bind.TransactOpts, date string, idx uint32, v Struct0) (*types.Transaction, error) {
	return _Video.contract.Transact(opts, "replaceNew", date, idx, v)
}

// ReplaceNew is a paid mutator transaction binding the contract method 0x4d1d90fe.
//
// Solidity: function replaceNew(string date, uint32 idx, Struct0 v) returns()
func (_Video *VideoSession) ReplaceNew(date string, idx uint32, v Struct0) (*types.Transaction, error) {
	return _Video.Contract.ReplaceNew(&_Video.TransactOpts, date, idx, v)
}

// ReplaceNew is a paid mutator transaction binding the contract method 0x4d1d90fe.
//
// Solidity: function replaceNew(string date, uint32 idx, Struct0 v) returns()
func (_Video *VideoTransactorSession) ReplaceNew(date string, idx uint32, v Struct0) (*types.Transaction, error) {
	return _Video.Contract.ReplaceNew(&_Video.TransactOpts, date, idx, v)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Video *VideoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Video.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Video *VideoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Video.Contract.TransferOwnership(&_Video.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Video *VideoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Video.Contract.TransferOwnership(&_Video.TransactOpts, newOwner)
}

// VideoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Video contract.
type VideoOwnershipTransferredIterator struct {
	Event *VideoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VideoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VideoOwnershipTransferred)
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
		it.Event = new(VideoOwnershipTransferred)
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
func (it *VideoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VideoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VideoOwnershipTransferred represents a OwnershipTransferred event raised by the Video contract.
type VideoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Video *VideoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VideoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Video.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VideoOwnershipTransferredIterator{contract: _Video.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Video *VideoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VideoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Video.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VideoOwnershipTransferred)
				if err := _Video.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Video *VideoFilterer) ParseOwnershipTransferred(log types.Log) (*VideoOwnershipTransferred, error) {
	event := new(VideoOwnershipTransferred)
	if err := _Video.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
