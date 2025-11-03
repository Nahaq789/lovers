package authController

import (
	"log/slog"
	"lovers/internal/use_cases/auth"
	"lovers/internal/use_cases/dto/authDto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger *slog.Logger
	signUp *auth.SignUp
}

func NewAuthController(l *slog.Logger, s *auth.SignUp) *AuthController {
	return &AuthController{
		logger: l,
		signUp: s,
	}
}

func (a *AuthController) SignUp(ctx *gin.Context, c *authDto.SignUpDto) {
	var auth authDto.SignUpDto
	if err := ctx.ShouldBind(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err := a.signUp.Execute(ctx, &auth)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user created",
	})
	return
}
