package doorbell

import (
	"github.com/brian-armstrong/gpio"
	log "github.com/sirupsen/logrus"
)

// Listens for the given pin going HIGH
func (s *Sensor) listenForPress() {
	log.Infof("Arm sensor listening to pin %d for press", s.Pin)

	watcher := gpio.NewWatcher()
	watcher.AddPin(uint(s.Pin))

	for e := range watcher.Notification {
		log.WithFields(log.Fields{
			"pin":   e.Pin,
			"value": e.Value,
		}).Debug("Pin triggered")

		switch e.Value {
		case 1:
			s.doorbellPressed()
		case 0:
			s.doorbellDepressed()
		}
	}
}
