package main

import "fmt"

type FilterFunc func(in int) bool
func (ff FilterFunc) Filte(in int) bool {
	return ff(in)
}

func main() {
	//f:= FilterFunc()
	FilterFunc().Filte(14)
	fmt.Println()
}