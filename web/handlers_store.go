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
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// StoreSearchHandler fetches the available snaps from the store
func (wb Service) StoreSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSONHeader)

	vars := mux.Vars(r)

	resp, err := get(wb.Settings.StoreURL + "snaps/search?q=" + vars["snapName"])
	if err != nil {
		fmt.Fprint(w, "{}")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprint(w, "{}")
		return
	}

	fmt.Fprint(w, string(body))
}

var get = func(u string) (*http.Response, error) {
	return http.Get(u)
}
