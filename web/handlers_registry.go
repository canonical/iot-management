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
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/everactive/iot-identity/service"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// RegDeviceList is the API method to list the registered devices
func (wb Service) RegDeviceList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the devices
	response := wb.Manage.RegDeviceList(vars["orgid"], user.Username, user.Role)
	if len(response.Code) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	_ = encodeResponse(response, w)
}

// RegDeviceGet is the API method to get a registered device
func (wb Service) RegDeviceGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the device
	response := wb.Manage.RegDeviceGet(vars["orgid"], user.Username, user.Role, vars["device"])
	if len(response.Code) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	_ = encodeResponse(response, w)
}

// RegDeviceGetDownload provides the download of the device data
func (wb Service) RegDeviceGetDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Fetch the device
	response := wb.Manage.RegDeviceGet(vars["orgid"], user.Username, user.Role, vars["device"])
	if len(response.Code) > 0 {
		formatStandardResponse(response.Code, response.Message, w)
		return
	}

	// Set the download headers and body for the file
	w.Header().Set("Content-Disposition", "attachment; filename=devicedata-"+response.Enrollment.ID)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	// Decode the base64 file
	data, err := base64.StdEncoding.DecodeString(response.Enrollment.DeviceData)
	if err != nil {
		log.Println("Error decoding the device data:", err)
		formatStandardResponse("DeviceData", "Error decoding the file", w)
		return
	}

	io.Copy(w, bytes.NewReader(data))
}

// RegDeviceUpdate is the API method to update a registered device status
func (wb Service) RegDeviceUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		formatStandardResponse("RegDevice", "error reading the request", w)
		return
	}

	// Get the devices
	response := wb.Manage.RegDeviceUpdate(vars["orgid"], user.Username, user.Role, vars["device"], b)
	if len(response.Code) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	_ = encodeResponse(response, w)
}

// RegisterDevice registers a new device with the Identity service
func (wb Service) RegisterDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsAdminAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	// Read the body of the request
	vars := mux.Vars(r)
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		formatStandardResponse("RegDevice", "error reading the request", w)
		return
	}

	// Decode the body and check we have a valid organization ID
	req, err := decodeDeviceRequest(b)
	if err != nil {
		formatStandardResponse("RegDevice", "error decoding the request", w)
		return
	}
	if req.OrganizationID != vars["orgid"] {
		formatStandardResponse("RegDevice", "the request organization ID is invalid", w)
		return
	}

	// Register the devices
	response := wb.Manage.RegisterDevice(vars["orgid"], user.Username, user.Role, b)
	if len(response.Code) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}
	_ = encodeResponse(response, w)
}

func decodeDeviceRequest(body []byte) (*service.RegisterDeviceRequest, error) {
	// Decode the JSON body
	dev := service.RegisterDeviceRequest{}
	err := json.Unmarshal(body, &dev)
	return &dev, err
}
