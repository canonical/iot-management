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
	"github.com/CanonicalLtd/iot-management/config"
	"github.com/CanonicalLtd/iot-management/datastore/memory"
	"github.com/CanonicalLtd/iot-management/service/manage"
	"github.com/juju/usso"
	"net/http"
	"net/url"
	"testing"
)

func TestLoginHandlerUSSORedirect(t *testing.T) {
	// Mock the services
	settings, _ := config.Config("../settings.yaml")
	db := memory.NewStore()
	m := manage.NewMockManagement(db)
	wb := NewService(settings, m)

	w := sendRequest("GET", "/login", nil, wb, wb.Settings.JwtSecret, 100)

	if w.Code != http.StatusFound {
		t.Errorf("Expected HTTP status '302', got: %v", w.Code)
	}

	u, err := url.Parse(w.Header().Get("Location"))
	if err != nil {
		t.Errorf("Error Parsing the redirect URL: %v", u)
	}

	// Check that the redirect is to the Ubuntu SSO service
	url := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	if url != usso.ProductionUbuntuSSOServer.LoginURL() {
		t.Errorf("Unexpected redirect URL: %v", url)
	}
}
