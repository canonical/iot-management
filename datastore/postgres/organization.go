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

package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/canonical/iot-management/datastore"
)

// createOrganizationTable creates the database table for devices with its indexes.
func (db *Store) createOrganizationTable() error {
	_, err := db.Exec(createOrganizationTableSQL)
	return err
}

// createOrganizationUserTable creates the database table for devices with its indexes.
func (db *Store) createOrganizationUserTable() error {
	_, err := db.Exec(createOrganizationUserTableSQL)
	return err
}

// OrganizationGet returns an organization
func (db *Store) OrganizationGet(orgID string) (datastore.Organization, error) {
	org := datastore.Organization{}

	err := db.QueryRow(getOrganizationSQL, orgID).Scan(&org.OrganizationID, &org.Name)
	if err != nil {
		log.Printf("Error retrieving organization %v: %v\n", orgID, err)
	}
	return org, err
}

// OrgUserAccess checks if the user has permissions to access the organization
func (db *Store) OrgUserAccess(orgID, username string, role int) bool {
	// Superusers can access all accounts
	if role == datastore.Superuser {
		return true
	}

	var linkExists bool
	err := db.QueryRow(organizationUserAccessSQL, orgID, username).Scan(&linkExists)
	if err != nil {
		log.Printf("Error verifying the account-user link: %v\n", err)
		return false
	}
	return linkExists
}

// OrganizationsForUser returns the organizations a user can access
func (db *Store) OrganizationsForUser(username string) ([]datastore.Organization, error) {
	var s string

	// Check if the user is a superuser
	user, err := db.GetUser(username)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %v", err)
	}

	// No restrictions for the superuser
	if user.Role == datastore.Superuser {
		s = listOrganizationsSQL
	} else {
		s = listUserOrganizationsSQL
	}

	rows, err := db.Query(s, username)

	if err != nil {
		log.Printf("Error retrieving database accounts: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	return rowsToOrganizations(rows)
}

// OrganizationForUserToggle toggles the user access to an organization
func (db *Store) OrganizationForUserToggle(orgID, username string) error {
	result, err := db.Exec(deleteOrganizationUserAccessSQL, orgID, username)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		// Create the link
		_, err := db.Exec(createOrganizationUserAccessSQL, orgID, username)
		if err != nil {
			return err
		}
	}

	return nil
}

func rowsToOrganizations(rows *sql.Rows) ([]datastore.Organization, error) {
	orgs := []datastore.Organization{}

	for rows.Next() {
		org := datastore.Organization{}
		err := rows.Scan(&org.OrganizationID, &org.Name)
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, org)
	}

	return orgs, nil
}

// OrganizationCreate creates a new organization
func (db *Store) OrganizationCreate(org datastore.Organization) error {
	var createdID int64

	err := db.QueryRow(createOrganizationSQL, org.OrganizationID, org.Name).Scan(&createdID)
	if err != nil {
		log.Printf("Error creating organization `%s`: %v\n", org.OrganizationID, err)
	}

	return err
}

// OrganizationUpdate updates an organization
func (db *Store) OrganizationUpdate(org datastore.Organization) error {
	_, err := db.Exec(updateOrganizationSQL, org.OrganizationID, org.Name)
	return err
}
