package mqtt_test

import (
	"room_read/internal/adapters/mqtt"
	"room_read/internal/infrastructure/configuration"
)

func NewMockMqttController() mqtt.MqttController {
	mockConfig := configuration.Configuration{}
	controller := mqtt.NewMQTTController(&mockConfig, nil)
	return controller
}
