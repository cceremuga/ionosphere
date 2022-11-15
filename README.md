<img src="./logo.png" alt="Ionosphere Logo" width="350">

![Build Status](https://github.com/cceremuga/ionosphere/actions/workflows/ci.yml/badge.svg) ![Dependency Status](https://github.com/cceremuga/ionosphere/actions/workflows/dependency-review.yml/badge.svg) ![GitHub](https://img.shields.io/github/license/cceremuga/ionosphere)

Receive, decode, log, upload [APRS](http://www.aprs.org/) packets using low cost [RTL-SDR](http://osmocom.org/projects/sdr/wiki/rtl-sdr) devices.

This project is the next-generation successor to [PyPacket](https://gihub.com/cceremuga/pypacket) with a number of enhancements and simplified cross-platform support. Please submit bug reports as you encounter them.

## Release Notes

* 11/6/2022 ([1.0.2 release](https://github.com/cceremuga/ionosphere/releases/tag/v1.0.2))
    * Fixed [an issue](https://github.com/cceremuga/ionosphere/issues/18) with config-supplied args for `rtl_fm` and `multimon-sg`.
    * Logging is now more verbose, including error output during the startup sequence for the underlying `rtl_fm`, `multimon-ng` processes.
* 11/4/2022 ([1.0.1 release](https://github.com/cceremuga/ionosphere/releases/tag/v1.0.1))
    * Fixed an infinite loop when an unexpected APRS-IS connection drop occurs.
    * Fixed panic crashes caused when parsing packet type names.
* 10/30/2022 🎃 ([1.0.0 release](https://github.com/cceremuga/ionosphere/releases/tag/v1.0.0))
    * Updated Golang version to latest.
    * Updated dependency packages to latest.
    * Added additional connection debugging for APRS-IS.
    * Fixed several APRS-IS protocol bugs.
    * Added APRS-IS digipeat output.
    * Documentation updates.
    * Fixed numerous small bugs.

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
  * If configured for automatic beaconing, you may edit the `comment` element to include a latitude, longitude, and symbol i.e. `!DDMM.hhN/DDDMM.hhWIhttp://ionosphere.xyz RX IGate`
  * You may find additional documentation on the [APRS protocol](http://www.aprs.net/vm/DOS/PROTOCOL.HTM) and [symbols](http://www.aprs.org/symbols.html) useful for custom comment formats.
* In a terminal, from the directory containing Ionosphere, run `./ionosphere`.

## Roadmap

* Unit tests. Shameful there are none yet! :sadpanda:
* Live map showing packets as they are received and uploaded.

## Security and Privacy

**The Automatic Packet Reporting System (APRS) is never private and never secure.** As an amateur radio mode, it is designed solely for experimental use by licensed operators to publicly communicate positions and messages. Encryption on amateur radio frequencies is forbidden in most localities. As such, **connections to APRS-IS are also unsecured and only intended for licensed amateur radio operators.**

## Contributing

You are welcome to contribute by submitting pull requests on GitHub if you would like. Feature / enhancement requests may be submitted via GitHub issues.

## License

MIT license, see `LICENSE.md` for more information.
