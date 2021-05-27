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

func TestClientAdapter_DeviceList(t *testing.T) {
	b1 := `{"devices": [{"deviceId":"a111"}]}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID string
		body  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", b1}, 1, ""},
		{"invalid-org", fields{""}, args{"invalid", b1}, 0, "MOCK error get"},
		{"invalid-body", fields{""}, args{"abc", ""}, 0, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.DeviceList(tt.args.orgID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.DeviceList() = %v, want %v", got.Message, tt.wantErr)
			}
			if len(got.Devices) != tt.want {
				t.Errorf("ClientAdapter.DeviceList() = %v, want %v", len(got.Devices), tt.want)
			}
		})
	}
}

func TestClientAdapter_DeviceGet(t *testing.T) {
	b1 := `{"device": {"deviceId":"a111"}}`
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
		want    string
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", b1}, "a111", ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", b1}, "", "MOCK error get"},
		{"invalid-body", fields{""}, args{"abc", "a111", ""}, "", "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.DeviceGet(tt.args.orgID, tt.args.deviceID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.DeviceGet() = %v, want %v", got.Message, tt.wantErr)
			}
			if got.Device.DeviceID != tt.want {
				t.Errorf("ClientAdapter.DeviceGet() = %v, want %v", got.Device.DeviceID, tt.want)
			}
		})
	}
}

func TestClientAdapter_DeviceLogs(t *testing.T) {
	b1 := `{"url":"", "https://upload.com/upload", "limit": 10}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		deviceID string
		body     string
		data     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", b1, []byte("{}")}, ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", b1, []byte("{}")}, "MOCK error post"},
		{"invalid-body", fields{""}, args{"abc", "a111", "", []byte("{}")}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.DeviceLogs(tt.args.orgID, tt.args.deviceID, tt.args.data)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.DeviceLogs() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}
