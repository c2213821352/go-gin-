package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"os"
	"wxApp/common"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	r := gin.Default()
	//用户使用get请求访问
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	//r.GET("/hello", sayHello)

	if port != "" {
		panic(r.Run(":" + port)) //启动失败，则抛出一个运行时异常。
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

//func sayHello(c *gin.Context) {
//	//c.HTML(200, "hello")
//	b,_ := c.GetRawData()  //获取数据
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//
//	//接受就可以校验
//	var m = map[string]interface{}
//	_ =  json.Unmarshal(b,&m)
//	c.JSON(http.StatusOK, m)
//	c.JSON(200, gin.H{
//		"message": "hello",
//		"age":     18,
//		"username": username,
//		"pwd":password,
//
//	})
//}
//
//func notNound(ctx *gin.Context)  {
//	ctx.HTML(http.StatusNotFound,"404.html")
//}
