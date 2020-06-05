package main

import (
	"fmt"
	"sync"
)

var s []int

func appendValue(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	s = append(s, i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go appendValue(i, &wg)
	}
	wg.Wait()

	for i, v := range s {
		fmt.Println(i, ":", v)
	}
}
