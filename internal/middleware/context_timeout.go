package middleware

import (
	"context"
	"time"

	"BlogService/global"
	"github.com/gin-gonic/gin"
)

func ContextTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		//设置context的超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), global.Config.Server.ContextTimeout*time.Second)
		defer cancel()
		//重新赋到请求中
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
