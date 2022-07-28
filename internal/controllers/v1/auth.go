package v1

import (
	"BlogService/global"
	"BlogService/internal/service"
	"BlogService/pkg/app"
	"BlogService/pkg/errcode"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, verrs := app.BindAndValid(c, &param)
	if !valid { //入参校验或绑定参数失败
		global.Logger.WithFields(log.Fields{
			"error": verrs.Error(),
		}).Error("BindAndValid failed!")
		errRsp := errcode.InvalidParams.WithDetails(verrs.Errors()...)
		response.ToErrorResponse(errRsp)
	} else { //参数校验或绑定参数成功
		svc := service.New(c)
		_, err := svc.GetAuth(&param)
		if err != nil { //auth不存在
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Get auth failed!")
			response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		}

		token, err := app.GenerateToken(param.AppKey, param.AppSecret)
		if err != nil {
			global.Logger.WithFields(log.Fields{
				"error": err.Error(),
			}).Error("Generate JWT token failed!")
			response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		}

		response.ToResponse(gin.H{
			"token": token,
		})
	}
}
