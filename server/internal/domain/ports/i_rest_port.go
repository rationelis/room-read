package ports

import (
	"context"
	"room_read/internal/domain/model"
)

type RestPort interface {
	GetMessages(ctx context.Context) ([]*model.Message, error)
}
