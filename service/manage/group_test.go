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
	"github.com/CanonicalLtd/iot-management/datastore/memory"
	"testing"

	"github.com/CanonicalLtd/iot-management/twinapi"
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
