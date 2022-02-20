package main

import "fmt"

func main()  {
  c := make(chan int, 2)

  c <- 42
  c <- 43

  fmt.Println(<-c) // the first one gets the first one and the 2nd the 2nd 
  fmt.Println(<-c)
}
