package middleware

import (
	"KBCommentAPI/helper"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		helper.Logger().Infof("[%s] %s | %3d | %13v | %15s |",
			reqMethod,
			reqURI,
			statusCode,
			latencyTime,
			clientIP,
		)
	}
}
