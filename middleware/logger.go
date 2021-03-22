package middleware

import (
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// LoggerToFile 中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		log.Infof("[%s %3d] %s %s [%v | %s]",
			reqMethod,
			statusCode,
			reqURI,
			reqBody,
			latencyTime,
			clientIP,
		)
	}
}
