package sast

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Clones the repository in /input_repos dir and returns path
func CloneRepository(repo_url string) (string, error) {
	if repo_url == "" {
		return "", fmt.Errorf("repository URL is empty")
	}

	// Extract repository name from URL
	repoParts := strings.Split(repo_url, "/")
	repoName := strings.TrimSuffix(repoParts[len(repoParts)-1], ".git")

	// Generate a unique directory path for cloning
	cloneDir := filepath.Join("./../input_repos", repoName)

	// Check if the directory already exists
	if _, err := os.Stat(cloneDir); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", repo_url, cloneDir)

		// Capture standard output and standard error
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("failed to clone repository: %v", err)
		}
	} else {
		return "", fmt.Errorf("directory '%s' already exists", cloneDir)
	}

	return cloneDir, nil
}

// fileExists checks if a file exists at the given path
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ExtractRepoName(url string) string {
	// Remove any leading or trailing whitespace
	url = strings.TrimSpace(url)

	// Strip off the ".git" suffix if present
	if strings.HasSuffix(url, ".git") {
		url = strings.TrimSuffix(url, ".git")
	}

	// Find the position of the last slash in the URL
	lastSlashIndex := strings.LastIndex(url, "/")
	if lastSlashIndex == -1 {
		// No slash found, return an empty string
		return ""
	}

	// Extract the repository name from the URL
	repoName := url[lastSlashIndex+1:]
	return repoName
}
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

		// Skip directories
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
