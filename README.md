# hrdwr
HRDWR (pronounces hardware) is a small monitoring API for Linux systems, written in Go

## Features
* Fetching and calculation of CPU usage
* Fetching total and used memory
* Fetching total and used disk storage

## Planned Features
* Web sockets for communication
* Fetching temperature data
* Fetching fan data
* Storing history
* Web client

## Download
1. Have go installed
1. Run `go get github.com/axnion/hrdwr`
1. Naviagate to project `$GOPATH/src/github.com/axnion/hrdwr`

## Notes
* Run all tests: `go test ./...`
* Build application: `go build` & `./hrdwr`