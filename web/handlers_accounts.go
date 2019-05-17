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

//// AccountsResponse is the JSON response from the API Accounts method
//type AccountsResponse struct {
//	Success      bool                `json:"success"`
//	ErrorCode    string              `json:"error_code"`
//	ErrorMessage string              `json:"message"`
//	Accounts     []datastore.Account `json:"accounts"`
//}
//
//// AccountResponse is the JSON response from the API Accounts method
//type AccountResponse struct {
//	Success      bool              `json:"success"`
//	ErrorCode    string            `json:"error_code"`
//	ErrorMessage string            `json:"message"`
//	Account      datastore.Account `json:"account"`
//}
//
//// UserAccount defines an account and whether it is selected for a user
//type UserAccount struct {
//	datastore.Account
//	Selected bool `json:"selected"`
//}
//
//// AccountsForUserResponse is the JSON response from the API Accounts method
//type AccountsForUserResponse struct {
//	Success      bool          `json:"success"`
//	ErrorCode    string        `json:"error_code"`
//	ErrorMessage string        `json:"message"`
//	Accounts     []UserAccount `json:"accounts"`
//}

//func formatAccountsResponse(success bool, errorCode, message string, accounts []datastore.Account, w http.ResponseWriter) error {
//	response := AccountsResponse{Success: success, ErrorCode: errorCode, ErrorMessage: message, Accounts: accounts}
//
//	// Encode the response as JSON
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		log.Println("Error forming the accounts response.")
//		return err
//	}
//	return nil
//}
//
//func formatAccountResponse(success bool, errorCode, message string, account datastore.Account, w http.ResponseWriter) error {
//	response := AccountResponse{Success: success, ErrorCode: errorCode, ErrorMessage: message, Account: account}
//
//	// Encode the response as JSON
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		log.Println("Error forming the account response.")
//		return err
//	}
//	return nil
//}
//
//func formatUserAccountsResponse(success bool, errorCode, message string, accounts []UserAccount, w http.ResponseWriter) error {
//	response := AccountsForUserResponse{Success: success, ErrorCode: errorCode, ErrorMessage: message, Accounts: accounts}
//
//	// Encode the response as JSON
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		log.Println("Error forming the accounts response.")
//		return err
//	}
//	return nil
//}
//
//// AccountListHandler returns the list of accounts
//func (wb Service) AccountListHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	authUser, err := wb.checkIsStandardAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatAccountsResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	accounts, err := datastore.Environ.DB.ListAllowedAccounts(authUser, "")
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		formatAccountsResponse(false, "error-accounts-json", err.Error(), nil, w)
//		return
//	}
//
//	// Format the model for output and return JSON response
//	w.WriteHeader(http.StatusOK)
//	formatAccountsResponse(true, "", "", accounts, w)
//
//}
//
//// AccountGetHandler returns a single account
//func (wb Service) AccountGetHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatAccountResponse(false, "error-auth", "", datastore.Account{}, w)
//		return
//	}
//
//	vars := mux.Vars(r)
//	accountID, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		w.WriteHeader(http.StatusNotFound)
//		formatAccountResponse(false, "error-account-invalid", "Cannot find an account for the ID", datastore.Account{}, w)
//		return
//	}
//
//	account, err := datastore.Environ.DB.GetAccount(accountID)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		formatAccountResponse(false, "error-account-invalid", err.Error(), datastore.Account{}, w)
//		return
//	}
//
//	// Format the model for output and return JSON response
//	w.WriteHeader(http.StatusOK)
//	formatAccountResponse(true, "", "", account, w)
//
//}
//
//// AccountCreateHandler handles account creation
//func (wb Service) AccountCreateHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatAccountsResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	account := datastore.Account{}
//	err = json.NewDecoder(r.Body).Decode(&account)
//	switch {
//	// Check we have some data
//	case err == io.EOF:
//		w.WriteHeader(http.StatusBadRequest)
//		formatAccountResponse(false, "error-account-data", "No account data supplied", account, w)
//		return
//		// Check for parsing errors
//	case err != nil:
//		w.WriteHeader(http.StatusBadRequest)
//		formatAccountResponse(false, "error-data-json", err.Error(), account, w)
//		return
//	}
//
//	err = datastore.Environ.DB.CreateAccount(account)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		formatAccountResponse(false, "error-account-create", err.Error(), account, w)
//		return
//	}
//
//	formatAccountResponse(true, "", "", account, w)
//}
//
//// AccountUpdateHandler handles account update
//func (wb Service) AccountUpdateHandler(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", JSONHeader)
//
//	_, err := wb.checkIsSuperuserAndGetUserFromJWT(w, r)
//	if err != nil {
//		formatAccountsResponse(false, "error-auth", "", nil, w)
//		return
//	}
//
//	account := datastore.Account{}
//	err = json.NewDecoder(r.Body).Decode(&account)
//	switch {
//	// Check we have some data
//	case err == io.EOF:
//		w.WriteHeader(http.StatusBadRequest)
//		formatAccountResponse(false, "error-account-data", "No account data supplied", account, w)
//		return
//		// Check for parsing errors
//	case err != nil:
//		w.WriteHeader(http.StatusBadRequest)
//		formatAccountResponse(false, "error-data-json", err.Error(), account, w)
//		return
//	}
//
//	account, err = datastore.Environ.DB.UpdateAccount(account)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		formatAccountResponse(false, "error-account-update", err.Error(), account, w)
//		return
//	}
//
//	formatAccountResponse(true, "", "", account, w)
//}
//
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
