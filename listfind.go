package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for it := l.Head; it != nil; it = it.Next {
		if comp(it.Data, ref) {
			return &it.Data
		}
	}
	return nil
}
