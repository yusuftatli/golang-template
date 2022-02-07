package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yusuftatli/go-template-with-redis-gorm-api/pkg/logging"
	"go.uber.org/zap"
)

func HttpLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// before request

		c.Next()

		// after request

		logger := logging.GetLogger()
		logger.ZapLogger.With(
			zap.Int("status", c.Writer.Status()),
			zap.Int64("latency_ns", time.Since(start).Nanoseconds()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.String("url", c.Request.URL.Path),
			zap.String("host", c.Request.Host),
			zap.String("correlationId", logger.CorrelationId.String()),
		).Info("request details")
	}
}
