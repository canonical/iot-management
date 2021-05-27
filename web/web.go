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
	"net/http"

	"github.com/everactive/iot-management/config"
	"github.com/everactive/iot-management/service/manage"
	muxlogrus "github.com/pytimer/mux-logrus"
	log "github.com/sirupsen/logrus"
)

// JSONHeader is the content-type header for JSON responses
const JSONHeader = "application/json; charset=UTF-8"

// Service is the implementation of the web API
type Service struct {
	Settings *config.Settings
	Manage   manage.Manage
}

// NewService returns a new web controller
func NewService(settings *config.Settings, srv manage.Manage) *Service {
	return &Service{
		Settings: settings,
		Manage:   srv,
	}
}

// Run starts the web service
func (wb Service) Run() error {
	fmt.Printf("Starting service on port :%s\n", wb.Settings.LocalPort)

	r := wb.Router()
	r.Use(muxlogrus.NewLogger(muxlogrus.LogOptions{Formatter: &log.JSONFormatter{}}).Middleware)

	return http.ListenAndServe(":"+wb.Settings.LocalPort, r)
}
