package main

import (
	"fmt"
	"math"
)

type square struct {
	lenght float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.lenght * s.width
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	c1 := circle{
		radius: 3,
	}

	sq1 := square{
		lenght: 3,
		width:  4,
	}

	info(c1)
	info(sq1)
}
