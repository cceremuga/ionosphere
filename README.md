<img src="./logo.png" alt="Ionosphere Logo" width="350">

[![Build Status](https://travis-ci.com/cceremuga/ionosphere.svg?branch=master)](https://travis-ci.com/cceremuga/ionosphere) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/12d209f6a0af41e594cdc4e881fd4d99)](https://www.codacy.com/gh/cceremuga/ionosphere/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=cceremuga/ionosphere&amp;utm_campaign=Badge_Grade) [![Coverage Status](https://coveralls.io/repos/github/cceremuga/ionosphere/badge.svg?branch=master)](https://coveralls.io/github/cceremuga/ionosphere?branch=master) ![GitHub](https://img.shields.io/github/license/cceremuga/ionosphere)

Receive, decode, log, upload [APRS](http://www.aprs.org/) packets using low cost [RTL-SDR](http://osmocom.org/projects/sdr/wiki/rtl-sdr) devices.

This project is the next-generation successor to [PyPacket](https://gihub.com/cceremuga/pypacket).

It is **very much under active development** and should be considered in a "beta" state at best. Please submit bug reports as you encounter them.

## Release Notes

* 4/??/2022 ([1.0.0-beta2 release]()
    * Fix beacon time interval to match requirement
    *
* 3/9/2021 ([1.0.0-beta1 release](https://github.com/cceremuga/ionosphere/releases/tag/v1.0.0-beta1))
    * Receives, decodes, logs APRS packets to terminal, warnings and errors to file.
    * Uploads APRS packets, periodic beacons to APRS-IS.
    * Allows for full configuration RTL-SDR, multimon-ng options via simple YAML.

## Requirements

To run Ionosphere, the following are required.

* An RTL-SDR compatible device.
* [rtl_fm](http://osmocom.org/projects/sdr/wiki/rtl-sdr)
* [multimon-ng](https://github.com/EliasOenal/multimon-ng)

We use a Raspberry Pi. They're simple to use, compatible, and frankly, just pretty awesome.

## Usage

* Make sure all software in the Requirements section is installed.
* Ensure your RTL-SDR device is connected.
* Download and extract the latest [binary release](https://github.com/cceremuga/ionosphere/releases/) for your OS.
* Edit `config/config.yml` to match your needs.
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
