package main

import "fmt"

func sum(xi ...int) (int, int) {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total, len(xi)
}

func avg(f func(xi ...int) (int, int), vi ...int) int {
	s, lenght := f(vi...)
	return s / lenght
}

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7}
	average := avg(sum, ii...)
	fmt.Println(average)
}
