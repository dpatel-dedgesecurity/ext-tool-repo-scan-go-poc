'npx hardhat clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/golom-protocol-contracts)
'npx hardhat clean --global' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/golom-protocol-contracts)
'npx hardhat compile --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/golom-protocol-contracts)

Reentrancy in VoteEscrow._removeTokenFrom(address,uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#265-278):
	External calls:
	- this.removeDelegation(_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#270)
Reentrancy in VoteEscrow._transferFrom(address,address,uint256,address) (contracts/vote-escrow/VoteEscrowDelegation.sol#243-261):
	External calls:
	- _removeTokenFrom(_from,_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#254)
		- this.removeDelegation(_tokenId) (contracts/vote-escrow/VoteEscrowDelegation.sol#270)
Reference:  

state variable: ERC20Permit._nonces (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#25) not initialized and not written in contract but be used in contract
state variable: ERC20Votes._checkpoints (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#37) not initialized and not written in contract but be used in contract
state variable: ERC20Votes._totalSupplyCheckpoints (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#38) not initialized and not written in contract but be used in contract
state variable: ERC721Mock._tokenIds (contracts/test/ERC721Mock.sol#12) not initialized and not written in contract but be used in contract
state variable: VoteEscrowCore.token (contracts/vote-escrow/VoteEscrowCore.sol#300) not initialized and not written in contract but be used in contract
state variable: VoteEscrowCore.supportedInterfaces (contracts/vote-escrow/VoteEscrowCore.sol#344) not initialized and not written in contract but be used in contract
Reference:  

	- WETH.withdraw(uint256) (contracts/test/WETH.sol#25-30)
Reference:  

	- Timelock.executeTransaction(address,uint256,string,bytes,uint256) (contracts/governance/Timelock.sol#114-145)
Reference:  

	- GolomTreasury.withdraw(address,uint256) (contracts/governance/GolomTreasury.sol#19-22)
	- DummyRewardDistributor.withdrawTokens(address,address) (contracts/rewards/DummyRewardDistributor.sol#39-41)
	- DummyRewardDistributor.withdrawEth(address) (contracts/rewards/DummyRewardDistributor.sol#43-45)
Reference:  

Potential price manipulation risk:
	- In function addFee
		-- tokenToEmit = (dailyEmission * (rewardToken.totalSupply() - rewardToken.balanceOf(address(ve)))) / rewardToken.totalSupply() (contracts/rewards/RewardDistributor.sol#121-122) have potential price manipulated risk from tokenToEmit and call rewardToken.balanceOf(address(ve)) which could influence variable:tokenToEmit
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

public mint or burn found in GenesisClaim.burn() (contracts/rewards/GenesisClaim.sol#213-216)public mint or burn found in GolomAirdrop.burn() (contracts/rewards/GolomAirdrop.sol#215-218)public mint or burn found in ERC721Mock.mint(address) (contracts/test/ERC721Mock.sol#16-23)Reference: check public mint method

VoteEscrowCore._checkpoint(uint256,LockedBalance,LockedBalance).u_old (contracts/vote-escrow/VoteEscrowCore.sol#694) is a local variable never initialized
VoteEscrowCore._deposit_for(uint256,uint256,uint256,LockedBalance,VoteEscrowCore.DepositType).old_locked (contracts/vote-escrow/VoteEscrowCore.sol#843) is a local variable never initialized
ERC20Votes._moveVotingPower(address,address,uint256).oldWeight_scope_0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#220) is a local variable never initialized
VoteEscrowCore._checkpoint(uint256,LockedBalance,LockedBalance).u_new (contracts/vote-escrow/VoteEscrowCore.sol#695) is a local variable never initialized
VoteEscrow._getCurrentDelegated(uint256).myArray (contracts/vote-escrow/VoteEscrowDelegation.sol#130) is a local variable never initialized
VoteEscrow._getPriorDelegated(uint256,uint256).myArray (contracts/vote-escrow/VoteEscrowDelegation.sol#143) is a local variable never initialized
VoteEscrow.removeElement(uint256[],uint256).i (contracts/vote-escrow/VoteEscrowDelegation.sol#214) is a local variable never initialized
Reference:  

GovernorAlpha._queueOrRevert(address,uint256,string,bytes,uint256) (contracts/governance/GovernerBravo.sol#250-262)have ignores return value in timelock.queueTransaction(target,value,signature,data,eta) (contracts/governance/GovernerBravo.sol#261)
GovernorAlpha.execute(uint256) (contracts/governance/GovernerBravo.sol#264-281)have ignores return value in timelock.executeTransaction{value: proposal.values[i]}(proposal.targets[i],proposal.values[i],proposal.signatures[i],proposal.calldatas[i],proposal.eta) (contracts/governance/GovernerBravo.sol#272-278)
GovernorAlpha.__queueSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#408-411)have ignores return value in timelock.queueTransaction(address(timelock),0,setPendingAdmin(address),abi.encode(newPendingAdmin),eta) (contracts/governance/GovernerBravo.sol#410)
GovernorAlpha.__executeSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#413-416)have ignores return value in timelock.executeTransaction(address(timelock),0,setPendingAdmin(address),abi.encode(newPendingAdmin),eta) (contracts/governance/GovernerBravo.sol#415)
Reference:  

GenesisClaim.claim(uint256,bytes32[],bool) (contracts/rewards/GenesisClaim.sol#189-211)have ignores return value in ve.create_lock_for(amount,maxLockDuration,msg.sender) (contracts/rewards/GenesisClaim.sol#205)
GolomAirdrop.claim(uint256,bytes32[],IGolomTrader.Order,bool) (contracts/rewards/GolomAirdrop.sol#188-213)have ignores return value in ve.create_lock_for(amount,maxLockDuration,msg.sender) (contracts/rewards/GolomAirdrop.sol#207)
Reference:  

function:Timelock.fallback() (contracts/governance/Timelock.sol#152)is empty 
Reference:  

function:ERC1155._beforeTokenTransfer(address,address,address,uint256[],uint256[],bytes) (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#423-430)is empty 
function:ERC1155._afterTokenTransfer(address,address,address,uint256[],uint256[],bytes) (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#452-459)is empty 
function:ERC20._beforeTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#358-362)is empty 
function:ERC20._afterTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#378-382)is empty 
function:ERC721._beforeTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#425-429)is empty 
function:ERC721._afterTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#442-446)is empty 
function:GolomTrader.fallback() (contracts/core/GolomTrader.sol#528)is empty 
function:DummyRewardDistributor.addFee(address[2],uint256) (contracts/rewards/DummyRewardDistributor.sol#37)is empty 
function:DummyRewardDistributor.fallback() (contracts/rewards/DummyRewardDistributor.sol#47)is empty 
function:RewardDistributor.fallback() (contracts/rewards/RewardDistributor.sol#282)is empty 
Reference:  

TokenUriHelper._tokenURI(uint256,uint256,uint256,uint256).output (contracts/vote-escrow/TokenUriHelper.sol#71) is written in both
	output = string(abi.encodePacked(output,<text y="392px" x="447px" fill="white" font-family="Lexend Deca, sans-serif" font-weight="400" font-size="24px" text-anchor="end">,toString(_value),</text><mask id="mask0_2_190" style="mask-type:alpha" maskUnits="userSpaceOnUse" x="399" y="73" width="58" height="58"><path d="M456.543 102.091C456.543 117.884 443.74 130.686 427.948 130.686C412.155 130.686 399.352 117.884 399.352 102.091C399.352 86.2981 412.155 73.4955 427.948 73.4955C443.74 73.4955 456.543 86.2981 456.543 102.091ZM405.91 102.091C405.91 114.262 415.777 124.128 427.948 124.128C440.118 124.128 449.985 114.262 449.985 102.091C449.985 89.9201 440.118 80.0537 427.948 80.0537C415.777 80.0537 405.91 89.9201 405.91 102.091Z" fill="#C4C4C4"/></mask><g mask="url(#mask0_2_190)"><path d="M458.138 73.4955H396.962V133.076H458.138V73.4955Z" fill="#FD7A7A"/><path d="M396.962 76.7614H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 80.266H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 83.7708H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 87.2754H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 90.7802H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 94.2848H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 97.7897H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 101.294H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 104.799H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 108.304H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 111.808H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 115.313H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 118.818H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 122.323H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 125.827H458.138" stroke="black" stroke-width="0.328567"/><path d="M396.962 129.332H458.138" stroke="black" stroke-width="0.328567"/></g><mask id="mask1_2_190" style="mask-type:alpha" maskUnits="userSpaceOnUse" x="399" y="73" width="58" height="58"><path d="M456.543 102.091C456.543 117.884 443.74 130.686 427.948 130.686C412.155 130.686 399.352 117.884 399.352 102.091C399.352 86.2981 412.155 73.4955 427.948 73.4955C443.74 73.4955 456.543 86.2981 456.543 102.091ZM405.91 102.091C405.91 114.262 415.777 124.128 427.948 124.128C440.118 124.128 449.985 114.262 449.985 102.091C449.985 89.9201 440.118 80.0537 427.948 80.0537C415.777 80.0537 405.91 89.9201 405.91 102.091Z" fill="white"/></mask><g mask="url(#mask1_2_190)"><path d="M456.543 102.091C456.543 117.884 443.74 130.686 427.948 130.686C412.155 130.686 399.352 117.884 399.352 102.091C399.352 86.2981 412.155 73.4955 427.948 73.4955C443.74 73.4955 456.543 86.2981 456.543 102.091ZM405.91 102.091C405.91 114.262 415.777 124.128 427.948 124.128C440.118 124.128 449.985 114.262 449.985 102.091C449.985 89.9201 440.118 80.0537 427.948 80.0537C415.777 80.0537 405.91 89.9201 405.91 102.091Z" stroke="black" stroke-width="0.876179"/><path d="M422.132 130.606H428.026" stroke="black" stroke-width="0.109522"/><path d="M422.132 73.5752H428.026" stroke="black" stroke-width="0.109522"/></g><mask id="mask2_2_190" style="mask-type:alpha" maskUnits="userSpaceOnUse" x="393" y="73" width="58" height="58"><path d="M450.648 102.091C450.648 117.884 437.845 130.686 422.053 130.686C406.26 130.686 393.457 117.884 393.457 102.091C393.457 86.2981 406.26 73.4955 422.053 73.4955C437.845 73.4955 450.648 86.2981 450.648 102.091ZM400.015 102.091C400.015 114.262 409.882 124.128 422.053 124.128C434.223 124.128 444.09 114.262 444.09 102.091C444.09 89.9201 434.223 80.0537 422.053 80.0537C409.882 80.0537 400.015 89.9201 400.015 102.091Z" fill="white"/></mask><g mask="url(#mask2_2_190)"><path d="M450.648 102.091C450.648 117.884 437.845 130.686 422.053 130.686C406.26 130.686 393.457 117.884 393.457 102.091C393.457 86.2981 406.26 73.4955 422.053 73.4955C437.845 73.4955 450.648 86.2981 450.648 102.091ZM400.015 102.091C400.015 114.262 409.882 124.128 422.053 124.128C434.223 124.128 444.09 114.262 444.09 102.091C444.09 89.9201 434.223 80.0537 422.053 80.0537C409.882 80.0537 400.015 89.9201 400.015 102.091Z" fill="white" stroke="black" stroke-width="0.876179"/></g></g><defs><filter id="filter0_f_2_190" x="-381" y="-552" width="1275" height="840" filterUnits="userSpaceOnUse" color-interpolation-filters="sRGB"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend mode="normal" in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feGaussianBlur stdDeviation="128" result="effect1_foregroundBlur_2_190"/></filter><clipPath id="clip0_2_190"><rect width="512" height="468" rx="40" fill="white"/></clipPath></defs></svg>)) (contracts/vote-escrow/TokenUriHelper.sol#97-104)
	output = string(abi.encodePacked(data:application/json;base64,,json)) (contracts/vote-escrow/TokenUriHelper.sol#125)
Reference:  

	- Timelock.acceptAdmin() (contracts/governance/Timelock.sol#64-70)
	- Timelock.queueTransaction(address,uint256,string,bytes,uint256) (contracts/governance/Timelock.sol#79-97)
	- Timelock.cancelTransaction(address,uint256,string,bytes,uint256) (contracts/governance/Timelock.sol#99-112)
Reference:  

	- GolomTrader.setDistributor(address) (contracts/core/GolomTrader.sol#524-526)
	- RewardDistributor.addFee(address[2],uint256) (contracts/rewards/RewardDistributor.sol#101-137)
	- RewardDistributor.changeVoteEscrow(address) (contracts/rewards/RewardDistributor.sol#278-280)
Reference:  

Setter function ERC20Mock.constructor() (contracts/test/ERC20Mock.sol#18-21) does not emit an event
Setter function ERC20Mock.setMinter(address) (contracts/test/ERC20Mock.sol#24-27) does not emit an event
Setter function ERC20Mock.transfer(address,uint256) (contracts/test/ERC20Mock.sol#57-59) does not emit an event
Setter function ERC20Mock.transferFrom(address,address,uint256) (contracts/test/ERC20Mock.sol#61-71) does not emit an event
Setter function ERC20Mock.mint(address,uint256) (contracts/test/ERC20Mock.sol#73-77) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Setter function WETH.transfer(address,uint256) (contracts/test/WETH.sol#42-44) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Setter function GovernorAlpha.castVote(uint256,uint256,bool) (contracts/governance/GovernerBravo.sol#348-355) does not emit an event
Setter function GovernorAlpha.__acceptAdmin() (contracts/governance/GovernerBravo.sol#398-401) does not emit an event
Setter function GovernorAlpha.__abdicate() (contracts/governance/GovernerBravo.sol#403-406) does not emit an event
Setter function GovernorAlpha.__queueSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#408-411) does not emit an event
Setter function GovernorAlpha.__executeSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#413-416) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Setter function GolomTrader.fillAsk(GolomTrader.Order,uint256,address,GolomTrader.Payment,address) (contracts/core/GolomTrader.sol#229-311) does not emit an event
Setter function GolomTrader.fillBid(GolomTrader.Order,uint256,address,GolomTrader.Payment) (contracts/core/GolomTrader.sol#337-375) does not emit an event
Setter function GolomTrader.fillCriteriaBid(GolomTrader.Order,uint256,uint256,bytes32[],address,GolomTrader.Payment) (contracts/core/GolomTrader.sol#404-448) does not emit an event
Setter function GolomTrader._settleBalances(GolomTrader.Order,uint256,address,GolomTrader.Payment) (contracts/core/GolomTrader.sol#455-483) does not emit an event
Setter function GolomTrader.setDistributor(address) (contracts/core/GolomTrader.sol#524-526) does not emit an event
Setter function GolomToken.mint(address,uint256) (contracts/governance/GolomToken.sol#37-39) does not emit an event
Setter function GolomToken.mintAirdrop(address) (contracts/governance/GolomToken.sol#43-47) does not emit an event
Setter function GolomToken.mintGenesisReward(address) (contracts/governance/GolomToken.sol#51-55) does not emit an event
Setter function GolomToken.mintTreasury(address) (contracts/governance/GolomToken.sol#59-63) does not emit an event
Setter function GolomToken.setMinter(address) (contracts/governance/GolomToken.sol#67-69) does not emit an event
Setter function GolomToken.onlyMinter() (contracts/governance/GolomToken.sol#24-27) does not emit an event
Setter function DummyRewardDistributor.addFee(address[2],uint256) (contracts/rewards/DummyRewardDistributor.sol#37) does not emit an event
Setter function DummyRewardDistributor.withdrawTokens(address,address) (contracts/rewards/DummyRewardDistributor.sol#39-41) does not emit an event
Setter function DummyRewardDistributor.withdrawEth(address) (contracts/rewards/DummyRewardDistributor.sol#43-45) does not emit an event
Setter function DummyRewardDistributor.onlyTrader() (contracts/rewards/DummyRewardDistributor.sol#26-29) does not emit an event
Setter function GenesisClaim.validateOrder(IGolomTrader.Order) (contracts/rewards/GenesisClaim.sol#138-144) does not emit an event
Setter function GenesisClaim.fixBoolean() (contracts/rewards/GenesisClaim.sol#152-154) does not emit an event
Setter function GenesisClaim.pauseAirdrop() (contracts/rewards/GenesisClaim.sol#156-158) does not emit an event
Setter function GenesisClaim.unpauseAirdrop() (contracts/rewards/GenesisClaim.sol#160-162) does not emit an event
Setter function GenesisClaim.changeLockDuration(uint256,uint256) (contracts/rewards/GenesisClaim.sol#184-187) does not emit an event
Setter function GolomAirdrop.validateOrder(IGolomTrader.Order) (contracts/rewards/GolomAirdrop.sol#138-144) does not emit an event
Setter function GolomAirdrop.fixBoolean() (contracts/rewards/GolomAirdrop.sol#151-153) does not emit an event
Setter function GolomAirdrop.pauseAirdrop() (contracts/rewards/GolomAirdrop.sol#155-157) does not emit an event
Setter function GolomAirdrop.unpauseAirdrop() (contracts/rewards/GolomAirdrop.sol#159-161) does not emit an event
Setter function GolomAirdrop.changeLockDuration(uint256,uint256) (contracts/rewards/GolomAirdrop.sol#183-186) does not emit an event
Setter function RewardDistributor.traderClaim(address,uint256[]) (contracts/rewards/RewardDistributor.sol#140-151) does not emit an event
Setter function RewardDistributor.exchangeClaim(address,uint256[]) (contracts/rewards/RewardDistributor.sol#154-165) does not emit an event
Setter function RewardDistributor.multiStakerClaim(uint256[],uint256[]) (contracts/rewards/RewardDistributor.sol#171-207) does not emit an event
Setter function RewardDistributor.changeTrader(address) (contracts/rewards/RewardDistributor.sol#272-274) does not emit an event
Setter function RewardDistributor.changeVoteEscrow(address) (contracts/rewards/RewardDistributor.sol#278-280) does not emit an event
Setter function RewardDistributor.onlyTrader() (contracts/rewards/RewardDistributor.sol#90-93) does not emit an event
Setter function ERC1155Mock.constructor() (contracts/test/ERC1155Mock.sol#14-20) does not emit an event
Setter function ERC1155Mock.mint(address) (contracts/test/ERC1155Mock.sol#22-25) does not emit an event
Setter function VoteEscrow._writeCheckpoint(uint256,uint256,uint256[]) (contracts/vote-escrow/VoteEscrowDelegation.sol#106-121) does not emit an event
Setter function VoteEscrow.removeDelegation(uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#225-235) does not emit an event
Setter function VoteEscrow._removeTokenFrom(address,uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#265-278) does not emit an event
Setter function VoteEscrow.changeMinVotingPower(uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#282-284) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

value assignment lack of validation	RewardDistributor.multiStakerClaim(uint256[],uint256[]) (contracts/rewards/RewardDistributor.sol#171-207)tokenowner = ve.ownerOf(tokenids[0]) (contracts/rewards/RewardDistributor.sol#176)
Reference: input validation

variable lacks a zero-check on 		- ERC20Mock.setMinter(address) (contracts/test/ERC20Mock.sol#24-27)
variable lacks a zero-check on 		- ERC20Mock.approve(address,uint256) (contracts/test/ERC20Mock.sol#29-33)
variable lacks a zero-check on 		- ERC20Mock.transferFrom(address,address,uint256) (contracts/test/ERC20Mock.sol#61-71)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

variable lacks a zero-check on 		- WETH.approve(address,uint256) (contracts/test/WETH.sol#36-40)
variable lacks a zero-check on 		- WETH.transferFrom(address,address,uint256) (contracts/test/WETH.sol#46-64)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

variable lacks a zero-check on 		- Timelock.constructor(address,uint256) (contracts/governance/Timelock.sol#47-53)
variable lacks a zero-check on 		- Timelock.setPendingAdmin(address) (contracts/governance/Timelock.sol#72-77)
variable lacks a zero-check on 		- Timelock.executeTransaction(address,uint256,string,bytes,uint256) (contracts/governance/Timelock.sol#114-145)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

variable lacks a zero-check on 		- GovernorAlpha.constructor(address,address,address,address) (contracts/governance/GovernerBravo.sol#149-160)
variable lacks a zero-check on 		- GovernorAlpha.__queueSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#408-411)
variable lacks a zero-check on 		- GovernorAlpha.__executeSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#413-416)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

variable lacks a zero-check on 		- Ownable.changeOwner(address) (contracts/utils/Ownable.sol#27-30)
variable lacks a zero-check on 		- GolomToken.setMinter(address) (contracts/governance/GolomToken.sol#67-69)
variable lacks a zero-check on 		- GolomTreasury.withdraw(address,uint256) (contracts/governance/GolomTreasury.sol#19-22)
variable lacks a zero-check on 		- DummyRewardDistributor.constructor(address,address) (contracts/rewards/DummyRewardDistributor.sol#21-24)
variable lacks a zero-check on 		- DummyRewardDistributor.withdrawTokens(address,address) (contracts/rewards/DummyRewardDistributor.sol#39-41)
variable lacks a zero-check on 		- DummyRewardDistributor.withdrawEth(address) (contracts/rewards/DummyRewardDistributor.sol#43-45)
variable lacks a zero-check on 		- RewardDistributor.constructor(address,address,address,address,address,uint256) (contracts/rewards/RewardDistributor.sol#74-88)
variable lacks a zero-check on 		- RewardDistributor.traderClaim(address,uint256[]) (contracts/rewards/RewardDistributor.sol#140-151)
variable lacks a zero-check on 		- RewardDistributor.exchangeClaim(address,uint256[]) (contracts/rewards/RewardDistributor.sol#154-165)
variable lacks a zero-check on 		- RewardDistributor.changeTrader(address) (contracts/rewards/RewardDistributor.sol#272-274)
variable lacks a zero-check on 		- VoteEscrowCore.safeTransferFrom(address,address,uint256,bytes) (contracts/vote-escrow/VoteEscrowCore.sol#593-615)
variable lacks a zero-check on 		- VoteEscrow.constructor(address) (contracts/vote-escrow/VoteEscrowDelegation.sol#52-65)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

Variable 'ERC20Votes._moveVotingPower(address,address,uint256).oldWeight (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#215)' in ERC20Votes._moveVotingPower(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#208-224) potentially used before declaration: (oldWeight,newWeight) = _writeCheckpoint(_checkpoints[dst],_add,amount) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#220)
Variable 'ECDSA.tryRecover(bytes32,bytes).r (node_modules/@openzeppelin/contracts/utils/cryptography/ECDSA.sol#62)' in ECDSA.tryRecover(bytes32,bytes) (node_modules/@openzeppelin/contracts/utils/cryptography/ECDSA.sol#57-86) potentially used before declaration: r = mload(uint256)(signature + 0x20) (node_modules/@openzeppelin/contracts/utils/cryptography/ECDSA.sol#79)
Reference:  

GovernorAlpha._castVote(uint256,uint256,bool) (contracts/governance/GovernerBravo.sol#374-396) compares to a boolean constant:
	-require(bool,string)(receipt.hasVoted == false,GovernorAlpha::_castVote: voter already voted) (contracts/governance/GovernerBravo.sol#382)
Reference:  

	- ERC20Mock.setMinter(address) (contracts/test/ERC20Mock.sol#24-27)
Reference:  

	- GovernorAlpha.propose(uint256,address[],uint256[],string[],bytes[],string) (contracts/governance/GovernerBravo.sol#162-232)
	- GovernorAlpha.cancel(uint256) (contracts/governance/GovernerBravo.sol#283-306)
	- GovernorAlpha.__abdicate() (contracts/governance/GovernerBravo.sol#403-406)
Reference:  

	- Ownable.changeOwner(address) (contracts/utils/Ownable.sol#27-30)
	- Ownable.acceptOwnership() (contracts/utils/Ownable.sol#33-38)
	- GolomToken.mintAirdrop(address) (contracts/governance/GolomToken.sol#43-47)
	- GolomToken.mintGenesisReward(address) (contracts/governance/GolomToken.sol#51-55)
	- GolomToken.mintTreasury(address) (contracts/governance/GolomToken.sol#59-63)
	- GolomToken.setMinter(address) (contracts/governance/GolomToken.sol#67-69)
	- GenesisClaim.setMerkleRoot(bytes32) (contracts/rewards/GenesisClaim.sol#146-150)
	- GenesisClaim.fixBoolean() (contracts/rewards/GenesisClaim.sol#152-154)
	- GenesisClaim.updateEndTimestamp(uint256) (contracts/rewards/GenesisClaim.sol#164-168)
	- GenesisClaim.changeLockDuration(uint256,uint256) (contracts/rewards/GenesisClaim.sol#184-187)
	- GenesisClaim.claim(uint256,bytes32[],bool) (contracts/rewards/GenesisClaim.sol#189-211)
	- GolomAirdrop.setMerkleRoot(bytes32) (contracts/rewards/GolomAirdrop.sol#146-150)
	- GolomAirdrop.fixBoolean() (contracts/rewards/GolomAirdrop.sol#151-153)
	- GolomAirdrop.updateEndTimestamp(uint256) (contracts/rewards/GolomAirdrop.sol#163-167)
	- GolomAirdrop.changeLockDuration(uint256,uint256) (contracts/rewards/GolomAirdrop.sol#183-186)
	- GolomAirdrop.claim(uint256,bytes32[],IGolomTrader.Order,bool) (contracts/rewards/GolomAirdrop.sol#188-213)
	- RewardDistributor.changeTrader(address) (contracts/rewards/RewardDistributor.sol#272-274)
	- VoteEscrow.changeMinVotingPower(uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#282-284)
Reference:  

	- ERC20Mock.mint(address,uint256) (contracts/test/ERC20Mock.sol#73-77)
Reference:  

	- GovernorAlpha.castVote(uint256,uint256,bool) (contracts/governance/GovernerBravo.sol#348-355)
	- GovernorAlpha.__acceptAdmin() (contracts/governance/GovernerBravo.sol#398-401)
	- GovernorAlpha.__queueSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#408-411)
	- GovernorAlpha.__executeSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#413-416)
Reference:  

	- GolomToken.mint(address,uint256) (contracts/governance/GolomToken.sol#37-39)
	- DummyRewardDistributor.addFee(address[2],uint256) (contracts/rewards/DummyRewardDistributor.sol#37)
	- GenesisClaim.pauseAirdrop() (contracts/rewards/GenesisClaim.sol#156-158)
	- GenesisClaim.unpauseAirdrop() (contracts/rewards/GenesisClaim.sol#160-162)
	- GolomAirdrop.pauseAirdrop() (contracts/rewards/GolomAirdrop.sol#155-157)
	- GolomAirdrop.unpauseAirdrop() (contracts/rewards/GolomAirdrop.sol#159-161)
Reference:  

GovernorAlpha.getChainId() (contracts/governance/GovernerBravo.sol#429-435) is never used and should be removed
Reference:  

GenesisClaim._hashOrder(IGolomTrader.Order,uint256[2]) (contracts/rewards/GenesisClaim.sol#114-136) is never used and should be removed
GenesisClaim.hashPayment(IGolomTrader.Payment) (contracts/rewards/GenesisClaim.sol#103-112) is never used and should be removed
GenesisClaim.validateOrder(IGolomTrader.Order) (contracts/rewards/GenesisClaim.sol#138-144) is never used and should be removed
Reference:  

require() missing error messages
	 - require(bool)(msg.sender == minter) (contracts/test/ERC20Mock.sol#25)
require() missing error messages
	 - require(bool)(msg.sender == minter) (contracts/test/ERC20Mock.sol#74)
Reference: https://dev.to/tawseef/require-vs-assert-in-solidity-5e9d

require() missing error messages
	 - require(bool)(msg.sender == o.reservedAddress) (contracts/core/GolomTrader.sol#247)
require() missing error messages
	 - require(bool)(o.totalAmt * amount >= (o.exchange.paymentAmt + o.prePayment.paymentAmt) * amount + p.paymentAmt) (contracts/core/GolomTrader.sol#343)
require() missing error messages
	 - require(bool)(msg.sender == o.reservedAddress) (contracts/core/GolomTrader.sol#347)
require() missing error messages
	 - require(bool)(o.orderType == 1) (contracts/core/GolomTrader.sol#349)
require() missing error messages
	 - require(bool)(status == 3) (contracts/core/GolomTrader.sol#351)
require() missing error messages
	 - require(bool)(amountRemaining >= amount) (contracts/core/GolomTrader.sol#352)
require() missing error messages
	 - require(bool)(o.signer == msg.sender) (contracts/core/GolomTrader.sol#380)
require() missing error messages
	 - require(bool)(o.totalAmt * amount >= (o.exchange.paymentAmt + o.prePayment.paymentAmt) * amount + p.paymentAmt) (contracts/core/GolomTrader.sol#412)
require() missing error messages
	 - require(bool)(msg.sender == o.reservedAddress) (contracts/core/GolomTrader.sol#415)
require() missing error messages
	 - require(bool)(o.orderType == 2) (contracts/core/GolomTrader.sol#417)
require() missing error messages
	 - require(bool)(status == 3) (contracts/core/GolomTrader.sol#419)
require() missing error messages
	 - require(bool)(amountRemaining >= amount) (contracts/core/GolomTrader.sol#420)
require() missing error messages
	 - require(bool)(msg.sender == trader) (contracts/rewards/DummyRewardDistributor.sol#27)
require() missing error messages
	 - require(bool)(epochs[index] < epoch) (contracts/rewards/RewardDistributor.sol#143)
require() missing error messages
	 - require(bool)(epochs[index] < epoch) (contracts/rewards/RewardDistributor.sol#157)
require() missing error messages
	 - require(bool)(msg.sender == trader) (contracts/rewards/RewardDistributor.sol#91)
require() missing error messages
	 - require(bool)(_isApprovedOrOwner(_sender,_tokenId)) (contracts/vote-escrow/VoteEscrowCore.sol#539)
require() missing error messages
	 - require(bool)(owner != address(0)) (contracts/vote-escrow/VoteEscrowCore.sol#645)
require() missing error messages
	 - require(bool)(_approved != owner) (contracts/vote-escrow/VoteEscrowCore.sol#647)
require() missing error messages
	 - require(bool)(senderIsOwner || senderIsApprovedForAll) (contracts/vote-escrow/VoteEscrowCore.sol#651)
require() missing error messages
	 - require(bool)(_value > 0) (contracts/vote-escrow/VoteEscrowCore.sol#925)
require() missing error messages
	 - require(bool)(_value > 0) (contracts/vote-escrow/VoteEscrowCore.sol#942)
require() missing error messages
	 - require(bool)(_isApprovedOrOwner(_sender,_tokenId)) (contracts/vote-escrow/VoteEscrowDelegation.sol#250)
require() missing error messages
	 - require(bool)(_entered_state == _not_entered) (contracts/vote-escrow/VoteEscrowCore.sol#360)
Reference: https://dev.to/tawseef/require-vs-assert-in-solidity-5e9d

keyERC721Mock (contracts/test/ERC721Mock.sol#10-28) does not specify SPDX license identifierkeyBase64 (contracts/vote-escrow/TokenUriHelper.sol#8-63) does not specify SPDX license identifierReference: https://certik-public-assets.s3.amazonaws.com/CertiK-Audit-for-Polylastic---Airdrop-and-Token-Swap.pdf

Potential price manipulation risk:
	- In function _delegate
		-- delegatorBalance = balanceOf(delegator) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#200) have potential price manipulated risk from delegatorBalance and call None which could influence variable:delegatorBalance
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

Potential DoS Gas Limit Attack occur inGovernorAlpha.queue(uint256) (contracts/governance/GovernerBravo.sol#234-248)BEGIN_LOOP (contracts/governance/GovernerBravo.sol#242-244)
Potential DoS Gas Limit Attack occur inGovernorAlpha.execute(uint256) (contracts/governance/GovernerBravo.sol#264-281)BEGIN_LOOP (contracts/governance/GovernerBravo.sol#271-279)
Reference: https://swcregistry.io/docs/SWC-128

Potential DoS Gas Limit Attack occur inRewardDistributor.traderClaim(address,uint256[]) (contracts/rewards/RewardDistributor.sol#140-151)BEGIN_LOOP (contracts/rewards/RewardDistributor.sol#142-149)
Potential DoS Gas Limit Attack occur inRewardDistributor.exchangeClaim(address,uint256[]) (contracts/rewards/RewardDistributor.sol#154-165)BEGIN_LOOP (contracts/rewards/RewardDistributor.sol#156-163)
Potential DoS Gas Limit Attack occur inRewardDistributor.multiStakerClaim(uint256[],uint256[]) (contracts/rewards/RewardDistributor.sol#171-207)BEGIN_LOOP (contracts/rewards/RewardDistributor.sol#179-204)
Potential DoS Gas Limit Attack occur inVoteEscrowCore._checkpoint(uint256,LockedBalance,LockedBalance) (contracts/vote-escrow/VoteEscrowCore.sol#689-824)BEGIN_LOOP (contracts/vote-escrow/VoteEscrowCore.sol#744-774)
Potential DoS Gas Limit Attack occur inVoteEscrow.removeElement(uint256[],uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#213-221)BEGIN_LOOP (contracts/vote-escrow/VoteEscrowDelegation.sol#214-220)
Reference: https://swcregistry.io/docs/SWC-128

function:Emitter.pushOrder(Emitter.Order) (contracts/core/Emitter.sol#117-129)is public and can be replaced with external 
Reference:  

function:WETH.deposit() (contracts/test/WETH.sol#20-23)is public and can be replaced with external 
function:WETH.withdraw(uint256) (contracts/test/WETH.sol#25-30)is public and can be replaced with external 
function:WETH.totalSupply() (contracts/test/WETH.sol#32-34)is public and can be replaced with external 
function:WETH.approve(address,uint256) (contracts/test/WETH.sol#36-40)is public and can be replaced with external 
function:WETH.transfer(address,uint256) (contracts/test/WETH.sol#42-44)is public and can be replaced with external 
Reference:  

function:Timelock.setDelay(uint256) (contracts/governance/Timelock.sol#55-62)is public and can be replaced with external 
function:Timelock.acceptAdmin() (contracts/governance/Timelock.sol#64-70)is public and can be replaced with external 
function:Timelock.setPendingAdmin(address) (contracts/governance/Timelock.sol#72-77)is public and can be replaced with external 
function:Timelock.queueTransaction(address,uint256,string,bytes,uint256) (contracts/governance/Timelock.sol#79-97)is public and can be replaced with external 
function:Timelock.cancelTransaction(address,uint256,string,bytes,uint256) (contracts/governance/Timelock.sol#99-112)is public and can be replaced with external 
function:Timelock.executeTransaction(address,uint256,string,bytes,uint256) (contracts/governance/Timelock.sol#114-145)is public and can be replaced with external 
Reference:  

function:GovernorAlpha.propose(uint256,address[],uint256[],string[],bytes[],string) (contracts/governance/GovernerBravo.sol#162-232)is public and can be replaced with external 
function:GovernorAlpha.queue(uint256) (contracts/governance/GovernerBravo.sol#234-248)is public and can be replaced with external 
function:GovernorAlpha.execute(uint256) (contracts/governance/GovernerBravo.sol#264-281)is public and can be replaced with external 
function:GovernorAlpha.cancel(uint256) (contracts/governance/GovernerBravo.sol#283-306)is public and can be replaced with external 
function:GovernorAlpha.getActions(uint256) (contracts/governance/GovernerBravo.sol#308-320)is public and can be replaced with external 
function:GovernorAlpha.getReceipt(uint256,uint256) (contracts/governance/GovernerBravo.sol#322-324)is public and can be replaced with external 
function:GovernorAlpha.castVote(uint256,uint256,bool) (contracts/governance/GovernerBravo.sol#348-355)is public and can be replaced with external 
function:GovernorAlpha.__acceptAdmin() (contracts/governance/GovernerBravo.sol#398-401)is public and can be replaced with external 
function:GovernorAlpha.__abdicate() (contracts/governance/GovernerBravo.sol#403-406)is public and can be replaced with external 
function:GovernorAlpha.__queueSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#408-411)is public and can be replaced with external 
function:GovernorAlpha.__executeSetTimelockPendingAdmin(address,uint256) (contracts/governance/GovernerBravo.sol#413-416)is public and can be replaced with external 
Reference:  

function:Pausable.paused() (node_modules/@openzeppelin/contracts/security/Pausable.sol#40-42)is public and can be replaced with external 
function:ERC1155.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#42-47)is public and can be replaced with external 
function:ERC1155.uri(uint256) (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#59-61)is public and can be replaced with external 
function:ERC1155.balanceOfBatch(address[],uint256[]) (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#82-98)is public and can be replaced with external 
function:ERC1155.setApprovalForAll(address,bool) (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#103-105)is public and can be replaced with external 
function:ERC1155.safeTransferFrom(address,address,uint256,uint256,bytes) (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#117-129)is public and can be replaced with external 
function:ERC1155.safeBatchTransferFrom(address,address,uint256[],uint256[],bytes) (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#134-146)is public and can be replaced with external 
function:ERC20.name() (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#62-64)is public and can be replaced with external 
function:ERC20.symbol() (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#70-72)is public and can be replaced with external 
function:ERC20.decimals() (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#87-89)is public and can be replaced with external 
function:ERC20.totalSupply() (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#94-96)is public and can be replaced with external 
function:ERC20.balanceOf(address) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#101-103)is public and can be replaced with external 
function:ERC20.transfer(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#113-117)is public and can be replaced with external 
function:ERC20.approve(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#136-140)is public and can be replaced with external 
function:ERC20.transferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#158-167)is public and can be replaced with external 
function:ERC20.increaseAllowance(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#181-185)is public and can be replaced with external 
function:ERC20.decreaseAllowance(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#201-210)is public and can be replaced with external 
function:ERC20Permit.permit(address,address,uint256,uint256,uint8,bytes32,bytes32) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#49-68)is public and can be replaced with external 
function:ERC20Permit.nonces(address) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#73-75)is public and can be replaced with external 
function:ERC20Votes.checkpoints(address,uint32) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#43-45)is public and can be replaced with external 
function:ERC20Votes.numCheckpoints(address) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#50-52)is public and can be replaced with external 
function:ERC20Votes.getVotes(address) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#64-67)is public and can be replaced with external 
function:ERC20Votes.getPastVotes(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#76-79)is public and can be replaced with external 
function:ERC20Votes.getPastTotalSupply(uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#89-92)is public and can be replaced with external 
function:ERC20Votes.delegate(address) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#126-128)is public and can be replaced with external 
function:ERC20Votes.delegateBySig(address,uint256,uint256,uint8,bytes32,bytes32) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#133-150)is public and can be replaced with external 
function:ERC721.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#52-57)is public and can be replaced with external 
function:ERC721.balanceOf(address) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#62-65)is public and can be replaced with external 
function:ERC721.ownerOf(uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#70-74)is public and can be replaced with external 
function:ERC721.name() (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#79-81)is public and can be replaced with external 
function:ERC721.symbol() (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#86-88)is public and can be replaced with external 
function:ERC721.tokenURI(uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#93-98)is public and can be replaced with external 
function:ERC721.approve(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#112-122)is public and can be replaced with external 
function:ERC721.setApprovalForAll(address,bool) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#136-138)is public and can be replaced with external 
function:ERC721.transferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#150-159)is public and can be replaced with external 
function:ERC721.safeTransferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#164-170)is public and can be replaced with external 
function:ERC165.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#26-28)is public and can be replaced with external 
function:Ownable.changeOwner(address) (contracts/utils/Ownable.sol#27-30)is public and can be replaced with external 
function:Ownable.acceptOwnership() (contracts/utils/Ownable.sol#33-38)is public and can be replaced with external 
function:GolomTrader.fillAsk(GolomTrader.Order,uint256,address,GolomTrader.Payment,address) (contracts/core/GolomTrader.sol#229-311)is public and can be replaced with external 
function:GolomTrader.fillBid(GolomTrader.Order,uint256,address,GolomTrader.Payment) (contracts/core/GolomTrader.sol#337-375)is public and can be replaced with external 
function:GolomTrader.cancelOrder(GolomTrader.Order) (contracts/core/GolomTrader.sol#379-387)is public and can be replaced with external 
function:GolomTrader.fillCriteriaBid(GolomTrader.Order,uint256,uint256,bytes32[],address,GolomTrader.Payment) (contracts/core/GolomTrader.sol#404-448)is public and can be replaced with external 
function:DummyRewardDistributor.addFee(address[2],uint256) (contracts/rewards/DummyRewardDistributor.sol#37)is public and can be replaced with external 
function:RewardDistributor.addFee(address[2],uint256) (contracts/rewards/RewardDistributor.sol#101-137)is public and can be replaced with external 
function:RewardDistributor.traderClaim(address,uint256[]) (contracts/rewards/RewardDistributor.sol#140-151)is public and can be replaced with external 
function:RewardDistributor.exchangeClaim(address,uint256[]) (contracts/rewards/RewardDistributor.sol#154-165)is public and can be replaced with external 
function:RewardDistributor.multiStakerClaim(uint256[],uint256[]) (contracts/rewards/RewardDistributor.sol#171-207)is public and can be replaced with external 
function:RewardDistributor.stakerRewards(uint256) (contracts/rewards/RewardDistributor.sol#211-247)is public and can be replaced with external 
function:RewardDistributor.traderRewards(address) (contracts/rewards/RewardDistributor.sol#251-257)is public and can be replaced with external 
function:RewardDistributor.exchangeRewards(address) (contracts/rewards/RewardDistributor.sol#261-267)is public and can be replaced with external 
function:ERC1155Mock.mint(address) (contracts/test/ERC1155Mock.sol#22-25)is public and can be replaced with external 
function:ERC721Mock.mint(address) (contracts/test/ERC721Mock.sol#16-23)is public and can be replaced with external 
function:ERC721Mock.current() (contracts/test/ERC721Mock.sol#25-27)is public and can be replaced with external 
function:TokenUriHelper._tokenURI(uint256,uint256,uint256,uint256) (contracts/vote-escrow/TokenUriHelper.sol#66-126)is public and can be replaced with external 
function:VoteEscrow.getPriorVotes(uint256,uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#198-208)is public and can be replaced with external 
Reference:  

VoteEscrowDelegateVotesChanged(address,uint256,uint256) (contracts/vote-escrow/VoteEscrowDelegation.sol#29) is never emitted in VoteEscrow (contracts/vote-escrow/VoteEscrowDelegation.sol#20-285)
Reference: https://certik-public-assets.s3.amazonaws.com/CertiK-Audit-for-Kitty-inu.pdf

	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.11. ^0.8.11 (contracts/governance/Timelock.sol#2)
Reference:  

	Pragma confirmed better, here is pragma solidity^0.8.11. ^0.8.11 (contracts/governance/GovernerBravo.sol#2)
Reference:  

	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/governance/utils/IVotes.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/security/Pausable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC1155/ERC1155.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC1155/IERC1155.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC1155/extensions/IERC1155MetadataURI.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (node_modules/@openzeppelin/contracts/utils/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Context.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Counters.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Strings.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/cryptography/ECDSA.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/cryptography/MerkleProof.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/cryptography/draft-EIP712.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/math/Math.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.11. ^0.8.11 (contracts/governance/GolomToken.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.11. ^0.8.11 (contracts/governance/GolomTreasury.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/test/ERC1155Mock.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/test/ERC721Mock.sol#5)
	Pragma confirmed better, here is pragma solidity>=0.4.22<0.9.0. >=0.4.22<0.9.0 (node_modules/hardhat/console.sol#2)
Reference:  

WETH.decimals (contracts/test/WETH.sol#7) should be constant
WETH.name (contracts/test/WETH.sol#5) should be constant
WETH.symbol (contracts/test/WETH.sol#6) should be constant
Reference:  

ERC20Permit._PERMIT_TYPEHASH_DEPRECATED_SLOT (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#37) should be constant
GolomToken.minterEnableDate (contracts/governance/GolomToken.sol#17) should be constant
GolomToken.pendingMinter (contracts/governance/GolomToken.sol#18) should be constant
GolomTrader.governance (contracts/core/GolomTrader.sol#72) should be constant
RewardDistributor.pendingTrader (contracts/rewards/RewardDistributor.sol#50) should be constant
RewardDistributor.pendingVoteEscrow (contracts/rewards/RewardDistributor.sol#53) should be constant
RewardDistributor.traderEnableDate (contracts/rewards/RewardDistributor.sol#51) should be constant
RewardDistributor.voteEscrowEnableDate (contracts/rewards/RewardDistributor.sol#54) should be constant
Reference:  

unnecessary reentrancy lock found inGolomTrader
	- GolomTrader.incrementNonce() (contracts/core/GolomTrader.sol#393-396)
unnecessary reentrancy lock found inVoteEscrow
	- VoteEscrowCore.create_lock_for(uint256,uint256,address) (contracts/vote-escrow/VoteEscrowCore.sol#957-963)
Reference: http://
. analyzed (70 contracts with 86 detectors), 326 result(s) found
INFO:Falcon:metatrust result: ../../scanned_output/golom-protocol-contracts_autogen_output/mwe-output.json generate success.
before get_detectors_and_printers
before main_impl
parse_args
set_colorization_enabled
before outputting_json
before choose_printers
before getLevelName
before for
before StreamHandler
before FormatterCryticCompile
before getLogger
try1
before compile_all
after compile all
before: for compilation in compilations:
begin to process ast
run process
check register
check printer_cls
begin run detectors
An error occurred: unhashable type: 'list'
An error occurred: unhashable type: 'list'
An error occurred: 
An error occurred: 
An error occurred: 
An error occurred: 
An error occurred: 
An error occurred: 
after for recrusion
Execution time for Falcon: 11m26.605017605s
../input_repos/golom-protocol-contracts/contracts/core/Emitter.sol: Smart Contract (1): 150 lines
../input_repos/golom-protocol-contracts/contracts/core/GolomTrader.sol: Smart Contract (1), Interface (4): 531 lines
../input_repos/golom-protocol-contracts/contracts/governance/GolomToken.sol: Smart Contract (1): 70 lines
../input_repos/golom-protocol-contracts/contracts/governance/GolomTreasury.sol: Smart Contract (1): 23 lines
../input_repos/golom-protocol-contracts/contracts/governance/GovernerBravo.sol: Smart Contract (1), Interface (3): 474 lines
../input_repos/golom-protocol-contracts/contracts/governance/Timelock.sol: Smart Contract (1): 155 lines
../input_repos/golom-protocol-contracts/contracts/rewards/DummyRewardDistributor.sol: Smart Contract (1), Interface (1): 50 lines
../input_repos/golom-protocol-contracts/contracts/rewards/GenesisClaim.sol: Smart Contract (1), Interface (2): 217 lines
../input_repos/golom-protocol-contracts/contracts/rewards/GolomAirdrop.sol: Smart Contract (1), Interface (2): 219 lines
../input_repos/golom-protocol-contracts/contracts/rewards/RewardDistributor.sol: Smart Contract (1), Interface (2): 285 lines
../input_repos/golom-protocol-contracts/contracts/utils/Ownable.sol: Abstract Contract (1): 46 lines
../input_repos/golom-protocol-contracts/contracts/vote-escrow/TokenUriHelper.sol: Library (2): 149 lines
../input_repos/golom-protocol-contracts/contracts/vote-escrow/VoteEscrowCore.sol: Smart Contract (1), Interface (5): 1234 lines
../input_repos/golom-protocol-contracts/contracts/vote-escrow/VoteEscrowDelegation.sol: Smart Contract (1), Interface (1): 285 lines

Summary:
Total number of Solidity files: 14
Abstract Contracts Files: 1
Smart Contracts Files: 4
Libraries Files: 1
Interfaces Files: 0
Multiple Types Files: 8
******************************************
Abstract Contracts: 1
Smart Contracts: 12
Libraries: 2
Interfaces: 20
LOC: 3888
