package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go tel(ch, done)
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-done:
			os.Exit(0)
		}
	}
}

func tel(ch chan int, done chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	done <- true
}
