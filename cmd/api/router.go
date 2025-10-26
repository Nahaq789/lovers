package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Router(r gin.IRouter) {
	v1 := r.Group("api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		slog.InfoContext(c.Request.Context(), "良い子のみんな〜")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
