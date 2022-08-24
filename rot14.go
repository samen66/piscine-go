package piscine

func Rot14(s string) string {
	arr := []rune(s)
	for i, v := range arr {
		if v >= 'a' && v <= 'l' {
			arr[i] = arr[i] + 14
		} else if v >= 'm' && v <= 'z' {
			arr[i] = arr[i] - 12
		} else if v >= 'A' && v <= 'L' {
			arr[i] = arr[i] + 14
		} else if v >= 'M' && v <= 'Z' {
			arr[i] = arr[i] - 12
		}
	}
	return string(arr)
}
