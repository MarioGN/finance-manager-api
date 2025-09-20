package main

import (
	"time"

	exp "github.com/MarioGN/finance-manager-api/core/expenses"
)

type Store interface {
	FindAllExpenses() ([]exp.Expense, error)
	SaveExpense(expense exp.Expense) error
	FindExpenseByID(id string) (*exp.Expense, error)
	UpdateExpense(expense exp.Expense) error
	DeleteExpense(id string) error
}

type InMemoryStore struct {
	expenses []exp.Expense
}

func NewInMemoryStore() *InMemoryStore {
	expenses := []exp.Expense{}

	ex, _ := exp.NewExpense(10000, "Electricity bill", time.Now(), exp.VariableExpense)
	expenses = append(expenses, *ex)

	ex, _ = exp.NewExpense(50000, "Rent", time.Now(), exp.FixedExpense)
	expenses = append(expenses, *ex)

	return &InMemoryStore{
		expenses: expenses,
	}
}

func (s *InMemoryStore) FindAllExpenses() ([]exp.Expense, error) {
	return s.expenses, nil
}

func (s *InMemoryStore) SaveExpense(expense exp.Expense) error {
	s.expenses = append(s.expenses, expense)
	return nil
}

func (s *InMemoryStore) FindExpenseByID(id string) (*exp.Expense, error) {
	for _, e := range s.expenses {
		if e.ID() == id {
			return &e, nil
		}
	}
	return nil, nil
}

func (s *InMemoryStore) UpdateExpense(expense exp.Expense) error {
	for i, e := range s.expenses {
		if e.ID() == expense.ID() {
			s.expenses[i] = expense
			return nil
		}
	}
	return nil
}

func (s *InMemoryStore) DeleteExpense(id string) error {
	for i, exp := range s.expenses {
		if exp.ID() == id {
			s.expenses = append(s.expenses[:i], s.expenses[i+1:]...)
			return nil
		}
	}
	return nil
}
