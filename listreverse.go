package piscine

import "fmt"

func ListReverse(l *List) {
	curr := l.Head
	first := l.Head

	// var next *NodeL
	var prev *NodeL

	for curr != nil {
		fmt.Printf("current := %v current.Next := %v \n", curr.Data, curr.Next.Data)
		curr, prev, curr.Next = curr.Next, curr, prev
		fmt.Printf("current := %v  prev := %v after\n", curr.Data, prev.Data)
	}
	l.Head = prev
	l.Tail = first
}
