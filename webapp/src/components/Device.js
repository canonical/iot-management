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
        if (!d.device) {return <div>Loading...</div>};

        return (
            <div className="row">
                <h1 className="tight">{d.device.brand} {d.device.model}</h1>
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
                                        <td>{T('model')}:</td><td>{d.device.brand} {d.device.model}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('serial-number')}:</td><td>{d.device.serial}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('os-version')}:</td><td>{d.device.version.osId} {d.device.version.osVersionId}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('series')}:</td><td>{d.device.version.series}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('version')}:</td><td>{d.device.version.version}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('kernel-version')}:</td><td>{d.device.version.kernelVersion}</td>
                                    </tr>
                                    <tr>
                                        <td>{T('on-classic')}:</td><td>{d.device.version.onClassic}</td>
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
