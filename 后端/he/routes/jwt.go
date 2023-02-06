package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sql/response"
)

func HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, e := c.Request.Cookie("user_cookie")
		if e == nil {
			c.Next()
		} else {
			response.Response(c, 502, 502, gin.H{}, "未登录不能执行此操作")
			c.Abort()
			//c.HTML(http.StatusUnauthorized, "401.tmpl", nil)
		}
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		log.Println("===================>", origin)
		c.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,mode,credentials")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
