package models

import (
	"encoding/json"
	"time"
)

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

const (
	EventSendMessage = "send_message"
	EventNewMessage  = "new_message"
)

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type NewMessageEvent struct {
	SendMessageEvent
	Sent time.Time `json:"sent"`
}
