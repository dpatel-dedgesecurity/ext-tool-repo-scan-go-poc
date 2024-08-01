package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	//	"time"

	"github.com/joho/godotenv"
)

// Global variables
var tool string
var repoURL string
var outputDirPath string
var clonedRepoDir string // New variable to store the path of the cloned repository
var frameworkwithpath string
var packagewithpath string

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

	//Call the CloneRepository function
	clonedRepoDir, err := CloneRepository(repoURL)
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		return
	}
	//clonedRepoDir = "/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance"
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
	frameworkwithpath, err := detectFramework(directory)
	if err != nil {
		fmt.Printf("Error detecting frameworks: %v\n", err)
		return
	}
	for dir, fwks := range frameworkwithpath {
		fmt.Printf("Directory: %s, Frameworks: %s\n", dir, strings.Join(fwks, ", "))
	}

	packagewithpath, err := detectPackageInstaller(directory)
	if err != nil {
		fmt.Printf("Error detecting frameworks: %v\n", err)
		return
	}
	// for dir, fwks := range packagewithpath {
	// 	fmt.Printf("packages path: %s, packages: %s\n", dir, strings.Join(fwks, ", "))
	// }

	if err := installDependencies(packagewithpath); err != nil {
		fmt.Printf("Error installing dependencies: %v\n", err)
		return
	}

	// Compile the framework
	results := compileRepositories(frameworkwithpath)

	// Handle the results
	for directory, result := range results {
		if result.Success {
			fmt.Printf("Compilation succeeded in directory %s\n", directory)
		} else {
			fmt.Printf("Compilation failed in directory %s: %v\n", directory, result.Err)
		}
	}
	// Run the specified tool
	// startTime := time.Now()
	// if err := runTool(tool, frameworkwithpath, outputDirPath, results); err != nil {
	// 	fmt.Printf("Error running %s: %v\n", tool, err)
	// 	return
	// }
	// endTime := time.Now()
	// executionTime := endTime.Sub(startTime)
	// fmt.Printf("Execution time: %s\n", executionTime)

	// Get details about Solidity files in the project
	if err := getSolidityFilesDetails(directory, outputDirPath); err != nil {
		fmt.Printf("Error getting Solidity files details: %v\n", err)
		return
	}

	fmt.Println("Process completed successfully.")
}
