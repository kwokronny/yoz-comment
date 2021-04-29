package middleware

import (
	"YozComment/util"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var resp = util.Response{}

// AuthCheck 用户校验中间件
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			resp.Error(c, util.ResponseAuthorized, "鉴权失败")
			c.Abort()
			return
		}
		token, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
			return []byte(util.Config.JWTEncrypt), nil
		})
		if err == nil && token.Valid {
			c.Next()
		} else {
			resp.Error(c, util.ResponseAuthorized, "鉴权失败")
			c.Abort()
			return
		}
	}
}

// GenerateToken 生成 JWT-TOKEN
func GenerateToken() (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"username": util.Config.AdminRoot,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	}
	tokenString, err := token.SignedString([]byte(util.Config.JWTEncrypt))
	return tokenString, err
}
