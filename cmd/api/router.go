package main

import (
	"log/slog"
	authDi "lovers/cmd/di/auth"
	"lovers/cmd/di/expense"
	"lovers/cmd/di/group"
	"lovers/cmd/di/template"
	"lovers/cmd/di/user"
	"lovers/internal/presentation/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r gin.IRouter, authSet authDi.AuthSet) {
	v1 := r.Group("api/v1")
	v1.Use(middleware.TraceMiddleware())
	v1.Use(middleware.LoggingMiddleware())
	{
		v1.POST("/auth/signup", authSet.AuthController.SignUp)
	}

	r.GET("/ping", func(c *gin.Context) {
		slog.InfoContext(c.Request.Context(), "良い子のみんな〜")
		time.Sleep(10 * time.Second)
		slog.InfoContext(c.Request.Context(), "hogehoge")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func UserRouter(r gin.IRouter, userSet user.UserSet) {
	v1 := r.Group("api/v1")
	v1.Use(middleware.TraceMiddleware())
	v1.Use(middleware.LoggingMiddleware())

	{
		v1.POST("/user/register", userSet.UserController.Registration)
	}
}

func GroupRouter(r gin.IRouter, groupSet group.GroupSet) {
	v1 := r.Group("api/v1")
	v1.Use(middleware.TraceMiddleware())
	v1.Use(middleware.LoggingMiddleware())
	v1.Use(middleware.AuthMiddleware())

	{
		v1.POST("/group/create", groupSet.GroupController.Create)
	}
}

func TemplateRouter(r gin.IRouter, templateSet template.TemplateSet) {
	v1 := r.Group("api/v1")
	v1.Use(middleware.TraceMiddleware())
	v1.Use(middleware.LoggingMiddleware())
	v1.Use(middleware.AuthMiddleware())

	{
		v1.POST("/template/create", templateSet.TemplateController.Create)
	}
}

func ExpenseRouter(r gin.IRouter, expenseSet expense.ExpenseSet) {
	v1 := r.Group("api/v1")
	v1.Use(middleware.TraceMiddleware())
	v1.Use(middleware.LoggingMiddleware())
	v1.Use(middleware.AuthMiddleware())

	{
		v1.POST("/expense/add", expenseSet.ExpenseController.Add)
	}
}
