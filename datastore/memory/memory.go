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
	"github.com/CanonicalLtd/iot-management/datastore"
	"github.com/juju/usso/openid"
	"sync"
)

// Store implements an in-memory store for testing
type Store struct {
	lock     sync.RWMutex
	Users    []datastore.User
	Orgs     []datastore.Organization
	OrgUsers []datastore.OrganizationUser
}

// NewStore creates a new memory store
func NewStore() *Store {
	return &Store{
		Users: []datastore.User{
			{Username: "jamesj", Name: "JJ", Role: 300},
		},
		Orgs:     []datastore.Organization{{OrganizationID: "abc", Name: "Example Org"}},
		OrgUsers: []datastore.OrganizationUser{{OrganizationID: "abc", Username: "jamesj"}},
	}
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

// OpenIDNonceStore returns an openid nonce store
func (mem *Store) OpenIDNonceStore() openid.NonceStore {
	return &NonceStore{DB: mem}
}

// createOpenidNonce stores a new nonce entry
func (mem *Store) createOpenidNonce(nonce datastore.OpenidNonce) error {
	// Delete the expired nonce
	// Create the nonce in the database
	return nil
}

// OrgUserAccess checks if the user has permissions to access the organization
func (mem *Store) OrgUserAccess(orgID, username string, role int) bool {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	// Superusers can access all accounts
	if role == datastore.Superuser {
		return true
	}

	for _, ou := range mem.OrgUsers {
		if ou.OrganizationID == orgID && ou.Username == username {
			return true
		}
	}
	return false
}
