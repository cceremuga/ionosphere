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

// Modified from https://github.com/pd0mz/go-aprs/blob/master/data_type.go
var (
	dataTypeName = map[aprs.DataType]string{
		0x1c: "Current Mic-E Data",
		0x1d: "Old Mic-E Data",
		'!':  "Position",
		'#':  "Peet Bros U-II Weather Station",
		'$':  "Raw GPS Data or Ultimeter 2000",
		'%':  "Agrelo DFJr / MicroFinder",
		'"':  "Old Mic-E Data",
		')':  "Item",
		'*':  "Peet Bros U-II Weather Station",
		',':  "Invalid Data or Test Data",
		'/':  "Position",
		':':  "Message",
		';':  "Object",
		'<':  "Station Capabilities",
		'=':  "Position",
		'>':  "Status",
		'?':  "Query",
		'@':  "Position",
		'T':  "Telemetry Data",
		'[':  "Maidenhead Grid Locator Beacon",
		'_':  "Weather Report",
		'`':  "Current Mic-E Data",
		'{':  "User-Defined APRS Packet Format",
		'}':  "Third-party Traffic",
	}
)

func PacketTypeName(t aprs.DataType) string {
	if s, ok := dataTypeName[t]; ok {
		return s
	}

	return fmt.Sprintf("Unknown Packet Type %#02x", byte(t))
}
