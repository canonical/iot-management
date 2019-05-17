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
