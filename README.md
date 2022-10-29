<img src="./logo.png" alt="Ionosphere Logo" width="350">

[![Build Status](https://app.travis-ci.com/cceremuga/ionosphere.svg?branch=master)](https://app.travis-ci.com/cceremuga/ionosphere) ![GitHub](https://img.shields.io/github/license/cceremuga/ionosphere)

Receive, decode, log, upload [APRS](http://www.aprs.org/) packets using low cost [RTL-SDR](http://osmocom.org/projects/sdr/wiki/rtl-sdr) devices.

This project is the next-generation successor to [PyPacket](https://gihub.com/cceremuga/pypacket). It should be considered in a "beta" state at best. Please submit bug reports as you encounter them.

## Release Notes

* 10/30/2022 🎃 ([1.0.0-beta2 release]()
    * Updated Golang version to latest.
    * Added additional connection debugging for APRS-IS.
    * Fixed several APRS-IS protocol bugs.
    * Added owncall highlighting to stdout.
    * Documentation updates.
* 4/12/2021
    * Fixed beacon interval.
    * Updated dependencies, removed some.
* 3/9/2021 ([1.0.0-beta1 release](https://github.com/cceremuga/ionosphere/releases/tag/v1.0.0-beta1))
    * Receives, decodes, logs APRS packets to terminal, warnings and errors to file.
    * Uploads APRS packets, periodic beacons to APRS-IS.
    * Allows for full configuration RTL-SDR, multimon-ng options via simple YAML.

## Requirements

To run Ionosphere, the following are required.

* An RTL-SDR compatible device.
* [rtl_fm](http://osmocom.org/projects/sdr/wiki/rtl-sdr)
* [multimon-ng](https://github.com/EliasOenal/multimon-ng)

If you're looking to set up Ionosphere on a Pi, there's a [helpful script here](https://github.com/g7gpr/rpiionosphereinstaller).

## Usage

* Make sure all software in the Requirements section is installed.
* Ensure your RTL-SDR device is connected.
* Download and extract the latest [binary release](https://github.com/cceremuga/ionosphere/releases/) for your OS.
* Edit `config/config.yml` to match your needs.
  * If configured for automatic beaconing, you may edit the config `comment` element to include a latitude, longitude, and symbol i.e. `!DDMM.hhN/DDDMM.hhWIhttp://ionosphere.xyz`
  * Additional documentation on the [APRS protocol](http://www.aprs.net/vm/DOS/PROTOCOL.HTM) and [symbols](http://www.aprs.org/symbols.html) is available.
* In a terminal, from the directory containing Ionosphere, run `./ionosphere`.

## Roadmap

* Unit tests. Shameful there are none yet! :sadpanda:
* Plugin framework.

## Security and Privacy

**The Automatic Packet Reporting System (APRS) is never private and never secure.** As an amateur radio mode, it is designed solely for experimental use by licensed operators to publicly communicate positions and messages. Encryption on amateur radio frequencies is forbidden in most localities. As such, **connections to APRS-IS are also unsecured and only intended for licensed amateur radio operators.**

## Contributing

You are welcome to contribute by submitting pull requests on GitHub if you would like. Feature / enhancement requests may be submitted via GitHub issues.

## License

MIT license, see `LICENSE.md` for more information.
