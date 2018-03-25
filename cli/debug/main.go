package main

import (
	"flag"
	"github.com/caius/go-doorbell/internal/doorbell"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func main() {

	var pin int
	flag.IntVar(&pin, "pin", 0, "GPIO Pin (Physical number) to listen for doorbell on")
	flag.Parse()

	log.Info("Welcome to doorbell debugger")

	pressChannel := make(chan doorbell.Event)
	sensor := doorbell.NewSensor(pin, pressChannel)
	go sensor.Start()

	for _ = range pressChannel {
		log.Info("Doorbell press received!")
	}

}
