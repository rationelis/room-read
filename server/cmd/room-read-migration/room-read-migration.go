package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"room_read/internal/infrastructure/configuration"
	"room_read/internal/infrastructure/logging"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type database struct {
	handle sql.DB
}

func main() {
	ctx := context.Background()

	configuration, err := configuration.CreateConfiguration()

	if err != nil {
		logging.Error(ctx, "error creating configuration", slog.String("error", err.Error()))
		os.Exit(1)
	}

	db, err := connect(configuration)

	if err != nil {
		logging.Error(ctx, "error connecting to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	err = runMigration(&db.handle, configuration.Database.Name, configuration.Database.MigrationPath)

	if err != nil {
		logging.Error(ctx, "error running migration", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func connect(configuration *configuration.Configuration) (*database, error) {
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

func runMigration(db *sql.DB, databaseName string, migrationPath string) error {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})

	if err != nil {
		return err
	}

	file, err := os.Open(migrationPath)

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+file.Name(),
		databaseName, driver)

	if err != nil {
		return err
	}

	err = m.Up()

	if err != migrate.ErrNoChange {
		return err
	}

	return nil
}
