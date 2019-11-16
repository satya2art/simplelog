package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Log(level string, msg ...interface{}) {

	file, err := os.OpenFile("/var/log/simplelogger.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr", err)
	}

	if level == "Info" {
		log.Info(msg)
	} else if level == "Debug" {
		log.Debug(msg)
	} else if level == "Error" {
		log.Error(msg)
	} else {
		log.Fatal("Non existent logging level")
	}
}
