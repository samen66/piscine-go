package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "galaxy" || arg == "01" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
