package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	gs := 100
	incrementer := 0
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			m.Unlock()
			fmt.Println("\n-------------\n", incrementer)
			fmt.Println(" ",runtime.NumGoroutine())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(incrementer)

}