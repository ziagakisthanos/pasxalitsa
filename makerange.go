package piscine

func MakeRange(min, max int) []int {
	if min > max || min == max {
		return nil
	}
	length := max - min
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		ans[i] = min + i
	}
	return ans
}
