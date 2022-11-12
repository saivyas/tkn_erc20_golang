// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8.0;

import "../openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

contract GoToken is ERC20 {
    constructor(uint256 totalSupply) ERC20("GoToken", "GOT") {
        _mint(msg.sender, totalSupply);
    }
}
