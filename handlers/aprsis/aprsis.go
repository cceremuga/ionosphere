// Package aprsis is a packet handler for APRS-IS
package aprsis

import (
	"strconv"

	"github.com/cceremuga/ionosphere/services/aprsis"
	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/log"
	"github.com/pd0mz/go-aprs"
)

// APRSIS helps fulfill the Handler interface contract.
type APRSIS struct{}

const id = "b67ac5d5-3612-4618-88a9-a63d36a1777c"
const name = "APRIS-IS"

var opts map[string]string

func init() {
	opts = config.HandlerOptions(id)
}

// ID defines the Id for this handler.
func (s APRSIS) ID() string {
	return id
}

// Name returns this handler's name.
func (s APRSIS) Name() string {
	return name
}

// Enabled determines if this handler is enabled.
func (s APRSIS) Enabled() bool {
	enabled, err := strconv.ParseBool(opts["enabled"])

	if err != nil {
		return false
	}

	if !enabled {
		log.Debug("APRS-IS handler inactive.")
	}

	return enabled
}

// Start initializes a connection to APRS-IS.
func (s APRSIS) Start() {
	aprsis.Connect(opts)
}

// Handle uploads packets to APRS-IS if configured.
func (s APRSIS) Handle(p *aprs.Packet) {
	aprsis.Upload(p)
}

// Stop disconnects from APRS-IS.
func (s APRSIS) Stop() {
	aprsis.Disconnect()
}
