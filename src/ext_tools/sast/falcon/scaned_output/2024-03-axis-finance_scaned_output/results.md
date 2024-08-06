'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-03-axis-finance/moonraker)
'forge config --json' running
'forge build --build-info --skip */test/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-03-axis-finance/moonraker)
Impossible to generate IR for Catalogue.getRouting:
 'TypeAliasTopLevel' object has no attribute 'type'

state variable: Auction.minAuctionDuration (src/modules/Auction.sol#104) not initialized and not written in contract but be used in contract
state variable: Derivative.tokenMetadata (src/modules/Derivative.sol#33) not initialized and not written in contract but be used in contract
Reference:  

EncryptedMarginalPriceAuctionModule._settle(uint96) (src/modules/auctions/EMPAM.sol#747-837) contains a tautology or contradiction:
	- i >= 0 (src/modules/auctions/EMPAM.sol#772)
Reference:  

Call with hardcoded gas amount occur at Function
ECIES._ecMul(Point,uint256) (src/lib/ECIES.sol#120-127)	(success,output) = address(0x07).staticcall{gas: 6000}(abi.encode(p.x,p.y,scalar)) (src/lib/ECIES.sol#121-122)
Reference: https://swcregistry.io/docs/SWC-134

Transfer.transferFrom(ERC20,address,address,uint256,bool).balanceBefore (src/lib/Transfer.sol#93) is a local variable never initialized
EncryptedMarginalPriceAuctionModule._refundBid(uint96,uint64,address).i (src/modules/auctions/EMPAM.sol#295) is a local variable never initialized
EncryptedMarginalPriceAuctionModule._claimBids(uint96,uint64[]).i (src/modules/auctions/EMPAM.sol#378) is a local variable never initialized
EncryptedMarginalPriceAuctionModule._decryptAndSortBids(uint96,uint64).i (src/modules/auctions/EMPAM.sol#474) is a local variable never initialized
Transfer.permit2TransferFrom(ERC20,address,address,address,uint256,Transfer.Permit2Approval,bool).balanceBefore (src/lib/Transfer.sol#117) is a local variable never initialized
EncryptedMarginalPriceAuctionModule._decryptAndSortBids(uint96,uint64).amountOut (src/modules/auctions/EMPAM.sol#479) is a local variable never initialized
Transfer.transfer(ERC20,address,uint256,bool).balanceBefore (src/lib/Transfer.sol#55) is a local variable never initialized
Reference:  

AuctionHouse._sendPayout(address,uint256,Auctioneer.Routing,bytes) (src/AuctionHouse.sol#800-831)have ignores return value in module.mint(recipient_,address(baseToken),routingParams_.derivativeParams,payoutAmount_,routingParams_.wrapDerivative) (src/AuctionHouse.sol#823-829)
BlastAuctionHouse.constructor(address,address,address) (src/blast/BlastAuctionHouse.sol#48-59)have ignores return value in _WETH.configure(YieldMode.CLAIMABLE) (src/blast/BlastAuctionHouse.sol#54)
BlastAuctionHouse.claimYieldAndGas() (src/blast/BlastAuctionHouse.sol#63-70)have ignores return value in _WETH.claim(_protocol,_WETH.getClaimableAmount(address(this))) (src/blast/BlastAuctionHouse.sol#65)
BlastAuctionHouse.claimModuleGas(Veecode) (src/blast/BlastAuctionHouse.sol#72-75)have ignores return value in _BLAST.claimMaxGas(address(_getModuleIfInstalled(reference_)),_protocol) (src/blast/BlastAuctionHouse.sol#74)
BaselinePreAsset._onClaimProceeds(uint96,uint96,uint96,bytes) (src/callbacks/liquidity/BaselinePreAsset.sol#187-240)have ignores return value in baselineFactory.deploy(address(this),name,symbol,address(reserve),salt,feeRecipient,initTick,initFloor,initDisc) (src/callbacks/liquidity/BaselinePreAsset.sol#226-236)
Reference:  

function:Module.TYPE() (src/modules/Modules.sol#381)is empty 
function:Module.VEECODE() (src/modules/Modules.sol#385)is empty 
function:Module.INIT() (src/modules/Modules.sol#391)is empty 
function:Auction.payoutFor(uint96,uint96) (src/modules/Auction.sol#235)is empty 
function:Auction.priceFor(uint96,uint96) (src/modules/Auction.sol#237)is empty 
function:Auction.maxPayout(uint96) (src/modules/Auction.sol#239)is empty 
function:Auction.maxAmountAccepted(uint96) (src/modules/Auction.sol#241)is empty 
function:MerkleAllowlist.__onPurchase(uint96,address,uint96,uint96,bool,bytes) (src/callbacks/allowlists/MerkleAllowlist.sol#76-83)is empty 
function:MerkleAllowlist.__onBid(uint96,uint64,address,uint96,bytes) (src/callbacks/allowlists/MerkleAllowlist.sol#99-105)is empty 
function:DirectToLiquidity._onCreate(uint96,address,address,address,uint96,bool,bytes) (src/callbacks/liquidity/DTL.sol#16-36)is empty 
function:DirectToLiquidity._onCancel(uint96,uint96,bool,bytes) (src/callbacks/liquidity/DTL.sol#38-40)is empty 
function:DirectToLiquidity._onCurate(uint96,uint96,bool,bytes) (src/callbacks/liquidity/DTL.sol#42-44)is empty 
function:DirectToLiquidity._onClaimProceeds(uint96,uint96,uint96,bytes) (src/callbacks/liquidity/DTL.sol#63-75)is empty 
function:FixedPriceAuctionModule._cancelAuction(uint96) (src/modules/auctions/FPAM.sol#113)is empty 
Reference:  

Timestamp.toPaddedString(uint48).num1 (src/lib/Timestamp.sol#21) is written in both
	num1 = num1 - (1461 * _year) / 4 + 31 (src/lib/Timestamp.sol#25)
	num1 = _month / 11 (src/lib/Timestamp.sol#28)
Reference:  

	- AuctionHouse.setFee(Keycode,FeeManager.FeeType,uint48) (src/AuctionHouse.sol#704-717)
	- BaselinePreAsset.setBAsset(address) (src/callbacks/liquidity/BaselinePreAsset.sol#271-280)
Reference:  

Setter function CappedMerkleAllowlist._onCreate(uint96,address,address,address,uint96,bool,bytes) (src/callbacks/allowlists/CappedMerkleAllowlist.sol#37-52) does not emit an event
Setter function CappedMerkleAllowlist._canBuy(uint96,address,uint96) (src/callbacks/allowlists/CappedMerkleAllowlist.sol#76-84) does not emit an event
Setter function TokenAllowlist._onCreate(uint96,address,address,address,uint96,bool,bytes) (src/callbacks/allowlists/TokenAllowlist.sol#50-73) does not emit an event
Setter function BaselinePreAsset._onCreate(uint96,address,address,address,uint96,bool,bytes) (src/callbacks/liquidity/BaselinePreAsset.sol#109-135) does not emit an event
Setter function BaselinePreAsset._onCurate(uint96,uint96,bool,bytes) (src/callbacks/liquidity/BaselinePreAsset.sol#154-168) does not emit an event
Setter function BaselinePreAsset._onClaimProceeds(uint96,uint96,uint96,bytes) (src/callbacks/liquidity/BaselinePreAsset.sol#187-240) does not emit an event
Setter function BaselinePreAsset.setBAsset(address) (src/callbacks/liquidity/BaselinePreAsset.sol#271-280) does not emit an event
Setter function FixedPriceAuctionModule._auction(uint96,Auction.Lot,bytes) (src/modules/auctions/FPAM.sol#78-106) does not emit an event
Setter function SoulboundCloneERC20.mint(address,uint256) (src/modules/derivatives/SoulboundCloneERC20.sol#42-44) does not emit an event
Setter function SoulboundCloneERC20.burn(address,uint256) (src/modules/derivatives/SoulboundCloneERC20.sol#48-50) does not emit an event
Setter function SoulboundCloneERC20.onlyOwner() (src/modules/derivatives/SoulboundCloneERC20.sol#33-36) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

value assignment lack of validation	ERC20.permit(address,address,uint256,uint256,uint8,bytes32,bytes32) (lib/solmate/src/tokens/ERC20.sol#116-160)recoveredAddress = ecrecover(bytes32,uint8,bytes32,bytes32)(keccak256(bytes)(abi.encodePacked(,DOMAIN_SEPARATOR(),keccak256(bytes)(abi.encode(keccak256(bytes)(Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)),owner,spender,value,nonces[owner] ++,deadline)))),v,r,s) (lib/solmate/src/tokens/ERC20.sol#130-152)
Reference: input validation

variable lacks a zero-check on 		- Catalogue.constructor(address) (src/Catalogue.sol#20-22)
variable lacks a zero-check on 		- AuctionHouse.setProtocol(address) (src/AuctionHouse.sol#720-722)
variable lacks a zero-check on 		- Owned.transferOwnership(address) (lib/solmate/src/auth/Owned.sol#39-43)
variable lacks a zero-check on 		- LinearVesting.redeemable(address,uint256) (src/modules/derivatives/LinearVesting.sol#400-438)
variable lacks a zero-check on 		- ERC6909.setOperator(address,bool) (lib/solmate/src/tokens/ERC6909.sol#87-93)
variable lacks a zero-check on 		- BaseCallback.setSeller(address) (src/callbacks/BaseCallback.sol#192-194)
variable lacks a zero-check on 		- ERC20.approve(address,uint256) (lib/solmate/src/tokens/ERC20.sol#68-74)
variable lacks a zero-check on 		- ERC20.transfer(address,uint256) (lib/solmate/src/tokens/ERC20.sol#76-88)
variable lacks a zero-check on 		- ERC20.transferFrom(address,address,uint256) (lib/solmate/src/tokens/ERC20.sol#90-110)
variable lacks a zero-check on 		- ERC20.permit(address,address,uint256,uint256,uint8,bytes32,bytes32) (lib/solmate/src/tokens/ERC20.sol#116-160)
variable lacks a zero-check on 		- Transfer.transfer(ERC20,address,uint256,bool) (src/lib/Transfer.sol#49-68)
variable lacks a zero-check on 		- Transfer.transferFrom(ERC20,address,address,uint256,bool) (src/lib/Transfer.sol#86-106)
variable lacks a zero-check on 		- Transfer.permit2TransferFrom(ERC20,address,address,address,uint256,Transfer.Permit2Approval,bool) (src/lib/Transfer.sol#108-140)
variable lacks a zero-check on 		- CloneERC20.increaseAllowance(address,uint256) (src/lib/clones/CloneERC20.sol#57-63)
variable lacks a zero-check on 		- CloneERC20.decreaseAllowance(address,uint256) (src/lib/clones/CloneERC20.sol#65-71)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

AuctionHouse.curate(uint96,bytes) (src/AuctionHouse.sol#634-699) compares to a boolean constant:
	-feeData.curated || module.hasEnded(lotId_) == true (src/AuctionHouse.sol#646)
Auctioneer.auction(Auctioneer.RoutingParams,Auction.AuctionParams,string) (src/bases/Auctioneer.sol#160-284) compares to a boolean constant:
	-requiresPrefunding == true (src/bases/Auctioneer.sol#252)
FeeManager._calculatePayoutFees(bool,uint48,uint96) (src/bases/FeeManager.sol#91-101) compares to a boolean constant:
	-curated_ == false (src/bases/FeeManager.sol#97)
Transfer.transfer(ERC20,address,uint256,bool) (src/lib/Transfer.sol#49-68) compares to a boolean constant:
	-validateBalance_ == true (src/lib/Transfer.sol#56)
Transfer.transferFrom(ERC20,address,address,uint256,bool) (src/lib/Transfer.sol#86-106) compares to a boolean constant:
	-validateBalance_ == true (src/lib/Transfer.sol#94)
Transfer.permit2TransferFrom(ERC20,address,address,address,uint256,Transfer.Permit2Approval,bool) (src/lib/Transfer.sol#108-140) compares to a boolean constant:
	-validateBalance_ == true (src/lib/Transfer.sol#118)
LinearVesting.deploy(address,bytes,bool) (src/modules/derivatives/LinearVesting.sol#160-178) compares to a boolean constant:
	-_validate(underlyingToken_,params) == false (src/modules/derivatives/LinearVesting.sol#169)
LinearVesting.mint(address,address,bytes,uint256,bool) (src/modules/derivatives/LinearVesting.sol#243-275) compares to a boolean constant:
	-_validate(underlyingToken_,params) == false (src/modules/derivatives/LinearVesting.sol#262)
LinearVesting._deployIfNeeded(address,LinearVesting.VestingParams,bool) (src/modules/derivatives/LinearVesting.sol#639-674) compares to a boolean constant:
	-token.exists == false (src/modules/derivatives/LinearVesting.sol#650)
LinearVesting.onlyValidTokenId(uint256) (src/modules/derivatives/LinearVesting.sol#107-110) compares to a boolean constant:
	-tokenMetadata[tokenId_].exists == false (src/modules/derivatives/LinearVesting.sol#108)
Reference:  

	- Owned.transferOwnership(address) (lib/solmate/src/auth/Owned.sol#39-43)
	- WithModules.installModule(Module) (src/modules/Modules.sol#160-186)
	- WithModules.sunsetModule(Keycode) (src/modules/Modules.sol#199-211)
	- WithModules.execOnModule(Veecode,bytes) (src/modules/Modules.sol#299-315)
	- AuctionHouse.setProtocol(address) (src/AuctionHouse.sol#720-722)
	- AuctionModule.auction(uint96,Auction.AuctionParams,uint8,uint8) (src/modules/Auction.sol#298-330)
	- AuctionModule.cancelAuction(uint96) (src/modules/Auction.sol#351-364)
	- AuctionModule.purchase(uint96,uint96,bytes) (src/modules/Auction.sol#388-413)
	- AuctionModule.settle(uint96) (src/modules/Auction.sol#589-612)
	- BaseCallback.onCreate(uint96,address,address,address,uint96,bool,bytes) (src/callbacks/BaseCallback.sol#51-74)
	- BaseCallback.setSeller(address) (src/callbacks/BaseCallback.sol#192-194)
Reference:  

	- AuctionModule.bid(uint96,address,address,uint96,bytes) (src/modules/Auction.sol#446-461)
	- AuctionModule.refundBid(uint96,uint64,address) (src/modules/Auction.sol#501-516)
	- AuctionModule.claimBids(uint96,uint64[]) (src/modules/Auction.sol#545-560)
	- AuctionModule.claimProceeds(uint96) (src/modules/Auction.sol#640-654)
	- Module.INIT() (src/modules/Modules.sol#391)
	- LinearVesting.mint(address,uint256,uint256,bool) (src/modules/derivatives/LinearVesting.sol#293-311)
	- LinearVesting.redeemMax(uint256) (src/modules/derivatives/LinearVesting.sol#355-364)
	- LinearVesting.redeem(uint256,uint256) (src/modules/derivatives/LinearVesting.sol#371-385)
	- LinearVesting.wrap(uint256,uint256) (src/modules/derivatives/LinearVesting.sol#462-481)
	- LinearVesting.unwrap(uint256,uint256) (src/modules/derivatives/LinearVesting.sol#488-506)
	- BaseCallback.onCancel(uint96,uint96,bool,bytes) (src/callbacks/BaseCallback.sol#86-97)
	- BaseCallback.onCurate(uint96,uint96,bool,bytes) (src/callbacks/BaseCallback.sol#106-117)
	- BaseCallback.onPurchase(uint96,address,uint96,uint96,bool,bytes) (src/callbacks/BaseCallback.sol#126-139)
	- BaseCallback.onBid(uint96,uint64,address,uint96,bytes) (src/callbacks/BaseCallback.sol#150-161)
	- BaseCallback.onClaimProceeds(uint96,uint96,uint96,bytes) (src/callbacks/BaseCallback.sol#171-181)
	- SoulboundCloneERC20.mint(address,uint256) (src/modules/derivatives/SoulboundCloneERC20.sol#42-44)
	- SoulboundCloneERC20.burn(address,uint256) (src/modules/derivatives/SoulboundCloneERC20.sol#48-50)
Reference:  

AuctionModule._revertIfLotStarted(uint96) (src/modules/Auction.sol#727-729) is never used and should be removed
WithModules._getModuleIfInstalled(Keycode,uint8) (src/modules/Modules.sol#251-268) is never used and should be removed
Reference:  

Potential price manipulation risk:
	- In function redeemMax
		-- redeemableAmount = redeemable(msg.sender,tokenId_) (src/modules/derivatives/LinearVesting.sol#357) have potential price manipulated risk from redeemableAmount and call None which could influence variable:redeemableAmount
	- In function redeemable
		-- wrappedBalance = SoulboundCloneERC20(token.wrapped).balanceOf(owner_) (src/modules/derivatives/LinearVesting.sol#413-414) have potential price manipulated risk from wrappedBalance and call None which could influence variable:vested
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

Potential DoS Gas Limit Attack occur inAuctionHouse.claimBids(uint96,uint64[]) (src/AuctionHouse.sol#403-448)BEGIN_LOOP (src/AuctionHouse.sol#420-447)
Potential DoS Gas Limit Attack occur inEncryptedMarginalPriceAuctionModule._refundBid(uint96,uint64,address) (src/modules/auctions/EMPAM.sol#284-305)BEGIN_LOOP (src/modules/auctions/EMPAM.sol#295-301)
Potential DoS Gas Limit Attack occur inEncryptedMarginalPriceAuctionModule._claimBids(uint96,uint64[]) (src/modules/auctions/EMPAM.sol#372-387)BEGIN_LOOP (src/modules/auctions/EMPAM.sol#378-384)
Potential DoS Gas Limit Attack occur inEncryptedMarginalPriceAuctionModule._decryptAndSortBids(uint96,uint64) (src/modules/auctions/EMPAM.sol#459-522)BEGIN_LOOP (src/modules/auctions/EMPAM.sol#474-513)
Potential DoS Gas Limit Attack occur inEncryptedMarginalPriceAuctionModule._getLotMarginalPrice(uint96) (src/modules/auctions/EMPAM.sol#595-728)BEGIN_LOOP (src/modules/auctions/EMPAM.sol#611-724)
Potential DoS Gas Limit Attack occur inEncryptedMarginalPriceAuctionModule._settle(uint96) (src/modules/auctions/EMPAM.sol#747-837)BEGIN_LOOP (src/modules/auctions/EMPAM.sol#772-781)
Potential DoS Gas Limit Attack occur inMaxPriorityQueue._swim(Queue,uint64) (src/lib/MaxPriorityQueue.sol#52-57)BEGIN_LOOP (src/lib/MaxPriorityQueue.sol#53-56)
Potential DoS Gas Limit Attack occur inMaxPriorityQueue._sink(Queue,uint64) (src/lib/MaxPriorityQueue.sol#60-72)BEGIN_LOOP (src/lib/MaxPriorityQueue.sol#61-71)
Potential DoS Gas Limit Attack occur inClonesWithImmutableArgs.clone(address,bytes) (src/lib/clones/ClonesWithImmutableArgs.sol#25-152)BEGIN_LOOP (src/lib/clones/ClonesWithImmutableArgs.sol#125-133)
Potential DoS Gas Limit Attack occur inClonesWithImmutableArgs.clone3(address,bytes,bytes32) (src/lib/clones/ClonesWithImmutableArgs.sol#160-338)BEGIN_LOOP (src/lib/clones/ClonesWithImmutableArgs.sol#267-275)
Reference: https://swcregistry.io/docs/SWC-128

CondenserModule (src/modules/Condenser.sol#13) does not implement functions:
	- Condenser.condense(bytes,bytes) (src/modules/Condenser.sol#7-10)
Reference:  

function:Owned.transferOwnership(address) (lib/solmate/src/auth/Owned.sol#39-43)is public and can be replaced with external 
function:ERC20.approve(address,uint256) (lib/solmate/src/tokens/ERC20.sol#68-74)is public and can be replaced with external 
function:ERC20.transfer(address,uint256) (lib/solmate/src/tokens/ERC20.sol#76-88)is public and can be replaced with external 
function:ERC20.transferFrom(address,address,uint256) (lib/solmate/src/tokens/ERC20.sol#90-110)is public and can be replaced with external 
function:ERC20.permit(address,address,uint256,uint256,uint8,bytes32,bytes32) (lib/solmate/src/tokens/ERC20.sol#116-160)is public and can be replaced with external 
function:ERC6909.transfer(address,uint256,uint256) (lib/solmate/src/tokens/ERC6909.sol#33-49)is public and can be replaced with external 
function:ERC6909.transferFrom(address,address,uint256,uint256) (lib/solmate/src/tokens/ERC6909.sol#51-73)is public and can be replaced with external 
function:ERC6909.approve(address,uint256,uint256) (lib/solmate/src/tokens/ERC6909.sol#75-85)is public and can be replaced with external 
function:ERC6909.setOperator(address,bool) (lib/solmate/src/tokens/ERC6909.sol#87-93)is public and can be replaced with external 
function:ERC6909.supportsInterface(bytes4) (lib/solmate/src/tokens/ERC6909.sol#99-103)is public and can be replaced with external 
function:FeeManager.calculateQuoteFees(uint96,uint96,bool,uint96) (src/bases/FeeManager.sol#68-88)is public and can be replaced with external 
function:EncryptedMarginalPriceAuctionModule.VEECODE() (src/modules/auctions/EMPAM.sol#124-126)is public and can be replaced with external 
function:EncryptedMarginalPriceAuctionModule.TYPE() (src/modules/auctions/EMPAM.sol#128-130)is public and can be replaced with external 
function:AuctionModule.hasEnded(uint96) (src/modules/Auction.sol#687-689)is public and can be replaced with external 
function:Module.TYPE() (src/modules/Modules.sol#381)is public and can be replaced with external 
function:Module.VEECODE() (src/modules/Modules.sol#385)is public and can be replaced with external 
function:Auction.payoutFor(uint96,uint96) (src/modules/Auction.sol#235)is public and can be replaced with external 
function:Auction.priceFor(uint96,uint96) (src/modules/Auction.sol#237)is public and can be replaced with external 
function:Auction.maxPayout(uint96) (src/modules/Auction.sol#239)is public and can be replaced with external 
function:Auction.maxAmountAccepted(uint96) (src/modules/Auction.sol#241)is public and can be replaced with external 
function:Auction.isLive(uint96) (src/modules/Auction.sol#250)is public and can be replaced with external 
function:Auction.hasEnded(uint96) (src/modules/Auction.sol#259)is public and can be replaced with external 
function:LinearVesting.TYPE() (src/modules/derivatives/LinearVesting.sol#96-98)is public and can be replaced with external 
function:LinearVesting.transfer(address,uint256,uint256) (src/modules/derivatives/LinearVesting.sol#123-125)is public and can be replaced with external 
function:LinearVesting.transferFrom(address,address,uint256,uint256) (src/modules/derivatives/LinearVesting.sol#129-136)is public and can be replaced with external 
function:LinearVesting.approve(address,uint256,uint256) (src/modules/derivatives/LinearVesting.sol#140-142)is public and can be replaced with external 
function:LinearVesting.validate(address,bytes) (src/modules/derivatives/LinearVesting.sol#546-554)is public and can be replaced with external 
function:LinearVesting.name(uint256) (src/modules/derivatives/LinearVesting.sol#719-732)is public and can be replaced with external 
function:LinearVesting.symbol(uint256) (src/modules/derivatives/LinearVesting.sol#737-750)is public and can be replaced with external 
function:LinearVesting.decimals(uint256) (src/modules/derivatives/LinearVesting.sol#755-767)is public and can be replaced with external 
function:ERC6909Metadata.name(uint256) (src/lib/ERC6909Metadata.sol#9)is public and can be replaced with external 
function:ERC6909Metadata.symbol(uint256) (src/lib/ERC6909Metadata.sol#15)is public and can be replaced with external 
function:ERC6909Metadata.decimals(uint256) (src/lib/ERC6909Metadata.sol#21)is public and can be replaced with external 
function:ECIES.decrypt(uint256,Point,uint256,uint256) (src/lib/ECIES.sol#59-74)is public and can be replaced with external 
function:ECIES.encrypt(uint256,Point,uint256,uint256) (src/lib/ECIES.sol#83-102)is public and can be replaced with external 
function:ECIES.isValid(Point) (src/lib/ECIES.sol#137-139)is public and can be replaced with external 
function:MaxPriorityQueue.initialize(Queue) (src/lib/MaxPriorityQueue.sol#27-29)is public and can be replaced with external 
function:MaxPriorityQueue.getNumBids(Queue) (src/lib/MaxPriorityQueue.sol#35-37)is public and can be replaced with external 
function:MaxPriorityQueue.getMax(Queue) (src/lib/MaxPriorityQueue.sol#40-44)is public and can be replaced with external 
function:MaxPriorityQueue.getMaxId(Queue) (src/lib/MaxPriorityQueue.sol#46-49)is public and can be replaced with external 
function:MaxPriorityQueue.insert(Queue,uint64,uint96,uint96) (src/lib/MaxPriorityQueue.sol#75-85)is public and can be replaced with external 
function:MaxPriorityQueue.delMax(Queue) (src/lib/MaxPriorityQueue.sol#96-105)is public and can be replaced with external 
function:Timestamp.toPaddedString(uint48) (src/lib/Timestamp.sol#7-44)is public and can be replaced with external 
function:Transfer.approve(ERC20,address,uint256) (src/lib/Transfer.sol#30-32)is public and can be replaced with external 
function:Transfer.transfer(ERC20,address,uint256,bool) (src/lib/Transfer.sol#49-68)is public and can be replaced with external 
function:Transfer.permit2OrTransferFrom(ERC20,address,address,address,uint256,Transfer.Permit2Approval,bool) (src/lib/Transfer.sol#142-161)is public and can be replaced with external 
function:Transfer.decodePermit2Approval(bytes) (src/lib/Transfer.sol#163-177)is public and can be replaced with external 
function:CloneERC20.approve(address,uint256) (src/lib/clones/CloneERC20.sol#49-55)is public and can be replaced with external 
function:CloneERC20.increaseAllowance(address,uint256) (src/lib/clones/CloneERC20.sol#57-63)is public and can be replaced with external 
function:CloneERC20.decreaseAllowance(address,uint256) (src/lib/clones/CloneERC20.sol#65-71)is public and can be replaced with external 
function:CloneERC20.transfer(address,uint256) (src/lib/clones/CloneERC20.sol#73-85)is public and can be replaced with external 
function:CloneERC20.transferFrom(address,address,uint256) (src/lib/clones/CloneERC20.sol#87-103)is public and can be replaced with external 
function:FixedPriceAuctionModule.VEECODE() (src/modules/auctions/FPAM.sol#53-55)is public and can be replaced with external 
function:FixedPriceAuctionModule.TYPE() (src/modules/auctions/FPAM.sol#58-60)is public and can be replaced with external 
function:SoulboundCloneERC20.transfer(address,uint256) (src/modules/derivatives/SoulboundCloneERC20.sol#54-56)is public and can be replaced with external 
function:SoulboundCloneERC20.transferFrom(address,address,uint256) (src/modules/derivatives/SoulboundCloneERC20.sol#58-60)is public and can be replaced with external 
function:SoulboundCloneERC20.approve(address,uint256) (src/modules/derivatives/SoulboundCloneERC20.sol#62-64)is public and can be replaced with external 
function:SoulboundCloneERC20.owner() (src/modules/derivatives/SoulboundCloneERC20.sol#78-80)is public and can be replaced with external 
Reference:  

	Pragma confirmed better, here is pragma solidity^0.8.4. ^0.8.4 (lib/solady/src/utils/MerkleProofLib.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (lib/solmate/src/auth/Owned.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (lib/solmate/src/tokens/ERC20.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (lib/solmate/src/tokens/ERC6909.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (lib/solmate/src/utils/FixedPointMathLib.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (lib/solmate/src/utils/ReentrancyGuard.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (lib/solmate/src/utils/SafeTransferLib.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.19. ^0.8.19 (src/callbacks/liquidity/DTL.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (src/interfaces/ICallback.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.19. ^0.8.19 (src/lib/Callbacks.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (src/lib/ECIES.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (src/lib/MaxPriorityQueue.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8. ^0.8 (src/lib/Uint2Str.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.4. ^0.8.4 (src/lib/clones/Clone.sol#2)
	Pragma confirmed better, here is pragma solidity>=0.8.0. >=0.8.0 (src/lib/clones/CloneERC20.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.4. ^0.8.4 (src/lib/clones/ClonesWithImmutableArgs.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.19. ^0.8.19 (src/lib/permit2/interfaces/IPermit2.sol#2)
Reference:  

unnecessary reentrancy lock found inBlastAuctionHouse
	- FeeManager.claimRewards(address) (src/bases/FeeManager.sol#123-129)
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
An error occurred: 'TypeAliasTopLevel' object has no attribute 'type'
An error occurred: maximum recursion depth exceeded while pickling an object
after for recrusion
Summary
 - [state-variable-not-initialized](#state-variable-not-initialized) (2 results) (High)
 - [constant-result](#constant-result) (1 results) (Medium)
 - [hardcode-gas-amount](#hardcode-gas-amount) (1 results) (Medium)
 - [uninitialized-local](#uninitialized-local) (7 results) (Medium)
 - [unused-return](#unused-return) (5 results) (Medium)
 - [void-function](#void-function) (14 results) (Medium)
 - [useless-write](#useless-write) (1 results) (Medium)
 - [centralized-risk-low](#centralized-risk-low) (2 results) (Low)
 - [pess-event-setter](#pess-event-setter) (11 results) (Low)
 - [input-validation](#input-validation) (1 results) (Low)
 - [missing-zero-check](#missing-zero-check) (15 results) (Low)
 - [unnecessary-boolean-compare](#unnecessary-boolean-compare) (10 results) (Informational)
 - [centralized-risk-informational](#centralized-risk-informational) (11 results) (Informational)
 - [centralized-risk-other](#centralized-risk-other) (17 results) (Informational)
 - [dead-function](#dead-function) (2 results) (Informational)
 - [price-manipulation-info](#price-manipulation-info) (1 results) (Informational)
 - [uncontrolled-resource-consumption](#uncontrolled-resource-consumption) (10 results) (Informational)
 - [no-derived-function](#no-derived-function) (1 results) (Informational)
 - [unnecessary-public-function-modifier](#unnecessary-public-function-modifier) (58 results) (Informational)
 - [version-only](#version-only) (17 results) (Informational)
 - [unnecessary-reentrancy-lock](#unnecessary-reentrancy-lock) (1 results) (Optimization)
## state-variable-not-initialized
Impact: High
Confidence: High
 - [ ] ID-0
state variable: [Derivative.tokenMetadata](src/modules/Derivative.sol#L33) not initialized and not written in contract but be used in contract

src/modules/Derivative.sol#L33


 - [ ] ID-1
state variable: [Auction.minAuctionDuration](src/modules/Auction.sol#L104) not initialized and not written in contract but be used in contract

src/modules/Auction.sol#L104


## constant-result
Impact: Medium
Confidence: High
 - [ ] ID-2
[EncryptedMarginalPriceAuctionModule._settle(uint96)](src/modules/auctions/EMPAM.sol#L747-L837) contains a tautology or contradiction:
	- [i >= 0](src/modules/auctions/EMPAM.sol#L772)

src/modules/auctions/EMPAM.sol#L747-L837


## hardcode-gas-amount
Impact: Medium
Confidence: High
 - [ ] ID-3
Call with hardcoded gas amount occur at Function
[ECIES._ecMul(Point,uint256)](src/lib/ECIES.sol#L120-L127)	[(success,output) = address(0x07).staticcall{gas: 6000}(abi.encode(p.x,p.y,scalar))](src/lib/ECIES.sol#L121-L122)

src/lib/ECIES.sol#L120-L127


## uninitialized-local
Impact: Medium
Confidence: Medium
 - [ ] ID-4
[EncryptedMarginalPriceAuctionModule._decryptAndSortBids(uint96,uint64).i](src/modules/auctions/EMPAM.sol#L474) is a local variable never initialized

src/modules/auctions/EMPAM.sol#L474


 - [ ] ID-5
[EncryptedMarginalPriceAuctionModule._claimBids(uint96,uint64[]).i](src/modules/auctions/EMPAM.sol#L378) is a local variable never initialized

src/modules/auctions/EMPAM.sol#L378


 - [ ] ID-6
[EncryptedMarginalPriceAuctionModule._refundBid(uint96,uint64,address).i](src/modules/auctions/EMPAM.sol#L295) is a local variable never initialized

src/modules/auctions/EMPAM.sol#L295


 - [ ] ID-7
[Transfer.transferFrom(ERC20,address,address,uint256,bool).balanceBefore](src/lib/Transfer.sol#L93) is a local variable never initialized

src/lib/Transfer.sol#L93


 - [ ] ID-8
[EncryptedMarginalPriceAuctionModule._decryptAndSortBids(uint96,uint64).amountOut](src/modules/auctions/EMPAM.sol#L479) is a local variable never initialized

src/modules/auctions/EMPAM.sol#L479


 - [ ] ID-9
[Transfer.permit2TransferFrom(ERC20,address,address,address,uint256,Transfer.Permit2Approval,bool).balanceBefore](src/lib/Transfer.sol#L117) is a local variable never initialized

src/lib/Transfer.sol#L117


 - [ ] ID-10
[Transfer.transfer(ERC20,address,uint256,bool).balanceBefore](src/lib/Transfer.sol#L55) is a local variable never initialized

src/lib/Transfer.sol#L55


## unused-return
Impact: Medium
Confidence: Medium
 - [ ] ID-11
[BaselinePreAsset._onClaimProceeds(uint96,uint96,uint96,bytes)](src/callbacks/liquidity/BaselinePreAsset.sol#L187-L240)have ignores return value in [baselineFactory.deploy(address(this),name,symbol,address(reserve),salt,feeRecipient,initTick,initFloor,initDisc)](src/callbacks/liquidity/BaselinePreAsset.sol#L226-L236)

src/callbacks/liquidity/BaselinePreAsset.sol#L187-L240


 - [ ] ID-12
[AuctionHouse._sendPayout(address,uint256,Auctioneer.Routing,bytes)](src/AuctionHouse.sol#L800-L831)have ignores return value in [module.mint(recipient_,address(baseToken),routingParams_.derivativeParams,payoutAmount_,routingParams_.wrapDerivative)](src/AuctionHouse.sol#L823-L829)

src/AuctionHouse.sol#L800-L831


 - [ ] ID-13
[BlastAuctionHouse.constructor(address,address,address)](src/blast/BlastAuctionHouse.sol#L48-L59)have ignores return value in [_WETH.configure(YieldMode.CLAIMABLE)](src/blast/BlastAuctionHouse.sol#L54)

src/blast/BlastAuctionHouse.sol#L48-L59


 - [ ] ID-14
[BlastAuctionHouse.claimModuleGas(Veecode)](src/blast/BlastAuctionHouse.sol#L72-L75)have ignores return value in [_BLAST.claimMaxGas(address(_getModuleIfInstalled(reference_)),_protocol)](src/blast/BlastAuctionHouse.sol#L74)

src/blast/BlastAuctionHouse.sol#L72-L75


 - [ ] ID-15
[BlastAuctionHouse.claimYieldAndGas()](src/blast/BlastAuctionHouse.sol#L63-L70)have ignores return value in [_WETH.claim(_protocol,_WETH.getClaimableAmount(address(this)))](src/blast/BlastAuctionHouse.sol#L65)

src/blast/BlastAuctionHouse.sol#L63-L70


## void-function
Impact: Medium
Confidence: High
 - [ ] ID-16
function:[MerkleAllowlist.__onPurchase(uint96,address,uint96,uint96,bool,bytes)](src/callbacks/allowlists/MerkleAllowlist.sol#L76-L83)is empty 

src/callbacks/allowlists/MerkleAllowlist.sol#L76-L83


 - [ ] ID-17
function:[Auction.priceFor(uint96,uint96)](src/modules/Auction.sol#L237)is empty 

src/modules/Auction.sol#L237


 - [ ] ID-18
function:[MerkleAllowlist.__onBid(uint96,uint64,address,uint96,bytes)](src/callbacks/allowlists/MerkleAllowlist.sol#L99-L105)is empty 

src/callbacks/allowlists/MerkleAllowlist.sol#L99-L105


 - [ ] ID-19
function:[FixedPriceAuctionModule._cancelAuction(uint96)](src/modules/auctions/FPAM.sol#L113)is empty 

src/modules/auctions/FPAM.sol#L113


 - [ ] ID-20
function:[Auction.maxAmountAccepted(uint96)](src/modules/Auction.sol#L241)is empty 

src/modules/Auction.sol#L241


 - [ ] ID-21
function:[Module.TYPE()](src/modules/Modules.sol#L381)is empty 

src/modules/Modules.sol#L381


 - [ ] ID-22
function:[DirectToLiquidity._onCreate(uint96,address,address,address,uint96,bool,bytes)](src/callbacks/liquidity/DTL.sol#L16-L36)is empty 

src/callbacks/liquidity/DTL.sol#L16-L36


 - [ ] ID-23
function:[Auction.payoutFor(uint96,uint96)](src/modules/Auction.sol#L235)is empty 

src/modules/Auction.sol#L235


 - [ ] ID-24
function:[Auction.maxPayout(uint96)](src/modules/Auction.sol#L239)is empty 

src/modules/Auction.sol#L239


 - [ ] ID-25
function:[Module.INIT()](src/modules/Modules.sol#L391)is empty 

src/modules/Modules.sol#L391


 - [ ] ID-26
function:[DirectToLiquidity._onCancel(uint96,uint96,bool,bytes)](src/callbacks/liquidity/DTL.sol#L38-L40)is empty 

src/callbacks/liquidity/DTL.sol#L38-L40


 - [ ] ID-27
function:[DirectToLiquidity._onCurate(uint96,uint96,bool,bytes)](src/callbacks/liquidity/DTL.sol#L42-L44)is empty 

src/callbacks/liquidity/DTL.sol#L42-L44


 - [ ] ID-28
function:[Module.VEECODE()](src/modules/Modules.sol#L385)is empty 

src/modules/Modules.sol#L385


 - [ ] ID-29
function:[DirectToLiquidity._onClaimProceeds(uint96,uint96,uint96,bytes)](src/callbacks/liquidity/DTL.sol#L63-L75)is empty 

src/callbacks/liquidity/DTL.sol#L63-L75


## useless-write
Impact: Medium
Confidence: High
 - [ ] ID-30
[Timestamp.toPaddedString(uint48).num1](src/lib/Timestamp.sol#L21) is written in both
	[num1 = num1 - (1461 * _year) / 4 + 31](src/lib/Timestamp.sol#L25)
	[num1 = _month / 11](src/lib/Timestamp.sol#L28)

src/lib/Timestamp.sol#L21


## centralized-risk-low
Impact: Low
Confidence: High
 - [ ] ID-31
	- [AuctionHouse.setFee(Keycode,FeeManager.FeeType,uint48)](src/AuctionHouse.sol#L704-L717)

src/AuctionHouse.sol#L704-L717


 - [ ] ID-32
	- [BaselinePreAsset.setBAsset(address)](src/callbacks/liquidity/BaselinePreAsset.sol#L271-L280)

src/callbacks/liquidity/BaselinePreAsset.sol#L271-L280


## pess-event-setter
Impact: Low
Confidence: Medium
 - [ ] ID-33
Setter function [CappedMerkleAllowlist._canBuy(uint96,address,uint96)](src/callbacks/allowlists/CappedMerkleAllowlist.sol#L76-L84) does not emit an event

src/callbacks/allowlists/CappedMerkleAllowlist.sol#L76-L84


 - [ ] ID-34
Setter function [SoulboundCloneERC20.onlyOwner()](src/modules/derivatives/SoulboundCloneERC20.sol#L33-L36) does not emit an event

src/modules/derivatives/SoulboundCloneERC20.sol#L33-L36


 - [ ] ID-35
Setter function [TokenAllowlist._onCreate(uint96,address,address,address,uint96,bool,bytes)](src/callbacks/allowlists/TokenAllowlist.sol#L50-L73) does not emit an event

src/callbacks/allowlists/TokenAllowlist.sol#L50-L73


 - [ ] ID-36
Setter function [SoulboundCloneERC20.mint(address,uint256)](src/modules/derivatives/SoulboundCloneERC20.sol#L42-L44) does not emit an event

src/modules/derivatives/SoulboundCloneERC20.sol#L42-L44


 - [ ] ID-37
Setter function [BaselinePreAsset.setBAsset(address)](src/callbacks/liquidity/BaselinePreAsset.sol#L271-L280) does not emit an event

src/callbacks/liquidity/BaselinePreAsset.sol#L271-L280


 - [ ] ID-38
Setter function [BaselinePreAsset._onCreate(uint96,address,address,address,uint96,bool,bytes)](src/callbacks/liquidity/BaselinePreAsset.sol#L109-L135) does not emit an event

src/callbacks/liquidity/BaselinePreAsset.sol#L109-L135


 - [ ] ID-39
Setter function [BaselinePreAsset._onClaimProceeds(uint96,uint96,uint96,bytes)](src/callbacks/liquidity/BaselinePreAsset.sol#L187-L240) does not emit an event

src/callbacks/liquidity/BaselinePreAsset.sol#L187-L240


 - [ ] ID-40
Setter function [BaselinePreAsset._onCurate(uint96,uint96,bool,bytes)](src/callbacks/liquidity/BaselinePreAsset.sol#L154-L168) does not emit an event

src/callbacks/liquidity/BaselinePreAsset.sol#L154-L168


 - [ ] ID-41
Setter function [FixedPriceAuctionModule._auction(uint96,Auction.Lot,bytes)](src/modules/auctions/FPAM.sol#L78-L106) does not emit an event

src/modules/auctions/FPAM.sol#L78-L106


 - [ ] ID-42
Setter function [CappedMerkleAllowlist._onCreate(uint96,address,address,address,uint96,bool,bytes)](src/callbacks/allowlists/CappedMerkleAllowlist.sol#L37-L52) does not emit an event

src/callbacks/allowlists/CappedMerkleAllowlist.sol#L37-L52


 - [ ] ID-43
Setter function [SoulboundCloneERC20.burn(address,uint256)](src/modules/derivatives/SoulboundCloneERC20.sol#L48-L50) does not emit an event

src/modules/derivatives/SoulboundCloneERC20.sol#L48-L50


## input-validation
Impact: Low
Confidence: Medium
 - [ ] ID-44
value assignment lack of validation	[ERC20.permit(address,address,uint256,uint256,uint8,bytes32,bytes32)](lib/solmate/src/tokens/ERC20.sol#L116-L160)[recoveredAddress = ecrecover(bytes32,uint8,bytes32,bytes32)(keccak256(bytes)(abi.encodePacked(,DOMAIN_SEPARATOR(),keccak256(bytes)(abi.encode(keccak256(bytes)(Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)),owner,spender,value,nonces[owner] ++,deadline)))),v,r,s)](lib/solmate/src/tokens/ERC20.sol#L130-L152)

lib/solmate/src/tokens/ERC20.sol#L116-L160


## missing-zero-check
Impact: Low
Confidence: Medium
 - [ ] ID-45
variable lacks a zero-check on 		- [ERC20.transferFrom(address,address,uint256)](lib/solmate/src/tokens/ERC20.sol#L90-L110)

lib/solmate/src/tokens/ERC20.sol#L90-L110


 - [ ] ID-46
variable lacks a zero-check on 		- [CloneERC20.increaseAllowance(address,uint256)](src/lib/clones/CloneERC20.sol#L57-L63)

src/lib/clones/CloneERC20.sol#L57-L63


 - [ ] ID-47
variable lacks a zero-check on 		- [Transfer.transferFrom(ERC20,address,address,uint256,bool)](src/lib/Transfer.sol#L86-L106)

src/lib/Transfer.sol#L86-L106


 - [ ] ID-48
variable lacks a zero-check on 		- [Transfer.permit2TransferFrom(ERC20,address,address,address,uint256,Transfer.Permit2Approval,bool)](src/lib/Transfer.sol#L108-L140)

src/lib/Transfer.sol#L108-L140


 - [ ] ID-49
variable lacks a zero-check on 		- [CloneERC20.decreaseAllowance(address,uint256)](src/lib/clones/CloneERC20.sol#L65-L71)

src/lib/clones/CloneERC20.sol#L65-L71


 - [ ] ID-50
variable lacks a zero-check on 		- [LinearVesting.redeemable(address,uint256)](src/modules/derivatives/LinearVesting.sol#L400-L438)

src/modules/derivatives/LinearVesting.sol#L400-L438


 - [ ] ID-51
variable lacks a zero-check on 		- [ERC20.transfer(address,uint256)](lib/solmate/src/tokens/ERC20.sol#L76-L88)

lib/solmate/src/tokens/ERC20.sol#L76-L88


 - [ ] ID-52
variable lacks a zero-check on 		- [Catalogue.constructor(address)](src/Catalogue.sol#L20-L22)

src/Catalogue.sol#L20-L22


 - [ ] ID-53
variable lacks a zero-check on 		- [ERC20.permit(address,address,uint256,uint256,uint8,bytes32,bytes32)](lib/solmate/src/tokens/ERC20.sol#L116-L160)

lib/solmate/src/tokens/ERC20.sol#L116-L160


 - [ ] ID-54
variable lacks a zero-check on 		- [BaseCallback.setSeller(address)](src/callbacks/BaseCallback.sol#L192-L194)

src/callbacks/BaseCallback.sol#L192-L194


 - [ ] ID-55
variable lacks a zero-check on 		- [ERC20.approve(address,uint256)](lib/solmate/src/tokens/ERC20.sol#L68-L74)

lib/solmate/src/tokens/ERC20.sol#L68-L74


 - [ ] ID-56
variable lacks a zero-check on 		- [Owned.transferOwnership(address)](lib/solmate/src/auth/Owned.sol#L39-L43)

lib/solmate/src/auth/Owned.sol#L39-L43


 - [ ] ID-57
variable lacks a zero-check on 		- [ERC6909.setOperator(address,bool)](lib/solmate/src/tokens/ERC6909.sol#L87-L93)

lib/solmate/src/tokens/ERC6909.sol#L87-L93


 - [ ] ID-58
variable lacks a zero-check on 		- [Transfer.transfer(ERC20,address,uint256,bool)](src/lib/Transfer.sol#L49-L68)

src/lib/Transfer.sol#L49-L68


 - [ ] ID-59
variable lacks a zero-check on 		- [AuctionHouse.setProtocol(address)](src/AuctionHouse.sol#L720-L722)

src/AuctionHouse.sol#L720-L722


## unnecessary-boolean-compare
Impact: Informational
Confidence: High
 - [ ] ID-60
[LinearVesting.mint(address,address,bytes,uint256,bool)](src/modules/derivatives/LinearVesting.sol#L243-L275) compares to a boolean constant:
	-[_validate(underlyingToken_,params) == false](src/modules/derivatives/LinearVesting.sol#L262)

src/modules/derivatives/LinearVesting.sol#L243-L275


 - [ ] ID-61
[AuctionHouse.curate(uint96,bytes)](src/AuctionHouse.sol#L634-L699) compares to a boolean constant:
	-[feeData.curated || module.hasEnded(lotId_) == true](src/AuctionHouse.sol#L646)

src/AuctionHouse.sol#L634-L699


 - [ ] ID-62
[LinearVesting._deployIfNeeded(address,LinearVesting.VestingParams,bool)](src/modules/derivatives/LinearVesting.sol#L639-L674) compares to a boolean constant:
	-[token.exists == false](src/modules/derivatives/LinearVesting.sol#L650)

src/modules/derivatives/LinearVesting.sol#L639-L674


 - [ ] ID-63
[FeeManager._calculatePayoutFees(bool,uint48,uint96)](src/bases/FeeManager.sol#L91-L101) compares to a boolean constant:
	-[curated_ == false](src/bases/FeeManager.sol#L97)

src/bases/FeeManager.sol#L91-L101


 - [ ] ID-64
[Auctioneer.auction(Auctioneer.RoutingParams,Auction.AuctionParams,string)](src/bases/Auctioneer.sol#L160-L284) compares to a boolean constant:
	-[requiresPrefunding == true](src/bases/Auctioneer.sol#L252)

src/bases/Auctioneer.sol#L160-L284


 - [ ] ID-65
[LinearVesting.deploy(address,bytes,bool)](src/modules/derivatives/LinearVesting.sol#L160-L178) compares to a boolean constant:
	-[_validate(underlyingToken_,params) == false](src/modules/derivatives/LinearVesting.sol#L169)

src/modules/derivatives/LinearVesting.sol#L160-L178


 - [ ] ID-66
[Transfer.transfer(ERC20,address,uint256,bool)](src/lib/Transfer.sol#L49-L68) compares to a boolean constant:
	-[validateBalance_ == true](src/lib/Transfer.sol#L56)

src/lib/Transfer.sol#L49-L68


 - [ ] ID-67
[LinearVesting.onlyValidTokenId(uint256)](src/modules/derivatives/LinearVesting.sol#L107-L110) compares to a boolean constant:
	-[tokenMetadata[tokenId_].exists == false](src/modules/derivatives/LinearVesting.sol#L108)

src/modules/derivatives/LinearVesting.sol#L107-L110


 - [ ] ID-68
[Transfer.permit2TransferFrom(ERC20,address,address,address,uint256,Transfer.Permit2Approval,bool)](src/lib/Transfer.sol#L108-L140) compares to a boolean constant:
	-[validateBalance_ == true](src/lib/Transfer.sol#L118)

src/lib/Transfer.sol#L108-L140


 - [ ] ID-69
[Transfer.transferFrom(ERC20,address,address,uint256,bool)](src/lib/Transfer.sol#L86-L106) compares to a boolean constant:
	-[validateBalance_ == true](src/lib/Transfer.sol#L94)

src/lib/Transfer.sol#L86-L106


## centralized-risk-informational
Impact: Informational
Confidence: High
 - [ ] ID-70
	- [WithModules.execOnModule(Veecode,bytes)](src/modules/Modules.sol#L299-L315)

src/modules/Modules.sol#L299-L315


 - [ ] ID-71
	- [BaseCallback.onCreate(uint96,address,address,address,uint96,bool,bytes)](src/callbacks/BaseCallback.sol#L51-L74)

src/callbacks/BaseCallback.sol#L51-L74


 - [ ] ID-72
	- [BaseCallback.setSeller(address)](src/callbacks/BaseCallback.sol#L192-L194)

src/callbacks/BaseCallback.sol#L192-L194


 - [ ] ID-73
	- [Owned.transferOwnership(address)](lib/solmate/src/auth/Owned.sol#L39-L43)

lib/solmate/src/auth/Owned.sol#L39-L43


 - [ ] ID-74
	- [AuctionModule.cancelAuction(uint96)](src/modules/Auction.sol#L351-L364)

src/modules/Auction.sol#L351-L364


 - [ ] ID-75
	- [AuctionHouse.setProtocol(address)](src/AuctionHouse.sol#L720-L722)

src/AuctionHouse.sol#L720-L722


 - [ ] ID-76
	- [AuctionModule.settle(uint96)](src/modules/Auction.sol#L589-L612)

src/modules/Auction.sol#L589-L612


 - [ ] ID-77
	- [AuctionModule.purchase(uint96,uint96,bytes)](src/modules/Auction.sol#L388-L413)

src/modules/Auction.sol#L388-L413


 - [ ] ID-78
	- [AuctionModule.auction(uint96,Auction.AuctionParams,uint8,uint8)](src/modules/Auction.sol#L298-L330)

src/modules/Auction.sol#L298-L330


 - [ ] ID-79
	- [WithModules.sunsetModule(Keycode)](src/modules/Modules.sol#L199-L211)

src/modules/Modules.sol#L199-L211


 - [ ] ID-80
	- [WithModules.installModule(Module)](src/modules/Modules.sol#L160-L186)

src/modules/Modules.sol#L160-L186


## centralized-risk-other
Impact: Informational
Confidence: High
 - [ ] ID-81
	- [BaseCallback.onClaimProceeds(uint96,uint96,uint96,bytes)](src/callbacks/BaseCallback.sol#L171-L181)

src/callbacks/BaseCallback.sol#L171-L181


 - [ ] ID-82
	- [BaseCallback.onPurchase(uint96,address,uint96,uint96,bool,bytes)](src/callbacks/BaseCallback.sol#L126-L139)

src/callbacks/BaseCallback.sol#L126-L139


 - [ ] ID-83
	- [LinearVesting.mint(address,uint256,uint256,bool)](src/modules/derivatives/LinearVesting.sol#L293-L311)

src/modules/derivatives/LinearVesting.sol#L293-L311


 - [ ] ID-84
	- [AuctionModule.claimBids(uint96,uint64[])](src/modules/Auction.sol#L545-L560)

src/modules/Auction.sol#L545-L560


 - [ ] ID-85
	- [BaseCallback.onCurate(uint96,uint96,bool,bytes)](src/callbacks/BaseCallback.sol#L106-L117)

src/callbacks/BaseCallback.sol#L106-L117


 - [ ] ID-86
	- [LinearVesting.redeemMax(uint256)](src/modules/derivatives/LinearVesting.sol#L355-L364)

src/modules/derivatives/LinearVesting.sol#L355-L364


 - [ ] ID-87
	- [AuctionModule.claimProceeds(uint96)](src/modules/Auction.sol#L640-L654)

src/modules/Auction.sol#L640-L654


 - [ ] ID-88
	- [LinearVesting.wrap(uint256,uint256)](src/modules/derivatives/LinearVesting.sol#L462-L481)

src/modules/derivatives/LinearVesting.sol#L462-L481


 - [ ] ID-89
	- [Module.INIT()](src/modules/Modules.sol#L391)

src/modules/Modules.sol#L391


 - [ ] ID-90
	- [SoulboundCloneERC20.burn(address,uint256)](src/modules/derivatives/SoulboundCloneERC20.sol#L48-L50)

src/modules/derivatives/SoulboundCloneERC20.sol#L48-L50


 - [ ] ID-91
	- [LinearVesting.redeem(uint256,uint256)](src/modules/derivatives/LinearVesting.sol#L371-L385)

src/modules/derivatives/LinearVesting.sol#L371-L385


 - [ ] ID-92
	- [AuctionModule.bid(uint96,address,address,uint96,bytes)](src/modules/Auction.sol#L446-L461)

src/modules/Auction.sol#L446-L461


 - [ ] ID-93
	- [BaseCallback.onCancel(uint96,uint96,bool,bytes)](src/callbacks/BaseCallback.sol#L86-L97)

src/callbacks/BaseCallback.sol#L86-L97


 - [ ] ID-94
	- [LinearVesting.unwrap(uint256,uint256)](src/modules/derivatives/LinearVesting.sol#L488-L506)

src/modules/derivatives/LinearVesting.sol#L488-L506


 - [ ] ID-95
	- [BaseCallback.onBid(uint96,uint64,address,uint96,bytes)](src/callbacks/BaseCallback.sol#L150-L161)

src/callbacks/BaseCallback.sol#L150-L161


 - [ ] ID-96
	- [SoulboundCloneERC20.mint(address,uint256)](src/modules/derivatives/SoulboundCloneERC20.sol#L42-L44)

src/modules/derivatives/SoulboundCloneERC20.sol#L42-L44


 - [ ] ID-97
	- [AuctionModule.refundBid(uint96,uint64,address)](src/modules/Auction.sol#L501-L516)

src/modules/Auction.sol#L501-L516


## dead-function
Impact: Informational
Confidence: Medium
 - [ ] ID-98
[WithModules._getModuleIfInstalled(Keycode,uint8)](src/modules/Modules.sol#L251-L268) is never used and should be removed

src/modules/Modules.sol#L251-L268


 - [ ] ID-99
[AuctionModule._revertIfLotStarted(uint96)](src/modules/Auction.sol#L727-L729) is never used and should be removed

src/modules/Auction.sol#L727-L729


## price-manipulation-info
Impact: Informational
Confidence: Medium
 - [ ] ID-100
Potential price manipulation risk:
	- In function redeemMax
		-- [redeemableAmount = redeemable(msg.sender,tokenId_)](src/modules/derivatives/LinearVesting.sol#L357) have potential price manipulated risk from redeemableAmount and call None which could influence variable:redeemableAmount
	- In function redeemable
		-- [wrappedBalance = SoulboundCloneERC20(token.wrapped).balanceOf(owner_)](src/modules/derivatives/LinearVesting.sol#L413-L414) have potential price manipulated risk from wrappedBalance and call None which could influence variable:vested

src/modules/derivatives/LinearVesting.sol#L357


## uncontrolled-resource-consumption
Impact: Informational
Confidence: Medium
 - [ ] ID-101
Potential DoS Gas Limit Attack occur in[ClonesWithImmutableArgs.clone(address,bytes)](src/lib/clones/ClonesWithImmutableArgs.sol#L25-L152)[BEGIN_LOOP](src/lib/clones/ClonesWithImmutableArgs.sol#L125-L133)

src/lib/clones/ClonesWithImmutableArgs.sol#L25-L152


 - [ ] ID-102
Potential DoS Gas Limit Attack occur in[AuctionHouse.claimBids(uint96,uint64[])](src/AuctionHouse.sol#L403-L448)[BEGIN_LOOP](src/AuctionHouse.sol#L420-L447)

src/AuctionHouse.sol#L403-L448


 - [ ] ID-103
Potential DoS Gas Limit Attack occur in[MaxPriorityQueue._sink(Queue,uint64)](src/lib/MaxPriorityQueue.sol#L60-L72)[BEGIN_LOOP](src/lib/MaxPriorityQueue.sol#L61-L71)

src/lib/MaxPriorityQueue.sol#L60-L72


 - [ ] ID-104
Potential DoS Gas Limit Attack occur in[EncryptedMarginalPriceAuctionModule._settle(uint96)](src/modules/auctions/EMPAM.sol#L747-L837)[BEGIN_LOOP](src/modules/auctions/EMPAM.sol#L772-L781)

src/modules/auctions/EMPAM.sol#L747-L837


 - [ ] ID-105
Potential DoS Gas Limit Attack occur in[EncryptedMarginalPriceAuctionModule._claimBids(uint96,uint64[])](src/modules/auctions/EMPAM.sol#L372-L387)[BEGIN_LOOP](src/modules/auctions/EMPAM.sol#L378-L384)

src/modules/auctions/EMPAM.sol#L372-L387


 - [ ] ID-106
Potential DoS Gas Limit Attack occur in[EncryptedMarginalPriceAuctionModule._refundBid(uint96,uint64,address)](src/modules/auctions/EMPAM.sol#L284-L305)[BEGIN_LOOP](src/modules/auctions/EMPAM.sol#L295-L301)

src/modules/auctions/EMPAM.sol#L284-L305


 - [ ] ID-107
Potential DoS Gas Limit Attack occur in[EncryptedMarginalPriceAuctionModule._getLotMarginalPrice(uint96)](src/modules/auctions/EMPAM.sol#L595-L728)[BEGIN_LOOP](src/modules/auctions/EMPAM.sol#L611-L724)

src/modules/auctions/EMPAM.sol#L595-L728


 - [ ] ID-108
Potential DoS Gas Limit Attack occur in[EncryptedMarginalPriceAuctionModule._decryptAndSortBids(uint96,uint64)](src/modules/auctions/EMPAM.sol#L459-L522)[BEGIN_LOOP](src/modules/auctions/EMPAM.sol#L474-L513)

src/modules/auctions/EMPAM.sol#L459-L522


 - [ ] ID-109
Potential DoS Gas Limit Attack occur in[MaxPriorityQueue._swim(Queue,uint64)](src/lib/MaxPriorityQueue.sol#L52-L57)[BEGIN_LOOP](src/lib/MaxPriorityQueue.sol#L53-L56)

src/lib/MaxPriorityQueue.sol#L52-L57


 - [ ] ID-110
Potential DoS Gas Limit Attack occur in[ClonesWithImmutableArgs.clone3(address,bytes,bytes32)](src/lib/clones/ClonesWithImmutableArgs.sol#L160-L338)[BEGIN_LOOP](src/lib/clones/ClonesWithImmutableArgs.sol#L267-L275)

src/lib/clones/ClonesWithImmutableArgs.sol#L160-L338


## no-derived-function
Impact: Informational
Confidence: High
 - [ ] ID-111
[CondenserModule](src/modules/Condenser.sol#L13) does not implement functions:
	- [Condenser.condense(bytes,bytes)](src/modules/Condenser.sol#L7-L10)

src/modules/Condenser.sol#L13


## unnecessary-public-function-modifier
Impact: Informational
Confidence: High
 - [ ] ID-112
function:[Auction.maxAmountAccepted(uint96)](src/modules/Auction.sol#L241)is public and can be replaced with external 

src/modules/Auction.sol#L241


 - [ ] ID-113
function:[LinearVesting.name(uint256)](src/modules/derivatives/LinearVesting.sol#L719-L732)is public and can be replaced with external 

src/modules/derivatives/LinearVesting.sol#L719-L732


 - [ ] ID-114
function:[LinearVesting.transferFrom(address,address,uint256,uint256)](src/modules/derivatives/LinearVesting.sol#L129-L136)is public and can be replaced with external 

src/modules/derivatives/LinearVesting.sol#L129-L136


 - [ ] ID-115
function:[Timestamp.toPaddedString(uint48)](src/lib/Timestamp.sol#L7-L44)is public and can be replaced with external 

src/lib/Timestamp.sol#L7-L44


 - [ ] ID-116
function:[ERC6909.transferFrom(address,address,uint256,uint256)](lib/solmate/src/tokens/ERC6909.sol#L51-L73)is public and can be replaced with external 

lib/solmate/src/tokens/ERC6909.sol#L51-L73


 - [ ] ID-117
function:[CloneERC20.increaseAllowance(address,uint256)](src/lib/clones/CloneERC20.sol#L57-L63)is public and can be replaced with external 

src/lib/clones/CloneERC20.sol#L57-L63


 - [ ] ID-118
function:[ERC20.approve(address,uint256)](lib/solmate/src/tokens/ERC20.sol#L68-L74)is public and can be replaced with external 

lib/solmate/src/tokens/ERC20.sol#L68-L74


 - [ ] ID-119
function:[AuctionModule.hasEnded(uint96)](src/modules/Auction.sol#L687-L689)is public and can be replaced with external 

src/modules/Auction.sol#L687-L689


 - [ ] ID-120
function:[MaxPriorityQueue.insert(Queue,uint64,uint96,uint96)](src/lib/MaxPriorityQueue.sol#L75-L85)is public and can be replaced with external 

src/lib/MaxPriorityQueue.sol#L75-L85


 - [ ] ID-121
function:[LinearVesting.decimals(uint256)](src/modules/derivatives/LinearVesting.sol#L755-L767)is public and can be replaced with external 

src/modules/derivatives/LinearVesting.sol#L755-L767


 - [ ] ID-122
function:[CloneERC20.transfer(address,uint256)](src/lib/clones/CloneERC20.sol#L73-L85)is public and can be replaced with external 

src/lib/clones/CloneERC20.sol#L73-L85


 - [ ] ID-123
function:[EncryptedMarginalPriceAuctionModule.TYPE()](src/modules/auctions/EMPAM.sol#L128-L130)is public and can be replaced with external 

src/modules/auctions/EMPAM.sol#L128-L130


 - [ ] ID-124
function:[ERC20.transferFrom(address,address,uint256)](lib/solmate/src/tokens/ERC20.sol#L90-L110)is public and can be replaced with external 

lib/solmate/src/tokens/ERC20.sol#L90-L110


 - [ ] ID-125
function:[Auction.priceFor(uint96,uint96)](src/modules/Auction.sol#L237)is public and can be replaced with external 

src/modules/Auction.sol#L237


 - [ ] ID-126
function:[MaxPriorityQueue.initialize(Queue)](src/lib/MaxPriorityQueue.sol#L27-L29)is public and can be replaced with external 

src/lib/MaxPriorityQueue.sol#L27-L29


 - [ ] ID-127
function:[SoulboundCloneERC20.transfer(address,uint256)](src/modules/derivatives/SoulboundCloneERC20.sol#L54-L56)is public and can be replaced with external 

src/modules/derivatives/SoulboundCloneERC20.sol#L54-L56


 - [ ] ID-128
function:[ERC6909.transfer(address,uint256,uint256)](lib/solmate/src/tokens/ERC6909.sol#L33-L49)is public and can be replaced with external 

lib/solmate/src/tokens/ERC6909.sol#L33-L49


 - [ ] ID-129
function:[ERC6909.setOperator(address,bool)](lib/solmate/src/tokens/ERC6909.sol#L87-L93)is public and can be replaced with external 

lib/solmate/src/tokens/ERC6909.sol#L87-L93


 - [ ] ID-130
function:[MaxPriorityQueue.delMax(Queue)](src/lib/MaxPriorityQueue.sol#L96-L105)is public and can be replaced with external 

src/lib/MaxPriorityQueue.sol#L96-L105


 - [ ] ID-131
function:[LinearVesting.TYPE()](src/modules/derivatives/LinearVesting.sol#L96-L98)is public and can be replaced with external 

src/modules/derivatives/LinearVesting.sol#L96-L98


 - [ ] ID-132
function:[Auction.isLive(uint96)](src/modules/Auction.sol#L250)is public and can be replaced with external 

src/modules/Auction.sol#L250


 - [ ] ID-133
function:[FeeManager.calculateQuoteFees(uint96,uint96,bool,uint96)](src/bases/FeeManager.sol#L68-L88)is public and can be replaced with external 

src/bases/FeeManager.sol#L68-L88


 - [ ] ID-134
function:[LinearVesting.transfer(address,uint256,uint256)](src/modules/derivatives/LinearVesting.sol#L123-L125)is public and can be replaced with external 

src/modules/derivatives/LinearVesting.sol#L123-L125


 - [ ] ID-135
function:[Transfer.decodePermit2Approval(bytes)](src/lib/Transfer.sol#L163-L177)is public and can be replaced with external 

src/lib/Transfer.sol#L163-L177


 - [ ] ID-136
function:[ERC6909Metadata.symbol(uint256)](src/lib/ERC6909Metadata.sol#L15)is public and can be replaced with external 

src/lib/ERC6909Metadata.sol#L15


 - [ ] ID-137
function:[Transfer.permit2OrTransferFrom(ERC20,address,address,address,uint256,Transfer.Permit2Approval,bool)](src/lib/Transfer.sol#L142-L161)is public and can be replaced with external 

src/lib/Transfer.sol#L142-L161


 - [ ] ID-138
function:[Module.VEECODE()](src/modules/Modules.sol#L385)is public and can be replaced with external 

src/modules/Modules.sol#L385


 - [ ] ID-139
function:[ECIES.decrypt(uint256,Point,uint256,uint256)](src/lib/ECIES.sol#L59-L74)is public and can be replaced with external 

src/lib/ECIES.sol#L59-L74


 - [ ] ID-140
function:[ERC6909Metadata.name(uint256)](src/lib/ERC6909Metadata.sol#L9)is public and can be replaced with external 

src/lib/ERC6909Metadata.sol#L9


 - [ ] ID-141
function:[SoulboundCloneERC20.approve(address,uint256)](src/modules/derivatives/SoulboundCloneERC20.sol#L62-L64)is public and can be replaced with external 

src/modules/derivatives/SoulboundCloneERC20.sol#L62-L64


 - [ ] ID-142
function:[MaxPriorityQueue.getMax(Queue)](src/lib/MaxPriorityQueue.sol#L40-L44)is public and can be replaced with external 

src/lib/MaxPriorityQueue.sol#L40-L44


 - [ ] ID-143
function:[Transfer.approve(ERC20,address,uint256)](src/lib/Transfer.sol#L30-L32)is public and can be replaced with external 

src/lib/Transfer.sol#L30-L32


 - [ ] ID-144
function:[ERC20.permit(address,address,uint256,uint256,uint8,bytes32,bytes32)](lib/solmate/src/tokens/ERC20.sol#L116-L160)is public and can be replaced with external 

lib/solmate/src/tokens/ERC20.sol#L116-L160


 - [ ] ID-145
function:[Auction.hasEnded(uint96)](src/modules/Auction.sol#L259)is public and can be replaced with external 

src/modules/Auction.sol#L259


 - [ ] ID-146
function:[MaxPriorityQueue.getMaxId(Queue)](src/lib/MaxPriorityQueue.sol#L46-L49)is public and can be replaced with external 

src/lib/MaxPriorityQueue.sol#L46-L49


 - [ ] ID-147
function:[ECIES.isValid(Point)](src/lib/ECIES.sol#L137-L139)is public and can be replaced with external 

src/lib/ECIES.sol#L137-L139


 - [ ] ID-148
. analyzed (52 contracts with 86 detectors), 188 result(s) found
INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/2024-03-axis-finance_scaned_output/mwe-output.json generate success.
function:[MaxPriorityQueue.getNumBids(Queue)](src/lib/MaxPriorityQueue.sol#L35-L37)is public and can be replaced with external 

src/lib/MaxPriorityQueue.sol#L35-L37


 - [ ] ID-149
function:[EncryptedMarginalPriceAuctionModule.VEECODE()](src/modules/auctions/EMPAM.sol#L124-L126)is public and can be replaced with external 

src/modules/auctions/EMPAM.sol#L124-L126


 - [ ] ID-150
function:[ECIES.encrypt(uint256,Point,uint256,uint256)](src/lib/ECIES.sol#L83-L102)is public and can be replaced with external 

src/lib/ECIES.sol#L83-L102


 - [ ] ID-151
function:[ERC20.transfer(address,uint256)](lib/solmate/src/tokens/ERC20.sol#L76-L88)is public and can be replaced with external 

lib/solmate/src/tokens/ERC20.sol#L76-L88


 - [ ] ID-152
function:[Auction.maxPayout(uint96)](src/modules/Auction.sol#L239)is public and can be replaced with external 

src/modules/Auction.sol#L239


 - [ ] ID-153
function:[CloneERC20.transferFrom(address,address,uint256)](src/lib/clones/CloneERC20.sol#L87-L103)is public and can be replaced with external 

src/lib/clones/CloneERC20.sol#L87-L103


 - [ ] ID-154
function:[LinearVesting.symbol(uint256)](src/modules/derivatives/LinearVesting.sol#L737-L750)is public and can be replaced with external 

src/modules/derivatives/LinearVesting.sol#L737-L750


 - [ ] ID-155
function:[CloneERC20.approve(address,uint256)](src/lib/clones/CloneERC20.sol#L49-L55)is public and can be replaced with external 

src/lib/clones/CloneERC20.sol#L49-L55


 - [ ] ID-156
function:[FixedPriceAuctionModule.TYPE()](src/modules/auctions/FPAM.sol#L58-L60)is public and can be replaced with external 

src/modules/auctions/FPAM.sol#L58-L60


 - [ ] ID-157
function:[Auction.payoutFor(uint96,uint96)](src/modules/Auction.sol#L235)is public and can be replaced with external 

src/modules/Auction.sol#L235


 - [ ] ID-158
function:[ERC6909Metadata.decimals(uint256)](src/lib/ERC6909Metadata.sol#L21)is public and can be replaced with external 

src/lib/ERC6909Metadata.sol#L21


 - [ ] ID-159
function:[Transfer.transfer(ERC20,address,uint256,bool)](src/lib/Transfer.sol#L49-L68)is public and can be replaced with external 

src/lib/Transfer.sol#L49-L68


 - [ ] ID-160
function:[ERC6909.approve(address,uint256,uint256)](lib/solmate/src/tokens/ERC6909.sol#L75-L85)is public and can be replaced with external 

lib/solmate/src/tokens/ERC6909.sol#L75-L85


 - [ ] ID-161
function:[SoulboundCloneERC20.transferFrom(address,address,uint256)](src/modules/derivatives/SoulboundCloneERC20.sol#L58-L60)is public and can be replaced with external 

src/modules/derivatives/SoulboundCloneERC20.sol#L58-L60


 - [ ] ID-162
function:[Module.TYPE()](src/modules/Modules.sol#L381)is public and can be replaced with external 

src/modules/Modules.sol#L381


 - [ ] ID-163
function:[LinearVesting.validate(address,bytes)](src/modules/derivatives/LinearVesting.sol#L546-L554)is public and can be replaced with external 

src/modules/derivatives/LinearVesting.sol#L546-L554


 - [ ] ID-164
function:[SoulboundCloneERC20.owner()](src/modules/derivatives/SoulboundCloneERC20.sol#L78-L80)is public and can be replaced with external 

src/modules/derivatives/SoulboundCloneERC20.sol#L78-L80


 - [ ] ID-165
function:[LinearVesting.approve(address,uint256,uint256)](src/modules/derivatives/LinearVesting.sol#L140-L142)is public and can be replaced with external 

src/modules/derivatives/LinearVesting.sol#L140-L142


 - [ ] ID-166
function:[FixedPriceAuctionModule.VEECODE()](src/modules/auctions/FPAM.sol#L53-L55)is public and can be replaced with external 

src/modules/auctions/FPAM.sol#L53-L55


 - [ ] ID-167
function:[CloneERC20.decreaseAllowance(address,uint256)](src/lib/clones/CloneERC20.sol#L65-L71)is public and can be replaced with external 

src/lib/clones/CloneERC20.sol#L65-L71


 - [ ] ID-168
function:[Owned.transferOwnership(address)](lib/solmate/src/auth/Owned.sol#L39-L43)is public and can be replaced with external 

lib/solmate/src/auth/Owned.sol#L39-L43


 - [ ] ID-169
function:[ERC6909.supportsInterface(bytes4)](lib/solmate/src/tokens/ERC6909.sol#L99-L103)is public and can be replaced with external 

lib/solmate/src/tokens/ERC6909.sol#L99-L103


## version-only
Impact: Informational
Confidence: High
 - [ ] ID-170
	Pragma confirmed better, here is pragma solidity^0.8. [^0.8](src/lib/Uint2Str.sol#L2)

src/lib/Uint2Str.sol#L2


 - [ ] ID-171
	Pragma confirmed better, here is pragma solidity^0.8.4. [^0.8.4](lib/solady/src/utils/MerkleProofLib.sol#L2)

lib/solady/src/utils/MerkleProofLib.sol#L2


 - [ ] ID-172
	Pragma confirmed better, here is pragma solidity^0.8.19. [^0.8.19](src/callbacks/liquidity/DTL.sol#L2)

src/callbacks/liquidity/DTL.sol#L2


 - [ ] ID-173
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](src/interfaces/ICallback.sol#L2)

src/interfaces/ICallback.sol#L2


 - [ ] ID-174
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](lib/solmate/src/auth/Owned.sol#L2)

lib/solmate/src/auth/Owned.sol#L2


 - [ ] ID-175
	Pragma confirmed better, here is pragma solidity^0.8.19. [^0.8.19](src/lib/Callbacks.sol#L2)

src/lib/Callbacks.sol#L2


 - [ ] ID-176
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](src/lib/ECIES.sol#L2)

src/lib/ECIES.sol#L2


 - [ ] ID-177
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](src/lib/MaxPriorityQueue.sol#L2)

src/lib/MaxPriorityQueue.sol#L2


 - [ ] ID-178
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](lib/solmate/src/utils/SafeTransferLib.sol#L2)

lib/solmate/src/utils/SafeTransferLib.sol#L2


 - [ ] ID-179
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](lib/solmate/src/tokens/ERC20.sol#L2)

lib/solmate/src/tokens/ERC20.sol#L2


 - [ ] ID-180
	Pragma confirmed better, here is pragma solidity^0.8.4. [^0.8.4](src/lib/clones/ClonesWithImmutableArgs.sol#L3)

src/lib/clones/ClonesWithImmutableArgs.sol#L3


 - [ ] ID-181
	Pragma confirmed better, here is pragma solidity^0.8.4. [^0.8.4](src/lib/clones/Clone.sol#L2)

src/lib/clones/Clone.sol#L2


 - [ ] ID-182
	Pragma confirmed better, here is pragma solidity^0.8.19. [^0.8.19](src/lib/permit2/interfaces/IPermit2.sol#L2)

src/lib/permit2/interfaces/IPermit2.sol#L2


 - [ ] ID-183
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](lib/solmate/src/utils/FixedPointMathLib.sol#L2)

lib/solmate/src/utils/FixedPointMathLib.sol#L2


 - [ ] ID-184
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](lib/solmate/src/utils/ReentrancyGuard.sol#L2)

lib/solmate/src/utils/ReentrancyGuard.sol#L2


 - [ ] ID-185
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](src/lib/clones/CloneERC20.sol#L2)

src/lib/clones/CloneERC20.sol#L2


 - [ ] ID-186
	Pragma confirmed better, here is pragma solidity>=0.8.0. [>=0.8.0](lib/solmate/src/tokens/ERC6909.sol#L2)

lib/solmate/src/tokens/ERC6909.sol#L2


## unnecessary-reentrancy-lock
Impact: Optimization
Confidence: High
 - [ ] ID-187
unnecessary reentrancy lock found inBlastAuctionHouse
	- [FeeManager.claimRewards(address)](src/bases/FeeManager.sol#L123-L129)

src/bases/FeeManager.sol#L123-L129


Execution time for Falcon: 15.31611731s
