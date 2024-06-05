package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

// 	var n32 int32 = int32(n & 0xFFFFFFFF)

// 	// Check if the original number was negative
// 	if n < 0 {
// 		// For negative numbers, we need to ensure the sign is preserved
// 		// This is a bit tricky because simply casting to int32 will not work correctly
// 		// We'll use bitwise operations to manually set the sign bit
// 		n32 = int32(n & 0x7FFFFFFF) // Clear the sign bit
// 		n32 = -n32                  // Negate the value to get the correct negative number
// 	}

// 	// Now, n32 is an int32 representation of the original int64 value
// 	// Proceed with printing as before, using the int32 value
// 	if n32 < 0 {
// 		z01.PrintRune('-')
// 		n32 = -n32 // Work with the absolute value for printing
// 	}
// 	if n32/10 != 0 {
// 		PrintNbr(int(n32 / 10)) // Recursively call PrintNbr with the next digit
// 	}
// 	z01.PrintRune(rune(n32%10 + '0'))
// }
