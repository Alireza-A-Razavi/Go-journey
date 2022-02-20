package main

import (
	"fmt"
	"runtime"
	"sync"
)
// run the program with -race argument in terminal
func main() {
	var wg sync.WaitGroup

	gs := 100
	incrementer := 0
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println("\n-------------\n", incrementer)
			fmt.Println(" ",runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)

}