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
import {T} from './Utils';

class Groups extends Component {

    renderRows(items) {
        return items.map((l) => {
          console.log(l)
          return (
            <tr key={l.id}>
                <td className="overflow">
                    <button>{l.name}</button>
                </td>
            </tr>
          );
        });
    }


    renderTable(items) {
        
        if (!items) {
            return
        }
        if (items.length > 0) {
            return (
            <div className="col-3">
                <table>
                <thead>
                    <tr>
                        <th>{T('name')}</th>
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
            <p>{T('no-groups')}</p>
            );
        }
    }

    render () {
        return (
            <div className="row">
                <section className="row">
                    <h2>{T('groups')}</h2>
                    <div className="col-12">
                        <AlertBox message={this.props.message} />
                    </div>
                </section>

                <section className="row spacer">
                    {this.renderTable(this.props.groups)}
                </section>
            </div>
        )
    }
}

export default Groups
