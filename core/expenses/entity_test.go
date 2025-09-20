package expenses

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewExpense_ExpensesCreation(t *testing.T) {
	testDate := time.Now()

	t.Run("Create valid expense", func(t *testing.T) {
		tests := []struct {
			name        string
			amount      int64
			description string
			date        time.Time
			expenseType ExpenseType
		}{
			{
				name:        "Valid fixed expense",
				amount:      150000, // $ 1500.00 in cents
				description: "Loan installment",
				date:        testDate,
				expenseType: FixedExpense,
			},
			{
				name:        "Valid variable expense",
				amount:      8500, // $ 85.00 in cents
				description: "Electricity bill",
				date:        testDate,
				expenseType: VariableExpense,
			},
			{
				name:        "Valid unplanned expense",
				amount:      2500, // $ 25.00 in cents
				description: "Restaurant dinner",
				date:        testDate,
				expenseType: UnplannedExpense,
			},
		}

		for _, tc := range tests {
			expense, err := NewExpense(tc.amount, tc.description, tc.date, tc.expenseType)
			assert.NoError(t, err, "Expected no error when creating a valid expense")

			assert.NotNil(t, expense, "Expense should not be nil")
			assert.NotEmpty(t, expense.id, "Expense ID should not be empty")

			_, err = uuid.Parse(expense.id)
			assert.NoError(t, err, "Expense ID should be a valid UUID")

			assert.Equal(t, tc.amount, expense.amount, "Expense amount should match")
			assert.Equal(t, tc.description, expense.description, "Expense description should match")
			assert.Equal(t, tc.date, expense.date, "Expense date should match")
			assert.Equal(t, tc.expenseType, expense.expenseType, "Expense type should match")
		}
	})

	t.Run("Should not allow expenses with zero or negative amount", func(t *testing.T) {
		tests := []struct {
			name        string
			amount      int64
			description string
			date        time.Time
			expenseType ExpenseType
		}{
			{
				name:        "Valid fixed expense",
				amount:      0,
				description: "Loan installment",
				date:        testDate,
				expenseType: FixedExpense,
			},
			{
				name:        "Valid variable expense",
				amount:      -10000,
				description: "Electricity bill",
				date:        testDate,
				expenseType: VariableExpense,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				expense, err := NewExpense(tc.amount, tc.description, tc.date, tc.expenseType)
				assert.Error(t, err, "Expected error when creating an expense with zero or negative amount")
				assert.Nil(t, expense, "Expense should be nil when creation fails")
			})
		}

	})

	t.Run("Should not allow expenses with no date set", func(t *testing.T) {
		testDate := time.Time{} // Zero value of time.Time

		expense, err := NewExpense(10000, "Electricity bill", testDate, VariableExpense)
		assert.Error(t, err, "Expected error when creating an expense with no date set")
		assert.Nil(t, expense, "Expense should be nil when creation fails")
	})

	t.Run("Should create expenses when expenseType is valid", func(t *testing.T) {
		tests := []struct {
			name        string
			amount      int64
			description string
			date        time.Time
			expenseType ExpenseType
		}{
			{
				name:        "Valid fixed expense",
				amount:      190000,
				description: "Loan installment",
				date:        testDate,
				expenseType: FixedExpense,
			},
			{
				name:        "Valid variable expense",
				amount:      8000,
				description: "Electricity bill",
				date:        testDate,
				expenseType: VariableExpense,
			},
			{
				name:        "Valid unplanned expense",
				amount:      2500,
				description: "Restaurant dinner",
				date:        testDate,
				expenseType: UnplannedExpense,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				expense, err := NewExpense(tc.amount, tc.description, tc.date, tc.expenseType)
				assert.NoError(t, err, "Expected no error when creating an expense with a valid expenseType")
				assert.NotNil(t, expense, "Expense should not be nil when creation succeeds")
			})
		}
	})

	t.Run("Should return an error when expenseType is not valid", func(t *testing.T) {
		invalidExpenseType := ExpenseType("invalid_type")

		expense, err := NewExpense(10000, "Electricity bill", testDate, invalidExpenseType)
		assert.Error(t, err, "Expected error when creating an expense with an invalid expenseType")
		assert.Nil(t, expense, "Expense should be nil when creation fails")
	})
}
