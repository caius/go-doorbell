package doorbell

import ()

type Event struct {
	Value int
}

func NewEvent() Event {
	return Event{}
}

func NewPressEvent() Event {
	e := NewEvent()
	e.Value = 1
	return e
}

func NewDepressEvent() Event {
	e := NewEvent()
	e.Value = 0
	return e
}

func (e *Event) RingRing() bool {
	return e.Value == 1
}
