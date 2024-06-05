package piscine

func Unmatch(a []int) int {
	c := make([]int, len(a))
	for i, num := range a {
		for j := 0; j < len(a); j++ {
			if num == a[j] {
				c[i]++
			}
		}
	}
	for i, count := range a {
		if c[i]%2 != 0 {
			return count
		}
	}
	return -1
}
