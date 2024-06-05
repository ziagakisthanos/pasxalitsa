package piscine

func SortIntegerTable(table []int) {
	if table == nil {
		return
	}
	for i := 0; i <= len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[j], table[i] = table[i], table[j]
			}
		}
	}
}
