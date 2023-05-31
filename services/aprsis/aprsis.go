// Package aprsis is a service for handling connections, uploads to APRS-IS.
package aprsis

import (
	"errors"
	"fmt"
	"io"
	"net/textproto"
	"strings"

	"github.com/cceremuga/ionosphere/framework/marshaler"
	"github.com/cceremuga/ionosphere/services/log"
	"github.com/fatih/color"
	"github.com/pd0mz/go-aprs"
)

var connected = false
var conn *textproto.Conn
var opts map[string]string

type config struct {
	server   string
	callsign string
	passcode string
}

// Connect connects to APRS-IS with the specified options.
func Connect(options map[string]string) {
	if connected {
		return
	}

	cyan := color.New(color.FgCyan).SprintFunc()

	config, err := validate(options)
	if err != nil {
		log.Fatal(err)
	}

	// Cache
	opts = options

	// Connect
	c, err := textproto.Dial("tcp", config.server)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Debug(fmt.Sprintf("Connected to APRS-IS: %s", config.server))
	}

	// Send auth
	err = c.PrintfLine("user %s pass %s vers Ionosphere 1.0.0-beta2 filter %s",
		config.callsign, config.passcode, opts["filter"])
	if err != nil {
		log.Fatal(err)
	}

	// Server replies with version
	resp, err := c.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	if strings.HasPrefix(resp, "# aprsc ") {
		log.Info(fmt.Sprintf("APRS-IS -> %s: %s", cyan(config.callsign), resp))
	}

	// Server replies with authentication response
	resp, err = c.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	// Validate authentication response.
	err = loggedIn(resp, config.callsign)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Debug("Successfully authenticated with APRS-IS.")
	}

	conn = c
	connected = true

	go func() {
		log.Debug("Listening for responses from APRS-IS.")

		for {
			message, err := c.ReadLine()

			if err != nil {
				log.Error(err)
				if err == io.EOF {
					log.Info("Restart Ionosphere to reconnect. " +
						"Auto-reconnect is not implemented yet.")
					Disconnect()
					break
				}
			} else if !isReadReceipt(message) {
				// These are typically other packets coming _from_ APRS-IS
				p, marshalErr := marshaler.Unmarshal(message)

				if marshalErr != nil {
					log.Info(fmt.Sprintf("%s %s", cyan("[APRS-IS OTHER]"), message))
					continue
				}

				log.Info(fmt.Sprintf(
					"%s %s", cyan("[APRS-IS DIGIPEAT]"), marshaler.ToLogFormat(p)))
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

	err := conn.PrintfLine("%s", s)

	if err != nil {
		log.Error(err)
	}
}

func validate(options map[string]string) (*config, error) {
	if options["server"] == "" {
		return nil, errors.New("no server address specified")
	}

	if options["call-sign"] == "" {
		return nil, errors.New("no callsign specified")
	}

	if options["passcode"] == "" {
		return nil, errors.New("no passcode specified")
	}

	return &config{
		options["server"],
		options["call-sign"],
		options["passcode"],
	}, nil
}

func loggedIn(resp, callsign string) error {
	if strings.HasPrefix(resp, fmt.Sprintf("# logresp %s verified", callsign)) {
		return nil
	}

	return errors.New(resp)
}

func isReadReceipt(message string) bool {
	if strings.HasPrefix(message, "# aprsc") {
		return true
	}

	if strings.HasPrefix(message, "# javAPRSSrvr") {
		return true
	}

	return false
}
