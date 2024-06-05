package piscine

func DescendAppendRange(max, min int) []int {
	ans := []int{}
	if min > max || min == max {
		return []int{}
	}
	for i := max; i > min; i-- {
		ans = append(ans, i)
	}
	return ans
}
