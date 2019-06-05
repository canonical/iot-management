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

	"github.com/CanonicalLtd/iot-management/datastore/memory"
	"github.com/CanonicalLtd/iot-management/service/manage"
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
		{"valid", "/v1/organizations", "jamesj", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/organizations", "jamesj", 0, http.StatusBadRequest, "UserAuth"},
		{"valid", "/v1/organizations", "unknown", 300, http.StatusOK, ""},
		{"invalid-user", "/v1/organizations", "invalid", 200, http.StatusBadRequest, "OrgList"},
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

func TestService_OrganizationCreateHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		data        []byte
		want        int
		wantErr     string
	}{
		{"valid", "/v1/organizations", "jamesj", 300, []byte(`{"orgid":"def","name":"Test Inc"}`), http.StatusOK, ""},
		{"invalid-duplicate", "/v1/organizations", "jamesj", 300, []byte(`{"orgid":"abc","name":"Test Inc"}`), http.StatusBadRequest, "OrgCreate"},
		{"invalid-permissions", "/v1/organizations", "jamesj", 200, []byte(`{"orgid":"def","name":"Test Inc"}`), http.StatusBadRequest, "UserAuth"},
		{"invalid-data", "/v1/organizations", "jamesj", 300, []byte(`\u1000`), http.StatusBadRequest, "OrgCreate"},
		{"invalid-data-empty", "/v1/organizations", "jamesj", 300, []byte(``), http.StatusBadRequest, "OrgCreate"},
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
				t.Errorf("Web.OrganizationCreateHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_OrganizationUpdateHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		data        []byte
		want        int
		wantErr     string
	}{
		{"valid", "/v1/organizations/abc", "jamesj", 300, []byte(`{"orgid":"abc","name":"Test Inc"}`), http.StatusOK, ""},
		{"invalid-org", "/v1/organizations/def", "jamesj", 300, []byte(`{"orgid":"def","name":"Test Inc"}`), http.StatusBadRequest, "OrgUpdate"},
		{"invalid-permissions", "/v1/organizations/abc", "jamesj", 200, []byte(`{"orgid":"abc","name":"Test Inc"}`), http.StatusBadRequest, "UserAuth"},
		{"invalid-data", "/v1/organizations/abc", "jamesj", 300, []byte(`\u1000`), http.StatusBadRequest, "OrgUpdate"},
		{"invalid-data-empty", "/v1/organizations/abc", "jamesj", 300, []byte(``), http.StatusBadRequest, "OrgUpdate"},
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
				t.Errorf("Web.OrganizationUpdateHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_OrganizationGetHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/organizations/abc", "jamesj", 300, http.StatusOK, ""},
		{"invalid-org", "/v1/organizations/invalid", "jamesj", 300, http.StatusBadRequest, "OrgGet"},
		{"invalid-permissions", "/v1/organizations/abc", "jamesj", 200, http.StatusBadRequest, "UserAuth"},
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
				t.Errorf("Web.OrganizationGetHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_OrganizationsForUserHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/users/jamesj/organizations", "jamesj", 300, http.StatusOK, ""},
		{"invalid-org", "/v1/users/invalid/organizations", "jamesj", 300, http.StatusBadRequest, "OrgList"},
		{"invalid-permissions", "/v1/users/jamesj/organizations", "jamesj", 200, http.StatusBadRequest, "UserAuth"},
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
				t.Errorf("Web.OrganizationsForUserHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_OrganizationUpdateForUserHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		username    string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/users/jamesj/organizations/abc", "jamesj", 300, http.StatusOK, ""},
		{"invalid-org", "/v1/users/invalid/organizations/abc", "jamesj", 300, http.StatusBadRequest, "UserOrg"},
		{"invalid-permissions", "/v1/users/jamesj/organizations/abc", "jamesj", 200, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("POST", tt.url, nil, wb, tt.username, wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.OrganizationsForUserHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}
