package middleWare

import "github.com/gin-gonic/gin"

// 定义一个中间件函数，用于在 Gin 上下文中存放 openid 和 session_key
func StoreOpenidAndSessionKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 openid 和 session_key
		openid := c.Query("openid")
		sessionKey := c.Query("session_key")

		// 将 openid 和 session_key 存储到 Gin 上下文中
		c.Set("openid", openid)
		c.Set("session_key", sessionKey)

		c.Next()
	}
}
