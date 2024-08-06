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
