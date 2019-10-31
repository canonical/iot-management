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

package twinapi

import (
	"testing"
)

func TestClientAdapter_ActionList(t *testing.T) {
	b1 := `{"actions": [{"deviceId":"a111"}]}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		deviceID string
		body     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", b1}, 1, ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", b1}, 0, "MOCK error get"},
		{"invalid-body", fields{""}, args{"abc", "a111", ""}, 0, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.ActionList(tt.args.orgID, tt.args.deviceID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.DeviceGet() = %v, want %v", got.Message, tt.wantErr)
			}
			if len(got.Actions) != tt.want {
				t.Errorf("ClientAdapter.DeviceGet() = %v, want %v", len(got.Actions), tt.want)
			}
		})
	}
}
