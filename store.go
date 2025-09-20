package main

import "time"

type Store interface {
	FindAllExpenses() ([]Expense, error)
	SaveExpense(expense Expense) error
}

type InMemoryStore struct {
	expenses []Expense
}

func NewInMemoryStore() *InMemoryStore {
	expenses := []Expense{}

	exp, _ := NewExpense(10000, "Electricity bill", time.Now(), VariableExpense)
	expenses = append(expenses, *exp)

	exp, _ = NewExpense(50000, "Rent", time.Now(), FixedExpense)
	expenses = append(expenses, *exp)

	return &InMemoryStore{
		expenses: expenses,
	}
}

func (s *InMemoryStore) FindAllExpenses() ([]Expense, error) {
	return s.expenses, nil
}

func (s *InMemoryStore) SaveExpense(expense Expense) error {
	s.expenses = append(s.expenses, expense)
	return nil
}
