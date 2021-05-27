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

func TestManagement_SnapList(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300, "a111"}, 1, ""},
		{"invalid-user", args{"abc", "invalid", 200, "a111"}, 0, "SnapsAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.SnapList(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID)
			if got.Code != tt.wantErr {
				t.Errorf("Management.SnapList() = %v, want %v", got.Code, tt.wantErr)
			}
			if len(got.Snaps) != tt.want {
				t.Errorf("Management.SnapList() = %v, want %v", len(got.Snaps), tt.want)
			}
		})
	}
}

func TestManagement_SnapInstall(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
		snap     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300, "a111", "helloworld"}, ""},
		{"invalid-user", args{"abc", "invalid", 200, "a111", "helloworld"}, "SnapAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.SnapInstall(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID, tt.args.snap)
			if got.Code != tt.wantErr {
				t.Errorf("Management.SnapInstall() = %v, want %v", got.Code, tt.wantErr)
			}
		})
	}
}

func TestManagement_SnapRemove(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
		snap     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid", args{"abc", "jamesj", 300, "a111", "helloworld"}, ""},
		{"invalid-user", args{"abc", "invalid", 200, "a111", "helloworld"}, "SnapAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.SnapRemove(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID, tt.args.snap)
			if got.Code != tt.wantErr {
				t.Errorf("Management.SnapRemove() = %v, want %v", got.Code, tt.wantErr)
			}
		})
	}
}

func TestManagement_SnapUpdate(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
		snap     string
		action   string
		body     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid-enable", args{"abc", "jamesj", 300, "a111", "helloworld", "enable", []byte("{}")}, ""},
		{"invalid-user", args{"abc", "invalid", 200, "a111", "helloworld", "enable", []byte("{}")}, "SnapAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.SnapUpdate(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID, tt.args.snap, tt.args.action, tt.args.body)
			if got.Code != tt.wantErr {
				t.Errorf("Management.SnapUpdate() = %v, want %v", got.Code, tt.wantErr)
			}
		})
	}
}

func TestManagement_SnapConfigSet(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
		snap     string
		config   []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid-enable", args{"abc", "jamesj", 300, "a111", "helloworld", []byte("{}")}, ""},
		{"invalid-user", args{"abc", "invalid", 200, "a111", "helloworld", []byte("{}")}, "SnapAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.SnapConfigSet(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID, tt.args.snap, tt.args.config)
			if got.Code != tt.wantErr {
				t.Errorf("Management.SnapConfigSet() = %v, want %v", got.Code, tt.wantErr)
			}
		})
	}
}

func TestManagement_SnapSnapshot(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
		deviceID string
		snap     string
		body     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{"valid-enable", args{"abc", "jamesj", 300, "a111", "helloworld", []byte("{}")}, ""},
		{"invalid-user", args{"abc", "invalid", 200, "a111", "helloworld", []byte("{}")}, "SnapAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := &Management{
				Settings: getSettings(),
				DB:       memory.NewStore(),
				TwinAPI:  twinapi.NewMockClient(""),
			}
			got := srv.SnapSnapshot(tt.args.orgID, tt.args.username, tt.args.role, tt.args.deviceID, tt.args.snap, tt.args.body)
			if got.Code != tt.wantErr {
				t.Errorf("Management.SnapSnapshot() = %v, want %v", got.Code, tt.wantErr)
			}
		})
	}
}
