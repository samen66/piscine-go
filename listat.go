package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0

	it := l
	for it != nil {

		if counter == pos {
			return it
		}
		counter++
		it = it.Next
	}
	return it
}
