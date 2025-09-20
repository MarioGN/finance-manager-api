package controller

import (
	"github.com/MarioGN/finance-manager-api/data"
	"github.com/MarioGN/finance-manager-api/internal/expenses/dto"
	"github.com/MarioGN/finance-manager-api/internal/expenses/usecase"
	"github.com/labstack/echo/v4"
)

type expenseController struct {
	store *data.Store
}

func ConfigureExpenseRoutes(group *echo.Group, store *data.Store) {
	ctrl := &expenseController{store: store}

	group.GET("", ctrl.handleGetExpenses)
	group.POST("", ctrl.handleCreateExpense)
	group.GET("/:id", ctrl.handleGetExpenseByID)
	group.PUT("/:id", ctrl.handleUpdateExpense)
	group.DELETE("/:id", ctrl.handleDeleteExpense)
}

func (ctrl *expenseController) handleGetExpenses(c echo.Context) error {
	uc := usecase.NewGetExpensesUseCase(*ctrl.store)

	res, err := uc.Execute()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to list expenses"})
	}

	return c.JSON(200, res)
}

func (ctrl *expenseController) handleCreateExpense(c echo.Context) error {
	var req dto.ExpenseDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request payload"})
	}

	uc := usecase.NewCreateExpenseUseCase(*ctrl.store)

	res, err := uc.Execute(req)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to create expense"})
	}

	return c.JSON(201, res)
}

func (ctrl *expenseController) handleGetExpenseByID(c echo.Context) error {
	id := c.Param("id")
	uc := usecase.NewGetExpenseUseCase(*ctrl.store)

	res, err := uc.Execute(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to get expense"})
	}

	if res == nil {
		return c.JSON(404, map[string]string{"error": "expense not found"})
	}

	return c.JSON(200, res)
}

func (ctrl *expenseController) handleUpdateExpense(c echo.Context) error {
	var req dto.ExpenseDTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request payload"})
	}

	id := c.Param("id")

	uc := usecase.NewUpdateExpenseUseCase(*ctrl.store)

	res, err := uc.Execute(id, req)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to update expense"})
	}

	return c.JSON(200, res)
}

func (ctrl *expenseController) handleDeleteExpense(c echo.Context) error {
	id := c.Param("id")

	uc := usecase.NewDeleteExpenseUseCase(*ctrl.store)

	err := uc.Execute(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to delete expense"})
	}

	return c.NoContent(204)
}
