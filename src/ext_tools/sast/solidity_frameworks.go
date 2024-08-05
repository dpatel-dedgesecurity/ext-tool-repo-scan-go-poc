package sast

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CompilationResult struct {
	Success []FrameworkCompilation `json:"success"`
	Error   []CompilationError     `json:"error"`
}

type FrameworkCompilation struct {
	DirectoryPath string `json:"directory_path"`
	Framework     string `json:"framework"`
	Force         bool   `json:"force"`
}

type CompilationError struct {
	DirectoryPath string `json:"directory_path"`
	ErrorBlob     string `json:"error_blob"`
}

func DetectFramework(clonedRepoPath string) (map[string][]string, error) {
	frameworks := make(map[string][]string) // Map to store frameworks and their directory paths

	err := filepath.Walk(clonedRepoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current path is a directory
		if info.IsDir() {
			// Skip the directory if it is named 'lib' or 'node_modules'
			dirName := filepath.Base(path)
			if dirName == "lib" || dirName == "node_modules" {
				return filepath.SkipDir
			}

			// Check for framework files in this directory
			files, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			for _, file := range files {
				if file.IsDir() {
					continue
				}

				// Determine which framework files are present
				var framework string
				switch filepath.Base(file.Name()) {
				case "hardhat.config.js", "hardhat.config.ts":
					framework = "hardhat"
				case "foundry.toml":
					framework = "foundry"
				case "brownie-config.yaml":
					framework = "brownie"
				}

				// Add the path to the corresponding framework in the map
				if framework != "" {
					frameworks[framework] = append(frameworks[framework], path)
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %v: %w", clonedRepoPath, err)
	}

	// Check if any frameworks were found
	if len(frameworks) == 0 {
		fmt.Println("No frameworks found.")
		return nil, nil
	}

	// Optionally handle multiple frameworks in the same directory
	// This logic is now more focused on detecting frameworks
	for framework, dirs := range frameworks {
		if len(dirs) > 1 {
			fmt.Printf("Framework %s is found in directories: %s\n", framework, strings.Join(dirs, ", "))
		}
	}

	return frameworks, nil
}

func CompileProject(frameworks_with_path map[string][]string) (*CompilationResult, error) {
	result := &CompilationResult{
		Success: []FrameworkCompilation{},
		Error:   []CompilationError{},
	}

	for framework, dirs := range frameworks_with_path {
		for _, dir := range dirs {
			// Skip directories named "node_modules" and "lib"
			if dir == "node_modules" || dir == "lib" {
				continue
			}

			var filepath string
			var err error
			switch framework {
			case "foundry":
				filepath, err = CompileFoundryProject(dir)
			case "hardhat":
				filepath, err = CompileHardhatProject(dir)
			default:
				err = fmt.Errorf("unknown framework: %s", framework)
			}

			if err != nil {
				result.Error = append(result.Error, CompilationError{
					DirectoryPath: dir,
					ErrorBlob:     err.Error(),
				})
			} else {
				result.Success = append(result.Success, FrameworkCompilation{
					DirectoryPath: filepath,
					Framework:     framework,
					Force:         false, // Adjust as needed
				})
			}
		}
	}

	return result, nil
}
