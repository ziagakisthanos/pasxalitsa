package piscine

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return true
	}
	for _, char := range s {
		if !((char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z')) {
			return false
		}
	}
	return true
}
