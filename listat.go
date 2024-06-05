package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	for c := 0; c < pos; c++ {
		if l == nil {
			return nil
		}
		l = l.Next
	}
	return l
}
