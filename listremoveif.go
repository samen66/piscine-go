package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL
	for it := l.Head; it != nil; it = it.Next {
		if it.Data == data_ref {
			if prev == nil {
				l.Head = it.Next
			} else {
				prev.Next = it.Next
			}
		} else {
			prev = it
		}
	}
}
