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

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/everactive/iot-management/crypt"
	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v2"
)

// Version is the application version
const (
	Version                 = "0.2"
	paramsPath              = "."
	paramsFilename          = "settings.yaml"
	defaultURLHost          = "management:8010"
	defaultURLScheme        = "http"
	defaultLocalPort        = "8010"
	defaultDriver           = "memory"
	defaultDataSource       = ""
	defaultDeviceTwinAPIUrl = "http://localhost:8040/v1/"
	defaultIdentityAPIUrl   = "http://localhost:8030/v1/"
	defaultStoreURL         = "https://api.snapcraft.io/api/v1/"
)

// Settings defines the parsed config file settings.
type Settings struct {
	Driver           string `yaml:"driver"`
	DataSource       string `yaml:"datasource"`
	JwtSecret        string `yaml:"jwtSecret"`
	LocalPort        string `yaml:"localport"`
	DeviceTwinAPIUrl string `yaml:"deviceTwinAPIServiceUrl"`
	IdentityAPIUrl   string `yaml:"defaultIdentityAPIUrl"`
	URLHost          string `yaml:"urlHost"`
	URLScheme        string `yaml:"urlScheme"`
	StoreURL         string `yaml:"storeURL"`
	Version          string
}

var settings *Settings

// Config parses the config file
func Config(filePath string) (*Settings, error) {
	settings = &Settings{}
	parseArgs(settings)

	source, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("Error opening the config file. Using default settings")
	} else {
		// File exists, so use it
		err = yaml.Unmarshal(source, settings)
		if err != nil {
			log.Errorf("Error parsing the config file.")
			return settings, err
		}
	}

	if len(settings.JwtSecret) == 0 {
		secret, err := crypt.CreateSecret(32)
		if err != nil {
			fmt.Println("Error generating JWT secret:", err)
			return nil, err
		}
		settings.JwtSecret = secret
		_ = Store(settings, filePath)
	}

	return settings, nil
}

// Store stores the configuration parameters on the filesystem
func Store(c *Settings, p string) error {
	// Create the output file
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	// Convert the parameters to JSON
	b, err := yaml.Marshal(c)
	if err != nil {
		log.Errorf("Error marshalling config parameters: %v", err)
		return err
	}

	// Output the JSON to the file
	_, err = f.Write(b)
	if err != nil {
		log.Errorf("Error storing config parameters: %v", err)
		return err
	}
	_ = f.Sync()

	// // Restrict access to the file
	// err = os.Chmod(p, 0600)
	return nil
}

// parseArgs set up the defaults from the env vars
func parseArgs(c *Settings) {
	var (
		driver           string
		datasource       string
		urlHost          string
		urlScheme        string
		urlDeviceTwinAPI string
		urlIdentityAPI   string
		storeURL         string
	)

	if len(os.Getenv("DRIVER")) > 0 {
		driver = os.Getenv("DRIVER")
	} else {
		driver = defaultDriver
	}
	if len(os.Getenv("DATASOURCE")) > 0 {
		datasource = os.Getenv("DATASOURCE")
	} else {
		datasource = defaultDataSource
	}
	if len(os.Getenv("HOST")) > 0 {
		urlHost = os.Getenv("HOST")
	} else {
		urlHost = defaultURLHost
	}
	if len(os.Getenv("SCHEME")) > 0 {
		urlScheme = os.Getenv("SCHEME")
	} else {
		urlScheme = defaultURLScheme
	}
	if len(os.Getenv("DEVICETWINAPI")) > 0 {
		urlDeviceTwinAPI = os.Getenv("DEVICETWINAPI")
	} else {
		urlDeviceTwinAPI = defaultDeviceTwinAPIUrl
	}
	if len(os.Getenv("IDENTITYAPI")) > 0 {
		urlIdentityAPI = os.Getenv("IDENTITYAPI")
	} else {
		urlIdentityAPI = defaultIdentityAPIUrl
	}
	if len(os.Getenv("STOREURL")) > 0 {
		storeURL = os.Getenv("STOREURL")
	} else {
		storeURL = defaultStoreURL
	}

	// Set the application settings from the environment/default parameters
	c.Version = Version
	c.LocalPort = defaultLocalPort
	c.Driver = driver
	c.DataSource = datasource
	c.URLHost = urlHost
	c.URLScheme = urlScheme
	c.DeviceTwinAPIUrl = urlDeviceTwinAPI
	c.IdentityAPIUrl = urlIdentityAPI
	c.StoreURL = storeURL
}

// GetPath returns the path to the settings file
func GetPath() string {
	return path.Join(paramsPath, paramsFilename)
}
