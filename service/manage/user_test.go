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
	"github.com/CanonicalLtd/iot-management/identityapi"
	"testing"

	"github.com/CanonicalLtd/iot-management/datastore/memory"
	"github.com/CanonicalLtd/iot-management/domain"
	"github.com/CanonicalLtd/iot-management/twinapi"
)

func TestManagement_UserWorkflow(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		count   int
		wantErr bool
	}{
		{"valid", args{"jamesj"}, "JJ", 1, false},
		{"invalid-user", args{"invalid"}, "", 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			got, err := srv.GetUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Management.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name != tt.want {
				t.Errorf("Management.GetUser() = %v, want %v", got.Name, tt.want)
			}

			if tt.wantErr {
				return
			}

			users, err := srv.UserList()
			if err != nil {
				t.Errorf("Management.UserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(users) != tt.count {
				t.Errorf("Management.UserList() = %v, want %v", len(users), tt.count)
			}
		})
	}
}

func TestManagement_OpenIDNonceStore(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"valid", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			got := srv.OpenIDNonceStore()
			if (got == nil) != tt.wantErr {
				t.Errorf("Management.OpenIDNonceStore() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestManagement_CreateUser(t *testing.T) {
	u1 := domain.User{Username: "jane", Name: "Jane Doe", Email: "jd@example.com", Role: 200}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{u1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			if err := srv.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Management.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManagement_UserUpdate(t *testing.T) {
	u1 := domain.User{Username: "jamesj", Name: "James Jones", Email: "jj@example.com", Role: 200}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{u1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			if err := srv.UserUpdate(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Management.UserUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManagement_UserDelete(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"jamesj"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""), &identityapi.MockIdentity{})
			if err := srv.UserDelete(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("Management.UserDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
