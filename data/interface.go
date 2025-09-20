package data

import (
	"github.com/MarioGN/finance-manager-api/core/expenses"
)

type ExpenseRepository interface {
	FindAll() ([]expenses.Expense, error)
	Save(expense expenses.Expense) error
	FindByID(id string) (*expenses.Expense, error)
	Update(expense expenses.Expense) error
	Delete(id string) error
}
