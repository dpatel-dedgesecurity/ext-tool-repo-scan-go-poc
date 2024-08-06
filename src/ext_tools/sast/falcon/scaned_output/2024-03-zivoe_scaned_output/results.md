'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-03-zivoe/zivoe-core-foundry)
'forge config --json' running
'forge build --build-info --skip */test/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-03-zivoe/zivoe-core-foundry)

Reentrancy in ZivoeTranches.depositBoth(uint256,address,uint256,address) (src/ZivoeTranches.sol#322-325):
	External calls:
	- depositSenior(amountSenior,assetSenior) (src/ZivoeTranches.sol#323)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- IERC20(asset).safeTransferFrom(depositor,IZivoeGlobals_ZivoeTranches(GBL).DAO(),amount) (src/ZivoeTranches.sol#304)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	- depositJunior(amountJunior,assetJunior) (src/ZivoeTranches.sol#324)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- IERC20(asset).safeTransferFrom(depositor,IZivoeGlobals_ZivoeTranches(GBL).DAO(),amount) (src/ZivoeTranches.sol#277)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	External calls sending eth:
	- depositSenior(amountSenior,assetSenior) (src/ZivoeTranches.sol#323)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	- depositJunior(amountJunior,assetJunior) (src/ZivoeTranches.sol#324)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
Reentrancy in OCL_ZVE.forwardYield() (src/lockers/OCL/OCL_ZVE.sol#287-305):
	External calls:
	- _forwardYield(amount,lp) (src/lockers/OCL/OCL_ZVE.sol#302)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	External calls sending eth:
	- _forwardYield(amount,lp) (src/lockers/OCL/OCL_ZVE.sol#302)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
Reentrancy in ZivoeRewards.fullWithdraw() (src/ZivoeRewards.sol#246-249):
	External calls:
	- withdraw(_balances[_msgSender()]) (src/ZivoeRewards.sol#247)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	- getRewards() (src/ZivoeRewards.sol#248)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
		- IERC20(_rewardsToken).safeTransfer(_msgSender(),reward) (src/ZivoeRewards.sol#292)
	External calls sending eth:
	- withdraw(_balances[_msgSender()]) (src/ZivoeRewards.sol#247)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	- getRewards() (src/ZivoeRewards.sol#248)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
Reentrancy in ZivoeRewardsVesting.fullWithdraw() (src/ZivoeRewardsVesting.sol#370-373):
	External calls:
	- withdraw() (src/ZivoeRewardsVesting.sol#371)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	- getRewards() (src/ZivoeRewardsVesting.sol#372)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
		- IERC20(_rewardsToken).safeTransfer(_msgSender(),reward) (src/ZivoeRewardsVesting.sol#495)
	External calls sending eth:
	- withdraw() (src/ZivoeRewardsVesting.sol#371)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	- getRewards() (src/ZivoeRewardsVesting.sol#372)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
Reference:  

state variable: Governor._governanceCall (lib/openzeppelin-contracts/contracts/governance/Governor.sol#53) not initialized and not written in contract but be used in contract
state variable: GovernorVotesQuorumFraction._quorumNumerator (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#19) not initialized and not written in contract but be used in contract
state variable: ERC20Permit._nonces (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#25) not initialized and not written in contract but be used in contract
state variable: ERC20Votes._checkpoints (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#37) not initialized and not written in contract but be used in contract
state variable: ERC20Votes._totalSupplyCheckpoints (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#38) not initialized and not written in contract but be used in contract
state variable: ZivoeVotes._checkpoints (src/libraries/ZivoeVotes.sol#17) not initialized and not written in contract but be used in contract
state variable: ZivoeVotes._totalSupplyCheckpoints (src/libraries/ZivoeVotes.sol#19) not initialized and not written in contract but be used in contract
Reference:  

	- Governor.relay(address,uint256,bytes) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#543-550)
Reference:  

public mint or burn found in MockStablecoin.mint(address,uint256) (src/misc/MockStablecoin.sol#42-44)Reference: check public mint method

OCL_ZVE.pushToLockerMulti(address[],uint256[],bytes[]).preBasis (src/lockers/OCL/OCL_ZVE.sol#188) is a local variable never initialized
ERC20Votes._moveVotingPower(address,address,uint256).newWeight_scope_1 (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#233) is a local variable never initialized
Reference:  

Governor._execute(uint256,address[],uint256[],bytes[],bytes32) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#316-328)have ignores return value in Address.verifyCallResult(success,returndata,errorMessage) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#326)
Governor.relay(address,uint256,bytes) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#543-550)have ignores return value in Address.verifyCallResult(success,returndata,Governor: relay reverted without message) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#549)
GovernorVotesQuorumFraction._updateQuorumNumerator(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#99-118)have ignores return value in _quorumNumeratorHistory.push(newQuorumNumerator) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#115)
OCY_Convex_A.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_A.sol#117-169)have ignores return value in IBasePool_OCY_Convex_A(curveBasePool).add_liquidity(_amounts,_min_mint_amountBP) (src/lockers/OCY/OCY_Convex_A.sol#131)
OCY_Convex_A.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_A.sol#174-207)have ignores return value in IBaseRewardPool_OCY_Convex_A(convexRewards).withdrawAndUnwrap(IERC20(convexRewards).balanceOf(address(this)),false) (src/lockers/OCY/OCY_Convex_A.sol#180-182)
OCY_Convex_A.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_A.sol#213-242)have ignores return value in IBaseRewardPool_OCY_Convex_A(convexRewards).withdrawAndUnwrap(amount,false) (src/lockers/OCY/OCY_Convex_A.sol#218)
OCY_Convex_A.claimRewards(bool) (src/lockers/OCY/OCY_Convex_A.sol#246-266)have ignores return value in IBaseRewardPool_OCY_Convex_A(convexRewards).getReward() (src/lockers/OCY/OCY_Convex_A.sol#247)
OCY_Convex_B.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_B.sol#112-159)have ignores return value in IBooster_OCY_Convex_B(convexDeposit).deposit(convexPoolID,balCurveBasePoolToken,true) (src/lockers/OCY/OCY_Convex_B.sol#157)
OCY_Convex_B.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_B.sol#164-191)have ignores return value in IBaseRewardPool_OCY_Convex_B(convexRewards).withdrawAndUnwrap(IERC20(convexRewards).balanceOf(address(this)),false) (src/lockers/OCY/OCY_Convex_B.sol#170-172)
OCY_Convex_B.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_B.sol#197-223)have ignores return value in IBaseRewardPool_OCY_Convex_B(convexRewards).withdrawAndUnwrap(amount,false) (src/lockers/OCY/OCY_Convex_B.sol#203)
OCY_Convex_B.claimRewards(bool) (src/lockers/OCY/OCY_Convex_B.sol#227-247)have ignores return value in IBaseRewardPool_OCY_Convex_B(convexRewards).getReward() (src/lockers/OCY/OCY_Convex_B.sol#228)
OCY_Convex_C.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_C.sol#110-141)have ignores return value in IBooster_OCY_Convex_C(convexDeposit).deposit(convexPoolID,balCurveBasePoolToken,true) (src/lockers/OCY/OCY_Convex_C.sol#139)
OCY_Convex_C.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_C.sol#146-169)have ignores return value in IBaseRewardPool_OCY_Convex_C(convexRewards).withdrawAndUnwrap(IERC20(convexRewards).balanceOf(address(this)),false) (src/lockers/OCY/OCY_Convex_C.sol#152-154)
OCY_Convex_C.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_C.sol#175-196)have ignores return value in IBaseRewardPool_OCY_Convex_C(convexRewards).withdrawAndUnwrap(amount,false) (src/lockers/OCY/OCY_Convex_C.sol#181)
OCY_Convex_C.claimRewards(bool) (src/lockers/OCY/OCY_Convex_C.sol#200-220)have ignores return value in IBaseRewardPool_OCY_Convex_C(convexRewards).getReward() (src/lockers/OCY/OCY_Convex_C.sol#201)
Reference:  

function:ERC20._beforeTokenTransfer(address,address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#364-368)is empty 
function:ERC20._afterTokenTransfer(address,address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#384-388)is empty 
function:Presale.fallback() (src/misc/Presale.sol#190)is empty 
Reference:  

Setter function Timers.setDeadline(Timers.Timestamp,uint64) (lib/openzeppelin-contracts/contracts/utils/Timers.sol#18-20) does not emit an event
Setter function Timers.setDeadline(Timers.BlockNumber,uint64) (lib/openzeppelin-contracts/contracts/utils/Timers.sol#50-52) does not emit an event
Setter function ZivoeGlobals.proposeZVL(address) (src/ZivoeGlobals.sol#210-212) does not emit an event
Setter function ZivoeGlobals.onlyZVL() (src/ZivoeGlobals.sol#112-115) does not emit an event
Setter function ZivoeGovernorV2.setVotingDelay(uint256) (src/ZivoeGovernorV2.sol#84-87) does not emit an event
Setter function ZivoeGovernorV2.setVotingPeriod(uint256) (src/ZivoeGovernorV2.sol#91-95) does not emit an event
Setter function ZivoeGovernorV2.setProposalThreshold(uint256) (src/ZivoeGovernorV2.sol#98-104) does not emit an event
Setter function ZivoeGovernorV2.updateQuorumNumerator(uint256) (src/ZivoeGovernorV2.sol#107-110) does not emit an event
Setter function ZivoeITO.depositBoth(uint256,address,uint256,address) (src/ZivoeITO.sol#305-309) does not emit an event
Setter function ZivoeRewards.fullWithdraw() (src/ZivoeRewards.sol#246-249) does not emit an event
Setter function ZivoeRewards.getRewards() (src/ZivoeRewards.sol#281-283) does not emit an event
Setter function ZivoeRewards.updateReward(address) (src/ZivoeRewards.sol#123-134) does not emit an event
Setter function ZivoeRewardsVesting.fullWithdraw() (src/ZivoeRewardsVesting.sol#370-373) does not emit an event
Setter function ZivoeRewardsVesting.getRewards() (src/ZivoeRewardsVesting.sol#484-486) does not emit an event
Setter function ZivoeRewardsVesting.onlyZVLOrITO() (src/ZivoeRewardsVesting.sol#186-193) does not emit an event
Setter function ZivoeRewardsVesting.updateReward(address) (src/ZivoeRewardsVesting.sol#197-208) does not emit an event
Setter function ZivoeToken.burn(uint256) (src/ZivoeToken.sol#35) does not emit an event
Setter function ZivoeTrancheToken.burn(uint256) (src/ZivoeTrancheToken.sol#72) does not emit an event
Setter function ZivoeTrancheToken.mint(address,uint256) (src/ZivoeTrancheToken.sol#87) does not emit an event
Setter function ZivoeTrancheToken.isMinterRole() (src/ZivoeTrancheToken.sol#54-57) does not emit an event
Setter function ZivoeTranches.pushToLocker(address,uint256,bytes) (src/ZivoeTranches.sol#180-186) does not emit an event
Setter function ZivoeTranches.depositBoth(uint256,address,uint256,address) (src/ZivoeTranches.sol#322-325) does not emit an event
Setter function ZivoeTranches.switchPause() (src/ZivoeTranches.sol#328-334) does not emit an event
Setter function ZivoeTranches.unlock() (src/ZivoeTranches.sol#404-410) does not emit an event
Setter function ZivoeTranches.onlyGovernance() (src/ZivoeTranches.sol#154-160) does not emit an event
Setter function ZivoeYDL.unlock() (src/ZivoeYDL.sol#321-352) does not emit an event
Setter function ZivoeTLC._schedule(bytes32,uint256) (src/libraries/ZivoeTLC.sol#275-279) does not emit an event
Setter function ZivoeTLC._afterCall(bytes32) (src/libraries/ZivoeTLC.sol#398-401) does not emit an event
Setter function ZivoeTLC._afterCallKeeper(bytes32) (src/libraries/ZivoeTLC.sol#414-417) does not emit an event
Setter function ZivoeTLC.onlyRoleOrOpenRole(bytes32) (src/libraries/ZivoeTLC.sol#121-126) does not emit an event
Setter function OCC_Modular.constructor(address,address,address,address,address) (src/lockers/OCC/OCC_Modular.sol#156-162) does not emit an event
Setter function OCC_Modular.isUnderwriter() (src/lockers/OCC/OCC_Modular.sol#381-384) does not emit an event
Setter function OCE_ZVE.constructor(address,address) (src/lockers/OCE/OCE_ZVE.sol#70-78) does not emit an event
Setter function OCE_ZVE.pushToLocker(address,uint256,bytes) (src/lockers/OCE/OCE_ZVE.sol#122-128) does not emit an event
Setter function OCE_ZVE.forwardEmissions() (src/lockers/OCE/OCE_ZVE.sol#131-135) does not emit an event
Setter function OCG_Defaults.constructor(address,address) (src/lockers/OCG/OCG_Defaults.sol#29-32) does not emit an event
Setter function OCG_Defaults.decreaseDefaults(uint256) (src/lockers/OCG/OCG_Defaults.sol#57-59) does not emit an event
Setter function OCG_Defaults.increaseDefaults(uint256) (src/lockers/OCG/OCG_Defaults.sol#64-66) does not emit an event
Setter function OCG_Defaults.onlyGovernance() (src/lockers/OCG/OCG_Defaults.sol#41-47) does not emit an event
Setter function OCG_ERC1155.constructor(address) (src/lockers/OCG/OCG_ERC1155.sol#15-17) does not emit an event
Setter function OCG_ERC20.constructor(address) (src/lockers/OCG/OCG_ERC20.sol#15-17) does not emit an event
Setter function OCG_ERC20_FreeClaim.constructor(address) (src/lockers/OCG/OCG_ERC20_FreeClaim.sol#20-22) does not emit an event
Setter function OCG_ERC20_FreeClaim.claim(address,uint256) (src/lockers/OCG/OCG_ERC20_FreeClaim.sol#63-65) does not emit an event
Setter function OCG_ERC721.constructor(address) (src/lockers/OCG/OCG_ERC721.sol#15-17) does not emit an event
Setter function OCL_ZVE.constructor(address,address,address,address,address,address) (src/lockers/OCL/OCL_ZVE.sol#111-118) does not emit an event
Setter function OCL_ZVE.forwardYield() (src/lockers/OCL/OCL_ZVE.sol#287-305) does not emit an event
Setter function OCR_Modular.constructor(address,address,address,uint16) (src/lockers/OCR/OCR_Modular.sol#87-94) does not emit an event
Setter function OCR_Modular.pushToLocker(address,uint256,bytes) (src/lockers/OCR/OCR_Modular.sol#180-185) does not emit an event
Setter function OCR_Modular.pullFromLocker(address,bytes) (src/lockers/OCR/OCR_Modular.sol#190-196) does not emit an event
Setter function OCR_Modular.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCR/OCR_Modular.sol#202-210) does not emit an event
Setter function OCT_DAO.constructor(address,address) (src/lockers/OCT/OCT_DAO.sol#41-44) does not emit an event
Setter function OCT_YDL.constructor(address,address) (src/lockers/OCT/OCT_YDL.sol#47-50) does not emit an event
Setter function OCT_ZVL.constructor(address,address) (src/lockers/OCT/OCT_ZVL.sol#40-43) does not emit an event
Setter function OCY_Convex_A.constructor(address,address,address) (src/lockers/OCY/OCY_Convex_A.sol#81-85) does not emit an event
Setter function OCY_Convex_A.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_A.sol#117-169) does not emit an event
Setter function OCY_Convex_A.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_A.sol#174-207) does not emit an event
Setter function OCY_Convex_A.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_A.sol#213-242) does not emit an event
Setter function OCY_Convex_B.constructor(address,address,address) (src/lockers/OCY/OCY_Convex_B.sol#76-80) does not emit an event
Setter function OCY_Convex_B.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_B.sol#112-159) does not emit an event
Setter function OCY_Convex_B.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_B.sol#164-191) does not emit an event
Setter function OCY_Convex_B.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_B.sol#197-223) does not emit an event
Setter function OCY_Convex_C.constructor(address,address,address) (src/lockers/OCY/OCY_Convex_C.sol#74-78) does not emit an event
Setter function OCY_Convex_C.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_C.sol#110-141) does not emit an event
Setter function OCY_Convex_C.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_C.sol#146-169) does not emit an event
Setter function OCY_Convex_C.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_C.sol#175-196) does not emit an event
Setter function OCY_OUSD.constructor(address,address,address) (src/lockers/OCY/OCY_OUSD.sol#44-48) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

value assignment lack of validation	ZivoeGTC.queue(address[],uint256[],bytes[],bytes32) (src/libraries/ZivoeGTC.sol#91-108)_timelockIds[proposalId] = _timelock.hashOperationBatch(targets,values,calldatas,0,descriptionHash) (src/libraries/ZivoeGTC.sol#102)
Reference: input validation

variable lacks a zero-check on 		- ZivoeDAO.constructor(address) (src/ZivoeDAO.sol#169)
variable lacks a zero-check on 		- ZivoeDAO.push(address,address,uint256,bytes) (src/ZivoeDAO.sol#239-247)
variable lacks a zero-check on 		- ZivoeDAO.pull(address,address,bytes) (src/ZivoeDAO.sol#255-259)
variable lacks a zero-check on 		- ZivoeDAO.pullPartial(address,address,uint256,bytes) (src/ZivoeDAO.sol#268-274)
variable lacks a zero-check on 		- ZivoeDAO.pushERC721(address,address,uint256,bytes) (src/ZivoeDAO.sol#344-357)
variable lacks a zero-check on 		- ZivoeDAO.pullERC721(address,address,uint256,bytes) (src/ZivoeDAO.sol#395-401)
variable lacks a zero-check on 		- ZivoeDAO.pushERC1155(address,address,uint256[],uint256[],bytes) (src/ZivoeDAO.sol#429-446)
variable lacks a zero-check on 		- ZivoeDAO.pullERC1155(address,address,uint256[],uint256[],bytes) (src/ZivoeDAO.sol#455-465)
variable lacks a zero-check on 		- ZivoeGlobals.proposeZVL(address) (src/ZivoeGlobals.sol#210-212)
variable lacks a zero-check on 		- ZivoeGlobals.updateIsKeeper(address,bool) (src/ZivoeGlobals.sol#226-229)
variable lacks a zero-check on 		- ZivoeGlobals.updateIsLocker(address,bool) (src/ZivoeGlobals.sol#235-238)
variable lacks a zero-check on 		- ZivoeGlobals.updateStablecoinWhitelist(address,bool) (src/ZivoeGlobals.sol#244-247)
variable lacks a zero-check on 		- ZivoeGlobals.updateYDL(address) (src/ZivoeGlobals.sol#252-255)
variable lacks a zero-check on 		- Governor.relay(address,uint256,bytes) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#543-550)
variable lacks a zero-check on 		- ZivoeGovernorV2.constructor(IVotes,ZivoeTLC,address) (src/ZivoeGovernorV2.sol#62-64)
variable lacks a zero-check on 		- ZivoeITO.constructor(address,address[]) (src/ZivoeITO.sol#123-126)
variable lacks a zero-check on 		- ZivoeITO.depositJunior(uint256,address) (src/ZivoeITO.sol#248-271)
variable lacks a zero-check on 		- ZivoeITO.depositSenior(uint256,address) (src/ZivoeITO.sol#277-298)
variable lacks a zero-check on 		- ZivoeRewards.constructor(address,address) (src/ZivoeRewards.sol#72-75)
variable lacks a zero-check on 		- ZivoeRewardsVesting.constructor(address,address) (src/ZivoeRewardsVesting.sol#107-111)
variable lacks a zero-check on 		- ZivoeTrancheToken.changeMinterRole(address,bool) (src/ZivoeTrancheToken.sol#78-81)
variable lacks a zero-check on 		- ZivoeLocker.pullFromLocker(address,bytes) (src/ZivoeLocker.sol#81-84)
variable lacks a zero-check on 		- ZivoeLocker.pushToLockerERC721(address,uint256,bytes) (src/ZivoeLocker.sol#135-138)
variable lacks a zero-check on 		- ZivoeLocker.pullFromLockerERC721(address,uint256,bytes) (src/ZivoeLocker.sol#144-147)
variable lacks a zero-check on 		- ZivoeLocker.pushToLockerERC1155(address,uint256[],uint256[],bytes) (src/ZivoeLocker.sol#180-185)
variable lacks a zero-check on 		- ZivoeLocker.pullFromLockerERC1155(address,uint256[],uint256[],bytes) (src/ZivoeLocker.sol#192-197)
variable lacks a zero-check on 		- ZivoeTranches.constructor(address) (src/ZivoeTranches.sol#96)
variable lacks a zero-check on 		- ZivoeTranches.depositJunior(uint256,address) (src/ZivoeTranches.sol#268-289)
variable lacks a zero-check on 		- ZivoeTranches.depositSenior(uint256,address) (src/ZivoeTranches.sol#295-315)
variable lacks a zero-check on 		- ZivoeYDL.constructor(address,address) (src/ZivoeYDL.sol#129-133)
variable lacks a zero-check on 		- ZivoeTLC.constructor(uint256,address[],address[],address) (src/libraries/ZivoeTLC.sol#82-113)
variable lacks a zero-check on 		- OCC_Modular.constructor(address,address,address,address,address) (src/lockers/OCC/OCC_Modular.sol#156-162)
variable lacks a zero-check on 		- OCE_ZVE.constructor(address,address) (src/lockers/OCE/OCE_ZVE.sol#70-78)
variable lacks a zero-check on 		- OCG_Defaults.constructor(address,address) (src/lockers/OCG/OCG_Defaults.sol#29-32)
variable lacks a zero-check on 		- OCL_ZVE.constructor(address,address,address,address,address,address) (src/lockers/OCL/OCL_ZVE.sol#111-118)
variable lacks a zero-check on 		- OCR_Modular.constructor(address,address,address,uint16) (src/lockers/OCR/OCR_Modular.sol#87-94)
variable lacks a zero-check on 		- OCT_DAO.constructor(address,address) (src/lockers/OCT/OCT_DAO.sol#41-44)
variable lacks a zero-check on 		- OCT_DAO.convertAndForward(address,address,bytes) (src/lockers/OCT/OCT_DAO.sol#87-99)
variable lacks a zero-check on 		- OCT_YDL.constructor(address,address) (src/lockers/OCT/OCT_YDL.sol#47-50)
variable lacks a zero-check on 		- OCT_YDL.convertAndForward(address,bytes) (src/lockers/OCT/OCT_YDL.sol#97-110)
variable lacks a zero-check on 		- OCT_ZVL.constructor(address,address) (src/lockers/OCT/OCT_ZVL.sol#40-43)
variable lacks a zero-check on 		- OCY_Convex_A.constructor(address,address,address) (src/lockers/OCY/OCY_Convex_A.sol#81-85)
variable lacks a zero-check on 		- OCY_Convex_B.constructor(address,address,address) (src/lockers/OCY/OCY_Convex_B.sol#76-80)
variable lacks a zero-check on 		- OCY_Convex_C.constructor(address,address,address) (src/lockers/OCY/OCY_Convex_C.sol#74-78)
variable lacks a zero-check on 		- OCY_OUSD.constructor(address,address,address) (src/lockers/OCY/OCY_OUSD.sol#44-48)
variable lacks a zero-check on 		- Presale.constructor(address[],address,address) (src/misc/Presale.sol#51-65)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

Variable 'ERC20Votes._moveVotingPower(address,address,uint256).oldWeight (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#228)' in ERC20Votes._moveVotingPower(address,address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#221-237) potentially used before declaration: (oldWeight,newWeight) = _writeCheckpoint(_checkpoints[dst],_add,amount) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#233)
Reference:  

	- OwnableLocked.transferOwnershipAndLock(address) (src/libraries/OwnableLocked.sol#40-44)
	- ZivoeGlobals.decreaseDefaults(uint256) (src/ZivoeGlobals.sol#157-162)
	- ZivoeGlobals.increaseDefaults(uint256) (src/ZivoeGlobals.sol#168-173)
	- ZivoeGlobals.initializeGlobals(address[],address[]) (src/ZivoeGlobals.sol#179-205)
	- ZivoeGlobals.proposeZVL(address) (src/ZivoeGlobals.sol#210-212)
	- ZivoeGlobals.acceptZVL() (src/ZivoeGlobals.sol#215-220)
	- ZivoeGlobals.updateIsKeeper(address,bool) (src/ZivoeGlobals.sol#226-229)
	- ZivoeGlobals.updateIsLocker(address,bool) (src/ZivoeGlobals.sol#235-238)
	- ZivoeGlobals.updateStablecoinWhitelist(address,bool) (src/ZivoeGlobals.sol#244-247)
	- ZivoeGlobals.updateYDL(address) (src/ZivoeGlobals.sol#252-255)
	- ZivoeITO.depositJunior(uint256,address) (src/ZivoeITO.sol#248-271)
	- ZivoeITO.depositSenior(uint256,address) (src/ZivoeITO.sol#277-298)
	- ZivoeITO.migrateDeposits() (src/ZivoeITO.sol#313-335)
	- ZivoeITO.commence() (src/ZivoeITO.sol#339-347)
	- ZivoeRewards.addReward(address,uint256) (src/ZivoeRewards.sol#208-223)
	- ZivoeRewardsVesting.addReward(address,uint256) (src/ZivoeRewardsVesting.sol#328-347)
	- ZivoeRewardsVesting.createVestingSchedule(address,uint256,uint256,uint256,bool) (src/ZivoeRewardsVesting.sol#381-425)
	- ZivoeRewardsVesting.revokeVestingSchedule(address) (src/ZivoeRewardsVesting.sol#429-467)
	- ZivoeTrancheToken.changeMinterRole(address,bool) (src/ZivoeTrancheToken.sol#78-81)
	- ZivoeTranches.switchPause() (src/ZivoeTranches.sol#328-334)
	- ZivoeTranches.updateLowerRatioIncentiveBIPS(uint256) (src/ZivoeTranches.sol#343-354)
	- ZivoeTranches.updateMaxTrancheRatio(uint256) (src/ZivoeTranches.sol#360-364)
	- ZivoeTranches.updateMaxZVEPerJTTMint(uint256) (src/ZivoeTranches.sol#368-373)
	- ZivoeTranches.updateMinZVEPerJTTMint(uint256) (src/ZivoeTranches.sol#377-381)
	- ZivoeTranches.updateUpperRatioIncentiveBIPS(uint256) (src/ZivoeTranches.sol#390-401)
	- ZivoeTranches.unlock() (src/ZivoeTranches.sol#404-410)
	- ZivoeYDL.unlock() (src/ZivoeYDL.sol#321-352)
	- ZivoeYDL.updateDistributedAsset(address) (src/ZivoeYDL.sol#356-371)
	- ZivoeYDL.updateProtocolEarningsRateBIPS(uint256) (src/ZivoeYDL.sol#375-386)
	- ZivoeYDL.updateRecipients(address[],uint256[],bool) (src/ZivoeYDL.sol#392-416)
	- ZivoeYDL.updateTargetAPYBIPS(uint256) (src/ZivoeYDL.sol#420-424)
	- ZivoeYDL.updateTargetRatioBIPS(uint256) (src/ZivoeYDL.sol#428-432)
	- ZivoeTLC.cancel(bytes32) (src/libraries/ZivoeTLC.sol#288-293)
	- OCC_Modular.acceptOffer(uint256) (src/lockers/OCC/OCC_Modular.sol#461-487)
	- OCC_Modular.callLoan(uint256) (src/lockers/OCC/OCC_Modular.sol#492-518)
	- OCC_Modular.cancelOffer(uint256) (src/lockers/OCC/OCC_Modular.sol#522-529)
	- OCC_Modular.createOffer(address,uint256,uint256,uint256,uint256,uint256,uint256,int8) (src/lockers/OCC/OCC_Modular.sol#540-570)
	- OCC_Modular.markDefault(uint256) (src/lockers/OCC/OCC_Modular.sol#606-618)
	- OCC_Modular.markRepaid(uint256) (src/lockers/OCC/OCC_Modular.sol#622-630)
	- OCC_Modular.processPayment(uint256) (src/lockers/OCC/OCC_Modular.sol#636-677)
	- OCC_Modular.updateOCTYDL(address) (src/lockers/OCC/OCC_Modular.sol#727-732)
	- OCC_Modular.applyCombine(uint256) (src/lockers/OCC/OCC_Modular.sol#742-808)
	- OCC_Modular.applyConversionToAmortization(uint256) (src/lockers/OCC/OCC_Modular.sol#812-824)
	- OCC_Modular.applyConversionToBullet(uint256) (src/lockers/OCC/OCC_Modular.sol#828-840)
	- OCC_Modular.applyExtension(uint256) (src/lockers/OCC/OCC_Modular.sol#844-855)
	- OCC_Modular.applyRefinance(uint256) (src/lockers/OCC/OCC_Modular.sol#859-869)
	- OCC_Modular.approveCombine(uint256[],uint256,uint256,uint256,uint256,int8) (src/lockers/OCC/OCC_Modular.sol#877-905)
	- OCC_Modular.approveConversionToAmortization(uint256) (src/lockers/OCC/OCC_Modular.sol#909-912)
	- OCC_Modular.approveConversionToBullet(uint256) (src/lockers/OCC/OCC_Modular.sol#916-919)
	- OCC_Modular.approveExtension(uint256,uint256) (src/lockers/OCC/OCC_Modular.sol#924-927)
	- OCC_Modular.approveRefinance(uint256,uint256) (src/lockers/OCC/OCC_Modular.sol#932-935)
	- OCC_Modular.unapproveCombine(uint256) (src/lockers/OCC/OCC_Modular.sol#939-942)
	- OCC_Modular.unapproveConversionToAmortization(uint256) (src/lockers/OCC/OCC_Modular.sol#946-949)
	- OCC_Modular.unapproveConversionToBullet(uint256) (src/lockers/OCC/OCC_Modular.sol#953-956)
	- OCC_Modular.unapproveExtension(uint256) (src/lockers/OCC/OCC_Modular.sol#960-963)
	- OCC_Modular.unapproveRefinance(uint256) (src/lockers/OCC/OCC_Modular.sol#967-970)
	- OCE_ZVE.updateDistributionRatioBIPS(uint256[3]) (src/lockers/OCE/OCE_ZVE.sol#163-177)
	- OCE_ZVE.updateExponentialDecayPerSecond(uint256) (src/lockers/OCE/OCE_ZVE.sol#183-194)
	- OCL_ZVE.pushToLockerMulti(address[],uint256[],bytes[]) (src/lockers/OCL/OCL_ZVE.sol#172-215)
	- OCL_ZVE.pullFromLocker(address,bytes) (src/lockers/OCL/OCL_ZVE.sol#220-247)
	- OCL_ZVE.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCL/OCL_ZVE.sol#253-284)
	- OCL_ZVE.forwardYield() (src/lockers/OCL/OCL_ZVE.sol#287-305)
	- OCL_ZVE.updateCompoundingRateBIPS(uint256) (src/lockers/OCL/OCL_ZVE.sol#347-356)
	- OCL_ZVE.updateOCTYDL(address) (src/lockers/OCL/OCL_ZVE.sol#361-369)
	- OCR_Modular.destroyRequest(uint256) (src/lockers/OCR/OCR_Modular.sol#232-251)
	- OCR_Modular.updateRedemptionsFeeBIPS(uint256) (src/lockers/OCR/OCR_Modular.sol#340-350)
	- OCY_Convex_A.updateOCTYDL(address) (src/lockers/OCY/OCY_Convex_A.sol#271-279)
	- OCY_Convex_B.updateOCTYDL(address) (src/lockers/OCY/OCY_Convex_B.sol#252-260)
	- OCY_Convex_C.updateOCTYDL(address) (src/lockers/OCY/OCY_Convex_C.sol#225-233)
	- OCY_OUSD.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_OUSD.sol#89-94)
	- OCY_OUSD.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_OUSD.sol#99-107)
	- OCY_OUSD.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_OUSD.sol#113-123)
	- OCY_OUSD.updateOCTYDL(address) (src/lockers/OCY/OCY_OUSD.sol#134-142)
Reference:  

	- AccessControl.grantRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#144-146)
	- AccessControl.revokeRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#159-161)
	- Ownable.renounceOwnership() (lib/openzeppelin-contracts/contracts/access/Ownable.sol#61-63)
	- Ownable.transferOwnership(address) (lib/openzeppelin-contracts/contracts/access/Ownable.sol#69-72)
	- GovernorSettings.setVotingDelay(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#61-63)
	- GovernorSettings.setVotingPeriod(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#70-72)
	- GovernorSettings.setProposalThreshold(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#79-81)
	- GovernorVotesQuorumFraction.updateQuorumNumerator(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#86-88)
	- OwnableLocked.renounceOwnership() (src/libraries/OwnableLocked.sol#25)
	- OwnableLocked.transferOwnership(address) (src/libraries/OwnableLocked.sol#31-34)
	- ZivoeDAO.push(address,address,uint256,bytes) (src/ZivoeDAO.sol#239-247)
	- ZivoeDAO.pull(address,address,bytes) (src/ZivoeDAO.sol#255-259)
	- ZivoeDAO.pullPartial(address,address,uint256,bytes) (src/ZivoeDAO.sol#268-274)
	- ZivoeDAO.pushMulti(address,address[],uint256[],bytes[]) (src/ZivoeDAO.sol#282-301)
	- ZivoeDAO.pullMulti(address,address[],bytes[]) (src/ZivoeDAO.sol#309-316)
	- ZivoeDAO.pullMultiPartial(address,address[],uint256[],bytes[]) (src/ZivoeDAO.sol#325-336)
	- ZivoeDAO.pushERC721(address,address,uint256,bytes) (src/ZivoeDAO.sol#344-357)
	- ZivoeDAO.pushMultiERC721(address,address[],uint256[],bytes[]) (src/ZivoeDAO.sol#365-387)
	- ZivoeDAO.pullERC721(address,address,uint256,bytes) (src/ZivoeDAO.sol#395-401)
	- ZivoeDAO.pullMultiERC721(address,address[],uint256[],bytes[]) (src/ZivoeDAO.sol#409-420)
	- ZivoeDAO.pushERC1155(address,address,uint256[],uint256[],bytes) (src/ZivoeDAO.sol#429-446)
	- ZivoeDAO.pullERC1155(address,address,uint256[],uint256[],bytes) (src/ZivoeDAO.sol#455-465)
	- ZivoeGovernorV2.setVotingDelay(uint256) (src/ZivoeGovernorV2.sol#84-87)
	- ZivoeGovernorV2.setVotingPeriod(uint256) (src/ZivoeGovernorV2.sol#91-95)
	- ZivoeGovernorV2.setProposalThreshold(uint256) (src/ZivoeGovernorV2.sol#98-104)
	- ZivoeGovernorV2.updateQuorumNumerator(uint256) (src/ZivoeGovernorV2.sol#107-110)
	- ZivoeLocker.pushToLocker(address,uint256,bytes) (src/ZivoeLocker.sol#73-76)
	- ZivoeLocker.pullFromLocker(address,bytes) (src/ZivoeLocker.sol#81-84)
	- ZivoeLocker.pullFromLockerPartial(address,uint256,bytes) (src/ZivoeLocker.sol#90-93)
	- ZivoeLocker.pushToLockerMulti(address[],uint256[],bytes[]) (src/ZivoeLocker.sol#99-106)
	- ZivoeLocker.pullFromLockerMulti(address[],bytes[]) (src/ZivoeLocker.sol#111-116)
	- ZivoeLocker.pullFromLockerMultiPartial(address[],uint256[],bytes[]) (src/ZivoeLocker.sol#122-129)
	- ZivoeLocker.pushToLockerERC721(address,uint256,bytes) (src/ZivoeLocker.sol#135-138)
	- ZivoeLocker.pullFromLockerERC721(address,uint256,bytes) (src/ZivoeLocker.sol#144-147)
	- ZivoeLocker.pushToLockerMultiERC721(address[],uint256[],bytes[]) (src/ZivoeLocker.sol#153-160)
	- ZivoeLocker.pullFromLockerMultiERC721(address[],uint256[],bytes[]) (src/ZivoeLocker.sol#166-173)
	- ZivoeLocker.pushToLockerERC1155(address,uint256[],uint256[],bytes) (src/ZivoeLocker.sol#180-185)
	- ZivoeLocker.pullFromLockerERC1155(address,uint256[],uint256[],bytes) (src/ZivoeLocker.sol#192-197)
	- ZivoeTranches.pushToLocker(address,uint256,bytes) (src/ZivoeTranches.sol#180-186)
	- ZivoeTLC.schedule(address,uint256,bytes,bytes32,bytes32,uint256) (src/libraries/ZivoeTLC.sol#232-243)
	- ZivoeTLC.scheduleBatch(address[],uint256[],bytes[],bytes32,bytes32,uint256) (src/libraries/ZivoeTLC.sol#254-270)
	- ZivoeTLC.execute(address,uint256,bytes,bytes32,bytes32) (src/libraries/ZivoeTLC.sol#307-328)
	- ZivoeTLC.executeBatch(address[],uint256[],bytes[],bytes32,bytes32) (src/libraries/ZivoeTLC.sol#339-373)
	- OCE_ZVE.pushToLocker(address,uint256,bytes) (src/lockers/OCE/OCE_ZVE.sol#122-128)
	- OCG_Defaults.decreaseDefaults(uint256) (src/lockers/OCG/OCG_Defaults.sol#57-59)
	- OCG_Defaults.increaseDefaults(uint256) (src/lockers/OCG/OCG_Defaults.sol#64-66)
	- OCR_Modular.pushToLocker(address,uint256,bytes) (src/lockers/OCR/OCR_Modular.sol#180-185)
	- OCR_Modular.pullFromLocker(address,bytes) (src/lockers/OCR/OCR_Modular.sol#190-196)
	- OCR_Modular.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCR/OCR_Modular.sol#202-210)
	- OCT_DAO.convertAndForward(address,address,bytes) (src/lockers/OCT/OCT_DAO.sol#87-99)
	- OCT_YDL.convertAndForward(address,bytes) (src/lockers/OCT/OCT_YDL.sol#97-110)
	- OCT_ZVL.claim() (src/lockers/OCT/OCT_ZVL.sol#81-87)
	- OCY_Convex_A.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_A.sol#117-169)
	- OCY_Convex_A.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_A.sol#174-207)
	- OCY_Convex_A.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_A.sol#213-242)
	- OCY_Convex_B.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_B.sol#112-159)
	- OCY_Convex_B.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_B.sol#164-191)
	- OCY_Convex_B.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_B.sol#197-223)
	- OCY_Convex_C.pushToLocker(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_C.sol#110-141)
	- OCY_Convex_C.pullFromLocker(address,bytes) (src/lockers/OCY/OCY_Convex_C.sol#146-169)
	- OCY_Convex_C.pullFromLockerPartial(address,uint256,bytes) (src/lockers/OCY/OCY_Convex_C.sol#175-196)
Reference:  

ZivoeVotes._add(uint256,uint256) (src/libraries/ZivoeVotes.sol#129-131) is never used and should be removed
ZivoeVotes._subtract(uint256,uint256) (src/libraries/ZivoeVotes.sol#133-135) is never used and should be removed
Reference:  

require() missing error messages
	 - require(bool)(denominator > prod1) (lib/openzeppelin-contracts/contracts/utils/math/Math.sol#78)
require() missing error messages
	 - require(bool)(_executor() == address(this)) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#86)
require() missing error messages
	 - require(bool)(_msgSender() == IZivoeGlobals_OCC(GBL).ZVL()) (src/lockers/OCC/OCC_Modular.sol#728)
require() missing error messages
	 - require(bool)(_OCT_YDL != address(0)) (src/lockers/OCC/OCC_Modular.sol#729)
Reference: https://dev.to/tawseef/require-vs-assert-in-solidity-5e9d

Potential price manipulation risk:
	- In function _delegate
		-- delegatorBalance = balanceOf(delegator) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#213) have potential price manipulated risk from delegatorBalance and call None which could influence variable:delegatorBalance
Potential price manipulation risk:
	- In function pushToLockerMulti
		-- (preBasis,None) = fetchBasis() (src/lockers/OCL/OCL_ZVE.sol#189) have potential price manipulated risk from preBasis and call None which could influence variable:basis
	- In function pushToLockerMulti
		-- (postBasis) = fetchBasis() (src/lockers/OCL/OCL_ZVE.sol#212) have potential price manipulated risk from postBasis and call None which could influence variable:basis
	- In function fetchBasis
		-- pairAssetBalance = IERC20(pairAsset).balanceOf(pool) (src/lockers/OCL/OCL_ZVE.sol#338) have potential price manipulated risk from pairAssetBalance and call None which could influence variable:pairAssetBalance
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

Potential DoS Gas Limit Attack occur inGovernor._execute(uint256,address[],uint256[],bytes[],bytes32) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#316-328)BEGIN_LOOP (lib/openzeppelin-contracts/contracts/governance/Governor.sol#324-327)
Potential DoS Gas Limit Attack occur inGovernor._beforeExecute(uint256,address[],uint256[],bytes[],bytes32) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#333-347)BEGIN_LOOP (lib/openzeppelin-contracts/contracts/governance/Governor.sol#341-345)
Potential DoS Gas Limit Attack occur inZivoeYDL.distributeYield() (src/ZivoeYDL.sol#213-310)BEGIN_LOOP (src/ZivoeYDL.sol#242-269)
Potential DoS Gas Limit Attack occur inOCR_Modular.tickEpoch() (src/lockers/OCR/OCR_Modular.sol#310-336)BEGIN_LOOP (src/lockers/OCR/OCR_Modular.sol#311-335)
Potential DoS Gas Limit Attack occur inOCY_Convex_A.claimRewards(bool) (src/lockers/OCY/OCY_Convex_A.sol#246-266)BEGIN_LOOP (src/lockers/OCY/OCY_Convex_A.sol#258-264)
Potential DoS Gas Limit Attack occur inOCY_Convex_B.claimRewards(bool) (src/lockers/OCY/OCY_Convex_B.sol#227-247)BEGIN_LOOP (src/lockers/OCY/OCY_Convex_B.sol#239-245)
Potential DoS Gas Limit Attack occur inOCY_Convex_C.claimRewards(bool) (src/lockers/OCY/OCY_Convex_C.sol#200-220)BEGIN_LOOP (src/lockers/OCY/OCY_Convex_C.sol#212-218)
Reference: https://swcregistry.io/docs/SWC-128

function:AccessControl.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#77-79)is public and can be replaced with external 
function:AccessControl.grantRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#144-146)is public and can be replaced with external 
function:AccessControl.revokeRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#159-161)is public and can be replaced with external 
function:AccessControl.renounceRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#179-183)is public and can be replaced with external 
function:Ownable.renounceOwnership() (lib/openzeppelin-contracts/contracts/access/Ownable.sol#61-63)is public and can be replaced with external 
function:Ownable.transferOwnership(address) (lib/openzeppelin-contracts/contracts/access/Ownable.sol#69-72)is public and can be replaced with external 
function:IGovernor.name() (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#76)is public and can be replaced with external 
function:IGovernor.version() (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#82)is public and can be replaced with external 
function:IGovernor.COUNTING_MODE() (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#107)is public and can be replaced with external 
function:IGovernor.hashProposal(address[],uint256[],bytes[],bytes32) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#113-118)is public and can be replaced with external 
function:IGovernor.state(uint256) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#124)is public and can be replaced with external 
function:IGovernor.proposalSnapshot(uint256) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#132)is public and can be replaced with external 
function:IGovernor.proposalDeadline(uint256) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#139)is public and can be replaced with external 
function:IGovernor.quorum(uint256) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#164)is public and can be replaced with external 
function:IGovernor.getVotes(address,uint256) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#173)is public and can be replaced with external 
function:IGovernor.getVotesWithParams(address,uint256,bytes) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#179-183)is public and can be replaced with external 
function:IGovernor.hasVoted(uint256,address) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#189)is public and can be replaced with external 
function:IGovernor.propose(address[],uint256[],bytes[],string) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#197-202)is public and can be replaced with external 
function:IGovernor.execute(address[],uint256[],bytes[],bytes32) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#212-217)is public and can be replaced with external 
function:IGovernor.castVote(uint256,uint8) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#224)is public and can be replaced with external 
function:IGovernor.castVoteWithReason(uint256,uint8,string) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#231-235)is public and can be replaced with external 
function:IGovernor.castVoteWithReasonAndParams(uint256,uint8,string,bytes) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#242-247)is public and can be replaced with external 
function:IGovernor.castVoteBySig(uint256,uint8,uint8,bytes32,bytes32) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#254-260)is public and can be replaced with external 
function:IGovernor.castVoteWithReasonAndParamsBySig(uint256,uint8,string,bytes,uint8,bytes32,bytes32) (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#267-275)is public and can be replaced with external 
function:Governor.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#92-104)is public and can be replaced with external 
function:Governor.name() (lib/openzeppelin-contracts/contracts/governance/Governor.sol#109-111)is public and can be replaced with external 
function:Governor.propose(address[],uint256[],bytes[],string) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#245-284)is public and can be replaced with external 
function:Governor.execute(address[],uint256[],bytes[],bytes32) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#289-311)is public and can be replaced with external 
function:Governor.getVotesWithParams(address,uint256,bytes) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#402-408)is public and can be replaced with external 
function:Governor.castVote(uint256,uint8) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#413-416)is public and can be replaced with external 
function:Governor.castVoteWithReason(uint256,uint8,string) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#421-428)is public and can be replaced with external 
function:Governor.castVoteWithReasonAndParams(uint256,uint8,string,bytes) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#433-441)is public and can be replaced with external 
function:Governor.castVoteBySig(uint256,uint8,uint8,bytes32,bytes32) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#446-460)is public and can be replaced with external 
function:Governor.castVoteWithReasonAndParamsBySig(uint256,uint8,string,bytes,uint8,bytes32,bytes32) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#465-492)is public and can be replaced with external 
function:Governor.onERC721Received(address,address,uint256,bytes) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#563-570)is public and can be replaced with external 
function:Governor.onERC1155Received(address,address,uint256,uint256,bytes) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#575-583)is public and can be replaced with external 
function:Governor.onERC1155BatchReceived(address,address,uint256[],uint256[],bytes) (lib/openzeppelin-contracts/contracts/governance/Governor.sol#588-596)is public and can be replaced with external 
function:IGovernor.votingDelay() (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#146)is public and can be replaced with external 
function:IGovernor.votingPeriod() (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#155)is public and can be replaced with external 
function:GovernorCountingSimple.COUNTING_MODE() (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#36-38)is public and can be replaced with external 
function:GovernorCountingSimple.hasVoted(uint256,address) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#43-45)is public and can be replaced with external 
function:GovernorCountingSimple.proposalVotes(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#50-62)is public and can be replaced with external 
function:Governor.proposalThreshold() (lib/openzeppelin-contracts/contracts/governance/Governor.sol#196-198)is public and can be replaced with external 
function:GovernorSettings.setVotingDelay(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#61-63)is public and can be replaced with external 
function:GovernorSettings.setVotingPeriod(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#70-72)is public and can be replaced with external 
function:GovernorSettings.setProposalThreshold(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#79-81)is public and can be replaced with external 
function:GovernorVotesQuorumFraction.quorum(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#72-74)is public and can be replaced with external 
function:IGovernorTimelock.timelock() (lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#16)is public and can be replaced with external 
function:IGovernorTimelock.proposalEta(uint256) (lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#18)is public and can be replaced with external 
function:IGovernorTimelock.queue(address[],uint256[],bytes[],bytes32) (lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#20-25)is public and can be replaced with external 
function:ERC1155Receiver.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Receiver.sol#16-18)is public and can be replaced with external 
function:ERC1155Holder.onERC1155Received(address,address,uint256,uint256,bytes) (lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#17-25)is public and can be replaced with external 
function:ERC1155Holder.onERC1155BatchReceived(address,address,uint256[],uint256[],bytes) (lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#27-35)is public and can be replaced with external 
function:ERC20.name() (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#62-64)is public and can be replaced with external 
function:ERC20.symbol() (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#70-72)is public and can be replaced with external 
function:ERC20.decimals() (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#87-89)is public and can be replaced with external 
function:ERC20.totalSupply() (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#94-96)is public and can be replaced with external 
function:ERC20.balanceOf(address) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#101-103)is public and can be replaced with external 
function:ERC20.transfer(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#113-117)is public and can be replaced with external 
function:ERC20.approve(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#136-140)is public and can be replaced with external 
function:ERC20.transferFrom(address,address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#158-167)is public and can be replaced with external 
function:ERC20.increaseAllowance(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#181-185)is public and can be replaced with external 
function:ERC20.decreaseAllowance(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#201-210)is public and can be replaced with external 
function:ERC20Permit.permit(address,address,uint256,uint256,uint8,bytes32,bytes32) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#49-68)is public and can be replaced with external 
function:ERC20Permit.nonces(address) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#73-75)is public and can be replaced with external 
function:ERC20Votes.checkpoints(address,uint32) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#43-45)is public and can be replaced with external 
function:ERC20Votes.numCheckpoints(address) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#50-52)is public and can be replaced with external 
function:ERC20Votes.getVotes(address) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#64-67)is public and can be replaced with external 
function:ERC20Votes.getPastVotes(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#76-79)is public and can be replaced with external 
function:ERC20Votes.getPastTotalSupply(uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#89-92)is public and can be replaced with external 
function:ERC20Votes.delegate(address) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#139-141)is public and can be replaced with external 
function:ERC20Votes.delegateBySig(address,uint256,uint256,uint8,bytes32,bytes32) (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#146-163)is public and can be replaced with external 
function:ERC721Holder.onERC721Received(address,address,uint256,bytes) (lib/openzeppelin-contracts/contracts/token/ERC721/utils/ERC721Holder.sol#20-27)is public and can be replaced with external 
function:ERC165.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#26-28)is public and can be replaced with external 
function:OwnableLocked.renounceOwnership() (src/libraries/OwnableLocked.sol#25)is public and can be replaced with external 
function:OwnableLocked.transferOwnership(address) (src/libraries/OwnableLocked.sol#31-34)is public and can be replaced with external 
function:OwnableLocked.transferOwnershipAndLock(address) (src/libraries/OwnableLocked.sol#40-44)is public and can be replaced with external 
function:ZivoeGTC.supportsInterface(bytes4) (src/libraries/ZivoeGTC.sol#46-48)is public and can be replaced with external 
function:ZivoeGTC.state(uint256) (src/libraries/ZivoeGTC.sol#53-71)is public and can be replaced with external 
function:ZivoeGTC.timelock() (src/libraries/ZivoeGTC.sol#76-78)is public and can be replaced with external 
function:ZivoeGTC.proposalEta(uint256) (src/libraries/ZivoeGTC.sol#83-86)is public and can be replaced with external 
function:ZivoeGTC.queue(address[],uint256[],bytes[],bytes32) (src/libraries/ZivoeGTC.sol#91-108)is public and can be replaced with external 
function:GovernorSettings.proposalThreshold() (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#52-54)is public and can be replaced with external 
function:ZivoeGovernorV2.setVotingDelay(uint256) (src/ZivoeGovernorV2.sol#84-87)is public and can be replaced with external 
function:ZivoeGovernorV2.setVotingPeriod(uint256) (src/ZivoeGovernorV2.sol#91-95)is public and can be replaced with external 
function:ZivoeGovernorV2.setProposalThreshold(uint256) (src/ZivoeGovernorV2.sol#98-104)is public and can be replaced with external 
function:ZivoeGovernorV2.updateQuorumNumerator(uint256) (src/ZivoeGovernorV2.sol#107-110)is public and can be replaced with external 
function:ZivoeGovernorV2.supportsInterface(bytes4) (src/ZivoeGovernorV2.sol#118-120)is public and can be replaced with external 
function:ZivoeMath.yieldTarget(uint256,uint256,uint256,uint256,uint256) (src/ZivoeMath.sol#121-123)is public and can be replaced with external 
function:ZivoeVotes.checkpoints(address,uint32) (src/libraries/ZivoeVotes.sol#24-26)is public and can be replaced with external 
function:ZivoeVotes.numCheckpoints(address) (src/libraries/ZivoeVotes.sol#31-33)is public and can be replaced with external 
function:ZivoeVotes.getVotes(address) (src/libraries/ZivoeVotes.sol#38-41)is public and can be replaced with external 
function:ZivoeVotes.getPastVotes(address,uint256) (src/libraries/ZivoeVotes.sol#50-53)is public and can be replaced with external 
function:ZivoeVotes.getPastTotalSupply(uint256) (src/libraries/ZivoeVotes.sol#63-66)is public and can be replaced with external 
function:ZivoeRewards.earned(address,address) (src/ZivoeRewards.sol#180-184)is public and can be replaced with external 
function:ZivoeRewardsVesting.earned(address,address) (src/ZivoeRewardsVesting.sol#300-304)is public and can be replaced with external 
function:ZivoeLocker.canPush() (src/ZivoeLocker.sol#34)is public and can be replaced with external 
function:ZivoeLocker.canPull() (src/ZivoeLocker.sol#37)is public and can be replaced with external 
function:ZivoeLocker.canPullPartial() (src/ZivoeLocker.sol#40)is public and can be replaced with external 
function:ZivoeTLC.supportsInterface(bytes4) (src/libraries/ZivoeTLC.sol#136-138)is public and can be replaced with external 
function:ZivoeTLC.schedule(address,uint256,bytes,bytes32,bytes32,uint256) (src/libraries/ZivoeTLC.sol#232-243)is public and can be replaced with external 
function:ZivoeTLC.scheduleBatch(address[],uint256[],bytes[],bytes32,bytes32,uint256) (src/libraries/ZivoeTLC.sol#254-270)is public and can be replaced with external 
function:ZivoeTLC.cancel(bytes32) (src/libraries/ZivoeTLC.sol#288-293)is public and can be replaced with external 
function:ZivoeTLC.execute(address,uint256,bytes,bytes32,bytes32) (src/libraries/ZivoeTLC.sol#307-328)is public and can be replaced with external 
function:ZivoeTLC.executeBatch(address[],uint256[],bytes[],bytes32,bytes32) (src/libraries/ZivoeTLC.sol#339-373)is public and can be replaced with external 
function:ZivoeTLC.onERC721Received(address,address,uint256,bytes) (src/libraries/ZivoeTLC.sol#440-447)is public and can be replaced with external 
function:ZivoeTLC.onERC1155Received(address,address,uint256,uint256,bytes) (src/libraries/ZivoeTLC.sol#452-460)is public and can be replaced with external 
function:ZivoeTLC.onERC1155BatchReceived(address,address,uint256[],uint256[],bytes) (src/libraries/ZivoeTLC.sol#465-473)is public and can be replaced with external 
function:OCG_Defaults.decreaseDefaults(uint256) (src/lockers/OCG/OCG_Defaults.sol#57-59)is public and can be replaced with external 
function:OCG_Defaults.increaseDefaults(uint256) (src/lockers/OCG/OCG_Defaults.sol#64-66)is public and can be replaced with external 
function:ZivoeLocker.canPushERC1155() (src/ZivoeLocker.sol#64)is public and can be replaced with external 
function:ZivoeLocker.canPullERC1155() (src/ZivoeLocker.sol#67)is public and can be replaced with external 
function:ZivoeLocker.canPushMulti() (src/ZivoeLocker.sol#43)is public and can be replaced with external 
function:ZivoeLocker.canPullMulti() (src/ZivoeLocker.sol#46)is public and can be replaced with external 
function:ZivoeLocker.canPullMultiPartial() (src/ZivoeLocker.sol#49)is public and can be replaced with external 
function:ZivoeLocker.canPushERC721() (src/ZivoeLocker.sol#52)is public and can be replaced with external 
function:ZivoeLocker.canPullERC721() (src/ZivoeLocker.sol#55)is public and can be replaced with external 
function:ZivoeLocker.canPushMultiERC721() (src/ZivoeLocker.sol#58)is public and can be replaced with external 
function:ZivoeLocker.canPullMultiERC721() (src/ZivoeLocker.sol#61)is public and can be replaced with external 
function:OCR_Modular.tickEpoch() (src/lockers/OCR/OCR_Modular.sol#310-336)is public and can be replaced with external 
function:OCY_OUSD.rebase() (src/lockers/OCY/OCY_OUSD.sol#127-129)is public and can be replaced with external 
function:MockStablecoin.decimals() (src/misc/MockStablecoin.sol#47-49)is public and can be replaced with external 
function:Presale.depositStablecoin(address,uint256) (src/misc/Presale.sol#147-165)is public and can be replaced with external 
function:Presale.depositETH() (src/misc/Presale.sol#168-184)is public and can be replaced with external 
Reference:  

	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/access/IAccessControl.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/access/Ownable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/governance/Governor.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotes.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/governance/utils/IVotes.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/security/ReentrancyGuard.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC1155/IERC1155.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC1155/IERC1155Receiver.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Receiver.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/IERC20Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC721/IERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC721/IERC721Receiver.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/token/ERC721/utils/ERC721Holder.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (lib/openzeppelin-contracts/contracts/utils/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/Checkpoints.sol#5)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/Context.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/Counters.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/Strings.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/Timers.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/cryptography/EIP712.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/introspection/IERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/math/Math.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/math/SafeCast.sol#5)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.4. ^0.8.4 (lib/openzeppelin-contracts/contracts/utils/structs/DoubleEndedQueue.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeDAO.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeGlobals.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeGovernorV2.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeITO.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeLocker.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeMath.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeRewards.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeRewardsVesting.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeToken.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeTrancheToken.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeTranches.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/ZivoeYDL.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/libraries/FloorMath.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/libraries/OwnableLocked.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (src/libraries/ZivoeGTC.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (src/libraries/ZivoeTLC.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (src/libraries/ZivoeVotes.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCC/OCC_Modular.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCE/OCE_ZVE.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCG/OCG_Defaults.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCG/OCG_ERC1155.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCG/OCG_ERC20.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCG/OCG_ERC20_FreeClaim.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCG/OCG_ERC721.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCL/OCL_ZVE.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCR/OCR_Modular.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCT/OCT_DAO.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCT/OCT_YDL.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCT/OCT_ZVL.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCY/OCY_Convex_A.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCY/OCY_Convex_B.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCY/OCY_Convex_C.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/OCY/OCY_OUSD.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/lockers/Utility/ZivoeSwapper.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/misc/BaseContractTemplate.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/misc/InterfacesAggregated.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/misc/MockStablecoin.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.17. ^0.8.17 (src/misc/Presale.sol#2)
Reference:  

ERC20Permit._PERMIT_TYPEHASH_DEPRECATED_SLOT (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#37) should be constant
GovernorVotesQuorumFraction._quorumNumerator (lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#19) should be constant
OCY_Convex_A.convexDeposit (src/lockers/OCY/OCY_Convex_A.sol#60) should be constant
OCY_Convex_A.convexPoolID (src/lockers/OCY/OCY_Convex_A.sol#64) should be constant
OCY_Convex_A.convexPoolToken (src/lockers/OCY/OCY_Convex_A.sol#62) should be constant
OCY_Convex_A.convexRewards (src/lockers/OCY/OCY_Convex_A.sol#61) should be constant
OCY_Convex_A.curveBasePool (src/lockers/OCY/OCY_Convex_A.sol#67) should be constant
OCY_Convex_A.curveBasePoolToken (src/lockers/OCY/OCY_Convex_A.sol#68) should be constant
OCY_Convex_A.curveMetaPool (src/lockers/OCY/OCY_Convex_A.sol#69) should be constant
OCY_Convex_B.convexDeposit (src/lockers/OCY/OCY_Convex_B.sol#56) should be constant
OCY_Convex_B.convexPoolID (src/lockers/OCY/OCY_Convex_B.sol#60) should be constant
OCY_Convex_B.convexPoolToken (src/lockers/OCY/OCY_Convex_B.sol#57) should be constant
OCY_Convex_B.convexRewards (src/lockers/OCY/OCY_Convex_B.sol#58) should be constant
OCY_Convex_B.curveBasePool (src/lockers/OCY/OCY_Convex_B.sol#63) should be constant
OCY_Convex_B.curveBasePoolToken (src/lockers/OCY/OCY_Convex_B.sol#64) should be constant
OCY_Convex_C.convexDeposit (src/lockers/OCY/OCY_Convex_C.sol#54) should be constant
OCY_Convex_C.convexPoolID (src/lockers/OCY/OCY_Convex_C.sol#58) should be constant
OCY_Convex_C.convexPoolToken (src/lockers/OCY/OCY_Convex_C.sol#55) should be constant
OCY_Convex_C.convexRewards (src/lockers/OCY/OCY_Convex_C.sol#56) should be constant
OCY_Convex_C.curveBasePool (src/lockers/OCY/OCY_Convex_C.sol#61) should be constant
OCY_Convex_C.curveBasePoolToken (src/lockers/OCY/OCY_Convex_C.sol#62) should be constant
Presale.pointsCeiling (src/misc/Presale.sol#31) should be constant
Presale.pointsFloor (src/misc/Presale.sol#29) should be constant
Presale.presaleDuration (src/misc/Presale.sol#35) should be constant
Reference:  
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
An error occurred: maximum recursion depth exceeded while calling a Python object
after for recrusion
Summary
 - [reentrancy-with-eth-transfer](#reentrancy-with-eth-transfer) (4 results) (Critical)
 - [state-variable-not-initialized](#state-variable-not-initialized) (7 results) (High)
 - [centralized-risk-medium](#centralized-risk-medium) (1 results) (Medium)
 - [public-mint-burn](#public-mint-burn) (1 results) (Medium)
 - [uninitialized-local](#uninitialized-local) (2 results) (Medium)
 - [unused-return](#unused-return) (15 results) (Medium)
 - [void-function](#void-function) (3 results) (Medium)
 - [pess-event-setter](#pess-event-setter) (66 results) (Low)
 - [input-validation](#input-validation) (1 results) (Low)
 - [missing-zero-check](#missing-zero-check) (46 results) (Low)
 - [variable-scope](#variable-scope) (1 results) (Low)
 - [centralized-risk-informational](#centralized-risk-informational) (73 results) (Informational)
 - [centralized-risk-other](#centralized-risk-other) (61 results) (Informational)
 - [dead-function](#dead-function) (2 results) (Informational)
 - [error-msg](#error-msg) (4 results) (Informational)
 - [price-manipulation-info](#price-manipulation-info) (2 results) (Informational)
 - [uncontrolled-resource-consumption](#uncontrolled-resource-consumption) (7 results) (Informational)
 - [unnecessary-public-function-modifier](#unnecessary-public-function-modifier) (124 results) (Informational)
 - [version-only](#version-only) (78 results) (Informational)
 - [state-should-be-constant](#state-should-be-constant) (24 results) (Optimization)
## reentrancy-with-eth-transfer
Impact: Critical
Confidence: Medium
 - [ ] ID-0
Reentrancy in [ZivoeRewardsVesting.fullWithdraw()](src/ZivoeRewardsVesting.sol#L370-L373):
	External calls:
	- [withdraw()](src/ZivoeRewardsVesting.sol#L371)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
	- [getRewards()](src/ZivoeRewardsVesting.sol#L372)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
		- [IERC20(_rewardsToken).safeTransfer(_msgSender(),reward)](src/ZivoeRewardsVesting.sol#L495)
	External calls sending eth:
	- [withdraw()](src/ZivoeRewardsVesting.sol#L371)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
	- [getRewards()](src/ZivoeRewardsVesting.sol#L372)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)

src/ZivoeRewardsVesting.sol#L370-L373


 - [ ] ID-1
Reentrancy in [OCL_ZVE.forwardYield()](src/lockers/OCL/OCL_ZVE.sol#L287-L305):
	External calls:
	- [_forwardYield(amount,lp)](src/lockers/OCL/OCL_ZVE.sol#L302)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
	External calls sending eth:
	- [_forwardYield(amount,lp)](src/lockers/OCL/OCL_ZVE.sol#L302)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)

src/lockers/OCL/OCL_ZVE.sol#L287-L305


 - [ ] ID-2
Reentrancy in [ZivoeRewards.fullWithdraw()](src/ZivoeRewards.sol#L246-L249):
	External calls:
	- [withdraw(_balances[_msgSender()])](src/ZivoeRewards.sol#L247)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
	- [getRewards()](src/ZivoeRewards.sol#L248)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
		- [IERC20(_rewardsToken).safeTransfer(_msgSender(),reward)](src/ZivoeRewards.sol#L292)
	External calls sending eth:
	- [withdraw(_balances[_msgSender()])](src/ZivoeRewards.sol#L247)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
	- [getRewards()](src/ZivoeRewards.sol#L248)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)

src/ZivoeRewards.sol#L246-L249


 - [ ] ID-3
Reentrancy in [ZivoeTranches.depositBoth(uint256,address,uint256,address)](src/ZivoeTranches.sol#L322-L325):
	External calls:
	- [depositSenior(amountSenior,assetSenior)](src/ZivoeTranches.sol#L323)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [IERC20(asset).safeTransferFrom(depositor,IZivoeGlobals_ZivoeTranches(GBL).DAO(),amount)](src/ZivoeTranches.sol#L304)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
	- [depositJunior(amountJunior,assetJunior)](src/ZivoeTranches.sol#L324)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [IERC20(asset).safeTransferFrom(depositor,IZivoeGlobals_ZivoeTranches(GBL).DAO(),amount)](src/ZivoeTranches.sol#L277)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
	External calls sending eth:
	- [depositSenior(amountSenior,assetSenior)](src/ZivoeTranches.sol#L323)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)
	- [depositJunior(amountJunior,assetJunior)](src/ZivoeTranches.sol#L324)
		- [(success,returndata) = target.call{value: value}(data)](lib/openzeppelin-contracts/contracts/utils/Address.sol#L135)

src/ZivoeTranches.sol#L322-L325


## state-variable-not-initialized
Impact: High
Confidence: High
 - [ ] ID-4
state variable: [ERC20Votes._totalSupplyCheckpoints](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L38) not initialized and not written in contract but be used in contract

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L38


 - [ ] ID-5
state variable: [ZivoeVotes._totalSupplyCheckpoints](src/libraries/ZivoeVotes.sol#L19) not initialized and not written in contract but be used in contract

src/libraries/ZivoeVotes.sol#L19


 - [ ] ID-6
state variable: [ERC20Votes._checkpoints](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L37) not initialized and not written in contract but be used in contract

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L37


 - [ ] ID-7
state variable: [Governor._governanceCall](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L53) not initialized and not written in contract but be used in contract

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L53


 - [ ] ID-8
state variable: [ERC20Permit._nonces](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L25) not initialized and not written in contract but be used in contract

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L25


 - [ ] ID-9
state variable: [GovernorVotesQuorumFraction._quorumNumerator](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L19) not initialized and not written in contract but be used in contract

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L19


 - [ ] ID-10
state variable: [ZivoeVotes._checkpoints](src/libraries/ZivoeVotes.sol#L17) not initialized and not written in contract but be used in contract

src/libraries/ZivoeVotes.sol#L17


## centralized-risk-medium
Impact: Medium
Confidence: High
 - [ ] ID-11
	- [Governor.relay(address,uint256,bytes)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L543-L550)

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L543-L550


## public-mint-burn
Impact: Medium
Confidence: Medium
 - [ ] ID-12
public mint or burn found in [MockStablecoin.mint(address,uint256)](src/misc/MockStablecoin.sol#L42-L44)
src/misc/MockStablecoin.sol#L42-L44


## uninitialized-local
Impact: Medium
Confidence: Medium
 - [ ] ID-13
[OCL_ZVE.pushToLockerMulti(address[],uint256[],bytes[]).preBasis](src/lockers/OCL/OCL_ZVE.sol#L188) is a local variable never initialized

src/lockers/OCL/OCL_ZVE.sol#L188


 - [ ] ID-14
[ERC20Votes._moveVotingPower(address,address,uint256).newWeight_scope_1](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L233) is a local variable never initialized

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L233


## unused-return
Impact: Medium
Confidence: Medium
 - [ ] ID-15
[OCY_Convex_C.claimRewards(bool)](src/lockers/OCY/OCY_Convex_C.sol#L200-L220)have ignores return value in [IBaseRewardPool_OCY_Convex_C(convexRewards).getReward()](src/lockers/OCY/OCY_Convex_C.sol#L201)

src/lockers/OCY/OCY_Convex_C.sol#L200-L220


 - [ ] ID-16
[OCY_Convex_C.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L146-L169)have ignores return value in [IBaseRewardPool_OCY_Convex_C(convexRewards).withdrawAndUnwrap(IERC20(convexRewards).balanceOf(address(this)),false)](src/lockers/OCY/OCY_Convex_C.sol#L152-L154)

src/lockers/OCY/OCY_Convex_C.sol#L146-L169


 - [ ] ID-17
[OCY_Convex_B.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L164-L191)have ignores return value in [IBaseRewardPool_OCY_Convex_B(convexRewards).withdrawAndUnwrap(IERC20(convexRewards).balanceOf(address(this)),false)](src/lockers/OCY/OCY_Convex_B.sol#L170-L172)

src/lockers/OCY/OCY_Convex_B.sol#L164-L191


 - [ ] ID-18
[GovernorVotesQuorumFraction._updateQuorumNumerator(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L99-L118)have ignores return value in [_quorumNumeratorHistory.push(newQuorumNumerator)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L115)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L99-L118


 - [ ] ID-19
[OCY_Convex_A.claimRewards(bool)](src/lockers/OCY/OCY_Convex_A.sol#L246-L266)have ignores return value in [IBaseRewardPool_OCY_Convex_A(convexRewards).getReward()](src/lockers/OCY/OCY_Convex_A.sol#L247)

src/lockers/OCY/OCY_Convex_A.sol#L246-L266


 - [ ] ID-20
[OCY_Convex_A.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L174-L207)have ignores return value in [IBaseRewardPool_OCY_Convex_A(convexRewards).withdrawAndUnwrap(IERC20(convexRewards).balanceOf(address(this)),false)](src/lockers/OCY/OCY_Convex_A.sol#L180-L182)

src/lockers/OCY/OCY_Convex_A.sol#L174-L207


 - [ ] ID-21
[OCY_Convex_B.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L197-L223)have ignores return value in [IBaseRewardPool_OCY_Convex_B(convexRewards).withdrawAndUnwrap(amount,false)](src/lockers/OCY/OCY_Convex_B.sol#L203)

src/lockers/OCY/OCY_Convex_B.sol#L197-L223


 - [ ] ID-22
[Governor._execute(uint256,address[],uint256[],bytes[],bytes32)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L316-L328)have ignores return value in [Address.verifyCallResult(success,returndata,errorMessage)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L326)

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L316-L328


 - [ ] ID-23
[OCY_Convex_A.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L213-L242)have ignores return value in [IBaseRewardPool_OCY_Convex_A(convexRewards).withdrawAndUnwrap(amount,false)](src/lockers/OCY/OCY_Convex_A.sol#L218)

src/lockers/OCY/OCY_Convex_A.sol#L213-L242


 - [ ] ID-24
[OCY_Convex_B.claimRewards(bool)](src/lockers/OCY/OCY_Convex_B.sol#L227-L247)have ignores return value in [IBaseRewardPool_OCY_Convex_B(convexRewards).getReward()](src/lockers/OCY/OCY_Convex_B.sol#L228)

src/lockers/OCY/OCY_Convex_B.sol#L227-L247


 - [ ] ID-25
[OCY_Convex_B.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L112-L159)have ignores return value in [IBooster_OCY_Convex_B(convexDeposit).deposit(convexPoolID,balCurveBasePoolToken,true)](src/lockers/OCY/OCY_Convex_B.sol#L157)

src/lockers/OCY/OCY_Convex_B.sol#L112-L159


 - [ ] ID-26
[OCY_Convex_A.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L117-L169)have ignores return value in [IBasePool_OCY_Convex_A(curveBasePool).add_liquidity(_amounts,_min_mint_amountBP)](src/lockers/OCY/OCY_Convex_A.sol#L131)

src/lockers/OCY/OCY_Convex_A.sol#L117-L169


 - [ ] ID-27
[Governor.relay(address,uint256,bytes)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L543-L550)have ignores return value in [Address.verifyCallResult(success,returndata,Governor: relay reverted without message)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L549)

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L543-L550


 - [ ] ID-28
[OCY_Convex_C.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L110-L141)have ignores return value in [IBooster_OCY_Convex_C(convexDeposit).deposit(convexPoolID,balCurveBasePoolToken,true)](src/lockers/OCY/OCY_Convex_C.sol#L139)

src/lockers/OCY/OCY_Convex_C.sol#L110-L141


 - [ ] ID-29
[OCY_Convex_C.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L175-L196)have ignores return value in [IBaseRewardPool_OCY_Convex_C(convexRewards).withdrawAndUnwrap(amount,false)](src/lockers/OCY/OCY_Convex_C.sol#L181)

src/lockers/OCY/OCY_Convex_C.sol#L175-L196


## void-function
Impact: Medium
Confidence: High
 - [ ] ID-30
function:[ERC20._afterTokenTransfer(address,address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L384-L388)is empty 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L384-L388


 - [ ] ID-31
function:[ERC20._beforeTokenTransfer(address,address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L364-L368)is empty 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L364-L368


 - [ ] ID-32
function:[Presale.fallback()](src/misc/Presale.sol#L190)is empty 

src/misc/Presale.sol#L190


## pess-event-setter
Impact: Low
Confidence: Medium
 - [ ] ID-33
Setter function [OCE_ZVE.constructor(address,address)](src/lockers/OCE/OCE_ZVE.sol#L70-L78) does not emit an event

src/lockers/OCE/OCE_ZVE.sol#L70-L78


 - [ ] ID-34
Setter function [OCY_OUSD.constructor(address,address,address)](src/lockers/OCY/OCY_OUSD.sol#L44-L48) does not emit an event

src/lockers/OCY/OCY_OUSD.sol#L44-L48


 - [ ] ID-35
Setter function [OCR_Modular.pushToLocker(address,uint256,bytes)](src/lockers/OCR/OCR_Modular.sol#L180-L185) does not emit an event

src/lockers/OCR/OCR_Modular.sol#L180-L185


 - [ ] ID-36
Setter function [ZivoeTranches.onlyGovernance()](src/ZivoeTranches.sol#L154-L160) does not emit an event

src/ZivoeTranches.sol#L154-L160


 - [ ] ID-37
Setter function [ZivoeGovernorV2.setVotingPeriod(uint256)](src/ZivoeGovernorV2.sol#L91-L95) does not emit an event

src/ZivoeGovernorV2.sol#L91-L95


 - [ ] ID-38
Setter function [ZivoeTrancheToken.burn(uint256)](src/ZivoeTrancheToken.sol#L72) does not emit an event

src/ZivoeTrancheToken.sol#L72


 - [ ] ID-39
Setter function [OCY_Convex_C.constructor(address,address,address)](src/lockers/OCY/OCY_Convex_C.sol#L74-L78) does not emit an event

src/lockers/OCY/OCY_Convex_C.sol#L74-L78


 - [ ] ID-40
Setter function [Timers.setDeadline(Timers.Timestamp,uint64)](lib/openzeppelin-contracts/contracts/utils/Timers.sol#L18-L20) does not emit an event

lib/openzeppelin-contracts/contracts/utils/Timers.sol#L18-L20


 - [ ] ID-41
Setter function [ZivoeYDL.unlock()](src/ZivoeYDL.sol#L321-L352) does not emit an event

src/ZivoeYDL.sol#L321-L352


 - [ ] ID-42
Setter function [OCC_Modular.isUnderwriter()](src/lockers/OCC/OCC_Modular.sol#L381-L384) does not emit an event

src/lockers/OCC/OCC_Modular.sol#L381-L384


 - [ ] ID-43
Setter function [ZivoeGovernorV2.updateQuorumNumerator(uint256)](src/ZivoeGovernorV2.sol#L107-L110) does not emit an event

src/ZivoeGovernorV2.sol#L107-L110


 - [ ] ID-44
Setter function [OCY_Convex_C.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L175-L196) does not emit an event

src/lockers/OCY/OCY_Convex_C.sol#L175-L196


 - [ ] ID-45
Setter function [OCG_Defaults.onlyGovernance()](src/lockers/OCG/OCG_Defaults.sol#L41-L47) does not emit an event

src/lockers/OCG/OCG_Defaults.sol#L41-L47


 - [ ] ID-46
Setter function [OCY_Convex_A.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L117-L169) does not emit an event

src/lockers/OCY/OCY_Convex_A.sol#L117-L169


 - [ ] ID-47
Setter function [OCY_Convex_C.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L110-L141) does not emit an event

src/lockers/OCY/OCY_Convex_C.sol#L110-L141


 - [ ] ID-48
Setter function [OCC_Modular.constructor(address,address,address,address,address)](src/lockers/OCC/OCC_Modular.sol#L156-L162) does not emit an event

src/lockers/OCC/OCC_Modular.sol#L156-L162


 - [ ] ID-49
Setter function [ZivoeTLC.onlyRoleOrOpenRole(bytes32)](src/libraries/ZivoeTLC.sol#L121-L126) does not emit an event

src/libraries/ZivoeTLC.sol#L121-L126


 - [ ] ID-50
Setter function [OCG_ERC1155.constructor(address)](src/lockers/OCG/OCG_ERC1155.sol#L15-L17) does not emit an event

src/lockers/OCG/OCG_ERC1155.sol#L15-L17


 - [ ] ID-51
Setter function [OCY_Convex_A.constructor(address,address,address)](src/lockers/OCY/OCY_Convex_A.sol#L81-L85) does not emit an event

src/lockers/OCY/OCY_Convex_A.sol#L81-L85


 - [ ] ID-52
Setter function [OCG_Defaults.decreaseDefaults(uint256)](src/lockers/OCG/OCG_Defaults.sol#L57-L59) does not emit an event

src/lockers/OCG/OCG_Defaults.sol#L57-L59


 - [ ] ID-53
Setter function [OCE_ZVE.pushToLocker(address,uint256,bytes)](src/lockers/OCE/OCE_ZVE.sol#L122-L128) does not emit an event

src/lockers/OCE/OCE_ZVE.sol#L122-L128


 - [ ] ID-54
Setter function [ZivoeTranches.switchPause()](src/ZivoeTranches.sol#L328-L334) does not emit an event

src/ZivoeTranches.sol#L328-L334


 - [ ] ID-55
Setter function [OCR_Modular.pullFromLocker(address,bytes)](src/lockers/OCR/OCR_Modular.sol#L190-L196) does not emit an event

src/lockers/OCR/OCR_Modular.sol#L190-L196


 - [ ] ID-56
Setter function [OCY_Convex_B.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L164-L191) does not emit an event

src/lockers/OCY/OCY_Convex_B.sol#L164-L191


 - [ ] ID-57
Setter function [ZivoeRewardsVesting.getRewards()](src/ZivoeRewardsVesting.sol#L484-L486) does not emit an event

src/ZivoeRewardsVesting.sol#L484-L486


 - [ ] ID-58
Setter function [ZivoeTLC._afterCallKeeper(bytes32)](src/libraries/ZivoeTLC.sol#L414-L417) does not emit an event

src/libraries/ZivoeTLC.sol#L414-L417


 - [ ] ID-59
Setter function [OCG_ERC20_FreeClaim.claim(address,uint256)](src/lockers/OCG/OCG_ERC20_FreeClaim.sol#L63-L65) does not emit an event

src/lockers/OCG/OCG_ERC20_FreeClaim.sol#L63-L65


 - [ ] ID-60
Setter function [OCR_Modular.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCR/OCR_Modular.sol#L202-L210) does not emit an event

src/lockers/OCR/OCR_Modular.sol#L202-L210


 - [ ] ID-61
Setter function [ZivoeTrancheToken.mint(address,uint256)](src/ZivoeTrancheToken.sol#L87) does not emit an event

src/ZivoeTrancheToken.sol#L87


 - [ ] ID-62
Setter function [OCG_ERC20.constructor(address)](src/lockers/OCG/OCG_ERC20.sol#L15-L17) does not emit an event

src/lockers/OCG/OCG_ERC20.sol#L15-L17


 - [ ] ID-63
Setter function [OCG_Defaults.constructor(address,address)](src/lockers/OCG/OCG_Defaults.sol#L29-L32) does not emit an event

src/lockers/OCG/OCG_Defaults.sol#L29-L32


 - [ ] ID-64
Setter function [OCY_Convex_B.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L112-L159) does not emit an event

src/lockers/OCY/OCY_Convex_B.sol#L112-L159


 - [ ] ID-65
Setter function [OCG_ERC721.constructor(address)](src/lockers/OCG/OCG_ERC721.sol#L15-L17) does not emit an event

src/lockers/OCG/OCG_ERC721.sol#L15-L17


 - [ ] ID-66
Setter function [ZivoeTrancheToken.isMinterRole()](src/ZivoeTrancheToken.sol#L54-L57) does not emit an event

src/ZivoeTrancheToken.sol#L54-L57


 - [ ] ID-67
Setter function [ZivoeGovernorV2.setProposalThreshold(uint256)](src/ZivoeGovernorV2.sol#L98-L104) does not emit an event

src/ZivoeGovernorV2.sol#L98-L104


 - [ ] ID-68
Setter function [OCY_Convex_B.constructor(address,address,address)](src/lockers/OCY/OCY_Convex_B.sol#L76-L80) does not emit an event

src/lockers/OCY/OCY_Convex_B.sol#L76-L80


 - [ ] ID-69
Setter function [ZivoeGlobals.proposeZVL(address)](src/ZivoeGlobals.sol#L210-L212) does not emit an event

src/ZivoeGlobals.sol#L210-L212


 - [ ] ID-70
Setter function [ZivoeRewardsVesting.updateReward(address)](src/ZivoeRewardsVesting.sol#L197-L208) does not emit an event

src/ZivoeRewardsVesting.sol#L197-L208


 - [ ] ID-71
Setter function [OCY_Convex_A.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L174-L207) does not emit an event

src/lockers/OCY/OCY_Convex_A.sol#L174-L207


 - [ ] ID-72
Setter function [ZivoeITO.depositBoth(uint256,address,uint256,address)](src/ZivoeITO.sol#L305-L309) does not emit an event

src/ZivoeITO.sol#L305-L309


 - [ ] ID-73
Setter function [Timers.setDeadline(Timers.BlockNumber,uint64)](lib/openzeppelin-contracts/contracts/utils/Timers.sol#L50-L52) does not emit an event

lib/openzeppelin-contracts/contracts/utils/Timers.sol#L50-L52


 - [ ] ID-74
Setter function [ZivoeRewards.getRewards()](src/ZivoeRewards.sol#L281-L283) does not emit an event

src/ZivoeRewards.sol#L281-L283


 - [ ] ID-75
Setter function [ZivoeRewardsVesting.onlyZVLOrITO()](src/ZivoeRewardsVesting.sol#L186-L193) does not emit an event

src/ZivoeRewardsVesting.sol#L186-L193


 - [ ] ID-76
Setter function [ZivoeTranches.depositBoth(uint256,address,uint256,address)](src/ZivoeTranches.sol#L322-L325) does not emit an event

src/ZivoeTranches.sol#L322-L325


 - [ ] ID-77
Setter function [OCY_Convex_B.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L197-L223) does not emit an event

src/lockers/OCY/OCY_Convex_B.sol#L197-L223


 - [ ] ID-78
Setter function [ZivoeTLC._schedule(bytes32,uint256)](src/libraries/ZivoeTLC.sol#L275-L279) does not emit an event

src/libraries/ZivoeTLC.sol#L275-L279


 - [ ] ID-79
Setter function [OCG_Defaults.increaseDefaults(uint256)](src/lockers/OCG/OCG_Defaults.sol#L64-L66) does not emit an event

src/lockers/OCG/OCG_Defaults.sol#L64-L66


 - [ ] ID-80
Setter function [OCT_YDL.constructor(address,address)](src/lockers/OCT/OCT_YDL.sol#L47-L50) does not emit an event

src/lockers/OCT/OCT_YDL.sol#L47-L50


 - [ ] ID-81
Setter function [OCR_Modular.constructor(address,address,address,uint16)](src/lockers/OCR/OCR_Modular.sol#L87-L94) does not emit an event

src/lockers/OCR/OCR_Modular.sol#L87-L94


 - [ ] ID-82
Setter function [OCE_ZVE.forwardEmissions()](src/lockers/OCE/OCE_ZVE.sol#L131-L135) does not emit an event

src/lockers/OCE/OCE_ZVE.sol#L131-L135


 - [ ] ID-83
Setter function [OCY_Convex_A.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L213-L242) does not emit an event

src/lockers/OCY/OCY_Convex_A.sol#L213-L242


 - [ ] ID-84
Setter function [OCY_Convex_C.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L146-L169) does not emit an event

src/lockers/OCY/OCY_Convex_C.sol#L146-L169


 - [ ] ID-85
Setter function [ZivoeGlobals.onlyZVL()](src/ZivoeGlobals.sol#L112-L115) does not emit an event

src/ZivoeGlobals.sol#L112-L115


 - [ ] ID-86
Setter function [ZivoeRewardsVesting.fullWithdraw()](src/ZivoeRewardsVesting.sol#L370-L373) does not emit an event

src/ZivoeRewardsVesting.sol#L370-L373


 - [ ] ID-87
Setter function [OCT_DAO.constructor(address,address)](src/lockers/OCT/OCT_DAO.sol#L41-L44) does not emit an event

src/lockers/OCT/OCT_DAO.sol#L41-L44


 - [ ] ID-88
Setter function [ZivoeTranches.unlock()](src/ZivoeTranches.sol#L404-L410) does not emit an event

src/ZivoeTranches.sol#L404-L410


 - [ ] ID-89
Setter function [ZivoeGovernorV2.setVotingDelay(uint256)](src/ZivoeGovernorV2.sol#L84-L87) does not emit an event

src/ZivoeGovernorV2.sol#L84-L87


 - [ ] ID-90
Setter function [ZivoeRewards.updateReward(address)](src/ZivoeRewards.sol#L123-L134) does not emit an event

src/ZivoeRewards.sol#L123-L134


 - [ ] ID-91
Setter function [OCL_ZVE.constructor(address,address,address,address,address,address)](src/lockers/OCL/OCL_ZVE.sol#L111-L118) does not emit an event

src/lockers/OCL/OCL_ZVE.sol#L111-L118


 - [ ] ID-92
Setter function [ZivoeToken.burn(uint256)](src/ZivoeToken.sol#L35) does not emit an event

src/ZivoeToken.sol#L35


 - [ ] ID-93
Setter function [OCL_ZVE.forwardYield()](src/lockers/OCL/OCL_ZVE.sol#L287-L305) does not emit an event

src/lockers/OCL/OCL_ZVE.sol#L287-L305


 - [ ] ID-94
Setter function [ZivoeTranches.pushToLocker(address,uint256,bytes)](src/ZivoeTranches.sol#L180-L186) does not emit an event

src/ZivoeTranches.sol#L180-L186


 - [ ] ID-95
Setter function [OCT_ZVL.constructor(address,address)](src/lockers/OCT/OCT_ZVL.sol#L40-L43) does not emit an event

src/lockers/OCT/OCT_ZVL.sol#L40-L43


 - [ ] ID-96
Setter function [ZivoeTLC._afterCall(bytes32)](src/libraries/ZivoeTLC.sol#L398-L401) does not emit an event

src/libraries/ZivoeTLC.sol#L398-L401


 - [ ] ID-97
Setter function [OCG_ERC20_FreeClaim.constructor(address)](src/lockers/OCG/OCG_ERC20_FreeClaim.sol#L20-L22) does not emit an event

src/lockers/OCG/OCG_ERC20_FreeClaim.sol#L20-L22


 - [ ] ID-98
Setter function [ZivoeRewards.fullWithdraw()](src/ZivoeRewards.sol#L246-L249) does not emit an event

src/ZivoeRewards.sol#L246-L249


## input-validation
Impact: Low
Confidence: Medium
 - [ ] ID-99
value assignment lack of validation	[ZivoeGTC.queue(address[],uint256[],bytes[],bytes32)](src/libraries/ZivoeGTC.sol#L91-L108)[_timelockIds[proposalId] = _timelock.hashOperationBatch(targets,values,calldatas,0,descriptionHash)](src/libraries/ZivoeGTC.sol#L102)

src/libraries/ZivoeGTC.sol#L91-L108


## missing-zero-check
Impact: Low
Confidence: Medium
 - [ ] ID-100
variable lacks a zero-check on 		- [ZivoeLocker.pushToLockerERC1155(address,uint256[],uint256[],bytes)](src/ZivoeLocker.sol#L180-L185)

src/ZivoeLocker.sol#L180-L185


 - [ ] ID-101
variable lacks a zero-check on 		- [ZivoeRewardsVesting.constructor(address,address)](src/ZivoeRewardsVesting.sol#L107-L111)

src/ZivoeRewardsVesting.sol#L107-L111


 - [ ] ID-102
variable lacks a zero-check on 		- [ZivoeITO.depositJunior(uint256,address)](src/ZivoeITO.sol#L248-L271)

src/ZivoeITO.sol#L248-L271


 - [ ] ID-103
variable lacks a zero-check on 		- [ZivoeTLC.constructor(uint256,address[],address[],address)](src/libraries/ZivoeTLC.sol#L82-L113)

src/libraries/ZivoeTLC.sol#L82-L113


 - [ ] ID-104
variable lacks a zero-check on 		- [OCG_Defaults.constructor(address,address)](src/lockers/OCG/OCG_Defaults.sol#L29-L32)

src/lockers/OCG/OCG_Defaults.sol#L29-L32


 - [ ] ID-105
variable lacks a zero-check on 		- [ZivoeDAO.pushERC721(address,address,uint256,bytes)](src/ZivoeDAO.sol#L344-L357)

src/ZivoeDAO.sol#L344-L357


 - [ ] ID-106
variable lacks a zero-check on 		- [ZivoeDAO.pullPartial(address,address,uint256,bytes)](src/ZivoeDAO.sol#L268-L274)

src/ZivoeDAO.sol#L268-L274


 - [ ] ID-107
variable lacks a zero-check on 		- [OCY_Convex_A.constructor(address,address,address)](src/lockers/OCY/OCY_Convex_A.sol#L81-L85)

src/lockers/OCY/OCY_Convex_A.sol#L81-L85


 - [ ] ID-108
variable lacks a zero-check on 		- [ZivoeTranches.depositSenior(uint256,address)](src/ZivoeTranches.sol#L295-L315)

src/ZivoeTranches.sol#L295-L315


 - [ ] ID-109
variable lacks a zero-check on 		- [ZivoeYDL.constructor(address,address)](src/ZivoeYDL.sol#L129-L133)

src/ZivoeYDL.sol#L129-L133


 - [ ] ID-110
variable lacks a zero-check on 		- [OCT_ZVL.constructor(address,address)](src/lockers/OCT/OCT_ZVL.sol#L40-L43)

src/lockers/OCT/OCT_ZVL.sol#L40-L43


 - [ ] ID-111
variable lacks a zero-check on 		- [ZivoeRewards.constructor(address,address)](src/ZivoeRewards.sol#L72-L75)

src/ZivoeRewards.sol#L72-L75


 - [ ] ID-112
variable lacks a zero-check on 		- [ZivoeDAO.pullERC1155(address,address,uint256[],uint256[],bytes)](src/ZivoeDAO.sol#L455-L465)

src/ZivoeDAO.sol#L455-L465


 - [ ] ID-113
variable lacks a zero-check on 		- [ZivoeLocker.pullFromLockerERC1155(address,uint256[],uint256[],bytes)](src/ZivoeLocker.sol#L192-L197)

src/ZivoeLocker.sol#L192-L197


 - [ ] ID-114
variable lacks a zero-check on 		- [ZivoeITO.depositSenior(uint256,address)](src/ZivoeITO.sol#L277-L298)

src/ZivoeITO.sol#L277-L298


 - [ ] ID-115
variable lacks a zero-check on 		- [OCT_DAO.constructor(address,address)](src/lockers/OCT/OCT_DAO.sol#L41-L44)

src/lockers/OCT/OCT_DAO.sol#L41-L44


 - [ ] ID-116
variable lacks a zero-check on 		- [OCY_OUSD.constructor(address,address,address)](src/lockers/OCY/OCY_OUSD.sol#L44-L48)

src/lockers/OCY/OCY_OUSD.sol#L44-L48


 - [ ] ID-117
variable lacks a zero-check on 		- [ZivoeGlobals.updateIsKeeper(address,bool)](src/ZivoeGlobals.sol#L226-L229)

src/ZivoeGlobals.sol#L226-L229


 - [ ] ID-118
variable lacks a zero-check on 		- [ZivoeTranches.constructor(address)](src/ZivoeTranches.sol#L96)

src/ZivoeTranches.sol#L96


 - [ ] ID-119
variable lacks a zero-check on 		- [ZivoeGlobals.updateIsLocker(address,bool)](src/ZivoeGlobals.sol#L235-L238)

src/ZivoeGlobals.sol#L235-L238


 - [ ] ID-120
variable lacks a zero-check on 		- [OCT_DAO.convertAndForward(address,address,bytes)](src/lockers/OCT/OCT_DAO.sol#L87-L99)

src/lockers/OCT/OCT_DAO.sol#L87-L99


 - [ ] ID-121
variable lacks a zero-check on 		- [ZivoeITO.constructor(address,address[])](src/ZivoeITO.sol#L123-L126)

src/ZivoeITO.sol#L123-L126


 - [ ] ID-122
variable lacks a zero-check on 		- [ZivoeTranches.depositJunior(uint256,address)](src/ZivoeTranches.sol#L268-L289)

src/ZivoeTranches.sol#L268-L289


 - [ ] ID-123
variable lacks a zero-check on 		- [OCY_Convex_B.constructor(address,address,address)](src/lockers/OCY/OCY_Convex_B.sol#L76-L80)

src/lockers/OCY/OCY_Convex_B.sol#L76-L80


 - [ ] ID-124
variable lacks a zero-check on 		- [ZivoeTrancheToken.changeMinterRole(address,bool)](src/ZivoeTrancheToken.sol#L78-L81)

src/ZivoeTrancheToken.sol#L78-L81


 - [ ] ID-125
variable lacks a zero-check on 		- [ZivoeLocker.pullFromLocker(address,bytes)](src/ZivoeLocker.sol#L81-L84)

src/ZivoeLocker.sol#L81-L84


 - [ ] ID-126
variable lacks a zero-check on 		- [ZivoeGovernorV2.constructor(IVotes,ZivoeTLC,address)](src/ZivoeGovernorV2.sol#L62-L64)

src/ZivoeGovernorV2.sol#L62-L64


 - [ ] ID-127
variable lacks a zero-check on 		- [ZivoeDAO.pushERC1155(address,address,uint256[],uint256[],bytes)](src/ZivoeDAO.sol#L429-L446)

src/ZivoeDAO.sol#L429-L446


 - [ ] ID-128
variable lacks a zero-check on 		- [ZivoeDAO.constructor(address)](src/ZivoeDAO.sol#L169)

src/ZivoeDAO.sol#L169


 - [ ] ID-129
variable lacks a zero-check on 		- [ZivoeDAO.push(address,address,uint256,bytes)](src/ZivoeDAO.sol#L239-L247)

src/ZivoeDAO.sol#L239-L247


 - [ ] ID-130
variable lacks a zero-check on 		- [Presale.constructor(address[],address,address)](src/misc/Presale.sol#L51-L65)

src/misc/Presale.sol#L51-L65


 - [ ] ID-131
variable lacks a zero-check on 		- [OCT_YDL.convertAndForward(address,bytes)](src/lockers/OCT/OCT_YDL.sol#L97-L110)

src/lockers/OCT/OCT_YDL.sol#L97-L110


 - [ ] ID-132
variable lacks a zero-check on 		- [ZivoeDAO.pullERC721(address,address,uint256,bytes)](src/ZivoeDAO.sol#L395-L401)

src/ZivoeDAO.sol#L395-L401


 - [ ] ID-133
variable lacks a zero-check on 		- [ZivoeLocker.pushToLockerERC721(address,uint256,bytes)](src/ZivoeLocker.sol#L135-L138)

src/ZivoeLocker.sol#L135-L138


 - [ ] ID-134
variable lacks a zero-check on 		- [OCT_YDL.constructor(address,address)](src/lockers/OCT/OCT_YDL.sol#L47-L50)

src/lockers/OCT/OCT_YDL.sol#L47-L50


 - [ ] ID-135
variable lacks a zero-check on 		- [OCC_Modular.constructor(address,address,address,address,address)](src/lockers/OCC/OCC_Modular.sol#L156-L162)

src/lockers/OCC/OCC_Modular.sol#L156-L162


 - [ ] ID-136
variable lacks a zero-check on 		- [ZivoeGlobals.updateYDL(address)](src/ZivoeGlobals.sol#L252-L255)

src/ZivoeGlobals.sol#L252-L255


 - [ ] ID-137
variable lacks a zero-check on 		- [ZivoeDAO.pull(address,address,bytes)](src/ZivoeDAO.sol#L255-L259)

src/ZivoeDAO.sol#L255-L259


 - [ ] ID-138
variable lacks a zero-check on 		- [ZivoeLocker.pullFromLockerERC721(address,uint256,bytes)](src/ZivoeLocker.sol#L144-L147)

src/ZivoeLocker.sol#L144-L147


 - [ ] ID-139
variable lacks a zero-check on 		- [ZivoeGlobals.proposeZVL(address)](src/ZivoeGlobals.sol#L210-L212)

src/ZivoeGlobals.sol#L210-L212


 - [ ] ID-140
variable lacks a zero-check on 		- [OCY_Convex_C.constructor(address,address,address)](src/lockers/OCY/OCY_Convex_C.sol#L74-L78)

src/lockers/OCY/OCY_Convex_C.sol#L74-L78


 - [ ] ID-141
variable lacks a zero-check on 		- [OCL_ZVE.constructor(address,address,address,address,address,address)](src/lockers/OCL/OCL_ZVE.sol#L111-L118)

src/lockers/OCL/OCL_ZVE.sol#L111-L118


 - [ ] ID-142
variable lacks a zero-check on 		- [ZivoeGlobals.updateStablecoinWhitelist(address,bool)](src/ZivoeGlobals.sol#L244-L247)

src/ZivoeGlobals.sol#L244-L247


 - [ ] ID-143
variable lacks a zero-check on 		- [Governor.relay(address,uint256,bytes)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L543-L550)

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L543-L550


 - [ ] ID-144
variable lacks a zero-check on 		- [OCE_ZVE.constructor(address,address)](src/lockers/OCE/OCE_ZVE.sol#L70-L78)

src/lockers/OCE/OCE_ZVE.sol#L70-L78


 - [ ] ID-145
variable lacks a zero-check on 		- [OCR_Modular.constructor(address,address,address,uint16)](src/lockers/OCR/OCR_Modular.sol#L87-L94)

src/lockers/OCR/OCR_Modular.sol#L87-L94


## variable-scope
Impact: Low
Confidence: High
 - [ ] ID-146
Variable '[ERC20Votes._moveVotingPower(address,address,uint256).oldWeight](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L228)' in [ERC20Votes._moveVotingPower(address,address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L221-L237) potentially used before declaration: [(oldWeight,newWeight) = _writeCheckpoint(_checkpoints[dst],_add,amount)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L233)

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L228


## centralized-risk-informational
Impact: Informational
Confidence: High
 - [ ] ID-147
	- [OCY_OUSD.updateOCTYDL(address)](src/lockers/OCY/OCY_OUSD.sol#L134-L142)

src/lockers/OCY/OCY_OUSD.sol#L134-L142


 - [ ] ID-148
	- [OCC_Modular.unapproveRefinance(uint256)](src/lockers/OCC/OCC_Modular.sol#L967-L970)

src/lockers/OCC/OCC_Modular.sol#L967-L970


 - [ ] ID-149
	- [OCL_ZVE.pushToLockerMulti(address[],uint256[],bytes[])](src/lockers/OCL/OCL_ZVE.sol#L172-L215)

src/lockers/OCL/OCL_ZVE.sol#L172-L215


 - [ ] ID-150
	- [ZivoeYDL.updateDistributedAsset(address)](src/ZivoeYDL.sol#L356-L371)

src/ZivoeYDL.sol#L356-L371


 - [ ] ID-151
	- [ZivoeYDL.unlock()](src/ZivoeYDL.sol#L321-L352)

src/ZivoeYDL.sol#L321-L352


 - [ ] ID-152
	- [ZivoeGlobals.updateYDL(address)](src/ZivoeGlobals.sol#L252-L255)

src/ZivoeGlobals.sol#L252-L255


 - [ ] ID-153
	- [OCC_Modular.applyCombine(uint256)](src/lockers/OCC/OCC_Modular.sol#L742-L808)

src/lockers/OCC/OCC_Modular.sol#L742-L808


 - [ ] ID-154
	- [OCC_Modular.unapproveExtension(uint256)](src/lockers/OCC/OCC_Modular.sol#L960-L963)

src/lockers/OCC/OCC_Modular.sol#L960-L963


 - [ ] ID-155
	- [ZivoeTranches.updateMaxTrancheRatio(uint256)](src/ZivoeTranches.sol#L360-L364)

src/ZivoeTranches.sol#L360-L364


 - [ ] ID-156
	- [OCC_Modular.unapproveConversionToBullet(uint256)](src/lockers/OCC/OCC_Modular.sol#L953-L956)

src/lockers/OCC/OCC_Modular.sol#L953-L956


 - [ ] ID-157
	- [OCY_OUSD.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_OUSD.sol#L99-L107)

src/lockers/OCY/OCY_OUSD.sol#L99-L107


 - [ ] ID-158
	- [OCL_ZVE.updateOCTYDL(address)](src/lockers/OCL/OCL_ZVE.sol#L361-L369)

src/lockers/OCL/OCL_ZVE.sol#L361-L369


 - [ ] ID-159
	- [ZivoeTranches.unlock()](src/ZivoeTranches.sol#L404-L410)

src/ZivoeTranches.sol#L404-L410


 - [ ] ID-160
	- [OCC_Modular.applyConversionToBullet(uint256)](src/lockers/OCC/OCC_Modular.sol#L828-L840)

src/lockers/OCC/OCC_Modular.sol#L828-L840


 - [ ] ID-161
	- [ZivoeGlobals.decreaseDefaults(uint256)](src/ZivoeGlobals.sol#L157-L162)

src/ZivoeGlobals.sol#L157-L162


 - [ ] ID-162
	- [ZivoeGlobals.proposeZVL(address)](src/ZivoeGlobals.sol#L210-L212)

src/ZivoeGlobals.sol#L210-L212


 - [ ] ID-163
	- [OCC_Modular.processPayment(uint256)](src/lockers/OCC/OCC_Modular.sol#L636-L677)

src/lockers/OCC/OCC_Modular.sol#L636-L677


 - [ ] ID-164
	- [ZivoeYDL.updateRecipients(address[],uint256[],bool)](src/ZivoeYDL.sol#L392-L416)

src/ZivoeYDL.sol#L392-L416


 - [ ] ID-165
	- [OCC_Modular.callLoan(uint256)](src/lockers/OCC/OCC_Modular.sol#L492-L518)

src/lockers/OCC/OCC_Modular.sol#L492-L518


 - [ ] ID-166
	- [ZivoeYDL.updateTargetAPYBIPS(uint256)](src/ZivoeYDL.sol#L420-L424)

src/ZivoeYDL.sol#L420-L424


 - [ ] ID-167
	- [ZivoeRewardsVesting.revokeVestingSchedule(address)](src/ZivoeRewardsVesting.sol#L429-L467)

src/ZivoeRewardsVesting.sol#L429-L467


 - [ ] ID-168
	- [OCC_Modular.markDefault(uint256)](src/lockers/OCC/OCC_Modular.sol#L606-L618)

src/lockers/OCC/OCC_Modular.sol#L606-L618


 - [ ] ID-169
	- [OwnableLocked.transferOwnershipAndLock(address)](src/libraries/OwnableLocked.sol#L40-L44)

src/libraries/OwnableLocked.sol#L40-L44


 - [ ] ID-170
	- [OCY_OUSD.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_OUSD.sol#L89-L94)

src/lockers/OCY/OCY_OUSD.sol#L89-L94


 - [ ] ID-171
	- [OCC_Modular.approveCombine(uint256[],uint256,uint256,uint256,uint256,int8)](src/lockers/OCC/OCC_Modular.sol#L877-L905)

src/lockers/OCC/OCC_Modular.sol#L877-L905


 - [ ] ID-172
	- [OCC_Modular.applyExtension(uint256)](src/lockers/OCC/OCC_Modular.sol#L844-L855)

src/lockers/OCC/OCC_Modular.sol#L844-L855


 - [ ] ID-173
	- [OCR_Modular.destroyRequest(uint256)](src/lockers/OCR/OCR_Modular.sol#L232-L251)

src/lockers/OCR/OCR_Modular.sol#L232-L251


 - [ ] ID-174
	- [OCE_ZVE.updateExponentialDecayPerSecond(uint256)](src/lockers/OCE/OCE_ZVE.sol#L183-L194)

src/lockers/OCE/OCE_ZVE.sol#L183-L194


 - [ ] ID-175
	- [ZivoeYDL.updateTargetRatioBIPS(uint256)](src/ZivoeYDL.sol#L428-L432)

src/ZivoeYDL.sol#L428-L432


 - [ ] ID-176
	- [OCC_Modular.updateOCTYDL(address)](src/lockers/OCC/OCC_Modular.sol#L727-L732)

src/lockers/OCC/OCC_Modular.sol#L727-L732


 - [ ] ID-177
	- [ZivoeGlobals.updateStablecoinWhitelist(address,bool)](src/ZivoeGlobals.sol#L244-L247)

src/ZivoeGlobals.sol#L244-L247


 - [ ] ID-178
	- [OCC_Modular.createOffer(address,uint256,uint256,uint256,uint256,uint256,uint256,int8)](src/lockers/OCC/OCC_Modular.sol#L540-L570)

src/lockers/OCC/OCC_Modular.sol#L540-L570


 - [ ] ID-179
	- [ZivoeITO.migrateDeposits()](src/ZivoeITO.sol#L313-L335)

src/ZivoeITO.sol#L313-L335


 - [ ] ID-180
	- [ZivoeITO.commence()](src/ZivoeITO.sol#L339-L347)

src/ZivoeITO.sol#L339-L347


 - [ ] ID-181
	- [ZivoeTranches.updateMaxZVEPerJTTMint(uint256)](src/ZivoeTranches.sol#L368-L373)

src/ZivoeTranches.sol#L368-L373


 - [ ] ID-182
	- [OCC_Modular.approveConversionToBullet(uint256)](src/lockers/OCC/OCC_Modular.sol#L916-L919)

src/lockers/OCC/OCC_Modular.sol#L916-L919


 - [ ] ID-183
	- [ZivoeGlobals.updateIsKeeper(address,bool)](src/ZivoeGlobals.sol#L226-L229)

src/ZivoeGlobals.sol#L226-L229


 - [ ] ID-184
	- [ZivoeGlobals.increaseDefaults(uint256)](src/ZivoeGlobals.sol#L168-L173)

src/ZivoeGlobals.sol#L168-L173


 - [ ] ID-185
	- [OCC_Modular.acceptOffer(uint256)](src/lockers/OCC/OCC_Modular.sol#L461-L487)

src/lockers/OCC/OCC_Modular.sol#L461-L487


 - [ ] ID-186
	- [OCC_Modular.unapproveCombine(uint256)](src/lockers/OCC/OCC_Modular.sol#L939-L942)

src/lockers/OCC/OCC_Modular.sol#L939-L942


 - [ ] ID-187
	- [OCY_OUSD.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_OUSD.sol#L113-L123)

src/lockers/OCY/OCY_OUSD.sol#L113-L123


 - [ ] ID-188
	- [OCC_Modular.applyRefinance(uint256)](src/lockers/OCC/OCC_Modular.sol#L859-L869)

src/lockers/OCC/OCC_Modular.sol#L859-L869


 - [ ] ID-189
	- [OCC_Modular.markRepaid(uint256)](src/lockers/OCC/OCC_Modular.sol#L622-L630)

src/lockers/OCC/OCC_Modular.sol#L622-L630


 - [ ] ID-190
	- [OCC_Modular.approveConversionToAmortization(uint256)](src/lockers/OCC/OCC_Modular.sol#L909-L912)

src/lockers/OCC/OCC_Modular.sol#L909-L912


 - [ ] ID-191
	- [OCC_Modular.unapproveConversionToAmortization(uint256)](src/lockers/OCC/OCC_Modular.sol#L946-L949)

src/lockers/OCC/OCC_Modular.sol#L946-L949


 - [ ] ID-192
	- [OCL_ZVE.forwardYield()](src/lockers/OCL/OCL_ZVE.sol#L287-L305)

src/lockers/OCL/OCL_ZVE.sol#L287-L305


 - [ ] ID-193
	- [ZivoeGlobals.acceptZVL()](src/ZivoeGlobals.sol#L215-L220)

src/ZivoeGlobals.sol#L215-L220


 - [ ] ID-194
	- [OCL_ZVE.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCL/OCL_ZVE.sol#L253-L284)

src/lockers/OCL/OCL_ZVE.sol#L253-L284


 - [ ] ID-195
	- [ZivoeRewardsVesting.addReward(address,uint256)](src/ZivoeRewardsVesting.sol#L328-L347)

src/ZivoeRewardsVesting.sol#L328-L347


 - [ ] ID-196
	- [ZivoeTrancheToken.changeMinterRole(address,bool)](src/ZivoeTrancheToken.sol#L78-L81)

src/ZivoeTrancheToken.sol#L78-L81


 - [ ] ID-197
	- [ZivoeTLC.cancel(bytes32)](src/libraries/ZivoeTLC.sol#L288-L293)

src/libraries/ZivoeTLC.sol#L288-L293


 - [ ] ID-198
	- [OCC_Modular.cancelOffer(uint256)](src/lockers/OCC/OCC_Modular.sol#L522-L529)

src/lockers/OCC/OCC_Modular.sol#L522-L529


 - [ ] ID-199
	- [ZivoeRewards.addReward(address,uint256)](src/ZivoeRewards.sol#L208-L223)

src/ZivoeRewards.sol#L208-L223


 - [ ] ID-200
	- [ZivoeTranches.switchPause()](src/ZivoeTranches.sol#L328-L334)

src/ZivoeTranches.sol#L328-L334


 - [ ] ID-201
	- [OCC_Modular.applyConversionToAmortization(uint256)](src/lockers/OCC/OCC_Modular.sol#L812-L824)

src/lockers/OCC/OCC_Modular.sol#L812-L824


 - [ ] ID-202
	- [OCY_Convex_B.updateOCTYDL(address)](src/lockers/OCY/OCY_Convex_B.sol#L252-L260)

src/lockers/OCY/OCY_Convex_B.sol#L252-L260


 - [ ] ID-203
	- [OCC_Modular.approveExtension(uint256,uint256)](src/lockers/OCC/OCC_Modular.sol#L924-L927)

src/lockers/OCC/OCC_Modular.sol#L924-L927


 - [ ] ID-204
	- [OCL_ZVE.pullFromLocker(address,bytes)](src/lockers/OCL/OCL_ZVE.sol#L220-L247)

src/lockers/OCL/OCL_ZVE.sol#L220-L247


 - [ ] ID-205
	- [ZivoeYDL.updateProtocolEarningsRateBIPS(uint256)](src/ZivoeYDL.sol#L375-L386)

src/ZivoeYDL.sol#L375-L386


 - [ ] ID-206
	- [ZivoeITO.depositSenior(uint256,address)](src/ZivoeITO.sol#L277-L298)

src/ZivoeITO.sol#L277-L298


 - [ ] ID-207
	- [ZivoeITO.depositJunior(uint256,address)](src/ZivoeITO.sol#L248-L271)

src/ZivoeITO.sol#L248-L271


 - [ ] ID-208
	- [ZivoeTranches.updateMinZVEPerJTTMint(uint256)](src/ZivoeTranches.sol#L377-L381)

src/ZivoeTranches.sol#L377-L381


 - [ ] ID-209
	- [OCR_Modular.updateRedemptionsFeeBIPS(uint256)](src/lockers/OCR/OCR_Modular.sol#L340-L350)

src/lockers/OCR/OCR_Modular.sol#L340-L350


 - [ ] ID-210
	- [OCC_Modular.approveRefinance(uint256,uint256)](src/lockers/OCC/OCC_Modular.sol#L932-L935)

src/lockers/OCC/OCC_Modular.sol#L932-L935


 - [ ] ID-211
	- [ZivoeTranches.updateLowerRatioIncentiveBIPS(uint256)](src/ZivoeTranches.sol#L343-L354)

src/ZivoeTranches.sol#L343-L354


 - [ ] ID-212
	- [OCE_ZVE.updateDistributionRatioBIPS(uint256[3])](src/lockers/OCE/OCE_ZVE.sol#L163-L177)

src/lockers/OCE/OCE_ZVE.sol#L163-L177


 - [ ] ID-213
	- [OCL_ZVE.updateCompoundingRateBIPS(uint256)](src/lockers/OCL/OCL_ZVE.sol#L347-L356)

src/lockers/OCL/OCL_ZVE.sol#L347-L356


 - [ ] ID-214
	- [ZivoeGlobals.updateIsLocker(address,bool)](src/ZivoeGlobals.sol#L235-L238)

src/ZivoeGlobals.sol#L235-L238


 - [ ] ID-215
	- [ZivoeRewardsVesting.createVestingSchedule(address,uint256,uint256,uint256,bool)](src/ZivoeRewardsVesting.sol#L381-L425)

src/ZivoeRewardsVesting.sol#L381-L425


 - [ ] ID-216
	- [OCY_Convex_A.updateOCTYDL(address)](src/lockers/OCY/OCY_Convex_A.sol#L271-L279)

src/lockers/OCY/OCY_Convex_A.sol#L271-L279


 - [ ] ID-217
	- [OCY_Convex_C.updateOCTYDL(address)](src/lockers/OCY/OCY_Convex_C.sol#L225-L233)

src/lockers/OCY/OCY_Convex_C.sol#L225-L233


 - [ ] ID-218
	- [ZivoeGlobals.initializeGlobals(address[],address[])](src/ZivoeGlobals.sol#L179-L205)

src/ZivoeGlobals.sol#L179-L205


 - [ ] ID-219
	- [ZivoeTranches.updateUpperRatioIncentiveBIPS(uint256)](src/ZivoeTranches.sol#L390-L401)

src/ZivoeTranches.sol#L390-L401


## centralized-risk-other
Impact: Informational
Confidence: High
 - [ ] ID-220
	- [ZivoeLocker.pushToLocker(address,uint256,bytes)](src/ZivoeLocker.sol#L73-L76)

src/ZivoeLocker.sol#L73-L76


 - [ ] ID-221
	- [ZivoeDAO.pushMultiERC721(address,address[],uint256[],bytes[])](src/ZivoeDAO.sol#L365-L387)

src/ZivoeDAO.sol#L365-L387


 - [ ] ID-222
	- [ZivoeLocker.pushToLockerMulti(address[],uint256[],bytes[])](src/ZivoeLocker.sol#L99-L106)

src/ZivoeLocker.sol#L99-L106


 - [ ] ID-223
	- [Ownable.renounceOwnership()](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L61-L63)

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L61-L63


 - [ ] ID-224
	- [OCY_Convex_B.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L164-L191)

src/lockers/OCY/OCY_Convex_B.sol#L164-L191


 - [ ] ID-225
	- [OCR_Modular.pullFromLocker(address,bytes)](src/lockers/OCR/OCR_Modular.sol#L190-L196)

src/lockers/OCR/OCR_Modular.sol#L190-L196


 - [ ] ID-226
	- [ZivoeGovernorV2.setVotingPeriod(uint256)](src/ZivoeGovernorV2.sol#L91-L95)

src/ZivoeGovernorV2.sol#L91-L95


 - [ ] ID-227
	- [Ownable.transferOwnership(address)](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L69-L72)

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L69-L72


 - [ ] ID-228
	- [ZivoeTLC.execute(address,uint256,bytes,bytes32,bytes32)](src/libraries/ZivoeTLC.sol#L307-L328)

src/libraries/ZivoeTLC.sol#L307-L328


 - [ ] ID-229
	- [GovernorSettings.setVotingPeriod(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L70-L72)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L70-L72


 - [ ] ID-230
	- [OCY_Convex_C.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L110-L141)

src/lockers/OCY/OCY_Convex_C.sol#L110-L141


 - [ ] ID-231
	- [ZivoeDAO.pullMultiPartial(address,address[],uint256[],bytes[])](src/ZivoeDAO.sol#L325-L336)

src/ZivoeDAO.sol#L325-L336


 - [ ] ID-232
	- [ZivoeLocker.pullFromLocker(address,bytes)](src/ZivoeLocker.sol#L81-L84)

src/ZivoeLocker.sol#L81-L84


 - [ ] ID-233
	- [OCY_Convex_A.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L117-L169)

src/lockers/OCY/OCY_Convex_A.sol#L117-L169


 - [ ] ID-234
	- [AccessControl.grantRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L144-L146)

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L144-L146


 - [ ] ID-235
	- [ZivoeDAO.pullPartial(address,address,uint256,bytes)](src/ZivoeDAO.sol#L268-L274)

src/ZivoeDAO.sol#L268-L274


 - [ ] ID-236
	- [ZivoeDAO.push(address,address,uint256,bytes)](src/ZivoeDAO.sol#L239-L247)

src/ZivoeDAO.sol#L239-L247


 - [ ] ID-237
	- [ZivoeLocker.pullFromLockerMultiERC721(address[],uint256[],bytes[])](src/ZivoeLocker.sol#L166-L173)

src/ZivoeLocker.sol#L166-L173


 - [ ] ID-238
	- [OCY_Convex_A.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L213-L242)

src/lockers/OCY/OCY_Convex_A.sol#L213-L242


 - [ ] ID-239
	- [ZivoeDAO.pullMultiERC721(address,address[],uint256[],bytes[])](src/ZivoeDAO.sol#L409-L420)

src/ZivoeDAO.sol#L409-L420


 - [ ] ID-240
	- [ZivoeLocker.pushToLockerMultiERC721(address[],uint256[],bytes[])](src/ZivoeLocker.sol#L153-L160)

src/ZivoeLocker.sol#L153-L160


 - [ ] ID-241
	- [ZivoeTLC.schedule(address,uint256,bytes,bytes32,bytes32,uint256)](src/libraries/ZivoeTLC.sol#L232-L243)

src/libraries/ZivoeTLC.sol#L232-L243


 - [ ] ID-242
	- [ZivoeLocker.pullFromLockerMultiPartial(address[],uint256[],bytes[])](src/ZivoeLocker.sol#L122-L129)

src/ZivoeLocker.sol#L122-L129


 - [ ] ID-243
	- [ZivoeDAO.pullERC721(address,address,uint256,bytes)](src/ZivoeDAO.sol#L395-L401)

src/ZivoeDAO.sol#L395-L401


 - [ ] ID-244
	- [ZivoeLocker.pullFromLockerMulti(address[],bytes[])](src/ZivoeLocker.sol#L111-L116)

src/ZivoeLocker.sol#L111-L116


 - [ ] ID-245
	- [OwnableLocked.transferOwnership(address)](src/libraries/OwnableLocked.sol#L31-L34)

src/libraries/OwnableLocked.sol#L31-L34


 - [ ] ID-246
	- [GovernorSettings.setVotingDelay(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L61-L63)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L61-L63


 - [ ] ID-247
	- [GovernorSettings.setProposalThreshold(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L79-L81)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L79-L81


 - [ ] ID-248
	- [OwnableLocked.renounceOwnership()](src/libraries/OwnableLocked.sol#L25)

src/libraries/OwnableLocked.sol#L25


 - [ ] ID-249
	- [ZivoeTranches.pushToLocker(address,uint256,bytes)](src/ZivoeTranches.sol#L180-L186)

src/ZivoeTranches.sol#L180-L186


 - [ ] ID-250
	- [OCT_ZVL.claim()](src/lockers/OCT/OCT_ZVL.sol#L81-L87)

src/lockers/OCT/OCT_ZVL.sol#L81-L87


 - [ ] ID-251
	- [ZivoeDAO.pushMulti(address,address[],uint256[],bytes[])](src/ZivoeDAO.sol#L282-L301)

src/ZivoeDAO.sol#L282-L301


 - [ ] ID-252
	- [ZivoeDAO.pullERC1155(address,address,uint256[],uint256[],bytes)](src/ZivoeDAO.sol#L455-L465)

src/ZivoeDAO.sol#L455-L465


 - [ ] ID-253
	- [ZivoeLocker.pullFromLockerERC721(address,uint256,bytes)](src/ZivoeLocker.sol#L144-L147)

src/ZivoeLocker.sol#L144-L147


 - [ ] ID-254
	- [ZivoeDAO.pushERC1155(address,address,uint256[],uint256[],bytes)](src/ZivoeDAO.sol#L429-L446)

src/ZivoeDAO.sol#L429-L446


 - [ ] ID-255
	- [OCY_Convex_A.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_A.sol#L174-L207)

src/lockers/OCY/OCY_Convex_A.sol#L174-L207


 - [ ] ID-256
	- [ZivoeLocker.pullFromLockerERC1155(address,uint256[],uint256[],bytes)](src/ZivoeLocker.sol#L192-L197)

src/ZivoeLocker.sol#L192-L197


 - [ ] ID-257
	- [OCR_Modular.pushToLocker(address,uint256,bytes)](src/lockers/OCR/OCR_Modular.sol#L180-L185)

src/lockers/OCR/OCR_Modular.sol#L180-L185


 - [ ] ID-258
	- [OCT_YDL.convertAndForward(address,bytes)](src/lockers/OCT/OCT_YDL.sol#L97-L110)

src/lockers/OCT/OCT_YDL.sol#L97-L110


 - [ ] ID-259
	- [OCE_ZVE.pushToLocker(address,uint256,bytes)](src/lockers/OCE/OCE_ZVE.sol#L122-L128)

src/lockers/OCE/OCE_ZVE.sol#L122-L128


 - [ ] ID-260
	- [ZivoeDAO.pushERC721(address,address,uint256,bytes)](src/ZivoeDAO.sol#L344-L357)

src/ZivoeDAO.sol#L344-L357


 - [ ] ID-261
	- [ZivoeGovernorV2.updateQuorumNumerator(uint256)](src/ZivoeGovernorV2.sol#L107-L110)

src/ZivoeGovernorV2.sol#L107-L110


 - [ ] ID-262
	- [ZivoeGovernorV2.setProposalThreshold(uint256)](src/ZivoeGovernorV2.sol#L98-L104)

src/ZivoeGovernorV2.sol#L98-L104


 - [ ] ID-263
	- [OCY_Convex_C.pullFromLocker(address,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L146-L169)

src/lockers/OCY/OCY_Convex_C.sol#L146-L169


 - [ ] ID-264
	- [GovernorVotesQuorumFraction.updateQuorumNumerator(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L86-L88)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L86-L88


 - [ ] ID-265
	- [OCR_Modular.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCR/OCR_Modular.sol#L202-L210)

src/lockers/OCR/OCR_Modular.sol#L202-L210


 - [ ] ID-266
	- [ZivoeTLC.scheduleBatch(address[],uint256[],bytes[],bytes32,bytes32,uint256)](src/libraries/ZivoeTLC.sol#L254-L270)

src/libraries/ZivoeTLC.sol#L254-L270


 - [ ] ID-267
	- [ZivoeDAO.pull(address,address,bytes)](src/ZivoeDAO.sol#L255-L259)

src/ZivoeDAO.sol#L255-L259


 - [ ] ID-268
	- [OCY_Convex_B.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L197-L223)

src/lockers/OCY/OCY_Convex_B.sol#L197-L223


 - [ ] ID-269
	- [OCY_Convex_C.pullFromLockerPartial(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_C.sol#L175-L196)

src/lockers/OCY/OCY_Convex_C.sol#L175-L196


 - [ ] ID-270
	- [ZivoeGovernorV2.setVotingDelay(uint256)](src/ZivoeGovernorV2.sol#L84-L87)

src/ZivoeGovernorV2.sol#L84-L87


 - [ ] ID-271
	- [OCG_Defaults.increaseDefaults(uint256)](src/lockers/OCG/OCG_Defaults.sol#L64-L66)

src/lockers/OCG/OCG_Defaults.sol#L64-L66


 - [ ] ID-272
	- [OCG_Defaults.decreaseDefaults(uint256)](src/lockers/OCG/OCG_Defaults.sol#L57-L59)

src/lockers/OCG/OCG_Defaults.sol#L57-L59


 - [ ] ID-273
	- [ZivoeTLC.executeBatch(address[],uint256[],bytes[],bytes32,bytes32)](src/libraries/ZivoeTLC.sol#L339-L373)

src/libraries/ZivoeTLC.sol#L339-L373


 - [ ] ID-274
	- [AccessControl.revokeRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L159-L161)

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L159-L161


 - [ ] ID-275
	- [OCT_DAO.convertAndForward(address,address,bytes)](src/lockers/OCT/OCT_DAO.sol#L87-L99)

src/lockers/OCT/OCT_DAO.sol#L87-L99


 - [ ] ID-276
	- [ZivoeLocker.pullFromLockerPartial(address,uint256,bytes)](src/ZivoeLocker.sol#L90-L93)

src/ZivoeLocker.sol#L90-L93


 - [ ] ID-277
	- [ZivoeDAO.pullMulti(address,address[],bytes[])](src/ZivoeDAO.sol#L309-L316)

src/ZivoeDAO.sol#L309-L316


 - [ ] ID-278
	- [ZivoeLocker.pushToLockerERC1155(address,uint256[],uint256[],bytes)](src/ZivoeLocker.sol#L180-L185)

src/ZivoeLocker.sol#L180-L185


 - [ ] ID-279
	- [OCY_Convex_B.pushToLocker(address,uint256,bytes)](src/lockers/OCY/OCY_Convex_B.sol#L112-L159)

src/lockers/OCY/OCY_Convex_B.sol#L112-L159


 - [ ] ID-280
	- [ZivoeLocker.pushToLockerERC721(address,uint256,bytes)](src/ZivoeLocker.sol#L135-L138)

src/ZivoeLocker.sol#L135-L138


## dead-function
Impact: Informational
Confidence: Medium
 - [ ] ID-281
[ZivoeVotes._subtract(uint256,uint256)](src/libraries/ZivoeVotes.sol#L133-L135) is never used and should be removed

src/libraries/ZivoeVotes.sol#L133-L135


 - [ ] ID-282
[ZivoeVotes._add(uint256,uint256)](src/libraries/ZivoeVotes.sol#L129-L131) is never used and should be removed

src/libraries/ZivoeVotes.sol#L129-L131


## error-msg
Impact: Informational
Confidence: Medium
 - [ ] ID-283
require() missing error messages
	 - [require(bool)(_OCT_YDL != address(0))](src/lockers/OCC/OCC_Modular.sol#L729)

src/lockers/OCC/OCC_Modular.sol#L729


 - [ ] ID-284
require() missing error messages
	 - [require(bool)(_msgSender() == IZivoeGlobals_OCC(GBL).ZVL())](src/lockers/OCC/OCC_Modular.sol#L728)

src/lockers/OCC/OCC_Modular.sol#L728


 - [ ] ID-285
require() missing error messages
	 - [require(bool)(denominator > prod1)](lib/openzeppelin-contracts/contracts/utils/math/Math.sol#L78)

lib/openzeppelin-contracts/contracts/utils/math/Math.sol#L78


 - [ ] ID-286
require() missing error messages
	 - [require(bool)(_executor() == address(this))](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L86)

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L86


## price-manipulation-info
Impact: Informational
Confidence: Medium
 - [ ] ID-287
Potential price manipulation risk:
	- In function _delegate
		-- [delegatorBalance = balanceOf(delegator)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L213) have potential price manipulated risk from delegatorBalance and call None which could influence variable:delegatorBalance

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L213


 - [ ] ID-288
Potential price manipulation risk:
	- In function pushToLockerMulti
		-- [(preBasis,None) = fetchBasis()](src/lockers/OCL/OCL_ZVE.sol#L189) have potential price manipulated risk from preBasis and call None which could influence variable:basis
	- In function pushToLockerMulti
		-- [(postBasis) = fetchBasis()](src/lockers/OCL/OCL_ZVE.sol#L212) have potential price manipulated risk from postBasis and call None which could influence variable:basis
	- In function fetchBasis
		-- [pairAssetBalance = IERC20(pairAsset).balanceOf(pool)](src/lockers/OCL/OCL_ZVE.sol#L338) have potential price manipulated risk from pairAssetBalance and call None which could influence variable:pairAssetBalance

src/lockers/OCL/OCL_ZVE.sol#L189


## uncontrolled-resource-consumption
Impact: Informational
Confidence: Medium
 - [ ] ID-289
Potential DoS Gas Limit Attack occur in[Governor._beforeExecute(uint256,address[],uint256[],bytes[],bytes32)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L333-L347)[BEGIN_LOOP](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L341-L345)

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L333-L347


 - [ ] ID-290
Potential DoS Gas Limit Attack occur in[OCY_Convex_A.claimRewards(bool)](src/lockers/OCY/OCY_Convex_A.sol#L246-L266)[BEGIN_LOOP](src/lockers/OCY/OCY_Convex_A.sol#L258-L264)

src/lockers/OCY/OCY_Convex_A.sol#L246-L266


 - [ ] ID-291
Potential DoS Gas Limit Attack occur in[ZivoeYDL.distributeYield()](src/ZivoeYDL.sol#L213-L310)[BEGIN_LOOP](src/ZivoeYDL.sol#L242-L269)

src/ZivoeYDL.sol#L213-L310


 - [ ] ID-292
Potential DoS Gas Limit Attack occur in[Governor._execute(uint256,address[],uint256[],bytes[],bytes32)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L316-L328)[BEGIN_LOOP](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L324-L327)

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L316-L328


 - [ ] ID-293
Potential DoS Gas Limit Attack occur in[OCR_Modular.tickEpoch()](src/lockers/OCR/OCR_Modular.sol#L310-L336)[BEGIN_LOOP](src/lockers/OCR/OCR_Modular.sol#L311-L335)

src/lockers/OCR/OCR_Modular.sol#L310-L336


 - [ ] ID-294
Potential DoS Gas Limit Attack occur in[OCY_Convex_C.claimRewards(bool)](src/lockers/OCY/OCY_Convex_C.sol#L200-L220)[BEGIN_LOOP](src/lockers/OCY/OCY_Convex_C.sol#L212-L218)

src/lockers/OCY/OCY_Convex_C.sol#L200-L220


 - [ ] ID-295
Potential DoS Gas Limit Attack occur in[OCY_Convex_B.claimRewards(bool)](src/lockers/OCY/OCY_Convex_B.sol#L227-L247)[BEGIN_LOOP](src/lockers/OCY/OCY_Convex_B.sol#L239-L245)

src/lockers/OCY/OCY_Convex_B.sol#L227-L247


## unnecessary-public-function-modifier
Impact: Informational
Confidence: High
 - [ ] ID-296
function:[Presale.depositETH()](src/misc/Presale.sol#L168-L184)is public and can be replaced with external 

src/misc/Presale.sol#L168-L184


 - [ ] ID-297
function:[ZivoeLocker.canPushERC721()](src/ZivoeLocker.sol#L52)is public and can be replaced with external 

src/ZivoeLocker.sol#L52


 - [ ] ID-298
function:[Governor.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L92-L104)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L92-L104


 - [ ] ID-299
function:[ERC20.symbol()](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L70-L72)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L70-L72


 - [ ] ID-300
function:[AccessControl.grantRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L144-L146)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L144-L146


 - [ ] ID-301
function:[ZivoeLocker.canPull()](src/ZivoeLocker.sol#L37)is public and can be replaced with external 

src/ZivoeLocker.sol#L37


 - [ ] ID-302
function:[ERC20Permit.nonces(address)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L73-L75)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L73-L75


 - [ ] ID-303
function:[ERC20Votes.checkpoints(address,uint32)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L43-L45)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L43-L45


 - [ ] ID-304
function:[ZivoeLocker.canPushERC1155()](src/ZivoeLocker.sol#L64)is public and can be replaced with external 

src/ZivoeLocker.sol#L64


 - [ ] ID-305
function:[ZivoeTLC.scheduleBatch(address[],uint256[],bytes[],bytes32,bytes32,uint256)](src/libraries/ZivoeTLC.sol#L254-L270)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L254-L270


 - [ ] ID-306
function:[Governor.propose(address[],uint256[],bytes[],string)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L245-L284)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L245-L284


 - [ ] ID-307
function:[IGovernor.castVoteWithReasonAndParamsBySig(uint256,uint8,string,bytes,uint8,bytes32,bytes32)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L267-L275)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L267-L275


 - [ ] ID-308
function:[IGovernor.castVoteWithReason(uint256,uint8,string)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L231-L235)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L231-L235


 - [ ] ID-309
function:[ERC20Votes.getPastTotalSupply(uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L89-L92)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L89-L92


 - [ ] ID-310
function:[ERC20.totalSupply()](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L94-L96)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L94-L96


 - [ ] ID-311
function:[Governor.proposalThreshold()](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L196-L198)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L196-L198


 - [ ] ID-312
function:[Governor.name()](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L109-L111)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L109-L111


 - [ ] ID-313
function:[Presale.depositStablecoin(address,uint256)](src/misc/Presale.sol#L147-L165)is public and can be replaced with external 

src/misc/Presale.sol#L147-L165


 - [ ] ID-314
function:[AccessControl.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L77-L79)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L77-L79


 - [ ] ID-315
function:[ERC20.approve(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L136-L140)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L136-L140


 - [ ] ID-316
function:[ERC20.balanceOf(address)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L101-L103)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L101-L103


 - [ ] ID-317
function:[IGovernor.quorum(uint256)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L164)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L164


 - [ ] ID-318
function:[ZivoeLocker.canPullMulti()](src/ZivoeLocker.sol#L46)is public and can be replaced with external 

src/ZivoeLocker.sol#L46


 - [ ] ID-319
function:[ERC1155Holder.onERC1155Received(address,address,uint256,uint256,bytes)](lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#L17-L25)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#L17-L25


 - [ ] ID-320
function:[IGovernor.name()](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L76)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L76


 - [ ] ID-321
function:[ZivoeVotes.checkpoints(address,uint32)](src/libraries/ZivoeVotes.sol#L24-L26)is public and can be replaced with external 

src/libraries/ZivoeVotes.sol#L24-L26


 - [ ] ID-322
function:[Governor.execute(address[],uint256[],bytes[],bytes32)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L289-L311)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L289-L311


 - [ ] ID-323
function:[ZivoeLocker.canPullMultiPartial()](src/ZivoeLocker.sol#L49)is public and can be replaced with external 

src/ZivoeLocker.sol#L49


 - [ ] ID-324
function:[IGovernor.hasVoted(uint256,address)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L189)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L189


 - [ ] ID-325
function:[IGovernor.castVoteBySig(uint256,uint8,uint8,bytes32,bytes32)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L254-L260)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L254-L260


 - [ ] ID-326
function:[AccessControl.revokeRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L159-L161)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L159-L161


 - [ ] ID-327
function:[IGovernor.COUNTING_MODE()](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L107)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L107


 - [ ] ID-328
function:[ZivoeLocker.canPushMultiERC721()](src/ZivoeLocker.sol#L58)is public and can be replaced with external 

src/ZivoeLocker.sol#L58


 - [ ] ID-329
function:[Governor.castVoteWithReasonAndParamsBySig(uint256,uint8,string,bytes,uint8,bytes32,bytes32)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L465-L492)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L465-L492


 - [ ] ID-330
function:[ZivoeVotes.getVotes(address)](src/libraries/ZivoeVotes.sol#L38-L41)is public and can be replaced with external 

src/libraries/ZivoeVotes.sol#L38-L41


 - [ ] ID-331
function:[ERC20Votes.getVotes(address)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L64-L67)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L64-L67


 - [ ] ID-332
function:[IGovernor.state(uint256)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L124)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L124


 - [ ] ID-333
function:[OwnableLocked.renounceOwnership()](src/libraries/OwnableLocked.sol#L25)is public and can be replaced with external 

src/libraries/OwnableLocked.sol#L25


 - [ ] ID-334
function:[GovernorSettings.setVotingPeriod(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L70-L72)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L70-L72


 - [ ] ID-335
function:[IGovernor.getVotes(address,uint256)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L173)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L173


 - [ ] ID-336
function:[ERC20Votes.getPastVotes(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L76-L79)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L76-L79


 - [ ] ID-337
function:[ERC20Votes.numCheckpoints(address)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L50-L52)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L50-L52


 - [ ] ID-338
function:[ZivoeTLC.onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)](src/libraries/ZivoeTLC.sol#L465-L473)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L465-L473


 - [ ] ID-339
function:[ZivoeGTC.state(uint256)](src/libraries/ZivoeGTC.sol#L53-L71)is public and can be replaced with external 

src/libraries/ZivoeGTC.sol#L53-L71


 - [ ] ID-340
function:[ERC20.transferFrom(address,address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L158-L167)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L158-L167


 - [ ] ID-341
function:[ZivoeGTC.timelock()](src/libraries/ZivoeGTC.sol#L76-L78)is public and can be replaced with external 

src/libraries/ZivoeGTC.sol#L76-L78


 - [ ] ID-342
function:[Ownable.transferOwnership(address)](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L69-L72)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L69-L72


 - [ ] ID-343
function:[ERC20.name()](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L62-L64)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L62-L64


 - [ ] ID-344
function:[ZivoeTLC.supportsInterface(bytes4)](src/libraries/ZivoeTLC.sol#L136-L138)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L136-L138


 - [ ] ID-345
function:[OCY_OUSD.rebase()](src/lockers/OCY/OCY_OUSD.sol#L127-L129)is public and can be replaced with external 

src/lockers/OCY/OCY_OUSD.sol#L127-L129


 - [ ] ID-346
function:[ZivoeVotes.numCheckpoints(address)](src/libraries/ZivoeVotes.sol#L31-L33)is public and can be replaced with external 

src/libraries/ZivoeVotes.sol#L31-L33


 - [ ] ID-347
function:[ZivoeRewards.earned(address,address)](src/ZivoeRewards.sol#L180-L184)is public and can be replaced with external 

src/ZivoeRewards.sol#L180-L184


 - [ ] ID-348
function:[ZivoeLocker.canPullPartial()](src/ZivoeLocker.sol#L40)is public and can be replaced with external 

src/ZivoeLocker.sol#L40


 - [ ] ID-349
function:[Ownable.renounceOwnership()](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L61-L63)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L61-L63


 - [ ] ID-350
function:[IGovernor.castVote(uint256,uint8)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L224)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L224


 - [ ] ID-351
function:[ERC20.increaseAllowance(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L181-L185)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L181-L185


 - [ ] ID-352
function:[ERC721Holder.onERC721Received(address,address,uint256,bytes)](lib/openzeppelin-contracts/contracts/token/ERC721/utils/ERC721Holder.sol#L20-L27)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/utils/ERC721Holder.sol#L20-L27


 - [ ] ID-353
function:[Governor.onERC721Received(address,address,uint256,bytes)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L563-L570)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L563-L570


 - [ ] ID-354
function:[OCG_Defaults.decreaseDefaults(uint256)](src/lockers/OCG/OCG_Defaults.sol#L57-L59)is public and can be replaced with external 

src/lockers/OCG/OCG_Defaults.sol#L57-L59


 - [ ] ID-355
function:[ERC1155Holder.onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)](lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#L27-L35)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#L27-L35


 - [ ] ID-356
function:[Governor.onERC1155Received(address,address,uint256,uint256,bytes)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L575-L583)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L575-L583


 - [ ] ID-357
function:[ZivoeGTC.queue(address[],uint256[],bytes[],bytes32)](src/libraries/ZivoeGTC.sol#L91-L108)is public and can be replaced with external 

src/libraries/ZivoeGTC.sol#L91-L108


 - [ ] ID-358
function:[OwnableLocked.transferOwnership(address)](src/libraries/OwnableLocked.sol#L31-L34)is public and can be replaced with external 

src/libraries/OwnableLocked.sol#L31-L34


 - [ ] ID-359
function:[ERC165.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#L26-L28)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#L26-L28


 - [ ] ID-360
function:[GovernorSettings.setVotingDelay(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L61-L63)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L61-L63


 - [ ] ID-361
function:[IGovernor.proposalDeadline(uint256)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L139)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L139


 - [ ] ID-362
function:[ERC20Votes.delegate(address)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L139-L141)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L139-L141


 - [ ] ID-363
function:[GovernorCountingSimple.hasVoted(uint256,address)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#L43-L45)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#L43-L45


 - [ ] ID-364
function:[GovernorVotesQuorumFraction.quorum(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L72-L74)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L72-L74


 - [ ] ID-365
function:[IGovernor.castVoteWithReasonAndParams(uint256,uint8,string,bytes)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L242-L247)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L242-L247


 - [ ] ID-366
function:[ZivoeLocker.canPullERC721()](src/ZivoeLocker.sol#L55)is public and can be replaced with external 

src/ZivoeLocker.sol#L55


 - [ ] ID-367
function:[ERC20Votes.delegateBySig(address,uint256,uint256,uint8,bytes32,bytes32)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L146-L163)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L146-L163


 - [ ] ID-368
function:[ZivoeVotes.getPastTotalSupply(uint256)](src/libraries/ZivoeVotes.sol#L63-L66)is public and can be replaced with external 

src/libraries/ZivoeVotes.sol#L63-L66


 - [ ] ID-369
function:[Governor.onERC1155BatchReceived(address,address,uint256[],uint256[],bytes)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L588-L596)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L588-L596


 - [ ] ID-370
function:[GovernorCountingSimple.proposalVotes(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#L50-L62)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#L50-L62


 - [ ] ID-371
function:[Governor.castVoteBySig(uint256,uint8,uint8,bytes32,bytes32)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L446-L460)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L446-L460


 - [ ] ID-372
function:[ZivoeRewardsVesting.earned(address,address)](src/ZivoeRewardsVesting.sol#L300-L304)is public and can be replaced with external 

src/ZivoeRewardsVesting.sol#L300-L304


 - [ ] ID-373
function:[Governor.getVotesWithParams(address,uint256,bytes)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L402-L408)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L402-L408


 - [ ] ID-374
function:[ZivoeTLC.schedule(address,uint256,bytes,bytes32,bytes32,uint256)](src/libraries/ZivoeTLC.sol#L232-L243)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L232-L243


 - [ ] ID-375
function:[IGovernorTimelock.timelock()](lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#L16)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#L16


 - [ ] ID-376
function:[IGovernorTimelock.proposalEta(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#L18)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#L18


 - [ ] ID-377
function:[ZivoeGovernorV2.updateQuorumNumerator(uint256)](src/ZivoeGovernorV2.sol#L107-L110)is public and can be replaced with external 

src/ZivoeGovernorV2.sol#L107-L110


 - [ ] ID-378
function:[IGovernor.getVotesWithParams(address,uint256,bytes)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L179-L183)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L179-L183


 - [ ] ID-379
function:[ZivoeLocker.canPullERC1155()](src/ZivoeLocker.sol#L67)is public and can be replaced with external 

src/ZivoeLocker.sol#L67


 - [ ] ID-380
function:[ZivoeGTC.supportsInterface(bytes4)](src/libraries/ZivoeGTC.sol#L46-L48)is public and can be replaced with external 

src/libraries/ZivoeGTC.sol#L46-L48


 - [ ] ID-381
function:[ZivoeLocker.canPushMulti()](src/ZivoeLocker.sol#L43)is public and can be replaced with external 

src/ZivoeLocker.sol#L43


 - [ ] ID-382
function:[OCG_Defaults.increaseDefaults(uint256)](src/lockers/OCG/OCG_Defaults.sol#L64-L66)is public and can be replaced with external 

src/lockers/OCG/OCG_Defaults.sol#L64-L66


 - [ ] ID-383
function:[IGovernor.proposalSnapshot(uint256)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L132)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L132


 - [ ] ID-384
function:[IGovernor.propose(address[],uint256[],bytes[],string)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L197-L202)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L197-L202


 - [ ] ID-385
function:[ZivoeVotes.getPastVotes(address,uint256)](src/libraries/ZivoeVotes.sol#L50-L53)is public and can be replaced with external 

src/libraries/ZivoeVotes.sol#L50-L53


 - [ ] ID-386
function:[ERC20.decreaseAllowance(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L201-L210)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L201-L210


 - [ ] ID-387
function:[ERC20.transfer(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L113-L117)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L113-L117


 - [ ] ID-388
function:[ZivoeTLC.executeBatch(address[],uint256[],bytes[],bytes32,bytes32)](src/libraries/ZivoeTLC.sol#L339-L373)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L339-L373


 - [ ] ID-389
function:[ERC1155Receiver.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Receiver.sol#L16-L18)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Receiver.sol#L16-L18


 - [ ] ID-390
function:[ZivoeGovernorV2.setVotingDelay(uint256)](src/ZivoeGovernorV2.sol#L84-L87)is public and can be replaced with external 

src/ZivoeGovernorV2.sol#L84-L87


 - [ ] ID-391
function:[IGovernor.hashProposal(address[],uint256[],bytes[],bytes32)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L113-L118)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L113-L118


 - [ ] ID-392
function:[ZivoeGovernorV2.setVotingPeriod(uint256)](src/ZivoeGovernorV2.sol#L91-L95)is public and can be replaced with external 

src/ZivoeGovernorV2.sol#L91-L95


 - [ ] ID-393
function:[IGovernor.execute(address[],uint256[],bytes[],bytes32)](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L212-L217)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L212-L217


 - [ ] ID-394
function:[ZivoeGTC.proposalEta(uint256)](src/libraries/ZivoeGTC.sol#L83-L86)is public and can be replaced with external 

src/libraries/ZivoeGTC.sol#L83-L86


 - [ ] ID-395
function:[IGovernor.votingDelay()](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L146)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L146


 - [ ] ID-396
function:[IGovernor.version()](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L82)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L82


 - [ ] ID-397
function:[GovernorCountingSimple.COUNTING_MODE()](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#L36-L38)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#L36-L38


 - [ ] ID-398
function:[ZivoeTLC.onERC1155Received(address,address,uint256,uint256,bytes)](src/libraries/ZivoeTLC.sol#L452-L460)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L452-L460


 - [ ] ID-399
function:[Governor.castVoteWithReason(uint256,uint8,string)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L421-L428)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L421-L428


 - [ ] ID-400
function:[ZivoeTLC.cancel(bytes32)](src/libraries/ZivoeTLC.sol#L288-L293)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L288-L293


 - [ ] ID-401
function:[IGovernor.votingPeriod()](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L155)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L155


 - [ ] ID-402
function:[ZivoeGovernorV2.supportsInterface(bytes4)](src/ZivoeGovernorV2.sol#L118-L120)is public and can be replaced with external 

src/ZivoeGovernorV2.sol#L118-L120


 - [ ] ID-403
function:[ZivoeTLC.execute(address,uint256,bytes,bytes32,bytes32)](src/libraries/ZivoeTLC.sol#L307-L328)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L307-L328


 - [ ] ID-404
function:[MockStablecoin.decimals()](src/misc/MockStablecoin.sol#L47-L49)is public and can be replaced with external 

src/misc/MockStablecoin.sol#L47-L49


 - [ ] ID-405
function:[ERC20.decimals()](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L87-L89)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L87-L89


 - [ ] ID-406
function:[OCR_Modular.tickEpoch()](src/lockers/OCR/OCR_Modular.sol#L310-L336)is public and can be replaced with external 

src/lockers/OCR/OCR_Modular.sol#L310-L336


 - [ ] ID-407
function:[ZivoeTLC.onERC721Received(address,address,uint256,bytes)](src/libraries/ZivoeTLC.sol#L440-L447)is public and can be replaced with external 

src/libraries/ZivoeTLC.sol#L440-L447


 - [ ] ID-408
function:[Governor.castVote(uint256,uint8)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L413-L416)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L413-L416


 - [ ] ID-409
function:[AccessControl.renounceRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L179-L183)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L179-L183


 - [ ] ID-410
function:[OwnableLocked.transferOwnershipAndLock(address)](src/libraries/OwnableLocked.sol#L40-L44)is public and can be replaced with external 

src/libraries/OwnableLocked.sol#L40-L44


 - [ ] ID-411
function:[ZivoeGovernorV2.setProposalThreshold(uint256)](src/ZivoeGovernorV2.sol#L98-L104)is public and can be replaced with external 

src/ZivoeGovernorV2.sol#L98-L104


 - [ ] ID-412
function:[ZivoeMath.yieldTarget(uint256,uint256,uint256,uint256,uint256)](src/ZivoeMath.sol#L121-L123)is public and can be replaced with external 

src/ZivoeMath.sol#L121-L123


 - [ ] ID-413
function:[ZivoeLocker.canPullMultiERC721()](src/ZivoeLocker.sol#L61)is public and can be replaced with external 

src/ZivoeLocker.sol#L61


 - [ ] ID-414
function:[GovernorSettings.setProposalThreshold(uint256)](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L79-L81)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L79-L81


 - [ ] ID-415
function:[Governor.castVoteWithReasonAndParams(uint256,uint8,string,bytes)](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L433-L441)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L433-L441


 - [ ] ID-416
function:[IGovernorTimelock.queue(address[],uint256[],bytes[],bytes32)](lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#L20-L25)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#L20-L25


 - [ ] ID-417
function:[GovernorSettings.proposalThreshold()](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L52-L54)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L52-L54


 - [ ] ID-418
function:[ZivoeLocker.canPush()](src/ZivoeLocker.sol#L34)is public and can be replaced with external 

src/ZivoeLocker.sol#L34


 - [ ] ID-419
function:[ERC20Permit.permit(address,address,uint256,uint256,uint8,bytes32,bytes32)](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L49-L68)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L49-L68


## version-only
Impact: Informational
Confidence: High
 - [ ] ID-420
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#L4


 - [ ] ID-421
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/cryptography/EIP712.sol#L4)

lib/openzeppelin-contracts/contracts/utils/cryptography/EIP712.sol#L4


 - [ ] ID-422
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/libraries/OwnableLocked.sol#L2)

src/libraries/OwnableLocked.sol#L2


 - [ ] ID-423
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCC/OCC_Modular.sol#L2)

src/lockers/OCC/OCC_Modular.sol#L2


 - [ ] ID-424
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeDAO.sol#L2)

src/ZivoeDAO.sol#L2


 - [ ] ID-425
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCY/OCY_OUSD.sol#L2)

src/lockers/OCY/OCY_OUSD.sol#L2


 - [ ] ID-426
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol#L4


 - [ ] ID-427
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeToken.sol#L2)

src/ZivoeToken.sol#L2


 - [ ] ID-428
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](src/libraries/ZivoeVotes.sol#L3)

src/libraries/ZivoeVotes.sol#L3


 - [ ] ID-429
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeGlobals.sol#L2)

src/ZivoeGlobals.sol#L2


 - [ ] ID-430
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeGovernorV2.sol#L2)

src/ZivoeGovernorV2.sol#L2


 - [ ] ID-431
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/ERC20Votes.sol#L4


 - [ ] ID-432
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/security/ReentrancyGuard.sol#L4)

lib/openzeppelin-contracts/contracts/security/ReentrancyGuard.sol#L4


 - [ ] ID-433
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCY/OCY_Convex_B.sol#L2)

src/lockers/OCY/OCY_Convex_B.sol#L2


 - [ ] ID-434
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC721/utils/ERC721Holder.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/utils/ERC721Holder.sol#L4


 - [ ] ID-435
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol#L4)

lib/openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol#L4


 - [ ] ID-436
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCG/OCG_ERC721.sol#L2)

src/lockers/OCG/OCG_ERC721.sol#L2


 - [ ] ID-437
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCT/OCT_ZVL.sol#L2)

src/lockers/OCT/OCT_ZVL.sol#L2


 - [ ] ID-438
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/Timers.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Timers.sol#L4


 - [ ] ID-439
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC721/IERC721Receiver.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/IERC721Receiver.sol#L4


 - [ ] ID-440
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L4)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorSettings.sol#L4


 - [ ] ID-441
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L4)

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L4


 - [ ] ID-442
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC1155/IERC1155Receiver.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC1155/IERC1155Receiver.sol#L4


 - [ ] ID-443
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeITO.sol#L2)

src/ZivoeITO.sol#L2


 - [ ] ID-444
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeTranches.sol#L2)

src/ZivoeTranches.sol#L2


 - [ ] ID-445
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/access/IAccessControl.sol#L4)

lib/openzeppelin-contracts/contracts/access/IAccessControl.sol#L4


 - [ ] ID-446
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/math/Math.sol#L4)

lib/openzeppelin-contracts/contracts/utils/math/Math.sol#L4


 - [ ] ID-447
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Holder.sol#L4


 - [ ] ID-448
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/governance/Governor.sol#L4)

lib/openzeppelin-contracts/contracts/governance/Governor.sol#L4


 - [ ] ID-449
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeYDL.sol#L2)

src/ZivoeYDL.sol#L2


 - [ ] ID-450
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/Strings.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Strings.sol#L4


 - [ ] ID-451
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCG/OCG_ERC20_FreeClaim.sol#L2)

src/lockers/OCG/OCG_ERC20_FreeClaim.sol#L2


 - [ ] ID-452
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/misc/Presale.sol#L2)

src/misc/Presale.sol#L2


 - [ ] ID-453
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/governance/utils/IVotes.sol#L3)

lib/openzeppelin-contracts/contracts/governance/utils/IVotes.sol#L3


 - [ ] ID-454
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeLocker.sol#L2)

src/ZivoeLocker.sol#L2


 - [ ] ID-455
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCY/OCY_Convex_A.sol#L2)

src/lockers/OCY/OCY_Convex_A.sol#L2


 - [ ] ID-456
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol#L4)

lib/openzeppelin-contracts/contracts/utils/math/SafeMath.sol#L4


 - [ ] ID-457
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](src/libraries/ZivoeTLC.sol#L3)

src/libraries/ZivoeTLC.sol#L3


 - [ ] ID-458
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCG/OCG_ERC1155.sol#L2)

src/lockers/OCG/OCG_ERC1155.sol#L2


 - [ ] ID-459
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCR/OCR_Modular.sol#L2)

src/lockers/OCR/OCR_Modular.sol#L2


 - [ ] ID-460
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#L4)

lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#L4


 - [ ] ID-461
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCY/OCY_Convex_C.sol#L2)

src/lockers/OCY/OCY_Convex_C.sol#L2


 - [ ] ID-462
	Pragma confirmed better, here is pragma solidity^0.8.4. [^0.8.4](lib/openzeppelin-contracts/contracts/utils/structs/DoubleEndedQueue.sol#L3)

lib/openzeppelin-contracts/contracts/utils/structs/DoubleEndedQueue.sol#L3


 - [ ] ID-463
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/misc/BaseContractTemplate.sol#L2)

src/misc/BaseContractTemplate.sol#L2


 - [ ] ID-464
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeTrancheToken.sol#L2)

src/ZivoeTrancheToken.sol#L2


 - [ ] ID-465
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Receiver.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC1155/utils/ERC1155Receiver.sol#L4


 - [ ] ID-466
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#L4)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorCountingSimple.sol#L4


 - [ ] ID-467
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCT/OCT_YDL.sol#L2)

src/lockers/OCT/OCT_YDL.sol#L2


 - [ ] ID-468
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/math/SafeCast.sol#L5)

lib/openzeppelin-contracts/contracts/utils/math/SafeCast.sol#L5


 - [ ] ID-469
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L4


 - [ ] ID-470
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](src/libraries/ZivoeGTC.sol#L3)

src/libraries/ZivoeGTC.sol#L3


 - [ ] ID-471
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/Counters.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Counters.sol#L4


 - [ ] ID-472
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L4


 - [ ] ID-473
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeRewardsVesting.sol#L2)

src/ZivoeRewardsVesting.sol#L2


 - [ ] ID-474
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeRewards.sol#L2)

src/ZivoeRewards.sol#L2


 - [ ] ID-475
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/Context.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Context.sol#L4


 - [ ] ID-476
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4


 - [ ] ID-477
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC721/IERC721.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/IERC721.sol#L4


 - [ ] ID-478
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/misc/InterfacesAggregated.sol#L2)

src/misc/InterfacesAggregated.sol#L2


 - [ ] ID-479
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/libraries/FloorMath.sol#L2)

src/libraries/FloorMath.sol#L2


 - [ ] ID-480
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L4)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L4


 - [ ] ID-481
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCG/OCG_ERC20.sol#L2)

src/lockers/OCG/OCG_ERC20.sol#L2


 - [ ] ID-482
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#L4)

lib/openzeppelin-contracts/contracts/governance/extensions/IGovernorTimelock.sol#L4


 - [ ] ID-483
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/Utility/ZivoeSwapper.sol#L2)

src/lockers/Utility/ZivoeSwapper.sol#L2


 - [ ] ID-484
	Pragma confirmed better, here is pragma solidity^0.8.1. [^0.8.1](lib/openzeppelin-contracts/contracts/utils/Address.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Address.sol#L4


 - [ ] ID-485
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/Checkpoints.sol#L5)

lib/openzeppelin-contracts/contracts/utils/Checkpoints.sol#L5


 - [ ] ID-486
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#L4


 - [ ] ID-487
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCE/OCE_ZVE.sol#L2)

src/lockers/OCE/OCE_ZVE.sol#L2


 - [ ] ID-488
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotes.sol#L4)

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotes.sol#L4


 - [ ] ID-489
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCL/OCL_ZVE.sol#L2)

src/lockers/OCL/OCL_ZVE.sol#L2


 - [ ] ID-490
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/misc/MockStablecoin.sol#L2)

src/misc/MockStablecoin.sol#L2


 - [ ] ID-491
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/token/ERC1155/IERC1155.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC1155/IERC1155.sol#L4


 - [ ] ID-492
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCG/OCG_Defaults.sol#L2)

src/lockers/OCG/OCG_Defaults.sol#L2


 - [ ] ID-493
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/lockers/OCT/OCT_DAO.sol#L2)

src/lockers/OCT/OCT_DAO.sol#L2


 - [ ] ID-494
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/utils/introspection/IERC165.sol#L4)

lib/openzeppelin-contracts/contracts/utils/introspection/IERC165.sol#L4


 - [ ] ID-495
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L4)

lib/openzeppelin-contracts/contracts/governance/IGovernor.sol#L4


 - [ ] ID-496
	Pragma confirmed better, here is pragma solidity^0.8.17. [^0.8.17](src/ZivoeMath.sol#L2)

src/ZivoeMath.sol#L2


 - [ ] ID-497
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L4)

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L4


## state-should-be-constant
Impact: Optimization
Confidence: High
 - [ ] ID-498
[OCY_Convex_B.convexDeposit](src/lockers/OCY/OCY_Convex_B.sol#L56) should be constant

src/lockers/OCY/OCY_Convex_B.sol#L56


 - [ ] ID-499
[Presale.presaleDuration](src/misc/Presale.sol#L35) should be constant

src/misc/Presale.sol#L35


 - [ ] ID-500
[OCY_Convex_C.curveBasePoolToken](src/lockers/OCY/OCY_Convex_C.sol#L62) should be constant

src/lockers/OCY/OCY_Convex_C.sol#L62


 - [ ] ID-501
[OCY_Convex_B.convexPoolID](src/lockers/OCY/OCY_Convex_B.sol#L60) should be constant

src/lockers/OCY/OCY_Convex_B.sol#L60


 - [ ] ID-502
[OCY_Convex_A.curveMetaPool](src/lockers/OCY/OCY_Convex_A.sol#L69) should be constant

src/lockers/OCY/OCY_Convex_A.sol#L69


 - [ ] ID-503
[OCY_Convex_A.convexPoolID](src/lockers/OCY/OCY_Convex_A.sol#L64) should be constant

src/lockers/OCY/OCY_Convex_A.sol#L64


 - [ ] ID-504
[OCY_Convex_A.curveBasePool](src/lockers/OCY/OCY_Convex_A.sol#L67) should be constant

src/lockers/OCY/OCY_Convex_A.sol#L67


 - [ ] ID-505
[Presale.pointsFloor](src/misc/Presale.sol#L29) should be constant

src/misc/Presale.sol#L29


 - [ ] ID-506
[GovernorVotesQuorumFraction._quorumNumerator](lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L19) should be constant

lib/openzeppelin-contracts/contracts/governance/extensions/GovernorVotesQuorumFraction.sol#L19


 - [ ] ID-507
[OCY_Convex_B.convexPoolToken](src/lockers/OCY/OCY_Convex_B.sol#L57) should be constant

src/lockers/OCY/OCY_Convex_B.sol#L57


 - [ ] ID-508
[OCY_Convex_B.curveBasePoolToken](src/lockers/OCY/OCY_Convex_B.sol#L64) should be constant

src/lockers/OCY/OCY_Convex_B.sol#L64


 - [ ] ID-509
[OCY_Convex_A.convexRewards](src/lockers/OCY/OCY_Convex_A.sol#L61) should be constant

src/lockers/OCY/OCY_Convex_A.sol#L61


 - [ ] ID-510
[OCY_Convex_C.curveBasePool](src/lockers/OCY/OCY_Convex_C.sol#L61) should be constant

src/lockers/OCY/OCY_Convex_C.sol#L61


 - [ ] ID-511
[OCY_Convex_C.convexRewards](src/lockers/OCY/OCY_Convex_C.sol#L56) should be constant

src/lockers/OCY/OCY_Convex_C.sol#L56


 - [ ] ID-512
[OCY_Convex_C.convexPoolToken](src/lockers/OCY/OCY_Convex_C.sol#L55) should be constant

src/lockers/OCY/OCY_Convex_C.sol#L55


 - [ ] ID-513
[OCY_Convex_B.convexRewards](src/lockers/OCY/OCY_Convex_B.sol#L58) should be constant

src/lockers/OCY/OCY_Convex_B.sol#L58


 - [ ] ID-514
[Presale.pointsCeiling](src/misc/Presale.sol#L31) should be constant

src/misc/Presale.sol#L31


 - [ ] ID-515
[OCY_Convex_A.convexDeposit](src/lockers/OCY/OCY_Convex_A.sol#L60) should be constant

src/lockers/OCY/OCY_Convex_A.sol#L60


 - [ ] ID-516
[OCY_Convex_A.convexPoolToken](src/lockers/OCY/OCY_Convex_A.sol#L62) should be constant

src/lockers/OCY/OCY_Convex_A.sol#L62


 - [ ] ID-517
. analyzed (160 contracts with 86 detectors), 522 result(s) found
INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/2024-03-zivoe_scaned_output/mwe-output.json generate success.
[OCY_Convex_C.convexDeposit](src/lockers/OCY/OCY_Convex_C.sol#L54) should be constant

src/lockers/OCY/OCY_Convex_C.sol#L54


 - [ ] ID-518
[OCY_Convex_A.curveBasePoolToken](src/lockers/OCY/OCY_Convex_A.sol#L68) should be constant

src/lockers/OCY/OCY_Convex_A.sol#L68


 - [ ] ID-519
[OCY_Convex_C.convexPoolID](src/lockers/OCY/OCY_Convex_C.sol#L58) should be constant

src/lockers/OCY/OCY_Convex_C.sol#L58


 - [ ] ID-520
[ERC20Permit._PERMIT_TYPEHASH_DEPRECATED_SLOT](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L37) should be constant

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L37


 - [ ] ID-521
[OCY_Convex_B.curveBasePool](src/lockers/OCY/OCY_Convex_B.sol#L63) should be constant

src/lockers/OCY/OCY_Convex_B.sol#L63


Execution time for Falcon: 46.039420343s
