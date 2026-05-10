package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "os"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/joho/godotenv"

    "geth/internal/blockchain"
)

func getenv(name string) string {
    v := strings.TrimSpace(os.Getenv(name))
    if v == "" {
        log.Fatalf("missing env %s", name)
    }
    return v
}

func main() {
    _ = godotenv.Load()

    rpc := getenv("ETH_RPC_URL")
    proxy := common.HexToAddress(getenv("NFT_AUCTION_PROXY_ADDRESS"))
    tokenId := new(big.Int)
    tokenId.SetString(getenv("TOKEN_ID"), 10)

    client, err := ethclient.Dial(rpc)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    auction, err := blockchain.NewNFTAuctionUUPSCaller(proxy, client)
    if err != nil {
        log.Fatal(err)
    }

    callOpts := &bind.CallOpts{Context: context.Background()}
    data, err := auction.AuctionData(callOpts, tokenId)
    if err != nil {
        log.Fatal("AuctionData failed:", err)
    }

    fmt.Printf("auctionData: seller=%s nft=%s tokenId=%s startPrice=%s startTime=%s endTime=%s highestBidder=%s highestBid=%s settled=%t cancelled=%t\n",
        data.Seller.Hex(), data.NftContract.Hex(), data.TokenId.String(), data.StartPrice.String(), data.StartTime.String(), data.EndTime.String(), data.HighestBidder.Hex(), data.HighestBid.String(), data.Settled, data.Cancelled)
}
