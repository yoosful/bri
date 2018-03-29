bri
===
bri is an IOT device management platform for accommodation.
Provides micro-usage tracking and payment, along with security.

## Overview

bri is a platform that will allow accommodation services to track fine-grained usage of IOT devices, charge micro payment in a secured way.

bri provides:
* Easy addition and deletion of devices
* Encrypted identity of devcies
* Safe keeping of data collected from devices
* Fine-grained usage tracking
* Micro-payment based on usage
* Real-time usage tracking with with web page

## Installing

First, get the source code:
```
$ go get github.com/pseohy/bri
```

You can simply build the app from the source code:
```
$ cd $GOPATH/src/github.com/pseohy/bri
$ go build
```

Run `bri`:
```
$ ./bri
bri is an IOT device management platform for accommodation.
Provides micro-usage tracking and payment, along with security

Usage:
  bri [command]

Available Commands:
  config      Configure IOT devices
  device      A device sends data to the server
  help        Help about any command
  serve       Run server that collects data from IOT devices
  summary     Summarize usage of IOT devices

Flags:
      --config string   config file (default is $HOME/.bri.yaml)
  -h, --help            help for bri
  -t, --toggle          Help message for toggle

Use "bri [command] --help" for more information about a command.
```

## Example

### Add and delete devices
Addition and deletion of a device is easy. You can add devices with `config` command.
```
$ ./bri config add <--type|-t> <dtype> <--id|-i> <did>
$ ./bri config delete <--type|-t> <dtype> <--id|-i> <did>
```

### Run server
You can run the server by running simple command:
```
$ ./bri serve
```

### Simulate device
You can simulate message generated from user or device:
```
$ ./bri device <--type|-t> <dtype> <--id|-i> <did> <--user|-user> <user> <--msg|-m> <on|off>
```
