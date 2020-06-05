package main

import "fmt"

func main() {
	var arr1 [3]int
	var arr2 = new([3]int)
	for i := range arr2 {
		arr2[i] = i
	}
	fmt.Println(arr2)
	arr1 = *arr2
	fmt.Println(arr1)

	for i := range arr2 {
		arr2[i] = i + 1
	}
	fmt.Println(arr2)
	fmt.Println(arr1)
}
