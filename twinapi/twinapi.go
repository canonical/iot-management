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
	"bytes"
	"net/http"
	"net/url"
	"path"

	"github.com/everactive/iot-devicetwin/web"
)

// Client is a client for the device twin API
type Client interface {
	DeviceList(orgID string) web.DevicesResponse
	DeviceGet(orgID, deviceID string) web.DeviceResponse
	DeviceDelete(orgID, deviceID string) web.StandardResponse
	DeviceLogs(orgID, deviceID string, body []byte) web.StandardResponse
	ActionList(orgID, deviceID string) web.ActionsResponse
	SnapList(orgID, deviceID string) web.SnapsResponse

	SnapSnapshot(orgID, deviceID, snap string, body []byte) web.StandardResponse
	SnapListOnDevice(orgID, deviceID string) web.StandardResponse
	SnapInstall(orgID, deviceID, snap string) web.StandardResponse
	SnapRemove(orgID, deviceID, snap string) web.StandardResponse
	SnapUpdate(orgID, deviceID, snap, action string, body []byte) web.StandardResponse
	SnapConfigSet(orgID, deviceID, snap string, config []byte) web.StandardResponse
	SnapServiceAction(orgID, deviceID, snap, action string, body []byte) web.StandardResponse

	GroupList(orgID string) web.GroupsResponse
	GroupCreate(orgID string, body []byte) web.StandardResponse
	GroupDevices(orgID, name string) web.DevicesResponse
	GroupExcludedDevices(orgID, name string) web.DevicesResponse
	GroupDeviceLink(orgID, name, deviceID string) web.StandardResponse
	GroupDeviceUnlink(orgID, name, deviceID string) web.StandardResponse
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

var delete = func(p string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, p, nil)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
