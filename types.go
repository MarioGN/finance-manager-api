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
	id          string
	amount      int64
	description string
	date        time.Time
	expenseType ExpenseType
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

type ApplicationError struct {
	Message string
}
