package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Exclude the program name from arguments
	// Sort the arguments in ASCII order
	for i := 0; i < len(args)-1; i++ {
		for j := i + 1; j < len(args); j++ {
			if compare(args[i], args[j]) > 0 {
				args[i], args[j] = args[j], args[i] // Swap
			}
		}
	}
	// Print the sorted arguments
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func compare(s1, s2 string) int {
	i := 0
	for i < len(s1) && i < len(s2) {
		if s1[i] != s2[i] {
			return int(s1[i]) - int(s2[i])
		}
		i++
	}
	return len(s1) - len(s2)
}
