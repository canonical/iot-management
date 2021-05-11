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

package manage

import (
	"github.com/everactive/iot-management/datastore"
	"github.com/everactive/iot-management/domain"
	"github.com/juju/usso/openid"
)

// OpenIDNonceStore fetches the OpenID nonce store
func (srv *Management) OpenIDNonceStore() openid.NonceStore {
	return srv.DB.OpenIDNonceStore()
}

// GetUser fetches a user from the database
func (srv *Management) GetUser(username string) (domain.User, error) {
	u, err := srv.DB.GetUser(username)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}, nil
}

// UserList fetches the existing users
func (srv *Management) UserList() ([]domain.User, error) {
	users, err := srv.DB.UserList()
	if err != nil {
		return nil, err
	}

	uu := []domain.User{}

	for _, u := range users {
		uu = append(uu, domain.User{
			ID:       u.ID,
			Name:     u.Name,
			Username: u.Username,
			Email:    u.Email,
			Role:     u.Role,
		})
	}
	return uu, nil
}

// CreateUser creates a new user
func (srv *Management) CreateUser(user domain.User) error {
	u := datastore.User{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
	}

	_, err := srv.DB.CreateUser(u)
	return err
}

// UserUpdate updates a new user
func (srv *Management) UserUpdate(user domain.User) error {
	u := datastore.User{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
	}

	return srv.DB.UserUpdate(u)
}

// UserDelete removes a user
func (srv *Management) UserDelete(username string) error {
	return srv.DB.UserDelete(username)
}
