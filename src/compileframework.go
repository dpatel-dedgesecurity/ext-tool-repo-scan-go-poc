package main

import (
	"fmt"
	"os/exec"
)

// CompilationResult represents the result of a compilation attempt
type CompilationResult struct {
	Success bool
	Err     error
}

// compileRepositories compiles repositories based on the given map of directories and frameworks
func compileRepositories(directories map[string][]string) map[string]CompilationResult {
//	var anyFrameworkSucceeded bool
	results := make(map[string]CompilationResult)

	for directory, frameworks := range directories {
		fmt.Printf("Compiling in directory: %s\n", directory)

		if len(frameworks) == 0 {
			fmt.Printf("No frameworks specified for directory: %s\n", directory)
			results[directory] = CompilationResult{Success: false, Err: fmt.Errorf("no frameworks specified")}
			continue
		}

		//var anyFrameworkSucceeded bool
		var hardhatSuccess, foundrySuccess bool

		for _, framework := range frameworks {
			var result CompilationResult

			switch framework {
			case "hardhat":
				result = compileHardhat(directory)
				if result.Success {
					hardhatSuccess = true
					fmt.Printf("Compilation succeeded for %s in directory: %s\n", framework, directory)
				} else {
					fmt.Printf("Compilation failed for %s in directory %s: %v\n", framework, directory, result.Err)
				}

			case "foundry":
				result = compileFoundry(directory)
				if result.Success {
					foundrySuccess = true
					fmt.Printf("Compilation succeeded for %s in directory: %s\n", framework, directory)
				} else {
					fmt.Printf("Compilation failed for %s in directory %s: %v\n", framework, directory, result.Err)
				}

			default:
				result.Err = fmt.Errorf("unknown framework: %s", framework)
			}

			// if result.Success {
			// 	anyFrameworkSucceeded = true
			// }
		}

		if hardhatSuccess && foundrySuccess {
			// Run the tool using Hardhat as the chosen framework
			if err := runTool(tool, directory, outputDirPath, "hardhat", results); err != nil {
				fmt.Printf("Error running Falcon in directory %s with Hardhat: %v\n", directory, err)
				results[directory] = CompilationResult{Success: false, Err: err}
			} else {
				results[directory] = CompilationResult{Success: true}
			}
		} else if hardhatSuccess {
			// Run the tool using Hardhat framework
			if err := runTool(tool, directory, outputDirPath, "hardhat", results); err != nil {
				fmt.Printf("Error running Falcon in directory %s with Hardhat: %v\n", directory, err)
				results[directory] = CompilationResult{Success: false, Err: err}
			} else {
				results[directory] = CompilationResult{Success: true}
			}
		} else if foundrySuccess {
			// Run the tool using Foundry framework
			if err := runTool(tool, directory, outputDirPath, "foundry", results); err != nil {
				fmt.Printf("Error running Falcon in directory %s with Foundry: %v\n", directory, err)
				results[directory] = CompilationResult{Success: false, Err: err}
			} else {
				results[directory] = CompilationResult{Success: true}
			}
		} else {
			results[directory] = CompilationResult{Success: false, Err: fmt.Errorf("all frameworks failed")}
		}
	}

	return results
}

// compileHardhat compiles a Hardhat project
func compileHardhat(repoDirectory string) CompilationResult {
	cmd := exec.Command("npx", "hardhat", "compile", "--force")
	cmd.Dir = repoDirectory
	if err := cmd.Run(); err != nil {
		return CompilationResult{Success: false, Err: err}
	}
	return CompilationResult{Success: true}
}

// compileFoundry compiles a Foundry project
func compileFoundry(repoDirectory string) CompilationResult {
	cmd := exec.Command("forge", "build")
	cmd.Dir = repoDirectory
	if err := cmd.Run(); err != nil {
		return CompilationResult{Success: false, Err: err}
	}
	return CompilationResult{Success: true}
}
