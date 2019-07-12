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

package identityapi

import (
	"encoding/json"
	"github.com/CanonicalLtd/iot-identity/service"
	"testing"
)

func TestClientAdapter_RegisterOrganization(t *testing.T) {
	b1 := `{"id": "def", "message": ""}`
	type fields struct {
		URL string
	}
	type args struct {
		name    string
		country string
		body    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr string
	}{
		{"valid", fields{""}, args{"Test Inc", "GB", b1}, "def", ""},
		{"invalid-org", fields{"invalid"}, args{"invalid", "GB", b1}, "", "MOCK error post"},
		{"invalid-body", fields{""}, args{"abc", "GB", ""}, "", "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			b, _ := json.Marshal(service.RegisterOrganizationRequest{Name: tt.args.name, CountryName: tt.args.country})

			got := a.RegisterOrganization(b)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.RegisterOrganization() = %v, want %v", got.Message, tt.wantErr)
			}
			if got.ID != tt.want {
				t.Errorf("ClientAdapter.RegisterOrganization() = %v, want ID %v", got.Code, tt.want)
			}
		})
	}
}

func TestClientAdapter_RegOrganizationList(t *testing.T) {
	b1 := `{"organizations": [{"id":"abc", "name":"Test Org Ltd"}]}`
	type fields struct {
		URL string
	}
	tests := []struct {
		name    string
		body    string
		fields  fields
		want    int
		wantErr string
	}{
		{"valid", b1, fields{""}, 1, ""},
		{"invalid-org", "", fields{"invalid"}, 0, "MOCK error get"},
		{"invalid-body", "", fields{""}, 0, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.RegOrganizationList()
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.RegisterOrganization() = %v, want %v", got.Message, tt.wantErr)
			}
			if len(got.Organizations) != tt.want {
				t.Errorf("ClientAdapter.RegisterOrganization() = %v, want ID %v", len(got.Organizations), tt.want)
			}
		})
	}
}
