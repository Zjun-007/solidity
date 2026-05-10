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

// FeedAuctionMetaData contains all meta data concerning the FeedAuction contract.
var FeedAuctionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_PRICE_DELAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"biddingToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"highestBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ended\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bidERC20\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bidETH\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createAuction\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"biddingToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBid\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBidUSD\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeeds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAggregatorV3Interface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuctionCreated\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"biddingToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionEnded\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"winner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BidPlaced\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"biddingToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceFeedSet\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"feed\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawn\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// FeedAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use FeedAuctionMetaData.ABI instead.
var FeedAuctionABI = FeedAuctionMetaData.ABI

// FeedAuction is an auto generated Go binding around an Ethereum contract.
type FeedAuction struct {
	FeedAuctionCaller     // Read-only binding to the contract
	FeedAuctionTransactor // Write-only binding to the contract
	FeedAuctionFilterer   // Log filterer for contract events
}

// FeedAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeedAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeedAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeedAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeedAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeedAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeedAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeedAuctionSession struct {
	Contract     *FeedAuction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeedAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeedAuctionCallerSession struct {
	Contract *FeedAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FeedAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeedAuctionTransactorSession struct {
	Contract     *FeedAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FeedAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeedAuctionRaw struct {
	Contract *FeedAuction // Generic contract binding to access the raw methods on
}

// FeedAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeedAuctionCallerRaw struct {
	Contract *FeedAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// FeedAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeedAuctionTransactorRaw struct {
	Contract *FeedAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeedAuction creates a new instance of FeedAuction, bound to a specific deployed contract.
func NewFeedAuction(address common.Address, backend bind.ContractBackend) (*FeedAuction, error) {
	contract, err := bindFeedAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeedAuction{FeedAuctionCaller: FeedAuctionCaller{contract: contract}, FeedAuctionTransactor: FeedAuctionTransactor{contract: contract}, FeedAuctionFilterer: FeedAuctionFilterer{contract: contract}}, nil
}

// NewFeedAuctionCaller creates a new read-only instance of FeedAuction, bound to a specific deployed contract.
func NewFeedAuctionCaller(address common.Address, caller bind.ContractCaller) (*FeedAuctionCaller, error) {
	contract, err := bindFeedAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeedAuctionCaller{contract: contract}, nil
}

// NewFeedAuctionTransactor creates a new write-only instance of FeedAuction, bound to a specific deployed contract.
func NewFeedAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*FeedAuctionTransactor, error) {
	contract, err := bindFeedAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeedAuctionTransactor{contract: contract}, nil
}

// NewFeedAuctionFilterer creates a new log filterer instance of FeedAuction, bound to a specific deployed contract.
func NewFeedAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*FeedAuctionFilterer, error) {
	contract, err := bindFeedAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeedAuctionFilterer{contract: contract}, nil
}

// bindFeedAuction binds a generic wrapper to an already deployed contract.
func bindFeedAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeedAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeedAuction *FeedAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeedAuction.Contract.FeedAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeedAuction *FeedAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeedAuction.Contract.FeedAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeedAuction *FeedAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeedAuction.Contract.FeedAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeedAuction *FeedAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeedAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeedAuction *FeedAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeedAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeedAuction *FeedAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeedAuction.Contract.contract.Transact(opts, method, params...)
}

