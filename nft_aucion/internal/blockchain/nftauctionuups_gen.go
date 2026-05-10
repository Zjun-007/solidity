// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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

// NFTAuctionUUPSMetaData contains all meta data concerning the NFTAuctionUUPS contract.
var NFTAuctionUUPSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctionData\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"highestBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"settled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"cancelled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bidAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimUnsoldNFT\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createAuction\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextAuctionId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftToken2AuctionId\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onERC721Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AuctionCancelled\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionCreated\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionEnded\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ActiveAuctionExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadySettled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionInactive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AuctionNotEnded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BidTooLow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotCancel\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ETHPaySellerFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ETHRefundFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTimeRange\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoBidsToSettle\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotBiddingWindow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSeller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSellerOrHasBids\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWinner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StartInPast\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// NFTAuctionUUPSABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTAuctionUUPSMetaData.ABI instead.
var NFTAuctionUUPSABI = NFTAuctionUUPSMetaData.ABI

// NFTAuctionUUPS is an auto generated Go binding around an Ethereum contract.
type NFTAuctionUUPS struct {
	NFTAuctionUUPSCaller     // Read-only binding to the contract
	NFTAuctionUUPSTransactor // Write-only binding to the contract
	NFTAuctionUUPSFilterer   // Log filterer for contract events
}

// NFTAuctionUUPSCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTAuctionUUPSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionUUPSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTAuctionUUPSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionUUPSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTAuctionUUPSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTAuctionUUPSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTAuctionUUPSSession struct {
	Contract     *NFTAuctionUUPS   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTAuctionUUPSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTAuctionUUPSCallerSession struct {
	Contract *NFTAuctionUUPSCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NFTAuctionUUPSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTAuctionUUPSTransactorSession struct {
	Contract     *NFTAuctionUUPSTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NFTAuctionUUPSRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTAuctionUUPSRaw struct {
	Contract *NFTAuctionUUPS // Generic contract binding to access the raw methods on
}

// NFTAuctionUUPSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTAuctionUUPSCallerRaw struct {
	Contract *NFTAuctionUUPSCaller // Generic read-only contract binding to access the raw methods on
}

// NFTAuctionUUPSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTAuctionUUPSTransactorRaw struct {
	Contract *NFTAuctionUUPSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTAuctionUUPS creates a new instance of NFTAuctionUUPS, bound to a specific deployed contract.
func NewNFTAuctionUUPS(address common.Address, backend bind.ContractBackend) (*NFTAuctionUUPS, error) {
	contract, err := bindNFTAuctionUUPS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPS{NFTAuctionUUPSCaller: NFTAuctionUUPSCaller{contract: contract}, NFTAuctionUUPSTransactor: NFTAuctionUUPSTransactor{contract: contract}, NFTAuctionUUPSFilterer: NFTAuctionUUPSFilterer{contract: contract}}, nil
}

// NewNFTAuctionUUPSCaller creates a new read-only instance of NFTAuctionUUPS, bound to a specific deployed contract.
func NewNFTAuctionUUPSCaller(address common.Address, caller bind.ContractCaller) (*NFTAuctionUUPSCaller, error) {
	contract, err := bindNFTAuctionUUPS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSCaller{contract: contract}, nil
}

// NewNFTAuctionUUPSTransactor creates a new write-only instance of NFTAuctionUUPS, bound to a specific deployed contract.
func NewNFTAuctionUUPSTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTAuctionUUPSTransactor, error) {
	contract, err := bindNFTAuctionUUPS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSTransactor{contract: contract}, nil
}

// NewNFTAuctionUUPSFilterer creates a new log filterer instance of NFTAuctionUUPS, bound to a specific deployed contract.
func NewNFTAuctionUUPSFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTAuctionUUPSFilterer, error) {
	contract, err := bindNFTAuctionUUPS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSFilterer{contract: contract}, nil
}

