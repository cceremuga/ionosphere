// Package beacon allows for periodic beacons to be uploaded to APRS-IS.
package beacon

import (
	"errors"
	"fmt"
	"time"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/log"
)

// Start will initiate a Ticker (if enabled) to upload beacons at intervals >= 10m.
func Start(c *config.Beacon) {
	err := validate(c)

	if err != nil && c.Enabled {
		log.Warn(err)
		c.Enabled = false
	}

	if !c.Enabled {
		log.Println("Beacon inactive.")
		return
	}

	startTicker(c)
}

func startTicker(c *config.Beacon) {
	log.Println(fmt.Sprintf("Beacon active, will upload every %s", c.Interval))

	ticker := time.NewTicker(c.Interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				tickerInterval(c)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func tickerInterval(c *config.Beacon) {
	log.Println("TODO: Actually upload beacon to APRS-IS.")
}

func validate(c *config.Beacon) error {
	if c.Interval < (time.Duration(10) * time.Minute) {
		return errors.New("Beacon may not be set to less than a 10 minute interval.")
	}

	if c.Latitude == 0 {
		return errors.New("Beacon latitude not configured.")
	}

	if c.Longitude == 0 {
		return errors.New("Beacon longitude not configured.")
	}

	return nil
}