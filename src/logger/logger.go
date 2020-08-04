package logger

import (
	"time"
	"fmt"
	utils "go-webservice/src/utils"
	logrus "github.com/sirupsen/logrus"
)

func InitLogger(filePrefix string) {
	// create file path
	fmt.Println(time.Now())
	dirPath := "./log/"
	fileName := filePrefix + "_" + time.Now().Format("2006-01-02") + ".log"
	logFile, _ := utils.CreateFile(dirPath, fileName)
	logrus.SetOutput(logFile)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetReportCaller(true)
}

func WriteLog(level string, msg string) {
	var Logger = logrus.WithFields(logrus.Fields{

	})
	switch logLevel := level; logLevel {
	case "error":
		Logger.Error(msg)
	case "warn":
		Logger.Warn(msg)
	default:
		Logger.Info(msg)
	}
}