package main

import "encoding/json"

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

const (
	EventSendMessage = "send_message"
)

type EventHandler func(event Event, c *Client) error

type SendMessage struct {
	Message string `json:"message"`
	From    string `json:"from"`
}
