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
import {T, isUserSuperuser, formatError} from './Utils';

class AccountEdit extends Component {
    
    constructor(props) {
        super(props);
        this.state = {
            title: null,
            error: null,
            hideForm: false,
            account: {active: true},
        };
    }

    componentDidMount() {

        if (this.props.id) {
            this.setTitle('edit-account');
            this.getAccount(this.props.id);
        } else {
            this.setTitle('new-account');
        }
    }

    getAccount(id) {
        api.accountsGet(id).then(response => {
            this.setState({account: response.data.organization});
        })
        .catch((e) => {
            this.setState({error: formatError(e.response.data), hideForm: true});
        })
    }

    setTitle(title) {
        this.setState({title: T(title)});
    }

    handleChangeCode = (e) => {
        var account = this.state.account;
        account['orgid'] = e.target.value;
        this.setState({account: account});
    }

    handleChangeName = (e) => {
        var account = this.state.account;
        account['name'] = e.target.value;
        this.setState({account: account});
    }

    handleChangeActive = (e) => {
        var account = this.state.account;
        account['active'] = e.target.checked;
        this.setState({account: account});
    }

    handleSaveClick = (e) => {
        e.preventDefault();

        if (this.props.id) {
            api.accountsUpdate(this.props.id, this.state.account).then(response => {
                window.location = '/accounts';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
        } else {
            api.accountsNew(this.state.account).then(response => {
                window.location = '/accounts';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
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

        if (!this.state.account.orgid) {
            return <div />
        }

        return (
            <div className="row">
                <section className="row">
                    <h2>{this.state.title}</h2>

                    <AlertBox message={this.state.error} />

                    <form>
                        <fieldset>
                            <label htmlFor="code">{T('code')}:
                                <input type="text" id="code" placeholder={T('account-code-desc')}
                                    value={this.state.account.orgid} onChange={this.handleChangeCode} />
                            </label>
                            <label htmlFor="name">{T('name')}:
                                <input type="text" id="name" placeholder={T('account-name-desc')}
                                    value={this.state.account.name} onChange={this.handleChangeName}/>
                            </label>
                        </fieldset>
                    </form>

                    <div>
                        <a href='/accounts' className="p-button--neutral">{T('cancel')}</a>
                        &nbsp;
                        <a href='/accounts' onClick={this.handleSaveClick} className="p-button--brand">{T('save')}</a>
                    </div>
                </section>
                <br />
            </div>
        )

    }
}

export default AccountEdit;
