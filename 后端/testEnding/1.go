package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	Sex      uint8     `json:"sex"`
	Birthday time.Time `json:"birthday"`
}

// 要想使用json.Marshal这个来把结构体解析成json格式，首先结构体里面变量名必须要首字母大写，而且要在后面接上相应的 `json:****`
func main() {
	u := user{
		Username: "坤坤",
		Nickname: "阿坤",
		Sex:      20,
		Birthday: time.Now(),
	}
	bs, err := json.Marshal(&u)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(bs))
}
