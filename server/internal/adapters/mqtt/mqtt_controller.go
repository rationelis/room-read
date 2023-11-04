package mqtt

import (
	"context"
	"room_read/internal/domain/model"
	"room_read/internal/domain/ports"
	"room_read/internal/infrastructure/configuration"

	"github.com/mochi-mqtt/server/v2/packets"
)

type MqttController interface {
	HandlePacket(clientID string, pk packets.Packet) error
}

type mqttController struct {
	configuration configuration.Configuration
	messagePort   ports.MessagePort
}

func NewMQTTController(configuration *configuration.Configuration, messagePort ports.MessagePort) MqttController {
	return &mqttController{
		configuration: *configuration,
		messagePort:   messagePort,
	}
}

func (h *mqttController) HandlePacket(clientID string, pk packets.Packet) error {
	ctx := context.Background()
	h.messagePort.ProcessMessage(ctx, model.NewMessage(clientID, pk))
	return nil
}
