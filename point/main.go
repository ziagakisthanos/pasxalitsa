package main

import (
	"github.com/01-edu/z01"
)

func setPoint(x, y *int) {
	*x = 42
	*y = 21
}

func intToDigits(num int) []rune {
	if num == 0 {
		return []rune{'0'}
	}

	var digits []rune
	for num > 0 {
		digit := '0' + rune(num%10)
		digits = append([]rune{digit}, digits...)
		num /= 10
	}
	return digits
}

func printDigits(digits []rune) {
	for _, digit := range digits {
		z01.PrintRune(digit)
	}
}

func main() {
	var x, y int

	setPoint(&x, &y)

	// Print x
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printDigits(intToDigits(x))
	z01.PrintRune(',')

	// Print y
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printDigits(intToDigits(y))
	z01.PrintRune('\n')
}
