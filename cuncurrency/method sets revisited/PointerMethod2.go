package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

type closed_shape interface {
	perimeter() float64
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func information(cs closed_shape) {
	fmt.Print("perimeter: ", cs.perimeter())
}

func (c circle) perimeter() float64 {
	return c.radius * 2 * math.Pi
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	c := circle{5}
	info(c)
	information(c)
}