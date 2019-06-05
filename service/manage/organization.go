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
	"github.com/CanonicalLtd/iot-management/datastore"
	"github.com/CanonicalLtd/iot-management/domain"
)

// OrganizationsForUser fetches the organizations for a user
func (srv *Management) OrganizationsForUser(username string) ([]domain.Organization, error) {
	orgs, err := srv.DB.OrganizationsForUser(username)
	if err != nil {
		return nil, err
	}

	oo := []domain.Organization{}
	for _, o := range orgs {
		oo = append(oo, domain.Organization{
			OrganizationID: o.OrganizationID,
			Name:           o.Name,
		})
	}
	return oo, nil
}

// OrganizationForUserToggle toggles organization access for a user
func (srv *Management) OrganizationForUserToggle(orgID, username string) error {
	return srv.DB.OrganizationForUserToggle(orgID, username)
}

// OrganizationGet fetches an organization
func (srv *Management) OrganizationGet(orgID string) (domain.Organization, error) {
	org, err := srv.DB.OrganizationGet(orgID)
	if err != nil {
		return domain.Organization{}, err
	}
	return domain.Organization{
		OrganizationID: org.OrganizationID,
		Name:           org.Name,
	}, nil
}

// OrganizationCreate creates a new organization
func (srv *Management) OrganizationCreate(org domain.Organization) error {
	o := datastore.Organization{
		OrganizationID: org.OrganizationID,
		Name:           org.Name,
	}

	return srv.DB.OrganizationCreate(o)
}

// OrganizationUpdate updates an organization
func (srv *Management) OrganizationUpdate(org domain.Organization) error {
	o := datastore.Organization{
		OrganizationID: org.OrganizationID,
		Name:           org.Name,
	}

	return srv.DB.OrganizationUpdate(o)
}
