package piscine

func CountIf(f func(string) bool, tab []string) int {
	c := 0
	for _, str := range tab {
		if f(str) {
			c++
		}
	}
	return c
}
