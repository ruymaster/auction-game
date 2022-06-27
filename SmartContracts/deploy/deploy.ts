import {HardhatRuntimeEnvironment} from 'hardhat/types';
import {DeployFunction} from 'hardhat-deploy/types';
const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  const {deployments, getNamedAccounts, ethers} = hre;
  const {deploy} = deployments;
  const signer = process.env.SIGNER_ADDRESS || '0x80aEa81791Ded20568221346C79B0ad4E0890FAA';
  const {deployer} = await getNamedAccounts();
  const contract = await deploy('BiddingWar', {
      from: deployer,
      args: [],
      log: true,
    });
  await delay(60000);
};

function delay(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export default func;
func.tags = ['all'];
