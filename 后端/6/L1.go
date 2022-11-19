package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type SQ2 struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	service := gin.Default()
	service.GET("/register", func(context *gin.Context) {
		var user SQ2
		user.Username = context.Query("username")
		user.Password = context.Query("password")
		context.String(http.StatusOK, "注册成功")
		var dns = "root:z0114b3586@tcp(127.0.0.1:3306)/mydata"
		db, err := sql.Open("mysql", dns)
		err = db.Ping()
		if err != nil {
			fmt.Println("==================>", err)
		} else {
			fmt.Println("==================>yes")
		}
		db.Exec("insert into user(username,password) values(?,?)", user.Username, user.Password)
		row, _ := db.Query("select * from user")
		for row.Next() {
			var name string
			var password string
			row.Scan(&name, &password)
			fmt.Println(name, password)
		}
		defer db.Close()
	})
	service.GET("/login", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		var dns = "root:z0114b3586@tcp(127.0.0.1:3306)/mydata"
		db, _ := sql.Open("mysql", dns)
		var user SQ2
		db.QueryRow("select * from user where username = ?", username).Scan(&user.Username, &user.Password)
		if user.Username == "" {
			context.String(http.StatusOK, "没有该用户")
			return
		}
		if user.Password == password {
			context.String(http.StatusOK, "登陆成功")
			return
		} else {
			context.String(http.StatusOK, "密码错误")
			return
		}

	})
	service.Run(":2222")
}
