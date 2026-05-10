// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract FeedAuction is Ownable, ReentrancyGuard {
    using Address for address payable;

    IERC721 public nft;
    uint256 public tokenId;
    address public seller;
    uint256 public startPrice;
    uint256 public endTime;
    bool public canceled;
    bool public ended;
    bool public nftDeposited; // 新增：标记 NFT 是否已转入

    address public highestBidder;
    uint256 public highestBid;
    mapping(address => uint256) public pendingReturns;

    event AuctionCreated(address indexed seller, uint256 tokenId, uint256 startPrice, uint256 endTime);
    event Bid(address indexed bidder, uint256 amount);
    event AuctionEnded(address winner, uint256 amount);
    event AuctionCanceled();
    event Withdraw(address indexed bidder, uint256 amount);
    event NFTDeposited(); // 新增：NFT 转入事件

    constructor(
        address _nft,
        uint256 _tokenId,
        uint256 _startPrice,
        uint256 _duration
    ) Ownable(msg.sender) {
        require(_nft != address(0), "Invalid NFT address");
        require(_startPrice > 0, "Start price must be > 0");
        require(_duration > 0, "Duration must be > 0");

        nft = IERC721(_nft);
        tokenId = _tokenId;
        seller = msg.sender;
        startPrice = _startPrice;
        endTime = block.timestamp + _duration;
        nftDeposited = false;

        // 不再立即 transferFrom，等待卖家调用 depositNFT
        emit AuctionCreated(seller, tokenId, startPrice, endTime);
    }

    /**
     * @dev 卖家调用此函数，将 NFT 转入合约（必须在拍卖开始前调用）
     * 要求：卖家已授权本合约操作该 NFT（approve 或 setApprovalForAll）
     */
    function depositNFT() external {
        require(msg.sender == seller, "Only seller can deposit");
        require(!nftDeposited, "NFT already deposited");
        require(!canceled && !ended, "Auction already ended or canceled");
        require(block.timestamp < endTime, "Auction expired");

        nft.transferFrom(seller, address(this), tokenId);
        nftDeposited = true;

        emit NFTDeposited();
    }

    modifier auctionActive() {
        require(!canceled, "Auction canceled");
        require(!ended, "Auction already ended");
        require(block.timestamp < endTime, "Auction expired");
        require(nftDeposited, "NFT not yet deposited");
        _;
    }

    function bid() external payable nonReentrant auctionActive {
        require(msg.value >= startPrice, "Bid too low");
        require(msg.value > highestBid, "There already is a higher bid");

        if (highestBid != 0) {
            pendingReturns[highestBidder] += highestBid;
        }

        highestBid = msg.value;
        highestBidder = msg.sender;

        emit Bid(msg.sender, msg.value);
    }

    function withdraw() external nonReentrant {
        uint256 amount = pendingReturns[msg.sender];
        require(amount > 0, "No pending returns");

        pendingReturns[msg.sender] = 0;
        payable(msg.sender).sendValue(amount);

        emit Withdraw(msg.sender, amount);
    }

    function endAuction() external nonReentrant {
        require(msg.sender == seller, "Only seller can end auction");
        require(!canceled, "Auction canceled");
        require(!ended, "Auction already ended");
        require(block.timestamp >= endTime, "Auction not yet ended");
        require(nftDeposited, "NFT not deposited");

        ended = true;

        if (highestBidder != address(0)) {
            nft.transferFrom(address(this), highestBidder, tokenId);
            payable(seller).sendValue(highestBid);
            emit AuctionEnded(highestBidder, highestBid);
        } else {
            nft.transferFrom(address(this), seller, tokenId);
            emit AuctionEnded(address(0), 0);
        }
    }

    function cancelAuction() external {
        require(msg.sender == seller, "Only seller can cancel");
        require(!ended, "Auction already ended");
        require(block.timestamp < endTime, "Auction already expired");

        canceled = true;

        if (nftDeposited) {
            nft.transferFrom(address(this), seller, tokenId);
        }

        emit AuctionCanceled();
    }

    function isActive() public view returns (bool) {
        return !canceled && !ended && block.timestamp < endTime && nftDeposited;
    }

    function getHighestBid() public view returns (address bidder, uint256 amount) {
        return (highestBidder, highestBid);
    }
}