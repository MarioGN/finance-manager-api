package server

import (
	"github.com/MarioGN/finance-manager-api/data"

	"github.com/MarioGN/finance-manager-api/internal/expenses/dto"
	"github.com/MarioGN/finance-manager-api/internal/expenses/usecase"

	"github.com/labstack/echo/v4"
)

type server struct {
	echo  *echo.Echo
	store *data.Store
}

func New(store *data.Store) *server {
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
	{
		expensesGroup.GET("", s.getExpensesHandler)
		expensesGroup.POST("", s.createExpenseHandler)

		expensesGroup.GET("/:id", s.getExpenseByIDHandler)
		expensesGroup.PUT("/:id", s.updateExpenseHandler)
		expensesGroup.DELETE("/:id", s.deleteExpenseHandler)
	}

}

func (s *server) getExpensesHandler(c echo.Context) error {
	uc := usecase.NewGetExpensesUseCase(*s.store)

	res, err := uc.Execute()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to list expenses"})
	}

	return c.JSON(200, res)
}

func (s *server) createExpenseHandler(c echo.Context) error {
	var req dto.CreateExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request payload"})
	}

	uc := usecase.NewCreateExpenseUseCase(*s.store)

	res, err := uc.Execute(req)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to create expense"})
	}

	return c.JSON(201, res)
}

func (s *server) getExpenseByIDHandler(c echo.Context) error {
	id := c.Param("id")
	uc := usecase.NewGetExpenseUseCase(*s.store)

	res, err := uc.Execute(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to get expense"})
	}

	if res == nil {
		return c.JSON(404, map[string]string{"error": "expense not found"})
	}

	return c.JSON(200, res)
}

func (s *server) updateExpenseHandler(c echo.Context) error {
	var req dto.UpdateExpenseRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request payload"})
	}

	id := c.Param("id")

	uc := usecase.NewUpdateExpenseUseCase(*s.store)

	res, err := uc.Execute(id, req)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to update expense"})
	}

	return c.JSON(200, res)
}

func (s *server) deleteExpenseHandler(c echo.Context) error {
	id := c.Param("id")

	uc := usecase.NewDeleteExpenseUseCase(*s.store)

	err := uc.Execute(id)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "failed to delete expense"})
	}

	return c.NoContent(204)
}
