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
	"github.com/CanonicalLtd/iot-management/config"
	"github.com/CanonicalLtd/iot-management/datastore"
	"github.com/CanonicalLtd/iot-management/domain"
	"github.com/CanonicalLtd/iot-management/twinapi"
	"github.com/juju/usso/openid"
)

// Manage interface for the service
type Manage interface {
	OpenIDNonceStore() openid.NonceStore
	GetUser(username string) (domain.User, error)
	UserList() ([]domain.User, error)

	DeviceList(orgID, username string, role int) web.DevicesResponse
	DeviceGet(orgID, username string, role int, deviceID string) web.DeviceResponse

	SnapList(orgID, username string, role int, deviceID string) web.SnapsResponse
	SnapInstall(orgID, username string, role int, deviceID, snap string) web.StandardResponse
	SnapRemove(orgID, username string, role int, deviceID, snap string) web.StandardResponse
	SnapUpdate(orgID, username string, role int, deviceID, snap, action string) web.StandardResponse
	SnapConfigSet(orgID, username string, role int, deviceID, snap string, config []byte) web.StandardResponse

	GroupList(orgID, username string, role int) web.GroupsResponse
}

// Management implementation of the management service use cases
type Management struct {
	Settings *config.Settings
	DB       datastore.DataStore
	TwinAPI  twinapi.Client
}

// NewManagement creates an implementation of the management use cases
func NewManagement(settings *config.Settings, db datastore.DataStore, api twinapi.Client) *Management {
	return &Management{
		Settings: settings,
		DB:       db,
		TwinAPI:  api,
	}
}
