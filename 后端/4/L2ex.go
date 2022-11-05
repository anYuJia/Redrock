package main

import (
	"fmt"
	"os"
	"time"
)

type Ticker1 struct {
	ticker *time.Ticker
	time   int
	info   string
	doing  bool
}

func main() {

	var tickers []Ticker1
	for {
		var x int
		fmt.Println("1.增加定时器,2.删除定时器，3.查看定时器,4.启动定时器,5.停止计时器   输入其它按键就退出！")
		_, err := fmt.Scanln(&x)
		if err != nil {
		}
		switch x {
		case 1:
			buildTicker(&tickers)
		case 2:
			removeTicker(&tickers)
		case 3:
			viewTicker(tickers)
		case 4:
			startTicker(&tickers)
		case 5:
			closeTicker(&tickers)
		default:
			fmt.Println("退出中......")
			time.Sleep(time.Second * 2)
			os.Exit(0)
		}
	}
}
func buildTicker(tickers *[]Ticker1) {
	var s int
	var name string
	fmt.Println("请输入定时时间(单位为s)")
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("请输入该定时器的名字")
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
	}
	t := time.NewTicker(time.Second * time.Duration(s))
	ticker := Ticker1{
		t, s, name, false,
	}
	*tickers = append(*tickers, ticker)
}
func viewTicker(tickers []Ticker1) {
	println("现在有如下定时器：")
	for n, ticker := range tickers {
		println(n+1, ticker.info, "计时时间：", ticker.time)
	}
}
func removeTicker(tickers *[]Ticker1) {
re:
	var i int
	fmt.Println("请输入要删除计时器的序号")
	_, err := fmt.Scanln(&i)
	if err != nil {
		fmt.Println(err)
	}
	if i < 1 || i > len(*tickers) {
		fmt.Println("你输入的序号错了")
		goto re
	}
	*tickers = append((*tickers)[:i-1], (*tickers)[i:]...)
}
func startTicker(tickers *[]Ticker1) {
restart:
	var j, g int
	fmt.Println("你要启动计时器的序号")
	_, err := fmt.Scanln(&j)
	if err != nil || j < 1 || j > len(*tickers) {
		println("请输入正确序号")
		goto restart
	}
	for i, _ := range *tickers {
		if i == j-1 {
			fmt.Println("输入1个数字任意退出计时器")
			var k rune
			(*tickers)[i].ticker = time.NewTicker(time.Second * time.Duration((*tickers)[i].time))
			go func() {
				for {
					select {
					case <-(*tickers)[i].ticker.C:
						g++
						(*tickers)[i].doing = true
						fmt.Println((*tickers)[i].info, "----正在计时中....第", g, "次")
					}
				}
			}()
			for {
				n, err := fmt.Scanln(&k)
				if n != 0 {
					return
				}
				if err != nil {
				}
			}
		}
	}
}
func closeTicker(tickers *[]Ticker1) {
	fmt.Println("现在正在运行的计时器有")
	for i, ticker := range *tickers {
		if ticker.doing {
			fmt.Println(i+1, ticker.info)
		}
	}
	fmt.Println("你要停止哪一个计时器")
	var x int
	fmt.Scanln(&x)
	(*tickers)[x-1].ticker.Stop()
	(*tickers)[x-1].doing = false
	fmt.Println((*tickers)[x-1].info, "已经停止")
}
