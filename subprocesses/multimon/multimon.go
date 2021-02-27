// Package multimon represents a subprocess for multimon-ng.
package multimon

import (
	"bufio"
	"io"
	"os/exec"

	"github.com/cceremuga/ionosphere/services/config"
	log "github.com/sirupsen/logrus"
)

// Build builds the Command for multimon-ng based upon config and flags.
func Build(c *config.Multimon) *exec.Cmd {
	args := []string{"-a", "AFSK1200", "-A", "-t", "raw", "-", c.AdditionalFlags}
	return exec.Command(c.Path, args...)
}

// Start starts the multimon-ng subprocess.
func Start(m *exec.Cmd, f func(reader io.Reader)) {
	stdout, err := m.StdoutPipe()

	if err != nil {
		log.Fatalf("Error obtaining multimon-ng ouput: %s", err.Error())
	}

	reader := bufio.NewReader(stdout)
	go f(reader)

	if err := m.Start(); err != nil {
		log.Fatalf("Error starting multimon-ng: %s", err.Error())
	}

	log.Println("Multimon-ng subprocess initialized.")
}
