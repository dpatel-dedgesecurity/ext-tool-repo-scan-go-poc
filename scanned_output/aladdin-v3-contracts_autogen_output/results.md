'npx hardhat clean' running (wd: /home/im/dedge/delete/ext-tool-repo-scan-go-poc/input_repos/aladdin-v3-contracts)
'npx hardhat clean --global' running (wd: /home/im/dedge/delete/ext-tool-repo-scan-go-poc/input_repos/aladdin-v3-contracts)
'npx hardhat compile --force' running (wd: /home/im/dedge/delete/ext-tool-repo-scan-go-poc/input_repos/aladdin-v3-contracts)
Missing function Incorrect ternary conversion _routes = routes[if _isETH(_fromToken) then WETH else _fromToken][WETH] contracts/zap/AladdinZap.sol#71
Traceback (most recent call last):
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 835, in main_impl
    ) = process_all(filename, args, detector_classes, printer_classes)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 100, in process_all
    ) = process_single(compilation, args, detector_classes, printer_classes)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 75, in process_single
    falcon = Falcon(target, ast_format=ast, **vars(args))
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/falcon.py", line 129, in __init__
    parser.analyze_contracts()
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/solc_parsing/falcon_compilation_unit_solc.py", line 507, in analyze_contracts
    self._convert_to_falconir()
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/solc_parsing/falcon_compilation_unit_solc.py", line 690, in _convert_to_falconir
    func.generate_falconir_and_analyze()
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/core/declarations/function.py", line 1701, in generate_falconir_and_analyze
    node.falconir_generation()
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/core/cfg/node.py", line 719, in falconir_generation
    self._irs = convert_expression(expression, self)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/ir/convert.py", line 116, in convert_expression
    visitor = ExpressionToFalconIR(expression, node)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/falconir/expression_to_falconir.py", line 143, in __init__
    self._visit_expression(self.expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 45, in _visit_expression
    self._visit_assignement_operation(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 101, in _visit_assignement_operation
    self._visit_expression(expression.expression_right)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 63, in _visit_expression
    self._visit_index_access(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 131, in _visit_index_access
    self._visit_expression(expression.expression_left)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 63, in _visit_expression
    self._visit_index_access(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 132, in _visit_index_access
    self._visit_expression(expression.expression_right)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 95, in _visit_expression
    self._post_visit(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 274, in _post_visit
    self._post_conditional_expression(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/falconir/expression_to_falconir.py", line 355, in _post_conditional_expression
    raise Exception(f"Ternary operator are not convertible to FalconIR {expression}")
Exception: Ternary operator are not convertible to FalconIR if _isETH(_fromToken) then WETH else _fromToken
None
Error in .
Traceback (most recent call last):
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 835, in main_impl
    ) = process_all(filename, args, detector_classes, printer_classes)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 100, in process_all
    ) = process_single(compilation, args, detector_classes, printer_classes)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 75, in process_single
    falcon = Falcon(target, ast_format=ast, **vars(args))
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/falcon.py", line 129, in __init__
    parser.analyze_contracts()
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/solc_parsing/falcon_compilation_unit_solc.py", line 507, in analyze_contracts
    self._convert_to_falconir()
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/solc_parsing/falcon_compilation_unit_solc.py", line 690, in _convert_to_falconir
    func.generate_falconir_and_analyze()
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/core/declarations/function.py", line 1701, in generate_falconir_and_analyze
    node.falconir_generation()
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/core/cfg/node.py", line 719, in falconir_generation
    self._irs = convert_expression(expression, self)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/ir/convert.py", line 116, in convert_expression
    visitor = ExpressionToFalconIR(expression, node)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/falconir/expression_to_falconir.py", line 143, in __init__
    self._visit_expression(self.expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 45, in _visit_expression
    self._visit_assignement_operation(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 101, in _visit_assignement_operation
    self._visit_expression(expression.expression_right)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 63, in _visit_expression
    self._visit_index_access(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 131, in _visit_index_access
    self._visit_expression(expression.expression_left)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 63, in _visit_expression
    self._visit_index_access(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 132, in _visit_index_access
    self._visit_expression(expression.expression_right)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 95, in _visit_expression
    self._post_visit(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/expression/expression.py", line 274, in _post_visit
    self._post_conditional_expression(expression)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/visitors/falconir/expression_to_falconir.py", line 355, in _post_conditional_expression
    raise Exception(f"Ternary operator are not convertible to FalconIR {expression}")
Exception: Ternary operator are not convertible to FalconIR if _isETH(_fromToken) then WETH else _fromToken

INFO:Falcon:metatrust result: ../../scanned_output/aladdin-v3-contracts_autogen_output/mwe-output.json generate success.
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
../input_repos/aladdin-v3-contracts/contracts/External.sol: : 8 lines
../input_repos/aladdin-v3-contracts/contracts/clever/CLeverCVXLocker.sol: Smart Contract (1): 1078 lines
../input_repos/aladdin-v3-contracts/contracts/clever/CLeverToken.sol: Smart Contract (1): 97 lines
../input_repos/aladdin-v3-contracts/contracts/clever/Furnace.sol: Smart Contract (1): 613 lines
../input_repos/aladdin-v3-contracts/contracts/clever/generic/MetaCLever.sol: Smart Contract (1): 1030 lines
../input_repos/aladdin-v3-contracts/contracts/clever/generic/MetaFurnace.sol: Smart Contract (1): 609 lines
../input_repos/aladdin-v3-contracts/contracts/clever/strategies/AladdinCRVStrategy.sol: Smart Contract (1): 90 lines
../input_repos/aladdin-v3-contracts/contracts/clever/strategies/ConcentratorBatchStrategy.sol: Smart Contract (1): 76 lines
../input_repos/aladdin-v3-contracts/contracts/clever/strategies/ConcentratorStrategy.sol: Smart Contract (1): 248 lines
../input_repos/aladdin-v3-contracts/contracts/clever/strategies/YieldStrategyBase.sol: Abstract Contract (1): 63 lines
../input_repos/aladdin-v3-contracts/contracts/clever/strategies/upgradable/ConcentratorStrategyUpgradable.sol: Smart Contract (1): 251 lines
../input_repos/aladdin-v3-contracts/contracts/clever/strategies/upgradable/StakeDAOCRVStrategyUpgradeable.sol: Smart Contract (1): 174 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/Diamond.sol: Smart Contract (1): 75 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/facets/DiamondCutFacet.sol: Smart Contract (1): 30 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/facets/DiamondLoupeFacet.sol: Smart Contract (1): 155 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/facets/OwnershipFacet.sol: Smart Contract (1): 21 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/interfaces/IDiamond.sol: Interface (1): 24 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/interfaces/IDiamondCut.sol: Interface (1): 23 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/interfaces/IDiamondLoupe.sol: Interface (1): 38 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/interfaces/IERC165.sol: Interface (1): 12 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/interfaces/IERC173.sol: Interface (1): 19 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/libraries/LibDiamond.sol: Library (1): 210 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/upgradeInitializers/DiamondInit.sol: Smart Contract (1): 42 lines
../input_repos/aladdin-v3-contracts/contracts/common/EIP2535/upgradeInitializers/DiamondMultiInit.sol: Smart Contract (1): 27 lines
../input_repos/aladdin-v3-contracts/contracts/common/ERC4626/IERC4626.sol: Interface (1): 242 lines
../input_repos/aladdin-v3-contracts/contracts/common/codec/WordCodec.sol: Library (1): 120 lines
../input_repos/aladdin-v3-contracts/contracts/common/fees/CustomFeeRate.sol: Abstract Contract (1): 181 lines
../input_repos/aladdin-v3-contracts/contracts/common/fees/FeeCustomization.sol: Abstract Contract (1): 121 lines
../input_repos/aladdin-v3-contracts/contracts/common/math/DecrementalFloatingPoint.sol: Library (1): 94 lines
../input_repos/aladdin-v3-contracts/contracts/common/math/ExponentialMovingAverageV7.sol: Library (1): 59 lines
../input_repos/aladdin-v3-contracts/contracts/common/math/ExponentialMovingAverageV8.sol: Library (1): 62 lines
../input_repos/aladdin-v3-contracts/contracts/common/math/GaussElimination.sol: Library (1): 43 lines
../input_repos/aladdin-v3-contracts/contracts/common/math/LogExpMathV7.sol: Library (1): 513 lines
../input_repos/aladdin-v3-contracts/contracts/common/math/LogExpMathV8.sol: Library (1): 527 lines
../input_repos/aladdin-v3-contracts/contracts/common/math/StableMath.sol: Library (1): 94 lines
../input_repos/aladdin-v3-contracts/contracts/common/rewards/accumulator/IMultipleRewardAccumulator.sol: Interface (1): 83 lines
../input_repos/aladdin-v3-contracts/contracts/common/rewards/accumulator/MultipleRewardAccumulator.sol: Abstract Contract (1): 309 lines
../input_repos/aladdin-v3-contracts/contracts/common/rewards/accumulator/MultipleRewardCompoundingAccumulator.sol: Abstract Contract (1): 364 lines
../input_repos/aladdin-v3-contracts/contracts/common/rewards/distributor/IMultipleRewardDistributor.sol: Interface (1): 84 lines
../input_repos/aladdin-v3-contracts/contracts/common/rewards/distributor/IRewardDistributor.sol: Interface (1): 35 lines
../input_repos/aladdin-v3-contracts/contracts/common/rewards/distributor/LinearMultipleRewardDistributor.sol: Abstract Contract (1): 209 lines
../input_repos/aladdin-v3-contracts/contracts/common/rewards/distributor/LinearReward.sol: Library (1): 84 lines
../input_repos/aladdin-v3-contracts/contracts/common/rewards/distributor/LinearRewardDistributor.sol: Abstract Contract (1): 125 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/AladdinCompounder.sol: Abstract Contract (1): 395 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/AladdinCompounderWithStrategy.sol: Abstract Contract (1): 212 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/ConcentratorBase.sol: Smart Contract (1): 58 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/ConcentratorBaseV2.sol: Abstract Contract (1): 224 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/ConcentratorCompounderBase.sol: Abstract Contract (1): 444 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/ConcentratorConvexVault.sol: Abstract Contract (1): 740 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/ConcentratorGeneralVault.sol: Abstract Contract (1): 821 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/CLeverAMOBase.sol: Abstract Contract (1): 539 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/RewardClaimable.sol: Abstract Contract (1): 102 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/cvx/AladdinCVX.sol: Smart Contract (1): 334 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/interfaces/ICLeverAMO.sol: Interface (1): 158 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/interfaces/ICLeverAMOStrategy.sol: Interface (1): 9 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/interfaces/ILegacyFurnace.sol: Interface (1): 13 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/math/AMOMath.sol: Library (1): 187 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/strategies/AMOConvexCurveStrategy.sol: Smart Contract (1): 18 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/clever/strategies/CLeverGaugeStrategy.sol: Smart Contract (1): 131 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvx/CvxCompounder.sol: Smart Contract (1): 57 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvx/CvxStakingStrategy.sol: Smart Contract (1): 129 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvxcrv/AladdinCRV.sol: Smart Contract (1): 405 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvxcrv/AladdinCRVConvexVault.sol: Smart Contract (1): 813 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvxcrv/AladdinCRVV2.sol: Smart Contract (1): 670 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvxcrv/ConcentratorIFOVault.sol: Smart Contract (1), Interface (1): 144 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvxcrv/CvxCrvStakingWrapperStrategy.sol: Smart Contract (1): 107 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvxcrv/CvxCrvWeightAdjuster.sol: Smart Contract (1): 277 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvxfxn/CvxFxnCompounder.sol: Smart Contract (1): 156 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/cvxfxn/CvxFxnStakingStrategy.sol: Smart Contract (1): 197 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/eth/AladdinETH.sol: Smart Contract (1): 34 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/eth/ConcentratorAladdinETHVault.sol: Smart Contract (1): 94 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/fx-protocol/FxUSDCompounder.sol: Smart Contract (1): 478 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/fx-protocol/FxUSDCompounder4626.sol: Smart Contract (1): 215 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/fx-protocol/FxUSDStandardizedYieldBase.sol: Abstract Contract (1): 332 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/fxs/AladdinFXS.sol: Smart Contract (1): 247 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/fxs/AladdinFXSConvexVault.sol: Smart Contract (1): 140 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/fxs/AladdinFXSV2.sol: Smart Contract (1): 351 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/fxs/CvxFxsStakingStrategy.sol: Smart Contract (1): 164 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/AladdinSdCRV.sol: Smart Contract (1): 281 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/ConcentratorSdCrvGaugeWrapper.sol: Smart Contract (1): 136 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/ConcentratorStakeDAOGaugeWrapper.sol: Abstract Contract (1): 304 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/ConcentratorStakeDAOLocker.sol: Smart Contract (1): 269 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/ConcentratorVaultForAsdCRV.sol: Smart Contract (1): 71 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/SdCRVBribeBurner.sol: Smart Contract (1): 186 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/SdCRVBribeBurnerV2.sol: Smart Contract (1): 181 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/SdCRVLocker.sol: Abstract Contract (1): 105 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/SdCrvCompounder.sol: Smart Contract (1), Interface (1): 404 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/StakeDAOCRVVault.sol: Smart Contract (1): 285 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/StakeDAOVaultBase.sol: Abstract Contract (1): 486 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stakedao/VeSDTDelegation.sol: Smart Contract (1): 420 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stash/ConcentratorCompounderStash.sol: Smart Contract (1): 165 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stash/ConcentratorStashBase.sol: Abstract Contract (1): 55 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stash/CvxCrvCompounderStash.sol: Smart Contract (1), Interface (1): 64 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stash/LegacyCompounderStash.sol: Smart Contract (1), Interface (1): 223 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/stash/StakeDAOGaugeWrapperStash.sol: Smart Contract (1): 53 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/AutoCompoundingConvexCurveStrategy.sol: Smart Contract (1): 87 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/AutoCompoundingConvexFraxStrategy.sol: Smart Contract (1): 305 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/AutoCompoundingStrategyBase.sol: Abstract Contract (1): 44 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/AutoCompoundingStrategyBaseV2.sol: Abstract Contract (1): 56 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/ConcentratorStrategyBase.sol: Abstract Contract (1): 175 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/ConcentratorStrategyBaseV2.sol: Abstract Contract (1): 36 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/ManualCompoundingConvexCurveStrategy.sol: Smart Contract (1): 83 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/ManualCompoundingCurveGaugeStrategy.sol: Smart Contract (1): 86 lines
../input_repos/aladdin-v3-contracts/contracts/concentrator/strategies/ManualCompoundingStrategyBase.sol: Abstract Contract (1): 38 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/FxVault.sol: Smart Contract (1): 351 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/arbitrage/FxUSDArbitrager.sol: Smart Contract (1): 118 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/funding-cost-market/MarketWithFundingCost.sol: Smart Contract (1): 51 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/funding-cost-market/TreasuryWithFundingCost.sol: Smart Contract (1): 166 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/funding-cost-market/funding-rate-adapter/CrvUSDBorrowRateAdapter.sol: Abstract Contract (1): 119 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/math/FxLowVolatilityMath.sol: Library (1): 406 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/math/FxStableMath.sol: Library (1): 246 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxBTCDerivativeOracleBase.sol: Abstract Contract (1): 144 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxCVXOracle.sol: Smart Contract (1): 43 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxChainlinkTwapOracle.sol: Smart Contract (1): 178 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxEETHOracleV2.sol: Smart Contract (1): 50 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxERC20OracleBase.sol: Abstract Contract (1): 225 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxEzETHOracleV2.sol: Smart Contract (1): 57 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxFrxETHOracleV2.sol: Smart Contract (1): 46 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxLSDOracleV2Base.sol: Abstract Contract (1): 229 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxSpotOracleBase.sol: Abstract Contract (1): 150 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxStETHOracleV2.sol: Smart Contract (1): 46 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/FxWBTCOracleV2.sol: Smart Contract (1): 45 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxCVXTwapOracle.sol: Smart Contract (1): 86 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxEETHTwapOracle.sol: Smart Contract (1): 88 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxEzETHTwapOracle.sol: Smart Contract (1): 92 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxFrxETHTwapOracle.sol: Smart Contract (1): 104 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxLSDOracleBase.sol: Abstract Contract (1): 79 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxPxETHTwapOracle.sol: Smart Contract (1): 97 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxStETHTwapOracle.sol: Smart Contract (1): 121 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxTwapOracleBase.sol: Abstract Contract (1): 67 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/oracle/archived/FxWBTCTwapOracle.sol: Smart Contract (1): 110 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rate-provider/BalancerV2CachedRateProvider.sol: Smart Contract (1): 29 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rate-provider/ERC4626RateProvider.sol: Smart Contract (1): 23 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rate-provider/WstETHRateProvider.sol: Smart Contract (1): 22 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rate-provider/wBETHRateProvider.sol: Smart Contract (1): 22 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/BoostableRebalancePool.sol: Smart Contract (1): 671 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/FxUSDShareableRebalancePool.sol: Smart Contract (1): 30 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/RebalancePool.sol: Smart Contract (1): 979 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/RebalancePoolGaugeClaimer.sol: Smart Contract (1): 272 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/RebalancePoolRegistry.sol: Smart Contract (1): 74 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/RebalancePoolSplitter.sol: Smart Contract (1), Interface (1): 231 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/RewardTokenWrapper.sol: Smart Contract (1): 73 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/ShareableRebalancePool.sol: Smart Contract (1): 911 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalance-pool/ShareableRebalancePoolV2.sol: Smart Contract (1): 83 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalancer/FxUSDRebalancer.sol: Smart Contract (1): 121 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/rebalancer/RebalanceWithBonusToken.sol: Smart Contract (1): 39 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/reserve-pool/ReservePool.sol: Smart Contract (1): 139 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/reserve-pool/ReservePoolV2.sol: Smart Contract (1): 153 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/steth/stETHGateway.sol: Smart Contract (1): 118 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/steth/stETHTreasury.sol: Smart Contract (1): 51 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v1/FractionalToken.sol: Smart Contract (1): 114 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v1/HarvestableTreasury.sol: Abstract Contract (1): 161 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v1/LeveragedToken.sol: Smart Contract (1): 90 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v1/Market.sol: Smart Contract (1): 824 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v1/Treasury.sol: Smart Contract (1): 670 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v1/WrappedTokenTreasury.sol: Smart Contract (1): 31 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v2/FractionalTokenV2.sol: Smart Contract (1): 78 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v2/FxInitialFund.sol: Smart Contract (1): 194 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v2/FxUSD.sol: Smart Contract (1): 503 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v2/LeveragedTokenV2.sol: Smart Contract (1): 178 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v2/MarketV2.sol: Smart Contract (1): 729 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v2/TreasuryV2.sol: Abstract Contract (1): 685 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/v2/WrappedTokenTreasuryV2.sol: Smart Contract (1): 101 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/wrapper/FxTokenBalancerV2Wrapper.sol: Smart Contract (1): 109 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/wrapper/LeveragedTokenWrapper.sol: Smart Contract (1): 74 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/wrapper/StETHAndxETHWrapper.sol: Smart Contract (1): 78 lines
../input_repos/aladdin-v3-contracts/contracts/f(x)/wrapper/wstETHWrapper.sol: Smart Contract (1): 54 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/AllInOneGateway.sol: Smart Contract (1), Interface (2): 308 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/BalancerLPGaugeGateway.sol: Smart Contract (1), Interface (1): 173 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/CLeverGateway.sol: Smart Contract (1), Interface (1): 91 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/CompounderGateway.sol: Smart Contract (1): 84 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/ConcentratorGateway.sol: Smart Contract (1): 54 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/FxGateway.sol: Smart Contract (1): 299 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/VeFeeGateway.sol: Smart Contract (1): 78 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/ZapGatewayBase.sol: Abstract Contract (1): 109 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/facets/ERC5115CompounderFacet.sol: Smart Contract (1): 55 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/facets/FxMarketV1Facet.sol: Smart Contract (1): 167 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/facets/FxUSDFacet.sol: Smart Contract (1): 392 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/facets/TokenConvertManagementFacet.sol: Smart Contract (1): 61 lines
../input_repos/aladdin-v3-contracts/contracts/gateways/libraries/LibGatewayRouter.sol: Library (1): 237 lines
../input_repos/aladdin-v3-contracts/contracts/harvester/concentrator/CLeverAMOHarvesterFacet.sol: Smart Contract (1): 18 lines
../input_repos/aladdin-v3-contracts/contracts/harvester/concentrator/ConcentratorHarvesterFacet.sol: Smart Contract (1): 100 lines
../input_repos/aladdin-v3-contracts/contracts/harvester/concentrator/FxUSDCompounderHarvestFacet.sol: Smart Contract (1): 25 lines
../input_repos/aladdin-v3-contracts/contracts/harvester/concentrator/StakeDaoHarvesterFacet.sol: Smart Contract (1): 33 lines
../input_repos/aladdin-v3-contracts/contracts/harvester/libraries/LibConcentratorHarvester.sol: Library (1): 106 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/MultiMerkleStash.sol: Smart Contract (1): 110 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/PlatformFeeSpliter.sol: Smart Contract (1): 285 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/TokenSale.sol: Smart Contract (1): 360 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/burner/BurnerBase.sol: Abstract Contract (1): 91 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/burner/ConvexFraxCompounderBurner.sol: Smart Contract (1), Interface (1): 33 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/burner/PlatformFeeBurner.sol: Smart Contract (1): 96 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/burner/StakeDAOCompounderBurner.sol: Smart Contract (1), Interface (1): 32 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/ConverterBase.sol: Abstract Contract (1): 112 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/ConverterRegistry.sol: Smart Contract (1): 85 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/CurveNGConverter.sol: Smart Contract (1): 117 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/ETHLSDConverter.sol: Smart Contract (1): 313 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/GeneralTokenConverter.sol: Smart Contract (1): 760 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/GeneralTokenConverterStorage.sol: Abstract Contract (1): 83 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/IConverterRegistry.sol: Interface (1): 42 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/ITokenConverter.sol: Interface (1): 117 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/LidoConverter.sol: Smart Contract (1): 114 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/MultiPathConverter.sol: Smart Contract (1): 56 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/UniswapV3Converter.sol: Smart Contract (1): 111 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/converter/WETHConverter.sol: Smart Contract (1): 65 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/IVesting.sol: Interface (1): 66 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/IVestingManager.sol: Interface (1): 46 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/ManageableVesting.sol: Smart Contract (1): 338 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/MultipleVestHelper.sol: Smart Contract (1): 35 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/Vesting.sol: Smart Contract (1): 142 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/VestingManagerProxy.sol: Smart Contract (1): 28 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/management/CvxFxnVestingManager.sol: Smart Contract (1): 81 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/management/PlainVestingManager.sol: Smart Contract (1): 64 lines
../input_repos/aladdin-v3-contracts/contracts/helpers/vesting/management/SdFxnVestingManager.sol: Smart Contract (1): 83 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IBalancerPool.sol: Interface (1): 56 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IBalancerV1Pool.sol: Interface (1): 76 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IBalancerVault.sol: Interface (1): 117 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IBalancerWeightedPoolFactory.sol: Interface (1): 14 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICVXMining.sol: Interface (1): 8 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IConvexBasicRewards.sol: Interface (1): 37 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IConvexBooster.sol: Interface (1): 29 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IConvexCRVDepositor.sol: Interface (1): 15 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IConvexCVXLocker.sol: Interface (1): 45 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IConvexCVXRewardPool.sol: Interface (1): 27 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IConvexFraxBooster.sol: Interface (1): 7 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IConvexToken.sol: Interface (1): 11 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurve3Pool.sol: Interface (1): 23 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveAPool.sol: Interface (4): 90 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveBasePool.sol: Interface (4): 72 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveCryptoPool.sol: Interface (3): 173 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveCryptoPoolFactory.sol: Interface (1): 23 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveETHPool.sol: Interface (1): 43 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveFactoryMetaPool.sol: Interface (4): 102 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveFactoryPlainPool.sol: Interface (4): 57 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveFactoryPool.sol: Interface (1): 22 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveGauge.sol: Interface (1): 84 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveMetaPool.sol: Interface (2): 89 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurvePlainPool.sol: Interface (1): 58 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurvePlainPoolFactory.sol: Interface (1): 17 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurvePoolOracle.sol: Interface (1): 52 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurvePoolRegistry.sol: Interface (1): 11 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveTokenMinter.sol: Interface (1): 24 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveTriCrypto.sol: Interface (1): 24 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveV2Pool.sol: Interface (1): 29 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveVoteEscrow.sol: Interface (1): 38 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICurveYPool.sol: Interface (8): 104 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ICvxCrvStakingWrapper.sol: Interface (1): 49 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IERC20Metadata.sol: Interface (1): 11 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IEllipsisMerkleDistributor.sol: Interface (1): 12 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IFraxUnifiedFarm.sol: Interface (1): 49 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ILidoStETH.sol: Interface (1): 17 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ILidoWstETH.sol: Interface (1): 54 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IMultiMerkleStash.sol: Interface (1): 30 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IMulticall2.sol: Interface (1): 29 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/ISnapshotDelegateRegistry.sol: Interface (1): 22 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IStakeDAOSdTokenDepositor.sol: Interface (1): 44 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IStakingProxyConvex.sol: Interface (1): 66 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IUniswapV2Pair.sol: Interface (1): 31 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IUniswapV3Pool.sol: Interface (1): 65 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IUniswapV3Quoter.sol: Interface (1): 52 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IUniswapV3Router.sol: Interface (1): 19 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IWBETH.sol: Interface (1): 24 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IWETH.sol: Interface (1): 9 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/IZap.sol: Interface (1): 27 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/balancer/IFlashLoanRecipient.sol: Interface (1): 21 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/clever/ICLeverCVXLocker.sol: Interface (1): 39 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/clever/ICLeverToken.sol: Interface (1): 13 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/clever/ICLeverYieldStrategy.sol: Interface (1): 87 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/clever/IFurnace.sol: Interface (1): 48 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/clever/IMetaCLever.sol: Interface (1): 105 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/clever/IMetaFurnace.sol: Interface (1): 59 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IAladdinCRV.sol: Interface (1): 92 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IAladdinCRVConvexVault.sol: Interface (1): 212 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IAladdinCompounder.sol: Interface (1): 146 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorBase.sol: Interface (1): 86 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorCompounder.sol: Interface (1): 81 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorConvexVault.sol: Interface (1): 173 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorGeneralVault.sol: Interface (1): 186 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorSdCrvGaugeWrapper.sol: Interface (1): 37 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorSdCrvVault.sol: Interface (1): 43 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorStakeDAOGaugeWrapper.sol: Interface (1): 142 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorStakeDAOLocker.sol: Interface (1): 40 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorStakeDAOVault.sol: Interface (1): 76 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IConcentratorStrategy.sol: Interface (1): 57 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/ICvxFxnCompounder.sol: Interface (1): 37 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/ICvxFxsCompounder.sol: Interface (1): 24 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IFxUSDCompounder.sol: Interface (1): 66 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/ISdCRVLocker.sol: Interface (1): 37 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/ISdCrvCompounder.sol: Interface (1): 31 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/concentrator/IStakeDAOBoostDelegation.sol: Interface (1): 54 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/ICommitUserSurrogate.sol: Interface (1): 7 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/IConvexFXNBooster.sol: Interface (1): 7 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/IConvexFXNDepositor.sol: Interface (1): 12 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/IConvexFXSDepositor.sol: Interface (1): 15 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/IConvexGaugeVotePlatform.sol: Interface (1): 40 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/IConvexVirtualBalanceRewardPool.sol: Interface (1): 9 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/ICvxFxnStaking.sol: Interface (1): 52 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/ICvxFxsStaking.sol: Interface (1): 50 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/ICvxRewardPool.sol: Interface (1): 33 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/IStakingProxyRebalancePool.sol: Interface (1): 51 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/convex/IStashTokenWrapper.sol: Interface (1): 23 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/curve/ICrvUSDAmm.sol: Interface (1): 7 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/curve/ICurveStableSwapMetaNG.sol: Interface (1): 305 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/curve/ICurveStableSwapNG.sol: Interface (1): 255 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/etherfi/IEtherFiLiquidityPool.sol: Interface (1): 34 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/etherfi/IEtherFiWeETH.sol: Interface (1): 29 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IAssetStrategy.sol: Interface (1): 9 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxBoostableRebalancePool.sol: Interface (1): 108 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxFractionalToken.sol: Interface (1): 40 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxFractionalTokenV2.sol: Interface (1): 33 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxLeveragedToken.sol: Interface (1): 26 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxLeveragedTokenV2.sol: Interface (1): 33 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxMarket.sol: Interface (1): 139 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxMarketV2.sol: Interface (1): 250 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxPriceOracle.sol: Interface (1): 20 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxPriceOracleV2.sol: Interface (1): 33 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxRateProvider.sol: Interface (1): 9 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxRebalancePool.sol: Interface (1): 128 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxRebalancePoolRegistry.sol: Interface (1): 27 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxRebalancePoolSplitter.sol: Interface (1): 58 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxReservePool.sol: Interface (1): 16 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxShareableRebalancePool.sol: Interface (1): 79 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxTokenWrapper.sol: Interface (1): 27 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxTreasury.sol: Interface (1): 192 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxTreasuryV2.sol: Interface (1): 248 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/f(x)/IFxUSD.sol: Interface (1): 197 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/frax-finance/IFrxETHMinter.sol: Interface (1): 22 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/kelp-dao/IKelpDAOLRTConfig.sol: Interface (1): 10 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/kelp-dao/IKelpDAOLRTDepositPool.sol: Interface (1): 30 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/paladin/IMultiMerkleDistributor.sol: Interface (1): 77 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/pendle/IStandardizedYield.sol: Interface (1): 171 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/pirex/IPirexEthMinter.sol: Interface (1): 86 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/renzo-protocol/IRenzoOracle.sol: Interface (1): 35 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/renzo-protocol/IRenzoRestakeManager.sol: Interface (1): 63 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/rocket-pool/IRocketDepositPool.sol: Interface (1): 9 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/rocket-pool/IRocketTokenRETH.sol: Interface (1): 24 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/stakedao/IStakeDAOCRVDepositor.sol: Interface (1): 36 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/IFeeDistributor.sol: Interface (1): 35 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/IGaugeController.sol: Interface (1): 85 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/IGovernanceToken.sol: Interface (1): 11 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/ILiquidityGauge.sol: Interface (1): 152 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/ILiquidityManager.sol: Interface (1): 98 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/ISharedLiquidityGauge.sol: Interface (1): 73 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/ITokenMinter.sol: Interface (1): 21 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/IVotingEscrow.sol: Interface (1): 89 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/IVotingEscrowBoost.sol: Interface (1): 246 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/IVotingEscrowHelper.sol: Interface (1): 41 lines
../input_repos/aladdin-v3-contracts/contracts/interfaces/voting-escrow/IVotingEscrowProxy.sol: Interface (1): 25 lines
../input_repos/aladdin-v3-contracts/contracts/misc/EmptyContract.sol: Smart Contract (1): 5 lines
../input_repos/aladdin-v3-contracts/contracts/misc/GaugeRewardDistributor.sol: Smart Contract (1): 307 lines
../input_repos/aladdin-v3-contracts/contracts/misc/ICurveGaugeV4V5.sol: Interface (1): 51 lines
../input_repos/aladdin-v3-contracts/contracts/misc/PlatformFeeDistributor.sol: Smart Contract (1): 209 lines
../input_repos/aladdin-v3-contracts/contracts/misc/RewardClaimHelper.sol: Smart Contract (1), Interface (2): 34 lines
../input_repos/aladdin-v3-contracts/contracts/misc/checker/CurveBasePoolChecker.sol: Smart Contract (1): 78 lines
../input_repos/aladdin-v3-contracts/contracts/misc/checker/CurveMetaPoolChecker.sol: Smart Contract (1): 68 lines
../input_repos/aladdin-v3-contracts/contracts/misc/checker/IPriceChecker.sol: Interface (1): 7 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockConcentratorGeneralVault.sol: Smart Contract (1): 72 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockConvexBasicRewards.sol: Smart Contract (1): 18 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockCurveGaugeV1V2V3.sol: Smart Contract (1): 11 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockCurveGaugeV4V5.sol: Smart Contract (1): 13 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockERC20.sol: Smart Contract (1): 19 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockFurnace.sol: Smart Contract (1): 61 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockTokenWrapper.sol: Smart Contract (1): 44 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockTwapOracle.sol: Smart Contract (1): 57 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockVestingManager.sol: Smart Contract (1): 54 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockYieldStrategy.sol: Smart Contract (1): 52 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockYieldStrategyForCLever.sol: Smart Contract (1): 88 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/MockYieldToken.sol: Smart Contract (1): 49 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/WETH9.sol: Smart Contract (1): 69 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/fees/MockCustomFeeRate.sol: Smart Contract (1): 15 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/math/MockDecrementalFloatingPoint.sol: Smart Contract (1): 37 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/math/MockExponentialMovingAverage.sol: Smart Contract (1): 26 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/math/MockGaussElimination.sol: Smart Contract (1): 12 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/math/MockLogExpMath.sol: Smart Contract (1): 11 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/rewards/MockLinearMultipleRewardDistributor.sol: Smart Contract (1): 21 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/rewards/MockLinearRewardDistributor.sol: Smart Contract (1): 21 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/rewards/MockMultipleRewardAccumulator.sol: Smart Contract (1): 52 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/common/rewards/MockMultipleRewardCompoundingAccumulator.sol: Smart Contract (1): 56 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/concentrator/MockConcentratorBaseV2.sol: Smart Contract (1): 25 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/concentrator/MockConcentratorCompounderBase.sol: Smart Contract (1): 63 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/concentrator/strategies/MockConcentratorStrategy.sol: Smart Contract (1): 37 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/f(x)/MockFxRateProvider.sol: Smart Contract (1): 13 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/f(x)/MockTreasuryForClaimer.sol: Smart Contract (1): 11 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/voting-escrow/gauges/MockLiquidityGauge.sol: Smart Contract (1): 23 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/voting-escrow/manager/MockLiquidityManager.sol: Smart Contract (1): 33 lines
../input_repos/aladdin-v3-contracts/contracts/mocks/voting-escrow/manager/MockLiquidityManagerImmutable.sol: Smart Contract (1): 37 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/AladdinPriceOracle.sol: Smart Contract (1): 40 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/ChainlinkPriceOracle.sol: Smart Contract (1): 60 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/CurveBasePoolPriceOracle.sol: Smart Contract (1): 86 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/CurveV2PriceOracle.sol: Smart Contract (1): 91 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/interfaces/AggregatorV3Interface.sol: Interface (1): 35 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/interfaces/IPriceOracle.sol: Interface (1): 13 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/interfaces/ISpotPriceOracle.sol: Interface (1): 53 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/interfaces/ITwapOracle.sol: Interface (1): 14 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/spot/SpotPriceOracle.sol: Smart Contract (1): 363 lines
../input_repos/aladdin-v3-contracts/contracts/price-oracle/twap/ChainlinkTwapOracleV3.sol: Smart Contract (1): 196 lines
../input_repos/aladdin-v3-contracts/contracts/voting/CLeverHermezSurrogate.sol: Smart Contract (1): 115 lines
../input_repos/aladdin-v3-contracts/contracts/voting/ISignatureVerifier.sol: Interface (1): 7 lines
../input_repos/aladdin-v3-contracts/contracts/voting/SignatureVerifier.sol: Smart Contract (1): 26 lines
../input_repos/aladdin-v3-contracts/contracts/voting/VoteProxy.sol: Smart Contract (1): 34 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/FeeDistributorAdmin.sol: Smart Contract (1): 63 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/SmartWalletWhitelist.sol: Smart Contract (1), Interface (1): 53 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/VotingEscrowBoost.sol: Smart Contract (1): 462 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/VotingEscrowHelper.sol: Smart Contract (1): 300 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/VotingEscrowProxy.sol: Smart Contract (1): 71 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/gauges/GaugeControllerOwner.sol: Smart Contract (1): 231 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/gauges/liquidity/DelegatedLiquidityGauge.sol: Smart Contract (1): 33 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/gauges/liquidity/LiquidityGauge.sol: Smart Contract (1): 539 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/gauges/liquidity/SharedLiquidityGauge.sol: Smart Contract (1): 195 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/manager/ConvexCurveManager.sol: Smart Contract (1): 226 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/manager/LiquidityManagerBase.sol: Abstract Contract (1): 212 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/manager/immutable/ConvexCurveManagerImmutable.sol: Smart Contract (1): 214 lines
../input_repos/aladdin-v3-contracts/contracts/voting-escrow/manager/immutable/LiquidityManagerBaseImmutable.sol: Abstract Contract (1): 204 lines
../input_repos/aladdin-v3-contracts/contracts/zap/AladdinCRVZap.sol: Smart Contract (1): 201 lines
../input_repos/aladdin-v3-contracts/contracts/zap/AladdinConvexVaultZap.sol: Smart Contract (1): 283 lines
../input_repos/aladdin-v3-contracts/contracts/zap/AladdinZap.sol: Smart Contract (1): 155 lines
../input_repos/aladdin-v3-contracts/contracts/zap/TokenZapLogic.sol: Smart Contract (1): 622 lines

Summary:
Total number of Solidity files: 416
Abstract Contracts Files: 40
Smart Contracts Files: 191
Libraries Files: 15
Interfaces Files: 157
Multiple Types Files: 12
******************************************
Abstract Contracts: 40
Smart Contracts: 203
Libraries: 15
Interfaces: 193
LOC: 56113
