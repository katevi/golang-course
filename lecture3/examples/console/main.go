package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineCount := 0
	wordCount := 0

	fmt.Println("Enter lines of text (press Enter on an empty line to finish):")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		lineCount++
		words := strings.Fields(line) // Split line into words
		wordCount += len(words)       // Add to total word count

		fmt.Println(strings.ToUpper(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}

	fmt.Printf("Total lines entered: %d\n", lineCount)
	fmt.Printf("Total words entered: %d\n", wordCount)
}
