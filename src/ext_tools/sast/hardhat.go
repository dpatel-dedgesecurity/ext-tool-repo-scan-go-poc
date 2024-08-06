package sast

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// setupHardhatProject sets up a temporary Hardhat project with a basic configuration and installs necessary packages.
// scanFiles reads Solidity files and returns unique compiler versions found in them.
func scanFiles(dir string) ([]string, error) {
	var versions []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sol") {
			content, err := ioutil.ReadFile(path)
			fmt.Println("🚀 ~ if!info.IsDir ~ content:", string(content))
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", path, err)
			}
			version := extractSolidityVersion(string(content))
			fmt.Println("🚀 ~ if!info.IsDir ~ version:", version)
			if version != "" && !contains(versions, version) {
				versions = append(versions, version)
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking through directory: %w", err)
	}

	return versions, nil
}

// extractSolidityVersion extracts the Solidity compiler version from the content of a .sol file.
func extractSolidityVersion(content string) string {
	re := regexp.MustCompile(`pragma solidity ([0-9]+\.[0-9]+(\.[0-9]+)?)(?:\^\d+\.\d+(\.\d+)?)?`)
	match := re.FindStringSubmatch(content)
	if len(match) > 1 {
		return match[1] // Return the first captured group, which is the version number
	}
	return ""
}

// contains checks if a slice contains a specific string.
func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}

// createHardhatConfig creates a hardhat.config.js file with the specified compiler versions.
func createHardhatConfig(dir, filePath string) error {
	// Get unique compiler versions from Solidity files
	versions, err := scanFiles(dir)
	fmt.Println("🚀 ~ funccreateHardhatConfig ~ dir:", dir)
	fmt.Println("🚀 ~ funccreateHardhatConfig ~ versions:", versions)
	if err != nil {
		return fmt.Errorf("error scanning files: %w", err)
	}

	// Generate compiler versions array for hardhat.config.js
	var compilers []string
	for _, version := range versions {
		compilers = append(compilers, fmt.Sprintf(`{ version: "%s" }`, version))
	}
	fmt.Println("🚀 ~ funccreateHardhatConfig ~ compilers:", compilers)

	// Define Hardhat configuration content
	hardhatConfigContent := fmt.Sprintf(`require("@nomiclabs/hardhat-ethers");

module.exports = {
  solidity: {
    compilers: [
      %s
    ],
  },
  paths: {
    // sources: "./**/*.sol",
    // sources: ".", // checks node_modules
    sources: "./contracts",
  },

};`, strings.Join(compilers, ",\n      "))

	// Write the content to the file
	err = ioutil.WriteFile(filePath, []byte(hardhatConfigContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file %s: %w", filePath, err)
	}

	fmt.Printf("Successfully created %s with specified compiler versions.\n", filePath)
	return nil
}

func MoveSolidityFiles(cloned_repo_path string) error {
	contractsDir := filepath.Join(cloned_repo_path, "contracts")

	// Check if the 'contracts' directory exists
	if _, err := os.Stat(contractsDir); !os.IsNotExist(err) {
		// 'contracts' directory already exists, skip the move process
		fmt.Printf("The 'contracts' directory already exists. Skipping the move process.\n")
		return nil
	}

	// Create the 'contracts' directory as it does not exist
	if err := os.MkdirAll(contractsDir, 0755); err != nil {
		return fmt.Errorf("failed to create 'contracts' directory: %w", err)
	}

	// Walk through the root directory to find all Solidity files
	err := filepath.WalkDir(cloned_repo_path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %w", path, err)
		}

		// Skip directoriesfilePath
		if d.IsDir() {
			return nil
		}

		// Check for Solidity files
		if filepath.Ext(path) == ".sol" {
			relativePath, err := filepath.Rel(cloned_repo_path, path)
			if err != nil {
				return fmt.Errorf("failed to get relative path for %q: %w", path, err)
			}

			// Construct the new path inside the 'contracts' directory
			newPath := filepath.Join(contractsDir, relativePath)

			// Create directories if needed
			if err := os.MkdirAll(filepath.Dir(newPath), 0755); err != nil {
				return fmt.Errorf("failed to create directory for %q: %w", newPath, err)
			}

			// Move the file to the new path
			if err := os.Rename(path, newPath); err != nil {
				return fmt.Errorf("failed to move file %q to %q: %w", path, newPath, err)
			}

			fmt.Printf("Moved %q to %q\n", path, newPath)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking the path %q: %w", cloned_repo_path, err)
	}

	fmt.Println("All Solidity files have been moved successfully.")
	return nil
}

// CompileHardhatProject compiles a Hardhat project located in the specified directory
func CompileHardhatProject(dir string) (string, error) {
	// Example command; adjust based on your needs
	cmd := exec.Command("npx", "hardhat", "compile")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output)) // Print output to console
	if err != nil {
		return "", fmt.Errorf("hardhat compile error: %s", err)
	}

	return dir, nil
}
