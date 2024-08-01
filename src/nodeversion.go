package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// handleNodeVersionError checks if the error message indicates a Node.js version issue,
// updates Node.js using nvm, and re-runs yarn install.
func handleNodeVersionError(err error, directory string) error {
	if err == nil {
		return nil
	}

	errMsg := err.Error()
	const versionErrorSubstring = `The engine "node" is incompatible with this module. Expected version "^14.0.0 || ^16.0.0 || ^18.0.0". Got`

	if strings.Contains(errMsg, versionErrorSubstring) {
		fmt.Println("Node.js version is incompatible. Updating Node.js...")

		// Define the correct Node.js version
		nodeVersion := "18" // Update to a version that matches the expected versions

		// Install the correct Node.js version using nvm
		cmd := exec.Command("nvm", "install", nodeVersion)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to install Node.js version %s: %w", nodeVersion, err)
		}

		// Use the correct Node.js version
		cmd = exec.Command("nvm", "use", nodeVersion)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to switch to Node.js version %s: %w", nodeVersion, err)
		}

		fmt.Printf("Node.js version switched to %s.\n", nodeVersion)

		// Re-run the installDependencies function
		// if err := installDependencies(packagewithpath); err != nil {
		// 	return fmt.Errorf("failed to install dependencies after updating Node.js: %w", err)
		// }
	} else {
		// If the error is not related to Node.js version, return the original error
		return fmt.Errorf("unexpected error: %w", err)
	}

	return nil
}
