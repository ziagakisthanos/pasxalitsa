package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if len(s)-i < len(toFind) {
			break
		}
		match := true
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
