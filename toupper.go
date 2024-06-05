package piscine

func ToUpper(s string) string {
	lower := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			lower += string(char - 32)
		} else {
			lower += string(char)
		}
	}
	return lower
}
