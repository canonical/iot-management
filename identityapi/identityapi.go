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

package identityapi

import (
	"bytes"
	"github.com/CanonicalLtd/iot-identity/web"
	"net/http"
	"net/url"
	"path"
)

// Client is a client for the identity API
type Client interface {
	RegDeviceList(orgID string) web.DevicesResponse
	RegisterDevice(body []byte) web.RegisterResponse
	RegDeviceGet(orgID, deviceID string) web.EnrollResponse
	RegDeviceUpdate(orgID, deviceID string, body []byte) web.StandardResponse
	RegisterOrganization(body []byte) web.RegisterResponse
	RegOrganizationList() web.OrganizationsResponse
}

// ClientAdapter adapts our expectations to device twin API
type ClientAdapter struct {
	URL string
}

var adapter *ClientAdapter

// NewClientAdapter creates an adapter to access the device twin service
func NewClientAdapter(u string) (*ClientAdapter, error) {
	if adapter == nil {
		adapter = &ClientAdapter{URL: u}
	}
	return adapter, nil
}

func (a *ClientAdapter) urlPath(p string) string {
	u, _ := url.Parse(a.URL)
	u.Path = path.Join(u.Path, p)
	return u.String()
}

var get = func(p string) (*http.Response, error) {
	return http.Get(p)
}

var post = func(p string, data []byte) (*http.Response, error) {
	return http.Post(p, "application/json", bytes.NewReader(data))
}

var put = func(p string, data []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, p, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
