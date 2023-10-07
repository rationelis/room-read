package model

type Message struct {
	DeviceId   string
	SensorType string
	Value      float64
	Timestamp  int64
}
