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
)

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

// OrganizationsForUser returns the organizations a user can access
func (mem *Store) OrganizationsForUser(username string) ([]datastore.Organization, error) {
	mem.lock.RLock()
	defer mem.lock.RLock()

	orgs := []datastore.Organization{}

	if username == "invalid" {
		return nil, fmt.Errorf("error finding user `%s`", username)
	}

	for _, ou := range mem.OrgUsers {
		if ou.Username != username {
			continue
		}

		o, err := mem.OrganizationGet(ou.OrganizationID)
		if err != nil {
			return nil, err
		}

		orgs = append(orgs, o)
	}
	return orgs, nil
}

// OrganizationGet returns an organization
func (mem *Store) OrganizationGet(orgID string) (datastore.Organization, error) {
	mem.lock.RLock()
	defer mem.lock.RUnlock()

	return mem.organizationGet(orgID)
}

// organizationGet returns an organization without locking the store
func (mem *Store) organizationGet(orgID string) (datastore.Organization, error) {
	for _, o := range mem.Orgs {
		if o.OrganizationID == orgID {
			return o, nil
		}
	}

	return datastore.Organization{}, fmt.Errorf("error finding organization `%s`", orgID)
}
