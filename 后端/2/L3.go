package main

import "fmt"

func main() {
	num := []int{9, 6, -8, 5, 2, 1, 11, 0, 54, 2, 1, 34, 1, 9}
	Sort(num)
	for _, value := range num {
		fmt.Printf("%3d", value)
	}
}

// 定义排序函数
func Sort(num []int) {
	i, j, min := 0, 0, 0
	for i = 0; i < len(num)-1; i++ {
		min = i
		for j = i + 1; j < len(num); j++ {
			if num[min] > num[j] {
				min = j
			}
		}
		if i != min {
			num[i], num[min] = num[min], num[i]
		}
	}
}
