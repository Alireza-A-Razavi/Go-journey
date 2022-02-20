package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func main() {

  c := make(chan int)

  wg.Add(2)
  // send
  go foo(c)

  // revieve
  go bar(c)
  wg.Wait()

  fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
  wg.Done()
  c <- 42
}

// recieve
func bar(c <-chan int) {
  wg.Done()
  fmt.Println(<-c)
}
