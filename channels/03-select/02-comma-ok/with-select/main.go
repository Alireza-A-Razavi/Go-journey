package main

import (
  "fmt"
)

func main() {
  even := make(chan int)
  odd := make(chan int)
  quit := make(chan int)

  go send(even , odd, quit)

  receive(even , odd, quit)
}

func send(e, o, q chan<- int) {
  for i := 0; i < 15; i++ {
    if i % 2 == 0 {
      e <- i
    } else {
      o <- i
    }
  }
  q <- 0
  close(q)
}

func receive(e, o, q <-chan int) {
  for {
    select {
    case v := <-e:
      fmt.Println("this is from even channel: ", v)
    case v := <-o:
      fmt.Println("this is from odd channel: ", v)
    case i, ok := <-q:
      if !ok {
        fmt.Println("from cooma ok", i, ok)
        return
      } else {
        fmt.Println("from comma ok", i)
      }
    }
  }
}
