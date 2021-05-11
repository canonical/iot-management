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
	"encoding/json"
	"fmt"
	"github.com/everactive/iot-devicetwin/web"
	"github.com/everactive/iot-management/config"
	"github.com/everactive/iot-management/web/usso"
	"github.com/juju/usso/openid"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
)

var settings *config.Settings

func getSettings() *config.Settings {
	if settings == nil {
		settings, _ = config.Config("../testing/memory.yaml")
	}
	return settings
}

func sendRequest(method, url string, data io.Reader, srv *Service, username, jwtSecret string, permissions int) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, data)

	if err := createJWTWithRole(username, jwtSecret, r, permissions); err != nil {
		log.Fatalf("Error creating JWT: %v", err)
	}

	srv.Router().ServeHTTP(w, r)

	return w
}

func createJWTWithRole(username, jwtSecret string, r *http.Request, role int) error {
	sreg := map[string]string{"nickname": username, "fullname": "JJ", "email": "jj@example.com"}
	resp := openid.Response{ID: "identity", Teams: []string{}, SReg: sreg}
	jwtToken, err := usso.NewJWTToken(jwtSecret, &resp, role)
	if err != nil {
		return fmt.Errorf("error creating a JWT: %v", err)
	}
	r.Header.Set("Authorization", "Bearer "+jwtToken)
	return nil
}

func parseStandardResponse(r io.Reader) (web.StandardResponse, error) {
	// Parse the response
	result := web.StandardResponse{}
	err := json.NewDecoder(r).Decode(&result)
	return result, err
}
