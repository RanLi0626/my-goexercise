package main

import "fmt"

func main() {
	var result [50]int
	func(n int) {
		for i := 0; i < n; i++ {
			if i <= 1 {
				result[i] = 1
				//result = append(result, 1)
			} else {
				result[i] = result[i-1] + result[i-2]
				//result = append(result, result[i-1]+result[i-2])
			}
		}
	}(50)
	fmt.Println(result)
}
