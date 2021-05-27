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
	"github.com/everactive/iot-devicetwin/web"
	idweb "github.com/everactive/iot-identity/web"
	"github.com/everactive/iot-management/config"
	"github.com/everactive/iot-management/datastore"
	"github.com/everactive/iot-management/domain"
	"github.com/everactive/iot-management/identityapi"
	"github.com/everactive/iot-management/twinapi"
	"github.com/juju/usso/openid"
)

// Manage interface for the service
type Manage interface {
	OpenIDNonceStore() openid.NonceStore
	CreateUser(user domain.User) error
	GetUser(username string) (domain.User, error)
	UserList() ([]domain.User, error)
	UserUpdate(user domain.User) error
	UserDelete(username string) error

	RegDeviceList(orgID, username string, role int) idweb.DevicesResponse
	RegisterDevice(orgID, username string, role int, body []byte) idweb.RegisterResponse
	RegDeviceGet(orgID, username string, role int, deviceID string) idweb.EnrollResponse
	RegDeviceUpdate(orgID, username string, role int, deviceID string, body []byte) idweb.StandardResponse

	DeviceList(orgID, username string, role int) web.DevicesResponse
	DeviceGet(orgID, username string, role int, deviceID string) web.DeviceResponse
	DeviceDelete(orgID, username string, role int, deviceID string) web.StandardResponse
	DeviceLogs(orgID, username string, role int, deviceID string, body []byte) web.StandardResponse
	ActionList(orgID, username string, role int, deviceID string) web.ActionsResponse

	SnapSnapshot(orgID, username string, role int, deviceID, snap string, body []byte) web.StandardResponse
	SnapList(orgID, username string, role int, deviceID string) web.SnapsResponse
	SnapListOnDevice(orgID, username string, role int, deviceID string) web.StandardResponse
	SnapInstall(orgID, username string, role int, deviceID, snap string) web.StandardResponse
	SnapRemove(orgID, username string, role int, deviceID, snap string) web.StandardResponse
	SnapUpdate(orgID, username string, role int, deviceID, snap, action string, body []byte) web.StandardResponse
	SnapConfigSet(orgID, username string, role int, deviceID, snap string, config []byte) web.StandardResponse
	SnapServiceAction(orgID, username string, role int, deviceID, snap, action string, body []byte) web.StandardResponse

	GroupList(orgID, username string, role int) web.GroupsResponse
	GroupCreate(orgID, username string, role int, body []byte) web.StandardResponse
	GroupDevices(orgID, username string, role int, name string) web.DevicesResponse
	GroupExcludedDevices(orgID, username string, role int, name string) web.DevicesResponse
	GroupDeviceLink(orgID, username string, role int, name, deviceID string) web.StandardResponse
	GroupDeviceUnlink(orgID, username string, role int, name, deviceID string) web.StandardResponse

	OrganizationsForUser(username string) ([]domain.Organization, error)
	OrganizationForUserToggle(orgID, username string) error
	OrganizationGet(orgID string) (domain.Organization, error)
	OrganizationCreate(org domain.OrganizationCreate) error
	OrganizationUpdate(org domain.Organization) error
}

// Management implementation of the management service use cases
type Management struct {
	Settings    *config.Settings
	DB          datastore.DataStore
	TwinAPI     twinapi.Client
	IdentityAPI identityapi.Client
}

// NewManagement creates an implementation of the management use cases
func NewManagement(settings *config.Settings, db datastore.DataStore, api twinapi.Client, id identityapi.Client) *Management {
	return &Management{
		Settings:    settings,
		DB:          db,
		TwinAPI:     api,
		IdentityAPI: id,
	}
}
