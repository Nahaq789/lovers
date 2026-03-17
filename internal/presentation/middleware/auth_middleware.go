package middleware

import (
	"lovers/internal/shared/infrastructure/security/userid"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := userid.WithContext(c.Request.Context(), "123e4567-e89b-12d3-a456-426614174000")
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
