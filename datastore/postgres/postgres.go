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
	_ "github.com/lib/pq" // postgresql driver
	"log"
)

// Store implements an in-memory store for testing
type Store struct {
	driver string
	*sql.DB
}

var pgStore *Store

// OpenStore returns an open database connection
func OpenStore(driver, dataSource string) *Store {
	if pgStore != nil {
		return pgStore
	}

	// Open the database
	pgStore = openDatabase(driver, dataSource)

	// Create the tables, if needed
	pgStore.createTables()

	return pgStore
}

// openDatabase return an open database connection for a postgreSQL database
func openDatabase(driver, dataSource string) *Store {
	// Open the database connection
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Fatalf("Error opening the database: %v\n", err)
	}

	// Check that we have a valid database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error accessing the database: %v\n", err)
	}

	return &Store{driver, db}
}

func (db *Store) createTables() {
	_ = db.createUserTable()
	_ = db.createNonceTable()
	_ = db.createOrganizationTable()
	_ = db.createOrganizationUserTable()
}
