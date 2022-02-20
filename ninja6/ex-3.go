package main

import "fmt"

func foo() {
	fmt.Println("Hello I am Foo")
}

func bar() {
	fmt.Println("And hello I am bar")
}

func main() {
	defer bar()
	foo()
}
