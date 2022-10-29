// Package marshaler unmarshals packets from strings.
package marshaler

import (
	"fmt"
	"strings"

	"github.com/cceremuga/ionosphere/services/log"
	"github.com/pd0mz/go-aprs"
)

const prefix = "APRS: "

// Unmarshals a raw APRS packet string to a Packet
func Unmarshal(raw string) (*aprs.Packet, error) {
	raw = strings.TrimLeft(raw, prefix)

	defer func() {
		if err := recover(); err != nil {
			log.Error(err, fmt.Sprintf(" (%s)", raw))
		}
	}()

	// Panic recovery above due to https://github.com/pd0mz/go-aprs/issues/5
	p, err := aprs.ParsePacket(raw)

	return &p, err
}
