'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts)
'forge config --json' running
'forge build --build-info --skip */tests/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts)

MinimalForwarderUpgradeable.execute(MinimalForwarderUpgradeable.ForwardRequest,bytes) (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#52-77) sends eth to arbitrary user
	Dangerous calls:
	- (success,returndata) = req.to.call{gas: req.gas,value: req.value}(abi.encodePacked(req.data,req.from)) (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#60-62)
Reference:  

Reentrancy in TellerV2._repayLoan(uint256,TellerV0Storage.Payment,uint256,bool) (contracts/TellerV2.sol#851-898):
	External calls:
	- _sendOrEscrowFunds(_bidId,_payment) (contracts/TellerV2.sol#887)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- (success,returndata) = target.call{value: value}(data) (node_modules/@openzeppelin/contracts/utils/Address.sol#135)
		- bid.loanDetails.lendingToken.transferFrom{gas: 100000}(_msgSenderForMarket(bid.marketplaceId),lender,_paymentAmount) (contracts/TellerV2.sol#909-947)
		- bid.loanDetails.lendingToken.safeTransferFrom(sender,address(this),_paymentAmount) (contracts/TellerV2.sol#924-928)
		- bid.loanDetails.lendingToken.approve(address(escrowVault),paymentAmountReceived) (contracts/TellerV2.sol#937-940)
		- ILoanRepaymentListener(loanRepaymentListener).repayLoanCallback{gas: 80000}(_bidId,_msgSenderForMarket(bid.marketplaceId),_payment.principal,_payment.interest) (contracts/TellerV2.sol#952-961)
	External calls sending eth:
	- _sendOrEscrowFunds(_bidId,_payment) (contracts/TellerV2.sol#887)
		- (success,returndata) = target.call{value: value}(data) (node_modules/@openzeppelin/contracts/utils/Address.sol#135)
Reference:  

Reentrancy in MarketLiquidityRewards.increaseAllocationAmount(uint256,uint256) (contracts/MarketLiquidityRewards.sol#154-163):
	External calls:
	- IERC20Upgradeable(allocatedRewards[_allocationId].rewardTokenAddress).transferFrom(msg.sender,address(this),_tokenAmount) (contracts/MarketLiquidityRewards.sol#158-159)
Reference:  

TellerASEIP712Verifier (contracts/EAS/TellerASEIP712Verifier.sol#9-128) TellerASEIP712Verifier.attest(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32) (contracts/EAS/TellerASEIP712Verifier.sol#68-101) missing protection (chainid or address(this)) against signature replay attacks.
Reference: https://swcregistry.io/docs/SWC-121

state variable: ReputationManager._delinquencies (contracts/ReputationManager.sol#18) not initialized and not written in contract but be used in contract
state variable: ReputationManager._defaults (contracts/ReputationManager.sol#19) not initialized and not written in contract but be used in contract
state variable: ReputationManager._currentDelinquencies (contracts/ReputationManager.sol#20) not initialized and not written in contract but be used in contract
state variable: ReputationManager._currentDefaults (contracts/ReputationManager.sol#21) not initialized and not written in contract but be used in contract
state variable: ERC20Permit._nonces (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#25) not initialized and not written in contract but be used in contract
state variable: ERC20Votes._checkpoints (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#37) not initialized and not written in contract but be used in contract
state variable: ERC20Votes._totalSupplyCheckpoints (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#38) not initialized and not written in contract but be used in contract
state variable: TellerV2Storage_G0._borrowerBidsActive (contracts/TellerV2Storage.sol#113) not initialized and not written in contract but be used in contract
state variable: TellerV2Storage_G1._approvedForwarderSenders (contracts/TellerV2Storage.sol#137-138) not initialized and not written in contract but be used in contract
state variable: TellerV2Storage_G0.marketRegistry (contracts/TellerV2Storage.sol#109) not initialized and not written in contract but be used in contract
state variable: TellerV2Storage_G2.lenderCommitmentForwarder (contracts/TellerV2Storage.sol#142) not initialized and not written in contract but be used in contract
state variable: LenderCommitmentForwarderMock.commitmentPrincipalAccepted (contracts/mock/LenderCommitmentForwarderMock.sol#36) not initialized and not written in contract but be used in contract
state variable: AccessControlEnumerable._roleMembers (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#16) not initialized and not written in contract but be used in contract
Reference:  

	- FlashRolloverLoan_G3.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#157-219)
	- FlashRolloverLoan_G1.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#137-199)
	- FlashRolloverLoan_G4.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#155-218)
	- FlashRolloverLoan_G5.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#156-219)
	- MarketLiquidityRewards.deallocateRewards(uint256,uint256) (contracts/MarketLiquidityRewards.sol#170-198)
	- WethMock.withdraw(uint256) (contracts/mock/WethMock.sol#42-47)
Reference:  

MarketRegistry.setMarketFeePercent(uint256,uint16) (contracts/MarketRegistry.sol#636-645) contains a tautology or contradiction:
	- require(bool,string)(_newPercent >= 0 && _newPercent <= 10000,invalid percent) (contracts/MarketRegistry.sol#640)
Reference:  

Contract locking ether found:
	Contract MarketRegistry (contracts/MarketRegistry.sol#19-1260) has payable functions:
	 - IASResolver.resolve(address,bytes,bytes,uint256,address) (contracts/interfaces/IASResolver.sol#25-31)
	 - MarketRegistry.resolve(address,bytes,bytes,uint256,address) (contracts/MarketRegistry.sol#452-471)
	But does not have a function to withdraw the ether
Reference:  

setCommitment(uint256,ILenderCommitmentForwarder.Commitment) and acceptCommitmentWithRecipient(uint256,uint256,uint256,uint256,address,address,uint16,uint32) have transaction order dependency caused by data race on state variables:LenderCommitmentForwarderMock.commitments
	- LenderCommitmentForwarderMock.setCommitment(uint256,ILenderCommitmentForwarder.Commitment) (contracts/mock/LenderCommitmentForwarderMock.sol#45-49)
	- LenderCommitmentForwarderMock.acceptCommitmentWithRecipient(uint256,uint256,uint256,uint256,address,address,uint16,uint32) (contracts/mock/LenderCommitmentForwarderMock.sol#94-115)
Reference: https://swcregistry.io/docs/SWC-114

CollateralEscrowV1._withdrawCollateral(Collateral,address,uint256,address).data (contracts/escrow/CollateralEscrowV1.sol#179) is a local variable never initialized
CollateralManager.getCollateralInfo(uint256).i (contracts/CollateralManager.sol#232) is a local variable never initialized
LenderCommitmentForwarder_G2.getRequiredCollateral(uint256,uint256,ILenderCommitmentForwarder.CommitmentCollateralType,address,address).collateralDecimals (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#595) is a local variable never initialized
LenderCommitmentForwarder_G1.getRequiredCollateral(uint256,uint256,LenderCommitmentForwarder_G1.CommitmentCollateralType,address,address).collateralDecimals (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#568) is a local variable never initialized
CollateralManager._withdraw(uint256,address).i (contracts/CollateralManager.sol#418) is a local variable never initialized
ERC20Votes._moveVotingPower(address,address,uint256).oldWeight_scope_0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#233) is a local variable never initialized
CollateralManager.deployAndDeposit(uint256).i (contracts/CollateralManager.sol#192) is a local variable never initialized
CollateralManager._checkBalances(address,Collateral[],bool).i (contracts/CollateralManager.sol#493) is a local variable never initialized
CollateralManager.commitCollateral(uint256,Collateral[]).i (contracts/CollateralManager.sol#129) is a local variable never initialized
ReputationManager.updateAccountReputation(address).i (contracts/ReputationManager.sol#81) is a local variable never initialized
CollateralManager._deposit(uint256,Collateral).data (contracts/CollateralManager.sol#382) is a local variable never initialized
CollateralEscrowV1._depositCollateral(CollateralType,address,uint256,uint256).data (contracts/escrow/CollateralEscrowV1.sol#137) is a local variable never initialized
Reference:  

CollateralManager._deposit(uint256,Collateral) (contracts/CollateralManager.sol#339-409)have ignores return value in IERC20Upgradeable(collateralInfo._collateralAddress).approve(escrowAddress,collateralInfo._amount) (contracts/CollateralManager.sol#355-358)
CollateralManager._commitCollateral(uint256,Collateral) (contracts/CollateralManager.sol#449-478)have ignores return value in collateral.collateralAddresses.add(_collateralInfo._collateralAddress) (contracts/CollateralManager.sol#467)
LenderCommitmentForwarder_G1._addBorrowersToCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#281-289)have ignores return value in commitmentBorrowersList[_commitmentId].add(_borrowerArray[i]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#286)
LenderCommitmentForwarder_G1._removeBorrowersFromCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#296-304)have ignores return value in commitmentBorrowersList[_commitmentId].remove(_borrowerArray[i]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#301)
LenderCommitmentForwarder_G2._addBorrowersToCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#248-256)have ignores return value in commitmentBorrowersList[_commitmentId].add(_borrowerArray[i]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#253)
LenderCommitmentForwarder_G2._removeBorrowersFromCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#263-271)have ignores return value in commitmentBorrowersList[_commitmentId].remove(_borrowerArray[i]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#268)
LenderCommitmentForwarder_U1._addBorrowersToCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#291-299)have ignores return value in commitmentBorrowersList[_commitmentId].add(_borrowerArray[i]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#296)
LenderCommitmentForwarder_U1._removeBorrowersFromCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#306-314)have ignores return value in commitmentBorrowersList[_commitmentId].remove(_borrowerArray[i]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#311)
CommitmentRolloverLoan.rolloverLoan(uint256,uint256,ICommitmentRolloverLoan.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#41-77)have ignores return value in lendingToken.approve(address(TELLER_V2),fundsReceived) (contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#68)
FlashRolloverLoan_G1.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#137-199)have ignores return value in IERC20Upgradeable(_flashToken).approve(address(POOL()),_flashAmount + _flashFees) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#174-177)
FlashRolloverLoan_G1._repayLoanFull(uint256,address,uint256) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#201-219)have ignores return value in IERC20Upgradeable(_principalToken).approve(address(TELLER_V2),_repayAmount) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#209-212)
FlashRolloverLoan_G3.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#157-219)have ignores return value in IERC20Upgradeable(_flashToken).approve(address(POOL()),_flashAmount + _flashFees) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#194-197)
FlashRolloverLoan_G3._repayLoanFull(uint256,address,uint256) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#235-253)have ignores return value in IERC20Upgradeable(_principalToken).approve(address(TELLER_V2),_repayAmount) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#243-246)
FlashRolloverLoan_G4.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#155-218)have ignores return value in IERC20Upgradeable(_flashToken).approve(address(POOL()),_flashAmount + _flashFees) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#193-196)
FlashRolloverLoan_G4._repayLoanFull(uint256,address,uint256) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#234-252)have ignores return value in IERC20Upgradeable(_principalToken).approve(address(TELLER_V2),_repayAmount) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#242-245)
FlashRolloverLoan_G5.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#156-219)have ignores return value in IERC20Upgradeable(_flashToken).approve(address(POOL()),_flashAmount + _flashFees) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#194-197)
FlashRolloverLoan_G5._repayLoanFull(uint256,address,uint256) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#235-253)have ignores return value in IERC20Upgradeable(_principalToken).approve(address(TELLER_V2),_repayAmount) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#243-246)
LenderCommitmentGroup_Smart.acceptFundsForAcceptBid(address,uint256,uint256,uint256,address,uint256,uint32,uint16) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#336-382)have ignores return value in principalToken.approve(address(TELLER_V2),_principalAmount) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#373)
LenderCommitmentGroup_Smart._acceptBidWithRepaymentListener(uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#384-391)have ignores return value in ITellerV2(TELLER_V2).lenderAcceptBid(_bidId) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#385)
MarketRegistry._attestStakeholderVerification(uint256,address,bytes32,bool) (contracts/MarketRegistry.sol#1118-1147)have ignores return value in markets[_marketId].verifiedLendersForMarket.add(_stakeholderAddress) (contracts/MarketRegistry.sol#1130-1132)
MarketRegistry._revokeStakeholderVerification(uint256,address,bool) (contracts/MarketRegistry.sol#1209-1235)have ignores return value in markets[_marketId].verifiedLendersForMarket.remove(_stakeholderAddress) (contracts/MarketRegistry.sol#1219-1221)
ReputationManager._addMark(address,uint256,RepMark) (contracts/ReputationManager.sol#115-127)have ignores return value in _delinquencies[_account].add(_bidId) (contracts/ReputationManager.sol#119)
ReputationManager._removeMark(address,uint256,RepMark) (contracts/ReputationManager.sol#129-139)have ignores return value in _currentDelinquencies[_account].remove(_bidId) (contracts/ReputationManager.sol#133)
TellerV2.lenderAcceptBid(uint256) (contracts/TellerV2.sol#481-576)have ignores return value in _borrowerBidsActive[bid.borrower].add(_bidId) (contracts/TellerV2.sol#569)
TellerV2._repayLoan(uint256,TellerV0Storage.Payment,uint256,bool) (contracts/TellerV2.sol#851-898)have ignores return value in _borrowerBidsActive[bid.borrower].remove(_bidId) (contracts/TellerV2.sol#874)
TellerV2._sendOrEscrowFunds(uint256,TellerV0Storage.Payment) (contracts/TellerV2.sol#901-963)have ignores return value in bid.loanDetails.lendingToken.approve(address(escrowVault),paymentAmountReceived) (contracts/TellerV2.sol#937-940)
TellerV2Autopay.autoPayLoanMinimum(uint256) (contracts/TellerV2Autopay.sol#109-143)have ignores return value in ERC20(lendingToken).approve(address(tellerV2),amountToRepayMinimum) (contracts/TellerV2Autopay.sol#137)
TellerV2Context.approveMarketForwarder(uint256,address) (contracts/TellerV2Context.sol#92-101)have ignores return value in _approvedForwarderSenders[_forwarder].add(_msgSender()) (contracts/TellerV2Context.sol#99)
TellerV2Context.renounceMarketForwarder(uint256,address) (contracts/TellerV2Context.sol#108-115)have ignores return value in _approvedForwarderSenders[_forwarder].remove(_msgSender()) (contracts/TellerV2Context.sol#112)
AccessControlEnumerable._grantRole(bytes32,address) (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#52-55)have ignores return value in _roleMembers[role].add(account) (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#54)
AccessControlEnumerable._revokeRole(bytes32,address) (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#60-63)have ignores return value in _roleMembers[role].remove(account) (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#62)
ERC1967Upgrade._upgradeToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#65-74)have ignores return value in Address.functionDelegateCall(newImplementation,data) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#72)
ERC1967Upgrade._upgradeBeaconToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#174-184)have ignores return value in Address.functionDelegateCall(IBeacon(newBeacon).implementation(),data) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#182)
Reference:  

function:ContextUpgradeable.__Context_init() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#18-19)is empty 
function:ContextUpgradeable.__Context_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#21-22)is empty 
function:EscrowVault.initialize() (contracts/EscrowVault.sol#25)is empty 
function:ERC20._beforeTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#364-368)is empty 
function:ERC20._afterTokenTransfer(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#384-388)is empty 
function:ERC721Upgradeable._afterTokenTransfer(address,address,uint256,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#502-507)is empty 
function:ERC165Upgradeable.__ERC165_init() (node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#24-25)is empty 
function:ERC165Upgradeable.__ERC165_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#27-28)is empty 
function:MarketLiquidityRewards.initialize() (contracts/MarketLiquidityRewards.sol#79)is empty 
function:MinimalForwarderUpgradeable.__MinimalForwarder_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#39)is empty 
function:CollateralManagerMock.withdraw(uint256) (contracts/mock/CollateralManagerMock.sol#74)is empty 
function:CollateralManagerMock.lenderClaimCollateral(uint256) (contracts/mock/CollateralManagerMock.sol#85)is empty 
function:CollateralManagerMock.liquidateCollateral(uint256,address) (contracts/mock/CollateralManagerMock.sol#93-95)is empty 
function:LenderCommitmentForwarderMock.createCommitment(ILenderCommitmentForwarder.Commitment,address[]) (contracts/mock/LenderCommitmentForwarderMock.sol#40-43)is empty 
function:MarketRegistryMock.initialize(TellerAS) (contracts/mock/MarketRegistryMock.sol#20)is empty 
function:MarketRegistryMock.getPaymentType(uint256) (contracts/mock/MarketRegistryMock.sol#107-111)is empty 
function:MarketRegistryMock.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,PaymentType,PaymentCycleType,string) (contracts/mock/MarketRegistryMock.sol#113-124)is empty 
function:MarketRegistryMock.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,string) (contracts/mock/MarketRegistryMock.sol#126-135)is empty 
function:MarketRegistryMock.closeMarket(uint256) (contracts/mock/MarketRegistryMock.sol#137)is empty 
function:ReputationManagerMock.initialize(address) (contracts/mock/ReputationManagerMock.sol#10)is empty 
function:ReputationManagerMock.getDelinquentLoanIds(address) (contracts/mock/ReputationManagerMock.sol#12-15)is empty 
function:ReputationManagerMock.getDefaultedLoanIds(address) (contracts/mock/ReputationManagerMock.sol#17-20)is empty 
function:ReputationManagerMock.getCurrentDelinquentLoanIds(address) (contracts/mock/ReputationManagerMock.sol#22-25)is empty 
function:ReputationManagerMock.getCurrentDefaultLoanIds(address) (contracts/mock/ReputationManagerMock.sol#27-30)is empty 
function:ReputationManagerMock.updateAccountReputation(address) (contracts/mock/ReputationManagerMock.sol#32)is empty 
function:TellerV2SolMock.liquidateLoanFull(uint256) (contracts/mock/TellerV2SolMock.sol#121)is empty 
function:TellerV2SolMock.liquidateLoanFullWithRecipient(uint256,address) (contracts/mock/TellerV2SolMock.sol#123-125)is empty 
function:TellerV2SolMock.repayLoanMinimum(uint256) (contracts/mock/TellerV2SolMock.sol#127)is empty 
function:TellerV2SolMock.getBorrowerActiveLoanIds(address) (contracts/mock/TellerV2SolMock.sol#241-245)is empty 
function:TellerV2SolMock.isLoanDefaulted(uint256) (contracts/mock/TellerV2SolMock.sol#247-252)is empty 
function:TellerV2SolMock.isLoanLiquidateable(uint256) (contracts/mock/TellerV2SolMock.sol#254-259)is empty 
function:TellerV2SolMock.isPaymentLate(uint256) (contracts/mock/TellerV2SolMock.sol#261)is empty 
function:TellerV2SolMock.getRepaymentListenerForBid(uint256) (contracts/mock/TellerV2SolMock.sol#337-341)is empty 
function:TellerV2SolMock.setRepaymentListenerForBid(uint256,address) (contracts/mock/TellerV2SolMock.sol#343-345)is empty 
function:AavePoolAddressProviderMock.setAddressAsProxy(bytes32,address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#166-168)is empty 
function:AavePoolAddressProviderMock.setPoolConfiguratorImpl(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#170-172)is empty 
function:AavePoolAddressProviderMock.setPoolImpl(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#174)is empty 
Reference:  

BokkyPooBahsDateTimeLibrary._daysToDate(uint256).L (contracts/libraries/DateTimeLib.sol#103) is written in both
	L = L - (1461 * _year) / 4 + 31 (contracts/libraries/DateTimeLib.sol#107)
	L = _month / 11 (contracts/libraries/DateTimeLib.sol#110)
BokkyPooBahsDateTimeLibrary.subMonths(uint256,uint256).year (contracts/libraries/DateTimeLib.sol#388) is written in both
	(year,month,day) = _daysToDate(timestamp / SECONDS_PER_DAY) (contracts/libraries/DateTimeLib.sol#388-390)
	year = yearMonth / 12 (contracts/libraries/DateTimeLib.sol#392)
Reference:  

	- LenderCommitmentGroup_Smart.repayLoanCallback(uint256,address,uint256,uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#700-709)
Reference:  

Setter function CollateralManager.initialize(address,address) (contracts/CollateralManager.sol#80-87) does not emit an event
Setter function CollateralManager.setCollateralEscrowBeacon(address) (contracts/CollateralManager.sol#93-98) does not emit an event
Setter function CollateralManager.commitCollateral(uint256,Collateral[]) (contracts/CollateralManager.sol#119-134) does not emit an event
Setter function CollateralManager.commitCollateral(uint256,Collateral) (contracts/CollateralManager.sol#142-152) does not emit an event
Setter function CollateralManager.liquidateCollateral(uint256,address) (contracts/CollateralManager.sol#291-303) does not emit an event
Setter function CollateralManager.onlyTellerV2() (contracts/CollateralManager.sol#68-71) does not emit an event
Setter function TellerASEIP712Verifier.attest(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32) (contracts/EAS/TellerASEIP712Verifier.sol#68-101) does not emit an event
Setter function TellerASEIP712Verifier.revoke(bytes32,address,uint8,bytes32,bytes32) (contracts/EAS/TellerASEIP712Verifier.sol#106-127) does not emit an event
Setter function EscrowVault.deposit(address,address,uint256) (contracts/EscrowVault.sol#32-41) does not emit an event
Setter function EscrowVault.withdraw(address,uint256) (contracts/EscrowVault.sol#43-48) does not emit an event
Setter function SmartCommitmentForwarder.acceptCommitmentWithRecipient(address,uint256,uint256,uint256,address,address,uint16,uint32) (contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#38-66) does not emit an event
Setter function CommitmentRolloverLoan.rolloverLoan(uint256,uint256,ICommitmentRolloverLoan.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#41-77) does not emit an event
Setter function CommitmentRolloverLoan._acceptCommitment(ICommitmentRolloverLoan.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#126-151) does not emit an event
Setter function FlashRolloverLoan_G4.rolloverLoanWithFlash(address,uint256,uint256,uint256,FlashRolloverLoan_G4.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#96-134) does not emit an event
Setter function FlashRolloverLoan_G4.onlyFlashLoanPool() (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#70-77) does not emit an event
Setter function FlashRolloverLoan_G5.rolloverLoanWithFlash(address,uint256,uint256,uint256,FlashRolloverLoan_G5.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#97-135) does not emit an event
Setter function FlashRolloverLoan_G5.onlyFlashLoanPool() (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#72-79) does not emit an event
Setter function LenderCommitmentGroupShares.mint(address,uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#17-19) does not emit an event
Setter function LenderCommitmentGroupShares.burn(address,uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#21-23) does not emit an event
Setter function LenderCommitmentGroup_Smart.initialize(address,address,uint256,uint32,uint16,uint16,uint16,uint16,uint24,uint32) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#158-213) does not emit an event
Setter function LenderCommitmentGroup_Smart._deployPoolSharesToken() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#215-239) does not emit an event
Setter function LenderCommitmentGroup_Smart.addPrincipalToCommitmentGroup(uint256,address) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#307-322) does not emit an event
Setter function LenderCommitmentGroup_Smart.acceptFundsForAcceptBid(address,uint256,uint256,uint256,address,uint256,uint32,uint16) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#336-382) does not emit an event
Setter function LenderCommitmentGroup_Smart.burnSharesToWithdrawEarnings(uint256,address) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#396-415) does not emit an event
Setter function LenderCommitmentGroup_Smart.liquidateDefaultedLoanWithIncentive(uint256,int256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#422-472) does not emit an event
Setter function LenderCommitmentGroup_Smart.repayLoanCallback(uint256,address,uint256,uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#700-709) does not emit an event
Setter function LenderCommitmentGroup_Smart.pauseBorrowing() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#793-795) does not emit an event
Setter function LenderCommitmentGroup_Smart.unpauseBorrowing() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#800-802) does not emit an event
Setter function LenderCommitmentGroup_Smart.onlySmartCommitmentForwarder() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#120-126) does not emit an event
Setter function LenderCommitmentGroup_Smart.onlyTellerV2() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#128-134) does not emit an event
Setter function LenderManager.initialize() (contracts/LenderManager.sol#26-28) does not emit an event
Setter function LenderManager.__LenderManager_init() (contracts/LenderManager.sol#30-33) does not emit an event
Setter function LenderManager.registerLoan(uint256,address) (contracts/LenderManager.sol#40-46) does not emit an event
Setter function MarketLiquidityRewards._decrementAllocatedAmount(uint256,uint256) (contracts/MarketLiquidityRewards.sol#363-367) does not emit an event
Setter function MarketLiquidityRewards.onlyMarketOwner(uint256) (contracts/MarketLiquidityRewards.sol#39-46) does not emit an event
Setter function MarketRegistry.initialize(TellerAS) (contracts/MarketRegistry.sol#106-117) does not emit an event
Setter function MarketRegistry.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,PaymentType,PaymentCycleType,string) (contracts/MarketRegistry.sol#132-156) does not emit an event
Setter function MarketRegistry.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,string) (contracts/MarketRegistry.sol#170-192) does not emit an event
Setter function MarketRegistry.attestLender(uint256,address,uint256) (contracts/MarketRegistry.sol#288-294) does not emit an event
Setter function MarketRegistry.revokeLender(uint256,address) (contracts/MarketRegistry.sol#323-325) does not emit an event
Setter function MarketRegistry.attestBorrower(uint256,address,uint256) (contracts/MarketRegistry.sol#366-372) does not emit an event
Setter function MarketRegistry.revokeBorrower(uint256,address) (contracts/MarketRegistry.sol#401-405) does not emit an event
Setter function MarketRegistry.updateMarketSettings(uint256,uint32,PaymentType,PaymentCycleType,uint32,uint32,uint16,bool,bool,string) (contracts/MarketRegistry.sol#502-526) does not emit an event
Setter function MarketRegistry._setMarketSettings(uint256,uint32,PaymentType,PaymentCycleType,uint32,uint32,uint16,bool,bool,string) (contracts/MarketRegistry.sol#965-985) does not emit an event
Setter function MarketRegistry._attestStakeholder(uint256,address,uint256,bool) (contracts/MarketRegistry.sol#1026-1057) does not emit an event
Setter function MarketRegistry._revokeStakeholder(uint256,address,bool) (contracts/MarketRegistry.sol#1156-1173) does not emit an event
Setter function MarketRegistry.ownsMarket(uint256) (contracts/MarketRegistry.sol#67-70) does not emit an event
Setter function MarketRegistry.withAttestingSchema(bytes32) (contracts/MarketRegistry.sol#72-76) does not emit an event
Setter function ProtocolFeeMock.initialize(uint16) (contracts/ProtocolFeeMock.sol#9-11) does not emit an event
Setter function ProtocolFeeMock.setProtocolFee(uint16) (contracts/ProtocolFeeMock.sol#13-25) does not emit an event
Setter function ReputationManager.initialize(address) (contracts/ReputationManager.sol#37-39) does not emit an event
Setter function TLR.mint(address,uint256) (contracts/TLR.sol#39-41) does not emit an event
Setter function TLR.burn(address,uint256) (contracts/TLR.sol#54-56) does not emit an event
Setter function TellerV2.initialize(uint16,address,address,address,address,address,address) (contracts/TellerV2.sol#188-227) does not emit an event
Setter function TellerV2._setLenderManager(address) (contracts/TellerV2.sol#234-243) does not emit an event
Setter function TellerV2._setEscrowVault(address) (contracts/TellerV2.sol#245-248) does not emit an event
Setter function TellerV2.claimLoanNFT(uint256) (contracts/TellerV2.sol#578-594) does not emit an event
Setter function TellerV2.setRepaymentListenerForBid(uint256,address) (contracts/TellerV2.sol#1248-1259) does not emit an event
Setter function TellerV2Autopay.setAutopayFee(uint16) (contracts/TellerV2Autopay.sol#70-72) does not emit an event
Setter function CollateralEscrowV1.initialize(uint256) (contracts/escrow/CollateralEscrowV1.sol#32-35) does not emit an event
Setter function CollateralEscrowV1._depositCollateral(CollateralType,address,uint256,uint256) (contracts/escrow/CollateralEscrowV1.sol#111-149) does not emit an event
Setter function CollateralManagerMock.deployAndDeposit(uint256) (contracts/mock/CollateralManagerMock.sol#37-39) does not emit an event
Setter function CollateralManagerMock.forceSetCommitCollateralValidation(bool) (contracts/mock/CollateralManagerMock.sol#97-99) does not emit an event
Setter function LenderCommitmentForwarderMock.setCommitment(uint256,ILenderCommitmentForwarder.Commitment) (contracts/mock/LenderCommitmentForwarderMock.sol#45-49) does not emit an event
Setter function LenderCommitmentForwarderMock.acceptCommitmentWithRecipient(uint256,uint256,uint256,uint256,address,address,uint16,uint32) (contracts/mock/LenderCommitmentForwarderMock.sol#94-115) does not emit an event
Setter function LenderCommitmentForwarderMock.acceptCommitmentWithRecipientAndProof(uint256,uint256,uint256,uint256,address,address,uint16,uint32,bytes32[]) (contracts/mock/LenderCommitmentForwarderMock.sol#117-139) does not emit an event
Setter function LenderManagerMock.registerLoan(uint256,address) (contracts/mock/LenderManagerMock.sol#15-20) does not emit an event
Setter function MarketRegistryMock.setMarketOwner(address) (contracts/mock/MarketRegistryMock.sol#99-101) does not emit an event
Setter function MarketRegistryMock.setMarketFeeRecipient(address) (contracts/mock/MarketRegistryMock.sol#103-105) does not emit an event
Setter function MarketRegistryMock.mock_setGlobalMarketsClosed(bool) (contracts/mock/MarketRegistryMock.sol#139-141) does not emit an event
Setter function MarketRegistryMock.mock_setBorrowerIsVerified(bool) (contracts/mock/MarketRegistryMock.sol#143-145) does not emit an event
Setter function MarketRegistryMock.mock_setLenderIsVerified(bool) (contracts/mock/MarketRegistryMock.sol#147-149) does not emit an event
Setter function TellerV2SolMock.setMarketRegistry(address) (contracts/mock/TellerV2SolMock.sol#30-32) does not emit an event
Setter function TellerV2SolMock.approveMarketForwarder(uint256,address) (contracts/mock/TellerV2SolMock.sol#43-47) does not emit an event
Setter function TellerV2SolMock.submitBid(address,uint256,uint256,uint32,uint16,string,address) (contracts/mock/TellerV2SolMock.sol#50-80) does not emit an event
Setter function TellerV2SolMock.submitBid(address,uint256,uint256,uint32,uint16,string,address,Collateral[]) (contracts/mock/TellerV2SolMock.sol#82-101) does not emit an event
Setter function TellerV2SolMock.lenderCloseLoan(uint256) (contracts/mock/TellerV2SolMock.sol#103-105) does not emit an event
Setter function TellerV2SolMock.lenderCloseLoanWithRecipient(uint256,address) (contracts/mock/TellerV2SolMock.sol#107-111) does not emit an event
Setter function TellerV2SolMock.repayLoanFull(uint256) (contracts/mock/TellerV2SolMock.sol#129-147) does not emit an event
Setter function TellerV2SolMock.repayLoan(uint256,uint256) (contracts/mock/TellerV2SolMock.sol#149-157) does not emit an event
Setter function TellerV2SolMock.lenderAcceptBid(uint256) (contracts/mock/TellerV2SolMock.sol#199-222) does not emit an event
Setter function TellerV2SolMock.setLastRepaidTimestamp(uint256,uint32) (contracts/mock/TellerV2SolMock.sol#323-325) does not emit an event
Setter function TellerV2SolMock.mock_setLoanDefaultTimestamp(uint256) (contracts/mock/TellerV2SolMock.sol#329-333) does not emit an event
Setter function TellerV2SolMock.setRepaymentListenerForBid(uint256,address) (contracts/mock/TellerV2SolMock.sol#343-345) does not emit an event
Setter function WethMock.transfer(address,uint256) (contracts/mock/WethMock.sol#59-61) does not emit an event
Setter function AavePoolAddressProviderMock.constructor(string,address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#35-38) does not emit an event
Setter function AavePoolAddressProviderMock.setMarketId(string) (contracts/mock/aave/AavePoolAddressProviderMock.sol#46-52) does not emit an event
Setter function AavePoolAddressProviderMock.setAddressAsProxy(bytes32,address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#166-168) does not emit an event
Setter function AavePoolAddressProviderMock.setPoolConfiguratorImpl(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#170-172) does not emit an event
Setter function AavePoolAddressProviderMock.setPoolImpl(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#174) does not emit an event
Setter function AavePoolMock.setShouldExecuteCallback(bool) (contracts/mock/aave/AavePoolMock.sol#12-14) does not emit an event
Setter function AavePoolMock.flashLoanSimple(address,address,uint256,bytes,uint16) (contracts/mock/aave/AavePoolMock.sol#16-52) does not emit an event
Setter function UniswapV3FactoryMock.setPoolMock(address) (contracts/mock/uniswap/UniswapV3FactoryMock.sol#11-13) does not emit an event
Setter function UniswapV3PoolMock.set_mockSqrtPriceX96(uint160) (contracts/mock/uniswap/UniswapV3PoolMock.sol#30-32) does not emit an event
Setter function UniswapV3PoolMock.set_mockToken0(address) (contracts/mock/uniswap/UniswapV3PoolMock.sol#34-36) does not emit an event
Setter function UniswapV3PoolMock.set_mockToken1(address) (contracts/mock/uniswap/UniswapV3PoolMock.sol#39-41) does not emit an event
Setter function UpgradeableBeacon._setImplementation(address) (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#61-64) does not emit an event
Setter function ERC20PresetMinterPauser.constructor(string,string) (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#38-43) does not emit an event
Setter function ERC20PresetMinterPauser.mint(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#54-57) does not emit an event
Setter function ERC20PresetMinterPauser.pause() (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#68-71) does not emit an event
Setter function ERC20PresetMinterPauser.unpause() (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#82-85) does not emit an event
Setter function ERC2771ContextUpgradeable._msgSender() (node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#25-35) does not emit an event
Setter function ERC2771ContextUpgradeable._msgData() (node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#37-43) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Condition variable is not initialized found in Initializable._disableInitializers() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#144-150)
Condition variable is not initialized found in Initializable._isInitializing() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#162-164)
Reference: initialize method should has permission check

value assignment lack of validation	LenderCommitmentForwarder_U1.createCommitmentWithUniswap(ILenderCommitmentForwarder_U1.Commitment,address[],ILenderCommitmentForwarder_U1.PoolRouteConfig[],uint16) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#182-219)commitmentUniswapPoolRoutes[commitmentId_].push(_poolRoutes[i]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#201)
value assignment lack of validation	MinimalForwarderUpgradeable.execute(MinimalForwarderUpgradeable.ForwardRequest,bytes) (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#52-77)_nonces[req.from] = req.nonce + 1 (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#58)
value assignment lack of validation	LenderCommitmentForwarderMock.setCommitment(uint256,ILenderCommitmentForwarder.Commitment) (contracts/mock/LenderCommitmentForwarderMock.sol#45-49)commitments[_commitmentId] = _commitment (contracts/mock/LenderCommitmentForwarderMock.sol#48)
value assignment lack of validation	TellerV2SolMock.setLastRepaidTimestamp(uint256,uint32) (contracts/mock/TellerV2SolMock.sol#323-325)bids[_bidId].loanDetails.lastRepaidTimestamp = _timestamp (contracts/mock/TellerV2SolMock.sol#324)
value assignment lack of validation	AavePoolMock.setShouldExecuteCallback(bool) (contracts/mock/aave/AavePoolMock.sol#12-14)shouldExecuteCallback = shouldExecute (contracts/mock/aave/AavePoolMock.sol#13)
Reference: input validation

variable lacks a zero-check on 		- CollateralManager.initialize(address,address) (contracts/CollateralManager.sol#80-87)
variable lacks a zero-check on 		- CollateralManager.setCollateralEscrowBeacon(address) (contracts/CollateralManager.sol#93-98)
variable lacks a zero-check on 		- TellerASEIP712Verifier.attest(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32) (contracts/EAS/TellerASEIP712Verifier.sol#68-101)
variable lacks a zero-check on 		- TellerASEIP712Verifier.revoke(bytes32,address,uint8,bytes32,bytes32) (contracts/EAS/TellerASEIP712Verifier.sol#106-127)
variable lacks a zero-check on 		- EscrowVault.deposit(address,address,uint256) (contracts/EscrowVault.sol#32-41)
variable lacks a zero-check on 		- EscrowVault.withdraw(address,uint256) (contracts/EscrowVault.sol#43-48)
variable lacks a zero-check on 		- LenderCommitmentForwarder_U1.getUniswapV3PoolAddress(address,address,uint24) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#706-717)
variable lacks a zero-check on 		- ExtensionsContextUpgradeable.revokeExtension(address) (contracts/LenderCommitmentForwarder/extensions/ExtensionsContextUpgradeable.sol#35-38)
variable lacks a zero-check on 		- FlashRolloverLoan_G3.rolloverLoanWithFlash(uint256,uint256,uint256,FlashRolloverLoan_G3.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#100-136)
variable lacks a zero-check on 		- FlashRolloverLoan_G3.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#157-219)
variable lacks a zero-check on 		- FlashRolloverLoan_G1.rolloverLoanWithFlash(uint256,uint256,uint256,FlashRolloverLoan_G1.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#96-132)
variable lacks a zero-check on 		- FlashRolloverLoan_G1.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#137-199)
variable lacks a zero-check on 		- FlashRolloverLoan_G4.constructor(address,address) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#62-68)
variable lacks a zero-check on 		- FlashRolloverLoan_G4.rolloverLoanWithFlash(address,uint256,uint256,uint256,FlashRolloverLoan_G4.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#96-134)
variable lacks a zero-check on 		- FlashRolloverLoan_G4.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#155-218)
variable lacks a zero-check on 		- FlashRolloverLoan_G5.constructor(address,address) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#64-70)
variable lacks a zero-check on 		- FlashRolloverLoan_G5.rolloverLoanWithFlash(address,uint256,uint256,uint256,FlashRolloverLoan_G5.AcceptCommitmentArgs) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#97-135)
variable lacks a zero-check on 		- FlashRolloverLoan_G5.executeOperation(address,uint256,uint256,address,bytes) (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#156-219)
variable lacks a zero-check on 		- LenderCommitmentGroup_Smart.constructor(address,address,address) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#143-151)
variable lacks a zero-check on 		- LenderCommitmentGroup_Smart.initialize(address,address,uint256,uint32,uint16,uint16,uint16,uint16,uint24,uint32) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#158-213)
variable lacks a zero-check on 		- LenderCommitmentGroup_Smart.addPrincipalToCommitmentGroup(uint256,address) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#307-322)
variable lacks a zero-check on 		- LenderCommitmentGroup_Smart.burnSharesToWithdrawEarnings(uint256,address) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#396-415)
variable lacks a zero-check on 		- MarketLiquidityRewards.constructor(address,address,address) (contracts/MarketLiquidityRewards.sol#69-77)
variable lacks a zero-check on 		- ReputationManager.updateAccountReputation(address) (contracts/ReputationManager.sol#77-84)
variable lacks a zero-check on 		- TellerV2Autopay.autoPayLoanMinimum(uint256) (contracts/TellerV2Autopay.sol#109-143)
variable lacks a zero-check on 		- MarketRegistryMock.setMarketOwner(address) (contracts/mock/MarketRegistryMock.sol#99-101)
variable lacks a zero-check on 		- MarketRegistryMock.setMarketFeeRecipient(address) (contracts/mock/MarketRegistryMock.sol#103-105)
variable lacks a zero-check on 		- TellerAS.attestByDelegation(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32) (contracts/EAS/TellerAS.sol#124-149)
variable lacks a zero-check on 		- TellerV2SolMock.approveMarketForwarder(uint256,address) (contracts/mock/TellerV2SolMock.sol#43-47)
variable lacks a zero-check on 		- WethMock.approve(address,uint256) (contracts/mock/WethMock.sol#53-57)
variable lacks a zero-check on 		- WethMock.transferFrom(address,address,uint256) (contracts/mock/WethMock.sol#63-83)
variable lacks a zero-check on 		- AavePoolMock.flashLoanSimple(address,address,uint256,bytes,uint16) (contracts/mock/aave/AavePoolMock.sol#16-52)
variable lacks a zero-check on 		- UniswapV3FactoryMock.setPoolMock(address) (contracts/mock/uniswap/UniswapV3FactoryMock.sol#11-13)
variable lacks a zero-check on 		- UniswapV3PoolMock.set_mockToken0(address) (contracts/mock/uniswap/UniswapV3PoolMock.sol#34-36)
variable lacks a zero-check on 		- UniswapV3PoolMock.set_mockToken1(address) (contracts/mock/uniswap/UniswapV3PoolMock.sol#39-41)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

Variable 'ERC20Votes._moveVotingPower(address,address,uint256).newWeight (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#228)' in ERC20Votes._moveVotingPower(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#221-237) potentially used before declaration: (oldWeight,newWeight) = _writeCheckpoint(_checkpoints[dst],_add,amount) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#233)
Reference:  

Reentrancy in EscrowVault.deposit(address,address,uint256) (contracts/EscrowVault.sol#32-41):
	External calls:
	- ERC20(token).safeTransferFrom(_msgSender(),address(this),amount) (contracts/EscrowVault.sol#37)
Reentrancy in AavePoolMock.flashLoanSimple(address,address,uint256,bytes,uint16) (contracts/mock/aave/AavePoolMock.sol#16-52):
	External calls:
	- IERC20(asset).transfer(receiverAddress,amount) (contracts/mock/aave/AavePoolMock.sol#25)
	- success = IFlashLoanSimpleReceiver(receiverAddress).executeOperation(asset,amount,premium,initiator,params) (contracts/mock/aave/AavePoolMock.sol#31-32)
	- IERC20(asset).transferFrom(receiverAddress,address(this),amount + premium) (contracts/mock/aave/AavePoolMock.sol#37-41)
Reference:  

LenderCommitmentGroup_Smart.bidIsActiveForGroup(uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#136-140) compares to a boolean constant:
	-require(bool,string)(activeBids[_bidId] == true,Bid is not active for group) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#137)
TellerV2.submitBid(address,uint256,uint256,uint32,uint16,string,address,Collateral[]) (contracts/TellerV2.sol#314-343) compares to a boolean constant:
	-require(bool,string)(validation == true,Collateral balance could not be validated) (contracts/TellerV2.sol#339-342)
AavePoolMock.flashLoanSimple(address,address,uint256,bytes,uint16) (contracts/mock/aave/AavePoolMock.sol#16-52) compares to a boolean constant:
	-require(bool,string)(success == true,executeOperation failed) (contracts/mock/aave/AavePoolMock.sol#34)
Reference:  

	- CollateralManager.deployAndDeposit(uint256) (contracts/CollateralManager.sol#184-206)
	- ProtocolFee.setProtocolFee(uint16) (contracts/ProtocolFee.sol#44-51)
	- ProtocolFeeMock.setProtocolFee(uint16) (contracts/ProtocolFeeMock.sol#13-25)
	- TellerV2Autopay.setAutoPayEnabled(uint256,bool) (contracts/TellerV2Autopay.sol#94-103)
	- TellerV2Context.setTrustedMarketForwarder(uint256,address) (contracts/TellerV2Context.sol#75-84)
	- CollateralEscrowV1.depositAsset(CollateralType,address,uint256,uint256) (contracts/escrow/CollateralEscrowV1.sol#51-76)
	- CollateralEscrowV1.withdraw(address,uint256,address) (contracts/escrow/CollateralEscrowV1.sol#84-103)
	- AavePoolAddressProviderMock.setAddress(bytes32,address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#60-68)
	- AavePoolAddressProviderMock.setPriceOracle(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#86-94)
	- AavePoolAddressProviderMock.setACLManager(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#102-106)
	- AavePoolAddressProviderMock.setACLAdmin(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#114-118)
	- AavePoolAddressProviderMock.setPriceOracleSentinel(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#126-137)
	- AavePoolAddressProviderMock.setPoolDataProvider(address) (contracts/mock/aave/AavePoolAddressProviderMock.sol#145-153)
Reference:  

	- OwnableUpgradeable.renounceOwnership() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#66-68)
	- OwnableUpgradeable.transferOwnership(address) (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#74-77)
	- CollateralManager.commitCollateral(uint256,Collateral[]) (contracts/CollateralManager.sol#119-134)
	- CollateralManager.commitCollateral(uint256,Collateral) (contracts/CollateralManager.sol#142-152)
	- CollateralManager.lenderClaimCollateral(uint256) (contracts/CollateralManager.sol#271-283)
	- CollateralManager.liquidateCollateral(uint256,address) (contracts/CollateralManager.sol#291-303)
	- Ownable.renounceOwnership() (node_modules/@openzeppelin/contracts/access/Ownable.sol#61-63)
	- Ownable.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable.sol#69-72)
	- LenderCommitmentGroupShares.mint(address,uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#17-19)
	- LenderCommitmentGroupShares.burn(address,uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#21-23)
	- LenderCommitmentGroup_Smart.pauseBorrowing() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#793-795)
	- LenderCommitmentGroup_Smart.unpauseBorrowing() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#800-802)
	- LenderManager.registerLoan(uint256,address) (contracts/LenderManager.sol#40-46)
	- TLR.mint(address,uint256) (contracts/TLR.sol#39-41)
	- TLR.burn(address,uint256) (contracts/TLR.sol#54-56)
	- TellerV2.cancelBid(uint256) (contracts/TellerV2.sol#428-440)
	- TellerV2.pauseProtocol() (contracts/TellerV2.sol#713-715)
	- TellerV2.unpauseProtocol() (contracts/TellerV2.sol#720-722)
	- TellerV2Autopay.setAutopayFee(uint16) (contracts/TellerV2Autopay.sol#70-72)
	- TellerV2Context.renounceMarketForwarder(uint256,address) (contracts/TellerV2Context.sol#108-115)
	- AavePoolAddressProviderMock.setMarketId(string) (contracts/mock/aave/AavePoolAddressProviderMock.sol#46-52)
	- AccessControl.grantRole(bytes32,address) (node_modules/@openzeppelin/contracts/access/AccessControl.sol#144-146)
	- AccessControl.revokeRole(bytes32,address) (node_modules/@openzeppelin/contracts/access/AccessControl.sol#159-161)
	- UpgradeableBeacon.upgradeTo(address) (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#49-52)
	- ERC20PresetMinterPauser.mint(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#54-57)
	- ERC20PresetMinterPauser.pause() (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#68-71)
	- ERC20PresetMinterPauser.unpause() (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#82-85)
Reference:  

BeaconProxy._beacon() (node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#37-39) is never used and should be removed
BeaconProxy._setBeacon(address,bytes) (node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#58-60) is never used and should be removed
Context._msgData() (node_modules/@openzeppelin/contracts/utils/Context.sol#21-23) is never used and should be removed
ContextUpgradeable.__Context_init() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#18-19) is never used and should be removed
ContextUpgradeable.__Context_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#21-22) is never used and should be removed
ContextUpgradeable._msgData() (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#27-29) is never used and should be removed
EIP712Upgradeable.__EIP712_init(string,string) (node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol#50-52) is never used and should be removed
ERC1967Upgrade._changeAdmin(address) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#133-136) is never used and should be removed
ERC1967Upgrade._getAdmin() (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#116-118) is never used and should be removed
ERC1967Upgrade._getImplementation() (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#38-40) is never used and should be removed
ERC1967Upgrade._setAdmin(address) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#123-126) is never used and should be removed
ERC1967Upgrade._setImplementation(address) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#45-48) is never used and should be removed
ERC1967Upgrade._upgradeTo(address) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#55-58) is never used and should be removed
ERC1967Upgrade._upgradeToAndCall(address,bytes,bool) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#65-74) is never used and should be removed
ERC1967Upgrade._upgradeToAndCallUUPS(address,bytes,bool) (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#81-99) is never used and should be removed
ERC2771ContextUpgradeable._msgData() (contracts/ERC2771ContextUpgradeable.sol#51-63) is never used and should be removed
Initializable._getInitializedVersion() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#155-157) is never used and should be removed
Initializable._isInitializing() (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#162-164) is never used and should be removed
LenderCommitmentForwarder_U1.getPriceX96FromSqrtPriceX96(uint160) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#806-812) is never used and should be removed
MinimalForwarderUpgradeable.__MinimalForwarder_init() (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#35-37) is never used and should be removed
MinimalForwarderUpgradeable.__MinimalForwarder_init_unchained() (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#39) is never used and should be removed
TellerV2._msgData() (contracts/TellerV2.sol#1327-1335) is never used and should be removed
TellerV2Context._msgDataForMarket(uint256) (contracts/TellerV2Context.sol#152-163) is never used and should be removed
TellerV2MarketForwarder_G1._submitBid(TellerV2MarketForwarder_G1.CreateLoanArgs,address) (contracts/TellerV2MarketForwarder_G1.sol#78-99) is never used and should be removed
TellerV2MarketForwarder_G3._acceptBidWithRepaymentListener(uint256,address,address) (contracts/TellerV2MarketForwarder_G3.sol#33-56) is never used and should be removed
Reference:  

require() missing error messages
	 - require(bool)(year >= 1970) (contracts/libraries/DateTimeLib.sol#61)
require() missing error messages
	 - require(bool)(newTimestamp >= timestamp) (contracts/libraries/DateTimeLib.sol#302)
require() missing error messages
	 - require(bool)(newTimestamp >= timestamp) (contracts/libraries/DateTimeLib.sol#324)
require() missing error messages
	 - require(bool)(newTimestamp >= timestamp) (contracts/libraries/DateTimeLib.sol#333)
require() missing error messages
	 - require(bool)(newTimestamp >= timestamp) (contracts/libraries/DateTimeLib.sol#342)
require() missing error messages
	 - require(bool)(newTimestamp >= timestamp) (contracts/libraries/DateTimeLib.sol#351)
require() missing error messages
	 - require(bool)(newTimestamp >= timestamp) (contracts/libraries/DateTimeLib.sol#360)
require() missing error messages
	 - require(bool)(newTimestamp <= timestamp) (contracts/libraries/DateTimeLib.sol#380)
require() missing error messages
	 - require(bool)(newTimestamp <= timestamp) (contracts/libraries/DateTimeLib.sol#402)
require() missing error messages
	 - require(bool)(newTimestamp <= timestamp) (contracts/libraries/DateTimeLib.sol#411)
require() missing error messages
	 - require(bool)(newTimestamp <= timestamp) (contracts/libraries/DateTimeLib.sol#420)
require() missing error messages
	 - require(bool)(newTimestamp <= timestamp) (contracts/libraries/DateTimeLib.sol#429)
require() missing error messages
	 - require(bool)(newTimestamp <= timestamp) (contracts/libraries/DateTimeLib.sol#438)
require() missing error messages
	 - require(bool)(fromTimestamp <= toTimestamp) (contracts/libraries/DateTimeLib.sol#446)
require() missing error messages
	 - require(bool)(fromTimestamp <= toTimestamp) (contracts/libraries/DateTimeLib.sol#457)
require() missing error messages
	 - require(bool)(fromTimestamp <= toTimestamp) (contracts/libraries/DateTimeLib.sol#472)
require() missing error messages
	 - require(bool)(fromTimestamp <= toTimestamp) (contracts/libraries/DateTimeLib.sol#481)
require() missing error messages
	 - require(bool)(fromTimestamp <= toTimestamp) (contracts/libraries/DateTimeLib.sol#490)
require() missing error messages
	 - require(bool)(fromTimestamp <= toTimestamp) (contracts/libraries/DateTimeLib.sol#499)
require() missing error messages
	 - require(bool)(denominator > 0) (contracts/libraries/uniswap/FullMath.sol#34)
require() missing error messages
	 - require(bool)(denominator > prod1) (contracts/libraries/uniswap/FullMath.sol#43)
require() missing error messages
	 - require(bool)(result < type()(uint256).max) (contracts/libraries/uniswap/FullMath.sol#121)
require() missing error messages
	 - require(bool)(balanceOf[msg.sender] >= wad) (contracts/mock/WethMock.sol#43)
require() missing error messages
	 - require(bool)(denominator > prod1) (node_modules/@openzeppelin/contracts/utils/math/Math.sol#78)
require() missing error messages
	 - require(bool)(denominator > prod1) (node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#78)
Reference: https://dev.to/tawseef/require-vs-assert-in-solidity-5e9d

keyIFlashRolloverLoan_G4 (contracts/interfaces/IFlashRolloverLoan_G4.sol#3-12) does not specify SPDX license identifierkeyAavePoolMock (contracts/mock/aave/AavePoolMock.sol#7-53) does not specify SPDX license identifierReference: https://certik-public-assets.s3.amazonaws.com/CertiK-Audit-for-Polylastic---Airdrop-and-Token-Swap.pdf

Potential price manipulation risk:
	- In function commitCollateral
		-- (validation_,None) = checkBalances(borrower,_collateralInfo) (contracts/CollateralManager.sol#125) have potential price manipulated risk from validation_ and call None which could influence variable:validation_
	- In function commitCollateral
		-- validation_ = _checkBalance(borrower,_collateralInfo) (contracts/CollateralManager.sol#148) have potential price manipulated risk from validation_ and call None which could influence variable:validation_
	- In function revalidateCollateral
		-- (validation_,None) = _checkBalances(borrower,collateralInfos,true) (contracts/CollateralManager.sol#165) have potential price manipulated risk from validation_ and call None which could influence variable:validation_
	- In function _checkBalances
		-- isValidated = _checkBalance(_borrowerAddress,_collateralInfo[i]) (contracts/CollateralManager.sol#494-497) have potential price manipulated risk from isValidated and call None which could influence variable:checks_
Potential price manipulation risk:
	- In function _delegate
		-- delegatorBalance = balanceOf(delegator) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#213) have potential price manipulated risk from delegatorBalance and call None which could influence variable:delegatorBalance
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

Potential DoS Gas Limit Attack occur inCollateralManager._withdraw(uint256,address) (contracts/CollateralManager.sol#416-442)BEGIN_LOOP (contracts/CollateralManager.sol#417-441)
Potential DoS Gas Limit Attack occur inCollateralManager._checkBalances(address,Collateral[],bool) (contracts/CollateralManager.sol#486-507)BEGIN_LOOP (contracts/CollateralManager.sol#493-506)
Potential DoS Gas Limit Attack occur inLenderCommitmentForwarder_G1._addBorrowersToCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#281-289)BEGIN_LOOP (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#285-287)
Potential DoS Gas Limit Attack occur inLenderCommitmentForwarder_G1._removeBorrowersFromCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#296-304)BEGIN_LOOP (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#300-302)
Potential DoS Gas Limit Attack occur inLenderCommitmentForwarder_U1.createCommitmentWithUniswap(ILenderCommitmentForwarder_U1.Commitment,address[],ILenderCommitmentForwarder_U1.PoolRouteConfig[],uint16) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#182-219)BEGIN_LOOP (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#200-202)
Potential DoS Gas Limit Attack occur inLenderCommitmentForwarder_U1._addBorrowersToCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#291-299)BEGIN_LOOP (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#295-297)
Potential DoS Gas Limit Attack occur inLenderCommitmentForwarder_U1._removeBorrowersFromCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#306-314)BEGIN_LOOP (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#310-312)
Potential DoS Gas Limit Attack occur inLenderCommitmentForwarder_G2._addBorrowersToCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#248-256)BEGIN_LOOP (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#252-254)
Potential DoS Gas Limit Attack occur inLenderCommitmentForwarder_G2._removeBorrowersFromCommitmentAllowlist(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#263-271)BEGIN_LOOP (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#267-269)
Potential DoS Gas Limit Attack occur inReputationManager.updateAccountReputation(address) (contracts/ReputationManager.sol#77-84)BEGIN_LOOP (contracts/ReputationManager.sol#81-83)
Reference: https://swcregistry.io/docs/SWC-128

function:OwnableUpgradeable.renounceOwnership() (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#66-68)is public and can be replaced with external 
function:OwnableUpgradeable.transferOwnership(address) (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#74-77)is public and can be replaced with external 
function:CollateralManager.commitCollateral(uint256,Collateral[]) (contracts/CollateralManager.sol#119-134)is public and can be replaced with external 
function:CollateralManager.commitCollateral(uint256,Collateral) (contracts/CollateralManager.sol#142-152)is public and can be replaced with external 
function:CollateralManager.getCollateralAmount(uint256,address) (contracts/CollateralManager.sol#243-251)is public and can be replaced with external 
function:TellerAS.attest(address,bytes32,uint256,bytes32,bytes) (contracts/EAS/TellerAS.sol#103-119)is public and can be replaced with external 
function:TellerAS.attestByDelegation(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32) (contracts/EAS/TellerAS.sol#124-149)is public and can be replaced with external 
function:TellerAS.revoke(bytes32) (contracts/EAS/TellerAS.sol#154-156)is public and can be replaced with external 
function:TellerAS.revokeByDelegation(bytes32,address,uint8,bytes32,bytes32) (contracts/EAS/TellerAS.sol#161-171)is public and can be replaced with external 
function:TellerAS.isAttestationActive(bytes32) (contracts/EAS/TellerAS.sol#200-211)is public and can be replaced with external 
function:EscrowVault.deposit(address,address,uint256) (contracts/EscrowVault.sol#32-41)is public and can be replaced with external 
function:LenderCommitmentForwarder_G1.createCommitment(LenderCommitmentForwarder_G1.Commitment,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#183-209)is public and can be replaced with external 
function:LenderCommitmentForwarder_G1.updateCommitment(uint256,LenderCommitmentForwarder_G1.Commitment) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#216-247)is public and can be replaced with external 
function:LenderCommitmentForwarder_G1.addCommitmentBorrowers(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#254-259)is public and can be replaced with external 
function:LenderCommitmentForwarder_G1.removeCommitmentBorrowers(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#266-274)is public and can be replaced with external 
function:LenderCommitmentForwarder_G1.deleteCommitment(uint256) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#310-317)is public and can be replaced with external 
function:TellerV2MarketForwarder_G1.getTellerV2() (contracts/TellerV2MarketForwarder_G1.sol#42-44)is public and can be replaced with external 
function:TellerV2MarketForwarder_G1.getTellerV2MarketOwner(uint256) (contracts/TellerV2MarketForwarder_G1.sol#50-52)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.createCommitmentWithUniswap(ILenderCommitmentForwarder_U1.Commitment,address[],ILenderCommitmentForwarder_U1.PoolRouteConfig[],uint16) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#182-219)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.updateCommitment(uint256,ILenderCommitmentForwarder_U1.Commitment) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#226-257)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.addCommitmentBorrowers(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#264-269)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.removeCommitmentBorrowers(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#276-284)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.deleteCommitment(uint256) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#320-327)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.acceptCommitment(uint256,uint256,uint256,uint256,address,uint16,uint32) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#371-391)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.acceptCommitmentWithProof(uint256,uint256,uint256,uint256,address,uint16,uint32,bytes32[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#454-476)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.getCommitmentUniswapPoolRoute(uint256,uint256) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#666-676)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.getAllCommitmentUniswapPoolRoutes(uint256) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#683-689)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.getCommitmentPoolOracleLtvRatio(uint256) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#696-702)is public and can be replaced with external 
function:LenderCommitmentForwarder_U1.getUniswapV3PoolAddress(address,address,uint24) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#706-717)is public and can be replaced with external 
function:TellerV2MarketForwarder_G2.getTellerV2() (contracts/TellerV2MarketForwarder_G2.sol#37-39)is public and can be replaced with external 
function:TellerV2MarketForwarder_G2.getTellerV2MarketOwner(uint256) (contracts/TellerV2MarketForwarder_G2.sol#45-47)is public and can be replaced with external 
function:LenderCommitmentForwarder_G2.createCommitment(ILenderCommitmentForwarder.Commitment,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#150-176)is public and can be replaced with external 
function:LenderCommitmentForwarder_G2.updateCommitment(uint256,ILenderCommitmentForwarder.Commitment) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#183-214)is public and can be replaced with external 
function:LenderCommitmentForwarder_G2.addCommitmentBorrowers(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#221-226)is public and can be replaced with external 
function:LenderCommitmentForwarder_G2.removeCommitmentBorrowers(uint256,address[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#233-241)is public and can be replaced with external 
function:LenderCommitmentForwarder_G2.deleteCommitment(uint256) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#277-284)is public and can be replaced with external 
function:LenderCommitmentForwarder_G2.acceptCommitment(uint256,uint256,uint256,uint256,address,uint16,uint32) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#328-348)is public and can be replaced with external 
function:LenderCommitmentForwarder_G2.acceptCommitmentWithProof(uint256,uint256,uint256,uint256,address,uint16,uint32,bytes32[]) (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#411-433)is public and can be replaced with external 
function:SmartCommitmentForwarder.acceptCommitmentWithRecipient(address,uint256,uint256,uint256,address,address,uint16,uint32) (contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#38-66)is public and can be replaced with external 
function:Ownable.renounceOwnership() (node_modules/@openzeppelin/contracts/access/Ownable.sol#61-63)is public and can be replaced with external 
function:Ownable.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable.sol#69-72)is public and can be replaced with external 
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
function:LenderCommitmentGroupShares.decimals() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#25-27)is public and can be replaced with external 
function:LenderCommitmentGroup_Smart.liquidateDefaultedLoanWithIncentive(uint256,int256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#422-472)is public and can be replaced with external 
function:LenderCommitmentGroup_Smart.getRequiredCollateral(uint256) (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#738-746)is public and can be replaced with external 
function:LenderCommitmentGroup_Smart.pauseBorrowing() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#793-795)is public and can be replaced with external 
function:LenderCommitmentGroup_Smart.unpauseBorrowing() (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#800-802)is public and can be replaced with external 
function:ERC721Upgradeable.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#57-62)is public and can be replaced with external 
function:ERC721Upgradeable.balanceOf(address) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#67-70)is public and can be replaced with external 
function:ERC721Upgradeable.ownerOf(uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#75-79)is public and can be replaced with external 
function:ERC721Upgradeable.name() (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#84-86)is public and can be replaced with external 
function:ERC721Upgradeable.symbol() (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#91-93)is public and can be replaced with external 
function:ERC721Upgradeable.tokenURI(uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#98-103)is public and can be replaced with external 
function:ERC721Upgradeable.approve(address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#117-127)is public and can be replaced with external 
function:ERC721Upgradeable.setApprovalForAll(address,bool) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#141-143)is public and can be replaced with external 
function:ERC721Upgradeable.transferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#155-164)is public and can be replaced with external 
function:ERC721Upgradeable.safeTransferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#169-175)is public and can be replaced with external 
function:LenderManager.registerLoan(uint256,address) (contracts/LenderManager.sol#40-46)is public and can be replaced with external 
function:MarketLiquidityRewards.allocateRewards(IMarketLiquidityRewards.RewardAllocation) (contracts/MarketLiquidityRewards.sol#86-116)is public and can be replaced with external 
function:MarketLiquidityRewards.updateAllocation(uint256,uint256,uint256,uint32,uint32) (contracts/MarketLiquidityRewards.sol#126-147)is public and can be replaced with external 
function:MarketLiquidityRewards.increaseAllocationAmount(uint256,uint256) (contracts/MarketLiquidityRewards.sol#154-163)is public and can be replaced with external 
function:MarketLiquidityRewards.deallocateRewards(uint256,uint256) (contracts/MarketLiquidityRewards.sol#170-198)is public and can be replaced with external 
function:MarketLiquidityRewards.getRewardTokenAmount(uint256) (contracts/MarketLiquidityRewards.sol#474-481)is public and can be replaced with external 
function:MarketRegistry.closeMarket(uint256) (contracts/MarketRegistry.sol#248-254)is public and can be replaced with external 
function:MarketRegistry.isMarketOpen(uint256) (contracts/MarketRegistry.sol#260-269)is public and can be replaced with external 
function:MarketRegistry.isMarketClosed(uint256) (contracts/MarketRegistry.sol#275-282)is public and can be replaced with external 
function:MarketRegistry.transferMarketOwnership(uint256,address) (contracts/MarketRegistry.sol#481-487)is public and can be replaced with external 
function:MarketRegistry.updateMarketSettings(uint256,uint32,PaymentType,PaymentCycleType,uint32,uint32,uint16,bool,bool,string) (contracts/MarketRegistry.sol#502-526)is public and can be replaced with external 
function:MarketRegistry.setMarketFeeRecipient(uint256,address) (contracts/MarketRegistry.sol#536-542)is public and can be replaced with external 
function:MarketRegistry.getMarketData(uint256) (contracts/MarketRegistry.sol#702-724)is public and can be replaced with external 
function:MarketRegistry.getMarketAttestationRequirements(uint256) (contracts/MarketRegistry.sol#730-742)is public and can be replaced with external 
function:MarketRegistry.getMarketOwner(uint256) (contracts/MarketRegistry.sol#749-757)is public and can be replaced with external 
function:MarketRegistry.getMarketFeeRecipient(uint256) (contracts/MarketRegistry.sol#778-791)is public and can be replaced with external 
function:MarketRegistry.getMarketURI(uint256) (contracts/MarketRegistry.sol#798-805)is public and can be replaced with external 
function:MarketRegistry.getPaymentCycle(uint256) (contracts/MarketRegistry.sol#813-823)is public and can be replaced with external 
function:MarketRegistry.getPaymentDefaultDuration(uint256) (contracts/MarketRegistry.sol#830-837)is public and can be replaced with external 
function:MarketRegistry.getPaymentType(uint256) (contracts/MarketRegistry.sol#844-851)is public and can be replaced with external 
function:MarketRegistry.getBidExpirationTime(uint256) (contracts/MarketRegistry.sol#853-860)is public and can be replaced with external 
function:MarketRegistry.getMarketplaceFee(uint256) (contracts/MarketRegistry.sol#867-874)is public and can be replaced with external 
function:MarketRegistry.isVerifiedLender(uint256,address) (contracts/MarketRegistry.sol#883-896)is public and can be replaced with external 
function:MarketRegistry.isVerifiedBorrower(uint256,address) (contracts/MarketRegistry.sol#905-918)is public and can be replaced with external 
function:MarketRegistry.getAllVerifiedLendersForMarket(uint256,uint256,uint256) (contracts/MarketRegistry.sol#927-936)is public and can be replaced with external 
function:MarketRegistry.getAllVerifiedBorrowersForMarket(uint256,uint256,uint256) (contracts/MarketRegistry.sol#945-953)is public and can be replaced with external 
function:MinimalForwarderUpgradeable.getNonce(address) (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#41-43)is public and can be replaced with external 
function:MinimalForwarderUpgradeable.execute(MinimalForwarderUpgradeable.ForwardRequest,bytes) (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#52-77)is public and can be replaced with external 
function:ProtocolFee.protocolFee() (contracts/ProtocolFee.sol#36-38)is public and can be replaced with external 
function:ReputationManager.getDelinquentLoanIds(address) (contracts/ReputationManager.sol#41-48)is public and can be replaced with external 
function:ReputationManager.getDefaultedLoanIds(address) (contracts/ReputationManager.sol#50-57)is public and can be replaced with external 
function:ReputationManager.getCurrentDelinquentLoanIds(address) (contracts/ReputationManager.sol#59-66)is public and can be replaced with external 
function:ReputationManager.getCurrentDefaultLoanIds(address) (contracts/ReputationManager.sol#68-75)is public and can be replaced with external 
function:ReputationManager.updateAccountReputation(address,uint256) (contracts/ReputationManager.sol#86-92)is public and can be replaced with external 
function:ERC20Votes.checkpoints(address,uint32) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#43-45)is public and can be replaced with external 
function:ERC20Votes.numCheckpoints(address) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#50-52)is public and can be replaced with external 
function:ERC20Votes.getVotes(address) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#64-67)is public and can be replaced with external 
function:ERC20Votes.getPastVotes(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#76-79)is public and can be replaced with external 
function:ERC20Votes.getPastTotalSupply(uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#89-92)is public and can be replaced with external 
function:ERC20Votes.delegate(address) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#139-141)is public and can be replaced with external 
function:ERC20Votes.delegateBySig(address,uint256,uint256,uint8,bytes32,bytes32) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#146-163)is public and can be replaced with external 
function:ERC20Permit.permit(address,address,uint256,uint256,uint8,bytes32,bytes32) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#49-68)is public and can be replaced with external 
function:ERC20Permit.nonces(address) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#73-75)is public and can be replaced with external 
function:TellerV2Context.hasApprovedMarketForwarder(uint256,address,address) (contracts/TellerV2Context.sol#59-67)is public and can be replaced with external 
function:TellerV2.getMetadataURI(uint256) (contracts/TellerV2.sol#255-271)is public and can be replaced with external 
function:TellerV2.submitBid(address,uint256,uint256,uint32,uint16,string,address) (contracts/TellerV2.sol#283-301)is public and can be replaced with external 
function:TellerV2.submitBid(address,uint256,uint256,uint32,uint16,string,address,Collateral[]) (contracts/TellerV2.sol#314-343)is public and can be replaced with external 
function:TellerV2.pauseProtocol() (contracts/TellerV2.sol#713-715)is public and can be replaced with external 
function:TellerV2.unpauseProtocol() (contracts/TellerV2.sol#720-722)is public and can be replaced with external 
function:TellerV2.calculateAmountOwed(uint256,uint256) (contracts/TellerV2.sol#973-994)is public and can be replaced with external 
function:TellerV2.calculateAmountDue(uint256,uint256) (contracts/TellerV2.sol#1001-1021)is public and can be replaced with external 
function:TellerV2.isPaymentLate(uint256) (contracts/TellerV2.sol#1049-1052)is public and can be replaced with external 
function:TellerV2.getLoanBorrower(uint256) (contracts/TellerV2.sol#1160-1166)is public and can be replaced with external 
function:TellerV2.getLoanDefaultTimestamp(uint256) (contracts/TellerV2.sol#1234-1246)is public and can be replaced with external 
function:TellerV2Autopay.setAutopayFee(uint16) (contracts/TellerV2Autopay.sol#70-72)is public and can be replaced with external 
function:CollateralEscrowV1.initialize(uint256) (contracts/escrow/CollateralEscrowV1.sol#32-35)is public and can be replaced with external 
function:V2Calculations.calculateNextDueDate(uint32,uint32,uint32,uint32,PaymentCycleType) (contracts/libraries/V2Calculations.sol#171-215)is public and can be replaced with external 
function:LenderCommitmentForwarderMock.setCommitment(uint256,ILenderCommitmentForwarder.Commitment) (contracts/mock/LenderCommitmentForwarderMock.sol#45-49)is public and can be replaced with external 
function:LenderCommitmentForwarderMock.getCommitmentLender(uint256) (contracts/mock/LenderCommitmentForwarderMock.sol#51-57)is public and can be replaced with external 
function:LenderCommitmentForwarderMock.getCommitmentMarketId(uint256) (contracts/mock/LenderCommitmentForwarderMock.sol#59-65)is public and can be replaced with external 
function:LenderCommitmentForwarderMock.getCommitmentAcceptedPrincipal(uint256) (contracts/mock/LenderCommitmentForwarderMock.sol#67-73)is public and can be replaced with external 
function:LenderCommitmentForwarderMock.getCommitmentMaxPrincipal(uint256) (contracts/mock/LenderCommitmentForwarderMock.sol#75-81)is public and can be replaced with external 
function:LenderCommitmentForwarderMock.acceptCommitmentWithRecipient(uint256,uint256,uint256,uint256,address,address,uint16,uint32) (contracts/mock/LenderCommitmentForwarderMock.sol#94-115)is public and can be replaced with external 
function:LenderCommitmentForwarderMock.acceptCommitmentWithRecipientAndProof(uint256,uint256,uint256,uint256,address,address,uint16,uint32,bytes32[]) (contracts/mock/LenderCommitmentForwarderMock.sol#117-139)is public and can be replaced with external 
function:LenderManagerMock.ownerOf(uint256) (contracts/mock/LenderManagerMock.sol#22-29)is public and can be replaced with external 
function:MarketRegistryMock.isVerifiedLender(uint256,address) (contracts/mock/MarketRegistryMock.sol#22-28)is public and can be replaced with external 
function:MarketRegistryMock.isMarketOpen(uint256) (contracts/mock/MarketRegistryMock.sol#30-32)is public and can be replaced with external 
function:MarketRegistryMock.isMarketClosed(uint256) (contracts/mock/MarketRegistryMock.sol#34-36)is public and can be replaced with external 
function:MarketRegistryMock.isVerifiedBorrower(uint256,address) (contracts/mock/MarketRegistryMock.sol#38-44)is public and can be replaced with external 
function:MarketRegistryMock.getMarketOwner(uint256) (contracts/mock/MarketRegistryMock.sol#46-53)is public and can be replaced with external 
function:MarketRegistryMock.getMarketFeeRecipient(uint256) (contracts/mock/MarketRegistryMock.sol#55-61)is public and can be replaced with external 
function:MarketRegistryMock.getMarketURI(uint256) (contracts/mock/MarketRegistryMock.sol#63-69)is public and can be replaced with external 
function:MarketRegistryMock.getPaymentCycle(uint256) (contracts/mock/MarketRegistryMock.sol#71-77)is public and can be replaced with external 
function:MarketRegistryMock.getPaymentDefaultDuration(uint256) (contracts/mock/MarketRegistryMock.sol#79-85)is public and can be replaced with external 
function:MarketRegistryMock.getBidExpirationTime(uint256) (contracts/mock/MarketRegistryMock.sol#87-93)is public and can be replaced with external 
function:MarketRegistryMock.getMarketplaceFee(uint256) (contracts/mock/MarketRegistryMock.sol#95-97)is public and can be replaced with external 
function:MarketRegistryMock.setMarketOwner(address) (contracts/mock/MarketRegistryMock.sol#99-101)is public and can be replaced with external 
function:MarketRegistryMock.setMarketFeeRecipient(address) (contracts/mock/MarketRegistryMock.sol#103-105)is public and can be replaced with external 
function:MarketRegistryMock.getPaymentType(uint256) (contracts/mock/MarketRegistryMock.sol#107-111)is public and can be replaced with external 
function:MarketRegistryMock.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,PaymentType,PaymentCycleType,string) (contracts/mock/MarketRegistryMock.sol#113-124)is public and can be replaced with external 
function:MarketRegistryMock.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,string) (contracts/mock/MarketRegistryMock.sol#126-135)is public and can be replaced with external 
function:MarketRegistryMock.closeMarket(uint256) (contracts/mock/MarketRegistryMock.sol#137)is public and can be replaced with external 
function:MarketRegistryMock.mock_setGlobalMarketsClosed(bool) (contracts/mock/MarketRegistryMock.sol#139-141)is public and can be replaced with external 
function:MarketRegistryMock.mock_setBorrowerIsVerified(bool) (contracts/mock/MarketRegistryMock.sol#143-145)is public and can be replaced with external 
function:MarketRegistryMock.mock_setLenderIsVerified(bool) (contracts/mock/MarketRegistryMock.sol#147-149)is public and can be replaced with external 
function:TellerASMock.isAttestationActive(bytes32) (contracts/mock/TellerASMock.sol#19-26)is public and can be replaced with external 
function:TellerV2SolMock.setMarketRegistry(address) (contracts/mock/TellerV2SolMock.sol#30-32)is public and can be replaced with external 
function:TellerV2SolMock.submitBid(address,uint256,uint256,uint32,uint16,string,address,Collateral[]) (contracts/mock/TellerV2SolMock.sol#82-101)is public and can be replaced with external 
function:TellerV2SolMock.repayLoan(uint256,uint256) (contracts/mock/TellerV2SolMock.sol#149-157)is public and can be replaced with external 
function:TellerV2SolMock.calculateAmountDue(uint256,uint256) (contracts/mock/TellerV2SolMock.sol#163-179)is public and can be replaced with external 
function:TellerV2SolMock.calculateAmountOwed(uint256,uint256) (contracts/mock/TellerV2SolMock.sol#181-197)is public and can be replaced with external 
function:TellerV2SolMock.lenderAcceptBid(uint256) (contracts/mock/TellerV2SolMock.sol#199-222)is public and can be replaced with external 
function:TellerV2SolMock.getBidState(uint256) (contracts/mock/TellerV2SolMock.sol#224-231)is public and can be replaced with external 
function:TellerV2SolMock.getLoanDetails(uint256) (contracts/mock/TellerV2SolMock.sol#233-239)is public and can be replaced with external 
function:TellerV2SolMock.getBorrowerActiveLoanIds(address) (contracts/mock/TellerV2SolMock.sol#241-245)is public and can be replaced with external 
function:TellerV2SolMock.isLoanDefaulted(uint256) (contracts/mock/TellerV2SolMock.sol#247-252)is public and can be replaced with external 
function:TellerV2SolMock.isLoanLiquidateable(uint256) (contracts/mock/TellerV2SolMock.sol#254-259)is public and can be replaced with external 
function:TellerV2SolMock.isPaymentLate(uint256) (contracts/mock/TellerV2SolMock.sol#261)is public and can be replaced with external 
function:TellerV2SolMock.setLastRepaidTimestamp(uint256,uint32) (contracts/mock/TellerV2SolMock.sol#323-325)is public and can be replaced with external 
function:TellerV2SolMock.getRepaymentListenerForBid(uint256) (contracts/mock/TellerV2SolMock.sol#337-341)is public and can be replaced with external 
function:TellerV2SolMock.setRepaymentListenerForBid(uint256,address) (contracts/mock/TellerV2SolMock.sol#343-345)is public and can be replaced with external 
function:WethMock.deposit() (contracts/mock/WethMock.sol#37-40)is public and can be replaced with external 
function:WethMock.withdraw(uint256) (contracts/mock/WethMock.sol#42-47)is public and can be replaced with external 
function:WethMock.totalSupply() (contracts/mock/WethMock.sol#49-51)is public and can be replaced with external 
function:WethMock.approve(address,uint256) (contracts/mock/WethMock.sol#53-57)is public and can be replaced with external 
function:WethMock.transfer(address,uint256) (contracts/mock/WethMock.sol#59-61)is public and can be replaced with external 
function:AavePoolMock.setShouldExecuteCallback(bool) (contracts/mock/aave/AavePoolMock.sol#12-14)is public and can be replaced with external 
function:UniswapV3FactoryMock.getPool(address,address,uint24) (contracts/mock/uniswap/UniswapV3FactoryMock.sol#4-9)is public and can be replaced with external 
function:UniswapV3FactoryMock.setPoolMock(address) (contracts/mock/uniswap/UniswapV3FactoryMock.sol#11-13)is public and can be replaced with external 
function:UniswapV3PoolMock.set_mockSqrtPriceX96(uint160) (contracts/mock/uniswap/UniswapV3PoolMock.sol#30-32)is public and can be replaced with external 
function:UniswapV3PoolMock.set_mockToken0(address) (contracts/mock/uniswap/UniswapV3PoolMock.sol#34-36)is public and can be replaced with external 
function:UniswapV3PoolMock.set_mockToken1(address) (contracts/mock/uniswap/UniswapV3PoolMock.sol#39-41)is public and can be replaced with external 
function:UniswapV3PoolMock.token0() (contracts/mock/uniswap/UniswapV3PoolMock.sol#43-45)is public and can be replaced with external 
function:UniswapV3PoolMock.token1() (contracts/mock/uniswap/UniswapV3PoolMock.sol#47-49)is public and can be replaced with external 
function:UniswapV3PoolMock.slot0() (contracts/mock/uniswap/UniswapV3PoolMock.sol#52-63)is public and can be replaced with external 
function:AccessControl.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts/access/AccessControl.sol#77-79)is public and can be replaced with external 
function:AccessControl.grantRole(bytes32,address) (node_modules/@openzeppelin/contracts/access/AccessControl.sol#144-146)is public and can be replaced with external 
function:AccessControl.revokeRole(bytes32,address) (node_modules/@openzeppelin/contracts/access/AccessControl.sol#159-161)is public and can be replaced with external 
function:AccessControl.renounceRole(bytes32,address) (node_modules/@openzeppelin/contracts/access/AccessControl.sol#179-183)is public and can be replaced with external 
function:AccessControlEnumerable.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#21-23)is public and can be replaced with external 
function:AccessControlEnumerable.getRoleMember(bytes32,uint256) (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#37-39)is public and can be replaced with external 
function:AccessControlEnumerable.getRoleMemberCount(bytes32) (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#45-47)is public and can be replaced with external 
function:UpgradeableBeacon.implementation() (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#35-37)is public and can be replaced with external 
function:UpgradeableBeacon.upgradeTo(address) (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#49-52)is public and can be replaced with external 
function:ERC20Burnable.burn(uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#20-22)is public and can be replaced with external 
function:ERC20Burnable.burnFrom(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#35-38)is public and can be replaced with external 
function:ERC20PresetMinterPauser.mint(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#54-57)is public and can be replaced with external 
function:ERC20PresetMinterPauser.pause() (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#68-71)is public and can be replaced with external 
function:ERC20PresetMinterPauser.unpause() (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#82-85)is public and can be replaced with external 
function:ERC165.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#26-28)is public and can be replaced with external 
function:ERC165Upgradeable.supportsInterface(bytes4) (node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#32-34)is public and can be replaced with external 
Reference:  

MarketRegistrySetPaymentCycleDuration(uint256,uint32) (contracts/MarketRegistry.sol#82) is never emitted in MarketRegistry (contracts/MarketRegistry.sol#19-1260)
Reference: https://certik-public-assets.s3.amazonaws.com/CertiK-Audit-for-Kitty-inu.pdf

	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/CollateralManager.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/EAS/TellerAS.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/EAS/TellerASEIP712Verifier.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/EAS/TellerASRegistry.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/EAS/TellerASResolver.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.9. ^0.8.9 (contracts/ERC2771ContextUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/EscrowVault.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/LenderCommitmentForwarderAlpha.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/LenderCommitmentForwarderStaging.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G3.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/ExtensionsContextUpgradeable.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G2.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/LenderManager.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/MarketLiquidityRewards.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/MarketRegistry.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.9. ^0.8.9 (contracts/MetaForwarder.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/ProtocolFee.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/ProtocolFeeMock.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/ReputationManager.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TLR.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TellerV0Storage.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TellerV2.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TellerV2Autopay.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TellerV2Context.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TellerV2MarketForwarder_G1.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TellerV2MarketForwarder_G2.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TellerV2MarketForwarder_G3.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/TellerV2Storage.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/Types.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/escrow/CollateralEscrowV1.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/IASRegistry.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/IASResolver.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ICollateralManager.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ICommitmentRolloverLoan.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/IEAS.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/IEASEIP712Verifier.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/IEscrowVault.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IExtensionsContext.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IFlashRolloverLoan.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IFlashRolloverLoan_G4.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ILenderCommitmentForwarder.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ILenderCommitmentForwarder_U1.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/ILenderCommitmentGroup.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ILenderManager.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ILoanRepaymentCallbacks.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ILoanRepaymentListener.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IMarketLiquidityRewards.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IMarketRegistry.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IProtocolFee.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IReputationManager.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/ISmartCommitment.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ISmartCommitmentForwarder.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/ITellerV2.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/ITellerV2Autopay.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/ITellerV2Context.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/ITellerV2MarketForwarder.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/ITellerV2Storage.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IUniswapV2Router.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/IWETH.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/aave/DataTypes.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/aave/IFlashLoanSimpleReceiver.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/aave/IPool.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/interfaces/aave/IPoolAddressesProvider.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/interfaces/escrow/ICollateralEscrowV1.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (contracts/interfaces/uniswap/IUniswapV3Factory.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.5.0. >=0.5.0 (contracts/interfaces/uniswap/IUniswapV3Pool.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.5.0. >=0.5.0 (contracts/interfaces/uniswap/pool/IUniswapV3PoolActions.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.5.0. >=0.5.0 (contracts/interfaces/uniswap/pool/IUniswapV3PoolDerivedState.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.5.0. >=0.5.0 (contracts/interfaces/uniswap/pool/IUniswapV3PoolEvents.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.5.0. >=0.5.0 (contracts/interfaces/uniswap/pool/IUniswapV3PoolImmutables.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.5.0. >=0.5.0 (contracts/interfaces/uniswap/pool/IUniswapV3PoolOwnerActions.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.5.0. >=0.5.0 (contracts/interfaces/uniswap/pool/IUniswapV3PoolState.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.6.0<0.9.0. >=0.6.0<0.9.0 (contracts/libraries/DateTimeLib.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/libraries/NumbersLib.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/libraries/V2Calculations.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/libraries/WadRayMath.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.4.0. >=0.4.0 (contracts/libraries/uniswap/FixedPoint96.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/libraries/uniswap/FullMath.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/libraries/uniswap/TickMath.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/CollateralManagerMock.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/LenderCommitmentForwarderMock.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/LenderManagerMock.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/MarketRegistryMock.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/ReputationManagerMock.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/TellerASMock.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/mock/TellerV2SolMock.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/WethMock.sol#20)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/aave/AavePoolAddressProviderMock.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (contracts/mock/aave/AavePoolMock.sol#1)
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. >=0.8.0<0.9.0 (contracts/type-imports.sol#1)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/access/AccessControl.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/access/IAccessControl.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/access/IAccessControlEnumerable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/access/Ownable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/governance/utils/IVotes.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/Proxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts/proxy/utils/Initializable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/security/Pausable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Pausable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (node_modules/@openzeppelin/contracts/utils/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Context.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Counters.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/Strings.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/cryptography/ECDSA.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/cryptography/EIP712.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/math/Math.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#5)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts/utils/structs/EnumerableSet.sol#5)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.9. ^0.8.9 (node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.2. ^0.8.2 (node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC1155/IERC1155Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/draft-IERC20PermitUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721ReceiverUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.1. ^0.8.1 (node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/MerkleProofUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol#5)
Reference:  

ERC20Permit._PERMIT_TYPEHASH_DEPRECATED_SLOT (node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#37) should be constant
LenderCommitmentForwarderMock.acceptBidWasCalled (contracts/mock/LenderCommitmentForwarderMock.sol#31) should be constant
LenderCommitmentForwarderMock.commitmentCount (contracts/mock/LenderCommitmentForwarderMock.sol#28) should be constant
LenderCommitmentForwarderMock.submitBidWasCalled (contracts/mock/LenderCommitmentForwarderMock.sol#32) should be constant
LenderCommitmentForwarderMock.submitBidWithCollateralWasCalled (contracts/mock/LenderCommitmentForwarderMock.sol#30) should be constant
MarketRegistry.version (contracts/MarketRegistry.sol#59) should be constant
TellerV2SolMock.amountOwedMockInterest (contracts/mock/TellerV2SolMock.sol#20) should be constant
TellerV2SolMock.amountOwedMockPrincipal (contracts/mock/TellerV2SolMock.sol#19) should be constant
TellerV2SolMock.globalBidPaymentCycleDuration (contracts/mock/TellerV2SolMock.sol#25) should be constant
TellerV2Storage_G0.__totalVolumeFilled (contracts/TellerV2Storage.sol#104) should be constant
TellerV2Storage_G0.version (contracts/TellerV2Storage.sol#126) should be constant
WethMock.decimals (contracts/mock/WethMock.sol#27) should be constant
WethMock.name (contracts/mock/WethMock.sol#25) should be constant
WethMock.symbol (contracts/mock/WethMock.sol#26) should be constant
Reference:  
. analyzed (171 contracts with 86 detectors), 742 result(s) found
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
Traceback (most recent call last):
  File "/usr/local/bin/falcon", line 33, in <module>
    sys.exit(load_entry_point('falcon-analyzer==0.2.28', 'console_scripts', 'falcon')())
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 712, in main
    main_impl(all_detector_classes=detectors, all_printer_classes=printers)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 905, in main_impl
    output_to_json(None if outputting_json_stdout else args.json, output_error, json_results)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/utils/output.py", line 63, in output_to_json
    with open(filename, "w", encoding="utf8") as f:
FileNotFoundError: [Errno 2] No such file or directory: '../../scanned_output/2024-04-teller-finance_autogen_output/output.json'
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/CollateralManager.sol: Smart Contract (1): 589 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/EAS/TellerAS.sol: Smart Contract (1): 514 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/EAS/TellerASEIP712Verifier.sol: Smart Contract (1): 128 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/EAS/TellerASRegistry.sol: Smart Contract (1): 81 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/EAS/TellerASResolver.sol: Abstract Contract (1): 21 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/ERC2771ContextUpgradeable.sol: Abstract Contract (1): 64 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/EscrowVault.sol: Smart Contract (1): 49 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/LenderCommitmentForwarder.sol: Smart Contract (1): 13 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/LenderCommitmentForwarderAlpha.sol: Smart Contract (1): 21 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/LenderCommitmentForwarderStaging.sol: Smart Contract (1): 17 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol: Smart Contract (1): 684 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol: Smart Contract (1): 693 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G3.sol: Smart Contract (1): 27 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol: Smart Contract (1): 901 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol: Smart Contract (1): 156 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol: Smart Contract (1): 192 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/ExtensionsContextUpgradeable.sol: Abstract Contract (1): 62 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan.sol: Smart Contract (1): 19 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol: Smart Contract (1): 274 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G2.sol: Smart Contract (1): 123 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol: Smart Contract (1): 434 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol: Smart Contract (1): 435 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol: Smart Contract (1): 467 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol: Smart Contract (1): 28 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol: Smart Contract (1): 803 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/LenderManager.sol: Smart Contract (1): 84 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/MarketLiquidityRewards.sol: Smart Contract (1): 482 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/MarketRegistry.sol: Smart Contract (1): 1260 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/MetaForwarder.sol: Smart Contract (1): 10 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/ProtocolFee.sol: Smart Contract (1): 52 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/ProtocolFeeMock.sol: Smart Contract (1): 26 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/ReputationManager.sol: Smart Contract (1): 140 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TLR.sol: Smart Contract (1): 57 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TellerV0Storage.sol: Smart Contract (1): 90 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TellerV2.sol: Smart Contract (1): 1336 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TellerV2Autopay.sol: Smart Contract (1): 157 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TellerV2Context.sol: Abstract Contract (1): 164 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TellerV2MarketForwarder_G1.sol: Abstract Contract (1): 156 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TellerV2MarketForwarder_G2.sol: Abstract Contract (1): 150 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TellerV2MarketForwarder_G3.sol: Abstract Contract (1): 60 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/TellerV2Storage.sol: Abstract Contract (8): 165 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/Types.sol: : 5 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/escrow/CollateralEscrowV1.sol: Smart Contract (1): 239 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IASRegistry.sol: Interface (1): 68 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IASResolver.sol: Interface (1): 32 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ICollateralManager.sol: Interface (1): 90 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ICommitmentRolloverLoan.sol: Interface (1): 14 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IEAS.sol: Interface (1): 307 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IEASEIP712Verifier.sol: Interface (1): 59 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IEscrowVault.sol: Interface (1): 12 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IExtensionsContext.sol: Interface (1): 13 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IFlashRolloverLoan.sol: Interface (1): 11 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IFlashRolloverLoan_G4.sol: Interface (1): 11 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ILenderCommitmentForwarder.sol: Interface (1): 92 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ILenderCommitmentForwarder_U1.sol: Interface (1): 102 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ILenderCommitmentGroup.sol: Interface (1): 30 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ILenderManager.sol: Abstract Contract (1): 13 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ILoanRepaymentCallbacks.sol: Interface (1): 13 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ILoanRepaymentListener.sol: Interface (1): 11 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IMarketLiquidityRewards.sol: Interface (1): 49 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IMarketRegistry.sol: Interface (1): 86 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IProtocolFee.sol: Interface (1): 6 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IReputationManager.sol: Interface (1): 34 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ISmartCommitment.sol: Interface (1): 54 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ISmartCommitmentForwarder.sol: Interface (1): 18 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ITellerV2.sol: Interface (1): 171 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ITellerV2Autopay.sol: Interface (1): 12 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ITellerV2Context.sol: Interface (1): 10 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ITellerV2MarketForwarder.sol: Interface (1): 17 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/ITellerV2Storage.sol: Interface (1): 6 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IUniswapV2Router.sol: Interface (1): 185 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/IWETH.sol: Interface (1): 33 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/aave/DataTypes.sol: Library (1): 269 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/aave/IFlashLoanSimpleReceiver.sol: Interface (1): 39 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/aave/IPool.sol: Interface (1): 796 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/aave/IPoolAddressesProvider.sol: Interface (1): 250 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/escrow/ICollateralEscrowV1.sol: Interface (1): 46 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/uniswap/IUniswapV3Factory.sol: Interface (1): 75 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/uniswap/IUniswapV3Pool.sol: Interface (1): 24 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/uniswap/pool/IUniswapV3PoolActions.sol: Interface (1): 103 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/uniswap/pool/IUniswapV3PoolDerivedState.sol: Interface (1): 43 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/uniswap/pool/IUniswapV3PoolEvents.sol: Interface (1): 131 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/uniswap/pool/IUniswapV3PoolImmutables.sol: Interface (1): 35 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/uniswap/pool/IUniswapV3PoolOwnerActions.sol: Interface (1): 23 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/interfaces/uniswap/pool/IUniswapV3PoolState.sol: Interface (1): 119 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/libraries/DateTimeLib.sol: Library (1): 502 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/libraries/NumbersLib.sol: Library (1): 135 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/libraries/V2Calculations.sol: Library (1): 216 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/libraries/WadRayMath.sol: Library (1): 121 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/libraries/uniswap/FixedPoint96.sol: Library (1): 10 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/libraries/uniswap/FullMath.sol: Library (1): 125 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/libraries/uniswap/TickMath.sol: Library (1): 250 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/CollateralManagerMock.sol: Smart Contract (1): 100 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/LenderCommitmentForwarderMock.sol: Smart Contract (1): 177 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/LenderManagerMock.sol: Smart Contract (1): 30 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/MarketRegistryMock.sol: Smart Contract (1): 150 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/ReputationManagerMock.sol: Smart Contract (1): 40 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/TellerASMock.sol: Smart Contract (1): 27 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/TellerV2SolMock.sol: Smart Contract (1): 375 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/WethMock.sol: Smart Contract (1): 84 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/aave/AavePoolAddressProviderMock.sol: Smart Contract (1): 175 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/aave/AavePoolMock.sol: Smart Contract (1): 53 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/uniswap/UniswapV3FactoryMock.sol: Smart Contract (1): 14 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/mock/uniswap/UniswapV3PoolMock.sol: Smart Contract (1): 88 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/contracts/type-imports.sol: : 7 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/CollateralEscrow_Override.sol: Smart Contract (1): 58 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/CollateralEscrow_Test.sol: Smart Contract (2): 671 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/CollateralManager_Override.sol: Smart Contract (1): 168 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/CollateralManager_Test.sol: Smart Contract (4): 1549 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/GetMetaDataURI_Test.sol: Smart Contract (1): 32 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/LenderCommitmentForwarder_Integration_Test.sol: Smart Contract (3): 239 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/LenderCommitmentForwarder_OracleLimited_Override.sol: Smart Contract (2): 87 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/LenderCommitmentForwarder_OracleLimited_Unit_Test.sol: Smart Contract (3): 2332 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/LenderCommitmentForwarder_Override.sol: Smart Contract (2): 82 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/LenderCommitmentForwarder_Unit_Test.sol: Smart Contract (3): 1467 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/ExtensionsContextUpgradeable_Test.sol: Smart Contract (3): 45 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/FlashRolloverLoan/FlashRolloverLoan_G2_Integration_Test.sol: Smart Contract (2): 240 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/FlashRolloverLoan/FlashRolloverLoan_G2_Unit_Test.sol: Smart Contract (3): 575 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/FlashRolloverLoan/FlashRolloverLoan_G3_Integration_Test.sol: Smart Contract (2): 263 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/FlashRolloverLoan/FlashRolloverLoan_G3_Unit_Test.sol: Smart Contract (3): 742 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/FlashRolloverLoan/FlashRolloverLoan_G4_Unit_Test.sol: Smart Contract (3): 751 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/FlashRolloverLoan/FlashRolloverLoan_G5_Unit_Test.sol: Smart Contract (3): 761 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/SmartCommitments/LenderCommitmentGroup_Smart_Override.sol: Smart Contract (1): 179 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderCommitmentForwarder/extensions/SmartCommitments/LenderCommitmentGroup_Smart_Test.sol: Smart Contract (3): 953 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderManager_Combined_Test.sol: Smart Contract (3): 166 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderManager_Override.sol: Smart Contract (1): 72 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/LenderManager_Test.sol: Smart Contract (3): 236 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/MarketForwarder_Combined_Test.sol: Smart Contract (3): 162 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/MarketForwarder_Override.sol: Smart Contract (1): 29 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/MarketForwarder_Test.sol: Smart Contract (3): 171 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/MarketLiquidityRewards_Combined_Test.sol: Smart Contract (3): 515 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/MarketLiquidityRewards_Override.sol: Smart Contract (1): 162 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/MarketLiquidityRewards_Test.sol: Smart Contract (3): 669 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/MarketRegistry_Override.sol: Smart Contract (1): 229 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/MarketRegistry_Test.sol: Smart Contract (3): 1144 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/NextDueDate_test.sol: Smart Contract (1): 70 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/PMT_Test.sol: Smart Contract (1): 56 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/ProtocolFee_Override.sol: Smart Contract (1): 17 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/ProtocolFee_Test.sol: Smart Contract (1): 82 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2/TellerV2_Override.sol: Smart Contract (1): 178 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2/TellerV2_Test.sol: Smart Contract (2): 281 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2/TellerV2_bids.sol: Smart Contract (2): 840 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2/TellerV2_getData.sol: Smart Contract (2): 877 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2/TellerV2_initialize.sol: Smart Contract (3): 191 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2/TellerV2_pause.sol: Smart Contract (2): 67 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2Autopay_Test.sol: Smart Contract (2): 299 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2Context/TellerV2Context_Override.sol: Smart Contract (1): 42 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2Context/TellerV2Context__msgSenderForMarket.sol: Smart Contract (1): 61 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2Context/TellerV2Context_approveMarketForwarder.sol: Smart Contract (1): 51 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2Context/TellerV2Context_hasApprovedMarketForwarder.sol: Smart Contract (1): 78 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2Context/TellerV2Context_isTrustedMarketForwarder.sol: Smart Contract (1): 56 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2Context/TellerV2Context_renounceMarketForwarder.sol: Smart Contract (1): 57 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/TellerV2Context/TellerV2Context_setTrustedMarketForwarder.sol: Smart Contract (2): 61 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/Test_Helpers.sol: Smart Contract (1): 149 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/Testable.sol: Abstract Contract (1): 9 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/V2Calculations_Test.sol: Smart Contract (1): 561 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/integration/IntegrationTestHelpers.sol: Library (1): 77 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/resolvers/TestASAttestationResolver.sol: Smart Contract (1): 52 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/resolvers/TestASAttesterResolver.sol: Smart Contract (1): 25 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/resolvers/TestASDataResolver.sol: Smart Contract (1): 20 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/resolvers/TestASExpirationTimeResolver.sol: Smart Contract (1): 25 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/resolvers/TestASPayingResolver.sol: Smart Contract (1): 31 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/resolvers/TestASRecipientResolver.sol: Smart Contract (1): 25 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/resolvers/TestASTokenResolver.sol: Smart Contract (1): 33 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/resolvers/TestASValueResolver.sol: Smart Contract (1): 29 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/tokens/TestERC1155Token.sol: Smart Contract (1): 27 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/tokens/TestERC20Token.sol: Smart Contract (1): 22 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/tokens/TestERC721Token.sol: Smart Contract (1): 18 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/upgrades/LenderCommitmentForwarder_UpgradeWithExtensions.sol: Smart Contract (1): 79 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/util/FoundryTest.sol: Abstract Contract (1): 38 lines
/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts/tests/util/StdChains.sol: Abstract Contract (1): 351 lines

Summary:
Total number of Solidity files: 171
Abstract Contracts Files: 12
Smart Contracts Files: 108
Libraries Files: 9
Interfaces Files: 40
Multiple Types Files: 0
******************************************
Abstract Contracts: 19
Smart Contracts: 156
Libraries: 9
Interfaces: 40
LOC: 37303
