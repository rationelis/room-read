package ports

import "room_read/internal/domain/model"

type MQTTPort interface {
	ProcessMessage(message model.Message) error
}
