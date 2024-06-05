package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		// If no arguments provided, read from stdin and print to stdout
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			printError("Error reading standard input: ", err)
			os.Exit(1)
		}
	} else {
		// If arguments provided, read each file and print its content
		for _, fileName := range os.Args[1:] {
			file, err := os.Open(fileName)
			if err != nil {
				printError("ERROR: open "+fileName+": no such file or directory", nil)
				os.Exit(1)
			}
			defer file.Close()

			_, err = io.Copy(os.Stdout, file)
			if err != nil {
				printError("ERROR: reading file "+fileName+": ", err)
				os.Exit(1)
			}
		}
	}
}

func printError(message string, err error) {
	for _, ch := range message {
		z01.PrintRune(ch)
	}
	if err != nil {
		for _, ch := range err.Error() {
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
