package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for _, s := range "zyxwvutsrqponmlkjihgfedcba" {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}
