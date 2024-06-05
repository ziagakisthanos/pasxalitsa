package piscine

func AppendRange(min, max int) []int {
	ans := []int{}
	if min > max || min == max {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		ans = append(ans, i)
	}
	return ans
}
