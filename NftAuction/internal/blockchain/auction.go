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

// BlockchainMetaData contains all meta data concerning the Blockchain contract.
var BlockchainMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_nft\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_startPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bid\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelAuction\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canceled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositNFT\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endAuction\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ended\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHighestBid\",\"inputs\":[],\"outputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBid\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"highestBidder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nft\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nftDeposited\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingReturns\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"seller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuctionCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionCreated\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"startPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionEnded\",\"inputs\":[{\"name\":\"winner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Bid\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTDeposited\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200275f3803806200275f8339818101604052810190620000369190620004b0565b335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000aa575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000a1919062000530565b60405180910390fd5b620000bb816200032060201b60201c565b506001620000de620000d2620003e160201b60201c565b6200040a60201b60201c565b5f01819055505f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff160362000155576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200014c90620005a9565b60405180910390fd5b5f82116200019a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001919062000617565b60405180910390fd5b5f8111620001df576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001d69062000685565b60405180910390fd5b8360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826002819055503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160048190555080426200027b9190620006d2565b6005819055505f600660026101000a81548160ff02191690831515021790555060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fbfc2001685f147773956c30c745f1b33e40df1ad4c1084b4d39c7c41bb7f66856002546004546005546040516200030e939291906200071d565b60405180910390a25050505062000758565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620004428262000417565b9050919050565b620004548162000436565b81146200045f575f80fd5b50565b5f81519050620004728162000449565b92915050565b5f819050919050565b6200048c8162000478565b811462000497575f80fd5b50565b5f81519050620004aa8162000481565b92915050565b5f805f8060808587031215620004cb57620004ca62000413565b5b5f620004da8782880162000462565b9450506020620004ed878288016200049a565b935050604062000500878288016200049a565b925050606062000513878288016200049a565b91505092959194509250565b6200052a8162000436565b82525050565b5f602082019050620005455f8301846200051f565b92915050565b5f82825260208201905092915050565b7f496e76616c6964204e46542061646472657373000000000000000000000000005f82015250565b5f620005916013836200054b565b91506200059e826200055b565b602082019050919050565b5f6020820190508181035f830152620005c28162000583565b9050919050565b7f5374617274207072696365206d757374206265203e20300000000000000000005f82015250565b5f620005ff6017836200054b565b91506200060c82620005c9565b602082019050919050565b5f6020820190508181035f8301526200063081620005f1565b9050919050565b7f4475726174696f6e206d757374206265203e20300000000000000000000000005f82015250565b5f6200066d6014836200054b565b91506200067a8262000637565b602082019050919050565b5f6020820190508181035f8301526200069e816200065f565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f620006de8262000478565b9150620006eb8362000478565b9250828201905080821115620007065762000705620006a5565b5b92915050565b620007178162000478565b82525050565b5f606082019050620007325f8301866200070c565b6200074160208301856200070c565b6200075060408301846200070c565b949350505050565b611ff980620007665f395ff3fe608060405260043610610129575f3560e01c806347ccca02116100aa5780638fa8b7901161006e5780638fa8b7901461035a57806391f9015714610370578063d57bde791461039a578063f1a9af89146103c4578063f2fde38b146103ee578063fe67a54b1461041657610129565b806347ccca02146102af5780634979440a146102d957806367e3c4d414610304578063715018a61461031a5780638da5cb5b1461033057610129565b806326b387bb116100f157806326b387bb146101df5780633197cbb61461021b5780633ca61eba146102455780633ccfd60b1461026f5780633f9942ff1461028557610129565b806308551a531461012d57806312fa6feb1461015757806317d70f7c146101815780631998aeef146101ab57806322f3e2d4146101b5575b5f80fd5b348015610138575f80fd5b5061014161042c565b60405161014e91906116e8565b60405180910390f35b348015610162575f80fd5b5061016b610451565b604051610178919061171b565b60405180910390f35b34801561018c575f80fd5b50610195610464565b6040516101a2919061174c565b60405180910390f35b6101b361046a565b005b3480156101c0575f80fd5b506101c961074e565b6040516101d6919061171b565b60405180910390f35b3480156101ea575f80fd5b5061020560048036038101906102009190611793565b6107a2565b604051610212919061174c565b60405180910390f35b348015610226575f80fd5b5061022f6107b7565b60405161023c919061174c565b60405180910390f35b348015610250575f80fd5b506102596107bd565b604051610266919061171b565b60405180910390f35b34801561027a575f80fd5b506102836107d0565b005b348015610290575f80fd5b5061029961091f565b6040516102a6919061171b565b60405180910390f35b3480156102ba575f80fd5b506102c3610931565b6040516102d09190611819565b60405180910390f35b3480156102e4575f80fd5b506102ed610956565b6040516102fb929190611832565b60405180910390f35b34801561030f575f80fd5b50610318610986565b005b348015610325575f80fd5b5061032e610c08565b005b34801561033b575f80fd5b50610344610c1b565b60405161035191906116e8565b60405180910390f35b348015610365575f80fd5b5061036e610c42565b005b34801561037b575f80fd5b50610384610e71565b60405161039191906116e8565b60405180910390f35b3480156103a5575f80fd5b506103ae610e97565b6040516103bb919061174c565b60405180910390f35b3480156103cf575f80fd5b506103d8610e9d565b6040516103e5919061174c565b60405180910390f35b3480156103f9575f80fd5b50610414600480360381019061040f9190611793565b610ea3565b005b348015610421575f80fd5b5061042a610f27565b005b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600660019054906101000a900460ff1681565b60025481565b6104726113b0565b60065f9054906101000a900460ff16156104c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b8906118b3565b60405180910390fd5b600660019054906101000a900460ff1615610511576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105089061191b565b60405180910390fd5b6005544210610555576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054c90611983565b60405180910390fd5b600660029054906101000a900460ff166105a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059b906119eb565b60405180910390fd5b6004543410156105e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e090611a53565b60405180910390fd5b600754341161062d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062490611abb565b60405180910390fd5b5f600754146106ae5760075460085f600660039054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106a69190611b06565b925050819055505b3460078190555033600660036101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503373ffffffffffffffffffffffffffffffffffffffff167fe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d23460405161073c919061174c565b60405180910390a261074c6113d2565b565b5f60065f9054906101000a900460ff161580156107785750600660019054906101000a900460ff16155b8015610785575060055442105b801561079d5750600660029054906101000a900460ff165b905090565b6008602052805f5260405f205f915090505481565b60055481565b600660029054906101000a900460ff1681565b6107d86113b0565b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f811161085b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085290611b83565b60405180910390fd5b5f60085f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055506108c6813373ffffffffffffffffffffffffffffffffffffffff166113ec90919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648260405161090c919061174c565b60405180910390a25061091d6113d2565b565b60065f9054906101000a900460ff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80600660039054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600754915091509091565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0c90611beb565b60405180910390fd5b600660029054906101000a900460ff1615610a65576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5c90611c53565b60405180910390fd5b60065f9054906101000a900460ff16158015610a8e5750600660019054906101000a900460ff16155b610acd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac490611ce1565b60405180910390fd5b6005544210610b11576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0890611983565b60405180910390fd5b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16306002546040518463ffffffff1660e01b8152600401610b9293929190611cff565b5f604051808303815f87803b158015610ba9575f80fd5b505af1158015610bbb573d5f803e3d5ffd5b505050506001600660026101000a81548160ff0219169083151502179055507ff598cc3c4e3735fcfdd2ee97fc90e46110fac921b6638cf16be65a27f72cbeac60405160405180910390a1565b610c106114a3565b610c195f61152a565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610cd1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc890611d7e565b60405180910390fd5b600660019054906101000a900460ff1615610d21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d189061191b565b60405180910390fd5b6005544210610d65576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5c90611de6565b60405180910390fd5b600160065f6101000a81548160ff021916908315150217905550600660029054906101000a900460ff1615610e435760015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166002546040518463ffffffff1660e01b8152600401610e1593929190611cff565b5f604051808303815f87803b158015610e2c575f80fd5b505af1158015610e3e573d5f803e3d5ffd5b505050505b7f3edab9d02702b82c68bc3a966d302af41abe4f3d8e2408b27b94e5f9bf2d722c60405160405180910390a1565b600660039054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b60045481565b610eab6114a3565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610f1b575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401610f1291906116e8565b60405180910390fd5b610f248161152a565b50565b610f2f6113b0565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb590611e4e565b60405180910390fd5b60065f9054906101000a900460ff161561100d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611004906118b3565b60405180910390fd5b600660019054906101000a900460ff161561105d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110549061191b565b60405180910390fd5b6005544210156110a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161109990611eb6565b60405180910390fd5b600660029054906101000a900460ff166110f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110e890611f1e565b60405180910390fd5b6001600660016101000a81548160ff0219169083151502179055505f73ffffffffffffffffffffffffffffffffffffffff16600660039054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146112be5760015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd30600660039054906101000a900473ffffffffffffffffffffffffffffffffffffffff166002546040518463ffffffff1660e01b81526004016111e393929190611cff565b5f604051808303815f87803b1580156111fa575f80fd5b505af115801561120c573d5f803e3d5ffd5b5050505061125c60075460035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166113ec90919063ffffffff16565b7fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda600660039054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007546040516112b1929190611832565b60405180910390a16113a6565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166002546040518463ffffffff1660e01b815260040161133f93929190611cff565b5f604051808303815f87803b158015611356575f80fd5b505af1158015611368573d5f803e3d5ffd5b505050507fdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda5f8060405161139d929190611f75565b60405180910390a15b6113ae6113d2565b565b6113b86115eb565b60026113ca6113c561162c565b611655565b5f0181905550565b60016113e46113df61162c565b611655565b5f0181905550565b804710156114335747816040517fcf47918100000000000000000000000000000000000000000000000000000000815260040161142a929190611f9c565b60405180910390fd5b61144c828260405180602001604052805f81525061165e565b61149f575f611459611674565b111561146c5761146761167b565b61149e565b6040517fd6bda27500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5050565b6114ab611686565b73ffffffffffffffffffffffffffffffffffffffff166114c9610c1b565b73ffffffffffffffffffffffffffffffffffffffff1614611528576114ec611686565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161151f91906116e8565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6115f361168d565b1561162a576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005f1b905090565b5f819050919050565b5f805f83516020850186885af190509392505050565b5f3d905090565b6040513d5f823e3d81fd5b5f33905090565b5f60026116a061169b61162c565b611655565b5f015414905090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6116d2826116a9565b9050919050565b6116e2816116c8565b82525050565b5f6020820190506116fb5f8301846116d9565b92915050565b5f8115159050919050565b61171581611701565b82525050565b5f60208201905061172e5f83018461170c565b92915050565b5f819050919050565b61174681611734565b82525050565b5f60208201905061175f5f83018461173d565b92915050565b5f80fd5b611772816116c8565b811461177c575f80fd5b50565b5f8135905061178d81611769565b92915050565b5f602082840312156117a8576117a7611765565b5b5f6117b58482850161177f565b91505092915050565b5f819050919050565b5f6117e16117dc6117d7846116a9565b6117be565b6116a9565b9050919050565b5f6117f2826117c7565b9050919050565b5f611803826117e8565b9050919050565b611813816117f9565b82525050565b5f60208201905061182c5f83018461180a565b92915050565b5f6040820190506118455f8301856116d9565b611852602083018461173d565b9392505050565b5f82825260208201905092915050565b7f41756374696f6e2063616e63656c6564000000000000000000000000000000005f82015250565b5f61189d601083611859565b91506118a882611869565b602082019050919050565b5f6020820190508181035f8301526118ca81611891565b9050919050565b7f41756374696f6e20616c726561647920656e64656400000000000000000000005f82015250565b5f611905601583611859565b9150611910826118d1565b602082019050919050565b5f6020820190508181035f830152611932816118f9565b9050919050565b7f41756374696f6e206578706972656400000000000000000000000000000000005f82015250565b5f61196d600f83611859565b915061197882611939565b602082019050919050565b5f6020820190508181035f83015261199a81611961565b9050919050565b7f4e4654206e6f7420796574206465706f736974656400000000000000000000005f82015250565b5f6119d5601583611859565b91506119e0826119a1565b602082019050919050565b5f6020820190508181035f830152611a02816119c9565b9050919050565b7f42696420746f6f206c6f770000000000000000000000000000000000000000005f82015250565b5f611a3d600b83611859565b9150611a4882611a09565b602082019050919050565b5f6020820190508181035f830152611a6a81611a31565b9050919050565b7f546865726520616c7265616479206973206120686967686572206269640000005f82015250565b5f611aa5601d83611859565b9150611ab082611a71565b602082019050919050565b5f6020820190508181035f830152611ad281611a99565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611b1082611734565b9150611b1b83611734565b9250828201905080821115611b3357611b32611ad9565b5b92915050565b7f4e6f2070656e64696e672072657475726e7300000000000000000000000000005f82015250565b5f611b6d601283611859565b9150611b7882611b39565b602082019050919050565b5f6020820190508181035f830152611b9a81611b61565b9050919050565b7f4f6e6c792073656c6c65722063616e206465706f7369740000000000000000005f82015250565b5f611bd5601783611859565b9150611be082611ba1565b602082019050919050565b5f6020820190508181035f830152611c0281611bc9565b9050919050565b7f4e465420616c7265616479206465706f736974656400000000000000000000005f82015250565b5f611c3d601583611859565b9150611c4882611c09565b602082019050919050565b5f6020820190508181035f830152611c6a81611c31565b9050919050565b7f41756374696f6e20616c726561647920656e646564206f722063616e63656c655f8201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b5f611ccb602183611859565b9150611cd682611c71565b604082019050919050565b5f6020820190508181035f830152611cf881611cbf565b9050919050565b5f606082019050611d125f8301866116d9565b611d1f60208301856116d9565b611d2c604083018461173d565b949350505050565b7f4f6e6c792073656c6c65722063616e2063616e63656c000000000000000000005f82015250565b5f611d68601683611859565b9150611d7382611d34565b602082019050919050565b5f6020820190508181035f830152611d9581611d5c565b9050919050565b7f41756374696f6e20616c726561647920657870697265640000000000000000005f82015250565b5f611dd0601783611859565b9150611ddb82611d9c565b602082019050919050565b5f6020820190508181035f830152611dfd81611dc4565b9050919050565b7f4f6e6c792073656c6c65722063616e20656e642061756374696f6e00000000005f82015250565b5f611e38601b83611859565b9150611e4382611e04565b602082019050919050565b5f6020820190508181035f830152611e6581611e2c565b9050919050565b7f41756374696f6e206e6f742079657420656e64656400000000000000000000005f82015250565b5f611ea0601583611859565b9150611eab82611e6c565b602082019050919050565b5f6020820190508181035f830152611ecd81611e94565b9050919050565b7f4e4654206e6f74206465706f73697465640000000000000000000000000000005f82015250565b5f611f08601183611859565b9150611f1382611ed4565b602082019050919050565b5f6020820190508181035f830152611f3581611efc565b9050919050565b5f819050919050565b5f611f5f611f5a611f5584611f3c565b6117be565b611734565b9050919050565b611f6f81611f45565b82525050565b5f604082019050611f885f8301856116d9565b611f956020830184611f66565b9392505050565b5f604082019050611faf5f83018561173d565b611fbc602083018461173d565b939250505056fea26469706673582212201ba99825d95bd43460848e6cc39757179cbc3eb13d29f4e0152b82d3b7aacf4e64736f6c63430008180033",
}

// BlockchainABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainMetaData.ABI instead.
var BlockchainABI = BlockchainMetaData.ABI

// BlockchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockchainMetaData.Bin instead.
var BlockchainBin = BlockchainMetaData.Bin

// DeployBlockchain deploys a new Ethereum contract, binding an instance of Blockchain to it.
func DeployBlockchain(auth *bind.TransactOpts, backend bind.ContractBackend, _nft common.Address, _tokenId *big.Int, _startPrice *big.Int, _duration *big.Int) (common.Address, *types.Transaction, *Blockchain, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockchainBin), backend, _nft, _tokenId, _startPrice, _duration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// Blockchain is an auto generated Go binding around an Ethereum contract.
type Blockchain struct {
	BlockchainCaller     // Read-only binding to the contract
	BlockchainTransactor // Write-only binding to the contract
	BlockchainFilterer   // Log filterer for contract events
}

// BlockchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainSession struct {
	Contract     *Blockchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainCallerSession struct {
	Contract *BlockchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlockchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainTransactorSession struct {
	Contract     *BlockchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainRaw struct {
	Contract *Blockchain // Generic contract binding to access the raw methods on
}

// BlockchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainCallerRaw struct {
	Contract *BlockchainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainTransactorRaw struct {
	Contract *BlockchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchain creates a new instance of Blockchain, bound to a specific deployed contract.
func NewBlockchain(address common.Address, backend bind.ContractBackend) (*Blockchain, error) {
	contract, err := bindBlockchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// NewBlockchainCaller creates a new read-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainCaller(address common.Address, caller bind.ContractCaller) (*BlockchainCaller, error) {
	contract, err := bindBlockchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainCaller{contract: contract}, nil
}

// NewBlockchainTransactor creates a new write-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainTransactor, error) {
	contract, err := bindBlockchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransactor{contract: contract}, nil
}

// NewBlockchainFilterer creates a new log filterer instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFilterer, error) {
	contract, err := bindBlockchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFilterer{contract: contract}, nil
}

// bindBlockchain binds a generic wrapper to an already deployed contract.
func bindBlockchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.BlockchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transact(opts, method, params...)
}

