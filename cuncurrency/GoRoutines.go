package main

import "fmt"

func DoSomeThing(x int) int {
	return x * 5
}

func main(){
	ch := make(chan int)
	go func() {
		ch <- DoSomeThing(5)
	}()
	fmt.Println(<-ch)
}

