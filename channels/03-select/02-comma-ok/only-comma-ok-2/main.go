package main

import (
  "fmt"
)


func main() {
  c := make(chan int)

  // send
  go func(){
    c <- 42
    close(c)
  }()

  v, ok := <-c
  fmt.Println(v, ok)

  v, ok = <-c
  fmt.Println(v, ok)

}