// MAXPRICEDELAY is a free data retrieval call binding the contract method 0x1f644d6d.
//
// Solidity: function MAX_PRICE_DELAY() view returns(uint256)
func (_FeedAuction *FeedAuctionCaller) MAXPRICEDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeedAuction.contract.Call(opts, &out, "MAX_PRICE_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPRICEDELAY is a free data retrieval call binding the contract method 0x1f644d6d.
//
// Solidity: function MAX_PRICE_DELAY() view returns(uint256)
func (_FeedAuction *FeedAuctionSession) MAXPRICEDELAY() (*big.Int, error) {
	return _FeedAuction.Contract.MAXPRICEDELAY(&_FeedAuction.CallOpts)
}

// MAXPRICEDELAY is a free data retrieval call binding the contract method 0x1f644d6d.
//
// Solidity: function MAX_PRICE_DELAY() view returns(uint256)
func (_FeedAuction *FeedAuctionCallerSession) MAXPRICEDELAY() (*big.Int, error) {
	return _FeedAuction.Contract.MAXPRICEDELAY(&_FeedAuction.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x1edbc5be.
//
// Solidity: function auctions(bytes32 ) view returns(address seller, address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 endTime, address highestBidder, uint256 highestBid, uint256 highestBidUSD, bool ended)
func (_FeedAuction *FeedAuctionCaller) Auctions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	BiddingToken  common.Address
	StartPrice    *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	HighestBidUSD *big.Int
	Ended         bool
}, error) {
	var out []interface{}
	err := _FeedAuction.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		Seller        common.Address
		NftContract   common.Address
		TokenId       *big.Int
		BiddingToken  common.Address
		StartPrice    *big.Int
		EndTime       *big.Int
		HighestBidder common.Address
		HighestBid    *big.Int
		HighestBidUSD *big.Int
		Ended         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BiddingToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.StartPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.HighestBid = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.HighestBidUSD = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x1edbc5be.
//
// Solidity: function auctions(bytes32 ) view returns(address seller, address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 endTime, address highestBidder, uint256 highestBid, uint256 highestBidUSD, bool ended)
func (_FeedAuction *FeedAuctionSession) Auctions(arg0 [32]byte) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	BiddingToken  common.Address
	StartPrice    *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	HighestBidUSD *big.Int
	Ended         bool
}, error) {
	return _FeedAuction.Contract.Auctions(&_FeedAuction.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x1edbc5be.
//
// Solidity: function auctions(bytes32 ) view returns(address seller, address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 endTime, address highestBidder, uint256 highestBid, uint256 highestBidUSD, bool ended)
func (_FeedAuction *FeedAuctionCallerSession) Auctions(arg0 [32]byte) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	BiddingToken  common.Address
	StartPrice    *big.Int
	EndTime       *big.Int
	HighestBidder common.Address
	HighestBid    *big.Int
	HighestBidUSD *big.Int
	Ended         bool
}, error) {
	return _FeedAuction.Contract.Auctions(&_FeedAuction.CallOpts, arg0)
}

