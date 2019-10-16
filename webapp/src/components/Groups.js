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


import React, {Component} from 'react';
import AlertBox from './AlertBox';
import {canUserAdministrate, formatError, T} from './Utils';
import api from "../models/api";

class Groups extends Component {
    constructor(props) {
        super(props)
        this.state = {
            name: null,
            devices: [],
            devicesExcluded: [],
        }
    }

    refresh(orgid, name) {
        this.getGroupDevices(orgid, name)
        this.getGroupExcludedDevices(orgid, name)
    }

    getGroupDevices(orgid, name) {
        api.groupsGetDevices(orgid, name).then(response => {
            this.setState({devices: response.data.devices})
        }).catch(e => {
            this.setState({message: formatError(e.response.data), devices: []});
        })
    }

    getGroupExcludedDevices (orgid, name) {
        api.groupsGetDevicesExcluded(orgid, name).then(response => {
            this.setState({devicesExcluded: response.data.devices})
        }).catch(e => {
            this.setState({message: formatError(e.response.data), devicesExcluded: []});
        })
    }

    handleGroupClick = (e) => {
        e.preventDefault();

        let selected = e.target.getAttribute('data-key')
        if (this.state.name === selected) {
            // Deselect the group
            this.setState({name: null, devices: [], devicesExcluded: []})
        } else {
            // Select the group
            this.setState({name: selected})
            this.refresh(this.props.account.orgid, selected)
        }
    }

    handleAddToGroup = (e) => {
        e.preventDefault();
        let deviceId = e.target.getAttribute('data-key')

        api.groupsDeviceLink(this.props.account.orgid, this.state.name, deviceId).then(response => {
            this.refresh(this.props.account.orgid, this.state.name)
        }).catch(e => {
            this.setState({message: formatError(e.response.data), devicesExcluded: []});
        })
    }

    handleRemoveFromGroup = (e) => {
        e.preventDefault();
        let deviceId = e.target.getAttribute('data-key')

        api.groupsDeviceUnlink(this.props.account.orgid, this.state.name, deviceId).then(response => {
            this.refresh(this.props.account.orgid, this.state.name)
        }).catch(e => {
            this.setState({message: formatError(e.response.data), devicesExcluded: []});
        })
    }

    renderRowsGroups(items) {
        return items.map((l) => {
          let selected = (l.name===this.state.name) ? 'p-button--brand' : 'p-button--neutral'
          return (
            <tr key={l.name}>
                <td className="overflow">
                    <button className={selected} onClick={this.handleGroupClick} data-key={l.name}>{l.name}</button>
                </td>
            </tr>
          );
        });
    }

    renderRowsDevices(items, excluded) {
        return items.map((l) => {
            return (
                <tr key={l.deviceId}>
                    <td className="overflow">
                        {excluded ? <button onClick={this.handleAddToGroup} data-key={l.deviceId} className="p-button--neutral xsmall"><i data-key={l.deviceId}  className="p-icon--plus" /></button>
                            : <button onClick={this.handleRemoveFromGroup} data-key={l.deviceId}  className="p-button--neutral xsmall"><i data-key={l.deviceId}  className="p-icon--close" /></button>}
                            &nbsp;
                        {l.brand} {l.model} {l.serial}
                    </td>
                </tr>
            );
        });
    }

    renderTableGroups(items) {
        if (!items) {
            return
        }
        if (items.length > 0) {
            return (
            <div className="col-4">
                <table>
                <thead>
                    <tr>
                        <th>{T('groups')}</th>
                    </tr>
                </thead>
                <tbody>
                    {this.renderRowsGroups(items)}
                </tbody>
                </table>
            </div>
            );
        } else {
            return (
            <p>{T('no-groups')}</p>
            );
        }
    }

    renderTableDevices(items) {
        if (!items) {
            return
        }
        if (items.length > 0) {
            return (
                <div className="col-4">
                    <table>
                        <thead>
                        <tr>
                            <th>{T('devices')}</th>
                        </tr>
                        </thead>
                        <tbody>
                        {this.renderRowsDevices(items, false)}
                        </tbody>
                    </table>
                </div>
            );
        } else {
            return (
                <div className="col-4">
                    <table>
                        <thead>
                        <tr>
                            <th>{T('devices')}</th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr><td>{T('no-devices')}</td></tr>
                        </tbody>
                    </table>
                </div>
            );
        }
    }

    renderTableDevicesExcluded(items) {
        if (!items) {
            return
        }
        if (items.length > 0) {
            return (
                <div className="col-3">
                    <table>
                        <thead>
                        <tr>
                            <th>{T('devices-excluded')}</th>
                        </tr>
                        </thead>
                        <tbody>
                        {this.renderRowsDevices(items, true)}
                        </tbody>
                    </table>
                </div>
            );
        } else {
            return (
                <div className="col-4">
                    <table>
                        <thead>
                        <tr>
                            <th>{T('devices-excluded')}</th>
                        </tr>
                        </thead>
                        <tbody>
                        <tr><td>{T('no-devices')}</td></tr>
                        </tbody>
                    </table>
                </div>
            );
        }
    }

    render () {
        return (
            <div className="row">
                <section className="row">
                    <h2>
                        {T('groups')}
                        {canUserAdministrate(this.props.token) ?
                            <div className="u-float-right">
                                <a href="/groups/new" className="p-button--brand" title={T('new-group')}>
                                    <i className="fa fa-plus" aria-hidden="true" />
                                </a>
                            </div>
                            : ''
                        }
                    </h2>
                    <div className="col-12">
                        <AlertBox message={this.props.message} />
                    </div>
                </section>

                <section className="row spacer">
                    {this.renderTableGroups(this.props.groups)}
                    {this.state.name ?
                        this.renderTableDevices(this.state.devices) : ''
                    }
                    {this.state.name ?
                        this.renderTableDevicesExcluded(this.state.devicesExcluded) : ''
                    }
                </section>
            </div>
        )
    }
}

export default Groups
