package main

import (
	"github.com/gin-gonic/gin"
	"sql/routes"
	"sql/sql"
)

func main() {
	sql.InitUser()
	defer sql.CloseDB()
	r := gin.Default()
	r.Use(routes.Cors())
	routes.CollectRouter(r)
	panic(r.Run(":9999"))
}
