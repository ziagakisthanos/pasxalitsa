package piscine

func BasicAtoi2(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	for _, char := range s {
		if char >= '0' && char <= '9' {
			n = n*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return n
}
