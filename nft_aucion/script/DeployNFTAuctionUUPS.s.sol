// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {NFTAuctionUUPS} from "../src/NFTAuctionUUPS.sol";

contract DeployNFTAuctionUUPS is Script {
    function run() external {
        address deployer = msg.sender;
        vm.startBroadcast();

        NFTAuctionUUPS impl = new NFTAuctionUUPS();
        bytes memory init = abi.encodeCall(NFTAuctionUUPS.initialize, (deployer));
        ERC1967Proxy proxy = new ERC1967Proxy(address(impl), init);

        console.log("Implementation:", address(impl));
        console.log("Proxy (use this address):", address(proxy));
        console.log("Owner:", deployer);

        vm.stopBroadcast();
    }
}
