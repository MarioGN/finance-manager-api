package data

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/MarioGN/finance-manager-api/internal/expenses/entity"
	_ "github.com/mattn/go-sqlite3"
)

type ExpensesSQLiteRepository struct {
	db *sql.DB
}

func NewExpensesSQLiteRepository(db *sql.DB) *ExpensesSQLiteRepository {
	return &ExpensesSQLiteRepository{db: db}
}

func (r *ExpensesSQLiteRepository) FindAll() ([]entity.Expense, error) {
	expenses := make([]entity.Expense, 0)

	rows, err := r.db.Query("SELECT id, amount, description, date, expense_type FROM expenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		expense, err := scanIntoExpense(rows)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, *expense)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpensesSQLiteRepository) Save(expense entity.Expense) error {
	res, err := r.db.Exec(
		"INSERT INTO expenses (id, amount, description, date, expense_type) VALUES (?, ?, ?, ?, ?)",
		expense.ID(),
		float64(expense.Amount())/100.0,
		expense.Description(),
		expense.Date().Format("2006-01-02"),
		string(expense.ExpenseType()),
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows affected")
	}

	return nil
}

func (r *ExpensesSQLiteRepository) FindByID(id string) (*entity.Expense, error) {
	rows, err := r.db.Query("SELECT id, amount, description, date, expense_type FROM expenses WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		return scanIntoExpense(rows)
	}

	return nil, fmt.Errorf("expense with id %s not found", id)
}

func (r *ExpensesSQLiteRepository) Update(expense entity.Expense) error {
	_, err := r.db.Exec(
		"UPDATE expenses SET amount = ?, description = ?, date = ?, expense_type = ? WHERE id = ?",
		float64(expense.Amount())/100.0,
		expense.Description(),
		expense.Date().Format("2006-01-02"),
		string(expense.ExpenseType()),
		expense.ID(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExpensesSQLiteRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM expenses WHERE id = ?", id)
	return err
}

func scanIntoExpense(rows *sql.Rows) (*entity.Expense, error) {
	type RowStruct struct {
		ID          string
		Amount      float64
		Description string
		Date        string
		ExpenseType string
	}

	var rowStruct RowStruct

	err := rows.Scan(
		&rowStruct.ID,
		&rowStruct.Amount,
		&rowStruct.Description,
		&rowStruct.Date,
		&rowStruct.ExpenseType,
	)

	if err != nil {
		return nil, err
	}

	date, err := time.Parse("2006-01-02", rowStruct.Date)
	if err != nil {
		return nil, err
	}

	expense, err := entity.NewExpense(
		int64(rowStruct.Amount)*100,
		rowStruct.Description,
		date,
		entity.ExpenseType(rowStruct.ExpenseType),
	)

	if err != nil {
		return nil, err
	}

	expense.SetID(rowStruct.ID)

	return expense, nil
}
