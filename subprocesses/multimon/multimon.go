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

// Build builds the Command for multimon-ng based upon config and flags.
func Build(c *config.Multimon) *exec.Cmd {
	requiredArgs := []string{"-a", "AFSK1200", "-A", "-t", "raw"}
	userArgs := strings.Fields(c.AdditionalFlags)
	args := append(requiredArgs, userArgs...)
	args = append(args, "-")

	return exec.Command(c.Path, args...)
}

// Start starts the multimon-ng subprocess.
func Start(m *exec.Cmd, f func(reader io.Reader)) {
	stdout, err := m.StdoutPipe()

	if err != nil {
		log.Fatalf("Error reading multimon-ng stdout: %s", err.Error())
	}

	stderr, rErr := m.StderrPipe()

	if rErr != nil {
		log.Fatalf("Error reading multimon-ng stderr: %s", err.Error())
	}

	go readStderr(stderr)

	reader := bufio.NewReader(stdout)
	go f(reader)

	if err := m.Start(); err != nil {
		log.Fatalf("Error starting multimon-ng: %s", err.Error())
	}

	log.Println("multimon-ng initialized.")
}

func readStderr(reader io.Reader) {
	r := bufio.NewReader(reader)

	for true {
		line, _, err := r.ReadLine()

		if line != nil {
			log.Info(string(line))
		}

		if err != nil {
			log.Error(err.Error())
			break
		}
	}
}
