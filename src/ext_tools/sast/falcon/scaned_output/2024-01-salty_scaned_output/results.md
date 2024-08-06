'forge clean' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-01-salty)
'forge config --json' running
'forge build --build-info --skip */test/** */script/** --force' running (wd: /home/im/dedge/ext-tool-repo-scan-go/src/ext_tools/sast/falcon/input_repos/2024-01-salty)

Reentrancy in Liquidizer._possiblyBurnUSDS() (src/stable/Liquidizer.sol#101-126):
	External calls:
	- _burnUSDS(usdsThatShouldBeBurned) (src/stable/Liquidizer.sol#112)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#122)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	External calls sending eth:
	- _burnUSDS(usdsThatShouldBeBurned) (src/stable/Liquidizer.sol#112)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
Reentrancy in Staking.transferStakedSaltFromAirdropToUser(address,uint256) (src/staking/Staking.sol#130-138):
	External calls:
	- _decreaseUserShare(msg.sender,PoolUtils.STAKED_SALT,amountToTransfer,false) (src/staking/Staking.sol#134)
		- returndata = address(token).functionCall(data,SafeERC20: low-level call failed) (lib/openzeppelin-contracts/contracts/token/ERC20/utils/SafeERC20.sol#122)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
	External calls sending eth:
	- _decreaseUserShare(msg.sender,PoolUtils.STAKED_SALT,amountToTransfer,false) (src/staking/Staking.sol#134)
		- (success,returndata) = target.call{value: value}(data) (lib/openzeppelin-contracts/contracts/utils/Address.sol#135)
Reference:  
