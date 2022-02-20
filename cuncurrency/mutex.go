package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("Go Routines: ", runtime.NumGoroutine())
	counter := 0
	var wg sync.WaitGroup
	const gs = 100
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func(){
			mu.Lock()
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("counter: ", counter)
}
