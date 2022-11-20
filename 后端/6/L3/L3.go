package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Count struct {
	Username string `json:"username"`
	Context  string `json:"context"`
}

func main() {
	service := gin.Default()
	service.LoadHTMLGlob("yuxi/6/L3/html/*")
	service.Static("yuxi/6/L3/img/", "./yuxi/6/l3/img/")
	service.GET("", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})
	service.POST("/login", func(context *gin.Context) {
		//username := context.PostForm("Username")
		var dns = "root:z0114b3586@tcp(127.0.0.1:3306)/mydata"
		db, _ := sql.Open("mysql", dns)
		rows, _ := db.Query("select * from liuyan")
		var count Count
		var Co []Count
		for rows.Next() {
			rows.Scan(&count.Username, &count.Context)
			Co = append(Co, count)
		}
		h, _ := json.Marshal(Co)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"c": h,
		})
	})
	service.POST("/liuyan", func(context *gin.Context) {
		user := context.PostForm("Username")
		con := context.PostForm("Context")
		var dns = "root:z0114b3586@tcp(127.0.0.1:3306)/mydata"
		db, _ := sql.Open("mysql", dns)
		defer db.Close()
		db.Exec("insert into liuyan values(?,?)", user, con)
		rows, _ := db.Query("select * from liuyan")
		var count Count
		var Co []Count
		for rows.Next() {
			rows.Scan(&count.Username, &count.Context)
			Co = append(Co, count)
		}
		h, _ := json.Marshal(Co)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"c": string(h),
		})
	})
	service.Run(":7777")
}
