package doorbell

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func (s *Sensor) listenForPress() {
	log.Info("Darwin sensor listening for press")
	s.doorbellPressed()

	ticker := time.NewTicker(time.Second * 10)

	for _ = range ticker.C {
		s.doorbellPressed()
		time.Sleep(3)
	}

}
