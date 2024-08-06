'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-05-munchables)
'forge config --json' running
'forge build --build-info --skip */test/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-05-munchables)
Top level EventDefinition not supported
Missing function Variable not found: ConfigUpdated() (context MockConfigNotifiable)

Reentrancy in BaseRNGProxy._callback(uint256,bytes) (src/rng/BaseRNGProxy.sol#34-54):
	External calls:
	- (success,returnData) = data.targetContract.call(callData) (src/rng/BaseRNGProxy.sol#45-47)
Reentrancy in LockManager._lock(address,uint256,address,address) (src/managers/LockManager.sol#311-398):
	External calls:
	- token_scope_0.transferFrom(_tokenOwner,address(this),_quantity) (src/managers/LockManager.sol#376)
Reference:  

state variable: BaseConfigStorage.configStorage (src/config/BaseConfigStorage.sol#7) not initialized and not written in contract but be used in contract
state variable: BaseConfigStorage._paused (src/config/BaseConfigStorage.sol#8) not initialized and not written in contract but be used in contract
Reference:  

AccountManager.spraySchnibblesPropose(address[],uint256[]) (src/managers/AccountManager.sol#135-177) deletes IAccountManager.SprayProposal (src/interfaces/IAccountManager.sol#15-18) which contains a array:
	-delete sprayProposals[proposer] (src/managers/AccountManager.sol#161)
AccountManager.execSprayProposal(address) (src/managers/AccountManager.sol#180-210) deletes IAccountManager.SprayProposal (src/interfaces/IAccountManager.sol#15-18) which contains a array:
	-delete sprayProposals[_proposer] (src/managers/AccountManager.sol#209)
AccountManager.removeSprayProposal(address) (src/managers/AccountManager.sol#213-228) deletes IAccountManager.SprayProposal (src/interfaces/IAccountManager.sol#15-18) which contains a array:
	-delete sprayProposals[_proposer] (src/managers/AccountManager.sol#227)
LockManager.proposeUSDPrice(uint256,address[]) (src/managers/LockManager.sol#142-174) deletes ILockManager.USDUpdateProposal (src/interfaces/ILockManager.sol#60-69) which contains a array:
	-delete usdUpdateProposal (src/managers/LockManager.sol#161)
LockManager.disapproveUSDPrice(uint256) (src/managers/LockManager.sol#210-242) deletes ILockManager.USDUpdateProposal (src/interfaces/ILockManager.sol#60-69) which contains a array:
	-delete usdUpdateProposal (src/managers/LockManager.sol#238)
LockManager._execUSDPriceUpdate() (src/managers/LockManager.sol#506-527) deletes ILockManager.USDUpdateProposal (src/interfaces/ILockManager.sol#60-69) which contains a array:
	-delete usdUpdateProposal (src/managers/LockManager.sol#525)
Reference:  

	- FundTreasuryDistributor.receiveTokens(IDistributor.TokenBag[]) (src/distributors/FundTreasuryDistributor.sol#26-50)
	- MigrationManager.lockFundsForAllMigration() (src/managers/MigrationManager.sol#221-244)
	- SnuggeryManager.importMunchable(uint256) (src/managers/SnuggeryManager.sol#79-106)
	- SnuggeryManager.exportMunchable(uint256) (src/managers/SnuggeryManager.sol#109-130)
Reference:  

LockManager.proposeUSDPrice(uint256,address[]) (src/managers/LockManager.sol#142-174) deletes ILockManager.USDUpdateProposal (src/interfaces/ILockManager.sol#60-69) which contains a mapping:
	-delete usdUpdateProposal (src/managers/LockManager.sol#161)
LockManager.disapproveUSDPrice(uint256) (src/managers/LockManager.sol#210-242) deletes ILockManager.USDUpdateProposal (src/interfaces/ILockManager.sol#60-69) which contains a mapping:
	-delete usdUpdateProposal (src/managers/LockManager.sol#238)
LockManager._execUSDPriceUpdate() (src/managers/LockManager.sol#506-527) deletes ILockManager.USDUpdateProposal (src/interfaces/ILockManager.sol#60-69) which contains a mapping:
	-delete usdUpdateProposal (src/managers/LockManager.sol#525)
Reference:  

public mint or burn found in MockMunchNFT.burn(uint256) (src/mock/MockMunchNFT.sol#7)public mint or burn found in MockMunchNFT.mint(address) (src/mock/MockMunchNFT.sol#9)public mint or burn found in TestERC20Token.mint(address,uint256) (src/tokens/TestERC20Token.sol#31-33)Reference: check public mint method

FundTreasuryDistributor.receiveTokens(IDistributor.TokenBag[]) (src/distributors/FundTreasuryDistributor.sol#26-50) use transfer in a loop: address(_treasury).transfer(amount) (src/distributors/FundTreasuryDistributor.sol#38)
Reference:  

NFTOverlord._calculateRaritySpecies(bytes,MunchablesCommonLib.Rarity).selectedSpeciesId (src/overlords/NFTOverlord.sol#376) is a local variable never initialized
MockNFTOverlord.levelUp(uint256,bytes).i (src/mock/MockNFTOverlord.sol#212) is a local variable never initialized
ConfigStorage.getSmallUintArray(StorageKey).i (src/config/ConfigStorage.sol#147) is a local variable never initialized
NFTOverlord.levelUp(uint256,bytes).currentValue (src/overlords/NFTOverlord.sol#260) is a local variable never initialized
RewardsManager.claimYieldForContracts(address[]).ongoingUSDB (src/managers/RewardsManager.sol#116) is a local variable never initialized
registrationDate is a member never initialized in LockManager._lock(address,uint256,address,address)._player (src/managers/LockManager.sol#319)
MunchadexManager.getMunchadexInfo(address).i_scope_0 (src/managers/MunchadexManager.sol#130) is a local variable never initialized
ConfigStorage.setUintArray(StorageKey,uint256[],bool).i (src/config/ConfigStorage.sol#77) is a local variable never initialized
LockManager.getLockedWeightedValue(address).i (src/managers/LockManager.sol#466) is a local variable never initialized
ConfigStorage.setAddressArray(StorageKey,address[],bool).i (src/config/ConfigStorage.sol#262) is a local variable never initialized
ConfigStorage.getUintArray(StorageKey).i (src/config/ConfigStorage.sol#104) is a local variable never initialized
RewardsManager.claimYieldForContracts(address[]).ongoingETH (src/managers/RewardsManager.sol#114) is a local variable never initialized
NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).i (src/managers/NFTAttributeManagerV1.sol#150) is a local variable never initialized
NFTOverlord.levelUp(uint256,bytes).zeroed (src/overlords/NFTOverlord.sol#244) is a local variable never initialized
AccountManager.spraySchnibblesPropose(address[],uint256[]).i (src/managers/AccountManager.sol#163) is a local variable never initialized
SnuggeryManager._findSnuggeryIndex(address,uint256).i (src/managers/SnuggeryManager.sol#310) is a local variable never initialized
maxSnuggerySize is a member never initialized in BonusManager._calculateLevelBonus(address).player (src/managers/BonusManager.sol#197)
RewardsManager.claimYieldForContracts(address[]).ongoingWETH (src/managers/RewardsManager.sol#115) is a local variable never initialized
MockSnuggeryManager.setSnuggeryForTest(address,MunchablesCommonLib.SnuggeryNFT[]).i (src/mock/MockSnuggeryManager.sol#18) is a local variable never initialized
FundTreasuryDistributor.receiveTokens(IDistributor.TokenBag[]).i (src/distributors/FundTreasuryDistributor.sol#31) is a local variable never initialized
AccountManager.execSprayProposal(address).i (src/managers/AccountManager.sol#197) is a local variable never initialized
tokenId is a member never initialized in MigrationManager._migrateNFTs(address,address,uint256[]).snapshot (src/managers/MigrationManager.sol#300)
ConfigStorage.setAddresses(StorageKey[],address[],bool).i (src/config/ConfigStorage.sol#243) is a local variable never initialized
referrer is a member never initialized in ClaimManager._claimPoints(address).player (src/managers/ClaimManager.sol#203)
NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).j (src/managers/NFTAttributeManagerV1.sol#164) is a local variable never initialized
MunchadexManager.getMunchadexInfo(address).i (src/managers/MunchadexManager.sol#124) is a local variable never initialized
BonusManager.getHarvestBonus(address)._migrationBonus (src/managers/BonusManager.sol#131) is a local variable never initialized
ConfigStorage.getAddressArray(StorageKey).i (src/config/ConfigStorage.sol#289) is a local variable never initialized
SnuggeryManager._recalculateChonks(address).i (src/managers/SnuggeryManager.sol#337) is a local variable never initialized
MockNFTOverlord.levelUp(uint256,bytes).currentValue (src/mock/MockNFTOverlord.sol#227) is a local variable never initialized
NFTOverlord.levelUp(uint256,bytes).i (src/overlords/NFTOverlord.sol#245) is a local variable never initialized
PrimordialManager._reconfigure().i (src/managers/PrimordialManager.sol#46) is a local variable never initialized
ConfigStorage.getSmallIntArray(StorageKey).i (src/config/ConfigStorage.sol#205) is a local variable never initialized
SnuggeryManager.getTotalChonk(address).i (src/managers/SnuggeryManager.sol#289) is a local variable never initialized
LockManager._execUSDPriceUpdate().i (src/managers/LockManager.sol#512) is a local variable never initialized
AccountManager.addSubAccount(address).i (src/managers/AccountManager.sol#244) is a local variable never initialized
MockNFTOverlord.levelUp(uint256,bytes).zeroed (src/mock/MockNFTOverlord.sol#211) is a local variable never initialized
MockNFTOverlord._calculateRaritySpecies(bytes,MunchablesCommonLib.Rarity).selectedSpeciesId (src/mock/MockNFTOverlord.sol#338) is a local variable never initialized
ClaimManager.burnNFTsForPoints(address,uint8[]).i (src/managers/ClaimManager.sol#123) is a local variable never initialized
ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool).i (src/config/ConfigStorage.sol#119) is a local variable never initialized
BonusManager._calculateMunchadexBonus(address).i (src/managers/BonusManager.sol#226) is a local variable never initialized
LockManager.setLockDuration(uint256).i (src/managers/LockManager.sol#252) is a local variable never initialized
ConfigStorage.setSmallIntArray(StorageKey,int16[],bool).i (src/config/ConfigStorage.sol#177) is a local variable never initialized
NFTAttributesManagerV1.setGameAttributes(uint256,MunchablesCommonLib.NFTGameAttribute[]).i (src/managers/NFTAttributeManagerV1.sol#92) is a local variable never initialized
NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).i_scope_0 (src/managers/NFTAttributeManagerV1.sol#157) is a local variable never initialized
Reference:  

ERC1967Utils.upgradeToAndCall(address,bytes) (lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#83-92)have ignores return value in Address.functionDelegateCall(newImplementation,data) (lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#88)
ERC1967Utils.upgradeBeaconToAndCall(address,bytes) (lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#173-182)have ignores return value in Address.functionDelegateCall(IBeacon(newBeacon).implementation(),data) (lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#178)
BaseBlastManager.__BaseBlastManager_reconfigure() (src/managers/BaseBlastManager.sol#30-90)have ignores return value in _USDB.configure(YieldMode.CLAIMABLE) (src/managers/BaseBlastManager.sol#73)
BaseBlastManager.claimERC20Yield(address,uint256) (src/managers/BaseBlastManager.sol#109-117)have ignores return value in IERC20Rebasing(_tokenContract).claim(configStorage.getAddress(StorageKey.RewardsManager),_amount) (src/managers/BaseBlastManager.sol#113-116)
MigrationManager.burnUnrevealedForPoints() (src/managers/MigrationManager.sol#135-146)have ignores return value in _claimManager.burnUnrevealedForPoints(msg.sender,amountToSwap) (src/managers/MigrationManager.sol#141)
MigrationManager.burnNFTs(address,uint32) (src/managers/MigrationManager.sol#148-182)have ignores return value in _claimManager.burnNFTsForPoints(_user,tokenIdsByRarity) (src/managers/MigrationManager.sol#179)
MigrationManager.burnRemainingPurchasedNFTs(address,uint32) (src/managers/MigrationManager.sol#184-219)have ignores return value in _claimManager.burnNFTsForPoints(_user,tokenIdsByRarity) (src/managers/MigrationManager.sol#217)
MigrationManager._migrateNFTs(address,address,uint256[]) (src/managers/MigrationManager.sol#290-349)have ignores return value in WETH.approve(address(_lockManager),quantity) (src/managers/MigrationManager.sol#342)
RewardsManager._forwardYield(uint256,uint256,uint256) (src/managers/RewardsManager.sol#219-264)have ignores return value in USDB.approve(address(yieldDistributor),ongoingUSDB) (src/managers/RewardsManager.sol#230)
MockMigrationManager.callMintForMigrationForTest(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[]) (src/mock/MockMigrationManager.sol#18-30)have ignores return value in _nftOverlord.mintForMigration(_player,_attributes,_immutableAttributes,_gameAttributes) (src/mock/MockMigrationManager.sol#24-29)
MockMigrationManager.burnNFTsForPoints(address,uint8[]) (src/mock/MockMigrationManager.sol#32-38)have ignores return value in IClaimManager(configStorage.getAddress(StorageKey.ClaimManager)).burnNFTsForPoints(_player,_rarities) (src/mock/MockMigrationManager.sol#36-37)
MockMigrationManager.burnUnrevealedForPoints(address,uint256) (src/mock/MockMigrationManager.sol#40-46)have ignores return value in IClaimManager(configStorage.getAddress(StorageKey.ClaimManager)).burnUnrevealedForPoints(_player,numUnrevealed) (src/mock/MockMigrationManager.sol#44-45)
MockNFTOverlord._reveal(address,MunchablesCommonLib.Rarity,uint16) (src/mock/MockNFTOverlord.sol#107-127)have ignores return value in IMunchNFT(munchNFT).mint(player) (src/mock/MockNFTOverlord.sol#117)
MockNFTOverlord.revealFromPrimordial(uint256,bytes) (src/mock/MockNFTOverlord.sol#141-162)have ignores return value in IMunchNFT(munchNFT).mint(player) (src/mock/MockNFTOverlord.sol#152)
MockNFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[]) (src/mock/MockNFTOverlord.sol#165-185)have ignores return value in IMunchNFT(munchNFT).mint(_player) (src/mock/MockNFTOverlord.sol#176)
MockRNGProxy.callRevealForTest(uint256,bytes) (src/mock/MockRNGProxy.sol#33-38)have ignores return value in nftOverlord.reveal(_player,_signature) (src/mock/MockRNGProxy.sol#37)
MockRNGProxy.callRevealFromPrimordialForTest(uint256,bytes) (src/mock/MockRNGProxy.sol#40-45)have ignores return value in nftOverlord.revealFromPrimordial(_player,_signature) (src/mock/MockRNGProxy.sol#44)
NFTOverlord.reveal(uint256,bytes) (src/overlords/NFTOverlord.sol#100-132)have ignores return value in IMunchNFT(munchNFT).mint(player) (src/overlords/NFTOverlord.sol#122)
NFTOverlord.revealFromPrimordial(uint256,bytes) (src/overlords/NFTOverlord.sol#154-186)have ignores return value in IMunchNFT(munchNFT).mint(player) (src/overlords/NFTOverlord.sol#176)
NFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[]) (src/overlords/NFTOverlord.sol#189-215)have ignores return value in IMunchNFT(munchNFT).mint(_player) (src/overlords/NFTOverlord.sol#206)
Reference:  

function:ERC165Upgradeable.__ERC165_init() (lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#22-23)is empty 
function:ERC165Upgradeable.__ERC165_init_unchained() (lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#25-26)is empty 
function:ContextUpgradeable.__Context_init() (lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#18-19)is empty 
function:ContextUpgradeable.__Context_init_unchained() (lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#21-22)is empty 
function:AccessControlUpgradeable.__AccessControl_init() (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#82-83)is empty 
function:AccessControlUpgradeable.__AccessControl_init_unchained() (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#85-86)is empty 
function:ERC721EnumerableUpgradeable.__ERC721Enumerable_init() (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#49-50)is empty 
function:ERC721EnumerableUpgradeable.__ERC721Enumerable_init_unchained() (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#52-53)is empty 
function:ERC721PausableUpgradeable.__ERC721Pausable_init_unchained() (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721PausableUpgradeable.sol#28-29)is empty 
function:ERC721URIStorageUpgradeable.__ERC721URIStorage_init() (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#37-38)is empty 
function:ERC721URIStorageUpgradeable.__ERC721URIStorage_init_unchained() (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#40-41)is empty 
function:BaseConfigStorageUpgradeable._authorizeUpgrade(address) (src/config/BaseConfigStorageUpgradeable.sol#14)is empty 
function:MockBlast.configureContract(address,YieldMode,GasMode,address) (src/mock/MockBlast.sol#11)is empty 
function:MockBlast.configure(YieldMode,GasMode,address) (src/mock/MockBlast.sol#12)is empty 
function:MockMunchNFT.burn(uint256) (src/mock/MockMunchNFT.sol#7)is empty 
function:MockMunchNFT.mint(address) (src/mock/MockMunchNFT.sol#9)is empty 
function:BaseRNGProxy.configUpdated() (src/rng/BaseRNGProxy.sol#19)is empty 
function:OldMunchNFT.configureContract(address) (src/tokens/OldMunchNFT.sol#93-98)is empty 
function:OldMunchNFT.revealNFT(address,string,bytes32,bytes,MunchablesCommonLib.Realm,MunchablesCommonLib.Rarity,uint16) (src/tokens/OldMunchNFT.sol#110-140)is empty 
function:OldMunchNFT.setURI(uint256,string) (src/tokens/OldMunchNFT.sol#145-160)is empty 
Reference:  

	- LockManager.configureLockdrop(ILockManager.Lockdrop) (src/managers/LockManager.sol#98-112)
	- LockManager.configureToken(address,ILockManager.ConfiguredToken) (src/managers/LockManager.sol#115-127)
	- LockManager.setLockDuration(uint256) (src/managers/LockManager.sol#245-272)
	- MigrationManager.loadMigrationSnapshot(address[],IMigrationManager.MigrationSnapshotData[]) (src/managers/MigrationManager.sol#78-114)
	- MigrationManager.sealData() (src/managers/MigrationManager.sol#116-120)
	- MigrationManager.burnNFTs(address,uint32) (src/managers/MigrationManager.sol#148-182)
Reference:  

Setter function AccessControl._checkRole(bytes32) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#86-88) does not emit an event
Setter function AccessControl.grantRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#122-124) does not emit an event
Setter function AccessControl.revokeRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#137-139) does not emit an event
Setter function AccessControl.renounceRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#157-163) does not emit an event
Setter function AccessControl.onlyRole(bytes32) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#63-66) does not emit an event
Setter function ProxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy,address,bytes) (lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#38-44) does not emit an event
Setter function ConfigStorage.constructor() (src/config/ConfigStorage.sol#25) does not emit an event
Setter function ConfigStorage.setRole(Role,address,address) (src/config/ConfigStorage.sol#27-33) does not emit an event
Setter function ConfigStorage.setUniversalRole(Role,address) (src/config/ConfigStorage.sol#35-37) does not emit an event
Setter function ConfigStorage.getRole(Role) (src/config/ConfigStorage.sol#39-41) does not emit an event
Setter function ConfigStorage.setUint(StorageKey,uint256,bool) (src/config/ConfigStorage.sol#55-62) does not emit an event
Setter function ConfigStorage.setUintArray(StorageKey,uint256[],bool) (src/config/ConfigStorage.sol#69-96) does not emit an event
Setter function ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool) (src/config/ConfigStorage.sol#111-139) does not emit an event
Setter function ConfigStorage.setSmallInt(StorageKey,int16,bool) (src/config/ConfigStorage.sol#155-162) does not emit an event
Setter function ConfigStorage.setSmallIntArray(StorageKey,int16[],bool) (src/config/ConfigStorage.sol#169-197) does not emit an event
Setter function ConfigStorage.setBool(StorageKey,bool,bool) (src/config/ConfigStorage.sol#213-220) does not emit an event
Setter function ConfigStorage.setAddress(StorageKey,address,bool) (src/config/ConfigStorage.sol#228-235) does not emit an event
Setter function ConfigStorage.setAddresses(StorageKey[],address[],bool) (src/config/ConfigStorage.sol#238-247) does not emit an event
Setter function ConfigStorage.setAddressArray(StorageKey,address[],bool) (src/config/ConfigStorage.sol#254-281) does not emit an event
Setter function ConfigStorage.setBytes32(StorageKey,bytes32,bool) (src/config/ConfigStorage.sol#297-304) does not emit an event
Setter function ConfigStorage.addNotifiableAddress(address) (src/config/ConfigStorage.sol#312-314) does not emit an event
Setter function ConfigStorage.addNotifiableAddresses(address[]) (src/config/ConfigStorage.sol#316-322) does not emit an event
Setter function ConfigStorage.removeNotifiableAddress(address) (src/config/ConfigStorage.sol#325-335) does not emit an event
Setter function ConfigStorage.manualNotify(uint8,uint8) (src/config/ConfigStorage.sol#345-350) does not emit an event
Setter function ConfigStorage.manualNotifyAddress(address) (src/config/ConfigStorage.sol#352-354) does not emit an event
Setter function FundTreasuryDistributor.configUpdated() (src/distributors/FundTreasuryDistributor.sol#17-19) does not emit an event
Setter function FundTreasuryDistributor._reconfigure() (src/distributors/FundTreasuryDistributor.sol#21-24) does not emit an event
Setter function BonusManager._reconfigure() (src/managers/BonusManager.sol#40-85) does not emit an event
Setter function BonusManager.configUpdated() (src/managers/BonusManager.sol#87-89) does not emit an event
Setter function MockAccountManager.giveSchnibbles(address,uint256) (src/mock/MockAccountManager.sol#12-17) does not emit an event
Setter function MockBlast.configureClaimableYield() (src/mock/MockBlast.sol#15-17) does not emit an event
Setter function MockBlast.configureClaimableYieldOnBehalf(address) (src/mock/MockBlast.sol#18-20) does not emit an event
Setter function MockBlast.configureAutomaticYield() (src/mock/MockBlast.sol#21-23) does not emit an event
Setter function MockBlast.configureAutomaticYieldOnBehalf(address) (src/mock/MockBlast.sol#24-26) does not emit an event
Setter function MockBlast.configureVoidYield() (src/mock/MockBlast.sol#27-29) does not emit an event
Setter function MockBlast.configureVoidYieldOnBehalf(address) (src/mock/MockBlast.sol#30-32) does not emit an event
Setter function MockBlast.configureClaimableGas() (src/mock/MockBlast.sol#33-35) does not emit an event
Setter function MockBlast.configureClaimableGasOnBehalf(address) (src/mock/MockBlast.sol#36-38) does not emit an event
Setter function MockBlast.configureVoidGas() (src/mock/MockBlast.sol#39-41) does not emit an event
Setter function MockBlast.configureVoidGasOnBehalf(address) (src/mock/MockBlast.sol#42-44) does not emit an event
Setter function MockBlast.configureGovernor(address) (src/mock/MockBlast.sol#45-47) does not emit an event
Setter function MockBlast.configureGovernorOnBehalf(address,address) (src/mock/MockBlast.sol#48-50) does not emit an event
Setter function MockBlast.claimYield(address,address,uint256) (src/mock/MockBlast.sol#53-56) does not emit an event
Setter function MockBlast.claimAllYield(address,address) (src/mock/MockBlast.sol#58-61) does not emit an event
Setter function MockBlast.claimAllGas(address,address) (src/mock/MockBlast.sol#64-67) does not emit an event
Setter function MockBlast.claimGasAtMinClaimRate(address,address,uint256) (src/mock/MockBlast.sol#68-75) does not emit an event
Setter function MockBlast.claimMaxGas(address,address) (src/mock/MockBlast.sol#76-79) does not emit an event
Setter function MockBlast.claimGas(address,address,uint256,uint256) (src/mock/MockBlast.sol#80-88) does not emit an event
Setter function MockClaimManager.givePoints(address,uint256) (src/mock/MockClaimManager.sol#12-14) does not emit an event
Setter function MockConfigNotifiable.configUpdated() (src/mock/MockConfigNotifiable.sol#17-20) does not emit an event
Setter function MockLockManager.setLockedTokenForTest(address,address,ILockManager.LockedToken) (src/mock/MockLockManager.sol#9-15) does not emit an event
Setter function MockMigrationManager.setUserMigrationCompletedDataForTest(address,IMigrationManager.MigrationTotals,bool) (src/mock/MockMigrationManager.sol#9-16) does not emit an event
Setter function MockMunchadexManager.setMunchadexForTest(address,uint256) (src/mock/MockMunchadexManager.sol#9-15) does not emit an event
Setter function MockMunchadexManager.setMunchadexNumInRealmForTest(address,MunchablesCommonLib.Realm,uint256) (src/mock/MockMunchadexManager.sol#17-24) does not emit an event
Setter function MockMunchadexManager.setMunchadexNumInRarityForTest(address,MunchablesCommonLib.Rarity,uint256) (src/mock/MockMunchadexManager.sol#26-33) does not emit an event
Setter function MockMunchadexManager.setMunchadexUniqueForTest(address,bytes32,uint256) (src/mock/MockMunchadexManager.sol#35-42) does not emit an event
Setter function MockNFTAttributesManagerV1.setImmutableAttributesForTest(uint256,MunchablesCommonLib.NFTImmutableAttributes) (src/mock/MockNFTAttributeManagerV1.sol#12-17) does not emit an event
Setter function MockNFTAttributesManagerV1.setAttributesForTest(uint256,MunchablesCommonLib.NFTAttributes) (src/mock/MockNFTAttributeManagerV1.sol#19-24) does not emit an event
Setter function MockNFTOverlord._reconfigure() (src/mock/MockNFTOverlord.sol#33-57) does not emit an event
Setter function MockNFTOverlord.configUpdated() (src/mock/MockNFTOverlord.sol#59-61) does not emit an event
Setter function MockNFTOverlord.addReveal(address,uint16) (src/mock/MockNFTOverlord.sol#64-67) does not emit an event
Setter function MockNFTOverlord.startReveal() (src/mock/MockNFTOverlord.sol#70-72) does not emit an event
Setter function MockNFTOverlord.mintFromPrimordial(address) (src/mock/MockNFTOverlord.sol#129-137) does not emit an event
Setter function MockNFTOverlord._populateDefaultProbabilities() (src/mock/MockNFTOverlord.sol#384-415) does not emit an event
Setter function MockNFTOverlord._populateDefaultRealmLookup() (src/mock/MockNFTOverlord.sol#417-425) does not emit an event
Setter function MockRNGProxy._reconfigure() (src/mock/MockRNGProxy.sol#15-19) does not emit an event
Setter function MockSnuggeryManager.setSnuggeryForTest(address,MunchablesCommonLib.SnuggeryNFT[]) (src/mock/MockSnuggeryManager.sol#11-22) does not emit an event
Setter function MockSnuggeryManager.setGlobalTotalChonk(uint256) (src/mock/MockSnuggeryManager.sol#24-26) does not emit an event
Setter function NFTOverlord._reconfigure() (src/overlords/NFTOverlord.sol#34-58) does not emit an event
Setter function NFTOverlord.configUpdated() (src/overlords/NFTOverlord.sol#60-62) does not emit an event
Setter function NFTOverlord.addReveal(address,uint16) (src/overlords/NFTOverlord.sol#65-71) does not emit an event
Setter function NFTOverlord.mintFromPrimordial(address) (src/overlords/NFTOverlord.sol#134-150) does not emit an event
Setter function NFTOverlord._populateDefaultProbabilities() (src/overlords/NFTOverlord.sol#422-453) does not emit an event
Setter function NFTOverlord._populateDefaultRealmLookup() (src/overlords/NFTOverlord.sol#455-463) does not emit an event
Setter function RNGProxyAPI3.requestRandom(address,bytes4,uint256) (src/rng/RNGProxyAPI3.sol#30-47) does not emit an event
Setter function RNGProxyAPI3.provideRandom(bytes32,bytes) (src/rng/RNGProxyAPI3.sol#51-58) does not emit an event
Setter function MunchNFT.transferFrom(address,address,uint256) (src/tokens/MunchNFT.sol#39-46) does not emit an event
Setter function MunchNFT._reconfigure() (src/tokens/MunchNFT.sol#48-60) does not emit an event
Setter function MunchNFT.configUpdated() (src/tokens/MunchNFT.sol#62-64) does not emit an event
Setter function MunchNFT.mint(address) (src/tokens/MunchNFT.sol#66-77) does not emit an event
Setter function MunchNFT.setTokenURI(uint256,string) (src/tokens/MunchNFT.sol#79-84) does not emit an event
Setter function MunchToken.configUpdated() (src/tokens/MunchToken.sol#24-26) does not emit an event
Setter function MunchToken.mint(address,uint256) (src/tokens/MunchToken.sol#28-33) does not emit an event
Setter function OldMunchNFT.initialize(address,address,address,address,address) (src/tokens/OldMunchNFT.sol#53-75) does not emit an event
Setter function OldMunchNFT.pause() (src/tokens/OldMunchNFT.sol#78-80) does not emit an event
Setter function OldMunchNFT.unpause() (src/tokens/OldMunchNFT.sol#83-85) does not emit an event
Setter function OldMunchNFT.burn(uint256) (src/tokens/OldMunchNFT.sol#87-89) does not emit an event
Setter function OldMunchNFT.configureContract(address) (src/tokens/OldMunchNFT.sol#93-98) does not emit an event
Setter function OldMunchNFT.setMigrationManager(address) (src/tokens/OldMunchNFT.sol#100-104) does not emit an event
Setter function OldMunchNFT.revealNFT(address,string,bytes32,bytes,MunchablesCommonLib.Realm,MunchablesCommonLib.Rarity,uint16) (src/tokens/OldMunchNFT.sol#110-140) does not emit an event
Setter function OldMunchNFT.setURI(uint256,string) (src/tokens/OldMunchNFT.sol#145-160) does not emit an event
Setter function OldMunchNFT.onlyMigrationManager() (src/tokens/OldMunchNFT.sol#39-42) does not emit an event
Setter function TestERC20Token.constructor() (src/tokens/TestERC20Token.sol#11-14) does not emit an event
Setter function TestERC20Token.configure(YieldMode) (src/tokens/TestERC20Token.sol#17-20) does not emit an event
Setter function TestERC20Token.claim(address,uint256) (src/tokens/TestERC20Token.sol#22-25) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

variable lacks a zero-check on 		- ProxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy,address,bytes) (lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#38-44)
variable lacks a zero-check on 		- BaseBlastManager.claimERC20Yield(address,uint256) (src/managers/BaseBlastManager.sol#109-117)
variable lacks a zero-check on 		- BonusManager.getFeedBonus(address,uint256) (src/managers/BonusManager.sol#91-123)
variable lacks a zero-check on 		- BonusManager.getHarvestBonus(address) (src/managers/BonusManager.sol#125-140)
variable lacks a zero-check on 		- BonusManager.getPetBonus(address) (src/managers/BonusManager.sol#142-148)
variable lacks a zero-check on 		- AccountManager.removeSprayProposal(address) (src/managers/AccountManager.sol#213-228)
variable lacks a zero-check on 		- AccountManager.updatePlayer(address,MunchablesCommonLib.Player) (src/managers/AccountManager.sol#284-295)
variable lacks a zero-check on 		- AccountManager.getDailySchnibbles(address) (src/managers/AccountManager.sol#363-370)
variable lacks a zero-check on 		- MockAccountManager.giveSchnibbles(address,uint256) (src/mock/MockAccountManager.sol#12-17)
variable lacks a zero-check on 		- ClaimManager.burnNFTsForPoints(address,uint8[]) (src/managers/ClaimManager.sol#113-129)
variable lacks a zero-check on 		- ClaimManager.burnUnrevealedForPoints(address,uint256) (src/managers/ClaimManager.sol#131-144)
variable lacks a zero-check on 		- ClaimManager.getPoints(address) (src/managers/ClaimManager.sol#196-200)
variable lacks a zero-check on 		- MockClaimManager.givePoints(address,uint256) (src/mock/MockClaimManager.sol#12-14)
variable lacks a zero-check on 		- MockLockManager.setLockedTokenForTest(address,address,ILockManager.LockedToken) (src/mock/MockLockManager.sol#9-15)
variable lacks a zero-check on 		- MockLockManager.callAddRevealForTest(address,uint8) (src/mock/MockLockManager.sol#17-19)
variable lacks a zero-check on 		- MigrationManager.loadMigrationSnapshot(address[],IMigrationManager.MigrationSnapshotData[]) (src/managers/MigrationManager.sol#78-114)
variable lacks a zero-check on 		- MockMigrationManager.setUserMigrationCompletedDataForTest(address,IMigrationManager.MigrationTotals,bool) (src/mock/MockMigrationManager.sol#9-16)
variable lacks a zero-check on 		- MockMigrationManager.callMintForMigrationForTest(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[]) (src/mock/MockMigrationManager.sol#18-30)
variable lacks a zero-check on 		- MockMigrationManager.burnNFTsForPoints(address,uint8[]) (src/mock/MockMigrationManager.sol#32-38)
variable lacks a zero-check on 		- MockMigrationManager.burnUnrevealedForPoints(address,uint256) (src/mock/MockMigrationManager.sol#40-46)
variable lacks a zero-check on 		- MockNFTOverlord.addReveal(address,uint16) (src/mock/MockNFTOverlord.sol#64-67)
variable lacks a zero-check on 		- MockNFTOverlord.mintFromPrimordial(address) (src/mock/MockNFTOverlord.sol#129-137)
variable lacks a zero-check on 		- MockNFTOverlord.revealFromPrimordial(uint256,bytes) (src/mock/MockNFTOverlord.sol#141-162)
variable lacks a zero-check on 		- MockNFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[]) (src/mock/MockNFTOverlord.sol#165-185)
variable lacks a zero-check on 		- MockNFTOverlord.getUnrevealedNFTs(address) (src/mock/MockNFTOverlord.sol#287-292)
variable lacks a zero-check on 		- MockPrimordialManager.callMintFromPrimordialForTest(address) (src/mock/MockPrimordialManager.sol#9-11)
variable lacks a zero-check on 		- RewardsManager.reassignBlastGovernor(address) (src/managers/RewardsManager.sol#159-167)
variable lacks a zero-check on 		- RewardsManager.isGovernorOfContract(address) (src/managers/RewardsManager.sol#169-173)
variable lacks a zero-check on 		- MockRewardsManager.claimAllGasForContract(address) (src/mock/MockRewardsManager.sol#10-19)
variable lacks a zero-check on 		- SnuggeryManager.getSnuggery(address,uint256) (src/managers/SnuggeryManager.sol#254-282)
variable lacks a zero-check on 		- MockSnuggeryManager.spendPoints(address,uint256) (src/mock/MockSnuggeryManager.sol#28-30)
variable lacks a zero-check on 		- MockSnuggeryManager.forceClaimPoints(address) (src/mock/MockSnuggeryManager.sol#32-34)
variable lacks a zero-check on 		- NFTOverlord.addReveal(address,uint16) (src/overlords/NFTOverlord.sol#65-71)
variable lacks a zero-check on 		- NFTOverlord.mintFromPrimordial(address) (src/overlords/NFTOverlord.sol#134-150)
variable lacks a zero-check on 		- NFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[]) (src/overlords/NFTOverlord.sol#189-215)
variable lacks a zero-check on 		- NFTOverlord.getUnrevealedNFTs(address) (src/overlords/NFTOverlord.sol#316-321)
variable lacks a zero-check on 		- ERC1967Proxy.constructor(address,bytes) (lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol#26-28)
variable lacks a zero-check on 		- RrpRequesterV0.constructor(address) (node_modules/@api3/airnode-protocol/contracts/rrp/requesters/RrpRequesterV0.sol#23-26)
variable lacks a zero-check on 		- RNGProxyAPI3.constructor(address,address,address,address,address,bytes32) (src/rng/RNGProxyAPI3.sol#16-28)
variable lacks a zero-check on 		- MunchNFT.transferFrom(address,address,uint256) (src/tokens/MunchNFT.sol#39-46)
variable lacks a zero-check on 		- MunchNFT.mint(address) (src/tokens/MunchNFT.sol#66-77)
variable lacks a zero-check on 		- OldMunchNFT.setMigrationManager(address) (src/tokens/OldMunchNFT.sol#100-104)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

	- ConfigStorage.setRole(Role,address,address) (src/config/ConfigStorage.sol#27-33)
	- ConfigStorage.setUniversalRole(Role,address) (src/config/ConfigStorage.sol#35-37)
	- ConfigStorage.setUint(StorageKey,uint256,bool) (src/config/ConfigStorage.sol#55-62)
	- ConfigStorage.setUintArray(StorageKey,uint256[],bool) (src/config/ConfigStorage.sol#69-96)
	- ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool) (src/config/ConfigStorage.sol#111-139)
	- ConfigStorage.setSmallInt(StorageKey,int16,bool) (src/config/ConfigStorage.sol#155-162)
	- ConfigStorage.setSmallIntArray(StorageKey,int16[],bool) (src/config/ConfigStorage.sol#169-197)
	- ConfigStorage.setBool(StorageKey,bool,bool) (src/config/ConfigStorage.sol#213-220)
	- ConfigStorage.setAddress(StorageKey,address,bool) (src/config/ConfigStorage.sol#228-235)
	- ConfigStorage.setAddresses(StorageKey[],address[],bool) (src/config/ConfigStorage.sol#238-247)
	- ConfigStorage.setAddressArray(StorageKey,address[],bool) (src/config/ConfigStorage.sol#254-281)
	- ConfigStorage.setBytes32(StorageKey,bytes32,bool) (src/config/ConfigStorage.sol#297-304)
	- ConfigStorage.addNotifiableAddress(address) (src/config/ConfigStorage.sol#312-314)
	- ConfigStorage.addNotifiableAddresses(address[]) (src/config/ConfigStorage.sol#316-322)
	- ConfigStorage.removeNotifiableAddress(address) (src/config/ConfigStorage.sol#325-335)
	- AccountManager.register(MunchablesCommonLib.Realm,address) (src/managers/AccountManager.sol#87-113)
	- AccountManager.spraySchnibblesPropose(address[],uint256[]) (src/managers/AccountManager.sol#135-177)
	- AccountManager.execSprayProposal(address) (src/managers/AccountManager.sol#180-210)
	- AccountManager.removeSprayProposal(address) (src/managers/AccountManager.sol#213-228)
	- AccountManager.addSubAccount(address) (src/managers/AccountManager.sol#231-251)
	- AccountManager.updatePlayer(address,MunchablesCommonLib.Player) (src/managers/AccountManager.sol#284-295)
	- ClaimManager.newPeriod() (src/managers/ClaimManager.sol#86-111)
	- ClaimManager.burnNFTsForPoints(address,uint8[]) (src/managers/ClaimManager.sol#113-129)
	- ClaimManager.burnUnrevealedForPoints(address,uint256) (src/managers/ClaimManager.sol#131-144)
	- ClaimManager.spendPoints(address,uint256) (src/managers/ClaimManager.sol#163-171)
	- ClaimManager.convertPointsToTokens(uint256) (src/managers/ClaimManager.sol#173-190)
	- LockManager.setUSDThresholds(uint8,uint8) (src/managers/LockManager.sol#129-139)
	- LockManager.proposeUSDPrice(uint256,address[]) (src/managers/LockManager.sol#142-174)
	- LockManager.approveUSDPrice(uint256) (src/managers/LockManager.sol#177-207)
	- LockManager.disapproveUSDPrice(uint256) (src/managers/LockManager.sol#210-242)
	- MigrationManager.loadUnrevealedSnapshot(address[],uint16[]) (src/managers/MigrationManager.sol#122-133)
	- MigrationManager.burnUnrevealedForPoints() (src/managers/MigrationManager.sol#135-146)
	- MigrationManager.burnRemainingPurchasedNFTs(address,uint32) (src/managers/MigrationManager.sol#184-219)
	- MunchadexManager.updateMunchadex(address,address,uint256) (src/managers/MunchadexManager.sol#41-100)
	- NFTAttributesManagerV1.setAttributes(uint256,MunchablesCommonLib.NFTAttributes) (src/managers/NFTAttributeManagerV1.sol#52-69)
	- NFTAttributesManagerV1.createWithImmutable(uint256,MunchablesCommonLib.NFTImmutableAttributes) (src/managers/NFTAttributeManagerV1.sol#71-82)
	- NFTAttributesManagerV1.setGameAttributes(uint256,MunchablesCommonLib.NFTGameAttribute[]) (src/managers/NFTAttributeManagerV1.sol#84-104)
	- PrimordialManager.claimPrimordial() (src/managers/PrimordialManager.sol#60-70)
	- PrimordialManager.feedPrimordial(uint256) (src/managers/PrimordialManager.sol#73-126)
	- BaseRNGProxy.requestRandom(address,bytes4,uint256) (src/rng/BaseRNGProxy.sol#21-32)
	- NFTOverlord.addReveal(address,uint16) (src/overlords/NFTOverlord.sol#65-71)
	- NFTOverlord.reveal(uint256,bytes) (src/overlords/NFTOverlord.sol#100-132)
	- NFTOverlord.mintFromPrimordial(address) (src/overlords/NFTOverlord.sol#134-150)
	- NFTOverlord.revealFromPrimordial(uint256,bytes) (src/overlords/NFTOverlord.sol#154-186)
	- NFTOverlord.levelUp(uint256,bytes) (src/overlords/NFTOverlord.sol#218-276)
	- NFTOverlord.munchableFed(uint256,address) (src/overlords/NFTOverlord.sol#279-313)
	- RNGProxyAPI3.requestRandom(address,bytes4,uint256) (src/rng/RNGProxyAPI3.sol#30-47)
	- MunchNFT.mint(address) (src/tokens/MunchNFT.sol#66-77)
	- OldMunchNFT.setMigrationManager(address) (src/tokens/OldMunchNFT.sol#100-104)
	- OldMunchNFT.safeMint(address,string) (src/tokens/OldMunchNFT.sol#184-193)
Reference:  

	- AccessControl.grantRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#122-124)
	- AccessControl.revokeRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#137-139)
	- Ownable.renounceOwnership() (lib/openzeppelin-contracts/contracts/access/Ownable.sol#76-78)
	- Ownable.transferOwnership(address) (lib/openzeppelin-contracts/contracts/access/Ownable.sol#84-89)
	- ProxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy,address,bytes) (lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#38-44)
	- UUPSUpgradeable.upgradeToAndCall(address,bytes) (lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#86-89)
	- AccessControlUpgradeable.grantRole(bytes32,address) (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#143-145)
	- AccessControlUpgradeable.revokeRole(bytes32,address) (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#158-160)
	- ConfigStorage.manualNotify(uint8,uint8) (src/config/ConfigStorage.sol#345-350)
	- ConfigStorage.manualNotifyAddress(address) (src/config/ConfigStorage.sol#352-354)
	- FundTreasuryDistributor.configUpdated() (src/distributors/FundTreasuryDistributor.sol#17-19)
	- BaseBlastManager.claimERC20Yield(address,uint256) (src/managers/BaseBlastManager.sol#109-117)
	- AccountManager.configUpdated() (src/managers/AccountManager.sol#82-84)
	- AccountManager.forceHarvest(address) (src/managers/AccountManager.sol#122-132)
	- AccountManager.removeSubAccount(address) (src/managers/AccountManager.sol#254-258)
	- BonusManager.configUpdated() (src/managers/BonusManager.sol#87-89)
	- ClaimManager.configUpdated() (src/managers/ClaimManager.sol#82-84)
	- ClaimManager.claimPoints() (src/managers/ClaimManager.sol#146-150)
	- ClaimManager.forceClaimPoints(address) (src/managers/ClaimManager.sol#153-161)
	- LockManager.configUpdated() (src/managers/LockManager.sol#85-87)
	- LockManager.lockOnBehalf(address,uint256,address) (src/managers/LockManager.sol#275-294)
	- LockManager.lock(address,uint256) (src/managers/LockManager.sol#297-309)
	- MigrationManager.configUpdated() (src/managers/MigrationManager.sol#74-76)
	- MigrationManager.migratePurchasedNFTs(uint256[]) (src/managers/MigrationManager.sol#267-288)
	- MunchadexManager.configUpdated() (src/managers/MunchadexManager.sol#37-39)
	- NFTAttributesManagerV1.configUpdated() (src/managers/NFTAttributeManagerV1.sol#48-50)
	- PrimordialManager.configUpdated() (src/managers/PrimordialManager.sol#55-57)
	- RewardsManager.configUpdated() (src/managers/RewardsManager.sol#54-56)
	- SnuggeryManager.configUpdated() (src/managers/SnuggeryManager.sol#74-76)
	- SnuggeryManager.feed(uint256,uint256) (src/managers/SnuggeryManager.sol#133-171)
	- MockNFTOverlord.configUpdated() (src/mock/MockNFTOverlord.sol#59-61)
	- MockNFTOverlord.reveal(uint256,bytes) (src/mock/MockNFTOverlord.sol#87-97)
	- RNGProxySelfHosted.provideRandom(uint256,bytes) (src/rng/RNGProxySelfHosted.sol#12-17)
	- NFTOverlord.configUpdated() (src/overlords/NFTOverlord.sol#60-62)
	- NFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[]) (src/overlords/NFTOverlord.sol#189-215)
	- RNGProxyAPI3.provideRandom(bytes32,bytes) (src/rng/RNGProxyAPI3.sol#51-58)
	- MunchNFT.configUpdated() (src/tokens/MunchNFT.sol#62-64)
	- MunchNFT.setTokenURI(uint256,string) (src/tokens/MunchNFT.sol#79-84)
	- MunchToken.configUpdated() (src/tokens/MunchToken.sol#24-26)
	- MunchToken.mint(address,uint256) (src/tokens/MunchToken.sol#28-33)
	- OldMunchNFT.pause() (src/tokens/OldMunchNFT.sol#78-80)
	- OldMunchNFT.unpause() (src/tokens/OldMunchNFT.sol#83-85)
	- OldMunchNFT.burn(uint256) (src/tokens/OldMunchNFT.sol#87-89)
	- OldMunchNFT.configureContract(address) (src/tokens/OldMunchNFT.sol#93-98)
Reference:  

BaseConfigStorageUpgradeable.__BaseConfigStorageUpgradable_reconfigure() (src/config/BaseConfigStorageUpgradeable.sol#20-22) is never used and should be removed
MockNFTOverlord._calculateRaritySpecies(bytes,MunchablesCommonLib.Rarity) (src/mock/MockNFTOverlord.sol#329-382) is never used and should be removed
MockNFTOverlord._getMainAccountRequireRegistered(address) (src/mock/MockNFTOverlord.sol#427-440) is never used and should be removed
Reference:  

require() missing error messages
	 - require(bool)(msg.data.length == 0) (src/managers/RewardsManager.sol#40)
Reference: https://dev.to/tawseef/require-vs-assert-in-solidity-5e9d

Potential price manipulation risk:
	- In function _addTokenToOwnerEnumeration
		-- length = balanceOf(to) - 1 (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#96) have potential price manipulated risk from length and call None which could influence variable:length
	- In function _removeTokenFromOwnerEnumeration
		-- lastTokenIndex = balanceOf(from) (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#122) have potential price manipulated risk from lastTokenIndex and call None which could influence variable:lastTokenIndex
Potential price manipulation risk:
	- In function _addTokenToOwnerEnumeration
		-- length = balanceOf(to) - 1 (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#118) have potential price manipulated risk from length and call None which could influence variable:length
	- In function _removeTokenFromOwnerEnumeration
		-- lastTokenIndex = balanceOf(from) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#146) have potential price manipulated risk from lastTokenIndex and call None which could influence variable:lastTokenIndex
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

Potential DoS Gas Limit Attack occur inConfigStorage.notify() (src/config/ConfigStorage.sol#357-361)BEGIN_LOOP (src/config/ConfigStorage.sol#358-360)
Potential DoS Gas Limit Attack occur inAccountManager._removeSubAccount(address,address) (src/managers/AccountManager.sol#260-281)BEGIN_LOOP (src/managers/AccountManager.sol#267-276)
Potential DoS Gas Limit Attack occur inLockManager._execUSDPriceUpdate() (src/managers/LockManager.sol#506-527)BEGIN_LOOP (src/managers/LockManager.sol#512-523)
Potential DoS Gas Limit Attack occur inMigrationManager.migrateAllNFTs(address,uint32) (src/managers/MigrationManager.sol#246-265)BEGIN_LOOP (src/managers/MigrationManager.sol#256-263)
Potential DoS Gas Limit Attack occur inPrimordialManager._reconfigure() (src/managers/PrimordialManager.sol#29-53)BEGIN_LOOP (src/managers/PrimordialManager.sol#46-49)
Potential DoS Gas Limit Attack occur inRewardsManager.claimGasFeeForContracts(address[]) (src/managers/RewardsManager.sol#145-157)BEGIN_LOOP (src/managers/RewardsManager.sol#147-149)
Potential DoS Gas Limit Attack occur inSnuggeryManager._recalculateChonks(address) (src/managers/SnuggeryManager.sol#334-348)BEGIN_LOOP (src/managers/SnuggeryManager.sol#337-341)
Potential DoS Gas Limit Attack occur inMockNFTOverlord.levelUp(uint256,bytes) (src/mock/MockNFTOverlord.sol#188-243)BEGIN_LOOP (src/mock/MockNFTOverlord.sol#212-237)
Potential DoS Gas Limit Attack occur inMockNFTOverlord._populateDefaultProbabilities() (src/mock/MockNFTOverlord.sol#384-415)BEGIN_LOOP (src/mock/MockNFTOverlord.sol#409-414)
Potential DoS Gas Limit Attack occur inMockNFTOverlord._populateDefaultRealmLookup() (src/mock/MockNFTOverlord.sol#417-425)BEGIN_LOOP (src/mock/MockNFTOverlord.sol#422-424)
Potential DoS Gas Limit Attack occur inMockSnuggeryManager.setSnuggeryForTest(address,MunchablesCommonLib.SnuggeryNFT[]) (src/mock/MockSnuggeryManager.sol#11-22)BEGIN_LOOP (src/mock/MockSnuggeryManager.sol#18-20)
Potential DoS Gas Limit Attack occur inNFTOverlord._populateDefaultProbabilities() (src/overlords/NFTOverlord.sol#422-453)BEGIN_LOOP (src/overlords/NFTOverlord.sol#447-452)
Potential DoS Gas Limit Attack occur inNFTOverlord._populateDefaultRealmLookup() (src/overlords/NFTOverlord.sol#455-463)BEGIN_LOOP (src/overlords/NFTOverlord.sol#460-462)
Reference: https://swcregistry.io/docs/SWC-128

MockConfigNotifiable (src/mock/MockConfigNotifiable.sol#10-30) does not implement functions:
	- BaseConfigStorage.__BaseConfigStorage_reconfigure() (src/config/BaseConfigStorage.sol#85-87)
	- BaseConfigStorage.__BaseConfigStorage_setConfigStorage(address) (src/config/BaseConfigStorage.sol#79-83)
	- MockConfigNotifiable.verifyDirtyUint(uint256) (src/mock/MockConfigNotifiable.sol#22-24)
	- MockConfigNotifiable.verifyUint(StorageKey,uint256) (src/mock/MockConfigNotifiable.sol#26-29)
Reference:  

function:AccessControl.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#71-73)is public and can be replaced with external 
function:AccessControl.grantRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#122-124)is public and can be replaced with external 
function:AccessControl.revokeRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#137-139)is public and can be replaced with external 
function:AccessControl.renounceRole(bytes32,address) (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#157-163)is public and can be replaced with external 
function:Ownable.renounceOwnership() (lib/openzeppelin-contracts/contracts/access/Ownable.sol#76-78)is public and can be replaced with external 
function:Ownable.transferOwnership(address) (lib/openzeppelin-contracts/contracts/access/Ownable.sol#84-89)is public and can be replaced with external 
function:ProxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy,address,bytes) (lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#38-44)is public and can be replaced with external 
function:UUPSUpgradeable.upgradeToAndCall(address,bytes) (lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#86-89)is public and can be replaced with external 
function:ERC20.name() (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#58-60)is public and can be replaced with external 
function:ERC20.symbol() (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#66-68)is public and can be replaced with external 
function:ERC20.decimals() (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#83-85)is public and can be replaced with external 
function:ERC20.totalSupply() (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#90-92)is public and can be replaced with external 
function:ERC20.balanceOf(address) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#97-99)is public and can be replaced with external 
function:ERC20.transfer(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#109-113)is public and can be replaced with external 
function:ERC20.approve(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#132-136)is public and can be replaced with external 
function:ERC20.transferFrom(address,address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#154-159)is public and can be replaced with external 
function:ERC721.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#47-52)is public and can be replaced with external 
function:ERC721.balanceOf(address) (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#57-62)is public and can be replaced with external 
function:ERC721.ownerOf(uint256) (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#67-69)is public and can be replaced with external 
function:ERC721.name() (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#74-76)is public and can be replaced with external 
function:ERC721.symbol() (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#81-83)is public and can be replaced with external 
function:ERC721.tokenURI(uint256) (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#88-93)is public and can be replaced with external 
function:ERC721.approve(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#107-109)is public and can be replaced with external 
function:ERC721.getApproved(uint256) (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#114-118)is public and can be replaced with external 
function:ERC721.setApprovalForAll(address,bool) (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#123-125)is public and can be replaced with external 
function:ERC721.safeTransferFrom(address,address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#152-154)is public and can be replaced with external 
function:ERC721Enumerable.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#39-41)is public and can be replaced with external 
function:ERC721Enumerable.tokenOfOwnerByIndex(address,uint256) (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#46-51)is public and can be replaced with external 
function:ERC721Enumerable.tokenByIndex(uint256) (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#63-68)is public and can be replaced with external 
function:ERC721URIStorage.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#27-29)is public and can be replaced with external 
function:ERC721URIStorage.tokenURI(uint256) (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#34-50)is public and can be replaced with external 
function:ERC165.supportsInterface(bytes4) (lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#24-26)is public and can be replaced with external 
function:AccessControlUpgradeable.supportsInterface(bytes4) (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#90-92)is public and can be replaced with external 
function:AccessControlUpgradeable.grantRole(bytes32,address) (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#143-145)is public and can be replaced with external 
function:AccessControlUpgradeable.revokeRole(bytes32,address) (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#158-160)is public and can be replaced with external 
function:AccessControlUpgradeable.renounceRole(bytes32,address) (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#178-184)is public and can be replaced with external 
function:ERC721Upgradeable.supportsInterface(bytes4) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#66-71)is public and can be replaced with external 
function:ERC721Upgradeable.balanceOf(address) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#76-82)is public and can be replaced with external 
function:ERC721Upgradeable.ownerOf(uint256) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#87-89)is public and can be replaced with external 
function:ERC721Upgradeable.name() (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#94-97)is public and can be replaced with external 
function:ERC721Upgradeable.symbol() (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#102-105)is public and can be replaced with external 
function:ERC721Upgradeable.tokenURI(uint256) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#110-115)is public and can be replaced with external 
function:ERC721Upgradeable.approve(address,uint256) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#129-131)is public and can be replaced with external 
function:ERC721Upgradeable.getApproved(uint256) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#136-140)is public and can be replaced with external 
function:ERC721Upgradeable.setApprovalForAll(address,bool) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#145-147)is public and can be replaced with external 
function:ERC721Upgradeable.safeTransferFrom(address,address,uint256) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#175-177)is public and can be replaced with external 
function:ERC721EnumerableUpgradeable.supportsInterface(bytes4) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#57-59)is public and can be replaced with external 
function:ERC721EnumerableUpgradeable.tokenOfOwnerByIndex(address,uint256) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#64-70)is public and can be replaced with external 
function:ERC721EnumerableUpgradeable.tokenByIndex(uint256) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#83-89)is public and can be replaced with external 
function:ERC721URIStorageUpgradeable.supportsInterface(bytes4) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#45-47)is public and can be replaced with external 
function:ERC721URIStorageUpgradeable.tokenURI(uint256) (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#52-69)is public and can be replaced with external 
function:ERC165Upgradeable.supportsInterface(bytes4) (lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#30-32)is public and can be replaced with external 
function:BaseConfigStorageUpgradeable.initialize(address) (src/config/BaseConfigStorageUpgradeable.sol#16-18)is public and can be replaced with external 
function:ConfigStorage.setRole(Role,address,address) (src/config/ConfigStorage.sol#27-33)is public and can be replaced with external 
function:ConfigStorage.setUniversalRole(Role,address) (src/config/ConfigStorage.sol#35-37)is public and can be replaced with external 
function:ConfigStorage.getRole(Role) (src/config/ConfigStorage.sol#39-41)is public and can be replaced with external 
function:ConfigStorage.getUniversalRole(Role) (src/config/ConfigStorage.sol#50-52)is public and can be replaced with external 
function:ConfigStorage.setUint(StorageKey,uint256,bool) (src/config/ConfigStorage.sol#55-62)is public and can be replaced with external 
function:ConfigStorage.getUint(StorageKey) (src/config/ConfigStorage.sol#65-67)is public and can be replaced with external 
function:ConfigStorage.setUintArray(StorageKey,uint256[],bool) (src/config/ConfigStorage.sol#69-96)is public and can be replaced with external 
function:ConfigStorage.getUintArray(StorageKey) (src/config/ConfigStorage.sol#98-109)is public and can be replaced with external 
function:ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool) (src/config/ConfigStorage.sol#111-139)is public and can be replaced with external 
function:ConfigStorage.getSmallUintArray(StorageKey) (src/config/ConfigStorage.sol#141-152)is public and can be replaced with external 
function:ConfigStorage.setSmallInt(StorageKey,int16,bool) (src/config/ConfigStorage.sol#155-162)is public and can be replaced with external 
function:ConfigStorage.getSmallInt(StorageKey) (src/config/ConfigStorage.sol#165-167)is public and can be replaced with external 
function:ConfigStorage.setSmallIntArray(StorageKey,int16[],bool) (src/config/ConfigStorage.sol#169-197)is public and can be replaced with external 
function:ConfigStorage.getSmallIntArray(StorageKey) (src/config/ConfigStorage.sol#199-210)is public and can be replaced with external 
function:ConfigStorage.setBool(StorageKey,bool,bool) (src/config/ConfigStorage.sol#213-220)is public and can be replaced with external 
function:ConfigStorage.getBool(StorageKey) (src/config/ConfigStorage.sol#223-225)is public and can be replaced with external 
function:ConfigStorage.setAddress(StorageKey,address,bool) (src/config/ConfigStorage.sol#228-235)is public and can be replaced with external 
function:ConfigStorage.setAddresses(StorageKey[],address[],bool) (src/config/ConfigStorage.sol#238-247)is public and can be replaced with external 
function:ConfigStorage.getAddress(StorageKey) (src/config/ConfigStorage.sol#250-252)is public and can be replaced with external 
function:ConfigStorage.setAddressArray(StorageKey,address[],bool) (src/config/ConfigStorage.sol#254-281)is public and can be replaced with external 
function:ConfigStorage.getAddressArray(StorageKey) (src/config/ConfigStorage.sol#283-294)is public and can be replaced with external 
function:ConfigStorage.setBytes32(StorageKey,bytes32,bool) (src/config/ConfigStorage.sol#297-304)is public and can be replaced with external 
function:ConfigStorage.getBytes32(StorageKey) (src/config/ConfigStorage.sol#307-309)is public and can be replaced with external 
function:ConfigStorage.addNotifiableAddress(address) (src/config/ConfigStorage.sol#312-314)is public and can be replaced with external 
function:ConfigStorage.addNotifiableAddresses(address[]) (src/config/ConfigStorage.sol#316-322)is public and can be replaced with external 
function:ConfigStorage.removeNotifiableAddress(address) (src/config/ConfigStorage.sol#325-335)is public and can be replaced with external 
function:ConfigStorage.manualNotify(uint8,uint8) (src/config/ConfigStorage.sol#345-350)is public and can be replaced with external 
function:ConfigStorage.manualNotifyAddress(address) (src/config/ConfigStorage.sol#352-354)is public and can be replaced with external 
function:BaseBlastManagerUpgradeable.initialize(address) (src/managers/BaseBlastManagerUpgradeable.sol#17-21)is public and can be replaced with external 
function:AccountManager.initialize(address) (src/managers/AccountManager.sol#48-51)is public and can be replaced with external 
function:ClaimManager.initialize(address) (src/managers/ClaimManager.sol#49-52)is public and can be replaced with external 
function:NFTAttributesManagerV1.getGameAttributeDataType(uint8) (src/managers/NFTAttributeManagerV1.sol#171-196)is public and can be replaced with external 
function:MockConfigNotifiable.verifyDirtyUint(uint256) (src/mock/MockConfigNotifiable.sol#22-24)is public and can be replaced with external 
function:MockConfigNotifiable.verifyUint(StorageKey,uint256) (src/mock/MockConfigNotifiable.sol#26-29)is public and can be replaced with external 
function:MockMunchNFT.burn(uint256) (src/mock/MockMunchNFT.sol#7)is public and can be replaced with external 
function:MockMunchNFT.mint(address) (src/mock/MockMunchNFT.sol#9)is public and can be replaced with external 
function:RNGProxySelfHosted.provideRandom(uint256,bytes) (src/rng/RNGProxySelfHosted.sol#12-17)is public and can be replaced with external 
function:BaseRNGProxy.requestRandom(address,bytes4,uint256) (src/rng/BaseRNGProxy.sol#21-32)is public and can be replaced with external 
function:MockRNGProxy.provideRandomForTest(uint256,bytes) (src/mock/MockRNGProxy.sol#25-27)is public and can be replaced with external 
function:MockRNGProxy.callLevelUpForTest(uint256,bytes) (src/mock/MockRNGProxy.sol#29-31)is public and can be replaced with external 
function:MockRNGProxy.callRevealForTest(uint256,bytes) (src/mock/MockRNGProxy.sol#33-38)is public and can be replaced with external 
function:MockRNGProxy.callRevealFromPrimordialForTest(uint256,bytes) (src/mock/MockRNGProxy.sol#40-45)is public and can be replaced with external 
function:RNGProxyAPI3.requestRandom(address,bytes4,uint256) (src/rng/RNGProxyAPI3.sol#30-47)is public and can be replaced with external 
function:MunchNFT.tokenURI(uint256) (src/tokens/MunchNFT.sol#114-123)is public and can be replaced with external 
function:MunchNFT.supportsInterface(bytes4) (src/tokens/MunchNFT.sol#125-134)is public and can be replaced with external 
function:OldMunchNFT.initialize(address,address,address,address,address) (src/tokens/OldMunchNFT.sol#53-75)is public and can be replaced with external 
function:OldMunchNFT.pause() (src/tokens/OldMunchNFT.sol#78-80)is public and can be replaced with external 
function:OldMunchNFT.unpause() (src/tokens/OldMunchNFT.sol#83-85)is public and can be replaced with external 
function:OldMunchNFT.burn(uint256) (src/tokens/OldMunchNFT.sol#87-89)is public and can be replaced with external 
function:OldMunchNFT.configureContract(address) (src/tokens/OldMunchNFT.sol#93-98)is public and can be replaced with external 
function:OldMunchNFT.revealNFT(address,string,bytes32,bytes,MunchablesCommonLib.Realm,MunchablesCommonLib.Rarity,uint16) (src/tokens/OldMunchNFT.sol#110-140)is public and can be replaced with external 
function:OldMunchNFT.setURI(uint256,string) (src/tokens/OldMunchNFT.sol#145-160)is public and can be replaced with external 
function:OldMunchNFT.fetchTokens(address) (src/tokens/OldMunchNFT.sol#166-179)is public and can be replaced with external 
function:OldMunchNFT.safeMint(address,string) (src/tokens/OldMunchNFT.sol#184-193)is public and can be replaced with external 
function:OldMunchNFT.supportsInterface(bytes4) (src/tokens/OldMunchNFT.sol#253-267)is public and can be replaced with external 
Reference:  

OldMunchNFTContractConfigured(address) (src/tokens/OldMunchNFT.sol#269) is never emitted in OldMunchNFT (src/tokens/OldMunchNFT.sol#13-281)
OldMunchNFTSetURI(uint256,string) (src/tokens/OldMunchNFT.sol#271) is never emitted in OldMunchNFT (src/tokens/OldMunchNFT.sol#13-281)
OldMunchNFTRevealNFT(address,uint256,string,bytes32,MunchablesCommonLib.Realm,MunchablesCommonLib.Rarity,uint16) (src/tokens/OldMunchNFT.sol#272-280) is never emitted in OldMunchNFT (src/tokens/OldMunchNFT.sol#13-281)
Reference: https://certik-public-assets.s3.amazonaws.com/CertiK-Audit-for-Kitty-inu.pdf

	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/access/AccessControl.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/access/IAccessControl.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/access/Ownable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/interfaces/IERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/interfaces/IERC1967.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/interfaces/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/interfaces/IERC4906.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/interfaces/IERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/interfaces/draft-IERC1822.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/interfaces/draft-IERC6093.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/proxy/Proxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/proxy/beacon/IBeacon.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC20/extensions/IERC20Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC721/IERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC721/IERC721Receiver.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Pausable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/IERC721Enumerable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/token/ERC721/extensions/IERC721Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/Context.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/Pausable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/StorageSlot.sol#5)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/Strings.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/introspection/IERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/math/Math.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts/contracts/utils/math/SignedMath.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721PausableUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/utils/PausableUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/utils/ReentrancyGuardUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IAirnodeRrpV0.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IAuthorizationUtilsV0.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/ITemplateUtilsV0.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IWithdrawalUtilsV0.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@api3/airnode-protocol/contracts/rrp/requesters/RrpRequesterV0.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (src/config/ConfigStorage.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/interfaces/IBlast.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (src/interfaces/IConfigNotifiable.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/interfaces/IConfigStorage.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/interfaces/IPrimordialManager.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/interfaces/IRNGProxy.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/interfaces/IRNGProxySelfHosted.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/libraries/SignatureVerifier.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/mock/MockBlast.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/mock/MockConfigNotifiable.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/rng/BaseRNGProxy.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/rng/RNGProxyAPI3.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.25. ^0.8.25 (src/rng/RNGProxySelfHosted.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (src/tokens/OldMunchNFT.sol#2)
Reference:  

BaseBlastManager._blastClaimableConfigured (src/managers/BaseBlastManager.sol#23) should be constant
MigrationManager.purchasedUnlockPrice (src/managers/MigrationManager.sol#23) should be constant
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
 - [reentrancy-without-eth-transfer](#reentrancy-without-eth-transfer) (2 results) (Critical)
 - [state-variable-not-initialized](#state-variable-not-initialized) (2 results) (High)
 - [invalid-array-deletion](#invalid-array-deletion) (9 results) (Medium)
 - [centralized-risk-medium](#centralized-risk-medium) (4 results) (Medium)
 - [public-mint-burn](#public-mint-burn) (3 results) (Medium)
 - [transfer-in-loop](#transfer-in-loop) (1 results) (Medium)
 - [uninitialized-local](#uninitialized-local) (45 results) (Medium)
 - [unused-return](#unused-return) (20 results) (Medium)
 - [void-function](#void-function) (20 results) (Medium)
 - [centralized-risk-low](#centralized-risk-low) (6 results) (Low)
 - [pess-event-setter](#pess-event-setter) (95 results) (Low)
 - [missing-zero-check](#missing-zero-check) (42 results) (Low)
 - [centralized-risk-informational](#centralized-risk-informational) (50 results) (Informational)
 - [centralized-risk-other](#centralized-risk-other) (44 results) (Informational)
 - [dead-function](#dead-function) (3 results) (Informational)
 - [error-msg](#error-msg) (1 results) (Informational)
 - [price-manipulation-info](#price-manipulation-info) (2 results) (Informational)
 - [uncontrolled-resource-consumption](#uncontrolled-resource-consumption) (13 results) (Informational)
 - [no-derived-function](#no-derived-function) (1 results) (Informational)
 - [unnecessary-public-function-modifier](#unnecessary-public-function-modifier) (108 results) (Informational)
 - [unused-event](#unused-event) (3 results) (Informational)
 - [version-only](#version-only) (67 results) (Informational)
 - [state-should-be-constant](#state-should-be-constant) (2 results) (Optimization)
## reentrancy-without-eth-transfer
Impact: Critical
Confidence: Medium
 - [ ] ID-0
Reentrancy in [LockManager._lock(address,uint256,address,address)](src/managers/LockManager.sol#L311-L398):
	External calls:
	- [token_scope_0.transferFrom(_tokenOwner,address(this),_quantity)](src/managers/LockManager.sol#L376)

src/managers/LockManager.sol#L311-L398


 - [ ] ID-1
Reentrancy in [BaseRNGProxy._callback(uint256,bytes)](src/rng/BaseRNGProxy.sol#L34-L54):
	External calls:
	- [(success,returnData) = data.targetContract.call(callData)](src/rng/BaseRNGProxy.sol#L45-L47)

src/rng/BaseRNGProxy.sol#L34-L54


## state-variable-not-initialized
Impact: High
Confidence: High
 - [ ] ID-2
state variable: [BaseConfigStorage.configStorage](src/config/BaseConfigStorage.sol#L7) not initialized and not written in contract but be used in contract

src/config/BaseConfigStorage.sol#L7


 - [ ] ID-3
state variable: [BaseConfigStorage._paused](src/config/BaseConfigStorage.sol#L8) not initialized and not written in contract but be used in contract

src/config/BaseConfigStorage.sol#L8


## invalid-array-deletion
Impact: Medium
Confidence: High
 - [ ] ID-4
[LockManager._execUSDPriceUpdate()](src/managers/LockManager.sol#L506-L527) deletes [ILockManager.USDUpdateProposal](src/interfaces/ILockManager.sol#L60-L69) which contains a array:
	-[delete usdUpdateProposal](src/managers/LockManager.sol#L525)

src/managers/LockManager.sol#L506-L527


 - [ ] ID-5
[AccountManager.execSprayProposal(address)](src/managers/AccountManager.sol#L180-L210) deletes [IAccountManager.SprayProposal](src/interfaces/IAccountManager.sol#L15-L18) which contains a array:
	-[delete sprayProposals[_proposer]](src/managers/AccountManager.sol#L209)

src/managers/AccountManager.sol#L180-L210


 - [ ] ID-6
[AccountManager.removeSprayProposal(address)](src/managers/AccountManager.sol#L213-L228) deletes [IAccountManager.SprayProposal](src/interfaces/IAccountManager.sol#L15-L18) which contains a array:
	-[delete sprayProposals[_proposer]](src/managers/AccountManager.sol#L227)

src/managers/AccountManager.sol#L213-L228


 - [ ] ID-7
[AccountManager.spraySchnibblesPropose(address[],uint256[])](src/managers/AccountManager.sol#L135-L177) deletes [IAccountManager.SprayProposal](src/interfaces/IAccountManager.sol#L15-L18) which contains a array:
	-[delete sprayProposals[proposer]](src/managers/AccountManager.sol#L161)

src/managers/AccountManager.sol#L135-L177


 - [ ] ID-8
[LockManager.proposeUSDPrice(uint256,address[])](src/managers/LockManager.sol#L142-L174) deletes [ILockManager.USDUpdateProposal](src/interfaces/ILockManager.sol#L60-L69) which contains a array:
	-[delete usdUpdateProposal](src/managers/LockManager.sol#L161)

src/managers/LockManager.sol#L142-L174


 - [ ] ID-9
[LockManager.disapproveUSDPrice(uint256)](src/managers/LockManager.sol#L210-L242) deletes [ILockManager.USDUpdateProposal](src/interfaces/ILockManager.sol#L60-L69) which contains a array:
	-[delete usdUpdateProposal](src/managers/LockManager.sol#L238)

src/managers/LockManager.sol#L210-L242


 - [ ] ID-10
[LockManager._execUSDPriceUpdate()](src/managers/LockManager.sol#L506-L527) deletes [ILockManager.USDUpdateProposal](src/interfaces/ILockManager.sol#L60-L69) which contains a mapping:
	-[delete usdUpdateProposal](src/managers/LockManager.sol#L525)

src/managers/LockManager.sol#L506-L527


 - [ ] ID-11
[LockManager.proposeUSDPrice(uint256,address[])](src/managers/LockManager.sol#L142-L174) deletes [ILockManager.USDUpdateProposal](src/interfaces/ILockManager.sol#L60-L69) which contains a mapping:
	-[delete usdUpdateProposal](src/managers/LockManager.sol#L161)

src/managers/LockManager.sol#L142-L174


 - [ ] ID-12
[LockManager.disapproveUSDPrice(uint256)](src/managers/LockManager.sol#L210-L242) deletes [ILockManager.USDUpdateProposal](src/interfaces/ILockManager.sol#L60-L69) which contains a mapping:
	-[delete usdUpdateProposal](src/managers/LockManager.sol#L238)

src/managers/LockManager.sol#L210-L242


## centralized-risk-medium
Impact: Medium
Confidence: High
 - [ ] ID-13
	- [SnuggeryManager.importMunchable(uint256)](src/managers/SnuggeryManager.sol#L79-L106)

src/managers/SnuggeryManager.sol#L79-L106


 - [ ] ID-14
	- [MigrationManager.lockFundsForAllMigration()](src/managers/MigrationManager.sol#L221-L244)

src/managers/MigrationManager.sol#L221-L244


 - [ ] ID-15
	- [FundTreasuryDistributor.receiveTokens(IDistributor.TokenBag[])](src/distributors/FundTreasuryDistributor.sol#L26-L50)

src/distributors/FundTreasuryDistributor.sol#L26-L50


 - [ ] ID-16
	- [SnuggeryManager.exportMunchable(uint256)](src/managers/SnuggeryManager.sol#L109-L130)

src/managers/SnuggeryManager.sol#L109-L130


## public-mint-burn
Impact: Medium
Confidence: Medium
 - [ ] ID-17
public mint or burn found in [MockMunchNFT.burn(uint256)](src/mock/MockMunchNFT.sol#L7)
src/mock/MockMunchNFT.sol#L7


 - [ ] ID-18
public mint or burn found in [MockMunchNFT.mint(address)](src/mock/MockMunchNFT.sol#L9)
src/mock/MockMunchNFT.sol#L9


 - [ ] ID-19
public mint or burn found in [TestERC20Token.mint(address,uint256)](src/tokens/TestERC20Token.sol#L31-L33)
src/tokens/TestERC20Token.sol#L31-L33


## transfer-in-loop
Impact: Medium
Confidence: Medium
 - [ ] ID-20
[FundTreasuryDistributor.receiveTokens(IDistributor.TokenBag[])](src/distributors/FundTreasuryDistributor.sol#L26-L50) use transfer in a loop: [address(_treasury).transfer(amount)](src/distributors/FundTreasuryDistributor.sol#L38)

src/distributors/FundTreasuryDistributor.sol#L26-L50


## uninitialized-local
Impact: Medium
Confidence: Medium
 - [ ] ID-21
[SnuggeryManager._recalculateChonks(address).i](src/managers/SnuggeryManager.sol#L337) is a local variable never initialized

src/managers/SnuggeryManager.sol#L337


 - [ ] ID-22
[LockManager.getLockedWeightedValue(address).i](src/managers/LockManager.sol#L466) is a local variable never initialized

src/managers/LockManager.sol#L466


 - [ ] ID-23
[BonusManager.getHarvestBonus(address)._migrationBonus](src/managers/BonusManager.sol#L131) is a local variable never initialized

src/managers/BonusManager.sol#L131


 - [ ] ID-24
[PrimordialManager._reconfigure().i](src/managers/PrimordialManager.sol#L46) is a local variable never initialized

src/managers/PrimordialManager.sol#L46


 - [ ] ID-25
[LockManager._execUSDPriceUpdate().i](src/managers/LockManager.sol#L512) is a local variable never initialized

src/managers/LockManager.sol#L512


 - [ ] ID-26
[MockNFTOverlord.levelUp(uint256,bytes).currentValue](src/mock/MockNFTOverlord.sol#L227) is a local variable never initialized

src/mock/MockNFTOverlord.sol#L227


 - [ ] ID-27
[ClaimManager.burnNFTsForPoints(address,uint8[]).i](src/managers/ClaimManager.sol#L123) is a local variable never initialized

src/managers/ClaimManager.sol#L123


 - [ ] ID-28
[NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).i](src/managers/NFTAttributeManagerV1.sol#L150) is a local variable never initialized

src/managers/NFTAttributeManagerV1.sol#L150


 - [ ] ID-29
[RewardsManager.claimYieldForContracts(address[]).ongoingUSDB](src/managers/RewardsManager.sol#L116) is a local variable never initialized

src/managers/RewardsManager.sol#L116


 - [ ] ID-30
[ConfigStorage.getAddressArray(StorageKey).i](src/config/ConfigStorage.sol#L289) is a local variable never initialized

src/config/ConfigStorage.sol#L289


 - [ ] ID-31
[NFTOverlord._calculateRaritySpecies(bytes,MunchablesCommonLib.Rarity).selectedSpeciesId](src/overlords/NFTOverlord.sol#L376) is a local variable never initialized

src/overlords/NFTOverlord.sol#L376


 - [ ] ID-32
[AccountManager.addSubAccount(address).i](src/managers/AccountManager.sol#L244) is a local variable never initialized

src/managers/AccountManager.sol#L244


 - [ ] ID-33
[BonusManager._calculateMunchadexBonus(address).i](src/managers/BonusManager.sol#L226) is a local variable never initialized

src/managers/BonusManager.sol#L226


 - [ ] ID-34
tokenId is a member never initialized in [MigrationManager._migrateNFTs(address,address,uint256[]).snapshot](src/managers/MigrationManager.sol#L300)

src/managers/MigrationManager.sol#L300


 - [ ] ID-35
[MockNFTOverlord._calculateRaritySpecies(bytes,MunchablesCommonLib.Rarity).selectedSpeciesId](src/mock/MockNFTOverlord.sol#L338) is a local variable never initialized

src/mock/MockNFTOverlord.sol#L338


 - [ ] ID-36
[RewardsManager.claimYieldForContracts(address[]).ongoingWETH](src/managers/RewardsManager.sol#L115) is a local variable never initialized

src/managers/RewardsManager.sol#L115


 - [ ] ID-37
[AccountManager.execSprayProposal(address).i](src/managers/AccountManager.sol#L197) is a local variable never initialized

src/managers/AccountManager.sol#L197


 - [ ] ID-38
registrationDate is a member never initialized in [LockManager._lock(address,uint256,address,address)._player](src/managers/LockManager.sol#L319)

src/managers/LockManager.sol#L319


 - [ ] ID-39
[ConfigStorage.setAddressArray(StorageKey,address[],bool).i](src/config/ConfigStorage.sol#L262) is a local variable never initialized

src/config/ConfigStorage.sol#L262


 - [ ] ID-40
[MockNFTOverlord.levelUp(uint256,bytes).i](src/mock/MockNFTOverlord.sol#L212) is a local variable never initialized

src/mock/MockNFTOverlord.sol#L212


 - [ ] ID-41
[ConfigStorage.getSmallUintArray(StorageKey).i](src/config/ConfigStorage.sol#L147) is a local variable never initialized

src/config/ConfigStorage.sol#L147


 - [ ] ID-42
referrer is a member never initialized in [ClaimManager._claimPoints(address).player](src/managers/ClaimManager.sol#L203)

src/managers/ClaimManager.sol#L203


 - [ ] ID-43
[SnuggeryManager.getTotalChonk(address).i](src/managers/SnuggeryManager.sol#L289) is a local variable never initialized

src/managers/SnuggeryManager.sol#L289


 - [ ] ID-44
[NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).i_scope_0](src/managers/NFTAttributeManagerV1.sol#L157) is a local variable never initialized

src/managers/NFTAttributeManagerV1.sol#L157


 - [ ] ID-45
[NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).j](src/managers/NFTAttributeManagerV1.sol#L164) is a local variable never initialized

src/managers/NFTAttributeManagerV1.sol#L164


 - [ ] ID-46
[SnuggeryManager._findSnuggeryIndex(address,uint256).i](src/managers/SnuggeryManager.sol#L310) is a local variable never initialized

src/managers/SnuggeryManager.sol#L310


 - [ ] ID-47
[FundTreasuryDistributor.receiveTokens(IDistributor.TokenBag[]).i](src/distributors/FundTreasuryDistributor.sol#L31) is a local variable never initialized

src/distributors/FundTreasuryDistributor.sol#L31


 - [ ] ID-48
[MockNFTOverlord.levelUp(uint256,bytes).zeroed](src/mock/MockNFTOverlord.sol#L211) is a local variable never initialized

src/mock/MockNFTOverlord.sol#L211


 - [ ] ID-49
[NFTOverlord.levelUp(uint256,bytes).currentValue](src/overlords/NFTOverlord.sol#L260) is a local variable never initialized

src/overlords/NFTOverlord.sol#L260


 - [ ] ID-50
[ConfigStorage.setSmallIntArray(StorageKey,int16[],bool).i](src/config/ConfigStorage.sol#L177) is a local variable never initialized

src/config/ConfigStorage.sol#L177


 - [ ] ID-51
[ConfigStorage.setUintArray(StorageKey,uint256[],bool).i](src/config/ConfigStorage.sol#L77) is a local variable never initialized

src/config/ConfigStorage.sol#L77


 - [ ] ID-52
[MunchadexManager.getMunchadexInfo(address).i_scope_0](src/managers/MunchadexManager.sol#L130) is a local variable never initialized

src/managers/MunchadexManager.sol#L130


 - [ ] ID-53
[ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool).i](src/config/ConfigStorage.sol#L119) is a local variable never initialized

src/config/ConfigStorage.sol#L119


 - [ ] ID-54
[ConfigStorage.getUintArray(StorageKey).i](src/config/ConfigStorage.sol#L104) is a local variable never initialized

src/config/ConfigStorage.sol#L104


 - [ ] ID-55
[NFTAttributesManagerV1.setGameAttributes(uint256,MunchablesCommonLib.NFTGameAttribute[]).i](src/managers/NFTAttributeManagerV1.sol#L92) is a local variable never initialized

src/managers/NFTAttributeManagerV1.sol#L92


 - [ ] ID-56
[NFTOverlord.levelUp(uint256,bytes).i](src/overlords/NFTOverlord.sol#L245) is a local variable never initialized

src/overlords/NFTOverlord.sol#L245


 - [ ] ID-57
[ConfigStorage.getSmallIntArray(StorageKey).i](src/config/ConfigStorage.sol#L205) is a local variable never initialized

src/config/ConfigStorage.sol#L205


 - [ ] ID-58
[NFTOverlord.levelUp(uint256,bytes).zeroed](src/overlords/NFTOverlord.sol#L244) is a local variable never initialized

src/overlords/NFTOverlord.sol#L244


 - [ ] ID-59
[ConfigStorage.setAddresses(StorageKey[],address[],bool).i](src/config/ConfigStorage.sol#L243) is a local variable never initialized

src/config/ConfigStorage.sol#L243


 - [ ] ID-60
[MockSnuggeryManager.setSnuggeryForTest(address,MunchablesCommonLib.SnuggeryNFT[]).i](src/mock/MockSnuggeryManager.sol#L18) is a local variable never initialized

src/mock/MockSnuggeryManager.sol#L18


 - [ ] ID-61
maxSnuggerySize is a member never initialized in [BonusManager._calculateLevelBonus(address).player](src/managers/BonusManager.sol#L197)

src/managers/BonusManager.sol#L197


 - [ ] ID-62
[AccountManager.spraySchnibblesPropose(address[],uint256[]).i](src/managers/AccountManager.sol#L163) is a local variable never initialized

src/managers/AccountManager.sol#L163


 - [ ] ID-63
[MunchadexManager.getMunchadexInfo(address).i](src/managers/MunchadexManager.sol#L124) is a local variable never initialized

src/managers/MunchadexManager.sol#L124


 - [ ] ID-64
[RewardsManager.claimYieldForContracts(address[]).ongoingETH](src/managers/RewardsManager.sol#L114) is a local variable never initialized

src/managers/RewardsManager.sol#L114


 - [ ] ID-65
[LockManager.setLockDuration(uint256).i](src/managers/LockManager.sol#L252) is a local variable never initialized

src/managers/LockManager.sol#L252


## unused-return
Impact: Medium
Confidence: Medium
 - [ ] ID-66
[MigrationManager.burnUnrevealedForPoints()](src/managers/MigrationManager.sol#L135-L146)have ignores return value in [_claimManager.burnUnrevealedForPoints(msg.sender,amountToSwap)](src/managers/MigrationManager.sol#L141)

src/managers/MigrationManager.sol#L135-L146


 - [ ] ID-67
[MockMigrationManager.burnUnrevealedForPoints(address,uint256)](src/mock/MockMigrationManager.sol#L40-L46)have ignores return value in [IClaimManager(configStorage.getAddress(StorageKey.ClaimManager)).burnUnrevealedForPoints(_player,numUnrevealed)](src/mock/MockMigrationManager.sol#L44-L45)

src/mock/MockMigrationManager.sol#L40-L46


 - [ ] ID-68
[MockNFTOverlord.revealFromPrimordial(uint256,bytes)](src/mock/MockNFTOverlord.sol#L141-L162)have ignores return value in [IMunchNFT(munchNFT).mint(player)](src/mock/MockNFTOverlord.sol#L152)

src/mock/MockNFTOverlord.sol#L141-L162


 - [ ] ID-69
[MockNFTOverlord._reveal(address,MunchablesCommonLib.Rarity,uint16)](src/mock/MockNFTOverlord.sol#L107-L127)have ignores return value in [IMunchNFT(munchNFT).mint(player)](src/mock/MockNFTOverlord.sol#L117)

src/mock/MockNFTOverlord.sol#L107-L127


 - [ ] ID-70
[ERC1967Utils.upgradeToAndCall(address,bytes)](lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#L83-L92)have ignores return value in [Address.functionDelegateCall(newImplementation,data)](lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#L88)

lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#L83-L92


 - [ ] ID-71
[NFTOverlord.revealFromPrimordial(uint256,bytes)](src/overlords/NFTOverlord.sol#L154-L186)have ignores return value in [IMunchNFT(munchNFT).mint(player)](src/overlords/NFTOverlord.sol#L176)

src/overlords/NFTOverlord.sol#L154-L186


 - [ ] ID-72
[NFTOverlord.reveal(uint256,bytes)](src/overlords/NFTOverlord.sol#L100-L132)have ignores return value in [IMunchNFT(munchNFT).mint(player)](src/overlords/NFTOverlord.sol#L122)

src/overlords/NFTOverlord.sol#L100-L132


 - [ ] ID-73
[MigrationManager.burnRemainingPurchasedNFTs(address,uint32)](src/managers/MigrationManager.sol#L184-L219)have ignores return value in [_claimManager.burnNFTsForPoints(_user,tokenIdsByRarity)](src/managers/MigrationManager.sol#L217)

src/managers/MigrationManager.sol#L184-L219


 - [ ] ID-74
[MockNFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[])](src/mock/MockNFTOverlord.sol#L165-L185)have ignores return value in [IMunchNFT(munchNFT).mint(_player)](src/mock/MockNFTOverlord.sol#L176)

src/mock/MockNFTOverlord.sol#L165-L185


 - [ ] ID-75
[MigrationManager.burnNFTs(address,uint32)](src/managers/MigrationManager.sol#L148-L182)have ignores return value in [_claimManager.burnNFTsForPoints(_user,tokenIdsByRarity)](src/managers/MigrationManager.sol#L179)

src/managers/MigrationManager.sol#L148-L182


 - [ ] ID-76
[ERC1967Utils.upgradeBeaconToAndCall(address,bytes)](lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#L173-L182)have ignores return value in [Address.functionDelegateCall(IBeacon(newBeacon).implementation(),data)](lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#L178)

lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#L173-L182


 - [ ] ID-77
[MockRNGProxy.callRevealFromPrimordialForTest(uint256,bytes)](src/mock/MockRNGProxy.sol#L40-L45)have ignores return value in [nftOverlord.revealFromPrimordial(_player,_signature)](src/mock/MockRNGProxy.sol#L44)

src/mock/MockRNGProxy.sol#L40-L45


 - [ ] ID-78
[RewardsManager._forwardYield(uint256,uint256,uint256)](src/managers/RewardsManager.sol#L219-L264)have ignores return value in [USDB.approve(address(yieldDistributor),ongoingUSDB)](src/managers/RewardsManager.sol#L230)

src/managers/RewardsManager.sol#L219-L264


 - [ ] ID-79
[MockMigrationManager.callMintForMigrationForTest(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[])](src/mock/MockMigrationManager.sol#L18-L30)have ignores return value in [_nftOverlord.mintForMigration(_player,_attributes,_immutableAttributes,_gameAttributes)](src/mock/MockMigrationManager.sol#L24-L29)

src/mock/MockMigrationManager.sol#L18-L30


 - [ ] ID-80
[BaseBlastManager.__BaseBlastManager_reconfigure()](src/managers/BaseBlastManager.sol#L30-L90)have ignores return value in [_USDB.configure(YieldMode.CLAIMABLE)](src/managers/BaseBlastManager.sol#L73)

src/managers/BaseBlastManager.sol#L30-L90


 - [ ] ID-81
[BaseBlastManager.claimERC20Yield(address,uint256)](src/managers/BaseBlastManager.sol#L109-L117)have ignores return value in [IERC20Rebasing(_tokenContract).claim(configStorage.getAddress(StorageKey.RewardsManager),_amount)](src/managers/BaseBlastManager.sol#L113-L116)

src/managers/BaseBlastManager.sol#L109-L117


 - [ ] ID-82
[MockMigrationManager.burnNFTsForPoints(address,uint8[])](src/mock/MockMigrationManager.sol#L32-L38)have ignores return value in [IClaimManager(configStorage.getAddress(StorageKey.ClaimManager)).burnNFTsForPoints(_player,_rarities)](src/mock/MockMigrationManager.sol#L36-L37)

src/mock/MockMigrationManager.sol#L32-L38


 - [ ] ID-83
[NFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[])](src/overlords/NFTOverlord.sol#L189-L215)have ignores return value in [IMunchNFT(munchNFT).mint(_player)](src/overlords/NFTOverlord.sol#L206)

src/overlords/NFTOverlord.sol#L189-L215


 - [ ] ID-84
[MigrationManager._migrateNFTs(address,address,uint256[])](src/managers/MigrationManager.sol#L290-L349)have ignores return value in [WETH.approve(address(_lockManager),quantity)](src/managers/MigrationManager.sol#L342)

src/managers/MigrationManager.sol#L290-L349


 - [ ] ID-85
[MockRNGProxy.callRevealForTest(uint256,bytes)](src/mock/MockRNGProxy.sol#L33-L38)have ignores return value in [nftOverlord.reveal(_player,_signature)](src/mock/MockRNGProxy.sol#L37)

src/mock/MockRNGProxy.sol#L33-L38


## void-function
Impact: Medium
Confidence: High
 - [ ] ID-86
function:[AccessControlUpgradeable.__AccessControl_init_unchained()](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L85-L86)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L85-L86


 - [ ] ID-87
function:[ContextUpgradeable.__Context_init_unchained()](lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#L21-L22)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#L21-L22


 - [ ] ID-88
function:[AccessControlUpgradeable.__AccessControl_init()](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L82-L83)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L82-L83


 - [ ] ID-89
function:[BaseRNGProxy.configUpdated()](src/rng/BaseRNGProxy.sol#L19)is empty 

src/rng/BaseRNGProxy.sol#L19


 - [ ] ID-90
function:[ERC165Upgradeable.__ERC165_init_unchained()](lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#L25-L26)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#L25-L26


 - [ ] ID-91
function:[OldMunchNFT.revealNFT(address,string,bytes32,bytes,MunchablesCommonLib.Realm,MunchablesCommonLib.Rarity,uint16)](src/tokens/OldMunchNFT.sol#L110-L140)is empty 

src/tokens/OldMunchNFT.sol#L110-L140


 - [ ] ID-92
function:[MockMunchNFT.mint(address)](src/mock/MockMunchNFT.sol#L9)is empty 

src/mock/MockMunchNFT.sol#L9


 - [ ] ID-93
function:[BaseConfigStorageUpgradeable._authorizeUpgrade(address)](src/config/BaseConfigStorageUpgradeable.sol#L14)is empty 

src/config/BaseConfigStorageUpgradeable.sol#L14


 - [ ] ID-94
function:[ERC721URIStorageUpgradeable.__ERC721URIStorage_init()](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L37-L38)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L37-L38


 - [ ] ID-95
function:[ERC721EnumerableUpgradeable.__ERC721Enumerable_init_unchained()](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L52-L53)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L52-L53


 - [ ] ID-96
function:[ERC721PausableUpgradeable.__ERC721Pausable_init_unchained()](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721PausableUpgradeable.sol#L28-L29)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721PausableUpgradeable.sol#L28-L29


 - [ ] ID-97
function:[ERC721URIStorageUpgradeable.__ERC721URIStorage_init_unchained()](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L40-L41)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L40-L41


 - [ ] ID-98
function:[ContextUpgradeable.__Context_init()](lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#L18-L19)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#L18-L19


 - [ ] ID-99
function:[MockBlast.configureContract(address,YieldMode,GasMode,address)](src/mock/MockBlast.sol#L11)is empty 

src/mock/MockBlast.sol#L11


 - [ ] ID-100
function:[OldMunchNFT.configureContract(address)](src/tokens/OldMunchNFT.sol#L93-L98)is empty 

src/tokens/OldMunchNFT.sol#L93-L98


 - [ ] ID-101
function:[ERC165Upgradeable.__ERC165_init()](lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#L22-L23)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#L22-L23


 - [ ] ID-102
function:[OldMunchNFT.setURI(uint256,string)](src/tokens/OldMunchNFT.sol#L145-L160)is empty 

src/tokens/OldMunchNFT.sol#L145-L160


 - [ ] ID-103
function:[ERC721EnumerableUpgradeable.__ERC721Enumerable_init()](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L49-L50)is empty 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L49-L50


 - [ ] ID-104
function:[MockBlast.configure(YieldMode,GasMode,address)](src/mock/MockBlast.sol#L12)is empty 

src/mock/MockBlast.sol#L12


 - [ ] ID-105
function:[MockMunchNFT.burn(uint256)](src/mock/MockMunchNFT.sol#L7)is empty 

src/mock/MockMunchNFT.sol#L7


## centralized-risk-low
Impact: Low
Confidence: High
 - [ ] ID-106
	- [MigrationManager.loadMigrationSnapshot(address[],IMigrationManager.MigrationSnapshotData[])](src/managers/MigrationManager.sol#L78-L114)

src/managers/MigrationManager.sol#L78-L114


 - [ ] ID-107
	- [LockManager.configureLockdrop(ILockManager.Lockdrop)](src/managers/LockManager.sol#L98-L112)

src/managers/LockManager.sol#L98-L112


 - [ ] ID-108
	- [LockManager.configureToken(address,ILockManager.ConfiguredToken)](src/managers/LockManager.sol#L115-L127)

src/managers/LockManager.sol#L115-L127


 - [ ] ID-109
	- [MigrationManager.burnNFTs(address,uint32)](src/managers/MigrationManager.sol#L148-L182)

src/managers/MigrationManager.sol#L148-L182


 - [ ] ID-110
	- [LockManager.setLockDuration(uint256)](src/managers/LockManager.sol#L245-L272)

src/managers/LockManager.sol#L245-L272


 - [ ] ID-111
	- [MigrationManager.sealData()](src/managers/MigrationManager.sol#L116-L120)

src/managers/MigrationManager.sol#L116-L120


## pess-event-setter
Impact: Low
Confidence: Medium
 - [ ] ID-112
Setter function [MockBlast.claimAllGas(address,address)](src/mock/MockBlast.sol#L64-L67) does not emit an event

src/mock/MockBlast.sol#L64-L67


 - [ ] ID-113
Setter function [ConfigStorage.setUintArray(StorageKey,uint256[],bool)](src/config/ConfigStorage.sol#L69-L96) does not emit an event

src/config/ConfigStorage.sol#L69-L96


 - [ ] ID-114
Setter function [MunchNFT.transferFrom(address,address,uint256)](src/tokens/MunchNFT.sol#L39-L46) does not emit an event

src/tokens/MunchNFT.sol#L39-L46


 - [ ] ID-115
Setter function [MockBlast.configureVoidGasOnBehalf(address)](src/mock/MockBlast.sol#L42-L44) does not emit an event

src/mock/MockBlast.sol#L42-L44


 - [ ] ID-116
Setter function [AccessControl.onlyRole(bytes32)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L63-L66) does not emit an event

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L63-L66


 - [ ] ID-117
Setter function [OldMunchNFT.revealNFT(address,string,bytes32,bytes,MunchablesCommonLib.Realm,MunchablesCommonLib.Rarity,uint16)](src/tokens/OldMunchNFT.sol#L110-L140) does not emit an event

src/tokens/OldMunchNFT.sol#L110-L140


 - [ ] ID-118
Setter function [MockSnuggeryManager.setGlobalTotalChonk(uint256)](src/mock/MockSnuggeryManager.sol#L24-L26) does not emit an event

src/mock/MockSnuggeryManager.sol#L24-L26


 - [ ] ID-119
Setter function [ConfigStorage.addNotifiableAddresses(address[])](src/config/ConfigStorage.sol#L316-L322) does not emit an event

src/config/ConfigStorage.sol#L316-L322


 - [ ] ID-120
Setter function [MockBlast.configureVoidYield()](src/mock/MockBlast.sol#L27-L29) does not emit an event

src/mock/MockBlast.sol#L27-L29


 - [ ] ID-121
Setter function [ConfigStorage.manualNotify(uint8,uint8)](src/config/ConfigStorage.sol#L345-L350) does not emit an event

src/config/ConfigStorage.sol#L345-L350


 - [ ] ID-122
Setter function [MockNFTOverlord.addReveal(address,uint16)](src/mock/MockNFTOverlord.sol#L64-L67) does not emit an event

src/mock/MockNFTOverlord.sol#L64-L67


 - [ ] ID-123
Setter function [NFTOverlord._populateDefaultProbabilities()](src/overlords/NFTOverlord.sol#L422-L453) does not emit an event

src/overlords/NFTOverlord.sol#L422-L453


 - [ ] ID-124
Setter function [ConfigStorage.setAddressArray(StorageKey,address[],bool)](src/config/ConfigStorage.sol#L254-L281) does not emit an event

src/config/ConfigStorage.sol#L254-L281


 - [ ] ID-125
Setter function [FundTreasuryDistributor._reconfigure()](src/distributors/FundTreasuryDistributor.sol#L21-L24) does not emit an event

src/distributors/FundTreasuryDistributor.sol#L21-L24


 - [ ] ID-126
Setter function [MockNFTOverlord._reconfigure()](src/mock/MockNFTOverlord.sol#L33-L57) does not emit an event

src/mock/MockNFTOverlord.sol#L33-L57


 - [ ] ID-127
Setter function [RNGProxyAPI3.provideRandom(bytes32,bytes)](src/rng/RNGProxyAPI3.sol#L51-L58) does not emit an event

src/rng/RNGProxyAPI3.sol#L51-L58


 - [ ] ID-128
Setter function [MockBlast.claimGasAtMinClaimRate(address,address,uint256)](src/mock/MockBlast.sol#L68-L75) does not emit an event

src/mock/MockBlast.sol#L68-L75


 - [ ] ID-129
Setter function [MockNFTOverlord._populateDefaultProbabilities()](src/mock/MockNFTOverlord.sol#L384-L415) does not emit an event

src/mock/MockNFTOverlord.sol#L384-L415


 - [ ] ID-130
Setter function [AccessControl._checkRole(bytes32)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L86-L88) does not emit an event

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L86-L88


 - [ ] ID-131
Setter function [MunchToken.mint(address,uint256)](src/tokens/MunchToken.sol#L28-L33) does not emit an event

src/tokens/MunchToken.sol#L28-L33


 - [ ] ID-132
Setter function [OldMunchNFT.configureContract(address)](src/tokens/OldMunchNFT.sol#L93-L98) does not emit an event

src/tokens/OldMunchNFT.sol#L93-L98


 - [ ] ID-133
Setter function [ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool)](src/config/ConfigStorage.sol#L111-L139) does not emit an event

src/config/ConfigStorage.sol#L111-L139


 - [ ] ID-134
Setter function [NFTOverlord.mintFromPrimordial(address)](src/overlords/NFTOverlord.sol#L134-L150) does not emit an event

src/overlords/NFTOverlord.sol#L134-L150


 - [ ] ID-135
Setter function [OldMunchNFT.unpause()](src/tokens/OldMunchNFT.sol#L83-L85) does not emit an event

src/tokens/OldMunchNFT.sol#L83-L85


 - [ ] ID-136
Setter function [ConfigStorage.setSmallInt(StorageKey,int16,bool)](src/config/ConfigStorage.sol#L155-L162) does not emit an event

src/config/ConfigStorage.sol#L155-L162


 - [ ] ID-137
Setter function [MockNFTOverlord._populateDefaultRealmLookup()](src/mock/MockNFTOverlord.sol#L417-L425) does not emit an event

src/mock/MockNFTOverlord.sol#L417-L425


 - [ ] ID-138
Setter function [MunchNFT.configUpdated()](src/tokens/MunchNFT.sol#L62-L64) does not emit an event

src/tokens/MunchNFT.sol#L62-L64


 - [ ] ID-139
Setter function [ConfigStorage.getRole(Role)](src/config/ConfigStorage.sol#L39-L41) does not emit an event

src/config/ConfigStorage.sol#L39-L41


 - [ ] ID-140
Setter function [MockBlast.configureClaimableGasOnBehalf(address)](src/mock/MockBlast.sol#L36-L38) does not emit an event

src/mock/MockBlast.sol#L36-L38


 - [ ] ID-141
Setter function [NFTOverlord.configUpdated()](src/overlords/NFTOverlord.sol#L60-L62) does not emit an event

src/overlords/NFTOverlord.sol#L60-L62


 - [ ] ID-142
Setter function [MockBlast.configureGovernorOnBehalf(address,address)](src/mock/MockBlast.sol#L48-L50) does not emit an event

src/mock/MockBlast.sol#L48-L50


 - [ ] ID-143
Setter function [ConfigStorage.setAddress(StorageKey,address,bool)](src/config/ConfigStorage.sol#L228-L235) does not emit an event

src/config/ConfigStorage.sol#L228-L235


 - [ ] ID-144
Setter function [ConfigStorage.constructor()](src/config/ConfigStorage.sol#L25) does not emit an event

src/config/ConfigStorage.sol#L25


 - [ ] ID-145
Setter function [ConfigStorage.setAddresses(StorageKey[],address[],bool)](src/config/ConfigStorage.sol#L238-L247) does not emit an event

src/config/ConfigStorage.sol#L238-L247


 - [ ] ID-146
Setter function [MockRNGProxy._reconfigure()](src/mock/MockRNGProxy.sol#L15-L19) does not emit an event

src/mock/MockRNGProxy.sol#L15-L19


 - [ ] ID-147
Setter function [MockNFTOverlord.mintFromPrimordial(address)](src/mock/MockNFTOverlord.sol#L129-L137) does not emit an event

src/mock/MockNFTOverlord.sol#L129-L137


 - [ ] ID-148
Setter function [MockSnuggeryManager.setSnuggeryForTest(address,MunchablesCommonLib.SnuggeryNFT[])](src/mock/MockSnuggeryManager.sol#L11-L22) does not emit an event

src/mock/MockSnuggeryManager.sol#L11-L22


 - [ ] ID-149
Setter function [TestERC20Token.constructor()](src/tokens/TestERC20Token.sol#L11-L14) does not emit an event

src/tokens/TestERC20Token.sol#L11-L14


 - [ ] ID-150
Setter function [MockMunchadexManager.setMunchadexUniqueForTest(address,bytes32,uint256)](src/mock/MockMunchadexManager.sol#L35-L42) does not emit an event

src/mock/MockMunchadexManager.sol#L35-L42


 - [ ] ID-151
Setter function [MockMunchadexManager.setMunchadexNumInRarityForTest(address,MunchablesCommonLib.Rarity,uint256)](src/mock/MockMunchadexManager.sol#L26-L33) does not emit an event

src/mock/MockMunchadexManager.sol#L26-L33


 - [ ] ID-152
Setter function [MockBlast.configureClaimableYieldOnBehalf(address)](src/mock/MockBlast.sol#L18-L20) does not emit an event

src/mock/MockBlast.sol#L18-L20


 - [ ] ID-153
Setter function [RNGProxyAPI3.requestRandom(address,bytes4,uint256)](src/rng/RNGProxyAPI3.sol#L30-L47) does not emit an event

src/rng/RNGProxyAPI3.sol#L30-L47


 - [ ] ID-154
Setter function [MockMunchadexManager.setMunchadexForTest(address,uint256)](src/mock/MockMunchadexManager.sol#L9-L15) does not emit an event

src/mock/MockMunchadexManager.sol#L9-L15


 - [ ] ID-155
Setter function [AccessControl.grantRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L122-L124) does not emit an event

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L122-L124


 - [ ] ID-156
Setter function [OldMunchNFT.burn(uint256)](src/tokens/OldMunchNFT.sol#L87-L89) does not emit an event

src/tokens/OldMunchNFT.sol#L87-L89


 - [ ] ID-157
Setter function [MockNFTOverlord.startReveal()](src/mock/MockNFTOverlord.sol#L70-L72) does not emit an event

src/mock/MockNFTOverlord.sol#L70-L72


 - [ ] ID-158
Setter function [MockBlast.configureVoidYieldOnBehalf(address)](src/mock/MockBlast.sol#L30-L32) does not emit an event

src/mock/MockBlast.sol#L30-L32


 - [ ] ID-159
Setter function [ConfigStorage.setSmallIntArray(StorageKey,int16[],bool)](src/config/ConfigStorage.sol#L169-L197) does not emit an event

src/config/ConfigStorage.sol#L169-L197


 - [ ] ID-160
Setter function [ConfigStorage.setRole(Role,address,address)](src/config/ConfigStorage.sol#L27-L33) does not emit an event

src/config/ConfigStorage.sol#L27-L33


 - [ ] ID-161
Setter function [ConfigStorage.setBool(StorageKey,bool,bool)](src/config/ConfigStorage.sol#L213-L220) does not emit an event

src/config/ConfigStorage.sol#L213-L220


 - [ ] ID-162
Setter function [MockMunchadexManager.setMunchadexNumInRealmForTest(address,MunchablesCommonLib.Realm,uint256)](src/mock/MockMunchadexManager.sol#L17-L24) does not emit an event

src/mock/MockMunchadexManager.sol#L17-L24


 - [ ] ID-163
Setter function [MockNFTOverlord.configUpdated()](src/mock/MockNFTOverlord.sol#L59-L61) does not emit an event

src/mock/MockNFTOverlord.sol#L59-L61


 - [ ] ID-164
Setter function [MockConfigNotifiable.configUpdated()](src/mock/MockConfigNotifiable.sol#L17-L20) does not emit an event

src/mock/MockConfigNotifiable.sol#L17-L20


 - [ ] ID-165
Setter function [ConfigStorage.addNotifiableAddress(address)](src/config/ConfigStorage.sol#L312-L314) does not emit an event

src/config/ConfigStorage.sol#L312-L314


 - [ ] ID-166
Setter function [MockLockManager.setLockedTokenForTest(address,address,ILockManager.LockedToken)](src/mock/MockLockManager.sol#L9-L15) does not emit an event

src/mock/MockLockManager.sol#L9-L15


 - [ ] ID-167
Setter function [MockMigrationManager.setUserMigrationCompletedDataForTest(address,IMigrationManager.MigrationTotals,bool)](src/mock/MockMigrationManager.sol#L9-L16) does not emit an event

src/mock/MockMigrationManager.sol#L9-L16


 - [ ] ID-168
Setter function [MockBlast.configureClaimableGas()](src/mock/MockBlast.sol#L33-L35) does not emit an event

src/mock/MockBlast.sol#L33-L35


 - [ ] ID-169
Setter function [OldMunchNFT.setMigrationManager(address)](src/tokens/OldMunchNFT.sol#L100-L104) does not emit an event

src/tokens/OldMunchNFT.sol#L100-L104


 - [ ] ID-170
Setter function [BonusManager._reconfigure()](src/managers/BonusManager.sol#L40-L85) does not emit an event

src/managers/BonusManager.sol#L40-L85


 - [ ] ID-171
Setter function [NFTOverlord.addReveal(address,uint16)](src/overlords/NFTOverlord.sol#L65-L71) does not emit an event

src/overlords/NFTOverlord.sol#L65-L71


 - [ ] ID-172
Setter function [TestERC20Token.claim(address,uint256)](src/tokens/TestERC20Token.sol#L22-L25) does not emit an event

src/tokens/TestERC20Token.sol#L22-L25


 - [ ] ID-173
Setter function [BonusManager.configUpdated()](src/managers/BonusManager.sol#L87-L89) does not emit an event

src/managers/BonusManager.sol#L87-L89


 - [ ] ID-174
Setter function [NFTOverlord._populateDefaultRealmLookup()](src/overlords/NFTOverlord.sol#L455-L463) does not emit an event

src/overlords/NFTOverlord.sol#L455-L463


 - [ ] ID-175
Setter function [MockBlast.claimGas(address,address,uint256,uint256)](src/mock/MockBlast.sol#L80-L88) does not emit an event

src/mock/MockBlast.sol#L80-L88


 - [ ] ID-176
Setter function [AccessControl.revokeRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L137-L139) does not emit an event

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L137-L139


 - [ ] ID-177
Setter function [ConfigStorage.manualNotifyAddress(address)](src/config/ConfigStorage.sol#L352-L354) does not emit an event

src/config/ConfigStorage.sol#L352-L354


 - [ ] ID-178
Setter function [ConfigStorage.setBytes32(StorageKey,bytes32,bool)](src/config/ConfigStorage.sol#L297-L304) does not emit an event

src/config/ConfigStorage.sol#L297-L304


 - [ ] ID-179
Setter function [MockBlast.configureClaimableYield()](src/mock/MockBlast.sol#L15-L17) does not emit an event

src/mock/MockBlast.sol#L15-L17


 - [ ] ID-180
Setter function [MockBlast.configureAutomaticYield()](src/mock/MockBlast.sol#L21-L23) does not emit an event

src/mock/MockBlast.sol#L21-L23


 - [ ] ID-181
Setter function [MockBlast.configureVoidGas()](src/mock/MockBlast.sol#L39-L41) does not emit an event

src/mock/MockBlast.sol#L39-L41


 - [ ] ID-182
Setter function [MockAccountManager.giveSchnibbles(address,uint256)](src/mock/MockAccountManager.sol#L12-L17) does not emit an event

src/mock/MockAccountManager.sol#L12-L17


 - [ ] ID-183
Setter function [ConfigStorage.removeNotifiableAddress(address)](src/config/ConfigStorage.sol#L325-L335) does not emit an event

src/config/ConfigStorage.sol#L325-L335


 - [ ] ID-184
Setter function [MunchNFT._reconfigure()](src/tokens/MunchNFT.sol#L48-L60) does not emit an event

src/tokens/MunchNFT.sol#L48-L60


 - [ ] ID-185
Setter function [MockBlast.claimMaxGas(address,address)](src/mock/MockBlast.sol#L76-L79) does not emit an event

src/mock/MockBlast.sol#L76-L79


 - [ ] ID-186
Setter function [MunchNFT.setTokenURI(uint256,string)](src/tokens/MunchNFT.sol#L79-L84) does not emit an event

src/tokens/MunchNFT.sol#L79-L84


 - [ ] ID-187
Setter function [OldMunchNFT.pause()](src/tokens/OldMunchNFT.sol#L78-L80) does not emit an event

src/tokens/OldMunchNFT.sol#L78-L80


 - [ ] ID-188
Setter function [MockClaimManager.givePoints(address,uint256)](src/mock/MockClaimManager.sol#L12-L14) does not emit an event

src/mock/MockClaimManager.sol#L12-L14


 - [ ] ID-189
Setter function [FundTreasuryDistributor.configUpdated()](src/distributors/FundTreasuryDistributor.sol#L17-L19) does not emit an event

src/distributors/FundTreasuryDistributor.sol#L17-L19


 - [ ] ID-190
Setter function [OldMunchNFT.onlyMigrationManager()](src/tokens/OldMunchNFT.sol#L39-L42) does not emit an event

src/tokens/OldMunchNFT.sol#L39-L42


 - [ ] ID-191
Setter function [MockNFTAttributesManagerV1.setImmutableAttributesForTest(uint256,MunchablesCommonLib.NFTImmutableAttributes)](src/mock/MockNFTAttributeManagerV1.sol#L12-L17) does not emit an event

src/mock/MockNFTAttributeManagerV1.sol#L12-L17


 - [ ] ID-192
Setter function [MockBlast.configureGovernor(address)](src/mock/MockBlast.sol#L45-L47) does not emit an event

src/mock/MockBlast.sol#L45-L47


 - [ ] ID-193
Setter function [MockBlast.configureAutomaticYieldOnBehalf(address)](src/mock/MockBlast.sol#L24-L26) does not emit an event

src/mock/MockBlast.sol#L24-L26


 - [ ] ID-194
Setter function [MunchToken.configUpdated()](src/tokens/MunchToken.sol#L24-L26) does not emit an event

src/tokens/MunchToken.sol#L24-L26


 - [ ] ID-195
Setter function [OldMunchNFT.initialize(address,address,address,address,address)](src/tokens/OldMunchNFT.sol#L53-L75) does not emit an event

src/tokens/OldMunchNFT.sol#L53-L75


 - [ ] ID-196
Setter function [MockNFTAttributesManagerV1.setAttributesForTest(uint256,MunchablesCommonLib.NFTAttributes)](src/mock/MockNFTAttributeManagerV1.sol#L19-L24) does not emit an event

src/mock/MockNFTAttributeManagerV1.sol#L19-L24


 - [ ] ID-197
Setter function [ConfigStorage.setUniversalRole(Role,address)](src/config/ConfigStorage.sol#L35-L37) does not emit an event

src/config/ConfigStorage.sol#L35-L37


 - [ ] ID-198
Setter function [MockBlast.claimYield(address,address,uint256)](src/mock/MockBlast.sol#L53-L56) does not emit an event

src/mock/MockBlast.sol#L53-L56


 - [ ] ID-199
Setter function [MockBlast.claimAllYield(address,address)](src/mock/MockBlast.sol#L58-L61) does not emit an event

src/mock/MockBlast.sol#L58-L61


 - [ ] ID-200
Setter function [ConfigStorage.setUint(StorageKey,uint256,bool)](src/config/ConfigStorage.sol#L55-L62) does not emit an event

src/config/ConfigStorage.sol#L55-L62


 - [ ] ID-201
Setter function [TestERC20Token.configure(YieldMode)](src/tokens/TestERC20Token.sol#L17-L20) does not emit an event

src/tokens/TestERC20Token.sol#L17-L20


 - [ ] ID-202
Setter function [AccessControl.renounceRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L157-L163) does not emit an event

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L157-L163


 - [ ] ID-203
Setter function [MunchNFT.mint(address)](src/tokens/MunchNFT.sol#L66-L77) does not emit an event

src/tokens/MunchNFT.sol#L66-L77


 - [ ] ID-204
Setter function [ProxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy,address,bytes)](lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L38-L44) does not emit an event

lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L38-L44


 - [ ] ID-205
Setter function [NFTOverlord._reconfigure()](src/overlords/NFTOverlord.sol#L34-L58) does not emit an event

src/overlords/NFTOverlord.sol#L34-L58


 - [ ] ID-206
Setter function [OldMunchNFT.setURI(uint256,string)](src/tokens/OldMunchNFT.sol#L145-L160) does not emit an event

src/tokens/OldMunchNFT.sol#L145-L160


## missing-zero-check
Impact: Low
Confidence: Medium
 - [ ] ID-207
variable lacks a zero-check on 		- [BonusManager.getFeedBonus(address,uint256)](src/managers/BonusManager.sol#L91-L123)

src/managers/BonusManager.sol#L91-L123


 - [ ] ID-208
variable lacks a zero-check on 		- [ClaimManager.burnNFTsForPoints(address,uint8[])](src/managers/ClaimManager.sol#L113-L129)

src/managers/ClaimManager.sol#L113-L129


 - [ ] ID-209
variable lacks a zero-check on 		- [ERC1967Proxy.constructor(address,bytes)](lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol#L26-L28)

lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol#L26-L28


 - [ ] ID-210
variable lacks a zero-check on 		- [MockSnuggeryManager.forceClaimPoints(address)](src/mock/MockSnuggeryManager.sol#L32-L34)

src/mock/MockSnuggeryManager.sol#L32-L34


 - [ ] ID-211
variable lacks a zero-check on 		- [SnuggeryManager.getSnuggery(address,uint256)](src/managers/SnuggeryManager.sol#L254-L282)

src/managers/SnuggeryManager.sol#L254-L282


 - [ ] ID-212
variable lacks a zero-check on 		- [AccountManager.getDailySchnibbles(address)](src/managers/AccountManager.sol#L363-L370)

src/managers/AccountManager.sol#L363-L370


 - [ ] ID-213
variable lacks a zero-check on 		- [MockMigrationManager.burnUnrevealedForPoints(address,uint256)](src/mock/MockMigrationManager.sol#L40-L46)

src/mock/MockMigrationManager.sol#L40-L46


 - [ ] ID-214
variable lacks a zero-check on 		- [MockRewardsManager.claimAllGasForContract(address)](src/mock/MockRewardsManager.sol#L10-L19)

src/mock/MockRewardsManager.sol#L10-L19


 - [ ] ID-215
variable lacks a zero-check on 		- [MockPrimordialManager.callMintFromPrimordialForTest(address)](src/mock/MockPrimordialManager.sol#L9-L11)

src/mock/MockPrimordialManager.sol#L9-L11


 - [ ] ID-216
variable lacks a zero-check on 		- [ClaimManager.getPoints(address)](src/managers/ClaimManager.sol#L196-L200)

src/managers/ClaimManager.sol#L196-L200


 - [ ] ID-217
variable lacks a zero-check on 		- [MockNFTOverlord.getUnrevealedNFTs(address)](src/mock/MockNFTOverlord.sol#L287-L292)

src/mock/MockNFTOverlord.sol#L287-L292


 - [ ] ID-218
variable lacks a zero-check on 		- [MockLockManager.setLockedTokenForTest(address,address,ILockManager.LockedToken)](src/mock/MockLockManager.sol#L9-L15)

src/mock/MockLockManager.sol#L9-L15


 - [ ] ID-219
variable lacks a zero-check on 		- [NFTOverlord.addReveal(address,uint16)](src/overlords/NFTOverlord.sol#L65-L71)

src/overlords/NFTOverlord.sol#L65-L71


 - [ ] ID-220
variable lacks a zero-check on 		- [RewardsManager.reassignBlastGovernor(address)](src/managers/RewardsManager.sol#L159-L167)

src/managers/RewardsManager.sol#L159-L167


 - [ ] ID-221
variable lacks a zero-check on 		- [MunchNFT.transferFrom(address,address,uint256)](src/tokens/MunchNFT.sol#L39-L46)

src/tokens/MunchNFT.sol#L39-L46


 - [ ] ID-222
variable lacks a zero-check on 		- [MockNFTOverlord.mintFromPrimordial(address)](src/mock/MockNFTOverlord.sol#L129-L137)

src/mock/MockNFTOverlord.sol#L129-L137


 - [ ] ID-223
variable lacks a zero-check on 		- [BonusManager.getHarvestBonus(address)](src/managers/BonusManager.sol#L125-L140)

src/managers/BonusManager.sol#L125-L140


 - [ ] ID-224
variable lacks a zero-check on 		- [NFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[])](src/overlords/NFTOverlord.sol#L189-L215)

src/overlords/NFTOverlord.sol#L189-L215


 - [ ] ID-225
variable lacks a zero-check on 		- [MockMigrationManager.setUserMigrationCompletedDataForTest(address,IMigrationManager.MigrationTotals,bool)](src/mock/MockMigrationManager.sol#L9-L16)

src/mock/MockMigrationManager.sol#L9-L16


 - [ ] ID-226
variable lacks a zero-check on 		- [AccountManager.removeSprayProposal(address)](src/managers/AccountManager.sol#L213-L228)

src/managers/AccountManager.sol#L213-L228


 - [ ] ID-227
variable lacks a zero-check on 		- [RrpRequesterV0.constructor(address)](node_modules/@api3/airnode-protocol/contracts/rrp/requesters/RrpRequesterV0.sol#L23-L26)

node_modules/@api3/airnode-protocol/contracts/rrp/requesters/RrpRequesterV0.sol#L23-L26


 - [ ] ID-228
variable lacks a zero-check on 		- [MockClaimManager.givePoints(address,uint256)](src/mock/MockClaimManager.sol#L12-L14)

src/mock/MockClaimManager.sol#L12-L14


 - [ ] ID-229
variable lacks a zero-check on 		- [MockMigrationManager.burnNFTsForPoints(address,uint8[])](src/mock/MockMigrationManager.sol#L32-L38)

src/mock/MockMigrationManager.sol#L32-L38


 - [ ] ID-230
variable lacks a zero-check on 		- [NFTOverlord.getUnrevealedNFTs(address)](src/overlords/NFTOverlord.sol#L316-L321)

src/overlords/NFTOverlord.sol#L316-L321


 - [ ] ID-231
variable lacks a zero-check on 		- [MockNFTOverlord.addReveal(address,uint16)](src/mock/MockNFTOverlord.sol#L64-L67)

src/mock/MockNFTOverlord.sol#L64-L67


 - [ ] ID-232
variable lacks a zero-check on 		- [AccountManager.updatePlayer(address,MunchablesCommonLib.Player)](src/managers/AccountManager.sol#L284-L295)

src/managers/AccountManager.sol#L284-L295


 - [ ] ID-233
variable lacks a zero-check on 		- [BonusManager.getPetBonus(address)](src/managers/BonusManager.sol#L142-L148)

src/managers/BonusManager.sol#L142-L148


 - [ ] ID-234
variable lacks a zero-check on 		- [ClaimManager.burnUnrevealedForPoints(address,uint256)](src/managers/ClaimManager.sol#L131-L144)

src/managers/ClaimManager.sol#L131-L144


 - [ ] ID-235
variable lacks a zero-check on 		- [MockMigrationManager.callMintForMigrationForTest(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[])](src/mock/MockMigrationManager.sol#L18-L30)

src/mock/MockMigrationManager.sol#L18-L30


 - [ ] ID-236
variable lacks a zero-check on 		- [MockAccountManager.giveSchnibbles(address,uint256)](src/mock/MockAccountManager.sol#L12-L17)

src/mock/MockAccountManager.sol#L12-L17


 - [ ] ID-237
variable lacks a zero-check on 		- [BaseBlastManager.claimERC20Yield(address,uint256)](src/managers/BaseBlastManager.sol#L109-L117)

src/managers/BaseBlastManager.sol#L109-L117


 - [ ] ID-238
variable lacks a zero-check on 		- [MockNFTOverlord.revealFromPrimordial(uint256,bytes)](src/mock/MockNFTOverlord.sol#L141-L162)

src/mock/MockNFTOverlord.sol#L141-L162


 - [ ] ID-239
variable lacks a zero-check on 		- [RNGProxyAPI3.constructor(address,address,address,address,address,bytes32)](src/rng/RNGProxyAPI3.sol#L16-L28)

src/rng/RNGProxyAPI3.sol#L16-L28


 - [ ] ID-240
variable lacks a zero-check on 		- [MockSnuggeryManager.spendPoints(address,uint256)](src/mock/MockSnuggeryManager.sol#L28-L30)

src/mock/MockSnuggeryManager.sol#L28-L30


 - [ ] ID-241
variable lacks a zero-check on 		- [MockLockManager.callAddRevealForTest(address,uint8)](src/mock/MockLockManager.sol#L17-L19)

src/mock/MockLockManager.sol#L17-L19


 - [ ] ID-242
variable lacks a zero-check on 		- [MockNFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[])](src/mock/MockNFTOverlord.sol#L165-L185)

src/mock/MockNFTOverlord.sol#L165-L185


 - [ ] ID-243
variable lacks a zero-check on 		- [RewardsManager.isGovernorOfContract(address)](src/managers/RewardsManager.sol#L169-L173)

src/managers/RewardsManager.sol#L169-L173


 - [ ] ID-244
variable lacks a zero-check on 		- [MigrationManager.loadMigrationSnapshot(address[],IMigrationManager.MigrationSnapshotData[])](src/managers/MigrationManager.sol#L78-L114)

src/managers/MigrationManager.sol#L78-L114


 - [ ] ID-245
variable lacks a zero-check on 		- [ProxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy,address,bytes)](lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L38-L44)

lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L38-L44


 - [ ] ID-246
variable lacks a zero-check on 		- [NFTOverlord.mintFromPrimordial(address)](src/overlords/NFTOverlord.sol#L134-L150)

src/overlords/NFTOverlord.sol#L134-L150


 - [ ] ID-247
variable lacks a zero-check on 		- [OldMunchNFT.setMigrationManager(address)](src/tokens/OldMunchNFT.sol#L100-L104)

src/tokens/OldMunchNFT.sol#L100-L104


 - [ ] ID-248
variable lacks a zero-check on 		- [MunchNFT.mint(address)](src/tokens/MunchNFT.sol#L66-L77)

src/tokens/MunchNFT.sol#L66-L77


## centralized-risk-informational
Impact: Informational
Confidence: High
 - [ ] ID-249
	- [ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool)](src/config/ConfigStorage.sol#L111-L139)

src/config/ConfigStorage.sol#L111-L139


 - [ ] ID-250
	- [OldMunchNFT.safeMint(address,string)](src/tokens/OldMunchNFT.sol#L184-L193)

src/tokens/OldMunchNFT.sol#L184-L193


 - [ ] ID-251
	- [ConfigStorage.addNotifiableAddress(address)](src/config/ConfigStorage.sol#L312-L314)

src/config/ConfigStorage.sol#L312-L314


 - [ ] ID-252
	- [NFTOverlord.munchableFed(uint256,address)](src/overlords/NFTOverlord.sol#L279-L313)

src/overlords/NFTOverlord.sol#L279-L313


 - [ ] ID-253
	- [ConfigStorage.setUintArray(StorageKey,uint256[],bool)](src/config/ConfigStorage.sol#L69-L96)

src/config/ConfigStorage.sol#L69-L96


 - [ ] ID-254
	- [MigrationManager.burnRemainingPurchasedNFTs(address,uint32)](src/managers/MigrationManager.sol#L184-L219)

src/managers/MigrationManager.sol#L184-L219


 - [ ] ID-255
	- [ConfigStorage.setAddressArray(StorageKey,address[],bool)](src/config/ConfigStorage.sol#L254-L281)

src/config/ConfigStorage.sol#L254-L281


 - [ ] ID-256
	- [NFTOverlord.levelUp(uint256,bytes)](src/overlords/NFTOverlord.sol#L218-L276)

src/overlords/NFTOverlord.sol#L218-L276


 - [ ] ID-257
	- [ClaimManager.burnUnrevealedForPoints(address,uint256)](src/managers/ClaimManager.sol#L131-L144)

src/managers/ClaimManager.sol#L131-L144


 - [ ] ID-258
	- [ConfigStorage.setRole(Role,address,address)](src/config/ConfigStorage.sol#L27-L33)

src/config/ConfigStorage.sol#L27-L33


 - [ ] ID-259
	- [ConfigStorage.setUniversalRole(Role,address)](src/config/ConfigStorage.sol#L35-L37)

src/config/ConfigStorage.sol#L35-L37


 - [ ] ID-260
	- [ConfigStorage.setBytes32(StorageKey,bytes32,bool)](src/config/ConfigStorage.sol#L297-L304)

src/config/ConfigStorage.sol#L297-L304


 - [ ] ID-261
	- [LockManager.disapproveUSDPrice(uint256)](src/managers/LockManager.sol#L210-L242)

src/managers/LockManager.sol#L210-L242


 - [ ] ID-262
	- [PrimordialManager.claimPrimordial()](src/managers/PrimordialManager.sol#L60-L70)

src/managers/PrimordialManager.sol#L60-L70


 - [ ] ID-263
	- [NFTAttributesManagerV1.setGameAttributes(uint256,MunchablesCommonLib.NFTGameAttribute[])](src/managers/NFTAttributeManagerV1.sol#L84-L104)

src/managers/NFTAttributeManagerV1.sol#L84-L104


 - [ ] ID-264
	- [NFTOverlord.mintFromPrimordial(address)](src/overlords/NFTOverlord.sol#L134-L150)

src/overlords/NFTOverlord.sol#L134-L150


 - [ ] ID-265
	- [ConfigStorage.setAddresses(StorageKey[],address[],bool)](src/config/ConfigStorage.sol#L238-L247)

src/config/ConfigStorage.sol#L238-L247


 - [ ] ID-266
	- [ConfigStorage.setSmallInt(StorageKey,int16,bool)](src/config/ConfigStorage.sol#L155-L162)

src/config/ConfigStorage.sol#L155-L162


 - [ ] ID-267
	- [LockManager.approveUSDPrice(uint256)](src/managers/LockManager.sol#L177-L207)

src/managers/LockManager.sol#L177-L207


 - [ ] ID-268
	- [ConfigStorage.removeNotifiableAddress(address)](src/config/ConfigStorage.sol#L325-L335)

src/config/ConfigStorage.sol#L325-L335


 - [ ] ID-269
	- [ClaimManager.burnNFTsForPoints(address,uint8[])](src/managers/ClaimManager.sol#L113-L129)

src/managers/ClaimManager.sol#L113-L129


 - [ ] ID-270
	- [ConfigStorage.setBool(StorageKey,bool,bool)](src/config/ConfigStorage.sol#L213-L220)

src/config/ConfigStorage.sol#L213-L220


 - [ ] ID-271
	- [ConfigStorage.setAddress(StorageKey,address,bool)](src/config/ConfigStorage.sol#L228-L235)

src/config/ConfigStorage.sol#L228-L235


 - [ ] ID-272
	- [ConfigStorage.addNotifiableAddresses(address[])](src/config/ConfigStorage.sol#L316-L322)

src/config/ConfigStorage.sol#L316-L322


 - [ ] ID-273
	- [AccountManager.removeSprayProposal(address)](src/managers/AccountManager.sol#L213-L228)

src/managers/AccountManager.sol#L213-L228


 - [ ] ID-274
	- [ClaimManager.convertPointsToTokens(uint256)](src/managers/ClaimManager.sol#L173-L190)

src/managers/ClaimManager.sol#L173-L190


 - [ ] ID-275
	- [AccountManager.spraySchnibblesPropose(address[],uint256[])](src/managers/AccountManager.sol#L135-L177)

src/managers/AccountManager.sol#L135-L177


 - [ ] ID-276
	- [BaseRNGProxy.requestRandom(address,bytes4,uint256)](src/rng/BaseRNGProxy.sol#L21-L32)

src/rng/BaseRNGProxy.sol#L21-L32


 - [ ] ID-277
	- [AccountManager.addSubAccount(address)](src/managers/AccountManager.sol#L231-L251)

src/managers/AccountManager.sol#L231-L251


 - [ ] ID-278
	- [MigrationManager.loadUnrevealedSnapshot(address[],uint16[])](src/managers/MigrationManager.sol#L122-L133)

src/managers/MigrationManager.sol#L122-L133


 - [ ] ID-279
	- [AccountManager.execSprayProposal(address)](src/managers/AccountManager.sol#L180-L210)

src/managers/AccountManager.sol#L180-L210


 - [ ] ID-280
	- [NFTOverlord.revealFromPrimordial(uint256,bytes)](src/overlords/NFTOverlord.sol#L154-L186)

src/overlords/NFTOverlord.sol#L154-L186


 - [ ] ID-281
	- [LockManager.proposeUSDPrice(uint256,address[])](src/managers/LockManager.sol#L142-L174)

src/managers/LockManager.sol#L142-L174


 - [ ] ID-282
	- [ClaimManager.spendPoints(address,uint256)](src/managers/ClaimManager.sol#L163-L171)

src/managers/ClaimManager.sol#L163-L171


 - [ ] ID-283
	- [NFTAttributesManagerV1.setAttributes(uint256,MunchablesCommonLib.NFTAttributes)](src/managers/NFTAttributeManagerV1.sol#L52-L69)

src/managers/NFTAttributeManagerV1.sol#L52-L69


 - [ ] ID-284
	- [MunchadexManager.updateMunchadex(address,address,uint256)](src/managers/MunchadexManager.sol#L41-L100)

src/managers/MunchadexManager.sol#L41-L100


 - [ ] ID-285
	- [LockManager.setUSDThresholds(uint8,uint8)](src/managers/LockManager.sol#L129-L139)

src/managers/LockManager.sol#L129-L139


 - [ ] ID-286
	- [AccountManager.register(MunchablesCommonLib.Realm,address)](src/managers/AccountManager.sol#L87-L113)

src/managers/AccountManager.sol#L87-L113


 - [ ] ID-287
	- [MunchNFT.mint(address)](src/tokens/MunchNFT.sol#L66-L77)

src/tokens/MunchNFT.sol#L66-L77


 - [ ] ID-288
	- [ConfigStorage.setSmallIntArray(StorageKey,int16[],bool)](src/config/ConfigStorage.sol#L169-L197)

src/config/ConfigStorage.sol#L169-L197


 - [ ] ID-289
	- [ConfigStorage.setUint(StorageKey,uint256,bool)](src/config/ConfigStorage.sol#L55-L62)

src/config/ConfigStorage.sol#L55-L62


 - [ ] ID-290
	- [NFTOverlord.addReveal(address,uint16)](src/overlords/NFTOverlord.sol#L65-L71)

src/overlords/NFTOverlord.sol#L65-L71


 - [ ] ID-291
	- [RNGProxyAPI3.requestRandom(address,bytes4,uint256)](src/rng/RNGProxyAPI3.sol#L30-L47)

src/rng/RNGProxyAPI3.sol#L30-L47


 - [ ] ID-292
	- [PrimordialManager.feedPrimordial(uint256)](src/managers/PrimordialManager.sol#L73-L126)

src/managers/PrimordialManager.sol#L73-L126


 - [ ] ID-293
	- [ClaimManager.newPeriod()](src/managers/ClaimManager.sol#L86-L111)

src/managers/ClaimManager.sol#L86-L111


 - [ ] ID-294
	- [NFTOverlord.reveal(uint256,bytes)](src/overlords/NFTOverlord.sol#L100-L132)

src/overlords/NFTOverlord.sol#L100-L132


 - [ ] ID-295
	- [NFTAttributesManagerV1.createWithImmutable(uint256,MunchablesCommonLib.NFTImmutableAttributes)](src/managers/NFTAttributeManagerV1.sol#L71-L82)

src/managers/NFTAttributeManagerV1.sol#L71-L82


 - [ ] ID-296
	- [OldMunchNFT.setMigrationManager(address)](src/tokens/OldMunchNFT.sol#L100-L104)

src/tokens/OldMunchNFT.sol#L100-L104


 - [ ] ID-297
	- [MigrationManager.burnUnrevealedForPoints()](src/managers/MigrationManager.sol#L135-L146)

src/managers/MigrationManager.sol#L135-L146


 - [ ] ID-298
	- [AccountManager.updatePlayer(address,MunchablesCommonLib.Player)](src/managers/AccountManager.sol#L284-L295)

src/managers/AccountManager.sol#L284-L295


## centralized-risk-other
Impact: Informational
Confidence: High
 - [ ] ID-299
	- [Ownable.renounceOwnership()](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L76-L78)

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L76-L78


 - [ ] ID-300
	- [RNGProxySelfHosted.provideRandom(uint256,bytes)](src/rng/RNGProxySelfHosted.sol#L12-L17)

src/rng/RNGProxySelfHosted.sol#L12-L17


 - [ ] ID-301
	- [PrimordialManager.configUpdated()](src/managers/PrimordialManager.sol#L55-L57)

src/managers/PrimordialManager.sol#L55-L57


 - [ ] ID-302
	- [ConfigStorage.manualNotify(uint8,uint8)](src/config/ConfigStorage.sol#L345-L350)

src/config/ConfigStorage.sol#L345-L350


 - [ ] ID-303
	- [ClaimManager.claimPoints()](src/managers/ClaimManager.sol#L146-L150)

src/managers/ClaimManager.sol#L146-L150


 - [ ] ID-304
	- [OldMunchNFT.configureContract(address)](src/tokens/OldMunchNFT.sol#L93-L98)

src/tokens/OldMunchNFT.sol#L93-L98


 - [ ] ID-305
	- [ProxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy,address,bytes)](lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L38-L44)

lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L38-L44


 - [ ] ID-306
	- [Ownable.transferOwnership(address)](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L84-L89)

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L84-L89


 - [ ] ID-307
	- [SnuggeryManager.feed(uint256,uint256)](src/managers/SnuggeryManager.sol#L133-L171)

src/managers/SnuggeryManager.sol#L133-L171


 - [ ] ID-308
	- [MunchToken.configUpdated()](src/tokens/MunchToken.sol#L24-L26)

src/tokens/MunchToken.sol#L24-L26


 - [ ] ID-309
	- [OldMunchNFT.burn(uint256)](src/tokens/OldMunchNFT.sol#L87-L89)

src/tokens/OldMunchNFT.sol#L87-L89


 - [ ] ID-310
	- [BonusManager.configUpdated()](src/managers/BonusManager.sol#L87-L89)

src/managers/BonusManager.sol#L87-L89


 - [ ] ID-311
	- [AccountManager.forceHarvest(address)](src/managers/AccountManager.sol#L122-L132)

src/managers/AccountManager.sol#L122-L132


 - [ ] ID-312
	- [AccessControl.grantRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L122-L124)

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L122-L124


 - [ ] ID-313
	- [MunchadexManager.configUpdated()](src/managers/MunchadexManager.sol#L37-L39)

src/managers/MunchadexManager.sol#L37-L39


 - [ ] ID-314
	- [OldMunchNFT.pause()](src/tokens/OldMunchNFT.sol#L78-L80)

src/tokens/OldMunchNFT.sol#L78-L80


 - [ ] ID-315
	- [NFTOverlord.configUpdated()](src/overlords/NFTOverlord.sol#L60-L62)

src/overlords/NFTOverlord.sol#L60-L62


 - [ ] ID-316
	- [UUPSUpgradeable.upgradeToAndCall(address,bytes)](lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#L86-L89)

lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#L86-L89


 - [ ] ID-317
	- [MockNFTOverlord.reveal(uint256,bytes)](src/mock/MockNFTOverlord.sol#L87-L97)

src/mock/MockNFTOverlord.sol#L87-L97


 - [ ] ID-318
	- [MunchNFT.configUpdated()](src/tokens/MunchNFT.sol#L62-L64)

src/tokens/MunchNFT.sol#L62-L64


 - [ ] ID-319
	- [FundTreasuryDistributor.configUpdated()](src/distributors/FundTreasuryDistributor.sol#L17-L19)

src/distributors/FundTreasuryDistributor.sol#L17-L19


 - [ ] ID-320
	- [BaseBlastManager.claimERC20Yield(address,uint256)](src/managers/BaseBlastManager.sol#L109-L117)

src/managers/BaseBlastManager.sol#L109-L117


 - [ ] ID-321
	- [LockManager.lock(address,uint256)](src/managers/LockManager.sol#L297-L309)

src/managers/LockManager.sol#L297-L309


 - [ ] ID-322
	- [MunchToken.mint(address,uint256)](src/tokens/MunchToken.sol#L28-L33)

src/tokens/MunchToken.sol#L28-L33


 - [ ] ID-323
	- [ConfigStorage.manualNotifyAddress(address)](src/config/ConfigStorage.sol#L352-L354)

src/config/ConfigStorage.sol#L352-L354


 - [ ] ID-324
	- [MunchNFT.setTokenURI(uint256,string)](src/tokens/MunchNFT.sol#L79-L84)

src/tokens/MunchNFT.sol#L79-L84


 - [ ] ID-325
	- [NFTAttributesManagerV1.configUpdated()](src/managers/NFTAttributeManagerV1.sol#L48-L50)

src/managers/NFTAttributeManagerV1.sol#L48-L50


 - [ ] ID-326
	- [AccountManager.removeSubAccount(address)](src/managers/AccountManager.sol#L254-L258)

src/managers/AccountManager.sol#L254-L258


 - [ ] ID-327
	- [AccessControlUpgradeable.grantRole(bytes32,address)](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L143-L145)

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L143-L145


 - [ ] ID-328
	- [NFTOverlord.mintForMigration(address,MunchablesCommonLib.NFTAttributes,MunchablesCommonLib.NFTImmutableAttributes,MunchablesCommonLib.NFTGameAttribute[])](src/overlords/NFTOverlord.sol#L189-L215)

src/overlords/NFTOverlord.sol#L189-L215


 - [ ] ID-329
	- [SnuggeryManager.configUpdated()](src/managers/SnuggeryManager.sol#L74-L76)

src/managers/SnuggeryManager.sol#L74-L76


 - [ ] ID-330
	- [OldMunchNFT.unpause()](src/tokens/OldMunchNFT.sol#L83-L85)

src/tokens/OldMunchNFT.sol#L83-L85


 - [ ] ID-331
	- [MockNFTOverlord.configUpdated()](src/mock/MockNFTOverlord.sol#L59-L61)

src/mock/MockNFTOverlord.sol#L59-L61


 - [ ] ID-332
	- [RewardsManager.configUpdated()](src/managers/RewardsManager.sol#L54-L56)

src/managers/RewardsManager.sol#L54-L56


 - [ ] ID-333
	- [MigrationManager.configUpdated()](src/managers/MigrationManager.sol#L74-L76)

src/managers/MigrationManager.sol#L74-L76


 - [ ] ID-334
	- [ClaimManager.configUpdated()](src/managers/ClaimManager.sol#L82-L84)

src/managers/ClaimManager.sol#L82-L84


 - [ ] ID-335
	- [MigrationManager.migratePurchasedNFTs(uint256[])](src/managers/MigrationManager.sol#L267-L288)

src/managers/MigrationManager.sol#L267-L288


 - [ ] ID-336
	- [AccessControlUpgradeable.revokeRole(bytes32,address)](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L158-L160)

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L158-L160


 - [ ] ID-337
	- [RNGProxyAPI3.provideRandom(bytes32,bytes)](src/rng/RNGProxyAPI3.sol#L51-L58)

src/rng/RNGProxyAPI3.sol#L51-L58


 - [ ] ID-338
	- [AccessControl.revokeRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L137-L139)

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L137-L139


 - [ ] ID-339
	- [LockManager.configUpdated()](src/managers/LockManager.sol#L85-L87)

src/managers/LockManager.sol#L85-L87


 - [ ] ID-340
	- [AccountManager.configUpdated()](src/managers/AccountManager.sol#L82-L84)

src/managers/AccountManager.sol#L82-L84


 - [ ] ID-341
	- [ClaimManager.forceClaimPoints(address)](src/managers/ClaimManager.sol#L153-L161)

src/managers/ClaimManager.sol#L153-L161


 - [ ] ID-342
	- [LockManager.lockOnBehalf(address,uint256,address)](src/managers/LockManager.sol#L275-L294)

src/managers/LockManager.sol#L275-L294


## dead-function
Impact: Informational
Confidence: Medium
 - [ ] ID-343
[BaseConfigStorageUpgradeable.__BaseConfigStorageUpgradable_reconfigure()](src/config/BaseConfigStorageUpgradeable.sol#L20-L22) is never used and should be removed

src/config/BaseConfigStorageUpgradeable.sol#L20-L22


 - [ ] ID-344
[MockNFTOverlord._calculateRaritySpecies(bytes,MunchablesCommonLib.Rarity)](src/mock/MockNFTOverlord.sol#L329-L382) is never used and should be removed

src/mock/MockNFTOverlord.sol#L329-L382


 - [ ] ID-345
[MockNFTOverlord._getMainAccountRequireRegistered(address)](src/mock/MockNFTOverlord.sol#L427-L440) is never used and should be removed

src/mock/MockNFTOverlord.sol#L427-L440


## error-msg
Impact: Informational
Confidence: Medium
 - [ ] ID-346
require() missing error messages
	 - [require(bool)(msg.data.length == 0)](src/managers/RewardsManager.sol#L40)

src/managers/RewardsManager.sol#L40


## price-manipulation-info
Impact: Informational
Confidence: Medium
 - [ ] ID-347
Potential price manipulation risk:
	- In function _addTokenToOwnerEnumeration
		-- [length = balanceOf(to) - 1](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L118) have potential price manipulated risk from length and call None which could influence variable:length
	- In function _removeTokenFromOwnerEnumeration
		-- [lastTokenIndex = balanceOf(from)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L146) have potential price manipulated risk from lastTokenIndex and call None which could influence variable:lastTokenIndex

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L118


 - [ ] ID-348
Potential price manipulation risk:
	- In function _addTokenToOwnerEnumeration
		-- [length = balanceOf(to) - 1](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L96) have potential price manipulated risk from length and call None which could influence variable:length
	- In function _removeTokenFromOwnerEnumeration
		-- [lastTokenIndex = balanceOf(from)](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L122) have potential price manipulated risk from lastTokenIndex and call None which could influence variable:lastTokenIndex

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L96


## uncontrolled-resource-consumption
Impact: Informational
Confidence: Medium
 - [ ] ID-349
Potential DoS Gas Limit Attack occur in[RewardsManager.claimGasFeeForContracts(address[])](src/managers/RewardsManager.sol#L145-L157)[BEGIN_LOOP](src/managers/RewardsManager.sol#L147-L149)

src/managers/RewardsManager.sol#L145-L157


 - [ ] ID-350
Potential DoS Gas Limit Attack occur in[MockNFTOverlord.levelUp(uint256,bytes)](src/mock/MockNFTOverlord.sol#L188-L243)[BEGIN_LOOP](src/mock/MockNFTOverlord.sol#L212-L237)

src/mock/MockNFTOverlord.sol#L188-L243


 - [ ] ID-351
Potential DoS Gas Limit Attack occur in[AccountManager._removeSubAccount(address,address)](src/managers/AccountManager.sol#L260-L281)[BEGIN_LOOP](src/managers/AccountManager.sol#L267-L276)

src/managers/AccountManager.sol#L260-L281


 - [ ] ID-352
Potential DoS Gas Limit Attack occur in[MockNFTOverlord._populateDefaultProbabilities()](src/mock/MockNFTOverlord.sol#L384-L415)[BEGIN_LOOP](src/mock/MockNFTOverlord.sol#L409-L414)

src/mock/MockNFTOverlord.sol#L384-L415


 - [ ] ID-353
Potential DoS Gas Limit Attack occur in[PrimordialManager._reconfigure()](src/managers/PrimordialManager.sol#L29-L53)[BEGIN_LOOP](src/managers/PrimordialManager.sol#L46-L49)

src/managers/PrimordialManager.sol#L29-L53


 - [ ] ID-354
Potential DoS Gas Limit Attack occur in[SnuggeryManager._recalculateChonks(address)](src/managers/SnuggeryManager.sol#L334-L348)[BEGIN_LOOP](src/managers/SnuggeryManager.sol#L337-L341)

src/managers/SnuggeryManager.sol#L334-L348


 - [ ] ID-355
Potential DoS Gas Limit Attack occur in[ConfigStorage.notify()](src/config/ConfigStorage.sol#L357-L361)[BEGIN_LOOP](src/config/ConfigStorage.sol#L358-L360)

src/config/ConfigStorage.sol#L357-L361


 - [ ] ID-356
Potential DoS Gas Limit Attack occur in[NFTOverlord._populateDefaultRealmLookup()](src/overlords/NFTOverlord.sol#L455-L463)[BEGIN_LOOP](src/overlords/NFTOverlord.sol#L460-L462)

src/overlords/NFTOverlord.sol#L455-L463


 - [ ] ID-357
Potential DoS Gas Limit Attack occur in[MigrationManager.migrateAllNFTs(address,uint32)](src/managers/MigrationManager.sol#L246-L265)[BEGIN_LOOP](src/managers/MigrationManager.sol#L256-L263)

src/managers/MigrationManager.sol#L246-L265


 - [ ] ID-358
Potential DoS Gas Limit Attack occur in[MockNFTOverlord._populateDefaultRealmLookup()](src/mock/MockNFTOverlord.sol#L417-L425)[BEGIN_LOOP](src/mock/MockNFTOverlord.sol#L422-L424)

src/mock/MockNFTOverlord.sol#L417-L425


 - [ ] ID-359
Potential DoS Gas Limit Attack occur in[NFTOverlord._populateDefaultProbabilities()](src/overlords/NFTOverlord.sol#L422-L453)[BEGIN_LOOP](src/overlords/NFTOverlord.sol#L447-L452)

src/overlords/NFTOverlord.sol#L422-L453


 - [ ] ID-360
Potential DoS Gas Limit Attack occur in[LockManager._execUSDPriceUpdate()](src/managers/LockManager.sol#L506-L527)[BEGIN_LOOP](src/managers/LockManager.sol#L512-L523)

src/managers/LockManager.sol#L506-L527


 - [ ] ID-361
Potential DoS Gas Limit Attack occur in[MockSnuggeryManager.setSnuggeryForTest(address,MunchablesCommonLib.SnuggeryNFT[])](src/mock/MockSnuggeryManager.sol#L11-L22)[BEGIN_LOOP](src/mock/MockSnuggeryManager.sol#L18-L20)

src/mock/MockSnuggeryManager.sol#L11-L22


## no-derived-function
Impact: Informational
Confidence: High
 - [ ] ID-362
[MockConfigNotifiable](src/mock/MockConfigNotifiable.sol#L10-L30) does not implement functions:
	- [BaseConfigStorage.__BaseConfigStorage_reconfigure()](src/config/BaseConfigStorage.sol#L85-L87)
	- [BaseConfigStorage.__BaseConfigStorage_setConfigStorage(address)](src/config/BaseConfigStorage.sol#L79-L83)
	- [MockConfigNotifiable.verifyDirtyUint(uint256)](src/mock/MockConfigNotifiable.sol#L22-L24)
	- [MockConfigNotifiable.verifyUint(StorageKey,uint256)](src/mock/MockConfigNotifiable.sol#L26-L29)

src/mock/MockConfigNotifiable.sol#L10-L30


## unnecessary-public-function-modifier
Impact: Informational
Confidence: High
 - [ ] ID-363
function:[ERC721.ownerOf(uint256)](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L67-L69)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L67-L69


 - [ ] ID-364
function:[ConfigStorage.setUintArray(StorageKey,uint256[],bool)](src/config/ConfigStorage.sol#L69-L96)is public and can be replaced with external 

src/config/ConfigStorage.sol#L69-L96


 - [ ] ID-365
function:[ERC721.name()](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L74-L76)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L74-L76


 - [ ] ID-366
function:[ERC721URIStorage.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L27-L29)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L27-L29


 - [ ] ID-367
function:[AccessControlUpgradeable.supportsInterface(bytes4)](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L90-L92)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L90-L92


 - [ ] ID-368
function:[ERC20.symbol()](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L66-L68)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L66-L68


 - [ ] ID-369
function:[ERC721Upgradeable.tokenURI(uint256)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L110-L115)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L110-L115


 - [ ] ID-370
function:[AccessControl.grantRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L122-L124)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L122-L124


 - [ ] ID-371
function:[ConfigStorage.setUint(StorageKey,uint256,bool)](src/config/ConfigStorage.sol#L55-L62)is public and can be replaced with external 

src/config/ConfigStorage.sol#L55-L62


 - [ ] ID-372
function:[OldMunchNFT.setURI(uint256,string)](src/tokens/OldMunchNFT.sol#L145-L160)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L145-L160


 - [ ] ID-373
function:[ERC721.getApproved(uint256)](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L114-L118)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L114-L118


 - [ ] ID-374
function:[OldMunchNFT.pause()](src/tokens/OldMunchNFT.sol#L78-L80)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L78-L80


 - [ ] ID-375
function:[OldMunchNFT.burn(uint256)](src/tokens/OldMunchNFT.sol#L87-L89)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L87-L89


 - [ ] ID-376
function:[ConfigStorage.getUniversalRole(Role)](src/config/ConfigStorage.sol#L50-L52)is public and can be replaced with external 

src/config/ConfigStorage.sol#L50-L52


 - [ ] ID-377
function:[ERC165Upgradeable.supportsInterface(bytes4)](lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#L30-L32)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#L30-L32


 - [ ] ID-378
function:[ConfigStorage.removeNotifiableAddress(address)](src/config/ConfigStorage.sol#L325-L335)is public and can be replaced with external 

src/config/ConfigStorage.sol#L325-L335


 - [ ] ID-379
function:[ConfigStorage.getUintArray(StorageKey)](src/config/ConfigStorage.sol#L98-L109)is public and can be replaced with external 

src/config/ConfigStorage.sol#L98-L109


 - [ ] ID-380
function:[ERC20.totalSupply()](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L90-L92)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L90-L92


 - [ ] ID-381
function:[ConfigStorage.setSmallInt(StorageKey,int16,bool)](src/config/ConfigStorage.sol#L155-L162)is public and can be replaced with external 

src/config/ConfigStorage.sol#L155-L162


 - [ ] ID-382
function:[OldMunchNFT.initialize(address,address,address,address,address)](src/tokens/OldMunchNFT.sol#L53-L75)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L53-L75


 - [ ] ID-383
function:[ERC721Upgradeable.symbol()](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L102-L105)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L102-L105


 - [ ] ID-384
function:[ProxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy,address,bytes)](lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L38-L44)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L38-L44


 - [ ] ID-385
function:[ClaimManager.initialize(address)](src/managers/ClaimManager.sol#L49-L52)is public and can be replaced with external 

src/managers/ClaimManager.sol#L49-L52


 - [ ] ID-386
function:[ERC721.balanceOf(address)](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L57-L62)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L57-L62


 - [ ] ID-387
function:[AccessControl.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L71-L73)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L71-L73


 - [ ] ID-388
function:[ConfigStorage.getBytes32(StorageKey)](src/config/ConfigStorage.sol#L307-L309)is public and can be replaced with external 

src/config/ConfigStorage.sol#L307-L309


 - [ ] ID-389
function:[ERC20.approve(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L132-L136)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L132-L136


 - [ ] ID-390
function:[AccessControlUpgradeable.grantRole(bytes32,address)](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L143-L145)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L143-L145


 - [ ] ID-391
function:[ERC20.balanceOf(address)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L97-L99)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L97-L99


 - [ ] ID-392
function:[ERC721EnumerableUpgradeable.tokenOfOwnerByIndex(address,uint256)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L64-L70)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L64-L70


 - [ ] ID-393
function:[ConfigStorage.setAddressArray(StorageKey,address[],bool)](src/config/ConfigStorage.sol#L254-L281)is public and can be replaced with external 

src/config/ConfigStorage.sol#L254-L281


 - [ ] ID-394
function:[AccessControl.revokeRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L137-L139)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L137-L139


 - [ ] ID-395
function:[ConfigStorage.setBytes32(StorageKey,bytes32,bool)](src/config/ConfigStorage.sol#L297-L304)is public and can be replaced with external 

src/config/ConfigStorage.sol#L297-L304


 - [ ] ID-396
function:[MockConfigNotifiable.verifyDirtyUint(uint256)](src/mock/MockConfigNotifiable.sol#L22-L24)is public and can be replaced with external 

src/mock/MockConfigNotifiable.sol#L22-L24


 - [ ] ID-397
function:[ConfigStorage.getUint(StorageKey)](src/config/ConfigStorage.sol#L65-L67)is public and can be replaced with external 

src/config/ConfigStorage.sol#L65-L67


 - [ ] ID-398
function:[ERC721Upgradeable.setApprovalForAll(address,bool)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L145-L147)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L145-L147


 - [ ] ID-399
function:[ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool)](src/config/ConfigStorage.sol#L111-L139)is public and can be replaced with external 

src/config/ConfigStorage.sol#L111-L139


 - [ ] ID-400
function:[ConfigStorage.getSmallInt(StorageKey)](src/config/ConfigStorage.sol#L165-L167)is public and can be replaced with external 

src/config/ConfigStorage.sol#L165-L167


 - [ ] ID-401
function:[ERC20.transferFrom(address,address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L154-L159)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L154-L159


 - [ ] ID-402
function:[ERC721Upgradeable.getApproved(uint256)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L136-L140)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L136-L140


 - [ ] ID-403
function:[RNGProxySelfHosted.provideRandom(uint256,bytes)](src/rng/RNGProxySelfHosted.sol#L12-L17)is public and can be replaced with external 

src/rng/RNGProxySelfHosted.sol#L12-L17


 - [ ] ID-404
function:[BaseBlastManagerUpgradeable.initialize(address)](src/managers/BaseBlastManagerUpgradeable.sol#L17-L21)is public and can be replaced with external 

src/managers/BaseBlastManagerUpgradeable.sol#L17-L21


 - [ ] ID-405
function:[OldMunchNFT.supportsInterface(bytes4)](src/tokens/OldMunchNFT.sol#L253-L267)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L253-L267


 - [ ] ID-406
function:[Ownable.transferOwnership(address)](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L84-L89)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L84-L89


 - [ ] ID-407
function:[ERC20.name()](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L58-L60)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L58-L60


 - [ ] ID-408
function:[MockConfigNotifiable.verifyUint(StorageKey,uint256)](src/mock/MockConfigNotifiable.sol#L26-L29)is public and can be replaced with external 

src/mock/MockConfigNotifiable.sol#L26-L29


 - [ ] ID-409
function:[ERC721.symbol()](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L81-L83)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L81-L83


 - [ ] ID-410
function:[ConfigStorage.manualNotify(uint8,uint8)](src/config/ConfigStorage.sol#L345-L350)is public and can be replaced with external 

src/config/ConfigStorage.sol#L345-L350


 - [ ] ID-411
function:[BaseConfigStorageUpgradeable.initialize(address)](src/config/BaseConfigStorageUpgradeable.sol#L16-L18)is public and can be replaced with external 

src/config/BaseConfigStorageUpgradeable.sol#L16-L18


 - [ ] ID-412
function:[ERC721Enumerable.tokenOfOwnerByIndex(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L46-L51)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L46-L51


 - [ ] ID-413
function:[OldMunchNFT.safeMint(address,string)](src/tokens/OldMunchNFT.sol#L184-L193)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L184-L193


 - [ ] ID-414
function:[MockMunchNFT.mint(address)](src/mock/MockMunchNFT.sol#L9)is public and can be replaced with external 

src/mock/MockMunchNFT.sol#L9


 - [ ] ID-415
function:[Ownable.renounceOwnership()](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L76-L78)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L76-L78


 - [ ] ID-416
function:[ConfigStorage.setUniversalRole(Role,address)](src/config/ConfigStorage.sol#L35-L37)is public and can be replaced with external 

src/config/ConfigStorage.sol#L35-L37


 - [ ] ID-417
function:[MunchNFT.tokenURI(uint256)](src/tokens/MunchNFT.sol#L114-L123)is public and can be replaced with external 

src/tokens/MunchNFT.sol#L114-L123


 - [ ] ID-418
function:[ERC721Upgradeable.supportsInterface(bytes4)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L66-L71)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L66-L71


 - [ ] ID-419
function:[ERC721URIStorageUpgradeable.tokenURI(uint256)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L52-L69)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L52-L69


 - [ ] ID-420
function:[ERC721Enumerable.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L39-L41)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L39-L41


 - [ ] ID-421
function:[RNGProxyAPI3.requestRandom(address,bytes4,uint256)](src/rng/RNGProxyAPI3.sol#L30-L47)is public and can be replaced with external 

src/rng/RNGProxyAPI3.sol#L30-L47


 - [ ] ID-422
function:[MockRNGProxy.callLevelUpForTest(uint256,bytes)](src/mock/MockRNGProxy.sol#L29-L31)is public and can be replaced with external 

src/mock/MockRNGProxy.sol#L29-L31


 - [ ] ID-423
function:[OldMunchNFT.fetchTokens(address)](src/tokens/OldMunchNFT.sol#L166-L179)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L166-L179


 - [ ] ID-424
function:[OldMunchNFT.revealNFT(address,string,bytes32,bytes,MunchablesCommonLib.Realm,MunchablesCommonLib.Rarity,uint16)](src/tokens/OldMunchNFT.sol#L110-L140)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L110-L140


 - [ ] ID-425
function:[ERC721Upgradeable.name()](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L94-L97)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L94-L97


 - [ ] ID-426
function:[ERC165.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#L24-L26)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#L24-L26


 - [ ] ID-427
function:[ConfigStorage.setSmallIntArray(StorageKey,int16[],bool)](src/config/ConfigStorage.sol#L169-L197)is public and can be replaced with external 

src/config/ConfigStorage.sol#L169-L197


 - [ ] ID-428
function:[ERC721.safeTransferFrom(address,address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L152-L154)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L152-L154


 - [ ] ID-429
function:[ERC721EnumerableUpgradeable.tokenByIndex(uint256)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L83-L89)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L83-L89


 - [ ] ID-430
function:[ERC721Upgradeable.balanceOf(address)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L76-L82)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L76-L82


 - [ ] ID-431
function:[ConfigStorage.addNotifiableAddress(address)](src/config/ConfigStorage.sol#L312-L314)is public and can be replaced with external 

src/config/ConfigStorage.sol#L312-L314


 - [ ] ID-432
function:[AccessControlUpgradeable.revokeRole(bytes32,address)](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L158-L160)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L158-L160


 - [ ] ID-433
function:[ConfigStorage.setAddress(StorageKey,address,bool)](src/config/ConfigStorage.sol#L228-L235)is public and can be replaced with external 

src/config/ConfigStorage.sol#L228-L235


 - [ ] ID-434
function:[UUPSUpgradeable.upgradeToAndCall(address,bytes)](lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#L86-L89)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#L86-L89


 - [ ] ID-435
function:[MockRNGProxy.callRevealForTest(uint256,bytes)](src/mock/MockRNGProxy.sol#L33-L38)is public and can be replaced with external 

src/mock/MockRNGProxy.sol#L33-L38


 - [ ] ID-436
function:[ERC721.supportsInterface(bytes4)](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L47-L52)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L47-L52


 - [ ] ID-437
function:[ERC721.tokenURI(uint256)](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L88-L93)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L88-L93


 - [ ] ID-438
function:[MockRNGProxy.callRevealFromPrimordialForTest(uint256,bytes)](src/mock/MockRNGProxy.sol#L40-L45)is public and can be replaced with external 

src/mock/MockRNGProxy.sol#L40-L45


 - [ ] ID-439
function:[AccountManager.initialize(address)](src/managers/AccountManager.sol#L48-L51)is public and can be replaced with external 

src/managers/AccountManager.sol#L48-L51


 - [ ] ID-440
function:[MockMunchNFT.burn(uint256)](src/mock/MockMunchNFT.sol#L7)is public and can be replaced with external 

src/mock/MockMunchNFT.sol#L7


 - [ ] ID-441
function:[ERC721.approve(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L107-L109)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L107-L109


 - [ ] ID-442
function:[ConfigStorage.addNotifiableAddresses(address[])](src/config/ConfigStorage.sol#L316-L322)is public and can be replaced with external 

src/config/ConfigStorage.sol#L316-L322


 - [ ] ID-443
function:[ERC20.transfer(address,uint256)](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L109-L113)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L109-L113


 - [ ] ID-444
function:[ERC721URIStorage.tokenURI(uint256)](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L34-L50)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L34-L50


 - [ ] ID-445
function:[MockRNGProxy.provideRandomForTest(uint256,bytes)](src/mock/MockRNGProxy.sol#L25-L27)is public and can be replaced with external 

src/mock/MockRNGProxy.sol#L25-L27


 - [ ] ID-446
function:[ERC721Upgradeable.approve(address,uint256)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L129-L131)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L129-L131


 - [ ] ID-447
function:[ConfigStorage.getBool(StorageKey)](src/config/ConfigStorage.sol#L223-L225)is public and can be replaced with external 

src/config/ConfigStorage.sol#L223-L225


 - [ ] ID-448
function:[ConfigStorage.getRole(Role)](src/config/ConfigStorage.sol#L39-L41)is public and can be replaced with external 

src/config/ConfigStorage.sol#L39-L41


 - [ ] ID-449
function:[ConfigStorage.getSmallIntArray(StorageKey)](src/config/ConfigStorage.sol#L199-L210)is public and can be replaced with external 

src/config/ConfigStorage.sol#L199-L210


 - [ ] ID-450
function:[ConfigStorage.getAddressArray(StorageKey)](src/config/ConfigStorage.sol#L283-L294)is public and can be replaced with external 

src/config/ConfigStorage.sol#L283-L294


 - [ ] ID-451
function:[ConfigStorage.setBool(StorageKey,bool,bool)](src/config/ConfigStorage.sol#L213-L220)is public and can be replaced with external 

src/config/ConfigStorage.sol#L213-L220


 - [ ] ID-452
function:[ERC721Enumerable.tokenByIndex(uint256)](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L63-L68)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L63-L68


 - [ ] ID-453
function:[ConfigStorage.manualNotifyAddress(address)](src/config/ConfigStorage.sol#L352-L354)is public and can be replaced with external 

src/config/ConfigStorage.sol#L352-L354


 - [ ] ID-454
function:[ConfigStorage.setAddresses(StorageKey[],address[],bool)](src/config/ConfigStorage.sol#L238-L247)is public and can be replaced with external 

src/config/ConfigStorage.sol#L238-L247


 - [ ] ID-455
function:[ERC721EnumerableUpgradeable.supportsInterface(bytes4)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L57-L59)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L57-L59


 - [ ] ID-456
function:[ERC721Upgradeable.ownerOf(uint256)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L87-L89)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L87-L89


 - [ ] ID-457
function:[BaseRNGProxy.requestRandom(address,bytes4,uint256)](src/rng/BaseRNGProxy.sol#L21-L32)is public and can be replaced with external 

src/rng/BaseRNGProxy.sol#L21-L32


 - [ ] ID-458
function:[ERC20.decimals()](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L83-L85)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L83-L85


 - [ ] ID-459
function:[NFTAttributesManagerV1.getGameAttributeDataType(uint8)](src/managers/NFTAttributeManagerV1.sol#L171-L196)is public and can be replaced with external 

src/managers/NFTAttributeManagerV1.sol#L171-L196


 - [ ] ID-460
function:[ERC721URIStorageUpgradeable.supportsInterface(bytes4)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L45-L47)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L45-L47


 - [ ] ID-461
function:[ERC721Upgradeable.safeTransferFrom(address,address,uint256)](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L175-L177)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L175-L177


 - [ ] ID-462
function:[AccessControlUpgradeable.renounceRole(bytes32,address)](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L178-L184)is public and can be replaced with external 

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L178-L184


 - [ ] ID-463
function:[OldMunchNFT.configureContract(address)](src/tokens/OldMunchNFT.sol#L93-L98)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L93-L98


 - [ ] ID-464
function:[AccessControl.renounceRole(bytes32,address)](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L157-L163)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L157-L163


 - [ ] ID-465
function:[ConfigStorage.setRole(Role,address,address)](src/config/ConfigStorage.sol#L27-L33)is public and can be replaced with external 

src/config/ConfigStorage.sol#L27-L33


 - [ ] ID-466
function:[OldMunchNFT.unpause()](src/tokens/OldMunchNFT.sol#L83-L85)is public and can be replaced with external 

src/tokens/OldMunchNFT.sol#L83-L85


 - [ ] ID-467
function:[MunchNFT.supportsInterface(bytes4)](src/tokens/MunchNFT.sol#L125-L134)is public and can be replaced with external 

src/tokens/MunchNFT.sol#L125-L134


 - [ ] ID-468
function:[ConfigStorage.getSmallUintArray(StorageKey)](src/config/ConfigStorage.sol#L141-L152)is public and can be replaced with external 

src/config/ConfigStorage.sol#L141-L152


 - [ ] ID-469
function:[ConfigStorage.getAddress(StorageKey)](src/config/ConfigStorage.sol#L250-L252)is public and can be replaced with external 

src/config/ConfigStorage.sol#L250-L252


 - [ ] ID-470
function:[ERC721.setApprovalForAll(address,bool)](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L123-L125)is public and can be replaced with external 

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L123-L125


## unused-event
Impact: Informational
Confidence: High
 - [ ] ID-471
[OldMunchNFTRevealNFT(address,uint256,string,bytes32,MunchablesCommonLib.Realm,MunchablesCommonLib.Rarity,uint16)](src/tokens/OldMunchNFT.sol#L272-L280) is never emitted in [OldMunchNFT](src/tokens/OldMunchNFT.sol#L13-L281)

src/tokens/OldMunchNFT.sol#L272-L280


 - [ ] ID-472
[OldMunchNFTContractConfigured(address)](src/tokens/OldMunchNFT.sol#L269) is never emitted in [OldMunchNFT](src/tokens/OldMunchNFT.sol#L13-L281)

src/tokens/OldMunchNFT.sol#L269


 - [ ] ID-473
[OldMunchNFTSetURI(uint256,string)](src/tokens/OldMunchNFT.sol#L271) is never emitted in [OldMunchNFT](src/tokens/OldMunchNFT.sol#L13-L281)

src/tokens/OldMunchNFT.sol#L271


## version-only
Impact: Informational
Confidence: High
 - [ ] ID-474
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/IERC721Metadata.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/IERC721Metadata.sol#L4


 - [ ] ID-475
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/utils/ContextUpgradeable.sol#L4


 - [ ] ID-476
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/IERC20.sol#L4


 - [ ] ID-477
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/utils/introspection/ERC165Upgradeable.sol#L4


 - [ ] ID-478
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/rng/RNGProxyAPI3.sol#L2)

src/rng/RNGProxyAPI3.sol#L2


 - [ ] ID-479
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/interfaces/IERC1967.sol#L4)

lib/openzeppelin-contracts/contracts/interfaces/IERC1967.sol#L4


 - [ ] ID-480
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/interfaces/draft-IERC6093.sol#L3)

lib/openzeppelin-contracts/contracts/interfaces/draft-IERC6093.sol#L3


 - [ ] ID-481
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol#L4


 - [ ] ID-482
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4


 - [ ] ID-483
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/access/AccessControlUpgradeable.sol#L4


 - [ ] ID-484
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC20/ERC20.sol#L4


 - [ ] ID-485
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/ERC721Upgradeable.sol#L4


 - [ ] ID-486
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Pausable.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Pausable.sol#L4


 - [ ] ID-487
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol#L4


 - [ ] ID-488
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/interfaces/draft-IERC1822.sol#L4)

lib/openzeppelin-contracts/contracts/interfaces/draft-IERC1822.sol#L4


 - [ ] ID-489
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](src/config/ConfigStorage.sol#L2)

src/config/ConfigStorage.sol#L2


 - [ ] ID-490
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IAirnodeRrpV0.sol#L2)

node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IAirnodeRrpV0.sol#L2


 - [ ] ID-491
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/math/Math.sol#L4)

lib/openzeppelin-contracts/contracts/utils/math/Math.sol#L4


 - [ ] ID-492
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/mock/MockBlast.sol#L2)

src/mock/MockBlast.sol#L2


 - [ ] ID-493
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/rng/RNGProxySelfHosted.sol#L2)

src/rng/RNGProxySelfHosted.sol#L2


 - [ ] ID-494
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/proxy/beacon/IBeacon.sol#L4)

lib/openzeppelin-contracts/contracts/proxy/beacon/IBeacon.sol#L4


 - [ ] ID-495
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/rng/BaseRNGProxy.sol#L2)

src/rng/BaseRNGProxy.sol#L2


 - [ ] ID-496
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/interfaces/IERC721.sol#L4)

lib/openzeppelin-contracts/contracts/interfaces/IERC721.sol#L4


 - [ ] ID-497
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/Address.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Address.sol#L4


 - [ ] ID-498
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/proxy/Proxy.sol#L4)

lib/openzeppelin-contracts/contracts/proxy/Proxy.sol#L4


 - [ ] ID-499
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L4)

lib/openzeppelin-contracts/contracts/access/AccessControl.sol#L4


 - [ ] ID-500
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol#L4)

lib/openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol#L4


 - [ ] ID-501
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IAuthorizationUtilsV0.sol#L2)

node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IAuthorizationUtilsV0.sol#L2


 - [ ] ID-502
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/mock/MockConfigNotifiable.sol#L2)

src/mock/MockConfigNotifiable.sol#L2


 - [ ] ID-503
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/Pausable.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Pausable.sol#L4


 - [ ] ID-504
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC721/IERC721Receiver.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/IERC721Receiver.sol#L4


 - [ ] ID-505
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L4)

lib/openzeppelin-contracts/contracts/proxy/transparent/ProxyAdmin.sol#L4


 - [ ] ID-506
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](src/tokens/OldMunchNFT.sol#L2)

src/tokens/OldMunchNFT.sol#L2


 - [ ] ID-507
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/access/Ownable.sol#L4)

lib/openzeppelin-contracts/contracts/access/Ownable.sol#L4


 - [ ] ID-508
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/interfaces/IConfigStorage.sol#L2)

src/interfaces/IConfigStorage.sol#L2


 - [ ] ID-509
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/interfaces/IERC4906.sol#L4)

lib/openzeppelin-contracts/contracts/interfaces/IERC4906.sol#L4


 - [ ] ID-510
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/interfaces/IERC165.sol#L4)

lib/openzeppelin-contracts/contracts/interfaces/IERC165.sol#L4


 - [ ] ID-511
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol#L4


 - [ ] ID-512
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/interfaces/IPrimordialManager.sol#L2)

src/interfaces/IPrimordialManager.sol#L2


 - [ ] ID-513
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/utils/PausableUpgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/utils/PausableUpgradeable.sol#L4


 - [ ] ID-514
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol#L4)

lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Proxy.sol#L4


 - [ ] ID-515
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721Enumerable.sol#L4


 - [ ] ID-516
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](src/interfaces/IConfigNotifiable.sol#L2)

src/interfaces/IConfigNotifiable.sol#L2


 - [ ] ID-517
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/libraries/SignatureVerifier.sol#L2)

src/libraries/SignatureVerifier.sol#L2


 - [ ] ID-518
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/Context.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Context.sol#L4


 - [ ] ID-519
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#L4)

lib/openzeppelin-contracts/contracts/proxy/ERC1967/ERC1967Utils.sol#L4


 - [ ] ID-520
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/StorageSlot.sol#L5)

lib/openzeppelin-contracts/contracts/utils/StorageSlot.sol#L5


 - [ ] ID-521
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721PausableUpgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721PausableUpgradeable.sol#L4


 - [ ] ID-522
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/introspection/IERC165.sol#L4)

lib/openzeppelin-contracts/contracts/utils/introspection/IERC165.sol#L4


 - [ ] ID-523
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC721/IERC721.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/IERC721.sol#L4


 - [ ] ID-524
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/IERC721Enumerable.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/IERC721Enumerable.sol#L4


 - [ ] ID-525
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol#L4


 - [ ] ID-526
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol#L4)

lib/openzeppelin-contracts/contracts/proxy/transparent/TransparentUpgradeableProxy.sol#L4


 - [ ] ID-527
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/ITemplateUtilsV0.sol#L2)

node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/ITemplateUtilsV0.sol#L2


 - [ ] ID-528
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L4)

lib/openzeppelin-contracts/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L4


 - [ ] ID-529
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts-upgradeable/contracts/utils/ReentrancyGuardUpgradeable.sol#L4)

lib/openzeppelin-contracts-upgradeable/contracts/utils/ReentrancyGuardUpgradeable.sol#L4


 - [ ] ID-530
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#L4)

lib/openzeppelin-contracts/contracts/utils/introspection/ERC165.sol#L4


 - [ ] ID-531
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/interfaces/IBlast.sol#L2)

src/interfaces/IBlast.sol#L2


 - [ ] ID-532
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/math/SignedMath.sol#L4)

lib/openzeppelin-contracts/contracts/utils/math/SignedMath.sol#L4


 - [ ] ID-533
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#L4)

lib/openzeppelin-contracts/contracts/proxy/utils/UUPSUpgradeable.sol#L4


 - [ ] ID-534
. analyzed (120 contracts with 86 detectors), 543 result(s) found
INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/2024-05-munchables_scaned_output/mwe-output.json generate success.
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IWithdrawalUtilsV0.sol#L2)

node_modules/@api3/airnode-protocol/contracts/rrp/interfaces/IWithdrawalUtilsV0.sol#L2


 - [ ] ID-535
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@api3/airnode-protocol/contracts/rrp/requesters/RrpRequesterV0.sol#L2)

node_modules/@api3/airnode-protocol/contracts/rrp/requesters/RrpRequesterV0.sol#L2


 - [ ] ID-536
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/interfaces/IRNGProxy.sol#L2)

src/interfaces/IRNGProxy.sol#L2


 - [ ] ID-537
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/interfaces/IERC20.sol#L4)

lib/openzeppelin-contracts/contracts/interfaces/IERC20.sol#L4


 - [ ] ID-538
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/utils/Strings.sol#L4)

lib/openzeppelin-contracts/contracts/utils/Strings.sol#L4


 - [ ] ID-539
	Pragma confirmed better, here is pragma solidity^0.8.25. [^0.8.25](src/interfaces/IRNGProxySelfHosted.sol#L2)

src/interfaces/IRNGProxySelfHosted.sol#L2


 - [ ] ID-540
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](lib/openzeppelin-contracts/contracts/access/IAccessControl.sol#L4)

lib/openzeppelin-contracts/contracts/access/IAccessControl.sol#L4


## state-should-be-constant
Impact: Optimization
Confidence: High
 - [ ] ID-541
[MigrationManager.purchasedUnlockPrice](src/managers/MigrationManager.sol#L23) should be constant

src/managers/MigrationManager.sol#L23


 - [ ] ID-542
[BaseBlastManager._blastClaimableConfigured](src/managers/BaseBlastManager.sol#L23) should be constant

src/managers/BaseBlastManager.sol#L23


Execution time for Falcon: 40.711857626s
