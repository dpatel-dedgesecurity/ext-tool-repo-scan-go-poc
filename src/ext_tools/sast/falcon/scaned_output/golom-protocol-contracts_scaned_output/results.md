'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/golom-protocol-contracts)
'forge config --json' running
warning: Found unknown config section in foundry.toml: [default]
This notation for profiles has been deprecated and may result in the profile not being registered in future versions.
Please use [profile.default] instead or run `forge config --fix`.
'forge build --build-info --skip */test/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/golom-protocol-contracts)
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
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/platform/foundry.py", line 86, in compile
    hardhat_like_parsing(crytic_compile, self._target, build_directory, self._target)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/platform/hardhat.py", line 52, in hardhat_like_parsing
    raise InvalidCompilation(txt)
crytic_compile.platform.exceptions.InvalidCompilation: Compilation failed. Can you run build command?
out/build-info is not a directory.
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
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/platform/foundry.py", line 86, in compile
    hardhat_like_parsing(crytic_compile, self._target, build_directory, self._target)
  File "/home/im/.local/lib/python3.10/site-packages/crytic_compile/platform/hardhat.py", line 52, in hardhat_like_parsing
    raise InvalidCompilation(txt)
crytic_compile.platform.exceptions.InvalidCompilation: Compilation failed. Can you run build command?
out/build-info is not a directory.

INFO:Falcon:metatrust result: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/scaned_output/golom-protocol-contracts_scaned_output/mwe-output.json generate success.
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
