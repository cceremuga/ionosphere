// Package aprsis is a packet handler for APRS-IS
package aprsis

import (
	"strconv"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/pd0mz/go-aprs"
	log "github.com/sirupsen/logrus"
)

// APRSIS helps fulfill the Handler interface contract.
type APRSIS struct{}

const id = "b67ac5d5-3612-4618-88a9-a63d36a1777c"
const name = "APRIS-IS"

var opts map[string]string

func init() {
	opts = config.HandlerOptions(id)
}

// Id defines the Id for this handler.
func (s APRSIS) Id() string {
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

	return enabled
}

// Start initializes a connection to APRS-IS.
func (s APRSIS) Start() {
	log.Println("TODO: Connect to APRS-IS.")
}

// Handle uploads packets to APRS-IS if configured.
func (s APRSIS) Handle(p *aprs.Packet) {
	// TODO: upload to APRS-IS here.
}

// Stop does nothing in this implementation.
func (s APRSIS) Stop() {}
