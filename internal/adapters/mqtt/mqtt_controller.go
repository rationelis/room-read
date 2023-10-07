package mqtt

type MqttController interface {
}

type mqttController struct {
	// ...
}

func NewMQTTController() MqttController {
	// ...
	return &mqttController{}
}
