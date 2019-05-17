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
	"log"
	"testing"
	"time"
)

func TestNonceStore_Accept(t *testing.T) {
	t1 := time.Now().UTC().Format(time.RFC3339)
	t2 := time.Now().AddDate(-2, 0, 0).UTC().Format(time.RFC3339)
	type args struct {
		endpoint string
		nonce    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"/login", t1}, false},
		{"invalid-short", args{"/login", "1234"}, true},
		{"invalid-value", args{"/login", "12345678901234567890"}, true},
		{"invalid-expired", args{"/login", t2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println("---", tt.args.nonce)
			mem := NewStore()
			s := mem.OpenIDNonceStore()
			if err := s.Accept(tt.args.endpoint, tt.args.nonce); (err != nil) != tt.wantErr {
				t.Errorf("NonceStore.Accept() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
