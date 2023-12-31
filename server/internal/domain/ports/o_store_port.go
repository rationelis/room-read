package ports

import (
	"context"
	"room_read/internal/domain/model"
)

type StorePort interface {
	GetMessages(ctx context.Context) ([]*model.Message, error)
	StoreMessage(ctx context.Context, message *model.Message) (*model.Message, error)
}
