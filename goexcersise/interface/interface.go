package main

import "fmt"

type A struct {
	year int
}

func (a A) Talk() string { return "a talk" }

func (a A) PrintTalk() { fmt.Printf("print talk %s", a.Talk()) }

func (a A) Greet() { fmt.Println("Hello GolangUK", a.year) }

type B struct {
	A
}

func (b B) Talk() string { return "b talk" }

func (b B) Greet() { fmt.Println("Welcome to GolangUK", b.year) }

func main() {
	var a A
	a.year = 2016
	var b B
	b.year = 2017
	a.Greet()     // Hello GolangUK 2016
	b.Greet()     // Welcome to GolangUK 2017
	b.A.Greet()   // Hello GolangUK 2017
	b.PrintTalk() // print talka talk
}
