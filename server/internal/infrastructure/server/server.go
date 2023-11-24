package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"room_read/internal/adapters/database"
	"room_read/internal/adapters/mqtt"
	"room_read/internal/adapters/rest"
	"room_read/internal/domain"
	"room_read/internal/infrastructure/configuration"
	"room_read/internal/infrastructure/logging"

	mqtt_server "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

type RoomReadServer struct {
	configuration.Configuration
	mqttServer *mqtt_server.Server
	httpServer *http.Server
}

func NewRoomReadServer(configuration *configuration.Configuration) (*RoomReadServer, error) {
	err := logging.SetupLogger(*configuration)
	if err != nil {
		slog.Info("Could not setup logger")
		os.Exit(1)
	}

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

	mqttServer := mqtt_server.New(nil)
	_ = mqttServer.AddHook(new(auth.AllowHook), nil)

	err = mqttServer.AddHook(hookWrapper, map[string]any{})
	if err != nil {
		return nil, err
	}

	connect := fmt.Sprintf(":%d", configuration.Mqtt.Port)
	slog.Info("Setting up MQTT server", "port", connect)

	tcp := listeners.NewTCP("t1", connect, nil)
	err = mqttServer.AddListener(tcp)
	if err != nil {
		return nil, err
	}

	serverMux := http.NewServeMux()

	slog.Info("Creating REST adapter")
	restController := rest.NewRestController(*configuration, domainService)

	serverMux.HandleFunc("/list", restController.RequestHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", configuration.Rest.Port),
		Handler: serverMux,
	}

	return &RoomReadServer{
		mqttServer: mqttServer,
		httpServer: server,
	}, nil
}

func (s *RoomReadServer) Start() error {
	err := s.mqttServer.Serve()
	if err != nil {
		return err
	}
	err = s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *RoomReadServer) Close() error {
	err := s.mqttServer.Close()
	if err != nil {
		return err
	}
	err = s.httpServer.Close()
	if err != nil {
		return err
	}
	return nil
}
