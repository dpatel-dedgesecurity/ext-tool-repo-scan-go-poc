package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// setupHardhatProject sets up a temporary Hardhat project with a basic configuration and installs necessary packages.
func setupHardhatProject(directory string) error {
	fmt.Println("Setting up a temporary Hardhat project...")

	// Define paths for package.json and hardhat.config.js
	packageJSONPath := filepath.Join(directory, "package.json")
	hardhatConfigPath := filepath.Join(directory, "hardhat.config.js")

	// Check if package.json exists, if not, create it
	if _, err := os.Stat(packageJSONPath); os.IsNotExist(err) {
		fmt.Println("Creating package.json...")
		packageJSONContent := `{
			"name": "temporary-hardhat-project",
			"version": "1.0.0",
			"main": "index.js",
			"scripts": {
				"test": "echo \"Error: no test specified\" && exit 1"
			},
			"dependencies": {},
			"devDependencies": {
				"hardhat": "^2.0.0",
				"@nomiclabs/hardhat-ethers": "^2.0.0",
  				"ethers": "^5.0.0"
			}
		}`
		if err := ioutil.WriteFile(packageJSONPath, []byte(packageJSONContent), 0644); err != nil {
			return fmt.Errorf("failed to create package.json: %w", err)
		}
	}

	// Check if hardhat.config.js exists, if not, create it
	if _, err := os.Stat(hardhatConfigPath); os.IsNotExist(err) {
		fmt.Println("Creating Hardhat config file...")
		hardhatConfigContent := `require("@nomiclabs/hardhat-ethers");

module.exports = {
  solidity: {
    compilers: [
      // { version: "0.8.26" },
      // { version: "0.8.25" },
      { version: "0.8.24" },
      { version: "0.8.23" },
      { version: "0.8.22" },
      { version: "0.8.21" },
      { version: "0.8.20" },
      { version: "0.8.19" },
      { version: "0.8.18" },
      { version: "0.8.17" },
      { version: "0.8.16" },
      { version: "0.8.15" },
      { version: "0.8.14" },
      { version: "0.8.13" },
      { version: "0.8.12" },
      { version: "0.8.11" },
      { version: "0.8.10" },
      { version: "0.8.9" },
      { version: "0.8.8" },
      { version: "0.8.7" },
      { version: "0.8.6" },
      { version: "0.8.5" },
      { version: "0.8.4" },
      { version: "0.8.3" },
      { version: "0.8.2" },
      { version: "0.8.1" },
      { version: "0.8.0" },
      { version: "0.7.6" },
      { version: "0.7.5" },
      { version: "0.7.4" },
      { version: "0.7.3" },
      { version: "0.7.2" },
      { version: "0.7.1" },
      { version: "0.7.0" },
      { version: "0.6.9" },
      { version: "0.6.8" },
      { version: "0.6.7" },
      { version: "0.6.6" },
      { version: "0.6.5" },
      { version: "0.6.4" },
      { version: "0.6.3" },
      { version: "0.6.2" },
      { version: "0.6.12" },
      { version: "0.6.11" },
      { version: "0.6.10" },
      { version: "0.6.1" },
      { version: "0.6.0" },
      { version: "0.5.9" },
      { version: "0.5.8" },
      { version: "0.5.7" },
      { version: "0.5.6" },
      { version: "0.5.5" },
      { version: "0.5.4" },
      { version: "0.5.3" },
      { version: "0.5.2" },
      { version: "0.5.17" },
      { version: "0.5.16" },
      { version: "0.5.15" },
      { version: "0.5.14" },
      { version: "0.5.13" },
      { version: "0.5.12" },
      { version: "0.5.11" },
      { version: "0.5.10" },
      { version: "0.5.1" },
      { version: "0.5.0" },
      { version: "0.4.9" },
      { version: "0.4.8" },
      { version: "0.4.7" },
      { version: "0.4.6" },
      { version: "0.4.5" },
      { version: "0.4.4" },
      { version: "0.4.3" },
      { version: "0.4.26" },
      { version: "0.4.25" },
      { version: "0.4.24" },
      { version: "0.4.23" },
      { version: "0.4.22" },
      { version: "0.4.21" },
      { version: "0.4.20" },
      { version: "0.4.2" },
      { version: "0.4.19" },
      { version: "0.4.18" },
      { version: "0.4.17" },
      { version: "0.4.16" },
      { version: "0.4.15" },
      { version: "0.4.14" },
      { version: "0.4.13" },
      { version: "0.4.12" },
      { version: "0.4.11" },
      { version: "0.4.10" },
      { version: "0.4.1" },
      { version: "0.4.0" },

    ],
  },
  paths: {
    // sources: "./**/*.sol",
    // sources: ".", // checks node_modules
    sources: "./contracts",
  },

};`

		if err := ioutil.WriteFile(hardhatConfigPath, []byte(hardhatConfigContent), 0644); err != nil {
			return fmt.Errorf("failed to create hardhat.config.js: %w", err)
		}
	}

	fmt.Println("Hardhat project initialized successfully.")
	return nil
}

// moveSolidityFiles moves all Solidity files into a 'contracts' directory,
// preserving their relative paths, but skips the move process if the 'contracts' directory already exists.
func moveSolidityFiles(rootDir string) error {
	contractsDir := filepath.Join(rootDir, "contracts")

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
	err := filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %q: %w", path, err)
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Check for Solidity files
		if filepath.Ext(path) == ".sol" {
			relativePath, err := filepath.Rel(rootDir, path)
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
		return fmt.Errorf("error walking the path %q: %w", rootDir, err)
	}

	fmt.Println("All Solidity files have been moved successfully.")
	return nil
}
