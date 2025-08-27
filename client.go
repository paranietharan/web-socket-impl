package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager

	egres chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egres:      make(chan []byte),
	}
}

func (client *Client) readMessages() {
	defer func() {
		client.manager.removeClient(client) // clean up the connection
	}()

	if err := client.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println("Error setting read deadline:", err)
	}

	client.connection.SetReadLimit(512)

	client.connection.SetPongHandler(client.pongHandler)

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

		for wsClient := range client.manager.clients {
			wsClient.egres <- payload
		}

		log.Println(messageType)
		log.Println(string(payload))
	}
}

func (client *Client) writeMessages() {
	defer func() {
		client.manager.removeClient(client)
	}()

	ticker := time.NewTicker(pingInterval)

	for {
		select {
		case message, ok := <-client.egres:
			if !ok {
				if err := client.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("writeMessages Error: ", err)
				}
				return
			}

			if err := client.connection.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Println("writeMessages Error in sending message: ", err)
			}
		case <-ticker.C:
			log.Println("writeMessages: ping")

			// Send a Ping to the Client
			if err := client.connection.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Println("writeMessages Error: ", err)
				return
			}
		}
	}
}

func (client *Client) pongHandler(pongMessage string) error {
	log.Println("pongHandler: ", pongMessage)
	return client.connection.SetReadDeadline(time.Now().Add(pongWait))
}
