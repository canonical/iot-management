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
import moment from 'moment';
import AlertBox from './AlertBox';
import If from './If';
import {T} from './Utils';

class Device extends Component {

    render () {
        var d = this.props.client;
        if (!d.info) {return <div>Loading...</div>};

        return (
            <div className="row">
                <h1 className="tight">{d.device.accountCode} {d.device.name}</h1>
                <h4 className="subtitle">{d.device.serial}</h4>

                <section className="row spacer">
                    <div className="col-12">
                        <AlertBox message={this.props.message} />
                    </div>
                </section>

                <If cond={!this.props.message}>
                    <section className="row spacer">
                        <div className="p-card--highlighted col-6">
                            <table className="p-card__content">
                                <tbody>
                                    <tr>
                                        <td>{T('last-update')}:</td><td>{moment(d.device.lastRefresh).format('llll')}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('registered')}:</td><td>{moment(d.device.created).format('llll')}</td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </section>

                    <section className="row spacer">
                        <div className="p-card">
                            <h3 className="p-card__title">{T('system-info')}</h3>
                            <table className="p-card__content">
                                <tbody>
                                    <tr>
                                        <td>{T('model')}:</td><td>{d.device.accountCode} {d.device.name}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('serial-number')}:</td><td>{d.device.serial}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('software-version')}:</td><td>{d.info.softwareversion}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('firmware-version')}:</td><td>{d.info.firmwareversion}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('current-time')}:</td><td>{d.info.currenttime}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('utc-offset')}:</td><td>{d.info.utcoffset}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('timezone')}:</td><td>{d.info.timezone}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('uptime')}:</td><td>{d.info.uptime}</td>
                                    </tr>
                                </tbody>
                            </table>
                        </div>
                    </section>
                </If>
            </div>
        )
    }
}

export default Device
