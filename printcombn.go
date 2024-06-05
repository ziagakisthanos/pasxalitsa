package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	// Print single digits
	for i, d := range digits {
		z01.PrintRune(d)
		if i < len(digits)-1 {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
