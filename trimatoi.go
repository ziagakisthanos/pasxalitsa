package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	new := 0

	for _, char := range s {
		if char == '-' && new == 0 {
			sign = -1
		}
		if char >= '0' && char <= '9' {
			new = new*10 + int(char-'0')
		}
	}
	return sign * new
}
