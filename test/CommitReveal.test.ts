
import { ethers } from "hardhat";
import { expect } from "chai";
import { CommitReveal } from "../typechain";

describe("CommitReveal", async () => {

  const PHASE_LENGTH_IN_SECONDS = 120
  let commitReveal: CommitReveal
  beforeEach("set up contract", async () => {
    const CommitRevealContractFactory = await ethers.getContractFactory("CommitReveal");
    commitReveal = (await CommitRevealContractFactory.deploy(PHASE_LENGTH_IN_SECONDS, "YES", "NO")) as CommitReveal;
  });

  it("should record a proper commit", async () => {
    const vote = "1~mybigsecret";
    const commit = ethers.encodeBytes32String(vote)
    await commitReveal.commitVote(commit);
    const commitsArray = await commitReveal.getVoteCommitsArray();
    const voteStatus = await commitReveal.voteStatuses(commit);
    const numberOfVotes = await commitReveal.numberOfVotesCast();
    expect(commitsArray[0]).to.equal(commit);
    expect(voteStatus).to.equal("Committed");
    expect(Number(numberOfVotes) == 1).to.be.true;
  });

  it("should record proper reveals", async () => {
    const vote1 = "1~mybigsecret";
    const vote2 = "2~mybigsecret";
    const commit1 = ethers.keccak256(ethers.toUtf8Bytes(vote1));
    const commit2 = ethers.keccak256(ethers.toUtf8Bytes(vote2));
    await commitReveal.commitVote(commit1);
    await commitReveal.commitVote(commit2);
    console.log('asdf')
    await ethers.provider.send("evm_increaseTime", [PHASE_LENGTH_IN_SECONDS]);
    await commitReveal.revealVote(vote1, commit1);
    await commitReveal.revealVote(vote2, commit2);
    const votesFor1 = await commitReveal.votesForChoice1();
    const votesFor2 = await commitReveal.votesForChoice2();
    const voteStatusCommit2 = await commitReveal.voteStatuses(commit2);
    const voteStatusCommit1 = await commitReveal.voteStatuses(commit1);
    expect(Number(votesFor2) == 1).to.be.true;
    expect(Number(votesFor1) == 1).to.be.true;
    expect(voteStatusCommit2).to.equal("Revealed");
    expect(voteStatusCommit1).to.equal("Revealed");
  });

  it("should not record an incorrect reveal", async () => {
    const badVote = "3~mybigsecret";
    const commit = ethers.encodeBytes32String(badVote);
    await commitReveal.commitVote(commit);
    await ethers.provider.send("evm_increaseTime", [PHASE_LENGTH_IN_SECONDS]);
    const tx = commitReveal.revealVote(badVote, commit);
    await expect(tx).to.be.rejectedWith(Error);
  });

  it("should correctly get winner", async () => {
    const vote = "1~mybigsecret";
    const commit = ethers.keccak256(ethers.toUtf8Bytes(vote));
    await commitReveal.commitVote(commit);
    await ethers.provider.send("evm_increaseTime", [PHASE_LENGTH_IN_SECONDS]);
    await commitReveal.revealVote(vote, commit);
    const winner = await commitReveal.getWinner();
    expect(winner).to.equal("YES");
  });
});
