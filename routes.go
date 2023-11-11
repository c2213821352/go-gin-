package main

import (
	"github.com/gin-gonic/gin"
	"wxApp/middleWare"
	"wxApp/server"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleWare.CorsMiddleWare(), middleWare.RecoveryMiddleware(), middleWare.StoreOpenidAndSessionKey())
	//r.POST("/api/auth/register", server.Register)
	r.GET("/login", server.Login)
	r.POST("/getOpenid", server.GetUserInfo)
	//
	//courseRouter 	CourseSelect

	return r
}
