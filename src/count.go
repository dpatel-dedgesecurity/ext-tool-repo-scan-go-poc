package main

import (
	"bufio"
	"os"
	"regexp"
)

func countOccurrences(filename, pattern string) int {
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		match, _ := regexp.MatchString(pattern, line)
		if match {
			count++
		}
	}

	return count
}

func countLines(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}
