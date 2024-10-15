// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract LAIRegistration is Ownable {
    
    uint256 fee = 0.001 ether;

    struct UserInfo {
            string apiKey;
            address userWallet;
        }
        constructor() Ownable(address(this)) { }

    mapping(address => string) public apiKeys;
    mapping(uint256 => UserInfo) public apiKeysTG;

    event RegistrationSuccessfulUI(address indexed wallet, string key);
    event RegistrationSuccessfulTG(address indexed wallet, string key, uint256 tgId);

  
function registration(string memory key, uint256 tgId) public payable {
        require(msg.value >= fee, "Fee is required");
        (bool sent, ) = owner().call{value: msg.value}("");
        require(sent, "Failed to send ETH to the owner");
        
if (tgId == 0) {
        apiKeys[msg.sender] = key;
        emit RegistrationSuccessfulUI(msg.sender, key);
} else {
        apiKeys[msg.sender] = key;
        apiKeysTG[tgId].apiKey = key;
        apiKeysTG[tgId].userWallet = msg.sender;
        emit RegistrationSuccessfulTG(msg.sender, key, tgId);
}
    }



    function withdraw() public onlyOwner {
        payable(owner()).transfer(address(this).balance);
    }

    function changeFee(uint256 _newFee) public onlyOwner {
        fee = _newFee;
    }


   function getUIApiKey(address wallet) public view returns (string memory) {
        return apiKeys[wallet];
    }

    
    function getTGApiKey(uint256 tgId) public view returns (string memory, address) {
        UserInfo memory userInfo = apiKeysTG[tgId];
        return (userInfo.apiKey, userInfo.userWallet);
    }

}