// Package rtlsdr represents a subprocess for rtl_fm.
package rtlsdr

import (
	"os"
	"os/exec"
	"strings"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/log"
)

// Build the Command for rtl_fm based upon config and flags.
func Build(c *config.Rtl) *exec.Cmd {
	requiredArgs := []string{
		"-f",
		c.Frequency,
		"-s",
		c.SampleRate,
		"-l",
		c.SquelchLevel,
		"-g",
		c.Gain,
		"-p",
		c.PpmError,
	}

	userArgs := strings.Fields(c.AdditionalFlags)
	args := append(requiredArgs, userArgs...)
	args = append(args, "-")
	return exec.Command(c.Path, args...)
}

// Start the rtl_fm subprocess.
func Start(r *exec.Cmd) {
	r.Stderr = os.Stderr

	if err := r.Start(); err != nil {
		log.Fatalf("Error starting rtl_fm: %s", err.Error())
	}

	log.Println("rtl_fm initialized.")
}
