package main

import (
	"fmt"
	"time"
)

func main() {
	start := "2017-01-13T00:00:00Z"
	t, err := time.Parse(time.RFC3339, start)
	if err != nil {
		fmt.Println(err)
	}
	t = t.AddDate(0, 0, -t.Day()+1)
	fmt.Println(t)
}
