package mqtt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMqttController(t *testing.T) {
	mockMqttController := NewMockMqttController()

	t.Run("TestMqttController", func(t *testing.T) {
		assert.NotNil(t, mockMqttController)
	})
}
