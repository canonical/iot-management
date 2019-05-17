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

import "testing"

func TestCreateSecret(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{"valid-10", args{10}, 10, false},
		{"valid-20", args{20}, 20, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateSecret(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSecret() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) < tt.wantLen {
				t.Errorf("CreateSecret() length = %v, want min. %v", len(got), tt.wantLen)
			}
		})
	}
}
