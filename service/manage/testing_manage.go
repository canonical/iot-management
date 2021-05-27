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
	"fmt"

	"github.com/everactive/iot-devicetwin/web"
	iddomain "github.com/everactive/iot-identity/domain"
	idweb "github.com/everactive/iot-identity/web"
	"github.com/everactive/iot-management/datastore"
	"github.com/everactive/iot-management/domain"
	"github.com/everactive/iot-management/twinapi"
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
	if username != "jamesj" {
		return domain.User{}, fmt.Errorf("User not found: %v", username)
	}

	return domain.User{Username: "jamesj", Name: "JJ", Role: 300}, nil
}

// UserList mocks fetching users
func (m *MockManage) UserList() ([]domain.User, error) {
	return []domain.User{{Username: "jamesj", Name: "JJ", Role: 300}}, nil
}

// UserDelete mocks removing a user
func (m *MockManage) UserDelete(username string) error {
	if username == "invalid" {
		return fmt.Errorf("MOCK error delete")
	}
	return nil
}

// CreateUser mocks creating a user
func (m *MockManage) CreateUser(user domain.User) error {
	if user.Username == "invalid" {
		return fmt.Errorf("MOCK error create")
	}
	return nil
}

// UserUpdate mocks updating a user
func (m *MockManage) UserUpdate(user domain.User) error {
	if user.Username == "invalid" {
		return fmt.Errorf("MOCK error update")
	}
	return nil
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

// ActionList mocks fetching action list
func (m *MockManage) ActionList(orgID, username string, role int, deviceID string) web.ActionsResponse {
	twin := twinapi.NewMockClient("")
	return twin.ActionList(orgID, deviceID)
}

// SnapList mocks listing snaps
func (m *MockManage) SnapList(orgID, username string, role int, deviceID string) web.SnapsResponse {
	twin := twinapi.NewMockClient("")
	return twin.SnapList(orgID, deviceID)
}

// SnapListOnDevice mocks listing snaps
func (m *MockManage) SnapListOnDevice(orgID, username string, role int, deviceID string) web.StandardResponse {
	return web.StandardResponse{}
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
func (m *MockManage) SnapUpdate(orgID, username string, role int, deviceID, snap, action string, body []byte) web.StandardResponse {
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

// GroupCreate mocks creating a group
func (m *MockManage) GroupCreate(orgID, username string, role int, body []byte) web.StandardResponse {
	twin := twinapi.NewMockClient("")
	return twin.GroupCreate(orgID, body)
}

// GroupDevices mocks listing devices for a group
func (m *MockManage) GroupDevices(orgID, username string, role int, name string) web.DevicesResponse {
	twin := twinapi.NewMockClient("")
	return twin.GroupDevices(orgID, name)
}

// GroupExcludedDevices mocks listing devices not in a group
func (m *MockManage) GroupExcludedDevices(orgID, username string, role int, name string) web.DevicesResponse {
	twin := twinapi.NewMockClient("")
	return twin.GroupExcludedDevices(orgID, name)
}

// GroupDeviceLink mocks linking a device to a group
func (m *MockManage) GroupDeviceLink(orgID, username string, role int, name, deviceID string) web.StandardResponse {
	twin := twinapi.NewMockClient("")
	return twin.GroupDeviceLink(orgID, name, deviceID)
}

// GroupDeviceUnlink mocks unlinking a device from a group
func (m *MockManage) GroupDeviceUnlink(orgID, username string, role int, name, deviceID string) web.StandardResponse {
	twin := twinapi.NewMockClient("")
	return twin.GroupDeviceUnlink(orgID, name, deviceID)
}

// OrganizationsForUser mocks organizations for a user
func (m *MockManage) OrganizationsForUser(username string) ([]domain.Organization, error) {
	if username == "invalid" {
		return nil, fmt.Errorf("MOCK error user org")
	}
	if username != "jamesj" {
		return []domain.Organization{}, nil
	}
	return []domain.Organization{{OrganizationID: "abc", Name: "Example Org"}}, nil
}

// OrganizationGet mocks fetching an organization
func (m *MockManage) OrganizationGet(orgID string) (domain.Organization, error) {
	if orgID == "invalid" {
		return domain.Organization{}, fmt.Errorf("MOCK error get")
	}
	if orgID != "abc" {
		return domain.Organization{}, nil
	}
	return domain.Organization{OrganizationID: "abc", Name: "Example Org"}, nil
}

// OrganizationCreate mocks creating an organization
func (m *MockManage) OrganizationCreate(org domain.OrganizationCreate) error {
	return nil
}

// OrganizationUpdate mocks updating an organization
func (m *MockManage) OrganizationUpdate(org domain.Organization) error {
	if org.OrganizationID != "abc" {
		return fmt.Errorf("MOCK error update")
	}
	return nil
}

// OrganizationForUserToggle mocks toggling user access to an organization
func (m *MockManage) OrganizationForUserToggle(orgID, username string) error {
	if orgID != "abc" || username != "jamesj" {
		return fmt.Errorf("MOCK error toggle")
	}
	return nil
}

// RegDeviceList mocks listing registered devices
func (m *MockManage) RegDeviceList(orgID, username string, role int) idweb.DevicesResponse {
	if orgID == "invalid" || role == 100 {
		return idweb.DevicesResponse{
			StandardResponse: idweb.StandardResponse{Code: "RegDeviceAuth", Message: "MOCK error devices"},
			Devices:          nil,
		}
	}
	return idweb.DevicesResponse{
		StandardResponse: idweb.StandardResponse{},
		Devices:          []iddomain.Enrollment{},
	}
}

// RegisterDevice mocks registering a device
func (m *MockManage) RegisterDevice(orgID, username string, role int, body []byte) idweb.RegisterResponse {
	if orgID == "invalid" || role == 100 {
		return idweb.RegisterResponse{
			StandardResponse: idweb.StandardResponse{Code: "RegDeviceAuth", Message: "MOCK error register"},
			ID:               "",
		}
	}
	return idweb.RegisterResponse{
		StandardResponse: idweb.StandardResponse{},
		ID:               "d444",
	}
}

// RegDeviceGet mocks fetching a registered device
func (m *MockManage) RegDeviceGet(orgID, username string, role int, deviceID string) idweb.EnrollResponse {
	if deviceID == "invalid" || role == 0 {
		return idweb.EnrollResponse{
			StandardResponse: idweb.StandardResponse{Code: "RegDeviceAuth", Message: "MOCK error get"},
			Enrollment:       iddomain.Enrollment{},
		}
	}
	return idweb.EnrollResponse{
		StandardResponse: idweb.StandardResponse{},
		Enrollment:       iddomain.Enrollment{},
	}
}

// RegDeviceUpdate mocks updating a registered device
func (m *MockManage) RegDeviceUpdate(orgID, username string, role int, deviceID string, body []byte) idweb.StandardResponse {
	if deviceID == "invalid" {
		return idweb.StandardResponse{Code: "RegDeviceUpdate", Message: "MOCK error update"}
	}
	return idweb.StandardResponse{}
}

// SnapShot mocks create a snap snapshot
func (m *MockManage) SnapShot(orgID, username string, role int, deviceID, snap, action string, body []byte) web.StandardResponse {
	return web.StandardResponse{}
}

// DeviceLogs mocks create a device logs
func (m *MockManage) DeviceLogs(orgID, username string, role int, deviceID, snap, action string, body []byte) web.StandardResponse {
	return web.StandardResponse{}
}
