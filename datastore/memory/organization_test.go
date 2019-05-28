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

package memory

import (
	"testing"
)

func TestStore_OrgUserAccess(t *testing.T) {
	type args struct {
		orgID    string
		username string
		role     int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"abc", "jamesj", 200}, true},
		{"valid-superuser", args{"abc", "jamesj", 300}, true},
		{"invalid-org", args{"invalid", "jamesj", 200}, false},
		{"invalid-user", args{"abc", "invalid", 200}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			if got := mem.OrgUserAccess(tt.args.orgID, tt.args.username, tt.args.role); got != tt.want {
				t.Errorf("Store.OrgUserAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_OrganizationsForUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"valid", args{"jamesj"}, 1, false},
		{"valid-not-found", args{"unknown"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			got, err := mem.OrganizationsForUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.OrganizationsForUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Store.OrganizationsForUser() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestStore_OrganizationGet(t *testing.T) {
	type args struct {
		orgID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"abc"}, false},
		{"invalid", args{"invalid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			got, err := mem.OrganizationGet(tt.args.orgID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.OrganizationGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got.OrganizationID != tt.args.orgID {
				t.Errorf("Store.OrganizationGet() = %v, want %v", got.OrganizationID, tt.args.orgID)
			}
		})
	}
}
