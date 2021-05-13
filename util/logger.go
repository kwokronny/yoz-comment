package util

import (
	"fmt"
	"io"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func init() {
	f := logFile()
	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&logFormatter{})
}

type logFormatter struct{}

// Format 设置日志格式
func (s *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())
	msg := fmt.Sprintf("%s [%s] %s \n", timestamp, level, entry.Message)
	return []byte(msg), nil
}

// Logger 日志记录到文件
func logFile() io.Writer {
	logFilePath := "./logs/app"

	writer, err := rotatelogs.New(
		logFilePath+"-%Y%m%d%H.log",
		rotatelogs.WithLinkName(logFilePath+".log"),
		rotatelogs.WithRotationTime(24*time.Hour),
		rotatelogs.WithRotationCount(30),
	)

	if err != nil {
		log.Errorf("日志模块配置本地文件错误: %v", err)
	}

	return writer
}
