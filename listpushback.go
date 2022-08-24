package piscine

func ListPushBack(l *List, data interface{}) {
	note := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = note
		l.Tail = note
	} else {
		l.Tail.Next = note
		l.Tail = note
	}
}
