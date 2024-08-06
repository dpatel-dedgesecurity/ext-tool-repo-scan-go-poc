'npx hardhat clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-02-openq)
'npx hardhat clean --global' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-02-openq)
'npx hardhat compile --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-02-openq)

ERC1967UpgradeUpgradeable._functionDelegateCall(address,bytes) (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#198-204) uses delegatecall to a input-controlled function id
	- (success,returndata) = target.delegatecall(data) (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#202)
Reference:  

state variable: BountyStorageCore.tokenAddresses (contracts/Bounty/Storage/BountyStorageCore.sol#58) not initialized and not written in contract but be used in contract
state variable: BountyStorageCore.bountyId (contracts/Bounty/Storage/BountyStorageCore.sol#33) not initialized and not written in contract but be used in contract
state variable: BountyStorageCore.status (contracts/Bounty/Storage/BountyStorageCore.sol#39) not initialized and not written in contract but be used in contract
state variable: BountyStorageCore.tokenId (contracts/Bounty/Storage/BountyStorageCore.sol#49) not initialized and not written in contract but be used in contract
state variable: BountyStorageCore.nftDeposits (contracts/Bounty/Storage/BountyStorageCore.sol#55) not initialized and not written in contract but be used in contract
state variable: BountyStorageCore.nftDepositLimit (contracts/Bounty/Storage/BountyStorageCore.sol#40) not initialized and not written in contract but be used in contract
state variable: TieredBountyStorageCore.payoutSchedule (contracts/Bounty/Storage/TieredBountyStorageCore.sol#15) not initialized and not written in contract but be used in contract
state variable: MockNft._tokenIdCounter (contracts/Mocks/MockNft.sol#12) not initialized and not written in contract but be used in contract
Reference:  

public mint or burn found in TestToken.mint(address,uint256) (contracts/Mocks/TestToken.sol#21-25)Reference: check public mint method

BountyCore.getLockedFunds(address).lockedFunds (contracts/Bounty/Implementations/BountyCore.sol#339) is a local variable never initialized
BountyFactory.mintBounty(string,address,string,address,address,OpenQDefinitions.InitOperation).beaconProxy (contracts/BountyFactory/BountyFactory.sol#59) is a local variable never initialized
Reference:  

ERC1967Upgrade._upgradeToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#65-74)have ignores return value in Address.functionDelegateCall(newImplementation,data) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#72)
ERC1967Upgrade._upgradeBeaconToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#174-184)have ignores return value in Address.functionDelegateCall(IBeacon(newBeacon).implementation(),data) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#182)
BountyCore.receiveFunds(address,address,uint256,uint256) (contracts/Bounty/Implementations/BountyCore.sol#21-58)have ignores return value in tokenAddresses.add(_tokenAddress) (contracts/Bounty/Implementations/BountyCore.sol#55)
Reference:  

function:ContextUpgradeable.__Context_init() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#18-19)is empty 
function:ContextUpgradeable.__Context_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#21-22)is empty 
function:ERC1967UpgradeUpgradeable.__ERC1967Upgrade_init() (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#21-22)is empty 
function:ERC1967UpgradeUpgradeable.__ERC1967Upgrade_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#24-25)is empty 
function:UUPSUpgradeable.__UUPSUpgradeable_init() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#23-24)is empty 
function:UUPSUpgradeable.__UUPSUpgradeable_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#26-27)is empty 
function:ERC721HolderUpgradeable.__ERC721Holder_init() (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#16-17)is empty 
function:ERC721HolderUpgradeable.__ERC721Holder_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#19-20)is empty 
function:ERC20._beforeTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#358-362)is empty 
function:ERC20._afterTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#378-382)is empty 
function:ERC721._beforeTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#432-436)is empty 
function:ERC721._afterTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#449-453)is empty 
function:ClaimManagerV1._authorizeUpgrade(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#368)is empty 
function:DepositManagerV1._authorizeUpgrade(address) (contracts/DepositManager/Implementations/DepositManagerV1.sol#229)is empty 
function:OpenQV1._authorizeUpgrade(address) (contracts/OpenQ/Implementations/OpenQV1.sol#445)is empty 
Reference:  

Setter function AtomicBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/AtomicBountyV1.sol#27-84) does not emit an event
Setter function AtomicBountyV1.claimBalance(address,address) (contracts/Bounty/Implementations/AtomicBountyV1.sol#89-98) does not emit an event
Setter function AtomicBountyV1.close(address,bytes) (contracts/Bounty/Implementations/AtomicBountyV1.sol#104-117) does not emit an event
Setter function AtomicBountyV1.receiveNft(address,address,uint256,uint256,bytes) (contracts/Bounty/Implementations/AtomicBountyV1.sol#125-152) does not emit an event
Setter function AtomicBountyV1.setInvoiceComplete(bytes) (contracts/Bounty/Implementations/AtomicBountyV1.sol#157-160) does not emit an event
Setter function AtomicBountyV1.setSupportingDocumentsComplete(bytes) (contracts/Bounty/Implementations/AtomicBountyV1.sol#165-171) does not emit an event
Setter function OngoingBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/OngoingBountyV1.sol#27-90) does not emit an event
Setter function OngoingBountyV1.claimOngoingPayout(address,bytes) (contracts/Bounty/Implementations/OngoingBountyV1.sol#96-112) does not emit an event
Setter function OngoingBountyV1.closeOngoing(address) (contracts/Bounty/Implementations/OngoingBountyV1.sol#116-125) does not emit an event
Setter function OngoingBountyV1.receiveNft(address,address,uint256,uint256,bytes) (contracts/Bounty/Implementations/OngoingBountyV1.sol#133-160) does not emit an event
Setter function OngoingBountyV1.setPayout(address,uint256) (contracts/Bounty/Implementations/OngoingBountyV1.sol#165-171) does not emit an event
Setter function OngoingBountyV1.setInvoiceComplete(bytes) (contracts/Bounty/Implementations/OngoingBountyV1.sol#176-183) does not emit an event
Setter function OngoingBountyV1.setSupportingDocumentsComplete(bytes) (contracts/Bounty/Implementations/OngoingBountyV1.sol#188-198) does not emit an event
Setter function TieredFixedBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#27-86) does not emit an event
Setter function TieredFixedBountyV1.claimTieredFixed(address,uint256) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#91-107) does not emit an event
Setter function TieredFixedBountyV1.closeCompetition() (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#110-118) does not emit an event
Setter function TieredFixedBountyV1.setFundingGoal(address,uint256) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#123-133) does not emit an event
Setter function TieredFixedBountyV1.setPayoutScheduleFixed(uint256[],address) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#138-171) does not emit an event
Setter function TieredPercentageBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#26-97) does not emit an event
Setter function TieredPercentageBountyV1.claimTiered(address,uint256,address) (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#104-120) does not emit an event
Setter function TieredPercentageBountyV1.closeCompetition() (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#123-136) does not emit an event
Setter function TieredPercentageBountyV1.setPayoutSchedule(uint256[]) (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#141-179) does not emit an event
Setter function BountyFactory.mintBounty(string,address,string,address,address,OpenQDefinitions.InitOperation) (contracts/BountyFactory/BountyFactory.sol#49-88) does not emit an event
Setter function ClaimManagerV1.initialize(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#20-24) does not emit an event
Setter function ClaimManagerV1._authorizeUpgrade(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#368) does not emit an event
Setter function ClaimManagerV1.transferOracle(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#372-375) does not emit an event
Setter function ClaimManagerV1.setOpenQ(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#378-380) does not emit an event
Setter function ClaimManagerV1.setKyc(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#384-386) does not emit an event
Setter function DepositManagerV1.initialize() (contracts/DepositManager/Implementations/DepositManagerV1.sol#15-18) does not emit an event
Setter function DepositManagerV1.setTokenWhitelist(address) (contracts/DepositManager/Implementations/DepositManagerV1.sol#22-28) does not emit an event
Setter function DepositManagerV1._authorizeUpgrade(address) (contracts/DepositManager/Implementations/DepositManagerV1.sol#229) does not emit an event
Setter function MockDai.constructor() (contracts/Mocks/MockDai.sol#10-13) does not emit an event
Setter function MockKyc.setIsValid(bool) (contracts/Mocks/MockKyc.sol#9-11) does not emit an event
Setter function MockLink.constructor() (contracts/Mocks/MockLink.sol#10-13) does not emit an event
Setter function MockNft.safeMint(address) (contracts/Mocks/MockNft.sol#19-24) does not emit an event
Setter function TokenFeeToken.constructor() (contracts/Mocks/TransferFeeToken.sol#10-13) does not emit an event
Setter function OpenQV1.initialize() (contracts/OpenQ/Implementations/OpenQV1.sol#14-18) does not emit an event
Setter function OpenQV1.setBountyFactory(address) (contracts/OpenQ/Implementations/OpenQV1.sol#63-69) does not emit an event
Setter function OpenQV1.setClaimManager(address) (contracts/OpenQ/Implementations/OpenQV1.sol#73-79) does not emit an event
Setter function OpenQV1.setDepositManager(address) (contracts/OpenQ/Implementations/OpenQV1.sol#83-89) does not emit an event
Setter function OpenQV1._authorizeUpgrade(address) (contracts/OpenQ/Implementations/OpenQV1.sol#445) does not emit an event
Setter function OpenQV1.transferOracle(address) (contracts/OpenQ/Implementations/OpenQV1.sol#455-458) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Condition variable is not initialized found in Initializable._disableInitializers() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#131-137)
Reference: initialize method should has permission check

variable lacks a zero-check on 		- BountyCore.setFundingGoal(address,uint256) (contracts/Bounty/Implementations/BountyCore.sol#141-149)
variable lacks a zero-check on 		- AtomicBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/AtomicBountyV1.sol#27-84)
variable lacks a zero-check on 		- OngoingBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/OngoingBountyV1.sol#27-90)
variable lacks a zero-check on 		- OngoingBountyV1.setPayout(address,uint256) (contracts/Bounty/Implementations/OngoingBountyV1.sol#165-171)
variable lacks a zero-check on 		- TieredFixedBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#27-86)
variable lacks a zero-check on 		- TieredFixedBountyV1.setFundingGoal(address,uint256) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#123-133)
variable lacks a zero-check on 		- TieredFixedBountyV1.setPayoutScheduleFixed(uint256[],address) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#138-171)
variable lacks a zero-check on 		- TieredPercentageBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#26-97)
variable lacks a zero-check on 		- BountyFactory.constructor(address,address,address,address,address) (contracts/BountyFactory/BountyFactory.sol#26-39)
variable lacks a zero-check on 		- ClaimManagerV1.claimBounty(address,address,bytes) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#31-67)
variable lacks a zero-check on 		- ClaimManagerV1.setOpenQ(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#378-380)
variable lacks a zero-check on 		- ClaimManagerV1.hasKYC(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#390-392)
variable lacks a zero-check on 		- DepositManagerV1.refundDeposit(address,bytes32) (contracts/DepositManager/Implementations/DepositManagerV1.sol#152-195)
variable lacks a zero-check on 		- DepositManagerV1.isWhitelisted(address) (contracts/DepositManager/Implementations/DepositManagerV1.sol#200-202)
variable lacks a zero-check on 		- OpenQV1.setClaimManager(address) (contracts/OpenQ/Implementations/OpenQV1.sol#73-79)
variable lacks a zero-check on 		- OpenQV1.setDepositManager(address) (contracts/OpenQ/Implementations/OpenQV1.sol#83-89)
variable lacks a zero-check on 		- OpenQV1.setFundingGoal(string,address,uint256) (contracts/OpenQ/Implementations/OpenQV1.sol#117-136)
variable lacks a zero-check on 		- OpenQV1.setPayout(string,address,uint256) (contracts/OpenQ/Implementations/OpenQV1.sol#256-275)
variable lacks a zero-check on 		- OpenQV1.associateExternalIdToAddress(string,address) (contracts/OpenQ/Implementations/OpenQV1.sol#464-488)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

	- BountyCore.receiveFunds(address,address,uint256,uint256) (contracts/Bounty/Implementations/BountyCore.sol#21-58)
	- BountyCore.refundDeposit(bytes32,address,uint256) (contracts/Bounty/Implementations/BountyCore.sol#64-93)
	- BountyCore.extendDeposit(bytes32,uint256,address) (contracts/Bounty/Implementations/BountyCore.sol#99-120)
	- BountyCore.setFundingGoal(address,uint256) (contracts/Bounty/Implementations/BountyCore.sol#141-149)
	- BountyCore.setKycRequired(bool) (contracts/Bounty/Implementations/BountyCore.sol#153-155)
	- BountyCore.setInvoiceRequired(bool) (contracts/Bounty/Implementations/BountyCore.sol#159-165)
	- BountyCore.setSupportingDocumentsRequired(bool) (contracts/Bounty/Implementations/BountyCore.sol#169-175)
	- AtomicBountyV1.close(address,bytes) (contracts/Bounty/Implementations/AtomicBountyV1.sol#104-117)
	- AtomicBountyV1.receiveNft(address,address,uint256,uint256,bytes) (contracts/Bounty/Implementations/AtomicBountyV1.sol#125-152)
	- AtomicBountyV1.setInvoiceComplete(bytes) (contracts/Bounty/Implementations/AtomicBountyV1.sol#157-160)
	- AtomicBountyV1.setSupportingDocumentsComplete(bytes) (contracts/Bounty/Implementations/AtomicBountyV1.sol#165-171)
	- OngoingBountyV1.claimOngoingPayout(address,bytes) (contracts/Bounty/Implementations/OngoingBountyV1.sol#96-112)
	- OngoingBountyV1.closeOngoing(address) (contracts/Bounty/Implementations/OngoingBountyV1.sol#116-125)
	- OngoingBountyV1.receiveNft(address,address,uint256,uint256,bytes) (contracts/Bounty/Implementations/OngoingBountyV1.sol#133-160)
	- OngoingBountyV1.setPayout(address,uint256) (contracts/Bounty/Implementations/OngoingBountyV1.sol#165-171)
	- OngoingBountyV1.setInvoiceComplete(bytes) (contracts/Bounty/Implementations/OngoingBountyV1.sol#176-183)
	- OngoingBountyV1.setSupportingDocumentsComplete(bytes) (contracts/Bounty/Implementations/OngoingBountyV1.sol#188-198)
	- TieredBountyCore.receiveNft(address,address,uint256,uint256,bytes) (contracts/Bounty/Implementations/TieredBountyCore.sol#18-48)
	- TieredBountyCore.setTierClaimed(uint256) (contracts/Bounty/Implementations/TieredBountyCore.sol#52-54)
	- TieredBountyCore.setTierWinner(string,uint256) (contracts/Bounty/Implementations/TieredBountyCore.sol#59-64)
	- TieredBountyCore.setInvoiceComplete(bytes) (contracts/Bounty/Implementations/TieredBountyCore.sol#69-75)
	- TieredBountyCore.setSupportingDocumentsComplete(bytes) (contracts/Bounty/Implementations/TieredBountyCore.sol#80-89)
	- TieredFixedBountyV1.closeCompetition() (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#110-118)
	- TieredFixedBountyV1.setFundingGoal(address,uint256) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#123-133)
	- TieredFixedBountyV1.setPayoutScheduleFixed(uint256[],address) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#138-171)
	- TieredPercentageBountyV1.closeCompetition() (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#123-136)
	- TieredPercentageBountyV1.setPayoutSchedule(uint256[]) (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#141-179)
	- ClaimManagerV1.setOpenQ(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#378-380)
	- ClaimManagerV1.setKyc(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#384-386)
	- DepositManagerV1.setTokenWhitelist(address) (contracts/DepositManager/Implementations/DepositManagerV1.sol#22-28)
	- OpenQV1.mintBounty(string,string,OpenQDefinitions.InitOperation) (contracts/OpenQ/Implementations/OpenQV1.sol#26-59)
	- OpenQV1.setBountyFactory(address) (contracts/OpenQ/Implementations/OpenQV1.sol#63-69)
	- OpenQV1.setClaimManager(address) (contracts/OpenQ/Implementations/OpenQV1.sol#73-79)
	- OpenQV1.setDepositManager(address) (contracts/OpenQ/Implementations/OpenQV1.sol#83-89)
	- OpenQV1.associateExternalIdToAddress(string,address) (contracts/OpenQ/Implementations/OpenQV1.sol#464-488)
	- TokenWhitelist.addToken(address) (contracts/TokenWhitelist/TokenWhitelist.sol#25-32)
	- TokenWhitelist.removeToken(address) (contracts/TokenWhitelist/TokenWhitelist.sol#36-43)
	- TokenWhitelist.setTokenAddressLimit(uint256) (contracts/TokenWhitelist/TokenWhitelist.sol#47-52)
Reference:  

	- OwnableUpgradeable.renounceOwnership() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#66-68)
	- OwnableUpgradeable.transferOwnership(address) (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#74-77)
	- UUPSUpgradeable.upgradeTo(address) (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#72-75)
	- UUPSUpgradeable.upgradeToAndCall(address,bytes) (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#85-88)
	- Ownable.renounceOwnership() (node_modules/@openzeppelin/contracts/access/Ownable.sol#61-63)
	- Ownable.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable.sol#69-72)
	- UpgradeableBeacon.upgradeTo(address) (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#49-52)
	- BountyCore.claimNft(address,bytes32) (contracts/Bounty/Implementations/BountyCore.sol#125-136)
	- AtomicBountyV1.claimBalance(address,address) (contracts/Bounty/Implementations/AtomicBountyV1.sol#89-98)
	- TieredFixedBountyV1.claimTieredFixed(address,uint256) (contracts/Bounty/Implementations/TieredFixedBountyV1.sol#91-107)
	- TieredPercentageBountyV1.claimTiered(address,uint256,address) (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#104-120)
	- BountyFactory.mintBounty(string,address,string,address,address,OpenQDefinitions.InitOperation) (contracts/BountyFactory/BountyFactory.sol#49-88)
	- ClaimManagerV1.initialize(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#20-24)
	- ClaimManagerV1.claimBounty(address,address,bytes) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#31-67)
	- ClaimManagerV1.permissionedClaimTieredBounty(address,bytes) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#75-116)
	- ClaimManagerV1.transferOracle(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#372-375)
	- DepositManagerV1.fundBountyToken(address,address,uint256,uint256,string) (contracts/DepositManager/Implementations/DepositManagerV1.sol#36-74)
	- DepositManagerV1.extendDeposit(address,bytes32,uint256) (contracts/DepositManager/Implementations/DepositManagerV1.sol#80-105)
	- DepositManagerV1.fundBountyNFT(address,address,uint256,uint256,bytes) (contracts/DepositManager/Implementations/DepositManagerV1.sol#113-147)
	- DepositManagerV1.refundDeposit(address,bytes32) (contracts/DepositManager/Implementations/DepositManagerV1.sol#152-195)
	- MockNft.safeMint(address) (contracts/Mocks/MockNft.sol#19-24)
	- OpenQV1.initialize() (contracts/OpenQ/Implementations/OpenQV1.sol#14-18)
	- OpenQV1.setFundingGoal(string,address,uint256) (contracts/OpenQ/Implementations/OpenQV1.sol#117-136)
	- OpenQV1.setKycRequired(string,bool) (contracts/OpenQ/Implementations/OpenQV1.sol#141-158)
	- OpenQV1.setInvoiceRequired(string,bool) (contracts/OpenQ/Implementations/OpenQV1.sol#163-180)
	- OpenQV1.setSupportingDocumentsRequired(string,bool) (contracts/OpenQ/Implementations/OpenQV1.sol#185-202)
	- OpenQV1.setInvoiceComplete(string,bytes) (contracts/OpenQ/Implementations/OpenQV1.sol#207-226)
	- OpenQV1.setSupportingDocumentsComplete(string,bytes) (contracts/OpenQ/Implementations/OpenQV1.sol#231-250)
	- OpenQV1.setPayout(string,address,uint256) (contracts/OpenQ/Implementations/OpenQV1.sol#256-275)
	- OpenQV1.setPayoutSchedule(string,uint256[]) (contracts/OpenQ/Implementations/OpenQV1.sol#281-299)
	- OpenQV1.setPayoutScheduleFixed(string,uint256[],address) (contracts/OpenQ/Implementations/OpenQV1.sol#305-324)
	- OpenQV1.transferOracle(address) (contracts/OpenQ/Implementations/OpenQV1.sol#455-458)
Reference:  

Potential DoS Gas Limit Attack occur inTieredPercentageBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation) (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#26-97)BEGIN_LOOP (contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#78-80)
Potential DoS Gas Limit Attack occur inClaimManagerV1._claimAtomicBounty(IAtomicBounty,address,bytes) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#123-166)BEGIN_LOOP (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#130-148)
Potential DoS Gas Limit Attack occur inClaimManagerV1._claimTieredPercentageBounty(IBounty,address,bytes) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#203-272)BEGIN_LOOP (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#230-249)
Potential DoS Gas Limit Attack occur inClaimManagerV1._claimTieredFixedBounty(IBounty,address,bytes) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#278-341)BEGIN_LOOP (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#320-338)
Reference: https://swcregistry.io/docs/SWC-128

function:OwnableUpgradeable.renounceOwnership() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#66-68)is public and can be replaced with external 
function:OwnableUpgradeable.transferOwnership(address) (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#74-77)is public and can be replaced with external 
function:ERC721HolderUpgradeable.onERC721Received(address,address,uint256,bytes) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#26-33)is public and can be replaced with external 
function:Ownable.renounceOwnership() (node_modules/@openzeppelin/contracts/access/Ownable.sol#61-63)is public and can be replaced with external 
function:Ownable.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable.sol#69-72)is public and can be replaced with external 
function:UpgradeableBeacon.implementation() (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#35-37)is public and can be replaced with external 
function:UpgradeableBeacon.upgradeTo(address) (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#49-52)is public and can be replaced with external 
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
function:BountyCore.getTokenAddresses() (contracts/Bounty/Implementations/BountyCore.sol#315-322)is public and can be replaced with external 
function:BountyCore.getLockedFunds(address) (contracts/Bounty/Implementations/BountyCore.sol#333-352)is public and can be replaced with external 
function:OnlyOpenQ.openQ() (contracts/OnlyOpenQ/OnlyOpenQ.sol#20-22)is public and can be replaced with external 
function:BountyCore.getTokenBalance(address) (contracts/Bounty/Implementations/BountyCore.sol#275-286)is public and can be replaced with external 
function:OngoingBountyV1.getClaimIds() (contracts/Bounty/Implementations/OngoingBountyV1.sol#224-226)is public and can be replaced with external 
function:ClaimManagerV1.bountyIsClaimable(address) (contracts/ClaimManager/Implementations/ClaimManagerV1.sol#345-365)is public and can be replaced with external 
function:MockNft.tokenURI(uint256) (contracts/Mocks/MockNft.sol#35-42)is public and can be replaced with external 
Reference:  

	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/interfaces/draft-IERC1822Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/proxy/beacon/IBeaconUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/draft-IERC20PermitUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721ReceiverUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/StorageSlotUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/access/Ownable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/Proxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (node_modules/@openzeppelin/contracts/utils/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Context.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Counters.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Strings.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#4)
Reference:  

MockNft.uri (contracts/Mocks/MockNft.sol#16-17) should be constant
Reference:  

unnecessary reentrancy lock found inAtomicBountyV1
	- BountyCore.claimNft(address,bytes32) (contracts/Bounty/Implementations/BountyCore.sol#125-136)
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
An error occurred: 
after for recrusion
Summary
 - [arbitrary-delegatecall](#arbitrary-delegatecall) (1 results) (Critical)
 - [state-variable-not-initialized](#state-variable-not-initialized) (8 results) (High)
 - [public-mint-burn](#public-mint-burn) (1 results) (Medium)
 - [uninitialized-local](#uninitialized-local) (2 results) (Medium)
 - [unused-return](#unused-return) (3 results) (Medium)
 - [void-function](#void-function) (15 results) (Medium)
 - [pess-event-setter](#pess-event-setter) (42 results) (Low)
 - [initialize-permission](#initialize-permission) (1 results) (Low)
 - [missing-zero-check](#missing-zero-check) (19 results) (Low)
 - [centralized-risk-informational](#centralized-risk-informational) (38 results) (Informational)
 - [centralized-risk-other](#centralized-risk-other) (32 results) (Informational)
 - [uncontrolled-resource-consumption](#uncontrolled-resource-consumption) (4 results) (Informational)
 - [unnecessary-public-function-modifier](#unnecessary-public-function-modifier) (36 results) (Informational)
 - [version-only](#version-only) (41 results) (Informational)
 - [state-should-be-constant](#state-should-be-constant) (1 results) (Optimization)
 - [unnecessary-reentrancy-lock](#unnecessary-reentrancy-lock) (1 results) (Optimization)
## arbitrary-delegatecall
Impact: Critical
Confidence: Medium
 - [ ] ID-0
[ERC1967UpgradeUpgradeable._functionDelegateCall(address,bytes)](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L198-L204) uses delegatecall to a input-controlled function id
	- [(success,returndata) = target.delegatecall(data)](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L202)

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L198-L204


## state-variable-not-initialized
Impact: High
Confidence: High
 - [ ] ID-1
state variable: [BountyStorageCore.nftDepositLimit](contracts/Bounty/Storage/BountyStorageCore.sol#L40) not initialized and not written in contract but be used in contract

contracts/Bounty/Storage/BountyStorageCore.sol#L40


 - [ ] ID-2
state variable: [BountyStorageCore.status](contracts/Bounty/Storage/BountyStorageCore.sol#L39) not initialized and not written in contract but be used in contract

contracts/Bounty/Storage/BountyStorageCore.sol#L39


 - [ ] ID-3
state variable: [MockNft._tokenIdCounter](contracts/Mocks/MockNft.sol#L12) not initialized and not written in contract but be used in contract

contracts/Mocks/MockNft.sol#L12


 - [ ] ID-4
state variable: [BountyStorageCore.tokenAddresses](contracts/Bounty/Storage/BountyStorageCore.sol#L58) not initialized and not written in contract but be used in contract

contracts/Bounty/Storage/BountyStorageCore.sol#L58


 - [ ] ID-5
state variable: [BountyStorageCore.bountyId](contracts/Bounty/Storage/BountyStorageCore.sol#L33) not initialized and not written in contract but be used in contract

contracts/Bounty/Storage/BountyStorageCore.sol#L33


 - [ ] ID-6
state variable: [BountyStorageCore.tokenId](contracts/Bounty/Storage/BountyStorageCore.sol#L49) not initialized and not written in contract but be used in contract

contracts/Bounty/Storage/BountyStorageCore.sol#L49


 - [ ] ID-7
state variable: [BountyStorageCore.nftDeposits](contracts/Bounty/Storage/BountyStorageCore.sol#L55) not initialized and not written in contract but be used in contract

contracts/Bounty/Storage/BountyStorageCore.sol#L55


 - [ ] ID-8
state variable: [TieredBountyStorageCore.payoutSchedule](contracts/Bounty/Storage/TieredBountyStorageCore.sol#L15) not initialized and not written in contract but be used in contract

contracts/Bounty/Storage/TieredBountyStorageCore.sol#L15


## public-mint-burn
Impact: Medium
Confidence: Medium
 - [ ] ID-9
public mint or burn found in [TestToken.mint(address,uint256)](contracts/Mocks/TestToken.sol#L21-L25)
contracts/Mocks/TestToken.sol#L21-L25


## uninitialized-local
Impact: Medium
Confidence: Medium
 - [ ] ID-10
[BountyFactory.mintBounty(string,address,string,address,address,OpenQDefinitions.InitOperation).beaconProxy](contracts/BountyFactory/BountyFactory.sol#L59) is a local variable never initialized

contracts/BountyFactory/BountyFactory.sol#L59


 - [ ] ID-11
[BountyCore.getLockedFunds(address).lockedFunds](contracts/Bounty/Implementations/BountyCore.sol#L339) is a local variable never initialized

contracts/Bounty/Implementations/BountyCore.sol#L339


## unused-return
Impact: Medium
Confidence: Medium
 - [ ] ID-12
[BountyCore.receiveFunds(address,address,uint256,uint256)](contracts/Bounty/Implementations/BountyCore.sol#L21-L58)have ignores return value in [tokenAddresses.add(_tokenAddress)](contracts/Bounty/Implementations/BountyCore.sol#L55)

contracts/Bounty/Implementations/BountyCore.sol#L21-L58


 - [ ] ID-13
[ERC1967Upgrade._upgradeToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L65-L74)have ignores return value in [Address.functionDelegateCall(newImplementation,data)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L72)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L65-L74


 - [ ] ID-14
[ERC1967Upgrade._upgradeBeaconToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L174-L184)have ignores return value in [Address.functionDelegateCall(IBeacon(newBeacon).implementation(),data)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L182)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L174-L184


## void-function
Impact: Medium
Confidence: High
 - [ ] ID-15
function:[ERC20._afterTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L378-L382)is empty 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L378-L382


 - [ ] ID-16
function:[ContextUpgradeable.__Context_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L21-L22)is empty 

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L21-L22


 - [ ] ID-17
function:[DepositManagerV1._authorizeUpgrade(address)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L229)is empty 

contracts/DepositManager/Implementations/DepositManagerV1.sol#L229


 - [ ] ID-18
function:[UUPSUpgradeable.__UUPSUpgradeable_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L26-L27)is empty 

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L26-L27


 - [ ] ID-19
function:[ClaimManagerV1._authorizeUpgrade(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L368)is empty 

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L368


 - [ ] ID-20
function:[ERC721HolderUpgradeable.__ERC721Holder_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#L19-L20)is empty 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#L19-L20


 - [ ] ID-21
function:[ERC721._afterTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L449-L453)is empty 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L449-L453


 - [ ] ID-22
function:[ERC721HolderUpgradeable.__ERC721Holder_init()](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#L16-L17)is empty 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#L16-L17


 - [ ] ID-23
function:[OpenQV1._authorizeUpgrade(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L445)is empty 

contracts/OpenQ/Implementations/OpenQV1.sol#L445


 - [ ] ID-24
function:[ContextUpgradeable.__Context_init()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L18-L19)is empty 

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L18-L19


 - [ ] ID-25
function:[ERC1967UpgradeUpgradeable.__ERC1967Upgrade_init()](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L21-L22)is empty 

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L21-L22


 - [ ] ID-26
function:[UUPSUpgradeable.__UUPSUpgradeable_init()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L23-L24)is empty 

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L23-L24


 - [ ] ID-27
function:[ERC20._beforeTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L358-L362)is empty 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L358-L362


 - [ ] ID-28
function:[ERC1967UpgradeUpgradeable.__ERC1967Upgrade_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L24-L25)is empty 

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L24-L25


 - [ ] ID-29
function:[ERC721._beforeTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L432-L436)is empty 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L432-L436


## pess-event-setter
Impact: Low
Confidence: Medium
 - [ ] ID-30
Setter function [ClaimManagerV1._authorizeUpgrade(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L368) does not emit an event

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L368


 - [ ] ID-31
Setter function [MockKyc.setIsValid(bool)](contracts/Mocks/MockKyc.sol#L9-L11) does not emit an event

contracts/Mocks/MockKyc.sol#L9-L11


 - [ ] ID-32
Setter function [TieredFixedBountyV1.setPayoutScheduleFixed(uint256[],address)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L138-L171) does not emit an event

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L138-L171


 - [ ] ID-33
Setter function [OpenQV1.initialize()](contracts/OpenQ/Implementations/OpenQV1.sol#L14-L18) does not emit an event

contracts/OpenQ/Implementations/OpenQV1.sol#L14-L18


 - [ ] ID-34
Setter function [TieredPercentageBountyV1.setPayoutSchedule(uint256[])](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L141-L179) does not emit an event

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L141-L179


 - [ ] ID-35
Setter function [OngoingBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L27-L90) does not emit an event

contracts/Bounty/Implementations/OngoingBountyV1.sol#L27-L90


 - [ ] ID-36
Setter function [TieredFixedBountyV1.setFundingGoal(address,uint256)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L123-L133) does not emit an event

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L123-L133


 - [ ] ID-37
Setter function [OngoingBountyV1.setSupportingDocumentsComplete(bytes)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L188-L198) does not emit an event

contracts/Bounty/Implementations/OngoingBountyV1.sol#L188-L198


 - [ ] ID-38
Setter function [ClaimManagerV1.setOpenQ(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L378-L380) does not emit an event

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L378-L380


 - [ ] ID-39
Setter function [ClaimManagerV1.initialize(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L20-L24) does not emit an event

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L20-L24


 - [ ] ID-40
Setter function [ClaimManagerV1.setKyc(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L384-L386) does not emit an event

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L384-L386


 - [ ] ID-41
Setter function [AtomicBountyV1.claimBalance(address,address)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L89-L98) does not emit an event

contracts/Bounty/Implementations/AtomicBountyV1.sol#L89-L98


 - [ ] ID-42
Setter function [OpenQV1.setDepositManager(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L83-L89) does not emit an event

contracts/OpenQ/Implementations/OpenQV1.sol#L83-L89


 - [ ] ID-43
Setter function [TieredPercentageBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L26-L97) does not emit an event

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L26-L97


 - [ ] ID-44
Setter function [DepositManagerV1.setTokenWhitelist(address)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L22-L28) does not emit an event

contracts/DepositManager/Implementations/DepositManagerV1.sol#L22-L28


 - [ ] ID-45
Setter function [AtomicBountyV1.close(address,bytes)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L104-L117) does not emit an event

contracts/Bounty/Implementations/AtomicBountyV1.sol#L104-L117


 - [ ] ID-46
Setter function [OngoingBountyV1.setPayout(address,uint256)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L165-L171) does not emit an event

contracts/Bounty/Implementations/OngoingBountyV1.sol#L165-L171


 - [ ] ID-47
Setter function [OpenQV1.setClaimManager(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L73-L79) does not emit an event

contracts/OpenQ/Implementations/OpenQV1.sol#L73-L79


 - [ ] ID-48
Setter function [TieredFixedBountyV1.closeCompetition()](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L110-L118) does not emit an event

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L110-L118


 - [ ] ID-49
Setter function [ClaimManagerV1.transferOracle(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L372-L375) does not emit an event

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L372-L375


 - [ ] ID-50
Setter function [DepositManagerV1._authorizeUpgrade(address)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L229) does not emit an event

contracts/DepositManager/Implementations/DepositManagerV1.sol#L229


 - [ ] ID-51
Setter function [MockDai.constructor()](contracts/Mocks/MockDai.sol#L10-L13) does not emit an event

contracts/Mocks/MockDai.sol#L10-L13


 - [ ] ID-52
Setter function [OngoingBountyV1.claimOngoingPayout(address,bytes)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L96-L112) does not emit an event

contracts/Bounty/Implementations/OngoingBountyV1.sol#L96-L112


 - [ ] ID-53
Setter function [DepositManagerV1.initialize()](contracts/DepositManager/Implementations/DepositManagerV1.sol#L15-L18) does not emit an event

contracts/DepositManager/Implementations/DepositManagerV1.sol#L15-L18


 - [ ] ID-54
Setter function [AtomicBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L27-L84) does not emit an event

contracts/Bounty/Implementations/AtomicBountyV1.sol#L27-L84


 - [ ] ID-55
Setter function [MockLink.constructor()](contracts/Mocks/MockLink.sol#L10-L13) does not emit an event

contracts/Mocks/MockLink.sol#L10-L13


 - [ ] ID-56
Setter function [TieredPercentageBountyV1.claimTiered(address,uint256,address)](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L104-L120) does not emit an event

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L104-L120


 - [ ] ID-57
Setter function [OngoingBountyV1.closeOngoing(address)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L116-L125) does not emit an event

contracts/Bounty/Implementations/OngoingBountyV1.sol#L116-L125


 - [ ] ID-58
Setter function [OpenQV1.transferOracle(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L455-L458) does not emit an event

contracts/OpenQ/Implementations/OpenQV1.sol#L455-L458


 - [ ] ID-59
Setter function [MockNft.safeMint(address)](contracts/Mocks/MockNft.sol#L19-L24) does not emit an event

contracts/Mocks/MockNft.sol#L19-L24


 - [ ] ID-60
Setter function [OpenQV1._authorizeUpgrade(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L445) does not emit an event

contracts/OpenQ/Implementations/OpenQV1.sol#L445


 - [ ] ID-61
Setter function [AtomicBountyV1.setSupportingDocumentsComplete(bytes)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L165-L171) does not emit an event

contracts/Bounty/Implementations/AtomicBountyV1.sol#L165-L171


 - [ ] ID-62
Setter function [TieredPercentageBountyV1.closeCompetition()](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L123-L136) does not emit an event

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L123-L136


 - [ ] ID-63
Setter function [AtomicBountyV1.setInvoiceComplete(bytes)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L157-L160) does not emit an event

contracts/Bounty/Implementations/AtomicBountyV1.sol#L157-L160


 - [ ] ID-64
Setter function [OpenQV1.setBountyFactory(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L63-L69) does not emit an event

contracts/OpenQ/Implementations/OpenQV1.sol#L63-L69


 - [ ] ID-65
Setter function [BountyFactory.mintBounty(string,address,string,address,address,OpenQDefinitions.InitOperation)](contracts/BountyFactory/BountyFactory.sol#L49-L88) does not emit an event

contracts/BountyFactory/BountyFactory.sol#L49-L88


 - [ ] ID-66
Setter function [OngoingBountyV1.receiveNft(address,address,uint256,uint256,bytes)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L133-L160) does not emit an event

contracts/Bounty/Implementations/OngoingBountyV1.sol#L133-L160


 - [ ] ID-67
Setter function [OngoingBountyV1.setInvoiceComplete(bytes)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L176-L183) does not emit an event

contracts/Bounty/Implementations/OngoingBountyV1.sol#L176-L183


 - [ ] ID-68
Setter function [TieredFixedBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L27-L86) does not emit an event

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L27-L86


 - [ ] ID-69
Setter function [TieredFixedBountyV1.claimTieredFixed(address,uint256)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L91-L107) does not emit an event

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L91-L107


 - [ ] ID-70
Setter function [AtomicBountyV1.receiveNft(address,address,uint256,uint256,bytes)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L125-L152) does not emit an event

contracts/Bounty/Implementations/AtomicBountyV1.sol#L125-L152


 - [ ] ID-71
Setter function [TokenFeeToken.constructor()](contracts/Mocks/TransferFeeToken.sol#L10-L13) does not emit an event

contracts/Mocks/TransferFeeToken.sol#L10-L13


## initialize-permission
Impact: Low
Confidence: Medium
 - [ ] ID-72
Condition variable is not initialized found in [Initializable._disableInitializers()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L131-L137)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L131-L137


## missing-zero-check
Impact: Low
Confidence: Medium
 - [ ] ID-73
variable lacks a zero-check on 		- [ClaimManagerV1.claimBounty(address,address,bytes)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L31-L67)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L31-L67


 - [ ] ID-74
variable lacks a zero-check on 		- [DepositManagerV1.isWhitelisted(address)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L200-L202)

contracts/DepositManager/Implementations/DepositManagerV1.sol#L200-L202


 - [ ] ID-75
variable lacks a zero-check on 		- [ClaimManagerV1.hasKYC(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L390-L392)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L390-L392


 - [ ] ID-76
variable lacks a zero-check on 		- [OngoingBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L27-L90)

contracts/Bounty/Implementations/OngoingBountyV1.sol#L27-L90


 - [ ] ID-77
variable lacks a zero-check on 		- [TieredPercentageBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L26-L97)

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L26-L97


 - [ ] ID-78
variable lacks a zero-check on 		- [OpenQV1.setClaimManager(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L73-L79)

contracts/OpenQ/Implementations/OpenQV1.sol#L73-L79


 - [ ] ID-79
variable lacks a zero-check on 		- [TieredFixedBountyV1.setFundingGoal(address,uint256)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L123-L133)

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L123-L133


 - [ ] ID-80
variable lacks a zero-check on 		- [TieredFixedBountyV1.setPayoutScheduleFixed(uint256[],address)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L138-L171)

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L138-L171


 - [ ] ID-81
variable lacks a zero-check on 		- [OpenQV1.setDepositManager(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L83-L89)

contracts/OpenQ/Implementations/OpenQV1.sol#L83-L89


 - [ ] ID-82
variable lacks a zero-check on 		- [OpenQV1.setFundingGoal(string,address,uint256)](contracts/OpenQ/Implementations/OpenQV1.sol#L117-L136)

contracts/OpenQ/Implementations/OpenQV1.sol#L117-L136


 - [ ] ID-83
variable lacks a zero-check on 		- [ClaimManagerV1.setOpenQ(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L378-L380)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L378-L380


 - [ ] ID-84
variable lacks a zero-check on 		- [BountyFactory.constructor(address,address,address,address,address)](contracts/BountyFactory/BountyFactory.sol#L26-L39)

contracts/BountyFactory/BountyFactory.sol#L26-L39


 - [ ] ID-85
variable lacks a zero-check on 		- [TieredFixedBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L27-L86)

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L27-L86


 - [ ] ID-86
variable lacks a zero-check on 		- [BountyCore.setFundingGoal(address,uint256)](contracts/Bounty/Implementations/BountyCore.sol#L141-L149)

contracts/Bounty/Implementations/BountyCore.sol#L141-L149


 - [ ] ID-87
variable lacks a zero-check on 		- [OngoingBountyV1.setPayout(address,uint256)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L165-L171)

contracts/Bounty/Implementations/OngoingBountyV1.sol#L165-L171


 - [ ] ID-88
variable lacks a zero-check on 		- [OpenQV1.associateExternalIdToAddress(string,address)](contracts/OpenQ/Implementations/OpenQV1.sol#L464-L488)

contracts/OpenQ/Implementations/OpenQV1.sol#L464-L488


 - [ ] ID-89
variable lacks a zero-check on 		- [DepositManagerV1.refundDeposit(address,bytes32)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L152-L195)

contracts/DepositManager/Implementations/DepositManagerV1.sol#L152-L195


 - [ ] ID-90
variable lacks a zero-check on 		- [OpenQV1.setPayout(string,address,uint256)](contracts/OpenQ/Implementations/OpenQV1.sol#L256-L275)

contracts/OpenQ/Implementations/OpenQV1.sol#L256-L275


 - [ ] ID-91
variable lacks a zero-check on 		- [AtomicBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L27-L84)

contracts/Bounty/Implementations/AtomicBountyV1.sol#L27-L84


## centralized-risk-informational
Impact: Informational
Confidence: High
 - [ ] ID-92
	- [BountyCore.refundDeposit(bytes32,address,uint256)](contracts/Bounty/Implementations/BountyCore.sol#L64-L93)

contracts/Bounty/Implementations/BountyCore.sol#L64-L93


 - [ ] ID-93
	- [OpenQV1.setClaimManager(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L73-L79)

contracts/OpenQ/Implementations/OpenQV1.sol#L73-L79


 - [ ] ID-94
	- [TokenWhitelist.removeToken(address)](contracts/TokenWhitelist/TokenWhitelist.sol#L36-L43)

contracts/TokenWhitelist/TokenWhitelist.sol#L36-L43


 - [ ] ID-95
	- [TieredPercentageBountyV1.setPayoutSchedule(uint256[])](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L141-L179)

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L141-L179


 - [ ] ID-96
	- [BountyCore.setKycRequired(bool)](contracts/Bounty/Implementations/BountyCore.sol#L153-L155)

contracts/Bounty/Implementations/BountyCore.sol#L153-L155


 - [ ] ID-97
	- [OngoingBountyV1.claimOngoingPayout(address,bytes)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L96-L112)

contracts/Bounty/Implementations/OngoingBountyV1.sol#L96-L112


 - [ ] ID-98
	- [DepositManagerV1.setTokenWhitelist(address)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L22-L28)

contracts/DepositManager/Implementations/DepositManagerV1.sol#L22-L28


 - [ ] ID-99
	- [OpenQV1.mintBounty(string,string,OpenQDefinitions.InitOperation)](contracts/OpenQ/Implementations/OpenQV1.sol#L26-L59)

contracts/OpenQ/Implementations/OpenQV1.sol#L26-L59


 - [ ] ID-100
	- [ClaimManagerV1.setKyc(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L384-L386)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L384-L386


 - [ ] ID-101
	- [TieredFixedBountyV1.setFundingGoal(address,uint256)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L123-L133)

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L123-L133


 - [ ] ID-102
	- [TieredBountyCore.setSupportingDocumentsComplete(bytes)](contracts/Bounty/Implementations/TieredBountyCore.sol#L80-L89)

contracts/Bounty/Implementations/TieredBountyCore.sol#L80-L89


 - [ ] ID-103
	- [BountyCore.setFundingGoal(address,uint256)](contracts/Bounty/Implementations/BountyCore.sol#L141-L149)

contracts/Bounty/Implementations/BountyCore.sol#L141-L149


 - [ ] ID-104
	- [OngoingBountyV1.setSupportingDocumentsComplete(bytes)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L188-L198)

contracts/Bounty/Implementations/OngoingBountyV1.sol#L188-L198


 - [ ] ID-105
	- [AtomicBountyV1.close(address,bytes)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L104-L117)

contracts/Bounty/Implementations/AtomicBountyV1.sol#L104-L117


 - [ ] ID-106
	- [TokenWhitelist.setTokenAddressLimit(uint256)](contracts/TokenWhitelist/TokenWhitelist.sol#L47-L52)

contracts/TokenWhitelist/TokenWhitelist.sol#L47-L52


 - [ ] ID-107
	- [BountyCore.extendDeposit(bytes32,uint256,address)](contracts/Bounty/Implementations/BountyCore.sol#L99-L120)

contracts/Bounty/Implementations/BountyCore.sol#L99-L120


 - [ ] ID-108
	- [OngoingBountyV1.closeOngoing(address)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L116-L125)

contracts/Bounty/Implementations/OngoingBountyV1.sol#L116-L125


 - [ ] ID-109
	- [BountyCore.receiveFunds(address,address,uint256,uint256)](contracts/Bounty/Implementations/BountyCore.sol#L21-L58)

contracts/Bounty/Implementations/BountyCore.sol#L21-L58


 - [ ] ID-110
	- [OngoingBountyV1.setInvoiceComplete(bytes)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L176-L183)

contracts/Bounty/Implementations/OngoingBountyV1.sol#L176-L183


 - [ ] ID-111
	- [TieredBountyCore.setTierClaimed(uint256)](contracts/Bounty/Implementations/TieredBountyCore.sol#L52-L54)

contracts/Bounty/Implementations/TieredBountyCore.sol#L52-L54


 - [ ] ID-112
	- [OpenQV1.associateExternalIdToAddress(string,address)](contracts/OpenQ/Implementations/OpenQV1.sol#L464-L488)

contracts/OpenQ/Implementations/OpenQV1.sol#L464-L488


 - [ ] ID-113
	- [TieredPercentageBountyV1.closeCompetition()](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L123-L136)

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L123-L136


 - [ ] ID-114
	- [AtomicBountyV1.setInvoiceComplete(bytes)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L157-L160)

contracts/Bounty/Implementations/AtomicBountyV1.sol#L157-L160


 - [ ] ID-115
	- [TieredBountyCore.receiveNft(address,address,uint256,uint256,bytes)](contracts/Bounty/Implementations/TieredBountyCore.sol#L18-L48)

contracts/Bounty/Implementations/TieredBountyCore.sol#L18-L48


 - [ ] ID-116
	- [TieredBountyCore.setTierWinner(string,uint256)](contracts/Bounty/Implementations/TieredBountyCore.sol#L59-L64)

contracts/Bounty/Implementations/TieredBountyCore.sol#L59-L64


 - [ ] ID-117
	- [TieredFixedBountyV1.setPayoutScheduleFixed(uint256[],address)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L138-L171)

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L138-L171


 - [ ] ID-118
	- [TieredFixedBountyV1.closeCompetition()](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L110-L118)

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L110-L118


 - [ ] ID-119
	- [TokenWhitelist.addToken(address)](contracts/TokenWhitelist/TokenWhitelist.sol#L25-L32)

contracts/TokenWhitelist/TokenWhitelist.sol#L25-L32


 - [ ] ID-120
	- [ClaimManagerV1.setOpenQ(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L378-L380)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L378-L380


 - [ ] ID-121
	- [OngoingBountyV1.setPayout(address,uint256)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L165-L171)

contracts/Bounty/Implementations/OngoingBountyV1.sol#L165-L171


 - [ ] ID-122
	- [AtomicBountyV1.receiveNft(address,address,uint256,uint256,bytes)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L125-L152)

contracts/Bounty/Implementations/AtomicBountyV1.sol#L125-L152


 - [ ] ID-123
	- [OngoingBountyV1.receiveNft(address,address,uint256,uint256,bytes)](contracts/Bounty/Implementations/OngoingBountyV1.sol#L133-L160)

contracts/Bounty/Implementations/OngoingBountyV1.sol#L133-L160


 - [ ] ID-124
	- [TieredBountyCore.setInvoiceComplete(bytes)](contracts/Bounty/Implementations/TieredBountyCore.sol#L69-L75)

contracts/Bounty/Implementations/TieredBountyCore.sol#L69-L75


 - [ ] ID-125
	- [AtomicBountyV1.setSupportingDocumentsComplete(bytes)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L165-L171)

contracts/Bounty/Implementations/AtomicBountyV1.sol#L165-L171


 - [ ] ID-126
	- [OpenQV1.setBountyFactory(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L63-L69)

contracts/OpenQ/Implementations/OpenQV1.sol#L63-L69


 - [ ] ID-127
	- [BountyCore.setInvoiceRequired(bool)](contracts/Bounty/Implementations/BountyCore.sol#L159-L165)

contracts/Bounty/Implementations/BountyCore.sol#L159-L165


 - [ ] ID-128
	- [OpenQV1.setDepositManager(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L83-L89)

contracts/OpenQ/Implementations/OpenQV1.sol#L83-L89


 - [ ] ID-129
	- [BountyCore.setSupportingDocumentsRequired(bool)](contracts/Bounty/Implementations/BountyCore.sol#L169-L175)

contracts/Bounty/Implementations/BountyCore.sol#L169-L175


## centralized-risk-other
Impact: Informational
Confidence: High
 - [ ] ID-130
	- [DepositManagerV1.refundDeposit(address,bytes32)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L152-L195)

contracts/DepositManager/Implementations/DepositManagerV1.sol#L152-L195


 - [ ] ID-131
	- [Ownable.renounceOwnership()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63


 - [ ] ID-132
	- [OpenQV1.setInvoiceComplete(string,bytes)](contracts/OpenQ/Implementations/OpenQV1.sol#L207-L226)

contracts/OpenQ/Implementations/OpenQV1.sol#L207-L226


 - [ ] ID-133
	- [OpenQV1.setKycRequired(string,bool)](contracts/OpenQ/Implementations/OpenQV1.sol#L141-L158)

contracts/OpenQ/Implementations/OpenQV1.sol#L141-L158


 - [ ] ID-134
	- [Ownable.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72


 - [ ] ID-135
	- [ClaimManagerV1.transferOracle(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L372-L375)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L372-L375


 - [ ] ID-136
	- [OpenQV1.setPayoutSchedule(string,uint256[])](contracts/OpenQ/Implementations/OpenQV1.sol#L281-L299)

contracts/OpenQ/Implementations/OpenQV1.sol#L281-L299


 - [ ] ID-137
	- [OpenQV1.setFundingGoal(string,address,uint256)](contracts/OpenQ/Implementations/OpenQV1.sol#L117-L136)

contracts/OpenQ/Implementations/OpenQV1.sol#L117-L136


 - [ ] ID-138
	- [OwnableUpgradeable.transferOwnership(address)](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77


 - [ ] ID-139
	- [TieredFixedBountyV1.claimTieredFixed(address,uint256)](contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L91-L107)

contracts/Bounty/Implementations/TieredFixedBountyV1.sol#L91-L107


 - [ ] ID-140
	- [UUPSUpgradeable.upgradeToAndCall(address,bytes)](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L85-L88)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L85-L88


 - [ ] ID-141
	- [ClaimManagerV1.permissionedClaimTieredBounty(address,bytes)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L75-L116)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L75-L116


 - [ ] ID-142
	- [DepositManagerV1.fundBountyToken(address,address,uint256,uint256,string)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L36-L74)

contracts/DepositManager/Implementations/DepositManagerV1.sol#L36-L74


 - [ ] ID-143
	- [OwnableUpgradeable.renounceOwnership()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68


 - [ ] ID-144
	- [OpenQV1.setInvoiceRequired(string,bool)](contracts/OpenQ/Implementations/OpenQV1.sol#L163-L180)

contracts/OpenQ/Implementations/OpenQV1.sol#L163-L180


 - [ ] ID-145
	- [DepositManagerV1.extendDeposit(address,bytes32,uint256)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L80-L105)

contracts/DepositManager/Implementations/DepositManagerV1.sol#L80-L105


 - [ ] ID-146
	- [OpenQV1.setPayout(string,address,uint256)](contracts/OpenQ/Implementations/OpenQV1.sol#L256-L275)

contracts/OpenQ/Implementations/OpenQV1.sol#L256-L275


 - [ ] ID-147
	- [UpgradeableBeacon.upgradeTo(address)](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L49-L52)

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L49-L52


 - [ ] ID-148
	- [OpenQV1.setPayoutScheduleFixed(string,uint256[],address)](contracts/OpenQ/Implementations/OpenQV1.sol#L305-L324)

contracts/OpenQ/Implementations/OpenQV1.sol#L305-L324


 - [ ] ID-149
	- [ClaimManagerV1.initialize(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L20-L24)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L20-L24


 - [ ] ID-150
	- [OpenQV1.setSupportingDocumentsComplete(string,bytes)](contracts/OpenQ/Implementations/OpenQV1.sol#L231-L250)

contracts/OpenQ/Implementations/OpenQV1.sol#L231-L250


 - [ ] ID-151
	- [OpenQV1.transferOracle(address)](contracts/OpenQ/Implementations/OpenQV1.sol#L455-L458)

contracts/OpenQ/Implementations/OpenQV1.sol#L455-L458


 - [ ] ID-152
	- [TieredPercentageBountyV1.claimTiered(address,uint256,address)](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L104-L120)

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L104-L120


 - [ ] ID-153
	- [OpenQV1.initialize()](contracts/OpenQ/Implementations/OpenQV1.sol#L14-L18)

contracts/OpenQ/Implementations/OpenQV1.sol#L14-L18


 - [ ] ID-154
	- [BountyFactory.mintBounty(string,address,string,address,address,OpenQDefinitions.InitOperation)](contracts/BountyFactory/BountyFactory.sol#L49-L88)

contracts/BountyFactory/BountyFactory.sol#L49-L88


 - [ ] ID-155
	- [OpenQV1.setSupportingDocumentsRequired(string,bool)](contracts/OpenQ/Implementations/OpenQV1.sol#L185-L202)

contracts/OpenQ/Implementations/OpenQV1.sol#L185-L202


 - [ ] ID-156
	- [DepositManagerV1.fundBountyNFT(address,address,uint256,uint256,bytes)](contracts/DepositManager/Implementations/DepositManagerV1.sol#L113-L147)

contracts/DepositManager/Implementations/DepositManagerV1.sol#L113-L147


 - [ ] ID-157
	- [AtomicBountyV1.claimBalance(address,address)](contracts/Bounty/Implementations/AtomicBountyV1.sol#L89-L98)

contracts/Bounty/Implementations/AtomicBountyV1.sol#L89-L98


 - [ ] ID-158
	- [ClaimManagerV1.claimBounty(address,address,bytes)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L31-L67)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L31-L67


 - [ ] ID-159
	- [UUPSUpgradeable.upgradeTo(address)](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L72-L75)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L72-L75


 - [ ] ID-160
	- [BountyCore.claimNft(address,bytes32)](contracts/Bounty/Implementations/BountyCore.sol#L125-L136)

contracts/Bounty/Implementations/BountyCore.sol#L125-L136


 - [ ] ID-161
	- [MockNft.safeMint(address)](contracts/Mocks/MockNft.sol#L19-L24)

contracts/Mocks/MockNft.sol#L19-L24


## uncontrolled-resource-consumption
Impact: Informational
Confidence: Medium
 - [ ] ID-162
Potential DoS Gas Limit Attack occur in[ClaimManagerV1._claimTieredPercentageBounty(IBounty,address,bytes)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L203-L272)[BEGIN_LOOP](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L230-L249)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L203-L272


 - [ ] ID-163
Potential DoS Gas Limit Attack occur in[ClaimManagerV1._claimTieredFixedBounty(IBounty,address,bytes)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L278-L341)[BEGIN_LOOP](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L320-L338)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L278-L341


 - [ ] ID-164
Potential DoS Gas Limit Attack occur in[ClaimManagerV1._claimAtomicBounty(IAtomicBounty,address,bytes)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L123-L166)[BEGIN_LOOP](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L130-L148)

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L123-L166


 - [ ] ID-165
Potential DoS Gas Limit Attack occur in[TieredPercentageBountyV1.initialize(string,address,string,address,address,address,OpenQDefinitions.InitOperation)](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L26-L97)[BEGIN_LOOP](contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L78-L80)

contracts/Bounty/Implementations/TieredPercentageBountyV1.sol#L26-L97


## unnecessary-public-function-modifier
Impact: Informational
Confidence: High
 - [ ] ID-166
function:[ERC721.ownerOf(uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L70-L74)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L70-L74


 - [ ] ID-167
function:[ERC721.name()](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L79-L81)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L79-L81


 - [ ] ID-168
function:[OwnableUpgradeable.renounceOwnership()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68


 - [ ] ID-169
function:[ERC20.symbol()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L70-L72)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L70-L72


 - [ ] ID-170
function:[OngoingBountyV1.getClaimIds()](contracts/Bounty/Implementations/OngoingBountyV1.sol#L224-L226)is public and can be replaced with external 

contracts/Bounty/Implementations/OngoingBountyV1.sol#L224-L226


 - [ ] ID-171
function:[ERC20.totalSupply()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L94-L96)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L94-L96


 - [ ] ID-172
function:[ERC721.balanceOf(address)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L62-L65)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L62-L65


 - [ ] ID-173
function:[ERC20.approve(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L136-L140)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L136-L140


 - [ ] ID-174
function:[ERC20.balanceOf(address)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L101-L103)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L101-L103


 - [ ] ID-175
function:[OwnableUpgradeable.transferOwnership(address)](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77


 - [ ] ID-176
function:[ERC721.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L150-L159)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L150-L159


 - [ ] ID-177
function:[ERC20.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L158-L167)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L158-L167


 - [ ] ID-178
function:[UpgradeableBeacon.upgradeTo(address)](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L49-L52)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L49-L52


 - [ ] ID-179
function:[Ownable.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72


 - [ ] ID-180
function:[ERC20.name()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L62-L64)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L62-L64


 - [ ] ID-181
function:[ERC721.symbol()](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L86-L88)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L86-L88


 - [ ] ID-182
function:[BountyCore.getLockedFunds(address)](contracts/Bounty/Implementations/BountyCore.sol#L333-L352)is public and can be replaced with external 

contracts/Bounty/Implementations/BountyCore.sol#L333-L352


 - [ ] ID-183
function:[Ownable.renounceOwnership()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63


 - [ ] ID-184
function:[ERC20.increaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L181-L185)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L181-L185


 - [ ] ID-185
function:[UpgradeableBeacon.implementation()](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L35-L37)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L35-L37


 - [ ] ID-186
function:[ERC721HolderUpgradeable.onERC721Received(address,address,uint256,bytes)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#L26-L33)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#L26-L33


 - [ ] ID-187
function:[ERC165.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L26-L28)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L26-L28


 - [ ] ID-188
function:[ERC721.safeTransferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L164-L170)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L164-L170


 - [ ] ID-189
function:[ERC721.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L52-L57)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L52-L57


 - [ ] ID-190
function:[ERC721.tokenURI(uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L93-L98)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L93-L98


 - [ ] ID-191
function:[OnlyOpenQ.openQ()](contracts/OnlyOpenQ/OnlyOpenQ.sol#L20-L22)is public and can be replaced with external 

contracts/OnlyOpenQ/OnlyOpenQ.sol#L20-L22


 - [ ] ID-192
function:[BountyCore.getTokenBalance(address)](contracts/Bounty/Implementations/BountyCore.sol#L275-L286)is public and can be replaced with external 

contracts/Bounty/Implementations/BountyCore.sol#L275-L286


 - [ ] ID-193
function:[BountyCore.getTokenAddresses()](contracts/Bounty/Implementations/BountyCore.sol#L315-L322)is public and can be replaced with external 

contracts/Bounty/Implementations/BountyCore.sol#L315-L322


 - [ ] ID-194
function:[ERC721.approve(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L112-L122)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L112-L122


 - [ ] ID-195
function:[ERC20.decreaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L201-L210)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L201-L210


 - [ ] ID-196
function:[ERC20.transfer(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L113-L117)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L113-L117


 - [ ] ID-197
function:[ERC721URIStorage.tokenURI(uint256)](node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L20-L36)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L20-L36


 - [ ] ID-198
function:[MockNft.tokenURI(uint256)](contracts/Mocks/MockNft.sol#L35-L42)is public and can be replaced with external 

contracts/Mocks/MockNft.sol#L35-L42


 - [ ] ID-199
function:[ERC20.decimals()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L87-L89)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L87-L89


 - [ ] ID-200
function:[ERC721.setApprovalForAll(address,bool)](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L136-L138)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L136-L138


 - [ ] ID-201
function:[ClaimManagerV1.bountyIsClaimable(address)](contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L345-L365)is public and can be replaced with external 

contracts/ClaimManager/Implementations/ClaimManagerV1.sol#L345-L365


## version-only
Impact: Informational
Confidence: High
 - [ ] ID-202
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/ERC721.sol#L4


 - [ ] ID-203
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/extensions/IERC721Metadata.sol#L4


 - [ ] ID-204
	Pragma confirmed better, here is pragma solidity^0.8.1. [^0.8.1](node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#L4


 - [ ] ID-205
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#L4


 - [ ] ID-206
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol#L4


 - [ ] ID-207
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L4)

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L4


 - [ ] ID-208
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/Proxy.sol#L4)

node_modules/@openzeppelin/contracts/proxy/Proxy.sol#L4


 - [ ] ID-209
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/draft-IERC20PermitUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/draft-IERC20PermitUpgradeable.sol#L4


 - [ ] ID-210
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#L4


 - [ ] ID-211
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/StorageSlotUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/StorageSlotUpgradeable.sol#L4


 - [ ] ID-212
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/proxy/beacon/IBeaconUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/beacon/IBeaconUpgradeable.sol#L4


 - [ ] ID-213
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L4


 - [ ] ID-214
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L4


 - [ ] ID-215
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#L4


 - [ ] ID-216
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/Counters.sol#L4)

node_modules/@openzeppelin/contracts/utils/Counters.sol#L4


 - [ ] ID-217
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol#L4


 - [ ] ID-218
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol#L4


 - [ ] ID-219
	Pragma confirmed better, here is pragma solidity^0.8.1. [^0.8.1](node_modules/@openzeppelin/contracts/utils/Address.sol#L4)

node_modules/@openzeppelin/contracts/utils/Address.sol#L4


 - [ ] ID-220
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/IERC721.sol#L4


 - [ ] ID-221
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#L4


 - [ ] ID-222
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol#L4


 - [ ] ID-223
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#L4)

node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#L4


 - [ ] ID-224
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/interfaces/draft-IERC1822Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/interfaces/draft-IERC1822Upgradeable.sol#L4


 - [ ] ID-225
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol#L4


 - [ ] ID-226
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4)

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-227
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L4)

node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L4


 - [ ] ID-228
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4


 - [ ] ID-229
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L4


 - [ ] ID-230
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#L4


 - [ ] ID-231
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#L4)

node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#L4


 - [ ] ID-232
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721ReceiverUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721ReceiverUpgradeable.sol#L4


 - [ ] ID-233
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/Strings.sol#L4)

node_modules/@openzeppelin/contracts/utils/Strings.sol#L4


 - [ ] ID-234
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#L4)

node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#L4


 - [ ] ID-235
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol#L4)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol#L4


 - [ ] ID-236
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L4)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L4


 - [ ] ID-237
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#L4)

node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#L4


 - [ ] ID-238
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L4


 - [ ] ID-239
. analyzed (87 contracts with 86 detectors), 245 result(s) found
INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/2023-02-openq_scaned_output/mwe-output.json generate success.
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#L4)

node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#L4


 - [ ] ID-240
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol#L4


 - [ ] ID-241
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/ERC1967/ERC1967UpgradeUpgradeable.sol#L4


 - [ ] ID-242
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


## state-should-be-constant
Impact: Optimization
Confidence: High
 - [ ] ID-243
[MockNft.uri](contracts/Mocks/MockNft.sol#L16-L17) should be constant

contracts/Mocks/MockNft.sol#L16-L17


## unnecessary-reentrancy-lock
Impact: Optimization
Confidence: High
 - [ ] ID-244
unnecessary reentrancy lock found inAtomicBountyV1
	- [BountyCore.claimNft(address,bytes32)](contracts/Bounty/Implementations/BountyCore.sol#L125-L136)

contracts/Bounty/Implementations/BountyCore.sol#L125-L136


Execution time for Falcon: 1m35.310486243s
