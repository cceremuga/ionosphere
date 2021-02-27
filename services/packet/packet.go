// Package packet is a service for interacting with packets.
package packet

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/cceremuga/ionosphere/services/handler"
	"github.com/pd0mz/go-aprs"
	log "github.com/sirupsen/logrus"
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

		// Attempt parse.
		p := unmarshal(raw)

		if p == nil {
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

const prefix = "APRS: "

func unmarshal(raw string) *aprs.Packet {
	raw = strings.TrimLeft(raw, prefix)

	defer func() {
		if err := recover(); err != nil {
			log.Error(err, fmt.Sprintf(" (%s)", raw))
		}
	}()

	// Panic recovery above due to https://github.com/pd0mz/go-aprs/issues/5
	p, err := aprs.ParsePacket(raw)

	if err != nil {
		log.Warn(err, fmt.Sprintf(" (%s)", raw))
		return nil
	}

	return &p
}
