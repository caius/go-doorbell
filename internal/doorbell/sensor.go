package doorbell

import (
	log "github.com/sirupsen/logrus"
)

type Sensor struct {
	Pin    int
	Output chan Event
}

func NewSensor(pin int, output chan Event) Sensor {
	return Sensor{
		Pin:    pin,
		Output: output,
	}
}

// Blocks
func (s *Sensor) Start() {
	log.Info("Sensor starting up")
	s.listenForPress()
}

func (s *Sensor) doorbellPressed() {
	log.Info("Doorbell pressed!")
	s.Output <- NewEventNow()
}
