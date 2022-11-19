package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type SQ3 struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	service := gin.Default()
	service.GET("/register", func(context *gin.Context) {
		var user SQ3
		user.Username = context.Query("username")
		user.Password = context.Query("password")
		user.Question = context.Query("question")
		user.Answer = context.Query("answer")

		var dns = "root:z0114b3586@tcp(127.0.0.1:3306)/mydata"
		db, err := sql.Open("mysql", dns)
		err = db.Ping()
		if err != nil {
			fmt.Println("==================>", err)
		} else {
			fmt.Println("==================>yes")
		}
		var username string
		db.QueryRow("select * from user where username = ?", user.Username).Scan(&username)
		if username != "" {
			context.String(http.StatusOK, "该用户名已经被用过了%s%s", user.Username)
			return
		}
		db.Exec("insert into user(username,password,question,answer) values(?,?,?,?)", user.Username, user.Password, user.Question, user.Answer)
		context.String(http.StatusOK, "注册成功")
		defer db.Close()
	})
	service.GET("/login", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		var dns = "root:z0114b3586@tcp(127.0.0.1:3306)/mydata"
		db, _ := sql.Open("mysql", dns)
		defer db.Close()
		var user SQ3
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
	service.GET("/renew", func(context *gin.Context) {
		username := context.Query("username")
		answer := context.Query("answer")
		newPassword := context.Query("newpassword")
		var dns = "root:z0114b3586@tcp(127.0.0.1:3306)/mydata"
		db, _ := sql.Open("mysql", dns)
		defer db.Close()
		var user SQ3
		db.QueryRow("select * from user where username = ?", username).Scan(&user.Username, &user.Password, &user.Question, &user.Answer)
		if user.Username == username && user.Answer == answer {
			db.Exec("update user set password =? where username=?", newPassword, username)
			context.String(http.StatusOK, "密码修改修改成功")
		} else {
			context.String(http.StatusOK, "密码修改失败")
		}
	})
	service.Run(":2222")
}
