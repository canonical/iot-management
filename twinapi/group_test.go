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

func TestClientAdapter_GroupList(t *testing.T) {
	b1 := `{"groups": [{"orgid":"abc", "name":"workshop"}]}`
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
			got := a.GroupList(tt.args.orgID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupList() = %v, want %v", got.Message, tt.wantErr)
			}
			if len(got.Groups) != tt.want {
				t.Errorf("ClientAdapter.GroupList() = %v, want %v", len(got.Groups), tt.want)
			}
		})
	}
}

func TestClientAdapter_GroupDevices(t *testing.T) {
	b1 := `{"devices": [{"id":1, "orgid":"abc", "brand":"example", "model":"drone-1000", "serial":"a111"}]}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID string
		name  string
		body  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "workshop", b1}, 1, ""},
		{"invalid-org", fields{""}, args{"invalid", "workshop", b1}, 0, "MOCK error get"},
		{"invalid-body", fields{""}, args{"abc", "", ""}, 0, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.GroupDevices(tt.args.orgID, tt.args.name)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupDevices() = %v, want %v", got.Message, tt.wantErr)
			}
			if len(got.Devices) != tt.want {
				t.Errorf("ClientAdapter.GroupDevices() = %v, want %v", len(got.Devices), tt.want)
			}
		})
	}
}

func TestClientAdapter_GroupCreate(t *testing.T) {
	b1 := `{"orgid":"abc", "name":"new-group"}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID string
		name  string
		body  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "workshop", b1}, ""},
		{"invalid-org", fields{""}, args{"invalid", "workshop", b1}, "MOCK error post"},
		{"invalid-body", fields{""}, args{"abc", "workshop", ""}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.GroupCreate(tt.args.orgID, []byte(tt.args.body))
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupCreate() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}

func TestClientAdapter_GroupExcludedDevices(t *testing.T) {
	b1 := `{"devices": [{"id":1, "orgid":"abc", "brand":"example", "model":"drone-1000", "serial":"a111"}]}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID string
		name  string
		body  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "workshop", b1}, 1, ""},
		{"invalid-org", fields{""}, args{"invalid", "workshop", b1}, 0, "MOCK error get"},
		{"invalid-body", fields{""}, args{"abc", "", ""}, 0, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.GroupExcludedDevices(tt.args.orgID, tt.args.name)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupExcludedDevices() = %v, want %v", got.Message, tt.wantErr)
			}
			if len(got.Devices) != tt.want {
				t.Errorf("ClientAdapter.GroupExcludedDevices() = %v, want %v", len(got.Devices), tt.want)
			}
		})
	}
}

func TestClientAdapter_GroupDeviceLink(t *testing.T) {
	b1 := `{"code": "", "message":""}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		name     string
		deviceID string
		body     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "workshop", "a111", b1}, ""},
		{"invalid-org", fields{""}, args{"invalid", "workshop", "a111", ""}, "MOCK error post"},
		{"invalid-body", fields{""}, args{"abc", "", "", ""}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.GroupDeviceLink(tt.args.orgID, tt.args.name, tt.args.deviceID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupDeviceLink() = %v, want %v", got.Message, tt.wantErr)
			}
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupDeviceLink() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}

func TestClientAdapter_GroupDeviceUnlink(t *testing.T) {
	b1 := `{"code": "", "message":""}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		name     string
		deviceID string
		body     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "workshop", "a111", b1}, ""},
		{"invalid-org", fields{""}, args{"invalid", "workshop", "a111", ""}, "MOCK error delete"},
		{"invalid-body", fields{""}, args{"abc", "", "", ""}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.GroupDeviceUnlink(tt.args.orgID, tt.args.name, tt.args.deviceID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupDeviceUnlink() = %v, want %v", got.Message, tt.wantErr)
			}
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupDeviceUnlink() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}
