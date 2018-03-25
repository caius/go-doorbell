package doorbell

import (
	"testing"
)

func TestNewEventNow(t *testing.T) {

	e := NewEventNow()
	if e.TriggeredAt.IsZero() {
		t.Fatal("Expected event to have non-zero time")
	}

}
