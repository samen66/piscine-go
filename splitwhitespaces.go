package piscine

func SplitWhiteSpaces(s string) []string {
	var arr []string
	first := 0

	for i, v := range s {
		if v == ' ' {
			if s[first:i] != "" {
				arr = append(arr, s[first:i])
			}

			first = i + 1
		}
	}
	arr = append(arr, s[first:])
	return arr
}
