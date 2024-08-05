package sast

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

// Clones the repository in /input_repos dir and returns path
func CloneRepository(repo_url string) (string, error) {
	if repo_url == "" {
		return "", fmt.Errorf("repository URL is empty")
	}

	// Extract repository name from URL
	repoParts := strings.Split(repo_url, "/")
	repoName := strings.TrimSuffix(repoParts[len(repoParts)-1], ".git")

	// Generate a unique directory path for cloning
	cloneDir := filepath.Join("./../input_repos", repoName)

	// Check if the directory already exists
	if _, err := os.Stat(cloneDir); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", repo_url, cloneDir)

		// Capture standard output and standard error
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("failed to clone repository: %v", err)
		}
	} else {
		return "", fmt.Errorf("directory '%s' already exists", cloneDir)
	}

	return cloneDir, nil
}

// below funtion traverse through cloned repo directory, fetches package.json and returns list of path where package.json is available
func DetectPackageJson(clonedRepoPath string) ([]string, error) {
	var paths []string

	err := filepath.Walk(clonedRepoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip "node_modules" and "lib" directories
		if info.IsDir() {
			dirName := filepath.Base(path)
			if dirName == "node_modules" || dirName == "lib" {
				return filepath.SkipDir // Skip the directory and its contents
			}
			return nil // Continue walking
		}

		// Check for "package.json" file
		if info.Name() == "package.json" {
			parentDir := filepath.Dir(path)
			paths = append(paths, parentDir)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking through directory tree: %w", err)
	}

	return paths, nil
}

// IdentifyPackageManager scans the given list of paths and returns a map
// of package managers and the paths containing their respective lock files.
// Also includes a "mockpackage" key for paths where no known package manager is found.
func IdentifyPackageManager(paths []string) map[string][]string {

	// Initialize the map to store package managers and their paths
	packageManagers := make(map[string][]string)

	// Iterate over each path in the input list
	for _, dirPath := range paths {
		// Check if the directory exists
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			fmt.Printf("Path does not exist: %s\n", dirPath)
			continue
		}

		// Check for presence of package manager files
		var manager string
		if fileExists(filepath.Join(dirPath, "yarn.lock")) {
			manager = "yarn"
		} else if fileExists(filepath.Join(dirPath, "package-lock.json")) {
			manager = "npm"
		}

		// If a manager was identified, add the path to the corresponding manager's list
		if manager != "" {
			packageManagers[manager] = append(packageManagers[manager], dirPath)
		} else {
			// If no package manager was identified, add to "mockpackage"
			packageManagers["nopackage"] = append(packageManagers["nopackage"], dirPath)
		}
	}

	return packageManagers
}

// fileExists checks if a file exists at the given path
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// // Pending to Test foundry.toml + .gitmodules while installing dependencies when they are in separate paths

// RunPackageManager installs dependencies for each directory based on the package manager specified.
func RunPackageManager(mapFetched map[string][]string) error {
	for pkgManager, dirs := range mapFetched {
		switch pkgManager {
		case "yarn":
			for _, dir := range dirs {
				cmd := exec.Command("yarn")
				cmd.Dir = dir
				output, err := cmd.CombinedOutput()
				fmt.Println(string(output)) // Print output to console
				if err != nil {
					return fmt.Errorf("failed to install dependencies with yarn in directory %s: %w", dir, err)
				}
			}
		case "npm":
			for _, dir := range dirs {
				cmd := exec.Command("npm", "install")
				cmd.Dir = dir
				output, err := cmd.CombinedOutput()
				fmt.Println(string(output)) // Print output to console
				if err != nil {
					return fmt.Errorf("failed to install dependencies with npm in directory %s: %w", dir, err)
				}
			}
		case "nopackage":
			// Assuming defaulting to yarn if no package manager is specified
			for _, dir := range dirs {
				cmd := exec.Command("yarn")
				cmd.Dir = dir
				output, err := cmd.CombinedOutput()
				fmt.Println(string(output)) // Print output to console
				if err != nil {
					return fmt.Errorf("failed to install dependencies with yarn in directory %s: %w", dir, err)
				}
			}
		default:
			return fmt.Errorf("unknown package manager: %s", pkgManager)
		}
	}
	return nil
}

// // RunPackageManager installs dependencies for each directory based on the package manager specified.
// func RunPackageManager(mapFetched map[string][]string) error {
// 	nodeVersionRegex := regexp.MustCompile(`Expected version "^12.0.0 || ^14.0.0 || ^16.0.0".`)

// 	fmt.Println("🚀 ~ funcRunPackageManager ~ nodeVersionRegex:", nodeVersionRegex)
// 	for pkgManager, dirs := range mapFetched {
// 		switch pkgManager {
// 		case "yarn":
// 			for _, dir := range dirs {
// 				cmd := exec.Command("yarn", "install")
// 				cmd.Dir = dir
// 				output, err := cmd.CombinedOutput() // Capture stdout and stderr
// 				// fmt.Println("🚀 ~ funcRunPackageManager ~ output:", string(output))
// 				if err != nil {
// 					errorMsg := string(output)
// 					fmt.Println("🚀 ~ funcRunPackageManager ~ errorMsg:", errorMsg)
// 					if nodeVersionRegex.MatchString(errorMsg) {
// 						// Extract the expected Node.js version from the error mesFailed to install dependencies with yarn in directory ../input_repos/golom-protocol-contracts:sage
// 						fmt.Println("🚀 ~ ifnodeVersionRegex.MatchString ")

