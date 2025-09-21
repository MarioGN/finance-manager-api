package entity

import (
	"errors"
	"time"

	dto "github.com/MarioGN/finance-manager-api/internal/expenses/dto"
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

func (e *Expense) SetAmount(amount int64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	e.amount = amount
	return nil
}

func (e *Expense) SetDescription(description string) {
	e.description = description
}

func (e *Expense) SetDate(date time.Time) error {
	if date.IsZero() {
		return errors.New("date must be a valid date")
	}
	e.date = date
	return nil
}

func (e *Expense) SetExpenseType(expenseType ExpenseType) error {
	if !expenseType.IsValid() {
		return errors.New("invalid expense type")
	}
	e.expenseType = expenseType
	return nil
}

func (e *Expense) ToDTO() *dto.ExpenseDTO {
	floatAmount := float64(e.amount) / 100.0

	return &dto.ExpenseDTO{
		ID:          e.id,
		Amount:      floatAmount,
		Description: e.description,
		Date:        e.date.Format("2006-01-02"),
		ExpenseType: string(e.expenseType),
	}
}

func (e *Expense) ID() string {
	return e.id
}

func (e *Expense) Amount() int64 {
	return e.amount
}

func (e *Expense) Description() string {
	return e.description
}

func (e *Expense) Date() time.Time {
	return e.date
}

func (e *Expense) ExpenseType() ExpenseType {
	return e.expenseType
}

func (e *Expense) SetID(id string) {
	e.id = id
}
