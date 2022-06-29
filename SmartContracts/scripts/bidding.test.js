const { ethers } = require('hardhat');
async function delay(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

async function main() {
    const [owner, bidder1, bidder2] = await ethers.getSigners();
    const BiddingWar = await ethers.getContractFactory('BiddingWar');
    const deployedContract = await BiddingWar.deploy();
    await deployedContract.deployed();

    console.log("owner", owner.address, "contract address:", deployedContract.address);
    console.log("current balance bidder1", ethers.utils.formatUnits(await bidder1.getBalance()));
    console.log("current balance bidder2", ethers.utils.formatUnits(await bidder2.getBalance()));

    await deployedContract.connect(owner).setGameParams(10, 2, 50); //update params for develop mode
    await deployedContract.connect(owner).start();    
    await deployedContract.connect(bidder1).bid({value: ethers.utils.parseEther("0.5")});

    console.log("--first bidding");
    console.log("total bidAmount",  ethers.utils.formatUnits(await deployedContract.totalWinnerAmount()));    
    ethers.provider.on("block", async (blockNum)=> {
        console.log(blockNum+ ": " + (await ethers.provider.getBlock(blockNum)).timestamp )
    })
    await delay(3000);
    console.log("--second bidding");
    await deployedContract.connect(bidder2).bid({value: ethers.utils.parseEther("0.55")});
    console.log("total bidAmount",  ethers.utils.formatUnits(await deployedContract.totalWinnerAmount()));
    await delay(8000);
    await deployedContract.connect(owner).rewardToWinner();
    console.log("current balance bidder2", ethers.utils.formatUnits(await bidder2.getBalance()));
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.log(err);
        process.exit(1);
    });
