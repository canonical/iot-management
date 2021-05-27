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
)

// DeviceList gets the devices a user can access for an organization
func (srv *Management) DeviceList(orgID, username string, role int) web.DevicesResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.DevicesResponse{
			StandardResponse: web.StandardResponse{
				Code:    "DevicesAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.TwinAPI.DeviceList(orgID)
}

// DeviceGet gets the device for an organization
func (srv *Management) DeviceGet(orgID, username string, role int, deviceID string) web.DeviceResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.DeviceResponse{
			StandardResponse: web.StandardResponse{
				Code:    "DeviceAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.TwinAPI.DeviceGet(orgID, deviceID)
}

// DeviceDelete deletes the device from an organization
func (srv *Management) DeviceDelete(orgID, username string, role int, deviceID string) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "DeviceAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	r1 := srv.TwinAPI.DeviceDelete(orgID, deviceID)
	r2 := srv.IdentityAPI.DeviceDelete(deviceID)

	message := fmt.Sprintf("twinapi: %s, identity: %s", r1.Message, r2.Message)
	return web.StandardResponse{Message: message}
}

func (srv *Management) DeviceLogs(orgID, username string, role int, deviceID string, body []byte) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "SnapAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.DeviceLogs(orgID, deviceID, body)
}
