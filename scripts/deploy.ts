import { ethers } from "hardhat";

async function main() {
  const ComomitReveal = await ethers.getContractFactory("CommitReveal")
  const commitReveal = await ComomitReveal.deploy(120, "YES", "NO");

  console.log(
    `Deployed CommitReveal at ${commitReveal.target}`
  );
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
