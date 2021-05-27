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
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/everactive/iot-devicetwin/web"
	"github.com/gorilla/mux"
)

const (
	snapRefreshURI = "refresh"
	snapEnableURI  = "enable"
	snapDisableURI = "disable"
	snapSwitchURI  = "switch"
)

// SnapListHandler fetches the list of installed snaps from the device
func (wb Service) SnapListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the device snaps
	response := wb.Manage.SnapList(vars["orgid"], user.Username, user.Role, vars["deviceid"])
	encodeResponse(response, w)
}

// SnapListOnDevice lists snaps on the device
func (wb Service) SnapListOnDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Install a snap on a device
	response := wb.Manage.SnapListOnDevice(vars["orgid"], user.Username, user.Role, vars["deviceid"])
	encodeResponse(response, w)
}

// SnapInstallHandler installs a snap on the device
func (wb Service) SnapInstallHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Install a snap on a device
	response := wb.Manage.SnapInstall(vars["orgid"], user.Username, user.Role, vars["deviceid"], vars["snap"])
	encodeResponse(response, w)
}

// SnapDeleteHandler uninstalls a snap from the device
func (wb Service) SnapDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Uninstall a snap on a device
	response := wb.Manage.SnapRemove(vars["orgid"], user.Username, user.Role, vars["deviceid"], vars["snap"])
	encodeResponse(response, w)
}

// SnapUpdateHandler updates a snap on the device
// Permitted actions are: enable, disable, refresh, or switch
func (wb Service) SnapUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		formatStandardResponse("SnapUpdate", err.Error(), w)
		return
	}

	if len(body) == 0 {
		body = []byte("{}")
	}

	defer r.Body.Close()

	vars := mux.Vars(r)
	var response web.StandardResponse

	switch vars["action"] {
	case snapEnableURI, snapDisableURI, snapRefreshURI, snapSwitchURI:
		response = wb.Manage.SnapUpdate(vars["orgid"], user.Username, user.Role, vars["deviceid"], vars["snap"], vars["action"], body)
	default:
		w.WriteHeader(http.StatusBadRequest)
		response = web.StandardResponse{Code: "SnapUpdate", Message: fmt.Sprintf("Invalid action provided: %s", vars["action"])}
	}
	encodeResponse(response, w)
}

func (wb Service) SnapSnapshotHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		formatStandardResponse("SnapPost", err.Error(), w)
		return
	}

	if len(body) == 0 {
		body = []byte("{}")
		return
	}

	defer r.Body.Close()

	vars := mux.Vars(r)

	response := wb.Manage.SnapSnapshot(vars["orgid"], user.Username, user.Role, vars["deviceid"], vars["snap"], body)

	encodeResponse(response, w)
}

// SnapConfigUpdateHandler gets the config for a snap on a device
func (wb Service) SnapConfigUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		formatStandardResponse("SnapUpdate", err.Error(), w)
		return
	}

	if len(body) == 0 {
		body = []byte("{}")
	}

	defer r.Body.Close()

	vars := mux.Vars(r)

	// Update a snap's config on a device
	response := wb.Manage.SnapConfigSet(vars["orgid"], user.Username, user.Role, vars["deviceid"], vars["snap"], body)
	encodeResponse(response, w)
}

// SnapServiceAction start/stop/restart a snap on device
func (wb Service) SnapServiceAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		formatStandardResponse("SnapUpdate", err.Error(), w)
		return
	}

	if len(body) == 0 {
		body = []byte("{}")
	}

	defer r.Body.Close()

	vars := mux.Vars(r)

	// Update a snap's config on a device
	response := wb.Manage.SnapServiceAction(vars["orgid"], user.Username, user.Role, vars["deviceid"], vars["snap"], vars["action"], body)
	encodeResponse(response, w)
}
