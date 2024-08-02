package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
)

// renameEnvExample takes a directory path as input, looks for a .env.example file
// within that directory, and renames it to .env if found.
func renameEnvExample(dirPath string) error {
	// Construct the full path to the .env.example file
	envExamplePath := filepath.Join(dirPath, ".env.example")

	// Check if the .env.example file exists
	_, err := os.Stat(envExamplePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(".env.example does not exist in the directory.")
			return nil // Return nil error if the file doesn't exist
		}
		return fmt.Errorf("error accessing .env.example: %w", err)
	}

	// Construct the destination path (.env)
	envPath := filepath.Join(dirPath, ".env")

	// Rename .env.example to .env
	err = os.Rename(envExamplePath, envPath)
	if err != nil {
		return fmt.Errorf("error renaming .env.example to .env: %w", err)
	}

	fmt.Println(".env.example renamed to .env")
	return nil
}

// generateRandomString generates a random alphanumeric string of fixed length
func generateRandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

// generateEnvFromFile finds the Hardhat config file in the given path, parses it for .env variables,
// initializes them with random values, and writes them to a new .env file at the specified output path.
func generateEnvFromFile(searchPath, outputPath string) error {
	configFilePath, err := findHardhatConfig(searchPath)
	if err != nil {
		return fmt.Errorf("error finding hardhat config file: %w", err)
	}

	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("error reading hardhat config file: %w", err)
	}

	envVars := parseEnvVariables(string(content))
	if len(envVars) == 0 {
		fmt.Println("No environment variables found.")
		return nil // Not an error condition, just informational
	}

	err = createEnvFile(envVars, outputPath)
	// console.log("🚀 ~ funcgenerateEnvFromFile ~ err:", err)
	if err != nil {
		return fmt.Errorf("error creating .env file: %w", err)
	}

	fmt.Printf("🚀    Created new .env file '%s' with %d variables.\n", outputPath, len(envVars))
	return nil
}

// findHardhatConfig searches for hardhat.config.ts or hardhat.config.js in the given path
func findHardhatConfig(path string) (string, error) {
	for _, filename := range []string{"hardhat.config.ts", "hardhat.config.js"} {
		filePath := filepath.Join(path, filename)
		if _, err := os.Stat(filePath); err == nil {
			return filePath, nil
		}
	}
	return "", fmt.Errorf("hardhat config file not found in %s", path)
}

// parseEnvVariables parses the given file content to find environment variable references
func parseEnvVariables(content string) []string {
	re := regexp.MustCompile(`process\.env\.([A-Za-z0-9_]+)`)
	matches := re.FindAllStringSubmatch(content, -1)
	var envVars []string
	for _, match := range matches {
		envVars = append(envVars, match[1])
	}
	return envVars
}

// createEnvFile creates a new .env file with variables initialized with random values
func createEnvFile(variables []string, outputPath string) error {
	content := ""
	for _, varName := range variables {
		content += fmt.Sprintf("%s=%s\n", varName, generateRandomString(10))
	}
	// Ensure the directory exists
	err := os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err != nil {
		return fmt.Errorf("error creating directories for .env file: %w", err)
	}
	return ioutil.WriteFile(outputPath, []byte(content), 0644)
	// return ioutil.WriteFile(outputPath, []byte(content), 0644)
}
