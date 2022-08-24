package piscine

func Map(f func(int) bool, a []int) []bool {
	var arrBool []bool

	for _, v := range a {
		arrBool = append(arrBool, f(v))
	}
	return arrBool
}
