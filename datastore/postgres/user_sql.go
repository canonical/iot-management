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

const createUserTableSQL = `
CREATE TABLE IF NOT EXISTS userinfo (
   id             serial primary key,
   created        timestamp default current_timestamp,
   modified       timestamp default current_timestamp,
   username       varchar(200) unique not null,
   name           varchar(200) not null,
   email          varchar(200) not null,
   user_role      int not null
)
`

const createUserSQL = "insert into userinfo (username, name, email, user_role) values ($1,$2,$3,$4) returning id"
const listUsersSQL = "select id, username, name, email, user_role from userinfo order by username"
const getUserSQL = "select id, username, name, email, user_role from userinfo where username=$1"
