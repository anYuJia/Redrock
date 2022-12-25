package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go findNumbers(i, &wg)
	}
	wg.Wait()
}
func findNumbers(digits int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 计算最大值
	max := int(math.Pow(10, float64(digits)))
	for i := 1; i < max; i++ {
		if isNumber(i, digits) {
			fmt.Println(i)
		}
	}
}
func isNumber(n, digits int) bool {
	sum := 0
	temp := n
	for temp > 0 {
		d := temp % 10
		sum += int(math.Pow(float64(d), float64(digits)))
		temp /= 10
	}
	return n == sum
}
