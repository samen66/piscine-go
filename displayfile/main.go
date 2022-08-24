package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileName := os.Args[1:]
	if !Cheker(fileName) {
		return
	}

	Read(fileName[0])
	// fmt.Print(CuntentFile)
}

func Cheker(args []string) bool {
	if len(args) == 1 {
		return true
	} else if len(args) == 0 {
		fmt.Println("File name missing")
		return false
	} else {
		fmt.Println("Too many arguments")
		return false
	}
}

func Read(filename string) {
	file, _ := os.Open(filename)

	cunt, _ := ioutil.ReadAll(file)
	fmt.Print(string(cunt))
	// return string(arr)
}
