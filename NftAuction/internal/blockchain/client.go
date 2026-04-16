// internal/blockchain/client.go
package blockchain

import (
    "context"
    "crypto/ecdsa"
    "fmt"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
    EthClient    *ethclient.Client
    Auction      *Blockchain          // 改为 Blockchain
    ContractAddr common.Address
    PrivateKey   *ecdsa.PrivateKey
    Auth         *bind.TransactOpts
    ChainID      *big.Int
}

func NewClient(rpcURL, contractAddr, privateKeyHex string, chainID int64) (*Client, error) {
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return nil, fmt.Errorf("dial eth client: %w", err)
    }

    contractAddress := common.HexToAddress(contractAddr)
    contract, err := NewBlockchain(contractAddress, client)   // 改为 NewBlockchain
    if err != nil {
        return nil, fmt.Errorf("bind contract: %w", err)
    }

    privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
    if err != nil {
        return nil, fmt.Errorf("parse private key: %w", err)
    }

    chainIDBig := big.NewInt(chainID)
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainIDBig)
    if err != nil {
        return nil, fmt.Errorf("create transactor: %w", err)
    }

    ctx := context.Background()
    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    nonce, err := client.PendingNonceAt(ctx, fromAddress)
    if err != nil {
        return nil, fmt.Errorf("get nonce: %w", err)
    }
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = 0 // 自动估算
    gasPrice, err := client.SuggestGasPrice(ctx)
    if err != nil {
        return nil, fmt.Errorf("suggest gas price: %w", err)
    }
    auth.GasPrice = gasPrice

    return &Client{
        EthClient:    client,
        Auction:      contract,
        ContractAddr: contractAddress,
        PrivateKey:   privateKey,
        Auth:         auth,
        ChainID:      chainIDBig,
    }, nil
}

func (c *Client) RefreshNonce() error {
    fromAddress := crypto.PubkeyToAddress(c.PrivateKey.PublicKey)
    nonce, err := c.EthClient.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        return err
    }
    c.Auth.Nonce = big.NewInt(int64(nonce))
    return nil
}

// DepositNFT 卖家存入 NFT（需预先授权）
func (c *Client) DepositNFT() (*types.Transaction, error) {
    tx, err := c.Auction.DepositNFT(c.Auth)
    if err != nil {
        return nil, fmt.Errorf("deposit nft: %w", err)
    }
    _ = c.RefreshNonce()
    return tx, nil
}

// Bid 出价
func (c *Client) Bid(value *big.Int) (*types.Transaction, error) {
    c.Auth.Value = value
    tx, err := c.Auction.Bid(c.Auth)
    c.Auth.Value = nil
    if err != nil {
        return nil, fmt.Errorf("bid: %w", err)
    }
    _ = c.RefreshNonce()
    return tx, nil
}

// EndAuction 结束拍卖（仅卖家）
func (c *Client) EndAuction() (*types.Transaction, error) {
    tx, err := c.Auction.EndAuction(c.Auth)
    if err != nil {
        return nil, fmt.Errorf("end auction: %w", err)
    }
    _ = c.RefreshNonce()
    return tx, nil
}

// Withdraw 提取被超越的出价
func (c *Client) Withdraw() (*types.Transaction, error) {
    tx, err := c.Auction.Withdraw(c.Auth)
    if err != nil {
        return nil, fmt.Errorf("withdraw: %w", err)
    }
    _ = c.RefreshNonce()
    return tx, nil
}

// CancelAuction 取消拍卖（仅卖家）
func (c *Client) CancelAuction() (*types.Transaction, error) {
    tx, err := c.Auction.CancelAuction(c.Auth)
    if err != nil {
        return nil, fmt.Errorf("cancel auction: %w", err)
    }
    _ = c.RefreshNonce()
    return tx, nil
}

// GetAuctionInfo 查询当前拍卖所有状态
func (c *Client) GetAuctionInfo(ctx context.Context) (*AuctionInfo, error) {
    opts := &bind.CallOpts{Context: ctx}
    nft, _ := c.Auction.Nft(opts)
    tokenId, _ := c.Auction.TokenId(opts)
    seller, _ := c.Auction.Seller(opts)
    startPrice, _ := c.Auction.StartPrice(opts)
    endTime, _ := c.Auction.EndTime(opts)
    highestBid, _ := c.Auction.HighestBid(opts)
    highestBidder, _ := c.Auction.HighestBidder(opts)
    ended, _ := c.Auction.Ended(opts)
    canceled, _ := c.Auction.Canceled(opts)
    nftDeposited, _ := c.Auction.NftDeposited(opts)
    isActive, _ := c.Auction.IsActive(opts)

    return &AuctionInfo{
        Nft:           nft,
        TokenId:       tokenId,
        Seller:        seller,
        StartPrice:    startPrice,
        EndTime:       endTime,
        HighestBid:    highestBid,
        HighestBidder: highestBidder,
        Ended:         ended,
        Canceled:      canceled,
        NftDeposited:  nftDeposited,
        IsActive:      isActive,
    }, nil
}

type AuctionInfo struct {
    Nft           common.Address
    TokenId       *big.Int
    Seller        common.Address
    StartPrice    *big.Int
    EndTime       *big.Int
    HighestBid    *big.Int
    HighestBidder common.Address
    Ended         bool
    Canceled      bool
    NftDeposited  bool
    IsActive      bool
}