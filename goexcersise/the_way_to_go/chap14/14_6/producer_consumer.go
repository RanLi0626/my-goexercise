package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	go producer(c, done)
	go consumer(c)
	<-done
}

func producer(c chan int, done chan bool) {
	for i := 0; i < 100; i = i + 10 {
		c <- i
	}
	done <- true
}

func consumer(c chan int) {
	for {
		fmt.Println(<-c)
	}
}
