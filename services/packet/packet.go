// Package packet is a service for interacting with packets.
package packet

import (
	"bufio"
	"io"

	"github.com/cceremuga/ionosphere/framework/marshaler"
	"github.com/cceremuga/ionosphere/services/handler"
	"github.com/cceremuga/ionosphere/services/log"
	"github.com/pd0mz/go-aprs"
)

// Decode packets in the IO stream output from multimon-ng.
func Decode(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		raw := s.Text()

		// Some basic minimums to sift out noise.
		if skip(raw) {
			continue
		}

		p, err := marshaler.Unmarshal(raw)

		if err != nil {
			log.Warn(err)
			continue
		}

		handle(p)
	}
}

const minPacket = 35

func skip(raw string) bool {
	return len(raw) < minPacket
}

func handle(p *aprs.Packet) {
	handlers := handler.All()

	for i := 0; i < len(handlers); i++ {
		handlers[i].Handle(p)
	}
}
