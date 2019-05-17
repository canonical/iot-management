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
	"fmt"
	"github.com/CanonicalLtd/iot-management/datastore/memory"
	"github.com/CanonicalLtd/iot-management/service/manage"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestService_StoreSearchHandler(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		permissions int
		want        int
	}{
		{"valid", "/v1/store/snaps/helloworld", 300, http.StatusOK},
		{"invalid-response", "/v1/store/snaps/invalid", 300, http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockGET(`[{}]`)
			db := memory.NewStore()
			wb := NewService(getSettings(), manage.NewMockManagement(db))
			w := sendRequest("GET", tt.url, nil, wb, wb.Settings.JwtSecret, tt.permissions)
			if w.Code != tt.want {
				t.Errorf("Expected HTTP status '%d', got: %v", tt.want, w.Code)
			}
		})
	}
}

func mockGET(body string) {
	// Mock the HTTP methods
	get = func(p string) (*http.Response, error) {
		log.Println("---", p)
		if strings.Contains(p, "invalid") {
			return nil, fmt.Errorf("MOCK error get")
		}
		return &http.Response{
			Body: ioutil.NopCloser(strings.NewReader(body)),
		}, nil
	}
}
