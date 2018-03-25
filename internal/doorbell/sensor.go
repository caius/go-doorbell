package doorbell

import ()

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
