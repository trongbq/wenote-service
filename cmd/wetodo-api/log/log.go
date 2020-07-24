package log

import (
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var FileLogName = "./log/wenote-api.log"

// InitLogrus sets up Logrus
func InitLogrus() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

	env := strings.ToLower(os.Getenv("ENV"))
	if env == "local" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if env == "test" {
		logrus.SetOutput(os.Stdout)
	} else {
		logFile := &lumberjack.Logger{
			Filename:   FileLogName,
			MaxSize:    200, //MB
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   false,
		}
		mw := io.MultiWriter(os.Stdout, logFile)
		logrus.SetOutput(mw)
	}

	logrus.Info("Logrus initialized")
}
