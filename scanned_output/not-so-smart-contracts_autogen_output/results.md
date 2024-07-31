'npx hardhat clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/not-so-smart-contracts)
'npx hardhat clean --global' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/not-so-smart-contracts)
'npx hardhat compile --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/input_repos/not-so-smart-contracts)
'npx' returned non-zero exit code 1
Downloading compiler 0.8.24
stdout: Downloading compiler 0.4.15
stdout: Downloading compiler 0.4.26
ParserError: Expected a state variable declaration. If you intended this as a fallback function or a function to handle plain ether transactions, use the "fallback" keyword or the "receive" keyword instead.
stderr:   --> contracts/contracts/bad_randomness/theRun_source_code/theRun.sol:34:20:
stderr:    |
stderr: 34 |         function() {
stderr:    |                    ^
stderr: 
stderr: 
stderr: ParserError: Expected '{' but got 'constant'
stderr:   --> contracts/honeypots/Lottery/Lottery.sol:59:49:
stderr:    |
stderr: 59 |     function luckyNumberOfAddress(address addr) constant returns(uint n){
stderr:    |                                                 ^^^^^^^^
stderr: 
stderr: 
stderr: ParserError: Expected '{' but got 'constant'
stderr:   --> contracts/reentrancy/DAO_source_code/DAO.sol:47:40:
stderr:    |
stderr: 47 |     function balanceOf(address _owner) constant returns (uint256 balance);
stderr:    |                                        ^^^^^^^^
stderr: 
stderr: 
stderr: Warning: SPDX license identifier not provided in source file. Before publishing, consider adding a comment containing "SPDX-License-Identifier: <SPDX-License>" to each source file. Use "SPDX-License-Identifier: UNLICENSED" for non-open-source code. Please see https://spdx.org for more information.
stderr: --> contracts/variable shadowing/inherited_state.sol
stderr: 
stderr: 
stderr: contracts/forced_ether_reception/coin.sol:8:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function owned() public {
stderr:     ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/forced_ether_reception/coin.sol:47:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function TokenERC20(
stderr:     ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/forced_ether_reception/coin.sol:153:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function MyAdvancedToken(
stderr:     ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/honeypots/KOTH/KOTH.sol:7:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function Ownable() public {owner = msg.sender;}
stderr:     ^---------------------------------------------^
stderr: 
stderr: contracts/honeypots/PrivateBank/PrivateBank.sol:11:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function Private_Bank(address _log)
stderr:     ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/honeypots/VarLoop/VarLoop.sol:24:13: Warning: Use of the "var" keyword is deprecated.
stderr:             var i1 = 1;
stderr:             ^----^
stderr: 
stderr: contracts/honeypots/VarLoop/VarLoop.sol:25:13: Warning: Use of the "var" keyword is deprecated.
stderr:             var i2 = 0;
stderr:             ^----^
stderr: 
stderr: contracts/honeypots/VarLoop/VarLoop.sol:26:13: Warning: Use of the "var" keyword is deprecated.
stderr:             var amX2 = msg.value*2;
stderr:             ^------^
stderr: 
stderr: contracts/race_condition/RaceCondition.sol:20:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function RaceCondition(uint _price, ERC20 _token)
stderr:     ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/reentrancy/Reentrancy.sol:18:13: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:             throw;
stderr:             ^---^
stderr: 
stderr: contracts/reentrancy/Reentrancy.sol:29:13: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:             throw;
stderr:             ^---^
stderr: 
stderr: contracts/reentrancy/ReentrancyExploit.sol:8:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function ReentranceExploit() public{
stderr:     ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/unchecked_external_call/KotET_source_code/KingOfTheEtherThrone.sol:65:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function KingOfTheEtherThrone() {
stderr:     ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/unprotected_function/Unprotected.sol:11:5: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:     function Unprotected()
stderr:     ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/unprotected_function/WalletLibrary_source_code/WalletLibrary.sol:126:5: Warning: Use of the "var" keyword is deprecated.
stderr:     var pending = m_pending[_operation];
stderr:     ^---------^
stderr: 
stderr: contracts/unprotected_function/WalletLibrary_source_code/WalletLibrary.sol:190:5: Warning: Use of the "var" keyword is deprecated.
stderr:     var pending = m_pending[_operation];
stderr:     ^---------^
stderr: 
stderr: contracts/unprotected_function/WalletLibrary_source_code/WalletLibrary.sol:240:11: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:           throw;
stderr:           ^---^
stderr: 
stderr: contracts/unprotected_function/WalletLibrary_source_code/WalletLibrary.sol:274:11: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:           throw;
stderr:           ^---^
stderr: 
stderr: contracts/unprotected_function/WalletLibrary_source_code/WalletLibrary.sol:291:5: Warning: Use of the "var" keyword is deprecated.
stderr:     var pending = m_pending[_operation];
stderr:     ^---------^
stderr: 
stderr: contracts/unprotected_function/WalletLibrary_source_code/WalletLibrary.sol:400:3: Warning: Defining constructors as functions with the same name as the contract is deprecated. Use "constructor(...) { ... }" instead.
stderr:   function Wallet(address[] _owners, uint _required, uint _daylimit) {
stderr:   ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/wrong_constructor_name/Rubixi_source_code/Rubixi.sol:76:41: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:                 if (collectedFees == 0) throw;
stderr:                                         ^---^
stderr: 
stderr: contracts/wrong_constructor_name/Rubixi_source_code/Rubixi.sol:86:41: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:                 if (collectedFees == 0) throw;
stderr:                                         ^---^
stderr: 
stderr: contracts/wrong_constructor_name/Rubixi_source_code/Rubixi.sol:93:57: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:                 if (collectedFees == 0 || _pcent > 100) throw;
stderr:                                                         ^---^
stderr: 
stderr: contracts/wrong_constructor_name/Rubixi_source_code/Rubixi.sol:106:49: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:                 if (_mult > 300 || _mult < 120) throw;
stderr:                                                 ^---^
stderr: 
stderr: contracts/wrong_constructor_name/Rubixi_source_code/Rubixi.sol:112:32: Warning: "throw" is deprecated in favour of "revert()", "require()" and "assert()".
stderr:                 if (_fee > 10) throw;
stderr:                                ^---^
stderr: 
stderr: contracts/race_condition/RaceCondition.sol:5:46: Warning: This declaration shadows an existing declaration.
stderr:     function totalSupply() constant returns (uint totalSupply);
stderr:                                              ^--------------^
stderr: contracts/race_condition/RaceCondition.sol:5:5: The shadowed declaration is here:
stderr:     function totalSupply() constant returns (uint totalSupply);
stderr:     ^---------------------------------------------------------^
stderr: 
stderr: 
stderr: contracts/unprotected_function/WalletLibrary_source_code/WalletLibrary.sol:258:3: Warning: Variable is shadowed in inline assembly by an instruction of the same name
stderr:   function create(uint _value, bytes _code) internal returns (address o_addr) {
stderr:   ^ (Relevant source part starts here and spans across multiple lines).
stderr: 
stderr: contracts/denial_of_service/auction.sol:51:5: Warning: Failure condition of 'send' ignored. Consider using 'transfer' instead.
stderr:     msg.sender.send(refund);
stderr:     ^---------------------^
stderr: 
stderr: contracts/denial_of_service/list_dos.sol:9:7: TypeError: No matching declaration found after argument-dependent lookup.
stderr:       require(refundAddresses[i].transfer(refundAmount[refundAddresses[i]]));
stderr:       ^-----^
stderr: 
stderr: Error HH600: Compilation failed
stderr: For more info go to https://hardhat.org/HH600 or run Hardhat with --show-stack-traces

transfer(address,uint256) has possible integer overflow/underflow:
	- balances[_to] += _value (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#70)
transferFrom(address,address,uint256) has possible integer overflow/underflow:
	- balances[_to] += _value (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#79)
Reference: https://swcregistry.io/docs/SWC-101

Setter function HumanStandardToken.HumanStandardToken(uint256,string,uint8,string) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#131-142) does not emit an event
Reference: https://github.com/pessimistic-io/slitherin/blob/master/docs/event_setter.md

variable lacks a zero-check on 		- StandardToken.transfer(address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#63-73)
variable lacks a zero-check on 		- StandardToken.transferFrom(address,address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#75-84)
variable lacks a zero-check on 		- StandardToken.approve(address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#90-94)
variable lacks a zero-check on 		- HumanStandardToken.approveAndCall(address,uint256,bytes) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#145-154)
Reference: https://github.com/crytic/slither/wiki/Detector-Documentation#missing-zero-address-validation

Deprecated standard detected require(bool)(_spender.call(bytes4(bytes32(sha3()(receiveApproval(address,uint256,address,bytes)))),msg.sender,_value,this,_extraData)) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#152):
	- Usage of "sha3()" should be replaced with "keccak256()"
Reference:  

require() missing error messages
	 - require(bool)(balances[msg.sender] >= _value) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#68)
require() missing error messages
	 - require(bool)(balances[_from] >= _value && allowed[_from][msg.sender] >= _value) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#78)
require() missing error messages
	 - require(bool)(_spender.call(bytes4(bytes32(sha3()(receiveApproval(address,uint256,address,bytes)))),msg.sender,_value,this,_extraData)) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#152)
Reference: https://dev.to/tawseef/require-vs-assert-in-solidity-5e9d

function:Token.balanceOf(address) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#22)is public and can be replaced with external 
function:Token.transfer(address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#28)is public and can be replaced with external 
function:Token.transferFrom(address,address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#35)is public and can be replaced with external 
function:Token.approve(address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#41)is public and can be replaced with external 
function:Token.allowance(address,address) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#46)is public and can be replaced with external 
function:StandardToken.transfer(address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#63-73)is public and can be replaced with external 
function:StandardToken.transferFrom(address,address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#75-84)is public and can be replaced with external 
function:StandardToken.balanceOf(address) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#86-88)is public and can be replaced with external 
function:StandardToken.approve(address,uint256) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#90-94)is public and can be replaced with external 
function:StandardToken.allowance(address,address) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#96-98)is public and can be replaced with external 
function:HumanStandardToken.approveAndCall(address,uint256,bytes) (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#145-154)is public and can be replaced with external 
Reference:  

HumanStandardToken.version (contracts/reentrancy/SpankChain_source_code/SpankChain.sol#129) should be constant
Reference:  
. analyzed (3 contracts with 86 detectors), 23 result(s) found
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
FileNotFoundError: [Errno 2] No such file or directory: '../input_repos/not-so-smart-contracts/output.json'
