package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	result := ""
	if len(arr) < 1 {
		z01.PrintRune('\n')
	} else {
		for i, v := range arr {
			if i != len(arr)-1 {
				result += v + " "
			} else {
				result += v
			}
		}

		arrRes := []rune(result)
		j := len(arrRes) - 1
		for i, v := range result {
			if isGlasnis(v) {
				for j > i {
					if isGlasnis(rune(result[j])) {
						arrRes[i], arrRes[j] = arrRes[j], arrRes[i]
						j--
						break
					}
					j--
				}
			}
		}
		for _, v := range arrRes {
			z01.PrintRune(v)
		}
		z01.PrintRune(' ')
		z01.PrintRune('\n')
	}
}

func isGlasnis(v rune) bool {
	if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
		return true
	}
	return false
}

///wvwervwerwervwerv ervrev
