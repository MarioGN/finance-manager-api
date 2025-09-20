package usecase

import (
	"fmt"
	"time"

	"github.com/MarioGN/finance-manager-api/data"
	"github.com/MarioGN/finance-manager-api/internal/expenses/dto"
	"github.com/MarioGN/finance-manager-api/internal/expenses/entity"
)

type UpdateExpenseUseCase struct {
	store data.Store
}

func NewUpdateExpenseUseCase(store data.Store) *UpdateExpenseUseCase {
	return &UpdateExpenseUseCase{store: store}
}

func (uc *UpdateExpenseUseCase) Execute(id string, input dto.ExpenseDTO) (result *dto.ExpenseDTO, err error) {
	dbExpense, err := uc.store.Expenses.FindByID(id)
	if err != nil {
		return nil, err
	}

	if dbExpense == nil {
		return nil, fmt.Errorf("expense not found")
	}

	err = dbExpense.SetAmount(int64(input.Amount * 100))
	if err != nil {
		return nil, fmt.Errorf("invalid amount: %w", err)
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}
	dbExpense.SetDate(date)

	err = dbExpense.SetExpenseType(entity.ExpenseType(input.ExpenseType))
	if err != nil {
		return nil, fmt.Errorf("invalid expense type: %w", err)
	}

	dbExpense.SetDescription(input.Description)

	if err := uc.store.Expenses.Update(*dbExpense); err != nil {
		return nil, fmt.Errorf("failed to save expense: %w", err)
	}

	return dbExpense.ToDTO(), nil
}
