package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func printNames(dir string) error {
	var files []string

	// Collect all files and directories
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Exclude the root directory
		if path != dir {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	// Sort the files and directories
	sort.Strings(files)

	// Print the sorted files and directories
	for _, file := range files {
		fmt.Println(file)
	}
	return nil
}

func main() {
	if err := printNames("pics"); err != nil {
		fmt.Println("Error:", err)
	}
}
