package doorbell

import (
	"time"
)

type Event struct {
	PressedAt time.Time
}

func NewEventNow() Event {
	return Event{PressedAt: time.Now()}
}
