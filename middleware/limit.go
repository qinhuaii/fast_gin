package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LimitMiddleWare(limit int) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("limit 请求")
		c.Next()
		fmt.Println("limit 响应")
	}
}
