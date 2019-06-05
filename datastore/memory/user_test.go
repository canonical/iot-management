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

	"github.com/CanonicalLtd/iot-management/datastore"
)

func TestStore_UserWorkflow(t *testing.T) {
	tests := []struct {
		name    string
		user    datastore.User
		want    int64
		find    string
		count   int
		wantErr bool
		findErr bool
	}{
		{"valid", datastore.User{Username: "jsmith", Name: "Joseph Smith", Role: 200}, 2, "jsmith", 2, false, false},
		{"invalid-find", datastore.User{Username: "jsmith", Name: "Joseph Smith", Role: 200}, 2, "not-exists", 1, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			got, err := mem.CreateUser(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Store.CreateUser() = %v, want %v", got, tt.want)
			}

			u, err := mem.GetUser(tt.find)
			if (err != nil) != tt.findErr {
				t.Errorf("Store.GetUser() error = %v, findErr %v", err, tt.findErr)
				return
			}
			if tt.findErr {
				return
			}
			if u.Username != tt.find {
				t.Errorf("Store.GetUser() = %v, want %v", u.Username, tt.find)
			}

			users, err := mem.UserList()
			if (err != nil) != tt.findErr {
				t.Errorf("Store.UserList() error = %v, findErr %v", err, tt.findErr)
				return
			}
			if len(users) != tt.count {
				t.Errorf("Store.UserList() = %v, want %v", len(users), tt.count)
			}
		})
	}
}

func TestStore_UserUpdate(t *testing.T) {
	tests := []struct {
		name    string
		user    datastore.User
		wantErr bool
	}{
		{"valid", datastore.User{Username: "jamesj", Name: "James Jones", Role: 200}, false},
		{"invalid", datastore.User{Username: "invalid", Name: "James Jones", Role: 200}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			if err := mem.UserUpdate(tt.user); (err != nil) != tt.wantErr {
				t.Errorf("Store.UserUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_UserDelete(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{"valid", "jamesj", false},
		{"invalid", "invalid", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			if err := mem.UserDelete(tt.username); (err != nil) != tt.wantErr {
				t.Errorf("Store.UserDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
