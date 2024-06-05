package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l.Head == nil {
		return nil
	}

	node := l.Head
	for node != nil {
		if comp(node.Data, ref) {
			return &node.Data
		}
		node = node.Next
	}
	return nil
}
