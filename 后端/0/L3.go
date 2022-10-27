package main

import "fmt"

func main() {
	var x int
	fmt.Println("请输入你要输入的数字个数：")
	fmt.Scanln(&x)
	fmt.Println("请输入数字")
	var num = make([]int, x, x)
	for i := 0; i < x; i++ {
		fmt.Scanln(&num[i])
	}
	Sort(num, x)
	for i := 0; i < x; i++ {
		fmt.Printf("%d  ", num[i])
	}
}
func Sort(num []int, x int) (num1 []int) {
	a, b := 0, 0
	for ; a < x; a++ {
		min := a
		for b = a; b < x; b++ {
			if num[b] < num[min] {
				min = b
			}
		}
		if a != min {
			num[a], num[min] = num[min], num[a]
		}
	}
	return num1
}
