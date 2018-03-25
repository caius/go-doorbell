package doorbell

import (
	"testing"
)

func TestNewSensor(t *testing.T) {

	c := make(chan Event)
	s := NewSensor(11, c)

	if s.Pin != 11 {
		t.Fatalf("Expected pin to equal 11, got %d", s.Pin)
	}

	if s.Output != c {
		t.Fatalf("Expected output channel to match given one, but it didn't")
	}

}
