// SPDX-License-Identifier: Academic Free License v3.0
pragma solidity >=0.8.0 <0.9.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract PojaleeshToken is Ownable, ERC20 {
    constructor(uint256 amount) ERC20("Pojaleesh", "POJ") {
        _mint(msg.sender, amount);
    }

    function mint(address account, uint256 amount) public onlyOwner {
        _mint(account, amount);
    }
}
