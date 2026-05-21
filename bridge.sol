// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


import "internal/blockchain/contract/dependencies/openzeppelin-contracts/contracts/token/ERC721/IERC721.sol";
import "internal/blockchain/contract/dependencies/openzeppelin-contracts/contracts/token/ERC721/utils/ERC721Holder.sol";

contract SimpleNFTTransfer is ERC721Holder {
    // keccak256("tokenURI(uint256)") => 0xc87b56dd
    bytes4 private constant TOKEN_URI_SELECTOR = 0xc87b56dd;

    event TokenLocked(
        uint tokenId,
        address sender,
        address addressInChainB,
        string tokenURIValue
    );

    /**
     * @notice Transfer any NFT to this contract and emit tokenURI (if available)
     * @dev User must approve this contract first using ERC721.approve()
     */
    function transferNFTToContract(
        address _nftContract,
        address ownerAddressInChainB,
        uint256 _tokenId
    ) external returns (bool) {
        IERC721 nft = IERC721(_nftContract);

        // Ensure caller owns the NFT
        require(nft.ownerOf(_tokenId) == msg.sender, "Not NFT owner");

        // Try to read tokenURI via low-level call; empty string if unsupported/reverts
        string memory uri = _safeTokenURI(_nftContract, _tokenId);

        // Transfer NFT to this contract
        nft.safeTransferFrom(msg.sender, address(this), _tokenId);

        // Emit event with tokenURI (empty if not supported)
        emit TokenLocked(_tokenId, msg.sender, ownerAddressInChainB, uri);
        return true;
    }

    /**
     * @notice Transfer NFT from this contract to any address
     */
    function transferNFTFromContract(
        address _nftContract,
        address _to,
        uint256 _tokenId
    ) external returns (bool) {
        IERC721 nft = IERC721(_nftContract);

        // Check if contract owns the NFT
        require(nft.ownerOf(_tokenId) == address(this), "Contract doesn't own this NFT");

        // Transfer NFT to destination
        nft.safeTransferFrom(address(this), _to, _tokenId);

        return true;
    }

    /**
     * @notice Check NFT ownership
     */
    function checkNFTOwnership(
        address _nftContract,
        uint256 _tokenId
    ) external view returns (address) {
        return IERC721(_nftContract).ownerOf(_tokenId);
    }

    /**
     * @dev Low-level, interface-free read of tokenURI.
     * Returns "" if the function doesn't exist or reverts.
     */
    function _safeTokenURI(address _nftContract, uint256 _tokenId) internal view returns (string memory) {
        (bool ok, bytes memory data) =
            _nftContract.staticcall(abi.encodeWithSelector(TOKEN_URI_SELECTOR, _tokenId));
        if (!ok || data.length == 0) {
            return "";
        }
        // If the contract returned something unexpected, abi.decode will revert.
        // That's fine — we guard by only calling this on read; callers can handle empty string.
        try this._decodeString(data) returns (string memory s) {
            return s;
        } catch {
            return "";
        }
    }

    // Helper to enable try/catch around decode (must be external for try/catch)
    function _decodeString(bytes calldata data) external pure returns (string memory) {
        return abi.decode(data, (string));
    }
}
