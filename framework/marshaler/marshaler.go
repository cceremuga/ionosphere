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

// Converts a packet to its reusable output format.
func ToLogFormat(p *aprs.Packet) string {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err, fmt.Sprintf(" (%s)", p.Raw))
		}
	}()

	// Panic recovery above due to https://github.com/pd0mz/go-aprs/issues/5
	packetType := p.Payload.Type()

	return fmt.Sprintf("%s -> %s [%s] (%f, %f) %s",
		p.Src.Call,
		p.Dst.Call,
		packetTypeName(packetType),
		p.Position.Latitude,
		p.Position.Longitude,
		p.Comment,
	)
}

// Modified from https://github.com/pd0mz/go-aprs/blob/master/data_type.go
var (
	dataTypeNames = map[aprs.DataType]string{
		0x1c: "Mic-E Data",
		0x1d: "Mic-E Data",
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
		'`':  "Mic-E Data",
		'{':  "User-Defined APRS Packet Format",
		'}':  "Third-party Traffic",
	}
)

func packetTypeName(t aprs.DataType) string {
	if typeName, ok := dataTypeNames[t]; ok {
		return typeName
	}

	return "Unknown Packet Type"
}
