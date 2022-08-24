package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args
	for _, v := range arr[1:] {
		for _, vv := range v {
			z01.PrintRune(vv)
		}
		z01.PrintRune('\n')
	}
}
