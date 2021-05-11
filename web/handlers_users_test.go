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

func TestService_UserListHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/users", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/users", 200, http.StatusBadRequest, "UserAuth"},
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
				t.Errorf("Web.UserListHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_UserCreateHandler(t *testing.T) {
	u1 := []byte(`{"username":"jane", "name":"Jane D", "email":"jd@example.com", "role":200}`)
	u2 := []byte(`{"username":"invalid", "name":"Invalid", "email":"jd@example.com", "role":200}`)
	u3 := []byte(``)
	u4 := []byte(`\u1000`)
	tests := []struct {
		name        string
		url         string
		permissions int
		data        []byte
		want        int
		wantErr     string
	}{
		{"valid", "/v1/users", 300, u1, http.StatusOK, ""},
		{"invalid-user", "/v1/users", 300, u2, http.StatusBadRequest, "UserAuth"},
		{"invalid-permissions", "/v1/users", 200, u1, http.StatusBadRequest, "UserAuth"},
		{"invalid-empty", "/v1/users", 300, u3, http.StatusBadRequest, "UserAuth"},
		{"invalid-data", "/v1/users", 300, u4, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("POST", tt.url, bytes.NewReader(tt.data), wb, "jamesj", wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.UserListHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_UserGetHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/users/jamesj", 300, http.StatusOK, ""},
		{"invalid-permissions", "/v1/users/jamesj", 200, http.StatusBadRequest, "UserAuth"},
		{"invalid-user", "/v1/users/invalid", 300, http.StatusBadRequest, "UserAuth"},
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
				t.Errorf("Web.UserListHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_UserUpdateHandler(t *testing.T) {
	u1 := []byte(`{"username":"jamesj", "name":"James Jone", "email":"jj@example.com", "role":200}`)
	u2 := []byte(`{"username":"invalid", "name":"Invalid", "email":"jd@example.com", "role":200}`)
	u3 := []byte(``)
	u4 := []byte(`\u1000`)
	tests := []struct {
		name        string
		url         string
		permissions int
		data        []byte
		want        int
		wantErr     string
	}{
		{"valid", "/v1/users/jamesj", 300, u1, http.StatusOK, ""},
		{"invalid-user", "/v1/users/invalid", 300, u2, http.StatusBadRequest, "UserUpdate"},
		{"invalid-permissions", "/v1/users/jamesj", 200, u1, http.StatusBadRequest, "UserAuth"},
		{"invalid-empty", "/v1/users/jamesj", 300, u3, http.StatusBadRequest, "UserAuth"},
		{"invalid-data", "/v1/users/jamesj", 300, u4, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("PUT", tt.url, bytes.NewReader(tt.data), wb, "jamesj", wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.UserUpdateHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}

func TestService_UserDeleteHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		permissions int
		want        int
		wantErr     string
	}{
		{"valid", "/v1/users/jamesj", 300, http.StatusOK, ""},
		{"invalid-user", "/v1/users/invalid", 300, http.StatusBadRequest, "UserDelete"},
		{"invalid-permissions", "/v1/users/jamesj", 200, http.StatusBadRequest, "UserAuth"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("DELETE", tt.url, nil, wb, "jamesj", wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}

			resp, err := parseStandardResponse(w.Body)
			if err != nil {
				t.Errorf("Error parsing response: %v", err)
			}
			if resp.Code != tt.wantErr {
				t.Errorf("Web.UserUpdateHandler() got = %v, want %v", resp.Code, tt.wantErr)
			}
		})
	}
}
