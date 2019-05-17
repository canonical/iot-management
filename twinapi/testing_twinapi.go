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
	"github.com/CanonicalLtd/iot-devicetwin/domain"
	"github.com/CanonicalLtd/iot-devicetwin/web"
	"io/ioutil"
	"net/http"
	"strings"
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
		Devices: []domain.Device{
			{OrganizationID: "abc", DeviceID: "a111", Brand: "example", Model: "drone-1000", SerialNumber: "DR1000A111", DeviceKey: "AAAAAAAAA", StoreID: "example-store"},
			{OrganizationID: "abc", DeviceID: "b222", Brand: "example", Model: "drone-1000", SerialNumber: "DR1000B222", DeviceKey: "BBBBBBBBB", StoreID: "example-store"},
			{OrganizationID: "abc", DeviceID: "c333", Brand: "canonical", Model: "ubuntu-core-18-amd64", SerialNumber: "d75f7300-abbf-4c11-bf0a-8b7103038490", DeviceKey: "CCCCCCCCC"},
		},
	}
}

// DeviceGet mocks the device fetch
func (m *MockClient) DeviceGet(orgID, deviceID string) web.DeviceResponse {
	return web.DeviceResponse{
		StandardResponse: web.StandardResponse{},
		Device:           domain.Device{OrganizationID: "abc", DeviceID: "b222", Brand: "example", Model: "drone-1000", SerialNumber: "DR1000B222", DeviceKey: "BBBBBBBBB", StoreID: "example-store"},
	}
}

// SnapList mocks the snap list
func (m *MockClient) SnapList(orgID, deviceID string) web.SnapsResponse {
	return web.SnapsResponse{
		StandardResponse: web.StandardResponse{},
		Snaps:            []domain.DeviceSnap{{Name: "example-snap", InstalledSize: 2000, Status: "active"}},
	}
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
func (m *MockClient) SnapUpdate(orgID, deviceID, snap, action string) web.StandardResponse {
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
