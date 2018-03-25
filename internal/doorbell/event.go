package doorbell

import (
	"time"
)

type Event struct {
	TriggeredAt time.Time
}

func NewEventNow() Event {
	return Event{TriggeredAt: time.Now()}
}
