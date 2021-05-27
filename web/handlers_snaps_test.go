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

func TestService_SnapListHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/device/abc/a111/snaps", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/device/abc/a111/snaps", 0, http.StatusBadRequest, "UserAuth"},
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
				t.Errorf("Web.SnapListHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_SnapWorkflow(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		url         string
		body        []byte
		permissions int
		want        int
		wantErr     string
	}{
		{"list-valid", "POST", "/v1/snaps/abc/a111/list", nil, 300, http.StatusOK, ""},
		{"list-invalid-permissions", "POST", "/v1/snaps/abc/a111/list", nil, 0, http.StatusBadRequest, "UserAuth"},
		{"install-valid", "POST", "/v1/snaps/abc/a111/helloworld", nil, 300, http.StatusOK, ""},
		{"install-invalid-permissions", "POST", "/v1/snaps/abc/a111/helloworld", nil, 0, http.StatusBadRequest, "UserAuth"},
		{"delete-valid", "DELETE", "/v1/snaps/abc/a111/helloworld", nil, 300, http.StatusOK, ""},
		{"delete-invalid-permissions", "DELETE", "/v1/snaps/abc/a111/helloworld", nil, 0, http.StatusBadRequest, "UserAuth"},
		{"update-valid-refresh", "PUT", "/v1/snaps/abc/a111/helloworld/refresh", nil, 300, http.StatusOK, ""},
		{"update-valid-enable", "PUT", "/v1/snaps/abc/a111/helloworld/enable", nil, 300, http.StatusOK, ""},
		{"update-valid-disable", "PUT", "/v1/snaps/abc/a111/helloworld/disable", nil, 300, http.StatusOK, ""},
		{"update-valid-switch", "PUT", "/v1/snaps/abc/a111/helloworld/switch", []byte("{}"), 300, http.StatusOK, ""},
		{"update-action-invalid", "PUT", "/v1/snaps/abc/a111/helloworld/invalid", nil, 300, http.StatusBadRequest, "SnapUpdate"},
		{"update-invalid-permissions", "PUT", "/v1/snaps/abc/a111/helloworld/refresh", nil, 0, http.StatusBadRequest, "UserAuth"},
		{"config-valid", "PUT", "/v1/snaps/abc/a111/helloworld/settings", []byte("{}"), 300, http.StatusOK, ""},
		{"config-valid-empty", "PUT", "/v1/snaps/abc/a111/helloworld/settings", nil, 300, http.StatusOK, ""},
		{"config-invalid-permissions", "PUT", "/v1/snaps/abc/a111/helloworld/settings", []byte("{}"), 0, http.StatusBadRequest, "UserAuth"},
		{"send-snapshot-valid", "POST", "/v1/snaps/abc/a111/helloworld/snapshot", []byte("{}"), 300, http.StatusOK, ""},
		{"send-snapshot-invalid-permissions", "POST", "/v1/snaps/abc/a111/helloworld/snapshot", []byte("{}"), 0, http.StatusBadRequest, "UserAuth"},
		{"send-snapshot-valid-empty", "POST", "/v1/snaps/abc/a111/helloworld/snapshot", nil, 300, http.StatusBadRequest, ""},
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
