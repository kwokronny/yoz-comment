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

//格式详情
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
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logFilePath+".log"),

		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		rotatelogs.WithRotationTime(24*time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		rotatelogs.WithRotationCount(30),
	)

	if err != nil {
		log.Errorf("日志模块配置本地文件错误: %v", err)
	}

	return writer
	// if err := os.MkdirAll(logFilePath, 0777); err != nil {
	// 	fmt.Println(err.Error())
	// }
	// now := time.Now()
	// logFileName := now.Format("2006-01-02") + ".log"

	// fileName := path.Join(logFilePath, logFileName)
	// src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
}
