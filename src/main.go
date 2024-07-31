package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Global variables
var tool string
var repoURL string
var outputDirPath string
var clonedRepoDir string // New variable to store the path of the cloned repository
var environment string

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the value of REPO_OUTPUT_PATH, REPO_URL, and TOOL environment variables
	outputDirPath = os.Getenv("REPO_OUTPUT_PATH")
	repoURL = os.Getenv("REPO_URL")
	tool = os.Getenv("TOOL")

	if outputDirPath == "" {
		log.Fatal("Error: REPO_OUTPUT_PATH environment variable is not set")
	}
	if tool == "" {
		log.Fatal("Error: TOOL environment variable is not set")
	}

	// Extract the repository name from repoURL
	repoName := extractRepoName(repoURL)
	if repoName == "" {
		log.Fatal("Error: Could not extract repository name from REPO_URL")
	}

	// Construct outputDirPath with the repository name and "autogen_output"
	outputDirPath = fmt.Sprintf("%s/%s_autogen_output", outputDirPath, repoName)
}

func main() {
	// Print the URL for debugging
	fmt.Printf("Cloning repository from URL: %s\n", repoURL)

	// Call the CloneRepository function
	clonedRepoDir, err := CloneRepository(repoURL)
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		return
	}
	//clonedRepoDir = "/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-05-munchables"
	fmt.Printf("Repository successfully cloned into: %s\n", clonedRepoDir)

	// Remove and create autogen_output directory
	if err := os.RemoveAll(outputDirPath); err != nil {
		fmt.Printf("Error removing autogen_output directory: %v\n", err)
		return
	}

	if err := os.Mkdir(outputDirPath, 0755); err != nil {
		fmt.Printf("Error creating autogen_output directory: %v\n", err)
		return
	}

	// Use the path of the cloned repository instead of REPO_INPUT_PATH
	directory := clonedRepoDir
	fmt.Println(directory, "using cloned repo directory")

	// Install dependencies based on environment
	environment = detectEnvironment(directory)

	fmt.Printf("Detected environment: %s\n", environment)
	if err := installDependencies(directory, environment); err != nil {
		fmt.Printf("Error installing dependencies: %v\n", err)
		return
	}

	// Compile the framework
	result := compileRepository(directory, environment)
	if result.Success {
		fmt.Println("Compilation succeeded")
	} else {
		fmt.Printf("Compilation failed: %v\n", result.Err)
		// handleGitError(result.err, directory, env) // Uncomment if needed
	}

	// Run the specified tool
	startTime := time.Now()
	if err := runTool(tool, directory, outputDirPath); err != nil {
		fmt.Printf("Error running %s: %v\n", tool, err)
		return
	}
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time: %s\n", executionTime)

	// Get details about Solidity files in the project
	if err := getSolidityFilesDetails(directory, outputDirPath); err != nil {
		fmt.Printf("Error getting Solidity files details: %v\n", err)
		return
	}

	fmt.Println("Process completed successfully.")
}
