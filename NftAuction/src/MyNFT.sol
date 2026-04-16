// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

// ==================== NFT 合约 ====================
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721, Ownable {
    uint256 private _tokenIds;

    // constructor() ERC721("MyNFT", "MNFT") {}
    constructor() ERC721("MyNFT", "MNFT") Ownable(msg.sender) {}

    /// @notice 铸造新NFT，仅合约拥有者可调用
    function mint(address to) external onlyOwner returns (uint256) {
        require(to != address(0), "MyNFT: mint to the zero address");
        _tokenIds++;
        uint256 newTokenId = _tokenIds;
        _safeMint(to, newTokenId);
        return newTokenId;
    }
}
