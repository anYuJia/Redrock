package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a, cmd, c int
	fmt.Scanf("%d%c%d\n", &a, &cmd, &c)
	r := rand.Intn(1)
	result := 0
	if r == 0 {
		fmt.Println(1)
		result = Cmd1(a, cmd, c)
	}
	if r == 1 {
		fmt.Println(2)
		result = Cmd2(a, cmd, c)
	}
	fmt.Println(result)
}
func Cmd1(a int, cmd int, c int) (result int) {
	if cmd == '+' {
		return a + c
	} else if cmd == '-' {
		return a - c
	} else if cmd == '*' {
		return a * c
	} else if cmd == '/' {
		return a / c
	} else {
		return 0
	}
}
func Cmd2(a int, cmd int, b int) (result int) {
	switch cmd {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	default:
		return 0
	}
}
