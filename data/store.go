package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	Expenses ExpenseRepository
	db       *sql.DB
}

func NewStore() (*Store, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	store := &Store{
		db:       db,
		Expenses: NewExpensesSQLiteRepository(db),
	}

	if err := store.init(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *Store) init() error {
	schema := `
	CREATE TABLE IF NOT EXISTS expenses (
		id TEXT PRIMARY KEY,
		amount REAL NOT NULL,
		description TEXT,
		date TEXT NOT NULL,
		expense_type TEXT NOT NULL
	);
	`
	_, err := s.db.Exec(schema)
	return err
}
