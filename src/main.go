package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

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
	rand.Seed(time.Now().UnixNano())

}

func main() {
	// Print the URL for debugging
	fmt.Printf("Cloning repository from URL: %s\n", repoURL)

	//Call the CloneRepository function
	clonedRepoDir, err := CloneRepository(repoURL)
	//clonedRepoDir = "/home/im/dedge/ext-tool-repo-scan-go/input_repos/2024-04-teller-finance"
	fmt.Printf("Repository successfully cloned into: %s\n", clonedRepoDir)

	// // search for .env.example in dir and rename it with .env if available
	// renameEnvExample(clonedRepoDir)
	// if err != nil {
	// 	fmt.Printf("Error cloning repository: %v\n", err)
	// 	return
	// }

	// // Example usage
	// jsPath := "./defaultHardhat.config.js"
	// tsPath := "./defaultHardhat.config.ts"

	// if err := replaceConfigFile(jsPath, tsPath, clonedRepoDir); err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// }

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
		err = generateEnvFromFile(dir, dir+"/.env")
		if err != nil {
			fmt.Println(err)
		}
	}

	// if there is no framework found
	if len(frameworkwithpath) == 0 {
		handleNoFrameworkRepo(directory)
	}

	//Dependecy Installtion
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

	err = configureFoundryTOML(frameworkwithpath)
	if err != nil {
		fmt.Println("🚀 ~ funcmain ~ err:", err)
	}

	// Compile the frameworks even if there are multiple framework configs it will compile for all of them
	results := compileRepositories(frameworkwithpath)

	// This loop is just to print logs on console
	for directory, result := range results {
		if result.Success {
			fmt.Printf("Compilation succeeded in directory %s\n", directory)
		} else {
			fmt.Printf("Compilation failed in directory %s: %v\n", directory, result.Err)
		}
	}

	// Run the specified tool
	startTime := time.Now()
	if err := runTool(tool, frameworkwithpath, outputDirPath, results); err != nil {
		fmt.Printf("Error running %s: %v\n", tool, err)
		return
	}
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time: %s\n", executionTime)

	// Get details about Solidity files in the project to ammend in output
	if err := getSolidityFilesDetails(directory, outputDirPath); err != nil {
		fmt.Printf("Error getting Solidity files details: %v\n", err)
		return
	}

	fmt.Println("Process completed successfully.")
}
