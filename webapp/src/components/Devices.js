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


class Devices extends Component {

    renderTable(items) {
        
        if (items.length > 0) {
            return (
            <div>
                <table>
                <thead>
                    <tr>
                    <th>{T('name')}</th><th>{T('reg-date')}</th><th>{T('address')}</th><th>{T('last-update')}</th>
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
                <td><a href={'/devices/' + l.endpoint+ '/info'}>{l.endpoint}</a></td>
                <td>{moment(l.registrationDate).format('llll')}</td>
                <td>{l.address}</td>
                <td>{moment(l.lastUpdate).format('llll')}</td>
            </tr>
          );
        });
    }

    render () {
        return (
            <div className="row">
                <section className="row">
                    <h2>{T('devices')}</h2>
                    <div className="col-12">
                        <AlertBox message={this.props.message} />
                    </div>
                </section>

                <section className="row spacer">
                    {this.renderTable(this.props.clients)}
                </section>
            </div>
        )
    }

}

export default Devices;