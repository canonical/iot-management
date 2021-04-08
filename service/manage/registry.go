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
	"github.com/canonical/iot-identity/web"
	idweb "github.com/canonical/iot-identity/web"
)

// RegDeviceList gets the registered devices a user can access for an organization
func (srv *Management) RegDeviceList(orgID, username string, role int) web.DevicesResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.DevicesResponse{
			StandardResponse: web.StandardResponse{
				Code:    "RegDevicesAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.IdentityAPI.RegDeviceList(orgID)
}

// RegisterDevice registers a new device
func (srv *Management) RegisterDevice(orgID, username string, role int, body []byte) web.RegisterResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.RegisterResponse{
			StandardResponse: web.StandardResponse{
				Code:    "RegDeviceAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.IdentityAPI.RegisterDevice(body)
}

// RegDeviceGet fetches a device registration
func (srv *Management) RegDeviceGet(orgID, username string, role int, deviceID string) web.EnrollResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.EnrollResponse{
			StandardResponse: web.StandardResponse{
				Code:    "RegDeviceAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.IdentityAPI.RegDeviceGet(orgID, deviceID)
}

// RegDeviceUpdate updates a device registration
func (srv *Management) RegDeviceUpdate(orgID, username string, role int, deviceID string, body []byte) idweb.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return idweb.StandardResponse{
			Code:    "RegDeviceAuth",
			Message: "the user does not have permissions for the organization",
		}
	}
	return srv.IdentityAPI.RegDeviceUpdate(orgID, deviceID, body)
}
