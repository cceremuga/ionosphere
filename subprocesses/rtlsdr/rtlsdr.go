// Package rtlsdr represents a subprocess for RTL-SDR.
package rtlsdr

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/log"
)

// Build builds the Command for RTL-SDR based upon config and flags.
func Build(c *config.Rtl) *exec.Cmd {
	requiredArgs := []string{"-f", c.Frequency, "-s", c.SampleRate, "-l",
		c.SquelchLevel, "-g", c.Gain, "-p", c.PpmError}
	userArgs := strings.Fields(c.AdditionalFlags)
	args := append(requiredArgs, userArgs...)
	args = append(args, "-")

	return exec.Command(c.Path, args...)
}

// Start starts the RTL-SDR subprocess.
func Start(r *exec.Cmd) {
	go readStderr(r)

	if err := r.Start(); err != nil {
		log.Fatalf("Error starting rtl_fm: %s", err.Error())
	}

	log.Println("rtl_fm initialized.")
}

func readStderr(cmd *exec.Cmd) {
	stderr, err := cmd.StderrPipe()

	if err != nil {
		log.Fatalf("Error reading rtl_fm stderr: %s", err.Error())
	}

	r := bufio.NewReader(stderr)

	for true {
		line, _, err := r.ReadLine()

		if line != nil {
			log.Info(string(line))
		}

		if err != nil {
			log.Error(err.Error())
			cmd.Wait() // prevent hanging subprocess
			break
		}
	}
}
