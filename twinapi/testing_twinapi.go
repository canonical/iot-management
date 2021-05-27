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
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/everactive/iot-devicetwin/domain"
	"github.com/everactive/iot-devicetwin/pkg/messages"
	"github.com/everactive/iot-devicetwin/web"
)

// MockClient mocks the device twin client
type MockClient struct{}

func mockHTTP(body string) {
	// Mock the HTTP methods
	get = func(p string) (*http.Response, error) {
		if strings.Contains(p, "invalid") {
			return nil, fmt.Errorf("MOCK error get")
		}
		return &http.Response{
			Body: ioutil.NopCloser(strings.NewReader(body)),
		}, nil
	}
	post = func(p string, data []byte) (*http.Response, error) {
		if strings.Contains(p, "invalid") {
			return nil, fmt.Errorf("MOCK error post")
		}
		return &http.Response{
			Body: ioutil.NopCloser(strings.NewReader(body)),
		}, nil
	}
	put = func(p string, data []byte) (*http.Response, error) {
		if strings.Contains(p, "invalid") {
			return nil, fmt.Errorf("MOCK error put")
		}
		return &http.Response{
			Body: ioutil.NopCloser(strings.NewReader(body)),
		}, nil
	}
	delete = func(p string) (*http.Response, error) {
		if strings.Contains(p, "invalid") {
			return nil, fmt.Errorf("MOCK error delete")
		}
		return &http.Response{
			Body: ioutil.NopCloser(strings.NewReader(body)),
		}, nil
	}
}

// NewMockClient creates a mock client instance
func NewMockClient(body string) *MockClient {
	mockHTTP(body)
	return &MockClient{}
}

// DeviceList mocks the device list call
func (m *MockClient) DeviceList(orgID string) web.DevicesResponse {
	return web.DevicesResponse{
		StandardResponse: web.StandardResponse{},
		Devices: []messages.Device{
			{OrgId: "abc", DeviceId: "a111", Brand: "example", Model: "drone-1000", Serial: "DR1000A111", DeviceKey: "AAAAAAAAA", Store: "example-store"},
			{OrgId: "abc", DeviceId: "b222", Brand: "example", Model: "drone-1000", Serial: "DR1000B222", DeviceKey: "BBBBBBBBB", Store: "example-store"},
			{OrgId: "abc", DeviceId: "c333", Brand: "canonical", Model: "ubuntu-core-18-amd64", Serial: "d75f7300-abbf-4c11-bf0a-8b7103038490", DeviceKey: "CCCCCCCCC"},
		},
	}
}

// DeviceGet mocks the device fetch
func (m *MockClient) DeviceGet(orgID, deviceID string) web.DeviceResponse {
	return web.DeviceResponse{
		StandardResponse: web.StandardResponse{},
		Device:           messages.Device{OrgId: "abc", DeviceId: "b222", Brand: "example", Model: "drone-1000", Serial: "DR1000B222", DeviceKey: "BBBBBBBBB", Store: "example-store"},
	}
}

// ActionList mocks the action list
func (m *MockClient) ActionList(orgID, deviceID string) web.ActionsResponse {
	return web.ActionsResponse{
		StandardResponse: web.StandardResponse{},
		Actions: []domain.Action{
			{OrganizationID: "abc", DeviceID: "b222", Action: "list", Status: "complete"},
		},
	}
}

// SnapList mocks the snap list
func (m *MockClient) SnapList(orgID, deviceID string) web.SnapsResponse {
	return web.SnapsResponse{
		StandardResponse: web.StandardResponse{},
		Snaps:            []messages.DeviceSnap{{Name: "example-snap", InstalledSize: 2000, Status: "active"}},
	}
}

// SnapListOnDevice mocks the snap list
func (m *MockClient) SnapListOnDevice(orgID, deviceID string) web.StandardResponse {
	return web.StandardResponse{}
}

// SnapInstall mocks a snap installation
func (m *MockClient) SnapInstall(orgID, deviceID, snap string) web.StandardResponse {
	return web.StandardResponse{}
}

// SnapRemove mocks a snap removal
func (m *MockClient) SnapRemove(orgID, deviceID, snap string) web.StandardResponse {
	return web.StandardResponse{}
}

// SnapUpdate mocks a snap update request
func (m *MockClient) SnapUpdate(orgID, deviceID, snap, action string, body []byte) web.StandardResponse {
	return web.StandardResponse{}
}

// SnapConfigSet mocks a snap config update
func (m *MockClient) SnapConfigSet(orgID, deviceID, snap string, config []byte) web.StandardResponse {
	return web.StandardResponse{}
}

// GroupList mocks listing device groups
func (m *MockClient) GroupList(orgID string) web.GroupsResponse {
	return web.GroupsResponse{
		StandardResponse: web.StandardResponse{},
		Groups:           []domain.Group{{OrganizationID: "abc", Name: "workshop"}},
	}
}

// GroupDevices mocks listing devices for a groups
func (m *MockClient) GroupDevices(orgID, name string) web.DevicesResponse {
	return web.DevicesResponse{
		StandardResponse: web.StandardResponse{},
		Devices: []messages.Device{
			{OrgId: "abc", DeviceId: "a111", Brand: "example", Model: "drone-1000", Serial: "DR1000A111", DeviceKey: "AAAAAAAAA", Store: "example-store"},
		},
	}
}

// GroupExcludedDevices mocks listing devices not in a groups
func (m *MockClient) GroupExcludedDevices(orgID, name string) web.DevicesResponse {
	return web.DevicesResponse{
		StandardResponse: web.StandardResponse{},
		Devices: []messages.Device{
			{OrgId: "abc", DeviceId: "b222", Brand: "example", Model: "drone-1000", Serial: "DR1000B222", DeviceKey: "BBBBBBBBB", Store: "example-store"},
			{OrgId: "abc", DeviceId: "c333", Brand: "canonical", Model: "ubuntu-core-18-amd64", Serial: "d75f7300-abbf-4c11-bf0a-8b7103038490", DeviceKey: "CCCCCCCCC"},
		},
	}
}

// GroupCreate mocks creating a device groups
func (m *MockClient) GroupCreate(orgID string, body []byte) web.StandardResponse {
	return web.StandardResponse{}
}

// GroupDeviceLink mocks linking a device to a group
func (m *MockClient) GroupDeviceLink(orgID, name, deviceID string) web.StandardResponse {
	if orgID == "invalid" || deviceID == "invalid" {
		return web.StandardResponse{Code: "GroupDevice", Message: "MOCK error link"}
	}
	return web.StandardResponse{}
}

// GroupDeviceUnlink mocks unlinking a device from a group
func (m *MockClient) GroupDeviceUnlink(orgID, name, deviceID string) web.StandardResponse {
	if orgID == "invalid" || deviceID == "invalid" {
		return web.StandardResponse{Code: "GroupDevice", Message: "MOCK error unlink"}
	}
	return web.StandardResponse{}
}

// SnapUpdate mocks a snap update request
func (m *MockClient) SnapSnapshot(orgID, deviceID, snap, body []byte) web.StandardResponse {
	return web.StandardResponse{}
}

// DeviceLogs mocks a device create logs request
func (m *MockClient) DeviceLogs(orgID, deviceID, body []byte) web.StandardResponse {
	return web.StandardResponse{}
}
