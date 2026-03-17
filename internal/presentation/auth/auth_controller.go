package auth

import (
	auth_signup "lovers/internal/usecases/auth"
	"lovers/internal/usecases/dto/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	signUp *auth_signup.SignUp
}

func NewAuthController(s *auth_signup.SignUp) *AuthController {
	return &AuthController{
		signUp: s,
	}
}

func (a *AuthController) SignUp(ctx *gin.Context) {
	var auth auth.SignUpDto
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
