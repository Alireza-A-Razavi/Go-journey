package main

import "fmt"

type user struct {
	first string
	last  string
	age   int
}

func (p user) speak() {
	fmt.Print("Hello I am ", p.first, "and I am ", p.age, " years old")
}

func main() {
	p1 := user{
		first: "Alireza",
		last:  "Alavi",
		age:   23,
	}
	p1.speak()
}
