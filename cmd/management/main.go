// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * This file is part of the IoT Management Service
 * Copyright 2019 Canonical Ltd.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU Affero General Public License version 3, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranties of MERCHANTABILITY,
 * SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"os"

	"github.com/everactive/iot-management/config"
	"github.com/everactive/iot-management/identityapi"
	"github.com/everactive/iot-management/service/factory"
	"github.com/everactive/iot-management/service/manage"
	"github.com/everactive/iot-management/twinapi"
	"github.com/everactive/iot-management/web"
	log "github.com/sirupsen/logrus"
)

func main() {
	logLevel := os.Getenv("LOG_LEVEL")
	if len(logLevel) > 0 {
		l, err := log.ParseLevel(logLevel)
		if err != nil {
			log.SetLevel(log.TraceLevel)
			log.Tracef("LOG_LEVEL environment variable is set to %s, could not parse to a valid log level. Using trace logging.", logLevel)
		} else {
			log.SetLevel(l)
			log.Infof("Using LOG_LEVEL %s", logLevel)
		}
	}

	logFormat := os.Getenv("LOG_FORMAT")
	if logFormat == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// Parse the command line arguments
	log.Infof("Open config file %s", config.GetPath())
	settings, err := config.Config(config.GetPath())
	if err != nil {
		log.Fatalf("Error parsing the config file: %v", err)
	}

	// Open the connection to the local database
	db, err := factory.CreateDataStore(settings)
	if err != nil {
		log.Fatalf("Error accessing data store: %v", settings.Driver)
	}

	// Initialize the device twin client
	twinAPI, err := twinapi.NewClientAdapter(settings.DeviceTwinAPIUrl)
	if err != nil {
		log.Fatalf("Error connecting to the device twin service: %v", err)
	}

	// Initialize the identity client
	idAPI, err := identityapi.NewClientAdapter(settings.IdentityAPIUrl)
	if err != nil {
		log.Fatalf("Error connecting to the identity service: %v", err)
	}

	// Create the main services
	srv := manage.NewManagement(settings, db, twinAPI, idAPI)

	// Start the web service
	www := web.NewService(settings, srv)
	log.Fatal(www.Run())
}
