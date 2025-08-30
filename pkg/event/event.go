package event

import "time"

type DeviceType int

const (
	Temp DeviceType = iota
	Motion
	Door
)

func (d DeviceType) String() string {
	switch d {
	case Temp:
		return "temperature sensor"
	case Motion:
		return "motion sensor"
	case Door:
		return "smart door"
	default:
		return "Unknown"
	}
}

type Event struct {
	DeviceID string
	Type     DeviceType
	Value    interface{}
	Time     time.Time
}
