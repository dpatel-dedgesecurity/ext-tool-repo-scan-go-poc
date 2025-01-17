package falcon

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"services/ext_tools/sast"
	"time"
)

func Execute(compilationResults []byte, outputDir string) (bool, error) {

	var result CompilationResult

	// Unmarshal the JSON data into the CompilationResult struct
	err := json.Unmarshal(compilationResults, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	err = RunTool(result, outputDir)
	if err != nil {
		log.Fatalf("Error runnig tool on project: %v", err)
	}
	return true, err

}

func RunTool(compilationResults CompilationResult, outputDir string) error {
	// Collect successful paths by framework
	successfulPathsByFramework := map[string][]string{
		"hardhat": {},
		"foundry": {},
	}
	var errorPaths []string

	// Collect successful paths and error paths
	for _, result := range compilationResults.Success {
		if result.Framework == "hardhat" || result.Framework == "foundry" {
			successfulPathsByFramework[result.Framework] = append(successfulPathsByFramework[result.Framework], result.DirectoryPath)
		}
	}

	for _, result := range compilationResults.Error {
		errorPaths = append(errorPaths, result.DirectoryPath)
		fmt.Printf("Compilation error in directory %s: %s\n", result.DirectoryPath, result.ErrorBlob)
	}

	// Run Falcon on successful paths
	for framework, paths := range successfulPathsByFramework {
		if len(paths) > 0 {
			for _, path := range paths {
				if framework == "hardhat" && len(successfulPathsByFramework["foundry"]) > 0 {
					// If both Hardhat and Foundry succeeded, prioritize Hardhat
					if err := runFalcon(path, "hardhat", outputDir); err != nil {
						fmt.Println("error running Falcon for Hardhat framework in path %s: %w", path, err)
						// fmt.Println("Re-Executing tool with foundry framework in path %s: %w", path, err)

						// if err = runFalcon(path, "foundry"); err != nil {
						// 	fmt.Println("Failed Execution with foundry framework execution . ")
						// }

					}
				} else {
					if err := runFalcon(path, framework, outputDir); err != nil {
						return fmt.Errorf("error running Falcon for framework %s in path %s: %w", framework, path, err)
					}
				}
			}
		}
	}

	// Run Slither on successful paths
	// for _, paths := range successfulPathsByFramework {
	// 	if len(paths) > 0 {
	// 		for _, path := range paths {
	// 			if err := runSlither(path); err != nil {
	// 				return fmt.Errorf("error running Slither in path %s: %w", path, err)
	// 			}
	// 		}
	// 	}
	// }

	// Print output for error paths without running Slither
	for _, e := range compilationResults.Error {
		fmt.Printf("Directory Path: %s\n", e.DirectoryPath)
		fmt.Printf("Error Blob: %s\n\n", e.ErrorBlob)
	}

	return nil
}

// runFalcon runs the Falcon tool on the specified path with the given framework
func runFalcon(pathToRunTool, framework string, outputDir string) error {

	resultsFile := filepath.Join(outputDir, "results.md")

	//configure .toml file if framework is foundary
	if framework == "foundry" {
		fmt.Println("🚀 ~ func runFalcon ~ framework:", framework)
		err := sast.ConfigureFoundryTOML(pathToRunTool) // Assuming sast.ConfigureFoundryTOML is defined elsewhere
		if err != nil {
			fmt.Println("Error in updating foundry.toml ~ func runFalcon ~ err:", err)
		} else {
			fmt.Println("🚀 ~ Updated foundry.toml out field")
		}
	}

	startTime := time.Now()
	cmd := exec.Command("falcon", ".", "--checklist", "--json", filepath.Join(outputDir, "output.json"), "--filter-paths", "\\.t\\.sol$,\\.s\\.sol$")

	switch framework {
	case "hardhat":
		cmd.Args = append(cmd.Args, "--compile-force-framework", "hardhat")
	case "foundry":
		cmd.Args = append(cmd.Args, "--compile-force-framework", "foundry")
	default:
		return fmt.Errorf("unknown framework: %s", framework)
	}
	cmd.Dir = pathToRunTool

	// Create output file and attach to command output
	outfile, err := os.Create(resultsFile)
	if err != nil {
		return fmt.Errorf("failed to create results file: %w", err)
	}
	defer outfile.Close()

	writer := io.MultiWriter(outfile, os.Stdout)
	cmd.Stdout = writer
	cmd.Stderr = writer

	// Run the command
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running Falcon: %v\n", err)
		return err
	}

	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time for Falcon: %s\n", executionTime)

	executionTimeStr := fmt.Sprintf("Execution time for Falcon: %s\n", executionTime)
	if err := appendToFile(resultsFile, executionTimeStr); err != nil {
		fmt.Printf("Error writing execution time to results file: %v\n", err)
		return err
	}

	return nil
}

func appendToFile(filePath, content string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
