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

package factory

import (
	"testing"

	"github.com/everactive/iot-management/config"
)

func TestCreateDataStore(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"valid", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings, err := config.Config("../../testing/memory.yaml")
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDataStore() settings = %v, wantErr %v", err, tt.wantErr)
				return
			}

			_, err = CreateDataStore(settings)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateDataStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
