package server

import (
	"github.com/MarioGN/finance-manager-api/data"
	controller "github.com/MarioGN/finance-manager-api/server/controllers"

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
	controller.ConfigureExpenseRoutes(expensesGroup, s.store)
}
