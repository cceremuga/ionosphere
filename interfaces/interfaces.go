// Package interfaces contains all interface contracts.
package interfaces

import "github.com/pd0mz/go-aprs"

// Handler represents an interface for handling Packets.
type Handler interface {
	// Id specifies a unique id (GUID) for this handler.
	Id() string

	// Name defines the english name of a given handler for messages.
	Name() string

	// Enabled determines if a given handler is (or should be) enabled.
	Enabled() bool

	// Start performs operations for initialization of a handler.
	Start()

	// Handle handles a single packet.
	Handle(p *aprs.Packet)

	// Start handles operations for the teardown of a handler.
	Stop()
}
