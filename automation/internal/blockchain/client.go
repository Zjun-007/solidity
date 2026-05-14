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
    FeedAuction     *FeedAuction
    FeedAuctionAddr common.Address
    PrivateKey   *ecdsa.PrivateKey
    Auth         *bind.TransactOpts
    ChainID      *big.Int
}

func NewClient(rpcURL, contractAddr, privateKeyHex string, chainID int64) (*Client, error) {
    client, err := ethclient.Dial(rpcURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
    }

    contractAddress := common.HexToAddress(contractAddr)
    contract, err := NewFeedAuction(contractAddress, client)
    if err != nil {
        return nil, fmt.Errorf("failed to instantiate contract: %v", err)
    }

    privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
    if err != nil {
        return nil, fmt.Errorf("invalid private key: %v", err)
    }

    chainIDBig := big.NewInt(chainID)
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainIDBig)
    if err != nil {
        return nil, fmt.Errorf("failed to create transactor: %v", err)
    }

    // 获取nonce
    fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        return nil, fmt.Errorf("failed to get nonce: %v", err)
    }
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(300000)
    auth.GasPrice, _ = client.SuggestGasPrice(context.Background())

    return &Client{
        EthClient:    client,
        FeedAuction:     contract,
        FeedAuctionAddr: contractAddress,
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

func (c *Client) CreateAuction(nftFeedAuction common.Address, tokenId, startPrice, duration *big.Int) (*types.Transaction, error) {
    tx, err := c.FeedAuction.CreateAuction(c.Auth, nftFeedAuction, tokenId, startPrice, duration)
    if err != nil {
        return nil, err
    }
    c.RefreshNonce()
    return tx, nil
}

func (c *Client) Bid(auctionId *big.Int, value *big.Int) (*types.Transaction, error) {
    c.Auth.Value = value
    tx, err := c.FeedAuction.Bid(c.Auth, auctionId)
    c.Auth.Value = nil
    if err != nil {
        return nil, err
    }
    c.RefreshNonce()
    return tx, nil
}

func (c *Client) EndAuction(auctionId *big.Int) (*types.Transaction, error) {
    tx, err := c.FeedAuction.EndAuction(c.Auth, auctionId)
    if err != nil {
        return nil, err
    }
    c.RefreshNonce()
    return tx, nil
}