package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func runTool(tool string, directory, directoryPath, framework string, compilationResults map[string]CompilationResult) error {
	switch tool {
	case "falcon":
		if err := runFalcon(directory, directoryPath, framework); err != nil {
			return fmt.Errorf("error running Falcon in directory %s for framework %s: %w", directory, framework, err)
		}
	case "slither":
		if err := runSlither(directory, directoryPath); err != nil {
			return fmt.Errorf("error running Slither in directory %s: %w", directory, err)
		}
	default:
		return fmt.Errorf("unknown tool: %s", tool)
	}
	return nil
}

func runFalcon(directory, directoryPath, framework string) error {
	resultsFile := filepath.Join(directoryPath, "results.md")

	startTime := time.Now()
	cmd := exec.Command("falcon", ".", "--json", filepath.Join("../", directoryPath, "output.json"), "--filter-paths", "\\.t\\.sol$,\\.s\\.sol$")

	// Add additional Falcon parameters based on the chosen framework
	switch framework {
	case "hardhat":
		cmd.Args = append(cmd.Args, "--compile-force-framework", "hardhat")
	case "foundry":
		cmd.Args = append(cmd.Args, "--compile-force-framework", "foundry")
	default:
		return fmt.Errorf("unknown framework: %s", framework)
	}

	cmd.Dir = directory

	outfile, err := os.Create(resultsFile)
	if err != nil {
		return err
	}
	defer outfile.Close()

	writer := io.MultiWriter(outfile, os.Stdout)
	cmd.Stdout = writer
	cmd.Stderr = writer

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running Falcon: %v\n", err)
		return err
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time: %s\n", executionTime)

	executionTimeStr := fmt.Sprintf("Execution time for Falcon: %s\n", executionTime)
	if err := appendToFile(resultsFile, executionTimeStr); err != nil {
		fmt.Printf("Error writing execution time to results file: %v\n", err)
		return err
	}

	return nil
}


// Run Slither tool
func runSlither(directory, directoryPath string) error {
	resultsFile := filepath.Join(directoryPath, "results.md")
	slitherPath := "/usr/local/bin/slither"

	startTime := time.Now()
	cmd := exec.Command(slitherPath, ".", "--json", "autogen_output/output.json", "--filter-paths", ".t\\.sol$", ".s\\.sol$")
	cmd.Dir = directory

	outfile, err := os.Create(resultsFile)
	if err != nil {
		return err
	}
	defer outfile.Close()

	writer := io.MultiWriter(outfile, os.Stdout)
	cmd.Stdout = writer
	cmd.Stderr = writer

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running Slither: %v\n", err)
		return err
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time: %s\n", executionTime)

	executionTimeStr := fmt.Sprintf("Execution time for Slither: %s\n", executionTime)
	if err := appendToFile(resultsFile, executionTimeStr); err != nil {
		fmt.Printf("Error writing execution time to results file: %v\n", err)
		return err
	}

	return nil
}
