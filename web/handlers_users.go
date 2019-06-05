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
	"encoding/json"
	"github.com/CanonicalLtd/iot-identity/web"
	"github.com/CanonicalLtd/iot-management/domain"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// UsersResponse defines the response to list users
type UsersResponse struct {
	web.StandardResponse
	Users []domain.User `json:"users"`
}

// UserResponse defines the response to get a user
type UserResponse struct {
	web.StandardResponse
	User domain.User `json:"user"`
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

// UserGetHandler is the API method to fetch a user
func (wb Service) UserGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the users
	user, err := wb.Manage.GetUser(vars["username"])
	if err != nil {
		formatStandardResponse("UserAuth", err.Error(), w)
		return
	}

	_ = encodeResponse(UserResponse{web.StandardResponse{}, user}, w)
}

// UserCreateHandler handles user creation
func (wb Service) UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	user := domain.User{}
	err = json.NewDecoder(r.Body).Decode(&user)
	switch {
	// Check we have some data
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		formatStandardResponse("UserAuth", "No user data supplied", w)
		return
	// Check for parsing errors
	case err != nil:
		w.WriteHeader(http.StatusBadRequest)
		formatStandardResponse("UserAuth", err.Error(), w)
		return
	}

	// Create the user
	err = wb.Manage.CreateUser(user)
	if err != nil {
		formatStandardResponse("UserAuth", err.Error(), w)
		return
	}

	formatStandardResponse("", "", w)
}

// UserUpdateHandler handles user update
func (wb Service) UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	user := domain.User{}
	err = json.NewDecoder(r.Body).Decode(&user)
	switch {
	// Check we have some data
	case err == io.EOF:
		w.WriteHeader(http.StatusBadRequest)
		formatStandardResponse("UserAuth", "No user data supplied", w)
		return
	// Check for parsing errors
	case err != nil:
		w.WriteHeader(http.StatusBadRequest)
		formatStandardResponse("UserAuth", err.Error(), w)
		return
	}

	// Create the user
	err = wb.Manage.UserUpdate(user)
	if err != nil {
		formatStandardResponse("UserUpdate", err.Error(), w)
		return
	}

	formatStandardResponse("", "", w)
}

// UserDeleteHandler handles user deletion
func (wb Service) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	if err := wb.Manage.UserDelete(vars["username"]); err != nil {
		formatStandardResponse("UserDelete", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}
