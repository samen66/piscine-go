package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args
	for i := len(arr) - 1; i >= 1; i-- {
		for _, v := range arr[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
