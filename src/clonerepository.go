package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CloneRepository(url string) (string, error) {
	if url == "" {
		return "", fmt.Errorf("repository URL is empty")
	}

	// Extract repository name from URL
	repoParts := strings.Split(url, "/")
	repoName := strings.TrimSuffix(repoParts[len(repoParts)-1], ".git")

	// Generate a unique directory path for cloning
	cloneDir := filepath.Join("./../input_repos", repoName)

	// Check if the directory already exists
	if _, err := os.Stat(cloneDir); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", url, cloneDir)

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
