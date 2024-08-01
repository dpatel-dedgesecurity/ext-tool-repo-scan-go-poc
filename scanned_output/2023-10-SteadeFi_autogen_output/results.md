'npx hardhat clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/golom-protocol-contracts)
'npx hardhat clean --global' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/golom-protocol-contracts)
Problem executing hardhat: You are using a version of Node.js that is not supported by Hardhat, and it may work incorrectly, or not work at all.

Please, make sure you are using a supported version of Node.js.

To learn more about which versions of Node.js are supported go to https://hardhat.org/nodejs-versions

'npx hardhat compile --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/golom-protocol-contracts)

Reentrancy in VoteEscrow._removeTokenFrom(address,uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#265-278):
	External calls:
	- this.removeDelegation(_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#270)
Reentrancy in VoteEscrow._transferFrom(address,address,uint256,address) (contracts/vote-escrow/VoteEscrowDelegation.sol#243-261):
	External calls:
	- _removeTokenFrom(_from,_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#254)
		- this.removeDelegation(_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#270)
Reference:  
