package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func NewCircle(r float64) *Circle {
	return &Circle{radius: r}
}

func main() {
	c := NewCircle(5)
	fmt.Println("radius: ", c.Area())
	fmt.Println("perimeter: ", c.Perimeter())
}
