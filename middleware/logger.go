package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

func Logger(log *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		now := time.Now()
		TraceID := uuid.New().String()

		if c.Request.URL.RawQuery != "" {
			path = fmt.Sprintf("%s?%s", path, c.Request.URL.RawQuery)
		}

		log.Infow("request started", "trace_id", TraceID, "method", c.Request.Method, "path", path,
			"remoteaddr", c.Request.RemoteAddr)

		// 处理请求
		c.Next()

		if len(c.Errors.Errors()) > 0 {
			log.Errorf("request error %s", c.Errors.Errors())
		}

		log.Infow("request completed", "trace_id", TraceID, "method", c.Request.Method, "path", path,
			"remoteaddr", c.Request.RemoteAddr, "since", time.Since(now))
	}
}
