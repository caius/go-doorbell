package doorbell

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func (s *Sensor) listenForPress() {
	log.Info("Darwin sensor listening for press")
	s.darwinPress()

	ticker := time.NewTicker(time.Second * 10)

	for _ = range ticker.C {
		s.darwinPress()
	}

}

func (s *Sensor) darwinPress() {
	s.doorbellPressed()
	time.Sleep(time.Second * 3)
	s.doorbellDepressed()
}
