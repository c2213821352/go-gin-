package server

import (
	"crypto/rand"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io"
	"math/big"
	"net/http"
	"time"
	"wxApp/common"
	"wxApp/model"
	"wxApp/response"
)

//	func createUser(c *gin.Context) {
//		var user User
//		if err := c.Bind(&user); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		// 校验通过，处理业务逻辑
//		// ...
//	}
//
// 生成随机ID
func generateRandomID() (string, error) {
	const idLength = 10                                                              // ID的长度
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 字符集

	// 生成随机字符序列
	idBytes := make([]byte, idLength)
	for i := 0; i < idLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		idBytes[i] = charset[randomIndex.Int64()]
	}

	return string(idBytes), nil
}

// 生成随机名称
func generateRandomName() (string, error) {
	const nameLength = 8                                                   // 名称的长度
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 字符集

	// 生成随机字符序列
	nameBytes := make([]byte, nameLength)
	for i := 0; i < nameLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		nameBytes[i] = charset[randomIndex.Int64()]
	}

	return string(nameBytes), nil
}

func GenerateToken(user UserInfo) (string, error) {
	// 创建一个新的 Token
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置 Token 的有效期为 100年
	token.Claims = jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"exp":  time.Now().AddDate(100, 0, 0).Unix(),
	}

	// 使用密钥对 Token 进行签名
	tokenString, err := token.SignedString([]byte("20ebdf09fe126ba1c56afdda1f5a3e9b"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser model.User
	ctx.Bind(&requestUser)
	code := ctx.Query("code")

	// Make the API request
	resp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=wx001db7b4d40e7790&secret=20ebdf09fe126ba1c56afdda1f5a3e9b&js_code=%s&grant_type=authorization_code", code))
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		response.Fail(ctx, "登录错误", nil)
		return
	}
	defer resp.Body.Close()
	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"登录错误2": err.Error()})
		return
	}
	randomID, err := generateRandomID()
	if err != nil {
		fmt.Println("Failed to generate random ID:", err)
		return
	}
	fmt.Println("Random ID:", randomID)

	// 生成随机名称
	randomName, err := generateRandomName()
	if err != nil {
		fmt.Println("Failed to generate random name:", err)
		return
	}

	user := UserInfo{
		ID:   randomID,
		Name: randomName,
	}

	// 生成 Token
	token, err := GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//requestUser.SessionKey = body.session_key
	//requestUser.Openid = body.openid
	print(body)
	type Session struct {
		SessionKey string `json:"session_key"`
		OpenID     string `json:"openid"`
	}
	requestUser.Token = token
	requestUser.Static = true
	requestUser.Id = user.ID
	requestUser.Name = user.Name
	//转换body的数据类型
	var session Session
	if err := json.Unmarshal(body, &session); err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	fmt.Println(session)
	requestUser.SessionKey = session.SessionKey
	requestUser.Openid = session.OpenID
	//requestUser.SessionKey = body
	//body2 := string(body)
	//if err := ctx.ShouldBindJSON(&body2); err != nil {
	//	ctx.JSON(400, gin.H{"error": err.Error()})
	//	return
	//}
	//body3 := body2
	//print(body3)
	// Return the response body 里面是openid和session_key
	if err := DB.Create(&requestUser).Error; err != nil {
		fmt.Println("登录失败", err)
		response.Fail(ctx, "登录失败", nil)
		return
	}
	response.Success(ctx, "登录成功", gin.H{"token": token})

	//ctx.Data(http.StatusOK, "application/json", body)
	print(body)

}
