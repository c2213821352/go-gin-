package server

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func Money(c *gin.Context) {
	// 构造请求参数
	params := map[string]interface{}{
		"nonce_str":    "value1", //随机字符串
		"sign":         "value2", //签名
		"mch_billno":   GenerateOrderNumberHandler(),
		"mch_id":       "", //注册过的商户号
		"wxappid":      "value2",
		"send_name":    "聪明的懒羊羊",
		"re_openid":    "value1", //用户openid
		"total_amount": "1",
		"total_num":    "1",
		"wishing":      "祝您发大财",
		"act_name":     "聪明的懒羊羊",
		"remark":       "视频红包",               //备注
		"notify_way":   "MINI_PROGRAM_JSAPI", //默认格式
		//"scene_id":     "PRODUCT_5",
	}

	// 将请求参数转换为JSON格式
	jsonData, err := json.Marshal(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 创建HTTP请求客户端
	client := &http.Client{}

	// 创建POST请求
	req, err := http.NewRequest("POST", "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendminiprogramhb", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// 读取响应
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回响应结果
	c.Data(http.StatusOK, "application/json", respData)
}

// 生成订单号
func GenerateOrderNumber() string {
	uuid := uuid.New().String()
	return uuid
}

// 生成订单号
func GenerateOrderNumberHandler() string {
	orderNumber := GenerateOrderNumber()
	return orderNumber
	//c.JSON(http.StatusOK, gin.H{"order_number": orderNumber})
}
