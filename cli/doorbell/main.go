package main

import (
	"flag"
	// "fmt"
	"github.com/caius/go-doorbell/internal/doorbell"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

type Configuration struct {
	Pin        int
	Name       string
	MQTTBroker string
	Valid      bool
	Verbose    bool
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}

func main() {
	// Configuration Section
	config := Configuration{Valid: true}

	// Doorbell Configuration
	flag.StringVar(&config.Name, "name", "", "Name for emitted events")
	flag.IntVar(&config.Pin, "pin", 0, "GPIO Pin (Physical number) to listen for doorbell on")

	// MQTT Configuration
	flag.StringVar(&config.MQTTBroker, "mqtt-broker", "", "MQTT server (eg mqtt.local:1883)")

	// Other configuration
	flag.BoolVar(&config.Verbose, "verbose", false, "Verbose output")

	flag.Parse()

	log.Info("Welcome to doorbell, where we will be greeting visitors shortly.")

	if config.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	// Check configuration is correct
	if config.Name == "" {
		log.Error("--name is a required argument")
		config.Valid = false
	}

	if config.Pin == 0 {
		log.Error("--pin is a required argument")
		config.Valid = false
	}

	if config.MQTTBroker == "" {
		log.Error("--mqtt-broker is a required argument")
		config.Valid = false
	}

	if config.Valid != true {
		log.Error("Configuration errors, please fix. Check logs for more information.")
		os.Exit(1)
	}

	pressChannel := make(chan doorbell.Event)
	sensor := doorbell.NewSensor(config.Pin, pressChannel)

	go sensor.Start()

	publisher := doorbell.NewMQTTPublisher(config.MQTTBroker, config.Name)
	go publisher.Start(pressChannel)

	// Trap and cleanup on interrupt (^C)
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			log.Info("Received interrupt, bringing everything down")

			sensor.Stop()
			publisher.Stop()

			cleanupDone <- true
		}
	}()
	<-cleanupDone
	log.Info("Goodbye!")
}
