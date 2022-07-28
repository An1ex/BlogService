package middleware

import (
	"BlogService/pkg/app"
	"BlogService/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			code  = errcode.Success
		)

		//从请求中获取token
		if v, ok := c.GetQuery("token"); ok {
			token = v
		} else {
			token = c.GetHeader("token")
		}

		//解析并校验token
		if token == "" { //token为空
			code = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil { //解析校验token失败
				code = err.(*errcode.Error)
			}
		}

		//响应失败
		if code != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(code)
			c.Abort()
			return
		}
		//响应成功
		c.Next()
	}
}
