package main

import (
	"context"
	"os"
	"os/signal"
	"room_read/internal/infrastructure/configuration"
	"room_read/internal/infrastructure/server"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type operation func(ctx context.Context) error

func main() {
	log.Info().Msg("Starting room-read service")

	_, err := configuration.CreateConfiguration()

	if err != nil {
		log.Fatal().Err(err).Msg("Could not create configuration")
		os.Exit(1)
	}

	log.Info().Msg("Starting MQTT server")

	s, err := server.NewRoomReadServer()

	if err != nil {
		log.Fatal().Err(err).Msg("Could not create MQTT server")
		os.Exit(1)
	}

	go func() {
		err = s.Start()
		// if errors.Is(err, http.ErrServerClosed) {
		// 	log.Info().Msg("MQTT server stopped")
		// }
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

		log.Info().Msg("Shutting down server")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Error().Msgf("Timeout of %d ms has elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})
		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			go func(key string, op operation) {
				defer wg.Done()
				log.Info().Msgf("Shutting down: %s", key)
				if err := op(ctx); err != nil {
					log.Error().Err(err).Msgf("%s: Cleanup failed", key)
					return
				}
				log.Info().Msgf("%s was shut down gracefully", key)
			}(key, op)
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
