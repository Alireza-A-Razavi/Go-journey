package main

import "fmt"

func main() {
	func() {
		fmt.Println("shitt anonymous")
	}()

	func(x int) {
		fmt.Println(x)
	}(42)
}
