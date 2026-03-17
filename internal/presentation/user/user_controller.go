package user

import (
	userDto "lovers/internal/usecases/dto/user"
	userRegistration "lovers/internal/usecases/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	registration *userRegistration.UserRegistration
}

func NewUserController(r *userRegistration.UserRegistration) *UserController {
	return &UserController{registration: r}
}

func (u *UserController) Registration(ctx *gin.Context) {
	var user userDto.UserRegistrationDto
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err := u.registration.Execute(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
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
