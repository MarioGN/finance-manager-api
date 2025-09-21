package main

import (
	"log"

	"github.com/MarioGN/finance-manager-api/data"
	"github.com/MarioGN/finance-manager-api/server"
)

func main() {
	store, err := data.NewStore()
	if err != nil {
		log.Fatal("Failed to initialize store:", err)
	}

	srv := server.New(store)

	srv.Start()
}
