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

package crypt

import (
	"crypto/rand"
	"encoding/base64"
	"regexp"
)

// RegexpAlpha is a regex pattern for alphabetic characters only
var RegexpAlpha = regexp.MustCompile("[^a-zA-Z]+")

// RegexpAlphanumeric is a regex pattern for alphabetic/numeric characters only
var RegexpAlphanumeric = regexp.MustCompile("[^a-zA-Z0-9]+")

// CreateSecret generates a secret that can be used for encryption
func CreateSecret(length int) (string, error) {
	rb := make([]byte, length)
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}

	return RegexpAlphanumeric.ReplaceAllString(base64.URLEncoding.EncodeToString(rb), ""), nil
}
