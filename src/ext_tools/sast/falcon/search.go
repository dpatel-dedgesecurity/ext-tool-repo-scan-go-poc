package falcon

import (
	"encoding/json"
	"fmt"
	"log"
	"services/ext_tools/sast"
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

func SearchTargets(cloned_repo_path string) ([]byte, error) {
	list_of_detected_paths, err := sast.DetectPackageJson(cloned_repo_path)
	if err != nil {
		fmt.Println("❗️ ~ funcmain ~ err ~ Error in Detecting package.json :", err)
	}
	fmt.Println("🚀 ~ funcmain ~ list_of_detected_paths:", list_of_detected_paths)

	// if list_of_detected_paths is not empty then go ahead with identify package manager
	if len(list_of_detected_paths) != 0 {
		package_manager_paths := sast.IdentifyPackageManager(list_of_detected_paths)
		if err != nil {
			fmt.Println("❗️ ~ funcmain ~ err ~ Error in identifying package managers :", err)
		}
		fmt.Println("🚀 ~ funcmain ~ package_manager_paths :", package_manager_paths)
		// return package_manager_paths;

		err := sast.RunPackageManager(package_manager_paths)
		if err != nil {
			fmt.Println("❗️ ~ Error running package managers:", err)
		} else {
			fmt.Println("🚀 ~ Dependencies installed successfully.")
		}
	} else {
		fmt.Println("🚀 No package.json file found")
	}

	framework_with_paths, err := sast.DetectFramework(cloned_repo_path)
	if err != nil {
		fmt.Println("❗️ ~ funcSearchTargets ~ err:", err)
	}
	fmt.Println("🚀 ~ funcSearchTargets ~ framework_with_paths:", framework_with_paths)

	for framework, dirs := range framework_with_paths {
		fmt.Printf("Framework: %s\n", framework)
		for _, dir := range dirs {
			fmt.Printf("  %s\n", dir)
		}
	}

	if len(framework_with_paths) == 0 {
		err := sast.MoveSolidityFiles(cloned_repo_path)
		if err != nil {
			log.Fatalf("Failed to set up Hardhat project: %v", err)
		}
		if err := sast.SetupHardhatProject(cloned_repo_path); err != nil {
			log.Fatalf("Failed to set up Hardhat project: %v", err)
		}
	}

	result, err := sast.CompileProject(framework_with_paths)
	if err != nil {
		log.Fatalf("Error compiling projects: %v", err)
	}
	resultJSON, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println("Compilation Results:")
	fmt.Println(string(resultJSON))

	return resultJSON, err

}
