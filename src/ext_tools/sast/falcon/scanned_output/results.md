'npx hardhat clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/golom-protocol-contracts)
'npx hardhat clean --global' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/golom-protocol-contracts)
'npx hardhat compile --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/golom-protocol-contracts)

Reentrancy in VoteEscrow._removeTokenFrom(address,uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#265-278):
	External calls:
	- this.removeDelegation(_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#270)
Reentrancy in VoteEscrow._transferFrom(address,address,uint256,address) (contracts/vote-escrow/VoteEscrowDelegation.sol#243-261):
	External calls:
	- _removeTokenFrom(_from,_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#254)
		- this.removeDelegation(_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#270)
Reference:  
