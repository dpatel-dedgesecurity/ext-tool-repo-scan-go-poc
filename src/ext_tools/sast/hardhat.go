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
	versions := make(map[string]struct{}) // Use a map to store unique version strings
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sol") {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %w", path, err)
			}
			version := extractSolidityVersion(string(content))
			if version != "" {
				parsedVersions := parseVersions(version)
				for _, v := range parsedVersions {
					versions[v] = struct{}{} // Store version in map to ensure uniqueness
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking through directory: %w", err)
	}

	// Convert map keys to a slice
	uniqueVersions := make([]string, 0, len(versions))
	for v := range versions {
		uniqueVersions = append(uniqueVersions, v)
	}

	return uniqueVersions, nil
}

// extractSolidityVersion extracts the Solidity compiler version from the content of a .sol file.
func extractSolidityVersion(content string) string {
	re := regexp.MustCompile(`pragma\s+solidity\s+((\^|~)?\d+\.\d+\.\d+|>=\d+\.\d+\.\d+( <\d+\.\d+\.\d+)?( \|\| >=\d+\.\d+\.\d+( <\d+\.\d+\.\d+)?)?);`)
	match := re.FindStringSubmatch(content)
	if len(match) > 1 {
		return match[1] // Return the first captured group, which is the version number
	}
	return ""
}

// parseVersions parses the version string into individual version strings.
func parseVersions(version string) []string {
	versions := []string{}
	re := regexp.MustCompile(`\d+\.\d+\.\d+`)
	matches := re.FindAllString(version, -1)
	for _, v := range matches {
		versions = append(versions, v)
	}
	return versions
}

// setupHardhatProject sets up a temporary Hardhat project with a basic configuration and installs necessary packages.
func SetupHardhatProject(directory string) error {
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

		// Run yarn install to install dependencies
		fmt.Println("Running 'yarn install' to install dependencies...")
		cmd := exec.Command("yarn", "install")
		cmd.Dir = directory

		// Capture and print the output
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to run 'yarn install': %s\nOutput:\n%s", err, output)
		}
		fmt.Printf("Installation logs:\n%s", output)
	}

	// // Check if hardhat.config.js exists, if not, create it
	// if _, err := os.Stat(hardhatConfigPath); os.IsNotExist(err) {
	// 	fmt.Println("Creating Hardhat config file...")
	// 	if err != nil {
	// 		return fmt.Errorf("failed to create hardhat.config.js: %w", err)
	// 	}

	// }
	err := createHardhatConfig(directory, hardhatConfigPath)
	if err != nil {
		return fmt.Errorf("🚀 ~ funcSetupHardhatProject ~ err: %w", err)
	}

	fmt.Println("Hardhat project initialized successfully.")
	return nil
}

// createHardhatConfig creates a hardhat.config.js file with the specified compiler versions.
func createHardhatConfig(dir string, filePath string) error {

	// Get unique compiler versions from Solidity files
	//filePath := dir + "/hardhat.config.js"
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

		//disable network in config function called here------------------

		return "", fmt.Errorf("hardhat compile error: %s", err)

	}

	return dir, nil
}
