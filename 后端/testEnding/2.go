package main

import "fmt"

func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()

	if a {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()
}
