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

package manage

import (
	"github.com/CanonicalLtd/iot-devicetwin/web"
	"github.com/CanonicalLtd/iot-management/datastore"
	"github.com/CanonicalLtd/iot-management/domain"
	"github.com/CanonicalLtd/iot-management/twinapi"
	"github.com/juju/usso/openid"
)

// MockManage mocks the management service
type MockManage struct {
	DB datastore.DataStore
}

// NewMockManagement creates an implementation of the management use cases
func NewMockManagement(db datastore.DataStore) *MockManage {
	return &MockManage{
		DB: db,
	}
}

// OpenIDNonceStore mocks the openid nonce store
func (m *MockManage) OpenIDNonceStore() openid.NonceStore {
	return m.DB.OpenIDNonceStore()
}

// GetUser mocks fetching a user
func (m *MockManage) GetUser(username string) (domain.User, error) {
	panic("implement me")
}

// DeviceList mocks fetching devices
func (m *MockManage) DeviceList(orgID, username string, role int) web.DevicesResponse {
	twin := twinapi.NewMockClient("")
	return twin.DeviceList(orgID)
}

// DeviceGet mocks fetching a device
func (m *MockManage) DeviceGet(orgID, username string, role int, deviceID string) web.DeviceResponse {
	twin := twinapi.NewMockClient("")
	return twin.DeviceGet(orgID, deviceID)
}

// SnapList mocks listing snaps
func (m *MockManage) SnapList(orgID, username string, role int, deviceID string) web.SnapsResponse {
	twin := twinapi.NewMockClient("")
	return twin.SnapList(orgID, deviceID)
}

// SnapInstall mocks installing a snap
func (m *MockManage) SnapInstall(orgID, username string, role int, deviceID, snap string) web.StandardResponse {
	return web.StandardResponse{}
}

// SnapRemove mocks uninstalling a snap
func (m *MockManage) SnapRemove(orgID, username string, role int, deviceID, snap string) web.StandardResponse {
	return web.StandardResponse{}
}

// SnapUpdate mocks updating a snap
func (m *MockManage) SnapUpdate(orgID, username string, role int, deviceID, snap, action string) web.StandardResponse {
	return web.StandardResponse{}
}

// SnapConfigSet mocks updating a snap config
func (m *MockManage) SnapConfigSet(orgID, username string, role int, deviceID, snap string, config []byte) web.StandardResponse {
	return web.StandardResponse{}
}

// GroupList mocks listing groups
func (m *MockManage) GroupList(orgID, username string, role int) web.GroupsResponse {
	twin := twinapi.NewMockClient("")
	return twin.GroupList(orgID)
}
