package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getSolidityFilesDetails(directory string, directory_path string) error {
	solidityFilesFile := filepath.Join(directory_path, "solidity_files.txt")
	resultsFile := filepath.Join(directory_path, "results.md")

	// Find all .sol files and list them, excluding certain directories
	fmt.Printf("Listing all Solidity files in directory: %s\n", directory_path)

	excludeDirs := []string{"node_modules", "lib", "test", "script"}
	file, err := os.Create(solidityFilesFile)
	if err != nil {
		return err
	}
	defer file.Close()

	err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			for _, dir := range excludeDirs {
				if info.Name() == dir {
					return filepath.SkipDir
				}
			}
		} else if strings.HasSuffix(info.Name(), ".sol") && !strings.HasSuffix(info.Name(), ".t.sol") && !strings.HasSuffix(info.Name(), ".s.sol") {
			_, err := file.WriteString(path + "\n")
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	err = processSolidityFiles(solidityFilesFile, resultsFile)
	if err != nil {
		fmt.Printf("Error running processSolidityFiles func:")
		return err
	}

	return nil
}
