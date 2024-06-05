package piscine

func AlphaCount(s string) int {
	r := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' || r[i] >= 'a' && r[i] <= 'z' {
			count++
		}
	}
	return count
}
