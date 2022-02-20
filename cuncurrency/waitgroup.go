package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	a := foo()
	go fmt.Printf(a)

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println(a)
}

func foo() string {
	wg.Done()
	fmt.Println("This is from println func")
	return "This is printed out from foo function, but has been RETURNED"
}
