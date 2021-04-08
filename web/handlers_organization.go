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
	"io"
	"net/http"

	"github.com/canonical/iot-devicetwin/web"
	"github.com/canonical/iot-management/domain"
	"github.com/gorilla/mux"
)

// OrganizationsResponse defines the response to list users
type OrganizationsResponse struct {
	web.StandardResponse
	Organizations []domain.Organization `json:"organizations"`
}

// OrganizationResponse defines the response to list users
type OrganizationResponse struct {
	web.StandardResponse
	Organization domain.Organization `json:"organization"`
}

// UserOrganization defines an organization and whether it is selected for a user
type UserOrganization struct {
	domain.Organization
	Selected bool `json:"selected"`
}

// UserOrganizationsResponse defines the response to list users
type UserOrganizationsResponse struct {
	web.StandardResponse
	Organizations []UserOrganization `json:"organizations"`
}

func formatOrganizationsResponse(orgs []domain.Organization, w http.ResponseWriter) {
	response := OrganizationsResponse{Organizations: orgs}
	_ = encodeResponse(response, w)
}

func formatOrganizationResponse(org domain.Organization, w http.ResponseWriter) {
	response := OrganizationResponse{Organization: org}
	_ = encodeResponse(response, w)
}

// OrganizationListHandler returns the list of accounts
func (wb Service) OrganizationListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	orgs, err := wb.Manage.OrganizationsForUser(user.Username)
	if err != nil {
		formatStandardResponse("OrgList", err.Error(), w)
		return
	}
	formatOrganizationsResponse(orgs, w)
}

// OrganizationGetHandler fetches an organization
func (wb Service) OrganizationGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	org, err := wb.Manage.OrganizationGet(vars["id"])
	if err != nil {
		formatStandardResponse("OrgGet", err.Error(), w)
		return
	}

	formatOrganizationResponse(org, w)
}

// OrganizationCreateHandler creates a new organization
func (wb Service) OrganizationCreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	org := domain.OrganizationCreate{}
	err = json.NewDecoder(r.Body).Decode(&org)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("OrgCreate", "No organization data supplied", w)
		return
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("OrgCreate", err.Error(), w)
		return
	}

	if err = wb.Manage.OrganizationCreate(org); err != nil {
		formatStandardResponse("OrgCreate", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}

// OrganizationUpdateHandler updates an organization
func (wb Service) OrganizationUpdateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	org := domain.Organization{}
	err = json.NewDecoder(r.Body).Decode(&org)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("OrgUpdate", "No organization data supplied", w)
		return
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("OrgUpdate", err.Error(), w)
		return
	}

	if err = wb.Manage.OrganizationUpdate(org); err != nil {
		formatStandardResponse("OrgUpdate", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}

// OrganizationsForUserHandler fetches the organizations for a user
func (wb Service) OrganizationsForUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	user, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	// Get the organization a user can access
	userOrgs, err := wb.Manage.OrganizationsForUser(vars["username"])
	if err != nil {
		formatStandardResponse("OrgList", err.Error(), w)
		return
	}

	// Get all the organizations
	allOrgs, err := wb.Manage.OrganizationsForUser(user.Username)
	if err != nil {
		formatStandardResponse("OrgList", err.Error(), w)
		return
	}

	oo := []UserOrganization{}
	for _, o := range allOrgs {
		found := false
		for _, u := range userOrgs {
			if o.OrganizationID == u.OrganizationID {
				found = true
				break
			}
		}
		oo = append(oo, UserOrganization{o, found})
	}

	_ = encodeResponse(UserOrganizationsResponse{Organizations: oo}, w)
}

// OrganizationUpdateForUserHandler fetches the organizations for a user
func (wb Service) OrganizationUpdateForUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)
	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
	if err != nil {
		formatStandardResponse("UserAuth", "", w)
		return
	}

	vars := mux.Vars(r)

	if err := wb.Manage.OrganizationForUserToggle(vars["orgid"], vars["username"]); err != nil {
		formatStandardResponse("UserOrg", "", w)
		return
	}
	formatStandardResponse("", "", w)
}
