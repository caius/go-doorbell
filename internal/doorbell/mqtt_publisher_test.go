package doorbell

import (
	"fmt"
	"os"
	"testing"
)

func TestMQTTPublisherClientId(t *testing.T) {
	pub := NewMQTTPublisher("mqtt.test:1883", "front")
	expected := fmt.Sprintf("doorbell_front_%d", os.Getpid())

	if pub.clientId() != expected {
		t.Fatalf("clientID expected %s, got %s", expected, pub.clientId())
	}
}

func TestNewMQTTPublisher(t *testing.T) {

	broker := "mqtt.test:1883"
	name := "front"
	pub := NewMQTTPublisher(broker, name)

	if pub.Broker != broker {
		t.Fatalf("Expected broker '%s', got '%s'", broker, pub.Broker)
	}

	if pub.Name != name {
		t.Fatalf("Expected name '%s', got '%s'", name, pub.Name)
	}

}
