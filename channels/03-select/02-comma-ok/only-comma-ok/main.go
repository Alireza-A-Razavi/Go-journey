package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup

func main() {
  c := make(chan string)

  // send
  go func(){
    c <- "42"
    c <- "43"
    c <- "44"
  }()

  v, ok := <-c
  fmt.Println(v, ok)

  v, ok = <-c
  fmt.Println(v, ok)

  v, ok = <-c
  fmt.Println(v, ok)

  v, ok = <-c
  fmt.Println(v, ok)

  wg.Done()

}
