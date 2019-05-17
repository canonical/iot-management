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

package web

import (
	"errors"
	"github.com/CanonicalLtd/iot-management/web/usso"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

// Logger Handle logging for the web web
func Logger(start time.Time, r *http.Request) {
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}

// Middleware to pre-process web web requests
func Middleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		Logger(start, r)

		inner.ServeHTTP(w, r)
	})
}

// JWTCheck extracts the JWT from the request, validates it and returns the token
func (wb Service) JWTCheck(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {

	// Get the JWT from the header or cookie
	jwtToken, err := usso.JWTExtractor(r)
	if err != nil {
		log.Println("Error in JWT extraction:", err.Error())
		return nil, errors.New("error in retrieving the authentication token")
	}

	// Verify the JWT string
	token, err := usso.VerifyJWT(wb.Settings.JwtSecret, jwtToken)
	if err != nil {
		log.Printf("JWT fails verification: %v", err.Error())
		return nil, errors.New("the authentication token is invalid")
	}

	if !token.Valid {
		log.Println("Invalid JWT")
		return nil, errors.New("the authentication token is invalid")
	}

	// Set up the bearer token in the header
	w.Header().Set("Authorization", "Bearer "+jwtToken)

	return token, nil
}
