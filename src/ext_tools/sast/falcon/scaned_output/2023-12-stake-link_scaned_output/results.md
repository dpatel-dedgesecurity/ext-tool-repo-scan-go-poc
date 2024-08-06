'npx hardhat clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-12-stake-link)
'npx hardhat clean --global' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-12-stake-link)
'npx hardhat compile --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-12-stake-link)

Operator.distributeFunds(address[],uint256[]) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#349-358) sends eth to arbitrary user
	Dangerous calls:
	- receivers[i].transfer(sendAmount) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#355)
Reference:  

Reentrancy in SDLPoolSecondary.executeQueuedOperations(uint256[]) (contracts/core/sdlPool/SDLPoolSecondary.sol#247-250):
	External calls:
	- _executeQueuedLockUpdates(msg.sender,_lockIds) (contracts/core/sdlPool/SDLPoolSecondary.sol#248)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#122)
		- (success,returndata) = target.call{value: value}(data) (node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#135)
	External calls sending eth:
	- _executeQueuedLockUpdates(msg.sender,_lockIds) (contracts/core/sdlPool/SDLPoolSecondary.sol#248)
		- (success,returndata) = target.call{value: value}(data) (node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#135)
Reentrancy in PoolOwnersV1.exit() (contracts/core/test/deprecated/PoolOwnersV1.sol#137-140):
	External calls:
	- withdraw(balanceOf(msg.sender)) (contracts/core/test/deprecated/PoolOwnersV1.sol#138)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- (success,returndata) = target.call{value: value}(data) (node_modules/@openzeppelin/contracts/utils/Address.sol#137)
	External calls sending eth:
	- withdraw(balanceOf(msg.sender)) (contracts/core/test/deprecated/PoolOwnersV1.sol#138)
		- (success,returndata) = target.call{value: value}(data) (node_modules/@openzeppelin/contracts/utils/Address.sol#137)
Reference:  
Process ForkProcess-1:
Traceback (most recent call last):
  File "/usr/lib/python3.10/multiprocessing/process.py", line 314, in _bootstrap
    self.run()
  File "/usr/lib/python3.10/multiprocessing/process.py", line 108, in run
    self._target(*self._args, **self._kwargs)
  File "/usr/lib/python3.10/concurrent/futures/process.py", line 240, in _process_worker
    call_item = call_queue.get(block=True)
  File "/usr/lib/python3.10/multiprocessing/queues.py", line 122, in get
    return _ForkingPickler.loads(res)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/core/declarations/contract.py", line 1396, in __hash__
    return self._id
AttributeError: 'Contract' object has no attribute '_id'
Process ForkProcess-17:
Traceback (most recent call last):
  File "/usr/lib/python3.10/multiprocessing/process.py", line 314, in _bootstrap
    self.run()
  File "/usr/lib/python3.10/multiprocessing/process.py", line 108, in run
    self._target(*self._args, **self._kwargs)
  File "/usr/lib/python3.10/concurrent/futures/process.py", line 240, in _process_worker
    call_item = call_queue.get(block=True)
  File "/usr/lib/python3.10/multiprocessing/queues.py", line 122, in get
    return _ForkingPickler.loads(res)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/core/declarations/contract.py", line 1396, in __hash__
    return self._id
AttributeError: 'Contract' object has no attribute '_id'
Process ForkProcess-33:
Traceback (most recent call last):
  File "/usr/lib/python3.10/multiprocessing/process.py", line 314, in _bootstrap
    self.run()
  File "/usr/lib/python3.10/multiprocessing/process.py", line 108, in run
    self._target(*self._args, **self._kwargs)
  File "/usr/lib/python3.10/concurrent/futures/process.py", line 240, in _process_worker
    call_item = call_queue.get(block=True)
  File "/usr/lib/python3.10/multiprocessing/queues.py", line 122, in get
    return _ForkingPickler.loads(res)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/core/declarations/contract.py", line 1396, in __hash__
    return self._id
AttributeError: 'Contract' object has no attribute '_id'
Process ForkProcess-49:
Traceback (most recent call last):
  File "/usr/lib/python3.10/multiprocessing/process.py", line 314, in _bootstrap
    self.run()
  File "/usr/lib/python3.10/multiprocessing/process.py", line 108, in run
    self._target(*self._args, **self._kwargs)
  File "/usr/lib/python3.10/concurrent/futures/process.py", line 240, in _process_worker
    call_item = call_queue.get(block=True)
  File "/usr/lib/python3.10/multiprocessing/queues.py", line 122, in get
    return _ForkingPickler.loads(res)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/core/declarations/contract.py", line 1396, in __hash__
    return self._id
AttributeError: 'Contract' object has no attribute '_id'

Signature Malleability Found in DepositContract.deposit(bytes,bytes,bytes,bytes32) (contracts/ethStaking/test/DepositContract.sol#89-145)State Variable:DepositContract.branch (contracts/ethStaking/test/DepositContract.sol#63)Potential Signature Variable:signature
Reference: https://swcregistry.io/docs/SWC-117

state variable: Router.s_offRamps (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#63) not initialized and not written in contract but be used in contract
Reference:  

state variable: SDLPool.lastLockId (contracts/core/sdlPool/base/SDLPool.sol#35) not initialized and not written in contract but be used in contract
state variable: SDLPool.locks (contracts/core/sdlPool/base/SDLPool.sol#36) not initialized and not written in contract but be used in contract
state variable: SDLPool.totalEffectiveBalance (contracts/core/sdlPool/base/SDLPool.sol#40) not initialized and not written in contract but be used in contract
state variable: OperatorController.totalAssignedValidators (contracts/ethStaking/base/OperatorController.sol#40) not initialized and not written in contract but be used in contract
Reference:  

Arbitrary Storage Found in SDLPoolSecondary.handleOutgoingUpdate() (contracts/core/sdlPool/SDLPoolSecondary.sol#313-328)
Reference: https://swcregistry.io/docs/SWC-124

	- Operator.withdraw(address,uint256) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#295-302)
Reference:  

	- Router.recoverTokens(address,address,uint256) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#303-312)
Reference:  

	- RESDLTokenBridge.transferRESDL(uint64,address,uint256,bool,uint256) (contracts/core/ccip/RESDLTokenBridge.sol#84-135)
	- DistributionOracle.withdrawLink(uint256) (contracts/core/priorityPool/DistributionOracle.sol#217-220)
	- EthStakingStrategy.deposit(uint256) (contracts/ethStaking/EthStakingStrategy.sol#212-217)
	- NWLOperatorController.removeKeyPairs(uint256,uint256,uint256[]) (contracts/ethStaking/NWLOperatorController.sol#100-148)
	- NWLOperatorController.assignNextValidators(uint256) (contracts/ethStaking/NWLOperatorController.sol#180-241)
	- RewardsReceiver.withdraw() (contracts/ethStaking/RewardsReceiver.sol#36-57)
Reference:  

Contract locking ether found:
	Contract DepositContract (contracts/ethStaking/test/DepositContract.sol#58-164) has payable functions:
	 - IDepositContract.deposit(bytes,bytes,bytes,bytes32) (contracts/ethStaking/test/DepositContract.sol#27-32)
	 - DepositContract.deposit(bytes,bytes,bytes,bytes32) (contracts/ethStaking/test/DepositContract.sol#89-145)
	But does not have a function to withdraw the ether
Reference:  

Contract locking ether found:
	Contract WrappedNative (contracts/core/test/chainlink/WrappedNative.sol#6-12) has payable functions:
	 - WrappedNative.deposit() (contracts/core/test/chainlink/WrappedNative.sol#9-11)
	But does not have a function to withdraw the ether
Contract locking ether found:
	Contract EthStakingStrategyMock (contracts/ethStaking/test/EthStakingStrategyMock.sol#6-22) has payable functions:
	 - EthStakingStrategyMock.receive() (contracts/ethStaking/test/EthStakingStrategyMock.sol#9)
	But does not have a function to withdraw the ether
Reference:  

Potential price manipulation risk:
	- In function withdraw
		-- toWithdraw = balanceOf(_account) (contracts/core/StakingPool.sol#104) have potential price manipulated risk from toWithdraw and call None which could influence variable:toWithdraw
Potential price manipulation risk:
	- In function _updateReward
		-- toMint = balanceOf(_account) - super.balanceOf(_account) (contracts/core/test/deprecated/RewardsPoolV1.sol#79) have potential price manipulated risk from toMint and call None which could influence variable:toMint
	- In function withdraw
		-- toWithdraw = balanceOf(_account) (contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#51) have potential price manipulated risk from toWithdraw and call None which could influence variable:toWithdraw
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

Operator.distributeFunds(address[],uint256[]) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#349-358) use transfer in a loop: receivers[i].transfer(sendAmount) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#355)
Reference:  

DepositContract.get_deposit_root().node (contracts/ethStaking/test/DepositContract.sol#75) is a local variable never initialized
Reference:  

VaultControllerStrategyUpgrade.checkUpkeep(bytes).firstNonFullVault (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#151) is a local variable never initialized
SDLPoolPrimary.handleIncomingUpdate(uint256,int256).mintStartIndex (contracts/core/sdlPool/SDLPoolPrimary.sol#236) is a local variable never initialized
EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]).nwlPubkeys (contracts/ethStaking/EthStakingStrategy.sol#150) is a local variable never initialized
EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]).nwlSignatures (contracts/ethStaking/EthStakingStrategy.sol#151) is a local variable never initialized
OperatorController.getAssignedKeys(uint256,uint256).index (contracts/ethStaking/base/OperatorController.sol#157) is a local variable never initialized
StakingPool._updateStrategyRewards(uint256[],bytes).totalFeeAmounts (contracts/core/StakingPool.sol#426) is a local variable never initialized
SDLPoolCCIPControllerPrimary._distributeRewards(uint64,address[],uint256[]).numRewardTokensToTransfer (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#252) is a local variable never initialized
RewardsInitiator.checkUpkeep(bytes).totalStrategiesToUpdate (contracts/core/RewardsInitiator.sol#52) is a local variable never initialized
SDLPool.getLockIdsByOwner(address).lockIdsFound (contracts/core/sdlPool/base/SDLPool.sol#180) is a local variable never initialized
LiquidSDIndexPool.getWithdrawalAmounts(uint256).totalTargetDepositDiffs (contracts/liquidSDIndex/LiquidSDIndexPool.sol#218) is a local variable never initialized
StakingPool._updateStrategyRewards(uint256[],bytes).totalFeeCount (contracts/core/StakingPool.sol#427) is a local variable never initialized
EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]).wlPubkeys (contracts/ethStaking/EthStakingStrategy.sol#168) is a local variable never initialized
SDLPoolCCIPControllerPrimary._distributeRewards(uint64,address[],uint256[]).tokensAdded (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#263) is a local variable never initialized
VaultControllerStrategyUpgrade.getPendingFees().totalFees (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#211) is a local variable never initialized
RewardsInitiator.checkUpkeep(bytes).strategiesAdded (contracts/core/RewardsInitiator.sol#64) is a local variable never initialized
NWLOperatorController.reportStoppedValidators(uint256[],uint256[],uint256[]).totalNewlyStoppedValidators (contracts/ethStaking/NWLOperatorController.sol#303) is a local variable never initialized
WLOperatorController.reportStoppedValidators(uint256[],uint256[]).totalNewlyStoppedValidators (contracts/ethStaking/WLOperatorController.sol#375) is a local variable never initialized
LiquidSDIndexPool.removeLSDToken(address,uint256[]).index (contracts/liquidSDIndex/LiquidSDIndexPool.sol#362) is a local variable never initialized
LiquidSDIndexPool.updateRewards().totalFeeAmounts (contracts/liquidSDIndex/LiquidSDIndexPool.sol#295) is a local variable never initialized
WLOperatorController.getNextValidators(uint256).operatorCount (contracts/ethStaking/WLOperatorController.sol#297) is a local variable never initialized
LiquidSDIndexPool.getRewards().totalFees (contracts/liquidSDIndex/LiquidSDIndexPool.sol#274) is a local variable never initialized
StakingPool._updateStrategyRewards(uint256[],bytes).feesPaidCount (contracts/core/StakingPool.sol#472) is a local variable never initialized
WLOperatorController.assignNextValidators(uint256[],uint256[],uint256).totalValidatorCount (contracts/ethStaking/WLOperatorController.sol#141) is a local variable never initialized
VaultControllerStrategy.getPendingFees().totalFees (contracts/linkStaking/base/VaultControllerStrategy.sol#140) is a local variable never initialized
OperatorVault.updateDeposits(uint256,address).opRewards (contracts/linkStaking/OperatorVault.sol#185) is a local variable never initialized
WLOperatorController.assignNextValidators(uint256[],uint256[],uint256).maxBatches (contracts/ethStaking/WLOperatorController.sol#142) is a local variable never initialized
WLOperatorController.getNextValidators(uint256).loopValidatorCount (contracts/ethStaking/WLOperatorController.sol#300) is a local variable never initialized
EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]).wlSignatures (contracts/ethStaking/EthStakingStrategy.sol#169) is a local variable never initialized
WLOperatorController.assignNextValidators(uint256[],uint256[],uint256).maxBatchOperatorId (contracts/ethStaking/WLOperatorController.sol#143) is a local variable never initialized
Reference:  

Chainlink.initialize(Chainlink.Request,bytes32,address,bytes4) (node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#33-44)have ignores return value in BufferChainlink.init(self.buf,defaultBufferSize) (node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#39)
Chainlink.setBuffer(Chainlink.Request,bytes) (node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#52-55)have ignores return value in BufferChainlink.init(self.buf,data.length) (node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#53)
CBORChainlink.encodeFixedNumeric(BufferChainlink.buffer,uint8,uint64) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#21-37)have ignores return value in buf.appendUint8(uint8((major << 5) | value)) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#23)
CBORChainlink.encodeIndefiniteLengthType(BufferChainlink.buffer,uint8) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#39-41)have ignores return value in buf.appendUint8(uint8((major << 5) | 31)) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#40)
CBORChainlink.encodeBytes(BufferChainlink.buffer,bytes) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#63-66)have ignores return value in buf.append(value) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#65)
CBORChainlink.encodeBigNum(BufferChainlink.buffer,uint256) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#68-71)have ignores return value in buf.appendUint8(uint8((MAJOR_TYPE_TAG << 5) | TAG_TYPE_BIGNUM)) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#69)
CBORChainlink.encodeSignedBigNum(BufferChainlink.buffer,int256) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#73-76)have ignores return value in buf.appendUint8(uint8((MAJOR_TYPE_TAG << 5) | TAG_TYPE_NEGATIVE_BIGNUM)) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#74)
CBORChainlink.encodeString(BufferChainlink.buffer,string) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#78-81)have ignores return value in buf.append(bytes(value)) (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#80)
ERC1967UpgradeUpgradeable._upgradeToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#65-70)have ignores return value in AddressUpgradeable.functionDelegateCall(newImplementation,data) (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#68)
ERC1967UpgradeUpgradeable._upgradeBeaconToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#156-162)have ignores return value in AddressUpgradeable.functionDelegateCall(IBeaconUpgradeable(newBeacon).implementation(),data) (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#160)
ERC1967Upgrade._upgradeToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#65-74)have ignores return value in Address.functionDelegateCall(newImplementation,data) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#72)
ERC1967Upgrade._upgradeBeaconToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#174-184)have ignores return value in Address.functionDelegateCall(IBeacon(newBeacon).implementation(),data) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#182)
Reference:  

RewardsPoolWSD.distributeRewards() (contracts/core/RewardsPoolWSD.sol#49-60)have ignores return value in token.transferAndCall(address(wsdToken),balance,0x) (contracts/core/RewardsPoolWSD.sol#53)
SDLPoolCCIPControllerPrimary.distributeRewards() (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#56-93)have ignores return value in IERC677(token).transferAndCall(wrappedToken,tokenBalance,) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#74)
WrappedTokenBridge.constructor(address,address,address,address) (contracts/core/ccip/WrappedTokenBridge.sol#60-74)have ignores return value in linkToken.approve(_router,type()(uint256).max) (contracts/core/ccip/WrappedTokenBridge.sol#71)
SDLPoolCCIPController.constructor(address,address,address,address,uint256) (contracts/core/ccip/base/SDLPoolCCIPController.sol#41-54)have ignores return value in linkToken.approve(_router,type()(uint256).max) (contracts/core/ccip/base/SDLPoolCCIPController.sol#52)
DelegatorPool.retireDelegatorPool(address[],address) (contracts/core/test/deprecated/DelegatorPool.sol#249-276)have ignores return value in allowanceToken.approve(_sdlPool,type()(uint256).max) (contracts/core/test/deprecated/DelegatorPool.sol#251)
OperatorController.onTokenTransfer(address,uint256,bytes) (contracts/ethStaking/base/OperatorController.sol#176-183)have ignores return value in sdToken.transferAndCall(address(rewardsPool),_value,0x) (contracts/ethStaking/base/OperatorController.sol#182)
EthStakingStrategyMock.depositEther(uint256) (contracts/ethStaking/test/EthStakingStrategyMock.sol#11-13)have ignores return value in INWLOperatorController(nwlOperatorController).assignNextValidators(_totalValidatorCount) (contracts/ethStaking/test/EthStakingStrategyMock.sol#12)
OperatorVault.deposit(uint256) (contracts/linkStaking/OperatorVault.sol#95-99)have ignores return value in IERC677(address(token)).transferAndCall(address(stakeController),_amount,) (contracts/linkStaking/OperatorVault.sol#98)
Vault.deposit(uint256) (contracts/linkStaking/base/Vault.sol#62-65)have ignores return value in IERC677(address(token)).transferAndCall(address(stakeController),_amount,) (contracts/linkStaking/base/Vault.sol#64)
OperatorVCSMock.addVault(address) (contracts/linkStaking/test/OperatorVCSMock.sol#47-50)have ignores return value in token.approve(_vault,type()(uint256).max) (contracts/linkStaking/test/OperatorVCSMock.sol#49)
StakingMockV1.migrate(bytes) (contracts/linkStaking/test/StakingMockV1.sol#75-77)have ignores return value in token.transferAndCall(migration,stakedBalances[msg.sender] + baseReward + delegationReward,abi.encode(msg.sender)) (contracts/linkStaking/test/StakingMockV1.sol#76)
VCSMock.addVaults(address[]) (contracts/linkStaking/test/VCSMock.sol#26-32)have ignores return value in token.approve(vault,type()(uint256).max) (contracts/linkStaking/test/VCSMock.sol#30)
VaultV1.deposit(uint256) (contracts/linkStaking/test/deprecated/VaultV1.sol#73-76)have ignores return value in IERC677(address(token)).transferAndCall(address(stakeController),_amount,0x00) (contracts/linkStaking/test/deprecated/VaultV1.sol#75)
OperatorVCSUpgrade.initialize(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[],address[]) (contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#21-43)have ignores return value in token.approve(vault,type()(uint256).max) (contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#41)
Reference:  

function:ChainlinkClient.validateChainlinkCallback(bytes32) (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#286-291)is empty 
function:ContextUpgradeable.__Context_init() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#18-19)is empty 
function:ContextUpgradeable.__Context_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#21-22)is empty 
function:ERC1967UpgradeUpgradeable.__ERC1967Upgrade_init() (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#20-21)is empty 
function:ERC1967UpgradeUpgradeable.__ERC1967Upgrade_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#23-24)is empty 
function:UUPSUpgradeable.__UUPSUpgradeable_init() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#23-24)is empty 
function:UUPSUpgradeable.__UUPSUpgradeable_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#26-27)is empty 
function:ERC20Upgradeable._beforeTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#353)is empty 
function:ERC20Upgradeable._afterTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#369)is empty 
function:ERC20._beforeTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#358-362)is empty 
function:ERC20._afterTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#378-382)is empty 
function:ERC721._beforeTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#432-436)is empty 
function:ERC721._afterTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#449-453)is empty 
Reference:  

function:StakingRewardsPool._authorizeUpgrade(address) (contracts/core/base/StakingRewardsPool.sol#209)is empty 
function:RewardsPoolController._authorizeUpgrade(address) (contracts/core/base/RewardsPoolController.sol#197)is empty 
function:Strategy._authorizeUpgrade(address) (contracts/core/base/Strategy.sol#84)is empty 
function:PriorityPool._authorizeUpgrade(address) (contracts/core/priorityPool/PriorityPool.sol#572)is empty 
function:CCIPTokenPoolMock.lockOrBurn() (contracts/core/test/chainlink/CCIPTokenPoolMock.sol#21)is empty 
function:RewardsPoolControllerV1._authorizeUpgrade(address) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#221)is empty 
function:OperatorController._authorizeUpgrade(address) (contracts/ethStaking/base/OperatorController.sol#434)is empty 
function:Vault._authorizeUpgrade(address) (contracts/linkStaking/base/Vault.sol#102)is empty 
function:VaultV1._authorizeUpgrade(address) (contracts/linkStaking/test/deprecated/VaultV1.sol#107)is empty 
function:LSDIndexAdapter._authorizeUpgrade(address) (contracts/liquidSDIndex/base/LSDIndexAdapter.sol#71)is empty 
Reference:  

	- RESDLTokenBridge.setExtraArgs(uint64,bytes) (contracts/core/ccip/RESDLTokenBridge.sol#161-164)
	- Strategy.__Strategy_init(address,address) (contracts/core/base/Strategy.sol#20-25)
	- EthStakingStrategy.reportBeaconState(uint256,uint256,uint256) (contracts/ethStaking/EthStakingStrategy.sol#96-124)
	- EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]) (contracts/ethStaking/EthStakingStrategy.sol#135-206)
	- EthStakingStrategy.updateDeposits(bytes) (contracts/ethStaking/EthStakingStrategy.sol#246-280)
	- OperatorController.__OperatorController_init(address,address) (contracts/ethStaking/base/OperatorController.sol#74-80)
	- OperatorController.initiateKeyPairValidation(address,uint256) (contracts/ethStaking/base/OperatorController.sol#205-213)
	- OperatorController.setOperatorName(uint256,string) (contracts/ethStaking/base/OperatorController.sol#219-222)
	- OperatorController.setOperatorOwner(uint256,address) (contracts/ethStaking/base/OperatorController.sol#231-241)
	- OperatorController.disableOperator(uint256) (contracts/ethStaking/base/OperatorController.sol#248-265)
	- OperatorController.setRewardsPool(address) (contracts/ethStaking/base/OperatorController.sol#289-292)
	- NWLOperatorController.reportKeyPairValidation(uint256,bool) (contracts/ethStaking/NWLOperatorController.sol#155-172)
	- NWLOperatorController.reportStoppedValidators(uint256[],uint256[],uint256[]) (contracts/ethStaking/NWLOperatorController.sol#293-342)
	- RewardsReceiver.setWithdrawalLimits(uint256,uint256) (contracts/ethStaking/RewardsReceiver.sol#64-69)
Reference:  

Setter function Operator.transferOwnableContracts(address[],address) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#235-240) does not emit an event
Setter function Operator.acceptAuthorizedReceivers(address[],address[]) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#281-287) does not emit an event
Setter function Operator.withdraw(address,uint256) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#295-302) does not emit an event
Setter function Operator.ownerForward(address,bytes) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#319-323) does not emit an event
Setter function Operator.ownerTransferAndCall(address,uint256,bytes) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#332-338) does not emit an event
Setter function Operator._verifyAndProcessOracleRequest(address,uint256,address,bytes4,uint256,uint256) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#444-460) does not emit an event
Setter function Operator._verifyOracleRequestAndProcessPayment(bytes32,uint256,address,bytes4,uint256,uint256) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#470-483) does not emit an event
Setter function Operator._canSetAuthorizedSenders() (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#525-527) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Setter function Router.ccipSend(uint64,Client.EVM2AnyMessage) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#104-141) does not emit an event
Setter function Router.setWrappedNative(address) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#228-230) does not emit an event
Setter function Router.recoverTokens(address,address,uint256) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#303-312) does not emit an event
Setter function Router.onlyOffRamp(uint64) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#320-324) does not emit an event
Setter function CCIPReceiver.ccipReceive(Client.Any2EVMMessage) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#27-29) does not emit an event
Setter function CCIPReceiver.onlyRouter() (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#48-51) does not emit an event
Setter function EnumerableMap.set(EnumerableMap.Bytes32ToBytes32Map,bytes32,bytes32) (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#73-80) does not emit an event
Setter function EnumerableMap.set(EnumerableMap.UintToUintMap,uint256,uint256) (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#176-182) does not emit an event
Setter function EnumerableMap.set(EnumerableMap.UintToAddressMap,uint256,address) (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#268-274) does not emit an event
Setter function EnumerableMap.set(EnumerableMap.AddressToUintMap,address,uint256) (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#360-366) does not emit an event
Setter function EnumerableMap.set(EnumerableMap.Bytes32ToUintMap,bytes32,uint256) (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#452-458) does not emit an event
Setter function Chainlink.setBuffer(Chainlink.Request,bytes) (node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#52-55) does not emit an event
Setter function ChainlinkClient.sendChainlinkRequestTo(address,Chainlink.Request,uint256) (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#91-110) does not emit an event
Setter function ChainlinkClient.sendOperatorRequestTo(address,Chainlink.Request,uint256) (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#135-153) does not emit an event
Setter function ChainlinkClient.setChainlinkOracle(address) (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#210-212) does not emit an event
Setter function ChainlinkClient.setChainlinkToken(address) (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#218-220) does not emit an event
Setter function ChainlinkClient.setPublicChainlinkToken() (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#226-228) does not emit an event
Setter function ChainlinkClient.addChainlinkExternalRequest(address,bytes32) (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#252-254) does not emit an event
Setter function ChainlinkClient.useChainlinkWithENS(address,bytes32) (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#262-269) does not emit an event
Setter function ChainlinkClient.validateChainlinkCallback(bytes32) (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#286-291) does not emit an event
Setter function OwnableUpgradeable.__Ownable_init() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#29-31) does not emit an event
Setter function OwnableUpgradeable.__Ownable_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#33-35) does not emit an event
Setter function OwnableUpgradeable._checkOwner() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#55-57) does not emit an event
Setter function OwnableUpgradeable.renounceOwnership() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#66-68) does not emit an event
Setter function OwnableUpgradeable.transferOwnership(address) (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#74-77) does not emit an event
Setter function OwnableUpgradeable.onlyOwner() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#40-43) does not emit an event
Setter function PausableUpgradeable.__Pausable_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#38-40) does not emit an event
Setter function ERC20Upgradeable.__ERC20_init_unchained(string,string) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#59-62) does not emit an event
Setter function ERC20Upgradeable.transfer(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#118-122) does not emit an event
Setter function ERC20Upgradeable.approve(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#141-145) does not emit an event
Setter function ERC20Upgradeable.transferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#163-168) does not emit an event
Setter function ERC20Upgradeable.increaseAllowance(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#182-186) does not emit an event
Setter function ERC20Upgradeable.decreaseAllowance(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#202-211) does not emit an event
Setter function ReentrancyGuard.nonReentrant() (node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol#50-62) does not emit an event
Setter function ERC20.transfer(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#113-117) does not emit an event
Setter function ERC20.approve(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#136-140) does not emit an event
Setter function ERC20.transferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#158-167) does not emit an event
Setter function ERC20.increaseAllowance(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#181-185) does not emit an event
Setter function ERC20.decreaseAllowance(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#201-210) does not emit an event
Setter function ERC721URIStorage._setTokenURI(uint256,string) (node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#45-48) does not emit an event
Setter function ERC721URIStorage._burn(uint256) (node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#55-61) does not emit an event
Setter function StakedotlinkCouncil.mint(address,uint256) (contracts/governance/StakedotlinkCouncil.sol#90-92) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Setter function MerkleDistributor.addDistributions(address[],bytes32[],uint256[],uint256[]) (contracts/airdrop/MerkleDistributor.sol#54-68) does not emit an event
Setter function MerkleDistributor.updateDistributions(address[],bytes32[],uint256[],uint256[]) (contracts/airdrop/MerkleDistributor.sol#102-116) does not emit an event
Setter function MerkleDistributor.withdrawUnclaimedTokens(address,bytes32,uint256) (contracts/airdrop/MerkleDistributor.sol#208-224) does not emit an event
Setter function MerkleDistributor.pauseForWithdrawal(address) (contracts/airdrop/MerkleDistributor.sol#232-239) does not emit an event
Setter function RewardsInitiator.updateRewards(uint256[],bytes) (contracts/core/RewardsInitiator.sol#38-42) does not emit an event
Setter function RewardsPoolWSD.updateReward(address) (contracts/core/RewardsPoolWSD.sol#66-72) does not emit an event
Setter function StakingPool.initialize(address,string,string,StakingPool.Fee[]) (contracts/core/StakingPool.sol#41-52) does not emit an event
Setter function StakingPool.deposit(address,uint256) (contracts/core/StakingPool.sol#80-88) does not emit an event
Setter function StakingPool.withdraw(address,address,uint256) (contracts/core/StakingPool.sol#97-116) does not emit an event
Setter function StakingPool.strategyDeposit(uint256,uint256) (contracts/core/StakingPool.sol#123-126) does not emit an event
Setter function StakingPool.strategyWithdraw(uint256,uint256) (contracts/core/StakingPool.sol#133-136) does not emit an event
Setter function StakingPool.addStrategy(address) (contracts/core/StakingPool.sol#226-230) does not emit an event
Setter function StakingPool.removeStrategy(uint256,bytes) (contracts/core/StakingPool.sol#237-255) does not emit an event
Setter function StakingPool.reorderStrategies(uint256[]) (contracts/core/StakingPool.sol#261-274) does not emit an event
Setter function StakingPool.addFee(address,uint256) (contracts/core/StakingPool.sol#281-284) does not emit an event
Setter function StakingPool.updateFee(uint256,address,uint256) (contracts/core/StakingPool.sol#292-308) does not emit an event
Setter function StakingPool.updateStrategyRewards(uint256[],bytes) (contracts/core/StakingPool.sol#345-348) does not emit an event
Setter function StakingPool.setPriorityPool(address) (contracts/core/StakingPool.sol#375-377) does not emit an event
Setter function StakingPool.setRewardsInitiator(address) (contracts/core/StakingPool.sol#384-386) does not emit an event
Setter function StakingPool.onlyPriorityPool() (contracts/core/StakingPool.sol#54-57) does not emit an event
Setter function RESDLTokenBridge.onlySDLPoolCCIPController() (contracts/core/ccip/RESDLTokenBridge.sol#71-74) does not emit an event
Setter function SDLPoolCCIPControllerPrimary.distributeRewards() (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#56-93) does not emit an event
Setter function SDLPoolCCIPControllerPrimary.handleOutgoingRESDL(uint64,address,uint256) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#103-116) does not emit an event
Setter function SDLPoolCCIPControllerPrimary.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#125-134) does not emit an event
Setter function SDLPoolCCIPControllerPrimary.approveRewardTokens(address[]) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#190-195) does not emit an event
Setter function SDLPoolCCIPControllerPrimary.setRewardsInitiator(address) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#234-236) does not emit an event
Setter function SDLPoolCCIPControllerPrimary.onlyRewardsInitiator() (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#48-51) does not emit an event
Setter function SDLPoolCCIPControllerSecondary.performUpkeep(bytes) (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#66-71) does not emit an event
Setter function SDLPoolCCIPControllerSecondary.handleOutgoingRESDL(uint64,address,uint256) (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#80-86) does not emit an event
Setter function SDLPoolCCIPControllerSecondary.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken) (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#94-102) does not emit an event
Setter function WrappedTokenBridge.onTokenTransfer(address,uint256,bytes) (contracts/core/ccip/WrappedTokenBridge.sol#83-96) does not emit an event
Setter function WrappedTokenBridge.transferTokens(uint64,address,uint256,bool,uint256) (contracts/core/ccip/WrappedTokenBridge.sol#106-117) does not emit an event
Setter function WrappedTokenBridge.recoverTokens(address[],address) (contracts/core/ccip/WrappedTokenBridge.sol#140-147) does not emit an event
Setter function DistributionOracle.pauseForUpdate() (contracts/core/priorityPool/DistributionOracle.sol#142-144) does not emit an event
Setter function DistributionOracle.requestUpdate() (contracts/core/priorityPool/DistributionOracle.sol#151-153) does not emit an event
Setter function DistributionOracle.fulfillRequest(bytes32,bytes32,bytes32,uint256,uint256) (contracts/core/priorityPool/DistributionOracle.sol#163-177) does not emit an event
Setter function DistributionOracle.executeManualVerification() (contracts/core/priorityPool/DistributionOracle.sol#182-192) does not emit an event
Setter function DistributionOracle.rejectManualVerificationAndRetry() (contracts/core/priorityPool/DistributionOracle.sol#197-201) does not emit an event
Setter function DistributionOracle.cancelRequest(bytes32,uint256) (contracts/core/priorityPool/DistributionOracle.sol#208-211) does not emit an event
Setter function DistributionOracle.withdrawLink(uint256) (contracts/core/priorityPool/DistributionOracle.sol#217-220) does not emit an event
Setter function DistributionOracle._pauseForUpdate() (contracts/core/priorityPool/DistributionOracle.sol#263-268) does not emit an event
Setter function DistributionOracle._requestUpdate() (contracts/core/priorityPool/DistributionOracle.sol#275-288) does not emit an event
Setter function PriorityPool.initialize(address,address,address,uint128,uint128) (contracts/core/priorityPool/PriorityPool.sol#93-110) does not emit an event
Setter function PriorityPool.deposit(uint256,bool) (contracts/core/priorityPool/PriorityPool.sol#209-213) does not emit an event
Setter function PriorityPool.pauseForUpdate() (contracts/core/priorityPool/PriorityPool.sol#429-431) does not emit an event
Setter function PriorityPool.setDistributionOracle(address) (contracts/core/priorityPool/PriorityPool.sol#468-470) does not emit an event
Setter function PriorityPool._withdraw(uint256) (contracts/core/priorityPool/PriorityPool.sol#519-534) does not emit an event
Setter function PriorityPool._authorizeUpgrade(address) (contracts/core/priorityPool/PriorityPool.sol#572) does not emit an event
Setter function PriorityPool.onlyDistributionOracle() (contracts/core/priorityPool/PriorityPool.sol#115-118) does not emit an event
Setter function SDLPoolPrimary.initialize(string,string,address,address) (contracts/core/sdlPool/SDLPoolPrimary.sol#30-41) does not emit an event
Setter function SDLPoolPrimary.onTokenTransfer(address,uint256,bytes) (contracts/core/sdlPool/SDLPoolPrimary.sol#61-80) does not emit an event
Setter function SDLPoolPrimary.extendLockDuration(uint256,uint64) (contracts/core/sdlPool/SDLPoolPrimary.sol#91-94) does not emit an event
Setter function SDLPoolPrimary.migrate(address,uint256,uint64) (contracts/core/sdlPool/SDLPoolPrimary.sol#264-272) does not emit an event
Setter function SDLPoolSecondary.initialize(string,string,address,address,uint256) (contracts/core/sdlPool/SDLPoolSecondary.sol#66-79) does not emit an event
Setter function SDLPoolSecondary.onTokenTransfer(address,uint256,bytes) (contracts/core/sdlPool/SDLPoolSecondary.sol#132-151) does not emit an event
Setter function SDLPoolSecondary.extendLockDuration(uint256,uint64) (contracts/core/sdlPool/SDLPoolSecondary.sol#163-166) does not emit an event
Setter function SDLPoolSecondary.executeQueuedOperations(uint256[]) (contracts/core/sdlPool/SDLPoolSecondary.sol#247-250) does not emit an event
Setter function CurveMock.exchange(int128,int128,uint256,uint256,address) (contracts/core/test/CurveMock.sol#26-56) does not emit an event
Setter function ERC677ReceiverMock.onTokenTransfer(address,uint256,bytes) (contracts/core/test/ERC677ReceiverMock.sol#11-17) does not emit an event
Setter function ERC721ReceiverMock.onERC721Received(address,address,uint256,bytes) (contracts/core/test/ERC721ReceiverMock.sol#24-32) does not emit an event
Setter function PriorityPoolMock.updateDistribution(bytes32,bytes32,uint256,uint256) (contracts/core/test/PriorityPoolMock.sol#22-34) does not emit an event
Setter function PriorityPoolMock.pauseForUpdate() (contracts/core/test/PriorityPoolMock.sol#36-38) does not emit an event
Setter function PriorityPoolMock.setDepositsSinceLastUpdate(uint256) (contracts/core/test/PriorityPoolMock.sol#40-42) does not emit an event
Setter function RewardsPoolControllerMock.initialize(address) (contracts/core/test/RewardsPoolControllerMock.sol#23-26) does not emit an event
Setter function RewardsPoolControllerMock.stake(uint256) (contracts/core/test/RewardsPoolControllerMock.sol#36-40) does not emit an event
Setter function RewardsPoolControllerMock.withdraw(uint256) (contracts/core/test/RewardsPoolControllerMock.sol#42-46) does not emit an event
Setter function SDLDependentMock.updateSDLBalance(address,uint256) (contracts/core/test/SDLDependentMock.sol#11-13) does not emit an event
Setter function SDLPoolCCIPControllerMock.handleOutgoingRESDL(uint64,address,uint256) (contracts/core/test/SDLPoolCCIPControllerMock.sol#29-35) does not emit an event
Setter function SDLPoolCCIPControllerMock.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken) (contracts/core/test/SDLPoolCCIPControllerMock.sol#37-45) does not emit an event
Setter function SDLPoolCCIPControllerMock.distributeRewards() (contracts/core/test/SDLPoolCCIPControllerMock.sol#47-49) does not emit an event
Setter function SDLPoolCCIPControllerMock.setRESDLTokenBridge(address) (contracts/core/test/SDLPoolCCIPControllerMock.sol#51-53) does not emit an event
Setter function SDLPoolCCIPControllerMock.onlyBridge() (contracts/core/test/SDLPoolCCIPControllerMock.sol#19-22) does not emit an event
Setter function SDLPoolMock.setEffectiveBalance(address,uint256) (contracts/core/test/SDLPoolMock.sol#15-17) does not emit an event
Setter function WrappedSDTokenMock.onTokenTransfer(address,uint256,bytes) (contracts/core/test/WrappedSDTokenMock.sol#24-32) does not emit an event
Setter function WrappedSDTokenMock.unwrap(uint256) (contracts/core/test/WrappedSDTokenMock.sol#38-43) does not emit an event
Setter function WrappedSDTokenMock.setMultiplier(uint256) (contracts/core/test/WrappedSDTokenMock.sol#63-65) does not emit an event
Setter function CCIPOffRampMock.executeSingleMessage(bytes32,uint64,bytes,address,Client.EVMTokenAmount[]) (contracts/core/test/chainlink/CCIPOffRampMock.sol#34-52) does not emit an event
Setter function CCIPOffRampMock.setTokenPool(address,address) (contracts/core/test/chainlink/CCIPOffRampMock.sol#54-56) does not emit an event
Setter function CCIPOnRampMock.forwardFromRouter(Client.EVM2AnyMessage,uint256,address) (contracts/core/test/chainlink/CCIPOnRampMock.sol#49-57) does not emit an event
Setter function CCIPOnRampMock.setTokenPool(address,address) (contracts/core/test/chainlink/CCIPOnRampMock.sol#59-61) does not emit an event
Setter function WrappedNative.deposit() (contracts/core/test/chainlink/WrappedNative.sol#9-11) does not emit an event
Setter function DelegatorPool.initialize(address,string,string,address[]) (contracts/core/test/deprecated/DelegatorPool.sol#56-74) does not emit an event
Setter function DelegatorPool.onTokenTransfer(address,uint256,bytes) (contracts/core/test/deprecated/DelegatorPool.sol#82-104) does not emit an event
Setter function DelegatorPool.staked(address) (contracts/core/test/deprecated/DelegatorPool.sol#124-126) does not emit an event
Setter function DelegatorPool.totalStaked() (contracts/core/test/deprecated/DelegatorPool.sol#143-145) does not emit an event
Setter function DelegatorPool.setLockedApproval(address,uint256) (contracts/core/test/deprecated/DelegatorPool.sol#200-203) does not emit an event
Setter function DelegatorPool.burnLockedBalance(address,uint256) (contracts/core/test/deprecated/DelegatorPool.sol#210-224) does not emit an event
Setter function DelegatorPool.setPoolRouter(address) (contracts/core/test/deprecated/DelegatorPool.sol#230-233) does not emit an event
Setter function DelegatorPool.setCommunityPool(address,bool) (contracts/core/test/deprecated/DelegatorPool.sol#240-243) does not emit an event
Setter function OwnersRewardsPoolV1.withdraw(uint256) (contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#33-42) does not emit an event
Setter function OwnersRewardsPoolV1.onTokenTransfer(address,uint256,bytes) (contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#65-72) does not emit an event
Setter function PoolAllowanceV1.mintAllowance(address,uint256) (contracts/core/test/deprecated/PoolAllowanceV1.sol#31-33) does not emit an event
Setter function PoolAllowanceV1.burnAllowance(address,uint256) (contracts/core/test/deprecated/PoolAllowanceV1.sol#40-42) does not emit an event
Setter function PoolAllowanceV1.onlyPoolOwners() (contracts/core/test/deprecated/PoolAllowanceV1.sol#21-24) does not emit an event
Setter function PoolOwnersV1.onTokenTransfer(address,uint256,bytes) (contracts/core/test/deprecated/PoolOwnersV1.sol#94-101) does not emit an event
Setter function PoolOwnersV1.stake(uint256) (contracts/core/test/deprecated/PoolOwnersV1.sol#107-110) does not emit an event
Setter function PoolOwnersV1.exit() (contracts/core/test/deprecated/PoolOwnersV1.sol#137-140) does not emit an event
Setter function PoolOwnersV1._mintAllowance(address) (contracts/core/test/deprecated/PoolOwnersV1.sol#212-222) does not emit an event
Setter function PoolOwnersV1._burnAllowance(address) (contracts/core/test/deprecated/PoolOwnersV1.sol#228-238) does not emit an event
Setter function LinkPoolNFT.mint(address) (contracts/core/tokens/LinkPoolNFT.sol#26-31) does not emit an event
Setter function LinkPoolNFT.setBaseURI(string) (contracts/core/tokens/LinkPoolNFT.sol#37-39) does not emit an event
Setter function StakingAllowance.mint(address,uint256) (contracts/core/tokens/StakingAllowance.sol#20-22) does not emit an event
Setter function StakingAllowance.mintToContract(address,address,uint256,bytes) (contracts/core/tokens/StakingAllowance.sol#30-38) does not emit an event
Setter function StakingAllowance.burn(uint256) (contracts/core/tokens/StakingAllowance.sol#44-46) does not emit an event
Setter function StakingAllowance.burnFrom(address,uint256) (contracts/core/tokens/StakingAllowance.sol#52-55) does not emit an event
Setter function StakingAllowance.transferAndCallWithSender(address,address,uint256,bytes) (contracts/core/tokens/StakingAllowance.sol#65-75) does not emit an event
Setter function WrappedSDToken.onTokenTransfer(address,uint256,bytes) (contracts/core/tokens/WrappedSDToken.sol#27-34) does not emit an event
Setter function WrappedSDToken.wrap(uint256) (contracts/core/tokens/WrappedSDToken.sol#40-43) does not emit an event
Setter function WrappedSDToken.unwrap(uint256) (contracts/core/tokens/WrappedSDToken.sol#49-54) does not emit an event
Setter function DepositController.depositEther(bytes32,bytes32,bytes32,uint256,uint256,uint256[],uint256[]) (contracts/ethStaking/DepositController.sol#46-64) does not emit an event
Setter function EthStakingStrategy.initialize(address,address,uint256,uint256,address,bytes32,uint256) (contracts/ethStaking/EthStakingStrategy.sol#69-84) does not emit an event
Setter function EthStakingStrategy.deposit(uint256) (contracts/ethStaking/EthStakingStrategy.sol#212-217) does not emit an event
Setter function EthStakingStrategy.withdraw(uint256) (contracts/ethStaking/EthStakingStrategy.sol#224-226) does not emit an event
Setter function EthStakingStrategy.nwlWithdraw(address,uint256) (contracts/ethStaking/EthStakingStrategy.sol#234-237) does not emit an event
Setter function EthStakingStrategy.updateDeposits(bytes) (contracts/ethStaking/EthStakingStrategy.sol#246-280) does not emit an event
Setter function KeyValidationOracle.onTokenTransfer(address,uint256,bytes) (contracts/ethStaking/KeyValidationOracle.sol#48-59) does not emit an event
Setter function KeyValidationOracle.reportKeyPairValidation(bytes32,uint256,bool,bool) (contracts/ethStaking/KeyValidationOracle.sol#68-79) does not emit an event
Setter function NWLOperatorController.initialize(address,address) (contracts/ethStaking/NWLOperatorController.sol#36-38) does not emit an event
Setter function NWLOperatorController.addOperator(string) (contracts/ethStaking/NWLOperatorController.sol#72-74) does not emit an event
Setter function NWLOperatorController.addKeyPairs(uint256,uint256,bytes,bytes) (contracts/ethStaking/NWLOperatorController.sol#83-92) does not emit an event
Setter function NWLOperatorController.assignNextValidators(uint256) (contracts/ethStaking/NWLOperatorController.sol#180-241) does not emit an event
Setter function OperatorWhitelist.useWhitelist(address) (contracts/ethStaking/OperatorWhitelist.sol#41-46) does not emit an event
Setter function OperatorWhitelist.addWhitelistEntries(address[]) (contracts/ethStaking/OperatorWhitelist.sol#52-58) does not emit an event
Setter function OperatorWhitelist.removeWhitelistEntries(address[]) (contracts/ethStaking/OperatorWhitelist.sol#64-70) does not emit an event
Setter function OperatorWhitelist.setWLOperatorController(address) (contracts/ethStaking/OperatorWhitelist.sol#76-78) does not emit an event
Setter function WLOperatorController.initialize(address,address,address,uint256) (contracts/ethStaking/WLOperatorController.sol#33-42) does not emit an event
Setter function WLOperatorController.addOperator(string) (contracts/ethStaking/WLOperatorController.sol#48-51) does not emit an event
Setter function WLOperatorController.addKeyPairs(uint256,uint256,bytes,bytes) (contracts/ethStaking/WLOperatorController.sol#60-68) does not emit an event
Setter function WLOperatorController.assignNextValidators(uint256[],uint256[],uint256) (contracts/ethStaking/WLOperatorController.sol#125-272) does not emit an event
Setter function WLOperatorController.setBatchSize(uint256) (contracts/ethStaking/WLOperatorController.sol#410-412) does not emit an event
Setter function WLOperatorController.setOperatorWhitelist(address) (contracts/ethStaking/WLOperatorController.sol#418-420) does not emit an event
Setter function EthStakingStrategyMock.nwlWithdraw(address,uint256) (contracts/ethStaking/test/EthStakingStrategyMock.sol#15-17) does not emit an event
Setter function EthStakingStrategyMock.setNWLOperatorController(address) (contracts/ethStaking/test/EthStakingStrategyMock.sol#19-21) does not emit an event
Setter function GovernanceController.revokeRole(uint256,address) (contracts/governance/GovernanceController.sol#138-140) does not emit an event
Setter function GovernanceController.renounceRole(uint256) (contracts/governance/GovernanceController.sol#146-148) does not emit an event
Setter function CommunityVCS.initialize(address,address,address,address,VaultControllerStrategy.Fee[],uint256,uint128,uint128) (contracts/linkStaking/CommunityVCS.sol#36-57) does not emit an event
Setter function CommunityVCS.addVaults(uint256) (contracts/linkStaking/CommunityVCS.sol#106-108) does not emit an event
Setter function OperatorVCS.initialize(address,address,address,address,VaultControllerStrategy.Fee[],uint256,uint256) (contracts/linkStaking/OperatorVCS.sol#46-76) does not emit an event
Setter function OperatorVCS.onTokenTransfer(address,uint256,bytes) (contracts/linkStaking/OperatorVCS.sol#84-90) does not emit an event
Setter function OperatorVCS.withdrawOperatorRewards(address,uint256) (contracts/linkStaking/OperatorVCS.sol#102-113) does not emit an event
Setter function OperatorVCS.updateDeposits(bytes) (contracts/linkStaking/OperatorVCS.sol#155-216) does not emit an event
Setter function OperatorVCS.setOperator(uint256,address) (contracts/linkStaking/OperatorVCS.sol#251-253) does not emit an event
Setter function OperatorVCS.setRewardsReceiver(uint256,address) (contracts/linkStaking/OperatorVCS.sol#261-263) does not emit an event
Setter function OperatorVCS.togglePreRelease() (contracts/linkStaking/OperatorVCS.sol#287-289) does not emit an event
Setter function OperatorVault.initialize(address,address,address,address,address,address,address) (contracts/linkStaking/OperatorVault.sol#50-72) does not emit an event
Setter function OperatorVault.deposit(uint256) (contracts/linkStaking/OperatorVault.sol#95-99) does not emit an event
Setter function OperatorVault.updateDeposits(uint256,address) (contracts/linkStaking/OperatorVault.sol#175-200) does not emit an event
Setter function OperatorVault.setOperator(address) (contracts/linkStaking/OperatorVault.sol#211-215) does not emit an event
Setter function OperatorVault.onlyOperator() (contracts/linkStaking/OperatorVault.sol#77-80) does not emit an event
Setter function OperatorVault.onlyRewardsReceiver() (contracts/linkStaking/OperatorVault.sol#85-88) does not emit an event
Setter function CommunityVaultV2Mock.initializeV2(uint256) (contracts/linkStaking/test/CommunityVaultV2Mock.sol#13-15) does not emit an event
Setter function OperatorVCSMock.deposit(uint256) (contracts/linkStaking/test/OperatorVCSMock.sol#33-36) does not emit an event
Setter function OperatorVCSMock.addVault(address) (contracts/linkStaking/test/OperatorVCSMock.sol#47-50) does not emit an event
Setter function OperatorVCSMock.setWithdrawalPercentage(uint256) (contracts/linkStaking/test/OperatorVCSMock.sol#52-54) does not emit an event
Setter function PFAlertsControllerMock.raiseAlert(address) (contracts/linkStaking/test/PFAlertsControllerMock.sol#17-19) does not emit an event
Setter function StakingMock.onTokenTransfer(address,uint256,bytes) (contracts/linkStaking/test/StakingMock.sol#39-51) does not emit an event
Setter function StakingMock.removePrincipal(address,uint256) (contracts/linkStaking/test/StakingMock.sol#77-80) does not emit an event
Setter function StakingMock.setActive(bool) (contracts/linkStaking/test/StakingMock.sol#90-92) does not emit an event
Setter function StakingMock.setMaxPoolSize(uint256) (contracts/linkStaking/test/StakingMock.sol#94-96) does not emit an event
Setter function StakingMock.setDepositLimits(uint256,uint256) (contracts/linkStaking/test/StakingMock.sol#98-101) does not emit an event
Setter function StakingMockV1.onTokenTransfer(address,uint256,bytes) (contracts/linkStaking/test/StakingMockV1.sol#30-37) does not emit an event
Setter function StakingMockV1.setActive(bool) (contracts/linkStaking/test/StakingMockV1.sol#55-57) does not emit an event
Setter function StakingMockV1.setMigration(address) (contracts/linkStaking/test/StakingMockV1.sol#71-73) does not emit an event
Setter function StakingMockV1.migrate(bytes) (contracts/linkStaking/test/StakingMockV1.sol#75-77) does not emit an event
Setter function StakingMockV1.setBaseReward(uint256) (contracts/linkStaking/test/StakingMockV1.sol#79-81) does not emit an event
Setter function StakingMockV1.setDelegationReward(uint256) (contracts/linkStaking/test/StakingMockV1.sol#87-89) does not emit an event
Setter function StakingMockV1.setPaused(bool) (contracts/linkStaking/test/StakingMockV1.sol#99-101) does not emit an event
Setter function StakingMockV1.raiseAlert() (contracts/linkStaking/test/StakingMockV1.sol#107-109) does not emit an event
Setter function StakingRewardsMock.claimReward() (contracts/linkStaking/test/StakingRewardsMock.sol#27-32) does not emit an event
Setter function StakingRewardsMock.setReward(address,uint256) (contracts/linkStaking/test/StakingRewardsMock.sol#34-36) does not emit an event
Setter function VCSMock.initialize(address,address,address,address,VaultControllerStrategy.Fee[]) (contracts/linkStaking/test/VCSMock.sol#16-24) does not emit an event
Setter function VCSMock.addVaults(address[]) (contracts/linkStaking/test/VCSMock.sol#26-32) does not emit an event
Setter function OperatorVaultV1.initialize(address,address,address,address) (contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#22-30) does not emit an event
Setter function OperatorVaultV1.setOperator(address) (contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#61-64) does not emit an event
Setter function OperatorVaultV1.onlyOperator() (contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#32-35) does not emit an event
Setter function OperatorVCSUpgrade.initialize(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[],address[]) (contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#21-43) does not emit an event
Setter function OperatorVCSUpgrade.setOperator(uint256,address) (contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#92-94) does not emit an event
Setter function LiquidSDIndexPool.initialize(string,string,uint256,uint256,LiquidSDIndexPool.Fee[],uint256) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#56-72) does not emit an event
Setter function LiquidSDIndexPool.setPaused(bool) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#480-483) does not emit an event
Setter function CoinbaseLSDIndexAdapter.initialize(address,address) (contracts/liquidSDIndex/adapters/CoinbaseLSDIndexAdapter.sol#17-19) does not emit an event
Setter function FraxLSDIndexAdapter.initialize(address,address) (contracts/liquidSDIndex/adapters/FraxLSDIndexAdapter.sol#17-19) does not emit an event
Setter function LidoLSDIndexAdapter.initialize(address,address) (contracts/liquidSDIndex/adapters/LidoLSDIndexAdapter.sol#16-18) does not emit an event
Setter function RocketPoolLSDIndexAdapter.initialize(address,address) (contracts/liquidSDIndex/adapters/RocketPoolLSDIndexAdapter.sol#17-19) does not emit an event
Setter function LSDIndexAdapterMock.initialize(address,address,uint256) (contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#13-20) does not emit an event
Setter function LSDIndexAdapterMock.setExchangeRate(uint256) (contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#30-32) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Condition variable is not initialized found in Initializable._disableInitializers() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#145-151)
Condition variable is not initialized found in Initializable._isInitializing() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#163-165)
Reference: initialize method should has permission check

Condition variable is not initialized found in SDLPoolCCIPControllerSecondary._initiateUpdate(uint64,address,bytes) (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#119-140)
Condition variable is not initialized found in KeyValidationOracle._initiateKeyPairValidation(address,uint256,bool) (contracts/ethStaking/KeyValidationOracle.sol#112-129)
Reference: initialize method should has permission check

value assignment lack of validation	SDLPoolMock.setEffectiveBalance(address,uint256) (contracts/core/test/SDLPoolMock.sol#15-17)effectiveBalances[_account] = _amount (contracts/core/test/SDLPoolMock.sol#16)
Reference: input validation

variable lacks a zero-check on 		- Operator.transferOwnableContracts(address[],address) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#235-240)
variable lacks a zero-check on 		- Operator.ownerTransferAndCall(address,uint256,bytes) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#332-338)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

variable lacks a zero-check on 		- Router.constructor(address,address) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#65-70)
variable lacks a zero-check on 		- Router.setWrappedNative(address) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#228-230)
variable lacks a zero-check on 		- VestingWallet.release(address) (node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#101-106)
variable lacks a zero-check on 		- VestingWallet.vestedAmount(address,uint64) (node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#118-120)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

Variable 'BytesLib.concatStorage(bytes,bytes).sc_concatStorage_asm_0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#147)' in BytesLib.concatStorage(bytes,bytes) (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#91-226) potentially used before declaration: sc_concatStorage_asm_0 = keccak256(uint256,uint256)(0x0,0x20) + slength_concatStorage_asm_0 / 32 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#195)
Variable 'BytesLib.concatStorage(bytes,bytes).submod_concatStorage_asm_0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#161)' in BytesLib.concatStorage(bytes,bytes) (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#91-226) potentially used before declaration: submod_concatStorage_asm_0 = 32 - slengthmod_concatStorage_asm_0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#204)
Variable 'BytesLib.concatStorage(bytes,bytes).mc_concatStorage_asm_0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#162)' in BytesLib.concatStorage(bytes,bytes) (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#91-226) potentially used before declaration: mc_concatStorage_asm_0 = _postBytes + submod_concatStorage_asm_0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#205)
Variable 'BytesLib.concatStorage(bytes,bytes).end_concatStorage_asm_0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#163)' in BytesLib.concatStorage(bytes,bytes) (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#91-226) potentially used before declaration: end_concatStorage_asm_0 = _postBytes + mlength_concatStorage_asm_0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#206)
Variable 'BytesLib.concatStorage(bytes,bytes).mask_concatStorage_asm_0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#164)' in BytesLib.concatStorage(bytes,bytes) (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#91-226) potentially used before declaration: mask_concatStorage_asm_0 = 0x100 ** submod_concatStorage_asm_0 - 1 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#207)
Reference:  

Reentrancy in LinkPoolNFT.mint(address) (contracts/core/tokens/LinkPoolNFT.sol#26-31):
	External calls:
	- _safeMint(_to,tokenId) (contracts/core/tokens/LinkPoolNFT.sol#29)
		- IERC721Receiver(to).onERC721Received(_msgSender(),from,tokenId,data) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#401-412)
Reference:  

AuthorizedReceiver.setAuthorizedSenders(address[]) (node_modules/@chainlink/contracts/src/v0.7/AuthorizedReceiver.sol#16-31) compares to a boolean constant:
	-require(bool,string)(s_authorizedSenders[senders[i_scope_0]] == false,Must not have duplicate senders) (node_modules/@chainlink/contracts/src/v0.7/AuthorizedReceiver.sol#25)
Reference:  

RESDLTokenBridge.transferRESDL(uint64,address,uint256,bool,uint256) (contracts/core/ccip/RESDLTokenBridge.sol#84-135) compares to a boolean constant:
	-_payNative == false && msg.value != 0 (contracts/core/ccip/RESDLTokenBridge.sol#93)
WrappedTokenBridge.transferTokens(uint64,address,uint256,bool,uint256) (contracts/core/ccip/WrappedTokenBridge.sol#106-117) compares to a boolean constant:
	-_payNative == false && msg.value != 0 (contracts/core/ccip/WrappedTokenBridge.sol#113)
DistributionOracle.withdrawLink(uint256) (contracts/core/priorityPool/DistributionOracle.sol#217-220) compares to a boolean constant:
	-link.transfer(msg.sender,_amount) != true (contracts/core/priorityPool/DistributionOracle.sol#219)
PriorityPool.withdraw(uint256,uint256,uint256,bytes32[],bool) (contracts/core/priorityPool/PriorityPool.sol#225-266) compares to a boolean constant:
	-_shouldUnqueue == true (contracts/core/priorityPool/PriorityPool.sol#237)
Reference:  

	- ConfirmedOwnerWithProposal.acceptOwnership() (node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#37-45)
	- Operator.transferOwnableContracts(address[],address) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#235-240)
Reference:  

	- Router.setWrappedNative(address) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#228-230)
	- Router.applyRampUpdates(Router.OnRamp[],Router.OffRamp[],Router.OffRamp[]) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#264-296)
	- StakedotlinkCouncil.mintWithTokenURI(address,uint256,string) (contracts/governance/StakedotlinkCouncil.sol#100-111)
	- StakedotlinkCouncil.burn(uint256) (contracts/governance/StakedotlinkCouncil.sol#131-152)
	- StakedotlinkCouncil.setTokenURI(uint256,string) (contracts/governance/StakedotlinkCouncil.sol#178-182)
Reference:  

	- MerkleDistributor.addDistribution(address,bytes32,uint256,uint256) (contracts/airdrop/MerkleDistributor.sol#77-93)
	- MerkleDistributor.updateDistribution(address,bytes32,uint256,uint256) (contracts/airdrop/MerkleDistributor.sol#128-143)
	- MerkleDistributor.withdrawUnclaimedTokens(address,bytes32,uint256) (contracts/airdrop/MerkleDistributor.sol#208-224)
	- MerkleDistributor.pauseForWithdrawal(address) (contracts/airdrop/MerkleDistributor.sol#232-239)
	- MerkleDistributor.setExpiryTimestamp(address,uint256) (contracts/airdrop/MerkleDistributor.sol#246-251)
	- RewardsInitiator.whitelistCaller(address,bool) (contracts/core/RewardsInitiator.sol#101-104)
	- StakingRewardsPool.__StakingRewardsPool_init(address,string,string) (contracts/core/base/StakingRewardsPool.sol#20-29)
	- StakingPool.deposit(address,uint256) (contracts/core/StakingPool.sol#80-88)
	- StakingPool.withdraw(address,address,uint256) (contracts/core/StakingPool.sol#97-116)
	- StakingPool.addStrategy(address) (contracts/core/StakingPool.sol#226-230)
	- StakingPool.removeStrategy(uint256,bytes) (contracts/core/StakingPool.sol#237-255)
	- StakingPool.reorderStrategies(uint256[]) (contracts/core/StakingPool.sol#261-274)
	- StakingPool.addFee(address,uint256) (contracts/core/StakingPool.sol#281-284)
	- StakingPool.updateFee(uint256,address,uint256) (contracts/core/StakingPool.sol#292-308)
	- StakingPool.setPriorityPool(address) (contracts/core/StakingPool.sol#375-377)
	- StakingPool.setRewardsInitiator(address) (contracts/core/StakingPool.sol#384-386)
	- RewardsPoolController.addToken(address,address) (contracts/core/base/RewardsPoolController.sol#154-165)
	- RewardsPoolController.removeToken(address) (contracts/core/base/RewardsPoolController.sol#171-185)
	- SDLPoolCCIPController.setMaxLINKFee(uint256) (contracts/core/ccip/base/SDLPoolCCIPController.sol#130-132)
	- SDLPoolCCIPController.setRESDLTokenBridge(address) (contracts/core/ccip/base/SDLPoolCCIPController.sol#138-140)
	- SDLPoolCCIPControllerPrimary.handleOutgoingRESDL(uint64,address,uint256) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#103-116)
	- SDLPoolCCIPControllerPrimary.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#125-134)
	- SDLPoolCCIPControllerPrimary.addWhitelistedChain(uint64,address,bytes,bytes) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#151-164)
	- SDLPoolCCIPControllerPrimary.removeWhitelistedChain(uint64) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#170-184)
	- SDLPoolCCIPControllerPrimary.setWrappedRewardToken(address,address) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#202-205)
	- SDLPoolCCIPControllerPrimary.setUpdateExtraArgs(uint64,bytes) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#212-216)
	- SDLPoolCCIPControllerPrimary.setRewardsExtraArgs(uint64,bytes) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#223-227)
	- SDLPoolCCIPControllerPrimary.setRewardsInitiator(address) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#234-236)
	- SDLPoolCCIPControllerSecondary.setExtraArgs(bytes) (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#108-111)
	- DistributionOracle.fulfillRequest(bytes32,bytes32,bytes32,uint256,uint256) (contracts/core/priorityPool/DistributionOracle.sol#163-177)
	- DistributionOracle.executeManualVerification() (contracts/core/priorityPool/DistributionOracle.sol#182-192)
	- DistributionOracle.rejectManualVerificationAndRetry() (contracts/core/priorityPool/DistributionOracle.sol#197-201)
	- DistributionOracle.cancelRequest(bytes32,uint256) (contracts/core/priorityPool/DistributionOracle.sol#208-211)
	- DistributionOracle.setUpdateParams(uint64,uint128,uint64) (contracts/core/priorityPool/DistributionOracle.sol#229-238)
	- DistributionOracle.toggleManualVerification() (contracts/core/priorityPool/DistributionOracle.sol#243-246)
	- DistributionOracle.setChainlinkParams(bytes32,uint256) (contracts/core/priorityPool/DistributionOracle.sol#253-257)
	- PriorityPool.updateDistribution(bytes32,bytes32,uint256,uint256) (contracts/core/priorityPool/PriorityPool.sol#409-424)
	- PriorityPool.setPoolStatus(PriorityPool.PoolStatus) (contracts/core/priorityPool/PriorityPool.sol#437-442)
	- PriorityPool.setPoolStatusClosed() (contracts/core/priorityPool/PriorityPool.sol#447-451)
	- PriorityPool.setQueueDepositParams(uint128,uint128) (contracts/core/priorityPool/PriorityPool.sol#458-462)
	- PriorityPool.setDistributionOracle(address) (contracts/core/priorityPool/PriorityPool.sol#468-470)
	- LinearBoostController.setMaxLockingDuration(uint64) (contracts/core/sdlPool/LinearBoostController.sol#45-48)
	- LinearBoostController.setMaxBoost(uint64) (contracts/core/sdlPool/LinearBoostController.sol#56-59)
	- SDLPool.__SDLPoolBase_init(string,string,address,address) (contracts/core/sdlPool/base/SDLPool.sol#93-104)
	- SDLPool.setBaseURI(string) (contracts/core/sdlPool/base/SDLPool.sol#363-365)
	- SDLPool.setBoostController(address) (contracts/core/sdlPool/base/SDLPool.sol#372-374)
	- SDLPool.setCCIPController(address) (contracts/core/sdlPool/base/SDLPool.sol#381-383)
	- SDLPoolPrimary.initiateUnlock(uint256) (contracts/core/sdlPool/SDLPoolPrimary.sol#107-121)
	- SDLPoolPrimary.withdraw(uint256,uint256) (contracts/core/sdlPool/SDLPoolPrimary.sol#134-164)
	- SDLPoolPrimary.handleOutgoingRESDL(address,uint256,address) (contracts/core/sdlPool/SDLPoolPrimary.sol#172-199)
	- SDLPoolPrimary.handleIncomingRESDL(address,uint256,SDLPool.Lock) (contracts/core/sdlPool/SDLPoolPrimary.sol#207-223)
	- SDLPoolPrimary.handleIncomingUpdate(uint256,int256) (contracts/core/sdlPool/SDLPoolPrimary.sol#231-253)
	- SDLPoolSecondary.initiateUnlock(uint256) (contracts/core/sdlPool/SDLPoolSecondary.sol#180-200)
	- SDLPoolSecondary.withdraw(uint256,uint256) (contracts/core/sdlPool/SDLPoolSecondary.sol#214-239)
	- SDLPoolSecondary.handleOutgoingRESDL(address,uint256,address) (contracts/core/sdlPool/SDLPoolSecondary.sol#259-281)
	- SDLPoolSecondary.handleIncomingRESDL(address,uint256,SDLPool.Lock) (contracts/core/sdlPool/SDLPoolSecondary.sol#289-307)
	- SDLPoolSecondary.handleOutgoingUpdate() (contracts/core/sdlPool/SDLPoolSecondary.sol#313-328)
	- SDLPoolSecondary.handleIncomingUpdate(uint256) (contracts/core/sdlPool/SDLPoolSecondary.sol#336-347)
	- StrategyMock.deposit(uint256) (contracts/core/test/StrategyMock.sol#45-49)
	- StrategyMock.withdraw(uint256) (contracts/core/test/StrategyMock.sol#51-56)
	- StrategyMock.updateDeposits(bytes) (contracts/core/test/StrategyMock.sol#58-79)
	- StrategyMock.setMaxDeposits(uint256) (contracts/core/test/StrategyMock.sol#101-103)
	- StrategyMock.setMinDeposits(uint256) (contracts/core/test/StrategyMock.sol#105-107)
	- RewardsPoolControllerV1.addToken(address,address) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#178-189)
	- RewardsPoolControllerV1.removeToken(address) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#195-209)
	- DelegatorPool.onTokenTransfer(address,uint256,bytes) (contracts/core/test/deprecated/DelegatorPool.sol#82-104)
	- DelegatorPool.setLockedApproval(address,uint256) (contracts/core/test/deprecated/DelegatorPool.sol#200-203)
	- DelegatorPool.burnLockedBalance(address,uint256) (contracts/core/test/deprecated/DelegatorPool.sol#210-224)
	- DelegatorPool.setPoolRouter(address) (contracts/core/test/deprecated/DelegatorPool.sol#230-233)
	- DelegatorPool.setCommunityPool(address,bool) (contracts/core/test/deprecated/DelegatorPool.sol#240-243)
	- DelegatorPool.retireDelegatorPool(address[],address) (contracts/core/test/deprecated/DelegatorPool.sol#249-276)
	- OwnersRewardsPoolV1.withdraw(address) (contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#48-60)
	- PoolOwnersV1.addRewardToken(address,address,address) (contracts/core/test/deprecated/PoolOwnersV1.sol#157-168)
	- PoolOwnersV1.removeRewardToken(uint16) (contracts/core/test/deprecated/PoolOwnersV1.sol#174-185)
	- LinkPoolNFT.mint(address) (contracts/core/tokens/LinkPoolNFT.sol#26-31)
	- LinkPoolNFT.setBaseURI(string) (contracts/core/tokens/LinkPoolNFT.sol#37-39)
	- EthStakingStrategy.setWLOperatorController(address) (contracts/ethStaking/EthStakingStrategy.sol#286-289)
	- EthStakingStrategy.setNWLOperatorController(address) (contracts/ethStaking/EthStakingStrategy.sol#295-298)
	- EthStakingStrategy.setBeaconOracle(address) (contracts/ethStaking/EthStakingStrategy.sol#304-307)
	- EthStakingStrategy.setMaxDeposits(uint256) (contracts/ethStaking/EthStakingStrategy.sol#337-340)
	- EthStakingStrategy.setMinDeposits(uint256) (contracts/ethStaking/EthStakingStrategy.sol#346-349)
	- EthStakingStrategy.setDepositController(address) (contracts/ethStaking/EthStakingStrategy.sol#355-358)
	- EthStakingStrategy.setRewardsReceiver(address) (contracts/ethStaking/EthStakingStrategy.sol#364-367)
	- KeyValidationOracle.setOracleConfig(address,bytes32,uint256) (contracts/ethStaking/KeyValidationOracle.sol#87-96)
	- OperatorController.setKeyValidationOracle(address) (contracts/ethStaking/base/OperatorController.sol#271-274)
	- OperatorController.setBeaconOracle(address) (contracts/ethStaking/base/OperatorController.sol#280-283)
	- NWLOperatorController.withdrawStake(uint256,uint256) (contracts/ethStaking/NWLOperatorController.sol#349-357)
	- OperatorWhitelist.useWhitelist(address) (contracts/ethStaking/OperatorWhitelist.sol#41-46)
	- OperatorWhitelist.addWhitelistEntries(address[]) (contracts/ethStaking/OperatorWhitelist.sol#52-58)
	- OperatorWhitelist.removeWhitelistEntries(address[]) (contracts/ethStaking/OperatorWhitelist.sol#64-70)
	- OperatorWhitelist.setWLOperatorController(address) (contracts/ethStaking/OperatorWhitelist.sol#76-78)
	- WLOperatorController.removeKeyPairs(uint256,uint256) (contracts/ethStaking/WLOperatorController.sol#75-93)
	- WLOperatorController.reportKeyPairValidation(uint256,bool) (contracts/ethStaking/WLOperatorController.sol#100-115)
	- WLOperatorController.assignNextValidators(uint256[],uint256[],uint256) (contracts/ethStaking/WLOperatorController.sol#125-272)
	- WLOperatorController.reportStoppedValidators(uint256[],uint256[]) (contracts/ethStaking/WLOperatorController.sol#369-404)
	- WLOperatorController.setBatchSize(uint256) (contracts/ethStaking/WLOperatorController.sol#410-412)
	- WLOperatorController.setOperatorWhitelist(address) (contracts/ethStaking/WLOperatorController.sol#418-420)
	- GovernanceController.addRole(string,address[],address[],bytes4[][]) (contracts/governance/GovernanceController.sol#99-120)
	- GovernanceController.grantRole(uint256,address) (contracts/governance/GovernanceController.sol#127-131)
	- GovernanceController.addRoleFunctions(uint256,address[],bytes4[][]) (contracts/governance/GovernanceController.sol#156-172)
	- GovernanceController.removeRoleFunctions(uint256,address[],bytes4[][]) (contracts/governance/GovernanceController.sol#180-196)
	- VaultControllerStrategy.__VaultControllerStrategy_init(address,address,address,address,VaultControllerStrategy.Fee[],uint256) (contracts/linkStaking/base/VaultControllerStrategy.sol#56-78)
	- VaultControllerStrategy.deposit(uint256) (contracts/linkStaking/base/VaultControllerStrategy.sol#93-110)
	- VaultControllerStrategy.updateDeposits(bytes) (contracts/linkStaking/base/VaultControllerStrategy.sol#157-191)
	- VaultControllerStrategy.addFee(address,uint256) (contracts/linkStaking/base/VaultControllerStrategy.sol#270-273)
	- VaultControllerStrategy.updateFee(uint256,address,uint256) (contracts/linkStaking/base/VaultControllerStrategy.sol#284-300)
	- VaultControllerStrategy.setMaxDepositSizeBP(uint256) (contracts/linkStaking/base/VaultControllerStrategy.sol#307-311)
	- VaultControllerStrategy.setVaultImplementation(address) (contracts/linkStaking/base/VaultControllerStrategy.sol#320-324)
	- CommunityVCS.setVaultDeploymentParams(uint128,uint128) (contracts/linkStaking/CommunityVCS.sol#115-119)
	- Vault.__Vault_init(address,address,address,address) (contracts/linkStaking/base/Vault.sol#36-48)
	- OperatorVCS.withdrawOperatorRewards(address,uint256) (contracts/linkStaking/OperatorVCS.sol#102-113)
	- OperatorVCS.updateDeposits(bytes) (contracts/linkStaking/OperatorVCS.sol#155-216)
	- OperatorVCS.addVault(address,address,address) (contracts/linkStaking/OperatorVCS.sol#225-243)
	- OperatorVCS.setOperatorRewardPercentage(uint256) (contracts/linkStaking/OperatorVCS.sol#274-281)
	- OperatorVCS.togglePreRelease() (contracts/linkStaking/OperatorVCS.sol#287-289)
	- OperatorVault.deposit(uint256) (contracts/linkStaking/OperatorVault.sol#95-99)
	- OperatorVault.withdrawRewards() (contracts/linkStaking/OperatorVault.sol#139-151)
	- OperatorVault.updateDeposits(uint256,address) (contracts/linkStaking/OperatorVault.sol#175-200)
	- OperatorVault.setOperator(address) (contracts/linkStaking/OperatorVault.sol#211-215)
	- OperatorVault.setRewardsReceiver(address) (contracts/linkStaking/OperatorVault.sol#227-233)
	- StakingMock.onTokenTransfer(address,uint256,bytes) (contracts/linkStaking/test/StakingMock.sol#39-51)
	- StakingMockV1.onTokenTransfer(address,uint256,bytes) (contracts/linkStaking/test/StakingMockV1.sol#30-37)
	- StakingRewardsMock.claimReward() (contracts/linkStaking/test/StakingRewardsMock.sol#27-32)
	- VaultV1.__Vault_init(address,address,address) (contracts/linkStaking/test/deprecated/VaultV1.sol#52-62)
	- VaultV1.migrate(bytes) (contracts/linkStaking/test/deprecated/VaultV1.sol#102-105)
	- OperatorVaultV1.setOperator(address) (contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#61-64)
	- VaultControllerStrategyUpgrade.__VaultControllerStrategy_init(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[]) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#86-109)
	- VaultControllerStrategyUpgrade.deposit(uint256) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#123-127)
	- VaultControllerStrategyUpgrade.updateDeposits(bytes) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#227-251)
	- VaultControllerStrategyUpgrade.addFee(address,uint256) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#315-318)
	- VaultControllerStrategyUpgrade.updateFee(uint256,address,uint256) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#326-342)
	- VaultControllerStrategyUpgrade.setMinDepositThreshold(uint256) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#349-354)
	- VaultControllerStrategyUpgrade.setVaultImplementation(address) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#360-364)
	- LiquidSDIndexPool.addLSDToken(address,address,uint256[]) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#328-347)
	- LiquidSDIndexPool.removeLSDToken(address,uint256[]) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#354-385)
	- LiquidSDIndexPool.setCompositionTargets(uint256[]) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#391-402)
	- LiquidSDIndexPool.setCompositionTolerance(uint256) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#411-415)
	- LiquidSDIndexPool.setCompositionEnforcementThreshold(uint256) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#425-428)
	- LiquidSDIndexPool.setWithdrawalFee(uint256) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#434-438)
	- LiquidSDIndexPool.addFee(address,uint256) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#445-449)
	- LiquidSDIndexPool.updateFee(uint256,address,uint256) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#457-474)
	- LiquidSDIndexPool.setPaused(bool) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#480-483)
	- LSDIndexAdapter.__LiquidSDAdapter_init(address,address) (contracts/liquidSDIndex/base/LSDIndexAdapter.sol#20-26)
	- Vesting.terminateVesting(address[]) (contracts/vesting/Vesting.sol#27-37)
Reference:  

	- ConfirmedOwnerWithProposal.transferOwnership(address) (node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#30-32)
	- LinkTokenReceiver.onTokenTransfer(address,uint256,bytes) (node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#13-27)
	- Operator.oracleRequest(address,uint256,bytes32,address,bytes4,uint256,uint256,bytes) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#98-117)
	- Operator.operatorRequest(address,uint256,bytes32,bytes4,uint256,uint256,bytes) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#131-149)
	- Operator.ownerForward(address,bytes) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#319-323)
	- Operator.ownerTransferAndCall(address,uint256,bytes) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#332-338)
Reference:  

	- Router.ccipSend(uint64,Client.EVM2AnyMessage) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#104-141)
	- Router.routeMessage(Client.Any2EVMMessage,uint16,uint256,address) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#155-214)
	- CCIPReceiver.ccipReceive(Client.Any2EVMMessage) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#27-29)
	- OwnableUpgradeable.renounceOwnership() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#66-68)
	- OwnableUpgradeable.transferOwnership(address) (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#74-77)
	- UUPSUpgradeable.upgradeTo(address) (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#74-77)
	- UUPSUpgradeable.upgradeToAndCall(address,bytes) (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#89-92)
	- Ownable.renounceOwnership() (node_modules/@openzeppelin/contracts/access/Ownable.sol#61-63)
	- Ownable.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable.sol#69-72)
	- StakedotlinkCouncil.mint(address,uint256) (contracts/governance/StakedotlinkCouncil.sol#90-92)
Reference:  

	- MerkleDistributor.addDistributions(address[],bytes32[],uint256[],uint256[]) (contracts/airdrop/MerkleDistributor.sol#54-68)
	- MerkleDistributor.updateDistributions(address[],bytes32[],uint256[],uint256[]) (contracts/airdrop/MerkleDistributor.sol#102-116)
	- RewardsInitiator.updateRewards(uint256[],bytes) (contracts/core/RewardsInitiator.sol#38-42)
	- RewardsPool.withdraw(address) (contracts/core/RewardsPool.sol#57-60)
	- RewardsPool.onTokenTransfer(address,uint256,bytes) (contracts/core/RewardsPool.sol#65-72)
	- ERC677Upgradeable.__ERC677_init(string,string,uint256) (contracts/core/tokens/base/ERC677Upgradeable.sol#9-16)
	- StakingPool.strategyDeposit(uint256,uint256) (contracts/core/StakingPool.sol#123-126)
	- StakingPool.strategyWithdraw(uint256,uint256) (contracts/core/StakingPool.sol#133-136)
	- StakingPool.updateStrategyRewards(uint256[],bytes) (contracts/core/StakingPool.sol#345-348)
	- RewardsPoolController.__RewardsPoolController_init() (contracts/core/base/RewardsPoolController.sol#28-31)
	- RESDLTokenBridge.ccipReceive(Client.Any2EVMMessage) (contracts/core/ccip/RESDLTokenBridge.sol#171-192)
	- SDLPoolCCIPController.ccipSend(uint64,Client.EVM2AnyMessage) (contracts/core/ccip/base/SDLPoolCCIPController.sol#89-100)
	- SDLPoolCCIPController.ccipReceive(Client.Any2EVMMessage) (contracts/core/ccip/base/SDLPoolCCIPController.sol#102-110)
	- SDLPoolCCIPController.recoverTokens(address[],address) (contracts/core/ccip/base/SDLPoolCCIPController.sol#117-124)
	- SDLPoolCCIPControllerPrimary.distributeRewards() (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#56-93)
	- SDLPoolCCIPControllerPrimary.approveRewardTokens(address[]) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#190-195)
	- SDLPoolCCIPControllerSecondary.handleOutgoingRESDL(uint64,address,uint256) (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#80-86)
	- SDLPoolCCIPControllerSecondary.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken) (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#94-102)
	- WrappedTokenBridge.onTokenTransfer(address,uint256,bytes) (contracts/core/ccip/WrappedTokenBridge.sol#83-96)
	- WrappedTokenBridge.transferTokens(uint64,address,uint256,bool,uint256) (contracts/core/ccip/WrappedTokenBridge.sol#106-117)
	- WrappedTokenBridge.recoverTokens(address[],address) (contracts/core/ccip/WrappedTokenBridge.sol#140-147)
	- DistributionOracle.pauseForUpdate() (contracts/core/priorityPool/DistributionOracle.sol#142-144)
	- DistributionOracle.requestUpdate() (contracts/core/priorityPool/DistributionOracle.sol#151-153)
	- PriorityPool.onTokenTransfer(address,uint256,bytes) (contracts/core/priorityPool/PriorityPool.sol#185-202)
	- PriorityPool.pauseForUpdate() (contracts/core/priorityPool/PriorityPool.sol#429-431)
	- SDLPool.addToken(address,address) (contracts/core/sdlPool/base/SDLPool.sol#336-339)
	- SDLPoolPrimary.onTokenTransfer(address,uint256,bytes) (contracts/core/sdlPool/SDLPoolPrimary.sol#61-80)
	- SDLPoolPrimary.migrate(address,uint256,uint64) (contracts/core/sdlPool/SDLPoolPrimary.sol#264-272)
	- SDLPoolSecondary.onTokenTransfer(address,uint256,bytes) (contracts/core/sdlPool/SDLPoolSecondary.sol#132-151)
	- SDLPoolCCIPControllerMock.handleOutgoingRESDL(uint64,address,uint256) (contracts/core/test/SDLPoolCCIPControllerMock.sol#29-35)
	- SDLPoolCCIPControllerMock.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken) (contracts/core/test/SDLPoolCCIPControllerMock.sol#37-45)
	- WrappedSDTokenMock.onTokenTransfer(address,uint256,bytes) (contracts/core/test/WrappedSDTokenMock.sol#24-32)
	- RewardsPoolControllerV1.__RewardsPoolController_init(string,string) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#30-37)
	- OwnersRewardsPoolV1.onTokenTransfer(address,uint256,bytes) (contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#65-72)
	- PoolAllowanceV1.mintAllowance(address,uint256) (contracts/core/test/deprecated/PoolAllowanceV1.sol#31-33)
	- PoolAllowanceV1.burnAllowance(address,uint256) (contracts/core/test/deprecated/PoolAllowanceV1.sol#40-42)
	- PoolOwnersV1.onTokenTransfer(address,uint256,bytes) (contracts/core/test/deprecated/PoolOwnersV1.sol#94-101)
	- LPLMigration.onTokenTransfer(address,uint256,bytes) (contracts/core/tokens/LPLMigration.sol#31-39)
	- StakingAllowance.mint(address,uint256) (contracts/core/tokens/StakingAllowance.sol#20-22)
	- StakingAllowance.mintToContract(address,address,uint256,bytes) (contracts/core/tokens/StakingAllowance.sol#30-38)
	- WrappedSDToken.onTokenTransfer(address,uint256,bytes) (contracts/core/tokens/WrappedSDToken.sol#27-34)
	- DepositController.depositEther(bytes32,bytes32,bytes32,uint256,uint256,uint256[],uint256[]) (contracts/ethStaking/DepositController.sol#46-64)
	- EthStakingStrategy.withdraw(uint256) (contracts/ethStaking/EthStakingStrategy.sol#224-226)
	- EthStakingStrategy.nwlWithdraw(address,uint256) (contracts/ethStaking/EthStakingStrategy.sol#234-237)
	- KeyValidationOracle.reportKeyPairValidation(bytes32,uint256,bool,bool) (contracts/ethStaking/KeyValidationOracle.sol#68-79)
	- OperatorController.onTokenTransfer(address,uint256,bytes) (contracts/ethStaking/base/OperatorController.sol#176-183)
	- NWLOperatorController.addKeyPairs(uint256,uint256,bytes,bytes) (contracts/ethStaking/NWLOperatorController.sol#83-92)
	- WLOperatorController.addKeyPairs(uint256,uint256,bytes,bytes) (contracts/ethStaking/WLOperatorController.sol#60-68)
	- EthStakingStrategyMock.nwlWithdraw(address,uint256) (contracts/ethStaking/test/EthStakingStrategyMock.sol#15-17)
	- GovernanceController.callFunction(uint256,address,bytes) (contracts/governance/GovernanceController.sol#73-90)
	- GovernanceController.revokeRole(uint256,address) (contracts/governance/GovernanceController.sol#138-140)
	- VaultControllerStrategy.upgradeVaults(uint256,uint256,bytes) (contracts/linkStaking/base/VaultControllerStrategy.sol#243-252)
	- CommunityVCS.addVaults(uint256) (contracts/linkStaking/CommunityVCS.sol#106-108)
	- Vault.deposit(uint256) (contracts/linkStaking/base/Vault.sol#62-65)
	- CommunityVault.claimRewards(uint256,address) (contracts/linkStaking/CommunityVault.sol#39-45)
	- OperatorVCS.onTokenTransfer(address,uint256,bytes) (contracts/linkStaking/OperatorVCS.sol#84-90)
	- OperatorVCS.setOperator(uint256,address) (contracts/linkStaking/OperatorVCS.sol#251-253)
	- OperatorVCS.setRewardsReceiver(uint256,address) (contracts/linkStaking/OperatorVCS.sol#261-263)
	- OperatorVault.raiseAlert(address) (contracts/linkStaking/OperatorVault.sol#114-124)
	- VaultV1.deposit(uint256) (contracts/linkStaking/test/deprecated/VaultV1.sol#73-76)
	- OperatorVaultV1.raiseAlert() (contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#51-55)
	- VaultControllerStrategyUpgrade.migrateVaults(uint256,uint256,bytes) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#274-283)
	- VaultControllerStrategyUpgrade.upgradeVaults(uint256,uint256,bytes) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#291-300)
	- OperatorVCSUpgrade.addVault(address) (contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#75-85)
	- OperatorVCSUpgrade.setOperator(uint256,address) (contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#92-94)
Reference:  

require() missing error messages
	 - require(bool)(len <= data.length) (node_modules/@chainlink/contracts/src/v0.8/vendor/BufferChainlink.sol#98)
require() missing error messages
	 - require(bool)(denominator > prod1) (node_modules/@openzeppelin/contracts/utils/math/Math.sol#78)
Reference: https://dev.to/tawseef/require-vs-assert-in-solidity-5e9d

Potential price manipulation risk:
	- In function canWithdraw
		-- stLINKCanWithdraw = MathUpgradeable.min(stakingPool.balanceOf(_account),stakingPool.canWithdraw() + totalQueued - canUnqueue) (contracts/core/priorityPool/PriorityPool.sol#171-174) have potential price manipulated risk from stLINKCanWithdraw and call None which could influence variable:stLINKCanWithdraw
Potential price manipulation risk:
	- In function getLockIdsByOwner
		-- lockCount = balanceOf(_owner) (contracts/core/sdlPool/base/SDLPool.sol#179) have potential price manipulated risk from lockCount and call None which could influence variable:lockIds
Potential price manipulation risk:
	- In function getEthBalance
		-- balance = addr.balance (contracts/core/test/Multicall3.sol#222) have potential price manipulated risk from balance and call balance(address) which could influence variable:balance
Potential price manipulation risk:
	- In function withdraw
		-- toWithdraw = balanceOf(msg.sender) (contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#37) have potential price manipulated risk from toWithdraw and call None which could influence variable:toWithdraw
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

Potential DoS Gas Limit Attack occur inDepositContract.constructor() (contracts/ethStaking/test/DepositContract.sol#68-72)BEGIN_LOOP (contracts/ethStaking/test/DepositContract.sol#70-71)
Potential DoS Gas Limit Attack occur inDepositContract.deposit(bytes,bytes,bytes,bytes32) (contracts/ethStaking/test/DepositContract.sol#89-145)BEGIN_LOOP (contracts/ethStaking/test/DepositContract.sol#134-141)
Reference: https://swcregistry.io/docs/SWC-128

Potential DoS Gas Limit Attack occur inOperator.distributeFunds(address[],uint256[]) (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#349-358)BEGIN_LOOP (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#352-356)
Reference: https://swcregistry.io/docs/SWC-128

Potential DoS Gas Limit Attack occur inBytesLib.concatStorage(bytes,bytes) (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#91-226)BEGIN_LOOP (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#177-185)
Reference: https://swcregistry.io/docs/SWC-128

Potential DoS Gas Limit Attack occur inMerkleDistributor.claimDistributions(address[],uint256[],address,uint256[],bytes32[][]) (contracts/airdrop/MerkleDistributor.sol#153-168)BEGIN_LOOP (contracts/airdrop/MerkleDistributor.sol#165-167)
Potential DoS Gas Limit Attack occur inRewardsInitiator.performUpkeep(bytes) (contracts/core/RewardsInitiator.sol#83-94)BEGIN_LOOP (contracts/core/RewardsInitiator.sol#89-91)
Potential DoS Gas Limit Attack occur inStakingPool.depositLiquidity() (contracts/core/StakingPool.sol#354-369)BEGIN_LOOP (contracts/core/StakingPool.sol#357-367)
Potential DoS Gas Limit Attack occur inStakingPool._withdrawLiquidity(uint256) (contracts/core/StakingPool.sol#402-417)BEGIN_LOOP (contracts/core/StakingPool.sol#405-416)
Potential DoS Gas Limit Attack occur inRewardsPoolController.distributeTokens(address[]) (contracts/core/base/RewardsPoolController.sol#102-106)BEGIN_LOOP (contracts/core/base/RewardsPoolController.sol#103-105)
Potential DoS Gas Limit Attack occur inRewardsPoolController._updateRewards(address) (contracts/core/base/RewardsPoolController.sol#191-195)BEGIN_LOOP (contracts/core/base/RewardsPoolController.sol#192-194)
Potential DoS Gas Limit Attack occur inSDLPoolCCIPControllerPrimary._distributeRewards(uint64,address[],uint256[]) (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#244-287)BEGIN_LOOP (contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#253-257)
Potential DoS Gas Limit Attack occur inSDLPoolCCIPControllerSecondary._ccipReceive(Client.Any2EVMMessage) (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#147-165)BEGIN_LOOP (contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#152-155)
Potential DoS Gas Limit Attack occur inSDLPoolSecondary._mintQueuedNewLocks(address) (contracts/core/sdlPool/SDLPoolSecondary.sol#384-419)BEGIN_LOOP (contracts/core/sdlPool/SDLPoolSecondary.sol#388-409)
Potential DoS Gas Limit Attack occur inSDLPoolSecondary._executeQueuedLockUpdates(address,uint256[]) (contracts/core/sdlPool/SDLPoolSecondary.sol#451-510)BEGIN_LOOP (contracts/core/sdlPool/SDLPoolSecondary.sol#454-509)
Potential DoS Gas Limit Attack occur inMulticall3.aggregate(Multicall3.Call[]) (contracts/core/test/Multicall3.sol#41-55)BEGIN_LOOP (contracts/core/test/Multicall3.sol#46-54)
Potential DoS Gas Limit Attack occur inMulticall3.tryAggregate(bool,Multicall3.Call[]) (contracts/core/test/Multicall3.sol#62-75)BEGIN_LOOP (contracts/core/test/Multicall3.sol#66-74)
Potential DoS Gas Limit Attack occur inMulticall3.aggregate3(Multicall3.Call3[]) (contracts/core/test/Multicall3.sol#118-145)BEGIN_LOOP (contracts/core/test/Multicall3.sol#122-144)
Potential DoS Gas Limit Attack occur inMulticall3.aggregate3Value(Multicall3.Call3Value[]) (contracts/core/test/Multicall3.sol#151-187)BEGIN_LOOP (contracts/core/test/Multicall3.sol#156-184)
Potential DoS Gas Limit Attack occur inCCIPOffRampMock.constructor(address,address[],address[]) (contracts/core/test/chainlink/CCIPOffRampMock.sol#23-32)BEGIN_LOOP (contracts/core/test/chainlink/CCIPOffRampMock.sol#29-31)
Potential DoS Gas Limit Attack occur inCCIPOnRampMock.constructor(address[],address[],address) (contracts/core/test/chainlink/CCIPOnRampMock.sol#22-31)BEGIN_LOOP (contracts/core/test/chainlink/CCIPOnRampMock.sol#27-29)
Potential DoS Gas Limit Attack occur inRewardsPoolControllerV1.distributeTokens(address[]) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#126-130)BEGIN_LOOP (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#127-129)
Potential DoS Gas Limit Attack occur inRewardsPoolControllerV1._updateRewards(address) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#215-219)BEGIN_LOOP (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#216-218)
Potential DoS Gas Limit Attack occur inPoolOwnersV1._mintAllowance(address) (contracts/core/test/deprecated/PoolOwnersV1.sol#212-222)BEGIN_LOOP (contracts/core/test/deprecated/PoolOwnersV1.sol#214-221)
Potential DoS Gas Limit Attack occur inPoolOwnersV1._burnAllowance(address) (contracts/core/test/deprecated/PoolOwnersV1.sol#228-238)BEGIN_LOOP (contracts/core/test/deprecated/PoolOwnersV1.sol#230-237)
Potential DoS Gas Limit Attack occur inOperatorController._addKeyPairs(uint256,uint256,bytes,bytes) (contracts/ethStaking/base/OperatorController.sol#312-337)BEGIN_LOOP (contracts/ethStaking/base/OperatorController.sol#324-331)
Potential DoS Gas Limit Attack occur inOperatorController._storeKeyPair(uint256,uint256,bytes,bytes) (contracts/ethStaking/base/OperatorController.sol#346-371)BEGIN_LOOP (contracts/ethStaking/base/OperatorController.sol#365-370)
Potential DoS Gas Limit Attack occur inOperatorWhitelist.constructor(address,address[]) (contracts/ethStaking/OperatorWhitelist.sol#19-25)BEGIN_LOOP (contracts/ethStaking/OperatorWhitelist.sol#22-24)
Potential DoS Gas Limit Attack occur inOperatorControllerMock.assignNextValidators(uint256[],uint256[],uint256) (contracts/ethStaking/test/OperatorControllerMock.sol#54-67)BEGIN_LOOP (contracts/ethStaking/test/OperatorControllerMock.sol#59-66)
Potential DoS Gas Limit Attack occur inOperatorWhitelistMock.constructor(address[]) (contracts/ethStaking/test/OperatorWhitelistMock.sol#12-16)BEGIN_LOOP (contracts/ethStaking/test/OperatorWhitelistMock.sol#13-15)
Potential DoS Gas Limit Attack occur inVaultControllerStrategy._depositToVaults(uint256,uint256,uint256,uint256) (contracts/linkStaking/base/VaultControllerStrategy.sol#334-362)BEGIN_LOOP (contracts/linkStaking/base/VaultControllerStrategy.sol#342-359)
Potential DoS Gas Limit Attack occur inCommunityVCS.claimRewards(uint256,uint256,uint256) (contracts/linkStaking/CommunityVCS.sol#65-74)BEGIN_LOOP (contracts/linkStaking/CommunityVCS.sol#71-73)
Potential DoS Gas Limit Attack occur inCommunityVCS._deployVaults(uint256) (contracts/linkStaking/CommunityVCS.sol#125-136)BEGIN_LOOP (contracts/linkStaking/CommunityVCS.sol#133-135)
Potential DoS Gas Limit Attack occur inOperatorVCS._updateStrategyRewards() (contracts/linkStaking/OperatorVCS.sol#296-303)BEGIN_LOOP (contracts/linkStaking/OperatorVCS.sol#299-301)
Potential DoS Gas Limit Attack occur inVCSMock.addVaults(address[]) (contracts/linkStaking/test/VCSMock.sol#26-32)BEGIN_LOOP (contracts/linkStaking/test/VCSMock.sol#27-31)
Potential DoS Gas Limit Attack occur inVaultControllerStrategyUpgrade._depositToVaults(uint256,uint256,uint256,uint256) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#387-411)BEGIN_LOOP (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#394-409)
Reference: https://swcregistry.io/docs/SWC-128

CCIPReceiver (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#11-52) does not implement functions:
	- CCIPReceiver._ccipReceive(Client.Any2EVMMessage) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#33)
UUPSUpgradeable (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#22-112) does not implement functions:
	- UUPSUpgradeable._authorizeUpgrade(address) (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#104)
Reference:  

function:ConfirmedOwnerWithProposal.transferOwnership(address) (node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#30-32)is public and can be replaced with external 
function:ConfirmedOwnerWithProposal.owner() (node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#50-52)is public and can be replaced with external 
function:LinkTokenReceiver.onTokenTransfer(address,uint256,bytes) (node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#13-27)is public and can be replaced with external 
function:LinkTokenReceiver.getChainlinkToken() (node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#29)is public and can be replaced with external 
function:Operator.getChainlinkToken() (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#420-422)is public and can be replaced with external 
Reference:  

function:CCIPReceiver.supportsInterface(bytes4) (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#22-24)is public and can be replaced with external 
function:CCIPReceiver.getRouter() (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#41-43)is public and can be replaced with external 
function:ENSResolver.addr(bytes32) (node_modules/@chainlink/contracts/src/v0.8/vendor/ENSResolver.sol#5)is public and can be replaced with external 
function:OwnableUpgradeable.renounceOwnership() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#66-68)is public and can be replaced with external 
function:OwnableUpgradeable.transferOwnership(address) (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#74-77)is public and can be replaced with external 
function:UUPSUpgradeable.upgradeTo(address) (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#74-77)is public and can be replaced with external 
function:UUPSUpgradeable.upgradeToAndCall(address,bytes) (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#89-92)is public and can be replaced with external 
function:ERC20Upgradeable.name() (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#67-69)is public and can be replaced with external 
function:ERC20Upgradeable.symbol() (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#75-77)is public and can be replaced with external 
function:ERC20Upgradeable.decimals() (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#92-94)is public and can be replaced with external 
function:ERC20Upgradeable.totalSupply() (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#99-101)is public and can be replaced with external 
function:ERC20Upgradeable.balanceOf(address) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#106-108)is public and can be replaced with external 
function:ERC20Upgradeable.transfer(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#118-122)is public and can be replaced with external 
function:ERC20Upgradeable.approve(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#141-145)is public and can be replaced with external 
function:ERC20Upgradeable.transferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#163-168)is public and can be replaced with external 
function:ERC20Upgradeable.increaseAllowance(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#182-186)is public and can be replaced with external 
function:ERC20Upgradeable.decreaseAllowance(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#202-211)is public and can be replaced with external 
function:Ownable.renounceOwnership() (node_modules/@openzeppelin/contracts/access/Ownable.sol#61-63)is public and can be replaced with external 
function:Ownable.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable.sol#69-72)is public and can be replaced with external 
function:VestingWallet.release() (node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#89-94)is public and can be replaced with external 
function:VestingWallet.release(address) (node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#101-106)is public and can be replaced with external 
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
function:ERC721URIStorage.tokenURI(uint256) (node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#20-36)is public and can be replaced with external 
function:ERC165.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#26-28)is public and can be replaced with external 
function:StakedotlinkCouncil.balanceOf(address) (contracts/governance/StakedotlinkCouncil.sol#57-59)is public and can be replaced with external 
function:StakedotlinkCouncil.transferFrom(address,address,uint256) (contracts/governance/StakedotlinkCouncil.sol#68-82)is public and can be replaced with external 
function:StakedotlinkCouncil.mint(address,uint256) (contracts/governance/StakedotlinkCouncil.sol#90-92)is public and can be replaced with external 
function:StakedotlinkCouncil.mintWithTokenURI(address,uint256,string) (contracts/governance/StakedotlinkCouncil.sol#100-111)is public and can be replaced with external 
function:StakedotlinkCouncil.burn(uint256) (contracts/governance/StakedotlinkCouncil.sol#131-152)is public and can be replaced with external 
function:StakedotlinkCouncil.totalSupply() (contracts/governance/StakedotlinkCouncil.sol#157-159)is public and can be replaced with external 
function:StakedotlinkCouncil.tokenURI(uint256) (contracts/governance/StakedotlinkCouncil.sol#166-170)is public and can be replaced with external 
function:StakedotlinkCouncil.setTokenURI(uint256,string) (contracts/governance/StakedotlinkCouncil.sol#178-182)is public and can be replaced with external 
Reference:  

function:RewardsPool.distributeRewards() (contracts/core/RewardsPool.sol#77-83)is public and can be replaced with external 
function:RewardsPool.updateReward(address) (contracts/core/RewardsPool.sol#89-95)is public and can be replaced with external 
function:StakingRewardsPool.totalSupply() (contracts/core/base/StakingRewardsPool.sol#35-37)is public and can be replaced with external 
function:StakingRewardsPool.sharesOf(address) (contracts/core/base/StakingRewardsPool.sol#58-60)is public and can be replaced with external 
function:ERC677Upgradeable.transferAndCall(address,uint256,bytes) (contracts/core/tokens/base/ERC677Upgradeable.sol#18-28)is public and can be replaced with external 
function:StakingPool.initialize(address,string,string,StakingPool.Fee[]) (contracts/core/StakingPool.sol#41-52)is public and can be replaced with external 
function:RewardsPoolController.__RewardsPoolController_init() (contracts/core/base/RewardsPoolController.sol#28-31)is public and can be replaced with external 
function:RewardsPoolController.distributeTokens(address[]) (contracts/core/base/RewardsPoolController.sol#102-106)is public and can be replaced with external 
function:RewardsPoolController.withdrawRewards(address[]) (contracts/core/base/RewardsPoolController.sol#142-147)is public and can be replaced with external 
function:RewardsPoolController.addToken(address,address) (contracts/core/base/RewardsPoolController.sol#154-165)is public and can be replaced with external 
function:StakingRewardsPool.__StakingRewardsPool_init(address,string,string) (contracts/core/base/StakingRewardsPool.sol#20-29)is public and can be replaced with external 
function:StakingRewardsPool.balanceOf(address) (contracts/core/base/StakingRewardsPool.sol#44-51)is public and can be replaced with external 
function:Strategy.__Strategy_init(address,address) (contracts/core/base/Strategy.sol#20-25)is public and can be replaced with external 
function:Strategy.canDeposit() (contracts/core/base/Strategy.sol#36-43)is public and can be replaced with external 
function:Strategy.canWithdraw() (contracts/core/base/Strategy.sol#49-56)is public and can be replaced with external 
function:DistributionOracle.fulfillRequest(bytes32,bytes32,bytes32,uint256,uint256) (contracts/core/priorityPool/DistributionOracle.sol#163-177)is public and can be replaced with external 
function:PriorityPool.initialize(address,address,address,uint128,uint128) (contracts/core/priorityPool/PriorityPool.sol#93-110)is public and can be replaced with external 
function:SDLPool.addToken(address,address) (contracts/core/sdlPool/base/SDLPool.sol#336-339)is public and can be replaced with external 
function:SDLPoolPrimary.initialize(string,string,address,address) (contracts/core/sdlPool/SDLPoolPrimary.sol#30-41)is public and can be replaced with external 
function:SDLPoolSecondary.initialize(string,string,address,address,uint256) (contracts/core/sdlPool/SDLPoolSecondary.sol#66-79)is public and can be replaced with external 
function:SDLPool.__SDLPoolBase_init(string,string,address,address) (contracts/core/sdlPool/base/SDLPool.sol#93-104)is public and can be replaced with external 
function:Multicall3.aggregate(Multicall3.Call[]) (contracts/core/test/Multicall3.sol#41-55)is public and can be replaced with external 
function:Multicall3.blockAndAggregate(Multicall3.Call[]) (contracts/core/test/Multicall3.sol#103-113)is public and can be replaced with external 
function:Multicall3.aggregate3(Multicall3.Call3[]) (contracts/core/test/Multicall3.sol#118-145)is public and can be replaced with external 
function:Multicall3.aggregate3Value(Multicall3.Call3Value[]) (contracts/core/test/Multicall3.sol#151-187)is public and can be replaced with external 
function:Multicall3.getBlockHash(uint256) (contracts/core/test/Multicall3.sol#191-193)is public and can be replaced with external 
function:Multicall3.getBlockNumber() (contracts/core/test/Multicall3.sol#196-198)is public and can be replaced with external 
function:Multicall3.getCurrentBlockCoinbase() (contracts/core/test/Multicall3.sol#201-203)is public and can be replaced with external 
function:Multicall3.getCurrentBlockDifficulty() (contracts/core/test/Multicall3.sol#206-208)is public and can be replaced with external 
function:Multicall3.getCurrentBlockGasLimit() (contracts/core/test/Multicall3.sol#211-213)is public and can be replaced with external 
function:Multicall3.getCurrentBlockTimestamp() (contracts/core/test/Multicall3.sol#216-218)is public and can be replaced with external 
function:Multicall3.getEthBalance(address) (contracts/core/test/Multicall3.sol#221-223)is public and can be replaced with external 
function:Multicall3.getLastBlockHash() (contracts/core/test/Multicall3.sol#226-230)is public and can be replaced with external 
function:Multicall3.getBasefee() (contracts/core/test/Multicall3.sol#234-236)is public and can be replaced with external 
function:Multicall3.getChainId() (contracts/core/test/Multicall3.sol#239-241)is public and can be replaced with external 
function:RewardsPoolControllerMock.initialize(address) (contracts/core/test/RewardsPoolControllerMock.sol#23-26)is public and can be replaced with external 
function:SDLPoolMock.setEffectiveBalance(address,uint256) (contracts/core/test/SDLPoolMock.sol#15-17)is public and can be replaced with external 
function:Strategy.getTotalDeposits() (contracts/core/base/Strategy.sol#70)is public and can be replaced with external 
function:Strategy.getMaxDeposits() (contracts/core/base/Strategy.sol#76)is public and can be replaced with external 
function:Strategy.getMinDeposits() (contracts/core/base/Strategy.sol#82)is public and can be replaced with external 
function:StrategyMock.initialize(address,address,uint256,uint256) (contracts/core/test/StrategyMock.sol#28-38)is public and can be replaced with external 
function:StrategyMock.createRewardsPool(address) (contracts/core/test/StrategyMock.sol#109-113)is public and can be replaced with external 
function:ERC677.transferAndCall(address,uint256,bytes) (contracts/core/tokens/base/ERC677.sol#17-27)is public and can be replaced with external 
function:CCIPOnRampMock.getPoolBySourceToken(address) (contracts/core/test/chainlink/CCIPOnRampMock.sol#45-47)is public and can be replaced with external 
function:RewardsPoolControllerV1.distributeTokens(address[]) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#126-130)is public and can be replaced with external 
function:RewardsPoolControllerV1.withdrawRewards(address[]) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#166-171)is public and can be replaced with external 
function:RewardsPoolControllerV1.addToken(address,address) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#178-189)is public and can be replaced with external 
function:DelegatorPool.initialize(address,string,string,address[]) (contracts/core/test/deprecated/DelegatorPool.sol#56-74)is public and can be replaced with external 
function:OwnersRewardsPoolV1.withdraw(uint256) (contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#33-42)is public and can be replaced with external 
function:PoolOwnersV1.totalSupply() (contracts/core/test/deprecated/PoolOwnersV1.sol#85-87)is public and can be replaced with external 
function:RewardsPoolControllerV1.__RewardsPoolController_init(string,string) (contracts/core/test/deprecated/RewardsPoolControllerV1.sol#30-37)is public and can be replaced with external 
function:RewardsPoolV1.withdraw(uint256) (contracts/core/test/deprecated/RewardsPoolV1.sol#60-63)is public and can be replaced with external 
function:LPLMigration.onTokenTransfer(address,uint256,bytes) (contracts/core/tokens/LPLMigration.sol#31-39)is public and can be replaced with external 
function:StakingAllowance.mint(address,uint256) (contracts/core/tokens/StakingAllowance.sol#20-22)is public and can be replaced with external 
function:StakingAllowance.mintToContract(address,address,uint256,bytes) (contracts/core/tokens/StakingAllowance.sol#30-38)is public and can be replaced with external 
function:StakingAllowance.burn(uint256) (contracts/core/tokens/StakingAllowance.sol#44-46)is public and can be replaced with external 
function:StakingAllowance.burnFrom(address,uint256) (contracts/core/tokens/StakingAllowance.sol#52-55)is public and can be replaced with external 
function:ERC677Upgradeable.__ERC677_init(string,string,uint256) (contracts/core/tokens/base/ERC677Upgradeable.sol#9-16)is public and can be replaced with external 
function:EthStakingStrategy.initialize(address,address,uint256,uint256,address,bytes32,uint256) (contracts/ethStaking/EthStakingStrategy.sol#69-84)is public and can be replaced with external 
function:OperatorController.staked(address) (contracts/ethStaking/base/OperatorController.sol#88-90)is public and can be replaced with external 
function:OperatorController.totalStaked() (contracts/ethStaking/base/OperatorController.sol#98-100)is public and can be replaced with external 
function:OperatorController.withdrawRewards() (contracts/ethStaking/base/OperatorController.sol#188-190)is public and can be replaced with external 
function:NWLOperatorController.initialize(address,address) (contracts/ethStaking/NWLOperatorController.sol#36-38)is public and can be replaced with external 
function:WLOperatorController.initialize(address,address,address,uint256) (contracts/ethStaking/WLOperatorController.sol#33-42)is public and can be replaced with external 
function:OperatorController.__OperatorController_init(address,address) (contracts/ethStaking/base/OperatorController.sol#74-80)is public and can be replaced with external 
function:OperatorControllerMock.initialize(address,address) (contracts/ethStaking/test/OperatorControllerMock.sol#19-21)is public and can be replaced with external 
function:CommunityVCS.initialize(address,address,address,address,VaultControllerStrategy.Fee[],uint256,uint128,uint128) (contracts/linkStaking/CommunityVCS.sol#36-57)is public and can be replaced with external 
function:Vault.getTotalDeposits() (contracts/linkStaking/base/Vault.sol#79-81)is public and can be replaced with external 
function:CommunityVault.initialize(address,address,address,address) (contracts/linkStaking/CommunityVault.sol#25-32)is public and can be replaced with external 
function:OperatorVCS.initialize(address,address,address,address,VaultControllerStrategy.Fee[],uint256,uint256) (contracts/linkStaking/OperatorVCS.sol#46-76)is public and can be replaced with external 
function:OperatorVCS.setOperatorRewardPercentage(uint256) (contracts/linkStaking/OperatorVCS.sol#274-281)is public and can be replaced with external 
function:OperatorVault.initialize(address,address,address,address,address,address,address) (contracts/linkStaking/OperatorVault.sol#50-72)is public and can be replaced with external 
function:OperatorVault.getPendingRewards() (contracts/linkStaking/OperatorVault.sol#157-165)is public and can be replaced with external 
function:OperatorVault.setRewardsReceiver(address) (contracts/linkStaking/OperatorVault.sol#227-233)is public and can be replaced with external 
function:Vault.__Vault_init(address,address,address,address) (contracts/linkStaking/base/Vault.sol#36-48)is public and can be replaced with external 
function:VaultControllerStrategy.__VaultControllerStrategy_init(address,address,address,address,VaultControllerStrategy.Fee[],uint256) (contracts/linkStaking/base/VaultControllerStrategy.sol#56-78)is public and can be replaced with external 
function:CommunityVaultV2Mock.initializeV2(uint256) (contracts/linkStaking/test/CommunityVaultV2Mock.sol#13-15)is public and can be replaced with external 
function:VCSMock.initialize(address,address,address,address,VaultControllerStrategy.Fee[]) (contracts/linkStaking/test/VCSMock.sol#16-24)is public and can be replaced with external 
function:VaultV1.getTotalDeposits() (contracts/linkStaking/test/deprecated/VaultV1.sol#89)is public and can be replaced with external 
function:VaultV1.getPrincipalDeposits() (contracts/linkStaking/test/deprecated/VaultV1.sol#95-97)is public and can be replaced with external 
function:OperatorVaultV1.initialize(address,address,address,address) (contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#22-30)is public and can be replaced with external 
function:OperatorVaultV1.getTotalDeposits() (contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#41-46)is public and can be replaced with external 
function:VaultV1.__Vault_init(address,address,address) (contracts/linkStaking/test/deprecated/VaultV1.sol#52-62)is public and can be replaced with external 
function:VaultControllerStrategyUpgrade.getVaultDepositLimits() (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#266)is public and can be replaced with external 
function:OperatorVCSUpgrade.initialize(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[],address[]) (contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#21-43)is public and can be replaced with external 
function:VaultControllerStrategyUpgrade.__VaultControllerStrategy_init(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[]) (contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#86-109)is public and can be replaced with external 
function:LiquidSDIndexPool.initialize(string,string,uint256,uint256,LiquidSDIndexPool.Fee[],uint256) (contracts/liquidSDIndex/LiquidSDIndexPool.sol#56-72)is public and can be replaced with external 
function:LSDIndexAdapter.getTotalDeposits() (contracts/liquidSDIndex/base/LSDIndexAdapter.sol#32-34)is public and can be replaced with external 
function:LSDIndexAdapter.getTotalDepositsLSD() (contracts/liquidSDIndex/base/LSDIndexAdapter.sol#40-42)is public and can be replaced with external 
function:LSDIndexAdapter.getLSDByUnderlying(uint256) (contracts/liquidSDIndex/base/LSDIndexAdapter.sol#58-60)is public and can be replaced with external 
function:LSDIndexAdapter.getExchangeRate() (contracts/liquidSDIndex/base/LSDIndexAdapter.sol#66)is public and can be replaced with external 
function:CoinbaseLSDIndexAdapter.initialize(address,address) (contracts/liquidSDIndex/adapters/CoinbaseLSDIndexAdapter.sol#17-19)is public and can be replaced with external 
function:FraxLSDIndexAdapter.initialize(address,address) (contracts/liquidSDIndex/adapters/FraxLSDIndexAdapter.sol#17-19)is public and can be replaced with external 
function:LidoLSDIndexAdapter.initialize(address,address) (contracts/liquidSDIndex/adapters/LidoLSDIndexAdapter.sol#16-18)is public and can be replaced with external 
function:RocketPoolLSDIndexAdapter.initialize(address,address) (contracts/liquidSDIndex/adapters/RocketPoolLSDIndexAdapter.sol#17-19)is public and can be replaced with external 
function:LSDIndexAdapter.__LiquidSDAdapter_init(address,address) (contracts/liquidSDIndex/base/LSDIndexAdapter.sol#20-26)is public and can be replaced with external 
function:LSDIndexAdapterMock.initialize(address,address,uint256) (contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#13-20)is public and can be replaced with external 
Reference:  

OperatorVCSWithdrawExtraRewards(address,uint256) (contracts/linkStaking/OperatorVCS.sol#22) is never emitted in OperatorVCS (contracts/linkStaking/OperatorVCS.sol#11-304)
Reference: https://certik-public-assets.s3.amazonaws.com/CertiK-Audit-for-Kitty-inu.pdf

	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/AuthorizedReceiver.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwner.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/Operator.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/interfaces/AuthorizedReceiverInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/interfaces/ChainlinkRequestInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/interfaces/LinkTokenInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/interfaces/OperatorInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/interfaces/OracleInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/interfaces/OwnableInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/interfaces/WithdrawalInterface.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.6.2<0.8.0. >=0.6.2<0.8.0 (node_modules/@chainlink/contracts/src/v0.7/vendor/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (node_modules/@chainlink/contracts/src/v0.7/vendor/SafeMathChainlink.sol#2)
	Pragma confirmed better, here is pragma solidity^0.7.0. ^0.7.0 (contracts/core/test/chainlink/CLContractImports0.7.sol#2)
Reference:  

	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ConfirmedOwner.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ConfirmedOwnerWithProposal.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IARM.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IAny2EVMMessageReceiver.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IEVM2AnyOnRamp.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IRouter.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IRouterClient.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IWrappedNative.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/pools/IPool.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/Client.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/Internal.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/MerkleMultiProof.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/interfaces/OwnableInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/interfaces/TypeAndVersionInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/shared/access/OwnerIsCreator.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/extensions/draft-IERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/utils/SafeERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/introspection/IERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#5)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableSet.sol#5)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/interfaces/ChainlinkRequestInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/interfaces/ENSInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/interfaces/OperatorInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/interfaces/OracleInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/interfaces/PointerInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/shared/interfaces/LinkTokenInterface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/vendor/BufferChainlink.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.4.19. >=0.4.19 (node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/vendor/ENSResolver.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/interfaces/IERC1967Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/interfaces/draft-IERC1822Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/proxy/beacon/IBeaconUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20PermitUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/StorageSlotUpgradeable.sol#5)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/MerkleProofUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/access/Ownable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/Proxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/security/Pausable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (node_modules/@openzeppelin/contracts/utils/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Context.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Strings.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/cryptography/MerkleProof.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/math/Math.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/core/test/chainlink/CLContractImports0.8.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.15. ^0.8.15 (contracts/governance/StakedotlinkCouncil.sol#3)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (node_modules/solidity-bytes-utils/contracts/BytesLib.sol#9)
Reference:  

DelegatorPool.feeCurve (contracts/core/test/deprecated/DelegatorPool.sol#36) should be constant
StakingPool.liquidityBuffer (contracts/core/StakingPool.sol#24) should be constant
StakingPool.poolIndex (contracts/core/StakingPool.sol#30) should be constant
Reference:  

unnecessary reentrancy lock found inOwnersRewardsPoolV1
	- RewardsPoolV1.updateReward(address) (contracts/core/test/deprecated/RewardsPoolV1.sol#52-54)
Reference: http://
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
An error occurred: A process in the process pool was terminated abruptly while the future was running or pending.
An error occurred: A process in the process pool was terminated abruptly while the future was running or pending.
An error occurred: A process in the process pool was terminated abruptly while the future was running or pending.
An error occurred: A process in the process pool was terminated abruptly while the future was running or pending.
An error occurred: list index out of range
after for recrusion
Summary
 - [arbitrary-transfer](#arbitrary-transfer) (1 results) (Critical)
 - [reentrancy-with-eth-transfer](#reentrancy-with-eth-transfer) (2 results) (Critical)
 - [signature-malleability](#signature-malleability) (1 results) (High)
 - [state-variable-not-initialized](#state-variable-not-initialized) (5 results) (High)
 - [arbitrary-storage-location](#arbitrary-storage-location) (1 results) (Medium)
 - [centralized-risk-medium](#centralized-risk-medium) (8 results) (Medium)
 - [ether-locked](#ether-locked) (3 results) (Medium)
 - [price-manipulation-medium](#price-manipulation-medium) (2 results) (Medium)
 - [transfer-in-loop](#transfer-in-loop) (1 results) (Medium)
 - [uninitialized-local](#uninitialized-local) (30 results) (Medium)
 - [unused-return](#unused-return) (26 results) (Medium)
 - [void-function](#void-function) (23 results) (Medium)
 - [centralized-risk-low](#centralized-risk-low) (14 results) (Low)
 - [pess-event-setter](#pess-event-setter) (235 results) (Low)
 - [initialize-permission](#initialize-permission) (4 results) (Low)
 - [input-validation](#input-validation) (1 results) (Low)
 - [missing-zero-check](#missing-zero-check) (6 results) (Low)
 - [variable-scope](#variable-scope) (5 results) (Low)
 - [reentrancy-same-effect](#reentrancy-same-effect) (1 results) (Low)
 - [unnecessary-boolean-compare](#unnecessary-boolean-compare) (5 results) (Informational)
 - [centralized-risk-informational](#centralized-risk-informational) (151 results) (Informational)
 - [centralized-risk-other](#centralized-risk-other) (81 results) (Informational)
 - [error-msg](#error-msg) (2 results) (Informational)
 - [price-manipulation-info](#price-manipulation-info) (4 results) (Informational)
 - [uncontrolled-resource-consumption](#uncontrolled-resource-consumption) (35 results) (Informational)
 - [no-derived-function](#no-derived-function) (2 results) (Informational)
 - [unnecessary-public-function-modifier](#unnecessary-public-function-modifier) (153 results) (Informational)
 - [unused-event](#unused-event) (1 results) (Informational)
 - [version-only](#version-only) (101 results) (Informational)
 - [state-should-be-constant](#state-should-be-constant) (3 results) (Optimization)
 - [unnecessary-reentrancy-lock](#unnecessary-reentrancy-lock) (1 results) (Optimization)
## arbitrary-transfer
Impact: Critical
Confidence: Medium
 - [ ] ID-0
[Operator.distributeFunds(address[],uint256[])](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L349-L358) sends eth to arbitrary user
	Dangerous calls:
	- [receivers[i].transfer(sendAmount)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L355)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L349-L358


## reentrancy-with-eth-transfer
Impact: Critical
Confidence: Medium
 - [ ] ID-1
Reentrancy in [PoolOwnersV1.exit()](contracts/core/test/deprecated/PoolOwnersV1.sol#L137-L140):
	External calls:
	- [withdraw(balanceOf(msg.sender))](contracts/core/test/deprecated/PoolOwnersV1.sol#L138)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [(success,returndata) = target.call{value: value}(data)](node_modules/@openzeppelin/contracts/utils/Address.sol#L137)
	External calls sending eth:
	- [withdraw(balanceOf(msg.sender))](contracts/core/test/deprecated/PoolOwnersV1.sol#L138)
		- [(success,returndata) = target.call{value: value}(data)](node_modules/@openzeppelin/contracts/utils/Address.sol#L137)

contracts/core/test/deprecated/PoolOwnersV1.sol#L137-L140


 - [ ] ID-2
Reentrancy in [SDLPoolSecondary.executeQueuedOperations(uint256[])](contracts/core/sdlPool/SDLPoolSecondary.sol#L247-L250):
	External calls:
	- [_executeQueuedLockUpdates(msg.sender,_lockIds)](contracts/core/sdlPool/SDLPoolSecondary.sol#L248)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#L122)
		- [(success,returndata) = target.call{value: value}(data)](node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#L135)
	External calls sending eth:
	- [_executeQueuedLockUpdates(msg.sender,_lockIds)](contracts/core/sdlPool/SDLPoolSecondary.sol#L248)
		- [(success,returndata) = target.call{value: value}(data)](node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#L135)

contracts/core/sdlPool/SDLPoolSecondary.sol#L247-L250


## signature-malleability
Impact: High
Confidence: Medium
 - [ ] ID-3
Signature Malleability Found in [DepositContract.deposit(bytes,bytes,bytes,bytes32)](contracts/ethStaking/test/DepositContract.sol#L89-L145)State Variable:[DepositContract.branch](contracts/ethStaking/test/DepositContract.sol#L63)Potential Signature Variable:signature

contracts/ethStaking/test/DepositContract.sol#L89-L145


## state-variable-not-initialized
Impact: High
Confidence: High
 - [ ] ID-4
state variable: [Router.s_offRamps](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L63) not initialized and not written in contract but be used in contract

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L63


 - [ ] ID-5
state variable: [SDLPool.locks](contracts/core/sdlPool/base/SDLPool.sol#L36) not initialized and not written in contract but be used in contract

contracts/core/sdlPool/base/SDLPool.sol#L36


 - [ ] ID-6
state variable: [SDLPool.totalEffectiveBalance](contracts/core/sdlPool/base/SDLPool.sol#L40) not initialized and not written in contract but be used in contract

contracts/core/sdlPool/base/SDLPool.sol#L40


 - [ ] ID-7
state variable: [SDLPool.lastLockId](contracts/core/sdlPool/base/SDLPool.sol#L35) not initialized and not written in contract but be used in contract

contracts/core/sdlPool/base/SDLPool.sol#L35


 - [ ] ID-8
state variable: [OperatorController.totalAssignedValidators](contracts/ethStaking/base/OperatorController.sol#L40) not initialized and not written in contract but be used in contract

contracts/ethStaking/base/OperatorController.sol#L40


## arbitrary-storage-location
Impact: Medium
Confidence: Low
 - [ ] ID-9
Arbitrary Storage Found in [SDLPoolSecondary.handleOutgoingUpdate()](contracts/core/sdlPool/SDLPoolSecondary.sol#L313-L328)

contracts/core/sdlPool/SDLPoolSecondary.sol#L313-L328


## centralized-risk-medium
Impact: Medium
Confidence: High
 - [ ] ID-10
	- [Operator.withdraw(address,uint256)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L295-L302)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L295-L302


 - [ ] ID-11
	- [Router.recoverTokens(address,address,uint256)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L303-L312)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L303-L312


 - [ ] ID-12
	- [RESDLTokenBridge.transferRESDL(uint64,address,uint256,bool,uint256)](contracts/core/ccip/RESDLTokenBridge.sol#L84-L135)

contracts/core/ccip/RESDLTokenBridge.sol#L84-L135


 - [ ] ID-13
	- [RewardsReceiver.withdraw()](contracts/ethStaking/RewardsReceiver.sol#L36-L57)

contracts/ethStaking/RewardsReceiver.sol#L36-L57


 - [ ] ID-14
	- [NWLOperatorController.assignNextValidators(uint256)](contracts/ethStaking/NWLOperatorController.sol#L180-L241)

contracts/ethStaking/NWLOperatorController.sol#L180-L241


 - [ ] ID-15
	- [NWLOperatorController.removeKeyPairs(uint256,uint256,uint256[])](contracts/ethStaking/NWLOperatorController.sol#L100-L148)

contracts/ethStaking/NWLOperatorController.sol#L100-L148


 - [ ] ID-16
	- [DistributionOracle.withdrawLink(uint256)](contracts/core/priorityPool/DistributionOracle.sol#L217-L220)

contracts/core/priorityPool/DistributionOracle.sol#L217-L220


 - [ ] ID-17
	- [EthStakingStrategy.deposit(uint256)](contracts/ethStaking/EthStakingStrategy.sol#L212-L217)

contracts/ethStaking/EthStakingStrategy.sol#L212-L217


## ether-locked
Impact: Medium
Confidence: High
 - [ ] ID-18
Contract locking ether found:
	Contract [DepositContract](contracts/ethStaking/test/DepositContract.sol#L58-L164) has payable functions:
	 - [IDepositContract.deposit(bytes,bytes,bytes,bytes32)](contracts/ethStaking/test/DepositContract.sol#L27-L32)
	 - [DepositContract.deposit(bytes,bytes,bytes,bytes32)](contracts/ethStaking/test/DepositContract.sol#L89-L145)
	But does not have a function to withdraw the ether

contracts/ethStaking/test/DepositContract.sol#L58-L164


 - [ ] ID-19
Contract locking ether found:
	Contract [WrappedNative](contracts/core/test/chainlink/WrappedNative.sol#L6-L12) has payable functions:
	 - [WrappedNative.deposit()](contracts/core/test/chainlink/WrappedNative.sol#L9-L11)
	But does not have a function to withdraw the ether

contracts/core/test/chainlink/WrappedNative.sol#L6-L12


 - [ ] ID-20
Contract locking ether found:
	Contract [EthStakingStrategyMock](contracts/ethStaking/test/EthStakingStrategyMock.sol#L6-L22) has payable functions:
	 - [EthStakingStrategyMock.receive()](contracts/ethStaking/test/EthStakingStrategyMock.sol#L9)
	But does not have a function to withdraw the ether

contracts/ethStaking/test/EthStakingStrategyMock.sol#L6-L22


## price-manipulation-medium
Impact: Medium
Confidence: Medium
 - [ ] ID-21
Potential price manipulation risk:
	- In function withdraw
		-- [toWithdraw = balanceOf(_account)](contracts/core/StakingPool.sol#L104) have potential price manipulated risk from toWithdraw and call None which could influence variable:toWithdraw

contracts/core/StakingPool.sol#L104


 - [ ] ID-22
Potential price manipulation risk:
	- In function _updateReward
		-- [toMint = balanceOf(_account) - super.balanceOf(_account)](contracts/core/test/deprecated/RewardsPoolV1.sol#L79) have potential price manipulated risk from toMint and call None which could influence variable:toMint
	- In function withdraw
		-- [toWithdraw = balanceOf(_account)](contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L51) have potential price manipulated risk from toWithdraw and call None which could influence variable:toWithdraw

contracts/core/test/deprecated/RewardsPoolV1.sol#L79


## transfer-in-loop
Impact: Medium
Confidence: Medium
 - [ ] ID-23
[Operator.distributeFunds(address[],uint256[])](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L349-L358) use transfer in a loop: [receivers[i].transfer(sendAmount)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L355)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L349-L358


## uninitialized-local
Impact: Medium
Confidence: Medium
 - [ ] ID-24
[DepositContract.get_deposit_root().node](contracts/ethStaking/test/DepositContract.sol#L75) is a local variable never initialized

contracts/ethStaking/test/DepositContract.sol#L75


 - [ ] ID-25
[WLOperatorController.getNextValidators(uint256).loopValidatorCount](contracts/ethStaking/WLOperatorController.sol#L300) is a local variable never initialized

contracts/ethStaking/WLOperatorController.sol#L300


 - [ ] ID-26
[SDLPoolCCIPControllerPrimary._distributeRewards(uint64,address[],uint256[]).tokensAdded](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L263) is a local variable never initialized

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L263


 - [ ] ID-27
[StakingPool._updateStrategyRewards(uint256[],bytes).totalFeeCount](contracts/core/StakingPool.sol#L427) is a local variable never initialized

contracts/core/StakingPool.sol#L427


 - [ ] ID-28
[EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]).wlPubkeys](contracts/ethStaking/EthStakingStrategy.sol#L168) is a local variable never initialized

contracts/ethStaking/EthStakingStrategy.sol#L168


 - [ ] ID-29
[OperatorVault.updateDeposits(uint256,address).opRewards](contracts/linkStaking/OperatorVault.sol#L185) is a local variable never initialized

contracts/linkStaking/OperatorVault.sol#L185


 - [ ] ID-30
[WLOperatorController.assignNextValidators(uint256[],uint256[],uint256).maxBatches](contracts/ethStaking/WLOperatorController.sol#L142) is a local variable never initialized

contracts/ethStaking/WLOperatorController.sol#L142


 - [ ] ID-31
[VaultControllerStrategyUpgrade.getPendingFees().totalFees](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L211) is a local variable never initialized

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L211


 - [ ] ID-32
[SDLPoolCCIPControllerPrimary._distributeRewards(uint64,address[],uint256[]).numRewardTokensToTransfer](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L252) is a local variable never initialized

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L252


 - [ ] ID-33
[LiquidSDIndexPool.getRewards().totalFees](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L274) is a local variable never initialized

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L274


 - [ ] ID-34
[WLOperatorController.assignNextValidators(uint256[],uint256[],uint256).totalValidatorCount](contracts/ethStaking/WLOperatorController.sol#L141) is a local variable never initialized

contracts/ethStaking/WLOperatorController.sol#L141


 - [ ] ID-35
[EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]).nwlPubkeys](contracts/ethStaking/EthStakingStrategy.sol#L150) is a local variable never initialized

contracts/ethStaking/EthStakingStrategy.sol#L150


 - [ ] ID-36
[LiquidSDIndexPool.updateRewards().totalFeeAmounts](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L295) is a local variable never initialized

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L295


 - [ ] ID-37
[VaultControllerStrategy.getPendingFees().totalFees](contracts/linkStaking/base/VaultControllerStrategy.sol#L140) is a local variable never initialized

contracts/linkStaking/base/VaultControllerStrategy.sol#L140


 - [ ] ID-38
[VaultControllerStrategyUpgrade.checkUpkeep(bytes).firstNonFullVault](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L151) is a local variable never initialized

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L151


 - [ ] ID-39
[StakingPool._updateStrategyRewards(uint256[],bytes).feesPaidCount](contracts/core/StakingPool.sol#L472) is a local variable never initialized

contracts/core/StakingPool.sol#L472


 - [ ] ID-40
[WLOperatorController.reportStoppedValidators(uint256[],uint256[]).totalNewlyStoppedValidators](contracts/ethStaking/WLOperatorController.sol#L375) is a local variable never initialized

contracts/ethStaking/WLOperatorController.sol#L375


 - [ ] ID-41
[WLOperatorController.assignNextValidators(uint256[],uint256[],uint256).maxBatchOperatorId](contracts/ethStaking/WLOperatorController.sol#L143) is a local variable never initialized

contracts/ethStaking/WLOperatorController.sol#L143


 - [ ] ID-42
[RewardsInitiator.checkUpkeep(bytes).strategiesAdded](contracts/core/RewardsInitiator.sol#L64) is a local variable never initialized

contracts/core/RewardsInitiator.sol#L64


 - [ ] ID-43
[WLOperatorController.getNextValidators(uint256).operatorCount](contracts/ethStaking/WLOperatorController.sol#L297) is a local variable never initialized

contracts/ethStaking/WLOperatorController.sol#L297


 - [ ] ID-44
[LiquidSDIndexPool.getWithdrawalAmounts(uint256).totalTargetDepositDiffs](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L218) is a local variable never initialized

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L218


 - [ ] ID-45
[RewardsInitiator.checkUpkeep(bytes).totalStrategiesToUpdate](contracts/core/RewardsInitiator.sol#L52) is a local variable never initialized

contracts/core/RewardsInitiator.sol#L52


 - [ ] ID-46
[EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]).nwlSignatures](contracts/ethStaking/EthStakingStrategy.sol#L151) is a local variable never initialized

contracts/ethStaking/EthStakingStrategy.sol#L151


 - [ ] ID-47
[EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[]).wlSignatures](contracts/ethStaking/EthStakingStrategy.sol#L169) is a local variable never initialized

contracts/ethStaking/EthStakingStrategy.sol#L169


 - [ ] ID-48
[NWLOperatorController.reportStoppedValidators(uint256[],uint256[],uint256[]).totalNewlyStoppedValidators](contracts/ethStaking/NWLOperatorController.sol#L303) is a local variable never initialized

contracts/ethStaking/NWLOperatorController.sol#L303


 - [ ] ID-49
[LiquidSDIndexPool.removeLSDToken(address,uint256[]).index](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L362) is a local variable never initialized

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L362


 - [ ] ID-50
[SDLPool.getLockIdsByOwner(address).lockIdsFound](contracts/core/sdlPool/base/SDLPool.sol#L180) is a local variable never initialized

contracts/core/sdlPool/base/SDLPool.sol#L180


 - [ ] ID-51
[StakingPool._updateStrategyRewards(uint256[],bytes).totalFeeAmounts](contracts/core/StakingPool.sol#L426) is a local variable never initialized

contracts/core/StakingPool.sol#L426


 - [ ] ID-52
[SDLPoolPrimary.handleIncomingUpdate(uint256,int256).mintStartIndex](contracts/core/sdlPool/SDLPoolPrimary.sol#L236) is a local variable never initialized

contracts/core/sdlPool/SDLPoolPrimary.sol#L236


 - [ ] ID-53
[OperatorController.getAssignedKeys(uint256,uint256).index](contracts/ethStaking/base/OperatorController.sol#L157) is a local variable never initialized

contracts/ethStaking/base/OperatorController.sol#L157


## unused-return
Impact: Medium
Confidence: Medium
 - [ ] ID-54
[CBORChainlink.encodeBytes(BufferChainlink.buffer,bytes)](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L63-L66)have ignores return value in [buf.append(value)](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L65)

node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L63-L66


 - [ ] ID-55
[Chainlink.initialize(Chainlink.Request,bytes32,address,bytes4)](node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L33-L44)have ignores return value in [BufferChainlink.init(self.buf,defaultBufferSize)](node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L39)

node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L33-L44


 - [ ] ID-56
[ERC1967UpgradeUpgradeable._upgradeToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L65-L70)have ignores return value in [AddressUpgradeable.functionDelegateCall(newImplementation,data)](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L68)

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L65-L70


 - [ ] ID-57
[ERC1967Upgrade._upgradeToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L65-L74)have ignores return value in [Address.functionDelegateCall(newImplementation,data)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L72)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L65-L74


 - [ ] ID-58
[CBORChainlink.encodeString(BufferChainlink.buffer,string)](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L78-L81)have ignores return value in [buf.append(bytes(value))](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L80)

node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L78-L81


 - [ ] ID-59
[ERC1967UpgradeUpgradeable._upgradeBeaconToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L156-L162)have ignores return value in [AddressUpgradeable.functionDelegateCall(IBeaconUpgradeable(newBeacon).implementation(),data)](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L160)

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L156-L162


 - [ ] ID-60
[ERC1967Upgrade._upgradeBeaconToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L174-L184)have ignores return value in [Address.functionDelegateCall(IBeacon(newBeacon).implementation(),data)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L182)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L174-L184


 - [ ] ID-61
[Chainlink.setBuffer(Chainlink.Request,bytes)](node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L52-L55)have ignores return value in [BufferChainlink.init(self.buf,data.length)](node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L53)

node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L52-L55


 - [ ] ID-62
[CBORChainlink.encodeBigNum(BufferChainlink.buffer,uint256)](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L68-L71)have ignores return value in [buf.appendUint8(uint8((MAJOR_TYPE_TAG << 5) | TAG_TYPE_BIGNUM))](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L69)

node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L68-L71


 - [ ] ID-63
[CBORChainlink.encodeSignedBigNum(BufferChainlink.buffer,int256)](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L73-L76)have ignores return value in [buf.appendUint8(uint8((MAJOR_TYPE_TAG << 5) | TAG_TYPE_NEGATIVE_BIGNUM))](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L74)

node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L73-L76


 - [ ] ID-64
[CBORChainlink.encodeFixedNumeric(BufferChainlink.buffer,uint8,uint64)](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L21-L37)have ignores return value in [buf.appendUint8(uint8((major << 5) | value))](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L23)

node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L21-L37


 - [ ] ID-65
[CBORChainlink.encodeIndefiniteLengthType(BufferChainlink.buffer,uint8)](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L39-L41)have ignores return value in [buf.appendUint8(uint8((major << 5) | 31))](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L40)

node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L39-L41


 - [ ] ID-66
[OperatorVCSUpgrade.initialize(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[],address[])](contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L21-L43)have ignores return value in [token.approve(vault,type()(uint256).max)](contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L41)

contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L21-L43


 - [ ] ID-67
[EthStakingStrategyMock.depositEther(uint256)](contracts/ethStaking/test/EthStakingStrategyMock.sol#L11-L13)have ignores return value in [INWLOperatorController(nwlOperatorController).assignNextValidators(_totalValidatorCount)](contracts/ethStaking/test/EthStakingStrategyMock.sol#L12)

contracts/ethStaking/test/EthStakingStrategyMock.sol#L11-L13


 - [ ] ID-68
[WrappedTokenBridge.constructor(address,address,address,address)](contracts/core/ccip/WrappedTokenBridge.sol#L60-L74)have ignores return value in [linkToken.approve(_router,type()(uint256).max)](contracts/core/ccip/WrappedTokenBridge.sol#L71)

contracts/core/ccip/WrappedTokenBridge.sol#L60-L74


 - [ ] ID-69
[SDLPoolCCIPController.constructor(address,address,address,address,uint256)](contracts/core/ccip/base/SDLPoolCCIPController.sol#L41-L54)have ignores return value in [linkToken.approve(_router,type()(uint256).max)](contracts/core/ccip/base/SDLPoolCCIPController.sol#L52)

contracts/core/ccip/base/SDLPoolCCIPController.sol#L41-L54


 - [ ] ID-70
[VaultV1.deposit(uint256)](contracts/linkStaking/test/deprecated/VaultV1.sol#L73-L76)have ignores return value in [IERC677(address(token)).transferAndCall(address(stakeController),_amount,0x00)](contracts/linkStaking/test/deprecated/VaultV1.sol#L75)

contracts/linkStaking/test/deprecated/VaultV1.sol#L73-L76


 - [ ] ID-71
[Vault.deposit(uint256)](contracts/linkStaking/base/Vault.sol#L62-L65)have ignores return value in [IERC677(address(token)).transferAndCall(address(stakeController),_amount,)](contracts/linkStaking/base/Vault.sol#L64)

contracts/linkStaking/base/Vault.sol#L62-L65


 - [ ] ID-72
[RewardsPoolWSD.distributeRewards()](contracts/core/RewardsPoolWSD.sol#L49-L60)have ignores return value in [token.transferAndCall(address(wsdToken),balance,0x)](contracts/core/RewardsPoolWSD.sol#L53)

contracts/core/RewardsPoolWSD.sol#L49-L60


 - [ ] ID-73
[DelegatorPool.retireDelegatorPool(address[],address)](contracts/core/test/deprecated/DelegatorPool.sol#L249-L276)have ignores return value in [allowanceToken.approve(_sdlPool,type()(uint256).max)](contracts/core/test/deprecated/DelegatorPool.sol#L251)

contracts/core/test/deprecated/DelegatorPool.sol#L249-L276


 - [ ] ID-74
[OperatorVCSMock.addVault(address)](contracts/linkStaking/test/OperatorVCSMock.sol#L47-L50)have ignores return value in [token.approve(_vault,type()(uint256).max)](contracts/linkStaking/test/OperatorVCSMock.sol#L49)

contracts/linkStaking/test/OperatorVCSMock.sol#L47-L50


 - [ ] ID-75
[OperatorVault.deposit(uint256)](contracts/linkStaking/OperatorVault.sol#L95-L99)have ignores return value in [IERC677(address(token)).transferAndCall(address(stakeController),_amount,)](contracts/linkStaking/OperatorVault.sol#L98)

contracts/linkStaking/OperatorVault.sol#L95-L99


 - [ ] ID-76
[VCSMock.addVaults(address[])](contracts/linkStaking/test/VCSMock.sol#L26-L32)have ignores return value in [token.approve(vault,type()(uint256).max)](contracts/linkStaking/test/VCSMock.sol#L30)

contracts/linkStaking/test/VCSMock.sol#L26-L32


 - [ ] ID-77
[StakingMockV1.migrate(bytes)](contracts/linkStaking/test/StakingMockV1.sol#L75-L77)have ignores return value in [token.transferAndCall(migration,stakedBalances[msg.sender] + baseReward + delegationReward,abi.encode(msg.sender))](contracts/linkStaking/test/StakingMockV1.sol#L76)

contracts/linkStaking/test/StakingMockV1.sol#L75-L77


 - [ ] ID-78
[SDLPoolCCIPControllerPrimary.distributeRewards()](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L56-L93)have ignores return value in [IERC677(token).transferAndCall(wrappedToken,tokenBalance,)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L74)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L56-L93


 - [ ] ID-79
[OperatorController.onTokenTransfer(address,uint256,bytes)](contracts/ethStaking/base/OperatorController.sol#L176-L183)have ignores return value in [sdToken.transferAndCall(address(rewardsPool),_value,0x)](contracts/ethStaking/base/OperatorController.sol#L182)

contracts/ethStaking/base/OperatorController.sol#L176-L183


## void-function
Impact: Medium
Confidence: High
 - [ ] ID-80
function:[ERC20Upgradeable._beforeTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L353)is empty 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L353


 - [ ] ID-81
function:[ERC20._afterTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L378-L382)is empty 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L378-L382


 - [ ] ID-82
function:[ContextUpgradeable.__Context_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L21-L22)is empty 

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L21-L22


 - [ ] ID-83
function:[UUPSUpgradeable.__UUPSUpgradeable_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L26-L27)is empty 

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L26-L27


 - [ ] ID-84
function:[ChainlinkClient.validateChainlinkCallback(bytes32)](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L286-L291)is empty 

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L286-L291


 - [ ] ID-85
function:[ERC721._afterTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L449-L453)is empty 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L449-L453


 - [ ] ID-86
function:[ContextUpgradeable.__Context_init()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L18-L19)is empty 

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L18-L19


 - [ ] ID-87
function:[ERC20Upgradeable._afterTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L369)is empty 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L369


 - [ ] ID-88
function:[ERC1967UpgradeUpgradeable.__ERC1967Upgrade_init()](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L20-L21)is empty 

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L20-L21


 - [ ] ID-89
function:[UUPSUpgradeable.__UUPSUpgradeable_init()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L23-L24)is empty 

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L23-L24


 - [ ] ID-90
function:[ERC20._beforeTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L358-L362)is empty 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L358-L362


 - [ ] ID-91
function:[ERC1967UpgradeUpgradeable.__ERC1967Upgrade_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L23-L24)is empty 

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L23-L24


 - [ ] ID-92
function:[ERC721._beforeTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L432-L436)is empty 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L432-L436


 - [ ] ID-93
function:[StakingRewardsPool._authorizeUpgrade(address)](contracts/core/base/StakingRewardsPool.sol#L209)is empty 

contracts/core/base/StakingRewardsPool.sol#L209


 - [ ] ID-94
function:[LSDIndexAdapter._authorizeUpgrade(address)](contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L71)is empty 

contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L71


 - [ ] ID-95
function:[RewardsPoolController._authorizeUpgrade(address)](contracts/core/base/RewardsPoolController.sol#L197)is empty 

contracts/core/base/RewardsPoolController.sol#L197


 - [ ] ID-96
function:[OperatorController._authorizeUpgrade(address)](contracts/ethStaking/base/OperatorController.sol#L434)is empty 

contracts/ethStaking/base/OperatorController.sol#L434


 - [ ] ID-97
function:[Strategy._authorizeUpgrade(address)](contracts/core/base/Strategy.sol#L84)is empty 

contracts/core/base/Strategy.sol#L84


 - [ ] ID-98
function:[CCIPTokenPoolMock.lockOrBurn()](contracts/core/test/chainlink/CCIPTokenPoolMock.sol#L21)is empty 

contracts/core/test/chainlink/CCIPTokenPoolMock.sol#L21


 - [ ] ID-99
function:[Vault._authorizeUpgrade(address)](contracts/linkStaking/base/Vault.sol#L102)is empty 

contracts/linkStaking/base/Vault.sol#L102


 - [ ] ID-100
function:[RewardsPoolControllerV1._authorizeUpgrade(address)](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L221)is empty 

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L221


 - [ ] ID-101
function:[PriorityPool._authorizeUpgrade(address)](contracts/core/priorityPool/PriorityPool.sol#L572)is empty 

contracts/core/priorityPool/PriorityPool.sol#L572


 - [ ] ID-102
function:[VaultV1._authorizeUpgrade(address)](contracts/linkStaking/test/deprecated/VaultV1.sol#L107)is empty 

contracts/linkStaking/test/deprecated/VaultV1.sol#L107


## centralized-risk-low
Impact: Low
Confidence: High
 - [ ] ID-103
	- [EthStakingStrategy.updateDeposits(bytes)](contracts/ethStaking/EthStakingStrategy.sol#L246-L280)

contracts/ethStaking/EthStakingStrategy.sol#L246-L280


 - [ ] ID-104
	- [NWLOperatorController.reportKeyPairValidation(uint256,bool)](contracts/ethStaking/NWLOperatorController.sol#L155-L172)

contracts/ethStaking/NWLOperatorController.sol#L155-L172


 - [ ] ID-105
	- [RewardsReceiver.setWithdrawalLimits(uint256,uint256)](contracts/ethStaking/RewardsReceiver.sol#L64-L69)

contracts/ethStaking/RewardsReceiver.sol#L64-L69


 - [ ] ID-106
	- [OperatorController.setOperatorOwner(uint256,address)](contracts/ethStaking/base/OperatorController.sol#L231-L241)

contracts/ethStaking/base/OperatorController.sol#L231-L241


 - [ ] ID-107
	- [EthStakingStrategy.depositEther(uint256,uint256,uint256[],uint256[])](contracts/ethStaking/EthStakingStrategy.sol#L135-L206)

contracts/ethStaking/EthStakingStrategy.sol#L135-L206


 - [ ] ID-108
	- [Strategy.__Strategy_init(address,address)](contracts/core/base/Strategy.sol#L20-L25)

contracts/core/base/Strategy.sol#L20-L25


 - [ ] ID-109
	- [OperatorController.__OperatorController_init(address,address)](contracts/ethStaking/base/OperatorController.sol#L74-L80)

contracts/ethStaking/base/OperatorController.sol#L74-L80


 - [ ] ID-110
	- [OperatorController.disableOperator(uint256)](contracts/ethStaking/base/OperatorController.sol#L248-L265)

contracts/ethStaking/base/OperatorController.sol#L248-L265


 - [ ] ID-111
	- [OperatorController.setOperatorName(uint256,string)](contracts/ethStaking/base/OperatorController.sol#L219-L222)

contracts/ethStaking/base/OperatorController.sol#L219-L222


 - [ ] ID-112
	- [RESDLTokenBridge.setExtraArgs(uint64,bytes)](contracts/core/ccip/RESDLTokenBridge.sol#L161-L164)

contracts/core/ccip/RESDLTokenBridge.sol#L161-L164


 - [ ] ID-113
	- [OperatorController.initiateKeyPairValidation(address,uint256)](contracts/ethStaking/base/OperatorController.sol#L205-L213)

contracts/ethStaking/base/OperatorController.sol#L205-L213


 - [ ] ID-114
	- [OperatorController.setRewardsPool(address)](contracts/ethStaking/base/OperatorController.sol#L289-L292)

contracts/ethStaking/base/OperatorController.sol#L289-L292


 - [ ] ID-115
	- [EthStakingStrategy.reportBeaconState(uint256,uint256,uint256)](contracts/ethStaking/EthStakingStrategy.sol#L96-L124)

contracts/ethStaking/EthStakingStrategy.sol#L96-L124


 - [ ] ID-116
	- [NWLOperatorController.reportStoppedValidators(uint256[],uint256[],uint256[])](contracts/ethStaking/NWLOperatorController.sol#L293-L342)

contracts/ethStaking/NWLOperatorController.sol#L293-L342


## pess-event-setter
Impact: Low
Confidence: Medium
 - [ ] ID-117
Setter function [Operator._canSetAuthorizedSenders()](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L525-L527) does not emit an event

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L525-L527


 - [ ] ID-118
Setter function [Operator._verifyOracleRequestAndProcessPayment(bytes32,uint256,address,bytes4,uint256,uint256)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L470-L483) does not emit an event

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L470-L483


 - [ ] ID-119
Setter function [Operator._verifyAndProcessOracleRequest(address,uint256,address,bytes4,uint256,uint256)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L444-L460) does not emit an event

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L444-L460


 - [ ] ID-120
Setter function [Operator.transferOwnableContracts(address[],address)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L235-L240) does not emit an event

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L235-L240


 - [ ] ID-121
Setter function [Operator.ownerTransferAndCall(address,uint256,bytes)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L332-L338) does not emit an event

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L332-L338


 - [ ] ID-122
Setter function [Operator.acceptAuthorizedReceivers(address[],address[])](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L281-L287) does not emit an event

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L281-L287


 - [ ] ID-123
Setter function [Operator.withdraw(address,uint256)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L295-L302) does not emit an event

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L295-L302


 - [ ] ID-124
Setter function [Operator.ownerForward(address,bytes)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L319-L323) does not emit an event

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L319-L323


 - [ ] ID-125
Setter function [ERC721URIStorage._setTokenURI(uint256,string)](node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L45-L48) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L45-L48


 - [ ] ID-126
Setter function [EnumerableMap.set(EnumerableMap.UintToAddressMap,uint256,address)](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L268-L274) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L268-L274


 - [ ] ID-127
Setter function [OwnableUpgradeable.renounceOwnership()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68


 - [ ] ID-128
Setter function [PausableUpgradeable.__Pausable_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#L38-L40) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#L38-L40


 - [ ] ID-129
Setter function [OwnableUpgradeable.transferOwnership(address)](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77


 - [ ] ID-130
Setter function [ChainlinkClient.validateChainlinkCallback(bytes32)](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L286-L291) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L286-L291


 - [ ] ID-131
Setter function [EnumerableMap.set(EnumerableMap.Bytes32ToUintMap,bytes32,uint256)](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L452-L458) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L452-L458


 - [ ] ID-132
Setter function [ERC20.approve(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L136-L140) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L136-L140


 - [ ] ID-133
Setter function [ERC20.decreaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L201-L210) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L201-L210


 - [ ] ID-134
Setter function [Router.onlyOffRamp(uint64)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L320-L324) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L320-L324


 - [ ] ID-135
Setter function [CCIPReceiver.onlyRouter()](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L48-L51) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L48-L51


 - [ ] ID-136
Setter function [ERC20Upgradeable.__ERC20_init_unchained(string,string)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L59-L62) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L59-L62


 - [ ] ID-137
Setter function [CCIPReceiver.ccipReceive(Client.Any2EVMMessage)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L27-L29) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L27-L29


 - [ ] ID-138
Setter function [ERC20Upgradeable.increaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L182-L186) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L182-L186


 - [ ] ID-139
Setter function [ChainlinkClient.setPublicChainlinkToken()](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L226-L228) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L226-L228


 - [ ] ID-140
Setter function [ERC20.transfer(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L113-L117) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L113-L117


 - [ ] ID-141
Setter function [Chainlink.setBuffer(Chainlink.Request,bytes)](node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L52-L55) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L52-L55


 - [ ] ID-142
Setter function [Router.recoverTokens(address,address,uint256)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L303-L312) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L303-L312


 - [ ] ID-143
Setter function [OwnableUpgradeable.__Ownable_init()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L29-L31) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L29-L31


 - [ ] ID-144
Setter function [Router.ccipSend(uint64,Client.EVM2AnyMessage)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L104-L141) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L104-L141


 - [ ] ID-145
Setter function [EnumerableMap.set(EnumerableMap.AddressToUintMap,address,uint256)](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L360-L366) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L360-L366


 - [ ] ID-146
Setter function [EnumerableMap.set(EnumerableMap.UintToUintMap,uint256,uint256)](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L176-L182) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L176-L182


 - [ ] ID-147
Setter function [Router.setWrappedNative(address)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L228-L230) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L228-L230


 - [ ] ID-148
Setter function [EnumerableMap.set(EnumerableMap.Bytes32ToBytes32Map,bytes32,bytes32)](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L73-L80) does not emit an event

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L73-L80


 - [ ] ID-149
Setter function [ERC20Upgradeable.approve(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L141-L145) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L141-L145


 - [ ] ID-150
Setter function [OwnableUpgradeable.onlyOwner()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L40-L43) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L40-L43


 - [ ] ID-151
Setter function [OwnableUpgradeable._checkOwner()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L55-L57) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L55-L57


 - [ ] ID-152
Setter function [ChainlinkClient.sendOperatorRequestTo(address,Chainlink.Request,uint256)](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L135-L153) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L135-L153


 - [ ] ID-153
Setter function [ChainlinkClient.setChainlinkOracle(address)](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L210-L212) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L210-L212


 - [ ] ID-154
Setter function [ChainlinkClient.sendChainlinkRequestTo(address,Chainlink.Request,uint256)](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L91-L110) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L91-L110


 - [ ] ID-155
Setter function [ERC20.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L158-L167) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L158-L167


 - [ ] ID-156
Setter function [ERC20Upgradeable.decreaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L202-L211) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L202-L211


 - [ ] ID-157
Setter function [ChainlinkClient.useChainlinkWithENS(address,bytes32)](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L262-L269) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L262-L269


 - [ ] ID-158
Setter function [ERC20.increaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L181-L185) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L181-L185


 - [ ] ID-159
Setter function [ERC20Upgradeable.transfer(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L118-L122) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L118-L122


 - [ ] ID-160
Setter function [ReentrancyGuard.nonReentrant()](node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol#L50-L62) does not emit an event

node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol#L50-L62


 - [ ] ID-161
Setter function [ERC721URIStorage._burn(uint256)](node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L55-L61) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L55-L61


 - [ ] ID-162
Setter function [OwnableUpgradeable.__Ownable_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L33-L35) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L33-L35


 - [ ] ID-163
Setter function [ChainlinkClient.setChainlinkToken(address)](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L218-L220) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L218-L220


 - [ ] ID-164
Setter function [StakedotlinkCouncil.mint(address,uint256)](contracts/governance/StakedotlinkCouncil.sol#L90-L92) does not emit an event

contracts/governance/StakedotlinkCouncil.sol#L90-L92


 - [ ] ID-165
Setter function [ERC20Upgradeable.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L163-L168) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L163-L168


 - [ ] ID-166
Setter function [ChainlinkClient.addChainlinkExternalRequest(address,bytes32)](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L252-L254) does not emit an event

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L252-L254


 - [ ] ID-167
Setter function [StakingAllowance.transferAndCallWithSender(address,address,uint256,bytes)](contracts/core/tokens/StakingAllowance.sol#L65-L75) does not emit an event

contracts/core/tokens/StakingAllowance.sol#L65-L75


 - [ ] ID-168
Setter function [OperatorVault.setOperator(address)](contracts/linkStaking/OperatorVault.sol#L211-L215) does not emit an event

contracts/linkStaking/OperatorVault.sol#L211-L215


 - [ ] ID-169
Setter function [RewardsInitiator.updateRewards(uint256[],bytes)](contracts/core/RewardsInitiator.sol#L38-L42) does not emit an event

contracts/core/RewardsInitiator.sol#L38-L42


 - [ ] ID-170
Setter function [EthStakingStrategy.initialize(address,address,uint256,uint256,address,bytes32,uint256)](contracts/ethStaking/EthStakingStrategy.sol#L69-L84) does not emit an event

contracts/ethStaking/EthStakingStrategy.sol#L69-L84


 - [ ] ID-171
Setter function [OperatorVault.initialize(address,address,address,address,address,address,address)](contracts/linkStaking/OperatorVault.sol#L50-L72) does not emit an event

contracts/linkStaking/OperatorVault.sol#L50-L72


 - [ ] ID-172
Setter function [StakingAllowance.burn(uint256)](contracts/core/tokens/StakingAllowance.sol#L44-L46) does not emit an event

contracts/core/tokens/StakingAllowance.sol#L44-L46


 - [ ] ID-173
Setter function [PoolAllowanceV1.burnAllowance(address,uint256)](contracts/core/test/deprecated/PoolAllowanceV1.sol#L40-L42) does not emit an event

contracts/core/test/deprecated/PoolAllowanceV1.sol#L40-L42


 - [ ] ID-174
Setter function [DepositController.depositEther(bytes32,bytes32,bytes32,uint256,uint256,uint256[],uint256[])](contracts/ethStaking/DepositController.sol#L46-L64) does not emit an event

contracts/ethStaking/DepositController.sol#L46-L64


 - [ ] ID-175
Setter function [OperatorVCSUpgrade.setOperator(uint256,address)](contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L92-L94) does not emit an event

contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L92-L94


 - [ ] ID-176
Setter function [EthStakingStrategyMock.nwlWithdraw(address,uint256)](contracts/ethStaking/test/EthStakingStrategyMock.sol#L15-L17) does not emit an event

contracts/ethStaking/test/EthStakingStrategyMock.sol#L15-L17


 - [ ] ID-177
Setter function [SDLPoolCCIPControllerPrimary.setRewardsInitiator(address)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L234-L236) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L234-L236


 - [ ] ID-178
Setter function [StakingMockV1.setDelegationReward(uint256)](contracts/linkStaking/test/StakingMockV1.sol#L87-L89) does not emit an event

contracts/linkStaking/test/StakingMockV1.sol#L87-L89


 - [ ] ID-179
Setter function [OperatorVCS.withdrawOperatorRewards(address,uint256)](contracts/linkStaking/OperatorVCS.sol#L102-L113) does not emit an event

contracts/linkStaking/OperatorVCS.sol#L102-L113


 - [ ] ID-180
Setter function [StakingPool.reorderStrategies(uint256[])](contracts/core/StakingPool.sol#L261-L274) does not emit an event

contracts/core/StakingPool.sol#L261-L274


 - [ ] ID-181
Setter function [LiquidSDIndexPool.initialize(string,string,uint256,uint256,LiquidSDIndexPool.Fee[],uint256)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L56-L72) does not emit an event

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L56-L72


 - [ ] ID-182
Setter function [SDLPoolCCIPControllerMock.onlyBridge()](contracts/core/test/SDLPoolCCIPControllerMock.sol#L19-L22) does not emit an event

contracts/core/test/SDLPoolCCIPControllerMock.sol#L19-L22


 - [ ] ID-183
Setter function [RocketPoolLSDIndexAdapter.initialize(address,address)](contracts/liquidSDIndex/adapters/RocketPoolLSDIndexAdapter.sol#L17-L19) does not emit an event

contracts/liquidSDIndex/adapters/RocketPoolLSDIndexAdapter.sol#L17-L19


 - [ ] ID-184
Setter function [StakingMockV1.setBaseReward(uint256)](contracts/linkStaking/test/StakingMockV1.sol#L79-L81) does not emit an event

contracts/linkStaking/test/StakingMockV1.sol#L79-L81


 - [ ] ID-185
Setter function [StakingMock.setDepositLimits(uint256,uint256)](contracts/linkStaking/test/StakingMock.sol#L98-L101) does not emit an event

contracts/linkStaking/test/StakingMock.sol#L98-L101


 - [ ] ID-186
Setter function [WLOperatorController.addKeyPairs(uint256,uint256,bytes,bytes)](contracts/ethStaking/WLOperatorController.sol#L60-L68) does not emit an event

contracts/ethStaking/WLOperatorController.sol#L60-L68


 - [ ] ID-187
Setter function [WrappedSDTokenMock.unwrap(uint256)](contracts/core/test/WrappedSDTokenMock.sol#L38-L43) does not emit an event

contracts/core/test/WrappedSDTokenMock.sol#L38-L43


 - [ ] ID-188
Setter function [EthStakingStrategy.nwlWithdraw(address,uint256)](contracts/ethStaking/EthStakingStrategy.sol#L234-L237) does not emit an event

contracts/ethStaking/EthStakingStrategy.sol#L234-L237


 - [ ] ID-189
Setter function [OperatorVCSMock.deposit(uint256)](contracts/linkStaking/test/OperatorVCSMock.sol#L33-L36) does not emit an event

contracts/linkStaking/test/OperatorVCSMock.sol#L33-L36


 - [ ] ID-190
Setter function [NWLOperatorController.addKeyPairs(uint256,uint256,bytes,bytes)](contracts/ethStaking/NWLOperatorController.sol#L83-L92) does not emit an event

contracts/ethStaking/NWLOperatorController.sol#L83-L92


 - [ ] ID-191
Setter function [NWLOperatorController.initialize(address,address)](contracts/ethStaking/NWLOperatorController.sol#L36-L38) does not emit an event

contracts/ethStaking/NWLOperatorController.sol#L36-L38


 - [ ] ID-192
Setter function [PriorityPool.pauseForUpdate()](contracts/core/priorityPool/PriorityPool.sol#L429-L431) does not emit an event

contracts/core/priorityPool/PriorityPool.sol#L429-L431


 - [ ] ID-193
Setter function [PFAlertsControllerMock.raiseAlert(address)](contracts/linkStaking/test/PFAlertsControllerMock.sol#L17-L19) does not emit an event

contracts/linkStaking/test/PFAlertsControllerMock.sol#L17-L19


 - [ ] ID-194
Setter function [WLOperatorController.addOperator(string)](contracts/ethStaking/WLOperatorController.sol#L48-L51) does not emit an event

contracts/ethStaking/WLOperatorController.sol#L48-L51


 - [ ] ID-195
Setter function [MerkleDistributor.updateDistributions(address[],bytes32[],uint256[],uint256[])](contracts/airdrop/MerkleDistributor.sol#L102-L116) does not emit an event

contracts/airdrop/MerkleDistributor.sol#L102-L116


 - [ ] ID-196
Setter function [StakingAllowance.burnFrom(address,uint256)](contracts/core/tokens/StakingAllowance.sol#L52-L55) does not emit an event

contracts/core/tokens/StakingAllowance.sol#L52-L55


 - [ ] ID-197
Setter function [SDLPoolMock.setEffectiveBalance(address,uint256)](contracts/core/test/SDLPoolMock.sol#L15-L17) does not emit an event

contracts/core/test/SDLPoolMock.sol#L15-L17


 - [ ] ID-198
Setter function [WrappedSDTokenMock.onTokenTransfer(address,uint256,bytes)](contracts/core/test/WrappedSDTokenMock.sol#L24-L32) does not emit an event

contracts/core/test/WrappedSDTokenMock.sol#L24-L32


 - [ ] ID-199
Setter function [SDLPoolSecondary.initialize(string,string,address,address,uint256)](contracts/core/sdlPool/SDLPoolSecondary.sol#L66-L79) does not emit an event

contracts/core/sdlPool/SDLPoolSecondary.sol#L66-L79


 - [ ] ID-200
Setter function [StakingPool.deposit(address,uint256)](contracts/core/StakingPool.sol#L80-L88) does not emit an event

contracts/core/StakingPool.sol#L80-L88


 - [ ] ID-201
Setter function [OperatorVaultV1.setOperator(address)](contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L61-L64) does not emit an event

contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L61-L64


 - [ ] ID-202
Setter function [SDLDependentMock.updateSDLBalance(address,uint256)](contracts/core/test/SDLDependentMock.sol#L11-L13) does not emit an event

contracts/core/test/SDLDependentMock.sol#L11-L13


 - [ ] ID-203
Setter function [DelegatorPool.totalStaked()](contracts/core/test/deprecated/DelegatorPool.sol#L143-L145) does not emit an event

contracts/core/test/deprecated/DelegatorPool.sol#L143-L145


 - [ ] ID-204
Setter function [DelegatorPool.onTokenTransfer(address,uint256,bytes)](contracts/core/test/deprecated/DelegatorPool.sol#L82-L104) does not emit an event

contracts/core/test/deprecated/DelegatorPool.sol#L82-L104


 - [ ] ID-205
Setter function [DistributionOracle.fulfillRequest(bytes32,bytes32,bytes32,uint256,uint256)](contracts/core/priorityPool/DistributionOracle.sol#L163-L177) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L163-L177


 - [ ] ID-206
Setter function [OperatorVCS.initialize(address,address,address,address,VaultControllerStrategy.Fee[],uint256,uint256)](contracts/linkStaking/OperatorVCS.sol#L46-L76) does not emit an event

contracts/linkStaking/OperatorVCS.sol#L46-L76


 - [ ] ID-207
Setter function [PriorityPool.setDistributionOracle(address)](contracts/core/priorityPool/PriorityPool.sol#L468-L470) does not emit an event

contracts/core/priorityPool/PriorityPool.sol#L468-L470


 - [ ] ID-208
Setter function [SDLPoolPrimary.onTokenTransfer(address,uint256,bytes)](contracts/core/sdlPool/SDLPoolPrimary.sol#L61-L80) does not emit an event

contracts/core/sdlPool/SDLPoolPrimary.sol#L61-L80


 - [ ] ID-209
Setter function [PriorityPool.initialize(address,address,address,uint128,uint128)](contracts/core/priorityPool/PriorityPool.sol#L93-L110) does not emit an event

contracts/core/priorityPool/PriorityPool.sol#L93-L110


 - [ ] ID-210
Setter function [EthStakingStrategy.updateDeposits(bytes)](contracts/ethStaking/EthStakingStrategy.sol#L246-L280) does not emit an event

contracts/ethStaking/EthStakingStrategy.sol#L246-L280


 - [ ] ID-211
Setter function [StakingPool.strategyWithdraw(uint256,uint256)](contracts/core/StakingPool.sol#L133-L136) does not emit an event

contracts/core/StakingPool.sol#L133-L136


 - [ ] ID-212
Setter function [SDLPoolCCIPControllerPrimary.handleOutgoingRESDL(uint64,address,uint256)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L103-L116) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L103-L116


 - [ ] ID-213
Setter function [MerkleDistributor.withdrawUnclaimedTokens(address,bytes32,uint256)](contracts/airdrop/MerkleDistributor.sol#L208-L224) does not emit an event

contracts/airdrop/MerkleDistributor.sol#L208-L224


 - [ ] ID-214
Setter function [StakingMock.removePrincipal(address,uint256)](contracts/linkStaking/test/StakingMock.sol#L77-L80) does not emit an event

contracts/linkStaking/test/StakingMock.sol#L77-L80


 - [ ] ID-215
Setter function [WrappedTokenBridge.recoverTokens(address[],address)](contracts/core/ccip/WrappedTokenBridge.sol#L140-L147) does not emit an event

contracts/core/ccip/WrappedTokenBridge.sol#L140-L147


 - [ ] ID-216
Setter function [PoolAllowanceV1.onlyPoolOwners()](contracts/core/test/deprecated/PoolAllowanceV1.sol#L21-L24) does not emit an event

contracts/core/test/deprecated/PoolAllowanceV1.sol#L21-L24


 - [ ] ID-217
Setter function [CoinbaseLSDIndexAdapter.initialize(address,address)](contracts/liquidSDIndex/adapters/CoinbaseLSDIndexAdapter.sol#L17-L19) does not emit an event

contracts/liquidSDIndex/adapters/CoinbaseLSDIndexAdapter.sol#L17-L19


 - [ ] ID-218
Setter function [DistributionOracle.rejectManualVerificationAndRetry()](contracts/core/priorityPool/DistributionOracle.sol#L197-L201) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L197-L201


 - [ ] ID-219
Setter function [SDLPoolCCIPControllerMock.setRESDLTokenBridge(address)](contracts/core/test/SDLPoolCCIPControllerMock.sol#L51-L53) does not emit an event

contracts/core/test/SDLPoolCCIPControllerMock.sol#L51-L53


 - [ ] ID-220
Setter function [CommunityVCS.addVaults(uint256)](contracts/linkStaking/CommunityVCS.sol#L106-L108) does not emit an event

contracts/linkStaking/CommunityVCS.sol#L106-L108


 - [ ] ID-221
Setter function [LidoLSDIndexAdapter.initialize(address,address)](contracts/liquidSDIndex/adapters/LidoLSDIndexAdapter.sol#L16-L18) does not emit an event

contracts/liquidSDIndex/adapters/LidoLSDIndexAdapter.sol#L16-L18


 - [ ] ID-222
Setter function [SDLPoolPrimary.extendLockDuration(uint256,uint64)](contracts/core/sdlPool/SDLPoolPrimary.sol#L91-L94) does not emit an event

contracts/core/sdlPool/SDLPoolPrimary.sol#L91-L94


 - [ ] ID-223
Setter function [SDLPoolCCIPControllerPrimary.approveRewardTokens(address[])](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L190-L195) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L190-L195


 - [ ] ID-224
Setter function [RewardsPoolControllerMock.stake(uint256)](contracts/core/test/RewardsPoolControllerMock.sol#L36-L40) does not emit an event

contracts/core/test/RewardsPoolControllerMock.sol#L36-L40


 - [ ] ID-225
Setter function [DelegatorPool.setLockedApproval(address,uint256)](contracts/core/test/deprecated/DelegatorPool.sol#L200-L203) does not emit an event

contracts/core/test/deprecated/DelegatorPool.sol#L200-L203


 - [ ] ID-226
Setter function [OperatorVCS.togglePreRelease()](contracts/linkStaking/OperatorVCS.sol#L287-L289) does not emit an event

contracts/linkStaking/OperatorVCS.sol#L287-L289


 - [ ] ID-227
Setter function [SDLPoolCCIPControllerMock.handleOutgoingRESDL(uint64,address,uint256)](contracts/core/test/SDLPoolCCIPControllerMock.sol#L29-L35) does not emit an event

contracts/core/test/SDLPoolCCIPControllerMock.sol#L29-L35


 - [ ] ID-228
Setter function [LinkPoolNFT.setBaseURI(string)](contracts/core/tokens/LinkPoolNFT.sol#L37-L39) does not emit an event

contracts/core/tokens/LinkPoolNFT.sol#L37-L39


 - [ ] ID-229
Setter function [OperatorVCS.setRewardsReceiver(uint256,address)](contracts/linkStaking/OperatorVCS.sol#L261-L263) does not emit an event

contracts/linkStaking/OperatorVCS.sol#L261-L263


 - [ ] ID-230
Setter function [OwnersRewardsPoolV1.onTokenTransfer(address,uint256,bytes)](contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L65-L72) does not emit an event

contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L65-L72


 - [ ] ID-231
Setter function [NWLOperatorController.assignNextValidators(uint256)](contracts/ethStaking/NWLOperatorController.sol#L180-L241) does not emit an event

contracts/ethStaking/NWLOperatorController.sol#L180-L241


 - [ ] ID-232
Setter function [StakingMockV1.setActive(bool)](contracts/linkStaking/test/StakingMockV1.sol#L55-L57) does not emit an event

contracts/linkStaking/test/StakingMockV1.sol#L55-L57


 - [ ] ID-233
Setter function [PriorityPoolMock.setDepositsSinceLastUpdate(uint256)](contracts/core/test/PriorityPoolMock.sol#L40-L42) does not emit an event

contracts/core/test/PriorityPoolMock.sol#L40-L42


 - [ ] ID-234
Setter function [RewardsPoolControllerMock.initialize(address)](contracts/core/test/RewardsPoolControllerMock.sol#L23-L26) does not emit an event

contracts/core/test/RewardsPoolControllerMock.sol#L23-L26


 - [ ] ID-235
Setter function [RESDLTokenBridge.onlySDLPoolCCIPController()](contracts/core/ccip/RESDLTokenBridge.sol#L71-L74) does not emit an event

contracts/core/ccip/RESDLTokenBridge.sol#L71-L74


 - [ ] ID-236
Setter function [GovernanceController.revokeRole(uint256,address)](contracts/governance/GovernanceController.sol#L138-L140) does not emit an event

contracts/governance/GovernanceController.sol#L138-L140


 - [ ] ID-237
Setter function [StakingPool.addFee(address,uint256)](contracts/core/StakingPool.sol#L281-L284) does not emit an event

contracts/core/StakingPool.sol#L281-L284


 - [ ] ID-238
Setter function [OperatorVCSMock.addVault(address)](contracts/linkStaking/test/OperatorVCSMock.sol#L47-L50) does not emit an event

contracts/linkStaking/test/OperatorVCSMock.sol#L47-L50


 - [ ] ID-239
Setter function [DistributionOracle.withdrawLink(uint256)](contracts/core/priorityPool/DistributionOracle.sol#L217-L220) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L217-L220


 - [ ] ID-240
Setter function [StakingPool.setRewardsInitiator(address)](contracts/core/StakingPool.sol#L384-L386) does not emit an event

contracts/core/StakingPool.sol#L384-L386


 - [ ] ID-241
Setter function [StakingPool.withdraw(address,address,uint256)](contracts/core/StakingPool.sol#L97-L116) does not emit an event

contracts/core/StakingPool.sol#L97-L116


 - [ ] ID-242
Setter function [StakingPool.setPriorityPool(address)](contracts/core/StakingPool.sol#L375-L377) does not emit an event

contracts/core/StakingPool.sol#L375-L377


 - [ ] ID-243
Setter function [SDLPoolPrimary.initialize(string,string,address,address)](contracts/core/sdlPool/SDLPoolPrimary.sol#L30-L41) does not emit an event

contracts/core/sdlPool/SDLPoolPrimary.sol#L30-L41


 - [ ] ID-244
Setter function [DistributionOracle._requestUpdate()](contracts/core/priorityPool/DistributionOracle.sol#L275-L288) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L275-L288


 - [ ] ID-245
Setter function [SDLPoolSecondary.onTokenTransfer(address,uint256,bytes)](contracts/core/sdlPool/SDLPoolSecondary.sol#L132-L151) does not emit an event

contracts/core/sdlPool/SDLPoolSecondary.sol#L132-L151


 - [ ] ID-246
Setter function [StakingMock.setActive(bool)](contracts/linkStaking/test/StakingMock.sol#L90-L92) does not emit an event

contracts/linkStaking/test/StakingMock.sol#L90-L92


 - [ ] ID-247
Setter function [CCIPOffRampMock.executeSingleMessage(bytes32,uint64,bytes,address,Client.EVMTokenAmount[])](contracts/core/test/chainlink/CCIPOffRampMock.sol#L34-L52) does not emit an event

contracts/core/test/chainlink/CCIPOffRampMock.sol#L34-L52


 - [ ] ID-248
Setter function [KeyValidationOracle.reportKeyPairValidation(bytes32,uint256,bool,bool)](contracts/ethStaking/KeyValidationOracle.sol#L68-L79) does not emit an event

contracts/ethStaking/KeyValidationOracle.sol#L68-L79


 - [ ] ID-249
Setter function [WLOperatorController.setOperatorWhitelist(address)](contracts/ethStaking/WLOperatorController.sol#L418-L420) does not emit an event

contracts/ethStaking/WLOperatorController.sol#L418-L420


 - [ ] ID-250
Setter function [WLOperatorController.setBatchSize(uint256)](contracts/ethStaking/WLOperatorController.sol#L410-L412) does not emit an event

contracts/ethStaking/WLOperatorController.sol#L410-L412


 - [ ] ID-251
Setter function [OperatorVaultV1.initialize(address,address,address,address)](contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L22-L30) does not emit an event

contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L22-L30


 - [ ] ID-252
Setter function [MerkleDistributor.pauseForWithdrawal(address)](contracts/airdrop/MerkleDistributor.sol#L232-L239) does not emit an event

contracts/airdrop/MerkleDistributor.sol#L232-L239


 - [ ] ID-253
Setter function [OperatorWhitelist.addWhitelistEntries(address[])](contracts/ethStaking/OperatorWhitelist.sol#L52-L58) does not emit an event

contracts/ethStaking/OperatorWhitelist.sol#L52-L58


 - [ ] ID-254
Setter function [VCSMock.initialize(address,address,address,address,VaultControllerStrategy.Fee[])](contracts/linkStaking/test/VCSMock.sol#L16-L24) does not emit an event

contracts/linkStaking/test/VCSMock.sol#L16-L24


 - [ ] ID-255
Setter function [PoolOwnersV1.exit()](contracts/core/test/deprecated/PoolOwnersV1.sol#L137-L140) does not emit an event

contracts/core/test/deprecated/PoolOwnersV1.sol#L137-L140


 - [ ] ID-256
Setter function [OperatorVault.onlyOperator()](contracts/linkStaking/OperatorVault.sol#L77-L80) does not emit an event

contracts/linkStaking/OperatorVault.sol#L77-L80


 - [ ] ID-257
Setter function [StakingMockV1.raiseAlert()](contracts/linkStaking/test/StakingMockV1.sol#L107-L109) does not emit an event

contracts/linkStaking/test/StakingMockV1.sol#L107-L109


 - [ ] ID-258
Setter function [CCIPOnRampMock.setTokenPool(address,address)](contracts/core/test/chainlink/CCIPOnRampMock.sol#L59-L61) does not emit an event

contracts/core/test/chainlink/CCIPOnRampMock.sol#L59-L61


 - [ ] ID-259
Setter function [OperatorWhitelist.setWLOperatorController(address)](contracts/ethStaking/OperatorWhitelist.sol#L76-L78) does not emit an event

contracts/ethStaking/OperatorWhitelist.sol#L76-L78


 - [ ] ID-260
Setter function [PriorityPool.onlyDistributionOracle()](contracts/core/priorityPool/PriorityPool.sol#L115-L118) does not emit an event

contracts/core/priorityPool/PriorityPool.sol#L115-L118


 - [ ] ID-261
Setter function [LiquidSDIndexPool.setPaused(bool)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L480-L483) does not emit an event

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L480-L483


 - [ ] ID-262
Setter function [EthStakingStrategyMock.setNWLOperatorController(address)](contracts/ethStaking/test/EthStakingStrategyMock.sol#L19-L21) does not emit an event

contracts/ethStaking/test/EthStakingStrategyMock.sol#L19-L21


 - [ ] ID-263
Setter function [DistributionOracle.requestUpdate()](contracts/core/priorityPool/DistributionOracle.sol#L151-L153) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L151-L153


 - [ ] ID-264
Setter function [SDLPoolCCIPControllerSecondary.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken)](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L94-L102) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L94-L102


 - [ ] ID-265
Setter function [OperatorVCS.updateDeposits(bytes)](contracts/linkStaking/OperatorVCS.sol#L155-L216) does not emit an event

contracts/linkStaking/OperatorVCS.sol#L155-L216


 - [ ] ID-266
Setter function [OperatorVCS.onTokenTransfer(address,uint256,bytes)](contracts/linkStaking/OperatorVCS.sol#L84-L90) does not emit an event

contracts/linkStaking/OperatorVCS.sol#L84-L90


 - [ ] ID-267
Setter function [VCSMock.addVaults(address[])](contracts/linkStaking/test/VCSMock.sol#L26-L32) does not emit an event

contracts/linkStaking/test/VCSMock.sol#L26-L32


 - [ ] ID-268
Setter function [StakingPool.removeStrategy(uint256,bytes)](contracts/core/StakingPool.sol#L237-L255) does not emit an event

contracts/core/StakingPool.sol#L237-L255


 - [ ] ID-269
Setter function [PoolOwnersV1._burnAllowance(address)](contracts/core/test/deprecated/PoolOwnersV1.sol#L228-L238) does not emit an event

contracts/core/test/deprecated/PoolOwnersV1.sol#L228-L238


 - [ ] ID-270
Setter function [OperatorVaultV1.onlyOperator()](contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L32-L35) does not emit an event

contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L32-L35


 - [ ] ID-271
Setter function [SDLPoolSecondary.extendLockDuration(uint256,uint64)](contracts/core/sdlPool/SDLPoolSecondary.sol#L163-L166) does not emit an event

contracts/core/sdlPool/SDLPoolSecondary.sol#L163-L166


 - [ ] ID-272
Setter function [GovernanceController.renounceRole(uint256)](contracts/governance/GovernanceController.sol#L146-L148) does not emit an event

contracts/governance/GovernanceController.sol#L146-L148


 - [ ] ID-273
Setter function [FraxLSDIndexAdapter.initialize(address,address)](contracts/liquidSDIndex/adapters/FraxLSDIndexAdapter.sol#L17-L19) does not emit an event

contracts/liquidSDIndex/adapters/FraxLSDIndexAdapter.sol#L17-L19


 - [ ] ID-274
Setter function [WrappedSDToken.onTokenTransfer(address,uint256,bytes)](contracts/core/tokens/WrappedSDToken.sol#L27-L34) does not emit an event

contracts/core/tokens/WrappedSDToken.sol#L27-L34


 - [ ] ID-275
Setter function [CommunityVCS.initialize(address,address,address,address,VaultControllerStrategy.Fee[],uint256,uint128,uint128)](contracts/linkStaking/CommunityVCS.sol#L36-L57) does not emit an event

contracts/linkStaking/CommunityVCS.sol#L36-L57


 - [ ] ID-276
Setter function [OperatorWhitelist.removeWhitelistEntries(address[])](contracts/ethStaking/OperatorWhitelist.sol#L64-L70) does not emit an event

contracts/ethStaking/OperatorWhitelist.sol#L64-L70


 - [ ] ID-277
Setter function [OperatorVault.onlyRewardsReceiver()](contracts/linkStaking/OperatorVault.sol#L85-L88) does not emit an event

contracts/linkStaking/OperatorVault.sol#L85-L88


 - [ ] ID-278
Setter function [RewardsPoolWSD.updateReward(address)](contracts/core/RewardsPoolWSD.sol#L66-L72) does not emit an event

contracts/core/RewardsPoolWSD.sol#L66-L72


 - [ ] ID-279
Setter function [EthStakingStrategy.withdraw(uint256)](contracts/ethStaking/EthStakingStrategy.sol#L224-L226) does not emit an event

contracts/ethStaking/EthStakingStrategy.sol#L224-L226


 - [ ] ID-280
Setter function [WLOperatorController.initialize(address,address,address,uint256)](contracts/ethStaking/WLOperatorController.sol#L33-L42) does not emit an event

contracts/ethStaking/WLOperatorController.sol#L33-L42


 - [ ] ID-281
Setter function [CCIPOnRampMock.forwardFromRouter(Client.EVM2AnyMessage,uint256,address)](contracts/core/test/chainlink/CCIPOnRampMock.sol#L49-L57) does not emit an event

contracts/core/test/chainlink/CCIPOnRampMock.sol#L49-L57


 - [ ] ID-282
Setter function [StakingMock.onTokenTransfer(address,uint256,bytes)](contracts/linkStaking/test/StakingMock.sol#L39-L51) does not emit an event

contracts/linkStaking/test/StakingMock.sol#L39-L51


 - [ ] ID-283
Setter function [OperatorVCSUpgrade.initialize(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[],address[])](contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L21-L43) does not emit an event

contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L21-L43


 - [ ] ID-284
Setter function [DelegatorPool.initialize(address,string,string,address[])](contracts/core/test/deprecated/DelegatorPool.sol#L56-L74) does not emit an event

contracts/core/test/deprecated/DelegatorPool.sol#L56-L74


 - [ ] ID-285
Setter function [StakingPool.initialize(address,string,string,StakingPool.Fee[])](contracts/core/StakingPool.sol#L41-L52) does not emit an event

contracts/core/StakingPool.sol#L41-L52


 - [ ] ID-286
Setter function [StakingPool.addStrategy(address)](contracts/core/StakingPool.sol#L226-L230) does not emit an event

contracts/core/StakingPool.sol#L226-L230


 - [ ] ID-287
Setter function [WrappedTokenBridge.onTokenTransfer(address,uint256,bytes)](contracts/core/ccip/WrappedTokenBridge.sol#L83-L96) does not emit an event

contracts/core/ccip/WrappedTokenBridge.sol#L83-L96


 - [ ] ID-288
Setter function [StakingMockV1.migrate(bytes)](contracts/linkStaking/test/StakingMockV1.sol#L75-L77) does not emit an event

contracts/linkStaking/test/StakingMockV1.sol#L75-L77


 - [ ] ID-289
Setter function [DelegatorPool.staked(address)](contracts/core/test/deprecated/DelegatorPool.sol#L124-L126) does not emit an event

contracts/core/test/deprecated/DelegatorPool.sol#L124-L126


 - [ ] ID-290
Setter function [DelegatorPool.burnLockedBalance(address,uint256)](contracts/core/test/deprecated/DelegatorPool.sol#L210-L224) does not emit an event

contracts/core/test/deprecated/DelegatorPool.sol#L210-L224


 - [ ] ID-291
Setter function [StakingPool.strategyDeposit(uint256,uint256)](contracts/core/StakingPool.sol#L123-L126) does not emit an event

contracts/core/StakingPool.sol#L123-L126


 - [ ] ID-292
Setter function [WrappedNative.deposit()](contracts/core/test/chainlink/WrappedNative.sol#L9-L11) does not emit an event

contracts/core/test/chainlink/WrappedNative.sol#L9-L11


 - [ ] ID-293
Setter function [StakingPool.updateStrategyRewards(uint256[],bytes)](contracts/core/StakingPool.sol#L345-L348) does not emit an event

contracts/core/StakingPool.sol#L345-L348


 - [ ] ID-294
Setter function [LSDIndexAdapterMock.initialize(address,address,uint256)](contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#L13-L20) does not emit an event

contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#L13-L20


 - [ ] ID-295
Setter function [StakingRewardsMock.setReward(address,uint256)](contracts/linkStaking/test/StakingRewardsMock.sol#L34-L36) does not emit an event

contracts/linkStaking/test/StakingRewardsMock.sol#L34-L36


 - [ ] ID-296
Setter function [PoolOwnersV1.onTokenTransfer(address,uint256,bytes)](contracts/core/test/deprecated/PoolOwnersV1.sol#L94-L101) does not emit an event

contracts/core/test/deprecated/PoolOwnersV1.sol#L94-L101


 - [ ] ID-297
Setter function [LinkPoolNFT.mint(address)](contracts/core/tokens/LinkPoolNFT.sol#L26-L31) does not emit an event

contracts/core/tokens/LinkPoolNFT.sol#L26-L31


 - [ ] ID-298
Setter function [PoolOwnersV1._mintAllowance(address)](contracts/core/test/deprecated/PoolOwnersV1.sol#L212-L222) does not emit an event

contracts/core/test/deprecated/PoolOwnersV1.sol#L212-L222


 - [ ] ID-299
Setter function [KeyValidationOracle.onTokenTransfer(address,uint256,bytes)](contracts/ethStaking/KeyValidationOracle.sol#L48-L59) does not emit an event

contracts/ethStaking/KeyValidationOracle.sol#L48-L59


 - [ ] ID-300
Setter function [SDLPoolCCIPControllerMock.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken)](contracts/core/test/SDLPoolCCIPControllerMock.sol#L37-L45) does not emit an event

contracts/core/test/SDLPoolCCIPControllerMock.sol#L37-L45


 - [ ] ID-301
Setter function [StakingPool.updateFee(uint256,address,uint256)](contracts/core/StakingPool.sol#L292-L308) does not emit an event

contracts/core/StakingPool.sol#L292-L308


 - [ ] ID-302
Setter function [DelegatorPool.setPoolRouter(address)](contracts/core/test/deprecated/DelegatorPool.sol#L230-L233) does not emit an event

contracts/core/test/deprecated/DelegatorPool.sol#L230-L233


 - [ ] ID-303
Setter function [WrappedTokenBridge.transferTokens(uint64,address,uint256,bool,uint256)](contracts/core/ccip/WrappedTokenBridge.sol#L106-L117) does not emit an event

contracts/core/ccip/WrappedTokenBridge.sol#L106-L117


 - [ ] ID-304
Setter function [StakingAllowance.mintToContract(address,address,uint256,bytes)](contracts/core/tokens/StakingAllowance.sol#L30-L38) does not emit an event

contracts/core/tokens/StakingAllowance.sol#L30-L38


 - [ ] ID-305
Setter function [CommunityVaultV2Mock.initializeV2(uint256)](contracts/linkStaking/test/CommunityVaultV2Mock.sol#L13-L15) does not emit an event

contracts/linkStaking/test/CommunityVaultV2Mock.sol#L13-L15


 - [ ] ID-306
Setter function [SDLPoolCCIPControllerPrimary.distributeRewards()](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L56-L93) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L56-L93


 - [ ] ID-307
Setter function [ERC721ReceiverMock.onERC721Received(address,address,uint256,bytes)](contracts/core/test/ERC721ReceiverMock.sol#L24-L32) does not emit an event

contracts/core/test/ERC721ReceiverMock.sol#L24-L32


 - [ ] ID-308
Setter function [PriorityPoolMock.updateDistribution(bytes32,bytes32,uint256,uint256)](contracts/core/test/PriorityPoolMock.sol#L22-L34) does not emit an event

contracts/core/test/PriorityPoolMock.sol#L22-L34


 - [ ] ID-309
Setter function [StakingMockV1.setMigration(address)](contracts/linkStaking/test/StakingMockV1.sol#L71-L73) does not emit an event

contracts/linkStaking/test/StakingMockV1.sol#L71-L73


 - [ ] ID-310
Setter function [SDLPoolCCIPControllerPrimary.onlyRewardsInitiator()](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L48-L51) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L48-L51


 - [ ] ID-311
Setter function [OperatorVCS.setOperator(uint256,address)](contracts/linkStaking/OperatorVCS.sol#L251-L253) does not emit an event

contracts/linkStaking/OperatorVCS.sol#L251-L253


 - [ ] ID-312
Setter function [WrappedSDToken.wrap(uint256)](contracts/core/tokens/WrappedSDToken.sol#L40-L43) does not emit an event

contracts/core/tokens/WrappedSDToken.sol#L40-L43


 - [ ] ID-313
Setter function [DistributionOracle._pauseForUpdate()](contracts/core/priorityPool/DistributionOracle.sol#L263-L268) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L263-L268


 - [ ] ID-314
Setter function [StakingMock.setMaxPoolSize(uint256)](contracts/linkStaking/test/StakingMock.sol#L94-L96) does not emit an event

contracts/linkStaking/test/StakingMock.sol#L94-L96


 - [ ] ID-315
Setter function [WrappedSDToken.unwrap(uint256)](contracts/core/tokens/WrappedSDToken.sol#L49-L54) does not emit an event

contracts/core/tokens/WrappedSDToken.sol#L49-L54


 - [ ] ID-316
Setter function [SDLPoolCCIPControllerSecondary.handleOutgoingRESDL(uint64,address,uint256)](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L80-L86) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L80-L86


 - [ ] ID-317
Setter function [CurveMock.exchange(int128,int128,uint256,uint256,address)](contracts/core/test/CurveMock.sol#L26-L56) does not emit an event

contracts/core/test/CurveMock.sol#L26-L56


 - [ ] ID-318
Setter function [NWLOperatorController.addOperator(string)](contracts/ethStaking/NWLOperatorController.sol#L72-L74) does not emit an event

contracts/ethStaking/NWLOperatorController.sol#L72-L74


 - [ ] ID-319
Setter function [PriorityPool._authorizeUpgrade(address)](contracts/core/priorityPool/PriorityPool.sol#L572) does not emit an event

contracts/core/priorityPool/PriorityPool.sol#L572


 - [ ] ID-320
Setter function [StakingAllowance.mint(address,uint256)](contracts/core/tokens/StakingAllowance.sol#L20-L22) does not emit an event

contracts/core/tokens/StakingAllowance.sol#L20-L22


 - [ ] ID-321
Setter function [EthStakingStrategy.deposit(uint256)](contracts/ethStaking/EthStakingStrategy.sol#L212-L217) does not emit an event

contracts/ethStaking/EthStakingStrategy.sol#L212-L217


 - [ ] ID-322
Setter function [PoolOwnersV1.stake(uint256)](contracts/core/test/deprecated/PoolOwnersV1.sol#L107-L110) does not emit an event

contracts/core/test/deprecated/PoolOwnersV1.sol#L107-L110


 - [ ] ID-323
Setter function [DistributionOracle.pauseForUpdate()](contracts/core/priorityPool/DistributionOracle.sol#L142-L144) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L142-L144


 - [ ] ID-324
Setter function [PriorityPool.deposit(uint256,bool)](contracts/core/priorityPool/PriorityPool.sol#L209-L213) does not emit an event

contracts/core/priorityPool/PriorityPool.sol#L209-L213


 - [ ] ID-325
Setter function [ERC677ReceiverMock.onTokenTransfer(address,uint256,bytes)](contracts/core/test/ERC677ReceiverMock.sol#L11-L17) does not emit an event

contracts/core/test/ERC677ReceiverMock.sol#L11-L17


 - [ ] ID-326
Setter function [SDLPoolCCIPControllerPrimary.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L125-L134) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L125-L134


 - [ ] ID-327
Setter function [SDLPoolCCIPControllerSecondary.performUpkeep(bytes)](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L66-L71) does not emit an event

contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L66-L71


 - [ ] ID-328
Setter function [OperatorVault.updateDeposits(uint256,address)](contracts/linkStaking/OperatorVault.sol#L175-L200) does not emit an event

contracts/linkStaking/OperatorVault.sol#L175-L200


 - [ ] ID-329
Setter function [PoolAllowanceV1.mintAllowance(address,uint256)](contracts/core/test/deprecated/PoolAllowanceV1.sol#L31-L33) does not emit an event

contracts/core/test/deprecated/PoolAllowanceV1.sol#L31-L33


 - [ ] ID-330
Setter function [SDLPoolSecondary.executeQueuedOperations(uint256[])](contracts/core/sdlPool/SDLPoolSecondary.sol#L247-L250) does not emit an event

contracts/core/sdlPool/SDLPoolSecondary.sol#L247-L250


 - [ ] ID-331
Setter function [PriorityPoolMock.pauseForUpdate()](contracts/core/test/PriorityPoolMock.sol#L36-L38) does not emit an event

contracts/core/test/PriorityPoolMock.sol#L36-L38


 - [ ] ID-332
Setter function [SDLPoolPrimary.migrate(address,uint256,uint64)](contracts/core/sdlPool/SDLPoolPrimary.sol#L264-L272) does not emit an event

contracts/core/sdlPool/SDLPoolPrimary.sol#L264-L272


 - [ ] ID-333
Setter function [OwnersRewardsPoolV1.withdraw(uint256)](contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L33-L42) does not emit an event

contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L33-L42


 - [ ] ID-334
Setter function [MerkleDistributor.addDistributions(address[],bytes32[],uint256[],uint256[])](contracts/airdrop/MerkleDistributor.sol#L54-L68) does not emit an event

contracts/airdrop/MerkleDistributor.sol#L54-L68


 - [ ] ID-335
Setter function [RewardsPoolControllerMock.withdraw(uint256)](contracts/core/test/RewardsPoolControllerMock.sol#L42-L46) does not emit an event

contracts/core/test/RewardsPoolControllerMock.sol#L42-L46


 - [ ] ID-336
Setter function [CCIPOffRampMock.setTokenPool(address,address)](contracts/core/test/chainlink/CCIPOffRampMock.sol#L54-L56) does not emit an event

contracts/core/test/chainlink/CCIPOffRampMock.sol#L54-L56


 - [ ] ID-337
Setter function [PriorityPool._withdraw(uint256)](contracts/core/priorityPool/PriorityPool.sol#L519-L534) does not emit an event

contracts/core/priorityPool/PriorityPool.sol#L519-L534


 - [ ] ID-338
Setter function [OperatorVCSMock.setWithdrawalPercentage(uint256)](contracts/linkStaking/test/OperatorVCSMock.sol#L52-L54) does not emit an event

contracts/linkStaking/test/OperatorVCSMock.sol#L52-L54


 - [ ] ID-339
Setter function [StakingMockV1.setPaused(bool)](contracts/linkStaking/test/StakingMockV1.sol#L99-L101) does not emit an event

contracts/linkStaking/test/StakingMockV1.sol#L99-L101


 - [ ] ID-340
Setter function [SDLPoolCCIPControllerMock.distributeRewards()](contracts/core/test/SDLPoolCCIPControllerMock.sol#L47-L49) does not emit an event

contracts/core/test/SDLPoolCCIPControllerMock.sol#L47-L49


 - [ ] ID-341
Setter function [OperatorWhitelist.useWhitelist(address)](contracts/ethStaking/OperatorWhitelist.sol#L41-L46) does not emit an event

contracts/ethStaking/OperatorWhitelist.sol#L41-L46


 - [ ] ID-342
Setter function [DistributionOracle.cancelRequest(bytes32,uint256)](contracts/core/priorityPool/DistributionOracle.sol#L208-L211) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L208-L211


 - [ ] ID-343
Setter function [StakingRewardsMock.claimReward()](contracts/linkStaking/test/StakingRewardsMock.sol#L27-L32) does not emit an event

contracts/linkStaking/test/StakingRewardsMock.sol#L27-L32


 - [ ] ID-344
Setter function [WLOperatorController.assignNextValidators(uint256[],uint256[],uint256)](contracts/ethStaking/WLOperatorController.sol#L125-L272) does not emit an event

contracts/ethStaking/WLOperatorController.sol#L125-L272


 - [ ] ID-345
Setter function [StakingPool.onlyPriorityPool()](contracts/core/StakingPool.sol#L54-L57) does not emit an event

contracts/core/StakingPool.sol#L54-L57


 - [ ] ID-346
Setter function [LSDIndexAdapterMock.setExchangeRate(uint256)](contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#L30-L32) does not emit an event

contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#L30-L32


 - [ ] ID-347
Setter function [StakingMockV1.onTokenTransfer(address,uint256,bytes)](contracts/linkStaking/test/StakingMockV1.sol#L30-L37) does not emit an event

contracts/linkStaking/test/StakingMockV1.sol#L30-L37


 - [ ] ID-348
Setter function [DistributionOracle.executeManualVerification()](contracts/core/priorityPool/DistributionOracle.sol#L182-L192) does not emit an event

contracts/core/priorityPool/DistributionOracle.sol#L182-L192


 - [ ] ID-349
Setter function [WrappedSDTokenMock.setMultiplier(uint256)](contracts/core/test/WrappedSDTokenMock.sol#L63-L65) does not emit an event

contracts/core/test/WrappedSDTokenMock.sol#L63-L65


 - [ ] ID-350
Setter function [OperatorVault.deposit(uint256)](contracts/linkStaking/OperatorVault.sol#L95-L99) does not emit an event

contracts/linkStaking/OperatorVault.sol#L95-L99


 - [ ] ID-351
Setter function [DelegatorPool.setCommunityPool(address,bool)](contracts/core/test/deprecated/DelegatorPool.sol#L240-L243) does not emit an event

contracts/core/test/deprecated/DelegatorPool.sol#L240-L243


## initialize-permission
Impact: Low
Confidence: Medium
 - [ ] ID-352
Condition variable is not initialized found in [Initializable._isInitializing()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L163-L165)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L163-L165


 - [ ] ID-353
Condition variable is not initialized found in [Initializable._disableInitializers()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L145-L151)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L145-L151


 - [ ] ID-354
Condition variable is not initialized found in [KeyValidationOracle._initiateKeyPairValidation(address,uint256,bool)](contracts/ethStaking/KeyValidationOracle.sol#L112-L129)

contracts/ethStaking/KeyValidationOracle.sol#L112-L129


 - [ ] ID-355
Condition variable is not initialized found in [SDLPoolCCIPControllerSecondary._initiateUpdate(uint64,address,bytes)](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L119-L140)

contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L119-L140


## input-validation
Impact: Low
Confidence: Medium
 - [ ] ID-356
value assignment lack of validation	[SDLPoolMock.setEffectiveBalance(address,uint256)](contracts/core/test/SDLPoolMock.sol#L15-L17)[effectiveBalances[_account] = _amount](contracts/core/test/SDLPoolMock.sol#L16)

contracts/core/test/SDLPoolMock.sol#L15-L17


## missing-zero-check
Impact: Low
Confidence: Medium
 - [ ] ID-357
variable lacks a zero-check on 		- [Operator.transferOwnableContracts(address[],address)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L235-L240)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L235-L240


 - [ ] ID-358
variable lacks a zero-check on 		- [Operator.ownerTransferAndCall(address,uint256,bytes)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L332-L338)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L332-L338


 - [ ] ID-359
variable lacks a zero-check on 		- [VestingWallet.vestedAmount(address,uint64)](node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L118-L120)

node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L118-L120


 - [ ] ID-360
variable lacks a zero-check on 		- [Router.constructor(address,address)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L65-L70)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L65-L70


 - [ ] ID-361
variable lacks a zero-check on 		- [Router.setWrappedNative(address)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L228-L230)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L228-L230


 - [ ] ID-362
variable lacks a zero-check on 		- [VestingWallet.release(address)](node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L101-L106)

node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L101-L106


## variable-scope
Impact: Low
Confidence: High
 - [ ] ID-363
Variable '[BytesLib.concatStorage(bytes,bytes).sc_concatStorage_asm_0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L147)' in [BytesLib.concatStorage(bytes,bytes)](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L91-L226) potentially used before declaration: [sc_concatStorage_asm_0 = keccak256(uint256,uint256)(0x0,0x20) + slength_concatStorage_asm_0 / 32](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L195)

node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L147


 - [ ] ID-364
Variable '[BytesLib.concatStorage(bytes,bytes).mc_concatStorage_asm_0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L162)' in [BytesLib.concatStorage(bytes,bytes)](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L91-L226) potentially used before declaration: [mc_concatStorage_asm_0 = _postBytes + submod_concatStorage_asm_0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L205)

node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L162


 - [ ] ID-365
Variable '[BytesLib.concatStorage(bytes,bytes).mask_concatStorage_asm_0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L164)' in [BytesLib.concatStorage(bytes,bytes)](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L91-L226) potentially used before declaration: [mask_concatStorage_asm_0 = 0x100 ** submod_concatStorage_asm_0 - 1](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L207)

node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L164


 - [ ] ID-366
Variable '[BytesLib.concatStorage(bytes,bytes).end_concatStorage_asm_0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L163)' in [BytesLib.concatStorage(bytes,bytes)](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L91-L226) potentially used before declaration: [end_concatStorage_asm_0 = _postBytes + mlength_concatStorage_asm_0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L206)

node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L163


 - [ ] ID-367
Variable '[BytesLib.concatStorage(bytes,bytes).submod_concatStorage_asm_0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L161)' in [BytesLib.concatStorage(bytes,bytes)](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L91-L226) potentially used before declaration: [submod_concatStorage_asm_0 = 32 - slengthmod_concatStorage_asm_0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L204)

node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L161


## reentrancy-same-effect
Impact: Low
Confidence: Medium
 - [ ] ID-368
Reentrancy in [LinkPoolNFT.mint(address)](contracts/core/tokens/LinkPoolNFT.sol#L26-L31):
	External calls:
	- [_safeMint(_to,tokenId)](contracts/core/tokens/LinkPoolNFT.sol#L29)
		- [IERC721Receiver(to).onERC721Received(_msgSender(),from,tokenId,data)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L401-L412)

contracts/core/tokens/LinkPoolNFT.sol#L26-L31


## unnecessary-boolean-compare
Impact: Informational
Confidence: High
 - [ ] ID-369
[AuthorizedReceiver.setAuthorizedSenders(address[])](node_modules/@chainlink/contracts/src/v0.7/AuthorizedReceiver.sol#L16-L31) compares to a boolean constant:
	-[require(bool,string)(s_authorizedSenders[senders[i_scope_0]] == false,Must not have duplicate senders)](node_modules/@chainlink/contracts/src/v0.7/AuthorizedReceiver.sol#L25)

node_modules/@chainlink/contracts/src/v0.7/AuthorizedReceiver.sol#L16-L31


 - [ ] ID-370
[PriorityPool.withdraw(uint256,uint256,uint256,bytes32[],bool)](contracts/core/priorityPool/PriorityPool.sol#L225-L266) compares to a boolean constant:
	-[_shouldUnqueue == true](contracts/core/priorityPool/PriorityPool.sol#L237)

contracts/core/priorityPool/PriorityPool.sol#L225-L266


 - [ ] ID-371
[RESDLTokenBridge.transferRESDL(uint64,address,uint256,bool,uint256)](contracts/core/ccip/RESDLTokenBridge.sol#L84-L135) compares to a boolean constant:
	-[_payNative == false && msg.value != 0](contracts/core/ccip/RESDLTokenBridge.sol#L93)

contracts/core/ccip/RESDLTokenBridge.sol#L84-L135


 - [ ] ID-372
[DistributionOracle.withdrawLink(uint256)](contracts/core/priorityPool/DistributionOracle.sol#L217-L220) compares to a boolean constant:
	-[link.transfer(msg.sender,_amount) != true](contracts/core/priorityPool/DistributionOracle.sol#L219)

contracts/core/priorityPool/DistributionOracle.sol#L217-L220


 - [ ] ID-373
[WrappedTokenBridge.transferTokens(uint64,address,uint256,bool,uint256)](contracts/core/ccip/WrappedTokenBridge.sol#L106-L117) compares to a boolean constant:
	-[_payNative == false && msg.value != 0](contracts/core/ccip/WrappedTokenBridge.sol#L113)

contracts/core/ccip/WrappedTokenBridge.sol#L106-L117


## centralized-risk-informational
Impact: Informational
Confidence: High
 - [ ] ID-374
	- [Operator.transferOwnableContracts(address[],address)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L235-L240)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L235-L240


 - [ ] ID-375
	- [ConfirmedOwnerWithProposal.acceptOwnership()](node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L37-L45)

node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L37-L45


 - [ ] ID-376
	- [StakedotlinkCouncil.burn(uint256)](contracts/governance/StakedotlinkCouncil.sol#L131-L152)

contracts/governance/StakedotlinkCouncil.sol#L131-L152


 - [ ] ID-377
	- [StakedotlinkCouncil.setTokenURI(uint256,string)](contracts/governance/StakedotlinkCouncil.sol#L178-L182)

contracts/governance/StakedotlinkCouncil.sol#L178-L182


 - [ ] ID-378
	- [Router.setWrappedNative(address)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L228-L230)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L228-L230


 - [ ] ID-379
	- [StakedotlinkCouncil.mintWithTokenURI(address,uint256,string)](contracts/governance/StakedotlinkCouncil.sol#L100-L111)

contracts/governance/StakedotlinkCouncil.sol#L100-L111


 - [ ] ID-380
	- [Router.applyRampUpdates(Router.OnRamp[],Router.OffRamp[],Router.OffRamp[])](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L264-L296)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L264-L296


 - [ ] ID-381
	- [LiquidSDIndexPool.setCompositionTolerance(uint256)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L411-L415)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L411-L415


 - [ ] ID-382
	- [SDLPoolCCIPControllerPrimary.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L125-L134)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L125-L134


 - [ ] ID-383
	- [RewardsPoolController.addToken(address,address)](contracts/core/base/RewardsPoolController.sol#L154-L165)

contracts/core/base/RewardsPoolController.sol#L154-L165


 - [ ] ID-384
	- [WLOperatorController.reportStoppedValidators(uint256[],uint256[])](contracts/ethStaking/WLOperatorController.sol#L369-L404)

contracts/ethStaking/WLOperatorController.sol#L369-L404


 - [ ] ID-385
	- [SDLPoolCCIPControllerPrimary.setRewardsInitiator(address)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L234-L236)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L234-L236


 - [ ] ID-386
	- [SDLPool.setCCIPController(address)](contracts/core/sdlPool/base/SDLPool.sol#L381-L383)

contracts/core/sdlPool/base/SDLPool.sol#L381-L383


 - [ ] ID-387
	- [SDLPoolPrimary.handleOutgoingRESDL(address,uint256,address)](contracts/core/sdlPool/SDLPoolPrimary.sol#L172-L199)

contracts/core/sdlPool/SDLPoolPrimary.sol#L172-L199


 - [ ] ID-388
	- [SDLPoolSecondary.initiateUnlock(uint256)](contracts/core/sdlPool/SDLPoolSecondary.sol#L180-L200)

contracts/core/sdlPool/SDLPoolSecondary.sol#L180-L200


 - [ ] ID-389
	- [RewardsPoolControllerV1.addToken(address,address)](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L178-L189)

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L178-L189


 - [ ] ID-390
	- [StakingPool.addFee(address,uint256)](contracts/core/StakingPool.sol#L281-L284)

contracts/core/StakingPool.sol#L281-L284


 - [ ] ID-391
	- [SDLPoolCCIPControllerPrimary.setWrappedRewardToken(address,address)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L202-L205)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L202-L205


 - [ ] ID-392
	- [KeyValidationOracle.setOracleConfig(address,bytes32,uint256)](contracts/ethStaking/KeyValidationOracle.sol#L87-L96)

contracts/ethStaking/KeyValidationOracle.sol#L87-L96


 - [ ] ID-393
	- [DistributionOracle.toggleManualVerification()](contracts/core/priorityPool/DistributionOracle.sol#L243-L246)

contracts/core/priorityPool/DistributionOracle.sol#L243-L246


 - [ ] ID-394
	- [MerkleDistributor.updateDistribution(address,bytes32,uint256,uint256)](contracts/airdrop/MerkleDistributor.sol#L128-L143)

contracts/airdrop/MerkleDistributor.sol#L128-L143


 - [ ] ID-395
	- [VaultControllerStrategyUpgrade.__VaultControllerStrategy_init(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[])](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L86-L109)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L86-L109


 - [ ] ID-396
	- [LiquidSDIndexPool.addLSDToken(address,address,uint256[])](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L328-L347)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L328-L347


 - [ ] ID-397
	- [PoolOwnersV1.removeRewardToken(uint16)](contracts/core/test/deprecated/PoolOwnersV1.sol#L174-L185)

contracts/core/test/deprecated/PoolOwnersV1.sol#L174-L185


 - [ ] ID-398
	- [VaultControllerStrategy.updateDeposits(bytes)](contracts/linkStaking/base/VaultControllerStrategy.sol#L157-L191)

contracts/linkStaking/base/VaultControllerStrategy.sol#L157-L191


 - [ ] ID-399
	- [LinkPoolNFT.mint(address)](contracts/core/tokens/LinkPoolNFT.sol#L26-L31)

contracts/core/tokens/LinkPoolNFT.sol#L26-L31


 - [ ] ID-400
	- [VaultControllerStrategyUpgrade.setVaultImplementation(address)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L360-L364)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L360-L364


 - [ ] ID-401
	- [OperatorController.setBeaconOracle(address)](contracts/ethStaking/base/OperatorController.sol#L280-L283)

contracts/ethStaking/base/OperatorController.sol#L280-L283


 - [ ] ID-402
	- [EthStakingStrategy.setBeaconOracle(address)](contracts/ethStaking/EthStakingStrategy.sol#L304-L307)

contracts/ethStaking/EthStakingStrategy.sol#L304-L307


 - [ ] ID-403
	- [DistributionOracle.fulfillRequest(bytes32,bytes32,bytes32,uint256,uint256)](contracts/core/priorityPool/DistributionOracle.sol#L163-L177)

contracts/core/priorityPool/DistributionOracle.sol#L163-L177


 - [ ] ID-404
	- [DelegatorPool.burnLockedBalance(address,uint256)](contracts/core/test/deprecated/DelegatorPool.sol#L210-L224)

contracts/core/test/deprecated/DelegatorPool.sol#L210-L224


 - [ ] ID-405
	- [OperatorVCS.updateDeposits(bytes)](contracts/linkStaking/OperatorVCS.sol#L155-L216)

contracts/linkStaking/OperatorVCS.sol#L155-L216


 - [ ] ID-406
	- [LinearBoostController.setMaxBoost(uint64)](contracts/core/sdlPool/LinearBoostController.sol#L56-L59)

contracts/core/sdlPool/LinearBoostController.sol#L56-L59


 - [ ] ID-407
	- [RewardsInitiator.whitelistCaller(address,bool)](contracts/core/RewardsInitiator.sol#L101-L104)

contracts/core/RewardsInitiator.sol#L101-L104


 - [ ] ID-408
	- [OperatorVault.setRewardsReceiver(address)](contracts/linkStaking/OperatorVault.sol#L227-L233)

contracts/linkStaking/OperatorVault.sol#L227-L233


 - [ ] ID-409
	- [VaultControllerStrategyUpgrade.setMinDepositThreshold(uint256)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L349-L354)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L349-L354


 - [ ] ID-410
	- [LiquidSDIndexPool.updateFee(uint256,address,uint256)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L457-L474)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L457-L474


 - [ ] ID-411
	- [CommunityVCS.setVaultDeploymentParams(uint128,uint128)](contracts/linkStaking/CommunityVCS.sol#L115-L119)

contracts/linkStaking/CommunityVCS.sol#L115-L119


 - [ ] ID-412
	- [EthStakingStrategy.setRewardsReceiver(address)](contracts/ethStaking/EthStakingStrategy.sol#L364-L367)

contracts/ethStaking/EthStakingStrategy.sol#L364-L367


 - [ ] ID-413
	- [EthStakingStrategy.setDepositController(address)](contracts/ethStaking/EthStakingStrategy.sol#L355-L358)

contracts/ethStaking/EthStakingStrategy.sol#L355-L358


 - [ ] ID-414
	- [StakingPool.addStrategy(address)](contracts/core/StakingPool.sol#L226-L230)

contracts/core/StakingPool.sol#L226-L230


 - [ ] ID-415
	- [VaultV1.__Vault_init(address,address,address)](contracts/linkStaking/test/deprecated/VaultV1.sol#L52-L62)

contracts/linkStaking/test/deprecated/VaultV1.sol#L52-L62


 - [ ] ID-416
	- [SDLPoolCCIPControllerPrimary.removeWhitelistedChain(uint64)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L170-L184)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L170-L184


 - [ ] ID-417
	- [WLOperatorController.assignNextValidators(uint256[],uint256[],uint256)](contracts/ethStaking/WLOperatorController.sol#L125-L272)

contracts/ethStaking/WLOperatorController.sol#L125-L272


 - [ ] ID-418
	- [VaultControllerStrategyUpgrade.deposit(uint256)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L123-L127)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L123-L127


 - [ ] ID-419
	- [StakingPool.updateFee(uint256,address,uint256)](contracts/core/StakingPool.sol#L292-L308)

contracts/core/StakingPool.sol#L292-L308


 - [ ] ID-420
	- [DelegatorPool.setCommunityPool(address,bool)](contracts/core/test/deprecated/DelegatorPool.sol#L240-L243)

contracts/core/test/deprecated/DelegatorPool.sol#L240-L243


 - [ ] ID-421
	- [StakingMock.onTokenTransfer(address,uint256,bytes)](contracts/linkStaking/test/StakingMock.sol#L39-L51)

contracts/linkStaking/test/StakingMock.sol#L39-L51


 - [ ] ID-422
	- [SDLPoolCCIPController.setRESDLTokenBridge(address)](contracts/core/ccip/base/SDLPoolCCIPController.sol#L138-L140)

contracts/core/ccip/base/SDLPoolCCIPController.sol#L138-L140


 - [ ] ID-423
	- [OwnersRewardsPoolV1.withdraw(address)](contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L48-L60)

contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L48-L60


 - [ ] ID-424
	- [SDLPoolSecondary.handleIncomingRESDL(address,uint256,SDLPool.Lock)](contracts/core/sdlPool/SDLPoolSecondary.sol#L289-L307)

contracts/core/sdlPool/SDLPoolSecondary.sol#L289-L307


 - [ ] ID-425
	- [MerkleDistributor.pauseForWithdrawal(address)](contracts/airdrop/MerkleDistributor.sol#L232-L239)

contracts/airdrop/MerkleDistributor.sol#L232-L239


 - [ ] ID-426
	- [WLOperatorController.reportKeyPairValidation(uint256,bool)](contracts/ethStaking/WLOperatorController.sol#L100-L115)

contracts/ethStaking/WLOperatorController.sol#L100-L115


 - [ ] ID-427
	- [VaultControllerStrategy.setVaultImplementation(address)](contracts/linkStaking/base/VaultControllerStrategy.sol#L320-L324)

contracts/linkStaking/base/VaultControllerStrategy.sol#L320-L324


 - [ ] ID-428
	- [DelegatorPool.onTokenTransfer(address,uint256,bytes)](contracts/core/test/deprecated/DelegatorPool.sol#L82-L104)

contracts/core/test/deprecated/DelegatorPool.sol#L82-L104


 - [ ] ID-429
	- [PriorityPool.setDistributionOracle(address)](contracts/core/priorityPool/PriorityPool.sol#L468-L470)

contracts/core/priorityPool/PriorityPool.sol#L468-L470


 - [ ] ID-430
	- [PriorityPool.updateDistribution(bytes32,bytes32,uint256,uint256)](contracts/core/priorityPool/PriorityPool.sol#L409-L424)

contracts/core/priorityPool/PriorityPool.sol#L409-L424


 - [ ] ID-431
	- [DistributionOracle.setUpdateParams(uint64,uint128,uint64)](contracts/core/priorityPool/DistributionOracle.sol#L229-L238)

contracts/core/priorityPool/DistributionOracle.sol#L229-L238


 - [ ] ID-432
	- [LiquidSDIndexPool.setPaused(bool)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L480-L483)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L480-L483


 - [ ] ID-433
	- [VaultControllerStrategyUpgrade.addFee(address,uint256)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L315-L318)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L315-L318


 - [ ] ID-434
	- [VaultControllerStrategy.__VaultControllerStrategy_init(address,address,address,address,VaultControllerStrategy.Fee[],uint256)](contracts/linkStaking/base/VaultControllerStrategy.sol#L56-L78)

contracts/linkStaking/base/VaultControllerStrategy.sol#L56-L78


 - [ ] ID-435
	- [NWLOperatorController.withdrawStake(uint256,uint256)](contracts/ethStaking/NWLOperatorController.sol#L349-L357)

contracts/ethStaking/NWLOperatorController.sol#L349-L357


 - [ ] ID-436
	- [OperatorVCS.togglePreRelease()](contracts/linkStaking/OperatorVCS.sol#L287-L289)

contracts/linkStaking/OperatorVCS.sol#L287-L289


 - [ ] ID-437
	- [MerkleDistributor.setExpiryTimestamp(address,uint256)](contracts/airdrop/MerkleDistributor.sol#L246-L251)

contracts/airdrop/MerkleDistributor.sol#L246-L251


 - [ ] ID-438
	- [StrategyMock.setMaxDeposits(uint256)](contracts/core/test/StrategyMock.sol#L101-L103)

contracts/core/test/StrategyMock.sol#L101-L103


 - [ ] ID-439
	- [PoolOwnersV1.addRewardToken(address,address,address)](contracts/core/test/deprecated/PoolOwnersV1.sol#L157-L168)

contracts/core/test/deprecated/PoolOwnersV1.sol#L157-L168


 - [ ] ID-440
	- [VaultV1.migrate(bytes)](contracts/linkStaking/test/deprecated/VaultV1.sol#L102-L105)

contracts/linkStaking/test/deprecated/VaultV1.sol#L102-L105


 - [ ] ID-441
	- [PriorityPool.setPoolStatusClosed()](contracts/core/priorityPool/PriorityPool.sol#L447-L451)

contracts/core/priorityPool/PriorityPool.sol#L447-L451


 - [ ] ID-442
	- [SDLPoolCCIPControllerPrimary.setRewardsExtraArgs(uint64,bytes)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L223-L227)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L223-L227


 - [ ] ID-443
	- [OperatorVault.deposit(uint256)](contracts/linkStaking/OperatorVault.sol#L95-L99)

contracts/linkStaking/OperatorVault.sol#L95-L99


 - [ ] ID-444
	- [LinkPoolNFT.setBaseURI(string)](contracts/core/tokens/LinkPoolNFT.sol#L37-L39)

contracts/core/tokens/LinkPoolNFT.sol#L37-L39


 - [ ] ID-445
	- [WLOperatorController.setBatchSize(uint256)](contracts/ethStaking/WLOperatorController.sol#L410-L412)

contracts/ethStaking/WLOperatorController.sol#L410-L412


 - [ ] ID-446
	- [OperatorVault.setOperator(address)](contracts/linkStaking/OperatorVault.sol#L211-L215)

contracts/linkStaking/OperatorVault.sol#L211-L215


 - [ ] ID-447
	- [StakingRewardsPool.__StakingRewardsPool_init(address,string,string)](contracts/core/base/StakingRewardsPool.sol#L20-L29)

contracts/core/base/StakingRewardsPool.sol#L20-L29


 - [ ] ID-448
	- [LiquidSDIndexPool.setWithdrawalFee(uint256)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L434-L438)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L434-L438


 - [ ] ID-449
	- [StakingMockV1.onTokenTransfer(address,uint256,bytes)](contracts/linkStaking/test/StakingMockV1.sol#L30-L37)

contracts/linkStaking/test/StakingMockV1.sol#L30-L37


 - [ ] ID-450
	- [GovernanceController.addRole(string,address[],address[],bytes4[][])](contracts/governance/GovernanceController.sol#L99-L120)

contracts/governance/GovernanceController.sol#L99-L120


 - [ ] ID-451
	- [VaultControllerStrategyUpgrade.updateFee(uint256,address,uint256)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L326-L342)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L326-L342


 - [ ] ID-452
	- [OperatorVCS.addVault(address,address,address)](contracts/linkStaking/OperatorVCS.sol#L225-L243)

contracts/linkStaking/OperatorVCS.sol#L225-L243


 - [ ] ID-453
	- [VaultControllerStrategyUpgrade.updateDeposits(bytes)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L227-L251)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L227-L251


 - [ ] ID-454
	- [GovernanceController.grantRole(uint256,address)](contracts/governance/GovernanceController.sol#L127-L131)

contracts/governance/GovernanceController.sol#L127-L131


 - [ ] ID-455
	- [OperatorWhitelist.addWhitelistEntries(address[])](contracts/ethStaking/OperatorWhitelist.sol#L52-L58)

contracts/ethStaking/OperatorWhitelist.sol#L52-L58


 - [ ] ID-456
	- [DistributionOracle.setChainlinkParams(bytes32,uint256)](contracts/core/priorityPool/DistributionOracle.sol#L253-L257)

contracts/core/priorityPool/DistributionOracle.sol#L253-L257


 - [ ] ID-457
	- [MerkleDistributor.addDistribution(address,bytes32,uint256,uint256)](contracts/airdrop/MerkleDistributor.sol#L77-L93)

contracts/airdrop/MerkleDistributor.sol#L77-L93


 - [ ] ID-458
	- [VaultControllerStrategy.updateFee(uint256,address,uint256)](contracts/linkStaking/base/VaultControllerStrategy.sol#L284-L300)

contracts/linkStaking/base/VaultControllerStrategy.sol#L284-L300


 - [ ] ID-459
	- [SDLPoolCCIPController.setMaxLINKFee(uint256)](contracts/core/ccip/base/SDLPoolCCIPController.sol#L130-L132)

contracts/core/ccip/base/SDLPoolCCIPController.sol#L130-L132


 - [ ] ID-460
	- [DistributionOracle.rejectManualVerificationAndRetry()](contracts/core/priorityPool/DistributionOracle.sol#L197-L201)

contracts/core/priorityPool/DistributionOracle.sol#L197-L201


 - [ ] ID-461
	- [Vesting.terminateVesting(address[])](contracts/vesting/Vesting.sol#L27-L37)

contracts/vesting/Vesting.sol#L27-L37


 - [ ] ID-462
	- [LiquidSDIndexPool.setCompositionEnforcementThreshold(uint256)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L425-L428)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L425-L428


 - [ ] ID-463
	- [LSDIndexAdapter.__LiquidSDAdapter_init(address,address)](contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L20-L26)

contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L20-L26


 - [ ] ID-464
	- [StakingPool.reorderStrategies(uint256[])](contracts/core/StakingPool.sol#L261-L274)

contracts/core/StakingPool.sol#L261-L274


 - [ ] ID-465
	- [DelegatorPool.setLockedApproval(address,uint256)](contracts/core/test/deprecated/DelegatorPool.sol#L200-L203)

contracts/core/test/deprecated/DelegatorPool.sol#L200-L203


 - [ ] ID-466
	- [DistributionOracle.cancelRequest(bytes32,uint256)](contracts/core/priorityPool/DistributionOracle.sol#L208-L211)

contracts/core/priorityPool/DistributionOracle.sol#L208-L211


 - [ ] ID-467
	- [LiquidSDIndexPool.addFee(address,uint256)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L445-L449)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L445-L449


 - [ ] ID-468
	- [StakingPool.setRewardsInitiator(address)](contracts/core/StakingPool.sol#L384-L386)

contracts/core/StakingPool.sol#L384-L386


 - [ ] ID-469
	- [DelegatorPool.setPoolRouter(address)](contracts/core/test/deprecated/DelegatorPool.sol#L230-L233)

contracts/core/test/deprecated/DelegatorPool.sol#L230-L233


 - [ ] ID-470
	- [EthStakingStrategy.setWLOperatorController(address)](contracts/ethStaking/EthStakingStrategy.sol#L286-L289)

contracts/ethStaking/EthStakingStrategy.sol#L286-L289


 - [ ] ID-471
	- [SDLPoolPrimary.handleIncomingRESDL(address,uint256,SDLPool.Lock)](contracts/core/sdlPool/SDLPoolPrimary.sol#L207-L223)

contracts/core/sdlPool/SDLPoolPrimary.sol#L207-L223


 - [ ] ID-472
	- [LiquidSDIndexPool.removeLSDToken(address,uint256[])](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L354-L385)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L354-L385


 - [ ] ID-473
	- [RewardsPoolControllerV1.removeToken(address)](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L195-L209)

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L195-L209


 - [ ] ID-474
	- [StakingPool.deposit(address,uint256)](contracts/core/StakingPool.sol#L80-L88)

contracts/core/StakingPool.sol#L80-L88


 - [ ] ID-475
	- [VaultControllerStrategy.setMaxDepositSizeBP(uint256)](contracts/linkStaking/base/VaultControllerStrategy.sol#L307-L311)

contracts/linkStaking/base/VaultControllerStrategy.sol#L307-L311


 - [ ] ID-476
	- [WLOperatorController.setOperatorWhitelist(address)](contracts/ethStaking/WLOperatorController.sol#L418-L420)

contracts/ethStaking/WLOperatorController.sol#L418-L420


 - [ ] ID-477
	- [WLOperatorController.removeKeyPairs(uint256,uint256)](contracts/ethStaking/WLOperatorController.sol#L75-L93)

contracts/ethStaking/WLOperatorController.sol#L75-L93


 - [ ] ID-478
	- [SDLPool.__SDLPoolBase_init(string,string,address,address)](contracts/core/sdlPool/base/SDLPool.sol#L93-L104)

contracts/core/sdlPool/base/SDLPool.sol#L93-L104


 - [ ] ID-479
	- [StrategyMock.setMinDeposits(uint256)](contracts/core/test/StrategyMock.sol#L105-L107)

contracts/core/test/StrategyMock.sol#L105-L107


 - [ ] ID-480
	- [OperatorWhitelist.useWhitelist(address)](contracts/ethStaking/OperatorWhitelist.sol#L41-L46)

contracts/ethStaking/OperatorWhitelist.sol#L41-L46


 - [ ] ID-481
	- [OperatorVault.updateDeposits(uint256,address)](contracts/linkStaking/OperatorVault.sol#L175-L200)

contracts/linkStaking/OperatorVault.sol#L175-L200


 - [ ] ID-482
	- [OperatorWhitelist.removeWhitelistEntries(address[])](contracts/ethStaking/OperatorWhitelist.sol#L64-L70)

contracts/ethStaking/OperatorWhitelist.sol#L64-L70


 - [ ] ID-483
	- [SDLPoolPrimary.handleIncomingUpdate(uint256,int256)](contracts/core/sdlPool/SDLPoolPrimary.sol#L231-L253)

contracts/core/sdlPool/SDLPoolPrimary.sol#L231-L253


 - [ ] ID-484
	- [StrategyMock.updateDeposits(bytes)](contracts/core/test/StrategyMock.sol#L58-L79)

contracts/core/test/StrategyMock.sol#L58-L79


 - [ ] ID-485
	- [EthStakingStrategy.setMinDeposits(uint256)](contracts/ethStaking/EthStakingStrategy.sol#L346-L349)

contracts/ethStaking/EthStakingStrategy.sol#L346-L349


 - [ ] ID-486
	- [SDLPoolPrimary.withdraw(uint256,uint256)](contracts/core/sdlPool/SDLPoolPrimary.sol#L134-L164)

contracts/core/sdlPool/SDLPoolPrimary.sol#L134-L164


 - [ ] ID-487
	- [EthStakingStrategy.setMaxDeposits(uint256)](contracts/ethStaking/EthStakingStrategy.sol#L337-L340)

contracts/ethStaking/EthStakingStrategy.sol#L337-L340


 - [ ] ID-488
	- [PriorityPool.setPoolStatus(PriorityPool.PoolStatus)](contracts/core/priorityPool/PriorityPool.sol#L437-L442)

contracts/core/priorityPool/PriorityPool.sol#L437-L442


 - [ ] ID-489
	- [LinearBoostController.setMaxLockingDuration(uint64)](contracts/core/sdlPool/LinearBoostController.sol#L45-L48)

contracts/core/sdlPool/LinearBoostController.sol#L45-L48


 - [ ] ID-490
	- [DistributionOracle.executeManualVerification()](contracts/core/priorityPool/DistributionOracle.sol#L182-L192)

contracts/core/priorityPool/DistributionOracle.sol#L182-L192


 - [ ] ID-491
	- [OperatorVaultV1.setOperator(address)](contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L61-L64)

contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L61-L64


 - [ ] ID-492
	- [SDLPoolSecondary.handleIncomingUpdate(uint256)](contracts/core/sdlPool/SDLPoolSecondary.sol#L336-L347)

contracts/core/sdlPool/SDLPoolSecondary.sol#L336-L347


 - [ ] ID-493
	- [SDLPoolCCIPControllerPrimary.handleOutgoingRESDL(uint64,address,uint256)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L103-L116)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L103-L116


 - [ ] ID-494
	- [StakingPool.removeStrategy(uint256,bytes)](contracts/core/StakingPool.sol#L237-L255)

contracts/core/StakingPool.sol#L237-L255


 - [ ] ID-495
	- [OperatorVCS.withdrawOperatorRewards(address,uint256)](contracts/linkStaking/OperatorVCS.sol#L102-L113)

contracts/linkStaking/OperatorVCS.sol#L102-L113


 - [ ] ID-496
	- [StakingRewardsMock.claimReward()](contracts/linkStaking/test/StakingRewardsMock.sol#L27-L32)

contracts/linkStaking/test/StakingRewardsMock.sol#L27-L32


 - [ ] ID-497
	- [StakingPool.withdraw(address,address,uint256)](contracts/core/StakingPool.sol#L97-L116)

contracts/core/StakingPool.sol#L97-L116


 - [ ] ID-498
	- [StrategyMock.deposit(uint256)](contracts/core/test/StrategyMock.sol#L45-L49)

contracts/core/test/StrategyMock.sol#L45-L49


 - [ ] ID-499
	- [SDLPool.setBoostController(address)](contracts/core/sdlPool/base/SDLPool.sol#L372-L374)

contracts/core/sdlPool/base/SDLPool.sol#L372-L374


 - [ ] ID-500
	- [VaultControllerStrategy.deposit(uint256)](contracts/linkStaking/base/VaultControllerStrategy.sol#L93-L110)

contracts/linkStaking/base/VaultControllerStrategy.sol#L93-L110


 - [ ] ID-501
	- [SDLPoolSecondary.handleOutgoingUpdate()](contracts/core/sdlPool/SDLPoolSecondary.sol#L313-L328)

contracts/core/sdlPool/SDLPoolSecondary.sol#L313-L328


 - [ ] ID-502
	- [SDLPool.setBaseURI(string)](contracts/core/sdlPool/base/SDLPool.sol#L363-L365)

contracts/core/sdlPool/base/SDLPool.sol#L363-L365


 - [ ] ID-503
	- [OperatorVCS.setOperatorRewardPercentage(uint256)](contracts/linkStaking/OperatorVCS.sol#L274-L281)

contracts/linkStaking/OperatorVCS.sol#L274-L281


 - [ ] ID-504
	- [DelegatorPool.retireDelegatorPool(address[],address)](contracts/core/test/deprecated/DelegatorPool.sol#L249-L276)

contracts/core/test/deprecated/DelegatorPool.sol#L249-L276


 - [ ] ID-505
	- [EthStakingStrategy.setNWLOperatorController(address)](contracts/ethStaking/EthStakingStrategy.sol#L295-L298)

contracts/ethStaking/EthStakingStrategy.sol#L295-L298


 - [ ] ID-506
	- [OperatorVault.withdrawRewards()](contracts/linkStaking/OperatorVault.sol#L139-L151)

contracts/linkStaking/OperatorVault.sol#L139-L151


 - [ ] ID-507
	- [RewardsPoolController.removeToken(address)](contracts/core/base/RewardsPoolController.sol#L171-L185)

contracts/core/base/RewardsPoolController.sol#L171-L185


 - [ ] ID-508
	- [SDLPoolPrimary.initiateUnlock(uint256)](contracts/core/sdlPool/SDLPoolPrimary.sol#L107-L121)

contracts/core/sdlPool/SDLPoolPrimary.sol#L107-L121


 - [ ] ID-509
	- [SDLPoolSecondary.handleOutgoingRESDL(address,uint256,address)](contracts/core/sdlPool/SDLPoolSecondary.sol#L259-L281)

contracts/core/sdlPool/SDLPoolSecondary.sol#L259-L281


 - [ ] ID-510
	- [GovernanceController.removeRoleFunctions(uint256,address[],bytes4[][])](contracts/governance/GovernanceController.sol#L180-L196)

contracts/governance/GovernanceController.sol#L180-L196


 - [ ] ID-511
	- [OperatorController.setKeyValidationOracle(address)](contracts/ethStaking/base/OperatorController.sol#L271-L274)

contracts/ethStaking/base/OperatorController.sol#L271-L274


 - [ ] ID-512
	- [OperatorWhitelist.setWLOperatorController(address)](contracts/ethStaking/OperatorWhitelist.sol#L76-L78)

contracts/ethStaking/OperatorWhitelist.sol#L76-L78


 - [ ] ID-513
	- [GovernanceController.addRoleFunctions(uint256,address[],bytes4[][])](contracts/governance/GovernanceController.sol#L156-L172)

contracts/governance/GovernanceController.sol#L156-L172


 - [ ] ID-514
	- [PriorityPool.setQueueDepositParams(uint128,uint128)](contracts/core/priorityPool/PriorityPool.sol#L458-L462)

contracts/core/priorityPool/PriorityPool.sol#L458-L462


 - [ ] ID-515
	- [StrategyMock.withdraw(uint256)](contracts/core/test/StrategyMock.sol#L51-L56)

contracts/core/test/StrategyMock.sol#L51-L56


 - [ ] ID-516
	- [MerkleDistributor.withdrawUnclaimedTokens(address,bytes32,uint256)](contracts/airdrop/MerkleDistributor.sol#L208-L224)

contracts/airdrop/MerkleDistributor.sol#L208-L224


 - [ ] ID-517
	- [LiquidSDIndexPool.setCompositionTargets(uint256[])](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L391-L402)

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L391-L402


 - [ ] ID-518
	- [SDLPoolSecondary.withdraw(uint256,uint256)](contracts/core/sdlPool/SDLPoolSecondary.sol#L214-L239)

contracts/core/sdlPool/SDLPoolSecondary.sol#L214-L239


 - [ ] ID-519
	- [SDLPoolCCIPControllerSecondary.setExtraArgs(bytes)](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L108-L111)

contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L108-L111


 - [ ] ID-520
	- [StakingPool.setPriorityPool(address)](contracts/core/StakingPool.sol#L375-L377)

contracts/core/StakingPool.sol#L375-L377


 - [ ] ID-521
	- [Vault.__Vault_init(address,address,address,address)](contracts/linkStaking/base/Vault.sol#L36-L48)

contracts/linkStaking/base/Vault.sol#L36-L48


 - [ ] ID-522
	- [SDLPoolCCIPControllerPrimary.addWhitelistedChain(uint64,address,bytes,bytes)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L151-L164)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L151-L164


 - [ ] ID-523
	- [VaultControllerStrategy.addFee(address,uint256)](contracts/linkStaking/base/VaultControllerStrategy.sol#L270-L273)

contracts/linkStaking/base/VaultControllerStrategy.sol#L270-L273


 - [ ] ID-524
	- [SDLPoolCCIPControllerPrimary.setUpdateExtraArgs(uint64,bytes)](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L212-L216)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L212-L216


## centralized-risk-other
Impact: Informational
Confidence: High
 - [ ] ID-525
	- [Operator.ownerTransferAndCall(address,uint256,bytes)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L332-L338)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L332-L338


 - [ ] ID-526
	- [Operator.ownerForward(address,bytes)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L319-L323)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L319-L323


 - [ ] ID-527
	- [ConfirmedOwnerWithProposal.transferOwnership(address)](node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L30-L32)

node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L30-L32


 - [ ] ID-528
	- [Operator.operatorRequest(address,uint256,bytes32,bytes4,uint256,uint256,bytes)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L131-L149)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L131-L149


 - [ ] ID-529
	- [LinkTokenReceiver.onTokenTransfer(address,uint256,bytes)](node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#L13-L27)

node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#L13-L27


 - [ ] ID-530
	- [Operator.oracleRequest(address,uint256,bytes32,address,bytes4,uint256,uint256,bytes)](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L98-L117)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L98-L117


 - [ ] ID-531
	- [StakedotlinkCouncil.mint(address,uint256)](contracts/governance/StakedotlinkCouncil.sol#L90-L92)

contracts/governance/StakedotlinkCouncil.sol#L90-L92


 - [ ] ID-532
	- [Ownable.renounceOwnership()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63


 - [ ] ID-533
	- [Ownable.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72


 - [ ] ID-534
	- [OwnableUpgradeable.transferOwnership(address)](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77


 - [ ] ID-535
	- [UUPSUpgradeable.upgradeToAndCall(address,bytes)](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L89-L92)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L89-L92


 - [ ] ID-536
	- [OwnableUpgradeable.renounceOwnership()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68


 - [ ] ID-537
	- [Router.ccipSend(uint64,Client.EVM2AnyMessage)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L104-L141)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L104-L141


 - [ ] ID-538
	- [CCIPReceiver.ccipReceive(Client.Any2EVMMessage)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L27-L29)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L27-L29


 - [ ] ID-539
	- [Router.routeMessage(Client.Any2EVMMessage,uint16,uint256,address)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L155-L214)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/Router.sol#L155-L214


 - [ ] ID-540
	- [UUPSUpgradeable.upgradeTo(address)](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L74-L77)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L74-L77


 - [ ] ID-541
	- [MerkleDistributor.addDistributions(address[],bytes32[],uint256[],uint256[])](contracts/airdrop/MerkleDistributor.sol#L54-L68)

contracts/airdrop/MerkleDistributor.sol#L54-L68


 - [ ] ID-542
	- [SDLPoolSecondary.onTokenTransfer(address,uint256,bytes)](contracts/core/sdlPool/SDLPoolSecondary.sol#L132-L151)

contracts/core/sdlPool/SDLPoolSecondary.sol#L132-L151


 - [ ] ID-543
	- [EthStakingStrategyMock.nwlWithdraw(address,uint256)](contracts/ethStaking/test/EthStakingStrategyMock.sol#L15-L17)

contracts/ethStaking/test/EthStakingStrategyMock.sol#L15-L17


 - [ ] ID-544
	- [OperatorVCS.onTokenTransfer(address,uint256,bytes)](contracts/linkStaking/OperatorVCS.sol#L84-L90)

contracts/linkStaking/OperatorVCS.sol#L84-L90


 - [ ] ID-545
	- [StakingAllowance.mint(address,uint256)](contracts/core/tokens/StakingAllowance.sol#L20-L22)

contracts/core/tokens/StakingAllowance.sol#L20-L22


 - [ ] ID-546
	- [SDLPoolCCIPControllerSecondary.handleOutgoingRESDL(uint64,address,uint256)](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L80-L86)

contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L80-L86


 - [ ] ID-547
	- [OperatorVCS.setOperator(uint256,address)](contracts/linkStaking/OperatorVCS.sol#L251-L253)

contracts/linkStaking/OperatorVCS.sol#L251-L253


 - [ ] ID-548
	- [RewardsPoolControllerV1.__RewardsPoolController_init(string,string)](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L30-L37)

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L30-L37


 - [ ] ID-549
	- [PriorityPool.onTokenTransfer(address,uint256,bytes)](contracts/core/priorityPool/PriorityPool.sol#L185-L202)

contracts/core/priorityPool/PriorityPool.sol#L185-L202


 - [ ] ID-550
	- [OperatorVCS.setRewardsReceiver(uint256,address)](contracts/linkStaking/OperatorVCS.sol#L261-L263)

contracts/linkStaking/OperatorVCS.sol#L261-L263


 - [ ] ID-551
	- [WrappedTokenBridge.recoverTokens(address[],address)](contracts/core/ccip/WrappedTokenBridge.sol#L140-L147)

contracts/core/ccip/WrappedTokenBridge.sol#L140-L147


 - [ ] ID-552
	- [WrappedTokenBridge.transferTokens(uint64,address,uint256,bool,uint256)](contracts/core/ccip/WrappedTokenBridge.sol#L106-L117)

contracts/core/ccip/WrappedTokenBridge.sol#L106-L117


 - [ ] ID-553
	- [PriorityPool.pauseForUpdate()](contracts/core/priorityPool/PriorityPool.sol#L429-L431)

contracts/core/priorityPool/PriorityPool.sol#L429-L431


 - [ ] ID-554
	- [CommunityVault.claimRewards(uint256,address)](contracts/linkStaking/CommunityVault.sol#L39-L45)

contracts/linkStaking/CommunityVault.sol#L39-L45


 - [ ] ID-555
	- [WrappedSDTokenMock.onTokenTransfer(address,uint256,bytes)](contracts/core/test/WrappedSDTokenMock.sol#L24-L32)

contracts/core/test/WrappedSDTokenMock.sol#L24-L32


 - [ ] ID-556
	- [SDLPoolCCIPController.ccipReceive(Client.Any2EVMMessage)](contracts/core/ccip/base/SDLPoolCCIPController.sol#L102-L110)

contracts/core/ccip/base/SDLPoolCCIPController.sol#L102-L110


 - [ ] ID-557
	- [RewardsInitiator.updateRewards(uint256[],bytes)](contracts/core/RewardsInitiator.sol#L38-L42)

contracts/core/RewardsInitiator.sol#L38-L42


 - [ ] ID-558
	- [PoolAllowanceV1.mintAllowance(address,uint256)](contracts/core/test/deprecated/PoolAllowanceV1.sol#L31-L33)

contracts/core/test/deprecated/PoolAllowanceV1.sol#L31-L33


 - [ ] ID-559
	- [DistributionOracle.pauseForUpdate()](contracts/core/priorityPool/DistributionOracle.sol#L142-L144)

contracts/core/priorityPool/DistributionOracle.sol#L142-L144


 - [ ] ID-560
	- [CommunityVCS.addVaults(uint256)](contracts/linkStaking/CommunityVCS.sol#L106-L108)

contracts/linkStaking/CommunityVCS.sol#L106-L108


 - [ ] ID-561
	- [SDLPoolCCIPController.ccipSend(uint64,Client.EVM2AnyMessage)](contracts/core/ccip/base/SDLPoolCCIPController.sol#L89-L100)

contracts/core/ccip/base/SDLPoolCCIPController.sol#L89-L100


 - [ ] ID-562
	- [SDLPoolCCIPController.recoverTokens(address[],address)](contracts/core/ccip/base/SDLPoolCCIPController.sol#L117-L124)

contracts/core/ccip/base/SDLPoolCCIPController.sol#L117-L124


 - [ ] ID-563
	- [GovernanceController.callFunction(uint256,address,bytes)](contracts/governance/GovernanceController.sol#L73-L90)

contracts/governance/GovernanceController.sol#L73-L90


 - [ ] ID-564
	- [DistributionOracle.requestUpdate()](contracts/core/priorityPool/DistributionOracle.sol#L151-L153)

contracts/core/priorityPool/DistributionOracle.sol#L151-L153


 - [ ] ID-565
	- [OperatorVCSUpgrade.addVault(address)](contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L75-L85)

contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L75-L85


 - [ ] ID-566
	- [EthStakingStrategy.nwlWithdraw(address,uint256)](contracts/ethStaking/EthStakingStrategy.sol#L234-L237)

contracts/ethStaking/EthStakingStrategy.sol#L234-L237


 - [ ] ID-567
	- [VaultControllerStrategy.upgradeVaults(uint256,uint256,bytes)](contracts/linkStaking/base/VaultControllerStrategy.sol#L243-L252)

contracts/linkStaking/base/VaultControllerStrategy.sol#L243-L252


 - [ ] ID-568
	- [Vault.deposit(uint256)](contracts/linkStaking/base/Vault.sol#L62-L65)

contracts/linkStaking/base/Vault.sol#L62-L65


 - [ ] ID-569
	- [StakingPool.strategyDeposit(uint256,uint256)](contracts/core/StakingPool.sol#L123-L126)

contracts/core/StakingPool.sol#L123-L126


 - [ ] ID-570
	- [RewardsPoolController.__RewardsPoolController_init()](contracts/core/base/RewardsPoolController.sol#L28-L31)

contracts/core/base/RewardsPoolController.sol#L28-L31


 - [ ] ID-571
	- [OperatorVaultV1.raiseAlert()](contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L51-L55)

contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L51-L55


 - [ ] ID-572
	- [StakingPool.updateStrategyRewards(uint256[],bytes)](contracts/core/StakingPool.sol#L345-L348)

contracts/core/StakingPool.sol#L345-L348


 - [ ] ID-573
	- [LPLMigration.onTokenTransfer(address,uint256,bytes)](contracts/core/tokens/LPLMigration.sol#L31-L39)

contracts/core/tokens/LPLMigration.sol#L31-L39


 - [ ] ID-574
	- [SDLPoolCCIPControllerPrimary.approveRewardTokens(address[])](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L190-L195)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L190-L195


 - [ ] ID-575
	- [SDLPoolCCIPControllerMock.handleOutgoingRESDL(uint64,address,uint256)](contracts/core/test/SDLPoolCCIPControllerMock.sol#L29-L35)

contracts/core/test/SDLPoolCCIPControllerMock.sol#L29-L35


 - [ ] ID-576
	- [OperatorVCSUpgrade.setOperator(uint256,address)](contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L92-L94)

contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L92-L94


 - [ ] ID-577
	- [KeyValidationOracle.reportKeyPairValidation(bytes32,uint256,bool,bool)](contracts/ethStaking/KeyValidationOracle.sol#L68-L79)

contracts/ethStaking/KeyValidationOracle.sol#L68-L79


 - [ ] ID-578
	- [SDLPoolPrimary.onTokenTransfer(address,uint256,bytes)](contracts/core/sdlPool/SDLPoolPrimary.sol#L61-L80)

contracts/core/sdlPool/SDLPoolPrimary.sol#L61-L80


 - [ ] ID-579
	- [VaultControllerStrategyUpgrade.migrateVaults(uint256,uint256,bytes)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L274-L283)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L274-L283


 - [ ] ID-580
	- [StakingPool.strategyWithdraw(uint256,uint256)](contracts/core/StakingPool.sol#L133-L136)

contracts/core/StakingPool.sol#L133-L136


 - [ ] ID-581
	- [RewardsPool.onTokenTransfer(address,uint256,bytes)](contracts/core/RewardsPool.sol#L65-L72)

contracts/core/RewardsPool.sol#L65-L72


 - [ ] ID-582
	- [NWLOperatorController.addKeyPairs(uint256,uint256,bytes,bytes)](contracts/ethStaking/NWLOperatorController.sol#L83-L92)

contracts/ethStaking/NWLOperatorController.sol#L83-L92


 - [ ] ID-583
	- [OperatorController.onTokenTransfer(address,uint256,bytes)](contracts/ethStaking/base/OperatorController.sol#L176-L183)

contracts/ethStaking/base/OperatorController.sol#L176-L183


 - [ ] ID-584
	- [PoolOwnersV1.onTokenTransfer(address,uint256,bytes)](contracts/core/test/deprecated/PoolOwnersV1.sol#L94-L101)

contracts/core/test/deprecated/PoolOwnersV1.sol#L94-L101


 - [ ] ID-585
	- [OperatorVault.raiseAlert(address)](contracts/linkStaking/OperatorVault.sol#L114-L124)

contracts/linkStaking/OperatorVault.sol#L114-L124


 - [ ] ID-586
	- [StakingAllowance.mintToContract(address,address,uint256,bytes)](contracts/core/tokens/StakingAllowance.sol#L30-L38)

contracts/core/tokens/StakingAllowance.sol#L30-L38


 - [ ] ID-587
	- [SDLPoolCCIPControllerSecondary.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken)](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L94-L102)

contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L94-L102


 - [ ] ID-588
	- [DepositController.depositEther(bytes32,bytes32,bytes32,uint256,uint256,uint256[],uint256[])](contracts/ethStaking/DepositController.sol#L46-L64)

contracts/ethStaking/DepositController.sol#L46-L64


 - [ ] ID-589
	- [GovernanceController.revokeRole(uint256,address)](contracts/governance/GovernanceController.sol#L138-L140)

contracts/governance/GovernanceController.sol#L138-L140


 - [ ] ID-590
	- [OwnersRewardsPoolV1.onTokenTransfer(address,uint256,bytes)](contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L65-L72)

contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L65-L72


 - [ ] ID-591
	- [SDLPool.addToken(address,address)](contracts/core/sdlPool/base/SDLPool.sol#L336-L339)

contracts/core/sdlPool/base/SDLPool.sol#L336-L339


 - [ ] ID-592
	- [SDLPoolCCIPControllerPrimary.distributeRewards()](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L56-L93)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L56-L93


 - [ ] ID-593
	- [VaultV1.deposit(uint256)](contracts/linkStaking/test/deprecated/VaultV1.sol#L73-L76)

contracts/linkStaking/test/deprecated/VaultV1.sol#L73-L76


 - [ ] ID-594
	- [WLOperatorController.addKeyPairs(uint256,uint256,bytes,bytes)](contracts/ethStaking/WLOperatorController.sol#L60-L68)

contracts/ethStaking/WLOperatorController.sol#L60-L68


 - [ ] ID-595
	- [SDLPoolCCIPControllerMock.handleIncomingRESDL(uint64,address,uint256,ISDLPool.RESDLToken)](contracts/core/test/SDLPoolCCIPControllerMock.sol#L37-L45)

contracts/core/test/SDLPoolCCIPControllerMock.sol#L37-L45


 - [ ] ID-596
	- [RewardsPool.withdraw(address)](contracts/core/RewardsPool.sol#L57-L60)

contracts/core/RewardsPool.sol#L57-L60


 - [ ] ID-597
	- [RESDLTokenBridge.ccipReceive(Client.Any2EVMMessage)](contracts/core/ccip/RESDLTokenBridge.sol#L171-L192)

contracts/core/ccip/RESDLTokenBridge.sol#L171-L192


 - [ ] ID-598
	- [MerkleDistributor.updateDistributions(address[],bytes32[],uint256[],uint256[])](contracts/airdrop/MerkleDistributor.sol#L102-L116)

contracts/airdrop/MerkleDistributor.sol#L102-L116


 - [ ] ID-599
	- [SDLPoolPrimary.migrate(address,uint256,uint64)](contracts/core/sdlPool/SDLPoolPrimary.sol#L264-L272)

contracts/core/sdlPool/SDLPoolPrimary.sol#L264-L272


 - [ ] ID-600
	- [PoolAllowanceV1.burnAllowance(address,uint256)](contracts/core/test/deprecated/PoolAllowanceV1.sol#L40-L42)

contracts/core/test/deprecated/PoolAllowanceV1.sol#L40-L42


 - [ ] ID-601
	- [EthStakingStrategy.withdraw(uint256)](contracts/ethStaking/EthStakingStrategy.sol#L224-L226)

contracts/ethStaking/EthStakingStrategy.sol#L224-L226


 - [ ] ID-602
	- [WrappedSDToken.onTokenTransfer(address,uint256,bytes)](contracts/core/tokens/WrappedSDToken.sol#L27-L34)

contracts/core/tokens/WrappedSDToken.sol#L27-L34


 - [ ] ID-603
	- [WrappedTokenBridge.onTokenTransfer(address,uint256,bytes)](contracts/core/ccip/WrappedTokenBridge.sol#L83-L96)

contracts/core/ccip/WrappedTokenBridge.sol#L83-L96


 - [ ] ID-604
	- [ERC677Upgradeable.__ERC677_init(string,string,uint256)](contracts/core/tokens/base/ERC677Upgradeable.sol#L9-L16)

contracts/core/tokens/base/ERC677Upgradeable.sol#L9-L16


 - [ ] ID-605
	- [VaultControllerStrategyUpgrade.upgradeVaults(uint256,uint256,bytes)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L291-L300)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L291-L300


## error-msg
Impact: Informational
Confidence: Medium
 - [ ] ID-606
require() missing error messages
	 - [require(bool)(denominator > prod1)](node_modules/@openzeppelin/contracts/utils/math/Math.sol#L78)

node_modules/@openzeppelin/contracts/utils/math/Math.sol#L78


 - [ ] ID-607
require() missing error messages
	 - [require(bool)(len <= data.length)](node_modules/@chainlink/contracts/src/v0.8/vendor/BufferChainlink.sol#L98)

node_modules/@chainlink/contracts/src/v0.8/vendor/BufferChainlink.sol#L98


## price-manipulation-info
Impact: Informational
Confidence: Medium
 - [ ] ID-608
Potential price manipulation risk:
	- In function withdraw
		-- [toWithdraw = balanceOf(msg.sender)](contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L37) have potential price manipulated risk from toWithdraw and call None which could influence variable:toWithdraw

contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L37


 - [ ] ID-609
Potential price manipulation risk:
	- In function getEthBalance
		-- [balance = addr.balance](contracts/core/test/Multicall3.sol#L222) have potential price manipulated risk from balance and call balance(address) which could influence variable:balance

contracts/core/test/Multicall3.sol#L222


 - [ ] ID-610
Potential price manipulation risk:
	- In function getLockIdsByOwner
		-- [lockCount = balanceOf(_owner)](contracts/core/sdlPool/base/SDLPool.sol#L179) have potential price manipulated risk from lockCount and call None which could influence variable:lockIds

contracts/core/sdlPool/base/SDLPool.sol#L179


 - [ ] ID-611
Potential price manipulation risk:
	- In function canWithdraw
		-- [stLINKCanWithdraw = MathUpgradeable.min(stakingPool.balanceOf(_account),stakingPool.canWithdraw() + totalQueued - canUnqueue)](contracts/core/priorityPool/PriorityPool.sol#L171-L174) have potential price manipulated risk from stLINKCanWithdraw and call None which could influence variable:stLINKCanWithdraw

contracts/core/priorityPool/PriorityPool.sol#L171-L174


## uncontrolled-resource-consumption
Impact: Informational
Confidence: Medium
 - [ ] ID-612
Potential DoS Gas Limit Attack occur in[DepositContract.deposit(bytes,bytes,bytes,bytes32)](contracts/ethStaking/test/DepositContract.sol#L89-L145)[BEGIN_LOOP](contracts/ethStaking/test/DepositContract.sol#L134-L141)

contracts/ethStaking/test/DepositContract.sol#L89-L145


 - [ ] ID-613
Potential DoS Gas Limit Attack occur in[DepositContract.constructor()](contracts/ethStaking/test/DepositContract.sol#L68-L72)[BEGIN_LOOP](contracts/ethStaking/test/DepositContract.sol#L70-L71)

contracts/ethStaking/test/DepositContract.sol#L68-L72


 - [ ] ID-614
Potential DoS Gas Limit Attack occur in[Operator.distributeFunds(address[],uint256[])](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L349-L358)[BEGIN_LOOP](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L352-L356)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L349-L358


 - [ ] ID-615
Potential DoS Gas Limit Attack occur in[BytesLib.concatStorage(bytes,bytes)](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L91-L226)[BEGIN_LOOP](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L177-L185)

node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L91-L226


 - [ ] ID-616
Potential DoS Gas Limit Attack occur in[SDLPoolCCIPControllerPrimary._distributeRewards(uint64,address[],uint256[])](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L244-L287)[BEGIN_LOOP](contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L253-L257)

contracts/core/ccip/SDLPoolCCIPControllerPrimary.sol#L244-L287


 - [ ] ID-617
Potential DoS Gas Limit Attack occur in[OperatorVCS._updateStrategyRewards()](contracts/linkStaking/OperatorVCS.sol#L296-L303)[BEGIN_LOOP](contracts/linkStaking/OperatorVCS.sol#L299-L301)

contracts/linkStaking/OperatorVCS.sol#L296-L303


 - [ ] ID-618
Potential DoS Gas Limit Attack occur in[OperatorController._storeKeyPair(uint256,uint256,bytes,bytes)](contracts/ethStaking/base/OperatorController.sol#L346-L371)[BEGIN_LOOP](contracts/ethStaking/base/OperatorController.sol#L365-L370)

contracts/ethStaking/base/OperatorController.sol#L346-L371


 - [ ] ID-619
Potential DoS Gas Limit Attack occur in[OperatorWhitelistMock.constructor(address[])](contracts/ethStaking/test/OperatorWhitelistMock.sol#L12-L16)[BEGIN_LOOP](contracts/ethStaking/test/OperatorWhitelistMock.sol#L13-L15)

contracts/ethStaking/test/OperatorWhitelistMock.sol#L12-L16


 - [ ] ID-620
Potential DoS Gas Limit Attack occur in[StakingPool.depositLiquidity()](contracts/core/StakingPool.sol#L354-L369)[BEGIN_LOOP](contracts/core/StakingPool.sol#L357-L367)

contracts/core/StakingPool.sol#L354-L369


 - [ ] ID-621
Potential DoS Gas Limit Attack occur in[RewardsPoolControllerV1.distributeTokens(address[])](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L126-L130)[BEGIN_LOOP](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L127-L129)

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L126-L130


 - [ ] ID-622
Potential DoS Gas Limit Attack occur in[CCIPOffRampMock.constructor(address,address[],address[])](contracts/core/test/chainlink/CCIPOffRampMock.sol#L23-L32)[BEGIN_LOOP](contracts/core/test/chainlink/CCIPOffRampMock.sol#L29-L31)

contracts/core/test/chainlink/CCIPOffRampMock.sol#L23-L32


 - [ ] ID-623
Potential DoS Gas Limit Attack occur in[RewardsPoolControllerV1._updateRewards(address)](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L215-L219)[BEGIN_LOOP](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L216-L218)

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L215-L219


 - [ ] ID-624
Potential DoS Gas Limit Attack occur in[VaultControllerStrategyUpgrade._depositToVaults(uint256,uint256,uint256,uint256)](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L387-L411)[BEGIN_LOOP](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L394-L409)

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L387-L411


 - [ ] ID-625
Potential DoS Gas Limit Attack occur in[PoolOwnersV1._burnAllowance(address)](contracts/core/test/deprecated/PoolOwnersV1.sol#L228-L238)[BEGIN_LOOP](contracts/core/test/deprecated/PoolOwnersV1.sol#L230-L237)

contracts/core/test/deprecated/PoolOwnersV1.sol#L228-L238


 - [ ] ID-626
Potential DoS Gas Limit Attack occur in[RewardsPoolController.distributeTokens(address[])](contracts/core/base/RewardsPoolController.sol#L102-L106)[BEGIN_LOOP](contracts/core/base/RewardsPoolController.sol#L103-L105)

contracts/core/base/RewardsPoolController.sol#L102-L106


 - [ ] ID-627
Potential DoS Gas Limit Attack occur in[CommunityVCS._deployVaults(uint256)](contracts/linkStaking/CommunityVCS.sol#L125-L136)[BEGIN_LOOP](contracts/linkStaking/CommunityVCS.sol#L133-L135)

contracts/linkStaking/CommunityVCS.sol#L125-L136


 - [ ] ID-628
Potential DoS Gas Limit Attack occur in[OperatorWhitelist.constructor(address,address[])](contracts/ethStaking/OperatorWhitelist.sol#L19-L25)[BEGIN_LOOP](contracts/ethStaking/OperatorWhitelist.sol#L22-L24)

contracts/ethStaking/OperatorWhitelist.sol#L19-L25


 - [ ] ID-629
Potential DoS Gas Limit Attack occur in[SDLPoolSecondary._executeQueuedLockUpdates(address,uint256[])](contracts/core/sdlPool/SDLPoolSecondary.sol#L451-L510)[BEGIN_LOOP](contracts/core/sdlPool/SDLPoolSecondary.sol#L454-L509)

contracts/core/sdlPool/SDLPoolSecondary.sol#L451-L510


 - [ ] ID-630
Potential DoS Gas Limit Attack occur in[OperatorController._addKeyPairs(uint256,uint256,bytes,bytes)](contracts/ethStaking/base/OperatorController.sol#L312-L337)[BEGIN_LOOP](contracts/ethStaking/base/OperatorController.sol#L324-L331)

contracts/ethStaking/base/OperatorController.sol#L312-L337


 - [ ] ID-631
Potential DoS Gas Limit Attack occur in[Multicall3.aggregate(Multicall3.Call[])](contracts/core/test/Multicall3.sol#L41-L55)[BEGIN_LOOP](contracts/core/test/Multicall3.sol#L46-L54)

contracts/core/test/Multicall3.sol#L41-L55


 - [ ] ID-632
Potential DoS Gas Limit Attack occur in[CommunityVCS.claimRewards(uint256,uint256,uint256)](contracts/linkStaking/CommunityVCS.sol#L65-L74)[BEGIN_LOOP](contracts/linkStaking/CommunityVCS.sol#L71-L73)

contracts/linkStaking/CommunityVCS.sol#L65-L74


 - [ ] ID-633
Potential DoS Gas Limit Attack occur in[Multicall3.aggregate3Value(Multicall3.Call3Value[])](contracts/core/test/Multicall3.sol#L151-L187)[BEGIN_LOOP](contracts/core/test/Multicall3.sol#L156-L184)

contracts/core/test/Multicall3.sol#L151-L187


 - [ ] ID-634
Potential DoS Gas Limit Attack occur in[SDLPoolCCIPControllerSecondary._ccipReceive(Client.Any2EVMMessage)](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L147-L165)[BEGIN_LOOP](contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L152-L155)

contracts/core/ccip/SDLPoolCCIPControllerSecondary.sol#L147-L165


 - [ ] ID-635
Potential DoS Gas Limit Attack occur in[CCIPOnRampMock.constructor(address[],address[],address)](contracts/core/test/chainlink/CCIPOnRampMock.sol#L22-L31)[BEGIN_LOOP](contracts/core/test/chainlink/CCIPOnRampMock.sol#L27-L29)

contracts/core/test/chainlink/CCIPOnRampMock.sol#L22-L31


 - [ ] ID-636
Potential DoS Gas Limit Attack occur in[MerkleDistributor.claimDistributions(address[],uint256[],address,uint256[],bytes32[][])](contracts/airdrop/MerkleDistributor.sol#L153-L168)[BEGIN_LOOP](contracts/airdrop/MerkleDistributor.sol#L165-L167)

contracts/airdrop/MerkleDistributor.sol#L153-L168


 - [ ] ID-637
Potential DoS Gas Limit Attack occur in[Multicall3.aggregate3(Multicall3.Call3[])](contracts/core/test/Multicall3.sol#L118-L145)[BEGIN_LOOP](contracts/core/test/Multicall3.sol#L122-L144)

contracts/core/test/Multicall3.sol#L118-L145


 - [ ] ID-638
Potential DoS Gas Limit Attack occur in[OperatorControllerMock.assignNextValidators(uint256[],uint256[],uint256)](contracts/ethStaking/test/OperatorControllerMock.sol#L54-L67)[BEGIN_LOOP](contracts/ethStaking/test/OperatorControllerMock.sol#L59-L66)

contracts/ethStaking/test/OperatorControllerMock.sol#L54-L67


 - [ ] ID-639
Potential DoS Gas Limit Attack occur in[VaultControllerStrategy._depositToVaults(uint256,uint256,uint256,uint256)](contracts/linkStaking/base/VaultControllerStrategy.sol#L334-L362)[BEGIN_LOOP](contracts/linkStaking/base/VaultControllerStrategy.sol#L342-L359)

contracts/linkStaking/base/VaultControllerStrategy.sol#L334-L362


 - [ ] ID-640
Potential DoS Gas Limit Attack occur in[VCSMock.addVaults(address[])](contracts/linkStaking/test/VCSMock.sol#L26-L32)[BEGIN_LOOP](contracts/linkStaking/test/VCSMock.sol#L27-L31)

contracts/linkStaking/test/VCSMock.sol#L26-L32


 - [ ] ID-641
Potential DoS Gas Limit Attack occur in[SDLPoolSecondary._mintQueuedNewLocks(address)](contracts/core/sdlPool/SDLPoolSecondary.sol#L384-L419)[BEGIN_LOOP](contracts/core/sdlPool/SDLPoolSecondary.sol#L388-L409)

contracts/core/sdlPool/SDLPoolSecondary.sol#L384-L419


 - [ ] ID-642
Potential DoS Gas Limit Attack occur in[RewardsPoolController._updateRewards(address)](contracts/core/base/RewardsPoolController.sol#L191-L195)[BEGIN_LOOP](contracts/core/base/RewardsPoolController.sol#L192-L194)

contracts/core/base/RewardsPoolController.sol#L191-L195


 - [ ] ID-643
Potential DoS Gas Limit Attack occur in[RewardsInitiator.performUpkeep(bytes)](contracts/core/RewardsInitiator.sol#L83-L94)[BEGIN_LOOP](contracts/core/RewardsInitiator.sol#L89-L91)

contracts/core/RewardsInitiator.sol#L83-L94


 - [ ] ID-644
Potential DoS Gas Limit Attack occur in[StakingPool._withdrawLiquidity(uint256)](contracts/core/StakingPool.sol#L402-L417)[BEGIN_LOOP](contracts/core/StakingPool.sol#L405-L416)

contracts/core/StakingPool.sol#L402-L417


 - [ ] ID-645
Potential DoS Gas Limit Attack occur in[Multicall3.tryAggregate(bool,Multicall3.Call[])](contracts/core/test/Multicall3.sol#L62-L75)[BEGIN_LOOP](contracts/core/test/Multicall3.sol#L66-L74)

contracts/core/test/Multicall3.sol#L62-L75


 - [ ] ID-646
Potential DoS Gas Limit Attack occur in[PoolOwnersV1._mintAllowance(address)](contracts/core/test/deprecated/PoolOwnersV1.sol#L212-L222)[BEGIN_LOOP](contracts/core/test/deprecated/PoolOwnersV1.sol#L214-L221)

contracts/core/test/deprecated/PoolOwnersV1.sol#L212-L222


## no-derived-function
Impact: Informational
Confidence: High
 - [ ] ID-647
[CCIPReceiver](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L11-L52) does not implement functions:
	- [CCIPReceiver._ccipReceive(Client.Any2EVMMessage)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L33)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L11-L52


 - [ ] ID-648
[UUPSUpgradeable](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L22-L112) does not implement functions:
	- [UUPSUpgradeable._authorizeUpgrade(address)](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L104)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L22-L112


## unnecessary-public-function-modifier
Impact: Informational
Confidence: High
 - [ ] ID-649
function:[LinkTokenReceiver.getChainlinkToken()](node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#L29)is public and can be replaced with external 

node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#L29


 - [ ] ID-650
function:[LinkTokenReceiver.onTokenTransfer(address,uint256,bytes)](node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#L13-L27)is public and can be replaced with external 

node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#L13-L27


 - [ ] ID-651
function:[ConfirmedOwnerWithProposal.owner()](node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L50-L52)is public and can be replaced with external 

node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L50-L52


 - [ ] ID-652
function:[Operator.getChainlinkToken()](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L420-L422)is public and can be replaced with external 

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L420-L422


 - [ ] ID-653
function:[ConfirmedOwnerWithProposal.transferOwnership(address)](node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L30-L32)is public and can be replaced with external 

node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L30-L32


 - [ ] ID-654
function:[StakedotlinkCouncil.mintWithTokenURI(address,uint256,string)](contracts/governance/StakedotlinkCouncil.sol#L100-L111)is public and can be replaced with external 

contracts/governance/StakedotlinkCouncil.sol#L100-L111


 - [ ] ID-655
function:[ERC20Upgradeable.transfer(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L118-L122)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L118-L122


 - [ ] ID-656
function:[ERC721.ownerOf(uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L70-L74)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L70-L74


 - [ ] ID-657
function:[ERC721.name()](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L79-L81)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L79-L81


 - [ ] ID-658
function:[ERC20Upgradeable.decimals()](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L92-L94)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L92-L94


 - [ ] ID-659
function:[StakedotlinkCouncil.tokenURI(uint256)](contracts/governance/StakedotlinkCouncil.sol#L166-L170)is public and can be replaced with external 

contracts/governance/StakedotlinkCouncil.sol#L166-L170


 - [ ] ID-660
function:[ERC20Upgradeable.approve(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L141-L145)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L141-L145


 - [ ] ID-661
function:[OwnableUpgradeable.renounceOwnership()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68


 - [ ] ID-662
function:[ERC20.symbol()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L70-L72)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L70-L72


 - [ ] ID-663
function:[ERC20Upgradeable.totalSupply()](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L99-L101)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L99-L101


 - [ ] ID-664
function:[ERC20.totalSupply()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L94-L96)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L94-L96


 - [ ] ID-665
function:[ERC721.balanceOf(address)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L62-L65)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L62-L65


 - [ ] ID-666
function:[ERC20Upgradeable.symbol()](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L75-L77)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L75-L77


 - [ ] ID-667
function:[ERC20.approve(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L136-L140)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L136-L140


 - [ ] ID-668
function:[ERC20.balanceOf(address)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L101-L103)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L101-L103


 - [ ] ID-669
function:[StakedotlinkCouncil.totalSupply()](contracts/governance/StakedotlinkCouncil.sol#L157-L159)is public and can be replaced with external 

contracts/governance/StakedotlinkCouncil.sol#L157-L159


 - [ ] ID-670
function:[StakedotlinkCouncil.transferFrom(address,address,uint256)](contracts/governance/StakedotlinkCouncil.sol#L68-L82)is public and can be replaced with external 

contracts/governance/StakedotlinkCouncil.sol#L68-L82


 - [ ] ID-671
function:[OwnableUpgradeable.transferOwnership(address)](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77


 - [ ] ID-672
function:[ERC20Upgradeable.name()](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L67-L69)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L67-L69


 - [ ] ID-673
function:[ERC721.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L150-L159)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L150-L159


 - [ ] ID-674
function:[StakedotlinkCouncil.burn(uint256)](contracts/governance/StakedotlinkCouncil.sol#L131-L152)is public and can be replaced with external 

contracts/governance/StakedotlinkCouncil.sol#L131-L152


 - [ ] ID-675
function:[ERC20.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L158-L167)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L158-L167


 - [ ] ID-676
function:[Ownable.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72


 - [ ] ID-677
function:[ERC20.name()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L62-L64)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L62-L64


 - [ ] ID-678
function:[StakedotlinkCouncil.balanceOf(address)](contracts/governance/StakedotlinkCouncil.sol#L57-L59)is public and can be replaced with external 

contracts/governance/StakedotlinkCouncil.sol#L57-L59


 - [ ] ID-679
function:[ERC721.symbol()](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L86-L88)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L86-L88


 - [ ] ID-680
function:[ERC20Upgradeable.balanceOf(address)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L106-L108)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L106-L108


 - [ ] ID-681
function:[Ownable.renounceOwnership()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63


 - [ ] ID-682
function:[ERC20.increaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L181-L185)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L181-L185


 - [ ] ID-683
function:[CCIPReceiver.supportsInterface(bytes4)](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L22-L24)is public and can be replaced with external 

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L22-L24


 - [ ] ID-684
function:[ERC165.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L26-L28)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L26-L28


 - [ ] ID-685
function:[ERC721.safeTransferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L164-L170)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L164-L170


 - [ ] ID-686
function:[StakedotlinkCouncil.setTokenURI(uint256,string)](contracts/governance/StakedotlinkCouncil.sol#L178-L182)is public and can be replaced with external 

contracts/governance/StakedotlinkCouncil.sol#L178-L182


 - [ ] ID-687
function:[UUPSUpgradeable.upgradeToAndCall(address,bytes)](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L89-L92)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L89-L92


 - [ ] ID-688
function:[ERC721.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L52-L57)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L52-L57


 - [ ] ID-689
function:[ERC721.tokenURI(uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L93-L98)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L93-L98


 - [ ] ID-690
function:[StakedotlinkCouncil.mint(address,uint256)](contracts/governance/StakedotlinkCouncil.sol#L90-L92)is public and can be replaced with external 

contracts/governance/StakedotlinkCouncil.sol#L90-L92


 - [ ] ID-691
function:[ERC721.approve(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L112-L122)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L112-L122


 - [ ] ID-692
function:[ERC20.decreaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L201-L210)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L201-L210


 - [ ] ID-693
function:[VestingWallet.release()](node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L89-L94)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L89-L94


 - [ ] ID-694
function:[ERC20.transfer(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L113-L117)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L113-L117


 - [ ] ID-695
function:[ERC721URIStorage.tokenURI(uint256)](node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L20-L36)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L20-L36


 - [ ] ID-696
function:[CCIPReceiver.getRouter()](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L41-L43)is public and can be replaced with external 

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L41-L43


 - [ ] ID-697
function:[ERC20Upgradeable.increaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L182-L186)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L182-L186


 - [ ] ID-698
function:[UUPSUpgradeable.upgradeTo(address)](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L74-L77)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L74-L77


 - [ ] ID-699
function:[ERC20.decimals()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L87-L89)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L87-L89


 - [ ] ID-700
function:[ERC20Upgradeable.decreaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L202-L211)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L202-L211


 - [ ] ID-701
function:[ERC20Upgradeable.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L163-L168)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L163-L168


 - [ ] ID-702
function:[ERC721.setApprovalForAll(address,bool)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L136-L138)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L136-L138


 - [ ] ID-703
function:[VestingWallet.release(address)](node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L101-L106)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L101-L106


 - [ ] ID-704
function:[ENSResolver.addr(bytes32)](node_modules/@chainlink/contracts/src/v0.8/vendor/ENSResolver.sol#L5)is public and can be replaced with external 

node_modules/@chainlink/contracts/src/v0.8/vendor/ENSResolver.sol#L5


 - [ ] ID-705
function:[PoolOwnersV1.totalSupply()](contracts/core/test/deprecated/PoolOwnersV1.sol#L85-L87)is public and can be replaced with external 

contracts/core/test/deprecated/PoolOwnersV1.sol#L85-L87


 - [ ] ID-706
function:[Multicall3.getBasefee()](contracts/core/test/Multicall3.sol#L234-L236)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L234-L236


 - [ ] ID-707
function:[Multicall3.aggregate3(Multicall3.Call3[])](contracts/core/test/Multicall3.sol#L118-L145)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L118-L145


 - [ ] ID-708
function:[VCSMock.initialize(address,address,address,address,VaultControllerStrategy.Fee[])](contracts/linkStaking/test/VCSMock.sol#L16-L24)is public and can be replaced with external 

contracts/linkStaking/test/VCSMock.sol#L16-L24


 - [ ] ID-709
function:[StakingAllowance.burnFrom(address,uint256)](contracts/core/tokens/StakingAllowance.sol#L52-L55)is public and can be replaced with external 

contracts/core/tokens/StakingAllowance.sol#L52-L55


 - [ ] ID-710
function:[StrategyMock.initialize(address,address,uint256,uint256)](contracts/core/test/StrategyMock.sol#L28-L38)is public and can be replaced with external 

contracts/core/test/StrategyMock.sol#L28-L38


 - [ ] ID-711
function:[SDLPool.__SDLPoolBase_init(string,string,address,address)](contracts/core/sdlPool/base/SDLPool.sol#L93-L104)is public and can be replaced with external 

contracts/core/sdlPool/base/SDLPool.sol#L93-L104


 - [ ] ID-712
function:[RocketPoolLSDIndexAdapter.initialize(address,address)](contracts/liquidSDIndex/adapters/RocketPoolLSDIndexAdapter.sol#L17-L19)is public and can be replaced with external 

contracts/liquidSDIndex/adapters/RocketPoolLSDIndexAdapter.sol#L17-L19


 - [ ] ID-713
function:[RewardsPoolController.distributeTokens(address[])](contracts/core/base/RewardsPoolController.sol#L102-L106)is public and can be replaced with external 

contracts/core/base/RewardsPoolController.sol#L102-L106


 - [ ] ID-714
function:[RewardsPool.updateReward(address)](contracts/core/RewardsPool.sol#L89-L95)is public and can be replaced with external 

contracts/core/RewardsPool.sol#L89-L95


 - [ ] ID-715
function:[OperatorController.totalStaked()](contracts/ethStaking/base/OperatorController.sol#L98-L100)is public and can be replaced with external 

contracts/ethStaking/base/OperatorController.sol#L98-L100


 - [ ] ID-716
function:[DistributionOracle.fulfillRequest(bytes32,bytes32,bytes32,uint256,uint256)](contracts/core/priorityPool/DistributionOracle.sol#L163-L177)is public and can be replaced with external 

contracts/core/priorityPool/DistributionOracle.sol#L163-L177


 - [ ] ID-717
function:[WLOperatorController.initialize(address,address,address,uint256)](contracts/ethStaking/WLOperatorController.sol#L33-L42)is public and can be replaced with external 

contracts/ethStaking/WLOperatorController.sol#L33-L42


 - [ ] ID-718
function:[Multicall3.getBlockHash(uint256)](contracts/core/test/Multicall3.sol#L191-L193)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L191-L193


 - [ ] ID-719
function:[LSDIndexAdapter.getExchangeRate()](contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L66)is public and can be replaced with external 

contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L66


 - [ ] ID-720
function:[Strategy.__Strategy_init(address,address)](contracts/core/base/Strategy.sol#L20-L25)is public and can be replaced with external 

contracts/core/base/Strategy.sol#L20-L25


 - [ ] ID-721
function:[OperatorVault.getPendingRewards()](contracts/linkStaking/OperatorVault.sol#L157-L165)is public and can be replaced with external 

contracts/linkStaking/OperatorVault.sol#L157-L165


 - [ ] ID-722
function:[Multicall3.getCurrentBlockCoinbase()](contracts/core/test/Multicall3.sol#L201-L203)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L201-L203


 - [ ] ID-723
function:[RewardsPoolV1.withdraw(uint256)](contracts/core/test/deprecated/RewardsPoolV1.sol#L60-L63)is public and can be replaced with external 

contracts/core/test/deprecated/RewardsPoolV1.sol#L60-L63


 - [ ] ID-724
function:[StakingAllowance.mintToContract(address,address,uint256,bytes)](contracts/core/tokens/StakingAllowance.sol#L30-L38)is public and can be replaced with external 

contracts/core/tokens/StakingAllowance.sol#L30-L38


 - [ ] ID-725
function:[OperatorVCSUpgrade.initialize(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[],address[])](contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L21-L43)is public and can be replaced with external 

contracts/linkStaking/v1.5/OperatorVCSUpgrade.sol#L21-L43


 - [ ] ID-726
function:[Strategy.getMaxDeposits()](contracts/core/base/Strategy.sol#L76)is public and can be replaced with external 

contracts/core/base/Strategy.sol#L76


 - [ ] ID-727
function:[StakingRewardsPool.sharesOf(address)](contracts/core/base/StakingRewardsPool.sol#L58-L60)is public and can be replaced with external 

contracts/core/base/StakingRewardsPool.sol#L58-L60


 - [ ] ID-728
function:[Multicall3.getCurrentBlockGasLimit()](contracts/core/test/Multicall3.sol#L211-L213)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L211-L213


 - [ ] ID-729
function:[CommunityVCS.initialize(address,address,address,address,VaultControllerStrategy.Fee[],uint256,uint128,uint128)](contracts/linkStaking/CommunityVCS.sol#L36-L57)is public and can be replaced with external 

contracts/linkStaking/CommunityVCS.sol#L36-L57


 - [ ] ID-730
function:[OperatorController.withdrawRewards()](contracts/ethStaking/base/OperatorController.sol#L188-L190)is public and can be replaced with external 

contracts/ethStaking/base/OperatorController.sol#L188-L190


 - [ ] ID-731
function:[StakingPool.initialize(address,string,string,StakingPool.Fee[])](contracts/core/StakingPool.sol#L41-L52)is public and can be replaced with external 

contracts/core/StakingPool.sol#L41-L52


 - [ ] ID-732
function:[SDLPool.addToken(address,address)](contracts/core/sdlPool/base/SDLPool.sol#L336-L339)is public and can be replaced with external 

contracts/core/sdlPool/base/SDLPool.sol#L336-L339


 - [ ] ID-733
function:[RewardsPoolControllerV1.distributeTokens(address[])](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L126-L130)is public and can be replaced with external 

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L126-L130


 - [ ] ID-734
function:[VaultV1.getTotalDeposits()](contracts/linkStaking/test/deprecated/VaultV1.sol#L89)is public and can be replaced with external 

contracts/linkStaking/test/deprecated/VaultV1.sol#L89


 - [ ] ID-735
function:[RewardsPoolController.withdrawRewards(address[])](contracts/core/base/RewardsPoolController.sol#L142-L147)is public and can be replaced with external 

contracts/core/base/RewardsPoolController.sol#L142-L147


 - [ ] ID-736
function:[LSDIndexAdapter.getLSDByUnderlying(uint256)](contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L58-L60)is public and can be replaced with external 

contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L58-L60


 - [ ] ID-737
function:[VaultControllerStrategyUpgrade.getVaultDepositLimits()](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L266)is public and can be replaced with external 

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L266


 - [ ] ID-738
function:[Multicall3.getLastBlockHash()](contracts/core/test/Multicall3.sol#L226-L230)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L226-L230


 - [ ] ID-739
function:[RewardsPoolControllerMock.initialize(address)](contracts/core/test/RewardsPoolControllerMock.sol#L23-L26)is public and can be replaced with external 

contracts/core/test/RewardsPoolControllerMock.sol#L23-L26


 - [ ] ID-740
function:[RewardsPoolControllerV1.withdrawRewards(address[])](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L166-L171)is public and can be replaced with external 

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L166-L171


 - [ ] ID-741
function:[OperatorVaultV1.initialize(address,address,address,address)](contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L22-L30)is public and can be replaced with external 

contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L22-L30


 - [ ] ID-742
function:[LSDIndexAdapter.getTotalDeposits()](contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L32-L34)is public and can be replaced with external 

contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L32-L34


 - [ ] ID-743
function:[RewardsPoolController.addToken(address,address)](contracts/core/base/RewardsPoolController.sol#L154-L165)is public and can be replaced with external 

contracts/core/base/RewardsPoolController.sol#L154-L165


 - [ ] ID-744
function:[Strategy.getTotalDeposits()](contracts/core/base/Strategy.sol#L70)is public and can be replaced with external 

contracts/core/base/Strategy.sol#L70


 - [ ] ID-745
function:[RewardsPoolController.__RewardsPoolController_init()](contracts/core/base/RewardsPoolController.sol#L28-L31)is public and can be replaced with external 

contracts/core/base/RewardsPoolController.sol#L28-L31


 - [ ] ID-746
function:[Multicall3.aggregate(Multicall3.Call[])](contracts/core/test/Multicall3.sol#L41-L55)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L41-L55


 - [ ] ID-747
function:[NWLOperatorController.initialize(address,address)](contracts/ethStaking/NWLOperatorController.sol#L36-L38)is public and can be replaced with external 

contracts/ethStaking/NWLOperatorController.sol#L36-L38


 - [ ] ID-748
function:[OperatorVCS.setOperatorRewardPercentage(uint256)](contracts/linkStaking/OperatorVCS.sol#L274-L281)is public and can be replaced with external 

contracts/linkStaking/OperatorVCS.sol#L274-L281


 - [ ] ID-749
function:[RewardsPoolControllerV1.addToken(address,address)](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L178-L189)is public and can be replaced with external 

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L178-L189


 - [ ] ID-750
function:[OperatorVaultV1.getTotalDeposits()](contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L41-L46)is public and can be replaced with external 

contracts/linkStaking/test/deprecated/OperatorVaultV1.sol#L41-L46


 - [ ] ID-751
function:[StakingRewardsPool.__StakingRewardsPool_init(address,string,string)](contracts/core/base/StakingRewardsPool.sol#L20-L29)is public and can be replaced with external 

contracts/core/base/StakingRewardsPool.sol#L20-L29


 - [ ] ID-752
function:[Vault.getTotalDeposits()](contracts/linkStaking/base/Vault.sol#L79-L81)is public and can be replaced with external 

contracts/linkStaking/base/Vault.sol#L79-L81


 - [ ] ID-753
function:[StakingRewardsPool.totalSupply()](contracts/core/base/StakingRewardsPool.sol#L35-L37)is public and can be replaced with external 

contracts/core/base/StakingRewardsPool.sol#L35-L37


 - [ ] ID-754
function:[ERC677Upgradeable.__ERC677_init(string,string,uint256)](contracts/core/tokens/base/ERC677Upgradeable.sol#L9-L16)is public and can be replaced with external 

contracts/core/tokens/base/ERC677Upgradeable.sol#L9-L16


 - [ ] ID-755
function:[RewardsPool.distributeRewards()](contracts/core/RewardsPool.sol#L77-L83)is public and can be replaced with external 

contracts/core/RewardsPool.sol#L77-L83


 - [ ] ID-756
function:[Strategy.canWithdraw()](contracts/core/base/Strategy.sol#L49-L56)is public and can be replaced with external 

contracts/core/base/Strategy.sol#L49-L56


 - [ ] ID-757
function:[LiquidSDIndexPool.initialize(string,string,uint256,uint256,LiquidSDIndexPool.Fee[],uint256)](contracts/liquidSDIndex/LiquidSDIndexPool.sol#L56-L72)is public and can be replaced with external 

contracts/liquidSDIndex/LiquidSDIndexPool.sol#L56-L72


 - [ ] ID-758
function:[Multicall3.getEthBalance(address)](contracts/core/test/Multicall3.sol#L221-L223)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L221-L223


 - [ ] ID-759
function:[Multicall3.aggregate3Value(Multicall3.Call3Value[])](contracts/core/test/Multicall3.sol#L151-L187)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L151-L187


 - [ ] ID-760
function:[FraxLSDIndexAdapter.initialize(address,address)](contracts/liquidSDIndex/adapters/FraxLSDIndexAdapter.sol#L17-L19)is public and can be replaced with external 

contracts/liquidSDIndex/adapters/FraxLSDIndexAdapter.sol#L17-L19


 - [ ] ID-761
function:[Strategy.canDeposit()](contracts/core/base/Strategy.sol#L36-L43)is public and can be replaced with external 

contracts/core/base/Strategy.sol#L36-L43


 - [ ] ID-762
function:[CoinbaseLSDIndexAdapter.initialize(address,address)](contracts/liquidSDIndex/adapters/CoinbaseLSDIndexAdapter.sol#L17-L19)is public and can be replaced with external 

contracts/liquidSDIndex/adapters/CoinbaseLSDIndexAdapter.sol#L17-L19


 - [ ] ID-763
function:[LSDIndexAdapter.getTotalDepositsLSD()](contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L40-L42)is public and can be replaced with external 

contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L40-L42


 - [ ] ID-764
function:[ERC677.transferAndCall(address,uint256,bytes)](contracts/core/tokens/base/ERC677.sol#L17-L27)is public and can be replaced with external 

contracts/core/tokens/base/ERC677.sol#L17-L27


 - [ ] ID-765
function:[VaultControllerStrategy.__VaultControllerStrategy_init(address,address,address,address,VaultControllerStrategy.Fee[],uint256)](contracts/linkStaking/base/VaultControllerStrategy.sol#L56-L78)is public and can be replaced with external 

contracts/linkStaking/base/VaultControllerStrategy.sol#L56-L78


 - [ ] ID-766
function:[OperatorController.__OperatorController_init(address,address)](contracts/ethStaking/base/OperatorController.sol#L74-L80)is public and can be replaced with external 

contracts/ethStaking/base/OperatorController.sol#L74-L80


 - [ ] ID-767
function:[OperatorVault.initialize(address,address,address,address,address,address,address)](contracts/linkStaking/OperatorVault.sol#L50-L72)is public and can be replaced with external 

contracts/linkStaking/OperatorVault.sol#L50-L72


 - [ ] ID-768
function:[StrategyMock.createRewardsPool(address)](contracts/core/test/StrategyMock.sol#L109-L113)is public and can be replaced with external 

contracts/core/test/StrategyMock.sol#L109-L113


 - [ ] ID-769
function:[LPLMigration.onTokenTransfer(address,uint256,bytes)](contracts/core/tokens/LPLMigration.sol#L31-L39)is public and can be replaced with external 

contracts/core/tokens/LPLMigration.sol#L31-L39


 - [ ] ID-770
function:[Vault.__Vault_init(address,address,address,address)](contracts/linkStaking/base/Vault.sol#L36-L48)is public and can be replaced with external 

contracts/linkStaking/base/Vault.sol#L36-L48


 - [ ] ID-771
function:[SDLPoolPrimary.initialize(string,string,address,address)](contracts/core/sdlPool/SDLPoolPrimary.sol#L30-L41)is public and can be replaced with external 

contracts/core/sdlPool/SDLPoolPrimary.sol#L30-L41


 - [ ] ID-772
function:[VaultV1.getPrincipalDeposits()](contracts/linkStaking/test/deprecated/VaultV1.sol#L95-L97)is public and can be replaced with external 

contracts/linkStaking/test/deprecated/VaultV1.sol#L95-L97


 - [ ] ID-773
function:[LSDIndexAdapter.__LiquidSDAdapter_init(address,address)](contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L20-L26)is public and can be replaced with external 

contracts/liquidSDIndex/base/LSDIndexAdapter.sol#L20-L26


 - [ ] ID-774
function:[CommunityVaultV2Mock.initializeV2(uint256)](contracts/linkStaking/test/CommunityVaultV2Mock.sol#L13-L15)is public and can be replaced with external 

contracts/linkStaking/test/CommunityVaultV2Mock.sol#L13-L15


 - [ ] ID-775
function:[SDLPoolMock.setEffectiveBalance(address,uint256)](contracts/core/test/SDLPoolMock.sol#L15-L17)is public and can be replaced with external 

contracts/core/test/SDLPoolMock.sol#L15-L17


 - [ ] ID-776
function:[PriorityPool.initialize(address,address,address,uint128,uint128)](contracts/core/priorityPool/PriorityPool.sol#L93-L110)is public and can be replaced with external 

contracts/core/priorityPool/PriorityPool.sol#L93-L110


 - [ ] ID-777
function:[DelegatorPool.initialize(address,string,string,address[])](contracts/core/test/deprecated/DelegatorPool.sol#L56-L74)is public and can be replaced with external 

contracts/core/test/deprecated/DelegatorPool.sol#L56-L74


 - [ ] ID-778
function:[Multicall3.blockAndAggregate(Multicall3.Call[])](contracts/core/test/Multicall3.sol#L103-L113)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L103-L113


 - [ ] ID-779
function:[VaultControllerStrategyUpgrade.__VaultControllerStrategy_init(address,address,address,address,uint256,VaultControllerStrategyUpgrade.Fee[])](contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L86-L109)is public and can be replaced with external 

contracts/linkStaking/v1.5/VaultControllerStrategyUpgrade.sol#L86-L109


 - [ ] ID-780
function:[CommunityVault.initialize(address,address,address,address)](contracts/linkStaking/CommunityVault.sol#L25-L32)is public and can be replaced with external 

contracts/linkStaking/CommunityVault.sol#L25-L32


 - [ ] ID-781
function:[Multicall3.getCurrentBlockDifficulty()](contracts/core/test/Multicall3.sol#L206-L208)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L206-L208


 - [ ] ID-782
function:[StakingAllowance.mint(address,uint256)](contracts/core/tokens/StakingAllowance.sol#L20-L22)is public and can be replaced with external 

contracts/core/tokens/StakingAllowance.sol#L20-L22


 - [ ] ID-783
function:[Multicall3.getChainId()](contracts/core/test/Multicall3.sol#L239-L241)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L239-L241


 - [ ] ID-784
function:[OperatorControllerMock.initialize(address,address)](contracts/ethStaking/test/OperatorControllerMock.sol#L19-L21)is public and can be replaced with external 

contracts/ethStaking/test/OperatorControllerMock.sol#L19-L21


 - [ ] ID-785
function:[ERC677Upgradeable.transferAndCall(address,uint256,bytes)](contracts/core/tokens/base/ERC677Upgradeable.sol#L18-L28)is public and can be replaced with external 

contracts/core/tokens/base/ERC677Upgradeable.sol#L18-L28


 - [ ] ID-786
function:[CCIPOnRampMock.getPoolBySourceToken(address)](contracts/core/test/chainlink/CCIPOnRampMock.sol#L45-L47)is public and can be replaced with external 

contracts/core/test/chainlink/CCIPOnRampMock.sol#L45-L47


 - [ ] ID-787
function:[OperatorVCS.initialize(address,address,address,address,VaultControllerStrategy.Fee[],uint256,uint256)](contracts/linkStaking/OperatorVCS.sol#L46-L76)is public and can be replaced with external 

contracts/linkStaking/OperatorVCS.sol#L46-L76


 - [ ] ID-788
function:[OperatorVault.setRewardsReceiver(address)](contracts/linkStaking/OperatorVault.sol#L227-L233)is public and can be replaced with external 

contracts/linkStaking/OperatorVault.sol#L227-L233


 - [ ] ID-789
function:[StakingAllowance.burn(uint256)](contracts/core/tokens/StakingAllowance.sol#L44-L46)is public and can be replaced with external 

contracts/core/tokens/StakingAllowance.sol#L44-L46


 - [ ] ID-790
function:[VaultV1.__Vault_init(address,address,address)](contracts/linkStaking/test/deprecated/VaultV1.sol#L52-L62)is public and can be replaced with external 

contracts/linkStaking/test/deprecated/VaultV1.sol#L52-L62


 - [ ] ID-791
function:[Multicall3.getCurrentBlockTimestamp()](contracts/core/test/Multicall3.sol#L216-L218)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L216-L218


 - [ ] ID-792
function:[OperatorController.staked(address)](contracts/ethStaking/base/OperatorController.sol#L88-L90)is public and can be replaced with external 

contracts/ethStaking/base/OperatorController.sol#L88-L90


 - [ ] ID-793
function:[LSDIndexAdapterMock.initialize(address,address,uint256)](contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#L13-L20)is public and can be replaced with external 

contracts/liquidSDIndex/test/LSDIndexAdapterMock.sol#L13-L20


 - [ ] ID-794
function:[StakingRewardsPool.balanceOf(address)](contracts/core/base/StakingRewardsPool.sol#L44-L51)is public and can be replaced with external 

contracts/core/base/StakingRewardsPool.sol#L44-L51


 - [ ] ID-795
function:[SDLPoolSecondary.initialize(string,string,address,address,uint256)](contracts/core/sdlPool/SDLPoolSecondary.sol#L66-L79)is public and can be replaced with external 

contracts/core/sdlPool/SDLPoolSecondary.sol#L66-L79


 - [ ] ID-796
function:[LidoLSDIndexAdapter.initialize(address,address)](contracts/liquidSDIndex/adapters/LidoLSDIndexAdapter.sol#L16-L18)is public and can be replaced with external 

contracts/liquidSDIndex/adapters/LidoLSDIndexAdapter.sol#L16-L18


 - [ ] ID-797
function:[OwnersRewardsPoolV1.withdraw(uint256)](contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L33-L42)is public and can be replaced with external 

contracts/core/test/deprecated/OwnersRewardsPoolV1.sol#L33-L42


 - [ ] ID-798
function:[Multicall3.getBlockNumber()](contracts/core/test/Multicall3.sol#L196-L198)is public and can be replaced with external 

contracts/core/test/Multicall3.sol#L196-L198


 - [ ] ID-799
function:[RewardsPoolControllerV1.__RewardsPoolController_init(string,string)](contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L30-L37)is public and can be replaced with external 

contracts/core/test/deprecated/RewardsPoolControllerV1.sol#L30-L37


 - [ ] ID-800
function:[Strategy.getMinDeposits()](contracts/core/base/Strategy.sol#L82)is public and can be replaced with external 

contracts/core/base/Strategy.sol#L82


 - [ ] ID-801
function:[EthStakingStrategy.initialize(address,address,uint256,uint256,address,bytes32,uint256)](contracts/ethStaking/EthStakingStrategy.sol#L69-L84)is public and can be replaced with external 

contracts/ethStaking/EthStakingStrategy.sol#L69-L84


## unused-event
Impact: Informational
Confidence: High
 - [ ] ID-802
[OperatorVCSWithdrawExtraRewards(address,uint256)](contracts/linkStaking/OperatorVCS.sol#L22) is never emitted in [OperatorVCS](contracts/linkStaking/OperatorVCS.sol#L11-L304)

contracts/linkStaking/OperatorVCS.sol#L22


## version-only
Impact: Informational
Confidence: High
 - [ ] ID-803
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](contracts/core/test/chainlink/CLContractImports0.7.sol#L2)

contracts/core/test/chainlink/CLContractImports0.7.sol#L2


 - [ ] ID-804
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/AuthorizedReceiver.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/AuthorizedReceiver.sol#L2


 - [ ] ID-805
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/interfaces/AuthorizedReceiverInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/interfaces/AuthorizedReceiverInterface.sol#L2


 - [ ] ID-806
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwnerWithProposal.sol#L2


 - [ ] ID-807
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/interfaces/WithdrawalInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/interfaces/WithdrawalInterface.sol#L2


 - [ ] ID-808
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/interfaces/LinkTokenInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/interfaces/LinkTokenInterface.sol#L2


 - [ ] ID-809
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/interfaces/OracleInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/interfaces/OracleInterface.sol#L2


 - [ ] ID-810
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/interfaces/ChainlinkRequestInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/interfaces/ChainlinkRequestInterface.sol#L2


 - [ ] ID-811
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/LinkTokenReceiver.sol#L2


 - [ ] ID-812
	Pragma confirmed better, here is pragma solidity>=0.6.2<0.8.0. [>=0.6.2<0.8.0](node_modules/@chainlink/contracts/src/v0.7/vendor/Address.sol#L4)

node_modules/@chainlink/contracts/src/v0.7/vendor/Address.sol#L4


 - [ ] ID-813
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/interfaces/OperatorInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/interfaces/OperatorInterface.sol#L2


 - [ ] ID-814
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/Operator.sol#L2


 - [ ] ID-815
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/vendor/SafeMathChainlink.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/vendor/SafeMathChainlink.sol#L2


 - [ ] ID-816
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwner.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/ConfirmedOwner.sol#L2


 - [ ] ID-817
	Pragma confirmed better, here is pragma solidity^0.7.0. [^0.7.0](node_modules/@chainlink/contracts/src/v0.7/interfaces/OwnableInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.7/interfaces/OwnableInterface.sol#L2


 - [ ] ID-818
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L4


 - [ ] ID-819
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol#L4


 - [ ] ID-820
	Pragma confirmed better, here is pragma solidity^0.8.1. [^0.8.1](node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#L4


 - [ ] ID-821
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/extensions/draft-IERC20Permit.sol#L4)

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/extensions/draft-IERC20Permit.sol#L4


 - [ ] ID-822
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IAny2EVMMessageReceiver.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IAny2EVMMessageReceiver.sol#L2


 - [ ] ID-823
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#L4


 - [ ] ID-824
	Pragma confirmed better, here is pragma solidity^0.8.15. [^0.8.15](contracts/governance/StakedotlinkCouncil.sol#L3)

contracts/governance/StakedotlinkCouncil.sol#L3


 - [ ] ID-825
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#L4


 - [ ] ID-826
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/interfaces/ENSInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/interfaces/ENSInterface.sol#L2


 - [ ] ID-827
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/cryptography/MerkleProof.sol#L4)

node_modules/@openzeppelin/contracts/utils/cryptography/MerkleProof.sol#L4


 - [ ] ID-828
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/IERC20.sol#L4)

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/IERC20.sol#L4


 - [ ] ID-829
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IRouter.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IRouter.sol#L2


 - [ ] ID-830
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#L4


 - [ ] ID-831
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/Proxy.sol#L4)

node_modules/@openzeppelin/contracts/proxy/Proxy.sol#L4


 - [ ] ID-832
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol#L4


 - [ ] ID-833
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#L4


 - [ ] ID-834
	Pragma confirmed better, here is pragma solidity^0.8.1. [^0.8.1](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/Address.sol#L4)

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/Address.sol#L4


 - [ ] ID-835
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IWrappedNative.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IWrappedNative.sol#L2


 - [ ] ID-836
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/ChainlinkClient.sol#L2


 - [ ] ID-837
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/MerkleProofUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/MerkleProofUpgradeable.sol#L4


 - [ ] ID-838
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/core/test/chainlink/CLContractImports0.8.sol#L2)

contracts/core/test/chainlink/CLContractImports0.8.sol#L2


 - [ ] ID-839
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/interfaces/ChainlinkRequestInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/interfaces/ChainlinkRequestInterface.sol#L2


 - [ ] ID-840
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#L4


 - [ ] ID-841
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/interfaces/IERC1967Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/interfaces/IERC1967Upgradeable.sol#L4


 - [ ] ID-842
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20PermitUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20PermitUpgradeable.sol#L4


 - [ ] ID-843
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/security/Pausable.sol#L4)

node_modules/@openzeppelin/contracts/security/Pausable.sol#L4


 - [ ] ID-844
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/proxy/beacon/IBeaconUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/beacon/IBeaconUpgradeable.sol#L4


 - [ ] ID-845
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L4


 - [ ] ID-846
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/Internal.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/Internal.sol#L2


 - [ ] ID-847
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/interfaces/OwnableInterface.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/interfaces/OwnableInterface.sol#L2


 - [ ] ID-848
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/shared/interfaces/LinkTokenInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/shared/interfaces/LinkTokenInterface.sol#L2


 - [ ] ID-849
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol#L4)

node_modules/@openzeppelin/contracts/security/ReentrancyGuard.sol#L4


 - [ ] ID-850
	Pragma confirmed better, here is pragma solidity>=0.4.19. [>=0.4.19](node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/vendor/CBORChainlink.sol#L2


 - [ ] ID-851
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L4


 - [ ] ID-852
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/introspection/IERC165.sol#L4)

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/introspection/IERC165.sol#L4


 - [ ] ID-853
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#L4


 - [ ] ID-854
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/interfaces/OperatorInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/interfaces/OperatorInterface.sol#L2


 - [ ] ID-855
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol#L4


 - [ ] ID-856
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/pools/IPool.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/pools/IPool.sol#L2


 - [ ] ID-857
	Pragma confirmed better, here is pragma solidity^0.8.1. [^0.8.1](node_modules/@openzeppelin/contracts/utils/Address.sol#L4)

node_modules/@openzeppelin/contracts/utils/Address.sol#L4


 - [ ] ID-858
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/Client.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/Client.sol#L2


 - [ ] ID-859
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/interfaces/TypeAndVersionInterface.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/interfaces/TypeAndVersionInterface.sol#L2


 - [ ] ID-860
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/interfaces/OracleInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/interfaces/OracleInterface.sol#L2


 - [ ] ID-861
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol#L4


 - [ ] ID-862
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#L4


 - [ ] ID-863
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/vendor/BufferChainlink.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/vendor/BufferChainlink.sol#L2


 - [ ] ID-864
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IARM.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IARM.sol#L2


 - [ ] ID-865
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IRouterClient.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IRouterClient.sol#L2


 - [ ] ID-866
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/shared/access/OwnerIsCreator.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/shared/access/OwnerIsCreator.sol#L2


 - [ ] ID-867
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L4


 - [ ] ID-868
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#L4)

node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#L4


 - [ ] ID-869
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/interfaces/draft-IERC1822Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/interfaces/draft-IERC1822Upgradeable.sol#L4


 - [ ] ID-870
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol#L4


 - [ ] ID-871
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/MerkleMultiProof.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/libraries/MerkleMultiProof.sol#L2


 - [ ] ID-872
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IEVM2AnyOnRamp.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/interfaces/IEVM2AnyOnRamp.sol#L2


 - [ ] ID-873
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol#L4


 - [ ] ID-874
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4)

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-875
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/utils/SafeERC20.sol#L4)

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/token/ERC20/utils/SafeERC20.sol#L4


 - [ ] ID-876
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ConfirmedOwner.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ConfirmedOwner.sol#L2


 - [ ] ID-877
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L9)

node_modules/solidity-bytes-utils/contracts/BytesLib.sol#L9


 - [ ] ID-878
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L4)

node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L4


 - [ ] ID-879
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4


 - [ ] ID-880
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L4


 - [ ] ID-881
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#L4


 - [ ] ID-882
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#L4)

node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#L4


 - [ ] ID-883
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/interfaces/PointerInterface.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/interfaces/PointerInterface.sol#L2


 - [ ] ID-884
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/Strings.sol#L4)

node_modules/@openzeppelin/contracts/utils/Strings.sol#L4


 - [ ] ID-885
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#L4


 - [ ] ID-886
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#L4)

node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#L4


 - [ ] ID-887
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol#L4)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol#L4


 - [ ] ID-888
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/StorageSlotUpgradeable.sol#L5)

node_modules/@openzeppelin/contracts-upgradeable/utils/StorageSlotUpgradeable.sol#L5


 - [ ] ID-889
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#L4)

node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#L4


 - [ ] ID-890
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/Chainlink.sol#L2


 - [ ] ID-891
. analyzed (304 contracts with 86 detectors), 908 result(s) found
INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/2023-12-stake-link_scaned_output/mwe-output.json generate success.
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L3)

node_modules/@openzeppelin/contracts/finance/VestingWallet.sol#L3


 - [ ] ID-892
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L4)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L4


 - [ ] ID-893
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#L4)

node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#L4


 - [ ] ID-894
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/vendor/ENSResolver.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/vendor/ENSResolver.sol#L2


 - [ ] ID-895
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/math/Math.sol#L4)

node_modules/@openzeppelin/contracts/utils/math/Math.sol#L4


 - [ ] ID-896
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableSet.sol#L5)

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableSet.sol#L5


 - [ ] ID-897
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L4


 - [ ] ID-898
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L4


 - [ ] ID-899
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L5)

node_modules/@chainlink/contracts-ccip/src/v0.8/vendor/openzeppelin-solidity/v4.8.0/utils/structs/EnumerableMap.sol#L5


 - [ ] ID-900
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ccip/applications/CCIPReceiver.sol#L2


 - [ ] ID-901
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts-ccip/src/v0.8/ConfirmedOwnerWithProposal.sol#L2)

node_modules/@chainlink/contracts-ccip/src/v0.8/ConfirmedOwnerWithProposal.sol#L2


 - [ ] ID-902
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L4


 - [ ] ID-903
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


## state-should-be-constant
Impact: Optimization
Confidence: High
 - [ ] ID-904
[StakingPool.liquidityBuffer](contracts/core/StakingPool.sol#L24) should be constant

contracts/core/StakingPool.sol#L24


 - [ ] ID-905
[DelegatorPool.feeCurve](contracts/core/test/deprecated/DelegatorPool.sol#L36) should be constant

contracts/core/test/deprecated/DelegatorPool.sol#L36


 - [ ] ID-906
[StakingPool.poolIndex](contracts/core/StakingPool.sol#L30) should be constant

contracts/core/StakingPool.sol#L30


## unnecessary-reentrancy-lock
Impact: Optimization
Confidence: High
 - [ ] ID-907
unnecessary reentrancy lock found inOwnersRewardsPoolV1
	- [RewardsPoolV1.updateReward(address)](contracts/core/test/deprecated/RewardsPoolV1.sol#L52-L54)

contracts/core/test/deprecated/RewardsPoolV1.sol#L52-L54


Execution time for Falcon: 1m48.957206762s
