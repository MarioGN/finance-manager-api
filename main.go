package main

func main() {
	store := NewInMemoryStore()
	srv := NewServer(store)

	srv.Start()
}
