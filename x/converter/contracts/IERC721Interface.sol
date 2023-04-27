// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IERC721PresetMinterPauser {
    function baseURI() external view returns (string memory);

    function classData() external view returns (string memory);

    function tokenData(uint256 tokenId) external view returns (string memory);
}
