package main

import (
	"fmt"
)

type SecretAgent struct {
	name string
	age int
	ltk bool
}

func main() {
	sa_1 := SecretAgent{
		age: 29,
		name: "james bond",
		ltk: true,
	}
	fmt.Println(*sa_1.Shoot())
}

func (sa *SecretAgent) Shoot() string {
	return sa.name + "\nThis is me shooting"
}