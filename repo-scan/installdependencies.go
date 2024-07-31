package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// installDependencies installs dependencies based on the detected environment
func installDependencies(directory, env string) error {
	switch env {
	case "hardhat":
		fmt.Println("Handling Hardhat environment")
		// Check for package-lock.json and yarn.lock to determine the correct package manager
		if _, err := os.Stat(filepath.Join(directory, "package-lock.json")); err == nil {
			fmt.Println("Found package-lock.json, using npm")
			if err := runCommand(directory, "npm", "install"); err != nil {
				fmt.Printf("npm install failed: %v\n", err)
				return err
			}
		} else if _, err := os.Stat(filepath.Join(directory, "yarn.lock")); err == nil {
			fmt.Println("Found yarn.lock, using yarn")
			if err := runCommand(directory, "yarn", "install"); err != nil {
				fmt.Printf("yarn install failed: %v\n", err)
				return err
			}
		} else if _, err := os.Stat(filepath.Join(directory, "package.json")); err == nil {
			fmt.Println("Found package.json but no lock files, deciding on package manager")
			// If no lock file is found, decide based on project preference or default to npm
			// Defaulting to npm if no lock files are found
			if err := runCommand(directory, "npm", "install"); err != nil {
				fmt.Printf("npm install failed: %v\n", err)
				return err
			}
		} else {
			fmt.Println("No package.json, package-lock.json, or yarn.lock found, no dependencies installed")
		}

	case "foundry":
		fmt.Println("Handling Foundry environment")
		if _, err := os.Stat(filepath.Join(directory, "foundry.toml")); err == nil {
			// Check for package-lock.json and yarn.lock in the Foundry project directory
			if _, err := os.Stat(filepath.Join(directory, "package-lock.json")); err == nil {
				fmt.Println("Found package-lock.json, using npm")
				if err := runCommand(directory, "npm", "install"); err != nil {
					fmt.Printf("npm install failed: %v\n", err)
					return err
				}
			} else if _, err := os.Stat(filepath.Join(directory, "yarn.lock")); err == nil {
				fmt.Println("Found yarn.lock, using yarn")
				if err := runCommand(directory, "yarn", "install"); err != nil {
					fmt.Printf("yarn install failed: %v\n", err)
					return err
				}
			} else if _, err := os.Stat(filepath.Join(directory, "package.json")); err == nil {
				fmt.Println("Found package.json but no lock files, deciding on package manager")
				// If no lock file is found, decide based on project preference or default to npm
				// Defaulting to npm if no lock files are found
				if err := runCommand(directory, "npm", "install"); err != nil {
					fmt.Printf("npm install failed: %v\n", err)
					return err
				}
			} else {
				fmt.Println("No package.json, package-lock.json, yarn.lock, or foundry.toml found, proceeding with forge install")
				if err := runCommand(directory, "forge", "install"); err != nil {
					fmt.Printf("forge install failed: %v\n", err)
					// Attempt alternative Foundry setup
					fmt.Println("Attempting alternative Foundry setup")
					if err := runCommand(directory, "forge", "update"); err != nil {
						fmt.Printf("forge update also failed: %v\n", err)
						return err
					}
				}
			}
		} else {
			fmt.Println("foundry.toml not found, skipping Foundry dependency installation")
		}

	case "brownie":
		fmt.Println("Handling Brownie environment")
		if _, err := os.Stat(filepath.Join(directory, "requirements.txt")); err == nil {
			if err := runCommand(directory, "pip", "install", "-r", filepath.Join(directory, "requirements.txt")); err != nil {
				fmt.Printf("pip install failed: %v\n", err)
				return err
			}
		} else {
			fmt.Println("requirements.txt not found, skipping Brownie dependency installation")
		}

	default:
		fmt.Println("Unknown environment:", env)
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
