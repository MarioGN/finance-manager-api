package main

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type ExpenseType string

func (e ExpenseType) IsValid() bool {
	switch e {
	case FixedExpense, VariableExpense, UnplannedExpense:
		return true
	default:
		return false
	}
}

const (
	FixedExpense     ExpenseType = "fixed"
	VariableExpense  ExpenseType = "variable"
	UnplannedExpense ExpenseType = "unplanned"
)

type Expense struct {
	id          string      `json:"id"`
	amount      int64       `json:"amount"`
	description string      `json:"description"`
	date        time.Time   `json:"date"`
	expenseType ExpenseType `json:"expense_type"`
}

func NewExpense(amount int64, description string, date time.Time, expeseType ExpenseType) (*Expense, error) {
	uuid := uuid.New().String()

	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	if date.IsZero() {
		return nil, errors.New("date must be a valid date")
	}

	if !expeseType.IsValid() {
		return nil, errors.New("invalid expense type")
	}

	return &Expense{
		id:          uuid,
		amount:      amount,
		description: description,
		date:        date,
		expenseType: expeseType,
	}, nil
}

func (e *Expense) ToDTO() *ListExpensesResponse {
	return &ListExpensesResponse{
		ID:          e.id,
		Amount:      e.amount,
		Description: e.description,
		Date:        e.date.Format("2006-01-02"),
		ExpenseType: e.expenseType,
	}
}

type ListExpensesResponse struct {
	ID          string      `json:"id"`
	Amount      int64       `json:"amount"`
	Description string      `json:"description"`
	Date        string      `json:"date"`
	ExpenseType ExpenseType `json:"expense_type"`
}

type CreateExpenseRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	ExpenseType string  `json:"expense_type"`
}
