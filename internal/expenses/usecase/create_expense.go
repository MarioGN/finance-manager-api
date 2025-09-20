package usecase

import (
	"fmt"
	"time"

	"github.com/MarioGN/finance-manager-api/data"
	"github.com/MarioGN/finance-manager-api/internal/expenses/dto"
	"github.com/MarioGN/finance-manager-api/internal/expenses/entity"
)

type CreateExpenseUseCase struct {
	store data.Store
}

func NewCreateExpenseUseCase(store data.Store) *CreateExpenseUseCase {
	return &CreateExpenseUseCase{store: store}
}

func (uc *CreateExpenseUseCase) Execute(input dto.CreateExpenseRequest) (result *dto.ListExpensesResponse, err error) {
	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}

	newExpense, err := entity.NewExpense(int64(input.Amount*100), input.Description, date, entity.ExpenseType(input.ExpenseType))
	if err != nil {
		return nil, fmt.Errorf("failed to create expense entity: %w", err)
	}

	if err := uc.store.Expenses.Save(*newExpense); err != nil {
		return nil, fmt.Errorf("failed to save expense: %w", err)
	}

	return newExpense.ToDTO(), nil
}
