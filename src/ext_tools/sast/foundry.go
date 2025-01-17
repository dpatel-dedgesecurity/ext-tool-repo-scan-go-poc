package sast

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Define structs according to your TOML structure
type FoundryConfig struct {
	Project struct {
		Src       string   `toml:"src"`
		Out       string   `toml:"out"`
		Libs      []string `toml:"libs"`
		Test      string   `toml:"test"`
		CachePath string   `toml:"cache_path"`
	} `toml:"project"`

	Doc struct {
		Homepage string   `toml:"homepage"`
		Ignore   []string `toml:"ignore"`
	} `toml:"doc"`
}

func ConfigureFoundryTOML(filePath string) error {
	// var filePath string

	filePath = filePath + "/foundry.toml"
	fmt.Println("Path for foundry.toml : ", filePath)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("failed to read file: %w", err)
	}

	// Convert byte slice to string for manipulation
	contentStr := string(content)

	// Step 2: Check and Update the 'out' key
	lines := strings.Split(contentStr, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "out =") {
			lines[i] = "out = \"out\""
			break // Stop searching once we've found and replaced the line
		}
	}

	// Join lines back into a single string
	newContentStr := strings.Join(lines, "\n")

	// Step 3: Write the updated content back to the file
	err = ioutil.WriteFile(filePath, []byte(newContentStr), 0644)
	if err != nil {
		fmt.Println("failed to write file: %w", err)
	}

	return nil
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
