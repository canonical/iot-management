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
	"errors"
	"github.com/everactive/iot-management/datastore"
	"github.com/everactive/iot-management/web/usso"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func (wb Service) checkIsStandardAndGetUserFromJWT(w http.ResponseWriter, r *http.Request) (datastore.User, error) {
	return wb.checkPermissionsAndGetUserFromJWT(w, r, datastore.Standard)
}

func (wb Service) checkIsAdminAndGetUserFromJWT(w http.ResponseWriter, r *http.Request) (datastore.User, error) {
	return wb.checkPermissionsAndGetUserFromJWT(w, r, datastore.Admin)
}

func (wb Service) checkIsSuperuserAndGetUserFromJWT(w http.ResponseWriter, r *http.Request) (datastore.User, error) {
	return wb.checkPermissionsAndGetUserFromJWT(w, r, datastore.Superuser)
}

func (wb Service) checkPermissionsAndGetUserFromJWT(w http.ResponseWriter, r *http.Request, minimumAuthorizedRole int) (datastore.User, error) {
	user, err := wb.getUserFromJWT(w, r)
	if err != nil {
		return user, err
	}
	err = checkUserPermissions(user, minimumAuthorizedRole)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (wb Service) getUserFromJWT(w http.ResponseWriter, r *http.Request) (datastore.User, error) {
	token, err := wb.JWTCheck(w, r)
	if err != nil {
		return datastore.User{}, err
	}

	// Null token is invalid
	if token == nil {
		return datastore.User{}, errors.New("No JWT provided")
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims[usso.ClaimsUsername].(string)
	role := int(claims[usso.ClaimsRole].(float64))

	return datastore.User{
		Username: username,
		Role:     role,
	}, nil
}

func checkUserPermissions(user datastore.User, minimumAuthorizedRole int) error {
	if user.Role < minimumAuthorizedRole {
		return errors.New("The user is not authorized")
	}
	return nil
}
