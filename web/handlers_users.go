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
	"github.com/CanonicalLtd/iot-identity/web"
	"github.com/CanonicalLtd/iot-management/domain"
	"net/http"
)

// UsersResponse defines the response to list users
type UsersResponse struct {
	web.StandardResponse
	Users []domain.User `json:"users"`
}

// UserListHandler is the API method to list the users
func (wb Service) UserListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	// Get the users
	users, err := wb.Manage.UserList()
	if err != nil {
		formatStandardResponse("UserAuth", err.Error(), w)
		return
	}

	_ = encodeResponse(UsersResponse{web.StandardResponse{}, users}, w)
}

//// UsersHandler is the API method to list the users
//func (wb Service) UsersHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUsersResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	users, err := datastore.Environ.DB.ListUsers()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		formatUsersResponse(false, "error-fetch-users", err.Error(), nil, w)
//		return
//	}
//
//	// Return successful JSON response with the list of users
//	w.WriteHeader(http.StatusOK)
//	formatUsersResponse(true, "", "", users, w)
//}
//
//// UserCreateHandler handles user creation
//func (wb Service) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUsersResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	user := datastore.User{}
//	err = json.NewDecoder(r.Body).Decode(&user)
//	switch {
//	// Check we have some data
//	case err == io.EOF:
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-user-data", "No user data supplied", user, w)
//		return
//		// Check for parsing errors
//	case err != nil:
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-data-json", err.Error(), user, w)
//		return
//	}
//
//	userID, err := datastore.Environ.DB.CreateUser(user)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-user-invalid", err.Error(), user, w)
//		return
//	}
//	user.ID = userID
//
//	formatUserResponse(true, "", "", user, w)
//}
//
//// UserGetHandler is the API method to retrieve user info
//func (wb Service) UserGetHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUserResponse(false, "error-auth", "", datastore.User{}, w)
//		return
//	}
//
//	vars := mux.Vars(r)
//	id, err := strconv.ParseInt(vars["id"], 10, 64)
//
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound)
//		errorMessage := fmt.Sprintf("%v", vars)
//		formatUserResponse(false, "error-user-invalid", errorMessage, datastore.User{}, w)
//		return
//	}
//
//	user, err := datastore.Environ.DB.GetUser(id)
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound)
//		errorMessage := fmt.Sprintf("User ID: %d.", id)
//		formatUserResponse(false, "error-get-user", errorMessage, datastore.User{ID: id}, w)
//		return
//	}
//
//	w.WriteHeader(http.StatusOK)
//	formatUserResponse(true, "", "", user, w)
//}
//
//// UserUpdateHandler handles user update
//func (wb Service) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUsersResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	user := datastore.User{}
//	err = json.NewDecoder(r.Body).Decode(&user)
//	switch {
//	// Check we have some data
//	case err == io.EOF:
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-user-data", "No user data supplied", user, w)
//		return
//		// Check for parsing errors
//	case err != nil:
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-data-json", err.Error(), user, w)
//		return
//	}
//
//	err = datastore.Environ.DB.UpdateUser(user)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserResponse(false, "error-user-invalid", err.Error(), user, w)
//		return
//	}
//
//	formatUserResponse(true, "", "", user, w)
//}