// bindNFTAuctionUUPS binds a generic wrapper to an already deployed contract.
func bindNFTAuctionUUPS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NFTAuctionUUPSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTAuctionUUPS *NFTAuctionUUPSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTAuctionUUPS.Contract.NFTAuctionUUPSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTAuctionUUPS *NFTAuctionUUPSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.NFTAuctionUUPSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTAuctionUUPS *NFTAuctionUUPSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.NFTAuctionUUPSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTAuctionUUPS *NFTAuctionUUPSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTAuctionUUPS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuctionUUPS *NFTAuctionUUPSCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTAuctionUUPS.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NFTAuctionUUPS.Contract.UPGRADEINTERFACEVERSION(&_NFTAuctionUUPS.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_NFTAuctionUUPS *NFTAuctionUUPSCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _NFTAuctionUUPS.Contract.UPGRADEINTERFACEVERSION(&_NFTAuctionUUPS.CallOpts)
}

// AuctionData is a free data retrieval call binding the contract method 0x55fc62d2.
//
// Solidity: function auctionData(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime, address highestBidder, uint256 highestBid, bool settled, bool cancelled)
func (_NFTAuctionUUPS *NFTAuctionUUPSCaller) AuctionData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartPrice    *big.Int
	StartTime     *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	Settled       bool
	Cancelled     bool
}, error) {
	var out []interface{}
	err := _NFTAuctionUUPS.contract.Call(opts, &out, "auctionData", arg0)

	outstruct := new(struct {
		Seller        common.Address
		NftContract   common.Address
		TokenId       *big.Int
		StartPrice    *big.Int
		StartTime     *big.Int
		EndTime       *big.Int
		HighestBidder common.Address
		HighestBid    *big.Int
		Settled       bool
		Cancelled     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Settled = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.Cancelled = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// AuctionData is a free data retrieval call binding the contract method 0x55fc62d2.
//
// Solidity: function auctionData(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime, address highestBidder, uint256 highestBid, bool settled, bool cancelled)
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) AuctionData(arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartPrice    *big.Int
	StartTime     *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	Settled       bool
	Cancelled     bool
}, error) {
	return _NFTAuctionUUPS.Contract.AuctionData(&_NFTAuctionUUPS.CallOpts, arg0)
}

// AuctionData is a free data retrieval call binding the contract method 0x55fc62d2.
//
// Solidity: function auctionData(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime, address highestBidder, uint256 highestBid, bool settled, bool cancelled)
func (_NFTAuctionUUPS *NFTAuctionUUPSCallerSession) AuctionData(arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartPrice    *big.Int
	StartTime     *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	Settled       bool
	Cancelled     bool
}, error) {
	return _NFTAuctionUUPS.Contract.AuctionData(&_NFTAuctionUUPS.CallOpts, arg0)
}

// NextAuctionId is a free data retrieval call binding the contract method 0xfc528482.
//
// Solidity: function nextAuctionId() view returns(uint256)
func (_NFTAuctionUUPS *NFTAuctionUUPSCaller) NextAuctionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuctionUUPS.contract.Call(opts, &out, "nextAuctionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextAuctionId is a free data retrieval call binding the contract method 0xfc528482.
//
// Solidity: function nextAuctionId() view returns(uint256)
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) NextAuctionId() (*big.Int, error) {
	return _NFTAuctionUUPS.Contract.NextAuctionId(&_NFTAuctionUUPS.CallOpts)
}

// NextAuctionId is a free data retrieval call binding the contract method 0xfc528482.
//
// Solidity: function nextAuctionId() view returns(uint256)
func (_NFTAuctionUUPS *NFTAuctionUUPSCallerSession) NextAuctionId() (*big.Int, error) {
	return _NFTAuctionUUPS.Contract.NextAuctionId(&_NFTAuctionUUPS.CallOpts)
}

// NftToken2AuctionId is a free data retrieval call binding the contract method 0x980fda47.
//
// Solidity: function nftToken2AuctionId(address , uint256 ) view returns(uint256)
func (_NFTAuctionUUPS *NFTAuctionUUPSCaller) NftToken2AuctionId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTAuctionUUPS.contract.Call(opts, &out, "nftToken2AuctionId", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftToken2AuctionId is a free data retrieval call binding the contract method 0x980fda47.
//
// Solidity: function nftToken2AuctionId(address , uint256 ) view returns(uint256)
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) NftToken2AuctionId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NFTAuctionUUPS.Contract.NftToken2AuctionId(&_NFTAuctionUUPS.CallOpts, arg0, arg1)
}

