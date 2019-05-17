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
	"fmt"
	"github.com/CanonicalLtd/iot-management/config"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/juju/usso"
	"github.com/juju/usso/openid"
)

var (
	teams    = "" // e.g. ce-web-logs,canonical
	required = "email,fullname,nickname"
	optional = ""
)

var client *openid.Client

// verify is used to perform the OpenID verification of the login
// response. This is declared as a variable so it can be overridden for
// testing.
var verify func(string) (*openid.Response, error)

func getClient(nonce openid.NonceStore) *openid.Client {
	if client != nil {
		return client
	}
	client = openid.NewClient(usso.ProductionUbuntuSSOServer, nonce, nil)
	verify = client.Verify
	return client
}

// LoginHandler processes the login for Ubuntu SSO
func LoginHandler(settings *config.Settings, nonce openid.NonceStore, w http.ResponseWriter, r *http.Request) (*openid.Response, *openid.Request, string, error) {
	getClient(nonce)
	r.ParseForm()

	url := *r.URL

	// Set the return URL: from the OpenID login with the full domain name
	url.Scheme = settings.URLScheme
	url.Host = settings.URLHost

	if r.Form.Get("openid.ns") == "" {
		req := openid.Request{
			ReturnTo:     url.String(),
			Teams:        strings.FieldsFunc(teams, isComma),
			SRegRequired: strings.FieldsFunc(required, isComma),
			SRegOptional: strings.FieldsFunc(optional, isComma),
		}
		url := client.RedirectURL(&req)
		http.Redirect(w, r, url, http.StatusFound)
		return nil, &req, "", nil
	}

	resp, err := verify(url.String())
	if err != nil {
		// A mangled OpenID response is suspicious, so leave a nasty response
		return nil, nil, "", fmt.Errorf("error verifying the OpenID response: %v", err)
	}

	return resp, nil, r.Form.Get("openid.sreg.nickname"), nil
}

func isComma(c rune) bool {
	return c == ','
}

// LogoutHandler logs the user out by removing the cookie and the JWT authorization header
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Remove the authorization header with contains the bearer token
	w.Header().Del("Authorization")

	// Create a new invalid token with an unauthorized user
	jwtToken, err := createJWT("INVALID", "INVALID", "Not Logged-In", "", "", 0, 0)
	if err != nil {
		log.Println("Error logging out:", err.Error())
	}

	// Update the cookie with the invalid token and expired date
	c, err := r.Cookie(JWTCookie)
	if err != nil {
		log.Println("Error logging out:", err.Error())
	}
	c.Value = jwtToken
	c.Expires = time.Now().AddDate(0, 0, -1)

	// Set the bearer token and the cookie
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
