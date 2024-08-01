package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// installDependencies installs dependencies based on the detected environment

func detectPackageInstaller(baseDir string) (map[string]string, error) {
	installers := make(map[string]string) // Map to store directory paths and their package installers

	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current path is a directory
		if info.IsDir() {
			// Skip the directory if it contains node_modules
			nodeModulesPath := filepath.Join(path, "node_modules")
			if _, err := os.Stat(nodeModulesPath); err == nil {
				// Skip this directory if node_modules exists
				return filepath.SkipDir
			}

			// Check for package installer files in this directory
			files, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			// Flags to check if specific installer files are present
			yarnFound := false
			npmFound := false
			forgeFound := false

			for _, file := range files {
				if file.IsDir() {
					continue
				}

				switch file.Name() {
				case "yarn.lock":
					yarnFound = true
				case "package-lock.json":
					npmFound = true
				case ".gitmodules":
					// Check if foundry.toml is also present
					if _, err := os.Stat(filepath.Join(path, "foundry.toml")); err == nil {
						forgeFound = true
					}
				}
			}

			// Determine which installer to use based on the flags
			if yarnFound {
				installers[path] = "yarn"
			} else if forgeFound {
				installers[path] = "forge"
			} else if npmFound {
				installers[path] = "npm"
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking the path %v: %w", baseDir, err)
	}

	// Check if any installers were found
	if len(installers) == 0 {
		fmt.Println("No package installers found.")
		return nil, nil
	}

	return installers, nil
}

func installDependencies(installers map[string]string) error {
	for directory, installer := range installers {
		// Check if node_modules exists and skip the directory if it does
		nodeModulesPath := filepath.Join(directory, "node_modules")
		if _, err := os.Stat(nodeModulesPath); err == nil {
			fmt.Printf("Skipping %s as node_modules already exists\n", directory)
			continue
		}

		fmt.Printf("Handling environment in directory: %s\n", directory)

		switch installer {
		case "npm":
			fmt.Println("Using npm")
			if err := runCommand(directory, "npm", "install"); err != nil {
				fmt.Printf("npm install failed in %s: %v\n", directory, err)
				return err
			}

		case "yarn":
			fmt.Println("Using yarn")
			if err := runCommand(directory, "yarn", "install"); err != nil {
				fmt.Printf("yarn install failed in %s: %v\n", directory, err)
				return err
			}

		case "pnpm":
			fmt.Println("Checking for pnpm...")
			if err := checkAndInstallPnpm(); err != nil {
				fmt.Printf("Failed to install pnpm: %v\n", err)
				return err
			}
			fmt.Println("Using pnpm")
			if err := runCommand(directory, "pnpm", "install"); err != nil {
				fmt.Printf("pnpm install failed in %s: %v\n", directory, err)
				return err
			}

		case "forge":
			fmt.Println("Using forge")
			if err := runCommand(directory, "forge", "install"); err != nil {
				fmt.Printf("forge install failed in %s: %v\n", directory, err)
				fmt.Println("Attempting alternative Foundry setup")
				if err := runCommand(directory, "forge", "update"); err != nil {
					fmt.Printf("forge update also failed in %s: %v\n", directory, err)
					return err
				}
			}

		case "pip":
			fmt.Println("Using pip")
			requirementsPath := filepath.Join(directory, "requirements.txt")
			if _, err := os.Stat(requirementsPath); err == nil {
				if err := runCommand(directory, "pip", "install", "-r", requirementsPath); err != nil {
					fmt.Printf("pip install failed in %s: %v\n", directory, err)
					return err
				}
			} else {
				fmt.Println("requirements.txt not found, skipping pip installation")
			}

		default:
			fmt.Printf("Unknown package installer %s in directory %s\n", installer, directory)
		}
	}

	return nil
}

// checkAndInstallPnpm checks if pnpm is installed and installs it if not
func checkAndInstallPnpm() error {
	// Check if pnpm is installed
	_, err := exec.Command("pnpm", "--version").Output()
	if err == nil {
		// pnpm is already installed
		return nil
	}

	// pnpm is not installed, try to install it using npm
	fmt.Println("pnpm not found. Installing pnpm...")
	cmd := exec.Command("npm", "install", "-g", "pnpm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to install pnpm: %w", err)
	}

	return nil
}

// runCommand executes a command in the given directory
func runCommand(directory, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = directory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
