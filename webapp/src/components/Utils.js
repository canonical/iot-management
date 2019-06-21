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

import Messages from './Messages'
import {Role} from './Constants'
import jwtDecode from 'jwt-decode'
import api from '../models/api';



const subLinks = {
    register: ['register', 'devices'],
    devices: ['info', 'snaps'],
    groups: ['groups', 'actions'],
    actions: ['groups', 'actions'],
}

export function T(message) {
    const msg = Messages[message] || message;
    return msg
}

// URL is in the form:
//  /section
//  /section/sectionId
//  /section/sectionId/subsection
export function parseRoute() {
    const parts = window.location.pathname.split('/')

    switch (parts.length) {
        case 2:
            return {section: parts[1]}
        case 3:
            return {section: parts[1], sectionId: parts[2]}
        case 4:
            return {section: parts[1], sectionId: parts[2], subsection: parts[3]}
        default:
            return {}
    }
}

export function sectionNavLinks(section, sectionId) {
    if (section === '') {
        return;
    }
    if ((section === 'devices') && (!sectionId)) {
        return subLinks['register']
    }
    return subLinks[section];
}

export function isLoggedIn(token) {
    return isUserStandard(token)
}

export function isUserStandard(token) {
    return isUser(Role.Standard, token)
}

export function isUserAdmin(token) {
    return isUser(Role.Admin, token)
}

export function isUserSuperuser(token) {
    return isUser(Role.Superuser, token)
}

export function roleAsString(role) {
    var str
    switch (role) {
        case Role.Standard:
            str = "Standard"	
            break;
        case Role.Admin:
            str = "Admin"
            break;
        case Role.Superuser:
            str = "Superuser"
            break
        default:
            str= "invalid role"
            break;
    }
    return str
}

function isUser(role, token) {
    if (!token) return false
    if (!token.role) return false

    return (token.role >= role)
}

export function getAuthToken(callback) {
    // Get a fresh token and return it to the callback
    // The token will be passed to the views
    api.getAuthToken().then((resp) => {

        var jwt = resp.headers.authorization

        if (!jwt) {
            callback({})
            return
        }
        var token = jwtDecode(jwt)

        if (!token) {
            callback({})
            return
        }
        callback(token)
    })
}

export function formatError(data) {
    var message = T(data.code);
    if (data.message) {
        message += ': ' + data.message;
    }
    return message;
}

export function saveAccount(account) {
    sessionStorage.setItem('accountCode', account.orgid);
    sessionStorage.setItem('accountName', account.name);
}

export function getAccount() {
    return {
        orgid: sessionStorage.getItem('accountCode'),
        name: sessionStorage.getItem('accountName'),
    }
}
