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
	"github.com/CanonicalLtd/iot-devicetwin/web"
	"path"
)

// GroupList lists the device groups
func (a *ClientAdapter) GroupList(orgID string) web.GroupsResponse {
	r := web.GroupsResponse{}
	p := path.Join("group", orgID)

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

// GroupCreate creates a device group
func (a *ClientAdapter) GroupCreate(orgID string, body []byte) web.StandardResponse {
	r := web.StandardResponse{}
	p := path.Join("group", orgID)

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

// GroupDevices lists the devices for a group
func (a *ClientAdapter) GroupDevices(orgID, name string) web.DevicesResponse {
	r := web.DevicesResponse{}
	p := path.Join("group", orgID, name, "devices")

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

// GroupExcludedDevices lists the devices for a group
func (a *ClientAdapter) GroupExcludedDevices(orgID, name string) web.DevicesResponse {
	r := web.DevicesResponse{}
	p := path.Join("group", orgID, name, "devices", "excluded")

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

// GroupDeviceLink links a device with a group
func (a *ClientAdapter) GroupDeviceLink(orgID, name, deviceID string) web.StandardResponse {
	r := web.StandardResponse{}
	p := path.Join("group", orgID, name, deviceID)

	resp, err := post(a.urlPath(p), []byte(""))
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

// GroupDeviceUnlink unlinks a device from a group
func (a *ClientAdapter) GroupDeviceUnlink(orgID, name, deviceID string) web.StandardResponse {
	r := web.StandardResponse{}
	p := path.Join("group", orgID, name, deviceID)

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
