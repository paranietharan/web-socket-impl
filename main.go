package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()

	setupAPI(ctx)
	log.Println("Server started......")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI(ctx context.Context) {
	manager := NewManager(ctx)
	http.Handle("/", http.FileServer(http.Dir("./front-end")))
	http.HandleFunc("/login", manager.loginHandler)
	http.HandleFunc("/ws", manager.serverWS)
}
