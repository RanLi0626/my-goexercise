package main

import "fmt"

func main() {
	m := map[string]string{"A": "a", "B": "b"}
	for k := range m {
		fmt.Println(k)
	}
}
