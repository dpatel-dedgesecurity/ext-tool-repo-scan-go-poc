'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts)
'forge config --json' running
'forge build --build-info --skip */tests/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-04-teller-finance/teller-protocol-v2-audit-2024/packages/contracts)

MinimalForwarderUpgradeable.execute(MinimalForwarderUpgradeable.ForwardRequest,bytes) (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#52-77) sends eth to arbitrary user
	Dangerous calls:
	- (success,returndata) = req.to.call{gas: req.gas,value: req.value}(abi.encodePacked(req.data,req.from)) (node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#60-62)
Reference:  

Reentrancy in TellerV2._repayLoan(uint256,TellerV0Storage.Payment,uint256,bool) (contracts/TellerV2.sol#851-898):
	External calls:
	- _sendOrEscrowFunds(_bidId,_payment) (contracts/TellerV2.sol#887)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#110)
		- bid.loanDetails.lendingToken.transferFrom{gas: 100000}(_msgSenderForMarket(bid.marketplaceId),lender,_paymentAmount) (contracts/TellerV2.sol#909-947)
		- (success,returndata) = target.call{value: value}(data) (node_modules/@openzeppelin/contracts/utils/Address.sol#135)
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

CollateralEscrowV1._depositCollateral(CollateralType,address,uint256,uint256).data (contracts/escrow/CollateralEscrowV1.sol#137) is a local variable never initialized
CollateralManager._withdraw(uint256,address).i (contracts/CollateralManager.sol#418) is a local variable never initialized
CollateralManager.deployAndDeposit(uint256).i (contracts/CollateralManager.sol#192) is a local variable never initialized
LenderCommitmentForwarder_G2.getRequiredCollateral(uint256,uint256,ILenderCommitmentForwarder.CommitmentCollateralType,address,address).collateralDecimals (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#595) is a local variable never initialized
ERC20Votes._moveVotingPower(address,address,uint256).newWeight_scope_1 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#233) is a local variable never initialized
ReputationManager.updateAccountReputation(address).i (contracts/ReputationManager.sol#81) is a local variable never initialized
CollateralManager.commitCollateral(uint256,Collateral[]).i (contracts/CollateralManager.sol#129) is a local variable never initialized
LenderCommitmentForwarder_G1.getRequiredCollateral(uint256,uint256,LenderCommitmentForwarder_G1.CommitmentCollateralType,address,address).collateralDecimals (contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#568) is a local variable never initialized
CollateralEscrowV1._withdrawCollateral(Collateral,address,uint256,address).data (contracts/escrow/CollateralEscrowV1.sol#179) is a local variable never initialized
CollateralManager._checkBalances(address,Collateral[],bool).i (contracts/CollateralManager.sol#493) is a local variable never initialized
CollateralManager.getCollateralInfo(uint256).i (contracts/CollateralManager.sol#232) is a local variable never initialized
CollateralManager._deposit(uint256,Collateral).data (contracts/CollateralManager.sol#382) is a local variable never initialized
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

Variable 'ERC20Votes._moveVotingPower(address,address,uint256).oldWeight (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#228)' in ERC20Votes._moveVotingPower(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#221-237) potentially used before declaration: (oldWeight,newWeight) = _writeCheckpoint(_checkpoints[dst],_add,amount) (node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#233)
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
 - [arbitrary-transfer](#arbitrary-transfer) (1 results) (Critical)
 - [reentrancy-with-eth-transfer](#reentrancy-with-eth-transfer) (1 results) (Critical)
 - [reentrancy-without-eth-transfer](#reentrancy-without-eth-transfer) (1 results) (Critical)
 - [sig-replay-attacks-protection](#sig-replay-attacks-protection) (1 results) (Critical)
 - [state-variable-not-initialized](#state-variable-not-initialized) (13 results) (High)
 - [centralized-risk-medium](#centralized-risk-medium) (6 results) (Medium)
 - [constant-result](#constant-result) (1 results) (Medium)
 - [ether-locked](#ether-locked) (1 results) (Medium)
 - [transaction-order-dependency-medium](#transaction-order-dependency-medium) (1 results) (Medium)
 - [uninitialized-local](#uninitialized-local) (12 results) (Medium)
 - [unused-return](#unused-return) (33 results) (Medium)
 - [void-function](#void-function) (37 results) (Medium)
 - [useless-write](#useless-write) (2 results) (Medium)
 - [centralized-risk-low](#centralized-risk-low) (1 results) (Low)
 - [pess-event-setter](#pess-event-setter) (103 results) (Low)
 - [initialize-permission](#initialize-permission) (2 results) (Low)
 - [input-validation](#input-validation) (5 results) (Low)
 - [missing-zero-check](#missing-zero-check) (35 results) (Low)
 - [variable-scope](#variable-scope) (1 results) (Low)
 - [reentrancy-same-effect](#reentrancy-same-effect) (2 results) (Low)
 - [unnecessary-boolean-compare](#unnecessary-boolean-compare) (3 results) (Informational)
 - [centralized-risk-informational](#centralized-risk-informational) (13 results) (Informational)
 - [centralized-risk-other](#centralized-risk-other) (27 results) (Informational)
 - [dead-function](#dead-function) (25 results) (Informational)
 - [error-msg](#error-msg) (25 results) (Informational)
 - [no-license](#no-license) (2 results) (Informational)
 - [price-manipulation-info](#price-manipulation-info) (2 results) (Informational)
 - [uncontrolled-resource-consumption](#uncontrolled-resource-consumption) (10 results) (Informational)
 - [unnecessary-public-function-modifier](#unnecessary-public-function-modifier) (197 results) (Informational)
 - [unused-event](#unused-event) (1 results) (Informational)
 - [version-only](#version-only) (164 results) (Informational)
 - [state-should-be-constant](#state-should-be-constant) (14 results) (Optimization)
## arbitrary-transfer
Impact: Critical
Confidence: Medium
 - [ ] ID-0
[MinimalForwarderUpgradeable.execute(MinimalForwarderUpgradeable.ForwardRequest,bytes)](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L52-L77) sends eth to arbitrary user
	Dangerous calls:
	- [(success,returndata) = req.to.call{gas: req.gas,value: req.value}(abi.encodePacked(req.data,req.from))](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L60-L62)

node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L52-L77


## reentrancy-with-eth-transfer
Impact: Critical
Confidence: Medium
 - [ ] ID-1
Reentrancy in [TellerV2._repayLoan(uint256,TellerV0Storage.Payment,uint256,bool)](contracts/TellerV2.sol#L851-L898):
	External calls:
	- [_sendOrEscrowFunds(_bidId,_payment)](contracts/TellerV2.sol#L887)
		- [returndata = address(token).functionCall(data,SafeERC20: low-level call failed)](node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#L110)
		- [bid.loanDetails.lendingToken.transferFrom{gas: 100000}(_msgSenderForMarket(bid.marketplaceId),lender,_paymentAmount)](contracts/TellerV2.sol#L909-L947)
		- [(success,returndata) = target.call{value: value}(data)](node_modules/@openzeppelin/contracts/utils/Address.sol#L135)
		- [bid.loanDetails.lendingToken.safeTransferFrom(sender,address(this),_paymentAmount)](contracts/TellerV2.sol#L924-L928)
		- [bid.loanDetails.lendingToken.approve(address(escrowVault),paymentAmountReceived)](contracts/TellerV2.sol#L937-L940)
		- [ILoanRepaymentListener(loanRepaymentListener).repayLoanCallback{gas: 80000}(_bidId,_msgSenderForMarket(bid.marketplaceId),_payment.principal,_payment.interest)](contracts/TellerV2.sol#L952-L961)
	External calls sending eth:
	- [_sendOrEscrowFunds(_bidId,_payment)](contracts/TellerV2.sol#L887)
		- [(success,returndata) = target.call{value: value}(data)](node_modules/@openzeppelin/contracts/utils/Address.sol#L135)

contracts/TellerV2.sol#L851-L898


## reentrancy-without-eth-transfer
Impact: Critical
Confidence: Medium
 - [ ] ID-2
Reentrancy in [MarketLiquidityRewards.increaseAllocationAmount(uint256,uint256)](contracts/MarketLiquidityRewards.sol#L154-L163):
	External calls:
	- [IERC20Upgradeable(allocatedRewards[_allocationId].rewardTokenAddress).transferFrom(msg.sender,address(this),_tokenAmount)](contracts/MarketLiquidityRewards.sol#L158-L159)

contracts/MarketLiquidityRewards.sol#L154-L163


## sig-replay-attacks-protection
Impact: Critical
Confidence: High
 - [ ] ID-3
[TellerASEIP712Verifier](contracts/EAS/TellerASEIP712Verifier.sol#L9-L128) [TellerASEIP712Verifier.attest(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32)](contracts/EAS/TellerASEIP712Verifier.sol#L68-L101) missing protection (chainid or address(this)) against signature replay attacks.

contracts/EAS/TellerASEIP712Verifier.sol#L9-L128


## state-variable-not-initialized
Impact: High
Confidence: High
 - [ ] ID-4
state variable: [ERC20Votes._totalSupplyCheckpoints](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L38) not initialized and not written in contract but be used in contract

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L38


 - [ ] ID-5
state variable: [TellerV2Storage_G2.lenderCommitmentForwarder](contracts/TellerV2Storage.sol#L142) not initialized and not written in contract but be used in contract

contracts/TellerV2Storage.sol#L142


 - [ ] ID-6
state variable: [ReputationManager._currentDelinquencies](contracts/ReputationManager.sol#L20) not initialized and not written in contract but be used in contract

contracts/ReputationManager.sol#L20


 - [ ] ID-7
state variable: [TellerV2Storage_G1._approvedForwarderSenders](contracts/TellerV2Storage.sol#L137-L138) not initialized and not written in contract but be used in contract

contracts/TellerV2Storage.sol#L137-L138


 - [ ] ID-8
state variable: [LenderCommitmentForwarderMock.commitmentPrincipalAccepted](contracts/mock/LenderCommitmentForwarderMock.sol#L36) not initialized and not written in contract but be used in contract

contracts/mock/LenderCommitmentForwarderMock.sol#L36


 - [ ] ID-9
state variable: [TellerV2Storage_G0.marketRegistry](contracts/TellerV2Storage.sol#L109) not initialized and not written in contract but be used in contract

contracts/TellerV2Storage.sol#L109


 - [ ] ID-10
state variable: [AccessControlEnumerable._roleMembers](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L16) not initialized and not written in contract but be used in contract

node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L16


 - [ ] ID-11
state variable: [ERC20Votes._checkpoints](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L37) not initialized and not written in contract but be used in contract

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L37


 - [ ] ID-12
state variable: [ERC20Permit._nonces](node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L25) not initialized and not written in contract but be used in contract

node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L25


 - [ ] ID-13
state variable: [ReputationManager._currentDefaults](contracts/ReputationManager.sol#L21) not initialized and not written in contract but be used in contract

contracts/ReputationManager.sol#L21


 - [ ] ID-14
state variable: [TellerV2Storage_G0._borrowerBidsActive](contracts/TellerV2Storage.sol#L113) not initialized and not written in contract but be used in contract

contracts/TellerV2Storage.sol#L113


 - [ ] ID-15
state variable: [ReputationManager._defaults](contracts/ReputationManager.sol#L19) not initialized and not written in contract but be used in contract

contracts/ReputationManager.sol#L19


 - [ ] ID-16
state variable: [ReputationManager._delinquencies](contracts/ReputationManager.sol#L18) not initialized and not written in contract but be used in contract

contracts/ReputationManager.sol#L18


## centralized-risk-medium
Impact: Medium
Confidence: High
 - [ ] ID-17
	- [FlashRolloverLoan_G5.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L156-L219)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L156-L219


 - [ ] ID-18
	- [FlashRolloverLoan_G1.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L137-L199)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L137-L199


 - [ ] ID-19
	- [FlashRolloverLoan_G4.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L155-L218)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L155-L218


 - [ ] ID-20
	- [FlashRolloverLoan_G3.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L157-L219)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L157-L219


 - [ ] ID-21
	- [WethMock.withdraw(uint256)](contracts/mock/WethMock.sol#L42-L47)

contracts/mock/WethMock.sol#L42-L47


 - [ ] ID-22
	- [MarketLiquidityRewards.deallocateRewards(uint256,uint256)](contracts/MarketLiquidityRewards.sol#L170-L198)

contracts/MarketLiquidityRewards.sol#L170-L198


## constant-result
Impact: Medium
Confidence: High
 - [ ] ID-23
[MarketRegistry.setMarketFeePercent(uint256,uint16)](contracts/MarketRegistry.sol#L636-L645) contains a tautology or contradiction:
	- [require(bool,string)(_newPercent >= 0 && _newPercent <= 10000,invalid percent)](contracts/MarketRegistry.sol#L640)

contracts/MarketRegistry.sol#L636-L645


## ether-locked
Impact: Medium
Confidence: High
 - [ ] ID-24
Contract locking ether found:
	Contract [MarketRegistry](contracts/MarketRegistry.sol#L19-L1260) has payable functions:
	 - [IASResolver.resolve(address,bytes,bytes,uint256,address)](contracts/interfaces/IASResolver.sol#L25-L31)
	 - [MarketRegistry.resolve(address,bytes,bytes,uint256,address)](contracts/MarketRegistry.sol#L452-L471)
	But does not have a function to withdraw the ether

contracts/MarketRegistry.sol#L19-L1260


## transaction-order-dependency-medium
Impact: Medium
Confidence: High
 - [ ] ID-25
setCommitment(uint256,ILenderCommitmentForwarder.Commitment) and acceptCommitmentWithRecipient(uint256,uint256,uint256,uint256,address,address,uint16,uint32) have transaction order dependency caused by data race on state variables:LenderCommitmentForwarderMock.commitments
	- [LenderCommitmentForwarderMock.setCommitment(uint256,ILenderCommitmentForwarder.Commitment)](contracts/mock/LenderCommitmentForwarderMock.sol#L45-L49)
	- [LenderCommitmentForwarderMock.acceptCommitmentWithRecipient(uint256,uint256,uint256,uint256,address,address,uint16,uint32)](contracts/mock/LenderCommitmentForwarderMock.sol#L94-L115)

contracts/mock/LenderCommitmentForwarderMock.sol#L45-L49


## uninitialized-local
Impact: Medium
Confidence: Medium
 - [ ] ID-26
[CollateralManager.deployAndDeposit(uint256).i](contracts/CollateralManager.sol#L192) is a local variable never initialized

contracts/CollateralManager.sol#L192


 - [ ] ID-27
[CollateralEscrowV1._withdrawCollateral(Collateral,address,uint256,address).data](contracts/escrow/CollateralEscrowV1.sol#L179) is a local variable never initialized

contracts/escrow/CollateralEscrowV1.sol#L179


 - [ ] ID-28
[CollateralManager._withdraw(uint256,address).i](contracts/CollateralManager.sol#L418) is a local variable never initialized

contracts/CollateralManager.sol#L418


 - [ ] ID-29
[LenderCommitmentForwarder_G2.getRequiredCollateral(uint256,uint256,ILenderCommitmentForwarder.CommitmentCollateralType,address,address).collateralDecimals](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L595) is a local variable never initialized

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L595


 - [ ] ID-30
[CollateralManager.getCollateralInfo(uint256).i](contracts/CollateralManager.sol#L232) is a local variable never initialized

contracts/CollateralManager.sol#L232


 - [ ] ID-31
[CollateralManager._checkBalances(address,Collateral[],bool).i](contracts/CollateralManager.sol#L493) is a local variable never initialized

contracts/CollateralManager.sol#L493


 - [ ] ID-32
[ReputationManager.updateAccountReputation(address).i](contracts/ReputationManager.sol#L81) is a local variable never initialized

contracts/ReputationManager.sol#L81


 - [ ] ID-33
[CollateralEscrowV1._depositCollateral(CollateralType,address,uint256,uint256).data](contracts/escrow/CollateralEscrowV1.sol#L137) is a local variable never initialized

contracts/escrow/CollateralEscrowV1.sol#L137


 - [ ] ID-34
[LenderCommitmentForwarder_G1.getRequiredCollateral(uint256,uint256,LenderCommitmentForwarder_G1.CommitmentCollateralType,address,address).collateralDecimals](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L568) is a local variable never initialized

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L568


 - [ ] ID-35
[ERC20Votes._moveVotingPower(address,address,uint256).newWeight_scope_1](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L233) is a local variable never initialized

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L233


 - [ ] ID-36
[CollateralManager._deposit(uint256,Collateral).data](contracts/CollateralManager.sol#L382) is a local variable never initialized

contracts/CollateralManager.sol#L382


 - [ ] ID-37
[CollateralManager.commitCollateral(uint256,Collateral[]).i](contracts/CollateralManager.sol#L129) is a local variable never initialized

contracts/CollateralManager.sol#L129


## unused-return
Impact: Medium
Confidence: Medium
 - [ ] ID-38
[CollateralManager._deposit(uint256,Collateral)](contracts/CollateralManager.sol#L339-L409)have ignores return value in [IERC20Upgradeable(collateralInfo._collateralAddress).approve(escrowAddress,collateralInfo._amount)](contracts/CollateralManager.sol#L355-L358)

contracts/CollateralManager.sol#L339-L409


 - [ ] ID-39
[TellerV2.lenderAcceptBid(uint256)](contracts/TellerV2.sol#L481-L576)have ignores return value in [_borrowerBidsActive[bid.borrower].add(_bidId)](contracts/TellerV2.sol#L569)

contracts/TellerV2.sol#L481-L576


 - [ ] ID-40
[FlashRolloverLoan_G5.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L156-L219)have ignores return value in [IERC20Upgradeable(_flashToken).approve(address(POOL()),_flashAmount + _flashFees)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L194-L197)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L156-L219


 - [ ] ID-41
[FlashRolloverLoan_G1.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L137-L199)have ignores return value in [IERC20Upgradeable(_flashToken).approve(address(POOL()),_flashAmount + _flashFees)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L174-L177)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L137-L199


 - [ ] ID-42
[ReputationManager._removeMark(address,uint256,RepMark)](contracts/ReputationManager.sol#L129-L139)have ignores return value in [_currentDelinquencies[_account].remove(_bidId)](contracts/ReputationManager.sol#L133)

contracts/ReputationManager.sol#L129-L139


 - [ ] ID-43
[ReputationManager._addMark(address,uint256,RepMark)](contracts/ReputationManager.sol#L115-L127)have ignores return value in [_delinquencies[_account].add(_bidId)](contracts/ReputationManager.sol#L119)

contracts/ReputationManager.sol#L115-L127


 - [ ] ID-44
[CollateralManager._commitCollateral(uint256,Collateral)](contracts/CollateralManager.sol#L449-L478)have ignores return value in [collateral.collateralAddresses.add(_collateralInfo._collateralAddress)](contracts/CollateralManager.sol#L467)

contracts/CollateralManager.sol#L449-L478


 - [ ] ID-45
[TellerV2._sendOrEscrowFunds(uint256,TellerV0Storage.Payment)](contracts/TellerV2.sol#L901-L963)have ignores return value in [bid.loanDetails.lendingToken.approve(address(escrowVault),paymentAmountReceived)](contracts/TellerV2.sol#L937-L940)

contracts/TellerV2.sol#L901-L963


 - [ ] ID-46
[AccessControlEnumerable._grantRole(bytes32,address)](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L52-L55)have ignores return value in [_roleMembers[role].add(account)](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L54)

node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L52-L55


 - [ ] ID-47
[LenderCommitmentForwarder_U1._addBorrowersToCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L291-L299)have ignores return value in [commitmentBorrowersList[_commitmentId].add(_borrowerArray[i])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L296)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L291-L299


 - [ ] ID-48
[FlashRolloverLoan_G4._repayLoanFull(uint256,address,uint256)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L234-L252)have ignores return value in [IERC20Upgradeable(_principalToken).approve(address(TELLER_V2),_repayAmount)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L242-L245)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L234-L252


 - [ ] ID-49
[MarketRegistry._attestStakeholderVerification(uint256,address,bytes32,bool)](contracts/MarketRegistry.sol#L1118-L1147)have ignores return value in [markets[_marketId].verifiedLendersForMarket.add(_stakeholderAddress)](contracts/MarketRegistry.sol#L1130-L1132)

contracts/MarketRegistry.sol#L1118-L1147


 - [ ] ID-50
[FlashRolloverLoan_G5._repayLoanFull(uint256,address,uint256)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L235-L253)have ignores return value in [IERC20Upgradeable(_principalToken).approve(address(TELLER_V2),_repayAmount)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L243-L246)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L235-L253


 - [ ] ID-51
[MarketRegistry._revokeStakeholderVerification(uint256,address,bool)](contracts/MarketRegistry.sol#L1209-L1235)have ignores return value in [markets[_marketId].verifiedLendersForMarket.remove(_stakeholderAddress)](contracts/MarketRegistry.sol#L1219-L1221)

contracts/MarketRegistry.sol#L1209-L1235


 - [ ] ID-52
[FlashRolloverLoan_G4.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L155-L218)have ignores return value in [IERC20Upgradeable(_flashToken).approve(address(POOL()),_flashAmount + _flashFees)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L193-L196)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L155-L218


 - [ ] ID-53
[TellerV2Autopay.autoPayLoanMinimum(uint256)](contracts/TellerV2Autopay.sol#L109-L143)have ignores return value in [ERC20(lendingToken).approve(address(tellerV2),amountToRepayMinimum)](contracts/TellerV2Autopay.sol#L137)

contracts/TellerV2Autopay.sol#L109-L143


 - [ ] ID-54
[LenderCommitmentForwarder_U1._removeBorrowersFromCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L306-L314)have ignores return value in [commitmentBorrowersList[_commitmentId].remove(_borrowerArray[i])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L311)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L306-L314


 - [ ] ID-55
[ERC1967Upgrade._upgradeToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L65-L74)have ignores return value in [Address.functionDelegateCall(newImplementation,data)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L72)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L65-L74


 - [ ] ID-56
[AccessControlEnumerable._revokeRole(bytes32,address)](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L60-L63)have ignores return value in [_roleMembers[role].remove(account)](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L62)

node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L60-L63


 - [ ] ID-57
[FlashRolloverLoan_G1._repayLoanFull(uint256,address,uint256)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L201-L219)have ignores return value in [IERC20Upgradeable(_principalToken).approve(address(TELLER_V2),_repayAmount)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L209-L212)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L201-L219


 - [ ] ID-58
[TellerV2Context.renounceMarketForwarder(uint256,address)](contracts/TellerV2Context.sol#L108-L115)have ignores return value in [_approvedForwarderSenders[_forwarder].remove(_msgSender())](contracts/TellerV2Context.sol#L112)

contracts/TellerV2Context.sol#L108-L115


 - [ ] ID-59
[ERC1967Upgrade._upgradeBeaconToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L174-L184)have ignores return value in [Address.functionDelegateCall(IBeacon(newBeacon).implementation(),data)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L182)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L174-L184


 - [ ] ID-60
[LenderCommitmentGroup_Smart._acceptBidWithRepaymentListener(uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L384-L391)have ignores return value in [ITellerV2(TELLER_V2).lenderAcceptBid(_bidId)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L385)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L384-L391


 - [ ] ID-61
[FlashRolloverLoan_G3._repayLoanFull(uint256,address,uint256)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L235-L253)have ignores return value in [IERC20Upgradeable(_principalToken).approve(address(TELLER_V2),_repayAmount)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L243-L246)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L235-L253


 - [ ] ID-62
[TellerV2Context.approveMarketForwarder(uint256,address)](contracts/TellerV2Context.sol#L92-L101)have ignores return value in [_approvedForwarderSenders[_forwarder].add(_msgSender())](contracts/TellerV2Context.sol#L99)

contracts/TellerV2Context.sol#L92-L101


 - [ ] ID-63
[LenderCommitmentForwarder_G2._removeBorrowersFromCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L263-L271)have ignores return value in [commitmentBorrowersList[_commitmentId].remove(_borrowerArray[i])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L268)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L263-L271


 - [ ] ID-64
[CommitmentRolloverLoan.rolloverLoan(uint256,uint256,ICommitmentRolloverLoan.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L41-L77)have ignores return value in [lendingToken.approve(address(TELLER_V2),fundsReceived)](contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L68)

contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L41-L77


 - [ ] ID-65
[TellerV2._repayLoan(uint256,TellerV0Storage.Payment,uint256,bool)](contracts/TellerV2.sol#L851-L898)have ignores return value in [_borrowerBidsActive[bid.borrower].remove(_bidId)](contracts/TellerV2.sol#L874)

contracts/TellerV2.sol#L851-L898


 - [ ] ID-66
[FlashRolloverLoan_G3.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L157-L219)have ignores return value in [IERC20Upgradeable(_flashToken).approve(address(POOL()),_flashAmount + _flashFees)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L194-L197)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L157-L219


 - [ ] ID-67
[LenderCommitmentForwarder_G2._addBorrowersToCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L248-L256)have ignores return value in [commitmentBorrowersList[_commitmentId].add(_borrowerArray[i])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L253)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L248-L256


 - [ ] ID-68
[LenderCommitmentForwarder_G1._removeBorrowersFromCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L296-L304)have ignores return value in [commitmentBorrowersList[_commitmentId].remove(_borrowerArray[i])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L301)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L296-L304


 - [ ] ID-69
[LenderCommitmentGroup_Smart.acceptFundsForAcceptBid(address,uint256,uint256,uint256,address,uint256,uint32,uint16)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L336-L382)have ignores return value in [principalToken.approve(address(TELLER_V2),_principalAmount)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L373)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L336-L382


 - [ ] ID-70
[LenderCommitmentForwarder_G1._addBorrowersToCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L281-L289)have ignores return value in [commitmentBorrowersList[_commitmentId].add(_borrowerArray[i])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L286)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L281-L289


## void-function
Impact: Medium
Confidence: High
 - [ ] ID-71
function:[MarketRegistryMock.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,PaymentType,PaymentCycleType,string)](contracts/mock/MarketRegistryMock.sol#L113-L124)is empty 

contracts/mock/MarketRegistryMock.sol#L113-L124


 - [ ] ID-72
function:[ERC20._afterTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L384-L388)is empty 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L384-L388


 - [ ] ID-73
function:[ContextUpgradeable.__Context_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L21-L22)is empty 

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L21-L22


 - [ ] ID-74
function:[TellerV2SolMock.isLoanDefaulted(uint256)](contracts/mock/TellerV2SolMock.sol#L247-L252)is empty 

contracts/mock/TellerV2SolMock.sol#L247-L252


 - [ ] ID-75
function:[ERC721Upgradeable._afterTokenTransfer(address,address,uint256,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L502-L507)is empty 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L502-L507


 - [ ] ID-76
function:[CollateralManagerMock.withdraw(uint256)](contracts/mock/CollateralManagerMock.sol#L74)is empty 

contracts/mock/CollateralManagerMock.sol#L74


 - [ ] ID-77
function:[TellerV2SolMock.liquidateLoanFullWithRecipient(uint256,address)](contracts/mock/TellerV2SolMock.sol#L123-L125)is empty 

contracts/mock/TellerV2SolMock.sol#L123-L125


 - [ ] ID-78
function:[MarketRegistryMock.getPaymentType(uint256)](contracts/mock/MarketRegistryMock.sol#L107-L111)is empty 

contracts/mock/MarketRegistryMock.sol#L107-L111


 - [ ] ID-79
function:[TellerV2SolMock.repayLoanMinimum(uint256)](contracts/mock/TellerV2SolMock.sol#L127)is empty 

contracts/mock/TellerV2SolMock.sol#L127


 - [ ] ID-80
function:[CollateralManagerMock.lenderClaimCollateral(uint256)](contracts/mock/CollateralManagerMock.sol#L85)is empty 

contracts/mock/CollateralManagerMock.sol#L85


 - [ ] ID-81
function:[TellerV2SolMock.isLoanLiquidateable(uint256)](contracts/mock/TellerV2SolMock.sol#L254-L259)is empty 

contracts/mock/TellerV2SolMock.sol#L254-L259


 - [ ] ID-82
function:[MarketRegistryMock.closeMarket(uint256)](contracts/mock/MarketRegistryMock.sol#L137)is empty 

contracts/mock/MarketRegistryMock.sol#L137


 - [ ] ID-83
function:[ReputationManagerMock.getDefaultedLoanIds(address)](contracts/mock/ReputationManagerMock.sol#L17-L20)is empty 

contracts/mock/ReputationManagerMock.sol#L17-L20


 - [ ] ID-84
function:[ERC165Upgradeable.__ERC165_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#L27-L28)is empty 

node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#L27-L28


 - [ ] ID-85
function:[MarketRegistryMock.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,string)](contracts/mock/MarketRegistryMock.sol#L126-L135)is empty 

contracts/mock/MarketRegistryMock.sol#L126-L135


 - [ ] ID-86
function:[MinimalForwarderUpgradeable.__MinimalForwarder_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L39)is empty 

node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L39


 - [ ] ID-87
function:[LenderCommitmentForwarderMock.createCommitment(ILenderCommitmentForwarder.Commitment,address[])](contracts/mock/LenderCommitmentForwarderMock.sol#L40-L43)is empty 

contracts/mock/LenderCommitmentForwarderMock.sol#L40-L43


 - [ ] ID-88
function:[TellerV2SolMock.liquidateLoanFull(uint256)](contracts/mock/TellerV2SolMock.sol#L121)is empty 

contracts/mock/TellerV2SolMock.sol#L121


 - [ ] ID-89
function:[TellerV2SolMock.isPaymentLate(uint256)](contracts/mock/TellerV2SolMock.sol#L261)is empty 

contracts/mock/TellerV2SolMock.sol#L261


 - [ ] ID-90
function:[TellerV2SolMock.getRepaymentListenerForBid(uint256)](contracts/mock/TellerV2SolMock.sol#L337-L341)is empty 

contracts/mock/TellerV2SolMock.sol#L337-L341


 - [ ] ID-91
function:[ReputationManagerMock.initialize(address)](contracts/mock/ReputationManagerMock.sol#L10)is empty 

contracts/mock/ReputationManagerMock.sol#L10


 - [ ] ID-92
function:[ReputationManagerMock.getDelinquentLoanIds(address)](contracts/mock/ReputationManagerMock.sol#L12-L15)is empty 

contracts/mock/ReputationManagerMock.sol#L12-L15


 - [ ] ID-93
function:[ReputationManagerMock.updateAccountReputation(address)](contracts/mock/ReputationManagerMock.sol#L32)is empty 

contracts/mock/ReputationManagerMock.sol#L32


 - [ ] ID-94
function:[AavePoolAddressProviderMock.setAddressAsProxy(bytes32,address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L166-L168)is empty 

contracts/mock/aave/AavePoolAddressProviderMock.sol#L166-L168


 - [ ] ID-95
function:[EscrowVault.initialize()](contracts/EscrowVault.sol#L25)is empty 

contracts/EscrowVault.sol#L25


 - [ ] ID-96
function:[ContextUpgradeable.__Context_init()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L18-L19)is empty 

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L18-L19


 - [ ] ID-97
function:[AavePoolAddressProviderMock.setPoolConfiguratorImpl(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L170-L172)is empty 

contracts/mock/aave/AavePoolAddressProviderMock.sol#L170-L172


 - [ ] ID-98
function:[TellerV2SolMock.getBorrowerActiveLoanIds(address)](contracts/mock/TellerV2SolMock.sol#L241-L245)is empty 

contracts/mock/TellerV2SolMock.sol#L241-L245


 - [ ] ID-99
function:[ERC165Upgradeable.__ERC165_init()](node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#L24-L25)is empty 

node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#L24-L25


 - [ ] ID-100
function:[CollateralManagerMock.liquidateCollateral(uint256,address)](contracts/mock/CollateralManagerMock.sol#L93-L95)is empty 

contracts/mock/CollateralManagerMock.sol#L93-L95


 - [ ] ID-101
function:[MarketRegistryMock.initialize(TellerAS)](contracts/mock/MarketRegistryMock.sol#L20)is empty 

contracts/mock/MarketRegistryMock.sol#L20


 - [ ] ID-102
function:[ERC20._beforeTokenTransfer(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L364-L368)is empty 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L364-L368


 - [ ] ID-103
function:[MarketLiquidityRewards.initialize()](contracts/MarketLiquidityRewards.sol#L79)is empty 

contracts/MarketLiquidityRewards.sol#L79


 - [ ] ID-104
function:[AavePoolAddressProviderMock.setPoolImpl(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L174)is empty 

contracts/mock/aave/AavePoolAddressProviderMock.sol#L174


 - [ ] ID-105
function:[ReputationManagerMock.getCurrentDelinquentLoanIds(address)](contracts/mock/ReputationManagerMock.sol#L22-L25)is empty 

contracts/mock/ReputationManagerMock.sol#L22-L25


 - [ ] ID-106
function:[ReputationManagerMock.getCurrentDefaultLoanIds(address)](contracts/mock/ReputationManagerMock.sol#L27-L30)is empty 

contracts/mock/ReputationManagerMock.sol#L27-L30


 - [ ] ID-107
function:[TellerV2SolMock.setRepaymentListenerForBid(uint256,address)](contracts/mock/TellerV2SolMock.sol#L343-L345)is empty 

contracts/mock/TellerV2SolMock.sol#L343-L345


## useless-write
Impact: Medium
Confidence: High
 - [ ] ID-108
[BokkyPooBahsDateTimeLibrary.subMonths(uint256,uint256).year](contracts/libraries/DateTimeLib.sol#L388) is written in both
	[(year,month,day) = _daysToDate(timestamp / SECONDS_PER_DAY)](contracts/libraries/DateTimeLib.sol#L388-L390)
	[year = yearMonth / 12](contracts/libraries/DateTimeLib.sol#L392)

contracts/libraries/DateTimeLib.sol#L388


 - [ ] ID-109
[BokkyPooBahsDateTimeLibrary._daysToDate(uint256).L](contracts/libraries/DateTimeLib.sol#L103) is written in both
	[L = L - (1461 * _year) / 4 + 31](contracts/libraries/DateTimeLib.sol#L107)
	[L = _month / 11](contracts/libraries/DateTimeLib.sol#L110)

contracts/libraries/DateTimeLib.sol#L103


## centralized-risk-low
Impact: Low
Confidence: High
 - [ ] ID-110
	- [LenderCommitmentGroup_Smart.repayLoanCallback(uint256,address,uint256,uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L700-L709)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L700-L709


## pess-event-setter
Impact: Low
Confidence: Medium
 - [ ] ID-111
Setter function [MarketRegistry.revokeBorrower(uint256,address)](contracts/MarketRegistry.sol#L401-L405) does not emit an event

contracts/MarketRegistry.sol#L401-L405


 - [ ] ID-112
Setter function [TellerV2._setEscrowVault(address)](contracts/TellerV2.sol#L245-L248) does not emit an event

contracts/TellerV2.sol#L245-L248


 - [ ] ID-113
Setter function [LenderCommitmentGroup_Smart.onlyTellerV2()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L128-L134) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L128-L134


 - [ ] ID-114
Setter function [LenderCommitmentGroup_Smart._deployPoolSharesToken()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L215-L239) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L215-L239


 - [ ] ID-115
Setter function [CommitmentRolloverLoan.rolloverLoan(uint256,uint256,ICommitmentRolloverLoan.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L41-L77) does not emit an event

contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L41-L77


 - [ ] ID-116
Setter function [TellerV2SolMock.submitBid(address,uint256,uint256,uint32,uint16,string,address)](contracts/mock/TellerV2SolMock.sol#L50-L80) does not emit an event

contracts/mock/TellerV2SolMock.sol#L50-L80


 - [ ] ID-117
Setter function [MarketRegistryMock.setMarketFeeRecipient(address)](contracts/mock/MarketRegistryMock.sol#L103-L105) does not emit an event

contracts/mock/MarketRegistryMock.sol#L103-L105


 - [ ] ID-118
Setter function [LenderCommitmentGroup_Smart.onlySmartCommitmentForwarder()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L120-L126) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L120-L126


 - [ ] ID-119
Setter function [LenderCommitmentGroup_Smart.unpauseBorrowing()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L800-L802) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L800-L802


 - [ ] ID-120
Setter function [MarketRegistry.updateMarketSettings(uint256,uint32,PaymentType,PaymentCycleType,uint32,uint32,uint16,bool,bool,string)](contracts/MarketRegistry.sol#L502-L526) does not emit an event

contracts/MarketRegistry.sol#L502-L526


 - [ ] ID-121
Setter function [MarketRegistryMock.mock_setGlobalMarketsClosed(bool)](contracts/mock/MarketRegistryMock.sol#L139-L141) does not emit an event

contracts/mock/MarketRegistryMock.sol#L139-L141


 - [ ] ID-122
Setter function [FlashRolloverLoan_G4.onlyFlashLoanPool()](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L70-L77) does not emit an event

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L70-L77


 - [ ] ID-123
Setter function [LenderCommitmentGroup_Smart.pauseBorrowing()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L793-L795) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L793-L795


 - [ ] ID-124
Setter function [ERC2771ContextUpgradeable._msgData()](node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#L37-L43) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#L37-L43


 - [ ] ID-125
Setter function [LenderManager.__LenderManager_init()](contracts/LenderManager.sol#L30-L33) does not emit an event

contracts/LenderManager.sol#L30-L33


 - [ ] ID-126
Setter function [MarketRegistry.attestLender(uint256,address,uint256)](contracts/MarketRegistry.sol#L288-L294) does not emit an event

contracts/MarketRegistry.sol#L288-L294


 - [ ] ID-127
Setter function [UpgradeableBeacon._setImplementation(address)](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L61-L64) does not emit an event

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L61-L64


 - [ ] ID-128
Setter function [LenderCommitmentGroupShares.mint(address,uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L17-L19) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L17-L19


 - [ ] ID-129
Setter function [AavePoolAddressProviderMock.setPoolImpl(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L174) does not emit an event

contracts/mock/aave/AavePoolAddressProviderMock.sol#L174


 - [ ] ID-130
Setter function [LenderCommitmentGroup_Smart.acceptFundsForAcceptBid(address,uint256,uint256,uint256,address,uint256,uint32,uint16)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L336-L382) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L336-L382


 - [ ] ID-131
Setter function [UniswapV3PoolMock.set_mockToken0(address)](contracts/mock/uniswap/UniswapV3PoolMock.sol#L34-L36) does not emit an event

contracts/mock/uniswap/UniswapV3PoolMock.sol#L34-L36


 - [ ] ID-132
Setter function [ERC20PresetMinterPauser.constructor(string,string)](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L38-L43) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L38-L43


 - [ ] ID-133
Setter function [MarketRegistry.initialize(TellerAS)](contracts/MarketRegistry.sol#L106-L117) does not emit an event

contracts/MarketRegistry.sol#L106-L117


 - [ ] ID-134
Setter function [MarketRegistry._attestStakeholder(uint256,address,uint256,bool)](contracts/MarketRegistry.sol#L1026-L1057) does not emit an event

contracts/MarketRegistry.sol#L1026-L1057


 - [ ] ID-135
Setter function [CommitmentRolloverLoan._acceptCommitment(ICommitmentRolloverLoan.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L126-L151) does not emit an event

contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L126-L151


 - [ ] ID-136
Setter function [FlashRolloverLoan_G5.rolloverLoanWithFlash(address,uint256,uint256,uint256,FlashRolloverLoan_G5.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L97-L135) does not emit an event

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L97-L135


 - [ ] ID-137
Setter function [AavePoolMock.flashLoanSimple(address,address,uint256,bytes,uint16)](contracts/mock/aave/AavePoolMock.sol#L16-L52) does not emit an event

contracts/mock/aave/AavePoolMock.sol#L16-L52


 - [ ] ID-138
Setter function [LenderCommitmentGroup_Smart.repayLoanCallback(uint256,address,uint256,uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L700-L709) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L700-L709


 - [ ] ID-139
Setter function [AavePoolMock.setShouldExecuteCallback(bool)](contracts/mock/aave/AavePoolMock.sol#L12-L14) does not emit an event

contracts/mock/aave/AavePoolMock.sol#L12-L14


 - [ ] ID-140
Setter function [MarketRegistryMock.mock_setBorrowerIsVerified(bool)](contracts/mock/MarketRegistryMock.sol#L143-L145) does not emit an event

contracts/mock/MarketRegistryMock.sol#L143-L145


 - [ ] ID-141
Setter function [ERC20PresetMinterPauser.unpause()](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L82-L85) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L82-L85


 - [ ] ID-142
Setter function [TellerV2SolMock.mock_setLoanDefaultTimestamp(uint256)](contracts/mock/TellerV2SolMock.sol#L329-L333) does not emit an event

contracts/mock/TellerV2SolMock.sol#L329-L333


 - [ ] ID-143
Setter function [UniswapV3PoolMock.set_mockToken1(address)](contracts/mock/uniswap/UniswapV3PoolMock.sol#L39-L41) does not emit an event

contracts/mock/uniswap/UniswapV3PoolMock.sol#L39-L41


 - [ ] ID-144
Setter function [ERC2771ContextUpgradeable._msgSender()](node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#L25-L35) does not emit an event

node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#L25-L35


 - [ ] ID-145
Setter function [FlashRolloverLoan_G4.rolloverLoanWithFlash(address,uint256,uint256,uint256,FlashRolloverLoan_G4.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L96-L134) does not emit an event

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L96-L134


 - [ ] ID-146
Setter function [MarketRegistry.withAttestingSchema(bytes32)](contracts/MarketRegistry.sol#L72-L76) does not emit an event

contracts/MarketRegistry.sol#L72-L76


 - [ ] ID-147
Setter function [CollateralEscrowV1._depositCollateral(CollateralType,address,uint256,uint256)](contracts/escrow/CollateralEscrowV1.sol#L111-L149) does not emit an event

contracts/escrow/CollateralEscrowV1.sol#L111-L149


 - [ ] ID-148
Setter function [CollateralManagerMock.forceSetCommitCollateralValidation(bool)](contracts/mock/CollateralManagerMock.sol#L97-L99) does not emit an event

contracts/mock/CollateralManagerMock.sol#L97-L99


 - [ ] ID-149
Setter function [MarketRegistryMock.setMarketOwner(address)](contracts/mock/MarketRegistryMock.sol#L99-L101) does not emit an event

contracts/mock/MarketRegistryMock.sol#L99-L101


 - [ ] ID-150
Setter function [FlashRolloverLoan_G5.onlyFlashLoanPool()](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L72-L79) does not emit an event

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L72-L79


 - [ ] ID-151
Setter function [ERC20PresetMinterPauser.pause()](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L68-L71) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L68-L71


 - [ ] ID-152
Setter function [AavePoolAddressProviderMock.setMarketId(string)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L46-L52) does not emit an event

contracts/mock/aave/AavePoolAddressProviderMock.sol#L46-L52


 - [ ] ID-153
Setter function [LenderCommitmentGroup_Smart.burnSharesToWithdrawEarnings(uint256,address)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L396-L415) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L396-L415


 - [ ] ID-154
Setter function [WethMock.transfer(address,uint256)](contracts/mock/WethMock.sol#L59-L61) does not emit an event

contracts/mock/WethMock.sol#L59-L61


 - [ ] ID-155
Setter function [MarketRegistry.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,string)](contracts/MarketRegistry.sol#L170-L192) does not emit an event

contracts/MarketRegistry.sol#L170-L192


 - [ ] ID-156
Setter function [TellerV2.initialize(uint16,address,address,address,address,address,address)](contracts/TellerV2.sol#L188-L227) does not emit an event

contracts/TellerV2.sol#L188-L227


 - [ ] ID-157
Setter function [LenderCommitmentGroup_Smart.initialize(address,address,uint256,uint32,uint16,uint16,uint16,uint16,uint24,uint32)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L158-L213) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L158-L213


 - [ ] ID-158
Setter function [LenderManagerMock.registerLoan(uint256,address)](contracts/mock/LenderManagerMock.sol#L15-L20) does not emit an event

contracts/mock/LenderManagerMock.sol#L15-L20


 - [ ] ID-159
Setter function [TellerV2.setRepaymentListenerForBid(uint256,address)](contracts/TellerV2.sol#L1248-L1259) does not emit an event

contracts/TellerV2.sol#L1248-L1259


 - [ ] ID-160
Setter function [LenderCommitmentGroup_Smart.addPrincipalToCommitmentGroup(uint256,address)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L307-L322) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L307-L322


 - [ ] ID-161
Setter function [CollateralManagerMock.deployAndDeposit(uint256)](contracts/mock/CollateralManagerMock.sol#L37-L39) does not emit an event

contracts/mock/CollateralManagerMock.sol#L37-L39


 - [ ] ID-162
Setter function [CollateralManager.setCollateralEscrowBeacon(address)](contracts/CollateralManager.sol#L93-L98) does not emit an event

contracts/CollateralManager.sol#L93-L98


 - [ ] ID-163
Setter function [CollateralManager.commitCollateral(uint256,Collateral)](contracts/CollateralManager.sol#L142-L152) does not emit an event

contracts/CollateralManager.sol#L142-L152


 - [ ] ID-164
Setter function [TellerV2._setLenderManager(address)](contracts/TellerV2.sol#L234-L243) does not emit an event

contracts/TellerV2.sol#L234-L243


 - [ ] ID-165
Setter function [TellerV2SolMock.lenderAcceptBid(uint256)](contracts/mock/TellerV2SolMock.sol#L199-L222) does not emit an event

contracts/mock/TellerV2SolMock.sol#L199-L222


 - [ ] ID-166
Setter function [MarketRegistry.attestBorrower(uint256,address,uint256)](contracts/MarketRegistry.sol#L366-L372) does not emit an event

contracts/MarketRegistry.sol#L366-L372


 - [ ] ID-167
Setter function [MarketRegistry.revokeLender(uint256,address)](contracts/MarketRegistry.sol#L323-L325) does not emit an event

contracts/MarketRegistry.sol#L323-L325


 - [ ] ID-168
Setter function [TellerV2SolMock.setRepaymentListenerForBid(uint256,address)](contracts/mock/TellerV2SolMock.sol#L343-L345) does not emit an event

contracts/mock/TellerV2SolMock.sol#L343-L345


 - [ ] ID-169
Setter function [TLR.burn(address,uint256)](contracts/TLR.sol#L54-L56) does not emit an event

contracts/TLR.sol#L54-L56


 - [ ] ID-170
Setter function [ProtocolFeeMock.initialize(uint16)](contracts/ProtocolFeeMock.sol#L9-L11) does not emit an event

contracts/ProtocolFeeMock.sol#L9-L11


 - [ ] ID-171
Setter function [TellerV2SolMock.approveMarketForwarder(uint256,address)](contracts/mock/TellerV2SolMock.sol#L43-L47) does not emit an event

contracts/mock/TellerV2SolMock.sol#L43-L47


 - [ ] ID-172
Setter function [MarketRegistryMock.mock_setLenderIsVerified(bool)](contracts/mock/MarketRegistryMock.sol#L147-L149) does not emit an event

contracts/mock/MarketRegistryMock.sol#L147-L149


 - [ ] ID-173
Setter function [MarketLiquidityRewards.onlyMarketOwner(uint256)](contracts/MarketLiquidityRewards.sol#L39-L46) does not emit an event

contracts/MarketLiquidityRewards.sol#L39-L46


 - [ ] ID-174
Setter function [TLR.mint(address,uint256)](contracts/TLR.sol#L39-L41) does not emit an event

contracts/TLR.sol#L39-L41


 - [ ] ID-175
Setter function [MarketRegistry.ownsMarket(uint256)](contracts/MarketRegistry.sol#L67-L70) does not emit an event

contracts/MarketRegistry.sol#L67-L70


 - [ ] ID-176
Setter function [MarketRegistry._revokeStakeholder(uint256,address,bool)](contracts/MarketRegistry.sol#L1156-L1173) does not emit an event

contracts/MarketRegistry.sol#L1156-L1173


 - [ ] ID-177
Setter function [ReputationManager.initialize(address)](contracts/ReputationManager.sol#L37-L39) does not emit an event

contracts/ReputationManager.sol#L37-L39


 - [ ] ID-178
Setter function [UniswapV3FactoryMock.setPoolMock(address)](contracts/mock/uniswap/UniswapV3FactoryMock.sol#L11-L13) does not emit an event

contracts/mock/uniswap/UniswapV3FactoryMock.sol#L11-L13


 - [ ] ID-179
Setter function [AavePoolAddressProviderMock.setAddressAsProxy(bytes32,address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L166-L168) does not emit an event

contracts/mock/aave/AavePoolAddressProviderMock.sol#L166-L168


 - [ ] ID-180
Setter function [TellerV2SolMock.setLastRepaidTimestamp(uint256,uint32)](contracts/mock/TellerV2SolMock.sol#L323-L325) does not emit an event

contracts/mock/TellerV2SolMock.sol#L323-L325


 - [ ] ID-181
Setter function [TellerASEIP712Verifier.attest(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32)](contracts/EAS/TellerASEIP712Verifier.sol#L68-L101) does not emit an event

contracts/EAS/TellerASEIP712Verifier.sol#L68-L101


 - [ ] ID-182
Setter function [TellerASEIP712Verifier.revoke(bytes32,address,uint8,bytes32,bytes32)](contracts/EAS/TellerASEIP712Verifier.sol#L106-L127) does not emit an event

contracts/EAS/TellerASEIP712Verifier.sol#L106-L127


 - [ ] ID-183
Setter function [LenderCommitmentForwarderMock.acceptCommitmentWithRecipientAndProof(uint256,uint256,uint256,uint256,address,address,uint16,uint32,bytes32[])](contracts/mock/LenderCommitmentForwarderMock.sol#L117-L139) does not emit an event

contracts/mock/LenderCommitmentForwarderMock.sol#L117-L139


 - [ ] ID-184
Setter function [LenderCommitmentGroup_Smart.liquidateDefaultedLoanWithIncentive(uint256,int256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L422-L472) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L422-L472


 - [ ] ID-185
Setter function [CollateralManager.liquidateCollateral(uint256,address)](contracts/CollateralManager.sol#L291-L303) does not emit an event

contracts/CollateralManager.sol#L291-L303


 - [ ] ID-186
Setter function [AavePoolAddressProviderMock.constructor(string,address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L35-L38) does not emit an event

contracts/mock/aave/AavePoolAddressProviderMock.sol#L35-L38


 - [ ] ID-187
Setter function [MarketRegistry._setMarketSettings(uint256,uint32,PaymentType,PaymentCycleType,uint32,uint32,uint16,bool,bool,string)](contracts/MarketRegistry.sol#L965-L985) does not emit an event

contracts/MarketRegistry.sol#L965-L985


 - [ ] ID-188
Setter function [LenderManager.initialize()](contracts/LenderManager.sol#L26-L28) does not emit an event

contracts/LenderManager.sol#L26-L28


 - [ ] ID-189
Setter function [ERC20PresetMinterPauser.mint(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L54-L57) does not emit an event

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L54-L57


 - [ ] ID-190
Setter function [EscrowVault.withdraw(address,uint256)](contracts/EscrowVault.sol#L43-L48) does not emit an event

contracts/EscrowVault.sol#L43-L48


 - [ ] ID-191
Setter function [LenderCommitmentForwarderMock.acceptCommitmentWithRecipient(uint256,uint256,uint256,uint256,address,address,uint16,uint32)](contracts/mock/LenderCommitmentForwarderMock.sol#L94-L115) does not emit an event

contracts/mock/LenderCommitmentForwarderMock.sol#L94-L115


 - [ ] ID-192
Setter function [TellerV2.claimLoanNFT(uint256)](contracts/TellerV2.sol#L578-L594) does not emit an event

contracts/TellerV2.sol#L578-L594


 - [ ] ID-193
Setter function [SmartCommitmentForwarder.acceptCommitmentWithRecipient(address,uint256,uint256,uint256,address,address,uint16,uint32)](contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#L38-L66) does not emit an event

contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#L38-L66


 - [ ] ID-194
Setter function [TellerV2SolMock.repayLoan(uint256,uint256)](contracts/mock/TellerV2SolMock.sol#L149-L157) does not emit an event

contracts/mock/TellerV2SolMock.sol#L149-L157


 - [ ] ID-195
Setter function [TellerV2SolMock.lenderCloseLoanWithRecipient(uint256,address)](contracts/mock/TellerV2SolMock.sol#L107-L111) does not emit an event

contracts/mock/TellerV2SolMock.sol#L107-L111


 - [ ] ID-196
Setter function [TellerV2Autopay.setAutopayFee(uint16)](contracts/TellerV2Autopay.sol#L70-L72) does not emit an event

contracts/TellerV2Autopay.sol#L70-L72


 - [ ] ID-197
Setter function [UniswapV3PoolMock.set_mockSqrtPriceX96(uint160)](contracts/mock/uniswap/UniswapV3PoolMock.sol#L30-L32) does not emit an event

contracts/mock/uniswap/UniswapV3PoolMock.sol#L30-L32


 - [ ] ID-198
Setter function [MarketLiquidityRewards._decrementAllocatedAmount(uint256,uint256)](contracts/MarketLiquidityRewards.sol#L363-L367) does not emit an event

contracts/MarketLiquidityRewards.sol#L363-L367


 - [ ] ID-199
Setter function [ProtocolFeeMock.setProtocolFee(uint16)](contracts/ProtocolFeeMock.sol#L13-L25) does not emit an event

contracts/ProtocolFeeMock.sol#L13-L25


 - [ ] ID-200
Setter function [TellerV2SolMock.submitBid(address,uint256,uint256,uint32,uint16,string,address,Collateral[])](contracts/mock/TellerV2SolMock.sol#L82-L101) does not emit an event

contracts/mock/TellerV2SolMock.sol#L82-L101


 - [ ] ID-201
Setter function [CollateralManager.onlyTellerV2()](contracts/CollateralManager.sol#L68-L71) does not emit an event

contracts/CollateralManager.sol#L68-L71


 - [ ] ID-202
Setter function [LenderCommitmentGroupShares.burn(address,uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L21-L23) does not emit an event

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L21-L23


 - [ ] ID-203
Setter function [CollateralEscrowV1.initialize(uint256)](contracts/escrow/CollateralEscrowV1.sol#L32-L35) does not emit an event

contracts/escrow/CollateralEscrowV1.sol#L32-L35


 - [ ] ID-204
Setter function [LenderManager.registerLoan(uint256,address)](contracts/LenderManager.sol#L40-L46) does not emit an event

contracts/LenderManager.sol#L40-L46


 - [ ] ID-205
Setter function [CollateralManager.commitCollateral(uint256,Collateral[])](contracts/CollateralManager.sol#L119-L134) does not emit an event

contracts/CollateralManager.sol#L119-L134


 - [ ] ID-206
Setter function [EscrowVault.deposit(address,address,uint256)](contracts/EscrowVault.sol#L32-L41) does not emit an event

contracts/EscrowVault.sol#L32-L41


 - [ ] ID-207
Setter function [TellerV2SolMock.setMarketRegistry(address)](contracts/mock/TellerV2SolMock.sol#L30-L32) does not emit an event

contracts/mock/TellerV2SolMock.sol#L30-L32


 - [ ] ID-208
Setter function [AavePoolAddressProviderMock.setPoolConfiguratorImpl(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L170-L172) does not emit an event

contracts/mock/aave/AavePoolAddressProviderMock.sol#L170-L172


 - [ ] ID-209
Setter function [TellerV2SolMock.lenderCloseLoan(uint256)](contracts/mock/TellerV2SolMock.sol#L103-L105) does not emit an event

contracts/mock/TellerV2SolMock.sol#L103-L105


 - [ ] ID-210
Setter function [CollateralManager.initialize(address,address)](contracts/CollateralManager.sol#L80-L87) does not emit an event

contracts/CollateralManager.sol#L80-L87


 - [ ] ID-211
Setter function [MarketRegistry.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,PaymentType,PaymentCycleType,string)](contracts/MarketRegistry.sol#L132-L156) does not emit an event

contracts/MarketRegistry.sol#L132-L156


 - [ ] ID-212
Setter function [TellerV2SolMock.repayLoanFull(uint256)](contracts/mock/TellerV2SolMock.sol#L129-L147) does not emit an event

contracts/mock/TellerV2SolMock.sol#L129-L147


 - [ ] ID-213
Setter function [LenderCommitmentForwarderMock.setCommitment(uint256,ILenderCommitmentForwarder.Commitment)](contracts/mock/LenderCommitmentForwarderMock.sol#L45-L49) does not emit an event

contracts/mock/LenderCommitmentForwarderMock.sol#L45-L49


## initialize-permission
Impact: Low
Confidence: Medium
 - [ ] ID-214
Condition variable is not initialized found in [Initializable._isInitializing()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L162-L164)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L162-L164


 - [ ] ID-215
Condition variable is not initialized found in [Initializable._disableInitializers()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L144-L150)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L144-L150


## input-validation
Impact: Low
Confidence: Medium
 - [ ] ID-216
value assignment lack of validation	[LenderCommitmentForwarder_U1.createCommitmentWithUniswap(ILenderCommitmentForwarder_U1.Commitment,address[],ILenderCommitmentForwarder_U1.PoolRouteConfig[],uint16)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L182-L219)[commitmentUniswapPoolRoutes[commitmentId_].push(_poolRoutes[i])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L201)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L182-L219


 - [ ] ID-217
value assignment lack of validation	[LenderCommitmentForwarderMock.setCommitment(uint256,ILenderCommitmentForwarder.Commitment)](contracts/mock/LenderCommitmentForwarderMock.sol#L45-L49)[commitments[_commitmentId] = _commitment](contracts/mock/LenderCommitmentForwarderMock.sol#L48)

contracts/mock/LenderCommitmentForwarderMock.sol#L45-L49


 - [ ] ID-218
value assignment lack of validation	[MinimalForwarderUpgradeable.execute(MinimalForwarderUpgradeable.ForwardRequest,bytes)](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L52-L77)[_nonces[req.from] = req.nonce + 1](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L58)

node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L52-L77


 - [ ] ID-219
value assignment lack of validation	[TellerV2SolMock.setLastRepaidTimestamp(uint256,uint32)](contracts/mock/TellerV2SolMock.sol#L323-L325)[bids[_bidId].loanDetails.lastRepaidTimestamp = _timestamp](contracts/mock/TellerV2SolMock.sol#L324)

contracts/mock/TellerV2SolMock.sol#L323-L325


 - [ ] ID-220
value assignment lack of validation	[AavePoolMock.setShouldExecuteCallback(bool)](contracts/mock/aave/AavePoolMock.sol#L12-L14)[shouldExecuteCallback = shouldExecute](contracts/mock/aave/AavePoolMock.sol#L13)

contracts/mock/aave/AavePoolMock.sol#L12-L14


## missing-zero-check
Impact: Low
Confidence: Medium
 - [ ] ID-221
variable lacks a zero-check on 		- [FlashRolloverLoan_G3.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L157-L219)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L157-L219


 - [ ] ID-222
variable lacks a zero-check on 		- [TellerV2Autopay.autoPayLoanMinimum(uint256)](contracts/TellerV2Autopay.sol#L109-L143)

contracts/TellerV2Autopay.sol#L109-L143


 - [ ] ID-223
variable lacks a zero-check on 		- [EscrowVault.deposit(address,address,uint256)](contracts/EscrowVault.sol#L32-L41)

contracts/EscrowVault.sol#L32-L41


 - [ ] ID-224
variable lacks a zero-check on 		- [TellerAS.attestByDelegation(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32)](contracts/EAS/TellerAS.sol#L124-L149)

contracts/EAS/TellerAS.sol#L124-L149


 - [ ] ID-225
variable lacks a zero-check on 		- [LenderCommitmentGroup_Smart.burnSharesToWithdrawEarnings(uint256,address)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L396-L415)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L396-L415


 - [ ] ID-226
variable lacks a zero-check on 		- [MarketRegistryMock.setMarketOwner(address)](contracts/mock/MarketRegistryMock.sol#L99-L101)

contracts/mock/MarketRegistryMock.sol#L99-L101


 - [ ] ID-227
variable lacks a zero-check on 		- [UniswapV3PoolMock.set_mockToken0(address)](contracts/mock/uniswap/UniswapV3PoolMock.sol#L34-L36)

contracts/mock/uniswap/UniswapV3PoolMock.sol#L34-L36


 - [ ] ID-228
variable lacks a zero-check on 		- [ReputationManager.updateAccountReputation(address)](contracts/ReputationManager.sol#L77-L84)

contracts/ReputationManager.sol#L77-L84


 - [ ] ID-229
variable lacks a zero-check on 		- [FlashRolloverLoan_G4.constructor(address,address)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L62-L68)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L62-L68


 - [ ] ID-230
variable lacks a zero-check on 		- [WethMock.transferFrom(address,address,uint256)](contracts/mock/WethMock.sol#L63-L83)

contracts/mock/WethMock.sol#L63-L83


 - [ ] ID-231
variable lacks a zero-check on 		- [MarketRegistryMock.setMarketFeeRecipient(address)](contracts/mock/MarketRegistryMock.sol#L103-L105)

contracts/mock/MarketRegistryMock.sol#L103-L105


 - [ ] ID-232
variable lacks a zero-check on 		- [FlashRolloverLoan_G5.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L156-L219)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L156-L219


 - [ ] ID-233
variable lacks a zero-check on 		- [LenderCommitmentForwarder_U1.getUniswapV3PoolAddress(address,address,uint24)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L706-L717)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L706-L717


 - [ ] ID-234
variable lacks a zero-check on 		- [CollateralManager.setCollateralEscrowBeacon(address)](contracts/CollateralManager.sol#L93-L98)

contracts/CollateralManager.sol#L93-L98


 - [ ] ID-235
variable lacks a zero-check on 		- [EscrowVault.withdraw(address,uint256)](contracts/EscrowVault.sol#L43-L48)

contracts/EscrowVault.sol#L43-L48


 - [ ] ID-236
variable lacks a zero-check on 		- [MarketLiquidityRewards.constructor(address,address,address)](contracts/MarketLiquidityRewards.sol#L69-L77)

contracts/MarketLiquidityRewards.sol#L69-L77


 - [ ] ID-237
variable lacks a zero-check on 		- [LenderCommitmentGroup_Smart.constructor(address,address,address)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L143-L151)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L143-L151


 - [ ] ID-238
variable lacks a zero-check on 		- [TellerASEIP712Verifier.attest(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32)](contracts/EAS/TellerASEIP712Verifier.sol#L68-L101)

contracts/EAS/TellerASEIP712Verifier.sol#L68-L101


 - [ ] ID-239
variable lacks a zero-check on 		- [FlashRolloverLoan_G4.rolloverLoanWithFlash(address,uint256,uint256,uint256,FlashRolloverLoan_G4.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L96-L134)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L96-L134


 - [ ] ID-240
variable lacks a zero-check on 		- [ExtensionsContextUpgradeable.revokeExtension(address)](contracts/LenderCommitmentForwarder/extensions/ExtensionsContextUpgradeable.sol#L35-L38)

contracts/LenderCommitmentForwarder/extensions/ExtensionsContextUpgradeable.sol#L35-L38


 - [ ] ID-241
variable lacks a zero-check on 		- [WethMock.approve(address,uint256)](contracts/mock/WethMock.sol#L53-L57)

contracts/mock/WethMock.sol#L53-L57


 - [ ] ID-242
variable lacks a zero-check on 		- [LenderCommitmentGroup_Smart.addPrincipalToCommitmentGroup(uint256,address)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L307-L322)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L307-L322


 - [ ] ID-243
variable lacks a zero-check on 		- [FlashRolloverLoan_G1.rolloverLoanWithFlash(uint256,uint256,uint256,FlashRolloverLoan_G1.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L96-L132)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L96-L132


 - [ ] ID-244
variable lacks a zero-check on 		- [FlashRolloverLoan_G5.rolloverLoanWithFlash(address,uint256,uint256,uint256,FlashRolloverLoan_G5.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L97-L135)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L97-L135


 - [ ] ID-245
variable lacks a zero-check on 		- [TellerV2SolMock.approveMarketForwarder(uint256,address)](contracts/mock/TellerV2SolMock.sol#L43-L47)

contracts/mock/TellerV2SolMock.sol#L43-L47


 - [ ] ID-246
variable lacks a zero-check on 		- [FlashRolloverLoan_G1.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L137-L199)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L137-L199


 - [ ] ID-247
variable lacks a zero-check on 		- [UniswapV3PoolMock.set_mockToken1(address)](contracts/mock/uniswap/UniswapV3PoolMock.sol#L39-L41)

contracts/mock/uniswap/UniswapV3PoolMock.sol#L39-L41


 - [ ] ID-248
variable lacks a zero-check on 		- [UniswapV3FactoryMock.setPoolMock(address)](contracts/mock/uniswap/UniswapV3FactoryMock.sol#L11-L13)

contracts/mock/uniswap/UniswapV3FactoryMock.sol#L11-L13


 - [ ] ID-249
variable lacks a zero-check on 		- [FlashRolloverLoan_G3.rolloverLoanWithFlash(uint256,uint256,uint256,FlashRolloverLoan_G3.AcceptCommitmentArgs)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L100-L136)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L100-L136


 - [ ] ID-250
variable lacks a zero-check on 		- [FlashRolloverLoan_G5.constructor(address,address)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L64-L70)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L64-L70


 - [ ] ID-251
variable lacks a zero-check on 		- [TellerASEIP712Verifier.revoke(bytes32,address,uint8,bytes32,bytes32)](contracts/EAS/TellerASEIP712Verifier.sol#L106-L127)

contracts/EAS/TellerASEIP712Verifier.sol#L106-L127


 - [ ] ID-252
variable lacks a zero-check on 		- [CollateralManager.initialize(address,address)](contracts/CollateralManager.sol#L80-L87)

contracts/CollateralManager.sol#L80-L87


 - [ ] ID-253
variable lacks a zero-check on 		- [FlashRolloverLoan_G4.executeOperation(address,uint256,uint256,address,bytes)](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L155-L218)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L155-L218


 - [ ] ID-254
variable lacks a zero-check on 		- [AavePoolMock.flashLoanSimple(address,address,uint256,bytes,uint16)](contracts/mock/aave/AavePoolMock.sol#L16-L52)

contracts/mock/aave/AavePoolMock.sol#L16-L52


 - [ ] ID-255
variable lacks a zero-check on 		- [LenderCommitmentGroup_Smart.initialize(address,address,uint256,uint32,uint16,uint16,uint16,uint16,uint24,uint32)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L158-L213)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L158-L213


## variable-scope
Impact: Low
Confidence: High
 - [ ] ID-256
Variable '[ERC20Votes._moveVotingPower(address,address,uint256).oldWeight](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L228)' in [ERC20Votes._moveVotingPower(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L221-L237) potentially used before declaration: [(oldWeight,newWeight) = _writeCheckpoint(_checkpoints[dst],_add,amount)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L233)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L228


## reentrancy-same-effect
Impact: Low
Confidence: Medium
 - [ ] ID-257
Reentrancy in [EscrowVault.deposit(address,address,uint256)](contracts/EscrowVault.sol#L32-L41):
	External calls:
	- [ERC20(token).safeTransferFrom(_msgSender(),address(this),amount)](contracts/EscrowVault.sol#L37)

contracts/EscrowVault.sol#L32-L41


 - [ ] ID-258
Reentrancy in [AavePoolMock.flashLoanSimple(address,address,uint256,bytes,uint16)](contracts/mock/aave/AavePoolMock.sol#L16-L52):
	External calls:
	- [IERC20(asset).transfer(receiverAddress,amount)](contracts/mock/aave/AavePoolMock.sol#L25)
	- [success = IFlashLoanSimpleReceiver(receiverAddress).executeOperation(asset,amount,premium,initiator,params)](contracts/mock/aave/AavePoolMock.sol#L31-L32)
	- [IERC20(asset).transferFrom(receiverAddress,address(this),amount + premium)](contracts/mock/aave/AavePoolMock.sol#L37-L41)

contracts/mock/aave/AavePoolMock.sol#L16-L52


## unnecessary-boolean-compare
Impact: Informational
Confidence: High
 - [ ] ID-259
[TellerV2.submitBid(address,uint256,uint256,uint32,uint16,string,address,Collateral[])](contracts/TellerV2.sol#L314-L343) compares to a boolean constant:
	-[require(bool,string)(validation == true,Collateral balance could not be validated)](contracts/TellerV2.sol#L339-L342)

contracts/TellerV2.sol#L314-L343


 - [ ] ID-260
[LenderCommitmentGroup_Smart.bidIsActiveForGroup(uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L136-L140) compares to a boolean constant:
	-[require(bool,string)(activeBids[_bidId] == true,Bid is not active for group)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L137)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L136-L140


 - [ ] ID-261
[AavePoolMock.flashLoanSimple(address,address,uint256,bytes,uint16)](contracts/mock/aave/AavePoolMock.sol#L16-L52) compares to a boolean constant:
	-[require(bool,string)(success == true,executeOperation failed)](contracts/mock/aave/AavePoolMock.sol#L34)

contracts/mock/aave/AavePoolMock.sol#L16-L52


## centralized-risk-informational
Impact: Informational
Confidence: High
 - [ ] ID-262
	- [AavePoolAddressProviderMock.setPoolDataProvider(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L145-L153)

contracts/mock/aave/AavePoolAddressProviderMock.sol#L145-L153


 - [ ] ID-263
	- [AavePoolAddressProviderMock.setPriceOracleSentinel(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L126-L137)

contracts/mock/aave/AavePoolAddressProviderMock.sol#L126-L137


 - [ ] ID-264
	- [TellerV2Context.setTrustedMarketForwarder(uint256,address)](contracts/TellerV2Context.sol#L75-L84)

contracts/TellerV2Context.sol#L75-L84


 - [ ] ID-265
	- [ProtocolFeeMock.setProtocolFee(uint16)](contracts/ProtocolFeeMock.sol#L13-L25)

contracts/ProtocolFeeMock.sol#L13-L25


 - [ ] ID-266
	- [AavePoolAddressProviderMock.setAddress(bytes32,address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L60-L68)

contracts/mock/aave/AavePoolAddressProviderMock.sol#L60-L68


 - [ ] ID-267
	- [AavePoolAddressProviderMock.setPriceOracle(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L86-L94)

contracts/mock/aave/AavePoolAddressProviderMock.sol#L86-L94


 - [ ] ID-268
	- [AavePoolAddressProviderMock.setACLManager(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L102-L106)

contracts/mock/aave/AavePoolAddressProviderMock.sol#L102-L106


 - [ ] ID-269
	- [TellerV2Autopay.setAutoPayEnabled(uint256,bool)](contracts/TellerV2Autopay.sol#L94-L103)

contracts/TellerV2Autopay.sol#L94-L103


 - [ ] ID-270
	- [CollateralManager.deployAndDeposit(uint256)](contracts/CollateralManager.sol#L184-L206)

contracts/CollateralManager.sol#L184-L206


 - [ ] ID-271
	- [CollateralEscrowV1.depositAsset(CollateralType,address,uint256,uint256)](contracts/escrow/CollateralEscrowV1.sol#L51-L76)

contracts/escrow/CollateralEscrowV1.sol#L51-L76


 - [ ] ID-272
	- [ProtocolFee.setProtocolFee(uint16)](contracts/ProtocolFee.sol#L44-L51)

contracts/ProtocolFee.sol#L44-L51


 - [ ] ID-273
	- [CollateralEscrowV1.withdraw(address,uint256,address)](contracts/escrow/CollateralEscrowV1.sol#L84-L103)

contracts/escrow/CollateralEscrowV1.sol#L84-L103


 - [ ] ID-274
	- [AavePoolAddressProviderMock.setACLAdmin(address)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L114-L118)

contracts/mock/aave/AavePoolAddressProviderMock.sol#L114-L118


## centralized-risk-other
Impact: Informational
Confidence: High
 - [ ] ID-275
	- [LenderCommitmentGroup_Smart.pauseBorrowing()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L793-L795)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L793-L795


 - [ ] ID-276
	- [LenderManager.registerLoan(uint256,address)](contracts/LenderManager.sol#L40-L46)

contracts/LenderManager.sol#L40-L46


 - [ ] ID-277
	- [Ownable.renounceOwnership()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63


 - [ ] ID-278
	- [Ownable.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72


 - [ ] ID-279
	- [LenderCommitmentGroupShares.mint(address,uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L17-L19)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L17-L19


 - [ ] ID-280
	- [AccessControl.grantRole(bytes32,address)](node_modules/@openzeppelin/contracts/access/AccessControl.sol#L144-L146)

node_modules/@openzeppelin/contracts/access/AccessControl.sol#L144-L146


 - [ ] ID-281
	- [TellerV2.pauseProtocol()](contracts/TellerV2.sol#L713-L715)

contracts/TellerV2.sol#L713-L715


 - [ ] ID-282
	- [OwnableUpgradeable.transferOwnership(address)](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77


 - [ ] ID-283
	- [CollateralManager.commitCollateral(uint256,Collateral)](contracts/CollateralManager.sol#L142-L152)

contracts/CollateralManager.sol#L142-L152


 - [ ] ID-284
	- [CollateralManager.commitCollateral(uint256,Collateral[])](contracts/CollateralManager.sol#L119-L134)

contracts/CollateralManager.sol#L119-L134


 - [ ] ID-285
	- [TellerV2.unpauseProtocol()](contracts/TellerV2.sol#L720-L722)

contracts/TellerV2.sol#L720-L722


 - [ ] ID-286
	- [TellerV2Context.renounceMarketForwarder(uint256,address)](contracts/TellerV2Context.sol#L108-L115)

contracts/TellerV2Context.sol#L108-L115


 - [ ] ID-287
	- [OwnableUpgradeable.renounceOwnership()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68


 - [ ] ID-288
	- [TLR.burn(address,uint256)](contracts/TLR.sol#L54-L56)

contracts/TLR.sol#L54-L56


 - [ ] ID-289
	- [UpgradeableBeacon.upgradeTo(address)](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L49-L52)

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L49-L52


 - [ ] ID-290
	- [CollateralManager.lenderClaimCollateral(uint256)](contracts/CollateralManager.sol#L271-L283)

contracts/CollateralManager.sol#L271-L283


 - [ ] ID-291
	- [TLR.mint(address,uint256)](contracts/TLR.sol#L39-L41)

contracts/TLR.sol#L39-L41


 - [ ] ID-292
	- [AavePoolAddressProviderMock.setMarketId(string)](contracts/mock/aave/AavePoolAddressProviderMock.sol#L46-L52)

contracts/mock/aave/AavePoolAddressProviderMock.sol#L46-L52


 - [ ] ID-293
	- [LenderCommitmentGroupShares.burn(address,uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L21-L23)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L21-L23


 - [ ] ID-294
	- [TellerV2Autopay.setAutopayFee(uint16)](contracts/TellerV2Autopay.sol#L70-L72)

contracts/TellerV2Autopay.sol#L70-L72


 - [ ] ID-295
	- [ERC20PresetMinterPauser.pause()](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L68-L71)

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L68-L71


 - [ ] ID-296
	- [CollateralManager.liquidateCollateral(uint256,address)](contracts/CollateralManager.sol#L291-L303)

contracts/CollateralManager.sol#L291-L303


 - [ ] ID-297
	- [AccessControl.revokeRole(bytes32,address)](node_modules/@openzeppelin/contracts/access/AccessControl.sol#L159-L161)

node_modules/@openzeppelin/contracts/access/AccessControl.sol#L159-L161


 - [ ] ID-298
	- [ERC20PresetMinterPauser.unpause()](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L82-L85)

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L82-L85


 - [ ] ID-299
	- [TellerV2.cancelBid(uint256)](contracts/TellerV2.sol#L428-L440)

contracts/TellerV2.sol#L428-L440


 - [ ] ID-300
	- [LenderCommitmentGroup_Smart.unpauseBorrowing()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L800-L802)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L800-L802


 - [ ] ID-301
	- [ERC20PresetMinterPauser.mint(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L54-L57)

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L54-L57


## dead-function
Impact: Informational
Confidence: Medium
 - [ ] ID-302
[ContextUpgradeable._msgData()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L27-L29) is never used and should be removed

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L27-L29


 - [ ] ID-303
[Initializable._isInitializing()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L162-L164) is never used and should be removed

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L162-L164


 - [ ] ID-304
[TellerV2Context._msgDataForMarket(uint256)](contracts/TellerV2Context.sol#L152-L163) is never used and should be removed

contracts/TellerV2Context.sol#L152-L163


 - [ ] ID-305
[LenderCommitmentForwarder_U1.getPriceX96FromSqrtPriceX96(uint160)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L806-L812) is never used and should be removed

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L806-L812


 - [ ] ID-306
[ERC1967Upgrade._upgradeToAndCallUUPS(address,bytes,bool)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L81-L99) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L81-L99


 - [ ] ID-307
[ERC1967Upgrade._changeAdmin(address)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L133-L136) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L133-L136


 - [ ] ID-308
[EIP712Upgradeable.__EIP712_init(string,string)](node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol#L50-L52) is never used and should be removed

node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol#L50-L52


 - [ ] ID-309
[ERC1967Upgrade._getImplementation()](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L38-L40) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L38-L40


 - [ ] ID-310
[ContextUpgradeable.__Context_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L21-L22) is never used and should be removed

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L21-L22


 - [ ] ID-311
[ERC1967Upgrade._upgradeToAndCall(address,bytes,bool)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L65-L74) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L65-L74


 - [ ] ID-312
[TellerV2._msgData()](contracts/TellerV2.sol#L1327-L1335) is never used and should be removed

contracts/TellerV2.sol#L1327-L1335


 - [ ] ID-313
[BeaconProxy._beacon()](node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#L37-L39) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#L37-L39


 - [ ] ID-314
[ERC2771ContextUpgradeable._msgData()](contracts/ERC2771ContextUpgradeable.sol#L51-L63) is never used and should be removed

contracts/ERC2771ContextUpgradeable.sol#L51-L63


 - [ ] ID-315
[ERC1967Upgrade._upgradeTo(address)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L55-L58) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L55-L58


 - [ ] ID-316
[TellerV2MarketForwarder_G3._acceptBidWithRepaymentListener(uint256,address,address)](contracts/TellerV2MarketForwarder_G3.sol#L33-L56) is never used and should be removed

contracts/TellerV2MarketForwarder_G3.sol#L33-L56


 - [ ] ID-317
[ERC1967Upgrade._getAdmin()](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L116-L118) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L116-L118


 - [ ] ID-318
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-319
[MinimalForwarderUpgradeable.__MinimalForwarder_init()](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L35-L37) is never used and should be removed

node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L35-L37


 - [ ] ID-320
[ERC1967Upgrade._setAdmin(address)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L123-L126) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L123-L126


 - [ ] ID-321
[Initializable._getInitializedVersion()](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L155-L157) is never used and should be removed

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L155-L157


 - [ ] ID-322
[MinimalForwarderUpgradeable.__MinimalForwarder_init_unchained()](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L39) is never used and should be removed

node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L39


 - [ ] ID-323
[BeaconProxy._setBeacon(address,bytes)](node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#L58-L60) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#L58-L60


 - [ ] ID-324
[TellerV2MarketForwarder_G1._submitBid(TellerV2MarketForwarder_G1.CreateLoanArgs,address)](contracts/TellerV2MarketForwarder_G1.sol#L78-L99) is never used and should be removed

contracts/TellerV2MarketForwarder_G1.sol#L78-L99


 - [ ] ID-325
[ERC1967Upgrade._setImplementation(address)](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L45-L48) is never used and should be removed

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L45-L48


 - [ ] ID-326
[ContextUpgradeable.__Context_init()](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L18-L19) is never used and should be removed

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L18-L19


## error-msg
Impact: Informational
Confidence: Medium
 - [ ] ID-327
require() missing error messages
	 - [require(bool)(newTimestamp >= timestamp)](contracts/libraries/DateTimeLib.sol#L360)

contracts/libraries/DateTimeLib.sol#L360


 - [ ] ID-328
require() missing error messages
	 - [require(bool)(fromTimestamp <= toTimestamp)](contracts/libraries/DateTimeLib.sol#L499)

contracts/libraries/DateTimeLib.sol#L499


 - [ ] ID-329
require() missing error messages
	 - [require(bool)(balanceOf[msg.sender] >= wad)](contracts/mock/WethMock.sol#L43)

contracts/mock/WethMock.sol#L43


 - [ ] ID-330
require() missing error messages
	 - [require(bool)(denominator > prod1)](node_modules/@openzeppelin/contracts/utils/math/Math.sol#L78)

node_modules/@openzeppelin/contracts/utils/math/Math.sol#L78


 - [ ] ID-331
require() missing error messages
	 - [require(bool)(newTimestamp <= timestamp)](contracts/libraries/DateTimeLib.sol#L402)

contracts/libraries/DateTimeLib.sol#L402


 - [ ] ID-332
require() missing error messages
	 - [require(bool)(year >= 1970)](contracts/libraries/DateTimeLib.sol#L61)

contracts/libraries/DateTimeLib.sol#L61


 - [ ] ID-333
require() missing error messages
	 - [require(bool)(newTimestamp >= timestamp)](contracts/libraries/DateTimeLib.sol#L302)

contracts/libraries/DateTimeLib.sol#L302


 - [ ] ID-334
require() missing error messages
	 - [require(bool)(denominator > 0)](contracts/libraries/uniswap/FullMath.sol#L34)

contracts/libraries/uniswap/FullMath.sol#L34


 - [ ] ID-335
require() missing error messages
	 - [require(bool)(newTimestamp >= timestamp)](contracts/libraries/DateTimeLib.sol#L351)

contracts/libraries/DateTimeLib.sol#L351


 - [ ] ID-336
require() missing error messages
	 - [require(bool)(newTimestamp <= timestamp)](contracts/libraries/DateTimeLib.sol#L380)

contracts/libraries/DateTimeLib.sol#L380


 - [ ] ID-337
require() missing error messages
	 - [require(bool)(denominator > prod1)](node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#L78)

node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#L78


 - [ ] ID-338
require() missing error messages
	 - [require(bool)(newTimestamp <= timestamp)](contracts/libraries/DateTimeLib.sol#L438)

contracts/libraries/DateTimeLib.sol#L438


 - [ ] ID-339
require() missing error messages
	 - [require(bool)(newTimestamp >= timestamp)](contracts/libraries/DateTimeLib.sol#L324)

contracts/libraries/DateTimeLib.sol#L324


 - [ ] ID-340
require() missing error messages
	 - [require(bool)(result < type()(uint256).max)](contracts/libraries/uniswap/FullMath.sol#L121)

contracts/libraries/uniswap/FullMath.sol#L121


 - [ ] ID-341
require() missing error messages
	 - [require(bool)(denominator > prod1)](contracts/libraries/uniswap/FullMath.sol#L43)

contracts/libraries/uniswap/FullMath.sol#L43


 - [ ] ID-342
require() missing error messages
	 - [require(bool)(fromTimestamp <= toTimestamp)](contracts/libraries/DateTimeLib.sol#L481)

contracts/libraries/DateTimeLib.sol#L481


 - [ ] ID-343
require() missing error messages
	 - [require(bool)(newTimestamp >= timestamp)](contracts/libraries/DateTimeLib.sol#L342)

contracts/libraries/DateTimeLib.sol#L342


 - [ ] ID-344
require() missing error messages
	 - [require(bool)(fromTimestamp <= toTimestamp)](contracts/libraries/DateTimeLib.sol#L446)

contracts/libraries/DateTimeLib.sol#L446


 - [ ] ID-345
require() missing error messages
	 - [require(bool)(fromTimestamp <= toTimestamp)](contracts/libraries/DateTimeLib.sol#L490)

contracts/libraries/DateTimeLib.sol#L490


 - [ ] ID-346
require() missing error messages
	 - [require(bool)(newTimestamp <= timestamp)](contracts/libraries/DateTimeLib.sol#L411)

contracts/libraries/DateTimeLib.sol#L411


 - [ ] ID-347
require() missing error messages
	 - [require(bool)(newTimestamp <= timestamp)](contracts/libraries/DateTimeLib.sol#L429)

contracts/libraries/DateTimeLib.sol#L429


 - [ ] ID-348
require() missing error messages
	 - [require(bool)(newTimestamp <= timestamp)](contracts/libraries/DateTimeLib.sol#L420)

contracts/libraries/DateTimeLib.sol#L420


 - [ ] ID-349
require() missing error messages
	 - [require(bool)(fromTimestamp <= toTimestamp)](contracts/libraries/DateTimeLib.sol#L472)

contracts/libraries/DateTimeLib.sol#L472


 - [ ] ID-350
require() missing error messages
	 - [require(bool)(newTimestamp >= timestamp)](contracts/libraries/DateTimeLib.sol#L333)

contracts/libraries/DateTimeLib.sol#L333


 - [ ] ID-351
require() missing error messages
	 - [require(bool)(fromTimestamp <= toTimestamp)](contracts/libraries/DateTimeLib.sol#L457)

contracts/libraries/DateTimeLib.sol#L457


## no-license
Impact: Informational
Confidence: High
 - [ ] ID-352
key[AavePoolMock](contracts/mock/aave/AavePoolMock.sol#L7-L53) does not specify SPDX license identifier
contracts/mock/aave/AavePoolMock.sol#L7-L53


 - [ ] ID-353
key[IFlashRolloverLoan_G4](contracts/interfaces/IFlashRolloverLoan_G4.sol#L3-L12) does not specify SPDX license identifier
contracts/interfaces/IFlashRolloverLoan_G4.sol#L3-L12


## price-manipulation-info
Impact: Informational
Confidence: Medium
 - [ ] ID-354
Potential price manipulation risk:
	- In function commitCollateral
		-- [(validation_,None) = checkBalances(borrower,_collateralInfo)](contracts/CollateralManager.sol#L125) have potential price manipulated risk from validation_ and call None which could influence variable:validation_
	- In function commitCollateral
		-- [validation_ = _checkBalance(borrower,_collateralInfo)](contracts/CollateralManager.sol#L148) have potential price manipulated risk from validation_ and call None which could influence variable:validation_
	- In function revalidateCollateral
		-- [(validation_,None) = _checkBalances(borrower,collateralInfos,true)](contracts/CollateralManager.sol#L165) have potential price manipulated risk from validation_ and call None which could influence variable:validation_
	- In function _checkBalances
		-- [isValidated = _checkBalance(_borrowerAddress,_collateralInfo[i])](contracts/CollateralManager.sol#L494-L497) have potential price manipulated risk from isValidated and call None which could influence variable:checks_

contracts/CollateralManager.sol#L125


 - [ ] ID-355
Potential price manipulation risk:
	- In function _delegate
		-- [delegatorBalance = balanceOf(delegator)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L213) have potential price manipulated risk from delegatorBalance and call None which could influence variable:delegatorBalance

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L213


## uncontrolled-resource-consumption
Impact: Informational
Confidence: Medium
 - [ ] ID-356
Potential DoS Gas Limit Attack occur in[LenderCommitmentForwarder_G1._addBorrowersToCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L281-L289)[BEGIN_LOOP](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L285-L287)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L281-L289


 - [ ] ID-357
Potential DoS Gas Limit Attack occur in[CollateralManager._checkBalances(address,Collateral[],bool)](contracts/CollateralManager.sol#L486-L507)[BEGIN_LOOP](contracts/CollateralManager.sol#L493-L506)

contracts/CollateralManager.sol#L486-L507


 - [ ] ID-358
Potential DoS Gas Limit Attack occur in[LenderCommitmentForwarder_U1.createCommitmentWithUniswap(ILenderCommitmentForwarder_U1.Commitment,address[],ILenderCommitmentForwarder_U1.PoolRouteConfig[],uint16)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L182-L219)[BEGIN_LOOP](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L200-L202)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L182-L219


 - [ ] ID-359
Potential DoS Gas Limit Attack occur in[CollateralManager._withdraw(uint256,address)](contracts/CollateralManager.sol#L416-L442)[BEGIN_LOOP](contracts/CollateralManager.sol#L417-L441)

contracts/CollateralManager.sol#L416-L442


 - [ ] ID-360
Potential DoS Gas Limit Attack occur in[LenderCommitmentForwarder_G2._removeBorrowersFromCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L263-L271)[BEGIN_LOOP](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L267-L269)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L263-L271


 - [ ] ID-361
Potential DoS Gas Limit Attack occur in[LenderCommitmentForwarder_G2._addBorrowersToCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L248-L256)[BEGIN_LOOP](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L252-L254)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L248-L256


 - [ ] ID-362
Potential DoS Gas Limit Attack occur in[LenderCommitmentForwarder_G1._removeBorrowersFromCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L296-L304)[BEGIN_LOOP](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L300-L302)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L296-L304


 - [ ] ID-363
Potential DoS Gas Limit Attack occur in[LenderCommitmentForwarder_U1._addBorrowersToCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L291-L299)[BEGIN_LOOP](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L295-L297)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L291-L299


 - [ ] ID-364
Potential DoS Gas Limit Attack occur in[LenderCommitmentForwarder_U1._removeBorrowersFromCommitmentAllowlist(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L306-L314)[BEGIN_LOOP](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L310-L312)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L306-L314


 - [ ] ID-365
Potential DoS Gas Limit Attack occur in[ReputationManager.updateAccountReputation(address)](contracts/ReputationManager.sol#L77-L84)[BEGIN_LOOP](contracts/ReputationManager.sol#L81-L83)

contracts/ReputationManager.sol#L77-L84


## unnecessary-public-function-modifier
Impact: Informational
Confidence: High
 - [ ] ID-366
function:[ERC20PresetMinterPauser.unpause()](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L82-L85)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L82-L85


 - [ ] ID-367
function:[MarketRegistry.getMarketOwner(uint256)](contracts/MarketRegistry.sol#L749-L757)is public and can be replaced with external 

contracts/MarketRegistry.sol#L749-L757


 - [ ] ID-368
function:[ERC20PresetMinterPauser.mint(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L54-L57)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L54-L57


 - [ ] ID-369
function:[MarketRegistry.isMarketClosed(uint256)](contracts/MarketRegistry.sol#L275-L282)is public and can be replaced with external 

contracts/MarketRegistry.sol#L275-L282


 - [ ] ID-370
function:[TellerV2.pauseProtocol()](contracts/TellerV2.sol#L713-L715)is public and can be replaced with external 

contracts/TellerV2.sol#L713-L715


 - [ ] ID-371
function:[MarketRegistry.isMarketOpen(uint256)](contracts/MarketRegistry.sol#L260-L269)is public and can be replaced with external 

contracts/MarketRegistry.sol#L260-L269


 - [ ] ID-372
function:[OwnableUpgradeable.renounceOwnership()](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L66-L68


 - [ ] ID-373
function:[ERC20.symbol()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L70-L72)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L70-L72


 - [ ] ID-374
function:[ERC721Upgradeable.tokenURI(uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L98-L103)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L98-L103


 - [ ] ID-375
function:[AccessControl.grantRole(bytes32,address)](node_modules/@openzeppelin/contracts/access/AccessControl.sol#L144-L146)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/AccessControl.sol#L144-L146


 - [ ] ID-376
function:[TellerV2SolMock.getRepaymentListenerForBid(uint256)](contracts/mock/TellerV2SolMock.sol#L337-L341)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L337-L341


 - [ ] ID-377
function:[SmartCommitmentForwarder.acceptCommitmentWithRecipient(address,uint256,uint256,uint256,address,address,uint16,uint32)](contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#L38-L66)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#L38-L66


 - [ ] ID-378
function:[MarketRegistry.getMarketURI(uint256)](contracts/MarketRegistry.sol#L798-L805)is public and can be replaced with external 

contracts/MarketRegistry.sol#L798-L805


 - [ ] ID-379
function:[LenderCommitmentForwarderMock.getCommitmentMaxPrincipal(uint256)](contracts/mock/LenderCommitmentForwarderMock.sol#L75-L81)is public and can be replaced with external 

contracts/mock/LenderCommitmentForwarderMock.sol#L75-L81


 - [ ] ID-380
function:[EscrowVault.deposit(address,address,uint256)](contracts/EscrowVault.sol#L32-L41)is public and can be replaced with external 

contracts/EscrowVault.sol#L32-L41


 - [ ] ID-381
function:[ReputationManager.updateAccountReputation(address,uint256)](contracts/ReputationManager.sol#L86-L92)is public and can be replaced with external 

contracts/ReputationManager.sol#L86-L92


 - [ ] ID-382
function:[ERC20Permit.nonces(address)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L73-L75)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L73-L75


 - [ ] ID-383
function:[ERC165Upgradeable.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#L32-L34)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#L32-L34


 - [ ] ID-384
function:[TellerV2SolMock.submitBid(address,uint256,uint256,uint32,uint16,string,address,Collateral[])](contracts/mock/TellerV2SolMock.sol#L82-L101)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L82-L101


 - [ ] ID-385
function:[ERC20PresetMinterPauser.pause()](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L68-L71)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L68-L71


 - [ ] ID-386
function:[MarketRegistryMock.getMarketplaceFee(uint256)](contracts/mock/MarketRegistryMock.sol#L95-L97)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L95-L97


 - [ ] ID-387
function:[CollateralManager.commitCollateral(uint256,Collateral[])](contracts/CollateralManager.sol#L119-L134)is public and can be replaced with external 

contracts/CollateralManager.sol#L119-L134


 - [ ] ID-388
function:[ERC20Votes.checkpoints(address,uint32)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L43-L45)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L43-L45


 - [ ] ID-389
function:[TellerV2.unpauseProtocol()](contracts/TellerV2.sol#L720-L722)is public and can be replaced with external 

contracts/TellerV2.sol#L720-L722


 - [ ] ID-390
function:[MarketLiquidityRewards.updateAllocation(uint256,uint256,uint256,uint32,uint32)](contracts/MarketLiquidityRewards.sol#L126-L147)is public and can be replaced with external 

contracts/MarketLiquidityRewards.sol#L126-L147


 - [ ] ID-391
function:[ERC20Votes.getPastTotalSupply(uint256)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L89-L92)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L89-L92


 - [ ] ID-392
function:[ERC20.totalSupply()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L94-L96)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L94-L96


 - [ ] ID-393
function:[ERC721Upgradeable.symbol()](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L91-L93)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L91-L93


 - [ ] ID-394
function:[LenderCommitmentForwarder_G2.deleteCommitment(uint256)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L277-L284)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L277-L284


 - [ ] ID-395
function:[LenderCommitmentGroup_Smart.pauseBorrowing()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L793-L795)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L793-L795


 - [ ] ID-396
function:[TellerV2.calculateAmountDue(uint256,uint256)](contracts/TellerV2.sol#L1001-L1021)is public and can be replaced with external 

contracts/TellerV2.sol#L1001-L1021


 - [ ] ID-397
function:[AccessControl.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts/access/AccessControl.sol#L77-L79)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/AccessControl.sol#L77-L79


 - [ ] ID-398
function:[UniswapV3PoolMock.set_mockToken0(address)](contracts/mock/uniswap/UniswapV3PoolMock.sol#L34-L36)is public and can be replaced with external 

contracts/mock/uniswap/UniswapV3PoolMock.sol#L34-L36


 - [ ] ID-399
function:[MarketRegistry.isVerifiedLender(uint256,address)](contracts/MarketRegistry.sol#L883-L896)is public and can be replaced with external 

contracts/MarketRegistry.sol#L883-L896


 - [ ] ID-400
function:[ERC20.approve(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L136-L140)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L136-L140


 - [ ] ID-401
function:[ERC20.balanceOf(address)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L101-L103)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L101-L103


 - [ ] ID-402
function:[LenderCommitmentForwarder_U1.getCommitmentUniswapPoolRoute(uint256,uint256)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L666-L676)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L666-L676


 - [ ] ID-403
function:[LenderCommitmentForwarderMock.acceptCommitmentWithRecipient(uint256,uint256,uint256,uint256,address,address,uint16,uint32)](contracts/mock/LenderCommitmentForwarderMock.sol#L94-L115)is public and can be replaced with external 

contracts/mock/LenderCommitmentForwarderMock.sol#L94-L115


 - [ ] ID-404
function:[MarketRegistryMock.mock_setLenderIsVerified(bool)](contracts/mock/MarketRegistryMock.sol#L147-L149)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L147-L149


 - [ ] ID-405
function:[TellerAS.attestByDelegation(address,bytes32,uint256,bytes32,bytes,address,uint8,bytes32,bytes32)](contracts/EAS/TellerAS.sol#L124-L149)is public and can be replaced with external 

contracts/EAS/TellerAS.sol#L124-L149


 - [ ] ID-406
function:[LenderCommitmentForwarderMock.acceptCommitmentWithRecipientAndProof(uint256,uint256,uint256,uint256,address,address,uint16,uint32,bytes32[])](contracts/mock/LenderCommitmentForwarderMock.sol#L117-L139)is public and can be replaced with external 

contracts/mock/LenderCommitmentForwarderMock.sol#L117-L139


 - [ ] ID-407
function:[MarketRegistryMock.getMarketFeeRecipient(uint256)](contracts/mock/MarketRegistryMock.sol#L55-L61)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L55-L61


 - [ ] ID-408
function:[OwnableUpgradeable.transferOwnership(address)](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L74-L77


 - [ ] ID-409
function:[ERC20Burnable.burnFrom(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#L35-L38)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#L35-L38


 - [ ] ID-410
function:[AccessControl.revokeRole(bytes32,address)](node_modules/@openzeppelin/contracts/access/AccessControl.sol#L159-L161)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/AccessControl.sol#L159-L161


 - [ ] ID-411
function:[TellerV2SolMock.setRepaymentListenerForBid(uint256,address)](contracts/mock/TellerV2SolMock.sol#L343-L345)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L343-L345


 - [ ] ID-412
function:[ReputationManager.getDefaultedLoanIds(address)](contracts/ReputationManager.sol#L50-L57)is public and can be replaced with external 

contracts/ReputationManager.sol#L50-L57


 - [ ] ID-413
function:[UniswapV3PoolMock.set_mockToken1(address)](contracts/mock/uniswap/UniswapV3PoolMock.sol#L39-L41)is public and can be replaced with external 

contracts/mock/uniswap/UniswapV3PoolMock.sol#L39-L41


 - [ ] ID-414
function:[ERC20Votes.getVotes(address)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L64-L67)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L64-L67


 - [ ] ID-415
function:[WethMock.deposit()](contracts/mock/WethMock.sol#L37-L40)is public and can be replaced with external 

contracts/mock/WethMock.sol#L37-L40


 - [ ] ID-416
function:[WethMock.approve(address,uint256)](contracts/mock/WethMock.sol#L53-L57)is public and can be replaced with external 

contracts/mock/WethMock.sol#L53-L57


 - [ ] ID-417
function:[TellerV2.getLoanBorrower(uint256)](contracts/TellerV2.sol#L1160-L1166)is public and can be replaced with external 

contracts/TellerV2.sol#L1160-L1166


 - [ ] ID-418
function:[ERC721Upgradeable.setApprovalForAll(address,bool)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L141-L143)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L141-L143


 - [ ] ID-419
function:[AccessControlEnumerable.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L21-L23)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L21-L23


 - [ ] ID-420
function:[MarketRegistry.getPaymentDefaultDuration(uint256)](contracts/MarketRegistry.sol#L830-L837)is public and can be replaced with external 

contracts/MarketRegistry.sol#L830-L837


 - [ ] ID-421
function:[ReputationManager.getCurrentDefaultLoanIds(address)](contracts/ReputationManager.sol#L68-L75)is public and can be replaced with external 

contracts/ReputationManager.sol#L68-L75


 - [ ] ID-422
function:[MarketRegistryMock.getPaymentCycle(uint256)](contracts/mock/MarketRegistryMock.sol#L71-L77)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L71-L77


 - [ ] ID-423
function:[TellerV2MarketForwarder_G2.getTellerV2()](contracts/TellerV2MarketForwarder_G2.sol#L37-L39)is public and can be replaced with external 

contracts/TellerV2MarketForwarder_G2.sol#L37-L39


 - [ ] ID-424
function:[TellerV2SolMock.setLastRepaidTimestamp(uint256,uint32)](contracts/mock/TellerV2SolMock.sol#L323-L325)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L323-L325


 - [ ] ID-425
function:[MarketRegistryMock.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,string)](contracts/mock/MarketRegistryMock.sol#L126-L135)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L126-L135


 - [ ] ID-426
function:[ERC20Votes.getPastVotes(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L76-L79)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L76-L79


 - [ ] ID-427
function:[TellerAS.revoke(bytes32)](contracts/EAS/TellerAS.sol#L154-L156)is public and can be replaced with external 

contracts/EAS/TellerAS.sol#L154-L156


 - [ ] ID-428
function:[TellerAS.revokeByDelegation(bytes32,address,uint8,bytes32,bytes32)](contracts/EAS/TellerAS.sol#L161-L171)is public and can be replaced with external 

contracts/EAS/TellerAS.sol#L161-L171


 - [ ] ID-429
function:[ERC20Votes.numCheckpoints(address)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L50-L52)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L50-L52


 - [ ] ID-430
function:[TellerV2SolMock.setMarketRegistry(address)](contracts/mock/TellerV2SolMock.sol#L30-L32)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L30-L32


 - [ ] ID-431
function:[UniswapV3PoolMock.token0()](contracts/mock/uniswap/UniswapV3PoolMock.sol#L43-L45)is public and can be replaced with external 

contracts/mock/uniswap/UniswapV3PoolMock.sol#L43-L45


 - [ ] ID-432
function:[ERC20.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L158-L167)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L158-L167


 - [ ] ID-433
function:[MinimalForwarderUpgradeable.execute(MinimalForwarderUpgradeable.ForwardRequest,bytes)](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L52-L77)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L52-L77


 - [ ] ID-434
function:[TellerV2.getMetadataURI(uint256)](contracts/TellerV2.sol#L255-L271)is public and can be replaced with external 

contracts/TellerV2.sol#L255-L271


 - [ ] ID-435
function:[TellerV2MarketForwarder_G2.getTellerV2MarketOwner(uint256)](contracts/TellerV2MarketForwarder_G2.sol#L45-L47)is public and can be replaced with external 

contracts/TellerV2MarketForwarder_G2.sol#L45-L47


 - [ ] ID-436
function:[MarketRegistry.getAllVerifiedLendersForMarket(uint256,uint256,uint256)](contracts/MarketRegistry.sol#L927-L936)is public and can be replaced with external 

contracts/MarketRegistry.sol#L927-L936


 - [ ] ID-437
function:[UpgradeableBeacon.upgradeTo(address)](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L49-L52)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L49-L52


 - [ ] ID-438
function:[LenderCommitmentForwarder_G1.deleteCommitment(uint256)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L310-L317)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L310-L317


 - [ ] ID-439
function:[Ownable.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable.sol#L69-L72


 - [ ] ID-440
function:[ERC20.name()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L62-L64)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L62-L64


 - [ ] ID-441
function:[LenderCommitmentForwarder_G1.createCommitment(LenderCommitmentForwarder_G1.Commitment,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L183-L209)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L183-L209


 - [ ] ID-442
function:[MarketRegistry.transferMarketOwnership(uint256,address)](contracts/MarketRegistry.sol#L481-L487)is public and can be replaced with external 

contracts/MarketRegistry.sol#L481-L487


 - [ ] ID-443
function:[MarketRegistry.getAllVerifiedBorrowersForMarket(uint256,uint256,uint256)](contracts/MarketRegistry.sol#L945-L953)is public and can be replaced with external 

contracts/MarketRegistry.sol#L945-L953


 - [ ] ID-444
function:[WethMock.transfer(address,uint256)](contracts/mock/WethMock.sol#L59-L61)is public and can be replaced with external 

contracts/mock/WethMock.sol#L59-L61


 - [ ] ID-445
function:[TellerV2SolMock.getLoanDetails(uint256)](contracts/mock/TellerV2SolMock.sol#L233-L239)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L233-L239


 - [ ] ID-446
function:[MarketRegistry.getMarketFeeRecipient(uint256)](contracts/MarketRegistry.sol#L778-L791)is public and can be replaced with external 

contracts/MarketRegistry.sol#L778-L791


 - [ ] ID-447
function:[LenderCommitmentForwarderMock.setCommitment(uint256,ILenderCommitmentForwarder.Commitment)](contracts/mock/LenderCommitmentForwarderMock.sol#L45-L49)is public and can be replaced with external 

contracts/mock/LenderCommitmentForwarderMock.sol#L45-L49


 - [ ] ID-448
function:[LenderCommitmentGroup_Smart.unpauseBorrowing()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L800-L802)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L800-L802


 - [ ] ID-449
function:[TellerAS.attest(address,bytes32,uint256,bytes32,bytes)](contracts/EAS/TellerAS.sol#L103-L119)is public and can be replaced with external 

contracts/EAS/TellerAS.sol#L103-L119


 - [ ] ID-450
function:[MarketLiquidityRewards.deallocateRewards(uint256,uint256)](contracts/MarketLiquidityRewards.sol#L170-L198)is public and can be replaced with external 

contracts/MarketLiquidityRewards.sol#L170-L198


 - [ ] ID-451
function:[LenderCommitmentForwarder_G2.acceptCommitmentWithProof(uint256,uint256,uint256,uint256,address,uint16,uint32,bytes32[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L411-L433)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L411-L433


 - [ ] ID-452
function:[LenderCommitmentForwarder_U1.addCommitmentBorrowers(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L264-L269)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L264-L269


 - [ ] ID-453
function:[MarketRegistry.getMarketplaceFee(uint256)](contracts/MarketRegistry.sol#L867-L874)is public and can be replaced with external 

contracts/MarketRegistry.sol#L867-L874


 - [ ] ID-454
function:[AccessControlEnumerable.getRoleMember(bytes32,uint256)](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L37-L39)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L37-L39


 - [ ] ID-455
function:[LenderCommitmentForwarder_U1.getUniswapV3PoolAddress(address,address,uint24)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L706-L717)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L706-L717


 - [ ] ID-456
function:[Ownable.renounceOwnership()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable.sol#L61-L63


 - [ ] ID-457
function:[UniswapV3FactoryMock.getPool(address,address,uint24)](contracts/mock/uniswap/UniswapV3FactoryMock.sol#L4-L9)is public and can be replaced with external 

contracts/mock/uniswap/UniswapV3FactoryMock.sol#L4-L9


 - [ ] ID-458
function:[ERC721Upgradeable.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L155-L164)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L155-L164


 - [ ] ID-459
function:[UniswapV3PoolMock.slot0()](contracts/mock/uniswap/UniswapV3PoolMock.sol#L52-L63)is public and can be replaced with external 

contracts/mock/uniswap/UniswapV3PoolMock.sol#L52-L63


 - [ ] ID-460
function:[ERC721Upgradeable.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L57-L62)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L57-L62


 - [ ] ID-461
function:[TellerV2.submitBid(address,uint256,uint256,uint32,uint16,string,address,Collateral[])](contracts/TellerV2.sol#L314-L343)is public and can be replaced with external 

contracts/TellerV2.sol#L314-L343


 - [ ] ID-462
function:[ERC20.increaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L181-L185)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L181-L185


 - [ ] ID-463
function:[MarketRegistryMock.getPaymentDefaultDuration(uint256)](contracts/mock/MarketRegistryMock.sol#L79-L85)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L79-L85


 - [ ] ID-464
function:[LenderCommitmentForwarder_U1.updateCommitment(uint256,ILenderCommitmentForwarder_U1.Commitment)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L226-L257)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L226-L257


 - [ ] ID-465
function:[MarketRegistry.updateMarketSettings(uint256,uint32,PaymentType,PaymentCycleType,uint32,uint32,uint16,bool,bool,string)](contracts/MarketRegistry.sol#L502-L526)is public and can be replaced with external 

contracts/MarketRegistry.sol#L502-L526


 - [ ] ID-466
function:[MarketRegistryMock.createMarket(address,uint32,uint32,uint32,uint16,bool,bool,PaymentType,PaymentCycleType,string)](contracts/mock/MarketRegistryMock.sol#L113-L124)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L113-L124


 - [ ] ID-467
function:[UniswapV3FactoryMock.setPoolMock(address)](contracts/mock/uniswap/UniswapV3FactoryMock.sol#L11-L13)is public and can be replaced with external 

contracts/mock/uniswap/UniswapV3FactoryMock.sol#L11-L13


 - [ ] ID-468
function:[MarketRegistry.getMarketAttestationRequirements(uint256)](contracts/MarketRegistry.sol#L730-L742)is public and can be replaced with external 

contracts/MarketRegistry.sol#L730-L742


 - [ ] ID-469
function:[TellerV2SolMock.getBorrowerActiveLoanIds(address)](contracts/mock/TellerV2SolMock.sol#L241-L245)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L241-L245


 - [ ] ID-470
function:[MarketRegistry.getMarketData(uint256)](contracts/MarketRegistry.sol#L702-L724)is public and can be replaced with external 

contracts/MarketRegistry.sol#L702-L724


 - [ ] ID-471
function:[LenderCommitmentGroupShares.decimals()](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L25-L27)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L25-L27


 - [ ] ID-472
function:[LenderCommitmentForwarder_G2.updateCommitment(uint256,ILenderCommitmentForwarder.Commitment)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L183-L214)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L183-L214


 - [ ] ID-473
function:[LenderCommitmentForwarder_U1.acceptCommitmentWithProof(uint256,uint256,uint256,uint256,address,uint16,uint32,bytes32[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L454-L476)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L454-L476


 - [ ] ID-474
function:[ProtocolFee.protocolFee()](contracts/ProtocolFee.sol#L36-L38)is public and can be replaced with external 

contracts/ProtocolFee.sol#L36-L38


 - [ ] ID-475
function:[MarketRegistryMock.getBidExpirationTime(uint256)](contracts/mock/MarketRegistryMock.sol#L87-L93)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L87-L93


 - [ ] ID-476
function:[ERC721Upgradeable.name()](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L84-L86)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L84-L86


 - [ ] ID-477
function:[UpgradeableBeacon.implementation()](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L35-L37)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L35-L37


 - [ ] ID-478
function:[LenderManager.registerLoan(uint256,address)](contracts/LenderManager.sol#L40-L46)is public and can be replaced with external 

contracts/LenderManager.sol#L40-L46


 - [ ] ID-479
function:[AccessControlEnumerable.getRoleMemberCount(bytes32)](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L45-L47)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L45-L47


 - [ ] ID-480
function:[LenderCommitmentGroup_Smart.liquidateDefaultedLoanWithIncentive(uint256,int256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L422-L472)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L422-L472


 - [ ] ID-481
function:[LenderCommitmentForwarder_U1.createCommitmentWithUniswap(ILenderCommitmentForwarder_U1.Commitment,address[],ILenderCommitmentForwarder_U1.PoolRouteConfig[],uint16)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L182-L219)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L182-L219


 - [ ] ID-482
function:[UniswapV3PoolMock.set_mockSqrtPriceX96(uint160)](contracts/mock/uniswap/UniswapV3PoolMock.sol#L30-L32)is public and can be replaced with external 

contracts/mock/uniswap/UniswapV3PoolMock.sol#L30-L32


 - [ ] ID-483
function:[MarketRegistryMock.isMarketClosed(uint256)](contracts/mock/MarketRegistryMock.sol#L34-L36)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L34-L36


 - [ ] ID-484
function:[ERC165.supportsInterface(bytes4)](node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L26-L28)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L26-L28


 - [ ] ID-485
function:[WethMock.withdraw(uint256)](contracts/mock/WethMock.sol#L42-L47)is public and can be replaced with external 

contracts/mock/WethMock.sol#L42-L47


 - [ ] ID-486
function:[MinimalForwarderUpgradeable.getNonce(address)](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L41-L43)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L41-L43


 - [ ] ID-487
function:[ReputationManager.getCurrentDelinquentLoanIds(address)](contracts/ReputationManager.sol#L59-L66)is public and can be replaced with external 

contracts/ReputationManager.sol#L59-L66


 - [ ] ID-488
function:[MarketRegistryMock.setMarketFeeRecipient(address)](contracts/mock/MarketRegistryMock.sol#L103-L105)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L103-L105


 - [ ] ID-489
function:[TellerV2.isPaymentLate(uint256)](contracts/TellerV2.sol#L1049-L1052)is public and can be replaced with external 

contracts/TellerV2.sol#L1049-L1052


 - [ ] ID-490
function:[WethMock.totalSupply()](contracts/mock/WethMock.sol#L49-L51)is public and can be replaced with external 

contracts/mock/WethMock.sol#L49-L51


 - [ ] ID-491
function:[MarketRegistryMock.isVerifiedLender(uint256,address)](contracts/mock/MarketRegistryMock.sol#L22-L28)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L22-L28


 - [ ] ID-492
function:[MarketLiquidityRewards.increaseAllocationAmount(uint256,uint256)](contracts/MarketLiquidityRewards.sol#L154-L163)is public and can be replaced with external 

contracts/MarketLiquidityRewards.sol#L154-L163


 - [ ] ID-493
function:[LenderCommitmentForwarderMock.getCommitmentMarketId(uint256)](contracts/mock/LenderCommitmentForwarderMock.sol#L59-L65)is public and can be replaced with external 

contracts/mock/LenderCommitmentForwarderMock.sol#L59-L65


 - [ ] ID-494
function:[LenderCommitmentForwarder_U1.acceptCommitment(uint256,uint256,uint256,uint256,address,uint16,uint32)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L371-L391)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L371-L391


 - [ ] ID-495
function:[LenderCommitmentForwarder_G2.removeCommitmentBorrowers(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L233-L241)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L233-L241


 - [ ] ID-496
function:[LenderCommitmentForwarder_G2.acceptCommitment(uint256,uint256,uint256,uint256,address,uint16,uint32)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L328-L348)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L328-L348


 - [ ] ID-497
function:[TellerV2MarketForwarder_G1.getTellerV2MarketOwner(uint256)](contracts/TellerV2MarketForwarder_G1.sol#L50-L52)is public and can be replaced with external 

contracts/TellerV2MarketForwarder_G1.sol#L50-L52


 - [ ] ID-498
function:[ERC20Votes.delegate(address)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L139-L141)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L139-L141


 - [ ] ID-499
function:[ERC721Upgradeable.balanceOf(address)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L67-L70)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L67-L70


 - [ ] ID-500
function:[MarketRegistryMock.isVerifiedBorrower(uint256,address)](contracts/mock/MarketRegistryMock.sol#L38-L44)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L38-L44


 - [ ] ID-501
function:[MarketLiquidityRewards.allocateRewards(IMarketLiquidityRewards.RewardAllocation)](contracts/MarketLiquidityRewards.sol#L86-L116)is public and can be replaced with external 

contracts/MarketLiquidityRewards.sol#L86-L116


 - [ ] ID-502
function:[TellerV2MarketForwarder_G1.getTellerV2()](contracts/TellerV2MarketForwarder_G1.sol#L42-L44)is public and can be replaced with external 

contracts/TellerV2MarketForwarder_G1.sol#L42-L44


 - [ ] ID-503
function:[LenderCommitmentForwarderMock.getCommitmentLender(uint256)](contracts/mock/LenderCommitmentForwarderMock.sol#L51-L57)is public and can be replaced with external 

contracts/mock/LenderCommitmentForwarderMock.sol#L51-L57


 - [ ] ID-504
function:[LenderCommitmentForwarder_G2.createCommitment(ILenderCommitmentForwarder.Commitment,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L150-L176)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L150-L176


 - [ ] ID-505
function:[ERC20Votes.delegateBySig(address,uint256,uint256,uint8,bytes32,bytes32)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L146-L163)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L146-L163


 - [ ] ID-506
function:[MarketRegistryMock.setMarketOwner(address)](contracts/mock/MarketRegistryMock.sol#L99-L101)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L99-L101


 - [ ] ID-507
function:[ReputationManager.getDelinquentLoanIds(address)](contracts/ReputationManager.sol#L41-L48)is public and can be replaced with external 

contracts/ReputationManager.sol#L41-L48


 - [ ] ID-508
function:[MarketRegistryMock.getPaymentType(uint256)](contracts/mock/MarketRegistryMock.sol#L107-L111)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L107-L111


 - [ ] ID-509
function:[LenderCommitmentForwarder_U1.removeCommitmentBorrowers(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L276-L284)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L276-L284


 - [ ] ID-510
function:[MarketRegistry.isVerifiedBorrower(uint256,address)](contracts/MarketRegistry.sol#L905-L918)is public and can be replaced with external 

contracts/MarketRegistry.sol#L905-L918


 - [ ] ID-511
function:[MarketLiquidityRewards.getRewardTokenAmount(uint256)](contracts/MarketLiquidityRewards.sol#L474-L481)is public and can be replaced with external 

contracts/MarketLiquidityRewards.sol#L474-L481


 - [ ] ID-512
function:[AavePoolMock.setShouldExecuteCallback(bool)](contracts/mock/aave/AavePoolMock.sol#L12-L14)is public and can be replaced with external 

contracts/mock/aave/AavePoolMock.sol#L12-L14


 - [ ] ID-513
function:[TellerV2SolMock.calculateAmountDue(uint256,uint256)](contracts/mock/TellerV2SolMock.sol#L163-L179)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L163-L179


 - [ ] ID-514
function:[MarketRegistryMock.getMarketOwner(uint256)](contracts/mock/MarketRegistryMock.sol#L46-L53)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L46-L53


 - [ ] ID-515
function:[LenderCommitmentForwarder_U1.getCommitmentPoolOracleLtvRatio(uint256)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L696-L702)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L696-L702


 - [ ] ID-516
function:[MarketRegistryMock.mock_setGlobalMarketsClosed(bool)](contracts/mock/MarketRegistryMock.sol#L139-L141)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L139-L141


 - [ ] ID-517
function:[ERC20.decreaseAllowance(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L201-L210)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L201-L210


 - [ ] ID-518
function:[LenderCommitmentForwarder_G1.updateCommitment(uint256,LenderCommitmentForwarder_G1.Commitment)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L216-L247)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L216-L247


 - [ ] ID-519
function:[TellerV2.submitBid(address,uint256,uint256,uint32,uint16,string,address)](contracts/TellerV2.sol#L283-L301)is public and can be replaced with external 

contracts/TellerV2.sol#L283-L301


 - [ ] ID-520
function:[ERC20.transfer(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L113-L117)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L113-L117


 - [ ] ID-521
function:[LenderCommitmentForwarder_G1.removeCommitmentBorrowers(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L266-L274)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L266-L274


 - [ ] ID-522
function:[LenderManagerMock.ownerOf(uint256)](contracts/mock/LenderManagerMock.sol#L22-L29)is public and can be replaced with external 

contracts/mock/LenderManagerMock.sol#L22-L29


 - [ ] ID-523
function:[LenderCommitmentGroup_Smart.getRequiredCollateral(uint256)](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L738-L746)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L738-L746


 - [ ] ID-524
function:[MarketRegistry.getPaymentType(uint256)](contracts/MarketRegistry.sol#L844-L851)is public and can be replaced with external 

contracts/MarketRegistry.sol#L844-L851


 - [ ] ID-525
function:[LenderCommitmentForwarder_U1.getAllCommitmentUniswapPoolRoutes(uint256)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L683-L689)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L683-L689


 - [ ] ID-526
function:[TellerV2SolMock.getBidState(uint256)](contracts/mock/TellerV2SolMock.sol#L224-L231)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L224-L231


 - [ ] ID-527
function:[UniswapV3PoolMock.token1()](contracts/mock/uniswap/UniswapV3PoolMock.sol#L47-L49)is public and can be replaced with external 

contracts/mock/uniswap/UniswapV3PoolMock.sol#L47-L49


 - [ ] ID-528
function:[ERC721Upgradeable.approve(address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L117-L127)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L117-L127


 - [ ] ID-529
function:[CollateralManager.commitCollateral(uint256,Collateral)](contracts/CollateralManager.sol#L142-L152)is public and can be replaced with external 

contracts/CollateralManager.sol#L142-L152


 - [ ] ID-530
function:[MarketRegistry.getBidExpirationTime(uint256)](contracts/MarketRegistry.sol#L853-L860)is public and can be replaced with external 

contracts/MarketRegistry.sol#L853-L860


 - [ ] ID-531
function:[TellerV2.getLoanDefaultTimestamp(uint256)](contracts/TellerV2.sol#L1234-L1246)is public and can be replaced with external 

contracts/TellerV2.sol#L1234-L1246


 - [ ] ID-532
function:[TellerV2SolMock.lenderAcceptBid(uint256)](contracts/mock/TellerV2SolMock.sol#L199-L222)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L199-L222


 - [ ] ID-533
function:[CollateralManager.getCollateralAmount(uint256,address)](contracts/CollateralManager.sol#L243-L251)is public and can be replaced with external 

contracts/CollateralManager.sol#L243-L251


 - [ ] ID-534
function:[TellerV2SolMock.isPaymentLate(uint256)](contracts/mock/TellerV2SolMock.sol#L261)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L261


 - [ ] ID-535
function:[MarketRegistryMock.mock_setBorrowerIsVerified(bool)](contracts/mock/MarketRegistryMock.sol#L143-L145)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L143-L145


 - [ ] ID-536
function:[V2Calculations.calculateNextDueDate(uint32,uint32,uint32,uint32,PaymentCycleType)](contracts/libraries/V2Calculations.sol#L171-L215)is public and can be replaced with external 

contracts/libraries/V2Calculations.sol#L171-L215


 - [ ] ID-537
function:[TellerAS.isAttestationActive(bytes32)](contracts/EAS/TellerAS.sol#L200-L211)is public and can be replaced with external 

contracts/EAS/TellerAS.sol#L200-L211


 - [ ] ID-538
function:[TellerV2Context.hasApprovedMarketForwarder(uint256,address,address)](contracts/TellerV2Context.sol#L59-L67)is public and can be replaced with external 

contracts/TellerV2Context.sol#L59-L67


 - [ ] ID-539
function:[TellerASMock.isAttestationActive(bytes32)](contracts/mock/TellerASMock.sol#L19-L26)is public and can be replaced with external 

contracts/mock/TellerASMock.sol#L19-L26


 - [ ] ID-540
function:[ERC721Upgradeable.ownerOf(uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L75-L79)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L75-L79


 - [ ] ID-541
function:[LenderCommitmentForwarder_U1.deleteCommitment(uint256)](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L320-L327)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L320-L327


 - [ ] ID-542
function:[TellerV2SolMock.isLoanLiquidateable(uint256)](contracts/mock/TellerV2SolMock.sol#L254-L259)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L254-L259


 - [ ] ID-543
function:[ERC20.decimals()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L87-L89)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L87-L89


 - [ ] ID-544
function:[ERC721Upgradeable.safeTransferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L169-L175)is public and can be replaced with external 

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L169-L175


 - [ ] ID-545
function:[MarketRegistry.getPaymentCycle(uint256)](contracts/MarketRegistry.sol#L813-L823)is public and can be replaced with external 

contracts/MarketRegistry.sol#L813-L823


 - [ ] ID-546
function:[LenderCommitmentForwarder_G2.addCommitmentBorrowers(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L221-L226)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L221-L226


 - [ ] ID-547
function:[MarketRegistry.setMarketFeeRecipient(uint256,address)](contracts/MarketRegistry.sol#L536-L542)is public and can be replaced with external 

contracts/MarketRegistry.sol#L536-L542


 - [ ] ID-548
function:[MarketRegistryMock.isMarketOpen(uint256)](contracts/mock/MarketRegistryMock.sol#L30-L32)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L30-L32


 - [ ] ID-549
function:[TellerV2.calculateAmountOwed(uint256,uint256)](contracts/TellerV2.sol#L973-L994)is public and can be replaced with external 

contracts/TellerV2.sol#L973-L994


 - [ ] ID-550
function:[TellerV2SolMock.repayLoan(uint256,uint256)](contracts/mock/TellerV2SolMock.sol#L149-L157)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L149-L157


 - [ ] ID-551
function:[CollateralEscrowV1.initialize(uint256)](contracts/escrow/CollateralEscrowV1.sol#L32-L35)is public and can be replaced with external 

contracts/escrow/CollateralEscrowV1.sol#L32-L35


 - [ ] ID-552
function:[AccessControl.renounceRole(bytes32,address)](node_modules/@openzeppelin/contracts/access/AccessControl.sol#L179-L183)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/AccessControl.sol#L179-L183


 - [ ] ID-553
function:[MarketRegistry.closeMarket(uint256)](contracts/MarketRegistry.sol#L248-L254)is public and can be replaced with external 

contracts/MarketRegistry.sol#L248-L254


 - [ ] ID-554
function:[LenderCommitmentForwarderMock.getCommitmentAcceptedPrincipal(uint256)](contracts/mock/LenderCommitmentForwarderMock.sol#L67-L73)is public and can be replaced with external 

contracts/mock/LenderCommitmentForwarderMock.sol#L67-L73


 - [ ] ID-555
function:[LenderCommitmentForwarder_G1.addCommitmentBorrowers(uint256,address[])](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L254-L259)is public and can be replaced with external 

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L254-L259


 - [ ] ID-556
function:[MarketRegistryMock.closeMarket(uint256)](contracts/mock/MarketRegistryMock.sol#L137)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L137


 - [ ] ID-557
function:[MarketRegistryMock.getMarketURI(uint256)](contracts/mock/MarketRegistryMock.sol#L63-L69)is public and can be replaced with external 

contracts/mock/MarketRegistryMock.sol#L63-L69


 - [ ] ID-558
function:[TellerV2Autopay.setAutopayFee(uint16)](contracts/TellerV2Autopay.sol#L70-L72)is public and can be replaced with external 

contracts/TellerV2Autopay.sol#L70-L72


 - [ ] ID-559
function:[TellerV2SolMock.isLoanDefaulted(uint256)](contracts/mock/TellerV2SolMock.sol#L247-L252)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L247-L252


 - [ ] ID-560
function:[ERC20Burnable.burn(uint256)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#L20-L22)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#L20-L22


 - [ ] ID-561
function:[TellerV2SolMock.calculateAmountOwed(uint256,uint256)](contracts/mock/TellerV2SolMock.sol#L181-L197)is public and can be replaced with external 

contracts/mock/TellerV2SolMock.sol#L181-L197


 - [ ] ID-562
function:[ERC20Permit.permit(address,address,uint256,uint256,uint8,bytes32,bytes32)](node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L49-L68)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L49-L68


## unused-event
Impact: Informational
Confidence: High
 - [ ] ID-563
[MarketRegistrySetPaymentCycleDuration(uint256,uint32)](contracts/MarketRegistry.sol#L82) is never emitted in [MarketRegistry](contracts/MarketRegistry.sol#L19-L1260)

contracts/MarketRegistry.sol#L82


## version-only
Impact: Informational
Confidence: High
 - [ ] ID-564
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan.sol#L2)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan.sol#L2


 - [ ] ID-565
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/EscrowVault.sol#L1)

contracts/EscrowVault.sol#L1


 - [ ] ID-566
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/ECDSAUpgradeable.sol#L4


 - [ ] ID-567
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/type-imports.sol#L1)

contracts/type-imports.sol#L1


 - [ ] ID-568
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IProtocolFee.sol#L2)

contracts/interfaces/IProtocolFee.sol#L2


 - [ ] ID-569
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TellerV2Storage.sol#L1)

contracts/TellerV2Storage.sol#L1


 - [ ] ID-570
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/CollateralManagerMock.sol#L1)

contracts/mock/CollateralManagerMock.sol#L1


 - [ ] ID-571
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/IEscrowVault.sol#L2)

contracts/interfaces/IEscrowVault.sol#L2


 - [ ] ID-572
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/ITellerV2.sol#L2)

contracts/interfaces/ITellerV2.sol#L2


 - [ ] ID-573
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/ReputationManagerMock.sol#L1)

contracts/mock/ReputationManagerMock.sol#L1


 - [ ] ID-574
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IMarketLiquidityRewards.sol#L2)

contracts/interfaces/IMarketLiquidityRewards.sol#L2


 - [ ] ID-575
	Pragma confirmed better, here is pragma solidity^0.8.1. [^0.8.1](node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/AddressUpgradeable.sol#L4


 - [ ] ID-576
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/IEAS.sol#L1)

contracts/interfaces/IEAS.sol#L1


 - [ ] ID-577
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol#L4


 - [ ] ID-578
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol#L4


 - [ ] ID-579
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/access/IAccessControlEnumerable.sol#L4)

node_modules/@openzeppelin/contracts/access/IAccessControlEnumerable.sol#L4


 - [ ] ID-580
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/IASRegistry.sol#L1)

contracts/interfaces/IASRegistry.sol#L1


 - [ ] ID-581
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/libraries/WadRayMath.sol#L2)

contracts/libraries/WadRayMath.sol#L2


 - [ ] ID-582
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/ITellerV2Context.sol#L2)

contracts/interfaces/ITellerV2Context.sol#L2


 - [ ] ID-583
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L4)

node_modules/@openzeppelin/contracts/access/AccessControlEnumerable.sol#L4


 - [ ] ID-584
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/EAS/TellerASEIP712Verifier.sol#L1)

contracts/EAS/TellerASEIP712Verifier.sol#L1


 - [ ] ID-585
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/StringsUpgradeable.sol#L4


 - [ ] ID-586
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L1)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G1.sol#L1


 - [ ] ID-587
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ILenderManager.sol#L1)

contracts/interfaces/ILenderManager.sol#L1


 - [ ] ID-588
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/governance/utils/IVotes.sol#L3)

node_modules/@openzeppelin/contracts/governance/utils/IVotes.sol#L3


 - [ ] ID-589
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/libraries/uniswap/FullMath.sol#L2)

contracts/libraries/uniswap/FullMath.sol#L2


 - [ ] ID-590
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/ProtocolFeeMock.sol#L1)

contracts/ProtocolFeeMock.sol#L1


 - [ ] ID-591
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/MarketRegistryMock.sol#L1)

contracts/mock/MarketRegistryMock.sol#L1


 - [ ] ID-592
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/cryptography/EIP712.sol#L4)

node_modules/@openzeppelin/contracts/utils/cryptography/EIP712.sol#L4


 - [ ] ID-593
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L4)

node_modules/@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol#L4


 - [ ] ID-594
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts/proxy/utils/Initializable.sol#L4)

node_modules/@openzeppelin/contracts/proxy/utils/Initializable.sol#L4


 - [ ] ID-595
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/mock/TellerV2SolMock.sol#L3)

contracts/mock/TellerV2SolMock.sol#L3


 - [ ] ID-596
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/access/IAccessControl.sol#L4)

node_modules/@openzeppelin/contracts/access/IAccessControl.sol#L4


 - [ ] ID-597
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol#L4


 - [ ] ID-598
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TLR.sol#L1)

contracts/TLR.sol#L1


 - [ ] ID-599
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L2)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G4.sol#L2


 - [ ] ID-600
	Pragma confirmed better, here is pragma solidity>=0.5.0. [>=0.5.0](contracts/interfaces/uniswap/pool/IUniswapV3PoolState.sol#L2)

contracts/interfaces/uniswap/pool/IUniswapV3PoolState.sol#L2


 - [ ] ID-601
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#L5)

node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#L5


 - [ ] ID-602
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#L4


 - [ ] ID-603
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/Proxy.sol#L4)

node_modules/@openzeppelin/contracts/proxy/Proxy.sol#L4


 - [ ] ID-604
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol#L4


 - [ ] ID-605
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/draft-IERC20PermitUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/draft-IERC20PermitUpgradeable.sol#L4


 - [ ] ID-606
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/escrow/CollateralEscrowV1.sol#L1)

contracts/escrow/CollateralEscrowV1.sol#L1


 - [ ] ID-607
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/LenderCommitmentForwarderMock.sol#L2)

contracts/mock/LenderCommitmentForwarderMock.sol#L2


 - [ ] ID-608
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol#L4


 - [ ] ID-609
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/EAS/TellerASResolver.sol#L1)

contracts/EAS/TellerASResolver.sol#L1


 - [ ] ID-610
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Votes.sol#L4


 - [ ] ID-611
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TellerV0Storage.sol#L1)

contracts/TellerV0Storage.sol#L1


 - [ ] ID-612
	Pragma confirmed better, here is pragma solidity^0.8.9. [^0.8.9](node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/metatx/ERC2771ContextUpgradeable.sol#L4


 - [ ] ID-613
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Pausable.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Pausable.sol#L4


 - [ ] ID-614
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/MerkleProofUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/cryptography/MerkleProofUpgradeable.sol#L4


 - [ ] ID-615
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G2.sol#L2)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G2.sol#L2


 - [ ] ID-616
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/libraries/uniswap/TickMath.sol#L2)

contracts/libraries/uniswap/TickMath.sol#L2


 - [ ] ID-617
	Pragma confirmed better, here is pragma solidity>=0.5.0. [>=0.5.0](contracts/interfaces/uniswap/pool/IUniswapV3PoolEvents.sol#L2)

contracts/interfaces/uniswap/pool/IUniswapV3PoolEvents.sol#L2


 - [ ] ID-618
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-IERC20Permit.sol#L4


 - [ ] ID-619
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L4


 - [ ] ID-620
	Pragma confirmed better, here is pragma solidity>=0.5.0. [>=0.5.0](contracts/interfaces/uniswap/pool/IUniswapV3PoolDerivedState.sol#L2)

contracts/interfaces/uniswap/pool/IUniswapV3PoolDerivedState.sol#L2


 - [ ] ID-621
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/security/Pausable.sol#L4)

node_modules/@openzeppelin/contracts/security/Pausable.sol#L4


 - [ ] ID-622
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/ILenderCommitmentGroup.sol#L2)

contracts/interfaces/ILenderCommitmentGroup.sol#L2


 - [ ] ID-623
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol#L4


 - [ ] ID-624
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/metatx/MinimalForwarderUpgradeable.sol#L4


 - [ ] ID-625
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/ITellerV2Storage.sol#L2)

contracts/interfaces/ITellerV2Storage.sol#L2


 - [ ] ID-626
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/LenderManager.sol#L1)

contracts/LenderManager.sol#L1


 - [ ] ID-627
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/cryptography/ECDSA.sol#L4)

node_modules/@openzeppelin/contracts/utils/cryptography/ECDSA.sol#L4


 - [ ] ID-628
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G3.sol#L1)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G3.sol#L1


 - [ ] ID-629
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L1)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_U1.sol#L1


 - [ ] ID-630
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L2)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G1.sol#L2


 - [ ] ID-631
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ISmartCommitmentForwarder.sol#L3)

contracts/interfaces/ISmartCommitmentForwarder.sol#L3


 - [ ] ID-632
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/CollateralManager.sol#L1)

contracts/CollateralManager.sol#L1


 - [ ] ID-633
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IReputationManager.sol#L2)

contracts/interfaces/IReputationManager.sol#L2


 - [ ] ID-634
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L4


 - [ ] ID-635
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ILoanRepaymentListener.sol#L2)

contracts/interfaces/ILoanRepaymentListener.sol#L2


 - [ ] ID-636
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol#L4


 - [ ] ID-637
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/IASResolver.sol#L1)

contracts/interfaces/IASResolver.sol#L1


 - [ ] ID-638
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/Counters.sol#L4)

node_modules/@openzeppelin/contracts/utils/Counters.sol#L4


 - [ ] ID-639
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IMarketRegistry.sol#L2)

contracts/interfaces/IMarketRegistry.sol#L2


 - [ ] ID-640
	Pragma confirmed better, here is pragma solidity>=0.5.0. [>=0.5.0](contracts/interfaces/uniswap/pool/IUniswapV3PoolActions.sol#L2)

contracts/interfaces/uniswap/pool/IUniswapV3PoolActions.sol#L2


 - [ ] ID-641
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol#L4


 - [ ] ID-642
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IWETH.sol#L2)

contracts/interfaces/IWETH.sol#L2


 - [ ] ID-643
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder.sol#L2)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder.sol#L2


 - [ ] ID-644
	Pragma confirmed better, here is pragma solidity^0.8.1. [^0.8.1](node_modules/@openzeppelin/contracts/utils/Address.sol#L4)

node_modules/@openzeppelin/contracts/utils/Address.sol#L4


 - [ ] ID-645
	Pragma confirmed better, here is pragma solidity^0.8.9. [^0.8.9](contracts/MetaForwarder.sol#L2)

contracts/MetaForwarder.sol#L2


 - [ ] ID-646
	Pragma confirmed better, here is pragma solidity>=0.5.0. [>=0.5.0](contracts/interfaces/uniswap/pool/IUniswapV3PoolOwnerActions.sol#L2)

contracts/interfaces/uniswap/pool/IUniswapV3PoolOwnerActions.sol#L2


 - [ ] ID-647
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TellerV2MarketForwarder_G3.sol#L1)

contracts/TellerV2MarketForwarder_G3.sol#L1


 - [ ] ID-648
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol#L5)

node_modules/@openzeppelin/contracts-upgradeable/utils/structs/EnumerableSetUpgradeable.sol#L5


 - [ ] ID-649
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/WethMock.sol#L20)

contracts/mock/WethMock.sol#L20


 - [ ] ID-650
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ICommitmentRolloverLoan.sol#L2)

contracts/interfaces/ICommitmentRolloverLoan.sol#L2


 - [ ] ID-651
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IFlashRolloverLoan.sol#L2)

contracts/interfaces/IFlashRolloverLoan.sol#L2


 - [ ] ID-652
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/aave/IPoolAddressesProvider.sol#L2)

contracts/interfaces/aave/IPoolAddressesProvider.sol#L2


 - [ ] ID-653
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L1)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarder_G2.sol#L1


 - [ ] ID-654
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/MarketLiquidityRewards.sol#L2)

contracts/MarketLiquidityRewards.sol#L2


 - [ ] ID-655
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#L4


 - [ ] ID-656
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/aave/DataTypes.sol#L2)

contracts/interfaces/aave/DataTypes.sol#L2


 - [ ] ID-657
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/LenderCommitmentForwarderStaging.sol#L2)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarderStaging.sol#L2


 - [ ] ID-658
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/presets/ERC20PresetMinterPauser.sol#L4


 - [ ] ID-659
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#L4)

node_modules/@openzeppelin/contracts/utils/StorageSlot.sol#L4


 - [ ] ID-660
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/escrow/ICollateralEscrowV1.sol#L2)

contracts/interfaces/escrow/ICollateralEscrowV1.sol#L2


 - [ ] ID-661
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ILenderCommitmentForwarder.sol#L2)

contracts/interfaces/ILenderCommitmentForwarder.sol#L2


 - [ ] ID-662
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/EAS/TellerAS.sol#L1)

contracts/EAS/TellerAS.sol#L1


 - [ ] ID-663
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/libraries/V2Calculations.sol#L1)

contracts/libraries/V2Calculations.sol#L1


 - [ ] ID-664
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/extensions/IERC20MetadataUpgradeable.sol#L4


 - [ ] ID-665
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ICollateralManager.sol#L2)

contracts/interfaces/ICollateralManager.sol#L2


 - [ ] ID-666
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ILenderCommitmentForwarder_U1.sol#L2)

contracts/interfaces/ILenderCommitmentForwarder_U1.sol#L2


 - [ ] ID-667
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IUniswapV2Router.sol#L2)

contracts/interfaces/IUniswapV2Router.sol#L2


 - [ ] ID-668
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/IEASEIP712Verifier.sol#L1)

contracts/interfaces/IEASEIP712Verifier.sol#L1


 - [ ] ID-669
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/Context.sol#L4)

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


 - [ ] ID-670
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/LenderCommitmentForwarderAlpha.sol#L2)

contracts/LenderCommitmentForwarder/LenderCommitmentForwarderAlpha.sol#L2


 - [ ] ID-671
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/aave/AavePoolMock.sol#L1)

contracts/mock/aave/AavePoolMock.sol#L1


 - [ ] ID-672
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IExtensionsContext.sol#L2)

contracts/interfaces/IExtensionsContext.sol#L2


 - [ ] ID-673
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/ReputationManager.sol#L2)

contracts/ReputationManager.sol#L2


 - [ ] ID-674
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC1155/IERC1155Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC1155/IERC1155Upgradeable.sol#L4


 - [ ] ID-675
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L4)

node_modules/@openzeppelin/contracts/utils/introspection/ERC165.sol#L4


 - [ ] ID-676
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TellerV2Autopay.sol#L1)

contracts/TellerV2Autopay.sol#L1


 - [ ] ID-677
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/introspection/ERC165Upgradeable.sol#L4


 - [ ] ID-678
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol#L4


 - [ ] ID-679
	Pragma confirmed better, here is pragma solidity>=0.5.0. [>=0.5.0](contracts/interfaces/uniswap/pool/IUniswapV3PoolImmutables.sol#L2)

contracts/interfaces/uniswap/pool/IUniswapV3PoolImmutables.sol#L2


 - [ ] ID-680
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4


 - [ ] ID-681
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol#L4


 - [ ] ID-682
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol#L4


 - [ ] ID-683
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/MarketRegistry.sol#L2)

contracts/MarketRegistry.sol#L2


 - [ ] ID-684
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#L4)

node_modules/@openzeppelin/contracts/interfaces/draft-IERC1822.sol#L4


 - [ ] ID-685
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721ReceiverUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/token/ERC721/IERC721ReceiverUpgradeable.sol#L4


 - [ ] ID-686
	Pragma confirmed better, here is pragma solidity>=0.5.0. [>=0.5.0](contracts/interfaces/uniswap/IUniswapV3Pool.sol#L2)

contracts/interfaces/uniswap/IUniswapV3Pool.sol#L2


 - [ ] ID-687
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/aave/AavePoolAddressProviderMock.sol#L2)

contracts/mock/aave/AavePoolAddressProviderMock.sol#L2


 - [ ] ID-688
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/EAS/TellerASRegistry.sol#L1)

contracts/EAS/TellerASRegistry.sol#L1


 - [ ] ID-689
	Pragma confirmed better, here is pragma solidity^0.8.9. [^0.8.9](contracts/ERC2771ContextUpgradeable.sol#L4)

contracts/ERC2771ContextUpgradeable.sol#L4


 - [ ] ID-690
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ILoanRepaymentCallbacks.sol#L2)

contracts/interfaces/ILoanRepaymentCallbacks.sol#L2


 - [ ] ID-691
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/Strings.sol#L4)

node_modules/@openzeppelin/contracts/utils/Strings.sol#L4


 - [ ] ID-692
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/libraries/NumbersLib.sol#L2)

contracts/libraries/NumbersLib.sol#L2


 - [ ] ID-693
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TellerV2.sol#L1)

contracts/TellerV2.sol#L1


 - [ ] ID-694
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/Types.sol#L1)

contracts/Types.sol#L1


 - [ ] ID-695
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/ISmartCommitment.sol#L2)

contracts/interfaces/ISmartCommitment.sol#L2


 - [ ] ID-696
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/math/MathUpgradeable.sol#L4


 - [ ] ID-697
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L2)

contracts/LenderCommitmentForwarder/extensions/CommitmentRolloverLoan.sol#L2


 - [ ] ID-698
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#L4)

node_modules/@openzeppelin/contracts/proxy/beacon/IBeacon.sol#L4


 - [ ] ID-699
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](contracts/interfaces/uniswap/IUniswapV3Factory.sol#L2)

contracts/interfaces/uniswap/IUniswapV3Factory.sol#L2


 - [ ] ID-700
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L1)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroupShares.sol#L1


 - [ ] ID-701
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/ExtensionsContextUpgradeable.sol#L2)

contracts/LenderCommitmentForwarder/extensions/ExtensionsContextUpgradeable.sol#L2


 - [ ] ID-702
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/structs/EnumerableSet.sol#L5)

node_modules/@openzeppelin/contracts/utils/structs/EnumerableSet.sol#L5


 - [ ] ID-703
	Pragma confirmed better, here is pragma solidity^0.8.2. [^0.8.2](node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L4)

node_modules/@openzeppelin/contracts/proxy/ERC1967/ERC1967Upgrade.sol#L4


 - [ ] ID-704
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#L4)

node_modules/@openzeppelin/contracts/utils/introspection/IERC165.sol#L4


 - [ ] ID-705
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TellerV2MarketForwarder_G1.sol#L1)

contracts/TellerV2MarketForwarder_G1.sol#L1


 - [ ] ID-706
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/ITellerV2Autopay.sol#L2)

contracts/interfaces/ITellerV2Autopay.sol#L2


 - [ ] ID-707
	Pragma confirmed better, here is pragma solidity>=0.6.0<0.9.0. [>=0.6.0<0.9.0](contracts/libraries/DateTimeLib.sol#L2)

contracts/libraries/DateTimeLib.sol#L2


 - [ ] ID-708
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#L2)

contracts/LenderCommitmentForwarder/SmartCommitmentForwarder.sol#L2


 - [ ] ID-709
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/LenderManagerMock.sol#L1)

contracts/mock/LenderManagerMock.sol#L1


 - [ ] ID-710
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/aave/IFlashLoanSimpleReceiver.sol#L2)

contracts/interfaces/aave/IFlashLoanSimpleReceiver.sol#L2


 - [ ] ID-711
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/IFlashRolloverLoan_G4.sol#L1)

contracts/interfaces/IFlashRolloverLoan_G4.sol#L1


 - [ ] ID-712
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4)

node_modules/@openzeppelin/contracts/utils/math/SafeMath.sol#L4


 - [ ] ID-713
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/utils/math/Math.sol#L4)

node_modules/@openzeppelin/contracts/utils/math/Math.sol#L4


 - [ ] ID-714
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/interfaces/aave/IPool.sol#L2)

contracts/interfaces/aave/IPool.sol#L2


 - [ ] ID-715
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L4)

node_modules/@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol#L4


 - [ ] ID-716
. analyzed (171 contracts with 86 detectors), 742 result(s) found
INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/2024-04-teller-finance_scaned_output/mwe-output.json generate success.
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TellerV2MarketForwarder_G2.sol#L1)

contracts/TellerV2MarketForwarder_G2.sol#L1


 - [ ] ID-717
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/access/AccessControl.sol#L4)

node_modules/@openzeppelin/contracts/access/AccessControl.sol#L4


 - [ ] ID-718
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#L4)

node_modules/@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol#L4


 - [ ] ID-719
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/TellerV2Context.sol#L1)

contracts/TellerV2Context.sol#L1


 - [ ] ID-720
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/mock/TellerASMock.sol#L2)

contracts/mock/TellerASMock.sol#L2


 - [ ] ID-721
	Pragma confirmed better, here is pragma solidity>=0.4.0. [>=0.4.0](contracts/libraries/uniswap/FixedPoint96.sol#L2)

contracts/libraries/uniswap/FixedPoint96.sol#L2


 - [ ] ID-722
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L2)

contracts/LenderCommitmentForwarder/extensions/LenderCommitmentGroup/LenderCommitmentGroup_Smart.sol#L2


 - [ ] ID-723
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/ProtocolFee.sol#L1)

contracts/ProtocolFee.sol#L1


 - [ ] ID-724
	Pragma confirmed better, here is pragma solidity>=0.8.0<0.9.0. [>=0.8.0<0.9.0](contracts/interfaces/ITellerV2MarketForwarder.sol#L2)

contracts/interfaces/ITellerV2MarketForwarder.sol#L2


 - [ ] ID-725
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L2)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G5.sol#L2


 - [ ] ID-726
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L2)

contracts/LenderCommitmentForwarder/extensions/FlashRolloverLoan_G3.sol#L2


 - [ ] ID-727
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


## state-should-be-constant
Impact: Optimization
Confidence: High
 - [ ] ID-728
[TellerV2SolMock.amountOwedMockInterest](contracts/mock/TellerV2SolMock.sol#L20) should be constant

contracts/mock/TellerV2SolMock.sol#L20


 - [ ] ID-729
[WethMock.symbol](contracts/mock/WethMock.sol#L26) should be constant

contracts/mock/WethMock.sol#L26


 - [ ] ID-730
[LenderCommitmentForwarderMock.submitBidWithCollateralWasCalled](contracts/mock/LenderCommitmentForwarderMock.sol#L30) should be constant

contracts/mock/LenderCommitmentForwarderMock.sol#L30


 - [ ] ID-731
[LenderCommitmentForwarderMock.commitmentCount](contracts/mock/LenderCommitmentForwarderMock.sol#L28) should be constant

contracts/mock/LenderCommitmentForwarderMock.sol#L28


 - [ ] ID-732
[TellerV2Storage_G0.version](contracts/TellerV2Storage.sol#L126) should be constant

contracts/TellerV2Storage.sol#L126


 - [ ] ID-733
[TellerV2SolMock.globalBidPaymentCycleDuration](contracts/mock/TellerV2SolMock.sol#L25) should be constant

contracts/mock/TellerV2SolMock.sol#L25


 - [ ] ID-734
[TellerV2Storage_G0.__totalVolumeFilled](contracts/TellerV2Storage.sol#L104) should be constant

contracts/TellerV2Storage.sol#L104


 - [ ] ID-735
[WethMock.name](contracts/mock/WethMock.sol#L25) should be constant

contracts/mock/WethMock.sol#L25


 - [ ] ID-736
[MarketRegistry.version](contracts/MarketRegistry.sol#L59) should be constant

contracts/MarketRegistry.sol#L59


 - [ ] ID-737
[WethMock.decimals](contracts/mock/WethMock.sol#L27) should be constant

contracts/mock/WethMock.sol#L27


 - [ ] ID-738
[LenderCommitmentForwarderMock.acceptBidWasCalled](contracts/mock/LenderCommitmentForwarderMock.sol#L31) should be constant

contracts/mock/LenderCommitmentForwarderMock.sol#L31


 - [ ] ID-739
[TellerV2SolMock.amountOwedMockPrincipal](contracts/mock/TellerV2SolMock.sol#L19) should be constant

contracts/mock/TellerV2SolMock.sol#L19


 - [ ] ID-740
[LenderCommitmentForwarderMock.submitBidWasCalled](contracts/mock/LenderCommitmentForwarderMock.sol#L32) should be constant

contracts/mock/LenderCommitmentForwarderMock.sol#L32


 - [ ] ID-741
[ERC20Permit._PERMIT_TYPEHASH_DEPRECATED_SLOT](node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L37) should be constant

node_modules/@openzeppelin/contracts/token/ERC20/extensions/draft-ERC20Permit.sol#L37


Execution time for Falcon: 42.726633399s
