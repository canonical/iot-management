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
	"github.com/CanonicalLtd/iot-management/config"
	"github.com/CanonicalLtd/iot-management/identityapi"
	"github.com/CanonicalLtd/iot-management/service/factory"
	"github.com/CanonicalLtd/iot-management/service/manage"
	"github.com/CanonicalLtd/iot-management/twinapi"
	"github.com/CanonicalLtd/iot-management/web"
	"log"
)

func main() {
	// Parse the command line arguments
	log.Println("Open config file", config.GetPath())
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
