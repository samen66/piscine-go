package main

import (
	"github.com/01-edu/z01"
)

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

type point struct {
	x int
	y int
}

func main() {
	points := &point{}

	setPoint(points)

	stringx := "x = "
	printStr(stringx)
	n := points.x
	printStr(ReversStr(IntToStr(n)))

	stringy := ", y = "
	printStr(stringy)
	printStr(ReversStr(IntToStr(points.y)))
	z01.PrintRune('\n')
}

func IntToStr(nbr int) string {
	s := ""
	for ; nbr > 0; nbr = nbr / 10 {
		s += string(rune(nbr%10) + 48)
	}
	return s
}

func ReversStr(s string) string {
	arrR := []rune(s)
	length := len(arrR)
	for i := 0; i < length/2; i++ {
		arrR[i], arrR[length-1-i] = arrR[length-1-i], arrR[i]
	}
	return string(arrR)
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
