package group

import (
	"lovers/internal/usecases/dto/group"
	groupCreate "lovers/internal/usecases/group"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	create *groupCreate.GroupCreate
}

func NewGroupController(c *groupCreate.GroupCreate) *GroupController {
	return &GroupController{create: c}
}

func (g *GroupController) Create(ctx *gin.Context) {
	var groupDto group.GroupCreateDto
	if err := ctx.ShouldBind(&groupDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err := g.create.Execute(ctx, &groupDto)
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
