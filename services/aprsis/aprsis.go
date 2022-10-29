// Package aprsis is a service for handling connections, uploads to APRS-IS.
package aprsis

import (
	"errors"
	"fmt"
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
	} else {
		log.Info(fmt.Sprintf("Connected to APRS-IS: %s", server))
	}

	// Send auth
	err = c.PrintfLine("user %s pass %s vers Ionosphere 1.0.0-beta2 filter %s",
		callsign, passcode, opts["filter"])
	if err != nil {
		log.Fatal(err)
	}

	// Server replies with version
	resp, err := c.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(resp, "# aprsc ") {
		log.Info(fmt.Sprintf("APRS-IS -> %s: %s", callsign, resp))
	}

	// Server replies with authentication response
	resp, err = c.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	// Validate authentication response.
	err = loggedIn(resp, callsign)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Successfully authenticated with APRS-IS.")
	}

	conn = c
	connected = true

	go func() {
		log.Info("Listening for responses from APRS-IS.")

		for {
			message, err := c.ReadLine()

			if err != nil {
				log.Error(err)
			} else if !strings.HasPrefix(message, "# aprsc") {
				// These are other packets coming _from_ APRS-IS
				log.Info(fmt.Sprintf("APRS-IS -> %s: %s", callsign, message))
			}
		}
	}()
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
	UploadRaw(p.Raw)
}

// UploadRaw sends a raw packet to APRS-IS.
func UploadRaw(s string) {
	if !connected && opts == nil {
		log.Warn("APRS-IS is not connected and not configured. Verify config.")
		return
	}

	if !connected {
		Connect(opts)
	}

	err := conn.PrintfLine(s)

	if err != nil {
		log.Error(err)
	}
}

func validate(options map[string]string) (string, string, string, error) {
	if options["server"] == "" {
		return "", "", "", errors.New("no server address specified")
	}

	if options["call-sign"] == "" {
		return "", "", "", errors.New("no callsign specified")
	}

	if options["passcode"] == "" {
		return "", "", "", errors.New("no passcode specified")
	}

	return options["server"], options["call-sign"], options["passcode"], nil
}

func loggedIn(resp, callsign string) error {
	if strings.HasPrefix(resp, fmt.Sprintf("# logresp %s verified", callsign)) {
		return nil
	}

	return errors.New(resp)
}
