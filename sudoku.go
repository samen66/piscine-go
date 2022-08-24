package piscine

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]

	allok, arrR := addAll(arr)

	if !allok {
		fmt.Println("Error")
		return
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !Row_and_Colum(i, j, arrR[i][j], arrR) {
				fmt.Println("Error")
				return
			}
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(string(arrR[i][j]), " ")
		}
		fmt.Println()
	}
	fmt.Println("------------------")

	for k := '1'; k <= '9'; k++ {
		if Solved(k, arrR) {
			Prindts(arrR)
			fmt.Println("------------------")
			break
		}
	}
}

func addAll(arr []string) (bool, [][]rune) {
	var result [][]rune
	if len(arr) == 9 { // check the Args has 9 arguments
		for _, argument := range arr {
			row := []rune(argument)
			if len(row) != 9 { // checking if argument has 9 rune
				return false, result
			} else { // if argument has 9 rune
				for _, r := range argument {
					if !isNameric(r) { // checks argument's rune is numeric or point
						return false, result
					}
				}
				result = append(result, row) // if all OK appends result slice
			}
		}
		return true, result
	}

	return false, result
}

func Solved(r rune, arr [][]rune) bool {
	ok, i, j := nextPoint(arr)
	if !ok {
		return true
	}
	if Cheker(i, j, r, arr) {
		arr[i][j] = r

		for k := '1'; k <= '9'; k++ {
			if Solved(k, arr) {
				return true
			}
		}
		arr[i][j] = '.'

	}
	return false
}

func arr_3to3(x, y int, r rune, arr [][]rune) bool {
	tatalrow := x - x%3
	tatalColum := y - y%3

	for i := tatalrow; i < tatalrow+3; i++ {
		for j := tatalColum; j < tatalColum+3; j++ {
			if arr[i][j] == r && arr[i][j] != '.' {
				return false
			}
		}
	}
	return true
}

func nextPoint(arr [][]rune) (bool, int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if arr[i][j] == '.' {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

func PrindsString(s string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Prindts(arr [][]rune) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(string(arr[i][j]), " ")
		}
		fmt.Println()
	}
}

/////////////////

func Cheker(x, y int, r rune, arr [][]rune) bool {
	if Row_and_Colum(x, y, r, arr) && arr_3to3(x, y, r, arr) {
		return true
	}
	return false
}

func isNameric(r rune) bool {
	if (r >= '1' && r <= '9') || r == '.' {
		return true
	}
	return false
}

func Row_and_Colum(x, y int, r rune, arr [][]rune) bool {
	for i := 0; i < 9; i++ {
		if arr[x][i] != '.' && i != y {
			if arr[x][i] == r {
				return false
			}
		}
		if arr[i][y] != '.' && x != i {
			if arr[i][y] == r {
				return false
			}
		}
	}
	return true
}
