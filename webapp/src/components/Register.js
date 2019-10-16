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
import {T, canUserAdministrate} from './Utils';
import {Status} from './Constants';


class Register extends Component {

    constructor(props) {
        super(props)
        this.state = {
            message: null,
        };
    }

    getStatus(s) {
        let status = T(Status[s])
        if (s === 1) {
            return <span><i className="fa fa-clock led-orange" title={status} />&nbsp;{status}</span>
        } else if (s === 2) {
            return <span><i className="fa fa-clipboard-check led-green" title={status} />&nbsp;{status}</span>
        } else {
            return <span><i className="fa fa-times-circle led-red" title={status} />&nbsp;{status}</span>
        }
    }

    copyToClipboard = (e) => {
        e.preventDefault()
        const el = document.createElement('textarea');
        el.value = e.target.getAttribute('data-key');
        document.body.appendChild(el);
        el.select();
        document.execCommand('copy');
        document.body.removeChild(el);
    }

    renderTable(items) {
        
        if (!items) {
            return
        }
        if (items.length > 0) {
            return (
            <div>
                <table>
                <thead>
                    <tr>
                        <th className="small" /><th>{T('id')}</th><th>{T('brand')}</th><th>{T('model')}</th><th>{T('serial')}</th><th className="overflow">{T('device-key')}</th>
                        <th className="xsmall">{T('status')}</th>
                    </tr>
                </thead>
                <tbody>
                    {this.renderRows(items)}
                </tbody>
                </table>
            </div>
            );
        } else {
            return (
            <p>{T('no-devices')}</p>
            );
        }
    }

    renderRows(items) {
        return items.map((l) => {
          return (
            <tr key={l.id}>
                <td>
                    <a href={'/register/' + l.id} className="p-button--neutral small"><i className="fa fa-edit" /></a>
                    {l.device.deviceKey ?
                        <a href="" onClick={this.copyToClipboard} data-key={l.device.deviceKey} className="p-button--neutral small" title={T('copy-device-key')}>
                        <i className="fa fa-clipboard" data-key={l.device.deviceKey} /></a> : ''}
                </td>
                <td className="overflow">{l.id}</td>
                <td className="overflow">{l.device.brand}</td>
                <td className="overflow">{l.device.model}</td>
                <td className="overflow">{l.device.serial}</td>
                <td className="overflow" title={l.device.deviceKey}>
                    {(l.device.deviceKey && l.device.deviceKey.substr(0,40)) || ''}
                </td>
                <td>{this.getStatus(l.status)}</td>
            </tr>
          );
        });
    }

    render () {
        return (
            <div className="row">
                <section className="row">
                    <h2>
                        {T('register-devices')}
                        {canUserAdministrate(this.props.token) ?
                            <div className="u-float-right">
                                <a href="/register/new" className="p-button--brand" title={T('new-device')}>
                                    <i className="fa fa-plus" aria-hidden="true" />
                                </a>
                            </div>
                            : ''
                        }
                    </h2>
                    <div className="col-12">
                        <AlertBox message={this.state.message} />
                    </div>
                </section>

                <section className="row spacer">
                    {this.renderTable(this.props.devices)}
                </section>
            </div>
        )
    }
}

export default Register;
