package expense

import (
	"lovers/internal/usecases/dto/expense"
	expenseCreate "lovers/internal/usecases/expense"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	create *expenseCreate.ExpenseAdd
}

func NewExpenseController(c *expenseCreate.ExpenseAdd) *ExpenseController {
	return &ExpenseController{create: c}
}

func (e *ExpenseController) Add(ctx *gin.Context) {
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
		"message": "expense created",
	})
	return

}
