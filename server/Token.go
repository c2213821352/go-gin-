package server

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"wxApp/common"
	"wxApp/model"
	"wxApp/response"
)

// 定义一个结构体用于存储用户信息
type UserInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetUserInfo(c *gin.Context) {
	DB := common.GetDB()
	token := c.GetHeader("Authorization")
	claims, err := getUserIDFromToken(token)
	if err != nil {
		fmt.Println("Failed to parse token:", err)
		return
	}

	fmt.Println("ID:", claims["id"])
	var UserInfo model.User
	DB.Where("id = ?", claims["id"]).First(&UserInfo)
	response.Success(c, "成功openid", gin.H{"userInfo": UserInfo.Openid})
	//user, err := findUserByID(claims["id"])
	fmt.Println("Name:", claims["name"])
	fmt.Println("Expiration Time:", claims["exp"])
	//c.JSON(200, user)
}

func getUserIDFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证密钥
		return []byte("20ebdf09fe126ba1c56afdda1f5a3e9b"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return claims, nil
}

// 定义一个生成 Token 的函数

// 定义一个处理函数，用于生成 Token
//func GenerateTokenHandler(c *gin.Context) {
//	// 假设从请求中获取用户信息
//	user := UserInfo{
//		ID:   '123',
//		Name: "Alice",
//	}
//
//	// 生成 Token
//	token, err := GenerateToken(user)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"token": token})
//}
