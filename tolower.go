package piscine

func ToLower(s string) string {
	upper := ""
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			upper += string(char + 32)
		} else {
			upper += string(char)
		}
	}
	return upper
}
