[![Build Status](https://travis-ci.org/axnion/hrdwr.svg?branch=master)](https://travis-ci.org/axnion/hrdwr)
# HRDWR
HRDWR (pronounced hardware, who needs vowels anyways?) is a small monitoring API for Linux systems written in Go. It's a small side project to help me improve as a programmer both in general but more specifically in Go, therefor suggestions and criticism are welcomed.

## Features
* Data aggregation
    * CPU usage
    * Memory usage
    * Disk space used
    * Temperature sensor data
    * Fan sensor data
    * Voltage sensor data

## Planned Features
* Web sockets for communication
* Storing history
* Web client

## Download
1. Have go installed
1. Run `go get github.com/axnion/hrdwr`
1. Naviagate to project `$GOPATH/src/github.com/axnion/hrdwr`

## Notes
* Run all tests: `go test ./...`
* Build application: `go build` & `./hrdwr`
