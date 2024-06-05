package piscine

func Abort(a, b, c, d, e int) int {
	median := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if median[i] > median[j] {
				median[i], median[j] = median[j], median[i]
			}
		}
	}
	return median[2]
}
