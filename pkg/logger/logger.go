package logger

import (
	"os"

	"BlogService/global"
	log "github.com/sirupsen/logrus"
)

func NewLogger() (*log.Logger, error) {
	l := log.New()
	logFilePath := global.Config.App.LogSavePath + "/" + global.Config.App.LogFileName + global.Config.App.LogFileExt
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	switch global.Config.Server.RunMode {
	case "debug":
		l.Level = log.DebugLevel
		l.Out = os.Stdout
		l.Formatter = &log.TextFormatter{ForceColors: true}
	case "release":
		l.Level = log.ErrorLevel
		l.Out = logFile
		l.Formatter = &log.JSONFormatter{}
	case "test":
		l.Level = log.WarnLevel
		l.Out = logFile
		l.Formatter = &log.JSONFormatter{}
	default:
	}
	return l, err
}

func InitLogger() error {
	var err error
	global.Logger, err = NewLogger()
	if err != nil {
		return err
	}
	return nil
}
