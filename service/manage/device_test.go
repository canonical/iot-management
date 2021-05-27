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

package manage

import (
	"testing"

	"github.com/everactive/iot-management/config"
	"github.com/everactive/iot-management/datastore/memory"
	"github.com/everactive/iot-management/identityapi"

	"github.com/everactive/iot-management/twinapi"
)

var settings *config.Settings

func getSettings() *config.Settings {
	if settings == nil {
		settings, _ = config.Config("../../testing/memory.yaml")
	}
	return settings
}

func TestManagement_DeviceList(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300}, 3, ""},
		{"invalid-user", args{"abc", "invalid", 200}, 0, "DevicesAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			got := srv.DeviceList(tt.args.orgID, tt.args.username, tt.args.role)
			if got.Code != tt.wantErr {
				t.Errorf("Management.DeviceList() = %v, want %v", got.Code, tt.wantErr)
			}
			if len(got.Devices) != tt.want {
				t.Errorf("Management.DeviceList() = %v, want %v", len(got.Devices), tt.want)
			}
		})
	}
}

func TestManagement_DeviceGet(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
	}
	tests := []struct {
		name       string
		args       args
		wantSerial string
		wantErr    string
	}{
		{"valid", args{"abc", "jamesj", 200, "b222"}, "DR1000B222", ""},
		{"invalid-user", args{"abc", "invalid", 200, "b222"}, "", "DeviceAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.DeviceGet(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID)
			if got.Code != tt.wantErr {
				t.Errorf("Management.DeviceGet() = %v, want %v", got.Code, tt.wantErr)
			}
			if got.Device.SerialNumber != tt.wantSerial {
				t.Errorf("Management.DeviceGet() = %v, want %v", got.Device.SerialNumber, tt.wantSerial)
			}
		})
	}
}

func TestManagement_DeviceLogs(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
		body     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid-enable", args{"abc", "jamesj", 300, "a111", []byte("{}")}, ""},
		{"invalid-user", args{"abc", "invalid", 200, "a111", []byte("{}")}, "SnapAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.DeviceLogs(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID, tt.args.snap, tt.args.body)
			if got.Code != tt.wantErr {
				t.Errorf("Management.DeviceLogs() = %v, want %v", got.Code, tt.wantErr)
			}
		})
	}
}
