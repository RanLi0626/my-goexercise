package main

import (
	"fmt"
	"math"
)

type Magnitude interface {
	Abs() float64
}

type Point2 struct {
	X, Y float64
}

type Point3 struct {
	X, Y, Z float64
}

type Polar struct {
	R, T float64
}

func (p *Point2) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func Scale(p *Point2, s float64) (q Point2) {
	q.X = p.X * s
	q.Y = p.Y * s
	return
}

func main() {
	p1 := new(Point2)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs())

	p2 := &Point2{4, 5}
	fmt.Printf("The length of the vector p2 is: %f\n", p2.Abs())

	q := Scale(p1, 5)
	fmt.Printf("The length of the vector q is: %f\n", q.Abs())
	fmt.Printf("Point p1 scaled by 5 has the following coordinates: X %f - Y %f", q.X, q.Y)
}
