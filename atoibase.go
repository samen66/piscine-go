package piscine

func AtoiBase(s string, base string) int {
	lenBase := len([]rune(base))

	if lenBase >= 2 {
		for i := 0; i < lenBase; i++ {
			if base[i] == '-' || base[i] == '+' {
				return 0
			}
			for j := i + 1; j < lenBase; j++ {
				if base[i] == base[j] {
					return 0
				}
			}
		}
		var pows int = 0
		result := 0
		for i := len(s) - 1; i >= 0; i-- {
			isfind := false
			for j := 0; j < len(base); j++ {
				if s[i] == base[j] {
					isfind = true
					powsBase := 1
					for k := 0; k < pows; k++ {
						powsBase = powsBase * lenBase
					}
					pows++
					result = result + j*powsBase

				}
			}
			if !isfind {
				return 0
			}
		}
		return result

	} else {
		return 0
	}
}
