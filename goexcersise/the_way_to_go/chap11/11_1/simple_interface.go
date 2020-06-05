package main

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	value int
}

func (s *Simple) Get() int {
	return s.value
}

func (s *Simple) Set(value int) {
	s.value = value
}

func method(s Simpler, value int) {
	s.Set(value)
	fmt.Println(s.Get())
}

func main() {
	var s Simple
	method(&s, 10)
}
