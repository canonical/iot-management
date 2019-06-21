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
import {T, formatError, isUserAdmin} from './Utils';

class GroupEdit extends Component {
    constructor(props) {
        super(props);
        this.state = {
            title: null,
            error: null,
            hideForm: false,
            nameTo: '',
        };
    }

    componentDidMount() {

        if (this.props.name) {
            this.setTitle('edit-group');
            this.getGroup(this.props.account.orgid, this.props.name);
        } else {
            this.setTitle('new-group');
        }
    }

    getGroup(org_id, name) {
        api.groupsGet(org_id, name).then(response => {
            this.setState({group: response.data.group, nameTo: response.data.group.name});
        })
        .catch((e) => {
            this.setState({error: formatError(e.response.data), hideForm: true});
        })
    }

    setTitle(title) {
        this.setState({title: T(title)});
    }

    handleChangeName = (e) => {
        this.setState({nameTo: e.target.value});
    }

    handleSaveClick = (e) => {
        e.preventDefault();

        if (this.props.name) {
            // Update the existing device
            api.groupsUpdate(this.props.account.orgid, this.props.name, this.state.nameTo).then(response => {
                window.location = '/groups';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
        } else {
            // Create a new device
            api.groupsCreate(this.props.account.orgid, this.state.nameTo).then(response => {
                window.location = '/groups';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
        }
    }

    render () {
        if (!isUserAdmin(this.props.token)) {
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

        let disabled = this.state.status ? true : false;

        return (
            <div className="row">
                <section className="row">
                    <h2>{this.state.title}</h2>

                    <AlertBox message={this.state.error} />

                    <form>
                        <fieldset>
                            <label htmlFor="name">{T('name')}:
                                <input type="text" id="name" placeholder={T('group-desc')}
                                       value={this.state.name} onChange={this.handleChangeName} disabled={disabled} />
                            </label>
                        </fieldset>
                    </form>

                    <div>
                        <a href='/groups' className="p-button--neutral">{T('cancel')}</a>
                        &nbsp;
                        <a href='/groups' onClick={this.handleSaveClick} className="p-button--brand">{T('save')}</a>
                    </div>
                </section>
                <br />
            </div>
        )

    }

}

export default GroupEdit;
