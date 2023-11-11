package middleWare

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			context.JSON(401, gin.H{"code": 401, "msg": "权限不足"})
			context.Abort()
			return
		}
	}
}
