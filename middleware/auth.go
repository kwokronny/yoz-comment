package middleware

import (
	"YozComment/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// LoggerToFile 日志记录到文件
func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func CreateToken() {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims["admin"] = Config.AdminRoot
	token.Claims["pass"] = Config.AdminPass
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString()
}
