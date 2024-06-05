package piscine

func Map(f func(int) bool, a []int) []bool {
	x := make([]bool, len(a))
	for i, char := range a {
		x[i] = f(char)
	}
	return x
}
