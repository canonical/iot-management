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

const createOrganizationTableSQL = `
CREATE TABLE IF NOT EXISTS organization (
	id            serial primary key,
	code          varchar(200) not null unique,
	name          varchar(200) not null,
	active        bool
)
`
const createOrganizationUserTableSQL = `
CREATE TABLE IF NOT EXISTS organization_user (
	id            serial primary key,
	org_id        int not null,
	user_id       int not null
)
`

const getOrganizationSQL = "select code, name from organization where code=$1"

const organizationUserAccessSQL = `
SELECT EXISTS(
	select ou.id from organization_user ou
	inner join organization o on o.id=ou.org_id
	inner join userinfo u on u.id=ou.user_id
	where o.code=$1 and u.username=$2
)
`
const listUserOrganizationsSQL = `
	select a.code, a.name
	from organization a
	inner join organization_user l on a.id = l.org_id
	inner join userinfo u on l.user_id = u.id
	where u.username=$1
`
