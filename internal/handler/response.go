package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一成功处理
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// 统一错误处理
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"code":  code,
		"error": msg,
	})
}
