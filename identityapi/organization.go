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
	"encoding/json"
	"github.com/everactive/iot-identity/web"
	"path"
)

// RegisterOrganization registers a new organization
func (a *ClientAdapter) RegisterOrganization(body []byte) web.RegisterResponse {
	r := web.RegisterResponse{}
	p := path.Join("organization")

	resp, err := post(a.urlPath(p), body)
	if err != nil {
		r.Message = err.Error()
		return r
	}

	// Parse the response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		r.Message = err.Error()
	}

	return r
}

// RegOrganizationList lists the organizations for an account
func (a *ClientAdapter) RegOrganizationList() web.OrganizationsResponse {
	r := web.OrganizationsResponse{}
	p := path.Join("organizations")

	resp, err := get(a.urlPath(p))
	if err != nil {
		r.StandardResponse.Message = err.Error()
		return r
	}

	// Parse the response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		r.StandardResponse.Message = err.Error()
	}

	return r
}
