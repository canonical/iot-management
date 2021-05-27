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

	"github.com/everactive/iot-management/datastore/memory"
	"github.com/everactive/iot-management/service/manage"
)

func TestService_DeviceHandlers(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/abc/devices", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/abc/devices", 0, http.StatusBadRequest, "UserAuth"},

		{"valid", "/v1/abc/devices/a111", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/abc/devices/a111", 0, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("GET", tt.url, nil, wb, "jamesj", wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.DeviceHandlers() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_ActionListHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/abc/devices/a111/actions", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/abc/devices/a111/actions", 0, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("GET", tt.url, nil, wb, "jamesj", wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.ActionListHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_Workflow(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		url         string
		body        []byte
		permissions int
		want        int
		wantErr     string
	}{
		{"send-logs-valid", "POST", "/v1/device/abc/a111/logs", []byte("{}"), 300, http.StatusOK, ""},
		{"send-logs-invalid-permissions", "POST", "/v1/device/abc/a111/logs", []byte("{}"), 0, http.StatusBadRequest, "UserAuth"},
		{"send-logs-valid-empty", "POST", "/v1/device/abc/a111/logs", nil, 300, http.StatusBadRequest, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest(tt.method, tt.url, bytes.NewReader(tt.body), wb, "jamesj", wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.SnapInstallHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}
