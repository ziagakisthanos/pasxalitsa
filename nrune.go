package piscine

func NRune(s string, n int) rune {
	convert := []rune(s)
	for n > 0 && n <= len(s) {
		return convert[n-1]
	}
	return 0
}