// NftToken2AuctionId is a free data retrieval call binding the contract method 0x980fda47.
//
// Solidity: function nftToken2AuctionId(address , uint256 ) view returns(uint256)
func (_NFTAuctionUUPS *NFTAuctionUUPSCallerSession) NftToken2AuctionId(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NFTAuctionUUPS.Contract.NftToken2AuctionId(&_NFTAuctionUUPS.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuctionUUPS *NFTAuctionUUPSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTAuctionUUPS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) Owner() (common.Address, error) {
	return _NFTAuctionUUPS.Contract.Owner(&_NFTAuctionUUPS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTAuctionUUPS *NFTAuctionUUPSCallerSession) Owner() (common.Address, error) {
	return _NFTAuctionUUPS.Contract.Owner(&_NFTAuctionUUPS.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuctionUUPS *NFTAuctionUUPSCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTAuctionUUPS.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) ProxiableUUID() ([32]byte, error) {
	return _NFTAuctionUUPS.Contract.ProxiableUUID(&_NFTAuctionUUPS.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_NFTAuctionUUPS *NFTAuctionUUPSCallerSession) ProxiableUUID() ([32]byte, error) {
	return _NFTAuctionUUPS.Contract.ProxiableUUID(&_NFTAuctionUUPS.CallOpts)
}

// BidAuction is a paid mutator transaction binding the contract method 0x64a7d7c7.
//
// Solidity: function bidAuction(uint256 auctionId) payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) BidAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "bidAuction", auctionId)
}

// BidAuction is a paid mutator transaction binding the contract method 0x64a7d7c7.
//
// Solidity: function bidAuction(uint256 auctionId) payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) BidAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.BidAuction(&_NFTAuctionUUPS.TransactOpts, auctionId)
}

// BidAuction is a paid mutator transaction binding the contract method 0x64a7d7c7.
//
// Solidity: function bidAuction(uint256 auctionId) payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) BidAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.BidAuction(&_NFTAuctionUUPS.TransactOpts, auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) CancelAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "cancelAuction", auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) CancelAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.CancelAuction(&_NFTAuctionUUPS.TransactOpts, auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) CancelAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.CancelAuction(&_NFTAuctionUUPS.TransactOpts, auctionId)
}

// ClaimUnsoldNFT is a paid mutator transaction binding the contract method 0xa792bf64.
//
// Solidity: function claimUnsoldNFT(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) ClaimUnsoldNFT(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "claimUnsoldNFT", auctionId)
}

// ClaimUnsoldNFT is a paid mutator transaction binding the contract method 0xa792bf64.
//
// Solidity: function claimUnsoldNFT(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) ClaimUnsoldNFT(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.ClaimUnsoldNFT(&_NFTAuctionUUPS.TransactOpts, auctionId)
}

