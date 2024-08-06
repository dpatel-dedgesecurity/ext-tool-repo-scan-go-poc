package sast

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func CreateEnvFileFromDotEnvStar(dotEnvStarPath string) {
	envFilePath := dotEnvStarPath[:strings.LastIndex(dotEnvStarPath, "/")] + "/.env"

	// Check if .env file already exists
	if _, err := os.Stat(envFilePath); err == nil {
		fmt.Printf(".env file already exists at %s\n", envFilePath)
		return
	}

	// Read the content of .env.* file
	dotEnvStarContent, err := ioutil.ReadFile(dotEnvStarPath)
	if err != nil {
		fmt.Printf("Error reading .env.* file at %s: %s\n", dotEnvStarPath, err)
		return
	}

	// Write the content to .env file
	err = ioutil.WriteFile(envFilePath, dotEnvStarContent, 0644)
	if err != nil {
		fmt.Printf("Error writing .env file at %s: %s\n", envFilePath, err)
		return
	}

	fmt.Printf("Created .env file at %s with content from .env.* file at %s\n", envFilePath, dotEnvStarPath)
}

// Config represents the structure of the Hardhat config file.
type Config struct {
	Networks map[string]interface{} `json:"networks,omitempty"`
}

// removeNetworksObject removes the 'networks' object and its trailing comma from a Hardhat configuration file.
func RemoveNetworksObject(filePath string) error {
	// Read the content of the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	// Convert content to string
	fileContent := string(content)

	// Find the start of the 'networks' object
	startIdx := strings.Index(fileContent, "networks: {")
	if startIdx == -1 {
		fmt.Printf("'networks' object not found in %s.\n", filePath)
		return nil
	}

	// Find the end of the 'networks' object
	// Use a stack to correctly handle nested braces
	openBraces := 0
	inString := false
	inChar := false
	inComment := false
	endIdx := -1

	for i := startIdx; i < len(fileContent); i++ {
		char := fileContent[i]

		// Toggle string/char state
		if char == '"' && !inChar && !inString {
			inString = true
		} else if char == '"' && inString {
			inString = false
		} else if char == '\'' && !inString && !inChar {
			inChar = true
		} else if char == '\'' && inChar {
			inChar = false
		}

		// Handle comments
		if !inString && !inChar {
			if i+1 < len(fileContent) && fileContent[i] == '/' && fileContent[i+1] == '/' {
				inComment = true
			} else if inComment && fileContent[i] == '\n' {
				inComment = false
			}
		}

		// Skip characters in comments and strings
		if inComment || inString || inChar {
			continue
		}

		// Track braces
		if char == '{' {
			openBraces++
		} else if char == '}' {
			openBraces--
			if openBraces == 0 {
				endIdx = i
				break
			}
		}
	}

	if endIdx == -1 {
		return fmt.Errorf("failed to find closing brace for 'networks' object in %s", filePath)
	}

	// Find trailing comma after 'networks' object if any
	trailingCommaIdx := endIdx + 1
	if trailingCommaIdx < len(fileContent) && fileContent[trailingCommaIdx] == ',' {
		trailingCommaIdx++
	}

	// Remove the 'networks' object and any trailing comma
	modifiedContent := fileContent[:startIdx] + fileContent[trailingCommaIdx:]

	// Write the modified content back to the file
	err = ioutil.WriteFile(filePath, []byte(modifiedContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	fmt.Printf("Successfully removed 'networks' object from %s.\n", filePath)
	return nil
}

// writeModifiedContent writes the modified content back to the file
func writeModifiedContent(filePath, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	fmt.Printf("Successfully removed 'networks' object from %s.\n", filePath)
	return nil
}
