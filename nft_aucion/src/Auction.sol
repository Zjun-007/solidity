// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/**
 * @title Auction
 * @dev 支持 ERC721 NFT 的拍卖，竞拍者可使用 ETH 或指定 ERC20 代币出价。
 *      卖家创建拍卖时需授权本合约转移 NFT。拍卖结束后，NFT 归最高出价者，
 *      资金归卖家，其他竞拍者可取回自己的出价。
 */
contract Auction is ReentrancyGuard {
    struct AuctionInfo {
        address seller;               // 卖家地址
        address nftContract;           // NFT 合约地址
        uint256 tokenId;               // NFT tokenId
        address biddingToken;          // 出价代币地址（address(0) 表示 ETH）
        uint256 startPrice;            // 起拍价
        uint256 endTime;               // 拍卖结束时间戳
        address highestBidder;         // 当前最高出价者
        uint256 highestBid;            // 当前最高出价金额
        bool ended;                    // 是否已结束（已结算）
        mapping(address => uint256) bids; // 每个地址的总出金额（用于退款）
    }

    // 以 NFT 合约地址和 tokenId 组合作为拍卖的唯一标识
    mapping(bytes32 => AuctionInfo) public auctions;

    event AuctionCreated(
        bytes32 indexed auctionId,
        address indexed seller,
        address nftContract,
        uint256 tokenId,
        address biddingToken,
        uint256 startPrice,
        uint256 endTime
    );
    event BidPlaced(
        bytes32 indexed auctionId,
        address indexed bidder,
        uint256 amount,
        address biddingToken
    );
    event AuctionEnded(
        bytes32 indexed auctionId,
        address winner,
        uint256 amount
    );
    event Withdrawn(
        bytes32 indexed auctionId,
        address indexed bidder,
        uint256 amount
    );

    modifier auctionExists(bytes32 auctionId) {
        require(auctions[auctionId].seller != address(0), "Auction does not exist");
        _;
    }

    modifier auctionNotEnded(bytes32 auctionId) {
        require(block.timestamp < auctions[auctionId].endTime, "Auction already ended");
        _;
    }

    modifier auctionEnded(bytes32 auctionId) {
        require(block.timestamp >= auctions[auctionId].endTime, "Auction not ended yet");
        _;
    }

    modifier notFinalized(bytes32 auctionId) {
        require(!auctions[auctionId].ended, "Auction already finalized");
        _;
    }

    /**
     * @notice 创建新拍卖
     * @param nftContract NFT 合约地址
     * @param tokenId NFT tokenId
     * @param biddingToken 出价代币地址（使用 address(0) 表示 ETH）
     * @param startPrice 起拍价（大于 0）
     * @param duration 拍卖持续时间（秒）
     * @return auctionId 拍卖唯一标识
     */
    function createAuction(
        address nftContract,
        uint256 tokenId,
        address biddingToken,
        uint256 startPrice,
        uint256 duration
    ) external returns (bytes32 auctionId) {
        require(startPrice > 0, "Start price must be > 0");
        require(duration > 0, "Duration must be > 0");

        IERC721 nft = IERC721(nftContract);
        // 检查调用者是否为 NFT 所有者
        require(nft.ownerOf(tokenId) == msg.sender, "Not owner of NFT");
        // 检查本合约是否被授权转移该 NFT（approve 或 setApprovalForAll）
        require(
            nft.getApproved(tokenId) == address(this) ||
            nft.isApprovedForAll(msg.sender, address(this)),
            "Contract not approved to transfer NFT"
        );

        bytes32 key = keccak256(abi.encodePacked(nftContract, tokenId));
        // 确保同一 NFT 没有活跃的拍卖（允许重新上架，但需等前一个结束）
        require(
            auctions[key].seller == address(0) || auctions[key].ended,
            "Auction already exists for this NFT"
        );

        AuctionInfo storage auction = auctions[key];
        auction.seller = msg.sender;
        auction.nftContract = nftContract;
        auction.tokenId = tokenId;
        auction.biddingToken = biddingToken;
        auction.startPrice = startPrice;
        auction.endTime = block.timestamp + duration;
        auction.ended = false;

        emit AuctionCreated(key, msg.sender, nftContract, tokenId, biddingToken, startPrice, auction.endTime);
        return key;
    }

    /**
     * @notice 使用 ETH 出价
     * @param auctionId 拍卖标识
     */
    function bidETH(bytes32 auctionId)external payable auctionExists(auctionId) auctionNotEnded(auctionId) notFinalized(auctionId) nonReentrant{
        AuctionInfo storage auction = auctions[auctionId];
        require(auction.biddingToken == address(0), "This auction uses ERC20");
        // require(msg.value > auction.highestBid, "Bid too low");
        require(
            msg.value >= auction.startPrice && auction.highestBid == 0 ||
            msg.value > auction.highestBid,
            "Bid must exceed current highest"
        );

        // 退还前一位最高出价者的全部出价
        if (auction.highestBidder != address(0)) {
            _refund(auctionId, auction.highestBidder);
        }

        // 更新最高出价
        auction.highestBidder = msg.sender;
        auction.highestBid = msg.value;
        auction.bids[msg.sender] += msg.value;

        emit BidPlaced(auctionId, msg.sender, msg.value, address(0));
    }

    /**
     * @notice 使用 ERC20 代币出价
     * @param auctionId 拍卖标识
     * @param amount 出价金额（需要已授权本合约）
     */
    function bidERC20(bytes32 auctionId, uint256 amount)external auctionExists(auctionId) auctionNotEnded(auctionId) notFinalized(auctionId) nonReentrant{
        AuctionInfo storage auction = auctions[auctionId];
        require(auction.biddingToken != address(0), "This auction uses ETH");
        // require(amount > auction.highestBid, "Bid too low");
        require(
            amount >= auction.startPrice && auction.highestBid == 0 ||
            amount > auction.highestBid,
            "Bid must exceed current highest"
        );

        // 从出价者账户转出代币到本合约
        IERC20 token = IERC20(auction.biddingToken);
        require(token.transferFrom(msg.sender, address(this), amount), "Transfer failed");

        // 退还前一位最高出价者
        if (auction.highestBidder != address(0)) {
            _refund(auctionId, auction.highestBidder);
        }

        auction.highestBidder = msg.sender;
        auction.highestBid = amount;
        auction.bids[msg.sender] += amount;

        emit BidPlaced(auctionId, msg.sender, amount, auction.biddingToken);
    }

    /**
     * @notice 内部函数：退还指定竞拍者的全部出价
     */
    function _refund(bytes32 auctionId, address bidder) private {
        AuctionInfo storage auction = auctions[auctionId];
        uint256 amount = auction.bids[bidder];
        require(amount > 0, "No bid to refund");
        auction.bids[bidder] = 0; //指定竞拍者的出价余额清零,CEI模式

        if (auction.biddingToken == address(0)) {
            (bool success, ) = payable(bidder).call{value: amount}("");
            require(success, "ETH refund failed");
        } else {
            // IERC20(auction.biddingToken).transfer(bidder, amount);
            require(IERC20(auction.biddingToken).transfer(bidder, amount), "ERC20 refund failed");
        }
        emit Withdrawn(auctionId, bidder, amount);
    }

    /**
     * @notice 结束拍卖，将 NFT 转给最高出价者，资金转给卖家
     * @param auctionId 拍卖标识
     */
    function finalize(bytes32 auctionId)external auctionExists(auctionId) auctionEnded(auctionId) notFinalized(auctionId) nonReentrant{
        AuctionInfo storage auction = auctions[auctionId];
        auction.ended = true;

        address winner = auction.highestBidder;
        uint256 winningBid = auction.highestBid;
        address seller = auction.seller;

        require(winner != address(0), "No bids placed");

        auction.bids[winner] = 0; // 清除胜出者的出价记录（避免混淆）

        // 转移 NFT 给胜出者
        IERC721(auction.nftContract).transferFrom(seller, winner, auction.tokenId);

        // 将最高出价转给卖家
        if (auction.biddingToken == address(0)) {
            (bool success, ) = payable(seller).call{value: winningBid}("");
            require(success, "ETH transfer to seller failed");
        } else {
            // IERC20(auction.biddingToken).transfer(seller, winningBid);
            require(IERC20(auction.biddingToken).transfer(seller, winningBid), "ERC20 transfer failed");
        }

        emit AuctionEnded(auctionId, winner, winningBid);
    }

    /**
     * @notice 竞拍者取回自己的出价（仅限非胜出者，或在拍卖结束前被超过后）
     *         实际上，由于_refund的存在，非胜出者通常在过程中已被退款，所以withdraw可能很少用到，但保留作为安全网
     * @param auctionId 拍卖标识
     */
    function withdraw(bytes32 auctionId) external nonReentrant {
        AuctionInfo storage auction = auctions[auctionId];
        require(auction.ended || block.timestamp >= auction.endTime, "Auction still active");
        require(msg.sender != auction.highestBidder, "Winner cannot withdraw until finalize");

        uint256 amount = auction.bids[msg.sender];
        require(amount > 0, "No bid to withdraw");
        auction.bids[msg.sender] = 0;

        if (auction.biddingToken == address(0)) {
            (bool success, ) = payable(msg.sender).call{value: amount}("");
            require(success, "ETH withdraw failed");
        } else {
            // IERC20(auction.biddingToken).transfer(msg.sender, amount);
            // 修复：检查 ERC20 transfer 返回值，确保转账成功
            require(IERC20(auction.biddingToken).transfer(msg.sender, amount), "ERC20 withdraw failed");
        }
        emit Withdrawn(auctionId, msg.sender, amount);
    }

    /**
     * @notice 查询指定竞拍者的总出金额
     */
    function getBid(bytes32 auctionId, address user) external view returns (uint256) {
        return auctions[auctionId].bids[user];
    }
}