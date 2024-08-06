'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-10-SteadeFi)
'forge config --json' running
'forge build --build-info --skip */test/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-10-SteadeFi)
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

UniswapV2Router02._swap(uint256[],address[],address).i (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#230) is a local variable never initialized
UniswapV2Library.getAmountsOut(address,uint256,address[]).i (contracts/mocks/gmx/MockUniswapV2/libraries/UniswapV2Library.sol#65) is a local variable never initialized
minTokenAAmt is a member never initialized in GMXDeposit.processDepositFailure(GMXTypes.Store,uint256,uint256)._rlp (contracts/strategy/gmx/GMXDeposit.sol#235)
GMXWithdraw.processWithdraw(GMXTypes.Store).reason (contracts/strategy/gmx/GMXWithdraw.sol#206) is a local variable never initialized
repayParams is a member never initialized in GMXWithdraw.withdraw(GMXTypes.Store,GMXTypes.WithdrawParams)._wc (contracts/strategy/gmx/GMXWithdraw.sol#63)
minTokenAAmt is a member never initialized in GMXWithdraw.withdraw(GMXTypes.Store,GMXTypes.WithdrawParams)._rlp (contracts/strategy/gmx/GMXWithdraw.sol#103)
minTokenAAmt is a member never initialized in GMXRebalance.rebalanceRemove(GMXTypes.Store,GMXTypes.RebalanceRemoveParams)._rlp (contracts/strategy/gmx/GMXRebalance.sol#173)
GMXReader.additionalCapacity(GMXTypes.Store)._additionalCapacity (contracts/strategy/gmx/GMXReader.sol#231) is a local variable never initialized
repayTokenAAmt is a member never initialized in GMXEmergency.emergencyClose(GMXTypes.Store,uint256)._rp (contracts/strategy/gmx/GMXEmergency.sol#118)
borrowParams is a member never initialized in GMXDeposit.deposit(GMXTypes.Store,GMXTypes.DepositParams,bool)._dc (contracts/strategy/gmx/GMXDeposit.sol#86)
UniswapV2Router02._swapSupportingFeeOnTransferTokens(address[],address).i (contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#351) is a local variable never initialized
GMXRebalance.processRebalanceAdd(GMXTypes.Store).reason (contracts/strategy/gmx/GMXRebalance.sol#114) is a local variable never initialized
GMXRebalance.processRebalanceRemove(GMXTypes.Store).reason (contracts/strategy/gmx/GMXRebalance.sol#234) is a local variable never initialized
GMXDeposit.processDeposit(GMXTypes.Store).reason (contracts/strategy/gmx/GMXDeposit.sol#182) is a local variable never initialized
GMXManager.calcBorrow(GMXTypes.Store,uint256)._borrowLongTokenAmt (contracts/strategy/gmx/GMXManager.sol#82) is a local variable never initialized
GMXManager.calcBorrow(GMXTypes.Store,uint256)._borrowShortTokenAmt (contracts/strategy/gmx/GMXManager.sol#83) is a local variable never initialized
MockGMXOracleWithAdjusts._getTokenPriceMinMaxFormatted(address)._tokenPriceMinMaxFormatted (contracts/mocks/MockGMXOracleWithAdjusts.sol#325) is a local variable never initialized
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
An error occurred: maximum recursion depth exceeded while calling a Python object
after for recrusion
Summary
 - [state-variable-not-initialized](#state-variable-not-initialized) (7 results) (High)
 - [centralized-risk-medium](#centralized-risk-medium) (1 results) (Medium)
 - [constant-result](#constant-result) (6 results) (Medium)
 - [price-manipulation-medium](#price-manipulation-medium) (1 results) (Medium)
 - [public-mint-burn](#public-mint-burn) (4 results) (Medium)
 - [uninitialized-local](#uninitialized-local) (17 results) (Medium)
 - [unused-return](#unused-return) (18 results) (Medium)
 - [centralized-risk-low](#centralized-risk-low) (8 results) (Low)
 - [pess-event-setter](#pess-event-setter) (132 results) (Low)
 - [initialize-permission](#initialize-permission) (1 results) (Low)
 - [missing-zero-check](#missing-zero-check) (44 results) (Low)
 - [price-manipulation-low](#price-manipulation-low) (1 results) (Low)
 - [centralized-risk-informational](#centralized-risk-informational) (32 results) (Informational)
 - [centralized-risk-other](#centralized-risk-other) (37 results) (Informational)
 - [dead-function](#dead-function) (2 results) (Informational)
 - [error-msg](#error-msg) (1 results) (Informational)
 - [price-manipulation-info](#price-manipulation-info) (2 results) (Informational)
 - [uncontrolled-resource-consumption](#uncontrolled-resource-consumption) (2 results) (Informational)
 - [unnecessary-public-function-modifier](#unnecessary-public-function-modifier) (77 results) (Informational)
 - [version-only](#version-only) (16 results) (Informational)
 - [state-should-be-constant](#state-should-be-constant) (4 results) (Optimization)
## state-variable-not-initialized
Impact: High
Confidence: High
 - [ ] ID-0
state variable: [MockExchangeRouter.eventProps](contracts/mocks/gmx/MockExchangeRouter.sol#L56) not initialized and not written in contract but be used in contract

contracts/mocks/gmx/MockExchangeRouter.sol#L56


 - [ ] ID-1
state variable: [MockExchangeRouter.withdrawalProps](contracts/mocks/gmx/MockExchangeRouter.sol#L55) not initialized and not written in contract but be used in contract

contracts/mocks/gmx/MockExchangeRouter.sol#L55


 - [ ] ID-2
state variable: [UniswapV2Factory.getPair](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L11) not initialized and not written in contract but be used in contract

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L11


 - [ ] ID-3
state variable: [MockSwapper.WETH](contracts/mocks/MockSwapper.sol#L27) not initialized and not written in contract but be used in contract

contracts/mocks/MockSwapper.sol#L27


 - [ ] ID-4
state variable: [MockExchangeRouter.depositProps](contracts/mocks/gmx/MockExchangeRouter.sol#L54) not initialized and not written in contract but be used in contract

contracts/mocks/gmx/MockExchangeRouter.sol#L54


 - [ ] ID-5
state variable: [MockSwapper.USDC](contracts/mocks/MockSwapper.sol#L28) not initialized and not written in contract but be used in contract

contracts/mocks/MockSwapper.sol#L28


 - [ ] ID-6
state variable: [UniswapV2Factory.allPairs](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L12) not initialized and not written in contract but be used in contract

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L12


## centralized-risk-medium
Impact: Medium
Confidence: High
 - [ ] ID-7
	- [GMXVault.receive()](contracts/strategy/gmx/GMXVault.sol#L697-L702)

contracts/strategy/gmx/GMXVault.sol#L697-L702


## constant-result
Impact: Medium
Confidence: High
 - [ ] ID-8
[MockChainlinkARBOracleWithAdjusts.addTokenMaxDelay(address,uint256)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L260-L266) contains a tautology or contradiction:
	- [maxDelay < 0](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L263)

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L260-L266


 - [ ] ID-9
[ChainlinkARBOracle.addTokenMaxDelay(address,uint256)](contracts/oracles/ChainlinkARBOracle.sol#L249-L255) contains a tautology or contradiction:
	- [maxDelay < 0](contracts/oracles/ChainlinkARBOracle.sol#L252)

contracts/oracles/ChainlinkARBOracle.sol#L249-L255


 - [ ] ID-10
[ChainlinkOracle.addTokenMaxDeviation(address,uint256)](contracts/oracles/ChainlinkOracle.sol#L230-L236) contains a tautology or contradiction:
	- [maxDeviation < 0](contracts/oracles/ChainlinkOracle.sol#L233)

contracts/oracles/ChainlinkOracle.sol#L230-L236


 - [ ] ID-11
[MockChainlinkARBOracleWithAdjusts.addTokenMaxDeviation(address,uint256)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L273-L279) contains a tautology or contradiction:
	- [maxDeviation < 0](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L276)

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L273-L279


 - [ ] ID-12
[ChainlinkOracle.addTokenMaxDelay(address,uint256)](contracts/oracles/ChainlinkOracle.sol#L217-L223) contains a tautology or contradiction:
	- [maxDelay < 0](contracts/oracles/ChainlinkOracle.sol#L220)

contracts/oracles/ChainlinkOracle.sol#L217-L223


 - [ ] ID-13
[ChainlinkARBOracle.addTokenMaxDeviation(address,uint256)](contracts/oracles/ChainlinkARBOracle.sol#L262-L268) contains a tautology or contradiction:
	- [maxDeviation < 0](contracts/oracles/ChainlinkARBOracle.sol#L265)

contracts/oracles/ChainlinkARBOracle.sol#L262-L268


## price-manipulation-medium
Impact: Medium
Confidence: Medium
 - [ ] ID-14
Potential price manipulation risk:
	- In function emergencyWithdraw
		-- [_userShareBalance = IERC20(address(self.vault)).balanceOf(msg.sender)](contracts/strategy/gmx/GMXEmergency.sol#L167) have potential price manipulated risk from _userShareBalance and call None which could influence variable:shareAmt

contracts/strategy/gmx/GMXEmergency.sol#L167


## public-mint-burn
Impact: Medium
Confidence: Medium
 - [ ] ID-15
public mint or burn found in [MockERC20.mint(address,uint256)](contracts/mocks/MockERC20.sol#L13-L15)
contracts/mocks/MockERC20.sol#L13-L15


 - [ ] ID-16
public mint or burn found in [MockWETH.burn(address,uint256)](contracts/mocks/MockWETH.sol#L14-L16)
contracts/mocks/MockWETH.sol#L14-L16


 - [ ] ID-17
public mint or burn found in [MockWETH.mint(address,uint256)](contracts/mocks/MockWETH.sol#L10-L12)
contracts/mocks/MockWETH.sol#L10-L12


 - [ ] ID-18
public mint or burn found in [MockERC20.burn(address,uint256)](contracts/mocks/MockERC20.sol#L17-L19)
contracts/mocks/MockERC20.sol#L17-L19


## uninitialized-local
Impact: Medium
Confidence: Medium
 - [ ] ID-19
minTokenAAmt is a member never initialized in [GMXRebalance.rebalanceRemove(GMXTypes.Store,GMXTypes.RebalanceRemoveParams)._rlp](contracts/strategy/gmx/GMXRebalance.sol#L173)

contracts/strategy/gmx/GMXRebalance.sol#L173


 - [ ] ID-20
[UniswapV2Router02._swap(uint256[],address[],address).i](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L230) is a local variable never initialized

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L230


 - [ ] ID-21
[UniswapV2Library.getAmountsOut(address,uint256,address[]).i](contracts/mocks/gmx/MockUniswapV2/libraries/UniswapV2Library.sol#L65) is a local variable never initialized

contracts/mocks/gmx/MockUniswapV2/libraries/UniswapV2Library.sol#L65


 - [ ] ID-22
minTokenAAmt is a member never initialized in [GMXWithdraw.withdraw(GMXTypes.Store,GMXTypes.WithdrawParams)._rlp](contracts/strategy/gmx/GMXWithdraw.sol#L103)

contracts/strategy/gmx/GMXWithdraw.sol#L103


 - [ ] ID-23
repayTokenAAmt is a member never initialized in [GMXEmergency.emergencyClose(GMXTypes.Store,uint256)._rp](contracts/strategy/gmx/GMXEmergency.sol#L118)

contracts/strategy/gmx/GMXEmergency.sol#L118


 - [ ] ID-24
borrowParams is a member never initialized in [GMXDeposit.deposit(GMXTypes.Store,GMXTypes.DepositParams,bool)._dc](contracts/strategy/gmx/GMXDeposit.sol#L86)

contracts/strategy/gmx/GMXDeposit.sol#L86


 - [ ] ID-25
[GMXDeposit.processDeposit(GMXTypes.Store).reason](contracts/strategy/gmx/GMXDeposit.sol#L182) is a local variable never initialized

contracts/strategy/gmx/GMXDeposit.sol#L182


 - [ ] ID-26
[GMXReader.additionalCapacity(GMXTypes.Store)._additionalCapacity](contracts/strategy/gmx/GMXReader.sol#L231) is a local variable never initialized

contracts/strategy/gmx/GMXReader.sol#L231


 - [ ] ID-27
[MockGMXOracleWithAdjusts._getTokenPriceMinMaxFormatted(address)._tokenPriceMinMaxFormatted](contracts/mocks/MockGMXOracleWithAdjusts.sol#L325) is a local variable never initialized

contracts/mocks/MockGMXOracleWithAdjusts.sol#L325


 - [ ] ID-28
[UniswapV2Router02._swapSupportingFeeOnTransferTokens(address[],address).i](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L351) is a local variable never initialized

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L351


 - [ ] ID-29
[GMXWithdraw.processWithdraw(GMXTypes.Store).reason](contracts/strategy/gmx/GMXWithdraw.sol#L206) is a local variable never initialized

contracts/strategy/gmx/GMXWithdraw.sol#L206


 - [ ] ID-30
[GMXRebalance.processRebalanceAdd(GMXTypes.Store).reason](contracts/strategy/gmx/GMXRebalance.sol#L114) is a local variable never initialized

contracts/strategy/gmx/GMXRebalance.sol#L114


 - [ ] ID-31
[GMXManager.calcBorrow(GMXTypes.Store,uint256)._borrowShortTokenAmt](contracts/strategy/gmx/GMXManager.sol#L83) is a local variable never initialized

contracts/strategy/gmx/GMXManager.sol#L83


 - [ ] ID-32
[GMXManager.calcBorrow(GMXTypes.Store,uint256)._borrowLongTokenAmt](contracts/strategy/gmx/GMXManager.sol#L82) is a local variable never initialized

contracts/strategy/gmx/GMXManager.sol#L82


 - [ ] ID-33
[GMXRebalance.processRebalanceRemove(GMXTypes.Store).reason](contracts/strategy/gmx/GMXRebalance.sol#L234) is a local variable never initialized

contracts/strategy/gmx/GMXRebalance.sol#L234


 - [ ] ID-34
repayParams is a member never initialized in [GMXWithdraw.withdraw(GMXTypes.Store,GMXTypes.WithdrawParams)._wc](contracts/strategy/gmx/GMXWithdraw.sol#L63)

contracts/strategy/gmx/GMXWithdraw.sol#L63


 - [ ] ID-35
minTokenAAmt is a member never initialized in [GMXDeposit.processDepositFailure(GMXTypes.Store,uint256,uint256)._rlp](contracts/strategy/gmx/GMXDeposit.sol#L235)

contracts/strategy/gmx/GMXDeposit.sol#L235


## unused-return
Impact: Medium
Confidence: Medium
 - [ ] ID-36
[GMXEmergency.emergencyResume(GMXTypes.Store)](contracts/strategy/gmx/GMXEmergency.sol#L72-L91)have ignores return value in [GMXManager.addLiquidity(self,_alp)](contracts/strategy/gmx/GMXEmergency.sol#L87-L90)

contracts/strategy/gmx/GMXEmergency.sol#L72-L91


 - [ ] ID-37
[GMXEmergency.emergencyClose(GMXTypes.Store,uint256)](contracts/strategy/gmx/GMXEmergency.sol#L111-L156)have ignores return value in [GMXManager.swapTokensForExactTokens(self,_sp)](contracts/strategy/gmx/GMXEmergency.sol#L141)

contracts/strategy/gmx/GMXEmergency.sol#L111-L156


 - [ ] ID-38
[GMXVault.constructor(string,string,GMXTypes.Store)](contracts/strategy/gmx/GMXVault.sol#L67-L132)have ignores return value in [_store.tokenA.approve(address(_store.router),type()(uint256).max)](contracts/strategy/gmx/GMXVault.sol#L118)

contracts/strategy/gmx/GMXVault.sol#L67-L132


 - [ ] ID-39
[TraderJoeSwap.swapExactTokensForTokens(ISwap.SwapParams)](contracts/swaps/TraderJoeSwap.sol#L60-L99)have ignores return value in [IERC20(sp.tokenIn).approve(address(router),sp.amountIn)](contracts/swaps/TraderJoeSwap.sol#L63)

contracts/swaps/TraderJoeSwap.sol#L60-L99


 - [ ] ID-40
[GMXTrove.constructor(address)](contracts/strategy/gmx/GMXTrove.sol#L32-L40)have ignores return value in [_store.tokenA.approve(address(vault),type()(uint256).max)](contracts/strategy/gmx/GMXTrove.sol#L38)

contracts/strategy/gmx/GMXTrove.sol#L32-L40


 - [ ] ID-41
[GMXEmergency.emergencyPause(GMXTypes.Store)](contracts/strategy/gmx/GMXEmergency.sol#L47-L66)have ignores return value in [GMXManager.removeLiquidity(self,_rlp)](contracts/strategy/gmx/GMXEmergency.sol#L58-L61)

contracts/strategy/gmx/GMXEmergency.sol#L47-L66


 - [ ] ID-42
[MockExchangeRouter.executeWithdrawal(address,address,address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L128-L147)have ignores return value in [uniswapRouter.removeLiquidity(tokenA,tokenB,IERC20(pair).balanceOf(address(this)),0,0,vault,block.timestamp + 1)](contracts/mocks/gmx/MockExchangeRouter.sol#L136-L144)

contracts/mocks/gmx/MockExchangeRouter.sol#L128-L147


 - [ ] ID-43
[GMXProcessWithdraw.processWithdraw(GMXTypes.Store)](contracts/strategy/gmx/GMXProcessWithdraw.sol#L24-L106)have ignores return value in [GMXManager.swapTokensForExactTokens(self,_sp)](contracts/strategy/gmx/GMXProcessWithdraw.sol#L52)

contracts/strategy/gmx/GMXProcessWithdraw.sol#L24-L106


 - [ ] ID-44
[UniswapSwap.swapExactTokensForTokens(ISwap.SwapParams)](contracts/swaps/UniswapSwap.sol#L60-L92)have ignores return value in [IERC20(sp.tokenIn).approve(address(router),sp.amountIn)](contracts/swaps/UniswapSwap.sol#L63)

contracts/swaps/UniswapSwap.sol#L60-L92


 - [ ] ID-45
[TraderJoeSwap.swapTokensForExactTokens(ISwap.SwapParams)](contracts/swaps/TraderJoeSwap.sol#L106-L151)have ignores return value in [IERC20(sp.tokenIn).approve(address(router),sp.amountIn)](contracts/swaps/TraderJoeSwap.sol#L113)

contracts/swaps/TraderJoeSwap.sol#L106-L151


 - [ ] ID-46
[UniswapV2Router02._addLiquidity(address,address,uint256,uint256,uint256,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L35-L62)have ignores return value in [IUniswapV2Factory(factory).createPair(tokenA,tokenB)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L45)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L35-L62


 - [ ] ID-47
[GMXWorker.swapTokensForExactTokens(GMXTypes.Store,ISwap.SwapParams)](contracts/strategy/gmx/GMXWorker.sol#L129-L136)have ignores return value in [IERC20(sp.tokenIn).approve(address(self.swapRouter),sp.amountIn)](contracts/strategy/gmx/GMXWorker.sol#L133)

contracts/strategy/gmx/GMXWorker.sol#L129-L136


 - [ ] ID-48
[UniswapSwap.swapTokensForExactTokens(ISwap.SwapParams)](contracts/swaps/UniswapSwap.sol#L99-L134)have ignores return value in [IERC20(sp.tokenIn).approve(address(router),sp.amountIn)](contracts/swaps/UniswapSwap.sol#L106)

contracts/swaps/UniswapSwap.sol#L99-L134


 - [ ] ID-49
[GMXCompound.compound(GMXTypes.Store,GMXTypes.CompoundParams)](contracts/strategy/gmx/GMXCompound.sol#L35-L107)have ignores return value in [GMXManager.swapExactTokensForTokens(self,_sp)](contracts/strategy/gmx/GMXCompound.sol#L72)

contracts/strategy/gmx/GMXCompound.sol#L35-L107


 - [ ] ID-50
[GMXWorker.swapExactTokensForTokens(GMXTypes.Store,ISwap.SwapParams)](contracts/strategy/gmx/GMXWorker.sol#L114-L121)have ignores return value in [IERC20(sp.tokenIn).approve(address(self.swapRouter),sp.amountIn)](contracts/strategy/gmx/GMXWorker.sol#L118)

contracts/strategy/gmx/GMXWorker.sol#L114-L121


 - [ ] ID-51
[MockExchangeRouter.executeDeposit(address,address,address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L70-L88)have ignores return value in [uniswapRouter.addLiquidity(tokenA,tokenB,IERC20(tokenA).balanceOf(address(this)),IERC20(tokenB).balanceOf(address(this)),vault,block.timestamp + 1)](contracts/mocks/gmx/MockExchangeRouter.sol#L78-L85)

contracts/mocks/gmx/MockExchangeRouter.sol#L70-L88


 - [ ] ID-52
[MockExchangeRouter._swapForDeposit(address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L249-L287)have ignores return value in [uniswapRouter.swapExactTokensForTokens(optimalSwapAmount,0,swapPathForOptimalDeposit,address(this),block.timestamp)](contracts/mocks/gmx/MockExchangeRouter.sol#L279-L285)

contracts/mocks/gmx/MockExchangeRouter.sol#L249-L287


 - [ ] ID-53
[GMXDeposit.processDepositFailureLiquidityWithdrawal(GMXTypes.Store)](contracts/strategy/gmx/GMXDeposit.sol#L282-L352)have ignores return value in [GMXManager.swapTokensForExactTokens(self,_sp)](contracts/strategy/gmx/GMXDeposit.sol#L317)

contracts/strategy/gmx/GMXDeposit.sol#L282-L352


## centralized-risk-low
Impact: Low
Confidence: High
 - [ ] ID-54
	- [GMXVault.updateCallback(address)](contracts/strategy/gmx/GMXVault.sol#L605-L608)

contracts/strategy/gmx/GMXVault.sol#L605-L608


 - [ ] ID-55
	- [GMXVault.updateFeePerSecond(uint256)](contracts/strategy/gmx/GMXVault.sol#L615-L618)

contracts/strategy/gmx/GMXVault.sol#L615-L618


 - [ ] ID-56
	- [GMXVault.updateParameterLimits(uint256,uint256,uint256,int256,int256)](contracts/strategy/gmx/GMXVault.sol#L629-L649)

contracts/strategy/gmx/GMXVault.sol#L629-L649


 - [ ] ID-57
	- [GMXVault.updateMinExecutionFee(uint256)](contracts/strategy/gmx/GMXVault.sol#L666-L669)

contracts/strategy/gmx/GMXVault.sol#L666-L669


 - [ ] ID-58
	- [GMXVault.updateMinSlippage(uint256)](contracts/strategy/gmx/GMXVault.sol#L656-L659)

contracts/strategy/gmx/GMXVault.sol#L656-L659


 - [ ] ID-59
	- [GMXVault.updateTreasury(address)](contracts/strategy/gmx/GMXVault.sol#L575-L578)

contracts/strategy/gmx/GMXVault.sol#L575-L578


 - [ ] ID-60
	- [GMXVault.updateSwapRouter(address)](contracts/strategy/gmx/GMXVault.sol#L585-L588)

contracts/strategy/gmx/GMXVault.sol#L585-L588


 - [ ] ID-61
	- [GMXVault.updateTrove(address)](contracts/strategy/gmx/GMXVault.sol#L595-L598)

contracts/strategy/gmx/GMXVault.sol#L595-L598


## pess-event-setter
Impact: Low
Confidence: Medium
 - [ ] ID-62
Setter function [GMXCompound.compound(GMXTypes.Store,GMXTypes.CompoundParams)](contracts/strategy/gmx/GMXCompound.sol#L35-L107) does not emit an event

contracts/strategy/gmx/GMXCompound.sol#L35-L107


 - [ ] ID-63
Setter function [GMXVault.processWithdrawFailureLiquidityAdded()](contracts/strategy/gmx/GMXVault.sol#L433-L435) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L433-L435


 - [ ] ID-64
Setter function [UniswapV2Router02.swapExactTokensForTokens(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L241-L255) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L241-L255


 - [ ] ID-65
Setter function [GMXCallback.afterDepositExecution(bytes32,IDeposit.Props,IEvent.Props)](contracts/strategy/gmx/GMXCallback.sol#L57-L89) does not emit an event

contracts/strategy/gmx/GMXCallback.sol#L57-L89


 - [ ] ID-66
Setter function [UniswapV2Router02.removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L131-L154) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L131-L154


 - [ ] ID-67
Setter function [MockChainlinkARBOracleWithAdjusts.changeAdjust(uint256)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L18-L22) does not emit an event

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L18-L22


 - [ ] ID-68
Setter function [GMXCallback.afterDepositCancellation(bytes32,IDeposit.Props,IEvent.Props)](contracts/strategy/gmx/GMXCallback.sol#L96-L115) does not emit an event

contracts/strategy/gmx/GMXCallback.sol#L96-L115


 - [ ] ID-69
Setter function [MockLendingVault.approveBorrower(address)](contracts/mocks/MockLendingVault.sol#L561-L565) does not emit an event

contracts/mocks/MockLendingVault.sol#L561-L565


 - [ ] ID-70
Setter function [GMXVault.processDepositCancellation()](contracts/strategy/gmx/GMXVault.sol#L371-L373) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L371-L373


 - [ ] ID-71
Setter function [GMXRebalance.rebalanceRemove(GMXTypes.Store,GMXTypes.RebalanceRemoveParams)](contracts/strategy/gmx/GMXRebalance.sol#L149-L210) does not emit an event

contracts/strategy/gmx/GMXRebalance.sol#L149-L210


 - [ ] ID-72
Setter function [MockLendingVault._burnShares(uint256)](contracts/mocks/MockLendingVault.sol#L432-L441) does not emit an event

contracts/mocks/MockLendingVault.sol#L432-L441


 - [ ] ID-73
Setter function [MockChainlinkARBOracleWithAdjusts.emergencyResume()](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L291-L293) does not emit an event

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L291-L293


 - [ ] ID-74
Setter function [GMXVault.processDepositFailureLiquidityWithdrawal()](contracts/strategy/gmx/GMXVault.sol#L393-L395) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L393-L395


 - [ ] ID-75
Setter function [MockLendingVault.revokeBorrower(address)](contracts/mocks/MockLendingVault.sol#L571-L575) does not emit an event

contracts/mocks/MockLendingVault.sol#L571-L575


 - [ ] ID-76
Setter function [GMXVault.rebalanceRemove(GMXTypes.RebalanceRemoveParams)](contracts/strategy/gmx/GMXVault.sol#L471-L475) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L471-L475


 - [ ] ID-77
Setter function [UniswapV2Router02.addLiquidity(address,address,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L65-L80) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L65-L80


 - [ ] ID-78
Setter function [MockLendingVault._updateVaultWithInterestsAndTimestamp(uint256)](contracts/mocks/MockLendingVault.sol#L447-L454) does not emit an event

contracts/mocks/MockLendingVault.sol#L447-L454


 - [ ] ID-79
Setter function [MockGMXOracleWithAdjusts.constructor(address,ISyntheticReader,IChainlinkOracle)](contracts/mocks/MockGMXOracleWithAdjusts.sol#L46-L58) does not emit an event

contracts/mocks/MockGMXOracleWithAdjusts.sol#L46-L58


 - [ ] ID-80
Setter function [GMXVault.processRebalanceAdd()](contracts/strategy/gmx/GMXVault.sol#L453-L455) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L453-L455


 - [ ] ID-81
Setter function [GMXVault.compound(GMXTypes.CompoundParams)](contracts/strategy/gmx/GMXVault.sol#L501-L503) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L501-L503


 - [ ] ID-82
Setter function [LendingVault.withdrawReserve(uint256)](contracts/lending/LendingVault.sol#L371-L382) does not emit an event

contracts/lending/LendingVault.sol#L371-L382


 - [ ] ID-83
Setter function [GMXVault.processWithdrawFailure(uint256,uint256)](contracts/strategy/gmx/GMXVault.sol#L421-L426) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L421-L426


 - [ ] ID-84
Setter function [GMXVault.processWithdrawCancellation()](contracts/strategy/gmx/GMXVault.sol#L411-L413) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L411-L413


 - [ ] ID-85
Setter function [LendingVault.onlyKeeper()](contracts/lending/LendingVault.sol#L97-L100) does not emit an event

contracts/lending/LendingVault.sol#L97-L100


 - [ ] ID-86
Setter function [UniswapV2Router02.removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L170-L183) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L170-L183


 - [ ] ID-87
Setter function [MockExchangeRouter.createWithdrawal(IExchangeRouter.CreateWithdrawalParams)](contracts/mocks/gmx/MockExchangeRouter.sol#L120-L126) does not emit an event

contracts/mocks/gmx/MockExchangeRouter.sol#L120-L126


 - [ ] ID-88
Setter function [GMXVault.processRebalanceAddCancellation()](contracts/strategy/gmx/GMXVault.sol#L462-L464) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L462-L464


 - [ ] ID-89
Setter function [ChainlinkARBOracle.addTokenMaxDelay(address,uint256)](contracts/oracles/ChainlinkARBOracle.sol#L249-L255) does not emit an event

contracts/oracles/ChainlinkARBOracle.sol#L249-L255


 - [ ] ID-90
Setter function [MockLendingVault.constructor(string,string,IERC20,bool,uint256,uint256,address,ILendingVault.InterestRate,ILendingVault.InterestRate)](contracts/mocks/MockLendingVault.sol#L117-L149) does not emit an event

contracts/mocks/MockLendingVault.sol#L117-L149


 - [ ] ID-91
Setter function [LendingVault._onlyBorrower()](contracts/lending/LendingVault.sol#L389-L391) does not emit an event

contracts/lending/LendingVault.sol#L389-L391


 - [ ] ID-92
Setter function [MockWETH.withdraw(uint256)](contracts/mocks/MockWETH.sol#L22-L26) does not emit an event

contracts/mocks/MockWETH.sol#L22-L26


 - [ ] ID-93
Setter function [MockChainlinkARBOracleWithAdjusts.addTokenMaxDelay(address,uint256)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L260-L266) does not emit an event

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L260-L266


 - [ ] ID-94
Setter function [MockSwapper.swap(ISwap.SwapParams)](contracts/mocks/MockSwapper.sol#L31-L53) does not emit an event

contracts/mocks/MockSwapper.sol#L31-L53


 - [ ] ID-95
Setter function [TraderJoeSwap.swapTokensForExactTokens(ISwap.SwapParams)](contracts/swaps/TraderJoeSwap.sol#L106-L151) does not emit an event

contracts/swaps/TraderJoeSwap.sol#L106-L151


 - [ ] ID-96
Setter function [GMXVault.mint(address,uint256)](contracts/strategy/gmx/GMXVault.sol#L677-L679) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L677-L679


 - [ ] ID-97
Setter function [UniswapV2Factory.setFeeTo(address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L41-L44) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L41-L44


 - [ ] ID-98
Setter function [MockSyntheticReader.setMarketTokenPrice(uint256)](contracts/mocks/gmx/MockSyntheticReader.sol#L20-L22) does not emit an event

contracts/mocks/gmx/MockSyntheticReader.sol#L20-L22


 - [ ] ID-99
Setter function [GMXRebalance.rebalanceAdd(GMXTypes.Store,GMXTypes.RebalanceAddParams)](contracts/strategy/gmx/GMXRebalance.sol#L35-L96) does not emit an event

contracts/strategy/gmx/GMXRebalance.sol#L35-L96


 - [ ] ID-100
Setter function [MockExchangeRouter.swapExactTokensForTokens(ISwap.SwapParams)](contracts/mocks/gmx/MockExchangeRouter.sol#L180-L199) does not emit an event

contracts/mocks/gmx/MockExchangeRouter.sol#L180-L199


 - [ ] ID-101
Setter function [TraderJoeSwap.swapExactTokensForTokens(ISwap.SwapParams)](contracts/swaps/TraderJoeSwap.sol#L60-L99) does not emit an event

contracts/swaps/TraderJoeSwap.sol#L60-L99


 - [ ] ID-102
Setter function [MockExchangeRouter.swapTokensForExactTokens(ISwap.SwapParams)](contracts/mocks/gmx/MockExchangeRouter.sol#L201-L220) does not emit an event

contracts/mocks/gmx/MockExchangeRouter.sol#L201-L220


 - [ ] ID-103
Setter function [UniswapV2Pair.constructor()](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L63-L65) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L63-L65


 - [ ] ID-104
Setter function [GMXVault.emergencyResume()](contracts/strategy/gmx/GMXVault.sol#L536-L538) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L536-L538


 - [ ] ID-105
Setter function [GMXVault.constructor(string,string,GMXTypes.Store)](contracts/strategy/gmx/GMXVault.sol#L67-L132) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L67-L132


 - [ ] ID-106
Setter function [GMXCallback.afterWithdrawalCancellation(bytes32,IWithdrawal.Props,IEvent.Props)](contracts/strategy/gmx/GMXCallback.sol#L152-L168) does not emit an event

contracts/strategy/gmx/GMXCallback.sol#L152-L168


 - [ ] ID-107
Setter function [UniswapV2Router02.swapTokensForExactETH(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L286-L305) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L286-L305


 - [ ] ID-108
Setter function [LendingVault.approveBorrower(address)](contracts/lending/LendingVault.sol#L554-L558) does not emit an event

contracts/lending/LendingVault.sol#L554-L558


 - [ ] ID-109
Setter function [UniswapV2Router02.receive()](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L30-L32) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L30-L32


 - [ ] ID-110
Setter function [MockChainlinkOracle.set(address,uint256,uint8)](contracts/mocks/MockChainlinkOracle.sol#L28-L31) does not emit an event

contracts/mocks/MockChainlinkOracle.sol#L28-L31


 - [ ] ID-111
Setter function [MockLendingVault.updateKeeper(address,bool)](contracts/mocks/MockLendingVault.sol#L582-L586) does not emit an event

contracts/mocks/MockLendingVault.sol#L582-L586


 - [ ] ID-112
Setter function [UniswapV2Router02.swapETHForExactTokens(uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L326-L346) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L326-L346


 - [ ] ID-113
Setter function [MockLendingVault.onlyKeeper()](contracts/mocks/MockLendingVault.sol#L99-L102) does not emit an event

contracts/mocks/MockLendingVault.sol#L99-L102


 - [ ] ID-114
Setter function [ChainlinkOracle.addTokenMaxDelay(address,uint256)](contracts/oracles/ChainlinkOracle.sol#L217-L223) does not emit an event

contracts/oracles/ChainlinkOracle.sol#L217-L223


 - [ ] ID-115
Setter function [LendingVault.updateTreasury(address)](contracts/lending/LendingVault.sol#L668-L672) does not emit an event

contracts/lending/LendingVault.sol#L668-L672


 - [ ] ID-116
Setter function [GMXChecks.beforeEmergencyWithdrawChecks(GMXTypes.Store,uint256)](contracts/strategy/gmx/GMXChecks.sol#L415-L427) does not emit an event

contracts/strategy/gmx/GMXChecks.sol#L415-L427


 - [ ] ID-117
Setter function [MockSwapper.setEthPrice(uint256)](contracts/mocks/MockSwapper.sol#L55-L57) does not emit an event

contracts/mocks/MockSwapper.sol#L55-L57


 - [ ] ID-118
Setter function [UniswapV2Pair.initialize(address,address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L68-L72) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L68-L72


 - [ ] ID-119
Setter function [UniswapV2Router02.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L210-L225) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L210-L225


 - [ ] ID-120
Setter function [MockExchangeRouter.createDeposit(IExchangeRouter.CreateDepositParams)](contracts/mocks/gmx/MockExchangeRouter.sol#L65-L68) does not emit an event

contracts/mocks/gmx/MockExchangeRouter.sol#L65-L68


 - [ ] ID-121
Setter function [GMXCallback.onlyController()](contracts/strategy/gmx/GMXCallback.sol#L31-L37) does not emit an event

contracts/strategy/gmx/GMXCallback.sol#L31-L37


 - [ ] ID-122
Setter function [MockStrategyVault.deposit(uint256)](contracts/mocks/MockStrategyVault.sol#L13-L15) does not emit an event

contracts/mocks/MockStrategyVault.sol#L13-L15


 - [ ] ID-123
Setter function [LendingVault.onlyBorrower()](contracts/lending/LendingVault.sol#L89-L92) does not emit an event

contracts/lending/LendingVault.sol#L89-L92


 - [ ] ID-124
Setter function [ChainlinkOracle.constructor()](contracts/oracles/ChainlinkOracle.sol#L39) does not emit an event

contracts/oracles/ChainlinkOracle.sol#L39


 - [ ] ID-125
Setter function [GMXVault.processDepositFailure(uint256,uint256)](contracts/strategy/gmx/GMXVault.sol#L381-L386) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L381-L386


 - [ ] ID-126
Setter function [UniswapSwap.swapExactTokensForTokens(ISwap.SwapParams)](contracts/swaps/UniswapSwap.sol#L60-L92) does not emit an event

contracts/swaps/UniswapSwap.sol#L60-L92


 - [ ] ID-127
Setter function [LendingVault._onlyKeeper()](contracts/lending/LendingVault.sol#L396-L398) does not emit an event

contracts/lending/LendingVault.sol#L396-L398


 - [ ] ID-128
Setter function [GMXVault.onlyVault()](contracts/strategy/gmx/GMXVault.sol#L48-L51) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L48-L51


 - [ ] ID-129
Setter function [ChainlinkARBOracle.constructor(address)](contracts/oracles/ChainlinkARBOracle.sol#L48-L52) does not emit an event

contracts/oracles/ChainlinkARBOracle.sol#L48-L52


 - [ ] ID-130
Setter function [GMXVault.emergencyClose(uint256)](contracts/strategy/gmx/GMXVault.sol#L555-L557) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L555-L557


 - [ ] ID-131
Setter function [LendingVault._updateVaultWithInterestsAndTimestamp(uint256)](contracts/lending/LendingVault.sol#L440-L447) does not emit an event

contracts/lending/LendingVault.sol#L440-L447


 - [ ] ID-132
Setter function [MockGMXOracleWithAdjusts.changeAdjust(uint256,uint256)](contracts/mocks/MockGMXOracleWithAdjusts.sol#L18-L24) does not emit an event

contracts/mocks/MockGMXOracleWithAdjusts.sol#L18-L24


 - [ ] ID-133
Setter function [ChainlinkOracle.emergencyResume()](contracts/oracles/ChainlinkOracle.sol#L248-L250) does not emit an event

contracts/oracles/ChainlinkOracle.sol#L248-L250


 - [ ] ID-134
Setter function [ChainlinkARBOracle.addTokenPriceFeed(address,address)](contracts/oracles/ChainlinkARBOracle.sol#L236-L242) does not emit an event

contracts/oracles/ChainlinkARBOracle.sol#L236-L242


 - [ ] ID-135
Setter function [MockLendingVault.updateTreasury(address)](contracts/mocks/MockLendingVault.sol#L675-L679) does not emit an event

contracts/mocks/MockLendingVault.sol#L675-L679


 - [ ] ID-136
Setter function [UniswapV2Router02.removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L113-L129) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L113-L129


 - [ ] ID-137
Setter function [GMXVault.processEmergencyResume()](contracts/strategy/gmx/GMXVault.sol#L545-L547) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L545-L547


 - [ ] ID-138
Setter function [MockChainlinkARBOracleWithAdjusts.constructor(address)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L58-L62) does not emit an event

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L58-L62


 - [ ] ID-139
Setter function [UniswapV2Router02.swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L368-L385) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L368-L385


 - [ ] ID-140
Setter function [MockChainlinkARBOracleWithAdjusts.emergencyPause()](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L284-L286) does not emit an event

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L284-L286


 - [ ] ID-141
Setter function [GMXCallback.afterWithdrawalExecution(bytes32,IWithdrawal.Props,IEvent.Props)](contracts/strategy/gmx/GMXCallback.sol#L122-L145) does not emit an event

contracts/strategy/gmx/GMXCallback.sol#L122-L145


 - [ ] ID-142
Setter function [UniswapV2Router02.removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L155-L169) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L155-L169


 - [ ] ID-143
Setter function [UniswapV2Pair.lock()](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L33-L38) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L33-L38


 - [ ] ID-144
Setter function [MockStrategyVault.withdraw(uint256)](contracts/mocks/MockStrategyVault.sol#L17-L19) does not emit an event

contracts/mocks/MockStrategyVault.sol#L17-L19


 - [ ] ID-145
Setter function [LendingVault._mintShares(uint256)](contracts/lending/LendingVault.sol#L405-L418) does not emit an event

contracts/lending/LendingVault.sol#L405-L418


 - [ ] ID-146
Setter function [UniswapV2Router02.removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L186-L209) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L186-L209


 - [ ] ID-147
Setter function [GMXVault.processCompoundCancellation()](contracts/strategy/gmx/GMXVault.sol#L519-L521) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L519-L521


 - [ ] ID-148
Setter function [GMXVault.receive()](contracts/strategy/gmx/GMXVault.sol#L697-L702) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L697-L702


 - [ ] ID-149
Setter function [MockLendingVault._mintShares(uint256)](contracts/mocks/MockLendingVault.sol#L412-L425) does not emit an event

contracts/mocks/MockLendingVault.sol#L412-L425


 - [ ] ID-150
Setter function [MockLendingVault.withdrawReserve(uint256)](contracts/mocks/MockLendingVault.sol#L378-L389) does not emit an event

contracts/mocks/MockLendingVault.sol#L378-L389


 - [ ] ID-151
Setter function [MockLendingVault.mockSetDebt(address,uint256)](contracts/mocks/MockLendingVault.sol#L152-L154) does not emit an event

contracts/mocks/MockLendingVault.sol#L152-L154


 - [ ] ID-152
Setter function [GMXVault.processDeposit()](contracts/strategy/gmx/GMXVault.sol#L362-L364) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L362-L364


 - [ ] ID-153
Setter function [LendingVault._burnShares(uint256)](contracts/lending/LendingVault.sol#L425-L434) does not emit an event

contracts/lending/LendingVault.sol#L425-L434


 - [ ] ID-154
Setter function [LendingVault.constructor(string,string,IERC20,bool,uint256,uint256,address,ILendingVault.InterestRate,ILendingVault.InterestRate)](contracts/lending/LendingVault.sol#L115-L147) does not emit an event

contracts/lending/LendingVault.sol#L115-L147


 - [ ] ID-155
Setter function [UniswapV2Factory.setFeeToSetter(address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L46-L49) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L46-L49


 - [ ] ID-156
Setter function [GMXVault.burn(address,uint256)](contracts/strategy/gmx/GMXVault.sol#L687-L689) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L687-L689


 - [ ] ID-157
Setter function [MockLendingVault._onlyBorrower()](contracts/mocks/MockLendingVault.sol#L396-L398) does not emit an event

contracts/mocks/MockLendingVault.sol#L396-L398


 - [ ] ID-158
Setter function [GMXVault.processRebalanceRemoveCancellation()](contracts/strategy/gmx/GMXVault.sol#L491-L493) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L491-L493


 - [ ] ID-159
Setter function [MockAggregatorV3.setPreviousRoundData(MockAggregatorV3.PreviousRoundData)](contracts/mocks/chainlink/MockAggregatorV3.sol#L32-L38) does not emit an event

contracts/mocks/chainlink/MockAggregatorV3.sol#L32-L38


 - [ ] ID-160
Setter function [GMXVault._onlyKeeper()](contracts/strategy/gmx/GMXVault.sol#L351-L353) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L351-L353


 - [ ] ID-161
Setter function [MockLendingVault._onlyKeeper()](contracts/mocks/MockLendingVault.sol#L403-L405) does not emit an event

contracts/mocks/MockLendingVault.sol#L403-L405


 - [ ] ID-162
Setter function [MockWETH.deposit()](contracts/mocks/MockWETH.sol#L18-L20) does not emit an event

contracts/mocks/MockWETH.sol#L18-L20


 - [ ] ID-163
Setter function [GMXVault.mintFee()](contracts/strategy/gmx/GMXVault.sol#L334-L337) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L334-L337


 - [ ] ID-164
Setter function [LendingVault.revokeBorrower(address)](contracts/lending/LendingVault.sol#L564-L568) does not emit an event

contracts/lending/LendingVault.sol#L564-L568


 - [ ] ID-165
Setter function [GMXVault.rebalanceAdd(GMXTypes.RebalanceAddParams)](contracts/strategy/gmx/GMXVault.sol#L442-L446) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L442-L446


 - [ ] ID-166
Setter function [UniswapV2Pair._mintFee(uint112,uint112)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L90-L108) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L90-L108


 - [ ] ID-167
Setter function [MockChainlinkARBOracleWithAdjusts.addTokenMaxDeviation(address,uint256)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L273-L279) does not emit an event

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L273-L279


 - [ ] ID-168
Setter function [GMXVault.emergencyPause()](contracts/strategy/gmx/GMXVault.sol#L528-L530) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L528-L530


 - [ ] ID-169
Setter function [MockAggregatorV3.setCurrentRoundData(MockAggregatorV3.CurrentRoundData)](contracts/mocks/chainlink/MockAggregatorV3.sol#L40-L46) does not emit an event

contracts/mocks/chainlink/MockAggregatorV3.sol#L40-L46


 - [ ] ID-170
Setter function [MockExchangeRouter.sendTokens(address,address,uint256)](contracts/mocks/gmx/MockExchangeRouter.sol#L237-L243) does not emit an event

contracts/mocks/gmx/MockExchangeRouter.sol#L237-L243


 - [ ] ID-171
Setter function [GMXVault.processCompound()](contracts/strategy/gmx/GMXVault.sol#L510-L512) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L510-L512


 - [ ] ID-172
Setter function [GMXEmergency.emergencyResume(GMXTypes.Store)](contracts/strategy/gmx/GMXEmergency.sol#L72-L91) does not emit an event

contracts/strategy/gmx/GMXEmergency.sol#L72-L91


 - [ ] ID-173
Setter function [UniswapSwap.swapTokensForExactTokens(ISwap.SwapParams)](contracts/swaps/UniswapSwap.sol#L99-L134) does not emit an event

contracts/swaps/UniswapSwap.sol#L99-L134


 - [ ] ID-174
Setter function [ChainlinkARBOracle.addTokenMaxDeviation(address,uint256)](contracts/oracles/ChainlinkARBOracle.sol#L262-L268) does not emit an event

contracts/oracles/ChainlinkARBOracle.sol#L262-L268


 - [ ] ID-175
Setter function [GMXVault.processRebalanceRemove()](contracts/strategy/gmx/GMXVault.sol#L482-L484) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L482-L484


 - [ ] ID-176
Setter function [MockLendingVault.onlyBorrower()](contracts/mocks/MockLendingVault.sol#L91-L94) does not emit an event

contracts/mocks/MockLendingVault.sol#L91-L94


 - [ ] ID-177
Setter function [GMXVault.onlyKeeper()](contracts/strategy/gmx/GMXVault.sol#L54-L57) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L54-L57


 - [ ] ID-178
Setter function [UniswapV2Router02.swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L409-L433) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L409-L433


 - [ ] ID-179
Setter function [ChainlinkOracle.addTokenPriceFeed(address,address)](contracts/oracles/ChainlinkOracle.sol#L204-L210) does not emit an event

contracts/oracles/ChainlinkOracle.sol#L204-L210


 - [ ] ID-180
Setter function [ChainlinkARBOracle.emergencyPause()](contracts/oracles/ChainlinkARBOracle.sol#L273-L275) does not emit an event

contracts/oracles/ChainlinkARBOracle.sol#L273-L275


 - [ ] ID-181
Setter function [UniswapV2Router02.swapExactTokensForETH(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L306-L325) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L306-L325


 - [ ] ID-182
Setter function [UniswapV2Router02.swapTokensForExactTokens(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L256-L270) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L256-L270


 - [ ] ID-183
Setter function [MockChainlinkOracle.constructor()](contracts/mocks/MockChainlinkOracle.sol#L25) does not emit an event

contracts/mocks/MockChainlinkOracle.sol#L25


 - [ ] ID-184
Setter function [GMXVault._onlyVault()](contracts/strategy/gmx/GMXVault.sol#L344-L346) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L344-L346


 - [ ] ID-185
Setter function [LendingVault.updateKeeper(address,bool)](contracts/lending/LendingVault.sol#L575-L579) does not emit an event

contracts/lending/LendingVault.sol#L575-L579


 - [ ] ID-186
Setter function [UniswapV2Router02.addLiquidityETH(address,uint256,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L82-L110) does not emit an event

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L82-L110


 - [ ] ID-187
Setter function [ChainlinkARBOracle.emergencyResume()](contracts/oracles/ChainlinkARBOracle.sol#L280-L282) does not emit an event

contracts/oracles/ChainlinkARBOracle.sol#L280-L282


 - [ ] ID-188
Setter function [ChainlinkOracle.emergencyPause()](contracts/oracles/ChainlinkOracle.sol#L241-L243) does not emit an event

contracts/oracles/ChainlinkOracle.sol#L241-L243


 - [ ] ID-189
Setter function [MockChainlinkARBOracleWithAdjusts.addTokenPriceFeed(address,address)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L247-L253) does not emit an event

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L247-L253


 - [ ] ID-190
Setter function [GMXVault.processWithdraw()](contracts/strategy/gmx/GMXVault.sol#L402-L404) does not emit an event

contracts/strategy/gmx/GMXVault.sol#L402-L404


 - [ ] ID-191
Setter function [TraderJoeSwap.constructor(ILBRouter,IChainlinkOracle)](contracts/swaps/TraderJoeSwap.sol#L43-L51) does not emit an event

contracts/swaps/TraderJoeSwap.sol#L43-L51


 - [ ] ID-192
Setter function [UniswapSwap.constructor(ISwapRouter,IChainlinkOracle)](contracts/swaps/UniswapSwap.sol#L43-L51) does not emit an event

contracts/swaps/UniswapSwap.sol#L43-L51


 - [ ] ID-193
Setter function [ChainlinkOracle.addTokenMaxDeviation(address,uint256)](contracts/oracles/ChainlinkOracle.sol#L230-L236) does not emit an event

contracts/oracles/ChainlinkOracle.sol#L230-L236


## initialize-permission
Impact: Low
Confidence: Medium
 - [ ] ID-194
Condition variable is not initialized found in [UniswapV2Pair.initialize(address,address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L68-L72)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L68-L72


## missing-zero-check
Impact: Low
Confidence: Medium
 - [ ] ID-195
variable lacks a zero-check on 		- [MockGMXOracleWithAdjusts.getAmountsOut(address,address,address,address,address,uint256)](contracts/mocks/MockGMXOracleWithAdjusts.sol#L73-L116)

contracts/mocks/MockGMXOracleWithAdjusts.sol#L73-L116


 - [ ] ID-196
variable lacks a zero-check on 		- [GMXReader.convertToUsdValue(GMXTypes.Store,address,uint256)](contracts/strategy/gmx/GMXReader.sol#L62-L70)

contracts/strategy/gmx/GMXReader.sol#L62-L70


 - [ ] ID-197
variable lacks a zero-check on 		- [UniswapV2Factory.constructor(address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L16-L18)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L16-L18


 - [ ] ID-198
variable lacks a zero-check on 		- [UniswapV2Router02.constructor(address,address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L25-L28)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L25-L28


 - [ ] ID-199
variable lacks a zero-check on 		- [GMXEmergency.emergencyClose(GMXTypes.Store,uint256)](contracts/strategy/gmx/GMXEmergency.sol#L111-L156)

contracts/strategy/gmx/GMXEmergency.sol#L111-L156


 - [ ] ID-200
variable lacks a zero-check on 		- [MockChainlinkOracle.set(address,uint256,uint8)](contracts/mocks/MockChainlinkOracle.sol#L28-L31)

contracts/mocks/MockChainlinkOracle.sol#L28-L31


 - [ ] ID-201
variable lacks a zero-check on 		- [MockExchangeRouter.executeMockWithdrawal(address,address,uint256,uint256,uint256,uint256,address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L149-L166)

contracts/mocks/gmx/MockExchangeRouter.sol#L149-L166


 - [ ] ID-202
variable lacks a zero-check on 		- [MockExchangeRouter.sendTokens(address,address,uint256)](contracts/mocks/gmx/MockExchangeRouter.sol#L237-L243)

contracts/mocks/gmx/MockExchangeRouter.sol#L237-L243


 - [ ] ID-203
variable lacks a zero-check on 		- [UniswapV2Pair.skim(address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L203-L208)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L203-L208


 - [ ] ID-204
variable lacks a zero-check on 		- [MockGMXOracle.getLpTokenReserves(address,address,address,address)](contracts/mocks/gmx/MockGMXOracle.sol#L39-L49)

contracts/mocks/gmx/MockGMXOracle.sol#L39-L49


 - [ ] ID-205
variable lacks a zero-check on 		- [GMXProcessWithdraw.processWithdraw(GMXTypes.Store)](contracts/strategy/gmx/GMXProcessWithdraw.sol#L24-L106)

contracts/strategy/gmx/GMXProcessWithdraw.sol#L24-L106


 - [ ] ID-206
variable lacks a zero-check on 		- [UniswapV2Router02.removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L131-L154)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L131-L154


 - [ ] ID-207
variable lacks a zero-check on 		- [MockUniswapV2Oracle.lpToken(address,address)](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L50-L55)

contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L50-L55


 - [ ] ID-208
variable lacks a zero-check on 		- [MockExchangeRouter.getMarketTokenInfo(address,address,address,address,bool,bool)](contracts/mocks/gmx/MockExchangeRouter.sol#L223-L235)

contracts/mocks/gmx/MockExchangeRouter.sol#L223-L235


 - [ ] ID-209
variable lacks a zero-check on 		- [UniswapV2Router02.addLiquidityETH(address,uint256,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L82-L110)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L82-L110


 - [ ] ID-210
variable lacks a zero-check on 		- [LendingVault.emergencyRepay(uint256,address)](contracts/lending/LendingVault.sol#L585-L614)

contracts/lending/LendingVault.sol#L585-L614


 - [ ] ID-211
variable lacks a zero-check on 		- [UniswapV2Router02.swapExactETHForTokensSupportingFeeOnTransferTokens(uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L386-L408)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L386-L408


 - [ ] ID-212
variable lacks a zero-check on 		- [MockExchangeRouter.executeWithdrawal(address,address,address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L128-L147)

contracts/mocks/gmx/MockExchangeRouter.sol#L128-L147


 - [ ] ID-213
variable lacks a zero-check on 		- [UniswapV2Router02.swapTokensForExactETH(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L286-L305)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L286-L305


 - [ ] ID-214
variable lacks a zero-check on 		- [MockLendingVault.emergencyRepay(uint256,address)](contracts/mocks/MockLendingVault.sol#L592-L621)

contracts/mocks/MockLendingVault.sol#L592-L621


 - [ ] ID-215
variable lacks a zero-check on 		- [MockGMXOracle.getLpTokenValue(address,address,address,address,bool,bool)](contracts/mocks/gmx/MockGMXOracle.sol#L19-L37)

contracts/mocks/gmx/MockGMXOracle.sol#L19-L37


 - [ ] ID-216
variable lacks a zero-check on 		- [UniswapV2Router02.removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L155-L169)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L155-L169


 - [ ] ID-217
variable lacks a zero-check on 		- [GMXOracle.getAmountsOut(address,address,address,address,address,uint256)](contracts/oracles/GMXOracle.sol#L61-L104)

contracts/oracles/GMXOracle.sol#L61-L104


 - [ ] ID-218
variable lacks a zero-check on 		- [MockLendingVault.mockSetDebt(address,uint256)](contracts/mocks/MockLendingVault.sol#L152-L154)

contracts/mocks/MockLendingVault.sol#L152-L154


 - [ ] ID-219
variable lacks a zero-check on 		- [UniswapV2Router02.removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L113-L129)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L113-L129


 - [ ] ID-220
variable lacks a zero-check on 		- [UniswapV2Pair.initialize(address,address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L68-L72)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L68-L72


 - [ ] ID-221
variable lacks a zero-check on 		- [UniswapV2Router02.removeLiquidityETHWithPermitSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L210-L225)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L210-L225


 - [ ] ID-222
variable lacks a zero-check on 		- [UniswapV2ERC20.permit(address,address,uint256,uint256,uint8,bytes32,bytes32)](contracts/mocks/gmx/MockUniswapV2/UniswapV2ERC20.sol#L76-L88)

contracts/mocks/gmx/MockUniswapV2/UniswapV2ERC20.sol#L76-L88


 - [ ] ID-223
variable lacks a zero-check on 		- [MockExchangeRouter.executeMockDeposit(address,address,uint256,uint256,uint256,uint256,address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L90-L107)

contracts/mocks/gmx/MockExchangeRouter.sol#L90-L107


 - [ ] ID-224
variable lacks a zero-check on 		- [UniswapV2Router02.removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L186-L209)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L186-L209


 - [ ] ID-225
variable lacks a zero-check on 		- [UniswapV2Factory.setFeeToSetter(address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L46-L49)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L46-L49


 - [ ] ID-226
variable lacks a zero-check on 		- [GMXVault.updateKeeper(address,bool)](contracts/strategy/gmx/GMXVault.sol#L565-L568)

contracts/strategy/gmx/GMXVault.sol#L565-L568


 - [ ] ID-227
variable lacks a zero-check on 		- [MockExchangeRouter.cancelDeposit(address,address,address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L109-L118)

contracts/mocks/gmx/MockExchangeRouter.sol#L109-L118


 - [ ] ID-228
variable lacks a zero-check on 		- [UniswapV2Router02.removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L170-L183)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L170-L183


 - [ ] ID-229
variable lacks a zero-check on 		- [MockExchangeRouter.executeDeposit(address,address,address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L70-L88)

contracts/mocks/gmx/MockExchangeRouter.sol#L70-L88


 - [ ] ID-230
variable lacks a zero-check on 		- [MockExchangeRouter.cancelWithdrawal(address,address,address,address)](contracts/mocks/gmx/MockExchangeRouter.sol#L168-L178)

contracts/mocks/gmx/MockExchangeRouter.sol#L168-L178


 - [ ] ID-231
variable lacks a zero-check on 		- [UniswapV2Factory.setFeeTo(address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L41-L44)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L41-L44


 - [ ] ID-232
variable lacks a zero-check on 		- [MockUniswapV2Oracle.getLpTokenValue(uint256,address,address,IUniswapV2Pair)](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L173-L206)

contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L173-L206


 - [ ] ID-233
variable lacks a zero-check on 		- [Ownable2Step.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L35-L38)

node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L35-L38


 - [ ] ID-234
variable lacks a zero-check on 		- [UniswapV2Router02.swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L409-L433)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L409-L433


 - [ ] ID-235
variable lacks a zero-check on 		- [GMXDeposit.processDepositFailureLiquidityWithdrawal(GMXTypes.Store)](contracts/strategy/gmx/GMXDeposit.sol#L282-L352)

contracts/strategy/gmx/GMXDeposit.sol#L282-L352


 - [ ] ID-236
variable lacks a zero-check on 		- [UniswapV2Router02.swapExactTokensForETH(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L306-L325)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L306-L325


 - [ ] ID-237
variable lacks a zero-check on 		- [UniswapV2Router02.swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L368-L385)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L368-L385


 - [ ] ID-238
variable lacks a zero-check on 		- [UniswapV2Router02.addLiquidity(address,address,uint256,uint256,address,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L65-L80)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L65-L80


## price-manipulation-low
Impact: Low
Confidence: Medium
 - [ ] ID-239
Potential price manipulation risk:
	- In function getLpTokenReserves
		-- [(reserve0,reserve1) = _pair.getReserves()](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L148) have potential price manipulated risk from reserve1 and call _pair.getReserves() which could influence variable:reserveB
	- In function getLpTokenValue
		-- [(totalReserveA,totalReserveB) = getLpTokenReserves(totalSupply,_tokenA,_tokenB,_pair)](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L181-L186) have potential price manipulated risk from totalReserveB and call _pair.getReserves() which could influence variable:lpFairValueIn18
	- In function getLpTokenAmount
		-- [lpTokenValue = getLpTokenValue(_pair.totalSupply(),_tokenA,_tokenB,_pair)](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L222-L227) have potential price manipulated risk from lpTokenValue and call _pair.getReserves() which could influence variable:lpTokenAmount
	- In function getAmountsOut
		-- [router.getAmountsOut(_amountIn,path)[1]](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L86)have potential price manipulated risk in return call router.getAmountsOut(_amountIn,path) could influence return value

contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L148


## centralized-risk-informational
Impact: Informational
Confidence: High
 - [ ] ID-240
	- [LendingVault.updateMaxCapacity(uint256)](contracts/lending/LendingVault.sol#L638-L642)

contracts/lending/LendingVault.sol#L638-L642


 - [ ] ID-241
	- [UniswapSwap.updateFee(address,address,uint24)](contracts/swaps/UniswapSwap.sol#L147-L153)

contracts/swaps/UniswapSwap.sol#L147-L153


 - [ ] ID-242
	- [GMXVault.updateKeeper(address,bool)](contracts/strategy/gmx/GMXVault.sol#L565-L568)

contracts/strategy/gmx/GMXVault.sol#L565-L568


 - [ ] ID-243
	- [ChainlinkOracle.addTokenMaxDelay(address,uint256)](contracts/oracles/ChainlinkOracle.sol#L217-L223)

contracts/oracles/ChainlinkOracle.sol#L217-L223


 - [ ] ID-244
	- [LendingVault.updateMaxInterestRate(ILendingVault.InterestRate)](contracts/lending/LendingVault.sol#L648-L662)

contracts/lending/LendingVault.sol#L648-L662


 - [ ] ID-245
	- [ChainlinkOracle.addTokenPriceFeed(address,address)](contracts/oracles/ChainlinkOracle.sol#L204-L210)

contracts/oracles/ChainlinkOracle.sol#L204-L210


 - [ ] ID-246
	- [MockGMXOracleWithAdjusts.changeAdjust(uint256,uint256)](contracts/mocks/MockGMXOracleWithAdjusts.sol#L18-L24)

contracts/mocks/MockGMXOracleWithAdjusts.sol#L18-L24


 - [ ] ID-247
	- [MockChainlinkARBOracleWithAdjusts.changeAdjust(uint256)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L18-L22)

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L18-L22


 - [ ] ID-248
	- [LendingVault.withdrawReserve(uint256)](contracts/lending/LendingVault.sol#L371-L382)

contracts/lending/LendingVault.sol#L371-L382


 - [ ] ID-249
	- [MockChainlinkARBOracleWithAdjusts.addTokenPriceFeed(address,address)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L247-L253)

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L247-L253


 - [ ] ID-250
	- [LendingVault.approveBorrower(address)](contracts/lending/LendingVault.sol#L554-L558)

contracts/lending/LendingVault.sol#L554-L558


 - [ ] ID-251
	- [MockLendingVault.updateTreasury(address)](contracts/mocks/MockLendingVault.sol#L675-L679)

contracts/mocks/MockLendingVault.sol#L675-L679


 - [ ] ID-252
	- [UniswapV2Pair.initialize(address,address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L68-L72)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L68-L72


 - [ ] ID-253
	- [MockLendingVault.updateMaxInterestRate(ILendingVault.InterestRate)](contracts/mocks/MockLendingVault.sol#L655-L669)

contracts/mocks/MockLendingVault.sol#L655-L669


 - [ ] ID-254
	- [UniswapV2Factory.setFeeToSetter(address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L46-L49)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L46-L49


 - [ ] ID-255
	- [ChainlinkARBOracle.addTokenMaxDeviation(address,uint256)](contracts/oracles/ChainlinkARBOracle.sol#L262-L268)

contracts/oracles/ChainlinkARBOracle.sol#L262-L268


 - [ ] ID-256
	- [LendingVault.emergencyRepay(uint256,address)](contracts/lending/LendingVault.sol#L585-L614)

contracts/lending/LendingVault.sol#L585-L614


 - [ ] ID-257
	- [MockChainlinkARBOracleWithAdjusts.addTokenMaxDeviation(address,uint256)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L273-L279)

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L273-L279


 - [ ] ID-258
	- [LendingVault.updatePerformanceFee(uint256)](contracts/lending/LendingVault.sol#L541-L548)

contracts/lending/LendingVault.sol#L541-L548


 - [ ] ID-259
	- [TraderJoeSwap.updatePairBinStep(address,address,uint256)](contracts/swaps/TraderJoeSwap.sol#L164-L170)

contracts/swaps/TraderJoeSwap.sol#L164-L170


 - [ ] ID-260
	- [LendingVault.updateKeeper(address,bool)](contracts/lending/LendingVault.sol#L575-L579)

contracts/lending/LendingVault.sol#L575-L579


 - [ ] ID-261
	- [UniswapV2Factory.setFeeTo(address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L41-L44)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Factory.sol#L41-L44


 - [ ] ID-262
	- [MockLendingVault.updateInterestRate(ILendingVault.InterestRate)](contracts/mocks/MockLendingVault.sol#L520-L542)

contracts/mocks/MockLendingVault.sol#L520-L542


 - [ ] ID-263
	- [LendingVault.updateInterestRate(ILendingVault.InterestRate)](contracts/lending/LendingVault.sol#L513-L535)

contracts/lending/LendingVault.sol#L513-L535


 - [ ] ID-264
	- [LendingVault.revokeBorrower(address)](contracts/lending/LendingVault.sol#L564-L568)

contracts/lending/LendingVault.sol#L564-L568


 - [ ] ID-265
	- [Ownable2Step.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L35-L38)

node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L35-L38


 - [ ] ID-266
	- [LendingVault.updateTreasury(address)](contracts/lending/LendingVault.sol#L668-L672)

contracts/lending/LendingVault.sol#L668-L672


 - [ ] ID-267
	- [ChainlinkOracle.addTokenMaxDeviation(address,uint256)](contracts/oracles/ChainlinkOracle.sol#L230-L236)

contracts/oracles/ChainlinkOracle.sol#L230-L236


 - [ ] ID-268
	- [MockLendingVault.updatePerformanceFee(uint256)](contracts/mocks/MockLendingVault.sol#L548-L555)

contracts/mocks/MockLendingVault.sol#L548-L555


 - [ ] ID-269
	- [ChainlinkARBOracle.addTokenMaxDelay(address,uint256)](contracts/oracles/ChainlinkARBOracle.sol#L249-L255)

contracts/oracles/ChainlinkARBOracle.sol#L249-L255


 - [ ] ID-270
	- [MockChainlinkARBOracleWithAdjusts.addTokenMaxDelay(address,uint256)](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L260-L266)

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L260-L266


 - [ ] ID-271
	- [ChainlinkARBOracle.addTokenPriceFeed(address,address)](contracts/oracles/ChainlinkARBOracle.sol#L236-L242)

contracts/oracles/ChainlinkARBOracle.sol#L236-L242


## centralized-risk-other
Impact: Informational
Confidence: High
 - [ ] ID-272
	- [GMXVault.rebalanceAdd(GMXTypes.RebalanceAddParams)](contracts/strategy/gmx/GMXVault.sol#L442-L446)

contracts/strategy/gmx/GMXVault.sol#L442-L446


 - [ ] ID-273
	- [ChainlinkARBOracle.emergencyResume()](contracts/oracles/ChainlinkARBOracle.sol#L280-L282)

contracts/oracles/ChainlinkARBOracle.sol#L280-L282


 - [ ] ID-274
	- [MockChainlinkARBOracleWithAdjusts.emergencyResume()](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L291-L293)

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L291-L293


 - [ ] ID-275
	- [Ownable.renounceOwnership()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L76-L78)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L76-L78


 - [ ] ID-276
	- [GMXVault.processRebalanceAddCancellation()](contracts/strategy/gmx/GMXVault.sol#L462-L464)

contracts/strategy/gmx/GMXVault.sol#L462-L464


 - [ ] ID-277
	- [ChainlinkOracle.emergencyResume()](contracts/oracles/ChainlinkOracle.sol#L248-L250)

contracts/oracles/ChainlinkOracle.sol#L248-L250


 - [ ] ID-278
	- [GMXVault.processWithdrawFailure(uint256,uint256)](contracts/strategy/gmx/GMXVault.sol#L421-L426)

contracts/strategy/gmx/GMXVault.sol#L421-L426


 - [ ] ID-279
	- [Ownable.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable.sol#L84-L89)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L84-L89


 - [ ] ID-280
	- [GMXVault.processDepositFailure(uint256,uint256)](contracts/strategy/gmx/GMXVault.sol#L381-L386)

contracts/strategy/gmx/GMXVault.sol#L381-L386


 - [ ] ID-281
	- [GMXVault.emergencyResume()](contracts/strategy/gmx/GMXVault.sol#L536-L538)

contracts/strategy/gmx/GMXVault.sol#L536-L538


 - [ ] ID-282
	- [GMXVault.compound(GMXTypes.CompoundParams)](contracts/strategy/gmx/GMXVault.sol#L501-L503)

contracts/strategy/gmx/GMXVault.sol#L501-L503


 - [ ] ID-283
	- [GMXVault.processDeposit()](contracts/strategy/gmx/GMXVault.sol#L362-L364)

contracts/strategy/gmx/GMXVault.sol#L362-L364


 - [ ] ID-284
	- [GMXVault.burn(address,uint256)](contracts/strategy/gmx/GMXVault.sol#L687-L689)

contracts/strategy/gmx/GMXVault.sol#L687-L689


 - [ ] ID-285
	- [ChainlinkOracle.emergencyPause()](contracts/oracles/ChainlinkOracle.sol#L241-L243)

contracts/oracles/ChainlinkOracle.sol#L241-L243


 - [ ] ID-286
	- [GMXCallback.afterDepositExecution(bytes32,IDeposit.Props,IEvent.Props)](contracts/strategy/gmx/GMXCallback.sol#L57-L89)

contracts/strategy/gmx/GMXCallback.sol#L57-L89


 - [ ] ID-287
	- [GMXVault.processRebalanceRemove()](contracts/strategy/gmx/GMXVault.sol#L482-L484)

contracts/strategy/gmx/GMXVault.sol#L482-L484


 - [ ] ID-288
	- [GMXVault.emergencyPause()](contracts/strategy/gmx/GMXVault.sol#L528-L530)

contracts/strategy/gmx/GMXVault.sol#L528-L530


 - [ ] ID-289
	- [ChainlinkARBOracle.emergencyPause()](contracts/oracles/ChainlinkARBOracle.sol#L273-L275)

contracts/oracles/ChainlinkARBOracle.sol#L273-L275


 - [ ] ID-290
	- [GMXVault.processCompound()](contracts/strategy/gmx/GMXVault.sol#L510-L512)

contracts/strategy/gmx/GMXVault.sol#L510-L512


 - [ ] ID-291
	- [GMXVault.processRebalanceRemoveCancellation()](contracts/strategy/gmx/GMXVault.sol#L491-L493)

contracts/strategy/gmx/GMXVault.sol#L491-L493


 - [ ] ID-292
	- [GMXCallback.afterWithdrawalCancellation(bytes32,IWithdrawal.Props,IEvent.Props)](contracts/strategy/gmx/GMXCallback.sol#L152-L168)

contracts/strategy/gmx/GMXCallback.sol#L152-L168


 - [ ] ID-293
	- [GMXVault.processRebalanceAdd()](contracts/strategy/gmx/GMXVault.sol#L453-L455)

contracts/strategy/gmx/GMXVault.sol#L453-L455


 - [ ] ID-294
	- [GMXVault.emergencyClose(uint256)](contracts/strategy/gmx/GMXVault.sol#L555-L557)

contracts/strategy/gmx/GMXVault.sol#L555-L557


 - [ ] ID-295
	- [MockChainlinkARBOracleWithAdjusts.emergencyPause()](contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L284-L286)

contracts/mocks/MockChainlinkARBOracleWithAdjusts.sol#L284-L286


 - [ ] ID-296
	- [GMXCallback.afterWithdrawalExecution(bytes32,IWithdrawal.Props,IEvent.Props)](contracts/strategy/gmx/GMXCallback.sol#L122-L145)

contracts/strategy/gmx/GMXCallback.sol#L122-L145


 - [ ] ID-297
	- [GMXVault.processDepositCancellation()](contracts/strategy/gmx/GMXVault.sol#L371-L373)

contracts/strategy/gmx/GMXVault.sol#L371-L373


 - [ ] ID-298
	- [MockLendingVault.emergencyResume()](contracts/mocks/MockLendingVault.sol#L635-L639)

contracts/mocks/MockLendingVault.sol#L635-L639


 - [ ] ID-299
	- [GMXCallback.afterDepositCancellation(bytes32,IDeposit.Props,IEvent.Props)](contracts/strategy/gmx/GMXCallback.sol#L96-L115)

contracts/strategy/gmx/GMXCallback.sol#L96-L115


 - [ ] ID-300
	- [GMXVault.processWithdrawCancellation()](contracts/strategy/gmx/GMXVault.sol#L411-L413)

contracts/strategy/gmx/GMXVault.sol#L411-L413


 - [ ] ID-301
	- [GMXVault.processCompoundCancellation()](contracts/strategy/gmx/GMXVault.sol#L519-L521)

contracts/strategy/gmx/GMXVault.sol#L519-L521


 - [ ] ID-302
	- [GMXVault.processDepositFailureLiquidityWithdrawal()](contracts/strategy/gmx/GMXVault.sol#L393-L395)

contracts/strategy/gmx/GMXVault.sol#L393-L395


 - [ ] ID-303
	- [GMXVault.processEmergencyResume()](contracts/strategy/gmx/GMXVault.sol#L545-L547)

contracts/strategy/gmx/GMXVault.sol#L545-L547


 - [ ] ID-304
	- [GMXVault.processWithdrawFailureLiquidityAdded()](contracts/strategy/gmx/GMXVault.sol#L433-L435)

contracts/strategy/gmx/GMXVault.sol#L433-L435


 - [ ] ID-305
	- [GMXVault.mint(address,uint256)](contracts/strategy/gmx/GMXVault.sol#L677-L679)

contracts/strategy/gmx/GMXVault.sol#L677-L679


 - [ ] ID-306
	- [GMXVault.processWithdraw()](contracts/strategy/gmx/GMXVault.sol#L402-L404)

contracts/strategy/gmx/GMXVault.sol#L402-L404


 - [ ] ID-307
	- [LendingVault.emergencyResume()](contracts/lending/LendingVault.sol#L628-L632)

contracts/lending/LendingVault.sol#L628-L632


 - [ ] ID-308
	- [GMXVault.rebalanceRemove(GMXTypes.RebalanceRemoveParams)](contracts/strategy/gmx/GMXVault.sol#L471-L475)

contracts/strategy/gmx/GMXVault.sol#L471-L475


## dead-function
Impact: Informational
Confidence: Medium
 - [ ] ID-309
[Context._msgData()](node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/Context.sol#L21-L23


 - [ ] ID-310
[ReentrancyGuard._reentrancyGuardEntered()](node_modules/@openzeppelin/contracts/utils/ReentrancyGuard.sol#L81-L83) is never used and should be removed

node_modules/@openzeppelin/contracts/utils/ReentrancyGuard.sol#L81-L83


## error-msg
Impact: Informational
Confidence: Medium
 - [ ] ID-311
require() missing error messages
	 - [require(bool)(balanceOf(msg.sender) >= wad)](contracts/mocks/MockWETH.sol#L23)

contracts/mocks/MockWETH.sol#L23


## price-manipulation-info
Impact: Informational
Confidence: Medium
 - [ ] ID-312
Potential price manipulation risk:
	- In function getLpTokenReserves
		-- [longTokenBalance = IERC20(longToken).balanceOf(marketToken)](contracts/mocks/gmx/MockGMXOracle.sol#L45) have potential price manipulated risk from longTokenBalance and call None which could influence variable:longTokenBalance
	- In function getLpTokenReserves
		-- [shortTokenBalance = IERC20(shortToken).balanceOf(marketToken)](contracts/mocks/gmx/MockGMXOracle.sol#L46) have potential price manipulated risk from shortTokenBalance and call None which could influence variable:shortTokenBalance

contracts/mocks/gmx/MockGMXOracle.sol#L45


 - [ ] ID-313
Potential price manipulation risk:
	- In function svTokenValue
		-- [equityValue_ = equityValue(self)](contracts/strategy/gmx/GMXReader.sol#L28) have potential price manipulated risk from equityValue_ and call None which could influence variable:equityValue_
	- In function equityValue
		-- [assetValue_ = assetValue(self)](contracts/strategy/gmx/GMXReader.sol#L131) have potential price manipulated risk from assetValue_ and call None which could influence variable:assetValue_
	- In function delta
		-- [(_tokenAAmt) = assetAmt(self)](contracts/strategy/gmx/GMXReader.sol#L195) have potential price manipulated risk from _tokenAAmt and call None which could influence variable:signedDelta
	- In function delta
		-- [equityValue_ = equityValue(self)](contracts/strategy/gmx/GMXReader.sol#L197) have potential price manipulated risk from equityValue_ and call None which could influence variable:signedDelta

contracts/strategy/gmx/GMXReader.sol#L28


## uncontrolled-resource-consumption
Impact: Informational
Confidence: Medium
 - [ ] ID-314
Potential DoS Gas Limit Attack occur in[UniswapV2Router02._swap(uint256[],address[],address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L229-L240)[BEGIN_LOOP](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L230-L239)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L229-L240


 - [ ] ID-315
Potential DoS Gas Limit Attack occur in[UniswapV2Router02._swapSupportingFeeOnTransferTokens(address[],address)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L350-L367)[BEGIN_LOOP](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L351-L366)

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L350-L367


## unnecessary-public-function-modifier
Impact: Informational
Confidence: High
 - [ ] ID-316
function:[GMXManager.removeLiquidity(GMXTypes.Store,GMXTypes.RemoveLiquidityParams)](contracts/strategy/gmx/GMXManager.sol#L276-L281)is public and can be replaced with external 

contracts/strategy/gmx/GMXManager.sol#L276-L281


 - [ ] ID-317
function:[ERC20.symbol()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L66-L68)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L66-L68


 - [ ] ID-318
function:[GMXReader.capacity(GMXTypes.Store)](contracts/strategy/gmx/GMXReader.sol#L282-L284)is public and can be replaced with external 

contracts/strategy/gmx/GMXReader.sol#L282-L284


 - [ ] ID-319
function:[MockLendingVault._onlyKeeper()](contracts/mocks/MockLendingVault.sol#L403-L405)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L403-L405


 - [ ] ID-320
function:[MockLendingVault.lvTokenValue()](contracts/mocks/MockLendingVault.sol#L188-L197)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L188-L197


 - [ ] ID-321
function:[ERC20.totalSupply()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L90-L92)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L90-L92


 - [ ] ID-322
function:[MockUniswapV2Oracle.lpToken(address,address)](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L50-L55)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L50-L55


 - [ ] ID-323
function:[LendingVault.updateInterestRate(ILendingVault.InterestRate)](contracts/lending/LendingVault.sol#L513-L535)is public and can be replaced with external 

contracts/lending/LendingVault.sol#L513-L535


 - [ ] ID-324
function:[MockLendingVault.deposit(uint256,uint256)](contracts/mocks/MockLendingVault.sol#L268-L282)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L268-L282


 - [ ] ID-325
function:[UniswapV2Router02.getAmountsIn(uint256,address[])](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L470-L478)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L470-L478


 - [ ] ID-326
function:[MockGMXOracleWithAdjusts.getLpTokenReserves(address,address,address,address)](contracts/mocks/MockGMXOracleWithAdjusts.sol#L211-L236)is public and can be replaced with external 

contracts/mocks/MockGMXOracleWithAdjusts.sol#L211-L236


 - [ ] ID-327
function:[ERC20.approve(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L132-L136)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L132-L136


 - [ ] ID-328
function:[ERC20.balanceOf(address)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L97-L99)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L97-L99


 - [ ] ID-329
function:[MockUniswapV2Oracle.getAmountsOut(uint256,address,address,IUniswapV2Pair)](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L65-L87)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L65-L87


 - [ ] ID-330
function:[Ownable2Step.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L35-L38)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L35-L38


 - [ ] ID-331
function:[MockUniswapV2Oracle.getLpTokenAmount(uint256,address,address,IUniswapV2Pair)](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L216-L232)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L216-L232


 - [ ] ID-332
function:[MockLendingVault.updateMaxInterestRate(ILendingVault.InterestRate)](contracts/mocks/MockLendingVault.sol#L655-L669)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L655-L669


 - [ ] ID-333
function:[UniswapV2Router02.getAmountsOut(uint256,address[])](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L460-L468)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L460-L468


 - [ ] ID-334
function:[LendingVault.withdraw(uint256,uint256)](contracts/lending/LendingVault.sol#L282-L303)is public and can be replaced with external 

contracts/lending/LendingVault.sol#L282-L303


 - [ ] ID-335
function:[GMXVault.assetValue()](contracts/strategy/gmx/GMXVault.sol#L203-L205)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L203-L205


 - [ ] ID-336
function:[MockERC20.decimals()](contracts/mocks/MockERC20.sol#L21-L23)is public and can be replaced with external 

contracts/mocks/MockERC20.sol#L21-L23


 - [ ] ID-337
function:[MockGMXOracleWithAdjusts.getAmountsIn(address,address,address,address,address,uint256)](contracts/mocks/MockGMXOracleWithAdjusts.sol#L131-L147)is public and can be replaced with external 

contracts/mocks/MockGMXOracleWithAdjusts.sol#L131-L147


 - [ ] ID-338
function:[ERC20.transferFrom(address,address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L154-L159)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L154-L159


 - [ ] ID-339
function:[Ownable.transferOwnership(address)](node_modules/@openzeppelin/contracts/access/Ownable.sol#L84-L89)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable.sol#L84-L89


 - [ ] ID-340
function:[ERC20.name()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L58-L60)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L58-L60


 - [ ] ID-341
function:[MockUniswapV2Oracle.getAmountsIn(uint256,address,address,IUniswapV2Pair)](contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L97-L119)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/MockUniswapV2Oracle.sol#L97-L119


 - [ ] ID-342
function:[MockLendingVault.updateInterestRate(ILendingVault.InterestRate)](contracts/mocks/MockLendingVault.sol#L520-L542)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L520-L542


 - [ ] ID-343
function:[MockLendingVault.lendingAPR()](contracts/mocks/MockLendingVault.sol#L211-L223)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L211-L223


 - [ ] ID-344
function:[GMXVault.lpAmt()](contracts/strategy/gmx/GMXVault.sol#L248-L250)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L248-L250


 - [ ] ID-345
function:[GMXVault.debtRatio()](contracts/strategy/gmx/GMXVault.sol#L275-L277)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L275-L277


 - [ ] ID-346
function:[GMXVault.assetAmt()](contracts/strategy/gmx/GMXVault.sol#L231-L233)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L231-L233


 - [ ] ID-347
function:[MockLendingVault.withdraw(uint256,uint256)](contracts/mocks/MockLendingVault.sol#L289-L310)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L289-L310


 - [ ] ID-348
function:[Ownable2Step.acceptOwnership()](node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L52-L58)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L52-L58


 - [ ] ID-349
function:[GMXVault.svTokenValue()](contracts/strategy/gmx/GMXVault.sol#L157-L159)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L157-L159


 - [ ] ID-350
function:[GMXVault.equityValue()](contracts/strategy/gmx/GMXVault.sol#L222-L224)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L222-L224


 - [ ] ID-351
function:[Ownable.renounceOwnership()](node_modules/@openzeppelin/contracts/access/Ownable.sol#L76-L78)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/access/Ownable.sol#L76-L78


 - [ ] ID-352
function:[GMXReader.svTokenValue(GMXTypes.Store)](contracts/strategy/gmx/GMXReader.sol#L27-L32)is public and can be replaced with external 

contracts/strategy/gmx/GMXReader.sol#L27-L32


 - [ ] ID-353
function:[GMXReader.valueToShares(GMXTypes.Store,uint256,uint256)](contracts/strategy/gmx/GMXReader.sol#L48-L56)is public and can be replaced with external 

contracts/strategy/gmx/GMXReader.sol#L48-L56


 - [ ] ID-354
function:[MockGMXOracle.getLpTokenValue(address,address,address,address,bool,bool)](contracts/mocks/gmx/MockGMXOracle.sol#L19-L37)is public and can be replaced with external 

contracts/mocks/gmx/MockGMXOracle.sol#L19-L37


 - [ ] ID-355
function:[UniswapV2Router02.getAmountIn(uint256,uint256,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L450-L458)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L450-L458


 - [ ] ID-356
function:[LendingVault.updateMaxInterestRate(ILendingVault.InterestRate)](contracts/lending/LendingVault.sol#L648-L662)is public and can be replaced with external 

contracts/lending/LendingVault.sol#L648-L662


 - [ ] ID-357
function:[UniswapV2Router02.getAmountOut(uint256,uint256,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L440-L448)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L440-L448


 - [ ] ID-358
function:[GMXReader.delta(GMXTypes.Store)](contracts/strategy/gmx/GMXReader.sol#L194-L214)is public and can be replaced with external 

contracts/strategy/gmx/GMXReader.sol#L194-L214


 - [ ] ID-359
function:[GMXVault.additionalCapacity()](contracts/strategy/gmx/GMXVault.sol#L283-L285)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L283-L285


 - [ ] ID-360
function:[MockWETH.deposit()](contracts/mocks/MockWETH.sol#L18-L20)is public and can be replaced with external 

contracts/mocks/MockWETH.sol#L18-L20


 - [ ] ID-361
function:[GMXManager.addLiquidity(GMXTypes.Store,GMXTypes.AddLiquidityParams)](contracts/strategy/gmx/GMXManager.sol#L263-L268)is public and can be replaced with external 

contracts/strategy/gmx/GMXManager.sol#L263-L268


 - [ ] ID-362
function:[GMXVault.debtAmt()](contracts/strategy/gmx/GMXVault.sol#L240-L242)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L240-L242


 - [ ] ID-363
function:[GMXVault.capacity()](contracts/strategy/gmx/GMXVault.sol#L291-L293)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L291-L293


 - [ ] ID-364
function:[GMXVault.store()](contracts/strategy/gmx/GMXVault.sol#L140-L142)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L140-L142


 - [ ] ID-365
function:[GMXVault.convertToUsdValue(address,uint256)](contracts/strategy/gmx/GMXVault.sol#L185-L187)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L185-L187


 - [ ] ID-366
function:[GMXVault.isTokenWhitelisted(address)](contracts/strategy/gmx/GMXVault.sol#L149-L151)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L149-L151


 - [ ] ID-367
function:[GMXReader.leverage(GMXTypes.Store)](contracts/strategy/gmx/GMXReader.sol#L185-L188)is public and can be replaced with external 

contracts/strategy/gmx/GMXReader.sol#L185-L188


 - [ ] ID-368
function:[MockGMXOracleWithAdjusts.getLpTokenAmount(uint256,address,address,address,address,bool,bool)](contracts/mocks/MockGMXOracleWithAdjusts.sol#L296-L315)is public and can be replaced with external 

contracts/mocks/MockGMXOracleWithAdjusts.sol#L296-L315


 - [ ] ID-369
function:[GMXVault.debtValue()](contracts/strategy/gmx/GMXVault.sol#L213-L215)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L213-L215


 - [ ] ID-370
function:[UniswapV2Router02.quote(uint256,uint256,uint256)](contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L436-L438)is public and can be replaced with external 

contracts/mocks/gmx/MockUniswapV2/UniswapV2Router02.sol#L436-L438


 - [ ] ID-371
function:[ERC20.transfer(address,uint256)](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L109-L113)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L109-L113


 - [ ] ID-372
function:[GMXReader.debtRatio(GMXTypes.Store)](contracts/strategy/gmx/GMXReader.sol#L220-L224)is public and can be replaced with external 

contracts/strategy/gmx/GMXReader.sol#L220-L224


 - [ ] ID-373
function:[GMXVault.pendingFee()](contracts/strategy/gmx/GMXVault.sol#L165-L167)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L165-L167


 - [ ] ID-374
function:[GMXVault.mintFee()](contracts/strategy/gmx/GMXVault.sol#L334-L337)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L334-L337


 - [ ] ID-375
function:[MockWETH.withdraw(uint256)](contracts/mocks/MockWETH.sol#L22-L26)is public and can be replaced with external 

contracts/mocks/MockWETH.sol#L22-L26


 - [ ] ID-376
function:[LendingVault.deposit(uint256,uint256)](contracts/lending/LendingVault.sol#L261-L275)is public and can be replaced with external 

contracts/lending/LendingVault.sol#L261-L275


 - [ ] ID-377
function:[GMXVault.delta()](contracts/strategy/gmx/GMXVault.sol#L266-L268)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L266-L268


 - [ ] ID-378
function:[ERC20.decimals()](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L83-L85)is public and can be replaced with external 

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L83-L85


 - [ ] ID-379
function:[GMXVault.leverage()](contracts/strategy/gmx/GMXVault.sol#L256-L258)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L256-L258


 - [ ] ID-380
function:[GMXOracle.getAmountsIn(address,address,address,address,address,uint256)](contracts/oracles/GMXOracle.sol#L119-L135)is public and can be replaced with external 

contracts/oracles/GMXOracle.sol#L119-L135


 - [ ] ID-381
function:[GMXOracle.getLpTokenAmount(uint256,address,address,address,address,bool,bool)](contracts/oracles/GMXOracle.sol#L280-L299)is public and can be replaced with external 

contracts/oracles/GMXOracle.sol#L280-L299


 - [ ] ID-382
function:[GMXVault.tokenWeights()](contracts/strategy/gmx/GMXVault.sol#L194-L196)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L194-L196


 - [ ] ID-383
function:[MockGMXOracle.getLpTokenReserves(address,address,address,address)](contracts/mocks/gmx/MockGMXOracle.sol#L39-L49)is public and can be replaced with external 

contracts/mocks/gmx/MockGMXOracle.sol#L39-L49


 - [ ] ID-384
function:[MockLendingVault._onlyBorrower()](contracts/mocks/MockLendingVault.sol#L396-L398)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L396-L398


 - [ ] ID-385
function:[GMXManager.borrow(GMXTypes.Store,uint256,uint256)](contracts/strategy/gmx/GMXManager.sol#L225-L236)is public and can be replaced with external 

contracts/strategy/gmx/GMXManager.sol#L225-L236


 - [ ] ID-386
function:[LendingVault.depositNative(uint256,uint256)](contracts/lending/LendingVault.sol#L238-L254)is public and can be replaced with external 

contracts/lending/LendingVault.sol#L238-L254


 - [ ] ID-387
function:[MockLendingVault.depositNative(uint256,uint256)](contracts/mocks/MockLendingVault.sol#L245-L261)is public and can be replaced with external 

contracts/mocks/MockLendingVault.sol#L245-L261


 - [ ] ID-388
function:[GMXManager.repay(GMXTypes.Store,uint256,uint256)](contracts/strategy/gmx/GMXManager.sol#L244-L255)is public and can be replaced with external 

contracts/strategy/gmx/GMXManager.sol#L244-L255


 - [ ] ID-389
function:[LendingVault.lendingAPR()](contracts/lending/LendingVault.sol#L204-L216)is public and can be replaced with external 

contracts/lending/LendingVault.sol#L204-L216


 - [ ] ID-390
function:[GMXOracle.getLpTokenReserves(address,address,address,address)](contracts/oracles/GMXOracle.sol#L197-L222)is public and can be replaced with external 

. analyzed (92 contracts with 86 detectors), 413 result(s) found
INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/2023-10-SteadeFi_scaned_output/mwe-output.json generate success.
contracts/oracles/GMXOracle.sol#L197-L222


 - [ ] ID-391
function:[GMXVault.valueToShares(uint256,uint256)](contracts/strategy/gmx/GMXVault.sol#L175-L177)is public and can be replaced with external 

contracts/strategy/gmx/GMXVault.sol#L175-L177


 - [ ] ID-392
function:[LendingVault.lvTokenValue()](contracts/lending/LendingVault.sol#L181-L190)is public and can be replaced with external 

contracts/lending/LendingVault.sol#L181-L190


## version-only
Impact: Informational
Confidence: High
 - [ ] ID-393
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#L5)

node_modules/@openzeppelin/contracts/utils/math/SafeCast.sol#L5


 - [ ] ID-394
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol#L4


 - [ ] ID-395
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/utils/Address.sol#L4)

node_modules/@openzeppelin/contracts/utils/Address.sol#L4


 - [ ] ID-396
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/access/Ownable.sol#L4)

node_modules/@openzeppelin/contracts/access/Ownable.sol#L4


 - [ ] ID-397
	Pragma confirmed better, here is pragma solidity^0.8.0. [^0.8.0](node_modules/@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol#L2)

node_modules/@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol#L2


 - [ ] ID-398
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol#L4


 - [ ] ID-399
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/utils/Pausable.sol#L4)

node_modules/@openzeppelin/contracts/utils/Pausable.sol#L4


 - [ ] ID-400
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/utils/math/Math.sol#L4)

node_modules/@openzeppelin/contracts/utils/math/Math.sol#L4


 - [ ] ID-401
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Permit.sol#L4


 - [ ] ID-402
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L4)

node_modules/@openzeppelin/contracts/access/Ownable2Step.sol#L4


 - [ ] ID-403
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol#L4


 - [ ] ID-404
	Pragma confirmed better, here is pragma solidity>=0.4.22<0.9.0. [>=0.4.22<0.9.0](lib/forge-std/src/console.sol#L2)

lib/forge-std/src/console.sol#L2


 - [ ] ID-405
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/interfaces/draft-IERC6093.sol#L3)

node_modules/@openzeppelin/contracts/interfaces/draft-IERC6093.sol#L3


 - [ ] ID-406
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/utils/ReentrancyGuard.sol#L4)

node_modules/@openzeppelin/contracts/utils/ReentrancyGuard.sol#L4


 - [ ] ID-407
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L4)

node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol#L4


 - [ ] ID-408
	Pragma confirmed better, here is pragma solidity^0.8.20. [^0.8.20](node_modules/@openzeppelin/contracts/utils/Context.sol#L4)

node_modules/@openzeppelin/contracts/utils/Context.sol#L4


## state-should-be-constant
Impact: Optimization
Confidence: High
 - [ ] ID-409
[MockSwapper.WETH](contracts/mocks/MockSwapper.sol#L27) should be constant

contracts/mocks/MockSwapper.sol#L27


 - [ ] ID-410
[UniswapV2Pair.price0CumulativeLast](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L28) should be constant

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L28


 - [ ] ID-411
[UniswapV2Pair.price1CumulativeLast](contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L29) should be constant

contracts/mocks/gmx/MockUniswapV2/UniswapV2Pair.sol#L29


 - [ ] ID-412
[MockSwapper.USDC](contracts/mocks/MockSwapper.sol#L28) should be constant

contracts/mocks/MockSwapper.sol#L28


Execution time for Falcon: 42.639630758s
