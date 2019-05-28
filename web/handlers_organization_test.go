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
	"github.com/CanonicalLtd/iot-management/datastore/memory"
	"github.com/CanonicalLtd/iot-management/service/manage"
	"net/http"
	"testing"
)

func TestService_OrganizationListHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/accounts", "jamesj", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/accounts", "jamesj", 0, http.StatusOK, "UserAuth"},
		{"valid", "/v1/accounts", "unknown", 300, http.StatusOK, ""},
		{"invalid-user", "/v1/accounts", "invalid", 200, http.StatusOK, "OrgList"},
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
				t.Errorf("Web.OrganizationListHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}
