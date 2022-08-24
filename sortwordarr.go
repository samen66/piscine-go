package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j][0] > a[j+1][0] {
				a[j], a[j+1] = a[j+1], a[j]
			} else if a[j][0] == a[j+1][0] {
				if CompareToStr(a[j], a[j+1]) {
					a[j], a[j+1] = a[j+1], a[j]
				}
			}
		}
	}
}

func CompareToStr(s1, s2 string) bool { // aa aab
	lenS1 := len(s1)
	lenS2 := len(s2)
	if lenS1 > lenS2 {
		for i := 0; i < lenS1; i++ {
			if i <= lenS2-1 {
				if s1[i] < s2[i] {
					return false
				} else if s1[i] > s2[i] {
					return true
				}
			} else {
				if s1[i] < s2[lenS2-1] {
					return false
				} else if s1[i] > s2[lenS2-1] {
					return true
				}
			}
		}
	} else {
		for i := 0; i < lenS2; i++ {
			if i <= lenS1-1 {
				if s1[i] < s2[i] {
					return false
				} else if s1[i] > s2[i] {
					return true
				}
			} else {
				if s1[lenS1-1] < s2[i] {
					return false
				} else if s1[lenS1-1] > s2[i] {
					return true
				}
			}
		}
	}
	return false
}
