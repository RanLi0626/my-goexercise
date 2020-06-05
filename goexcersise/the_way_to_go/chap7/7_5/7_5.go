package main

import "fmt"

func main() {
	sl := make([]byte, 10)
	fmt.Println(cap(sl))
	b := [10]byte{'H', 'E', 'L', 'L', 'O', 'L', 'I', 'R', 'A', 'N'}
	for _, v := range b {
		sl = append(sl, v)
	}
	fmt.Println(string(sl))
	fmt.Println(cap(sl))
}

func Append(slice, data []byte) []byte {
	return []byte{}
}
