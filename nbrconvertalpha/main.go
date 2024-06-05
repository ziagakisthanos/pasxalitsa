package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:] // Remove the --upper flag
	}

	for _, arg := range args {
		n := atoi(arg)
		if n >= 1 && n <= 26 {
			letter := rune('a' - 1 + n)
			if upper {
				letter -= 32 // Convert to uppercase
			}
			z01.PrintRune(letter)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func atoi(s string) int {
	res := 0
	for _, c := range s {
		res = res*10 + int(c-'0')
	}
	return res
}