// 						expectedVersion := "16.0.0"
// 						fmt.Println("🚀 ~ ifnodeVersionRegex.MatchString ~ expectedVersion:", expectedVersion)
// 						cmd := exec.Command("yarn", "install")
// 						cmd.Dir = dir
// 						output, err := cmd.CombinedOutput()
// 						if err != nil {
// 							fmt.Println("🚀 ~ ifnodeVersionRegex.MatchString ~ err:", err)
// 						}
// 						fmt.Println("🚀 ~ ifnodeVersionRegex.MatchString ~ output: yarn installed ", string(output))
// 						updateNodeVersion(expectedVersion)
// 						cmd1 := exec.Command("yarn", "install")
// 						cmd1.Dir = dir
// 						output, err = cmd1.CombinedOutput()
// 						if err != nil {
// 							fmt.Println("🚀 ~ ifnodeVersionRegex.MatchString ~ err:", err)
// 						}
// 						fmt.Println("🚀 ~ ifnodeVersionRegex.MatchString ~ output: yarn installed ", string(output))
// 					} else {
// 						// fmt.Printf("Failed to install dependencies with yarn in directory %s:\n%s\nError: %v\n", dir, errorMsg, err)
// 						return fmt.Errorf("failed to install dependencies with yarn in directory %s: %w", dir, err)
// 					}
// 				} else {
// 					fmt.Printf("Successfully installed dependencies with yarn in directory %s:\n%s\n", dir, string(output))
// 				}
// 			}
// 		case "npm":
// 			for _, dir := range dirs {
// 				cmd := exec.Command("npm", "install")
// 				cmd.Dir = dir
// 				output, err := cmd.CombinedOutput() // Capture stdout and stderr
// 				if err != nil {
// 					fmt.Printf("Failed to install dependencies with npm in directory %s:\n%s\nError: %v\n", dir, string(output), err)
// 					return fmt.Errorf("failed to install dependencies with npm in directory %s: %w", dir, err)
// 				} else {
// 					fmt.Printf("Successfully installed dependencies with npm in directory %s:\n%s\n", dir, string(output))
// 				}
// 			}
// 		case "nopackage":
// 			// Assuming defaulting to yarn if no package manager is specified
// 			for _, dir := range dirs {
// 				cmd := exec.Command("yarn", "install")
// 				cmd.Dir = dir
// 				output, err := cmd.CombinedOutput() // Capture stdout and stderr
// 				if err != nil {
// 					fmt.Printf("Failed to install dependencies with yarn in directory %s:\n%s\nError: %v\n", dir, string(output), err)
// 					return fmt.Errorf("failed to install dependencies with yarn in directory %s: %w", dir, err)
// 				} else {
// 					fmt.Printf("Successfully installed dependencies with yarn in directory %s:\n%s\n", dir, string(output))
// 				}
// 			}
// 		default:
// 			return fmt.Errorf("unknown package manager: %s", pkgManager)
// 		}
// 	}
// 	return nil
// }

// func updateNodeVersion(version string) {
// 	fmt.Printf("Attempting to update Node.js to version %s...\n", version)

// 	// Source nvm and install the specified Node.js version
// 	cmd := exec.Command("bash", "-c", fmt.Sprintf("source ~/.nvm/nvm.sh && nvm install %s", version))
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatalf("Failed to install Node.js version %s: %v\nOutput: %s", version, err, output)
// 	}

// 	// Source nvm and switch to the newly installed version
// 	cmd = exec.Command("bash", "-c", fmt.Sprintf("source ~/.nvm/nvm.sh && nvm use %s", version))
// 	output, err = cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatalf("Failed to switch to Node.js version %s: %v\nOutput: %s", version, err, output)
// 	}

// 	// Source nvm and set the newly installed version as the default globally
// 	cmd = exec.Command("bash", "-c", fmt.Sprintf("source ~/.nvm/nvm.sh && nvm alias default %s", version))
// 	output, err = cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatalf("Failed to set Node.js version %s as default: %v\nOutput: %s", version, err, output)
// 	}

// 	fmt.Printf("Successfully updated Node.js to version %s and set it as default.\n", version)
// }

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

// CompileFoundryProject compiles a Foundry project located in the specified directory
func CompileFoundryProject(dir string) (string, error) {
	// Example command; adjust based on your needs
	cmd := exec.Command("forge", "build")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output)) // Print output to console
	if err != nil {
		return "", fmt.Errorf("foundry compile error: %s", err)
	}

	return dir, nil
}

// CompileHardhatProject compiles a Hardhat project located in the specified directory
func CompileHardhatProject(dir string) (string, error) {
	// Example command; adjust based on your needs
	cmd := exec.Command("npx", "hardhat", "compile")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output)) // Print output to console
	if err != nil {
		return "", fmt.Errorf("hardhat compile error: %s", err)
	}

	return dir, nil
}

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

// func CompileProject() {

// }

// func CompileFoundryProject(target_path string) (bool, error) {

// }

// func CompileHardhatProject(target_path string) (bool, error) {

// }

// func RunTool() {

// }
