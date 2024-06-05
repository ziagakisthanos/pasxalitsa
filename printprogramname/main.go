package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) >= 0 {
		// Find the last occurrence of the path separator
		lastSlashIndex := -1
		for i := len(args[0]) - 1; i >= 0; i-- {
			if args[0][i] == '/' {
				lastSlashIndex = i
				break
			}
		}
		// Extract the program name
		programName := args[0][lastSlashIndex+1:]
		// Print the program name directly
		for _, char := range programName {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
