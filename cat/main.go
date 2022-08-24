package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if Cheker(args) {
		io.Copy(os.Stdout, os.Stdin)
		// os.Stdout = os.Stdin
	} else {
		for i, v := range args {
			if i == len(args)-1 {
				if !Read(v) {
					os.Exit(1)
				}
			} else if Read(v) {
			}
		}
	}
}

func Cheker(args []string) bool {
	if len(args) == 0 {
		return true
	}

	return false
}

func Read(filename string) bool {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		// fmt.Print("ERROR: ")
		printStr("ERROR: ")
		printStr(err.Error())
		printStr("\n")
		os.Exit(1)
		return false
	}

	cunt, _ := ioutil.ReadAll(file)
	printStr(string(cunt))
	// printStr("\n")
	// fmt.Print(string(cunt))
	// return string(arr)
	return true
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
