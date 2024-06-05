package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		l.Tail = nil
		return
	}
	prev := l.Head
	for curr := l.Head.Next; curr != nil; curr = curr.Next {
		if curr.Data == data_ref {
			prev.Next = curr.Next
			if curr == l.Tail {
				l.Tail = prev
			}
		} else {
			prev = curr
		}
	}
}
