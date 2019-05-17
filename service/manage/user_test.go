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

	"github.com/CanonicalLtd/iot-management/datastore/memory"
	"github.com/CanonicalLtd/iot-management/twinapi"
)

func TestManagement_GetUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"valid", args{"jamesj"}, "JJ", false},
		{"invalid-user", args{"invalid"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""))
			got, err := srv.GetUser(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Management.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Name != tt.want {
				t.Errorf("Management.GetUser() = %v, want %v", got.Name, tt.want)
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
			srv := NewManagement(getSettings(), memory.NewStore(), twinapi.NewMockClient(""))
			got := srv.OpenIDNonceStore()
			if (got == nil) != tt.wantErr {
				t.Errorf("Management.OpenIDNonceStore() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
