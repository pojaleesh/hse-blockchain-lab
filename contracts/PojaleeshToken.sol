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

    struct Actor {
        string id;
        int firstParam;
        int secondParam;
        int thirdParam;
    }

    mapping(string => Actor) contractMapping;
    event ActorAdded(string id, int firstParam, int secondParam, int thirdParam);
    event ActorDeleted(string id);

    function addToActor(string memory id, int firstParam, int secondParam, int thirdParam) public {
        contractMapping[id] = Actor(id, firstParam, secondParam, thirdParam);
        emit ActorAdded(id, firstParam, secondParam, thirdParam);
    }

    function delFromActor(string memory id) public {
        delete contractMapping[id];
        emit ActorDeleted(id);
    }

}
