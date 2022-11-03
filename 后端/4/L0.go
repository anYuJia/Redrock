package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("开始计时")
	//timer := time.NewTimer(time.Second * 3) //声明一个timer， time.Second是time包里声明的一个Duration
	// 等待设置的时间
	//fmt.Println("时间为3秒之后", <-timer.C)
	fmt.Println("aa")
	t := time.AfterFunc(time.Second*4, func() {
		fmt.Println("延迟两秒")
	})
	time.Sleep(time.Second * 5)
	fmt.Println("sda")
	defer t.Stop()
}