// Canceled is a free data retrieval call binding the contract method 0x3f9942ff.
//
// Solidity: function canceled() view returns(bool)
func (_Blockchain *BlockchainCaller) Canceled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "canceled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Canceled is a free data retrieval call binding the contract method 0x3f9942ff.
//
// Solidity: function canceled() view returns(bool)
func (_Blockchain *BlockchainSession) Canceled() (bool, error) {
	return _Blockchain.Contract.Canceled(&_Blockchain.CallOpts)
}

// Canceled is a free data retrieval call binding the contract method 0x3f9942ff.
//
// Solidity: function canceled() view returns(bool)
func (_Blockchain *BlockchainCallerSession) Canceled() (bool, error) {
	return _Blockchain.Contract.Canceled(&_Blockchain.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Blockchain *BlockchainCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Blockchain *BlockchainSession) EndTime() (*big.Int, error) {
	return _Blockchain.Contract.EndTime(&_Blockchain.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) EndTime() (*big.Int, error) {
	return _Blockchain.Contract.EndTime(&_Blockchain.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Blockchain *BlockchainCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Blockchain *BlockchainSession) Ended() (bool, error) {
	return _Blockchain.Contract.Ended(&_Blockchain.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Blockchain *BlockchainCallerSession) Ended() (bool, error) {
	return _Blockchain.Contract.Ended(&_Blockchain.CallOpts)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() view returns(address bidder, uint256 amount)
func (_Blockchain *BlockchainCaller) GetHighestBid(opts *bind.CallOpts) (struct {
	Bidder common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getHighestBid")

	outstruct := new(struct {
		Bidder common.Address
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bidder = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() view returns(address bidder, uint256 amount)
func (_Blockchain *BlockchainSession) GetHighestBid() (struct {
	Bidder common.Address
	Amount *big.Int
}, error) {
	return _Blockchain.Contract.GetHighestBid(&_Blockchain.CallOpts)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() view returns(address bidder, uint256 amount)
func (_Blockchain *BlockchainCallerSession) GetHighestBid() (struct {
	Bidder common.Address
	Amount *big.Int
}, error) {
	return _Blockchain.Contract.GetHighestBid(&_Blockchain.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Blockchain *BlockchainCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Blockchain *BlockchainSession) HighestBid() (*big.Int, error) {
	return _Blockchain.Contract.HighestBid(&_Blockchain.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) HighestBid() (*big.Int, error) {
	return _Blockchain.Contract.HighestBid(&_Blockchain.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Blockchain *BlockchainCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Blockchain *BlockchainSession) HighestBidder() (common.Address, error) {
	return _Blockchain.Contract.HighestBidder(&_Blockchain.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Blockchain *BlockchainCallerSession) HighestBidder() (common.Address, error) {
	return _Blockchain.Contract.HighestBidder(&_Blockchain.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Blockchain *BlockchainCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Blockchain *BlockchainSession) IsActive() (bool, error) {
	return _Blockchain.Contract.IsActive(&_Blockchain.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_Blockchain *BlockchainCallerSession) IsActive() (bool, error) {
	return _Blockchain.Contract.IsActive(&_Blockchain.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Blockchain *BlockchainCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Blockchain *BlockchainSession) Nft() (common.Address, error) {
	return _Blockchain.Contract.Nft(&_Blockchain.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Blockchain *BlockchainCallerSession) Nft() (common.Address, error) {
	return _Blockchain.Contract.Nft(&_Blockchain.CallOpts)
}

// NftDeposited is a free data retrieval call binding the contract method 0x3ca61eba.
//
// Solidity: function nftDeposited() view returns(bool)
func (_Blockchain *BlockchainCaller) NftDeposited(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "nftDeposited")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NftDeposited is a free data retrieval call binding the contract method 0x3ca61eba.
//
// Solidity: function nftDeposited() view returns(bool)
func (_Blockchain *BlockchainSession) NftDeposited() (bool, error) {
	return _Blockchain.Contract.NftDeposited(&_Blockchain.CallOpts)
}

// NftDeposited is a free data retrieval call binding the contract method 0x3ca61eba.
//
// Solidity: function nftDeposited() view returns(bool)
func (_Blockchain *BlockchainCallerSession) NftDeposited() (bool, error) {
	return _Blockchain.Contract.NftDeposited(&_Blockchain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainCallerSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// PendingReturns is a free data retrieval call binding the contract method 0x26b387bb.
//
// Solidity: function pendingReturns(address ) view returns(uint256)
func (_Blockchain *BlockchainCaller) PendingReturns(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "pendingReturns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReturns is a free data retrieval call binding the contract method 0x26b387bb.
//
// Solidity: function pendingReturns(address ) view returns(uint256)
func (_Blockchain *BlockchainSession) PendingReturns(arg0 common.Address) (*big.Int, error) {
	return _Blockchain.Contract.PendingReturns(&_Blockchain.CallOpts, arg0)
}

// PendingReturns is a free data retrieval call binding the contract method 0x26b387bb.
//
// Solidity: function pendingReturns(address ) view returns(uint256)
func (_Blockchain *BlockchainCallerSession) PendingReturns(arg0 common.Address) (*big.Int, error) {
	return _Blockchain.Contract.PendingReturns(&_Blockchain.CallOpts, arg0)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Blockchain *BlockchainCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Blockchain *BlockchainSession) Seller() (common.Address, error) {
	return _Blockchain.Contract.Seller(&_Blockchain.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Blockchain *BlockchainCallerSession) Seller() (common.Address, error) {
	return _Blockchain.Contract.Seller(&_Blockchain.CallOpts)
}

// StartPrice is a free data retrieval call binding the contract method 0xf1a9af89.
//
// Solidity: function startPrice() view returns(uint256)
func (_Blockchain *BlockchainCaller) StartPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "startPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartPrice is a free data retrieval call binding the contract method 0xf1a9af89.
//
// Solidity: function startPrice() view returns(uint256)
func (_Blockchain *BlockchainSession) StartPrice() (*big.Int, error) {
	return _Blockchain.Contract.StartPrice(&_Blockchain.CallOpts)
}

// StartPrice is a free data retrieval call binding the contract method 0xf1a9af89.
//
// Solidity: function startPrice() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) StartPrice() (*big.Int, error) {
	return _Blockchain.Contract.StartPrice(&_Blockchain.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Blockchain *BlockchainCaller) TokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "tokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Blockchain *BlockchainSession) TokenId() (*big.Int, error) {
	return _Blockchain.Contract.TokenId(&_Blockchain.CallOpts)
}

// TokenId is a free data retrieval call binding the contract method 0x17d70f7c.
//
// Solidity: function tokenId() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) TokenId() (*big.Int, error) {
	return _Blockchain.Contract.TokenId(&_Blockchain.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Blockchain *BlockchainTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Blockchain *BlockchainSession) Bid() (*types.Transaction, error) {
	return _Blockchain.Contract.Bid(&_Blockchain.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Blockchain *BlockchainTransactorSession) Bid() (*types.Transaction, error) {
	return _Blockchain.Contract.Bid(&_Blockchain.TransactOpts)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x8fa8b790.
//
// Solidity: function cancelAuction() returns()
func (_Blockchain *BlockchainTransactor) CancelAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "cancelAuction")
}

// CancelAuction is a paid mutator transaction binding the contract method 0x8fa8b790.
//
// Solidity: function cancelAuction() returns()
func (_Blockchain *BlockchainSession) CancelAuction() (*types.Transaction, error) {
	return _Blockchain.Contract.CancelAuction(&_Blockchain.TransactOpts)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x8fa8b790.
//
// Solidity: function cancelAuction() returns()
func (_Blockchain *BlockchainTransactorSession) CancelAuction() (*types.Transaction, error) {
	return _Blockchain.Contract.CancelAuction(&_Blockchain.TransactOpts)
}

// DepositNFT is a paid mutator transaction binding the contract method 0x67e3c4d4.
//
// Solidity: function depositNFT() returns()
func (_Blockchain *BlockchainTransactor) DepositNFT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "depositNFT")
}

// DepositNFT is a paid mutator transaction binding the contract method 0x67e3c4d4.
//
// Solidity: function depositNFT() returns()
func (_Blockchain *BlockchainSession) DepositNFT() (*types.Transaction, error) {
	return _Blockchain.Contract.DepositNFT(&_Blockchain.TransactOpts)
}

// DepositNFT is a paid mutator transaction binding the contract method 0x67e3c4d4.
//
// Solidity: function depositNFT() returns()
func (_Blockchain *BlockchainTransactorSession) DepositNFT() (*types.Transaction, error) {
	return _Blockchain.Contract.DepositNFT(&_Blockchain.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Blockchain *BlockchainTransactor) EndAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "endAuction")
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Blockchain *BlockchainSession) EndAuction() (*types.Transaction, error) {
	return _Blockchain.Contract.EndAuction(&_Blockchain.TransactOpts)
}

// EndAuction is a paid mutator transaction binding the contract method 0xfe67a54b.
//
// Solidity: function endAuction() returns()
func (_Blockchain *BlockchainTransactorSession) EndAuction() (*types.Transaction, error) {
	return _Blockchain.Contract.EndAuction(&_Blockchain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockchain.Contract.RenounceOwnership(&_Blockchain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockchain.Contract.RenounceOwnership(&_Blockchain.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.TransferOwnership(&_Blockchain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.TransferOwnership(&_Blockchain.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Blockchain *BlockchainTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Blockchain *BlockchainSession) Withdraw() (*types.Transaction, error) {
	return _Blockchain.Contract.Withdraw(&_Blockchain.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Blockchain *BlockchainTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Blockchain.Contract.Withdraw(&_Blockchain.TransactOpts)
}

// BlockchainAuctionCanceledIterator is returned from FilterAuctionCanceled and is used to iterate over the raw logs and unpacked data for AuctionCanceled events raised by the Blockchain contract.
type BlockchainAuctionCanceledIterator struct {
	Event *BlockchainAuctionCanceled // Event containing the contract specifics and raw log

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
func (it *BlockchainAuctionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainAuctionCanceled)
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
		it.Event = new(BlockchainAuctionCanceled)
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
func (it *BlockchainAuctionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainAuctionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainAuctionCanceled represents a AuctionCanceled event raised by the Blockchain contract.
type BlockchainAuctionCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAuctionCanceled is a free log retrieval operation binding the contract event 0x3edab9d02702b82c68bc3a966d302af41abe4f3d8e2408b27b94e5f9bf2d722c.
//
// Solidity: event AuctionCanceled()
func (_Blockchain *BlockchainFilterer) FilterAuctionCanceled(opts *bind.FilterOpts) (*BlockchainAuctionCanceledIterator, error) {

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "AuctionCanceled")
	if err != nil {
		return nil, err
	}
	return &BlockchainAuctionCanceledIterator{contract: _Blockchain.contract, event: "AuctionCanceled", logs: logs, sub: sub}, nil
}

// WatchAuctionCanceled is a free log subscription operation binding the contract event 0x3edab9d02702b82c68bc3a966d302af41abe4f3d8e2408b27b94e5f9bf2d722c.
//
// Solidity: event AuctionCanceled()
func (_Blockchain *BlockchainFilterer) WatchAuctionCanceled(opts *bind.WatchOpts, sink chan<- *BlockchainAuctionCanceled) (event.Subscription, error) {

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "AuctionCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainAuctionCanceled)
				if err := _Blockchain.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
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

// ParseAuctionCanceled is a log parse operation binding the contract event 0x3edab9d02702b82c68bc3a966d302af41abe4f3d8e2408b27b94e5f9bf2d722c.
//
// Solidity: event AuctionCanceled()
func (_Blockchain *BlockchainFilterer) ParseAuctionCanceled(log types.Log) (*BlockchainAuctionCanceled, error) {
	event := new(BlockchainAuctionCanceled)
	if err := _Blockchain.contract.UnpackLog(event, "AuctionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Blockchain contract.
type BlockchainAuctionCreatedIterator struct {
	Event *BlockchainAuctionCreated // Event containing the contract specifics and raw log

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
func (it *BlockchainAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainAuctionCreated)
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
		it.Event = new(BlockchainAuctionCreated)
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
func (it *BlockchainAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainAuctionCreated represents a AuctionCreated event raised by the Blockchain contract.
type BlockchainAuctionCreated struct {
	Seller     common.Address
	TokenId    *big.Int
	StartPrice *big.Int
	EndTime    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xbfc2001685f147773956c30c745f1b33e40df1ad4c1084b4d39c7c41bb7f6685.
//
// Solidity: event AuctionCreated(address indexed seller, uint256 tokenId, uint256 startPrice, uint256 endTime)
func (_Blockchain *BlockchainFilterer) FilterAuctionCreated(opts *bind.FilterOpts, seller []common.Address) (*BlockchainAuctionCreatedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "AuctionCreated", sellerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainAuctionCreatedIterator{contract: _Blockchain.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xbfc2001685f147773956c30c745f1b33e40df1ad4c1084b4d39c7c41bb7f6685.
//
// Solidity: event AuctionCreated(address indexed seller, uint256 tokenId, uint256 startPrice, uint256 endTime)
func (_Blockchain *BlockchainFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *BlockchainAuctionCreated, seller []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "AuctionCreated", sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainAuctionCreated)
				if err := _Blockchain.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xbfc2001685f147773956c30c745f1b33e40df1ad4c1084b4d39c7c41bb7f6685.
//
// Solidity: event AuctionCreated(address indexed seller, uint256 tokenId, uint256 startPrice, uint256 endTime)
func (_Blockchain *BlockchainFilterer) ParseAuctionCreated(log types.Log) (*BlockchainAuctionCreated, error) {
	event := new(BlockchainAuctionCreated)
	if err := _Blockchain.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the Blockchain contract.
type BlockchainAuctionEndedIterator struct {
	Event *BlockchainAuctionEnded // Event containing the contract specifics and raw log

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
func (it *BlockchainAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainAuctionEnded)
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
		it.Event = new(BlockchainAuctionEnded)
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
func (it *BlockchainAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainAuctionEnded represents a AuctionEnded event raised by the Blockchain contract.
type BlockchainAuctionEnded struct {
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Blockchain *BlockchainFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*BlockchainAuctionEndedIterator, error) {

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &BlockchainAuctionEndedIterator{contract: _Blockchain.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Blockchain *BlockchainFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *BlockchainAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainAuctionEnded)
				if err := _Blockchain.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0xdaec4582d5d9595688c8c98545fdd1c696d41c6aeaeb636737e84ed2f5c00eda.
//
// Solidity: event AuctionEnded(address winner, uint256 amount)
func (_Blockchain *BlockchainFilterer) ParseAuctionEnded(log types.Log) (*BlockchainAuctionEnded, error) {
	event := new(BlockchainAuctionEnded)
	if err := _Blockchain.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Blockchain contract.
type BlockchainBidIterator struct {
	Event *BlockchainBid // Event containing the contract specifics and raw log

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
func (it *BlockchainBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainBid)
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
		it.Event = new(BlockchainBid)
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
func (it *BlockchainBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainBid represents a Bid event raised by the Blockchain contract.
type BlockchainBid struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 amount)
func (_Blockchain *BlockchainFilterer) FilterBid(opts *bind.FilterOpts, bidder []common.Address) (*BlockchainBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainBidIterator{contract: _Blockchain.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 amount)
func (_Blockchain *BlockchainFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *BlockchainBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainBid)
				if err := _Blockchain.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 amount)
func (_Blockchain *BlockchainFilterer) ParseBid(log types.Log) (*BlockchainBid, error) {
	event := new(BlockchainBid)
	if err := _Blockchain.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainNFTDepositedIterator is returned from FilterNFTDeposited and is used to iterate over the raw logs and unpacked data for NFTDeposited events raised by the Blockchain contract.
type BlockchainNFTDepositedIterator struct {
	Event *BlockchainNFTDeposited // Event containing the contract specifics and raw log

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
func (it *BlockchainNFTDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainNFTDeposited)
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
		it.Event = new(BlockchainNFTDeposited)
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
func (it *BlockchainNFTDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainNFTDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainNFTDeposited represents a NFTDeposited event raised by the Blockchain contract.
type BlockchainNFTDeposited struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNFTDeposited is a free log retrieval operation binding the contract event 0xf598cc3c4e3735fcfdd2ee97fc90e46110fac921b6638cf16be65a27f72cbeac.
//
// Solidity: event NFTDeposited()
func (_Blockchain *BlockchainFilterer) FilterNFTDeposited(opts *bind.FilterOpts) (*BlockchainNFTDepositedIterator, error) {

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "NFTDeposited")
	if err != nil {
		return nil, err
	}
	return &BlockchainNFTDepositedIterator{contract: _Blockchain.contract, event: "NFTDeposited", logs: logs, sub: sub}, nil
}

// WatchNFTDeposited is a free log subscription operation binding the contract event 0xf598cc3c4e3735fcfdd2ee97fc90e46110fac921b6638cf16be65a27f72cbeac.
//
// Solidity: event NFTDeposited()
func (_Blockchain *BlockchainFilterer) WatchNFTDeposited(opts *bind.WatchOpts, sink chan<- *BlockchainNFTDeposited) (event.Subscription, error) {

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "NFTDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainNFTDeposited)
				if err := _Blockchain.contract.UnpackLog(event, "NFTDeposited", log); err != nil {
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

// ParseNFTDeposited is a log parse operation binding the contract event 0xf598cc3c4e3735fcfdd2ee97fc90e46110fac921b6638cf16be65a27f72cbeac.
//
// Solidity: event NFTDeposited()
func (_Blockchain *BlockchainFilterer) ParseNFTDeposited(log types.Log) (*BlockchainNFTDeposited, error) {
	event := new(BlockchainNFTDeposited)
	if err := _Blockchain.contract.UnpackLog(event, "NFTDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Blockchain contract.
type BlockchainOwnershipTransferredIterator struct {
	Event *BlockchainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockchainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainOwnershipTransferred)
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
		it.Event = new(BlockchainOwnershipTransferred)
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
func (it *BlockchainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainOwnershipTransferred represents a OwnershipTransferred event raised by the Blockchain contract.
type BlockchainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockchainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainOwnershipTransferredIterator{contract: _Blockchain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockchainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainOwnershipTransferred)
				if err := _Blockchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Blockchain *BlockchainFilterer) ParseOwnershipTransferred(log types.Log) (*BlockchainOwnershipTransferred, error) {
	event := new(BlockchainOwnershipTransferred)
	if err := _Blockchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Blockchain contract.
type BlockchainWithdrawIterator struct {
	Event *BlockchainWithdraw // Event containing the contract specifics and raw log

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
func (it *BlockchainWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainWithdraw)
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
		it.Event = new(BlockchainWithdraw)
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
func (it *BlockchainWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainWithdraw represents a Withdraw event raised by the Blockchain contract.
type BlockchainWithdraw struct {
	Bidder common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 amount)
func (_Blockchain *BlockchainFilterer) FilterWithdraw(opts *bind.FilterOpts, bidder []common.Address) (*BlockchainWithdrawIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "Withdraw", bidderRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainWithdrawIterator{contract: _Blockchain.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 amount)
func (_Blockchain *BlockchainFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BlockchainWithdraw, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "Withdraw", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainWithdraw)
				if err := _Blockchain.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed bidder, uint256 amount)
func (_Blockchain *BlockchainFilterer) ParseWithdraw(log types.Log) (*BlockchainWithdraw, error) {
	event := new(BlockchainWithdraw)
	if err := _Blockchain.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
