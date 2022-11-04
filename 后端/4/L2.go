package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("芜湖！起飞！\n")
			}
		}
	}()
	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				h, min, s := time.Now().Clock()
				if h == 18 && min == 14 && s == 50 {
					fmt.Println("我还能再战4小时！")
				}
				if h == 6 && min == 0 && s == 0 {
					fmt.Println("我要去图书馆开卷！")
				}
			}
		}

	}()
	for {
	}
}
