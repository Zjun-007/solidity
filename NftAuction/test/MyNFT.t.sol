// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import {Test} from "forge-std/Test.sol";
import {MyNFT} from "../src/MyNFT.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract MyNFTTest is Test {
    MyNFT public nft; // 声明一个公共的 MyNFT 合约变量
    address public owner;
    address public user1;
    address public user2;

    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);

    // 每个测试用例前执行，部署合约并初始化测试账户
    function setUp() public {
        owner = address(this); // 测试合约本身就是部署者，也是所有者
        user1 = address(0x1);
        user2 = address(0x2);

        // 部署 NFT 合约
        nft = new MyNFT();
    }

    // 测试构造函数设置正确的名称和符号
    function testNameAndSymbol() public view {
        assertEq(nft.name(), "MyNFT");
        assertEq(nft.symbol(), "MNFT");
    }

    // 测试只有所有者可以铸造
    function testMintOnlyOwner() public {
        // 非所有者尝试铸造，应该失败
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, user1));
        nft.mint(user1);

        // 所有者铸造成功
        uint256 tokenId = nft.mint(user1);
        assertEq(tokenId, 1); // 第一个 tokenId 应为 1
        assertEq(nft.ownerOf(tokenId), user1);
        assertEq(nft.balanceOf(user1), 1);
    }

    // 测试连续铸造 tokenId 递增
    function testMintIncrementsTokenId() public {
        uint256 firstToken = nft.mint(user1);
        assertEq(firstToken, 1);

        uint256 secondToken = nft.mint(user2);
        assertEq(secondToken, 2);

        assertEq(nft.ownerOf(1), user1);
        assertEq(nft.ownerOf(2), user2);
    }

    // 测试 Transfer 事件在铸造时被触发
    function testMintEmitsTransferEvent() public {
        vm.expectEmit(true, true, true, true);
        // 预期的事件：Transfer(0x0, user1, tokenId)
        emit Transfer(address(0), user1, 1);
        nft.mint(user1);
    }

    // 测试转移功能：所有者或授权账户可以转移
    function testTransfer() public {
        uint256 tokenId = nft.mint(user1);

        // user1 将自己的 NFT 转移给 user2
        vm.prank(user1);
        nft.transferFrom(user1, user2, tokenId);

        assertEq(nft.ownerOf(tokenId), user2);
        assertEq(nft.balanceOf(user1), 0);
        assertEq(nft.balanceOf(user2), 1);
    }

    // 测试安全转账 safeTransferFrom
    function testSafeTransfer() public {
        uint256 tokenId = nft.mint(user1);

        vm.prank(user1);
        nft.safeTransferFrom(user1, user2, tokenId);

        assertEq(nft.ownerOf(tokenId), user2);
    }

    // 测试 Transfer 事件在转移时被触发
    function testTransferEmitsTransferEvent() public {
        uint256 tokenId = nft.mint(user1);

        vm.expectEmit(true, true, true, true);
        emit Transfer(user1, user2, tokenId);

        vm.prank(user1);
        nft.transferFrom(user1, user2, tokenId);
    }

    // 测试未授权转移失败
    function testUnauthorizedTransferFails() public {
        uint256 tokenId = nft.mint(user1);

        // user2 尝试转移 user1 的 NFT，应该失败
        vm.prank(user2);
        vm.expectRevert("ERC721: caller is not token owner or approved");
        nft.transferFrom(user1, user2, tokenId);
    }

    // 测试批准和转移
    function testApprove() public {
        uint256 tokenId = nft.mint(user1);

        // user1 授权 user2 操作该 token
        vm.prank(user1);
        nft.approve(user2, tokenId);

        // user2 现在可以转移
        vm.prank(user2);
        nft.transferFrom(user1, user2, tokenId);

        assertEq(nft.ownerOf(tokenId), user2);
    }

    // 测试 setApprovalForAll
    function testSetApprovalForAll() public {
        uint256 tokenId1 = nft.mint(user1);
        uint256 tokenId2 = nft.mint(user1);

        // user1 授权 user2 管理所有 token
        vm.prank(user1);
        nft.setApprovalForAll(user2, true);

        // user2 可以转移任意 token
        vm.prank(user2);
        nft.transferFrom(user1, user2, tokenId1);

        vm.prank(user2);
        nft.transferFrom(user1, user2, tokenId2);

        assertEq(nft.ownerOf(tokenId1), user2);
        assertEq(nft.ownerOf(tokenId2), user2);
    }

    // 测试无法向零地址铸造
    function testCannotMintToZeroAddress() public {
        vm.expectRevert("MyNFT: mint to the zero address");
        nft.mint(address(0));
    }

    // 测试 tokenURI 默认实现（可选，如果合约重写了 _baseURI 可验证）
    // 当前合约未重写 tokenURI，因此返回空字符串
    function testTokenURIDefault() public view {
        uint256 tokenId = 1; // 假设已铸造
        // 如果不重写，默认 tokenURI 会尝试拼接 baseURI，而 baseURI 为空
        // 这里仅验证调用不会 revert
        nft.tokenURI(tokenId); // 应返回空字符串
    }

}

