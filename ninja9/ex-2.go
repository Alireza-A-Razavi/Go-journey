package main

import "fmt"

type person struct {
	name string
	age int
	height float64
}

type human interface {
	speak() string
}

func main() {
	p1 := person{
		name:   "bungo",
		age:    12,
		height: 124.54,
	}
	SaySomething(&p1)
}

func (p person) speak() string{
	return "This is a person speaking: I am " + p.name + " "
}

func SaySomething(h human) {
	fmt.Println(h.speak())
}
