package main

import "fmt"

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(id int) {
	b.id = id
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary float64
}

func main() {
	employee := new(Employee)
	employee.Person.SetId(10)
	fmt.Println(employee.Person.Id())

}
