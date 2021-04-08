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

import "github.com/canonical/iot-devicetwin/web"

// GroupList lists the device groups
func (srv *Management) GroupList(orgID, username string, role int) web.GroupsResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.GroupsResponse{
			StandardResponse: web.StandardResponse{
				Code:    "GroupAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.TwinAPI.GroupList(orgID)
}

// GroupCreate creates a device group
func (srv *Management) GroupCreate(orgID, username string, role int, body []byte) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "GroupAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.GroupCreate(orgID, body)
}

// GroupDevices lists the devices for a groups
func (srv *Management) GroupDevices(orgID, username string, role int, name string) web.DevicesResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.DevicesResponse{
			StandardResponse: web.StandardResponse{
				Code:    "GroupAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.TwinAPI.GroupDevices(orgID, name)
}

// GroupExcludedDevices lists the devices for a groups
func (srv *Management) GroupExcludedDevices(orgID, username string, role int, name string) web.DevicesResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.DevicesResponse{
			StandardResponse: web.StandardResponse{
				Code:    "GroupAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.TwinAPI.GroupExcludedDevices(orgID, name)
}

// GroupDeviceLink links a device to a group
func (srv *Management) GroupDeviceLink(orgID, username string, role int, name, deviceID string) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "GroupAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.GroupDeviceLink(orgID, name, deviceID)
}

// GroupDeviceUnlink unlinks a device from a group
func (srv *Management) GroupDeviceUnlink(orgID, username string, role int, name, deviceID string) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "GroupAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.GroupDeviceUnlink(orgID, name, deviceID)
}
