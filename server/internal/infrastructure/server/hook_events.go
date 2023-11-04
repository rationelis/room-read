package server

import (
	"bytes"
	"room_read/internal/adapters/mqtt"

	mqtt_server "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type HookWrapper struct {
	mqtt_server.HookBase
	controller mqtt.MqttController
}

func (h *HookWrapper) ID() string {
	return "messages-example"
}

func (h *HookWrapper) Provides(b byte) bool {
	return bytes.Contains([]byte{
		mqtt_server.OnPublish,
	}, []byte{b})
}

func (h *HookWrapper) OnPublish(cl *mqtt_server.Client, pk packets.Packet) (packets.Packet, error) {
	h.controller.HandlePacket(cl.ID, pk)
	return pk, nil
}
