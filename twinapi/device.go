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

package twinapi

import (
	"encoding/json"
	"path"

	"github.com/everactive/iot-devicetwin/web"
)

// DeviceList lists the devices for an account
func (a *ClientAdapter) DeviceList(orgID string) web.DevicesResponse {
	r := web.DevicesResponse{}
	p := path.Join("device", orgID)

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

// DeviceGet fetches a device for an account
func (a *ClientAdapter) DeviceGet(orgID, deviceID string) web.DeviceResponse {
	r := web.DeviceResponse{}
	p := path.Join("device", orgID, deviceID)

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

// DeviceDelete fetches a device for an account
func (a *ClientAdapter) DeviceDelete(orgID, deviceID string) web.StandardResponse {
	r := web.StandardResponse{}
	p := path.Join("device", orgID, deviceID)

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

func (a *ClientAdapter) DeviceLogs(orgID, deviceID string, body []byte) web.StandardResponse {
	r := web.StandardResponse{}
	p := path.Join("device", orgID, deviceID, "logs")

	resp, err := post(a.urlPath(p), body)
	if err != nil {
		r.Code = "DeviceLogs"
		r.Message = err.Error()
		return r
	}

	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		r.Code = "DeviceLogs"
		r.Message = err.Error()
		return r
	}

	return r
}
