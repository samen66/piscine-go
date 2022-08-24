package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			n := f(a[j], a[j+1])
			if n > 0 {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
