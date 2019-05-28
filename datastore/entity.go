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

package datastore

import "time"

// Available user roles:
//
// * Invalid:	default value set in case there is no authentication previous process for this user and thus not got a valid role.
// * Standard:	role for regular users. This is the less privileged role
// * Admin:		role for createsuperuser users, including standard role permissions but not superuser ones
// * Superuser:	role for users having all the permissions
const (
	Invalid   = iota       // 0
	Standard  = 100 * iota // 100
	Admin                  // 200
	Superuser              // 300
)

// User holds user personal, authentication and authorization info
type User struct {
	ID       int64
	Username string
	Name     string
	Email    string
	Role     int
}

// OpenidNonceMaxAge is the maximum age of stored nonces. Any nonces older
// than this will automatically be rejected. Stored nonces older
// than this will periodically be purged from the database.
const OpenidNonceMaxAge = MaxNonceAgeInSeconds * time.Second

// MaxNonceAgeInSeconds is the nonce age
const MaxNonceAgeInSeconds = 60

// OpenidNonce holds the details of the nonce, combining a timestamp and random text
type OpenidNonce struct {
	ID        int64
	Nonce     string
	Endpoint  string
	TimeStamp int64
}

// Organization holds details of the organization
type Organization struct {
	OrganizationID string
	Name           string
}

// OrganizationUser holds links a user and organization
type OrganizationUser struct {
	OrganizationID string
	Username       string
}
