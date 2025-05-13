package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// InitLogger initializes the logger
func InitLogger(level string) {
	log = logrus.New()

	// 设置输出
	log.Out = os.Stdout

	// 设置格式
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05Z07:00",
	}

	// 设置日志级别
	lvl, err := logrus.ParseLevel(level)
	if err != nil {

		log.Warnf("Invalid log level, using default level: %s", logrus.InfoLevel.String())
		lvl = logrus.InfoLevel
	}
	log.Level = lvl
	log.Info("Logger initialized")
}

// 获取日志记录器
func GetLogger() *logrus.Logger {
	if log == nil {
		InitLogger("info")
	}
	return log
}
