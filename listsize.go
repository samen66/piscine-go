package piscine

func ListSize(l *List) int {
	size := 0
	it := l.Head
	for it != nil {
		size++
		it = it.Next
	}
	return size
}
