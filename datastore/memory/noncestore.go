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
	"time"

	"github.com/canonical/iot-management/datastore"
	"gopkg.in/errgo.v1"
)

// NonceStore is a nonce store backed by database
type NonceStore struct {
	DB *Store
}

// Accept implements openid.NonceStore.Accept
func (s *NonceStore) Accept(endpoint, nonce string) error {
	return s.accept(endpoint, nonce, time.Now())
}

// accept is the implementation of Accept. The third parameter is the
// current time, useful for testing.
func (s *NonceStore) accept(endpoint, nonce string, now time.Time) error {
	// From the openid specification:
	//
	// openid.response_nonce
	//
	// Value: A string 255 characters or less in length, that MUST be
	// unique to this particular successful authentication response.
	// The nonce MUST start with the current time on the server, and
	// MAY contain additional ASCII characters in the range 33-126
	// inclusive (printable non-whitespace characters), as necessary
	// to make each response unique. The date and time MUST be
	// formatted as specified in section 5.6 of [RFC3339], with the
	// following restrictions:
	//
	// + All times must be in the UTC timezone, indicated with a "Z".
	//
	// + No fractional seconds are allowed
	//
	// For example: 2005-05-15T17:11:51ZUNIQUE

	if len(nonce) < 20 {
		return fmt.Errorf("%q does not contain a valid timestamp", nonce)
	}
	t, err := time.Parse(time.RFC3339, nonce[:20])
	if err != nil {
		return fmt.Errorf("%q does not contain a valid timestamp: %v", nonce, err)
	}

	// Check if the nonce has expired
	diff := now.Sub(t)
	if diff > datastore.OpenidNonceMaxAge {
		return fmt.Errorf("%q too old", nonce)
	}

	openidNonce := datastore.OpenidNonce{Nonce: nonce, Endpoint: endpoint, TimeStamp: t.Unix()}
	err = s.DB.createOpenidNonce(openidNonce)
	return errgo.Mask(err)
}
