package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[3:]
	n := StrToInt(os.Args[2:3][0])
	if Cheker(args) {
		return
	}
	find := false
	for i, v := range args {
		if !Read(v, n, i) {
			find = true
		}
	}
	if find {
		os.Exit(1)
	}
}

func Cheker(args []string) bool {
	if len(args) == 0 {
		return true
	}

	return false
}

func Read(filename string, n, i int) bool {
	file, err := os.OpenFile(filename, os.O_RDONLY, 7777)
	defer file.Close()
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return false
	}
	data := make([]byte, 64)
	var FinalData []byte

	for {
		cont, err := file.Read(data)
		if err != nil { // если конец файла
			break // выходим из цикла
		}
		for _, v := range data[:cont] {
			FinalData = append(FinalData, v)
		}

	}
	if i != 0 {
		fmt.Printf("\n")
	}
	fmt.Printf("==> %v <==\n", file.Name())
	if n >= len(FinalData) {
		fmt.Printf("%v", string(FinalData))
	} else {
		fmt.Printf("%v", string(FinalData)[len(FinalData)-n:])
	}

	// printStr("\n")
	// fmt.Print(string(cunt))
	// return string(arr)
	return true
}

func StrToInt(nbr string) int {
	var s int = 0
	for _, v := range nbr {
		s = s*10 + int(v-48)
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
