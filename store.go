package main

import "time"

type Store interface {
	FindAllExpenses() ([]Expense, error)
	SaveExpense(expense Expense) error
	FindExpenseByID(id string) (*Expense, error)
	UpdateExpense(expense Expense) error
	DeleteExpense(id string) error
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

func (s *InMemoryStore) FindExpenseByID(id string) (*Expense, error) {
	for _, exp := range s.expenses {
		if exp.id == id {
			return &exp, nil
		}
	}
	return nil, nil
}

func (s *InMemoryStore) UpdateExpense(expense Expense) error {
	for i, exp := range s.expenses {
		if exp.id == expense.id {
			s.expenses[i] = expense
			return nil
		}
	}
	return nil
}

func (s *InMemoryStore) DeleteExpense(id string) error {
	for i, exp := range s.expenses {
		if exp.id == id {
			s.expenses = append(s.expenses[:i], s.expenses[i+1:]...)
			return nil
		}
	}
	return nil
}
