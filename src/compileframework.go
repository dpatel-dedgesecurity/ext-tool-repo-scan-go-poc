package main

import (
	"fmt"
	"os/exec"
)

type CompilationResult struct {
	Success bool
	Err     error
}

func compileRepository(repoDirectory string, framework string) CompilationResult {
	var result CompilationResult

	switch framework {
	case "hardhat":
		result = compileHardhat(repoDirectory)
	case "foundry":
		result = compileFoundry(repoDirectory)
	case "brownie":
		result = compileFoundry(repoDirectory)
	default:
		result.Err = fmt.Errorf("unknown framework: %s", framework)
	}

	return result
}

func compileHardhat(repoDirectory string) CompilationResult {

	// cmd := exec.Command("yarn")
	// cmd.Dir = repoDirectory

	// if err := cmd.Run(); err != nil {
	//     return CompilationResult{Success: false, Err: err}
	// }

	cmd := exec.Command("npx", "hardhat", "compile", "--force")
	cmd.Dir = repoDirectory
	if err := cmd.Run(); err != nil {
		return CompilationResult{Success: false, Err: err}
	}

	return CompilationResult{Success: true}
}

func compileFoundry(repoDirectory string) CompilationResult {
	cmd := exec.Command("forge", "build")
	cmd.Dir = repoDirectory

	if err := cmd.Run(); err != nil {
		return CompilationResult{Success: false, Err: err}
	}

	return CompilationResult{Success: true}
}
