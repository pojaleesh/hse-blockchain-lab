const { expect } = require("chai");
const { ethers } = require("hardhat");


describe("PojaleeshToken.sol", () => {
    let contract;
    let owner, firstUser, secondUser;
    let initialSupply = 100000;
    let ownerAddress;
    let firstAddress;
    let secondAddress;

    beforeEach(async () => {
        let contractFactory = await ethers.getContractFactory("PojaleeshToken");
        contract = await contractFactory.deploy(initialSupply);
        [owner, firstUser, secondUser] = await ethers.getSigners();

        ownerAddress = await owner.getAddress();
        firstAddress = await firstUser.getAddress();
        secondAddress = await secondUser.getAddress();
    });

    describe("Init", () => {
        it("Correct Name", async () => {
            expect(await contract.name()).to.equal("Pojaleesh");
        });

        it("Correct supply", async () => {
            expect(await contract.totalSupply()).to.equal(initialSupply);
            expect(await contract.balanceOf(ownerAddress)).to.equal(initialSupply);
        });
    });

    describe("Core", () => {
        it("Correct transfer two sides", async () => {
            await contract.transfer(firstAddress, 1337);

            expect(await contract.balanceOf(secondAddress)).to.equal(0);
            expect(await contract.balanceOf(firstAddress)).to.equal(1337);

            await contract.connect(firstUser).transfer(secondAddress, 1337);

            expect(await contract.balanceOf(secondAddress)).to.equal(1337);
            expect(await contract.balanceOf(firstAddress)).to.equal(0);
        });

        it("Correct mint", async () => {
            await expect(contract.connect(firstUser).mint(firstAddress, 1))
                .to.be.revertedWith("Ownable: caller is not the owner");

            await contract.mint(firstAddress, 1337);

            expect(await contract.totalSupply()).to.equal(initialSupply + 1337);
            expect(await contract.balanceOf(secondAddress)).to.equal(0);
            expect(await contract.balanceOf(firstAddress)).to.equal(1337);
            expect(await contract.balanceOf(ownerAddress)).to.equal(initialSupply);

        });

        it('Correct approve and transferFrom', async function () {
            await contract.transfer(firstAddress, 1337);
            expect(await contract.balanceOf(firstAddress)).to.be.equal(1337);

            await contract.connect(firstUser).approve(ownerAddress, 1337);
            await contract.transferFrom(firstAddress, ownerAddress, 1337);

            expect(await contract.balanceOf(ownerAddress)).to.equal(initialSupply);
            expect(await contract.balanceOf(firstAddress)).to.equal(0);
            expect(await contract.balanceOf(secondAddress)).to.equal(0);
        });
    });
});
