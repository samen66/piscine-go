package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args
	var arrInsert []string
	var arrOrder []string
	if len(arr) > 1 {
		if isIt(arr[1], "-i") {
			if len(arr[1]) > 2 && arr[1][2] == '=' {
				arrInsert = append(arrInsert, arr[1][3:])
			}
			AddToOrderINsert(arr, &arrOrder, &arrInsert, 2)
			for i, v := range arrOrder {
				arrOrder[i] = SortingByASCII(v)
			}
			if len(arrOrder) > 0 {
				arrInsert = append(arrInsert, arrOrder...)
				res := ""
				for _, v := range arrInsert {
					res += v
				}
				res = SortingByASCII(res)
				fmt.Println(res)
			} else {
				for i := len(arrInsert) - 1; i >= 0; i-- {
					fmt.Print(arrInsert[i])
				}
				for _, arrOv := range arrOrder {
					fmt.Print(arrOv)
				}
				fmt.Println()
			}

		} else if isIt(arr[1], "--insert") {
			if len(arr[1]) > 8 && arr[1][8] == '=' {
				arrInsert = append(arrInsert, arr[1][9:])
			}
			AddToOrderINsert(arr, &arrOrder, &arrInsert, 2)
			for i, v := range arrOrder {
				arrOrder[i] = SortingByASCII(v)
			}

			if len(arrOrder) > 0 {
				arrInsert = append(arrInsert, arrOrder...)
				res := ""
				for _, v := range arrInsert {
					res += v
				}
				res = SortingByASCII(res)
				fmt.Println(res)

			} else {
				for i := len(arrInsert) - 1; i >= 0; i-- {
					fmt.Print(arrInsert[i])
				}
				for _, arrOv := range arrOrder {
					fmt.Print(arrOv)
				}
				fmt.Println()
			}

		} else if isIt(arr[1], "--help") {
			PrintHelpPage()
		} else if isIt(arr[1], "-h") {
			PrintHelpPage()
		} else {
			AddToOrderINsert(arr, &arrOrder, &arrInsert, 1)
			for i, v := range arrOrder {
				arrOrder[i] = SortingByASCII(v)
			}
			for _, arrOv := range arrOrder {
				fmt.Print(arrOv)
			}
			fmt.Println()
		}
	} else {
		PrintHelpPage()
	}
}

func PrintsResult(arrInsert, arrOrder []string) {
}

func isIt(s string, is string) bool {
	arrS := []rune(s)
	arrIs := []rune(is)
	if len(arrS)-len(arrIs) >= 0 {
		if s[0:len(arrIs)] == is {
			return true
		}
	}
	return false
}

func AddToOrderINsert(arr []string, arrOrder, arrInsert *[]string, starts int) {
	findOrder := false
	for _, v := range arr[starts:] {
		if isIt(v, "--order") {
			findOrder = true
			if len(v) > 7 && v[8] == '=' {
				*arrOrder = append(*arrOrder, v[8:])
			}
		} else if isIt(v, "-o") {
			findOrder = true
			if len(v) > 2 && v[3] == '=' {
				*arrOrder = append(*arrOrder, v[4:])
			}
		} else if findOrder {
			*arrOrder = append(*arrOrder, v)
		} else {
			*arrInsert = append(*arrInsert, v)
		}
	}
}

func SortingByASCII(s string) string {
	var arrRune []rune
	for _, v := range s {
		arrRune = append(arrRune, v)
	}

	for i := 0; i < len(arrRune); i++ {
		for j := 0; j < len(arrRune)-1; j++ {
			if arrRune[j] > arrRune[j+1] {
				arrRune[j], arrRune[j+1] = arrRune[j+1], arrRune[j]
			}
		}
	}
	return string(arrRune)
}

func PrintHelpPage() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("	 This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
}
