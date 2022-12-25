package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	stay := make(chan int)
	go Work("goroutine1", stay)
	<-stay
	go Work("goroutine2", stay)
	<-stay
	go Work("goroutine3", stay)
	<-stay
	wg.Wait()
	fmt.Println("successful")
}

func Work(workName string, stay chan<- int) {
	time.Sleep(time.Second) // 模拟逻辑处理
	fmt.Println(workName)
	wg.Done()
	stay <- 1
}
