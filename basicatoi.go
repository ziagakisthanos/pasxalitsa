package piscine

func BasicAtoi(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for _, char := range s {
		n = n*10 + int(char-'0')
	}
	return n
}
