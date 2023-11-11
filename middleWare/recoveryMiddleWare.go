package middleWare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"wxApp/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil { //如果捕获到了panic，则使用response.Fail函数返回一个错误响应。
				response.Fail(c, fmt.Sprint(err), nil)
				return
			}
		}()
		c.Next()
	}
}
