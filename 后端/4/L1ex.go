package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var n int
	fmt.Println("介绍：输入数字")
	skillNames := make(map[string]string)
	for {
		fmt.Printf("1.添加技能  2.释放技能   3.初始化技能   4.输入其它退出程序\n")
		_, err := fmt.Scanln(&n)
		if err != nil {
			println(err)
		}
		switch n {
		case 1:
			CustomSkillName(skillNames)
		case 2:
			Release(skillNames)
		case 3:
			InitSkill(skillNames)
		default:
			fmt.Println("game over!!!!!!!!!")
			goto end
		}
	}
end:
}
func ReleaseSkill1(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
func Release(skillName map[string]string) {
	var i int
	m := rand.Intn(len(skillName))
	for key, value := range skillName {
		if m == i {
			ReleaseSkill1(key, func(skillName string) {
				fmt.Println("是的，", value, "看我给你终极大绝招------", skillName)
			})
		}
		i++
	}
}
func CustomSkillName(skillName map[string]string) {
	var name, word string
	fmt.Println("请输入技能名")
	fmt.Scanln(&name)
	fmt.Println("请输入文字")
	fmt.Scanln(&word)
	var words []string = []string{"kun", "i", "坤", "饼", "篮球"} //敏感词库
	for _, value := range words {
		if strings.ContainsAny(value, word) || strings.ContainsAny(value, name) {
			println("不要这么黑人家")
			return
		}
	}
	skillName[name] = word
}
func InitSkill(skillName map[string]string) {
	skillName["坤坤之球"] = "大家好，我是练习时长两年半的个人练习生"
	skillName["坤之rap"] = "你食不食油饼"
	skillName["爱kun撒娇"] = "哎哟！！！你干嘛！！"
}
