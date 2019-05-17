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
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	settings, err := Config("../settings.yaml")
	if err != nil {
		t.Errorf("Error reading config file: %v", err)
	}
	if len(settings.JwtSecret) == 0 {
		t.Errorf("Error generating JWT secret: %v", err)
	}
}

func TestReadConfigNew(t *testing.T) {
	settings, err := Config("./settings.yaml")
	if err != nil {
		t.Errorf("Error reading config file: %v", err)
	}
	if len(settings.JwtSecret) == 0 {
		t.Errorf("Error generating JWT secret: %v", err)
	}
	os.Remove("./settings.yaml")
}
