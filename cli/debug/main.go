package main

import (
	"github.com/caius/go-doorbell/internal/doorbell"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func main() {

	pressChannel := make(chan doorbell.Event)
	sensor := doorbell.NewSensor(11, pressChannel)

	go sensor.Start()

	for _ = range pressChannel {
		log.Info("Doorbell press received!")
	}

}
