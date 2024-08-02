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
func compileRepositories(directoriesWithFrameworks map[string][]string) map[string]CompilationResult {
	//	var anyFrameworkSucceeded bool
	results := make(map[string]CompilationResult)

	for directoryWithFramework, frameworks := range directoriesWithFrameworks {
		fmt.Printf("Compiling in directory: %s\n", directoryWithFramework)

		if len(frameworks) == 0 {
			fmt.Printf("No frameworks specified for directory: %s\n", directoryWithFramework)
			results[directoryWithFramework] = CompilationResult{Success: false, Err: fmt.Errorf("no frameworks specified")}
			continue
		}

		//var anyFrameworkSucceeded bool
		var hardhatSuccess, foundrySuccess bool

		// There could be multiple Frameworks in on Directory
		for _, framework := range frameworks {
			var result CompilationResult

			switch framework {
			case "hardhat":
				result = compileHardhat(directoryWithFramework)
				if result.Success {
					hardhatSuccess = true
					fmt.Printf("Compilation succeeded for %s in directory: %s\n", framework, directoryWithFramework)
				} else {
					fmt.Printf("Compilation failed for %s in directory %s: %v\n", framework, directoryWithFramework, result.Err)
				}

			case "foundry":
				result = compileFoundry(directoryWithFramework)
				if result.Success {
					foundrySuccess = true
					fmt.Printf("Compilation succeeded for %s in directory: %s\n", framework, directoryWithFramework)
				} else {
					fmt.Printf("Compilation failed for %s in directory %s: %v\n", framework, directoryWithFramework, result.Err)
				}

			default:
				result.Err = fmt.Errorf("unknown framework: %s", framework)
			}

			// if result.Success {
			// 	anyFrameworkSucceeded = true
			// }
		}

		// If code compilese with both frameworks we will give priority to Hardhad else we will try others as well

		if hardhatSuccess && foundrySuccess {
			
			// Run the tool using Hardhat as the chosen framework
			if err := runTool(tool, directoryWithFramework, outputDirPath, "hardhat", results); err != nil {
				fmt.Printf("Error running Falcon in directory %s with Hardhat: %v\n", directoryWithFramework, err)
				results[directoryWithFramework] = CompilationResult{Success: false, Err: err}
			} else {
				results[directoryWithFramework] = CompilationResult{Success: true}
			}
		} else if hardhatSuccess {
			// Run the tool using Hardhat framework
			if err := runTool(tool, directoryWithFramework, outputDirPath, "hardhat", results); err != nil {
				fmt.Printf("Error running Falcon in directory %s with Hardhat: %v\n", directoryWithFramework, err)
				results[directoryWithFramework] = CompilationResult{Success: false, Err: err}
			} else {
				results[directoryWithFramework] = CompilationResult{Success: true}
			}
		} else if foundrySuccess {
			// Run the tool using Foundry framework
			if err := runTool(tool, directoryWithFramework, outputDirPath, "foundry", results); err != nil {
				fmt.Printf("Error running Falcon in directory %s with Foundry: %v\n", directoryWithFramework, err)
				results[directoryWithFramework] = CompilationResult{Success: false, Err: err}
			} else {
				results[directoryWithFramework] = CompilationResult{Success: true}
			}
		} else {
			results[directoryWithFramework] = CompilationResult{Success: false, Err: fmt.Errorf("all frameworks failed")}
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
