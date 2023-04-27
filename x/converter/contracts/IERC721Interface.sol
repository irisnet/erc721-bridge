// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

interface IERC721Base {
    function name() external view returns (string memory);

    function symbol() external view returns (string memory);
}

interface IERC721Common {
    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function baseURI() external view returns (string memory);
}

interface IERC721PresetMinterPauser {
    function name() external view returns (string memory);

    function symbol() external view returns (string memory);

    function baseURI() external view returns (string memory);

    function classData() external view returns (string memory);

    function tokenData(uint256 tokenId) external view returns (string memory);
}
