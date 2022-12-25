package main

import (
	"fmt"
)

func main() {
	// 输入发出船的数量n
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("船只数量不合法！")
		return
	}
	// 最终结果
	result := make([]int, n)

	// 存储满足条件的船只
	shipMap := make(map[int][]int)

	// 存储不同品质的蛋
	eggMap := make(map[int]struct{})

	// 记录当前发出的船的时间
	var t int

	// 遍历n艘船
	for i := 0; i < n; i++ {
		// 输入船的信息
		var z int //读取上次输入缓冲区的内容
		_, _ = fmt.Scanf("%d", &z)
		_, err = fmt.Scanf("%d", &t)
		var k int
		_, err = fmt.Scanf("%d", &k)
		if err != nil {
			fmt.Println("输入不合法，从新输入")
		}
		// 将满足条件的船存入shipMap
		if t >= 0 {
			shipMap[t-86400] = append(shipMap[t-86400], k)
		}
		// 将每批蛋的品质存入eggMap
		for j := 0; j < k; j++ {
			var x int
			_, err = fmt.Scanf("%d", &x)
			if err != nil {
				fmt.Println("输入不合法")
			}
			eggMap[x] = struct{}{}
		}

		// 计算结果
		result[i] = len(eggMap)
	}

	// 输出结果
	for _, v := range result {
		fmt.Println(v)
	}
}
