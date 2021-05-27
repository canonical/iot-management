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

func TestClientAdapter_SnapList(t *testing.T) {
	b1 := `{"snaps": [{"deviceId":"a111", "name":"helloworld"}]}`
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
			got := a.SnapList(tt.args.orgID, tt.args.deviceID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.SnapList() = %v, want %v", got.Message, tt.wantErr)
			}
			if len(got.Snaps) != tt.want {
				t.Errorf("ClientAdapter.SnapList() = %v, want %v", len(got.Snaps), tt.want)
			}
		})
	}
}

func TestClientAdapter_SnapInstall(t *testing.T) {
	b1 := `{"code":"", "message":""}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		deviceID string
		snap     string
		body     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", "helloworld", b1}, ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", "helloworld", b1}, "MOCK error post"},
		{"invalid-body", fields{""}, args{"abc", "a111", "helloworld", ""}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.SnapInstall(tt.args.orgID, tt.args.deviceID, tt.args.snap)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.SnapInstall() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}

func TestClientAdapter_SnapRemove(t *testing.T) {
	b1 := `{"code":"", "message":""}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		deviceID string
		snap     string
		body     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", "helloworld", b1}, ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", "helloworld", b1}, "MOCK error delete"},
		{"invalid-body", fields{""}, args{"abc", "a111", "helloworld", ""}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.SnapRemove(tt.args.orgID, tt.args.deviceID, tt.args.snap)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.SnapRemove() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}

func TestClientAdapter_SnapUpdate(t *testing.T) {
	b1 := `{"code":"", "message":""}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		deviceID string
		snap     string
		action   string
		body     string
		data     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", "helloworld", "refresh", b1, []byte("{}")}, ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", "helloworld", "refresh,", b1, []byte("{}")}, "MOCK error put"},
		{"invalid-body", fields{""}, args{"abc", "a111", "helloworld", "refresh", "", []byte("{}")}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.SnapUpdate(tt.args.orgID, tt.args.deviceID, tt.args.snap, tt.args.action, tt.args.data)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.SnapUpdate() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}

func TestClientAdapter_SnapConfigSet(t *testing.T) {
	config := []byte(`{"title":"Hello World!"}`)
	b1 := `{"code":"", "message":""}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		deviceID string
		snap     string
		config   []byte
		body     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", "helloworld", config, b1}, ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", "helloworld", config, b1}, "MOCK error put"},
		{"invalid-body", fields{""}, args{"abc", "a111", "helloworld", config, ""}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.SnapConfigSet(tt.args.orgID, tt.args.deviceID, tt.args.snap, tt.args.config)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.SnapConfigSet() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}

func TestClientAdapter_SnapListOnDevice(t *testing.T) {
	b1 := `{"code":"", "message":""}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		deviceID string
		snap     string
		action   string
		body     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", "helloworld", "refresh", b1}, ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", "helloworld", "refresh", b1}, "MOCK error post"},
		{"invalid-body", fields{""}, args{"abc", "a111", "helloworld", "refresh", ""}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.SnapListOnDevice(tt.args.orgID, tt.args.deviceID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.SnapListOnDevice() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestClientAdapter_SnapSnapshot(t *testing.T) {
	b1 := `{"url":"", "https://upload.com/upload":""}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID    string
		deviceID string
		snap     string
		body     string
		data     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", "a111", "helloworld", b1, []byte("{}")}, ""},
		{"invalid-org", fields{""}, args{"invalid", "a111", "helloworld", b1, []byte("{}")}, "MOCK error put"},
		{"invalid-body", fields{""}, args{"abc", "a111", "helloworld", "", []byte("{}")}, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.SnapSnapshot(tt.args.orgID, tt.args.deviceID, tt.args.snap, tt.args.data)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.SnapUpdate() = %v, want %v", got.Message, tt.wantErr)
			}
		})
	}
}
