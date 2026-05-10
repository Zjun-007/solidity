// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {ReentrancyGuardTransient} from "@openzeppelin/contracts/utils/ReentrancyGuardTransient.sol";
import {ERC721Holder} from "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import {IERC721} from "@openzeppelin/contracts/token/ERC721/IERC721.sol";

/**
 * @title NFTAuctionUUPS
 * @notice 单合约管理全部 NFT 拍卖；ETH 出价；UUPS 可升级。
 *         NFT 创建时托管在本合约；结束后胜者 `endAuction` 取货并打款给卖家。
 */
contract NFTAuctionUUPS is
    Initializable,
    OwnableUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardTransient,
    ERC721Holder
{
    struct Auction {
        address seller;
        address nftContract;
        uint256 tokenId;
        uint256 startPrice;
        uint256 startTime;
        uint256 endTime;
        address highestBidder;
        uint256 highestBid;
        bool settled;
        bool cancelled;
    }

    /// @dev NFT 合约 => tokenId => 当前进行中的拍卖 id（已结束/取消后为 0）
    mapping(address => mapping(uint256 => uint256)) public nftToken2AuctionId;
    /// @dev 拍卖 id => 数据
    mapping(uint256 => Auction) public auctionData;
    /// @dev 下一个拍卖 id（从 1 开始）
    uint256 public nextAuctionId;

    /// @custom:oz-upgrades-unsafe-allow state-variable-immutable
    uint256[47] private __gap;

    event AuctionCreated(
        uint256 indexed auctionId,
        address indexed seller,
        address indexed nftContract,
        uint256 tokenId,
        uint256 startPrice,
        uint256 startTime,
        uint256 endTime
    );
    event BidAuction(uint256 indexed auctionId, address indexed bidder, uint256 amount);
    event AuctionEnded(uint256 indexed auctionId, address indexed winner, uint256 amount, address indexed seller);
    event AuctionCancelled(uint256 indexed auctionId, address indexed seller);

    error InvalidTimeRange();
    error StartInPast();
    error ActiveAuctionExists();
    error NotSeller();
    error NotBiddingWindow();
    error AuctionInactive();
    error BidTooLow();
    error NotWinner();
    error AuctionNotEnded();
    error AlreadySettled();
    error NoBidsToSettle();
    error CannotCancel();
    error NotSellerOrHasBids();
    error ETHRefundFailed();
    error ETHPaySellerFailed();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(address initialOwner) external initializer {
        __Ownable_init(initialOwner);
    }

    function _authorizeUpgrade(address newImplementation) internal override onlyOwner {}

    /**
     * @notice 为指定 NFT 创建拍卖（需已 approve / setApprovalForAll 本合约）
     * @param nftContract ERC721 地址
     * @param tokenId tokenId
     * @param startPrice wei 起拍价
     * @param startTime 允许出价的开始时间戳（在此之前卖家可 cancel）
     * @param endTime 拍卖结束时间戳（必须 > startTime）
     */
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 startPrice,
        uint256 startTime,
        uint256 endTime
    ) external nonReentrant returns (uint256 auctionId) {
        if (endTime <= startTime) revert InvalidTimeRange();
        if (startTime < block.timestamp) revert StartInPast();
        if (startPrice == 0) revert BidTooLow();
        if (nftToken2AuctionId[nftContract][tokenId] != 0) revert ActiveAuctionExists();

        IERC721 nft = IERC721(nftContract);
        if (nft.ownerOf(tokenId) != msg.sender) revert NotSeller();

        auctionId = ++nextAuctionId;
        nftToken2AuctionId[nftContract][tokenId] = auctionId;

        Auction storage a = auctionData[auctionId];
        a.seller = msg.sender;
        a.nftContract = nftContract;
        a.tokenId = tokenId;
        a.startPrice = startPrice;
        a.startTime = startTime;
        a.endTime = endTime;

        nft.safeTransferFrom(msg.sender, address(this), tokenId);

        emit AuctionCreated(auctionId, msg.sender, nftContract, tokenId, startPrice, startTime, endTime);
    }

    /**
     * @notice 在 [startTime, endTime) 内以 ETH 出价；被超过则立即退回前一最高出价者。
     */
    function bidAuction(uint256 auctionId) external payable nonReentrant {
        Auction storage a = auctionData[auctionId];
        if (a.seller == address(0)) revert AuctionInactive();
        if (a.settled || a.cancelled) revert AuctionInactive();
        if (block.timestamp < a.startTime || block.timestamp >= a.endTime) revert NotBiddingWindow();

        uint256 value = msg.value;
        if (a.highestBidder == address(0)) {
            if (value < a.startPrice) revert BidTooLow();
        } else {
            if (value <= a.highestBid) revert BidTooLow();
        }

        address prev = a.highestBidder;
        uint256 prevAmt = a.highestBid;
        a.highestBidder = msg.sender;
        a.highestBid = value;

        if (prev != address(0) && prevAmt > 0) {
            (bool ok, ) = payable(prev).call{value: prevAmt}("");
            if (!ok) revert ETHRefundFailed();
        }

        emit BidAuction(auctionId, msg.sender, value);
    }

    /**
     * @notice 拍卖结束后由最高出价者调用：NFT 转给胜者，ETH 给卖家。
     */
    function endAuction(uint256 auctionId) external nonReentrant {
        Auction storage a = auctionData[auctionId];
        if (a.seller == address(0)) revert AuctionInactive();
        if (a.settled || a.cancelled) revert AlreadySettled();
        if (block.timestamp < a.endTime) revert AuctionNotEnded();
        if (a.highestBidder == address(0)) revert NoBidsToSettle();
        if (msg.sender != a.highestBidder) revert NotWinner();

        a.settled = true;
        address seller = a.seller;
        address winner = a.highestBidder;
        uint256 pay = a.highestBid;
        address nftC = a.nftContract;
        uint256 tid = a.tokenId;

        nftToken2AuctionId[nftC][tid] = 0;

        IERC721(nftC).safeTransferFrom(address(this), winner, tid);

        (bool paid, ) = payable(seller).call{value: pay}("");
        if (!paid) revert ETHPaySellerFailed();

        emit AuctionEnded(auctionId, winner, pay, seller);
    }

    /**
     * @notice 已结束但无人出价时，卖家取回 NFT。
     */
    function claimUnsoldNFT(uint256 auctionId) external nonReentrant {
        Auction storage a = auctionData[auctionId];
        if (a.seller == address(0)) revert AuctionInactive();
        if (a.settled || a.cancelled) revert AlreadySettled();
        if (block.timestamp < a.endTime) revert AuctionNotEnded();
        if (a.highestBidder != address(0)) revert NotSellerOrHasBids();
        if (msg.sender != a.seller) revert NotSeller();

        a.settled = true;
        address nftC = a.nftContract;
        uint256 tid = a.tokenId;

        nftToken2AuctionId[nftC][tid] = 0;

        IERC721(nftC).safeTransferFrom(address(this), a.seller, tid);
    }

    /**
     * @notice 在 startTime 之前卖家可取消，NFT 退回卖家。
     */
    function cancelAuction(uint256 auctionId) external nonReentrant {
        Auction storage a = auctionData[auctionId];
        if (a.seller == address(0)) revert AuctionInactive();
        if (a.settled || a.cancelled) revert AlreadySettled();
        if (block.timestamp >= a.startTime) revert CannotCancel();
        if (msg.sender != a.seller) revert NotSeller();

        a.cancelled = true;
        address nftC = a.nftContract;
        uint256 tid = a.tokenId;

        nftToken2AuctionId[nftC][tid] = 0;

        IERC721(nftC).safeTransferFrom(address(this), a.seller, tid);

        emit AuctionCancelled(auctionId, a.seller);
    }

    receive() external payable {}
}
