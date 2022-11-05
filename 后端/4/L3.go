package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// 用户第一次登录提醒修改密码
var k int

type UserData struct {
	UserName interface{} `json:"userName"`
	PassWord interface{} `json:"passWord"`
}

func main() {
	for {
		initTerminal()
	}

}

func CreateUser(userdata UserData) {
	filePtr, err := os.OpenFile("4/users.data.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	dataMap := new([]UserData)
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range *dataMap {
		//fmt.Println("用户名:", user.UserName, "密码:", user.PassWord)
		if user.UserName == userdata.UserName {
			fmt.Println("这个用户名已经注册过了，换一个用户名吧！")
			return
		}
	}
	*dataMap = append(*dataMap, userdata)
	err = filePtr.Close()
	if err != nil {
		return
	}
	filePtr, err = os.OpenFile("4/users.data.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer filePtr.Close()
	if err != nil {
		return
	}
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(*dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("注册成功！！！！！")
}
func InputUser() {
	var userdata UserData
	var name, word string
	fmt.Println("请输入你的用户名...")
	_, err := fmt.Scanln(&name)
	if err != nil {
		println("错误", err)
		return
	}
	fmt.Println("请输入密码")
	_, err = fmt.Scanln(&word)
	if err != nil {
		return
	}
	userdata.UserName = name
	userdata.PassWord = word
	passWordStr := fmt.Sprintf("%v", userdata.PassWord)
	var a, A, _1 int
	for _, char := range passWordStr {
		if char <= 'z' && char >= 'a' {
			a++
		}
		if char <= 'Z' && char >= 'A' {
			A++
		}
		if char >= '1' && char <= '9' {
			_1++
		}
	}
	if a == 0 || A == 0 || _1 == 0 {
		fmt.Println("密码中必须要包含大写字母(A-Z)小写字母(a-z)数字(0-9)")
		return
	}
	if len(passWordStr) < 6 {
		fmt.Println("密码必须大于6位")
	}
	CreateUser(userdata)
}
func initTerminal() {
	var a string
	fmt.Println("请输入你的选项：1.注册  2.登录  3.退出")
	fmt.Scanln(&a)
	switch a {
	case "1":
		InputUser()
	case "2":
		Login()
	case "3":
		fmt.Println("感谢使用我的垃圾登陆界面")
		time.Sleep(time.Second * 1)
		fmt.Println("正在退出中.....")
		time.Sleep(time.Second * 1)
		fmt.Println("已经退出.....")
		os.Exit(0)
	default:
		fmt.Println("不要乱输")
	}
}
func Login() {
restart:
	var count int
	var userdata UserData
	var name, word string
	fmt.Println("请输入你的用户名...")
	_, err := fmt.Scanln(&name)
	if err != nil {
		println("错误", err)
		return
	}
reInput:
	if count != 0 {
		fmt.Println("密码错误！！请输入密码，密码为q返回主界面，密码为f重新输入用户名")
	} else {
		fmt.Println("请输入密码")
	}

	_, err = fmt.Scanln(&word)
	if err != nil {
		return
	}
	userdata.UserName = name
	userdata.PassWord = word
	if word == "f" {
		goto restart
	}
	if word == "q" {
		return
	}
	filePtr, err := os.OpenFile("4/users.data.json", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	dataMap := new([]UserData)
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range *dataMap {
		if user.UserName == userdata.UserName {
			if user.PassWord == userdata.PassWord {
				fmt.Println("正在登录....")
				time.Sleep(time.Second * 1)
				fmt.Println("正在加载界面.....")
				time.Sleep(time.Second * 1)
				fmt.Println("恭喜你登陆成功")

				if k == 0 {
					fmt.Println("你的密码有点拉，是否要改密码！(y/n)")
				} else {
					fmt.Println("你的密码很安全！输入q退出")
				}
				var t string
				fmt.Scanln(&t)
				if t == "y" {
					ChangePassword(user)
					println("修改密码成功，请重新登陆")
					k++
					goto reInput
				} else if t == "q" {
					return
				}
			} else {
				count++
				goto reInput
			}
		}
	}
	fmt.Println("用户名不存在！请先注册用户！！")
}
func ChangePassword(userdata UserData) {
	filePtr, err := os.OpenFile("4/users.data.json", os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return
	}
	dataMap := new([]UserData)
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, user := range *dataMap {
		//fmt.Println("用户名:", user.UserName, "密码:", user.PassWord)
		if user.UserName == userdata.UserName {
			var newPassWord1, newPassWord2 string
		reInput:
			fmt.Println("请输入你的新密码")
			_, err = fmt.Scanln(&newPassWord1)
			if err != nil {
				return
			}
			fmt.Println("请再次输入你的新密码")
			_, err = fmt.Scanln(&newPassWord2)
			if err != nil {
				return
			}
			if newPassWord1 == newPassWord2 {
				var a, A, _1 int
				for _, char := range newPassWord1 {
					if char <= 'z' && char >= 'a' {
						a++
					}
					if char <= 'Z' && char >= 'A' {
						A++
					}
					if char >= '1' && char <= '9' {
						_1++
					}
				}
				if a == 0 || A == 0 || _1 == 0 {
					fmt.Println("密码中必须要包含大写字母(A-Z)小写字母(a-z)数字(0-9)")
					goto reInput
				} //密码规范
				if len(newPassWord1) < 6 {
					fmt.Println("密码必须大于6位")
					goto reInput
				} //密码大于6
				(*dataMap)[i].PassWord = newPassWord1
				err = filePtr.Close()
				if err != nil {
					return
				}
				filePtr, err = os.OpenFile("4/users.data.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
				if err != nil {
					return
				}
				encoder := json.NewEncoder(filePtr)
				err = encoder.Encode(*dataMap)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer filePtr.Close()
				fmt.Println("密码修改成功！！！！！")
			} else {
				fmt.Println("两次密码不相等")
				goto reInput
			}
		}
	}
}
