package mqtt_test

import "room_read/internal/adapters/mqtt"

func NewMockMqttController() mqtt.MqttController {
	controller := mqtt.NewMQTTController()
	return controller
}
