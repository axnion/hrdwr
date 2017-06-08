[![Travis](https://img.shields.io/travis/rust-lang/rust.svg)](https://travis-ci.org/axnion/hrdwr.svg?branch=master)
# HRDWR
HRDWR (pronounced hardware, who needs vowels anyways?) is a small monitoring API for Linux systems. It's written in Go in my spare time as a fun side project.

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