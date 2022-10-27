package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func add(ch chan int64) {
	var i int64
	for i = 0; i < 50000; i++ {
		i++
	}
	ch <- i
	wg.Done()
}
func main() {
	ch := make(chan int64)
	wg.Add(2)
	go add(ch)
	go add(ch)
	x := <-ch + <-ch
	wg.Wait()
	fmt.Println(x)
}
