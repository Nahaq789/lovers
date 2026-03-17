package template

import (
	"lovers/internal/usecases/dto/template"
	templateCreate "lovers/internal/usecases/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemplateController struct {
	create *templateCreate.TemplateCreate
}

func NewTemplateController(c *templateCreate.TemplateCreate) *TemplateController {
	return &TemplateController{create: c}
}

func (t *TemplateController) Create(ctx *gin.Context) {
	var templateDto template.TemplateCreateDto
	if err := ctx.ShouldBind(&templateDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err := t.create.Execute(ctx, &templateDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "template created",
	})
	return
}
