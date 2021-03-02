// Package rtlsdr represents a subprocess for RTL-SDR.
package rtlsdr

import (
	"os/exec"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/log"
)

// Build builds the Command for RTL-SDR based upon config and flags.
func Build(c *config.Rtl) *exec.Cmd {
	args := []string{"-f", c.Frequency, "-s", c.SampleRate, "-l",
		c.SquelchLevel, "-g", c.Gain, "-p", c.PpmError, "-", c.AdditionalFlags}
	return exec.Command(c.Path, args...)
}

// Start starts the RTL-SDR subprocess.
func Start(r *exec.Cmd) {
	if err := r.Start(); err != nil {
		log.Fatalf("Error starting RTL-SDR: %s", err.Error())
	}

	log.Println("RTL-SDR subprocess initialized.")
}
