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

import axios from 'axios'
import constants from './constants'

var service = {

    version: (query, cancelCallback) => {
        return axios.get(constants.baseUrl + 'version');
    },

    getToken: () => {
        return axios.get(constants.baseUrl + 'token')
    },

    getAuthToken: () => {
        return axios.get(constants.baseUrl + 'authtoken')
    },

    accountsList: (query, cancelCallback) => {
        return axios.get(constants.baseUrl + 'organizations');
    },

    accountsNew: (query, cancelCallback) => {
        return axios.post(constants.baseUrl + 'organizations', query);
    },

    accountsGet: (id, cancelCallback) => {
        return axios.get(constants.baseUrl + 'organizations/' + id);
    },

    accountsUpdate: (id, query, cancelCallback) => {
        return axios.put(constants.baseUrl + 'organizations/' + id, query);
    },

    accountsForUsers: (username, cancelCallback) => {
        return axios.get(constants.baseUrl + 'users/' + username + '/organizations');
    },

    accountsUpdateForUser: (userId, accountId, cancelCallback) => {
        return axios.post(constants.baseUrl + 'users/' + userId + '/organizations/' + accountId);
    },

    groupsList: (account, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/groups');
    },

    snapsList: (account, query, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/' + query + '/snaps');
    },

    snapsRemove: (account,device, snap, cancelCallback) => {
        return axios.delete(constants.baseUrl + 'snaps/' + device.device.accountCode + '/' + device.device.name + '/' + device.device.serial + '/' + snap);
    },

    snapsInstall: (account, device, snap, cancelCallback) => {
        return axios.post(constants.baseUrl + 'snaps/' + device.device.accountCode + '/' + device.device.name + '/' + device.device.serial + '/' + snap);
    },

    snapsUpdate: (account, device, snap, query, cancelCallback) => {
        return axios.put(constants.baseUrl + 'snaps/' + device.device.accountCode + '/' + device.device.name + '/' + device.device.serial + '/' + snap, query);
    },

    snapsSettings: (account, device, snap, cancelCallback) => {
        return axios.get(constants.baseUrl + 'snaps/' + device.device.accountCode + '/' + device.device.name + '/' + device.device.serial + '/' + snap + '/settings');
    },

    snapsSettingsUpdate: (account, device, snap, settings, cancelCallback) => {
        return axios.put(constants.baseUrl + 'snaps/' + device.device.accountCode + '/'  + device.device.name + '/' + device.device.serial + '/' + snap + '/settings', settings);
    },

    storeSearch: (snapName,cancelCallback) => {
        return axios.get(constants.baseUrl + 'store/snaps/' + snapName);
    },

    clientsList: (account, query, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/clients');
    },

    clientsDetail: (account, query, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/clients/' + query);
    },

    clientsDeviceObject: (account, query, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/clients/' + query + '/device');
    },

    devicesList: (account, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/devices');
    },

    devicesGet: (account, id, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/devices/' + id);
    },

    devicesNew: (account, device, cancelCallback) => {
        return axios.post(constants.baseUrl + account + '/devices', device);
    },

    devicesUpdate: (account, device, cancelCallback) => {
        return axios.put(constants.baseUrl + device.accountCode + '/devices/' + device.id, device);
    },

    usersList: (query, cancelCallback) => {
        return axios.get(constants.baseUrl + 'users');
    },

    usersNew: (query, cancelCallback) => {
        return axios.post(constants.baseUrl + 'users', query);
    },

    usersGet: (id, cancelCallback) => {
        return axios.get(constants.baseUrl + 'users/' + id);
    },

    usersUpdate: (id, query, cancelCallback) => {
        return axios.put(constants.baseUrl + 'users/' + id, query);
    }
}

export default service
