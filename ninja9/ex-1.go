package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("This from a normal one")
	wg.Done()
}

func main()  {
	go func() {
		fmt.Println("This is from an ananymous function")
		wg.Done()
	}()

	go foo()

	wg.Add(2)
	wg.Wait()
}




