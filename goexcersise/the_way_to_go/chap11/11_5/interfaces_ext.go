package main

import (
	"fmt"
	"math"
)

type PeriInterface interface {
	Perimeter()
}

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return sq.side * 4
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type Circle struct {
	radius float32
}

func (c *Circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

type AreaInterface interface {
	Area(x, y float64) float64
}

type triangle struct {
	base   float64
	height float64
}

func (t *triangle) Area() float64 {
	return t.base * t.height * 0.5
}

func main() {
	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	c := &Circle{2.5}
	fmt.Println("Looping through shapes for area ...")
	// shapes := []Shaper{Shaper(r), Shaper(q), Shaper(c)}
	shapes := []Shaper{r, q, c}
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
