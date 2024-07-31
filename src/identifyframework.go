package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func detectEnvironment(directory string) string {
	var DIRECTORY = directory
	var environment string

	// Check for existence of hardhat config files
	_, err := os.Stat(filepath.Join(DIRECTORY, "hardhat.config.js"))
	_, err1 := os.Stat(filepath.Join(DIRECTORY, "hardhat.config.ts"))
	if err == nil || err1 == nil {
		environment = "hardhat"
	} else {
		// Only check for foundry.toml if hardhat config JS doesn't exist
		_, err = os.Stat(filepath.Join(DIRECTORY, "foundry.toml"))
		if err == nil {
			environment = "foundry"
		} else {
			// Check for existence of brownie-config.yaml
			_, err = os.Stat(filepath.Join(DIRECTORY, "brownie-config.yaml"))
			if err == nil {
				environment = "brownie"
			} else {

			}
		}
	}

	err = filepath.Walk(DIRECTORY, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current path is a directory
		if info.IsDir() {
			return nil
		}

		// Check for hardhat config files
		if filepath.Base(path) == "hardhat.config.js" || filepath.Base(path) == "hardhat.config.ts" {
			environment = "hardhat"
			return filepath.SkipDir // Stop further walking in subdirectories
		}

		// Check for foundry.toml
		if filepath.Base(path) == "foundry.toml" {
			foundryTomlPath := filepath.Join(directory, "foundry.toml")
			if _, err := os.Stat(foundryTomlPath); os.IsNotExist(err) {
				return fmt.Errorf("foundry.toml not found in %s", directory)
			}
			environment = "foundry"
			file, err := os.ReadFile(foundryTomlPath)
			if err != nil {
				return fmt.Errorf("failed to read foundry.toml: %w", err)
			}

			// Update foundry.toml to include `out = "out"`
			content := string(file)
			if !containsOutConfiguration(content) {
				content += "\nout = \"out\"\n"
				err = os.WriteFile(foundryTomlPath, []byte(content), 0644)
				if err != nil {
					return fmt.Errorf("failed to write foundry.toml: %w", err)
				}
			}
			return filepath.SkipDir // Stop further walking in subdirectories
		}

		// Check for brownie-config.yaml
		if filepath.Base(path) == "brownie-config.yaml" {
			environment = "brownie"
			return filepath.SkipDir // Stop further walking in subdirectories
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", directory, err)
		return "unknown"
	}

	if environment == "" {
		environment = "unknown"
	}

	fmt.Println("Detected Environment:", environment)

	// check if environment is unknown
	if environment == "unknown" {
		// Create a temporary Hardhat project
		if err := moveSolidityFiles(directory); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		if err := setupHardhatProject(directory); err != nil {
			fmt.Printf("Error setting up Hardhat project: %v\n", err)
			return "unknown"
		}
		environment = "hardhat"
	}

	return environment

}
func containsOutConfiguration(content string) bool {
	return strings.Contains(content, "out =")
}
