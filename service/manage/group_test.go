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

	"github.com/everactive/iot-management/datastore/memory"
	"github.com/everactive/iot-management/twinapi"
)

func TestManagement_GroupList(t *testing.T) {
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
		{"valid", args{"abc", "jamesj", 300}, 1, ""},
		{"invalid-user", args{"abc", "invalid", 200}, 0, "GroupAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.GroupList(tt.args.orgID, tt.args.username, tt.args.role)
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupList() = %v, want %v", got.Code, tt.wantErr)
			}
			if len(got.Groups) != tt.want {
				t.Errorf("Management.GroupList() = %v, want %v", len(got.Groups), tt.want)
			}
		})
	}
}

func TestManagement_GroupCreate(t *testing.T) {
	d1 := []byte(`{"orgid":"abc", "name":"new-group"}`)
	type args struct {
		orgID    string
		username string
		role     int
		body     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300, d1}, ""},
		{"invalid-user", args{"abc", "invalid", 200, d1}, "GroupAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.GroupCreate(tt.args.orgID, tt.args.username, tt.args.role, tt.args.body)
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupCreate() = %v, want %v", got.Code, tt.wantErr)
			}
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupCreate() = %v, want %v", len(got.Code), tt.wantErr)
			}
		})
	}
}

func TestManagement_GroupDevices(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		name     string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300, "workshop"}, 1, ""},
		{"invalid-user", args{"abc", "invalid", 200, "workshop"}, 0, "GroupAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.GroupDevices(tt.args.orgID, tt.args.username, tt.args.role, tt.args.name)
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupDevices() = %v, want %v", got.Code, tt.wantErr)
			}
			if len(got.Devices) != tt.want {
				t.Errorf("Management.GroupDevices() = %v, want %v", len(got.Devices), tt.want)
			}
		})
	}
}

func TestManagement_GroupExcludedDevices(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		name     string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300, "workshop"}, 2, ""},
		{"invalid-user", args{"abc", "invalid", 200, "workshop"}, 0, "GroupAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.GroupExcludedDevices(tt.args.orgID, tt.args.username, tt.args.role, tt.args.name)
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupExcludedDevices() = %v, want %v", got.Code, tt.wantErr)
			}
			if len(got.Devices) != tt.want {
				t.Errorf("Management.GroupExcludedDevices() = %v, want %v", len(got.Devices), tt.want)
			}
		})
	}
}

func TestManagement_GroupDeviceLink(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		name     string
		deviceID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300, "workshop", "a111"}, ""},
		{"invalid-user", args{"abc", "invalid", 200, "workshop", "a111"}, "GroupAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.GroupDeviceLink(tt.args.orgID, tt.args.username, tt.args.role, tt.args.name, tt.args.deviceID)
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupDeviceLink() = %v, want %v", got.Code, tt.wantErr)
			}
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupDeviceLink() = %v, want %v", got.Code, tt.wantErr)
			}
		})
	}
}

func TestManagement_GroupDeviceUnlink(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		name     string
		deviceID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300, "workshop", "a111"}, ""},
		{"invalid-user", args{"abc", "invalid", 200, "workshop", "a111"}, "GroupAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.GroupDeviceUnlink(tt.args.orgID, tt.args.username, tt.args.role, tt.args.name, tt.args.deviceID)
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupDeviceUnlink() = %v, want %v", got.Code, tt.wantErr)
			}
			if got.Code != tt.wantErr {
				t.Errorf("Management.GroupDeviceUnlink() = %v, want %v", got.Code, tt.wantErr)
			}
		})
	}
}
