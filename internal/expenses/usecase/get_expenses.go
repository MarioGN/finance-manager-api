package usecase

import (
	"fmt"

	"github.com/MarioGN/finance-manager-api/data"
	"github.com/MarioGN/finance-manager-api/internal/expenses/dto"
)

type GetExpensesUseCase struct {
	store data.Store
}

func NewGetExpensesUseCase(store data.Store) *GetExpensesUseCase {
	return &GetExpensesUseCase{
		store: store,
	}
}

func (uc *GetExpensesUseCase) Execute() (result []dto.ExpenseDTO, err error) {
	expenses, err := uc.store.Expenses.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to list expenses: %w", err)
	}

	if len(expenses) == 0 {
		return []dto.ExpenseDTO{}, nil
	}

	for _, e := range expenses {
		dto := e.ToDTO()
		result = append(result, *dto)
	}

	return result, nil
}
