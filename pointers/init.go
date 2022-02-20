package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)

	fmt.Println(*b)

	fmt.Println(*&a)

	*b = 98123
	fmt.Println(a)

	// & gives you the address and * gives you the the value that is stored in address

}
