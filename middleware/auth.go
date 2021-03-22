package middleware

import (
	"YozComment/util"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.HTML(http.StatusOK, "login.html", gin.H{})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GenerateToken() (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims = jwt.MapClaims{
		"username": util.Config.AdminRoot,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(util.Config.JWTEncrypt)
	return tokenString, err
}
