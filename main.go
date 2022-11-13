// Package main creates subprocesses, a pipeline to pass stdout from RTL-SDR to
// stdin of multimon-ng. Then, stdout from multimon-ng is decoded and handled by
// packet handlers.
package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/handler"
	"github.com/cceremuga/ionosphere/services/log"
	"github.com/cceremuga/ionosphere/services/packet"
	"github.com/cceremuga/ionosphere/subprocesses/multimon"
	"github.com/cceremuga/ionosphere/subprocesses/rtlfm"
	"github.com/cceremuga/ionosphere/tasks/beacon"
	"github.com/fatih/color"
)

const logo = `
   ____                        __
  /  _/__  ___  ___  ___ ___  / /  ___ _______
 _/ // _ \/ _ \/ _ \(_-</ _ \/ _ \/ -_) __/ -_)
/___/\___/_//_/\___/___/ .__/_//_/\__/_/  \__/
                      /_/

`

func main() {
	color.Cyan(logo)
	log.Debug("Ionosphere initializing.")

	c := config.Load()

	// Pipe io from rtl_fm to multimon-ng
	rtl := rtlfm.Build(&c.Rtl)
	mult := multimon.Build(&c.Multimon)

	r, w := io.Pipe()
	rtl.Stdout = w
	mult.Stdin = r

	// Start handlers, then subprocesses.
	handler.Start()
	rtlfm.Start(rtl)
	multimon.Start(mult, packet.Decode)
	beacon.Start(&c.Beacon)

	// Listen for exit signals, set up for a clean exit.
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigchan
		log.Debug(fmt.Sprintf("Signal (%s) caught, terminating processes.", s))

		w.Close()
		rtl.Process.Kill()
		mult.Process.Kill()

		log.Debug("Ionosphere exiting!")
		os.Exit(0)
	}()

	log.Debug("Listening for packets.")

	// Perform a clean exit if one of the subprocesses terms early.
	go func() {
		rtl.Wait()
		w.Close()
	}()

	mult.Wait()
}
