package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	log := logrus.New()

	// 设置日志级别
	log.SetLevel(logrus.DebugLevel)

	// 设置日志输出格式为 JSON 格式
	log.SetFormatter(&logrus.JSONFormatter{})

	// 设置日志输出位置，如标准输出或文件
	log.SetOutput(os.Stdout)

	Log = log
}
