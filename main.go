package main

import (
	"log"
	"net/http"
)

func main() {
	setupAPI()
	log.Println("Server started......")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {
	manager := NewManager()
	http.Handle("/", http.FileServer(http.Dir("./front-end")))
	http.HandleFunc("/ws", manager.serverWS)
}
