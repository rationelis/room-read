package database

import (
	"context"
	"database/sql"
	"fmt"
	"room_read/internal/domain/model"
	"room_read/internal/infrastructure/configuration"
	"room_read/internal/infrastructure/logging"

	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	StoreMessage(ctx context.Context, message *model.Message) (*model.Message, error)
}

type database struct {
	handle sql.DB
}

func Connect(configuration *configuration.Configuration) (Database, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}

	return &database{
		handle: *db,
	}, nil
}

func (d *database) StoreMessage(ctx context.Context, message *model.Message) (*model.Message, error) {
	logging.Info(ctx, fmt.Sprintf("Storing message: %s", message.String()))
	return message, nil
}
