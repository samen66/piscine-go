package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		l = &NodeI{Data: data_ref}

		return l
	}
	it := l
	first := it
	var prev *NodeI
	temp := &NodeI{Data: data_ref}

	for it != nil {
		if it.Data > data_ref {
			if first != it {
				nexts := prev.Next
				prev.Next = temp
				temp.Next = nexts
				// fmt.Println("reue1")
				return first
			} else {
				newtemp := &NodeI{Data: it.Data}
				newtemp.Next = it.Next
				temp.Next = newtemp
				// fmt.Println("reue2")
				return temp
			}

			// next := it.Next
			// it.Next = &temp
			// temp.Next = next
		}
		prev = it

		it = it.Next

	}
	prev.Next = temp

	return first
}
