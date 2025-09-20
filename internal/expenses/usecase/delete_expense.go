package usecase

import (
	"fmt"

	"github.com/MarioGN/finance-manager-api/data"
)

type DeleteExpenseUseCase struct {
	store data.Store
}

func NewDeleteExpenseUseCase(store data.Store) *DeleteExpenseUseCase {
	return &DeleteExpenseUseCase{store: store}
}

func (uc *DeleteExpenseUseCase) Execute(id string) error {
	dbExpense, err := uc.store.Expenses.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to find expense by ID: %w", err)
	}

	if dbExpense == nil {
		return fmt.Errorf("expense not found")
	}

	if err := uc.store.Expenses.Delete(id); err != nil {
		return fmt.Errorf("failed to delete expense: %w", err)
	}

	return nil
}
