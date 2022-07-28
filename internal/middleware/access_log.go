package middleware

import (
	"bytes"
	"time"

	"BlogService/global"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()
		global.Logger.WithFields(log.Fields{
			"request":  c.Request.PostForm.Encode(),
			"method":   c.Request.Method,
			"response": bodyWriter.body.String(),
			"status":   bodyWriter.Status(),
			"begin":    beginTime,
			"end":      endTime,
		}).Info("A request response completed")
	}
}
