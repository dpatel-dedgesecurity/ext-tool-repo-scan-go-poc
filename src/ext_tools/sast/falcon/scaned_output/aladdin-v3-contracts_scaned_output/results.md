'npx hardhat clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/aladdin-v3-contracts)
'npx hardhat clean --global' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/aladdin-v3-contracts)
'npx hardhat compile --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/aladdin-v3-contracts)
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

INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/aladdin-v3-contracts_scaned_output/mwe-output.json generate success.
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
