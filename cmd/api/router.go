package main

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(r gin.IRouter) {
	v1 := r.Group("api/v1")
	v1.GET("/ping", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		slog.InfoContext(c.Request.Context(), "良い子のみんな〜")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
