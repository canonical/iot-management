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

// OrganizationCreate creates a new organization
func (mem *Store) OrganizationCreate(org datastore.Organization) error {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	_, err := mem.organizationGet(org.OrganizationID)
	if err == nil {
		return fmt.Errorf("organization already exists `%s`", org.OrganizationID)
	}

	mem.Orgs = append(mem.Orgs, org)
	return nil
}

// OrganizationUpdate updates an organization
func (mem *Store) OrganizationUpdate(org datastore.Organization) error {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	found := false
	orgs := []datastore.Organization{}

	for _, o := range mem.Orgs {
		if o.OrganizationID == org.OrganizationID {
			orgs = append(orgs, org)
			found = true
			continue
		}
		orgs = append(orgs, o)
	}

	if !found {
		return fmt.Errorf("organization not found: %s", org.OrganizationID)
	}

	mem.Orgs = orgs
	return nil
}

//OrganizationForUserToggle toggles the user access for an organization
func (mem *Store) OrganizationForUserToggle(orgID, username string) error {
	err := mem.removeOrgUserAccess(orgID, username)
	if err != nil {
		// Create it as it didn't exist
		return mem.addOrgUserAccess(orgID, username)
	}
	return nil
}

// removeOrgUserAccess remove access to an organization for a user
func (mem *Store) removeOrgUserAccess(orgID, username string) error {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	found := false
	oo := []datastore.OrganizationUser{}
	for _, ou := range mem.OrgUsers {
		if ou.OrganizationID == orgID && ou.Username == username {
			found = true
			continue
		}
		oo = append(oo, ou)
	}
	if !found {
		return fmt.Errorf("record not found")
	}
	mem.OrgUsers = oo
	return nil
}

// addOrgUserAccess adds access to an organization for a user
func (mem *Store) addOrgUserAccess(orgID, username string) error {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	mem.OrgUsers = append(mem.OrgUsers, datastore.OrganizationUser{
		OrganizationID: orgID,
		Username:       username,
	})
	return nil
}
