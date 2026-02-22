package expense

import (
	"lovers/internal/usecases/dto/expense"
	expenseCreate "lovers/internal/usecases/expense"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	create *expenseCreate.ExpenseCreate
}

func NewExpenseController(c *expenseCreate.ExpenseCreate) *ExpenseController {
	return &ExpenseController{create: c}
}

func (e *ExpenseController) Create(ctx *gin.Context) {
	var expenseDto expense.ExpenseCreateDto
	if err := ctx.ShouldBind(&expenseDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	err := e.create.Execute(ctx, &expenseDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "group created",
	})
	return

}
