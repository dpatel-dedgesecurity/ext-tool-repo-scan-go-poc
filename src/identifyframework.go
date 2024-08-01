package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// detectFramework checks each directory and subdirectory for specific framework configuration files
func detectFramework(baseDir string) (map[string][]string, error) {
	frameworks := make(map[string][]string) // Map to store directory paths and their frameworks

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current path is a directory
		if info.IsDir() {
			// Skip the directory if it is named 'lib'
			if filepath.Base(path) == "lib" {
				return filepath.SkipDir
			}

			// Check for framework files in this directory
			files, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			var currentFrameworks []string
			for _, file := range files {
				if file.IsDir() {
					continue
				}

				switch filepath.Base(file.Name()) {
				case "hardhat.config.js", "hardhat.config.ts":
					currentFrameworks = append(currentFrameworks, "hardhat")
				case "foundry.toml":
					currentFrameworks = append(currentFrameworks, "foundry")
				case "brownie-config.yaml":
					currentFrameworks = append(currentFrameworks, "brownie")
				}
			}

			// If we found any framework files in this directory
			if len(currentFrameworks) > 0 {
				frameworks[path] = currentFrameworks
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %v: %w", baseDir, err)
	}

	// Check if any frameworks were found
	if len(frameworks) == 0 {
		fmt.Println("No frameworks found.")
		return nil, nil
	}

	// Handle multiple frameworks in the same directory
	for dir, fwks := range frameworks {
		if len(fwks) > 1 {
			fmt.Printf("Directory %s contains multiple frameworks: %s\n", dir, strings.Join(fwks, ", "))
		}
	}

	return frameworks, nil
}
