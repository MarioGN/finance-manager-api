package usecase

import (
	"fmt"

	"github.com/MarioGN/finance-manager-api/data"
	"github.com/MarioGN/finance-manager-api/internal/expenses/dto"
)

type GetExpenseUseCase struct {
	store data.Store
}

func NewGetExpenseUseCase(store data.Store) *GetExpenseUseCase {
	return &GetExpenseUseCase{store: store}
}

func (uc *GetExpenseUseCase) Execute(id string) (result *dto.ListExpensesResponse, err error) {
	expense, err := uc.store.Expenses.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find expense by ID: %w", err)
	}

	if expense == nil {
		return nil, fmt.Errorf("expense not found")
	}

	return expense.ToDTO(), nil
}
