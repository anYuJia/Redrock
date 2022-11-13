package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/cookie", func(c *gin.Context) {

		cookie, _ := c.Cookie("gin_cookie")
		if cookie == "ok" {
			c.String(http.StatusOK, "你的cookie为%s", cookie)
		} else {
			c.String(http.StatusOK, "你还没有登陆啊")
		}
	})

	router.Run()
}
