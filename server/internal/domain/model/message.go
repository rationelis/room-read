package model

import (
	"fmt"
	"time"

	"github.com/mochi-mqtt/server/v2/packets"
)

type Message struct {
	ClientID  string
	Topic     string
	Payload   string
	Timestamp time.Time
}

func NewMessage(clientID string, pk packets.Packet) *Message {
	return &Message{
		ClientID:  clientID,
		Topic:     pk.TopicName,
		Payload:   string(pk.Payload),
		Timestamp: time.Unix(pk.Created, 0),
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("ClientID: %s, Topic: %s, Payload: %s, Timestamp: %s", m.ClientID, m.Topic, m.Payload, m.Timestamp)
}
