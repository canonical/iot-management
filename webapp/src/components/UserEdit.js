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
import {T, formatError, isUserSuperuser} from './Utils';

class UserEdit extends Component {

    constructor(props) {
        super(props)
        this.state = {
            title: null,
            user: {},
            error: null,
            hideForm: false,
        }
    }

    componentDidMount() {
        if (this.props.username) {
            this.setTitle('edit-user');
            this.getUser(this.props.username);
        } else {
            this.setTitle('new-user');
        }
    }

    setTitle(title) {
        this.setState({title: T(title)});
    }

    getUser(userId) {
        api.usersGet(userId).then((response) => {
            this.setState({
                user: response.data.user, 
                hideForm: false,
            });
        })
        .catch((e) => {
            this.setState({error: formatError(e.response.data)});
        })
    }

    handleChangeUsername = (e) => {
        var user = this.state.user;
        user.username = e.target.value;
        this.setState({user: user});
    }

    handleChangeName = (e) => {
        var user = this.state.user;
        user.name = e.target.value;
        this.setState({user: user});
    }

    handleChangeEmail = (e) => {
        var user = this.state.user;
        user.email = e.target.value;
        this.setState({user: user});
    }

    handleChangeRole = (e) => {
        var user = this.state.user;
        user.role = parseInt(e.target.value, 10);
        this.setState({user: user});
    }

    handleSaveClick = (e) => {
        e.preventDefault();

        if (this.props.username) {
            // Update the existing user
            api.usersUpdate(this.props.username, this.state.user).then(response => {
                window.location = '/users';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
        } else {
            // Create a new user
            api.usersNew(this.state.user).then(response => {
                window.location = '/users';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
        }
    }

    renderError() {
        if (this.state.error) {
            return (
                <AlertBox message={this.state.error} />
            );
        }
    }

    render() {
        if (!isUserSuperuser(this.props.token)) {
            return (
                <div className="row">
                <AlertBox message={T('error-no-permissions')} />
                </div>
            )
        }

        if (this.state.hideForm) {
            return (
                <div className="row">
                <AlertBox message={this.state.error} />
                </div>
            )
        }

        return (
            <div className="row">
                <section className="row">
                      <h2>{this.state.title}</h2>

                        <AlertBox message={this.state.error} />

                        <form>
                            <fieldset>
                                <label htmlFor="username">{T('username')}:
                                    <input type="text" id="username" placeholder={T('user-username-desc')}
                                        value={this.state.user.username} onChange={this.handleChangeUsername} />
                                </label>
                                <label htmlFor="name">{T('name')}:
                                    <input type="text" id="name" placeholder={T('user-name-desc')}
                                        value={this.state.user.name} onChange={this.handleChangeName}/>
                                </label>
                                <label htmlFor="email">{T('email')}:
                                    <input type="text" id="email" placeholder={T('user-email-desc')}
                                        value={this.state.user.email} onChange={this.handleChangeEmail}/>
                                </label>
                                <label htmlFor="role">{T('role')}:
                                    <select value={this.state.user.role} id="role" onChange={this.handleChangeRole}>
                                        <option />
                                        <option key="standard" value="100">Standard</option>
                                        <option key="admin" value="200">Admin</option>
                                        <option key="superuser" value="300">Superuser</option>
                                    </select>
                                </label>
                            </fieldset>
                        </form>

                        <div>
                            <a href='/users' className="p-button--neutral">{T('cancel')}</a>
                            &nbsp;
                            <a href='/users' onClick={this.handleSaveClick} className="p-button--brand">{T('save')}</a>
                        </div>
                </section>
                <br />
            </div>
        )
    }
}

export default UserEdit;
