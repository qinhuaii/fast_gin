package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(c *gin.Context) {
	fmt.Println("auth 请求")
	c.Next()
	fmt.Println("auth 响应")
}
