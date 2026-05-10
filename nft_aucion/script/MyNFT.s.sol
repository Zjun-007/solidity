// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script} from "forge-std/Script.sol";
import "forge-std/console.sol";
import {MyNFT} from "../src/MyNFT.sol";

contract DeployMyNFT is Script {
    function run() external {
        // 开始广播交易（使用部署私钥）
        vm.startBroadcast();

        // 部署 MyNFT 合约
        MyNFT nft = new MyNFT();

        // 可选：记录合约地址
        console.log("MyNFT deployed to:", address(nft));

        vm.stopBroadcast();
    }
}