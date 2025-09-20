package data

import (
	"time"

	"github.com/MarioGN/finance-manager-api/internal/expenses/entity"
)

type ExpenseInMemoryRepository struct {
	expenses []entity.Expense
}

func NewExpenseInMemoryRepository() *ExpenseInMemoryRepository {
	r := &ExpenseInMemoryRepository{}
	r._seedData()
	return r
}

func (r *ExpenseInMemoryRepository) _seedData() {
	ex, _ := entity.NewExpense(10000, "Electricity bill", time.Now(), entity.VariableExpense)
	r.expenses = append(r.expenses, *ex)

	ex, _ = entity.NewExpense(50000, "Rent", time.Now(), entity.FixedExpense)
	r.expenses = append(r.expenses, *ex)
}

func (r *ExpenseInMemoryRepository) FindAll() ([]entity.Expense, error) {
	return r.expenses, nil
}

func (r *ExpenseInMemoryRepository) Save(expense entity.Expense) error {
	r.expenses = append(r.expenses, expense)
	return nil
}

func (r *ExpenseInMemoryRepository) FindByID(id string) (*entity.Expense, error) {
	for _, e := range r.expenses {
		if e.ID() == id {
			return &e, nil
		}
	}
	return nil, nil
}

func (r *ExpenseInMemoryRepository) Update(expense entity.Expense) error {
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
