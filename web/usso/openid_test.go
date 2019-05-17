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
	"github.com/CanonicalLtd/iot-management/config"
	"github.com/CanonicalLtd/iot-management/datastore/memory"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"fmt"

	"github.com/juju/usso"
	"github.com/juju/usso/openid"
)

func TestLoginHandlerUSSORedirect(t *testing.T) {
	// Mock the database
	settings, _ := config.Config("../../settings.yaml")
	db := memory.NewStore()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/login", nil)
	_, req, _, err := LoginHandler(settings, db.OpenIDNonceStore(), w, r)
	if err != nil {
		t.Errorf("Error with the redirect URL: %v", err)
	}
	if req == nil {
		t.Error("Expected a redirect request")
	}

	u, err := url.Parse(w.Header().Get("Location"))
	if err != nil {
		t.Errorf("Error Parsing the redirect URL: %v", u)
	}

	// Check that the redirect is to the Ubuntu SSO service
	url := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	if url != usso.ProductionUbuntuSSOServer.LoginURL() {
		t.Errorf("Unexpected redirect URL: %v", url)
	}
}

func TestLoginHandlerReturn(t *testing.T) {
	// Response parameters from OpenID login
	const url = "/login?openid.ns=http://specs.openid.net/auth/2.0&openid.mode=id_res&openid.op_endpoint=https://login.ubuntu.com/%2Bopenid&openid.claimed_id=https://login.ubuntu.com/%2Bid/AAAAAA&openid.identity=https://login.ubuntu.com/%2Bid/AAAAAA&openid.return_to=http://return.to&openid.response_nonce=2005-05-15T17:11:51ZUNIQUE&openid.assoc_handle=1&openid.signed=op_endpoint,return_to,response_nonce,assoc_handle,claimed_id,identity,sreg.email,sreg.fullname&openid.sig=AAAA&openid.ns.sreg=http://openid.net/extensions/sreg/1.1&openid.sreg.email=a@example.org&openid.sreg.fullname=A&openid.sreg.nickname=jamesj"

	// Mock the database and OpenID verification
	settings, _ := config.Config("../../settings.yaml")
	db := memory.NewStore()
	verify = verifySuccess

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", url, nil)
	resp, req, s, err := LoginHandler(settings, db.OpenIDNonceStore(), w, r)

	if err != nil {
		t.Errorf("Expected success, got: %v", err)
	}
	if req != nil {
		t.Errorf("Expected response, got redirect: %v", req)
	}
	if s != "jamesj" {
		t.Errorf("Expected nickname `jamesj`, got: %v", s)
	}
	if resp == nil {
		t.Error("Expected response with claims")
	}
}

func TestLoginHandlerReturnBad(t *testing.T) {
	// Response parameters from OpenID login (with empty nickname)
	const url = "/login?openid.ns=http://specs.openid.net/auth/2.0&openid.mode=id_res&openid.op_endpoint=https://login.ubuntu.com/%2Bopenid&openid.claimed_id=https://login.ubuntu.com/%2Bid/AAAAAA&openid.identity=https://login.ubuntu.com/%2Bid/AAAAAA&openid.return_to=http://return.to&openid.response_nonce=2005-05-15T17:11:51ZUNIQUE&openid.assoc_handle=1&openid.signed=op_endpoint,return_to,response_nonce,assoc_handle,claimed_id,identity,sreg.email,sreg.fullname&openid.sig=AAAA&openid.ns.sreg=http://openid.net/extensions/sreg/1.1&openid.sreg.email=a@example.org&openid.sreg.fullname=A&openid.sreg.nickname="

	// Mock the database and OpenID verification
	settings, _ := config.Config("../../settings.yaml")
	db := memory.NewStore()
	verify = verifyFail

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", url, nil)
	resp, req, s, err := LoginHandler(settings, db.OpenIDNonceStore(), w, r)

	if err == nil {
		t.Error("Expected error, got success")
	}
	if req != nil {
		t.Errorf("Expected response, got redirect: %v", req)
	}
	if s != "" {
		t.Errorf("Expected nickname ``, got: %v", s)
	}
	if resp != nil {
		t.Errorf("Expected response without claims, got: %v", resp)
	}
}

func TestLogoutHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/logout", nil)

	// Create a test cookie
	expireCookie := time.Now().Add(time.Hour * 1)
	c := http.Cookie{Name: JWTCookie, Value: "test-cookie", Expires: expireCookie, HttpOnly: true}

	// Add auth header and cookie to the request
	r.Header.Set("Authorization", "test-header")
	r.AddCookie(&c)

	// Send the logout request
	http.HandlerFunc(LogoutHandler).ServeHTTP(w, r)

	if w.Code != http.StatusTemporaryRedirect {
		t.Errorf("Expected HTTP status '307', got: %v", w.Code)
	}

	if w.Header().Get("Location") != "/" {
		t.Errorf("Expected redirect to / but got: %v", w.Header().Get("Location"))
	}

	// Copy the headers over to a new Request
	request := &http.Request{Header: w.Header()}

	// Check the headers
	if request.Header.Get("Authorization") != "" {
		t.Errorf("Expected no Authorization header, got: %v", request.Header.Get("Authorization"))
	}
	_, err := request.Cookie(JWTCookie)
	if err == nil {
		t.Error("Expected the JWT cookie to have been deleted")
	}
}

func verifySuccess(requestURL string) (*openid.Response, error) {
	params := make(map[string]string)

	tokens := strings.Split(requestURL, "&")
	for _, t := range tokens {
		tks := strings.Split(t, "=")

		if len(tks) == 2 {
			params[strings.TrimSpace(tks[0])] = tks[1]
		}
	}

	r := openid.Response{
		ID:    params["openid.sig"],
		Teams: []string{"ce-web-logs"},
		SReg: map[string]string{
			"nickname": params["openid.sreg.nickname"],
			"fullname": params["openid.sreg.fullname"],
			"email":    params["openid.sreg.email"],
		},
	}

	return &r, nil
}

func verifyFail(requestURL string) (*openid.Response, error) {
	return nil, errors.New("MOCK error from OpenID verification")
}
