package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SQ struct {
	username string
	password string
}

func main() {
	service := gin.Default()
	var users []SQ
	service.GET("/register", func(context *gin.Context) {
		var user SQ
		user.username = context.Query("username")
		user.password = context.Query("password")
		users = append(users, user)
		context.String(http.StatusOK, "注册成功")
	})
	service.GET("/login", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		ok := false
		for _, user := range users {
			if username == user.username {
				ok = true
				if password == user.password {
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
