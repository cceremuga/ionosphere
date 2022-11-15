// Package rtlfm represents a subprocess for rtl_fm.
package rtlfm

import (
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
func Start(c *exec.Cmd) {
	c.Stderr = log.StderrLogger{}

	if err := c.Start(); err != nil {
		log.Fatalf("Error starting rtl_fm: %s", err.Error())
	}

	log.Debug("rtl_fm initialized.")
}
