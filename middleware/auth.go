package middleware

import (
	"encoding/json"
	"kwok-comment/helper"
	"time"

	"github.com/gin-gonic/gin"
)

// Auth 鉴权结构体
type Auth struct {
	OpenID    string `json:"openId"`
	AppName   string `json:"appName"`
	Timestarp int64  `json:"timestarp"`
	Role      int    `json:"role"`
}

const cookieName string = "ZHsPpgVaJVKorA2jmNsLP6W"
const aesKey string = "iok;9ELPLJixsHDzRYhb4LYc"

// CheckAuth 是验证登陆中间件
func CheckAuth(onlyDecrypt bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		encryptText := c.GetHeader(cookieName)
		if encryptText == "" && onlyDecrypt == true {
			helper.Response{}.Error(c, helper.ResponseAuthorized, "鉴权失败")
			c.Abort()
		}
		var auth Auth
		err := json.Unmarshal([]byte(helper.AesDecrypt(encryptText, aesKey)), &auth)
		if err == nil && auth.Timestarp+7200 >= time.Now().Unix() {
			c.Set("auth", auth)
		} else if onlyDecrypt == true {
			helper.Response{}.Error(c, helper.ResponseAuthorized, "鉴权失败")
			c.Abort()
		}
	}
}

// SetAuth 设置鉴权信息
func SetAuth(c *gin.Context, auth Auth) (encryptResult string, err error) {
	encryptText, err := json.Marshal(auth)
	if err == nil {
		encryptResult = helper.AesEncrypt(string(encryptText), aesKey)
	}
	return
}
