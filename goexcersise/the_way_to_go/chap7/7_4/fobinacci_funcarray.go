package main

import "fmt"

func main() {
	result := fib(10)
	fmt.Println(result)
}

func fib(n int) []int {
	slice := make([]int, n)
	for i := 0; i < len(slice); i++ {
		if i <= 1 {
			slice[i] = 1
		} else {
			slice[i] = slice[i-1] + slice[i-2]
		}
	}
	return slice
}
