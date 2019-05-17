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

package twinapi

import (
	"testing"
)

func TestClientAdapter_GroupList(t *testing.T) {
	b1 := `{"groups": [{"orgid":"abc", "name":"workshop"}]}`
	type fields struct {
		URL string
	}
	type args struct {
		orgID string
		body  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr string
	}{
		{"valid", fields{""}, args{"abc", b1}, 1, ""},
		{"invalid-org", fields{""}, args{"invalid", b1}, 0, "MOCK error get"},
		{"invalid-body", fields{""}, args{"abc", ""}, 0, "EOF"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockHTTP(tt.args.body)
			a := &ClientAdapter{
				URL: tt.fields.URL,
			}
			got := a.GroupList(tt.args.orgID)
			if got.Message != tt.wantErr {
				t.Errorf("ClientAdapter.GroupList() = %v, want %v", got.Message, tt.wantErr)
			}
			if len(got.Groups) != tt.want {
				t.Errorf("ClientAdapter.GroupList() = %v, want %v", len(got.Groups), tt.want)
			}
		})
	}
}
