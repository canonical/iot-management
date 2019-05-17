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
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/juju/usso/openid"
)

type testJWT struct {
	resp     openid.Response
	expected []string
	role     int
}

const jwtSecret = "jwt_secret"

func TestNewJWTToken(t *testing.T) {
	test1 := testJWT{
		resp:     openid.Response{ID: "id", Teams: []string{"teamone", "team2"}},
		expected: []string{"", "", ""},
		role:     100,
	}
	test2 := testJWT{
		resp:     openid.Response{ID: "id", Teams: []string{"teamone", "team2"}, SReg: map[string]string{"nickname": "jwt"}},
		expected: []string{"jwt", "", ""},
		role:     200,
	}
	test3 := testJWT{
		resp: openid.Response{
			ID:    "id",
			Teams: []string{"teamone", "team2"},
			SReg:  map[string]string{"nickname": "jwt", "email": "jwt@example.com", "fullname": "John W Thompson"},
		},
		expected: []string{"jwt", "jwt@example.com", "John W Thompson"},
		role:     300,
	}

	for _, r := range []testJWT{test1, test2, test3} {

		// adding arbitrary role value for second parameter
		jwtToken, err := NewJWTToken(jwtSecret, &r.resp, r.role)
		if err != nil {
			t.Errorf("Error creating JWT: %v", err)
		}

		expectedToken(t, jwtToken, &r.resp, r.expected[0], r.expected[1], r.expected[2], r.role)

	}
}

func expectedToken(t *testing.T, jwtToken string, resp *openid.Response, username, email, name string, role int) {
	token, err := VerifyJWT(jwtSecret, jwtToken)
	if err != nil {
		t.Errorf("Error validating JWT: %v", err)
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims[ClaimsIdentity] != resp.ID {
		t.Errorf("JWT ID does not match: %v", claims[ClaimsIdentity])
	}
	if claims[ClaimsUsername] != username {
		t.Errorf("JWT username does not match: %v", claims[ClaimsUsername])
	}
	if claims[ClaimsEmail] != email {
		t.Errorf("JWT email does not match: %v", claims[ClaimsEmail])
	}
	if claims[ClaimsName] != name {
		t.Errorf("JWT name does not match: %v", claims[ClaimsName])
	}
	if int(claims[ClaimsRole].(float64)) != role {
		t.Errorf("JWT role does not match: expected %v but got %v", role, claims[ClaimsRole])
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {

}

func TestAddJWTCookie(t *testing.T) {
	w := httptest.NewRecorder()
	AddJWTCookie("ThisShouldBeAJWT", w)

	// Copy the Cookie over to a new Request
	request := &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}

	// Extract the cookie from the request
	jwtToken, err := JWTExtractor(request)
	if err != nil {
		t.Errorf("Error getting the JWT cookie: %v", err)
	}
	if jwtToken != "ThisShouldBeAJWT" {
		t.Errorf("Expected 'ThisShouldBeAJWT', got '%v'", jwtToken)
	}
}
