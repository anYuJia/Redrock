package main

import (
	"fmt"
)

type coder struct {
	name       string
	student_id string
	phone_id   string
	code_rows  int
	birthday   string
}

// 输入成绩每输入一个值按下一个回车
// 连按两次回车就结束输入，未输入完全的学生不计入
// code_rows作为分数的标准，150行为及格
// 查找是一个循环，也可以同时查找多个符合条件的人，比如相同名字的人
// 不查找就自动排序然后输出
func main() {
	var stu []coder                //声明空结构体
	stu = AppendStu(stu)            //添加学生
	fmt.Printf("现在的学生有：\n")
	PrintStu(stu)                      //输出学生2
	ok_stu, n := CountStu(stu)          //统计达到150的
	fmt.Printf("一共有%d个人及格：\n", n)
	PrintStu(ok_stu)                   //输出及格的
	FindStu(stu)                       //查找函数
	stu = ChangeStu(stu)               //更改信息函数，不能改姓名
	stu = SortStu(stu)                 //排序
	fmt.Println("排序之后：")
	PrintStu(stu)                      //输出结果
}

// 添加人数
func AppendStu(stu []coder) []coder {
	for i := 0; ; i++ {
		var stu_n coder
		fmt.Println("请输入学生的名字 学号 手机号 代码行数 生日")
		n, err := fmt.Scanln(&stu_n.name)
		if n == 0 {
			break
		}
		n, err = fmt.Scanln(&stu_n.student_id)
		if n == 0 {
			break
		}
		n, err = fmt.Scanln(&stu_n.phone_id)
		if n == 0 {
			break
		}
		n, err = fmt.Scanln(&stu_n.code_rows)
		if n == 0 {
			break
		}
		n, err = fmt.Scanln(&stu_n.birthday)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		stu = append(stu, stu_n)
	}
	return stu
}

// 输出学生
func PrintStu(stu []coder) {
	for i := 0; i < len(stu); i++ {
		fmt.Println("名字：", stu[i].name, "学号", stu[i].student_id, "电话号码", stu[i].phone_id, "代码行数", stu[i].code_rows, "生日", stu[i].birthday)
	}
}

// 统计
func CountStu(stu []coder) ([]coder, int) {
	var stu1 []coder
	var n = 0
	for _, s := range stu {
		if s.code_rows >= 150 {
			n++
			stu1 = append(stu1, s)
		}
	}
	return stu1, n
}

// 查找
func FindStu(stu []coder) {
	var flag string
	for {
		fmt.Println("输入'c'可以按姓名或者学号查找，，输入'f'就不查找，")
		fmt.Scanln(&flag)
		if flag == "f" {
			break
		} else if flag == "c" {
			var F string
			var stu1 []coder
			n := 0
			fmt.Scanln(&F)
			for _, s := range stu {
				if F == s.name || F == s.student_id {
					stu1 = append(stu1, s)
					n++
				}
			}
			fmt.Printf("找到了%d个人，如下\n", n)
			PrintStu(stu1)
		} else {
			fmt.Println("请重新输入噢！")
			continue
		}
	}
}

// 排序
func SortStu(stu []coder) []coder {
	for i := 0; i < len(stu)-1; i++ {
		max := i
		for j := i; j < len(stu); j++ {
			if stu[max].code_rows < stu[j].code_rows {
				max = j
			}
		}
		if i != max {
			stu[i], stu[max] = stu[max], stu[i]
		}
	}
	return stu
}

// 修改
func ChangeStu(stu []coder) []coder {
	for {
		var flag string
		fmt.Println("如果要修改就输入'q'，不修改就其他，然后输出")
		fmt.Scanln(&flag)
		if flag == "q" {
			var name string
			var key string
			fmt.Println("请输入要修改的人的名字")
			fmt.Scanln(&name)
			fmt.Println("输入要修改的属性有：student_id,phone_id,code_rows,birthday")
			fmt.Scanln(&key)

			for i := 0; i < len(stu); i++ {
				if name == stu[i].name {
					if key == "code_rows" {
						var value int
						fmt.Println("你要修改为什么")
						fmt.Scanln(&value)
						stu[i].code_rows = value
						fmt.Println("修改成功")
					} else {
						var value string
						fmt.Println("你要修改为什么")
						fmt.Scanln(&value)
						switch key {
						case "student_id":
							stu[i].student_id = value
							fmt.Println("修改成功")
						case "name":
							fmt.Println("姓名不准修改")
						case "phone_id":
							stu[i].phone_id = value
							fmt.Println("修改成功")
						case "birthday":
							stu[i].birthday = value
							fmt.Println("修改成功")
						default:
							fmt.Println("输入有误")
						}

					}
				}
			}
		} else {
			break
		}
	}
	fmt.Printf("这是修改之后的名单")
	PrintStu(stu)
	return stu
}
