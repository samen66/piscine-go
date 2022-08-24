package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	if args[0][0] == '-' {
		if !(IsNomeric(args[0][1:])) {
			return
		}
	} else {
		if !(IsNomeric(args[0])) {
			return
		}
	}
	if args[2][0] == '-' {
		if !(IsNomeric(args[2][1:])) {
			return
		}
	} else {
		if !(IsNomeric(args[2])) {
			return
		}
	}
	if !(IsOperater(args[1])) {
		return
	}

	nbr1, isTrue := StrToInt(args[0])
	nbr2, isTrue2 := StrToInt(args[2])
	if !isTrue || !isTrue2 {
		return
	}

	result, ok := ActioinOperator(nbr1, nbr2, rune(args[1][0]))
	if !ok {
		return
	}

	fmt.Println(result)
}

func ActioinOperator(n1, n2 int, operator rune) (int, bool) {
	firstEdition := 0
	secondEdition := 0
	if n1 >= n2 {
		firstEdition = n1
		secondEdition = n2
	} else {
		firstEdition = n2
		secondEdition = n1
	}

	switch operator {
	case '+':
		{
			// fmt.Println("treu +")
			therd := Plus(n1, n2)
			if firstEdition <= 0 {
				if firstEdition < therd {
					return 0, false
				}
			} else if firstEdition > 0 {
				if secondEdition < 0 {
					if firstEdition < therd {
						return 0, false
					}
				} else {
					if firstEdition > therd {
						return 0, false
					}
				}
			}
			return therd, true
		}
	case '-':
		{
			// fmt.Println("treu -")
			therd := Minus(n1, n2)
			if firstEdition < 0 {
				if firstEdition < therd {
					return 0, false
				}
			} else if firstEdition >= 0 {
				if secondEdition < 0 {
					// fmt.Println("treu <0", firstEdition)
					if firstEdition < therd {
						return 0, false
					}
				} else {
					if firstEdition < therd {
						return 0, false
					}
				}
			}
			return therd, true
		}
	case '/':
		{
			// fmt.Println("treu /")
			if n2 == 0 {
				fmt.Println("No division by 0")
				return 0, false
			}
			therd := Devide(n1, n2)

			return therd, true
		}
	case '*':
		{
			// fmt.Println("treu *")
			therd := Multiply(n1, n2)
			if firstEdition < 0 {
				if firstEdition > therd {
					return 0, false
				}
			} else if firstEdition > 0 {
				if secondEdition > 0 {
					if firstEdition > therd {
						return 0, false
					}
				} else {
					if firstEdition < therd {
						return 0, false
					}
				}
			}
			return therd, true
		}
	case '%':
		{
			// fmt.Println("treu %")
			if n2 == 0 {
				fmt.Println("No modulo by 0")
				return 0, false
			}
			therd := Modulo(n1, n2)

			return therd, true
		}
	}
	return 0, false
}

func StrToInt(s string) (int, bool) {
	nbr := 1
	if s == "0" {
		return 0, true
	}
	if s[0] == '-' {
		nbr = nbr * (-1)
		s = s[1:]
		if s == "0" {
			return 0, true
		}
	}
	kob := 1
	for i, v := range s {
		firstEdition := nbr
		if i == 0 {
			nbr = nbr * int(v-48)
			if nbr < 0 {
				kob = kob * -1
			}

		} else {
			nbr = Plus(nbr*10, int(v-48)*kob)
		}

		if firstEdition < 0 {
			if firstEdition < nbr {
				return 0, false
			}
		} else if firstEdition > 0 {
			if firstEdition > nbr {
				return 0, false
			}
		}
	}
	return nbr, true
}

func Plus(n1, n2 int) int {
	return n1 + n2
}

func Minus(n1, n2 int) int {
	return n1 - n2
}

func Multiply(n1, n2 int) int {
	return n1 * n2
}

func Devide(n1, n2 int) int {
	return n1 / n2
}

func Modulo(n1, n2 int) int {
	return n1 % n2
}

func IsNomeric(s string) bool {
	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			return false
		}
	}
	return true
}

func IsOperater(s string) bool {
	for _, v := range s {
		if !(v == '+' || v == '-' || v == '/' || v == '*' || v == '%') {
			return false
		}
	}
	return true
}
