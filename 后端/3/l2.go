package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	for i := 1; i <= 100; i++ {
		go Print(ch)
		go Print(ch)
		ch <- i
		i++
		ch <- i
		time.Sleep(time.Millisecond)
	}
}
func Print(ch chan int) {
	fmt.Printf("%d  ", <-ch)
}
