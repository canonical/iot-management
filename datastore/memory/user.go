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

package memory

import (
	"fmt"
	"github.com/everactive/iot-management/datastore"
)

// UserList lists existing users
func (mem *Store) UserList() ([]datastore.User, error) {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	return mem.Users, nil
}

// CreateUser creates a new user
func (mem *Store) CreateUser(user datastore.User) (int64, error) {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	user.ID = int64(len(mem.Users) + 1)
	mem.Users = append(mem.Users, user)
	return user.ID, nil
}

// GetUser gets an existing user
func (mem *Store) GetUser(username string) (datastore.User, error) {
	mem.lock.RLock()
	defer mem.lock.RUnlock()

	for _, u := range mem.Users {
		if u.Username == username {
			return u, nil
		}
	}

	return datastore.User{}, fmt.Errorf("cannot find the user `%s`", username)
}

// UserUpdate updates a user
func (mem *Store) UserUpdate(user datastore.User) error {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	var index = -1

	for i, u := range mem.Users {
		if u.Username == user.Username {
			user.ID = u.ID
			index = i
			break
		}
	}

	if index < 0 {
		return fmt.Errorf("error finding user")
	}
	mem.Users[index] = user
	return nil
}

// UserDelete removes a user
func (mem *Store) UserDelete(username string) error {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	var index = -1

	for i, u := range mem.Users {
		if u.Username == username {
			index = i
			break
		}
	}

	if index < 0 {
		return fmt.Errorf("error finding user")
	}

	// Remove the element
	mem.Users = append(mem.Users[:index], mem.Users[index+1:]...)
	return nil
}
