// Package beacon allows for periodic beacons to be uploaded to APRS-IS.
package beacon

import (
	"errors"
	"fmt"
	"time"

	"github.com/cceremuga/ionosphere/framework/beacon"
	"github.com/cceremuga/ionosphere/services/aprsis"
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
	log.Println("Uploading beacon.")
	b := beacon.Beacon{
		Src:     c.Call,
		Comment: c.Comment,
	}
	aprsis.UploadRaw(b.String())
}

func validate(c *config.Beacon) error {
	if c.Interval < (time.Duration(30) * time.Minute) {
		return errors.New("interval cannot be < 10m")
	}

	if c.Call == "" {
		return errors.New("beacon call-sign not configured")
	}

	return nil
}
