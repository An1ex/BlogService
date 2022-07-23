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
	l.Out = logFile
	l.Formatter = &log.JSONFormatter{}
	l.Level = log.DebugLevel
	l.ReportCaller = true
	return l, err
}