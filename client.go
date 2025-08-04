package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}

func (client *Client) readMessages() {
	defer func() {
		client.manager.removeClient(client) // clean up the connection
	}()

	for {
		messageType, payload, err := client.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Println("readMessages Error: ", err)
			}
			break
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}
