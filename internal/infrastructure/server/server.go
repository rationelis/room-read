package server

import (
	"log"

	mqtt_server "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

type RoomReadServer struct {
	server *mqtt_server.Server
}

func NewRoomReadServer() (*RoomReadServer, error) {
	server := mqtt_server.New(nil)

	_ = server.AddHook(new(auth.AllowHook), nil)

	tcp := listeners.NewTCP("t1", ":1883", nil)
	err := server.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
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
