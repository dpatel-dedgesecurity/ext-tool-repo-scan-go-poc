'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi)
'forge config --json' running
'forge build --build-info --skip */test/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi)
Impossible to generate IR for UniswapV2Factory.createPair:
 'CompilationUnit' object has no attribute 'bytecode_init'

state variable: MockSwapper.WETH (contracts/mocks/MockSwapper.sol#27) not initialized and not written in contract but be used in contract
state variable: MockSwapper.USDC (contracts/mocks/MockSwapper.sol#28) not initialized and not written in contract but be used in contract
state variable: MockExchangeRouter.depositProps (contracts/mocks/gmx/MockExchangeRouter.sol#54) not initialized and not written in contract but be used in contract
state variable: MockExchangeRouter.withdrawalProps (contracts/mocks/gmx/MockExchangeRouter.sol#55) not initialized and not written in contract but be used in contract
state variable: MockExchangeRouter.eventProps (contracts/mocks/gmx/MockExchangeRouter.sol#56) not initialized and not written in contract but be used in contract
state variable: UniswapV2Factory.getPair (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#11) not initialized and not written in contract but be used in contract
state variable: UniswapV2Factory.allPairs (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#12) not initialized and not written in contract but be used in contract
Reference:  

	- GMXVault.receive() (contracts/strategy/gmx/GMXVault.sol#697-702)
Reference:  

MockChainlinkARBOracleWithAdjusts.addTokenMaxDelay(address,uint256) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#260-266) contains a tautology or contradiction:
	- maxDelay < 0 (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#263)
MockChainlinkARBOracleWithAdjusts.addTokenMaxDeviation(address,uint256) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#273-279) contains a tautology or contradiction:
	- maxDeviation < 0 (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#276)
ChainlinkARBOracle.addTokenMaxDelay(address,uint256) (contracts/oracles/ChainlinkARBOracle.sol#249-255) contains a tautology or contradiction:
	- maxDelay < 0 (contracts/oracles/ChainlinkARBOracle.sol#252)
ChainlinkARBOracle.addTokenMaxDeviation(address,uint256) (contracts/oracles/ChainlinkARBOracle.sol#262-268) contains a tautology or contradiction:
	- maxDeviation < 0 (contracts/oracles/ChainlinkARBOracle.sol#265)
ChainlinkOracle.addTokenMaxDelay(address,uint256) (contracts/oracles/ChainlinkOracle.sol#217-223) contains a tautology or contradiction:
	- maxDelay < 0 (contracts/oracles/ChainlinkOracle.sol#220)
ChainlinkOracle.addTokenMaxDeviation(address,uint256) (contracts/oracles/ChainlinkOracle.sol#230-236) contains a tautology or contradiction:
	- maxDeviation < 0 (contracts/oracles/ChainlinkOracle.sol#233)
Reference:  

Potential price manipulation risk:
	- In function emergencyWithdraw
		-- _userShareBalance = IERC20(address(self.vault)).balanceOf(msg.sender) (contracts/strategy/gmx/GMXEmergency.sol#167) have potential price manipulated risk from _userShareBalance and call None which could influence variable:shareAmt
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

public mint or burn found in MockERC20.mint(address,uint256) (contracts/mocks/MockERC20.sol#13-15)public mint or burn found in MockERC20.burn(address,uint256) (contracts/mocks/MockERC20.sol#17-19)public mint or burn found in MockWETH.mint(address,uint256) (contracts/mocks/MockWETH.sol#10-12)public mint or burn found in MockWETH.burn(address,uint256) (contracts/mocks/MockWETH.sol#14-16)Reference: check public mint method

minTokenAAmt is a member never initialized in GMXRebalance.rebalanceRemove(GMXTypes.Store,GMXTypes.RebalanceRemoveParams)._rlp (contracts/strategy/gmx/GMXRebalance.sol#173)
repayParams is a member never initialized in GMXWithdraw.withdraw(GMXTypes.Store,GMXTypes.WithdrawParams)._wc (contracts/strategy/gmx/GMXWithdraw.sol#63)
minTokenAAmt is a member never initialized in GMXWithdraw.withdraw(GMXTypes.Store,GMXTypes.WithdrawParams)._rlp (contracts/strategy/gmx/GMXWithdraw.sol#103)
GMXDeposit.processDeposit(GMXTypes.Store).reason (contracts/strategy/gmx/GMXDeposit.sol#182) is a local variable never initialized
repayTokenAAmt is a member never initialized in GMXEmergency.emergencyClose(GMXTypes.Store,uint256)._rp (contracts/strategy/gmx/GMXEmergency.sol#118)
MockGMXOracleWithAdjusts._getTokenPriceMinMaxFormatted(address)._tokenPriceMinMaxFormatted (contracts/mocks/MockGMXOracleWithAdjusts.sol#325) is a local variable never initialized
GMXRebalance.processRebalanceRemove(GMXTypes.Store).reason (contracts/strategy/gmx/GMXRebalance.sol#234) is a local variable never initialized
GMXManager.calcBorrow(GMXTypes.Store,uint256)._borrowLongTokenAmt (contracts/strategy/gmx/GMXManager.sol#82) is a local variable never initialized
borrowParams is a member never initialized in GMXDeposit.deposit(GMXTypes.Store,GMXTypes.DepositParams,bool)._dc (contracts/strategy/gmx/GMXDeposit.sol#86)
UniswapV2Library.getAmountsOut(address,uint256,address[]).i (contracts/mocks/gmx/MockUniswapV2/libraries/UniswapV2Library.sol#65) is a local variable never initialized
GMXManager.calcBorrow(GMXTypes.Store,uint256)._borrowShortTokenAmt (contracts/strategy/gmx/GMXManager.sol#83) is a local variable never initialized
GMXReader.additionalCapacity(GMXTypes.Store)._additionalCapacity (contracts/strategy/gmx/GMXReader.sol#231) is a local variable never initialized
GMXWithdraw.processWithdraw(GMXTypes.Store).reason (contracts/strategy/gmx/GMXWithdraw.sol#206) is a local variable never initialized
UniswapV2Router02._swapSupportingFeeOnTransferTokens(address[],address).i (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#351) is a local variable never initialized
GMXRebalance.processRebalanceAdd(GMXTypes.Store).reason (contracts/strategy/gmx/GMXRebalance.sol#114) is a local variable never initialized
minTokenAAmt is a member never initialized in GMXDeposit.processDepositFailure(GMXTypes.Store,uint256,uint256)._rlp (contracts/strategy/gmx/GMXDeposit.sol#235)
UniswapV2Router02._swap(uint256[],address[],address).i (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#230) is a local variable never initialized
Reference:  

MockExchangeRouter.executeDeposit(address,address,address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#70-88)have ignores return value in uniswapRouter.addLiquidity(tokenA,tokenB,IERC20(tokenA).balanceOf(address(this)),IERC20(tokenB).balanceOf(address(this)),vault,block.timestamp + 1) (contracts/mocks/gmx/MockExchangeRouter.sol#78-85)
MockExchangeRouter.executeWithdrawal(address,address,address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#128-147)have ignores return value in uniswapRouter.removeLiquidity(tokenA,tokenB,IERC20(pair).balanceOf(address(this)),0,0,vault,block.timestamp + 1) (contracts/mocks/gmx/MockExchangeRouter.sol#136-144)
MockExchangeRouter._swapForDeposit(address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#249-287)have ignores return value in uniswapRouter.swapExactTokensForTokens(optimalSwapAmount,0,swapPathForOptimalDeposit,address(this),block.timestamp) (contracts/mocks/gmx/MockExchangeRouter.sol#279-285)
UniswapV2Router02._addLiquidity(address,address,uint256,uint256,uint256,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#35-62)have ignores return value in IUniswapV2Factory(factory).createPair(tokenA,tokenB) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#45)
GMXCompound.compound(GMXTypes.Store,GMXTypes.CompoundParams) (contracts/strategy/gmx/GMXCompound.sol#35-107)have ignores return value in GMXManager.swapExactTokensForTokens(self,_sp) (contracts/strategy/gmx/GMXCompound.sol#72)
GMXDeposit.processDepositFailureLiquidityWithdrawal(GMXTypes.Store) (contracts/strategy/gmx/GMXDeposit.sol#282-352)have ignores return value in GMXManager.swapTokensForExactTokens(self,_sp) (contracts/strategy/gmx/GMXDeposit.sol#317)
GMXEmergency.emergencyPause(GMXTypes.Store) (contracts/strategy/gmx/GMXEmergency.sol#47-66)have ignores return value in GMXManager.removeLiquidity(self,_rlp) (contracts/strategy/gmx/GMXEmergency.sol#58-61)
GMXEmergency.emergencyResume(GMXTypes.Store) (contracts/strategy/gmx/GMXEmergency.sol#72-91)have ignores return value in GMXManager.addLiquidity(self,_alp) (contracts/strategy/gmx/GMXEmergency.sol#87-90)
GMXEmergency.emergencyClose(GMXTypes.Store,uint256) (contracts/strategy/gmx/GMXEmergency.sol#111-156)have ignores return value in GMXManager.swapTokensForExactTokens(self,_sp) (contracts/strategy/gmx/GMXEmergency.sol#141)
GMXProcessWithdraw.processWithdraw(GMXTypes.Store) (contracts/strategy/gmx/GMXProcessWithdraw.sol#24-106)have ignores return value in GMXManager.swapTokensForExactTokens(self,_sp) (contracts/strategy/gmx/GMXProcessWithdraw.sol#52)
GMXTrove.constructor(address) (contracts/strategy/gmx/GMXTrove.sol#32-40)have ignores return value in _store.tokenA.approve(address(vault),type()(uint256).max) (contracts/strategy/gmx/GMXTrove.sol#38)
GMXVault.constructor(string,string,GMXTypes.Store) (contracts/strategy/gmx/GMXVault.sol#67-132)have ignores return value in _store.tokenA.approve(address(_store.router),type()(uint256).max) (contracts/strategy/gmx/GMXVault.sol#118)
GMXWorker.swapExactTokensForTokens(GMXTypes.Store,ISwap.SwapParams) (contracts/strategy/gmx/GMXWorker.sol#114-121)have ignores return value in IERC20(sp.tokenIn).approve(address(self.swapRouter),sp.amountIn) (contracts/strategy/gmx/GMXWorker.sol#118)
GMXWorker.swapTokensForExactTokens(GMXTypes.Store,ISwap.SwapParams) (contracts/strategy/gmx/GMXWorker.sol#129-136)have ignores return value in IERC20(sp.tokenIn).approve(address(self.swapRouter),sp.amountIn) (contracts/strategy/gmx/GMXWorker.sol#133)
TraderJoeSwap.swapExactTokensForTokens(ISwap.SwapParams) (contracts/swaps/TraderJoeSwap.sol#60-99)have ignores return value in IERC20(sp.tokenIn).approve(address(router),sp.amountIn) (contracts/swaps/TraderJoeSwap.sol#63)
TraderJoeSwap.swapTokensForExactTokens(ISwap.SwapParams) (contracts/swaps/TraderJoeSwap.sol#106-151)have ignores return value in IERC20(sp.tokenIn).approve(address(router),sp.amountIn) (contracts/swaps/TraderJoeSwap.sol#113)
UniswapSwap.swapExactTokensForTokens(ISwap.SwapParams) (contracts/swaps/UniswapSwap.sol#60-92)have ignores return value in IERC20(sp.tokenIn).approve(address(router),sp.amountIn) (contracts/swaps/UniswapSwap.sol#63)
UniswapSwap.swapTokensForExactTokens(ISwap.SwapParams) (contracts/swaps/UniswapSwap.sol#99-134)have ignores return value in IERC20(sp.tokenIn).approve(address(router),sp.amountIn) (contracts/swaps/UniswapSwap.sol#106)
Reference:  

	- GMXVault.updateTreasury(address) (contracts/strategy/gmx/GMXVault.sol#575-578)
	- GMXVault.updateSwapRouter(address) (contracts/strategy/gmx/GMXVault.sol#585-588)
	- GMXVault.updateTrove(address) (contracts/strategy/gmx/GMXVault.sol#595-598)
	- GMXVault.updateCallback(address) (contracts/strategy/gmx/GMXVault.sol#605-608)
	- GMXVault.updateFeePerSecond(uint256) (contracts/strategy/gmx/GMXVault.sol#615-618)
	- GMXVault.updateParameterLimits(uint256,uint256,uint256,int256,int256) (contracts/strategy/gmx/GMXVault.sol#629-649)
	- GMXVault.updateMinSlippage(uint256) (contracts/strategy/gmx/GMXVault.sol#656-659)
	- GMXVault.updateMinExecutionFee(uint256) (contracts/strategy/gmx/GMXVault.sol#666-669)
Reference:  

Setter function LendingVault.constructor(string,string,IERC20,bool,uint256,uint256,address,ILendingVault.InterestRate,ILendingVault.InterestRate) (contracts/lending/LendingVault.sol#115-147) does not emit an event
Setter function LendingVault.withdrawReserve(uint256) (contracts/lending/LendingVault.sol#371-382) does not emit an event
Setter function LendingVault._onlyBorrower() (contracts/lending/LendingVault.sol#389-391) does not emit an event
Setter function LendingVault._onlyKeeper() (contracts/lending/LendingVault.sol#396-398) does not emit an event
Setter function LendingVault._mintShares(uint256) (contracts/lending/LendingVault.sol#405-418) does not emit an event
Setter function LendingVault._burnShares(uint256) (contracts/lending/LendingVault.sol#425-434) does not emit an event
Setter function LendingVault._updateVaultWithInterestsAndTimestamp(uint256) (contracts/lending/LendingVault.sol#440-447) does not emit an event
Setter function LendingVault.approveBorrower(address) (contracts/lending/LendingVault.sol#554-558) does not emit an event
Setter function LendingVault.revokeBorrower(address) (contracts/lending/LendingVault.sol#564-568) does not emit an event
Setter function LendingVault.updateKeeper(address,bool) (contracts/lending/LendingVault.sol#575-579) does not emit an event
Setter function LendingVault.updateTreasury(address) (contracts/lending/LendingVault.sol#668-672) does not emit an event
Setter function LendingVault.onlyBorrower() (contracts/lending/LendingVault.sol#89-92) does not emit an event
Setter function LendingVault.onlyKeeper() (contracts/lending/LendingVault.sol#97-100) does not emit an event
Setter function MockChainlinkARBOracleWithAdjusts.changeAdjust(uint256) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#18-22) does not emit an event
Setter function MockChainlinkARBOracleWithAdjusts.constructor(address) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#58-62) does not emit an event
Setter function MockChainlinkARBOracleWithAdjusts.addTokenPriceFeed(address,address) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#247-253) does not emit an event
Setter function MockChainlinkARBOracleWithAdjusts.addTokenMaxDelay(address,uint256) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#260-266) does not emit an event
Setter function MockChainlinkARBOracleWithAdjusts.addTokenMaxDeviation(address,uint256) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#273-279) does not emit an event
Setter function MockChainlinkARBOracleWithAdjusts.emergencyPause() (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#284-286) does not emit an event
Setter function MockChainlinkARBOracleWithAdjusts.emergencyResume() (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#291-293) does not emit an event
Setter function MockChainlinkOracle.constructor() (contracts/mocks/MockChainlinkOracle.sol#25) does not emit an event
Setter function MockChainlinkOracle.set(address,uint256,uint8) (contracts/mocks/MockChainlinkOracle.sol#28-31) does not emit an event
Setter function MockGMXOracleWithAdjusts.changeAdjust(uint256,uint256) (contracts/mocks/MockGMXOracleWithAdjusts.sol#18-24) does not emit an event
Setter function MockGMXOracleWithAdjusts.constructor(address,ISyntheticReader,IChainlinkOracle) (contracts/mocks/MockGMXOracleWithAdjusts.sol#46-58) does not emit an event
Setter function MockLendingVault.constructor(string,string,IERC20,bool,uint256,uint256,address,ILendingVault.InterestRate,ILendingVault.InterestRate) (contracts/mocks/MockLendingVault.sol#117-149) does not emit an event
Setter function MockLendingVault.mockSetDebt(address,uint256) (contracts/mocks/MockLendingVault.sol#152-154) does not emit an event
Setter function MockLendingVault.withdrawReserve(uint256) (contracts/mocks/MockLendingVault.sol#378-389) does not emit an event
Setter function MockLendingVault._onlyBorrower() (contracts/mocks/MockLendingVault.sol#396-398) does not emit an event
Setter function MockLendingVault._onlyKeeper() (contracts/mocks/MockLendingVault.sol#403-405) does not emit an event
Setter function MockLendingVault._mintShares(uint256) (contracts/mocks/MockLendingVault.sol#412-425) does not emit an event
Setter function MockLendingVault._burnShares(uint256) (contracts/mocks/MockLendingVault.sol#432-441) does not emit an event
Setter function MockLendingVault._updateVaultWithInterestsAndTimestamp(uint256) (contracts/mocks/MockLendingVault.sol#447-454) does not emit an event
Setter function MockLendingVault.approveBorrower(address) (contracts/mocks/MockLendingVault.sol#561-565) does not emit an event
Setter function MockLendingVault.revokeBorrower(address) (contracts/mocks/MockLendingVault.sol#571-575) does not emit an event
Setter function MockLendingVault.updateKeeper(address,bool) (contracts/mocks/MockLendingVault.sol#582-586) does not emit an event
Setter function MockLendingVault.updateTreasury(address) (contracts/mocks/MockLendingVault.sol#675-679) does not emit an event
Setter function MockLendingVault.onlyBorrower() (contracts/mocks/MockLendingVault.sol#91-94) does not emit an event
Setter function MockLendingVault.onlyKeeper() (contracts/mocks/MockLendingVault.sol#99-102) does not emit an event
Setter function MockStrategyVault.deposit(uint256) (contracts/mocks/MockStrategyVault.sol#13-15) does not emit an event
Setter function MockStrategyVault.withdraw(uint256) (contracts/mocks/MockStrategyVault.sol#17-19) does not emit an event
Setter function MockSwapper.swap(ISwap.SwapParams) (contracts/mocks/MockSwapper.sol#31-53) does not emit an event
Setter function MockSwapper.setEthPrice(uint256) (contracts/mocks/MockSwapper.sol#55-57) does not emit an event
Setter function MockWETH.deposit() (contracts/mocks/MockWETH.sol#18-20) does not emit an event
Setter function MockWETH.withdraw(uint256) (contracts/mocks/MockWETH.sol#22-26) does not emit an event
Setter function MockAggregatorV3.setPreviousRoundData(MockAggregatorV3.PreviousRoundData) (contracts/mocks/chainlink/MockAggregatorV3.sol#32-38) does not emit an event
Setter function MockAggregatorV3.setCurrentRoundData(MockAggregatorV3.CurrentRoundData) (contracts/mocks/chainlink/MockAggregatorV3.sol#40-46) does not emit an event
Setter function MockExchangeRouter.createDeposit(IExchangeRouter.CreateDepositParams) (contracts/mocks/gmx/MockExchangeRouter.sol#65-68) does not emit an event
Setter function MockExchangeRouter.createWithdrawal(IExchangeRouter.CreateWithdrawalParams) (contracts/mocks/gmx/MockExchangeRouter.sol#120-126) does not emit an event
Setter function MockExchangeRouter.swapExactTokensForTokens(ISwap.SwapParams) (contracts/mocks/gmx/MockExchangeRouter.sol#180-199) does not emit an event
Setter function MockExchangeRouter.swapTokensForExactTokens(ISwap.SwapParams) (contracts/mocks/gmx/MockExchangeRouter.sol#201-220) does not emit an event
Setter function MockExchangeRouter.sendTokens(address,address,uint256) (contracts/mocks/gmx/MockExchangeRouter.sol#237-243) does not emit an event
Setter function MockSyntheticReader.setMarketTokenPrice(uint256) (contracts/mocks/gmx/MockSyntheticReader.sol#20-22) does not emit an event
Setter function UniswapV2Factory.setFeeTo(address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#41-44) does not emit an event
Setter function UniswapV2Factory.setFeeToSetter(address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#46-49) does not emit an event
Setter function UniswapV2Pair.constructor() (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#63-65) does not emit an event
Setter function UniswapV2Pair.initialize(address,address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#68-72) does not emit an event
Setter function UniswapV2Pair._mintFee(uint112,uint112) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#90-108) does not emit an event
Setter function UniswapV2Pair.lock() (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#33-38) does not emit an event
Setter function UniswapV2Router02.receive() (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#30-32) does not emit an event
Setter function UniswapV2Router02.addLiquidity(address,address,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#65-80) does not emit an event
Setter function UniswapV2Router02.addLiquidityETH(address,uint256,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#82-110) does not emit an event
Setter function UniswapV2Router02.removeLiquidity(address,address,uint256,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#113-129) does not emit an event
Setter function UniswapV2Router02.removeLiquidityETH(address,uint256,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#131-154) does not emit an event
Setter function UniswapV2Router02.removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#155-169) does not emit an event
Setter function UniswapV2Router02.removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#170-183) does not emit an event
Setter function UniswapV2Router02.removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#186-209) does not emit an event
Setter function UniswapV2Router02.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#210-225) does not emit an event
Setter function UniswapV2Router02.swapExactTokensForTokens(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#241-255) does not emit an event
Setter function UniswapV2Router02.swapTokensForExactTokens(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#256-270) does not emit an event
Setter function UniswapV2Router02.swapTokensForExactETH(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#286-305) does not emit an event
Setter function UniswapV2Router02.swapExactTokensForETH(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#306-325) does not emit an event
Setter function UniswapV2Router02.swapETHForExactTokens(uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#326-346) does not emit an event
Setter function UniswapV2Router02.swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#368-385) does not emit an event
Setter function UniswapV2Router02.swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#409-433) does not emit an event
Setter function ChainlinkARBOracle.constructor(address) (contracts/oracles/ChainlinkARBOracle.sol#48-52) does not emit an event
Setter function ChainlinkARBOracle.addTokenPriceFeed(address,address) (contracts/oracles/ChainlinkARBOracle.sol#236-242) does not emit an event
Setter function ChainlinkARBOracle.addTokenMaxDelay(address,uint256) (contracts/oracles/ChainlinkARBOracle.sol#249-255) does not emit an event
Setter function ChainlinkARBOracle.addTokenMaxDeviation(address,uint256) (contracts/oracles/ChainlinkARBOracle.sol#262-268) does not emit an event
Setter function ChainlinkARBOracle.emergencyPause() (contracts/oracles/ChainlinkARBOracle.sol#273-275) does not emit an event
Setter function ChainlinkARBOracle.emergencyResume() (contracts/oracles/ChainlinkARBOracle.sol#280-282) does not emit an event
Setter function ChainlinkOracle.constructor() (contracts/oracles/ChainlinkOracle.sol#39) does not emit an event
Setter function ChainlinkOracle.addTokenPriceFeed(address,address) (contracts/oracles/ChainlinkOracle.sol#204-210) does not emit an event
Setter function ChainlinkOracle.addTokenMaxDelay(address,uint256) (contracts/oracles/ChainlinkOracle.sol#217-223) does not emit an event
Setter function ChainlinkOracle.addTokenMaxDeviation(address,uint256) (contracts/oracles/ChainlinkOracle.sol#230-236) does not emit an event
Setter function ChainlinkOracle.emergencyPause() (contracts/oracles/ChainlinkOracle.sol#241-243) does not emit an event
Setter function ChainlinkOracle.emergencyResume() (contracts/oracles/ChainlinkOracle.sol#248-250) does not emit an event
Setter function GMXCallback.afterDepositExecution(bytes32,IDeposit.Props,IEvent.Props) (contracts/strategy/gmx/GMXCallback.sol#57-89) does not emit an event
Setter function GMXCallback.afterDepositCancellation(bytes32,IDeposit.Props,IEvent.Props) (contracts/strategy/gmx/GMXCallback.sol#96-115) does not emit an event
Setter function GMXCallback.afterWithdrawalExecution(bytes32,IWithdrawal.Props,IEvent.Props) (contracts/strategy/gmx/GMXCallback.sol#122-145) does not emit an event
Setter function GMXCallback.afterWithdrawalCancellation(bytes32,IWithdrawal.Props,IEvent.Props) (contracts/strategy/gmx/GMXCallback.sol#152-168) does not emit an event
Setter function GMXCallback.onlyController() (contracts/strategy/gmx/GMXCallback.sol#31-37) does not emit an event
Setter function GMXChecks.beforeEmergencyWithdrawChecks(GMXTypes.Store,uint256) (contracts/strategy/gmx/GMXChecks.sol#415-427) does not emit an event
Setter function GMXCompound.compound(GMXTypes.Store,GMXTypes.CompoundParams) (contracts/strategy/gmx/GMXCompound.sol#35-107) does not emit an event
Setter function GMXEmergency.emergencyResume(GMXTypes.Store) (contracts/strategy/gmx/GMXEmergency.sol#72-91) does not emit an event
Setter function GMXRebalance.rebalanceAdd(GMXTypes.Store,GMXTypes.RebalanceAddParams) (contracts/strategy/gmx/GMXRebalance.sol#35-96) does not emit an event
Setter function GMXRebalance.rebalanceRemove(GMXTypes.Store,GMXTypes.RebalanceRemoveParams) (contracts/strategy/gmx/GMXRebalance.sol#149-210) does not emit an event
Setter function GMXVault.constructor(string,string,GMXTypes.Store) (contracts/strategy/gmx/GMXVault.sol#67-132) does not emit an event
Setter function GMXVault.mintFee() (contracts/strategy/gmx/GMXVault.sol#334-337) does not emit an event
Setter function GMXVault._onlyVault() (contracts/strategy/gmx/GMXVault.sol#344-346) does not emit an event
Setter function GMXVault._onlyKeeper() (contracts/strategy/gmx/GMXVault.sol#351-353) does not emit an event
Setter function GMXVault.processDeposit() (contracts/strategy/gmx/GMXVault.sol#362-364) does not emit an event
Setter function GMXVault.processDepositCancellation() (contracts/strategy/gmx/GMXVault.sol#371-373) does not emit an event
Setter function GMXVault.processDepositFailure(uint256,uint256) (contracts/strategy/gmx/GMXVault.sol#381-386) does not emit an event
Setter function GMXVault.processDepositFailureLiquidityWithdrawal() (contracts/strategy/gmx/GMXVault.sol#393-395) does not emit an event
Setter function GMXVault.processWithdraw() (contracts/strategy/gmx/GMXVault.sol#402-404) does not emit an event
Setter function GMXVault.processWithdrawCancellation() (contracts/strategy/gmx/GMXVault.sol#411-413) does not emit an event
Setter function GMXVault.processWithdrawFailure(uint256,uint256) (contracts/strategy/gmx/GMXVault.sol#421-426) does not emit an event
Setter function GMXVault.processWithdrawFailureLiquidityAdded() (contracts/strategy/gmx/GMXVault.sol#433-435) does not emit an event
Setter function GMXVault.rebalanceAdd(GMXTypes.RebalanceAddParams) (contracts/strategy/gmx/GMXVault.sol#442-446) does not emit an event
Setter function GMXVault.processRebalanceAdd() (contracts/strategy/gmx/GMXVault.sol#453-455) does not emit an event
Setter function GMXVault.processRebalanceAddCancellation() (contracts/strategy/gmx/GMXVault.sol#462-464) does not emit an event
Setter function GMXVault.rebalanceRemove(GMXTypes.RebalanceRemoveParams) (contracts/strategy/gmx/GMXVault.sol#471-475) does not emit an event
Setter function GMXVault.processRebalanceRemove() (contracts/strategy/gmx/GMXVault.sol#482-484) does not emit an event
Setter function GMXVault.processRebalanceRemoveCancellation() (contracts/strategy/gmx/GMXVault.sol#491-493) does not emit an event
Setter function GMXVault.compound(GMXTypes.CompoundParams) (contracts/strategy/gmx/GMXVault.sol#501-503) does not emit an event
Setter function GMXVault.processCompound() (contracts/strategy/gmx/GMXVault.sol#510-512) does not emit an event
Setter function GMXVault.processCompoundCancellation() (contracts/strategy/gmx/GMXVault.sol#519-521) does not emit an event
Setter function GMXVault.emergencyPause() (contracts/strategy/gmx/GMXVault.sol#528-530) does not emit an event
Setter function GMXVault.emergencyResume() (contracts/strategy/gmx/GMXVault.sol#536-538) does not emit an event
Setter function GMXVault.processEmergencyResume() (contracts/strategy/gmx/GMXVault.sol#545-547) does not emit an event
Setter function GMXVault.emergencyClose(uint256) (contracts/strategy/gmx/GMXVault.sol#555-557) does not emit an event
Setter function GMXVault.mint(address,uint256) (contracts/strategy/gmx/GMXVault.sol#677-679) does not emit an event
Setter function GMXVault.burn(address,uint256) (contracts/strategy/gmx/GMXVault.sol#687-689) does not emit an event
Setter function GMXVault.receive() (contracts/strategy/gmx/GMXVault.sol#697-702) does not emit an event
Setter function GMXVault.onlyVault() (contracts/strategy/gmx/GMXVault.sol#48-51) does not emit an event
Setter function GMXVault.onlyKeeper() (contracts/strategy/gmx/GMXVault.sol#54-57) does not emit an event
Setter function TraderJoeSwap.constructor(ILBRouter,IChainlinkOracle) (contracts/swaps/TraderJoeSwap.sol#43-51) does not emit an event
Setter function TraderJoeSwap.swapExactTokensForTokens(ISwap.SwapParams) (contracts/swaps/TraderJoeSwap.sol#60-99) does not emit an event
Setter function TraderJoeSwap.swapTokensForExactTokens(ISwap.SwapParams) (contracts/swaps/TraderJoeSwap.sol#106-151) does not emit an event
Setter function UniswapSwap.constructor(ISwapRouter,IChainlinkOracle) (contracts/swaps/UniswapSwap.sol#43-51) does not emit an event
Setter function UniswapSwap.swapExactTokensForTokens(ISwap.SwapParams) (contracts/swaps/UniswapSwap.sol#60-92) does not emit an event
Setter function UniswapSwap.swapTokensForExactTokens(ISwap.SwapParams) (contracts/swaps/UniswapSwap.sol#99-134) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

Condition variable is not initialized found in UniswapV2Pair.initialize(address,address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#68-72)
Reference: initialize method should has permission check

variable lacks a zero-check on 		- Ownable2Step.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#35-38)
variable lacks a zero-check on 		- LendingVault.emergencyRepay(uint256,address) (contracts/lending/LendingVault.sol#585-614)
variable lacks a zero-check on 		- MockChainlinkOracle.set(address,uint256,uint8) (contracts/mocks/MockChainlinkOracle.sol#28-31)
variable lacks a zero-check on 		- MockGMXOracleWithAdjusts.getAmountsOut(address,address,address,address,address,uint256) (contracts/mocks/MockGMXOracleWithAdjusts.sol#73-116)
variable lacks a zero-check on 		- MockLendingVault.mockSetDebt(address,uint256) (contracts/mocks/MockLendingVault.sol#152-154)
variable lacks a zero-check on 		- MockLendingVault.emergencyRepay(uint256,address) (contracts/mocks/MockLendingVault.sol#592-621)
variable lacks a zero-check on 		- MockExchangeRouter.executeDeposit(address,address,address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#70-88)
variable lacks a zero-check on 		- MockExchangeRouter.executeMockDeposit(address,address,uint256,uint256,uint256,uint256,address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#90-107)
variable lacks a zero-check on 		- MockExchangeRouter.cancelDeposit(address,address,address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#109-118)
variable lacks a zero-check on 		- MockExchangeRouter.executeWithdrawal(address,address,address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#128-147)
variable lacks a zero-check on 		- MockExchangeRouter.executeMockWithdrawal(address,address,uint256,uint256,uint256,uint256,address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#149-166)
variable lacks a zero-check on 		- MockExchangeRouter.cancelWithdrawal(address,address,address,address) (contracts/mocks/gmx/MockExchangeRouter.sol#168-178)
variable lacks a zero-check on 		- MockExchangeRouter.getMarketTokenInfo(address,address,address,address,bool,bool) (contracts/mocks/gmx/MockExchangeRouter.sol#223-235)
variable lacks a zero-check on 		- MockExchangeRouter.sendTokens(address,address,uint256) (contracts/mocks/gmx/MockExchangeRouter.sol#237-243)
variable lacks a zero-check on 		- MockGMXOracle.getLpTokenValue(address,address,address,address,bool,bool) (contracts/mocks/gmx/MockGMXOracle.sol#19-37)
variable lacks a zero-check on 		- MockGMXOracle.getLpTokenReserves(address,address,address,address) (contracts/mocks/gmx/MockGMXOracle.sol#39-49)
variable lacks a zero-check on 		- MockUniswapV2Oracle.lpToken(address,address) (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#50-55)
variable lacks a zero-check on 		- MockUniswapV2Oracle.getLpTokenValue(uint256,address,address,IUniswapV2Pair) (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#173-206)
variable lacks a zero-check on 		- UniswapV2Factory.constructor(address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#16-18)
variable lacks a zero-check on 		- UniswapV2Factory.setFeeTo(address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#41-44)
variable lacks a zero-check on 		- UniswapV2Factory.setFeeToSetter(address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#46-49)
variable lacks a zero-check on 		- UniswapV2ERC20.permit(address,address,uint256,uint256,uint8,bytes32,bytes32) (contracts/mocks/gmx/MockUniswapV2/UniswapV2ERC20.sol#76-88)
variable lacks a zero-check on 		- UniswapV2Pair.initialize(address,address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#68-72)
variable lacks a zero-check on 		- UniswapV2Pair.skim(address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#203-208)
variable lacks a zero-check on 		- UniswapV2Router02.constructor(address,address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#25-28)
variable lacks a zero-check on 		- UniswapV2Router02.addLiquidity(address,address,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#65-80)
variable lacks a zero-check on 		- UniswapV2Router02.addLiquidityETH(address,uint256,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#82-110)
variable lacks a zero-check on 		- UniswapV2Router02.removeLiquidity(address,address,uint256,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#113-129)
variable lacks a zero-check on 		- UniswapV2Router02.removeLiquidityETH(address,uint256,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#131-154)
variable lacks a zero-check on 		- UniswapV2Router02.removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#155-169)
variable lacks a zero-check on 		- UniswapV2Router02.removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#170-183)
variable lacks a zero-check on 		- UniswapV2Router02.removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#186-209)
variable lacks a zero-check on 		- UniswapV2Router02.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#210-225)
variable lacks a zero-check on 		- UniswapV2Router02.swapTokensForExactETH(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#286-305)
variable lacks a zero-check on 		- UniswapV2Router02.swapExactTokensForETH(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#306-325)
variable lacks a zero-check on 		- UniswapV2Router02.swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#368-385)
variable lacks a zero-check on 		- UniswapV2Router02.swapExactETHForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#386-408)
variable lacks a zero-check on 		- UniswapV2Router02.swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#409-433)
variable lacks a zero-check on 		- GMXOracle.getAmountsOut(address,address,address,address,address,uint256) (contracts/oracles/GMXOracle.sol#61-104)
variable lacks a zero-check on 		- GMXDeposit.processDepositFailureLiquidityWithdrawal(GMXTypes.Store) (contracts/strategy/gmx/GMXDeposit.sol#282-352)
variable lacks a zero-check on 		- GMXEmergency.emergencyClose(GMXTypes.Store,uint256) (contracts/strategy/gmx/GMXEmergency.sol#111-156)
variable lacks a zero-check on 		- GMXProcessWithdraw.processWithdraw(GMXTypes.Store) (contracts/strategy/gmx/GMXProcessWithdraw.sol#24-106)
variable lacks a zero-check on 		- GMXReader.convertToUsdValue(GMXTypes.Store,address,uint256) (contracts/strategy/gmx/GMXReader.sol#62-70)
variable lacks a zero-check on 		- GMXVault.updateKeeper(address,bool) (contracts/strategy/gmx/GMXVault.sol#565-568)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

Potential price manipulation risk:
	- In function getLpTokenReserves
		-- (reserve0,reserve1) = _pair.getReserves() (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#148) have potential price manipulated risk from reserve1 and call _pair.getReserves() which could influence variable:reserveB
	- In function getLpTokenValue
		-- (totalReserveA,totalReserveB) = getLpTokenReserves(totalSupply,_tokenA,_tokenB,_pair) (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#181-186) have potential price manipulated risk from totalReserveB and call _pair.getReserves() which could influence variable:lpFairValueIn18
	- In function getLpTokenAmount
		-- lpTokenValue = getLpTokenValue(_pair.totalSupply(),_tokenA,_tokenB,_pair) (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#222-227) have potential price manipulated risk from lpTokenValue and call _pair.getReserves() which could influence variable:lpTokenAmount
	- In function getAmountsOut
		-- router.getAmountsOut(_amountIn,path)[1] (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#86)have potential price manipulated risk in return call router.getAmountsOut(_amountIn,path) could influence return value
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

	- Ownable2Step.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#35-38)
	- LendingVault.withdrawReserve(uint256) (contracts/lending/LendingVault.sol#371-382)
	- LendingVault.updateInterestRate(ILendingVault.InterestRate) (contracts/lending/LendingVault.sol#513-535)
	- LendingVault.updatePerformanceFee(uint256) (contracts/lending/LendingVault.sol#541-548)
	- LendingVault.approveBorrower(address) (contracts/lending/LendingVault.sol#554-558)
	- LendingVault.revokeBorrower(address) (contracts/lending/LendingVault.sol#564-568)
	- LendingVault.updateKeeper(address,bool) (contracts/lending/LendingVault.sol#575-579)
	- LendingVault.emergencyRepay(uint256,address) (contracts/lending/LendingVault.sol#585-614)
	- LendingVault.updateMaxCapacity(uint256) (contracts/lending/LendingVault.sol#638-642)
	- LendingVault.updateMaxInterestRate(ILendingVault.InterestRate) (contracts/lending/LendingVault.sol#648-662)
	- LendingVault.updateTreasury(address) (contracts/lending/LendingVault.sol#668-672)
	- MockChainlinkARBOracleWithAdjusts.changeAdjust(uint256) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#18-22)
	- MockChainlinkARBOracleWithAdjusts.addTokenPriceFeed(address,address) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#247-253)
	- MockChainlinkARBOracleWithAdjusts.addTokenMaxDelay(address,uint256) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#260-266)
	- MockChainlinkARBOracleWithAdjusts.addTokenMaxDeviation(address,uint256) (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#273-279)
	- MockGMXOracleWithAdjusts.changeAdjust(uint256,uint256) (contracts/mocks/MockGMXOracleWithAdjusts.sol#18-24)
	- MockLendingVault.updateInterestRate(ILendingVault.InterestRate) (contracts/mocks/MockLendingVault.sol#520-542)
	- MockLendingVault.updatePerformanceFee(uint256) (contracts/mocks/MockLendingVault.sol#548-555)
	- MockLendingVault.updateMaxInterestRate(ILendingVault.InterestRate) (contracts/mocks/MockLendingVault.sol#655-669)
	- MockLendingVault.updateTreasury(address) (contracts/mocks/MockLendingVault.sol#675-679)
	- UniswapV2Factory.setFeeTo(address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#41-44)
	- UniswapV2Factory.setFeeToSetter(address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#46-49)
	- UniswapV2Pair.initialize(address,address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#68-72)
	- ChainlinkARBOracle.addTokenPriceFeed(address,address) (contracts/oracles/ChainlinkARBOracle.sol#236-242)
	- ChainlinkARBOracle.addTokenMaxDelay(address,uint256) (contracts/oracles/ChainlinkARBOracle.sol#249-255)
	- ChainlinkARBOracle.addTokenMaxDeviation(address,uint256) (contracts/oracles/ChainlinkARBOracle.sol#262-268)
	- ChainlinkOracle.addTokenPriceFeed(address,address) (contracts/oracles/ChainlinkOracle.sol#204-210)
	- ChainlinkOracle.addTokenMaxDelay(address,uint256) (contracts/oracles/ChainlinkOracle.sol#217-223)
	- ChainlinkOracle.addTokenMaxDeviation(address,uint256) (contracts/oracles/ChainlinkOracle.sol#230-236)
	- GMXVault.updateKeeper(address,bool) (contracts/strategy/gmx/GMXVault.sol#565-568)
	- TraderJoeSwap.updatePairBinStep(address,address,uint256) (contracts/swaps/TraderJoeSwap.sol#164-170)
	- UniswapSwap.updateFee(address,address,uint24) (contracts/swaps/UniswapSwap.sol#147-153)
Reference:  

	- Ownable.renounceOwnership() (node_modules/@openzeppelin/contracts/access/Ownable.sol#76-78)
	- Ownable.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable.sol#84-89)
	- LendingVault.emergencyResume() (contracts/lending/LendingVault.sol#628-632)
	- MockChainlinkARBOracleWithAdjusts.emergencyPause() (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#284-286)
	- MockChainlinkARBOracleWithAdjusts.emergencyResume() (contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#291-293)
	- MockLendingVault.emergencyResume() (contracts/mocks/MockLendingVault.sol#635-639)
	- ChainlinkARBOracle.emergencyPause() (contracts/oracles/ChainlinkARBOracle.sol#273-275)
	- ChainlinkARBOracle.emergencyResume() (contracts/oracles/ChainlinkARBOracle.sol#280-282)
	- ChainlinkOracle.emergencyPause() (contracts/oracles/ChainlinkOracle.sol#241-243)
	- ChainlinkOracle.emergencyResume() (contracts/oracles/ChainlinkOracle.sol#248-250)
	- GMXCallback.afterDepositExecution(bytes32,IDeposit.Props,IEvent.Props) (contracts/strategy/gmx/GMXCallback.sol#57-89)
	- GMXCallback.afterDepositCancellation(bytes32,IDeposit.Props,IEvent.Props) (contracts/strategy/gmx/GMXCallback.sol#96-115)
	- GMXCallback.afterWithdrawalExecution(bytes32,IWithdrawal.Props,IEvent.Props) (contracts/strategy/gmx/GMXCallback.sol#122-145)
	- GMXCallback.afterWithdrawalCancellation(bytes32,IWithdrawal.Props,IEvent.Props) (contracts/strategy/gmx/GMXCallback.sol#152-168)
	- GMXVault.processDeposit() (contracts/strategy/gmx/GMXVault.sol#362-364)
	- GMXVault.processDepositCancellation() (contracts/strategy/gmx/GMXVault.sol#371-373)
	- GMXVault.processDepositFailure(uint256,uint256) (contracts/strategy/gmx/GMXVault.sol#381-386)
	- GMXVault.processDepositFailureLiquidityWithdrawal() (contracts/strategy/gmx/GMXVault.sol#393-395)
	- GMXVault.processWithdraw() (contracts/strategy/gmx/GMXVault.sol#402-404)
	- GMXVault.processWithdrawCancellation() (contracts/strategy/gmx/GMXVault.sol#411-413)
	- GMXVault.processWithdrawFailure(uint256,uint256) (contracts/strategy/gmx/GMXVault.sol#421-426)
	- GMXVault.processWithdrawFailureLiquidityAdded() (contracts/strategy/gmx/GMXVault.sol#433-435)
	- GMXVault.rebalanceAdd(GMXTypes.RebalanceAddParams) (contracts/strategy/gmx/GMXVault.sol#442-446)
	- GMXVault.processRebalanceAdd() (contracts/strategy/gmx/GMXVault.sol#453-455)
	- GMXVault.processRebalanceAddCancellation() (contracts/strategy/gmx/GMXVault.sol#462-464)
	- GMXVault.rebalanceRemove(GMXTypes.RebalanceRemoveParams) (contracts/strategy/gmx/GMXVault.sol#471-475)
	- GMXVault.processRebalanceRemove() (contracts/strategy/gmx/GMXVault.sol#482-484)
	- GMXVault.processRebalanceRemoveCancellation() (contracts/strategy/gmx/GMXVault.sol#491-493)
	- GMXVault.compound(GMXTypes.CompoundParams) (contracts/strategy/gmx/GMXVault.sol#501-503)
	- GMXVault.processCompound() (contracts/strategy/gmx/GMXVault.sol#510-512)
	- GMXVault.processCompoundCancellation() (contracts/strategy/gmx/GMXVault.sol#519-521)
	- GMXVault.emergencyPause() (contracts/strategy/gmx/GMXVault.sol#528-530)
	- GMXVault.emergencyResume() (contracts/strategy/gmx/GMXVault.sol#536-538)
	- GMXVault.processEmergencyResume() (contracts/strategy/gmx/GMXVault.sol#545-547)
	- GMXVault.emergencyClose(uint256) (contracts/strategy/gmx/GMXVault.sol#555-557)
	- GMXVault.mint(address,uint256) (contracts/strategy/gmx/GMXVault.sol#677-679)
	- GMXVault.burn(address,uint256) (contracts/strategy/gmx/GMXVault.sol#687-689)
Reference:  

Context._msgData() (node_modules/@openzeppelin/contracts/utils/Context.sol#21-23) is never used and should be removed
ReentrancyGuard._reentrancyGuardEntered() (node_modules/@openzeppelin/contracts/utils/ReentrancyGuard.sol#81-83) is never used and should be removed
Reference:  

require() missing error messages
	 - require(bool)(balanceOf(msg.sender) >= wad) (contracts/mocks/MockWETH.sol#23)
Reference: https://dev.to/tawseef/require-vs-assert-in-solidity-5e9d

Potential price manipulation risk:
	- In function getLpTokenReserves
		-- longTokenBalance = IERC20(longToken).balanceOf(marketToken) (contracts/mocks/gmx/MockGMXOracle.sol#45) have potential price manipulated risk from longTokenBalance and call None which could influence variable:longTokenBalance
	- In function getLpTokenReserves
		-- shortTokenBalance = IERC20(shortToken).balanceOf(marketToken) (contracts/mocks/gmx/MockGMXOracle.sol#46) have potential price manipulated risk from shortTokenBalance and call None which could influence variable:shortTokenBalance
Potential price manipulation risk:
	- In function emergencyWithdraw
		-- _userShareBalance = IERC20(address(self.vault)).balanceOf(msg.sender) (contracts/strategy/gmx/GMXEmergency.sol#167) have potential price manipulated risk from _userShareBalance and call None which could influence variable:_withdrawAmtTokenB
Potential price manipulation risk:
	- In function svTokenValue
		-- equityValue_ = equityValue(self) (contracts/strategy/gmx/GMXReader.sol#28) have potential price manipulated risk from equityValue_ and call None which could influence variable:equityValue_
	- In function equityValue
		-- assetValue_ = assetValue(self) (contracts/strategy/gmx/GMXReader.sol#131) have potential price manipulated risk from assetValue_ and call None which could influence variable:assetValue_
	- In function delta
		-- (_tokenAAmt) = assetAmt(self) (contracts/strategy/gmx/GMXReader.sol#195) have potential price manipulated risk from _tokenAAmt and call None which could influence variable:signedDelta
	- In function delta
		-- equityValue_ = equityValue(self) (contracts/strategy/gmx/GMXReader.sol#197) have potential price manipulated risk from equityValue_ and call None which could influence variable:signedDelta
Reference:  https://metatrust.feishu.cn/wiki/wikcnley0RNMaoaSzdjcCpYxYoD

Potential DoS Gas Limit Attack occur inUniswapV2Router02._swap(uint256[],address[],address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#229-240)BEGIN_LOOP (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#230-239)
Potential DoS Gas Limit Attack occur inUniswapV2Router02._swapSupportingFeeOnTransferTokens(address[],address) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#350-367)BEGIN_LOOP (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#351-366)
Reference: https://swcregistry.io/docs/SWC-128

function:Ownable2Step.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#35-38)is public and can be replaced with external 
function:Ownable2Step.acceptOwnership() (node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#52-58)is public and can be replaced with external 
function:Ownable.renounceOwnership() (node_modules/@openzeppelin/contracts/access/Ownable.sol#76-78)is public and can be replaced with external 
function:Ownable.transferOwnership(address) (node_modules/@openzeppelin/contracts/access/Ownable.sol#84-89)is public and can be replaced with external 
function:ERC20.name() (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#58-60)is public and can be replaced with external 
function:ERC20.symbol() (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#66-68)is public and can be replaced with external 
function:ERC20.decimals() (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#83-85)is public and can be replaced with external 
function:ERC20.transfer(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#109-113)is public and can be replaced with external 
function:ERC20.approve(address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#132-136)is public and can be replaced with external 
function:ERC20.transferFrom(address,address,uint256) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#154-159)is public and can be replaced with external 
function:LendingVault.lvTokenValue() (contracts/lending/LendingVault.sol#181-190)is public and can be replaced with external 
function:LendingVault.lendingAPR() (contracts/lending/LendingVault.sol#204-216)is public and can be replaced with external 
function:LendingVault.depositNative(uint256,uint256) (contracts/lending/LendingVault.sol#238-254)is public and can be replaced with external 
function:LendingVault.deposit(uint256,uint256) (contracts/lending/LendingVault.sol#261-275)is public and can be replaced with external 
function:LendingVault.withdraw(uint256,uint256) (contracts/lending/LendingVault.sol#282-303)is public and can be replaced with external 
function:LendingVault.updateInterestRate(ILendingVault.InterestRate) (contracts/lending/LendingVault.sol#513-535)is public and can be replaced with external 
function:LendingVault.updateMaxInterestRate(ILendingVault.InterestRate) (contracts/lending/LendingVault.sol#648-662)is public and can be replaced with external 
function:ERC20.totalSupply() (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#90-92)is public and can be replaced with external 
function:ERC20.balanceOf(address) (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#97-99)is public and can be replaced with external 
function:MockERC20.decimals() (contracts/mocks/MockERC20.sol#21-23)is public and can be replaced with external 
function:MockGMXOracleWithAdjusts.getAmountsIn(address,address,address,address,address,uint256) (contracts/mocks/MockGMXOracleWithAdjusts.sol#131-147)is public and can be replaced with external 
function:MockGMXOracleWithAdjusts.getLpTokenReserves(address,address,address,address) (contracts/mocks/MockGMXOracleWithAdjusts.sol#211-236)is public and can be replaced with external 
function:MockGMXOracleWithAdjusts.getLpTokenAmount(uint256,address,address,address,address,bool,bool) (contracts/mocks/MockGMXOracleWithAdjusts.sol#296-315)is public and can be replaced with external 
function:MockLendingVault.lvTokenValue() (contracts/mocks/MockLendingVault.sol#188-197)is public and can be replaced with external 
function:MockLendingVault.lendingAPR() (contracts/mocks/MockLendingVault.sol#211-223)is public and can be replaced with external 
function:MockLendingVault.depositNative(uint256,uint256) (contracts/mocks/MockLendingVault.sol#245-261)is public and can be replaced with external 
function:MockLendingVault.deposit(uint256,uint256) (contracts/mocks/MockLendingVault.sol#268-282)is public and can be replaced with external 
function:MockLendingVault.withdraw(uint256,uint256) (contracts/mocks/MockLendingVault.sol#289-310)is public and can be replaced with external 
function:MockLendingVault._onlyBorrower() (contracts/mocks/MockLendingVault.sol#396-398)is public and can be replaced with external 
function:MockLendingVault._onlyKeeper() (contracts/mocks/MockLendingVault.sol#403-405)is public and can be replaced with external 
function:MockLendingVault.updateInterestRate(ILendingVault.InterestRate) (contracts/mocks/MockLendingVault.sol#520-542)is public and can be replaced with external 
function:MockLendingVault.updateMaxInterestRate(ILendingVault.InterestRate) (contracts/mocks/MockLendingVault.sol#655-669)is public and can be replaced with external 
function:MockWETH.deposit() (contracts/mocks/MockWETH.sol#18-20)is public and can be replaced with external 
function:MockWETH.withdraw(uint256) (contracts/mocks/MockWETH.sol#22-26)is public and can be replaced with external 
function:MockGMXOracle.getLpTokenValue(address,address,address,address,bool,bool) (contracts/mocks/gmx/MockGMXOracle.sol#19-37)is public and can be replaced with external 
function:MockGMXOracle.getLpTokenReserves(address,address,address,address) (contracts/mocks/gmx/MockGMXOracle.sol#39-49)is public and can be replaced with external 
function:MockUniswapV2Oracle.lpToken(address,address) (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#50-55)is public and can be replaced with external 
function:MockUniswapV2Oracle.getAmountsOut(uint256,address,address,IUniswapV2Pair) (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#65-87)is public and can be replaced with external 
function:MockUniswapV2Oracle.getAmountsIn(uint256,address,address,IUniswapV2Pair) (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#97-119)is public and can be replaced with external 
function:MockUniswapV2Oracle.getLpTokenAmount(uint256,address,address,IUniswapV2Pair) (contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#216-232)is public and can be replaced with external 
function:UniswapV2Router02.quote(uint256,uint256,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#436-438)is public and can be replaced with external 
function:UniswapV2Router02.getAmountOut(uint256,uint256,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#440-448)is public and can be replaced with external 
function:UniswapV2Router02.getAmountIn(uint256,uint256,uint256) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#450-458)is public and can be replaced with external 
function:UniswapV2Router02.getAmountsOut(uint256,address[]) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#460-468)is public and can be replaced with external 
function:UniswapV2Router02.getAmountsIn(uint256,address[]) (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#470-478)is public and can be replaced with external 
function:GMXOracle.getAmountsIn(address,address,address,address,address,uint256) (contracts/oracles/GMXOracle.sol#119-135)is public and can be replaced with external 
function:GMXOracle.getLpTokenReserves(address,address,address,address) (contracts/oracles/GMXOracle.sol#197-222)is public and can be replaced with external 
function:GMXOracle.getLpTokenAmount(uint256,address,address,address,address,bool,bool) (contracts/oracles/GMXOracle.sol#280-299)is public and can be replaced with external 
function:GMXManager.borrow(GMXTypes.Store,uint256,uint256) (contracts/strategy/gmx/GMXManager.sol#225-236)is public and can be replaced with external 
function:GMXManager.repay(GMXTypes.Store,uint256,uint256) (contracts/strategy/gmx/GMXManager.sol#244-255)is public and can be replaced with external 
function:GMXManager.addLiquidity(GMXTypes.Store,GMXTypes.AddLiquidityParams) (contracts/strategy/gmx/GMXManager.sol#263-268)is public and can be replaced with external 
function:GMXManager.removeLiquidity(GMXTypes.Store,GMXTypes.RemoveLiquidityParams) (contracts/strategy/gmx/GMXManager.sol#276-281)is public and can be replaced with external 
function:GMXReader.svTokenValue(GMXTypes.Store) (contracts/strategy/gmx/GMXReader.sol#27-32)is public and can be replaced with external 
function:GMXReader.valueToShares(GMXTypes.Store,uint256,uint256) (contracts/strategy/gmx/GMXReader.sol#48-56)is public and can be replaced with external 
function:GMXReader.leverage(GMXTypes.Store) (contracts/strategy/gmx/GMXReader.sol#185-188)is public and can be replaced with external 
function:GMXReader.delta(GMXTypes.Store) (contracts/strategy/gmx/GMXReader.sol#194-214)is public and can be replaced with external 
function:GMXReader.debtRatio(GMXTypes.Store) (contracts/strategy/gmx/GMXReader.sol#220-224)is public and can be replaced with external 
function:GMXReader.capacity(GMXTypes.Store) (contracts/strategy/gmx/GMXReader.sol#282-284)is public and can be replaced with external 
function:GMXVault.store() (contracts/strategy/gmx/GMXVault.sol#140-142)is public and can be replaced with external 
function:GMXVault.isTokenWhitelisted(address) (contracts/strategy/gmx/GMXVault.sol#149-151)is public and can be replaced with external 
function:GMXVault.svTokenValue() (contracts/strategy/gmx/GMXVault.sol#157-159)is public and can be replaced with external 
function:GMXVault.pendingFee() (contracts/strategy/gmx/GMXVault.sol#165-167)is public and can be replaced with external 
function:GMXVault.valueToShares(uint256,uint256) (contracts/strategy/gmx/GMXVault.sol#175-177)is public and can be replaced with external 
function:GMXVault.convertToUsdValue(address,uint256) (contracts/strategy/gmx/GMXVault.sol#185-187)is public and can be replaced with external 
function:GMXVault.tokenWeights() (contracts/strategy/gmx/GMXVault.sol#194-196)is public and can be replaced with external 
function:GMXVault.assetValue() (contracts/strategy/gmx/GMXVault.sol#203-205)is public and can be replaced with external 
function:GMXVault.debtValue() (contracts/strategy/gmx/GMXVault.sol#213-215)is public and can be replaced with external 
function:GMXVault.equityValue() (contracts/strategy/gmx/GMXVault.sol#222-224)is public and can be replaced with external 
function:GMXVault.assetAmt() (contracts/strategy/gmx/GMXVault.sol#231-233)is public and can be replaced with external 
function:GMXVault.debtAmt() (contracts/strategy/gmx/GMXVault.sol#240-242)is public and can be replaced with external 
function:GMXVault.lpAmt() (contracts/strategy/gmx/GMXVault.sol#248-250)is public and can be replaced with external 
function:GMXVault.leverage() (contracts/strategy/gmx/GMXVault.sol#256-258)is public and can be replaced with external 
function:GMXVault.delta() (contracts/strategy/gmx/GMXVault.sol#266-268)is public and can be replaced with external 
function:GMXVault.debtRatio() (contracts/strategy/gmx/GMXVault.sol#275-277)is public and can be replaced with external 
function:GMXVault.additionalCapacity() (contracts/strategy/gmx/GMXVault.sol#283-285)is public and can be replaced with external 
function:GMXVault.capacity() (contracts/strategy/gmx/GMXVault.sol#291-293)is public and can be replaced with external 
function:GMXVault.mintFee() (contracts/strategy/gmx/GMXVault.sol#334-337)is public and can be replaced with external 
Reference:  

	Pragma confirmed better, here is pragma solidity>=0.4.22<0.9.0. >=0.4.22<0.9.0 (lib/forge-std/src/console.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.0. ^0.8.0 (node_modules/@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol#2)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/access/Ownable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/interfaces/draft-IERC6093.sol#3)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/utils/Address.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/utils/Context.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/utils/Pausable.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/utils/ReentrancyGuard.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/utils/math/Math.sol#4)
	Pragma confirmed better, here is pragma solidity^0.8.20. ^0.8.20 (node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#5)
Reference:  

MockSwapper.USDC (contracts/mocks/MockSwapper.sol#28) should be constant
MockSwapper.WETH (contracts/mocks/MockSwapper.sol#27) should be constant
UniswapV2Pair.price0CumulativeLast (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#28) should be constant
UniswapV2Pair.price1CumulativeLast (contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#29) should be constant
Reference:  
. analyzed (92 contracts with 86 detectors), 414 result(s) found
INFO:Falcon:metatrust result: ../../repos-output/2023-10-SteadeFi_autogen_output/mwe-output.json generate success.
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
found function
found swap
found function
found swap
An error occurred: list index out of range
An error occurred: maximum recursion depth exceeded in comparison
after for recrusion
Execution time for Falcon: 43.939447578s
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/lending/ILendingVault.sol: Interface (1): 54 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/oracles/IChainlinkOracle.sol: Interface (1): 12 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/oracles/IGMXOracle.sol: Interface (1): 77 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IDeposit.sol: Interface (1): 54 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IDepositCallbackReceiver.sol: Interface (1): 27 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IDepositHandler.sol: Interface (1): 24 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IEvent.sol: Interface (1): 119 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IExchangeRouter.sol: Interface (1): 114 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IOrder.sol: Interface (1): 118 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IOrderCallbackReceiver.sol: Interface (1): 36 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IRoleStore.sol: Interface (1): 6 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/ISyntheticReader.sol: Interface (1): 153 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IWithdrawal.sol: Interface (1): 51 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/gmx/IWithdrawalCallbackReceiver.sol: Interface (1): 27 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/trader-joe/ILBRouter.sol: Interface (1): 99 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/protocols/uniswap/ISwapRouter.sol: Interface (1): 58 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/strategy/gmx/IGMXVault.sol: Interface (1): 72 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/strategy/gmx/IGMXVaultEvents.sol: Interface (1): 77 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/swap/ISwap.sol: Interface (1): 27 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/interfaces/tokens/IWNT.sol: Interface (1): 10 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/lending/LendingVault.sol: Smart Contract (1): 683 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol: Smart Contract (1): 294 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/MockChainlinkOracle.sol: Smart Contract (1): 57 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/MockERC20.sol: Smart Contract (1): 24 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/MockGMXOracleWithAdjusts.sol: Smart Contract (1): 335 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/MockLendingVault.sol: Smart Contract (1): 690 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/MockStrategyVault.sol: Smart Contract (1): 28 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/MockSwapper.sol: Smart Contract (1), Interface (1): 58 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/MockWETH.sol: Smart Contract (1): 28 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/chainlink/MockAggregatorV3.sol: Smart Contract (1): 97 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockArbSys.sol: Smart Contract (1): 12 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockExchangeRouter.sol: Smart Contract (1), Interface (1): 336 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockGMXOracle.sol: Smart Contract (1): 50 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockRealtimeFeedVerifier.sol: Smart Contract (1), Interface (1): 13 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockSyntheticReader.sol: Smart Contract (1): 23 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol: Smart Contract (1): 233 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/UniswapV2ERC20.sol: Smart Contract (1): 89 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol: Smart Contract (1): 50 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol: Smart Contract (1): 214 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol: Smart Contract (1): 479 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/interfaces/IERC20.sol: Interface (1): 18 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/interfaces/IUniswapV2Callee.sol: Interface (1): 6 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/interfaces/IUniswapV2ERC20.sol: Interface (1): 22 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/interfaces/IUniswapV2Factory.sol: Interface (1): 18 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/interfaces/IUniswapV2Pair.sol: Interface (1): 58 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/interfaces/IUniswapV2Router01.sol: Interface (1): 95 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/interfaces/IUniswapV2Router02.sol: Interface (1): 45 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/interfaces/IWETH.sol: Interface (1): 8 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/libraries/Math.sol: Library (1): 24 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/mocks/gmx/MockUniswapV2/libraries/UniswapV2Library.sol: Library (1): 81 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/oracles/ChainlinkARBOracle.sol: Smart Contract (1): 283 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/oracles/ChainlinkOracle.sol: Smart Contract (1): 251 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/oracles/GMXOracle.sol: Smart Contract (1): 316 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXCallback.sol: Smart Contract (1): 169 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXChecks.sol: Library (1): 452 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXCompound.sol: Library (1): 136 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXDeposit.sol: Library (1): 353 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXEmergency.sol: Library (1): 203 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXManager.sol: Library (1): 316 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXProcessDeposit.sol: Library (1): 34 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXProcessWithdraw.sol: Library (1): 107 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXReader.sol: Library (1): 285 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXRebalance.sol: Library (1): 258 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXTrove.sol: Smart Contract (1): 41 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXTypes.sol: Library (1): 332 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXVault.sol: Smart Contract (1): 703 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXWithdraw.sol: Library (1): 288 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/strategy/gmx/GMXWorker.sol: Library (1): 137 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/swaps/TraderJoeSwap.sol: Smart Contract (1): 171 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/swaps/UniswapSwap.sol: Smart Contract (1): 154 lines
/home/im/dedge/ext-tool-repo-scan-go-poc/input-repos/2023-10-SteadeFi/contracts/utils/Errors.sol: Library (1): 85 lines

Summary:
Total number of Solidity files: 71
Abstract Contracts Files: 0
Smart Contracts Files: 25
Libraries Files: 15
Interfaces Files: 28
Multiple Types Files: 3
******************************************
Abstract Contracts: 0
Smart Contracts: 28
Libraries: 15
Interfaces: 31
LOC: 10457
