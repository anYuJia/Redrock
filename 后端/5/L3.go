package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type SQ1 struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	service := gin.Default()
	service.GET("/register", func(context *gin.Context) {
		var user SQ1
		user.Username = context.Query("username")
		user.Password = context.Query("password")
		context.String(http.StatusOK, "注册成功")
		filePtr, _ := os.OpenFile("usersdata.json", os.O_CREATE|os.O_RDWR, 0666)
		userMap := new([]SQ1)
		decoder := json.NewDecoder(filePtr)
		_ = decoder.Decode(userMap)
		filePtr.Close()
		filePtr, _ = os.OpenFile("usersdata.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
		*userMap = append(*userMap, user)
		encoder := json.NewEncoder(filePtr)
		_ = encoder.Encode(*userMap)
		filePtr.Close()
	})
	service.GET("/login", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		filePtr, _ := os.OpenFile("usersdata.json", os.O_CREATE|os.O_RDWR, 0666)
		userMap := new([]SQ1)
		decoder := json.NewDecoder(filePtr)
		_ = decoder.Decode(userMap)
		ok := false
		for _, user := range *userMap {
			if username == user.Username {
				ok = true
				if password == user.Password {
					context.String(http.StatusOK, "登陆成功")
					break
				} else {
					context.String(http.StatusOK, "登陆失败，密码错误")
					break
				}
			}
		}
		if !ok {
			context.String(http.StatusOK, "没有该用户")
		}
	})
	service.Run(":2222")
}
