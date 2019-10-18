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

class Actions extends Component {
    constructor(props) {
        super(props)
        this.state = {
            name: null,
            groups: [],
            actions: [],
        }
    }

    refresh(orgid, name) {
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
                    {this.renderTableGroups(this.props.groups)}
                </section>
            </div>
        )
    }
}

export default Actions
