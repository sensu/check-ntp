[![Sensu Bonsai Asset](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/nixwiz/check-ntp)
![Go Test](https://github.com/nixwiz/check-ntp/workflows/Go%20Test/badge.svg)
![goreleaser](https://github.com/nixwiz/check-ntp/workflows/goreleaser/badge.svg)

# Sensu NTP offset check for Linux (ntpd or chrony)

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Check definition](#check-definition)
- [Installation from source](#installation-from-source)
- [Contributing](#contributing)

## Overview

The Sensu NTP Check is a [Sensu Check][1] that provides alerting on NTP offset
and a set of NTP metrics. Metrics are provided in [nagios_perfdata][5] format.
This check supports both ntpd and chrony NTP services running under Linux.

## Usage examples

```
Check NTP offset and provide metrics

Usage:
  check-ntp [flags]
  check-ntp [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -c, --critical float   Critical threshold for offset in ms (default 100)
  -w, --warning float    Warning threshold for offset in ms (default 10)
  -h, --help             help for check-ntp

Use "check-ntp [command] --help" for more information about a command.
```

## Configuration

### Asset registration

[Sensu Assets][2] are the best way to make use of this plugin. If you're not
using an asset, please consider doing so! If you're using sensuctl 5.13 with
Sensu Backend 5.13 or later, you can use the following command to add the asset:

```
sensuctl asset add nixwiz/check-ntp
```

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index][3].

### Check definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: check-ntp
  namespace: default
spec:
  command: >-
    check-ntp
    --critical 20
    --warning 10
  output_metric_format: nagios_perfdata
  output_metric_handlers:
    - influxdb
  subscriptions:
  - system
  runtime_assets:
  - nixwiz/check-ntp
```

## Installation from source

The preferred way of installing and deploying this plugin is to use it as an
Asset. If you would like to compile and install the plugin from source or
contribute to it, download the latest version or create an executable from this
source.

From the local path of the check-ntp repository:

```
go build
```

## Contributing

For more information about contributing to this plugin, see [Contributing][4].

[1]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[2]: https://docs.sensu.io/sensu-go/latest/reference/assets/
[3]: https://bonsai.sensu.io/assets/nixwiz/check-ntp
[4]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[5]: https://docs.sensu.io/sensu-go/latest/observability-pipeline/observe-schedule/collect-metrics-with-checks/#supported-output-metric-formats
