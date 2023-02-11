require("@nomiclabs/hardhat-ethers");
require("@nomiclabs/hardhat-waffle");

require("dotenv").config()

const ALCHEMY_API_KEY = process.env.API_KEY;
const GOERLI_PRIVATE_KEY = process.env.PRIVATE_KEY; 

module.exports = {
  solidity: "0.8.17",
  networks: {
      goerli: {
      url: `https://eth-goerli.alchemyapi.io/v2/${ALCHEMY_API_KEY}`,
      accounts: [GOERLI_PRIVATE_KEY]
    }
  }
};
