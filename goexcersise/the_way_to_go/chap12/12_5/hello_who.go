package main

import (
	"fmt"
	"os"
)

func main() {
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	fmt.Print("Hello")
	for _, v := range args {
		fmt.Print(" ", v)
	}
}
