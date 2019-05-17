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

package usso

import (
	"errors"
	"log"
	"strings"
	"time"

	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/juju/usso/openid"
)

func createJWT(jwtSecret, username, name, email, identity string, role int, expires int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[ClaimsUsername] = username
	claims[ClaimsName] = name
	claims[ClaimsEmail] = email
	claims[ClaimsIdentity] = identity
	claims[ClaimsRole] = role
	//claims[ClaimsAccountFilter] = datastore.Environ.Config.AccountFilter
	claims[StandardClaimExpiresAt] = expires

	if len(jwtSecret) == 0 {
		return "", errors.New("JWT secret empty value. Please configure it properly")
	}

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Printf("Error signing the JWT: %v", err.Error())
	}
	return tokenString, err
}

// NewJWTToken creates a new JWT from the verified OpenID response
func NewJWTToken(jwtSecret string, resp *openid.Response, role int) (string, error) {

	return createJWT(jwtSecret, resp.SReg["nickname"], resp.SReg["fullname"], resp.SReg["email"], resp.ID, role, time.Now().Add(time.Hour*24).Unix())
}

// VerifyJWT checks that we have a valid token
func VerifyJWT(jwtSecret, jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (i interface{}, e error) {
		if len(jwtSecret) == 0 {
			return []byte{}, errors.New("JWT secret empty value. Please configure it properly")
		}
		return []byte(jwtSecret), nil
	})

	return token, err
}

// AddJWTCookie sets the JWT as a cookie
func AddJWTCookie(jwtToken string, w http.ResponseWriter) {
	// Set the JWT as a bearer token
	// (In practice, the cookie will be used more as clicking on a page link will not send the auth header)
	w.Header().Set("Authorization", "Bearer "+jwtToken)

	expireCookie := time.Now().Add(time.Hour * 1)
	cookie := http.Cookie{Name: JWTCookie, Value: jwtToken, Expires: expireCookie, HttpOnly: true}
	http.SetCookie(w, &cookie)
}

// JWTExtractor extracts the JWT from a request and returns the token string.
// The token is not verified.
func JWTExtractor(r *http.Request) (string, error) {
	// Get the JWT from the header
	jwtToken := r.Header.Get("Authorization")
	splitToken := strings.Split(jwtToken, " ")
	if len(splitToken) != 2 {
		jwtToken = ""
	} else {
		jwtToken = splitToken[1]
	}

	// Check the cookie if we don't have a bearer token in the header
	if jwtToken == "" {
		cookie, err := r.Cookie(JWTCookie)
		if err != nil {
			log.Println("Cannot find the JWT")
			return "", errors.New("Cannot find the JWT")
		}
		jwtToken = cookie.Value
	}

	return jwtToken, nil
}
