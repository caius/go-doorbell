package doorbell

import (
	"testing"
)

func TestNewEventNow(t *testing.T) {

	e := NewEventNow()
	if e.PressedAt.IsZero() {
		t.Fatal("Expected event to have non-zero time")
	}

}
