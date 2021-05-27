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
	"io/ioutil"
	"net/http"

	dtwin "github.com/everactive/iot-devicetwin/web"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func formatStandardResponse(errorCode, message string, w http.ResponseWriter) {
	response := dtwin.StandardResponse{Code: errorCode, Message: message}
	if len(errorCode) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	_ = encodeResponse(response, w)
}

// DevicesListHandler is the API method to list the registered devices
func (wb Service) DevicesListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the devices
	response := wb.Manage.DeviceList(vars["orgid"], user.Username, user.Role)
	log.Tracef("Sending response back: %+v", response)
	_ = encodeResponse(response, w)
}

// DeviceDeleteHandler is the API method to delete a registered device
func (wb Service) DeviceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsAdminAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Delete the device
	response := wb.Manage.DeviceDelete(vars["orgid"], user.Username, user.Role, vars["deviceid"])
	_ = encodeResponse(response, w)
}

// DeviceGetHandler is the API method to get a registered device
func (wb Service) DeviceGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsAdminAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the device
	response := wb.Manage.DeviceGet(vars["orgid"], user.Username, user.Role, vars["deviceid"])
	_ = encodeResponse(response, w)
}

// ActionListHandler is the API method to get actions for a device
func (wb Service) ActionListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsAdminAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the device
	response := wb.Manage.ActionList(vars["orgid"], user.Username, user.Role, vars["deviceid"])
	_ = encodeResponse(response, w)
}

// DeviceLogsHandler is the API method to get logs for a device
func (wb Service) DeviceLogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsAdminAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		formatStandardResponse("DeviceLogs", err.Error(), w)
		return
	}

	if len(body) == 0 {
		body = []byte("{}")
	}

	defer r.Body.Close()

	vars := mux.Vars(r)

	// Get the device
	response := wb.Manage.DeviceLogs(vars["orgid"], user.Username, user.Role, vars["deviceid"], body)
	_ = encodeResponse(response, w)
}
