package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Run the specified tool (either Falcon or Slither)
func runTool(tool, directory, directoryPath string) error {
	switch tool {
	case "falcon":
		return runFalcon(directory, directoryPath)
	case "slither":
		return runSlither(directory, directoryPath)
	default:
		return fmt.Errorf("unknown tool: %s", tool)
	}
}

func runFalcon(directory, directoryPath string) error {
	resultsFile := filepath.Join(directoryPath, "results.md")

	startTime := time.Now()
	cmd := exec.Command("falcon", ".", "--json", filepath.Join("../", directoryPath, "output.json"), "--filter-paths", "\\.t\\.sol$,\\.s\\.sol$")
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
