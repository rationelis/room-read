package ports

import "room_read/internal/domain/model"

type StorePort interface {
	StoreMessage(message model.Message) error
}
