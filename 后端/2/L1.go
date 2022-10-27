package main

import (
	"fmt"
	"os"
)

// 定义猫猫的结构体
type Cat struct {
	id    int    //编号
	name  string //名字
	color string //颜色
	sex   rune   //性别
}

// 定义信息输出函数
func (cat Cat) say() error {
	fmt.Println("id:", cat.id)
	fmt.Println("name:", cat.name)
	fmt.Println("color:", cat.color)
	fmt.Printf("sex:%c\n", cat.sex)
	return nil
}

// 实现更改猫猫信息的函数
func (cat *Cat) Change(key string, value interface{}) (err string) {
	switch key {
	case "id":
		cat.id = value.(int)
	case "name":
		cat.name = value.(string)
	case "color":
		cat.color = value.(string)
	case "sex":
		cat.sex = value.(rune)
	default:
		return "抱歉，这个猫猫没有这个属性"
	}
	return ""
}
func main() {
	cat := Cat{12, "五仁", "橘色", 'f'}
	err := cat.say() //原本属性输出
	if err != nil {
		fmt.Println(err)
	}
	e := cat.Change("id", 15)
	e = cat.Change("name", "顺拐")
	//e = cat.Change("s", "s") //更改属性
	if e != "" {
		fmt.Println(e) //如果输入错了报错终止
		os.Exit(0)
	}
	cat.say()
}
