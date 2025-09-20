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

func (uc *GetExpensesUseCase) Execute() (list []dto.ListExpensesResponse, err error) {
	expenses, err := uc.store.Expenses.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to list expenses: %w", err)
	}

	for _, e := range expenses {
		result := e.ToDTO()
		list = append(list, *result)
	}

	return list, nil
}
