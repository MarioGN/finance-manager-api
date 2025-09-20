package main

import (
	"github.com/MarioGN/finance-manager-api/data"
	"github.com/MarioGN/finance-manager-api/server"
)

func main() {
	store := data.NewStore()
	srv := server.New(store)

	srv.Start()
}
