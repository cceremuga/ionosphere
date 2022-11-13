// Package multimon represents a subprocess for multimon-ng.
package multimon

import (
	"bufio"
	"io"
	"os/exec"
	"strings"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/log"
)

// Build the Command for multimon-ng based upon config and flags.
func Build(c *config.Multimon) *exec.Cmd {
	requiredArgs := []string{
		"-a",
		"AFSK1200",
		"-A",
		"-t",
		"raw",
	}

	userArgs := strings.Fields(c.AdditionalFlags)
	args := append(requiredArgs, userArgs...)
	args = append(args, "-")
	return exec.Command(c.Path, args...)
}

// Start the multimon-ng subprocess.
func Start(c *exec.Cmd, f func(reader io.Reader)) {
	stdout, err := c.StdoutPipe()

	if err != nil {
		log.Fatalf("Error reading multimon-ng stdout: %s", err.Error())
	}

	c.Stderr = log.StderrLogger{}

	reader := bufio.NewReader(stdout)
	go f(reader)

	if err := c.Start(); err != nil {
		log.Fatalf("Error starting multimon-ng: %s", err.Error())
	}

	log.Debug("multimon-ng initialized.")
}
