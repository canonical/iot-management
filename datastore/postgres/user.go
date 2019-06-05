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
	"github.com/CanonicalLtd/iot-management/datastore"
	"log"
)

// createUserTable creates the database table for devices with its indexes.
func (db *Store) createUserTable() error {
	_, err := db.Exec(createUserTableSQL)
	return err
}

// CreateUser creates a new user
func (db *Store) CreateUser(user datastore.User) (int64, error) {
	var createdUserID int64

	err := db.QueryRow(createUserSQL, user.Username, user.Name, user.Email, user.Role).Scan(&createdUserID)
	if err != nil {
		log.Printf("Error creating user `%s`: %v\n", user.Username, err)
	}

	return createdUserID, err
}

// UserUpdate updates a user
func (db *Store) UserUpdate(user datastore.User) error {
	_, err := db.Exec(updateUserSQL, user.Username, user.Name, user.Email, user.Role)
	return err
}

// UserDelete removes a user
func (db *Store) UserDelete(username string) error {
	_, err := db.Exec(deleteUserSQL, username)
	return err
}

// UserList lists existing users
func (db *Store) UserList() ([]datastore.User, error) {
	rows, err := db.Query(listUsersSQL)
	if err != nil {
		log.Printf("Error retrieving database users: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	return db.rowsToUsers(rows)
}

// GetUser gets an existing user
func (db *Store) GetUser(username string) (datastore.User, error) {
	row := db.QueryRow(getUserSQL, username)
	user, err := db.rowToUser(row)
	if err != nil {
		log.Printf("Error retrieving user %v: %v\n", username, err)
	}
	return user, err
}

func (db *Store) rowToUser(row *sql.Row) (datastore.User, error) {
	user := datastore.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Role)
	if err != nil {
		return datastore.User{}, err
	}

	return user, nil
}

func (db *Store) rowsToUser(rows *sql.Rows) (datastore.User, error) {
	user := datastore.User{}
	err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Role)
	if err != nil {
		return datastore.User{}, err
	}

	return user, nil
}

func (db *Store) rowsToUsers(rows *sql.Rows) ([]datastore.User, error) {
	users := []datastore.User{}

	for rows.Next() {
		user, err := db.rowsToUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
