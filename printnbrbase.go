package piscine

func PrintNbrBase(nbr int, base string) string {
	result := ""
	lenBase := len(base)
	var arrBase []int
	if lenBase >= 2 {
		for i := 0; i < lenBase; i++ {
			if base[i] == '-' || base[i] == '+' {
				return result
			}
			for j := i + 1; j < lenBase; j++ {
				if base[i] == base[j] {
					return result
				}
			}
		}
		for nbr > 0 {

			kaldyl := nbr % lenBase
			arrBase = append(arrBase, kaldyl)
			nbr = nbr / lenBase
		}

		for i := len(arrBase) - 1; i >= 0; i-- {
			for j, v := range base {
				if arrBase[i] == j {
					result += string(v)
				}
			}
		}
		return result

	} else {
		return result
	}
}
