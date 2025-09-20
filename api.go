package main

import (
	"time"

	"github.com/labstack/echo/v4"
)

type server struct {
	echo  *echo.Echo
	store Store
}

func NewServer(store Store) *server {
	return &server{
		echo:  echo.New(),
		store: store,
	}
}

func (s *server) Start() error {
	s.configureRoutes()

	s.echo.Logger.Fatal(s.echo.Start(":3000"))
	return nil
}

func (s *server) configureRoutes() {
	expensesGroup := s.echo.Group("/expenses")
	expensesGroup.GET("", s.getExpensesHandler)
	expensesGroup.POST("", s.createExpenseHandler)
}

func (s *server) getExpensesHandler(c echo.Context) error {
	expenses, err := s.store.FindAllExpenses()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to list expenses"})
	}

	var list []any
	for _, exp := range expenses {
		list = append(list, exp.ToDTO())
	}

	return c.JSON(200, list)
}

func (s *server) createExpenseHandler(c echo.Context) error {
	var req CreateExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request payload"})
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid date format, expected YYYY-MM-DD"})
	}

	newExpense, err := NewExpense(int64(req.Amount*100), req.Description, date, ExpenseType(req.ExpenseType))
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	err = s.store.SaveExpense(*newExpense)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "something went wrong while saving the expense"})
	}

	return c.JSON(201, newExpense.ToDTO())
}
