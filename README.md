[![Build Status](https://travis-ci.org/axnion/hrdwr.svg?branch=master)](https://travis-ci.org/axnion/hrdwr)
[![Codecov](https://img.shields.io/codecov/c/github/axnion/hrdwr.svg)](https://codecov.io/gh/axnion/hrdwr)
# HRDWR
HRDWR (pronounced hardware, because who needs vowels anyways?) gathers data on hardware used in Linux systems and stores it in InfluxDB.

## Features
* Data aggregation
    * CPU usage
    * Memory usage
    * Disk space used
    * Temperature sensor data
    * Fan sensor data
    * Voltage sensor data
* Storing data
    * InfluxDB (not implemented)

## Install 
```bash
$ go get github.com/axnion/hrdwr
```