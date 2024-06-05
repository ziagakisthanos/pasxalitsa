package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) >= 0 {
		for i := 1; i < len(args); i++ {
			for _, char := range args[i] {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
	}
}
