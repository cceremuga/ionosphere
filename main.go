// Package main creates subprocesses, a pipeline to pass stdout from RTL-SDR to
// stdin of multimon-ng. Then, stdout from multimon-ng is decoded and handled by
// packet handlers.
package main

import (
	"io"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/handler"
	"github.com/cceremuga/ionosphere/services/log"
	"github.com/cceremuga/ionosphere/services/packet"
	"github.com/cceremuga/ionosphere/subprocesses/multimon"
	"github.com/cceremuga/ionosphere/subprocesses/rtlsdr"
	"github.com/cceremuga/ionosphere/tasks/beacon"
	"github.com/gookit/color"
)

const logo = `
   ____                        __
  /  _/__  ___  ___  ___ ___  / /  ___ _______
 _/ // _ \/ _ \/ _ \(_-</ _ \/ _ \/ -_) __/ -_)
/___/\___/_//_/\___/___/ .__/_//_/\__/_/  \__/
                      /_/
`

func main() {
	color.LightBlue.Println(logo)
	c := config.Load()

	rtl := rtlsdr.Build(&c.Rtl)
	mult := multimon.Build(&c.Multimon)

	r, w := io.Pipe()
	rtl.Stdout = w
	mult.Stdin = r

	// Start handlers, then subprocesses.
	handler.Start()
	rtlsdr.Start(rtl)
	multimon.Start(mult, packet.Decode)
	beacon.Start(&c.Beacon)

	log.Println("Listening for packets.")

	go func() {
		defer w.Close()

		rtl.Wait()
	}()

	mult.Wait()
}
