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

	// API routes: registered devices
	router.Handle("/v1/{orgid}/register/devices", Middleware(http.HandlerFunc(wb.RegDeviceList))).Methods("GET")
	router.Handle("/v1/{orgid}/register/devices", Middleware(http.HandlerFunc(wb.RegisterDevice))).Methods("POST")
	router.Handle("/v1/{orgid}/register/devices/{device}", Middleware(http.HandlerFunc(wb.RegDeviceGet))).Methods("GET")
	router.Handle("/v1/{orgid}/register/devices/{device}", Middleware(http.HandlerFunc(wb.RegDeviceUpdate))).Methods("PUT")
	router.Handle("/v1/{orgid}/register/devices/{device}/download", Middleware(http.HandlerFunc(wb.RegDeviceGetDownload))).Methods("GET")

	// API routes: devices
	router.Handle("/v1/{orgid}/devices", Middleware(http.HandlerFunc(wb.DevicesListHandler))).Methods("GET")
	router.Handle("/v1/{orgid}/devices/{deviceid}", Middleware(http.HandlerFunc(wb.DeviceGetHandler))).Methods("GET")
	router.Handle("/v1/{orgid}/devices/{deviceid}/actions", Middleware(http.HandlerFunc(wb.ActionListHandler))).Methods("GET")
	router.Handle("/v1/{orgid}/devices/{deviceid}", Middleware(http.HandlerFunc(wb.DeviceDeleteHandler))).Methods("DELETE")
	router.Handle("/v1/{orgid}/devices/{deviceid}/logs", Middleware(http.HandlerFunc(wb.DeviceLogsHandler))).Methods("POST")

	// API routes: groups
	router.Handle("/v1/{orgid}/groups", Middleware(http.HandlerFunc(wb.GroupListHandler))).Methods("GET")
	router.Handle("/v1/{orgid}/groups/{name}/{device}", Middleware(http.HandlerFunc(wb.GroupDeviceLinkHandler))).Methods("POST")
	router.Handle("/v1/{orgid}/groups/{name}/{device}", Middleware(http.HandlerFunc(wb.GroupDeviceUnlinkHandler))).Methods("DELETE")
	router.Handle("/v1/{orgid}/groups/{name}/devices", Middleware(http.HandlerFunc(wb.GroupDevicesHandler))).Methods("GET")
	router.Handle("/v1/{orgid}/groups/{name}/devices/excluded", Middleware(http.HandlerFunc(wb.GroupExcludedDevicesHandler))).Methods("GET")
	router.Handle("/v1/{orgid}/groups", Middleware(http.HandlerFunc(wb.GroupCreateHandler))).Methods("POST")

	// API routes: snap functionality
	router.Handle("/v1/device/{orgid}/{deviceid}/snaps", Middleware(http.HandlerFunc(wb.SnapListHandler))).Methods("GET")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/list", Middleware(http.HandlerFunc(wb.SnapListOnDevice))).Methods("POST")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}", Middleware(http.HandlerFunc(wb.SnapInstallHandler))).Methods("POST")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}", Middleware(http.HandlerFunc(wb.SnapDeleteHandler))).Methods("DELETE")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}/settings", Middleware(http.HandlerFunc(wb.SnapConfigUpdateHandler))).Methods("PUT")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/services/{snap}/{action}", Middleware(http.HandlerFunc(wb.SnapServiceAction))).Methods("POST")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}/{action}", Middleware(http.HandlerFunc(wb.SnapUpdateHandler))).Methods("PUT")
	router.Handle("/v1/snaps/{orgid}/{deviceid}/{snap}/snapshot", Middleware(http.HandlerFunc(wb.SnapSnapshotHandler))).Methods("POST")

	// API routes: store functionality
	router.Handle("/v1/store/snaps/{snapName}", Middleware(http.HandlerFunc(wb.StoreSearchHandler))).Methods("GET")

	// API routes: accounts
	router.Handle("/v1/organizations", Middleware(http.HandlerFunc(wb.OrganizationListHandler))).Methods("GET")
	router.Handle("/v1/organizations/{id}", Middleware(http.HandlerFunc(wb.OrganizationGetHandler))).Methods("GET")
	router.Handle("/v1/organizations/{id}", Middleware(http.HandlerFunc(wb.OrganizationUpdateHandler))).Methods("PUT")
	router.Handle("/v1/organizations", Middleware(http.HandlerFunc(wb.OrganizationCreateHandler))).Methods("POST")

	// API routes: users
	router.Handle("/v1/users", Middleware(http.HandlerFunc(wb.UserListHandler))).Methods("GET")
	router.Handle("/v1/users", Middleware(http.HandlerFunc(wb.UserCreateHandler))).Methods("POST")
	router.Handle("/v1/users/{username}", Middleware(http.HandlerFunc(wb.UserGetHandler))).Methods("GET")
	router.Handle("/v1/users/{username}", Middleware(http.HandlerFunc(wb.UserUpdateHandler))).Methods("PUT")
	router.Handle("/v1/users/{username}", Middleware(http.HandlerFunc(wb.UserDeleteHandler))).Methods("DELETE")
	router.Handle("/v1/users/{username}/organizations", Middleware(http.HandlerFunc(wb.OrganizationsForUserHandler))).Methods("GET")
	router.Handle("/v1/users/{username}/organizations/{orgid}", Middleware(http.HandlerFunc(wb.OrganizationUpdateForUserHandler))).Methods("POST")

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
	router.PathPrefix("/actions").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.PathPrefix("/groups").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.PathPrefix("/register").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.PathPrefix("/users").Handler(Middleware(http.HandlerFunc(wb.IndexHandler)))
	router.Handle("/", Middleware(http.HandlerFunc(wb.IndexHandler))).Methods("GET")

	return router
}
