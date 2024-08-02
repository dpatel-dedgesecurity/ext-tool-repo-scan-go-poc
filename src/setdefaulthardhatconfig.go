package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// replaceConfigFile replaces the content of the existing hardhat configuration file
// in the target repository with the content from the given file.
func replaceConfigFile(jsPath, tsPath, targetRepo string) error {
	// Define the paths for the target repository's configuration files
	jsFilePath := filepath.Join(targetRepo, "hardhat.config.js")
	tsFilePath := filepath.Join(targetRepo, "hardhat.config.ts")

	// Determine which configuration file is present and should be replaced
	var configFilePath string
	if configFileExists(jsFilePath) {
		configFilePath = jsFilePath
	} else if configFileExists(tsFilePath) {
		configFilePath = tsFilePath
	} else {
		return fmt.Errorf("neither hardhat.config.js nor hardhat.config.ts found in target repository")
	}

	// Determine the source file to use based on the existing configuration file type
	var sourceFilePath string
	if configFileExists(jsPath) && filepath.Ext(configFilePath) == ".js" {
		sourceFilePath = jsPath
	} else if configFileExists(tsPath) && filepath.Ext(configFilePath) == ".ts" {
		sourceFilePath = tsPath
	} else {
		return fmt.Errorf("source file for the corresponding configuration type not found")
	}

	// Read the content from the source file
	sourceContent, err := ioutil.ReadFile(sourceFilePath)
	if err != nil {
		return fmt.Errorf("error reading source file %s: %v", sourceFilePath, err)
	}

	// Write the content to the target configuration file
	if err := ioutil.WriteFile(configFilePath, sourceContent, 0644); err != nil {
		return fmt.Errorf("error writing to target file %s: %v", configFilePath, err)
	}

	fmt.Printf("Successfully replaced %s with %s\n", configFilePath, sourceFilePath)
	return nil
}

// configFileExists checks if a file exists at the given path.
func configFileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
