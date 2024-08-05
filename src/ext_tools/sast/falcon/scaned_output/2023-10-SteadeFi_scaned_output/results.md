'npx hardhat clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-10-SteadeFi)
'npx hardhat clean --global' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-10-SteadeFi)
'npx hardhat compile --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2023-10-SteadeFi)
Traceback (most recent call last):
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 835, in main_impl
    ) = process_all(filename, args, detector_classes, printer_classes)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 87, in process_all
    compilations = compile_all(target, **vars(args))
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/crytic_compile.py", line 722, in compile_all
    compilations.append(CryticCompile(target, **kwargs))
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/crytic_compile.py", line 211, in __init__
    self._compile(**kwargs)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/crytic_compile.py", line 633, in _compile
    self._platform.compile(self, **kwargs)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/platform/hardhat.py", line 183, in compile
    hardhat_like_parsing(crytic_compile, self._target, build_directory, hardhat_working_dir)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/platform/hardhat.py", line 97, in hardhat_like_parsing
    path = convert_filename(
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/utils/naming.py", line 169, in convert_filename
    filename = _verify_filename_existence(filename, cwd)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/utils/naming.py", line 119, in _verify_filename_existence
    raise InvalidCompilation(f"Unknown file: {filename}")
crytic_compile.platform.exceptions.InvalidCompilation: Unknown file: forge-std/console.sol
None
Error in .
Traceback (most recent call last):
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 835, in main_impl
    ) = process_all(filename, args, detector_classes, printer_classes)
  File "/usr/local/lib/python3.10/dist-packages/falcon_analyzer-0.2.28-py3.10.egg/falcon/__main__.py", line 87, in process_all
    compilations = compile_all(target, **vars(args))
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/crytic_compile.py", line 722, in compile_all
    compilations.append(CryticCompile(target, **kwargs))
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/crytic_compile.py", line 211, in __init__
    self._compile(**kwargs)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/crytic_compile.py", line 633, in _compile
    self._platform.compile(self, **kwargs)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/platform/hardhat.py", line 183, in compile
    hardhat_like_parsing(crytic_compile, self._target, build_directory, hardhat_working_dir)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/platform/hardhat.py", line 97, in hardhat_like_parsing
    path = convert_filename(
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/utils/naming.py", line 169, in convert_filename
    filename = _verify_filename_existence(filename, cwd)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/utils/naming.py", line 119, in _verify_filename_existence
    raise InvalidCompilation(f"Unknown file: {filename}")
crytic_compile.platform.exceptions.InvalidCompilation: Unknown file: forge-std/console.sol

INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/2023-10-SteadeFi_scaned_output/mwe-output.json generate success.
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
