package piscine

func IsPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
