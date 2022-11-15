// Package discord is a packet handler for posting to discord.
package discord

import (
	"fmt"
	"strconv"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/log"

	"github.com/fatih/color"
	"github.com/pd0mz/go-aprs"
)

// Discord helps fulfill the Handler interface contract.
type Discord struct{}

const id = "b1425d50-04a2-4070-bd1b-851a41377648"
const name = "Discord"

var opts map[string]string

func init() {
	opts = config.HandlerOptions(id)
}

// ID defines the Id of this handler.
func (s Discord) ID() string {
	return id
}

// Name defines the name of this handler.
func (s Discord) Name() string {
	return name
}

// Enabled determines if this handler is enabled based upon the config.
func (s Discord) Enabled() bool {
	enabled, err := strconv.ParseBool(opts["enabled"])

	if err != nil {
		return false
	}

	if !enabled {
		log.Debug("Discord handler inactive.")
	}

	return enabled
}

// Start connects to Discord.
func (s Discord) Start() {
	log.Debug("TODO: Discord handler connecting.")
}

// Handle prints packet information to stdout via the log package.
func (s Discord) Handle(p *aprs.Packet) {
	blue := color.New(color.FgHiBlue).SprintFunc()
	log.Info(fmt.Sprintf("%s %s", blue("[DISCORD]"), p.Raw))
}

// Stop does not do anything in this implementation.
func (s Discord) Stop() {
	log.Debug("TODO: Discord handler disconnecting.")
}
