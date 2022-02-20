package main

import (
  "fmt"
)

// you can not use runtime WaitGroup
// and just remove Go routine for bar function

func main() {

  c := make(chan int)

  // send
  go foo(c)

  // revieve
  bar(c)

  fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
  c <- 42
}

// recieve
func bar(c <-chan int) {
  fmt.Println(<-c)
}
