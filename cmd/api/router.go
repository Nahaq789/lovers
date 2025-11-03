package main

import (
	"log/slog"
	authDi "lovers/cmd/di/auth"
	"lovers/internal/presentations/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r gin.IRouter, authSet authDi.AuthSet) {
	v1 := r.Group("api/v1")
	v1.Use(middleware.TraceMiddleware())
	v1.Use(middleware.LoggingMiddleware())
	{
		v1.POST("/auth/signup", authSet.AuthController.SignUp)
	}

	v1.GET("/ping", func(c *gin.Context) {
		slog.InfoContext(c.Request.Context(), "良い子のみんな〜")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
