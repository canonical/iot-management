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

package web

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/canonical/iot-management/datastore/memory"
	"github.com/canonical/iot-management/service/manage"
)

func TestService_RegDeviceList(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/abc/register/devices", "jamesj", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/abc/register/devices", "jamesj", 0, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("GET", tt.url, nil, wb, tt.username, wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.RegDeviceList() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_RegDeviceGet(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/abc/register/devices/a111", "jamesj", 300, http.StatusOK, ""},
		{"invalid-org", "/v1/abc/register/devices/invalid", "jamesj", 300, http.StatusBadRequest, "RegDeviceAuth"},
		{"invalid-permissions", "/v1/abc/register/devices/a111", "jamesj", 0, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("GET", tt.url, nil, wb, tt.username, wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.RegDeviceGet() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_RegisterDevice(t *testing.T) {
	d1 := []byte(`{"orgid":"abc", "brand":"deviceinc", "model":"A1000", "serial":"d1234"}`)
	tests := []struct {
		name        string
		url         string
		data        []byte
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/abc/register/devices", d1, "jamesj", 300, http.StatusOK, ""},
		{"invalid-org", "/v1/bbb/register/devices", d1, "jamesj", 300, http.StatusBadRequest, "RegDevice"},
		{"invalid-permissions", "/v1/abc/register/devices", d1, "jamesj", 0, http.StatusBadRequest, "UserAuth"},
		{"invalid-data", "/v1/abc/register/devices", []byte(`\u1000`), "jamesj", 300, http.StatusBadRequest, "RegDevice"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("POST", tt.url, bytes.NewReader(tt.data), wb, tt.username, wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.RegisterDevice() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_RegDeviceUpdate(t *testing.T) {
	d1 := []byte(`{"status":3}`)
	tests := []struct {
		name        string
		url         string
		data        []byte
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/abc/register/devices/a111", d1, "jamesj", 300, http.StatusOK, ""},
		{"invalid-device", "/v1/abc/register/devices/invalid", d1, "jamesj", 300, http.StatusBadRequest, "RegDeviceUpdate"},
		{"invalid-permissions", "/v1/abc/register/devices/a111", d1, "jamesj", 0, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("PUT", tt.url, bytes.NewReader(tt.data), wb, tt.username, wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.RegDeviceUpdate() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_RegDeviceGetDownload(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/abc/register/devices/a111/download", "jamesj", 300, http.StatusOK, ""},
		{"invalid-org", "/v1/abc/register/devices/invalid/download", "jamesj", 300, http.StatusBadRequest, "RegDeviceAuth"},
		{"invalid-permissions", "/v1/abc/register/devices/a111/download", "jamesj", 0, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("GET", tt.url, nil, wb, tt.username, wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}
		})
	}
}
