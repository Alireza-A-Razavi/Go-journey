package main

import "fmt"

func main()  {
  cs := make(chan<- int) // send

  cs <- 42

  fmt.Println(<-cs) // throws recive from send-only channel Error 
  fmt.Println("-----------")
  fmt.Printf("%T\n", cs)\
}
