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
	"path"

	"github.com/everactive/iot-identity/web"
)

// RegDeviceList lists the devices for an account
func (a *ClientAdapter) RegDeviceList(orgID string) web.DevicesResponse {
	r := web.DevicesResponse{}
	p := path.Join("devices", orgID)

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

// RegisterDevice registers a new device
func (a *ClientAdapter) RegisterDevice(body []byte) web.RegisterResponse {
	r := web.RegisterResponse{}
	p := path.Join("device")

	resp, err := post(a.urlPath(p), body)
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

// RegDeviceGet fetches a device registration
func (a *ClientAdapter) RegDeviceGet(orgID, deviceID string) web.EnrollResponse {
	r := web.EnrollResponse{}
	p := path.Join("devices", orgID, deviceID)

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

// RegDeviceUpdate updates a device registration
func (a *ClientAdapter) RegDeviceUpdate(orgID, deviceID string, body []byte) web.StandardResponse {
	r := web.StandardResponse{}
	p := path.Join("devices", orgID, deviceID)

	resp, err := put(a.urlPath(p), body)
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

// DeviceDelete deletes device registration
func (a *ClientAdapter) DeviceDelete(deviceID string) web.StandardResponse {
	r := web.StandardResponse{}
	p := path.Join("device", deviceID)

	resp, err := delete(a.urlPath(p))
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
