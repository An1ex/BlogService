package global

import (
	"BlogService/pkg/setting"
	log "github.com/sirupsen/logrus"
)

var (
	Config setting.Config
	Logger *log.Logger
)
