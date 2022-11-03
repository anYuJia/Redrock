package main

import "fmt"

func main() {
	var n int
	var skill_name string
	fmt.Println("请输入你的释放技能语言")
	fmt.Println("1.大家好，我是练习时长两年半的个人练习生")
	fmt.Println("2.你食不食油饼")
	fmt.Println("3.哎哟！！！你干嘛！！")
	fmt.Scanln(&n)
	fmt.Println("请输入技能名字")
	fmt.Scanln(&skill_name)
	switch n {
	case 1:
		ReleaseSkill(skill_name, func(skillName string) {
			fmt.Println("大家好，我是练习时长两年半的个人练习生", skillName)
		})
	case 2:
		ReleaseSkill(skill_name, func(skillName string) {
			fmt.Println("你食不食油饼", skillName)
		})
	case 3:
		ReleaseSkill(skill_name, func(skillName string) {
			fmt.Println("哎哟！！！你干嘛！！", skillName)
		})
	default:
		fmt.Println("重新来过")
	}
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
