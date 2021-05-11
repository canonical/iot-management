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
	"flag"
	"fmt"
	"github.com/everactive/iot-management/config"
	"github.com/everactive/iot-management/datastore"
	"github.com/everactive/iot-management/service/factory"
	"log"
	"os"
)

var username, name, email string

func main() {
	// Parse the command line arguments
	settings, err := config.Config(config.GetPath())
	if err != nil {
		log.Fatalf("Error parsing the config file: %v", err)
	}

	// Open the connection to the local database
	db, err := factory.CreateDataStore(settings)
	if err != nil {
		log.Fatalf("Error accessing data store: %v", settings.Driver)
	}

	// Get the command line parameters
	parseFlags()

	// Create the user
	err = run(db, username, name, email)
	if err != nil {
		fmt.Println("Error creating user:", err.Error())
		os.Exit(1)
	}
}

func run(db datastore.DataStore, username, name, email string) error {
	if len(username) == 0 {
		return fmt.Errorf("the username must be supplied")
	}

	// Create the user
	user := datastore.User{
		Username: username,
		Name:     name,
		Email:    email,
		Role:     datastore.Superuser,
	}
	_, err := db.CreateUser(user)
	return err
}

var parseFlags = func() {
	flag.StringVar(&username, "username", "", "Ubuntu SSO username of the user (https://login.ubuntu.com/)")
	flag.StringVar(&name, "name", "Super User", "Full name of the user")
	flag.StringVar(&email, "email", "user@example.com", "Email address of the user")
	flag.Parse()
}
