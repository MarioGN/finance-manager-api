package dto

type ExpenseDTO struct {
	ID          string  `json:"id,omitempty"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
	ExpenseType string  `json:"expense_type"`
}
