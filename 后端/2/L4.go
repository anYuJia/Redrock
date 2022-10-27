package main

type Interface interface {
	Len() int           // 获取元素数量
	Less(i, j int) bool // i，j是序列元素的指数。
	Swap(i, j int)      // 交换元素
}

//没读懂题（抱歉
