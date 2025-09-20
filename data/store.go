package data

type Store struct {
	Expenses ExpenseRepository
}

func NewStore() *Store {
	return &Store{
		Expenses: NewExpenseInMemoryRepository(),
	}
}
