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

	"github.com/canonical/iot-management/datastore/memory"
	"github.com/canonical/iot-management/identityapi"
	"github.com/canonical/iot-management/twinapi"
)

func TestManagement_RegDeviceList(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid", args{"abc", "jamesj", 300}, "", false},
		{"invalid-permissions", args{"abc", "invalid", 0}, "RegDevicesAuth", true},
		{"invalid", args{"invalid", "jamesj", 300}, "RegDeviceAuth", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			resp := srv.RegDeviceList(tt.args.orgID, tt.args.username, tt.args.role)
			if (len(resp.Code) > 0) != tt.wantErr {
				t.Errorf("Management.RegDeviceList() error = %v, wantErr %v", resp.Code, tt.wantErr)
				return
			}
			if resp.Code != tt.want {
				t.Errorf("Management.OrganizationsForUser() = %v, want %v", resp.Code, tt.want)
			}
		})
	}
}

func TestManagement_RegisterDevice(t *testing.T) {
	d1 := []byte(`{"orgid":"abc", "brand":"deviceinc", "model":"A1000", "serial":"d1234"}`)
	type args struct {
		orgID    string
		username string
		role     int
		body     []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid", args{"abc", "jamesj", 300, d1}, "", false},
		{"invalid-permissions", args{"abc", "invalid", 100, d1}, "RegDeviceAuth", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			got := srv.RegisterDevice(tt.args.orgID, tt.args.username, tt.args.role, tt.args.body)
			if got.Code != tt.want {
				t.Errorf("Management.RegisterDevice() = %v, want %v", got.Code, tt.want)
			}
		})
	}
}

func TestManagement_RegDeviceGet(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid", args{"abc", "jamesj", 300, "a111"}, "", false},
		{"invalid-device", args{"abc", "jamesj", 300, "invalid"}, "RegDeviceAuth", true},
		{"invalid-permissions", args{"abc", "invalid", 100, "a111"}, "RegDeviceAuth", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			got := srv.RegDeviceGet(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID)
			if got.Code != tt.want {
				t.Errorf("Management.RegDeviceGet() = %v, want %v", got.Code, tt.want)
			}
		})
	}
}

func TestManagement_RegDeviceUpdate(t *testing.T) {
	d1 := []byte(`{"status":3}`)
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
		body     []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"valid", args{"abc", "jamesj", 300, "a111", d1}, ""},
		{"invalid-device", args{"abc", "jamesj", 300, "invalid", d1}, "RegDeviceUpdate"},
		{"invalid-permissions", args{"abc", "invalid", 100, "a111", d1}, "RegDeviceAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			got := srv.RegDeviceUpdate(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID, tt.args.body)
			if got.Code != tt.want {
				t.Errorf("Management.RegDeviceUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
