package main

import "fmt"

func main() {
	c := make(chan int)
	go sum(1, 2, c)
	fmt.Println(<-c)
}

func sum(a int, b int, c chan int) {
	c <- a + b
}
