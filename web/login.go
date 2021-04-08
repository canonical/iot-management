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
	"html/template"
	"log"
	"net/http"

	"github.com/canonical/iot-management/datastore"
	"github.com/canonical/iot-management/web/usso"
)

// LoginHandler processes the login for Ubuntu SSO
func (wb Service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get the openid nonce store
	nonce := wb.Manage.OpenIDNonceStore()

	// Call the openid login handler
	resp, req, username, err := usso.LoginHandler(wb.Settings, nonce, w, r)
	if err != nil {
		log.Printf("Error verifying the OpenID response: %v\n", err)
		replyHTTPError(w, http.StatusBadRequest, err)
		return
	}
	if req != nil {
		// Redirect is handled by the SSO handler
		return
	}

	// Check that the user is registered
	user, err := wb.Manage.GetUser(username)
	if err != nil {
		// Cannot find the user, so redirect to the login page
		log.Printf("Error retrieving user from datastore: %v\n", err)
		http.Redirect(w, r, "/notfound", http.StatusTemporaryRedirect)
		return
	}

	// Verify role value is valid
	if user.Role != datastore.Standard && user.Role != datastore.Admin && user.Role != datastore.Superuser {
		log.Printf("Role obtained from database for user %v has not a valid value: %v\n", username, user.Role)
		http.Redirect(w, r, "/notfound", http.StatusTemporaryRedirect)
		return
	}

	// Build the JWT
	jwtToken, err := usso.NewJWTToken(wb.Settings.JwtSecret, resp, user.Role)
	if err != nil {
		// Unexpected that this should occur, so leave the detailed response
		log.Printf("Error creating the JWT: %v", err)
		replyHTTPError(w, http.StatusBadRequest, err)
		return
	}

	// Set a cookie with the JWT
	usso.AddJWTCookie(jwtToken, w)

	// Redirect to the homepage with the JWT
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// LogoutHandler logs the user out by removing the cookie and the JWT authorization header
func (wb Service) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	usso.LogoutHandler(w, r)
}

func replyHTTPError(w http.ResponseWriter, returnCode int, err error) {
	w.Header().Set("ContentType", "text/html")
	w.WriteHeader(returnCode)
	errorTemplate.Execute(w, err)
}

var errorTemplate = template.Must(template.New("failure").Parse(`<html>
 <head><title>Login Error</title></head>
 <body>{{.}}</body>
 </html>
 `))
