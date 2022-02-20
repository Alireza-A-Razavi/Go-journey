package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Money Penny",
		age:   28,
	}

	people := []person{
		p1,
		p2,
	}

	fmt.Println(people)
	bs, err := json.Marshal(people) // bs is byte slice
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs)) // converts our byte slice to string
}
