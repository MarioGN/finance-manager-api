package main

import "github.com/MarioGN/finance-manager-api/data"

func main() {
	store := data.NewStore()
	srv := NewServer(store)

	srv.Start()
}
