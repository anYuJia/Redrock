package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num0 := make([]int, 0)
	num1 := make([]int, 0)
	num2 := make([]int, 0)
	for i := 0; i < 100; i++ {
		num0 = append(num0, rand.Intn(1000))
	}
	for i := 0; i < 100; i++ {
		num1 = append(num1, rand.Intn(1000))
	}
	for i := 0; i < 100; i++ {
		num2 = append(num2, rand.Intn(1000))
	}
	fmt.Println("算法一")
	Sort0(num0)
	fmt.Println(num0)
	fmt.Println("算法二")
	Sort1(num1)
	fmt.Println(num1)
	fmt.Println("算法三")
	Sort2(num2)
	fmt.Println(num0)

}
func Sort0(num []int) {
	a, b := 0, 0
	for a = 0; a < len(num); a++ {
		for b = 0; b < len(num); b++ {
			if num[a] < num[b] {
				num[a], num[b] = num[b], num[a]
			}
		}
	}
}
func Sort1(num []int) {
	a, b, min := 0, 0, 0
	for a = 0; a < len(num)-1; a++ {
		min = a
		for b = a + 1; b < len(num); b++ {
			if num[min] > num[b] {
				min = b
			}
		}

	}
	if a != min {
		num[a], num[min] = num[min], num[min]
	}
}
func Sort2(num []int) {
	i := 0
	j := 0
	k := 0
	temp := 0
	for k = len(num) >> 1; k > 0; k >>= 1 {
		for i = k; i < len(num); i++ {
			temp = num[i]
			for j = i - k; j >= 0 && num[j] > temp; j -= k {
				num[j+k] = num[j]
			}
			num[j+k] = temp
		}
	}
}
