// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.21;

import "internal/contract/dependencies/@openzeppelin/contracts/utils/Counters.sol";
import "internal/contract/dependencies/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "internal/contract/dependencies/@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";

contract MyWrappedNFT is ERC721URIStorage, ERC721Enumerable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    constructor() ERC721("My Wrapped NFT", "MWNFT") {}

    function createToken(address owner, string memory _tokenURI) public returns (uint) {
        _tokenIds.increment();
        uint256 newItemId = _tokenIds.current();
        _mint(owner, newItemId);
        _setTokenURI(newItemId, _tokenURI); // Updated to use `_tokenURI`
        return newItemId;
    }

    function tokenURI(
        uint256 tokenId
    ) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        virtual
        override(ERC721Enumerable, ERC721URIStorage)
        returns (bool)
    {
        return (ERC721Enumerable.supportsInterface(interfaceId) ||
            ERC721URIStorage.supportsInterface(interfaceId)); // Checks URI storage interfaces
    }

    // Resolve the _update conflict
    function _update(
        address to,
        uint256 tokenId,
        address auth
    ) internal override(ERC721, ERC721Enumerable) returns (address) {
        return super._update(to, tokenId, auth);
    }

    // Override _increaseBalance (if needed)
    function _increaseBalance(
        address account,
        uint128 value
    ) internal virtual override(ERC721, ERC721Enumerable) {
        super._increaseBalance(account, value);
    }
}
