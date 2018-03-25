package doorbell

import (
	"github.com/brian-armstrong/gpio"
)

// Listens for the given pin going HIGH
func (s *Sensor) listenForPress() {
	log.Infof("Arm sensor listening to pin %d for press", s.Pin)

	watcher := gpio.NewWatcher()
	watcher.AddPin(s.Pin)

	for e := range watcher.Notification {
		log.WithFields(log.Fields{
			"pin":   e.Pin,
			"value": e.Value,
		}).Info("Pin triggered")
	}

}
