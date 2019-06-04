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
import {T, isUserAdmin, isUserSuperuser} from './Utils';

class Accounts extends Component {
    
    constructor(props) {
        super(props);
        this.state = {
            message: null,
            accounts: [],
        };
    }

    renderTable(items) {
        
        if (items.length > 0) {
            return (
            <div>
                <table>
                <thead>
                    <tr>
                    <th>{T('code')}</th><th>{T('name')}</th>
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
            <p>{T('no-accounts')}</p>
            );
        }
    }

    renderRows(items) {
        return items.map((l) => {
          let isSuperuser = isUserSuperuser(this.props.token);
          return (
            <tr key={l.id}>
                <td>{isSuperuser ? <a href={'/accounts/' + l.orgid}>{l.orgid}</a> : l.orgid}</td>
                <td>{l.name}</td>
            </tr>
          );
        });
    }

    render () {

        if (!isUserAdmin(this.props.token)) {
            return (
                <div className="row">
                <AlertBox message={T('error-no-permissions')} />
                </div>
            )
        }
        let isSuperuser = isUserSuperuser(this.props.token);

        return (
            <div className="row">
                <section className="row">
                    <div>
                        <h2>{T('accounts')}
                            {isSuperuser ?
                                <div className="u-float-right">
                                    <a href="/accounts/new" className="p-button--brand" title={T('new-account')}>
                                        <i className="fa fa-plus" aria-hidden="true" />
                                    </a>
                                </div>
                                : ''
                            }
                        </h2>
                    </div>
                    <div className="col-12">
                        <AlertBox message={this.state.message} />
                    </div>
                </section>

                <section className="row">
                    {this.renderTable(this.props.accounts)}
                </section>
            </div>
        )
    }

}

export default Accounts;
