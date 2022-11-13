package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	ginService := gin.Default()
	ginService.Use(favicon.New("lib/ios.ico"))

	ginService.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello,world"})
	})
	ginService.POST("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "hello,post user"})
	})
	ginService.Run(":5555")
}
