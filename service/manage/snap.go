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
)

// SnapList lists the snaps for a device
func (srv *Management) SnapList(orgID, username string, role int, deviceID string) web.SnapsResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.SnapsResponse{
			StandardResponse: web.StandardResponse{
				Code:    "SnapsAuth",
				Message: "the user does not have permissions for the organization",
			},
		}
	}

	return srv.TwinAPI.SnapList(orgID, deviceID)
}

// SnapListOnDevice lists snaps on a device
func (srv *Management) SnapListOnDevice(orgID, username string, role int, deviceID string) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "SnapAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.SnapListOnDevice(orgID, deviceID)
}

// SnapInstall installs a snap on a device
func (srv *Management) SnapInstall(orgID, username string, role int, deviceID, snap string) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "SnapAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.SnapInstall(orgID, deviceID, snap)
}

// SnapRemove uninstalls a snap on a device
func (srv *Management) SnapRemove(orgID, username string, role int, deviceID, snap string) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "SnapAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.SnapRemove(orgID, deviceID, snap)
}

// SnapUpdate enables/disables/refreshes/swtich a snap on a device
func (srv *Management) SnapUpdate(orgID, username string, role int, deviceID, snap, action string, body []byte) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "SnapAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.SnapUpdate(orgID, deviceID, snap, action, body)
}

// SnapConfigSet updates a snap config on a device
func (srv *Management) SnapConfigSet(orgID, username string, role int, deviceID, snap string, config []byte) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "SnapAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.SnapConfigSet(orgID, deviceID, snap, config)
}

func (srv *Management) SnapServiceAction(orgID, username string, role int, deviceID, snap, action string, body []byte) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "SnapAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.SnapServiceAction(orgID, deviceID, snap, action, body)
}

func (srv *Management) SnapSnapshot(orgID, username string, role int, deviceID, snap string, body []byte) web.StandardResponse {
	hasAccess := srv.DB.OrgUserAccess(orgID, username, role)
	if !hasAccess {
		return web.StandardResponse{
			Code:    "SnapAuth",
			Message: "the user does not have permissions for the organization",
		}
	}

	return srv.TwinAPI.SnapSnapshot(orgID, deviceID, snap, body)
}
