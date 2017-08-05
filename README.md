[![Build Status](https://travis-ci.org/axnion/hrdwr.svg?branch=master)](https://travis-ci.org/axnion/hrdwr)
[![Codecov](https://img.shields.io/codecov/c/github/axnion/hrdwr.svg)](https://codecov.io/gh/axnion/hrdwr)
# HRDWR
HRDWR (hardware) is a library for fetching health data on Linux systems.

## Features
* CPU usage
* Memory usage
* Disk space used
* Temperature sensor data
* Fan sensor data
* Voltage sensor data

## Install 
```bash
$ go get github.com/axnion/hrdwr
```

## Usage
Check the example en the example folder for more details
```go
cpus, err := hrdwr.GetCpus()

disks, err := hrdwr.GetDisks()

mem, err := hrdwr.GetMemory()

sensors := hrdwr.GetSensors()
```