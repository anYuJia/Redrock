package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(context *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	//context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(context *gin.Context, data gin.H, msg string) {
	//context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"msg":  msg,
	})
}

func Fail(context *gin.Context, data gin.H, msg string) {
	//context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	context.JSON(http.StatusOK, gin.H{
		"code": 400,
		"data": data,
		"msg":  msg,
	})
}
