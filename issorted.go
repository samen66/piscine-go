package piscine

// 0 1 2 3 4 5
func IsSorted(f func(a, b int) int, a []int) bool {
	isSortedToHigh := false
	isSortedToLow := false
	// 0, 1, 2, 3, 4, 5
	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) > 0 {
			if isSortedToLow {
				return false
			}
			isSortedToHigh = true

		} else if f(a[i], a[i+1]) < 0 {
			if isSortedToHigh {
				return false
			}
			isSortedToLow = true
		}
	}
	return true
}
