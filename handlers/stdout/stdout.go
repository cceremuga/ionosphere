// Package stdout is a packet handler for stdout.
package stdout

import (
	"fmt"

	"github.com/cceremuga/ionosphere/framework/marshaler"
	"github.com/cceremuga/ionosphere/services/log"

	"github.com/fatih/color"
	"github.com/pd0mz/go-aprs"
)

// Stdout helps fulfill the Handler interface contract.
type Stdout struct{}

const id = "4967ade5-7a97-416f-86bf-6e2ae8a5e581"
const name = "Stdout"

// ID defines the Id of this handler.
func (s Stdout) ID() string {
	return id
}

// Name defines the name of this handler.
func (s Stdout) Name() string {
	return name
}

// Enabled specifies that this handler is always enabled.
func (s Stdout) Enabled() bool {
	return true
}

// Start does not do anything in this implementation.
func (s Stdout) Start() {}

// Handle prints packet information to stdout via the log package.
func (s Stdout) Handle(p *aprs.Packet) {
	green := color.New(color.FgHiGreen).SprintFunc()
	log.Info(fmt.Sprintf(
		"%s %s", green("[PACKET]"), marshaler.ToLogFormat(p)))
}

// Stop does not do anything in this implementation.
func (s Stdout) Stop() {}
