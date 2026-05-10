// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Test} from "forge-std/Test.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {ERC721} from "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {NFTAuctionUUPS} from "../src/NFTAuctionUUPS.sol";

contract TinyNFT is ERC721 {
    uint256 private _id;
    constructor() ERC721("Tiny", "TIN") {}

    function mint(address to) external returns (uint256) {
        unchecked {
            ++_id;
        }
        _safeMint(to, _id);
        return _id;
    }
}

contract NFTAuctionUUPSTest is Test {
    NFTAuctionUUPS auction;
    TinyNFT nft;
    address seller = address(0xA11);
    address bidder1 = address(0xB01);
    address bidder2 = address(0xB02);

    function setUp() public {
        vm.deal(seller, 100 ether);
        vm.deal(bidder1, 100 ether);
        vm.deal(bidder2, 100 ether);

        NFTAuctionUUPS impl = new NFTAuctionUUPS();
        bytes memory init = abi.encodeCall(NFTAuctionUUPS.initialize, (address(this)));
        auction = NFTAuctionUUPS(payable(address(new ERC1967Proxy(address(impl), init))));

        vm.prank(seller);
        nft = new TinyNFT();
        uint256 tid = nft.mint(seller);
        assertEq(tid, 1);

        vm.prank(seller);
        nft.setApprovalForAll(address(auction), true);
    }

    function test_cancel_before_start() public {
        uint256 start = block.timestamp + 1 days;
        uint256 end = start + 2 days;
        vm.prank(seller);
        uint256 aid = auction.createAuction(address(nft), 1, 1 ether, start, end);
        assertEq(aid, 1);

        vm.prank(seller);
        auction.cancelAuction(aid);
        assertEq(nft.ownerOf(1), seller);
        assertEq(auction.nftToken2AuctionId(address(nft), 1), 0);
    }

    function test_bid_and_end() public {
        uint256 start = block.timestamp;
        uint256 end = start + 2 days;
        vm.prank(seller);
        uint256 aid = auction.createAuction(address(nft), 1, 1 ether, start, end);

        vm.prank(bidder1);
        auction.bidAuction{value: 1 ether}(aid);

        vm.prank(bidder2);
        auction.bidAuction{value: 2 ether}(aid);
        assertEq(bidder1.balance, 100 ether); // refunded

        vm.warp(end);
        vm.prank(bidder2);
        auction.endAuction(aid);

        assertEq(nft.ownerOf(1), bidder2);
        assertEq(seller.balance, 100 ether + 2 ether);
        assertEq(auction.nftToken2AuctionId(address(nft), 1), 0);
    }

    function test_claim_unsold() public {
        uint256 start = block.timestamp;
        uint256 end = start + 1 hours;
        vm.prank(seller);
        uint256 aid = auction.createAuction(address(nft), 1, 5 ether, start, end);

        vm.warp(end);
        vm.prank(seller);
        auction.claimUnsoldNFT(aid);
        assertEq(nft.ownerOf(1), seller);
    }
}
