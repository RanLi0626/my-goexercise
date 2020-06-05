package main

import "fmt"

func main() {
	var m map[string]string
	fmt.Println(m)
	m1 := map[string]string{}
	fmt.Println(m1)
	v, ok := m["test"]
	fmt.Println(ok)
	fmt.Println(v)
}
