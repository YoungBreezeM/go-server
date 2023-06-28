package log

import (
	"os"
	"server/config"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	log := logrus.New()
	log.SetLevel(logrus.Level(config.Cfg.Log.Level))
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	Log = log
}
