// Package marshaler unmarshals packets from strings.
package marshaler

import (
	"fmt"
	"strings"

	"github.com/cceremuga/ionosphere/services/log"
	"github.com/pd0mz/go-aprs"
)

const prefix = "APRS: "

func Unmarshal(raw string) *aprs.Packet {
	raw = strings.TrimLeft(raw, prefix)

	defer func() {
		if err := recover(); err != nil {
			log.Error(err, fmt.Sprintf(" (%s)", raw))
		}
	}()

	// Panic recovery above due to https://github.com/pd0mz/go-aprs/issues/5
	p, err := aprs.ParsePacket(raw)

	if err != nil {
		return nil
	}

	return &p
}
