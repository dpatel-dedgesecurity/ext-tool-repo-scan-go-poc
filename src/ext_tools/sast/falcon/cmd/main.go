package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"services/ext_tools/sast"
	"services/ext_tools/sast/falcon"
	"time"

	"github.com/joho/godotenv"
)

// Global variables
var tool string
var repoURL string
var outputDir string
var clonedRepoDir string // New variable to store the path of the cloned repository
var frameworkwithpath string
var packagewithpath string

// var outputDir string
var list_of_detected_paths []string
var compiled_paths []string

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve the value of REPO_OUTPUT_PATH, REPO_URL, and TOOL environment variables
	outputDir = os.Getenv("REPO_OUTPUT_PATH")
	repoURL = os.Getenv("REPO_URL")
	tool = os.Getenv("TOOL")

	// Create the scanned_output directory in the root path

	if outputDir == "" {
		log.Fatal("Error: REPO_OUTPUT_PATH environment variable is not set")
	}
	if tool == "" {
		log.Fatal("Error: TOOL environment variable is not set")
	}

	// // Construct outputDirPath with the repository name and "autogen_output"
	// outputDirPath = fmt.Sprintf("%s/%s_autogen_output", outputDirPath, repoName)
	// rand.Seed(time.Now().UnixNano())

	// Extract the repository name from repoURL
	repoName := sast.ExtractRepoName(repoURL)
	if repoName == "" {
		log.Fatal("Error: Could not extract repository name from REPO_URL")
	}

	// Construct outputDirPath with the repository name and "autogen_output"
	outputDir = fmt.Sprintf("%s/%s_scaned_output", outputDir, repoName)
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// outputDirpath = filepath.Join(outputDir, "scaned_output")
	// if err := os.MkdirAll(outputDirpath, 0755); err != nil {
	// 	return
	// }
	

	if err := os.RemoveAll(outputDir); err != nil {
		fmt.Printf("Error removing autogen_output directory: %v\n", err)
		return
	}

	if err := os.Mkdir(outputDir, 0755); err != nil {
		fmt.Printf("Error creating autogen_output directory: %v\n", err)
		return
	}

	
	clonedRepoDir, err := sast.CloneRepository(repoURL)
	if err != nil {
		fmt.Println("❗️ ~ funcmain ~ err ~ Error in cloning repository :", err)
		return
	}
	fmt.Println("🚀 ~ funcmain ~ clonedRepoDir:", clonedRepoDir)

	compiled_paths, err := falcon.SearchTargets(clonedRepoDir)
	if err != nil {
		fmt.Println("❗️ ~ funcmain ~ err: Error while searching target paths ", err)
	}
	// fmt.Println("🚀 ~ funcmain ~ compiled_paths:", compiled_paths)

	is_executed, err := falcon.Execute(compiled_paths,outputDir)
	if err != nil {
		fmt.Println("Failed to Execute Tool on given repository !!!")
	} else if is_executed {
		fmt.Println("Execution Completed Successfully ... ")
	}

}
