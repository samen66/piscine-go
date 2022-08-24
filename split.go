package piscine

func Split(s, sep string) []string { // helloHIhow // HI
	var arr []string
	first := 0

	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			arr = append(arr, s[first:i])
			first = i + len(sep)
		}
	}
	arr = append(arr, s[first:])
	return arr
}
