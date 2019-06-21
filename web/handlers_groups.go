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

package web

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// GroupListHandler is the API method to list the groups
func (wb Service) GroupListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the devices
	response := wb.Manage.GroupList(vars["orgid"], user.Username, user.Role)
	encodeResponse(response, w)
}

// GroupCreateHandler is the API method to create a group
func (wb Service) GroupCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	// Read the body of the request
	vars := mux.Vars(r)
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		formatStandardResponse("GroupCreate", "error reading the request", w)
		return
	}

	// Get the devices
	response := wb.Manage.GroupCreate(vars["orgid"], user.Username, user.Role, b)
	encodeResponse(response, w)
}

// GroupDevicesHandler is the API method to list the devices for a group
func (wb Service) GroupDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the devices
	response := wb.Manage.GroupDevices(vars["orgid"], user.Username, user.Role, vars["name"])
	encodeResponse(response, w)
}

// GroupExcludedDevicesHandler is the API method to list the devices not in a group
func (wb Service) GroupExcludedDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the devices
	response := wb.Manage.GroupExcludedDevices(vars["orgid"], user.Username, user.Role, vars["name"])
	encodeResponse(response, w)
}

// GroupDeviceLinkHandler is the API method to link a device to a group
func (wb Service) GroupDeviceLinkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the devices
	response := wb.Manage.GroupDeviceLink(vars["orgid"], user.Username, user.Role, vars["name"], vars["device"])
	encodeResponse(response, w)
}

// GroupDeviceUnlinkHandler is the API method to unlink a device from a group
func (wb Service) GroupDeviceUnlinkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the devices
	response := wb.Manage.GroupDeviceUnlink(vars["orgid"], user.Username, user.Role, vars["name"], vars["device"])
	encodeResponse(response, w)
}
