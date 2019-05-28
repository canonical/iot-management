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

const createNonceTableSQL = `
 CREATE TABLE IF NOT EXISTS openidnonce (
	 id             serial primary key,
	 nonce          varchar(255) not null,
	 endpoint       varchar(255) not null,
	 timestamp      int not null
 )
`

const createOpenidNonceSQL = "INSERT INTO openidnonce (nonce, endpoint, timestamp) VALUES ($1, $2, $3)"
const deleteExpiredOpenidNonceSQL = "DELETE FROM openidnonce where timestamp<$1"
