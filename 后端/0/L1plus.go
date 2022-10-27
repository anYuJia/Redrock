package main

import "fmt"

func main() {
	var a, b, c int
	var result int
	_, err := fmt.Scanf("%d%c%d", &a, &c, &b)
	if err != nil {
		return
	}
	if c == '+' {
		result = a + b
	}
	if c == '-' {
		result = a - b
	}
	if c == '*' {
		result = a * b
	}
	if c == '/' {
		result = a / b
	}
	fmt.Printf("%d%c%d=%d", a, c, b, result)
}
