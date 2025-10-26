package middleware

import (
	"context"
	"lovers/internal/shared/infrastructure/tracing/key"

	"github.com/gin-gonic/gin"
)

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		contextTrace := key.NewContextTrace()

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, key.ContextTraceKey, contextTrace.GenerateID())

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
