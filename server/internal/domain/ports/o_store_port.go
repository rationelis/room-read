package ports

import (
	"context"
	"room_read/internal/domain/model"
)

type StorePort interface {
	StoreMessage(ctx context.Context, message *model.Message) (*model.Message, error)
}
