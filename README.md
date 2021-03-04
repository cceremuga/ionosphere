<img src="./logo.png" alt="Ionosphere Logo" width="350">

[![Build Status](https://travis-ci.org/cceremuga/ionosphere.svg?branch=master)](https://travis-ci.org/cceremuga/ionosphere) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/12d209f6a0af41e594cdc4e881fd4d99)](https://www.codacy.com/gh/cceremuga/ionosphere/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cceremuga/ionosphere&amp;utm_campaign=Badge_Grade) ![GitHub](https://img.shields.io/github/license/cceremuga/ionosphere) 

Receive, decode, log, upload [APRS](http://www.aprs.org/) packets using low cost [RTL-SDR](http://osmocom.org/projects/sdr/wiki/rtl-sdr) devices.

This project is the next-generation successor to [PyPacket](https://gihub.com/cceremuga/pypacket).

It is **very much under active development** and is not a feature-matched replacement yet.

## What it does so far

* Receive, decode, log APRS packets to terminal, warnings and errors to file.
* Allow for full configuration RTL-SDR, multimon-ng options via simple YAML.

## Requirements

To run Ionosphere, the following are required.

* An RTL-SDR compatible device.
* [rtl_fm](http://osmocom.org/projects/sdr/wiki/rtl-sdr)
* [multimon-ng](https://github.com/EliasOenal/multimon-ng)

We use a Raspberry Pi. They're simple to use, compatible, and frankly, just pretty awesome.

## Usage

* Make sure all software in the Requirements section is installed.
* Ensure your RTL-SDR device is connected.
* Download and extract the latest release matching your system.
* Edit `config/config.yml` to match your needs.
* In a terminal, from the directory containing Ionosphere, run `./ionosphere`.

## Roadmap

* Upload packets to APRS-IS.
* Upload timed beacons to APRS-IS (lets you show up on aprs.fi).
* Unit tests. Shameful there are none yet! :sadpanda:
* Plugin framework.

## Security and Privacy

**The Automatic Packet Reporting System (APRS) is never private and never secure.** As an amateur radio mode, it is designed solely for experimental use by licensed operators to publicly communicate positions and messages. Encryption on amateur radio frequencies is forbidden in most localities. As such, **connections to APRS-IS are also unsecured and only intended for licensed amateur radio operators.**

## Contributing

You are welcome to contribute by submitting pull requests on GitHub if you would like. Feature / enhancement requests may be submitted via GitHub issues.

## License

MIT license, see `LICENSE.md` for more information.