// GetBid is a free data retrieval call binding the contract method 0xd0ba5264.
//
// Solidity: function getBid(bytes32 auctionId, address user) view returns(uint256)
func (_FeedAuction *FeedAuctionCaller) GetBid(opts *bind.CallOpts, auctionId [32]byte, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeedAuction.contract.Call(opts, &out, "getBid", auctionId, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBid is a free data retrieval call binding the contract method 0xd0ba5264.
//
// Solidity: function getBid(bytes32 auctionId, address user) view returns(uint256)
func (_FeedAuction *FeedAuctionSession) GetBid(auctionId [32]byte, user common.Address) (*big.Int, error) {
	return _FeedAuction.Contract.GetBid(&_FeedAuction.CallOpts, auctionId, user)
}

// GetBid is a free data retrieval call binding the contract method 0xd0ba5264.
//
// Solidity: function getBid(bytes32 auctionId, address user) view returns(uint256)
func (_FeedAuction *FeedAuctionCallerSession) GetBid(auctionId [32]byte, user common.Address) (*big.Int, error) {
	return _FeedAuction.Contract.GetBid(&_FeedAuction.CallOpts, auctionId, user)
}

// GetBidUSD is a free data retrieval call binding the contract method 0x4f580867.
//
// Solidity: function getBidUSD(bytes32 auctionId, address user) view returns(uint256)
func (_FeedAuction *FeedAuctionCaller) GetBidUSD(opts *bind.CallOpts, auctionId [32]byte, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeedAuction.contract.Call(opts, &out, "getBidUSD", auctionId, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidUSD is a free data retrieval call binding the contract method 0x4f580867.
//
// Solidity: function getBidUSD(bytes32 auctionId, address user) view returns(uint256)
func (_FeedAuction *FeedAuctionSession) GetBidUSD(auctionId [32]byte, user common.Address) (*big.Int, error) {
	return _FeedAuction.Contract.GetBidUSD(&_FeedAuction.CallOpts, auctionId, user)
}

// GetBidUSD is a free data retrieval call binding the contract method 0x4f580867.
//
// Solidity: function getBidUSD(bytes32 auctionId, address user) view returns(uint256)
func (_FeedAuction *FeedAuctionCallerSession) GetBidUSD(auctionId [32]byte, user common.Address) (*big.Int, error) {
	return _FeedAuction.Contract.GetBidUSD(&_FeedAuction.CallOpts, auctionId, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeedAuction *FeedAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeedAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeedAuction *FeedAuctionSession) Owner() (common.Address, error) {
	return _FeedAuction.Contract.Owner(&_FeedAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeedAuction *FeedAuctionCallerSession) Owner() (common.Address, error) {
	return _FeedAuction.Contract.Owner(&_FeedAuction.CallOpts)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_FeedAuction *FeedAuctionCaller) PriceFeeds(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _FeedAuction.contract.Call(opts, &out, "priceFeeds", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_FeedAuction *FeedAuctionSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _FeedAuction.Contract.PriceFeeds(&_FeedAuction.CallOpts, arg0)
}

// PriceFeeds is a free data retrieval call binding the contract method 0x9dcb511a.
//
// Solidity: function priceFeeds(address ) view returns(address)
func (_FeedAuction *FeedAuctionCallerSession) PriceFeeds(arg0 common.Address) (common.Address, error) {
	return _FeedAuction.Contract.PriceFeeds(&_FeedAuction.CallOpts, arg0)
}

// BidERC20 is a paid mutator transaction binding the contract method 0x830d974d.
//
// Solidity: function bidERC20(bytes32 auctionId, uint256 amount) returns()
func (_FeedAuction *FeedAuctionTransactor) BidERC20(opts *bind.TransactOpts, auctionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _FeedAuction.contract.Transact(opts, "bidERC20", auctionId, amount)
}

// BidERC20 is a paid mutator transaction binding the contract method 0x830d974d.
//
// Solidity: function bidERC20(bytes32 auctionId, uint256 amount) returns()
func (_FeedAuction *FeedAuctionSession) BidERC20(auctionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _FeedAuction.Contract.BidERC20(&_FeedAuction.TransactOpts, auctionId, amount)
}

// BidERC20 is a paid mutator transaction binding the contract method 0x830d974d.
//
// Solidity: function bidERC20(bytes32 auctionId, uint256 amount) returns()
func (_FeedAuction *FeedAuctionTransactorSession) BidERC20(auctionId [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _FeedAuction.Contract.BidERC20(&_FeedAuction.TransactOpts, auctionId, amount)
}

// BidETH is a paid mutator transaction binding the contract method 0x09c4b91f.
//
// Solidity: function bidETH(bytes32 auctionId) payable returns()
func (_FeedAuction *FeedAuctionTransactor) BidETH(opts *bind.TransactOpts, auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.contract.Transact(opts, "bidETH", auctionId)
}

// BidETH is a paid mutator transaction binding the contract method 0x09c4b91f.
//
// Solidity: function bidETH(bytes32 auctionId) payable returns()
func (_FeedAuction *FeedAuctionSession) BidETH(auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.Contract.BidETH(&_FeedAuction.TransactOpts, auctionId)
}

// BidETH is a paid mutator transaction binding the contract method 0x09c4b91f.
//
// Solidity: function bidETH(bytes32 auctionId) payable returns()
func (_FeedAuction *FeedAuctionTransactorSession) BidETH(auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.Contract.BidETH(&_FeedAuction.TransactOpts, auctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xf2df5047.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 duration) returns(bytes32 auctionId)
func (_FeedAuction *FeedAuctionTransactor) CreateAuction(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, biddingToken common.Address, startPrice *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _FeedAuction.contract.Transact(opts, "createAuction", nftContract, tokenId, biddingToken, startPrice, duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xf2df5047.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 duration) returns(bytes32 auctionId)
func (_FeedAuction *FeedAuctionSession) CreateAuction(nftContract common.Address, tokenId *big.Int, biddingToken common.Address, startPrice *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _FeedAuction.Contract.CreateAuction(&_FeedAuction.TransactOpts, nftContract, tokenId, biddingToken, startPrice, duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0xf2df5047.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 duration) returns(bytes32 auctionId)
func (_FeedAuction *FeedAuctionTransactorSession) CreateAuction(nftContract common.Address, tokenId *big.Int, biddingToken common.Address, startPrice *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _FeedAuction.Contract.CreateAuction(&_FeedAuction.TransactOpts, nftContract, tokenId, biddingToken, startPrice, duration)
}

// Finalize is a paid mutator transaction binding the contract method 0x92584d80.
//
// Solidity: function finalize(bytes32 auctionId) returns()
func (_FeedAuction *FeedAuctionTransactor) Finalize(opts *bind.TransactOpts, auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.contract.Transact(opts, "finalize", auctionId)
}

// Finalize is a paid mutator transaction binding the contract method 0x92584d80.
//
// Solidity: function finalize(bytes32 auctionId) returns()
func (_FeedAuction *FeedAuctionSession) Finalize(auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.Contract.Finalize(&_FeedAuction.TransactOpts, auctionId)
}

// Finalize is a paid mutator transaction binding the contract method 0x92584d80.
//
// Solidity: function finalize(bytes32 auctionId) returns()
func (_FeedAuction *FeedAuctionTransactorSession) Finalize(auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.Contract.Finalize(&_FeedAuction.TransactOpts, auctionId)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address token, address feed) returns()
func (_FeedAuction *FeedAuctionTransactor) SetPriceFeed(opts *bind.TransactOpts, token common.Address, feed common.Address) (*types.Transaction, error) {
	return _FeedAuction.contract.Transact(opts, "setPriceFeed", token, feed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address token, address feed) returns()
func (_FeedAuction *FeedAuctionSession) SetPriceFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _FeedAuction.Contract.SetPriceFeed(&_FeedAuction.TransactOpts, token, feed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address token, address feed) returns()
func (_FeedAuction *FeedAuctionTransactorSession) SetPriceFeed(token common.Address, feed common.Address) (*types.Transaction, error) {
	return _FeedAuction.Contract.SetPriceFeed(&_FeedAuction.TransactOpts, token, feed)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 auctionId) returns()
func (_FeedAuction *FeedAuctionTransactor) Withdraw(opts *bind.TransactOpts, auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 auctionId) returns()
func (_FeedAuction *FeedAuctionSession) Withdraw(auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.Contract.Withdraw(&_FeedAuction.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 auctionId) returns()
func (_FeedAuction *FeedAuctionTransactorSession) Withdraw(auctionId [32]byte) (*types.Transaction, error) {
	return _FeedAuction.Contract.Withdraw(&_FeedAuction.TransactOpts, auctionId)
}

// FeedAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the FeedAuction contract.
type FeedAuctionAuctionCreatedIterator struct {
	Event *FeedAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *FeedAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeedAuctionAuctionCreated)
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
		it.Event = new(FeedAuctionAuctionCreated)
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
func (it *FeedAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeedAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeedAuctionAuctionCreated represents a AuctionCreated event raised by the FeedAuction contract.
type FeedAuctionAuctionCreated struct {
	AuctionId    [32]byte
	Seller       common.Address
	NftContract  common.Address
	TokenId      *big.Int
	BiddingToken common.Address
	StartPrice   *big.Int
	EndTime      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xe092418694f6ddd70867c5b8e9b49c398bad1f0d784ef7eab3f6e7dd495b522f.
//
// Solidity: event AuctionCreated(bytes32 indexed auctionId, address indexed seller, address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 endTime)
func (_FeedAuction *FeedAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts, auctionId [][32]byte, seller []common.Address) (*FeedAuctionAuctionCreatedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _FeedAuction.contract.FilterLogs(opts, "AuctionCreated", auctionIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &FeedAuctionAuctionCreatedIterator{contract: _FeedAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xe092418694f6ddd70867c5b8e9b49c398bad1f0d784ef7eab3f6e7dd495b522f.
//
// Solidity: event AuctionCreated(bytes32 indexed auctionId, address indexed seller, address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 endTime)
func (_FeedAuction *FeedAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *FeedAuctionAuctionCreated, auctionId [][32]byte, seller []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _FeedAuction.contract.WatchLogs(opts, "AuctionCreated", auctionIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeedAuctionAuctionCreated)
				if err := _FeedAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xe092418694f6ddd70867c5b8e9b49c398bad1f0d784ef7eab3f6e7dd495b522f.
//
// Solidity: event AuctionCreated(bytes32 indexed auctionId, address indexed seller, address nftContract, uint256 tokenId, address biddingToken, uint256 startPrice, uint256 endTime)
func (_FeedAuction *FeedAuctionFilterer) ParseAuctionCreated(log types.Log) (*FeedAuctionAuctionCreated, error) {
	event := new(FeedAuctionAuctionCreated)
	if err := _FeedAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeedAuctionAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the FeedAuction contract.
type FeedAuctionAuctionEndedIterator struct {
	Event *FeedAuctionAuctionEnded // Event containing the contract specifics and raw log

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
func (it *FeedAuctionAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeedAuctionAuctionEnded)
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
		it.Event = new(FeedAuctionAuctionEnded)
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
func (it *FeedAuctionAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeedAuctionAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeedAuctionAuctionEnded represents a AuctionEnded event raised by the FeedAuction contract.
type FeedAuctionAuctionEnded struct {
	AuctionId [32]byte
	Winner    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0x4898413e58dd21ed9179dd98f8c9313f8359aae688bff5403ed56a6677bcfe97.
//
// Solidity: event AuctionEnded(bytes32 indexed auctionId, address winner, uint256 amount)
func (_FeedAuction *FeedAuctionFilterer) FilterAuctionEnded(opts *bind.FilterOpts, auctionId [][32]byte) (*FeedAuctionAuctionEndedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _FeedAuction.contract.FilterLogs(opts, "AuctionEnded", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &FeedAuctionAuctionEndedIterator{contract: _FeedAuction.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0x4898413e58dd21ed9179dd98f8c9313f8359aae688bff5403ed56a6677bcfe97.
//
// Solidity: event AuctionEnded(bytes32 indexed auctionId, address winner, uint256 amount)
func (_FeedAuction *FeedAuctionFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *FeedAuctionAuctionEnded, auctionId [][32]byte) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _FeedAuction.contract.WatchLogs(opts, "AuctionEnded", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeedAuctionAuctionEnded)
				if err := _FeedAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0x4898413e58dd21ed9179dd98f8c9313f8359aae688bff5403ed56a6677bcfe97.
//
// Solidity: event AuctionEnded(bytes32 indexed auctionId, address winner, uint256 amount)
func (_FeedAuction *FeedAuctionFilterer) ParseAuctionEnded(log types.Log) (*FeedAuctionAuctionEnded, error) {
	event := new(FeedAuctionAuctionEnded)
	if err := _FeedAuction.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeedAuctionBidPlacedIterator is returned from FilterBidPlaced and is used to iterate over the raw logs and unpacked data for BidPlaced events raised by the FeedAuction contract.
type FeedAuctionBidPlacedIterator struct {
	Event *FeedAuctionBidPlaced // Event containing the contract specifics and raw log

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
func (it *FeedAuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeedAuctionBidPlaced)
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
		it.Event = new(FeedAuctionBidPlaced)
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
func (it *FeedAuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeedAuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeedAuctionBidPlaced represents a BidPlaced event raised by the FeedAuction contract.
type FeedAuctionBidPlaced struct {
	AuctionId    [32]byte
	Bidder       common.Address
	Amount       *big.Int
	BiddingToken common.Address
	UsdValue     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBidPlaced is a free log retrieval operation binding the contract event 0xc7ca85bf2ee479bcf6ee898e400ce0913ca201a494ae7aa8d4e5d011c09e2243.
//
// Solidity: event BidPlaced(bytes32 indexed auctionId, address indexed bidder, uint256 amount, address biddingToken, uint256 usdValue)
func (_FeedAuction *FeedAuctionFilterer) FilterBidPlaced(opts *bind.FilterOpts, auctionId [][32]byte, bidder []common.Address) (*FeedAuctionBidPlacedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FeedAuction.contract.FilterLogs(opts, "BidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FeedAuctionBidPlacedIterator{contract: _FeedAuction.contract, event: "BidPlaced", logs: logs, sub: sub}, nil
}

// WatchBidPlaced is a free log subscription operation binding the contract event 0xc7ca85bf2ee479bcf6ee898e400ce0913ca201a494ae7aa8d4e5d011c09e2243.
//
// Solidity: event BidPlaced(bytes32 indexed auctionId, address indexed bidder, uint256 amount, address biddingToken, uint256 usdValue)
func (_FeedAuction *FeedAuctionFilterer) WatchBidPlaced(opts *bind.WatchOpts, sink chan<- *FeedAuctionBidPlaced, auctionId [][32]byte, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FeedAuction.contract.WatchLogs(opts, "BidPlaced", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeedAuctionBidPlaced)
				if err := _FeedAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
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

// ParseBidPlaced is a log parse operation binding the contract event 0xc7ca85bf2ee479bcf6ee898e400ce0913ca201a494ae7aa8d4e5d011c09e2243.
//
// Solidity: event BidPlaced(bytes32 indexed auctionId, address indexed bidder, uint256 amount, address biddingToken, uint256 usdValue)
func (_FeedAuction *FeedAuctionFilterer) ParseBidPlaced(log types.Log) (*FeedAuctionBidPlaced, error) {
	event := new(FeedAuctionBidPlaced)
	if err := _FeedAuction.contract.UnpackLog(event, "BidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeedAuctionPriceFeedSetIterator is returned from FilterPriceFeedSet and is used to iterate over the raw logs and unpacked data for PriceFeedSet events raised by the FeedAuction contract.
type FeedAuctionPriceFeedSetIterator struct {
	Event *FeedAuctionPriceFeedSet // Event containing the contract specifics and raw log

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
func (it *FeedAuctionPriceFeedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeedAuctionPriceFeedSet)
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
		it.Event = new(FeedAuctionPriceFeedSet)
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
func (it *FeedAuctionPriceFeedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeedAuctionPriceFeedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeedAuctionPriceFeedSet represents a PriceFeedSet event raised by the FeedAuction contract.
type FeedAuctionPriceFeedSet struct {
	Token common.Address
	Feed  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPriceFeedSet is a free log retrieval operation binding the contract event 0xd2d8394cf7549a5ddbc2ba3dd7b2de8d53c891472d1f2907008ed6a10045fdae.
//
// Solidity: event PriceFeedSet(address indexed token, address indexed feed)
func (_FeedAuction *FeedAuctionFilterer) FilterPriceFeedSet(opts *bind.FilterOpts, token []common.Address, feed []common.Address) (*FeedAuctionPriceFeedSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var feedRule []interface{}
	for _, feedItem := range feed {
		feedRule = append(feedRule, feedItem)
	}

	logs, sub, err := _FeedAuction.contract.FilterLogs(opts, "PriceFeedSet", tokenRule, feedRule)
	if err != nil {
		return nil, err
	}
	return &FeedAuctionPriceFeedSetIterator{contract: _FeedAuction.contract, event: "PriceFeedSet", logs: logs, sub: sub}, nil
}

// WatchPriceFeedSet is a free log subscription operation binding the contract event 0xd2d8394cf7549a5ddbc2ba3dd7b2de8d53c891472d1f2907008ed6a10045fdae.
//
// Solidity: event PriceFeedSet(address indexed token, address indexed feed)
func (_FeedAuction *FeedAuctionFilterer) WatchPriceFeedSet(opts *bind.WatchOpts, sink chan<- *FeedAuctionPriceFeedSet, token []common.Address, feed []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var feedRule []interface{}
	for _, feedItem := range feed {
		feedRule = append(feedRule, feedItem)
	}

	logs, sub, err := _FeedAuction.contract.WatchLogs(opts, "PriceFeedSet", tokenRule, feedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeedAuctionPriceFeedSet)
				if err := _FeedAuction.contract.UnpackLog(event, "PriceFeedSet", log); err != nil {
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

// ParsePriceFeedSet is a log parse operation binding the contract event 0xd2d8394cf7549a5ddbc2ba3dd7b2de8d53c891472d1f2907008ed6a10045fdae.
//
// Solidity: event PriceFeedSet(address indexed token, address indexed feed)
func (_FeedAuction *FeedAuctionFilterer) ParsePriceFeedSet(log types.Log) (*FeedAuctionPriceFeedSet, error) {
	event := new(FeedAuctionPriceFeedSet)
	if err := _FeedAuction.contract.UnpackLog(event, "PriceFeedSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeedAuctionWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the FeedAuction contract.
type FeedAuctionWithdrawnIterator struct {
	Event *FeedAuctionWithdrawn // Event containing the contract specifics and raw log

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
func (it *FeedAuctionWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeedAuctionWithdrawn)
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
		it.Event = new(FeedAuctionWithdrawn)
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
func (it *FeedAuctionWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeedAuctionWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeedAuctionWithdrawn represents a Withdrawn event raised by the FeedAuction contract.
type FeedAuctionWithdrawn struct {
	AuctionId [32]byte
	Bidder    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x04eda370f8b8612fa7266d7ebbd41af9d694e19793fe9d9ff31b3ddbd99b08e1.
//
// Solidity: event Withdrawn(bytes32 indexed auctionId, address indexed bidder, uint256 amount)
func (_FeedAuction *FeedAuctionFilterer) FilterWithdrawn(opts *bind.FilterOpts, auctionId [][32]byte, bidder []common.Address) (*FeedAuctionWithdrawnIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FeedAuction.contract.FilterLogs(opts, "Withdrawn", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &FeedAuctionWithdrawnIterator{contract: _FeedAuction.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x04eda370f8b8612fa7266d7ebbd41af9d694e19793fe9d9ff31b3ddbd99b08e1.
//
// Solidity: event Withdrawn(bytes32 indexed auctionId, address indexed bidder, uint256 amount)
func (_FeedAuction *FeedAuctionFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *FeedAuctionWithdrawn, auctionId [][32]byte, bidder []common.Address) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _FeedAuction.contract.WatchLogs(opts, "Withdrawn", auctionIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeedAuctionWithdrawn)
				if err := _FeedAuction.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x04eda370f8b8612fa7266d7ebbd41af9d694e19793fe9d9ff31b3ddbd99b08e1.
//
// Solidity: event Withdrawn(bytes32 indexed auctionId, address indexed bidder, uint256 amount)
func (_FeedAuction *FeedAuctionFilterer) ParseWithdrawn(log types.Log) (*FeedAuctionWithdrawn, error) {
	event := new(FeedAuctionWithdrawn)
	if err := _FeedAuction.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
