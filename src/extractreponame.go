package main

import (
	"strings"
)

// extractRepoName extracts the repository name from the given URL.
// It handles both HTTP(S) and SSH URL formats.
func extractRepoName(url string) string {
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
