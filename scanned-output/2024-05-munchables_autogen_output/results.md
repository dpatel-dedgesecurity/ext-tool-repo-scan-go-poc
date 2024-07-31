'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables)
'forge config --json' running
'forge build --build-info --skip */test/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables)
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

NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).j (src/managers/NFTAttributeManagerV1.sol#164) is a local variable never initialized
AccountManager.spraySchnibblesPropose(address[],uint256[]).i (src/managers/AccountManager.sol#163) is a local variable never initialized
NFTOverlord.levelUp(uint256,bytes).currentValue (src/overlords/NFTOverlord.sol#260) is a local variable never initialized
BonusManager._calculateMunchadexBonus(address).i (src/managers/BonusManager.sol#226) is a local variable never initialized
NFTAttributesManagerV1.setGameAttributes(uint256,MunchablesCommonLib.NFTGameAttribute[]).i (src/managers/NFTAttributeManagerV1.sol#92) is a local variable never initialized
PrimordialManager._reconfigure().i (src/managers/PrimordialManager.sol#46) is a local variable never initialized
BonusManager.getHarvestBonus(address)._migrationBonus (src/managers/BonusManager.sol#131) is a local variable never initialized
SnuggeryManager.getTotalChonk(address).i (src/managers/SnuggeryManager.sol#289) is a local variable never initialized
AccountManager.execSprayProposal(address).i (src/managers/AccountManager.sol#197) is a local variable never initialized
MunchadexManager.getMunchadexInfo(address).i (src/managers/MunchadexManager.sol#124) is a local variable never initialized
SnuggeryManager._findSnuggeryIndex(address,uint256).i (src/managers/SnuggeryManager.sol#310) is a local variable never initialized
registrationDate is a member never initialized in LockManager._lock(address,uint256,address,address)._player (src/managers/LockManager.sol#319)
MockNFTOverlord.levelUp(uint256,bytes).currentValue (src/mock/MockNFTOverlord.sol#227) is a local variable never initialized
MockNFTOverlord._calculateRaritySpecies(bytes,MunchablesCommonLib.Rarity).selectedSpeciesId (src/mock/MockNFTOverlord.sol#338) is a local variable never initialized
tokenId is a member never initialized in MigrationManager._migrateNFTs(address,address,uint256[]).snapshot (src/managers/MigrationManager.sol#300)
NFTOverlord.levelUp(uint256,bytes).zeroed (src/overlords/NFTOverlord.sol#244) is a local variable never initialized
referrer is a member never initialized in ClaimManager._claimPoints(address).player (src/managers/ClaimManager.sol#203)
ConfigStorage.getUintArray(StorageKey).i (src/config/ConfigStorage.sol#104) is a local variable never initialized
SnuggeryManager._recalculateChonks(address).i (src/managers/SnuggeryManager.sol#337) is a local variable never initialized
NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).i_scope_0 (src/managers/NFTAttributeManagerV1.sol#157) is a local variable never initialized
ConfigStorage.getSmallUintArray(StorageKey).i (src/config/ConfigStorage.sol#147) is a local variable never initialized
maxSnuggerySize is a member never initialized in BonusManager._calculateLevelBonus(address).player (src/managers/BonusManager.sol#197)
MockNFTOverlord.levelUp(uint256,bytes).zeroed (src/mock/MockNFTOverlord.sol#211) is a local variable never initialized
LockManager._execUSDPriceUpdate().i (src/managers/LockManager.sol#512) is a local variable never initialized
ConfigStorage.setSmallUintArray(StorageKey,uint8[],bool).i (src/config/ConfigStorage.sol#119) is a local variable never initialized
ConfigStorage.setUintArray(StorageKey,uint256[],bool).i (src/config/ConfigStorage.sol#77) is a local variable never initialized
ConfigStorage.getAddressArray(StorageKey).i (src/config/ConfigStorage.sol#289) is a local variable never initialized
ConfigStorage.setSmallIntArray(StorageKey,int16[],bool).i (src/config/ConfigStorage.sol#177) is a local variable never initialized
AccountManager.addSubAccount(address).i (src/managers/AccountManager.sol#244) is a local variable never initialized
RewardsManager.claimYieldForContracts(address[]).ongoingETH (src/managers/RewardsManager.sol#114) is a local variable never initialized
RewardsManager.claimYieldForContracts(address[]).ongoingWETH (src/managers/RewardsManager.sol#115) is a local variable never initialized
NFTOverlord.levelUp(uint256,bytes).i (src/overlords/NFTOverlord.sol#245) is a local variable never initialized
FundTreasuryDistributor.receiveTokens(IDistributor.TokenBag[]).i (src/distributors/FundTreasuryDistributor.sol#31) is a local variable never initialized
LockManager.getLockedWeightedValue(address).i (src/managers/LockManager.sol#466) is a local variable never initialized
RewardsManager.claimYieldForContracts(address[]).ongoingUSDB (src/managers/RewardsManager.sol#116) is a local variable never initialized
MockNFTOverlord.levelUp(uint256,bytes).i (src/mock/MockNFTOverlord.sol#212) is a local variable never initialized
ClaimManager.burnNFTsForPoints(address,uint8[]).i (src/managers/ClaimManager.sol#123) is a local variable never initialized
MunchadexManager.getMunchadexInfo(address).i_scope_0 (src/managers/MunchadexManager.sol#130) is a local variable never initialized
ConfigStorage.setAddresses(StorageKey[],address[],bool).i (src/config/ConfigStorage.sol#243) is a local variable never initialized
LockManager.setLockDuration(uint256).i (src/managers/LockManager.sol#252) is a local variable never initialized
ConfigStorage.getSmallIntArray(StorageKey).i (src/config/ConfigStorage.sol#205) is a local variable never initialized
NFTOverlord._calculateRaritySpecies(bytes,MunchablesCommonLib.Rarity).selectedSpeciesId (src/overlords/NFTOverlord.sol#376) is a local variable never initialized
MockSnuggeryManager.setSnuggeryForTest(address,MunchablesCommonLib.SnuggeryNFT[]).i (src/mock/MockSnuggeryManager.sol#18) is a local variable never initialized
ConfigStorage.setAddressArray(StorageKey,address[],bool).i (src/config/ConfigStorage.sol#262) is a local variable never initialized
NFTAttributesManagerV1.getGameAttributes(uint256,MunchablesCommonLib.GameAttributeIndex[]).i (src/managers/NFTAttributeManagerV1.sol#150) is a local variable never initialized
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
. analyzed (120 contracts with 86 detectors), 543 result(s) found
INFO:Falcon:metatrust result: ../../repos-output/2024-05-munchables_autogen_output/mwe-output.json generate success.
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
An error occurred: 
after for recrusion
Execution time for Falcon: 45.125212202s
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/config/BaseConfigStorage.sol: Abstract Contract (1): 88 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/config/BaseConfigStorageUpgradeable.sol: Abstract Contract (1): 23 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/config/ConfigStorage.sol: Smart Contract (1): 362 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/distributors/FundTreasuryDistributor.sol: Smart Contract (1): 51 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IAccountManager.sol: Interface (1): 218 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IBaseBlastManager.sol: Interface (1): 6 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IBlast.sol: Interface (3): 121 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IBonusManager.sol: Interface (1): 24 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IClaimManager.sol: Interface (1): 134 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IConfigNotifiable.sol: Interface (1): 8 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IConfigStorage.sol: Interface (1): 208 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IDistributor.sol: Interface (1): 22 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IERC20YieldClaimable.sol: Interface (1): 8 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IHoldsGovernorship.sol: Interface (1): 12 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/ILockManager.sol: Interface (1): 309 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IMigrationManager.sol: Interface (2): 168 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IMunchNFT.sol: Interface (1): 17 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IMunchToken.sol: Interface (1): 7 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IMunchadexManager.sol: Interface (1): 39 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/INFTAttributesManager.sol: Interface (1): 103 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/INFTOverlord.sol: Interface (1): 165 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IPrimordialManager.sol: Interface (1): 60 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IRNGProxy.sol: Interface (1): 28 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IRNGProxySelfHosted.sol: Interface (1): 6 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/IRewardsManager.sol: Interface (1): 28 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/interfaces/ISnuggeryManager.sol: Interface (1): 164 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/libraries/MunchablesCommonLib.sol: Library (1): 173 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/libraries/SignatureVerifier.sol: Library (1): 28 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/AccountManager.sol: Smart Contract (1): 405 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/BaseBlastManager.sol: Abstract Contract (1): 122 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/BaseBlastManagerUpgradeable.sol: Abstract Contract (1): 26 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/BonusManager.sol: Smart Contract (1): 261 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/ClaimManager.sol: Smart Contract (1): 251 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/LockManager.sol: Smart Contract (1): 528 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/MigrationManager.sol: Smart Contract (1): 389 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/MunchadexManager.sol: Smart Contract (1): 138 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/NFTAttributeManagerV1.sol: Smart Contract (1): 197 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/PrimordialManager.sol: Smart Contract (1): 173 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/RewardsManager.sol: Smart Contract (1): 265 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/managers/SnuggeryManager.sol: Smart Contract (1): 349 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockAccountManager.sol: Smart Contract (1): 18 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockBlast.sol: Smart Contract (1): 118 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockClaimManager.sol: Smart Contract (1): 15 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockConfigNotifiable.sol: Smart Contract (1): 30 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockLockManager.sol: Smart Contract (1): 20 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockMigrationManager.sol: Smart Contract (1): 47 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockMunchNFT.sol: Smart Contract (1): 10 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockMunchadexManager.sol: Smart Contract (1): 43 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockNFTAttributeManagerV1.sol: Smart Contract (1): 25 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockNFTOverlord.sol: Smart Contract (1): 441 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockPrimordialManager.sol: Smart Contract (1): 12 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockRNGProxy.sol: Smart Contract (1): 46 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockRNGRequester.sol: Smart Contract (1): 27 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockRewardsManager.sol: Smart Contract (1): 20 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/mock/MockSnuggeryManager.sol: Smart Contract (1): 42 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/overlords/NFTOverlord.sol: Smart Contract (1): 479 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/proxy/ProxyFactory.sol: Smart Contract (1): 12 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/rng/BaseRNGProxy.sol: Abstract Contract (1): 55 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/rng/RNGProxyAPI3.sol: Smart Contract (1): 59 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/rng/RNGProxySelfHosted.sol: Smart Contract (1): 18 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/tokens/MunchNFT.sol: Smart Contract (1): 135 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/tokens/MunchToken.sol: Smart Contract (1): 34 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/tokens/OldMunchNFT.sol: Smart Contract (1): 281 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2024-05-munchables/src/tokens/TestERC20Token.sol: Smart Contract (1): 34 lines

Summary:
Total number of Solidity files: 64
Abstract Contracts Files: 5
Smart Contracts Files: 35
Libraries Files: 2
Interfaces Files: 22
Multiple Types Files: 0
******************************************
Abstract Contracts: 5
Smart Contracts: 35
Libraries: 2
Interfaces: 25
LOC: 7705
