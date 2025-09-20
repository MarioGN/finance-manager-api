package data

import "github.com/MarioGN/finance-manager-api/internal/expenses/entity"

type ExpenseRepository interface {
	FindAll() ([]entity.Expense, error)
	Save(expense entity.Expense) error
	FindByID(id string) (*entity.Expense, error)
	Update(expense entity.Expense) error
	Delete(id string) error
}
