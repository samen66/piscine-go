package piscine

func ConcatParams(args []string) string {
	res := ""
	for i, v := range args {
		if i != len(args)-1 {
			res += v + "\n"
		} else {
			res += v
		}
	}
	return res
}
