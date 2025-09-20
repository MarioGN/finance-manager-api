package data

import (
	"time"

	"github.com/MarioGN/finance-manager-api/core/expenses"
)

type ExpenseInMemoryRepository struct {
	expenses []expenses.Expense
}

func NewExpenseInMemoryRepository() *ExpenseInMemoryRepository {
	r := &ExpenseInMemoryRepository{}
	r._seedData()
	return r
}

func (r *ExpenseInMemoryRepository) _seedData() {
	ex, _ := expenses.NewExpense(10000, "Electricity bill", time.Now(), expenses.VariableExpense)
	r.expenses = append(r.expenses, *ex)

	ex, _ = expenses.NewExpense(50000, "Rent", time.Now(), expenses.FixedExpense)
	r.expenses = append(r.expenses, *ex)
}

func (r *ExpenseInMemoryRepository) FindAll() ([]expenses.Expense, error) {
	return r.expenses, nil
}

func (r *ExpenseInMemoryRepository) Save(expense expenses.Expense) error {
	r.expenses = append(r.expenses, expense)
	return nil
}

func (r *ExpenseInMemoryRepository) FindByID(id string) (*expenses.Expense, error) {
	for _, e := range r.expenses {
		if e.ID() == id {
			return &e, nil
		}
	}
	return nil, nil
}

func (r *ExpenseInMemoryRepository) Update(expense expenses.Expense) error {
	for i, e := range r.expenses {
		if e.ID() == expense.ID() {
			r.expenses[i] = expense
			return nil
		}
	}
	return nil
}

func (r *ExpenseInMemoryRepository) Delete(id string) error {
	for i, e := range r.expenses {
		if e.ID() == id {
			r.expenses = append(r.expenses[:i], r.expenses[i+1:]...)
			return nil
		}
	}
	return nil
}
