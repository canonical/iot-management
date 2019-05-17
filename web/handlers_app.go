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
	"github.com/CanonicalLtd/iot-management/config"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/csrf"
)

var indexTemplate = "/static/app.html"

// Page is the page details for the web application
type Page struct {
	Title string
	Logo  string
}

// VersionResponse is the JSON response from the API Version method
type VersionResponse struct {
	Version string `json:"version"`
}

// IndexHandler is the front page of the web application
func (wb Service) IndexHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "IoT Management Service", Logo: ""}

	path := []string{".", indexTemplate}
	t, err := template.ParseFiles(strings.Join(path, ""))
	if err != nil {
		log.Printf("Error loading the application template: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// VersionHandler is the API method to return the version of the web
func (wb Service) VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	w.WriteHeader(http.StatusOK)

	response := VersionResponse{Version: config.Version}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding the version response: %v\n", err)
	}
}

// TokenHandler returns CSRF protection new token in a X-CSRF-Token response header
// This method is also used by the /authtoken endpoint to return the JWT. The method
// indicates to the UI whether OpenID user auth is enabled
func (wb Service) TokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	w.Header().Set("X-CSRF-Token", csrf.Token(r))

	// Check the JWT and return it in the authorization header, if valid
	wb.JWTCheck(w, r)

	response := VersionResponse{Version: config.Version}

	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding the token response: %v", err)
	}
}
