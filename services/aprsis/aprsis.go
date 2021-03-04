// Package aprsis is a service for handling connections, uploads to APRS-IS.
package aprsis

import (
	"errors"
	"net/textproto"
	"strings"

	"github.com/cceremuga/ionosphere/services/log"
	"github.com/pd0mz/go-aprs"
)

var connected = false
var conn *textproto.Conn
var opts map[string]string

// Connect connects to APRS-IS with the specified options.
func Connect(options map[string]string) {
	if connected {
		return
	}

	// TODO: Change to struct some time.
	server, callsign, passcode, err := validate(options)
	if err != nil {
		log.Fatal(err)
	}

	// Cache
	opts = options

	// Connect
	c, err := textproto.Dial("tcp", server)
	if err != nil {
		log.Fatal(err)
	}

	// Auth
	err = c.PrintfLine("user %s pass %s vers Ionosphere 0.1%s",
		callsign, passcode, opts["filter"])
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	// Validate connection status.
	err = loggedIn(resp)
	if err != nil {
		log.Fatal(err)
	}

	conn = c
	connected = true
}

// Disconnect disconnects from APRS-IS.
func Disconnect() {
	if !connected {
		return
	}

	conn.Close()
	connected = false
}

// Upload sends a packet to APRS-IS.
func Upload(p *aprs.Packet) {
	if !connected && opts == nil {
		log.Warn("APRS-IS is not connected and not configured. Verify config.")
		return
	}

	if !connected {
		Connect(opts)
	}

	err := conn.PrintfLine(p.Raw)

	if err != nil {
		log.Error(err)
	}
}

func validate(options map[string]string) (string, string, string, error) {
	if options["server"] == "" {
		return "", "", "", errors.New("No server address specified.")
	}

	if options["call-sign"] == "" {
		return "", "", "", errors.New("No callsign specified.")
	}

	if options["passcode"] == "" {
		return "", "", "", errors.New("No passcode specified.")
	}

	return options["server"], options["call-sign"], options["passcode"], nil
}

func loggedIn(resp string) error {
	// TODO: Switch?
	if strings.HasPrefix(resp, "# logresp ") {
		return errors.New(resp)
	}

	if strings.HasPrefix(resp, "# invalid ") {
		return errors.New(resp)
	}

	if strings.HasPrefix(resp, "# login by user not allowed") {
		return errors.New(resp)
	}

	return nil
}
