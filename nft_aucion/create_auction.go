package main

import (
    "context"
    "encoding/hex"
    "fmt"
    "log"
    "math/big"
    "os"
    "strings"
    "time"

    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/joho/godotenv"

    "geth/internal/blockchain"
)

const erc721ABI = `[
    {"constant":true,"inputs":[{"name":"_tokenId","type":"uint256"}],"name":"ownerOf","outputs":[{"name":"_owner","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},
    {"constant":true,"inputs":[{"name":"_tokenId","type":"uint256"}],"name":"getApproved","outputs":[{"name":"_approved","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},
    {"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_operator","type":"address"}],"name":"isApprovedForAll","outputs":[{"name":"_approved","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},
    {"constant":false,"inputs":[{"name":"_approved","type":"address"},{"name":"_tokenId","type":"uint256"}],"name":"approve","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}
]`

func getenv(name string) string {
    v := strings.TrimSpace(os.Getenv(name))
    if v == "" {
        log.Fatalf("missing env %s", name)
    }
    return v
}

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("no .env file loaded")
    }

    rpc := getenv("ETH_RPC_URL")
    pkHex := getenv("PRIVATE_KEY")
    proxy := common.HexToAddress(getenv("NFT_AUCTION_PROXY_ADDRESS"))
    nft := common.HexToAddress(getenv("NFT_CONTRACT_ADDRESS"))
    tokenIdStr := getenv("TOKEN_ID")
    tokenId := new(big.Int)
    tokenId.SetString(strings.TrimSpace(tokenIdStr), 10)
    if tokenId.Sign() <= 0 {
        log.Fatalf("invalid TOKEN_ID: %s", tokenIdStr)
    }

    client, err := ethclient.Dial(rpc)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    keyBytes, err := hexutil.Decode(pkHex)
    if err != nil {
        keyBytes, err = hex.DecodeString(strings.TrimPrefix(pkHex, "0x"))
        if err != nil {
            log.Fatal(err)
        }
    }
    key, err := crypto.ToECDSA(keyBytes)
    if err != nil {
        log.Fatal(err)
    }

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
    if err != nil {
        log.Fatal(err)
    }
    auth.Context = context.Background()

    owner := crypto.PubkeyToAddress(key.PublicKey)
    fmt.Println("sender:", owner.Hex())

    nftABI, err := abi.JSON(strings.NewReader(erc721ABI))
    if err != nil {
        log.Fatal(err)
    }
    nftContract := bind.NewBoundContract(nft, nftABI, client, client, client)

    callOpts := &bind.CallOpts{Context: context.Background()}

    ownerResult := []interface{}{new(common.Address)}
    err = nftContract.Call(callOpts, &ownerResult, "ownerOf", tokenId)
    if err != nil {
        log.Fatal("ownerOf failed:", err)
    }
    currentOwner := *ownerResult[0].(*common.Address)
    fmt.Println("ownerOf token", tokenId.String(), "=", currentOwner.Hex())
    if currentOwner != owner {
        log.Fatalf("sender does not own token %s", tokenId.String())
    }

    approvedResult := []interface{}{new(common.Address)}
    err = nftContract.Call(callOpts, &approvedResult, "getApproved", tokenId)
    if err != nil {
        log.Fatal("getApproved failed:", err)
    }
    approved := *approvedResult[0].(*common.Address)
    fmt.Println("getApproved:", approved.Hex())

    isAllResult := []interface{}{new(bool)}
    err = nftContract.Call(callOpts, &isAllResult, "isApprovedForAll", owner, proxy)
    if err != nil {
        log.Fatal("isApprovedForAll failed:", err)
    }
    isAll := *isAllResult[0].(*bool)
    fmt.Println("isApprovedForAll:", isAll)

    if approved != proxy && !isAll {
        fmt.Println("approving auction proxy for token", tokenId.String())
        tx, err := nftContract.Transact(auth, "approve", proxy, tokenId)
        if err != nil {
            log.Fatal("approve failed:", err)
        }
        fmt.Println("approve tx:", tx.Hash().Hex())
        fmt.Println("waiting for approve confirmation...")
        receipt, err := bind.WaitMined(context.Background(), client, tx)
        if err != nil {
            log.Fatal(err)
        }
        if receipt.Status != 1 {
            log.Fatal("approve tx failed")
        }
    }

    auction, err := blockchain.NewNFTAuctionUUPSTransactor(proxy, client)
    if err != nil {
        log.Fatal(err)
    }

    now := time.Now().Unix()
    start := big.NewInt(now + 60)
    end := big.NewInt(now + 3600)
    startPrice := new(big.Int)
    startPrice.SetString("1000000000000000", 10)

    fmt.Println("createAuction with startTime", start.String(), "endTime", end.String())
    tx, err := auction.CreateAuction(auth, nft, tokenId, startPrice, start, end)
    if err != nil {
        log.Fatal("createAuction failed:", err)
    }
    fmt.Println("createAuction tx:", tx.Hash().Hex())
}
