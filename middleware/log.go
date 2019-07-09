package middleware

import (
	"bytes"
	"net/url"
	"time"

	"api.service.com/lib"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// LogMiddleware 记录log
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				lib.GetLogger().Error(e)
			}
		} else {

			pGet := c.Request.URL.Query()
			var pPost url.Values
			if c.Request.Method == "POST" {
				c.Request.ParseForm()
				pPost = c.Request.PostForm
			}

			lib.GetLogger().Info(path,
				zap.String("method", c.Request.Method),
				zap.Int("status", c.Writer.Status()),
				zap.Any("get", pGet),
				zap.Any("post", pPost),
				//zap.ByteString("response", blw.body.Bytes()),
				zap.String("ip", c.ClientIP()),
				zap.Duration("latency", latency),
			)
		}
	}
}
