package dto

type ListExpensesResponse struct {
	ID          string `json:"id"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Date        string `json:"date"`
	ExpenseType string `json:"expense_type"`
}

type CreateExpenseRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	ExpenseType string  `json:"expense_type"`
}

type UpdateExpenseRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	ExpenseType string  `json:"expense_type"`
}
