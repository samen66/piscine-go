package piscine

func Compact(ptr *[]string) int {
	var arr []string

	for _, val := range *ptr {
		if val != "" {
			arr = append(arr, val)
		}
	}
	*ptr = arr

	return len(*ptr)
}