// ClaimUnsoldNFT is a paid mutator transaction binding the contract method 0xa792bf64.
//
// Solidity: function claimUnsoldNFT(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) ClaimUnsoldNFT(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.ClaimUnsoldNFT(&_NFTAuctionUUPS.TransactOpts, auctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime) returns(uint256 auctionId)
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) CreateAuction(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, startPrice *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "createAuction", nftContract, tokenId, startPrice, startTime, endTime)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime) returns(uint256 auctionId)
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) CreateAuction(nftContract common.Address, tokenId *big.Int, startPrice *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.CreateAuction(&_NFTAuctionUUPS.TransactOpts, nftContract, tokenId, startPrice, startTime, endTime)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime) returns(uint256 auctionId)
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) CreateAuction(nftContract common.Address, tokenId *big.Int, startPrice *big.Int, startTime *big.Int, endTime *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.CreateAuction(&_NFTAuctionUUPS.TransactOpts, nftContract, tokenId, startPrice, startTime, endTime)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) EndAuction(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "endAuction", auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.EndAuction(&_NFTAuctionUUPS.TransactOpts, auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 auctionId) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.EndAuction(&_NFTAuctionUUPS.TransactOpts, auctionId)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.Initialize(&_NFTAuctionUUPS.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.Initialize(&_NFTAuctionUUPS.TransactOpts, initialOwner)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.OnERC721Received(&_NFTAuctionUUPS.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.OnERC721Received(&_NFTAuctionUUPS.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.RenounceOwnership(&_NFTAuctionUUPS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.RenounceOwnership(&_NFTAuctionUUPS.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.TransferOwnership(&_NFTAuctionUUPS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.TransferOwnership(&_NFTAuctionUUPS.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.UpgradeToAndCall(&_NFTAuctionUUPS.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.UpgradeToAndCall(&_NFTAuctionUUPS.TransactOpts, newImplementation, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTAuctionUUPS.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSSession) Receive() (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.Receive(&_NFTAuctionUUPS.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NFTAuctionUUPS *NFTAuctionUUPSTransactorSession) Receive() (*types.Transaction, error) {
	return _NFTAuctionUUPS.Contract.Receive(&_NFTAuctionUUPS.TransactOpts)
}

// NFTAuctionUUPSAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSAuctionCancelledIterator struct {
	Event *NFTAuctionUUPSAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *NFTAuctionUUPSAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionUUPSAuctionCancelled)
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
		it.Event = new(NFTAuctionUUPSAuctionCancelled)
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
func (it *NFTAuctionUUPSAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionUUPSAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionUUPSAuctionCancelled represents a AuctionCancelled event raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSAuctionCancelled struct {
	AuctionId *big.Int
	Seller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x10ac9f0bb365b5d22d7bec500408692f23fdf83eadfec71615ef88b4c1134f0e.
//
// Solidity: event AuctionCancelled(uint256 indexed auctionId, address indexed seller)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) FilterAuctionCancelled(opts *bind.FilterOpts, auctionId []*big.Int, seller []common.Address) (*NFTAuctionUUPSAuctionCancelledIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.FilterLogs(opts, "AuctionCancelled", auctionIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSAuctionCancelledIterator{contract: _NFTAuctionUUPS.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x10ac9f0bb365b5d22d7bec500408692f23fdf83eadfec71615ef88b4c1134f0e.
//
// Solidity: event AuctionCancelled(uint256 indexed auctionId, address indexed seller)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *NFTAuctionUUPSAuctionCancelled, auctionId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.WatchLogs(opts, "AuctionCancelled", auctionIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionUUPSAuctionCancelled)
				if err := _NFTAuctionUUPS.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x10ac9f0bb365b5d22d7bec500408692f23fdf83eadfec71615ef88b4c1134f0e.
//
// Solidity: event AuctionCancelled(uint256 indexed auctionId, address indexed seller)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) ParseAuctionCancelled(log types.Log) (*NFTAuctionUUPSAuctionCancelled, error) {
	event := new(NFTAuctionUUPSAuctionCancelled)
	if err := _NFTAuctionUUPS.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionUUPSAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSAuctionCreatedIterator struct {
	Event *NFTAuctionUUPSAuctionCreated // Event containing the contract specifics and raw log

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
func (it *NFTAuctionUUPSAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionUUPSAuctionCreated)
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
		it.Event = new(NFTAuctionUUPSAuctionCreated)
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
func (it *NFTAuctionUUPSAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionUUPSAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionUUPSAuctionCreated represents a AuctionCreated event raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSAuctionCreated struct {
	AuctionId   *big.Int
	Seller      common.Address
	NftContract common.Address
	TokenId     *big.Int
	StartPrice  *big.Int
	StartTime   *big.Int
	EndTime     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xcaf0ae751fb2b122e8718bf7c0d4b7584d1418a853a4d0cdaba45418d3da138b.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) FilterAuctionCreated(opts *bind.FilterOpts, auctionId []*big.Int, seller []common.Address, nftContract []common.Address) (*NFTAuctionUUPSAuctionCreatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.FilterLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSAuctionCreatedIterator{contract: _NFTAuctionUUPS.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xcaf0ae751fb2b122e8718bf7c0d4b7584d1418a853a4d0cdaba45418d3da138b.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *NFTAuctionUUPSAuctionCreated, auctionId []*big.Int, seller []common.Address, nftContract []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.WatchLogs(opts, "AuctionCreated", auctionIdRule, sellerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionUUPSAuctionCreated)
				if err := _NFTAuctionUUPS.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xcaf0ae751fb2b122e8718bf7c0d4b7584d1418a853a4d0cdaba45418d3da138b.
//
// Solidity: event AuctionCreated(uint256 indexed auctionId, address indexed seller, address indexed nftContract, uint256 tokenId, uint256 startPrice, uint256 startTime, uint256 endTime)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) ParseAuctionCreated(log types.Log) (*NFTAuctionUUPSAuctionCreated, error) {
	event := new(NFTAuctionUUPSAuctionCreated)
	if err := _NFTAuctionUUPS.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionUUPSAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSAuctionEndedIterator struct {
	Event *NFTAuctionUUPSAuctionEnded // Event containing the contract specifics and raw log

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
func (it *NFTAuctionUUPSAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionUUPSAuctionEnded)
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
		it.Event = new(NFTAuctionUUPSAuctionEnded)
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
func (it *NFTAuctionUUPSAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionUUPSAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionUUPSAuctionEnded represents a AuctionEnded event raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSAuctionEnded struct {
	AuctionId *big.Int
	Winner    common.Address
	Amount    *big.Int
	Seller    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xbb6764412c29916bdf4a5c6fe6b1c079de35682160b2289928ce003ab459a749.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 amount, address indexed seller)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) FilterAuctionEnded(opts *bind.FilterOpts, auctionId []*big.Int, winner []common.Address, seller []common.Address) (*NFTAuctionUUPSAuctionEndedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.FilterLogs(opts, "AuctionEnded", auctionIdRule, winnerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSAuctionEndedIterator{contract: _NFTAuctionUUPS.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xbb6764412c29916bdf4a5c6fe6b1c079de35682160b2289928ce003ab459a749.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 amount, address indexed seller)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *NFTAuctionUUPSAuctionEnded, auctionId []*big.Int, winner []common.Address, seller []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.WatchLogs(opts, "AuctionEnded", auctionIdRule, winnerRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionUUPSAuctionEnded)
				if err := _NFTAuctionUUPS.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xbb6764412c29916bdf4a5c6fe6b1c079de35682160b2289928ce003ab459a749.
//
// Solidity: event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 amount, address indexed seller)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) ParseAuctionEnded(log types.Log) (*NFTAuctionUUPSAuctionEnded, error) {
	event := new(NFTAuctionUUPSAuctionEnded)
	if err := _NFTAuctionUUPS.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionUUPSBidAuctionIterator is returned from FilterBidAuction and is used to iterate over the raw logs and unpacked data for BidAuction events raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSBidAuctionIterator struct {
	Event *NFTAuctionUUPSBidAuction // Event containing the contract specifics and raw log

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
func (it *NFTAuctionUUPSBidAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionUUPSBidAuction)
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
		it.Event = new(NFTAuctionUUPSBidAuction)
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
func (it *NFTAuctionUUPSBidAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionUUPSBidAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionUUPSBidAuction represents a BidAuction event raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSBidAuction struct {
	AuctionId *big.Int
	Bidder    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBidAuction is a free log retrieval operation binding the contract event 0x014de1dd43627559120d21162d1154078a340acf94db1e6c44d3f1933ff03df9.
//
// Solidity: event BidAuction(uint256 indexed auctionId, address indexed bidder, uint256 amount)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) FilterBidAuction(opts *bind.FilterOpts, auctionId []*big.Int, bidder []common.Address) (*NFTAuctionUUPSBidAuctionIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.FilterLogs(opts, "BidAuction", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSBidAuctionIterator{contract: _NFTAuctionUUPS.contract, event: "BidAuction", logs: logs, sub: sub}, nil
}

// WatchBidAuction is a free log subscription operation binding the contract event 0x014de1dd43627559120d21162d1154078a340acf94db1e6c44d3f1933ff03df9.
//
// Solidity: event BidAuction(uint256 indexed auctionId, address indexed bidder, uint256 amount)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) WatchBidAuction(opts *bind.WatchOpts, sink chan<- *NFTAuctionUUPSBidAuction, auctionId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.WatchLogs(opts, "BidAuction", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionUUPSBidAuction)
				if err := _NFTAuctionUUPS.contract.UnpackLog(event, "BidAuction", log); err != nil {
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

// ParseBidAuction is a log parse operation binding the contract event 0x014de1dd43627559120d21162d1154078a340acf94db1e6c44d3f1933ff03df9.
//
// Solidity: event BidAuction(uint256 indexed auctionId, address indexed bidder, uint256 amount)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) ParseBidAuction(log types.Log) (*NFTAuctionUUPSBidAuction, error) {
	event := new(NFTAuctionUUPSBidAuction)
	if err := _NFTAuctionUUPS.contract.UnpackLog(event, "BidAuction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionUUPSInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSInitializedIterator struct {
	Event *NFTAuctionUUPSInitialized // Event containing the contract specifics and raw log

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
func (it *NFTAuctionUUPSInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionUUPSInitialized)
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
		it.Event = new(NFTAuctionUUPSInitialized)
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
func (it *NFTAuctionUUPSInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionUUPSInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionUUPSInitialized represents a Initialized event raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) FilterInitialized(opts *bind.FilterOpts) (*NFTAuctionUUPSInitializedIterator, error) {

	logs, sub, err := _NFTAuctionUUPS.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSInitializedIterator{contract: _NFTAuctionUUPS.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NFTAuctionUUPSInitialized) (event.Subscription, error) {

	logs, sub, err := _NFTAuctionUUPS.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionUUPSInitialized)
				if err := _NFTAuctionUUPS.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) ParseInitialized(log types.Log) (*NFTAuctionUUPSInitialized, error) {
	event := new(NFTAuctionUUPSInitialized)
	if err := _NFTAuctionUUPS.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionUUPSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSOwnershipTransferredIterator struct {
	Event *NFTAuctionUUPSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NFTAuctionUUPSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionUUPSOwnershipTransferred)
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
		it.Event = new(NFTAuctionUUPSOwnershipTransferred)
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
func (it *NFTAuctionUUPSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionUUPSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionUUPSOwnershipTransferred represents a OwnershipTransferred event raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTAuctionUUPSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSOwnershipTransferredIterator{contract: _NFTAuctionUUPS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTAuctionUUPSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionUUPSOwnershipTransferred)
				if err := _NFTAuctionUUPS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) ParseOwnershipTransferred(log types.Log) (*NFTAuctionUUPSOwnershipTransferred, error) {
	event := new(NFTAuctionUUPSOwnershipTransferred)
	if err := _NFTAuctionUUPS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTAuctionUUPSUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSUpgradedIterator struct {
	Event *NFTAuctionUUPSUpgraded // Event containing the contract specifics and raw log

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
func (it *NFTAuctionUUPSUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTAuctionUUPSUpgraded)
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
		it.Event = new(NFTAuctionUUPSUpgraded)
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
func (it *NFTAuctionUUPSUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTAuctionUUPSUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTAuctionUUPSUpgraded represents a Upgraded event raised by the NFTAuctionUUPS contract.
type NFTAuctionUUPSUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*NFTAuctionUUPSUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &NFTAuctionUUPSUpgradedIterator{contract: _NFTAuctionUUPS.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *NFTAuctionUUPSUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _NFTAuctionUUPS.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTAuctionUUPSUpgraded)
				if err := _NFTAuctionUUPS.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_NFTAuctionUUPS *NFTAuctionUUPSFilterer) ParseUpgraded(log types.Log) (*NFTAuctionUUPSUpgraded, error) {
	event := new(NFTAuctionUUPSUpgraded)
	if err := _NFTAuctionUUPS.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
