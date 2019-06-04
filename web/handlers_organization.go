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
	"github.com/CanonicalLtd/iot-devicetwin/web"
	"github.com/CanonicalLtd/iot-management/domain"
	"github.com/gorilla/mux"
	"io"
	"net/http"
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

	org := domain.Organization{}
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

//// AccountsForUserHandler returns the list of accounts a user can access
//func (wb Service) AccountsForUserHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	authUser, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatUserAccountsResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	// Get the username from the parameters
//	vars := mux.Vars(r)
//	username := vars["username"]
//	if len(username) == 0 {
//		w.WriteHeader(http.StatusBadRequest)
//		formatUserAccountsResponse(false, "error-user-invalid", "Username not supplied", nil, w)
//		return
//	}
//
//	// Get the accounts that the user can access
//	accountsForUser, err := datastore.Environ.DB.ListAllowedAccounts(authUser, username)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		formatUserAccountsResponse(false, "error-accounts-json", err.Error(), nil, w)
//		return
//	}
//
//	// Get all the available accounts
//	allAccounts, err := datastore.Environ.DB.ListAllowedAccounts(authUser, "")
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		formatUserAccountsResponse(false, "error-accounts-json", err.Error(), nil, w)
//		return
//	}
//
//	userAccounts := []UserAccount{}
//	for _, a := range allAccounts {
//		found := false
//		for _, u := range accountsForUser {
//			if a.Code == u.Code {
//				found = true
//				break
//			}
//		}
//		ua := UserAccount{a, found}
//		userAccounts = append(userAccounts, ua)
//	}
//
//	// Format the model for output and return JSON response
//	w.WriteHeader(http.StatusOK)
//	formatUserAccountsResponse(true, "", "", userAccounts, w)
//
//}
//
//// AccountUpdateForUserHandler updates the account access for the user
//func (wb Service) AccountUpdateForUserHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatResponse(false, "error-auth", "", w)
//		return
//	}
//
//	vars := mux.Vars(r)
//	accountID, err := strconv.Atoi(vars["account_id"])
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound)
//		formatResponse(false, "error-account-invalid", "Cannot find an account for the ID", w)
//		return
//	}
//
//	userID, err := strconv.Atoi(vars["user_id"])
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound)
//		formatResponse(false, "error-user-invalid", "Cannot find an user for the ID", w)
//		return
//	}
//
//	// Toggle the account-user link
//	err = datastore.Environ.DB.UpdateAccountUserToggle(accountID, userID)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		formatResponse(false, "error-accounts-json", err.Error(), w)
//		return
//	}
//}
