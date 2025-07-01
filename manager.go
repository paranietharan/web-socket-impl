package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	//
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) serverWS(w http.ResponseWriter, r *http.Request) {
	log.Println("new Connection")

	// upgrade the regualr http to web socket
	conn, err := webSocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	conn.Close()
}
