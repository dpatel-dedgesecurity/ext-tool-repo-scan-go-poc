package sast

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

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
