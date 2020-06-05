package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

func (c *Car) numberOfWheels() {
	fmt.Printf("car's number of wheels %v", c.wheelCount)
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

type Mecedes struct {
	Car
}

func (m *Mecedes) sayHiToMerkel() {
	fmt.Println("Hi Merkel!")
}

func main() {
	mecedes := new(Mecedes)
	mecedes.sayHiToMerkel()
	mecedes.wheelCount = 3
	mecedes.numberOfWheels()
}
