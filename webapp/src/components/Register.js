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
import {T} from './Utils';


class Register extends Component {

    constructor(props) {
        super(props)
        this.state = {
            message: null,
        };
    }

    getAge(m) {
        var start = moment(m);
        var end = moment()
        var dur = moment.duration(end.diff(start));
        var d = dur.asMinutes()
        if (d < 2) {
            return <i className="fa fa-clock led-green" title={start.format('llll')} />
        } else if (d < 5) {
            return <i className="fa fa-clock led-orange" title={start.format('llll')} />
        } else {
            return <i className="fa fa-clock led-red" title={start.format('llll')} />
        }
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
                        <th>{T('brand')}</th><th>{T('model')}</th><th>{T('serial')}</th>
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
            <tr key={l.registrationId}>
                <td className="overflow"><a href={'/devices/' + l.deviceId+ '/info'}>{l.brand}</a></td>
                <td className="overflow"><a href={'/devices/' + l.deviceId+ '/info'}>{l.model}</a></td>
                <td className="overflow"><a href={'/devices/' + l.deviceId+ '/info'}>{l.serial}</a></td>
                <td>
                    {this.getAge(l.lastRefresh)}
                </td>
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
