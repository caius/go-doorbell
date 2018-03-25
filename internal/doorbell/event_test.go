package doorbell

import (
	"testing"
)

func TestNewEvent(t *testing.T) {

	e := NewEvent()
	if e.Value != 0 {
		t.Fatal("Expected value to be zero")
	}
}

func TestNewPressEvent(t *testing.T) {
	e := NewPressEvent()

	if e.Value != 1 {
		t.Fatalf("Expected value to be 1, got %d", e.Value)
	}
}

func TestNewDepressEvent(t *testing.T) {
	e := NewDepressEvent()

	if e.Value != 0 {
		t.Fatalf("Expected value to be 0, got %d", e.Value)
	}
}

func TestRingRing(t *testing.T) {
	press := NewPressEvent()
	if press.RingRing() != true {
		t.Fatal("Expected press to RingRing")
	}

	depress := NewDepressEvent()
	if depress.RingRing() != false {
		t.Fatal("Expected press NOT to RingRing")
	}
}
