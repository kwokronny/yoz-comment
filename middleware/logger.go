package middleware

import (
	"YozComment/util"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	f := logFile()
	Logger.SetOutput(f)
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&logFormatter{})
}

type logFormatter struct{}

//格式详情
func (s *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var file string
	var line int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}
	level := strings.ToUpper(entry.Level.String())
	content := entry.Data
	msg := fmt.Sprintf("%s [%s] [%s:%d] %s #content:%v\n", timestamp, level, file, line, entry.Message, content)
	return []byte(msg), nil
}

// LoggerToFile 中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		Logger.Infof("[%s %3d] %s | %13v | %15s |",
			reqMethod,
			statusCode,
			reqURI,
			latencyTime,
			clientIP,
		)
	}
}

// Logger 日志记录到文件
func logFile() io.Writer {
	logFilePath := util.Config.LogFilePath

	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	now := time.Now()
	logFileName := now.Format("2006-01-02") + ".log"

	fileName := path.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
	}
	return src
}
