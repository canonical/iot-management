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
import api from '../models/api';
import AlertBox from './AlertBox';
import DialogBox from './DialogBox';
import {T, formatError, isUserAdmin, roleAsString} from './Utils';

class Users extends Component {
    
    constructor(props) {
        super(props)
        this.state = {
            message: null,
            users: [],
            accounts: [],
            confirmDelete: null,
            showAccountsId: null,
        };

        this.getUsers();
    }

    getUsers () {
        api.usersList().then(response => {
            this.setState({users: response.data.users})
        })
    }

    getAccountsForUser (username, userId) {
        api.accountsForUsers(username).then(response => {
            this.setState({accounts: response.data.accounts, showAccountsId: userId})
        })
    }

    handleDelete = (e) => {
        e.preventDefault();
        this.setState({confirmDelete: parseInt(e.target.getAttribute('data-key'), 10)});
    }

    handleDeleteUser = (e) => {
        e.preventDefault();
        var users = this.state.users.filter((user) => {
          return user.id === this.state.confirmDelete;
        });
        if (users.length === 0) {
          return;
        }
    
        api.deleteUser(users[0]).then(response => {
            window.location = '/users';
        })
        .catch((e) => {
            this.setState({error: formatError(e.response.data)});
        })
    }

    handleDeleteUserCancel = (e) => {
        e.preventDefault();
        this.setState({confirmDelete: null});
    }

    handleClickAccount = (e) => {
        e.preventDefault();
        var accountId = e.target.getAttribute('data-account');
        var userId = parseInt(e.target.getAttribute('data-user'), 10);
        var username = e.target.getAttribute('data-username');

        api.accountsUpdateForUser(userId, accountId).then(response => {
            this.getAccountsForUser(username, userId)
        })
    }

    showAccounts = (e) => {
        e.preventDefault();

        var id = parseInt(e.target.getAttribute('data-key'), 10);
        var username = e.target.getAttribute('data-user');
        if (this.state.showAccountsId === id) {
            this.setState({showAccountsId: null});
        } else {
            this.getAccountsForUser(username, id)
        }
    }

    renderTable(items) {

        if (items.length > 0) {
            return (
            <div>
                <table>
                <thead>
                    <tr>
                        <th /><th>{T('username')}</th><th>{T('name')}</th><th>{T('email')}</th><th>{T('role')}</th><th>{T('accounts')}</th>
                    </tr>
                </thead>
                    {this.renderRows(items)}
                </table>
            </div>
            );
        } else {
            return (
            <p>{T('no-users')}</p>
            );
        }
    }

    renderActions(user) {
        if (user.id !== this.state.confirmDelete) {
            return (
                <div>
                    <a href={'/users/'.concat(user.id)} className="p-button--brand small" title={T('edit-user')}><i className="fa fa-edit" /></a>
                    &nbsp;
                    <button onClick={this.handleDelete} data-key={user.id} className="p-button--neutral small" title={T('delete-user')}>
                        <i className="fa fa-trash" data-key={user.id} /></button>
                </div>
            );
        } else {
            return (
                <DialogBox message={T('confirm-user-delete')} handleYesClick={this.handleDeleteUser} handleCancelClick={this.handleDeleteUserCancel} />
            );
        }
    }

    renderRows(items) {
        return items.map((l) => {
          var c = "";
          if (l.id === this.state.showAccountsId) {
            c = 'borderless';
          }
          return (
              <tbody>
                <tr key={l.id} className={c}>
                    <td>
                    {this.renderActions(l)}
                    </td>
                    <td className="overflow" title={l.username}>{l.username}</td>
                    <td className="overflow" title={l.name}>{l.name}</td>
                    <td className="overflow" title={l.email}>{l.email}</td>
                    <td className="overflow" title={roleAsString(l.role)}>{roleAsString(l.role)}</td>
                    <td className="overflow">
                        <button className="p-button--neutral small" title={T('view-accounts')} data-key={l.id} data-user={l.username} onClick={this.showAccounts}>
                            <i className="fa fa-building" aria-hidden="true" data-key={l.id} data-user={l.username} />
                        </button>
                    </td>
                </tr>
                {this.renderAccounts(l)}
              </tbody>
          );
        });
    }
    
    renderAccounts(item) {
        if (item.id !== this.state.showAccountsId) {
            return
        }

        var style;

        var accounts = this.state.accounts.map(acc => {
            if (acc.selected) {
                style = 'p-button--brand';
            } else {
                style = 'p-button--neutral';
            }
            return (
                <button key={acc.code} data-account={acc.id} data-user={item.id} data-username={item.username} onClick={this.handleClickAccount} className={style}>
                {acc.name}
                </button>
            )
        })

        return (
            <tr>
                <td colSpan="6">Select accounts for {item.name}:<br /> {accounts}</td>
            </tr>
        )
    }

    render () {

        if (!isUserAdmin(this.props.token)) {
            return (
                <div className="row">
                <AlertBox message={T('error-no-permissions')} />
                </div>
            )
        }

        return (
            <div className="row">
                <section className="row">
                    <div>
                        <h2>{T('users')}
                            <div className="u-float-right">
                                <a href="/users/new" className="p-button--brand" title={T('new-user')}>
                                    <i className="fa fa-plus" aria-hidden="true" />
                                </a>
                            </div>
                        </h2>
                    </div>
                    <div className="col-12">
                        <AlertBox message={this.state.message} />
                    </div>
                </section>

                <section className="row">
                    {this.renderTable(this.state.users)}
                </section>
            </div>
        )
    }

}

export default Users;
