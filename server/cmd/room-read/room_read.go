package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"room_read/internal/infrastructure/configuration"
	"room_read/internal/infrastructure/logging"
	"room_read/internal/infrastructure/server"
	"sync"
	"syscall"
	"time"
)

type operation func(ctx context.Context) error

func main() {
	ctx := context.Background()

	slog.Info("Starting room-read service")

	configuration, err := configuration.CreateConfiguration()

	if err != nil {
		logging.WithError(ctx, errors.Join(errors.New("Could not create configuration"), err))
		os.Exit(1)
	}

	slog.Info("Starting MQTT server")

	s, err := server.NewRoomReadServer(configuration)

	if err != nil {
		logging.WithError(ctx, errors.Join(errors.New("Could not create MQTT server"), err))
		os.Exit(1)
	}

	go func() {
		err = s.Start()
	}()

	wait := gracefulShutdown(context.Background(), 5*time.Second, map[string]operation{
		"MQTT": func(ctx context.Context) error {
			return s.Close()
		},
	})

	<-wait
}

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL)
		<-signalChan

		logging.Info(ctx, "Shutting down server")

		timeoutFunc := time.AfterFunc(timeout, func() {
			logging.Error(ctx, fmt.Sprintf("Timeout of %d ms has elapsed, force exit", timeout.Milliseconds()))
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			go func(key string, op operation) {
				defer wg.Done()
				logging.Info(ctx, fmt.Sprintf("Shutting down: %s", key))
				if err := op(ctx); err != nil {
					logging.Error(ctx, fmt.Sprintf("%s: Cleanup failed", key))
					return
				}
				logging.Info(ctx, fmt.Sprintf("%s was shut down gracefully", key))
			}(key, op)
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
