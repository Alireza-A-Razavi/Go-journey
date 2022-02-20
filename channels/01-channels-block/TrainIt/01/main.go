package main

import "fmt"

func sum(s []int, c chan int, ch chan string){
  sum := 0
  for _, v := range s {
    sum += v
  }
  c <- sum //send sum to c
  ch <- "Hello there! from beyond the channel"
}

func main() {
    s := []int{7, 4, -2, 10, -10., -9}

    c := make(chan int)
    ch := make(chan string)
    go sum(s, c, ch)
    go sum(s[:len(s)/2], c, ch)
    x, y := <- c, <- ch
    t, z := <- c, <- ch
    fmt.Println(x, y)
    fmt.Println(t, z)
}
