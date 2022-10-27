package main

import "fmt"

func main() {
	var zh string
	var passWorld string
	data := map[string]string{"2211": "2211", "4244": "4244", "sdk432": "sdk432"}
	fmt.Println("请输入你的账号")
	fmt.Scanln(&zh)
	fmt.Println("请输入你的密码")
	fmt.Scanln(&passWorld)
	if data[zh] == passWorld {
		fmt.Println("登陆成功")
	} else {
		fmt.Println("密码错误")
	}
}
