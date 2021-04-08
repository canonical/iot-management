[![Build Status][travis-image]][travis-url]
[![Go Report Card][goreportcard-image]][goreportcard-url]
[![codecov][codecov-image]][codecov-url]
# IoT Management Service

The Management service is the end-user web interface to monitor and manage IoT devices.
The service integrates with the [IoT Identity](https://github.com/canonical/iot-identity) and 
[IoT Device Twin](https://github.com/canonical/iot-devicetwin) services to provide device management
for Ubuntu devices.

 
 ## Design
 ![IoT Management Solution Overview](docs/IoTManagement.svg)
 
 ## Build
 The project uses vendorized dependencies using `govendor`. Development has been done on minimum Go version 1.12.1.
 ```bash
 $ go get github.com/canonical/iot-management
 $ cd iot-management
 $ ./get-deps.sh
 $ go build ./...
 ```
 
 ## Run
 ```bash
 go run cmd/management/main.go
 ```
 
The service uses a settings.yaml file for configuration.
 
 ## Contributing
 Before contributing you should sign [Canonical's contributor agreement](https://www.ubuntu.com/legal/contributors), it’s the easiest way for you to give us permission to use your contributions.

[travis-image]: https://travis-ci.org/canonical/iot-management.svg?branch=master
[travis-url]: https://travis-ci.org/canonical/iot-management
[goreportcard-image]: https://goreportcard.com/badge/github.com/canonical/iot-management
[goreportcard-url]: https://goreportcard.com/report/github.com/canonical/iot-management
[codecov-url]: https://codecov.io/gh/canonical/iot-management
[codecov-image]: https://codecov.io/gh/canonical/iot-management/branch/master/graph/badge.svg
