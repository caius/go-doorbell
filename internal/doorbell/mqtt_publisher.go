package doorbell

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
)

type MQTTPublisher struct {
	Broker string
	Name   string
	client mqtt.Client
}

func NewMQTTPublisher(broker string, name string) MQTTPublisher {
	return MQTTPublisher{
		Broker: broker,
		Name:   name,
	}
}

func (p *MQTTPublisher) subscribeToEvents(events <-chan Event) {
	log.Debug("MQTTPublisher subscribing to events")

	for event := range events {
		log.WithFields(log.Fields{
			"event": event,
		}).Info("MQTTPublisher received event")

		p.publishEvent(event)
	}

	log.Debug("MQTTPublisher finished listening for events")
}

// Publishes a 'RING' message to topic 'stat/NAME/RING'
func (p *MQTTPublisher) publishEvent(event Event) {
	topic := fmt.Sprintf("stat/%s/RING", p.Name)

	log.WithFields(log.Fields{
		"topic": topic,
	}).Debug("MQTTPublisher publishing")

	if p.client.IsConnected() {
		msg := "OFF"
		if event.RingRing() {
			msg = "ON"
		}
		p.client.Publish(topic, 0, false, msg).Wait()
	}
}

func (p *MQTTPublisher) clientId() string {
	return fmt.Sprintf("doorbell_%s_%d", p.Name, os.Getpid())
}

func (p *MQTTPublisher) Start(events <-chan Event) {
	log.WithFields(log.Fields{
		"broker":   p.Broker,
		"location": p.Name,
	}).Info("MQTTPublisher publishing")

	willTopic := fmt.Sprintf("tele/%s/LWT", p.Name)

	mqttParams := mqtt.NewClientOptions()
	mqttParams.AddBroker(fmt.Sprintf("tcp://%s", p.Broker))
	mqttParams.SetClientID(p.clientId())
	mqttParams.SetWill(willTopic, "Offline", 0, true)

	p.client = mqtt.NewClient(mqttParams)

	p.client.Connect().Wait()
	log.WithFields(log.Fields{
		"broker": p.Broker,
	}).Info("MQTTPublisher connected to broker")

	p.client.Publish(willTopic, 0, true, "Online").Wait()

	// Do our job
	p.subscribeToEvents(events)
}

func (p *MQTTPublisher) Stop() {
	log.Info("MQTTPublisher received stop")
}
