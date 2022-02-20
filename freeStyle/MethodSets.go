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

type closed interface {
	perimeter() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) diameter() float64 {
	return c.radius * 2
}

func (c *circle) perimeter() float64 {
	return math.Pi * c.radius * 2
}

func info(s shape) {
	fmt.Println("area:\t:", s.area())
	fmt.Println("circle:\t", s.perimeter())
	fmt.Println("diameter:\t")
}

func main() {
	c := circle{radius: 3.2}
	// info(c)
	fmt.Println(c.area())
	fmt.Println()
}
