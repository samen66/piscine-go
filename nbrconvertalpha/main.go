package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args
	if len(arr) > 1 {
		if arr[1] == "--upper" {
			PrintNumberToAlpha(arr[2:], true)
		} else {
			PrintNumberToAlpha(arr[1:], false)
		}
	}
}

func PrintNumberToAlpha(arr []string, isUpper bool) {
	for _, val := range arr {
		arrArgs := []rune(val)
		if len(arrArgs) == 1 {
			if arrArgs[0] >= '1' && arrArgs[0] <= '9' {
				if isUpper {
					Prints(arrArgs[0], 16)
				} else {
					Prints(arrArgs[0], 48)
				}
			} else {
				z01.PrintRune(' ')
			}
		} else if len(arrArgs) == 2 {
			if arrArgs[0] == '1' {
				if arrArgs[1] >= '0' && arrArgs[1] <= '9' {
					if isUpper {
						Prints(arrArgs[1], 26)
					} else {
						Prints(arrArgs[1], 58)
					}
				} else {
					z01.PrintRune(' ')
				}
			} else if arrArgs[0] == '2' {
				if arrArgs[1] >= '0' && arrArgs[1] <= '6' {
					if isUpper {
						Prints(arrArgs[1], 36)
					} else {
						Prints(arrArgs[1], 68)
					}
				} else {
					z01.PrintRune(' ')
				}
			} else {
				z01.PrintRune(' ')
			}
		}

	}
	z01.PrintRune('\n')
}

func Prints(R rune, n int) {
	z01.PrintRune(rune(int(R) + n))
}
