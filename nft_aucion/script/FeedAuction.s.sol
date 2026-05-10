// script/FeedAuction.s.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script} from "forge-std/Script.sol";
import {FeedAuction} from "../src/FeedAuction.sol";
import "forge-std/console.sol";

contract FeedAuctionScript is Script {
    function run() external {
        // 从环境变量读取部署参数
        address nftAddress = vm.envAddress("NFT_CONTRACT_ADDRESS");
        uint256 tokenId = vm.envUint("TOKEN_ID");
        uint256 startPrice = vm.envUint("START_PRICE");
        uint256 duration = vm.envUint("DURATION");

        // 直接使用命令行传入的私钥（--private-key）
        vm.startBroadcast();

        FeedAuction auction = new FeedAuction(nftAddress, tokenId, startPrice, duration);

        vm.stopBroadcast();

        console.log("FeedAuction deployed at:", address(auction));
        console.log("Seller address:", auction.seller());
        console.log("End time:", auction.endTime());
    }
}