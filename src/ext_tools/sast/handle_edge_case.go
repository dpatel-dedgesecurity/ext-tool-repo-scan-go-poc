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
