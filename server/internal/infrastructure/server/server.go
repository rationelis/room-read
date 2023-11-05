package server

import (
	"fmt"
	"log/slog"
	"room_read/internal/adapters/database"
	"room_read/internal/adapters/mqtt"
	"room_read/internal/domain"
	"room_read/internal/infrastructure/configuration"

	mqtt_server "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

type RoomReadServer struct {
	configuration.Configuration
	server *mqtt_server.Server
}

func NewRoomReadServer(configuration *configuration.Configuration) (*RoomReadServer, error) {
	server := mqtt_server.New(nil)
	_ = server.AddHook(new(auth.AllowHook), nil)

	slog.Info("Connecting to database")
	db, err := database.Connect(configuration)

	if err != nil {
		return nil, err
	}

	slog.Info("Starting domain service")
	domainService := domain.NewRoomReadServer(db)

	slog.Info("Creating MQTT controller")
	mqttController := mqtt.NewMQTTController(configuration, domainService)
	hookWrapper := &mqtt.HookWrapper{
		Controller: mqttController,
	}

	err = server.AddHook(hookWrapper, map[string]any{})
	if err != nil {
		return nil, err
	}

	connect := fmt.Sprintf(":%d", configuration.Mqtt.Port)
	slog.Info("Setting up MQTT server", "port", connect)

	tcp := listeners.NewTCP("t1", connect, nil)
	err = server.AddListener(tcp)
	if err != nil {
		return nil, err
	}

	return &RoomReadServer{
		server: server,
	}, nil
}

func (s *RoomReadServer) Start() error {
	err := s.server.Serve()
	if err != nil {
		return err
	}
	return nil
}

func (s *RoomReadServer) Close() error {
	err := s.server.Close()
	if err != nil {
		return err
	}
	return nil
}
