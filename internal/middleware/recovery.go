package middleware

import (
	"fmt"
	"time"

	"BlogService/global"
	"BlogService/pkg/app"
	"BlogService/pkg/email"
	"BlogService/pkg/errcode"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.Config.Email.Host,
		Port:     global.Config.Email.Port,
		IsSSL:    global.Config.Email.IsSSL,
		UserName: global.Config.Email.UserName,
		Password: global.Config.Email.Password,
		From:     global.Config.Email.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithFields(log.Fields{
					"error": err,
				}).Error("Panic recover")

				err = defaultMailer.SendMail(
					global.Config.Email.To,
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)
				if err != nil {
					app.NewResponse(c).ToErrorResponse(errcode.ServerError)
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}
