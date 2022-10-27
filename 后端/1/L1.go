package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Println("input:")
	fmt.Scanf("%s", &str)
	str_p := str
	str_len := len([]rune(str))
	for i := 0; i < (str_len-1)/2; i++ {
		if str[i] != str_p[str_len-1-i] {
			fmt.Println("erorr!")
			os.Exit(0)
		}
	}
	for i := 0; i < (str_len)/2; i++ {
		if int(str[i]) < 10000 {
			fmt.Printf("%c", str[i])
		} else {
			fmt.Printf("%s", str[i])
		}

	}

}
