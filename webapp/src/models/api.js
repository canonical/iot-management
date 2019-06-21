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

    groupsGet: (account, name, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/groups/' + name);
    },

    groupsGetDevices: (account, name, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/groups/' + name + '/devices');
    },

    groupsGetDevicesExcluded: (account, name, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/groups/' + name + '/devices/excluded');
    },

    groupsCreate: (account, name, cancelCallback) => {
        return axios.post(constants.baseUrl + account + '/groups', {orgid: account, name: name});
    },

    groupsUpdate: (account, nameFrom, nameTo, cancelCallback) => {
        return axios.put(constants.baseUrl + account + '/groups', {orgid: account, nameFrom: nameFrom, nameTo: nameTo});
    },

    groupsDeviceLink: (account, name, deviceId, cancelCallback) => {
        return axios.post(constants.baseUrl + account + '/groups/' + name + '/' + deviceId);
    },

    groupsDeviceUnlink: (account, name, deviceId, cancelCallback) => {
        return axios.delete(constants.baseUrl + account + '/groups/' + name + '/' + deviceId);
    },

    snapsList: (account, device, cancelCallback) => {
        return axios.get(constants.baseUrl + 'device/' + account + '/' + device + '/snaps');
    },
    
    snapsInstallRefresh: (account, device, cancelCallback) => {
        return axios.post(constants.baseUrl + 'snaps/' + account + '/' + device.device.deviceId + '/list');
    },

    snapsRemove: (account,device, snap, cancelCallback) => {
        return axios.delete(constants.baseUrl + 'snaps/' + account + '/' + device.device.deviceId + '/' + snap);
    },

    snapsInstall: (account, device, snap, cancelCallback) => {
        return axios.post(constants.baseUrl + 'snaps/' + account + '/' + device.device.deviceId + '/' + snap);
    },

    snapsUpdate: (account, device, snap, action, cancelCallback) => {
        return axios.put(constants.baseUrl + 'snaps/' + account + '/' + device.device.deviceId + '/' + snap + '/' + action);
    },

    snapsSettings: (account, device, snap, cancelCallback) => {
        return axios.get(constants.baseUrl + 'snaps/' + account + '/' + device.device.deviceId + '/' + snap + '/settings');
    },

    snapsSettingsUpdate: (account, device, snap, settings, cancelCallback) => {
        return axios.put(constants.baseUrl + 'snaps/' + account + '/' + device.device.deviceId + '/' + snap + '/settings', settings);
    },

    storeSearch: (snapName,cancelCallback) => {
        return axios.get(constants.baseUrl + 'store/snaps/' + snapName);
    },

    clientsList: (account, query, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/register/devices');
    },

    clientsGet: (account, device, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/register/devices/' + device);
    },

    clientsNew: (account, device, cancelCallback) => {
        return axios.post(constants.baseUrl + account + '/register/devices', device);
    },

    clientsDeviceObject: (account, query, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/clients/' + query + '/device');
    },

    clientsUpdate: (account, deviceId, status, cancelCallback) => {
        return axios.put(constants.baseUrl + account + '/register/devices/' + deviceId, {status: status});
    },

    devicesUpdate: (account, device, cancelCallback) => {
        return axios.put(constants.baseUrl + device.accountCode + '/devices/' + device.id, device);
    },

    devicesList: (account, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/devices');
    },

    devicesGet: (account, id, cancelCallback) => {
        return axios.get(constants.baseUrl + account + '/devices/' + id);
    },

    usersList: (query, cancelCallback) => {
        return axios.get(constants.baseUrl + 'users');
    },

    usersNew: (query, cancelCallback) => {
        return axios.post(constants.baseUrl + 'users', query);
    },

    usersGet: (username, cancelCallback) => {
        return axios.get(constants.baseUrl + 'users/' + username);
    },

    usersUpdate: (id, query, cancelCallback) => {
        return axios.put(constants.baseUrl + 'users/' + id, query);
    },

    usersDelete: (username, cancelCallback) => {
        return axios.delete(constants.baseUrl + 'users/' + username);
    }
}

export default service
