package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if a directory path is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <directory>")
		return
	}

	root := os.Args[1] // Get the directory path from command-line arguments

	// Initialize counters
	var totalFiles int
	var totalSize int64

	// Walk the directory
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Print file details
		fmt.Printf("File: %s, Size: %d bytes\n", path, info.Size())

		// Update counters
		totalFiles++
		totalSize += info.Size()

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory: %v\n", err)
		return
	}

	// Print summary
	fmt.Printf("\nTotal files: %d\n", totalFiles)
	fmt.Printf("Total size: %d bytes\n", totalSize)
}
