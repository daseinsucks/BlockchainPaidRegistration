// SPDX-License-Identifier: MIT
const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("LAIRegistration", function () {
    let LAIRegistration, laiRegistration, owner, addr1, addr2;

    const fee = ethers.utils.parseEther("0.001");

    beforeEach(async function () {
        LAIRegistration = await ethers.getContractFactory("LAIRegistration");
        laiRegistration = await LAIRegistration.deploy();
        await laiRegistration.deployed();

        [owner, addr1, addr2] = await ethers.getSigners();
    });

    describe("Registration", function () {
        it("Should register user with correct fee and emit event", async function () {
            await expect(
                laiRegistration.connect(addr1).registration("testApiKey", 0, { value: fee })
            )
            .to.emit(laiRegistration, "RegistrationSuccessfulUI")
            .withArgs(addr1.address, "testApiKey");

            const apiKey = await laiRegistration.getUIApiKey(addr1.address);
            expect(apiKey).to.equal("testApiKey");
        });

        it("Should register user with TG ID and emit event", async function () {
            await expect(
                laiRegistration.connect(addr1).registration("testApiKeyTG", 123, { value: fee })
            )
            .to.emit(laiRegistration, "RegistrationSuccessfulTG")
            .withArgs(addr1.address, "testApiKeyTG", 123);

            const apiKey = await laiRegistration.getTGApiKey(123);
            expect(apiKey[0]).to.equal("testApiKeyTG");
            expect(apiKey[1]).to.equal(addr1.address);
        });

        it("Should require fee for registration", async function () {
            await expect(
                laiRegistration.connect(addr1).registration("testApiKeyNoFee", 0)
            ).to.be.revertedWith("Fee is required");
        });

        it("Should transfer ETH to the owner", async function () {
            const initialOwnerBalance = await ethers.provider.getBalance(owner.address);
            await laiRegistration.connect(addr1).registration("testApiKey", 0, { value: fee });
            const finalOwnerBalance = await ethers.provider.getBalance(owner.address);

            expect(finalOwnerBalance).to.equal(initialOwnerBalance.add(fee));
        });
    });

    describe("Owner functions", function () {
        it("Should allow the owner to withdraw funds", async function () {
            await laiRegistration.connect(addr1).registration("testApiKey", 0, { value: fee });
            const initialOwnerBalance = await ethers.provider.getBalance(owner.address);

            await laiRegistration.withdraw();
            const finalOwnerBalance = await ethers.provider.getBalance(owner.address);

            expect(finalOwnerBalance).to.equal(initialOwnerBalance.add(fee));
        });

        it("Should allow the owner to change the fee", async function () {
            await laiRegistration.changeFee(ethers.utils.parseEther("0.002"));
            const newFee = await laiRegistration.fee();
            expect(newFee).to.equal(ethers.utils.parseEther("0.002"));
        });
    });
});
