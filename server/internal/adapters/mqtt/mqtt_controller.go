package mqtt

import (
	"context"
	"room_read/internal/domain/model"
	"room_read/internal/domain/ports"
	"room_read/internal/infrastructure/configuration"
	"room_read/internal/infrastructure/logging"

	"github.com/google/uuid"
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
	ctx := context.WithValue(context.Background(), "traceId", uuid.New().String())

	message := model.NewMessage(clientID, pk)

	if message.ClientID == "" || message.Topic == "" || message.Payload == nil || message.Timestamp.IsZero() {
		logging.Error(ctx, "message has empty fields")
		return nil
	}

	_, err := h.messagePort.ProcessMessage(ctx, message)
	if err != nil {
		logging.WithError(ctx, err)
		return err
	}

	return nil
}
