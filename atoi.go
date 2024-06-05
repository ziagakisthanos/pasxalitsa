package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	n := 0
	j := 0
	if s[0] == '-' || s[0] == '+' {
		j += 1
	}
	for _, char := range s[j:] {
		if char >= '0' && char <= '9' {
			n = n*10 + int(char-'0')
		} else {
			return 0
		}
	}
	if s[0] == '-' {
		n *= -1
	}
	return n
}
