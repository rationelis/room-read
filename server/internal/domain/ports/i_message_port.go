package ports

import (
	"context"
	"room_read/internal/domain/model"
)

type MessagePort interface {
	ProcessMessage(ctx context.Context, message *model.Message) (*model.Message, error)
}
