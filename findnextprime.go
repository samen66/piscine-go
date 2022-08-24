package piscine

func FindNextPrime(nb int) int {
	isP := true
	if nb <= 2147483648 {
		if nb <= 2 {
			return 2
		} else if nb%2 == 0 {
			isP = false
		} else {
			for i := 3; i <= nb/3; i += 2 {
				if nb%i == 0 {
					isP = false
				}
			}
		}
		if isP {
			return nb
		} else {
			return FindNextPrime(nb + 1)
		}

	}
	return 0
}
