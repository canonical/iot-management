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
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Router returns the application route handler for administrating the application
func (wb Service) Router() *mux.Router {
	// Start the web web router
	router := mux.NewRouter()

	// API routes: CSRF token and auth token
	router.Handle("/v1/token", Middleware(http.HandlerFunc(wb.TokenHandler))).Methods("GET")
	router.Handle("/v1/authtoken", Middleware(http.HandlerFunc(wb.TokenHandler))).Methods("GET")
	router.Handle("/v1/version", Middleware(http.HandlerFunc(wb.VersionHandler))).Methods("GET")

	// API routes: clients
	//router.Handle("/v1/{account_code}/clients/{endpoint}", Middleware(http.HandlerFunc(DeviceGetByNameHandler))).Methods("GET")

	//// API routes: client subsection
	//router.Handle("/v1/{account_code}/clients/{endpoint}/device", Middleware(http.HandlerFunc(DeviceGetByNameHandler))).Methods("GET")
	//router.Handle("/v1/{account_code}/devices/{endpoint}", Middleware(http.HandlerFunc(DeviceGetByNameHandler))).Methods("GET")

	// API routes: devices
	router.Handle("/v1/{orgid}/devices", Middleware(http.HandlerFunc(wb.DevicesListHandler))).Methods("GET")
	router.Handle("/v1/{orgid}/devices/{deviceid}", Middleware(http.HandlerFunc(wb.DeviceGetHandler))).Methods("GET")
	//router.Handle("/v1/{account_code}/devices", Middleware(http.HandlerFunc(DeviceCreateHandler))).Methods("POST")
	//router.Handle("/v1/{account_code}/devices/{id}", Middleware(http.HandlerFunc(DeviceUpdateHandler))).Methods("PUT")

	// API routes: groups
	router.Handle("/v1/{orgid}/groups", Middleware(http.HandlerFunc(wb.GroupListHandler))).Methods("GET")

	// API routes: snap functionality
	router.Handle("/v1/device/{orgid}/{deviceid}/snaps", Middleware(http.HandlerFunc(wb.SnapListHandler))).Methods("GET")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}", Middleware(http.HandlerFunc(wb.SnapInstallHandler))).Methods("POST")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}", Middleware(http.HandlerFunc(wb.SnapDeleteHandler))).Methods("DELETE")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}/settings", Middleware(http.HandlerFunc(wb.SnapConfigUpdateHandler))).Methods("PUT")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}/{action}", Middleware(http.HandlerFunc(wb.SnapUpdateHandler))).Methods("PUT")

	// API routes: store functionality
	router.Handle("/v1/store/snaps/{snapName}", Middleware(http.HandlerFunc(wb.StoreSearchHandler))).Methods("GET")

	// API routes: accounts
	router.Handle("/v1/organizations", Middleware(http.HandlerFunc(wb.OrganizationListHandler))).Methods("GET")
	router.Handle("/v1/organizations/{id}", Middleware(http.HandlerFunc(wb.OrganizationGetHandler))).Methods("GET")
	router.Handle("/v1/organizations/{id}", Middleware(http.HandlerFunc(wb.OrganizationUpdateHandler))).Methods("PUT")
	router.Handle("/v1/organizations", Middleware(http.HandlerFunc(wb.OrganizationCreateHandler))).Methods("POST")

	//// API routes: users
	router.Handle("/v1/users", Middleware(http.HandlerFunc(wb.UserListHandler))).Methods("GET")
	//router.Handle("/v1/users", Middleware(http.HandlerFunc(UserCreateHandler))).Methods("POST")
	//router.Handle("/v1/users/{id}", Middleware(http.HandlerFunc(UserGetHandler))).Methods("GET")
	//router.Handle("/v1/users/{id}", Middleware(http.HandlerFunc(UserUpdateHandler))).Methods("PUT")
	//router.Handle("/v1/users/{username}/accounts", Middleware(http.HandlerFunc(AccountsForUserHandler))).Methods("GET")
	//router.Handle("/v1/users/{user_id}/accounts/{account_id}", Middleware(http.HandlerFunc(AccountUpdateForUserHandler))).Methods("POST")

	// OpenID routes: using Ubuntu SSO
	router.Handle("/login", Middleware(http.HandlerFunc(wb.LoginHandler)))
	router.Handle("/logout", Middleware(http.HandlerFunc(wb.LogoutHandler)))

	// Web application routes
	path := []string{".", "/static/"}
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(strings.Join(path, ""))))
	router.PathPrefix("/static/").Handler(fs)
	router.PathPrefix("/notfound").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.PathPrefix("/accounts").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.PathPrefix("/devices").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.PathPrefix("/groups").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.PathPrefix("/register").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.PathPrefix("/users").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.Handle("/", Middleware(http.HandlerFunc(wb.IndexHandler))).Methods("GET")

	return router
}
