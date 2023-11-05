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
	db, err := sql.Open("sqlite3", configuration.Database.Path)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &database{
		handle: *db,
	}, nil
}

func (d *database) StoreMessage(ctx context.Context, message *model.Message) (*model.Message, error) {
	logging.Info(ctx, fmt.Sprintf("Storing message: %s", message.String()))

	tx, err := d.handle.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	queries := New(&d.handle)
	qtx := queries.WithTx(tx)

	_, err = qtx.persistMessage(ctx, persistMessageParams{
		ClientID:  message.ClientID,
		Topic:     message.Topic,
		Payload:   message.Payload,
		Timestamp: message.Timestamp,
	})

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return message, nil
}
