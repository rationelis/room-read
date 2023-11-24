package domain

import (
	"context"
	"room_read/internal/domain/model"
	"room_read/internal/domain/ports"
)

type RoomReadServer struct {
	messageStore ports.StorePort
}

func NewRoomReadServer(messageStore ports.StorePort) *RoomReadServer {
	return &RoomReadServer{
		messageStore: messageStore,
	}
}

func (s *RoomReadServer) ProcessMessage(ctx context.Context, message *model.Message) (*model.Message, error) {
	return s.messageStore.StoreMessage(ctx, message)
}

func (s *RoomReadServer) GetMessages(ctx context.Context) ([]*model.Message, error) {
	return s.messageStore.GetMessages(ctx)
}
