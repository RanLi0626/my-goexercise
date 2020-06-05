package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "decimal:"
	keysAndParams := strings.Split(s, ":")
	params := strings.Split(keysAndParams[1], ",")
	var ps []interface{}
	for _, param := range params {
		i, err := strconv.Atoi(param)
		if err == nil {
			ps = append(ps, i)
			continue
		}
		b, err := strconv.ParseBool(param)
		if err == nil {
			ps = append(ps, b)
			continue
		}
		ps = append(ps, param)
	}
	for _, param := range ps {
		fmt.Printf("%T, %v\n", param, param)
	}
}
