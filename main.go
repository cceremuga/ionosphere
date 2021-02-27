// Package main creates subprocesses, a pipeline to pass stdout from RTL-SDR to
// stdin of multimon-ng. Then, stdout from multimon-ng is decoded and handled by
// packet handlers.
package main

import (
	"io"
	"os"

	"github.com/cceremuga/ionosphere/services/config"
	"github.com/cceremuga/ionosphere/services/handler"
	"github.com/cceremuga/ionosphere/services/packet"
	"github.com/cceremuga/ionosphere/subprocesses/multimon"
	"github.com/cceremuga/ionosphere/subprocesses/rtlsdr"
	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
)

const logo = `
   ____                        __
  /  _/__  ___  ___  ___ ___  / /  ___ _______
 _/ // _ \/ _ \/ _ \(_-</ _ \/ _ \/ -_) __/ -_)
/___/\___/_//_/\___/___/ .__/_//_/\__/_/  \__/
                      /_/
`

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	color.LightBlue.Println(logo)
	c := config.Load()

	rtl := rtlsdr.Build(&c.Rtl)
	mult := multimon.Build(&c.Multimon)

	r, w := io.Pipe()
	rtl.Stdout = w
	mult.Stdin = r

	// Start each subprocess, then all handlers in bulk.
	rtlsdr.Start(rtl)
	multimon.Start(mult, packet.Decode)
	handler.Start()

	log.Println("Listening for packets.")

	go func() {
		defer w.Close()

		rtl.Wait()
	}()

	mult.Wait()
}
