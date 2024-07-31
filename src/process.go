package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func processSolidityFiles(solidityFilesFile, resultsFile string) error {

	abstractContractFiles := []string{}
	smartContractFiles := []string{}
	libraryFiles := []string{}
	interfaceFiles := []string{}
	multipleTypeFiles := []string{}

	abstractContractCount := 0
	smartContractCount := 0
	libraryCount := 0
	interfaceCount := 0
	multipleTypeCount := 0

	abstractContractTotal := 0
	smartContractTotal := 0
	libraryTotal := 0
	interfaceTotal := 0
	loctotal := 0

	file, err := os.Open(solidityFilesFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		solfile := scanner.Text()

		// Count occurrences of each type
		abstractContractCountFile := countOccurrences(solfile, `abstract\s+contract\s+\w+`)
		smartContractCountFile := countOccurrences(solfile, `^\s*contract\s+\w+`)
		libraryCountFile := countOccurrences(solfile, `^\s*library\s+\w+`)
		interfaceCountFile := countOccurrences(solfile, `^\s*interface\s+\w+`)

		// Update cumulative counts
		abstractContractTotal += abstractContractCountFile
		smartContractTotal += smartContractCountFile
		libraryTotal += libraryCountFile
		interfaceTotal += interfaceCountFile

		loctotalFile := countLines(solfile)
		loctotal += loctotalFile

		// Determine file types
		types := []string{}
		if abstractContractCountFile > 0 {
			types = append(types, fmt.Sprintf("Abstract Contract (%d)", abstractContractCountFile))
		}
		if smartContractCountFile > 0 {
			types = append(types, fmt.Sprintf("Smart Contract (%d)", smartContractCountFile))
		}
		if libraryCountFile > 0 {
			types = append(types, fmt.Sprintf("Library (%d)", libraryCountFile))
		}
		if interfaceCountFile > 0 {
			types = append(types, fmt.Sprintf("Interface (%d)", interfaceCountFile))
		}
		fileTypes := strings.Join(types, ", ")

		// Store files based on their types
		if len(types) > 1 {
			multipleTypeFiles = append(multipleTypeFiles, solfile)
			multipleTypeCount++
		} else {
			if abstractContractCountFile > 0 {
				abstractContractFiles = append(abstractContractFiles, solfile)
				abstractContractCount++
			}
			if smartContractCountFile > 0 {
				smartContractFiles = append(smartContractFiles, solfile)
				smartContractCount++
			}
			if libraryCountFile > 0 {
				libraryFiles = append(libraryFiles, solfile)
				libraryCount++
			}
			if interfaceCountFile > 0 {
				interfaceFiles = append(interfaceFiles, solfile)
				interfaceCount++
			}
		}

		// Write details to results file
		result := fmt.Sprintf("%s: %s: %d lines\n", solfile, fileTypes, loctotalFile)
		err := appendToFile(resultsFile, result)
		if err != nil {
			return err
		}
	}

	// Display counts
	err = appendToFile(resultsFile, "\nSummary:\n")
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Total number of Solidity files: %d\n", countLines(solidityFilesFile)))
	if err != nil {
		fmt.Println(" error in total no of soli files")
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Abstract Contracts Files: %d\n", abstractContractCount))
	if err != nil {
		fmt.Println(" error in abstract files")
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Smart Contracts Files: %d\n", smartContractCount))
	if err != nil {
		fmt.Println(" error in smart contracts files")
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Libraries Files: %d\n", libraryCount))
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Interfaces Files: %d\n", interfaceCount))
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Multiple Types Files: %d\n", multipleTypeCount))
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, "******************************************\n")
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Abstract Contracts: %d\n", abstractContractTotal))
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Smart Contracts: %d\n", smartContractTotal))
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Libraries: %d\n", libraryTotal))
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("Interfaces: %d\n", interfaceTotal))
	if err != nil {
		return err
	}
	err = appendToFile(resultsFile, fmt.Sprintf("LOC: %d\n", loctotal))
	if err != nil {
		return err
	}

	return nil
}
