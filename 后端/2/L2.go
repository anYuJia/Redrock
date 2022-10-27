package main

import "fmt"

type Animal interface {
	say()
}

// 单身狗
type SingleDog struct {
}

// 狗
type Dog struct {
}

func (singledog SingleDog) say() {
	fmt.Println("呜呜呜~~~")
}
func (dog Dog) say() {
	fmt.Println("汪汪汪~~~")
}
func main() {
	dog := Animal(&Dog{})
	dog.say()
	singledog := SingleDog{}
	singledog.say()
}
